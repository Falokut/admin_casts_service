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
	ID     int32  `db:"movie_id"`
	Label  string `db:"label"`
	Actors string `db:"actors_ids"`
}

type CastsRepository interface {
	GetCast(ctx context.Context, id int32) (Cast, error)
	GetAllCasts(ctx context.Context, limit, offset int32) ([]Cast, error)
	GetCasts(ctx context.Context, ids []int32, limit, offset int32) ([]Cast, error)
	SearchCastByLabel(ctx context.Context, label string, limit, offset int32) ([]Cast, error)
	IsCastExist(ctx context.Context, id int32) (bool, int32, error)
	CreateCast(ctx context.Context, id int32, label string, actors []int32) error
	DeleteCast(ctx context.Context, id int32) error
	RemoveActorFromCasts(ctx context.Context, actorID int32) (err error)
	UpdateLabelForCast(ctx context.Context, id int32, label string) error
	AddActorsToTheCast(ctx context.Context, id int32, actorsIDs []int32) error
	RemoveActorsFromCast(ctx context.Context, id int32, actorsIDs []int32) error
}
