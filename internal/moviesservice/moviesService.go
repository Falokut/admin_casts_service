package moviesservice

import (
	"context"
	"errors"

	"github.com/Falokut/admin_casts_service/internal/config"
	"github.com/Falokut/admin_casts_service/internal/models"
	movies_service "github.com/Falokut/admin_movies_service/pkg/admin_movies_service/v1/protos"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type moviesService struct {
	client movies_service.MoviesServiceV1Client
	logger *logrus.Logger
	conn   *grpc.ClientConn
}

func NewMoviesService(addr string, cfg config.ConnectionSecureConfig, logger *logrus.Logger) (*moviesService, error) {
	conn, err := getGrpcConnection(addr, cfg)
	if err != nil {
		return nil, err
	}

	return &moviesService{
		client: movies_service.NewMoviesServiceV1Client(conn),
		conn:   conn,
		logger: logger,
	}, nil
}

func (c *moviesService) IsMovieExists(ctx context.Context, id int32) (exists bool, err error) {
	defer c.handleError(ctx, &err, "IsMovieExists")
	res, err := c.client.IsMovieExists(ctx, &movies_service.IsMovieExistsRequest{MovieID: id})
	if err != nil {
		return false, err
	}

	return res.Exists, nil
}

func (c *moviesService) Shutdown() {
	err := c.conn.Close()
	if err != nil {
		c.logger.Errorf("error while shutting down movies service %v", err)
	}
}

func (c *moviesService) handleError(ctx context.Context, err *error, functionName string) {
	if ctx.Err() != nil {
		var code models.ErrorCode
		switch {
		case errors.Is(ctx.Err(), context.Canceled):
			code = models.Canceled
		case errors.Is(ctx.Err(), context.DeadlineExceeded):
			code = models.DeadlineExceeded
		}
		*err = models.Error(code, ctx.Err().Error())
		return
	}
	if err == nil || *err == nil {
		return
	}

	c.logError(*err, functionName)
	e := *err
	switch status.Code(*err) {
	case codes.Canceled:
		*err = models.Error(models.Canceled, e.Error())
	case codes.DeadlineExceeded:
		*err = models.Error(models.DeadlineExceeded, e.Error())
	case codes.Internal:
		*err = models.Error(models.Internal, "can't check existence")
	default:
		*err = models.Error(models.Unknown, e.Error())
	}
}

func (c *moviesService) logError(err error, functionName string) {
	if err == nil {
		return
	}

	var sericeErr = &models.ServiceError{}
	if errors.As(err, &sericeErr) {
		c.logger.WithFields(
			logrus.Fields{
				"error.function.name": functionName,
				"error.msg":           sericeErr.Msg,
				"code":                sericeErr.Code,
			},
		).Error("movies service error occurred")
	} else {
		c.logger.WithFields(
			logrus.Fields{
				"error.function.name": functionName,
				"error.msg":           err.Error(),
			},
		).Error("movies service error occurred")
	}
}

func getGrpcConnection(addr string, cfg config.ConnectionSecureConfig) (*grpc.ClientConn, error) {
	creds, err := cfg.GetGrpcTransportCredentials()
	if err != nil {
		return nil, err
	}
	return grpc.Dial(addr,
		creds,
		grpc.WithUnaryInterceptor(
			otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
		grpc.WithStreamInterceptor(
			otgrpc.OpenTracingStreamClientInterceptor(opentracing.GlobalTracer())),
	)
}
