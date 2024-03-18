package handler

import (
	"context"
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/Falokut/admin_casts_service/internal/models"
	"github.com/Falokut/admin_casts_service/internal/service"
	casts_service "github.com/Falokut/admin_casts_service/pkg/admin_casts_service/v1/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ExistanceChecker interface {
	CheckPersons(ctx context.Context, persons []*casts_service.PersonParam) error
	CheckExistance(ctx context.Context, persons []*casts_service.PersonParam, movieID int32) error
}

type castsServiceHandler struct {
	casts_service.UnimplementedCastsServiceV1Server
	s service.CastsService
}

func NewCastsService(s service.CastsService) *castsServiceHandler {
	return &castsServiceHandler{s: s}
}

func (h *castsServiceHandler) GetCast(ctx context.Context, in *casts_service.GetCastRequest) (cast *casts_service.Cast, err error) {
	defer h.handleError(&err)

	res, err := h.s.GetCast(ctx, in.MovieId)
	if err != nil {
		return
	}

	cast = &casts_service.Cast{
		MovieID:   res[0].ID,
		CastLabel: res[0].Label,
		Persons:   make([]*casts_service.Person, len(res))}
	for i := range res {
		cast.Persons[i] = &casts_service.Person{
			ID: res[i].PersonID,
			Profession: &casts_service.Profession{
				ID:   res[i].ProfessionID,
				Name: res[i].ProfessionName,
			},
		}
	}

	return
}

func (h *castsServiceHandler) GetCasts(ctx context.Context,
	in *casts_service.GetCastsRequest) (casts *casts_service.Casts, err error) {
	defer h.handleError(&err)

	offset := in.Limit * (in.Page - 1)
	var res []models.Cast

	var ids []int32
	if in.MoviesIDs != "" {
		ids, err = h.convertStringToSlice(in.MoviesIDs)
		if err != nil {
			return
		}
	}

	res, err = h.s.GetCasts(ctx, ids, in.Limit, offset)
	if err != nil {
		return
	}

	return convertCastsToProto(res), nil
}

func (h *castsServiceHandler) SearchCastByLabel(ctx context.Context,
	in *casts_service.SearchCastByLabelRequest) (labels *casts_service.CastsLabels, err error) {
	defer h.handleError(&err)

	offset := in.Limit * (in.Page - 1)

	casts, err := h.s.SearchCastByLabel(ctx, in.Label, in.Limit, offset)
	if err != nil {
		return
	}

	labels = &casts_service.CastsLabels{
		Casts: make([]*casts_service.CastLabel, len(casts)),
	}

	for i, cast := range casts {
		labels.Casts[i] = &casts_service.CastLabel{
			MovieID:   cast.ID,
			CastLabel: cast.Label,
		}
	}

	return
}

func (h *castsServiceHandler) CreateCast(ctx context.Context, in *casts_service.CreateCastRequest) (_ *emptypb.Empty, err error) {
	defer h.handleError(&err)

	var persons = convertGrpcToModelsPersons(in.Persons)
	if err = h.s.CreateCast(ctx, in.MovieID, in.CastLabel, persons); err != nil {
		return
	}

	return &emptypb.Empty{}, nil
}

func (h *castsServiceHandler) UpdateLabelForCast(ctx context.Context,
	in *casts_service.UpdateLabelForCastRequest) (_ *emptypb.Empty, err error) {
	defer h.handleError(&err)
	err = h.s.UpdateLabelForCast(ctx, in.MovieID, in.Label)
	if err != nil {
		return
	}

	return &emptypb.Empty{}, nil
}

func (h *castsServiceHandler) AddPersonsToTheCast(ctx context.Context,
	in *casts_service.AddPersonsToTheCastRequest) (_ *emptypb.Empty, err error) {
	defer h.handleError(&err)

	var persons = convertGrpcToModelsPersons(in.Persons)
	err = h.s.AddPersonsToTheCast(ctx, in.MovieID, persons)
	if err != nil {
		return
	}

	return &emptypb.Empty{}, nil
}

func (h *castsServiceHandler) RemovePersonsFromTheCast(ctx context.Context,
	in *casts_service.RemovePersonsFromTheCastRequest) (_ *emptypb.Empty, err error) {
	defer h.handleError(&err)

	var persons = convertGrpcToModelsPersons(in.Persons)
	err = h.s.RemovePersonsFromCast(ctx, in.MovieID, persons)
	if err != nil {
		return
	}

	return &emptypb.Empty{}, nil
}

func (h *castsServiceHandler) DeleteCast(ctx context.Context,
	in *casts_service.DeleteCastRequest) (_ *emptypb.Empty, err error) {
	defer h.handleError(&err)

	err = h.s.DeleteCast(ctx, in.MovieId)
	if err != nil {
		return
	}

	return &emptypb.Empty{}, nil
}

func (h *castsServiceHandler) GetProfessions(ctx context.Context,
	_ *emptypb.Empty) (professions *casts_service.Professions, err error) {
	defer h.handleError(&err)

	res, err := h.s.GetProfessions(ctx)
	if err != nil {
		return
	}

	professions = &casts_service.Professions{
		Professions: make([]*casts_service.Profession, len(res)),
	}

	for i := range res {
		professions.Professions[i] = &casts_service.Profession{ID: res[i].ID, Name: res[i].Name}
	}

	return
}

func (h *castsServiceHandler) CreateProfession(ctx context.Context,
	in *casts_service.CreateProfessionRequest) (res *casts_service.CreateProfessionResponse, err error) {
	defer h.handleError(&err)

	id, err := h.s.CreateProfession(ctx, in.Name)
	if err != nil {
		return
	}

	return &casts_service.CreateProfessionResponse{Id: id}, nil
}

func (h *castsServiceHandler) DeleteProfession(ctx context.Context,
	in *casts_service.DeleteProfessionRequest) (_ *emptypb.Empty, err error) {
	defer h.handleError(&err)

	err = h.s.DeleteProfession(ctx, in.Id)
	if err != nil {
		return
	}

	return &emptypb.Empty{}, nil
}

func (h *castsServiceHandler) UpdateProfession(ctx context.Context,
	in *casts_service.UpdateProfessionRequest) (_ *emptypb.Empty, err error) {
	defer h.handleError(&err)

	err = h.s.UpdateProfession(ctx, in.Id, in.Name)
	if err != nil {
		return
	}

	return &emptypb.Empty{}, nil
}

func (h *castsServiceHandler) handleError(err *error) {
	if err == nil || *err == nil {
		return
	}

	serviceErr := &models.ServiceError{}
	if errors.As(*err, &serviceErr) {
		*err = status.Error(convertServiceErrCodeToGrpc(serviceErr.Code), serviceErr.Msg)
	} else if _, ok := status.FromError(*err); !ok {
		e := *err
		*err = status.Error(codes.Unknown, e.Error())
	}
}

func convertServiceErrCodeToGrpc(code models.ErrorCode) codes.Code {
	switch code {
	case models.Internal:
		return codes.Internal
	case models.InvalidArgument:
		return codes.InvalidArgument
	case models.Unauthenticated:
		return codes.Unauthenticated
	case models.Conflict:
		return codes.AlreadyExists
	case models.NotFound:
		return codes.NotFound
	case models.Canceled:
		return codes.Canceled
	case models.DeadlineExceeded:
		return codes.DeadlineExceeded
	case models.PermissionDenied:
		return codes.PermissionDenied
	default:
		return codes.Unknown
	}
}

func (h *castsServiceHandler) convertStringToSlice(str string) ([]int32, error) {
	str = strings.TrimSpace(strings.ReplaceAll(str, `"`, ""))
	if str == "" {
		return nil, status.Error(codes.InvalidArgument, "ids param mustn't be empty")
	}

	ok := regexp.MustCompile(`^\d+(,\d+)*$`).MatchString(str)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "ids must contain only digits and commas")
	}

	strs := strings.Split(str, ",")
	var nums = make([]int32, 0, len(str))
	for _, str := range strs {
		num, _ := strconv.Atoi(str) // already checked by regexp
		nums = append(nums, int32(num))
	}
	return nums, nil
}

func convertGrpcToModelsPersons(persons []*casts_service.PersonParam) []models.Person {
	var converted = make([]models.Person, len(persons))
	for i := range persons {
		converted[i] = models.Person{
			ID:           persons[i].ID,
			ProfessionID: persons[i].ProfessionID,
		}
	}
	return converted
}
func convertCastsToProto(casts []models.Cast) *casts_service.Casts {
	type CastInfo struct {
		Persons   []*casts_service.Person
		CastLabel string
	}
	const preallocateCastPersonsCapacity = 6

	var protoPersonsByCast = make(map[int32]CastInfo)
	for i := range casts {
		castInfo, ok := protoPersonsByCast[casts[i].ID]
		if !ok {
			castInfo = CastInfo{
				CastLabel: casts[i].Label}
			castInfo.Persons = make([]*casts_service.Person, 0, preallocateCastPersonsCapacity)
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
