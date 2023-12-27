package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/Falokut/admin_casts_service/internal/repository"
	casts_service "github.com/Falokut/admin_casts_service/pkg/admin_casts_service/v1/protos"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type castsService struct {
	casts_service.UnimplementedCastsServiceV1Server
	logger       *logrus.Logger
	repo         repository.CastsRepository
	errorHandler errorHandler
}

func NewCastsService(logger *logrus.Logger, repo repository.CastsRepository) *castsService {
	errorHandler := newErrorHandler(logger)
	return &castsService{
		logger:       logger,
		repo:         repo,
		errorHandler: errorHandler,
	}
}

func (s *castsService) GetCast(ctx context.Context, in *casts_service.GetCastRequest) (*casts_service.Cast, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsService.GetCast")
	defer span.Finish()

	cast, err := s.repo.GetCast(ctx, in.MovieId)
	if errors.Is(err, repository.ErrNotFound) {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrNotFound, "")
	} else if err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInternal, err.Error())
	}

	span.SetTag("grpc.status", codes.OK)
	return convertCastToProto(cast), nil
}

func (s *castsService) GetCasts(ctx context.Context, in *casts_service.GetCastsRequest) (*casts_service.Casts, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsService.GetCasts")
	defer span.Finish()

	offset := in.Limit * (in.Page - 1)
	if err := validateLimitAndPage(in.Page, in.Limit); err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInvalidArgument, err.Error())
	}

	var casts []repository.Cast
	var err error
	if in.MoviesIDs == "" {
		casts, err = s.repo.GetAllCasts(ctx, in.Limit, offset)
	} else {
		ids, er := s.convertStringToSlice(in.MoviesIDs)
		if er != nil {
			ext.LogError(span, er)
			span.SetTag("grpc.status", status.Code(er))
			return nil, er
		}

		casts, err = s.repo.GetCasts(ctx, ids, in.Limit, offset)

	}

	if errors.Is(err, repository.ErrNotFound) {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrNotFound, "")
	} else if err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInternal, err.Error())
	}

	span.SetTag("grpc.status", codes.OK)
	return convertCastsToProto(casts), nil
}

func (s *castsService) SearchCastByLabel(ctx context.Context,
	in *casts_service.SearchCastByLabelRequest) (*casts_service.Casts, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsService.SearchCastByLabel")
	defer span.Finish()

	offset := in.Limit * (in.Page - 1)
	if err := validateLimitAndPage(in.Page, in.Limit); err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInvalidArgument, err.Error())
	}

	casts, err := s.repo.SearchCastByLabel(ctx, in.Label, in.Limit, offset)
	if errors.Is(err, repository.ErrNotFound) {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrNotFound, "")
	} else if err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInternal, err.Error())
	}

	span.SetTag("grpc.status", codes.OK)
	return convertCastsToProto(casts), nil
}

func (s *castsService) CreateCast(ctx context.Context, in *casts_service.CreateCastRequest) (*emptypb.Empty, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsService.CreateCast")
	defer span.Finish()

	exist, id, err := s.repo.IsCastExist(ctx, in.MovieID)
	if err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInternal, err.Error())
	} else if exist {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrAlreadyExists,
			fmt.Sprintf("cast is already exists check cast with id %d", id))
	}

	if err := s.repo.CreateCast(ctx, in.MovieID, in.CastLabel, in.ActorsIDs); err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInternal, err.Error())
	}

	span.SetTag("grpc.status", codes.OK)
	return &emptypb.Empty{}, nil
}

func (s *castsService) UpdateLabelForCast(ctx context.Context,
	in *casts_service.UpdateLabelForCastRequest) (*emptypb.Empty, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsService.UpdateLabelForCast")
	defer span.Finish()

	_, err := s.checkCastExisting(ctx, in.MovieID)
	if err != nil {
		ext.LogError(span, err)
		span.SetTag("grpc.status", status.Code(err))
		return nil, err
	}

	if err := s.repo.UpdateLabelForCast(ctx, in.MovieID, in.Label); err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInternal, err.Error())
	}

	span.SetTag("grpc.status", codes.OK)
	return &emptypb.Empty{}, nil
}

func (s *castsService) AddActorsToTheCast(ctx context.Context,
	in *casts_service.AddActorsToTheCastRequest) (*emptypb.Empty, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsService.AddActorsToTheCast")
	defer span.Finish()

	_, err := s.checkCastExisting(ctx, in.MovieID)
	if err != nil {
		ext.LogError(span, err)
		span.SetTag("grpc.status", status.Code(err))
		return nil, err
	}

	if err := s.repo.AddActorsToTheCast(ctx, in.MovieID, in.ActorsIDs); err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInternal, err.Error())
	}

	span.SetTag("grpc.status", codes.OK)
	return &emptypb.Empty{}, nil
}

func (s *castsService) RemoveActorsFromTheCast(ctx context.Context,
	in *casts_service.RemoveActorsFromTheCastRequest) (*emptypb.Empty, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsService.RemoveActorsFromTheCast")
	defer span.Finish()

	_, err := s.checkCastExisting(ctx, in.MovieID)
	if err != nil {
		ext.LogError(span, err)
		span.SetTag("grpc.status", status.Code(err))
		return nil, err
	}

	ids, err := s.convertStringToSlice(in.ActorsIDs)
	if err != nil {
		ext.LogError(span, err)
		span.SetTag("grpc.status", status.Code(err))
		return nil, err
	}

	if err := s.repo.RemoveActorsFromCast(ctx, in.MovieID, ids); err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInternal, err.Error())
	}

	span.SetTag("grpc.status", codes.OK)
	return &emptypb.Empty{}, nil
}

func (s *castsService) DeleteCast(ctx context.Context,
	in *casts_service.DeleteCastRequest) (*emptypb.Empty, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsService.DeleteCast")
	defer span.Finish()

	_, err := s.checkCastExisting(ctx, in.MovieId)
	if err != nil {
		ext.LogError(span, err)
		span.SetTag("grpc.status", status.Code(err))
		return nil, err
	}

	if err := s.repo.DeleteCast(ctx, in.MovieId); err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInternal, err.Error())
	}

	span.SetTag("grpc.status", codes.OK)
	return &emptypb.Empty{}, nil
}

func (s *castsService) convertStringToSlice(str string) ([]int32, error) {
	str = strings.TrimSpace(strings.ReplaceAll(str, `"`, ""))
	if str == "" {
		return nil, s.errorHandler.createErrorResponce(ErrInvalidArgument, "ids param mustn't be empty")
	}

	if err := checkParam(str); err != nil {
		return nil, s.errorHandler.createErrorResponce(ErrInvalidArgument, err.Error())
	}

	strs := strings.Split(str, ",")
	var nums = make([]int32, 0, len(str))
	for _, str := range strs {
		num, _ := strconv.Atoi(str) // already checked by regexp
		nums = append(nums, int32(num))
	}
	return nums, nil
}

func (s *castsService) checkCastExisting(ctx context.Context, id int32) (int32, error) {
	exist, id, err := s.repo.IsCastExist(ctx, id)
	if err != nil {
		return 0, s.errorHandler.createErrorResponce(ErrInternal, err.Error())
	} else if !exist {
		return 0, s.errorHandler.createErrorResponce(ErrNotFound, "")
	}

	return id, nil
}

func convertCastsToProto(casts []repository.Cast) *casts_service.Casts {
	proto := &casts_service.Casts{}
	proto.Casts = make(map[int32]*casts_service.Cast, len(casts))

	for _, cast := range casts {
		proto.Casts[cast.ID] = convertCastToProto(cast)
	}

	return proto
}

func convertCastToProto(cast repository.Cast) *casts_service.Cast {
	actorsIDs := strings.Split(cast.Actors, ",")
	var Actors = make([]int32, 0, len(actorsIDs))
	for _, id := range actorsIDs {
		actorID, err := strconv.Atoi(id)
		if err != nil {
			continue
		}
		Actors = append(Actors, int32(actorID))
	}
	return &casts_service.Cast{
		MovieID:   cast.ID,
		CastLabel: cast.Label,
		ActorsIDs: Actors,
	}
}
