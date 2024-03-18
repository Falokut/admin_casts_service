package service

import (
	"context"

	"github.com/Falokut/admin_casts_service/internal/models"
	"github.com/Falokut/admin_casts_service/internal/repository"
	casts_service "github.com/Falokut/admin_casts_service/pkg/admin_casts_service/v1/protos"
	"github.com/sirupsen/logrus"
)

type ExistanceChecker interface {
	CheckPersons(ctx context.Context, persons []models.Person) error
	CheckExistance(ctx context.Context, persons []models.Person, movieID int32) error
}

type CastsService interface {
	GetCast(ctx context.Context, id int32) ([]models.Cast, error)
	GetCasts(ctx context.Context, ids []int32, limit, offset int32) ([]models.Cast, error)
	SearchCastByLabel(ctx context.Context, label string, limit, offset int32) ([]models.CastLabel, error)
	CreateCast(ctx context.Context, id int32, label string, persons []models.Person) error
	DeleteCast(ctx context.Context, id int32) error

	UpdateLabelForCast(ctx context.Context, id int32, label string) error
	AddPersonsToTheCast(ctx context.Context, id int32, persons []models.Person) error
	RemovePersonsFromCast(ctx context.Context, id int32, persons []models.Person) error

	CreateProfession(ctx context.Context, name string) (int32, error)
	UpdateProfession(ctx context.Context, id int32, name string) error
	DeleteProfession(ctx context.Context, id int32) error
	GetProfessions(ctx context.Context) ([]models.Profession, error)
}

type castsService struct {
	casts_service.UnimplementedCastsServiceV1Server
	checker ExistanceChecker
	logger  *logrus.Logger
	repo    repository.CastsRepository
}

func NewCastsService(logger *logrus.Logger, repo repository.CastsRepository, checker ExistanceChecker) *castsService {
	return &castsService{
		logger:  logger,
		repo:    repo,
		checker: checker,
	}
}

func (s *castsService) GetCast(ctx context.Context, id int32) ([]models.Cast, error) {
	return s.repo.GetCast(ctx, id)
}

func validateOffsetAndLimit(limit, offset int32) error {
	if limit < 10 || limit > 100 {
		return models.Error(models.InvalidArgument, "limit must be bigger than or equal 10 and less than or equal 100")
	}
	if offset < 0 {
		return models.Error(models.InvalidArgument, "invalid offset value, it must be bigger than or equal 0")
	}
	return nil
}

func (s *castsService) GetCasts(ctx context.Context, ids []int32, limit, offset int32) (casts []models.Cast, err error) {
	err = validateOffsetAndLimit(limit, offset)
	if err != nil {
		return
	}

	if len(ids) == 0 {
		return s.repo.GetAllCasts(ctx, limit, offset)
	}

	return s.repo.GetCasts(ctx, ids, limit, offset)
}

func (s *castsService) SearchCastByLabel(ctx context.Context, label string,
	limit, offset int32) (labels []models.CastLabel, err error) {
	err = validateOffsetAndLimit(limit, offset)
	if err != nil {
		return
	}

	return s.repo.SearchCastByLabel(ctx, label, limit, offset)
}

func (s *castsService) CreateCast(ctx context.Context, movieID int32,
	castLabel string, persons []models.Person) (err error) {
	if len(persons) == 0 {
		err = models.Error(models.InvalidArgument, "persons mustn't be empty")
		return
	}

	exist, err := s.repo.IsCastExist(ctx, movieID)
	if err != nil {
		return
	}
	if exist {
		err = models.Errorf(models.Conflict, "cast is already exists check cast with id %d", movieID)
		return
	}

	err = s.checker.CheckExistance(ctx, persons, movieID)
	if err != nil {
		return
	}

	return s.repo.CreateCast(ctx, movieID, castLabel, persons)
}

func (s *castsService) UpdateLabelForCast(ctx context.Context, movieID int32, castLabel string) (err error) {
	err = s.checkCastExisting(ctx, movieID)
	if err != nil {
		return
	}

	return s.repo.UpdateLabelForCast(ctx, movieID, castLabel)
}

func (s *castsService) AddPersonsToTheCast(ctx context.Context, movieID int32, persons []models.Person) (err error) {
	if len(persons) == 0 {
		err = models.Error(models.InvalidArgument, "persons mustn't be empty")
		return
	}

	err = s.checkCastExisting(ctx, movieID)
	if err != nil {
		return
	}

	err = s.checker.CheckPersons(ctx, persons)
	if err != nil {
		return
	}

	return s.repo.AddPersonsToTheCast(ctx, movieID, persons)
}

func (s *castsService) RemovePersonsFromCast(ctx context.Context, movieID int32, persons []models.Person) (err error) {
	if len(persons) == 0 {
		err = models.Error(models.InvalidArgument, "persons mustn't be empty")
		return
	}

	err = s.checkCastExisting(ctx, movieID)
	if err != nil {
		return
	}

	return s.repo.RemovePersonsFromCast(ctx, movieID, persons)
}

func (s *castsService) DeleteCast(ctx context.Context, movieID int32) (err error) {
	err = s.checkCastExisting(ctx, movieID)
	if err != nil {
		return
	}

	return s.repo.DeleteCast(ctx, movieID)
}

func (s *castsService) checkCastExisting(ctx context.Context, id int32) error {
	exist, err := s.repo.IsCastExist(ctx, id)
	if err != nil {
		return err
	}
	if !exist {
		return models.Errorf(models.NotFound, "cast with %d id not found", id)
	}

	return nil
}

func (s *castsService) GetProfessions(ctx context.Context) ([]models.Profession, error) {
	return s.repo.GetAllProfessions(ctx)
}

func (s *castsService) CreateProfession(ctx context.Context, name string) (id int32, err error) {
	exist, id, err := s.repo.IsProfessionWithNameExists(ctx, name)
	if err != nil {
		return
	}
	if exist {
		err = models.Errorf(models.Conflict, "profession already exists, check profession with id %d", id)
		return
	}

	return s.repo.CreateProfession(ctx, name)
}

func (s *castsService) DeleteProfession(ctx context.Context, id int32) error {
	exist, err := s.repo.IsProfessionExists(ctx, id)
	if err != nil {
		return err
	}
	if !exist {
		return models.Errorf(models.NotFound, "profession with id %d not found", id)
	}

	return s.repo.DeleteProfession(ctx, id)
}

func (s *castsService) UpdateProfession(ctx context.Context, professionID int32, name string) error {
	exist, id, err := s.repo.IsProfessionWithNameExists(ctx, name)
	if err != nil {
		return err
	}
	if exist {
		return models.Errorf(models.Conflict, "profession already exists, check profession with id %d", id)
	}

	exist, err = s.repo.IsProfessionExists(ctx, professionID)
	if err != nil {
		return err
	}
	if !exist {
		return models.Errorf(models.NotFound, "profession with id %d not found", professionID)
	}

	return s.repo.UpdateProfession(ctx, professionID, name)
}
