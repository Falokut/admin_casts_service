package service

import (
	"context"
	"fmt"
	"strings"
	"sync"

	casts_service "github.com/Falokut/admin_casts_service/pkg/admin_casts_service/v1/protos"
	movies_persons_service "github.com/Falokut/admin_movies_persons_service/pkg/admin_movies_persons_service/v1/protos"
	movies_service "github.com/Falokut/admin_movies_service/pkg/admin_movies_service/v1/protos"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/sirupsen/logrus"
)

type personsChecker struct {
	client movies_persons_service.MoviesPersonsServiceV1Client
	conn   *grpc.ClientConn
}

func NewPersonsChecker(addr string) (*personsChecker, error) {
	conn, err := getGrpcConnection(addr)
	if err != nil {
		return nil, err
	}

	return &personsChecker{client: movies_persons_service.NewMoviesPersonsServiceV1Client(conn), conn: conn}, nil
}

func (c *personsChecker) Shutdown() error {
	return c.conn.Close()
}

func (c *personsChecker) IsPersonsExists(ctx context.Context, ids []int32) (exists bool, notExistsIDs []int32, err error) {
	res, err := c.client.IsPersonsExists(ctx, &movies_persons_service.IsPersonsExistsRequest{PersonsIDs: convertIntSliceIntoString(ids)})
	if err != nil {
		return false, []int32{}, err
	}

	return res.PersonsExists, res.NotExistIDs, nil
}

type moviesChecker struct {
	client movies_service.MoviesServiceV1Client
	conn   *grpc.ClientConn
}

func NewMoviesChecker(addr string) (*moviesChecker, error) {
	conn, err := getGrpcConnection(addr)
	if err != nil {
		return nil, err
	}

	return &moviesChecker{client: movies_service.NewMoviesServiceV1Client(conn), conn: conn}, nil
}

func (c *moviesChecker) IsMovieExists(ctx context.Context, id int32) (exists bool, err error) {
	res, err := c.client.IsMovieExists(ctx, &movies_service.IsMovieExistsRequest{MovieID: id})
	if err != nil {
		return false, err
	}

	return res.Exists, nil
}

func (c *moviesChecker) Shutdown() error {
	return c.conn.Close()
}

type PersonsChecker interface {
	IsPersonsExists(ctx context.Context, ids []int32) (exists bool, notExistsIDs []int32, err error)
}

type MoviesChecker interface {
	IsMovieExists(ctx context.Context, id int32) (exists bool, err error)
}

type existanceChecker struct {
	moviesCheck  MoviesChecker
	personsCheck PersonsChecker
	errorHandler errorHandler
	logger       *logrus.Logger
}

func NewExistanceChecker(moviesCheck MoviesChecker, personsCheck PersonsChecker, logger *logrus.Logger) *existanceChecker {
	errorHandler := newErrorHandler(logger)
	return &existanceChecker{
		logger:       logger,
		errorHandler: errorHandler,
		moviesCheck:  moviesCheck,
		personsCheck: personsCheck,
	}
}

func convertIntSliceIntoString(nums []int32) string {
	var str = make([]string, len(nums))
	for i, num := range nums {
		str[i] = fmt.Sprint(num)
	}
	return strings.Join(str, ",")
}

func (c *existanceChecker) CheckPersons(ctx context.Context, persons []*casts_service.ActorParam) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "existanceChecker.CheckPersons")
	defer span.Finish()

	var ids = make([]int32, 0, len(persons))
	for _, person := range persons {
		if person == nil {
			continue
		}

		ids = append(ids, person.Id)
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		exists, notFound, err := c.personsCheck.IsPersonsExists(ctx, ids)
		if err != nil {
			return c.errorHandler.createErrorResponceWithSpan(span, ErrInternal, "can't check persons existence "+err.Error())
		} else if !exists {
			return c.errorHandler.createExtendedErrorResponceWithSpan(span, ErrInvalidArgument, "",
				fmt.Sprintf("persons with ids %s not found", convertIntSliceIntoString(notFound)))
		}
	}

	return nil
}

func (c *existanceChecker) CheckMovie(ctx context.Context, id int32) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "existanceChecker.CheckPersons")
	defer span.Finish()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		exists, err := c.moviesCheck.IsMovieExists(ctx, id)
		if err != nil {
			return c.errorHandler.createErrorResponceWithSpan(span, ErrInternal, "can't check movie existence "+err.Error())
		} else if !exists {
			return c.errorHandler.createExtendedErrorResponceWithSpan(span, ErrInvalidArgument, "",
				fmt.Sprintf("movie with id %d not found", id))
		}
	}

	return nil
}

func (c *existanceChecker) CheckExistance(ctx context.Context, persons []*casts_service.ActorParam, movieID int32) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "existanceChecker.CheckExistance")
	defer span.Finish()

	c.logger.Info("start checking existance")
	var wg sync.WaitGroup
	reqCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	var errCh = make(chan error, 2)

	wg.Add(2)
	go func() {
		defer wg.Done()
		errCh <- c.CheckMovie(reqCtx, movieID)
	}()
	go func() {
		defer wg.Done()
		errCh <- c.CheckPersons(reqCtx, persons)
	}()
	go func() {
		wg.Wait()
		c.logger.Debug("channel closed")
		close(errCh)
	}()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err, open := <-errCh:
			if !open {
				return nil
			} else if err != nil {
				return err
			}
		}
	}
}

func getGrpcConnection(addr string) (*grpc.ClientConn, error) {
	return grpc.Dial(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(
			otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
		grpc.WithStreamInterceptor(
			otgrpc.OpenTracingStreamClientInterceptor(opentracing.GlobalTracer())),
	)
}
