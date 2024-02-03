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

type ExistanceChecker interface {
	CheckPersons(ctx context.Context, persons []*casts_service.PersonParam) error
	CheckExistance(ctx context.Context, persons []*casts_service.PersonParam, movieID int32) error
}

type castsService struct {
	casts_service.UnimplementedCastsServiceV1Server
	checker      ExistanceChecker
	logger       *logrus.Logger
	repo         repository.CastsRepository
	errorHandler errorHandler
}

func NewCastsService(logger *logrus.Logger, repo repository.CastsRepository, checker ExistanceChecker) *castsService {
	errorHandler := newErrorHandler(logger)
	return &castsService{
		logger:       logger,
		repo:         repo,
		checker:      checker,
		errorHandler: errorHandler,
	}
}

func (s *castsService) GetCast(ctx context.Context, in *casts_service.GetCastRequest) (*casts_service.Cast, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsService.GetCast")
	defer span.Finish()

	cast, err := s.repo.GetCast(ctx, in.MovieId)
	if err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInternal, err.Error())
	} else if len(cast) == 0 {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrNotFound, "")
	}

	span.SetTag("grpc.status", codes.OK)

	var protoCast = &casts_service.Cast{MovieID: cast[0].ID, CastLabel: cast[0].Label}
	for i := range cast {
		protoCast.Persons = append(protoCast.Persons, &casts_service.Person{
			ID: cast[i].PersonID,
			Profession: &casts_service.Profession{
				ID:   cast[i].ProfessionID,
				Name: cast[i].ProfessionName,
			},
		})
	}

	return protoCast, nil
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
	in *casts_service.SearchCastByLabelRequest) (*casts_service.CastsLabels, error) {
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

	proto := &casts_service.CastsLabels{}
	proto.Casts = make([]*casts_service.CastLabel, len(casts))
	for i, cast := range casts {
		proto.Casts[i] = &casts_service.CastLabel{
			MovieID:   cast.ID,
			CastLabel: cast.Label,
		}
	}

	span.SetTag("grpc.status", codes.OK)
	return proto, nil
}

func (s *castsService) CreateCast(ctx context.Context, in *casts_service.CreateCastRequest) (*emptypb.Empty, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsService.CreateCast")
	defer span.Finish()

	if len(in.Persons) == 0 {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInvalidArgument, "persons musn't be empty")
	}
	exist, id, err := s.repo.IsCastExist(ctx, in.MovieID)
	if err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInternal, err.Error())
	} else if exist {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrAlreadyExists,
			fmt.Sprintf("cast is already exists check cast with id %d", id))
	}

	err = s.checker.CheckExistance(ctx, in.Persons, in.MovieID)
	if err != nil {
		ext.LogError(span, err)
		span.SetTag("grpc.status", status.Code(err))
		return nil, err
	}

	var actors = make([]repository.Person, len(in.Persons))
	for i, actor := range in.Persons {
		actors[i] = repository.Person{
			ID:           actor.Id,
			ProfessionID: actor.ProfessionID,
		}
	}
	if err := s.repo.CreateCast(ctx, in.MovieID, in.CastLabel, actors); err != nil {
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

func (s *castsService) AddPersonsToTheCast(ctx context.Context,
	in *casts_service.AddPersonsToTheCastRequest) (*emptypb.Empty, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsService.AddPersonsToTheCast")
	defer span.Finish()

	if len(in.Persons) == 0 {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInvalidArgument, "actors mustn't be emtpy")
	}
	_, err := s.checkCastExisting(ctx, in.MovieID)
	if err != nil {
		ext.LogError(span, err)
		span.SetTag("grpc.status", status.Code(err))
		return nil, err
	}

	err = s.checker.CheckPersons(ctx, in.Persons)
	if err != nil {
		ext.LogError(span, err)
		span.SetTag("grpc.status", status.Code(err))
		return nil, err
	}

	var actors = make([]repository.Person, len(in.Persons))
	for i, actor := range in.Persons {
		actors[i] = repository.Person{
			ID:           actor.Id,
			ProfessionID: actor.ProfessionID,
		}
	}

	if err := s.repo.AddPersonsToTheCast(ctx, in.MovieID, actors); err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInternal, err.Error())
	}

	span.SetTag("grpc.status", codes.OK)
	return &emptypb.Empty{}, nil
}

func (s *castsService) RemovePersonsFromTheCast(ctx context.Context,
	in *casts_service.RemovePersonsFromTheCastRequest) (*emptypb.Empty, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsService.RemovePersonsFromTheCast")
	defer span.Finish()

	if len(in.Persons) == 0 {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInvalidArgument, "actors musn't be empty")
	}

	_, err := s.checkCastExisting(ctx, in.MovieID)
	if err != nil {
		ext.LogError(span, err)
		span.SetTag("grpc.status", status.Code(err))
		return nil, err
	}

	var actors = make([]repository.Person, len(in.Persons))
	for i, actor := range in.Persons {
		actors[i] = repository.Person{
			ID:           actor.Id,
			ProfessionID: actor.ProfessionID,
		}
	}

	if err := s.repo.RemovePersonsFromCast(ctx, in.MovieID, actors); err != nil {
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
	type CastInfo struct {
		Persons   []*casts_service.Person
		CastLabel string
	}

	var protoPersonsByCast = make(map[int32]CastInfo)
	for i := range casts {
		castInfo, ok := protoPersonsByCast[casts[i].ID]
		if !ok {
			castInfo = CastInfo{
				CastLabel: casts[i].Label}
			castInfo.Persons = make([]*casts_service.Person, 0, 6)
		}
		castInfo.Persons = append(castInfo.Persons, &casts_service.Person{
			ID: casts[i].PersonID,
			Profession: &casts_service.Profession{
				ID:   casts[i].ProfessionID,
				Name: casts[i].ProfessionName,
			},
		})
		protoPersonsByCast[casts[i].ID] = castInfo
	}
	proto := &casts_service.Casts{}
	proto.Casts = make([]*casts_service.Cast, 0, len(casts))
	for castID, info := range protoPersonsByCast {
		proto.Casts = append(proto.Casts, &casts_service.Cast{
			MovieID:   castID,
			CastLabel: info.CastLabel,
			Persons:   info.Persons,
		})
	}

	return proto
}

func (s *castsService) GetProfessions(ctx context.Context,
	in *emptypb.Empty) (*casts_service.Professions, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsService.GetProfessions")
	defer span.Finish()

	professions, err := s.repo.GetAllProfessions(ctx)
	if err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInternal, err.Error())
	}
	if len(professions) == 0 {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrNotFound, "professions not found, table might be empty")
	}

	var protoProfessions = &casts_service.Professions{}
	protoProfessions.Professions = make([]*casts_service.Profession, len(professions))
	for i, profession := range professions {
		protoProfessions.Professions[i] = &casts_service.Profession{ID: profession.ID, Name: profession.Name}
	}

	span.SetTag("grpc.status", codes.OK)
	return protoProfessions, nil
}

func (s *castsService) CreateProfession(ctx context.Context,
	in *casts_service.CreateProfessionRequest) (*casts_service.CreateProfessionResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsService.CreateProfession")
	defer span.Finish()

	exist, id, err := s.repo.IsProfessionWithNameExists(ctx, in.Name)
	if err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInternal, err.Error())
	}
	if exist {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrAlreadyExists, fmt.Sprintf("profession already exists, check profession with id %d", id))
	}

	id, err = s.repo.CreateProfession(ctx, in.Name)
	if err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInternal, err.Error())
	}

	span.SetTag("grpc.status", codes.OK)
	return &casts_service.CreateProfessionResponse{Id: id}, nil
}

func (s *castsService) DeleteProfession(ctx context.Context,
	in *casts_service.DeleteProfessionRequest) (*emptypb.Empty, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsService.DeleteProfession")
	defer span.Finish()

	err := s.repo.DeleteProfession(ctx, in.Id)
	if errors.Is(err, repository.ErrNotFound) {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrNotFound, "")
	} else if err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInternal, err.Error())
	}

	span.SetTag("grpc.status", codes.OK)
	return &emptypb.Empty{}, nil
}

func (s *castsService) UpdateProfession(ctx context.Context,
	in *casts_service.UpdateProfessionRequest) (*emptypb.Empty, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsService.UpdateProfession")
	defer span.Finish()

	exist, id, err := s.repo.IsProfessionWithNameExists(ctx, in.Name)
	if err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInternal, err.Error())
	}
	if exist {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrAlreadyExists, fmt.Sprintf("profession already exists, check profession with id %d", id))
	}

	exist, err = s.repo.IsProfessionExists(ctx, in.Id)
	if err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInternal, err.Error())
	}
	if !exist {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrNotFound, "profession not found")
	}

	err = s.repo.UpdateProfession(ctx, in.Id, in.Name)
	if err != nil {
		return nil, s.errorHandler.createErrorResponceWithSpan(span, ErrInternal, err.Error())
	}

	span.SetTag("grpc.status", codes.OK)
	return &emptypb.Empty{}, nil
}
