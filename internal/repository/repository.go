package repository

import (
	"context"
	"errors"
)

type DBConfig struct {
	Host                     string `yaml:"host" env:"DB_HOST"`
	Port                     string `yaml:"port" env:"DB_PORT"`
	Username                 string `yaml:"username" env:"DB_USERNAME"`
	Password                 string `yaml:"password" env:"DB_PASSWORD"`
	DBName                   string `yaml:"db_name" env:"DB_NAME"`
	SSLMode                  string `yaml:"ssl_mode" env:"DB_SSL_MODE"`
	EnablePreparedStatements bool   `yaml:"enable_prepared_statements" env:"DB_ENABLE_PREPARED_STATEMENTS"`
}

var ErrNotFound = errors.New("entity not found")
var ErrInvalidArgument = errors.New("invalid input data")

type Cast struct {
	ID             int32  `db:"movie_id"`
	Label          string `db:"label"`
	PersonID        int32  `json:"person_id" db:"person_id"`
	ProfessionID   int32  `json:"profession_id" db:"profession_id"`
	ProfessionName string `json:"profession_name" db:"profession_name"`
}

type Person struct {
	ID           int32 `json:"id" db:"person_id"`
	ProfessionID int32 `json:"profession_id" db:"profession_id"`
}

type Profession struct {
	ID   int32  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type CastLabel struct {
	ID    int32  `db:"movie_id"`
	Label string `db:"label"`
}

type CastsRepository interface {
	GetCast(ctx context.Context, id int32) ([]Cast, error)
	GetAllCasts(ctx context.Context, limit, offset int32) ([]Cast, error)
	GetCasts(ctx context.Context, ids []int32, limit, offset int32) ([]Cast, error)
	SearchCastByLabel(ctx context.Context, label string, limit, offset int32) ([]CastLabel, error)
	IsCastExist(ctx context.Context, id int32) (bool, int32, error)
	CreateCast(ctx context.Context, id int32, label string, persons []Person) error
	DeleteCast(ctx context.Context, id int32) error
	RemovePersonFromCasts(ctx context.Context, personID int32) (err error)
	UpdateLabelForCast(ctx context.Context, id int32, label string) error
	AddPersonsToTheCast(ctx context.Context, id int32, persons []Person) error
	RemovePersonsFromCast(ctx context.Context, id int32, persons []Person) error

	IsProfessionWithNameExists(ctx context.Context, name string) (bool, int32, error)
	IsProfessionExists(ctx context.Context, id int32) (bool, error)

	CreateProfession(ctx context.Context, name string) (int32, error)
	UpdateProfession(ctx context.Context, id int32, name string) error
	DeleteProfession(ctx context.Context, id int32) error
	GetAllProfessions(ctx context.Context) ([]Profession, error)
}
