package personsservice

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/Falokut/admin_casts_service/internal/config"
	"github.com/Falokut/admin_casts_service/internal/models"
	movies_persons_service "github.com/Falokut/admin_movies_persons_service/pkg/admin_movies_persons_service/v1/protos"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type personsService struct {
	client movies_persons_service.MoviesPersonsServiceV1Client
	logger *logrus.Logger
	conn   *grpc.ClientConn
}

func NewPersonsService(addr string, cfg config.ConnectionSecureConfig, logger *logrus.Logger) (*personsService, error) {
	conn, err := getGrpcConnection(addr, cfg)
	if err != nil {
		return nil, err
	}

	return &personsService{
		client: movies_persons_service.NewMoviesPersonsServiceV1Client(conn),
		conn:   conn,
		logger: logger,
	}, nil
}

func (c *personsService) Shutdown() {
	err := c.conn.Close()
	if err != nil {
		c.logger.Errorf("error while shutting down persons checker %v", err)
	}
}

func (c *personsService) IsPersonsExists(ctx context.Context, ids []int32) (exists bool, notExistsIDs []int32, err error) {
	defer c.handleError(ctx, &err, "IsPersonsExists")

	res, err := c.client.IsPersonsExists(ctx,
		&movies_persons_service.IsPersonsExistsRequest{
			PersonsIDs: convertIntSliceIntoString(ids),
		})
	if err != nil {
		return false, []int32{}, err
	}

	return res.PersonsExists, res.NotExistIDs, nil
}

func (c *personsService) logError(err error, functionName string) {
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
		).Error("persons checker error occurred")
	} else {
		c.logger.WithFields(
			logrus.Fields{
				"error.function.name": functionName,
				"error.msg":           err.Error(),
			},
		).Error("persons checker error occurred")
	}
}

func (c *personsService) handleError(ctx context.Context, err *error, functionName string) {
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
		*err = models.Error(models.Internal, "can't check persons existence")
	default:
		*err = models.Error(models.Unknown, e.Error())
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

func convertIntSliceIntoString(nums []int32) string {
	var str = make([]string, len(nums))
	for i, num := range nums {
		str[i] = fmt.Sprint(num)
	}
	return strings.Join(str, ",")
}
