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
	professionsTableName = "professions"
)

func (r *castsRepository) GetCast(ctx context.Context, id int32) ([]Cast, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsRepository.GetCast")
	defer span.Finish()
	var err error
	defer span.SetTag("error", err != nil && errors.Is(err, sql.ErrNoRows))

	query := fmt.Sprintf("SELECT %[1]s.movie_id, COALESCE(label,'') AS label,"+
		"person_id, COALESCE(%[3]s.name,'') AS profession_name,COALESCE(%[3]s.id,-1) AS profession_id "+
		"FROM %[1]s LEFT JOIN %[2]s "+
		"ON %[1]s.movie_id=%[2]s.movie_id LEFT JOIN %[3]s ON profession_id=%[3]s.id "+
		"WHERE %[1]s.movie_id=$1",
		castsTableName, castsLabelsTableName, professionsTableName)

	var cast []Cast
	err = r.db.SelectContext(ctx, &cast, query, id)
	if err != nil {
		r.logger.Errorf("%v query: %s args: %v", err.Error(), query, id)
		return []Cast{}, err
	}
	if len(cast) == 0 {
		return []Cast{}, ErrNotFound
	}

	return cast, nil
}

func (r *castsRepository) GetAllCasts(ctx context.Context, limit, offset int32) ([]Cast, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsRepository.GetAllCasts")
	defer span.Finish()
	var err error
	defer span.SetTag("error", err != nil)

	query := fmt.Sprintf("SELECT %[1]s.movie_id, COALESCE(label,'') AS label,"+
		"person_id, COALESCE(%[3]s.name,'') AS profession_name,COALESCE(%[3]s.id,-1) AS profession_id "+
		"FROM %[1]s LEFT JOIN %[2]s "+
		"ON %[1]s.movie_id=%[2]s.movie_id LEFT JOIN %[3]s ON profession_id=%[3]s.id "+
		"LIMIT %[4]d OFFSET %[5]d;",
		castsTableName, castsLabelsTableName, professionsTableName, limit, offset)

	var casts []Cast
	err = r.db.SelectContext(ctx, &casts, query)
	if err != nil {
		r.logger.Errorf("%v query: %s", err.Error(), query)
		return []Cast{}, err
	}
	if len(casts) == 0 {
		return []Cast{}, ErrNotFound
	}

	return casts, nil
}

func (r *castsRepository) GetCasts(ctx context.Context, ids []int32, limit, offset int32) ([]Cast, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsRepository.GetCasts")
	defer span.Finish()
	var err error
	defer span.SetTag("error", err != nil)

	query := fmt.Sprintf("SELECT %[1]s.movie_id, COALESCE(label,'') AS label,"+
		"person_id, COALESCE(%[3]s.name,'') AS profession_name,COALESCE(%[3]s.id,-1) AS profession_id "+
		"FROM %[1]s LEFT JOIN %[2]s "+
		"ON %[1]s.movie_id=%[2]s.movie_id LEFT JOIN %[3]s ON profession_id=%[3]s.id "+
		"WHERE %[1]s.movie_id=ANY($1) "+
		"LIMIT %[4]d OFFSET %[5]d;",
		castsTableName, castsLabelsTableName, professionsTableName, limit, offset)

	var casts []Cast
	err = r.db.SelectContext(ctx, &casts, query, ids)
	if err != nil {
		r.logger.Errorf("%v query: %s args: %v", err.Error(), query, ids)
		return []Cast{}, err
	}
	if len(casts) == 0 {
		return []Cast{}, ErrNotFound
	}

	return casts, nil
}

func (r *castsRepository) SearchCastByLabel(ctx context.Context, label string, limit, offset int32) ([]CastLabel, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsRepository.SearchCastByLabel")
	defer span.Finish()
	var err error
	defer span.SetTag("error", err != nil)

	query := fmt.Sprintf("SELECT %[1]s.movie_id as movie_id, label "+
		"FROM %[1]s JOIN %[2]s ON %[1]s.movie_id=%[2]s.movie_id WHERE label LIKE($1) "+
		"LIMIT %[3]d OFFSET %[4]d;", castsTableName, castsLabelsTableName, limit, offset)

	var casts []CastLabel
	label += "%"
	err = r.db.SelectContext(ctx, &casts, query, label)
	if err != nil {
		r.logger.Errorf("%v query: %s args: %v", err.Error(), query, label)
		return []CastLabel{}, err
	}
	if len(casts) == 0 {
		return []CastLabel{}, ErrNotFound
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

func (r *castsRepository) CreateCast(ctx context.Context, id int32, label string, persons []Person) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsRepository.CreateCast")
	defer span.Finish()
	var err error
	defer span.SetTag("error", err != nil)

	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	values := make([]string, 0, len(persons))
	for _, person := range persons {
		values = append(values, fmt.Sprintf("(%d,%d,%d)", id, person.ID, person.ProfessionID))
	}

	query := fmt.Sprintf("INSERT INTO %s (movie_id,person_id,profession_id) VALUES %s ON CONFLICT DO NOTHING;", castsTableName,
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

func (r *castsRepository) RemovePersonFromCasts(ctx context.Context, personID int32) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsRepository.RemovePersonFromCasts")
	defer span.Finish()
	var err error
	defer span.SetTag("error", err != nil)

	query := fmt.Sprintf("DELETE FROM %s WHERE person_id=$1", castsTableName)
	_, err = r.db.ExecContext(ctx, query, personID)
	if err != nil {
		r.logger.Errorf("%v query: %s args: %v", err.Error(), query, personID)
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

func (r *castsRepository) AddPersonsToTheCast(ctx context.Context, id int32, persons []Person) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsRepository.AddPersonsToTheCast")
	defer span.Finish()
	var err error
	defer span.SetTag("error", err != nil && !errors.Is(err, sql.ErrNoRows))

	values := make([]string, 0, len(persons))
	for _, person := range persons {
		values = append(values, fmt.Sprintf("(%d,%d,%d)", id, person.ID, person.ProfessionID))
	}

	query := fmt.Sprintf("INSERT INTO %s (movie_id,person_id,profession_id) VALUES %s ON CONFLICT DO NOTHING;",
		castsTableName, strings.Join(values, ","))

	_, err = r.db.ExecContext(ctx, query)
	if err != nil {
		r.logger.Errorf("%v query: %s", err.Error(), query)
		return err
	}

	return nil
}

func (r *castsRepository) RemovePersonsFromCast(ctx context.Context, id int32, persons []Person) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsRepository.DeletePersonsFromTheCast")
	defer span.Finish()
	var err error
	defer span.SetTag("error", err != nil)

	statements := make([]string, 0, len(persons))
	args := make([]any, 0, len(persons)*3)
	args = append(args, id)
	for _, person := range persons {
		args = append(args, person.ID)
		statements = append(statements, fmt.Sprintf("(person_id=$%d AND profession_id=$%d)", len(args), len(args)+1))
		args = append(args, person.ProfessionID)
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE movie_id=$1 AND(%s)", castsTableName, strings.Join(statements, " OR "))
	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		r.logger.Errorf("%v query: %s args: %v", err.Error(), query, args)
		return err
	}

	return nil
}

func (r *castsRepository) DeleteProfession(ctx context.Context, id int32) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsRepository.DeleteProfession")
	defer span.Finish()
	var err error
	defer span.SetTag("error", err != nil)

	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 RETURNING id", professionsTableName)
	var delId int32
	err = r.db.GetContext(ctx, &delId, query, id)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrNotFound
	}
	if err != nil {
		r.logger.Errorf("%v query: %s args: %v", err.Error(), query, id)
		return err
	}

	return nil
}

func (r *castsRepository) CreateProfession(ctx context.Context, name string) (int32, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsRepository.CreateProfession")
	defer span.Finish()
	var err error
	defer span.SetTag("error", err != nil)

	query := fmt.Sprintf("INSERT INTO %s (name) VALUES($1) RETURNING id", professionsTableName)
	var createdID int32
	err = r.db.GetContext(ctx, &createdID, query, name)
	if err != nil {
		r.logger.Errorf("%v query: %s args: %v", err.Error(), query, name)
		return 0, err
	}

	return createdID, nil
}

func (r *castsRepository) IsProfessionWithNameExists(ctx context.Context, name string) (bool, int32, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsRepository.IsProfessionWithNameExists")
	defer span.Finish()
	var err error
	defer span.SetTag("error", err != nil)

	query := fmt.Sprintf("SELECT id FROM %s WHERE name=$1", professionsTableName)
	var id int32
	err = r.db.GetContext(ctx, &id, query, name)
	if errors.Is(err, sql.ErrNoRows) {
		return false, 0, nil
	}
	if err != nil {
		r.logger.Errorf("%v query: %s args: %v", err.Error(), query, name)
		return false, 0, err
	}

	return true, id, nil
}

func (r *castsRepository) UpdateProfession(ctx context.Context, id int32, name string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsRepository.UpdateProfession")
	defer span.Finish()
	var err error
	defer span.SetTag("error", err != nil)

	query := fmt.Sprintf("UPDATE %s SET name=$1 WHERE id=$2", professionsTableName)
	_, err = r.db.ExecContext(ctx, query, name, id)
	if err != nil {
		r.logger.Errorf("%v query: %s args: %v", err.Error(), query, id)
		return err
	}

	return nil
}

func (r *castsRepository) GetAllProfessions(ctx context.Context) ([]Profession, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsRepository.IsProfessionWithNameExists")
	defer span.Finish()
	var err error
	defer span.SetTag("error", err != nil)

	query := fmt.Sprintf("SELECT * FROM %s", professionsTableName)
	var professions []Profession
	err = r.db.SelectContext(ctx, &professions, query)
	if err != nil {
		r.logger.Errorf("%v query: %s", err.Error(), query)
		return []Profession{}, err
	}

	return professions, nil
}

func (r *castsRepository) IsProfessionExists(ctx context.Context, id int32) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "castsRepository.IsProfessionExists")
	defer span.Finish()
	var err error
	defer span.SetTag("error", err != nil)

	query := fmt.Sprintf("SELECT id FROM %s WHERE id=$1", professionsTableName)
	var findedid int32
	err = r.db.GetContext(ctx, &findedid, query, id)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		r.logger.Errorf("%v query: %s args: %v", err.Error(), query, id)
		return false, err
	}

	return true, nil
}
