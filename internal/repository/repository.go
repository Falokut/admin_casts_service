package repository

import (
	"context"

	"github.com/Falokut/admin_casts_service/internal/models"
)

type DBConfig struct {
	Host     string `yaml:"host" env:"DB_HOST"`
	Port     string `yaml:"port" env:"DB_PORT"`
	Username string `yaml:"username" env:"DB_USERNAME"`
	Password string `yaml:"password" env:"DB_PASSWORD"`
	DBName   string `yaml:"db_name" env:"DB_NAME"`
	SSLMode  string `yaml:"ssl_mode" env:"DB_SSL_MODE"`
}

type CastsRepository interface {
	GetCast(ctx context.Context, id int32) ([]models.Cast, error)
	GetAllCasts(ctx context.Context, limit, offset int32) ([]models.Cast, error)
	GetCasts(ctx context.Context, ids []int32, limit, offset int32) ([]models.Cast, error)
	SearchCastByLabel(ctx context.Context, label string, limit, offset int32) ([]models.CastLabel, error)
	IsCastExist(ctx context.Context, id int32) (bool, error)
	CreateCast(ctx context.Context, id int32, label string, persons []models.Person) error
	DeleteCast(ctx context.Context, id int32) error

	RemovePersonFromCasts(ctx context.Context, personID int32) (err error)
	UpdateLabelForCast(ctx context.Context, id int32, label string) error
	AddPersonsToTheCast(ctx context.Context, id int32, persons []models.Person) error
	RemovePersonsFromCast(ctx context.Context, id int32, persons []models.Person) error

	IsProfessionWithNameExists(ctx context.Context, name string) (bool, int32, error)
	IsProfessionExists(ctx context.Context, id int32) (bool, error)

	CreateProfession(ctx context.Context, name string) (int32, error)
	UpdateProfession(ctx context.Context, id int32, name string) error
	DeleteProfession(ctx context.Context, id int32) error
	GetAllProfessions(ctx context.Context) ([]models.Profession, error)
}
