package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type castsRepository struct {
	db     *sqlx.DB
	logger *logrus.Logger
}

func NewCastsRepository(db *sqlx.DB, logger *logrus.Logger) *castsRepository {
	return &castsRepository{db: db, logger: logger}
}

func NewPostgreDB(cfg DBConfig) (*sqlx.DB, error) {
	conStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode)

	if !cfg.EnablePreparedStatements {
		stdlib.RegisterConnConfig(&pgx.ConnConfig{
			DefaultQueryExecMode: pgx.QueryExecModeSimpleProtocol,
		})
	}

	return sqlx.Connect("pgx", conStr)
}

func (r *castsRepository) Shutdown() error {
	return r.db.Close()
}

const (
	castsTableName       = "casts"
	castsLabelsTableName = "casts_labels"
)

func (r *castsRepository) GetCast(ctx context.Context, id int32) (Cast, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsRepository.GetCast")
	defer span.Finish()
	var err error
	defer span.SetTag("error", err != nil && errors.Is(err, sql.ErrNoRows))

	query := fmt.Sprintf("SELECT %[1]s.movie_id, label, ARRAY_AGG(actor_id) as actors_ids FROM %[1]s JOIN %[2]s "+
		"ON %[1]s.movie_id=%[2]s.movie_id "+
		"WHERE %[1]s.movie_id=$1 "+
		"GROUP BY %[1]s.movie_id, label;", castsTableName, castsLabelsTableName)

	var cast Cast
	err = r.db.GetContext(ctx, &cast, query, id)
	if errors.Is(err, sql.ErrNoRows) {
		return Cast{}, ErrNotFound
	} else if err != nil {
		r.logger.Errorf("%v query: %s args: %v", err.Error(), query, id)
		return Cast{}, err
	}

	cast.Actors = strings.Trim(cast.Actors, "{}")
	return cast, nil
}

func (r *castsRepository) GetAllCasts(ctx context.Context, limit, offset int32) ([]Cast, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsRepository.GetAllCasts")
	defer span.Finish()
	var err error
	defer span.SetTag("error", err != nil)

	query := fmt.Sprintf("SELECT %[1]s.movie_id, label, ARRAY_AGG(actor_id) as actors_ids FROM %[1]s JOIN %[2]s "+
		"ON %[1]s.movie_id=%[2]s.movie_id "+
		"GROUP BY %[1]s.movie_id, label LIMIT %[3]d OFFSET %[4]d;", castsTableName, castsLabelsTableName, limit, offset)

	var casts []Cast
	err = r.db.SelectContext(ctx, &casts, query)
	if err != nil {
		r.logger.Errorf("%v query: %s", err.Error(), query)
		return []Cast{}, err
	}
	if len(casts) == 0 {
		return []Cast{}, ErrNotFound
	}

	for i := 0; i < len(casts); i++ {
		casts[i].Actors = strings.Trim(casts[i].Actors, "{}")

	}

	return casts, nil
}

func (r *castsRepository) GetCasts(ctx context.Context, ids []int32, limit, offset int32) ([]Cast, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsRepository.GetCasts")
	defer span.Finish()
	var err error
	defer span.SetTag("error", err != nil)

	query := fmt.Sprintf("SELECT %[1]s.movie_id, label, ARRAY_AGG(actor_id) as actors_ids FROM %[1]s JOIN %[2]s "+
		"ON %[1]s.movie_id=%[2]s.movie_id "+
		"WHERE %[1]s.movie_id=ANY($1) "+
		"GROUP BY %[1]s.movie_id, label LIMIT %[3]d OFFSET %[4]d;", castsTableName, castsLabelsTableName, limit, offset)

	var casts []Cast
	err = r.db.SelectContext(ctx, &casts, query, ids)
	if err != nil {
		r.logger.Errorf("%v query: %s args: %v", err.Error(), query, ids)
		return []Cast{}, err
	}
	if len(casts) == 0 {
		return []Cast{}, ErrNotFound
	}

	for i := 0; i < len(casts); i++ {
		casts[i].Actors = strings.Trim(casts[i].Actors, "{}")

	}

	return casts, nil
}

func (r *castsRepository) SearchCastByLabel(ctx context.Context, label string, limit, offset int32) ([]Cast, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsRepository.SearchCastByLabel")
	defer span.Finish()
	var err error
	defer span.SetTag("error", err != nil)

	query := fmt.Sprintf("SELECT %[1]s.movie_id, label, ARRAY_AGG(actor_id) as actors_ids FROM %[1]s JOIN %[2]s "+
		"ON %[1]s.movie_id=%[2]s.movie_id "+
		"WHERE label=LIKE($1) "+
		"GROUP BY %[1]s.movie_id, label LIMIT %[3]d OFFSET %[4]d;", castsTableName, castsLabelsTableName, limit, offset)

	var casts []Cast
	label += "%"
	err = r.db.SelectContext(ctx, &casts, query, label)
	if err != nil {
		r.logger.Errorf("%v query: %s args: %v", err.Error(), query, label)
		return []Cast{}, err
	}
	if len(casts) == 0 {
		return []Cast{}, ErrNotFound
	}

	for i := 0; i < len(casts); i++ {
		casts[i].Actors = strings.Trim(casts[i].Actors, "{}")

	}

	return casts, nil
}

func (r *castsRepository) IsCastExist(ctx context.Context, id int32) (bool, int32, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsRepository.IsCastExist")
	defer span.Finish()
	var err error
	defer span.SetTag("error", err != nil && !errors.Is(err, sql.ErrNoRows))

	query := fmt.Sprintf("SELECT movie_id FROM %s WHERE movie_id=$1 LIMIT 1", castsTableName)
	var foundedID int32
	err = r.db.GetContext(ctx, &foundedID, query, id)
	if errors.Is(err, sql.ErrNoRows) {
		return false, 0, nil
	} else if err != nil {
		r.logger.Errorf("%v query: %s args: %v", err.Error(), query, id)
		return false, 0, err
	}

	return true, foundedID, nil
}

func (r *castsRepository) CreateCast(ctx context.Context, id int32, label string, actors []int32) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsRepository.CreateCast")
	defer span.Finish()
	var err error
	defer span.SetTag("error", err != nil)

	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	values := make([]string, 0, len(actors))
	for _, actorID := range actors {
		values = append(values, fmt.Sprintf("(%d,%d)", id, actorID))
	}

	query := fmt.Sprintf("INSERT INTO %s (movie_id,actor_id) VALUES %s ON CONFLICT DO NOTHING;", castsTableName,
		strings.Join(values, ","))
	_, err = tx.ExecContext(ctx, query)
	if err != nil {
		r.logger.Errorf("%v query: %s args: %v", err.Error(), query, id)
		tx.Rollback()
		return err
	}

	query = fmt.Sprintf("INSERT INTO %s (movie_id,label) VALUES($1,$2);", castsLabelsTableName)
	_, err = tx.ExecContext(ctx, query, id, label)
	if err != nil {
		r.logger.Errorf("%v query: %s args: %v", err.Error(), query, id)
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *castsRepository) DeleteCast(ctx context.Context, id int32) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsRepository.CreateCast")
	defer span.Finish()
	var err error
	defer span.SetTag("error", err != nil)

	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE movie_id=$1", castsTableName)
	_, err = tx.ExecContext(ctx, query, id)
	if err != nil {
		r.logger.Errorf("%v query: %s args: %v", err.Error(), query, id)
		tx.Rollback()
		return err
	}

	query = fmt.Sprintf("DELETE FROM %s WHERE movie_id=$1", castsLabelsTableName)
	_, err = tx.ExecContext(ctx, query, id)
	if err != nil {
		r.logger.Errorf("%v query: %s args: %v", err.Error(), query, id)
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *castsRepository) RemoveActorFromCasts(ctx context.Context, actorID int32) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsRepository.RemoveActorFromCasts")
	defer span.Finish()
	var err error
	defer span.SetTag("error", err != nil)

	query := fmt.Sprintf("DELETE FROM %s WHERE actor_id=$1", castsTableName)
	_, err = r.db.ExecContext(ctx, query, actorID)
	if err != nil {
		r.logger.Errorf("%v query: %s args: %v", err.Error(), query, actorID)
		return err
	}

	return nil
}

func (r *castsRepository) UpdateLabelForCast(ctx context.Context, id int32, label string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsRepository.UpdateLabelForCast")
	defer span.Finish()
	var err error
	defer span.SetTag("error", err != nil && !errors.Is(err, sql.ErrNoRows))

	query := fmt.Sprintf("UPDATE %s SET label=$1 WHERE movie_id=$2", castsLabelsTableName)

	_, err = r.db.ExecContext(ctx, query, label, id)
	if err != nil {
		r.logger.Errorf("%v query: %s args: movie_id: %v label: %v", err.Error(), query, id, label)
		return err
	}

	return nil
}

func (r *castsRepository) AddActorsToTheCast(ctx context.Context, id int32, actorsIDs []int32) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsRepository.AddActorsToTheCast")
	defer span.Finish()
	var err error
	defer span.SetTag("error", err != nil && !errors.Is(err, sql.ErrNoRows))

	values := make([]string, 0, len(actorsIDs))
	for _, actorID := range actorsIDs {
		values = append(values, fmt.Sprintf("(%d,%d)", id, actorID))
	}
	query := fmt.Sprintf("INSERT INTO %s (movie_id,actor_id) VALUES %s ON CONFLICT DO NOTHING;",
		castsTableName, strings.Join(values, ","))

	_, err = r.db.ExecContext(ctx, query)
	if err != nil {
		r.logger.Errorf("%v query: %s", err.Error(), query)
		return err
	}

	return nil
}

func (r *castsRepository) RemoveActorsFromCast(ctx context.Context, id int32, actorsIDs []int32) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsRepository.DeleteActorsFromTheCast")
	defer span.Finish()
	var err error
	defer span.SetTag("error", err != nil)

	query := fmt.Sprintf("DELETE FROM %s WHERE movie_id=$1 AND actor_id=ANY($2)", castsTableName)
	_, err = r.db.ExecContext(ctx, query, id, actorsIDs)
	if err != nil {
		r.logger.Errorf("%v query: %s args: movie_id: %v actors_ids: %v", err.Error(), query, id, actorsIDs)
		return err
	}

	return nil
}
