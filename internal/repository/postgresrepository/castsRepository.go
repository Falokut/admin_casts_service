package postgresrepository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/Falokut/admin_casts_service/internal/models"
	"github.com/Falokut/admin_casts_service/internal/repository"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type castsRepository struct {
	db     *sqlx.DB
	logger *logrus.Logger
}

func NewCastsRepository(db *sqlx.DB, logger *logrus.Logger) *castsRepository {
	return &castsRepository{db: db, logger: logger}
}

func NewPostgreDB(cfg *repository.DBConfig) (*sqlx.DB, error) {
	conStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode)
	db, err := sqlx.Connect("pgx", conStr)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func (r *castsRepository) Shutdown() {
	err := r.db.Close()
	if err != nil {
		r.logger.Errorf("error while shutting down database %v", err)
	}
}

const (
	castsTableName       = "casts"
	castsLabelsTableName = "casts_labels"
	professionsTableName = "professions"
)

func (r *castsRepository) GetCast(ctx context.Context, id int32) (cast []models.Cast, err error) {
	defer r.handleError(ctx, &err, "GetCast")

	query := fmt.Sprintf(`SELECT 
		 %[1]s.movie_id, COALESCE(label,'') AS label,
		 person_id, COALESCE(%[3]s.name,'') AS profession_name,
		 COALESCE(%[3]s.id,-1) AS profession_id 
		 FROM %[1]s LEFT JOIN %[2]s ON %[1]s.movie_id=%[2]s.movie_id LEFT JOIN %[3]s ON profession_id=%[3]s.id 
		 WHERE %[1]s.movie_id=$1`,
		castsTableName, castsLabelsTableName, professionsTableName)

	err = r.db.SelectContext(ctx, &cast, query, id)
	return
}

func (r *castsRepository) GetAllCasts(ctx context.Context, limit, offset int32) (casts []models.Cast, err error) {
	defer r.handleError(ctx, &err, "GetAllCasts")

	query := fmt.Sprintf("SELECT %[1]s.movie_id, COALESCE(label,'') AS label,"+
		"person_id, COALESCE(%[3]s.name,'') AS profession_name,COALESCE(%[3]s.id,-1) AS profession_id "+
		"FROM %[1]s LEFT JOIN %[2]s "+
		"ON %[1]s.movie_id=%[2]s.movie_id LEFT JOIN %[3]s ON profession_id=%[3]s.id "+
		"LIMIT %[4]d OFFSET %[5]d;",
		castsTableName, castsLabelsTableName, professionsTableName, limit, offset)

	err = r.db.SelectContext(ctx, &casts, query)
	return
}

func (r *castsRepository) GetCasts(ctx context.Context, ids []int32, limit, offset int32) (casts []models.Cast, err error) {
	defer r.handleError(ctx, &err, "GetCasts")

	query := fmt.Sprintf("SELECT %[1]s.movie_id, COALESCE(label,'') AS label,"+
		"person_id, COALESCE(%[3]s.name,'') AS profession_name,COALESCE(%[3]s.id,-1) AS profession_id "+
		"FROM %[1]s LEFT JOIN %[2]s "+
		"ON %[1]s.movie_id=%[2]s.movie_id LEFT JOIN %[3]s ON profession_id=%[3]s.id "+
		"WHERE %[1]s.movie_id=ANY($1) "+
		"LIMIT %[4]d OFFSET %[5]d;",
		castsTableName, castsLabelsTableName, professionsTableName, limit, offset)

	err = r.db.SelectContext(ctx, &casts, query, ids)
	return
}

func (r *castsRepository) SearchCastByLabel(ctx context.Context, label string, limit, offset int32) (labels []models.CastLabel, err error) {
	defer r.handleError(ctx, &err, "SearchCastByLabel")

	query := fmt.Sprintf("SELECT DISTINCT %[1]s.movie_id AS movie_id, label "+
		"FROM %[1]s JOIN %[2]s ON %[1]s.movie_id=%[2]s.movie_id WHERE label LIKE($1) "+
		"LIMIT %[3]d OFFSET %[4]d;", castsTableName, castsLabelsTableName, limit, offset)

	label += "%"
	err = r.db.SelectContext(ctx, &labels, query, label)
	return
}

func (r *castsRepository) IsCastExist(ctx context.Context, id int32) (exist bool, err error) {
	defer r.handleError(ctx, &err, "IsCastExist")

	query := fmt.Sprintf("SELECT movie_id FROM %s WHERE movie_id=$1 LIMIT 1", castsTableName)
	var foundedID int32
	err = r.db.GetContext(ctx, &foundedID, query, id)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}

	return
}

func (r *castsRepository) CreateCast(ctx context.Context, id int32, label string, persons []models.Person) (err error) {
	defer r.handleError(ctx, &err, "CreateCast")

	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return
	}

	values := make([]string, 0, len(persons))
	for _, person := range persons {
		values = append(values, fmt.Sprintf("(%d,%d,%d)", id, person.ID, person.ProfessionID))
	}

	query := fmt.Sprintf("INSERT INTO %s (movie_id,person_id,profession_id) VALUES %s ON CONFLICT DO NOTHING;", castsTableName,
		strings.Join(values, ","))
	_, err = tx.ExecContext(ctx, query)
	if err != nil {
		err = tx.Rollback()
		return
	}

	query = fmt.Sprintf("INSERT INTO %s (movie_id,label) VALUES($1,$2);", castsLabelsTableName)
	_, err = tx.ExecContext(ctx, query, id, label)
	if err != nil {
		err = tx.Rollback()
		return
	}

	err = tx.Commit()
	return
}

func (r *castsRepository) DeleteCast(ctx context.Context, id int32) (err error) {
	defer r.handleError(ctx, &err, "DeleteCast")

	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE movie_id=$1", castsTableName)
	_, err = tx.ExecContext(ctx, query, id)
	if err != nil {
		err = tx.Rollback()
		return
	}

	query = fmt.Sprintf("DELETE FROM %s WHERE movie_id=$1", castsLabelsTableName)
	_, err = tx.ExecContext(ctx, query, id)
	if err != nil {
		err = tx.Rollback()
		return
	}

	err = tx.Commit()
	return
}

func (r *castsRepository) RemovePersonFromCasts(ctx context.Context, personID int32) (err error) {
	defer r.handleError(ctx, &err, "RemovePersonFromCasts")

	query := fmt.Sprintf("DELETE FROM %s WHERE person_id=$1", castsTableName)
	_, err = r.db.ExecContext(ctx, query, personID)
	return
}

func (r *castsRepository) UpdateLabelForCast(ctx context.Context, id int32, label string) (err error) {
	defer r.handleError(ctx, &err, "UpdateLabelForCast")

	query := fmt.Sprintf("UPDATE %s SET label=$1 WHERE movie_id=$2", castsLabelsTableName)

	_, err = r.db.ExecContext(ctx, query, label, id)
	return
}

func (r *castsRepository) AddPersonsToTheCast(ctx context.Context, id int32, persons []models.Person) (err error) {
	defer r.handleError(ctx, &err, "AddPersonsToTheCast")

	values := make([]string, 0, len(persons))
	for _, person := range persons {
		values = append(values, fmt.Sprintf("(%d,%d,%d)", id, person.ID, person.ProfessionID))
	}

	query := fmt.Sprintf("INSERT INTO %s (movie_id,person_id,profession_id) VALUES %s ON CONFLICT DO NOTHING;",
		castsTableName, strings.Join(values, ","))

	_, err = r.db.ExecContext(ctx, query)
	return
}

func (r *castsRepository) RemovePersonsFromCast(ctx context.Context, id int32, persons []models.Person) (err error) {
	defer r.handleError(ctx, &err, "RemovePersonsFromCast")

	statements := make([]string, 0, len(persons))
	args := make([]any, 0, len(persons)*2+1)
	args = append(args, id)
	for _, person := range persons {
		statements = append(statements, fmt.Sprintf("(person_id=$%d AND profession_id=$%d)", len(args)+1, len(args)+2))
		args = append(args, person.ID, person.ProfessionID)
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE movie_id=$1 AND(%s)", castsTableName, strings.Join(statements, " OR "))
	_, err = r.db.ExecContext(ctx, query, args...)
	return
}

func (r *castsRepository) DeleteProfession(ctx context.Context, id int32) (err error) {
	defer r.handleError(ctx, &err, "DeleteProfession")

	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 RETURNING id", professionsTableName)
	var deletedID int32
	err = r.db.GetContext(ctx, &deletedID, query, id)
	return
}

func (r *castsRepository) CreateProfession(ctx context.Context, name string) (id int32, err error) {
	defer r.handleError(ctx, &err, "CreateProfession")

	query := fmt.Sprintf("INSERT INTO %s (name) VALUES($1) RETURNING id", professionsTableName)
	err = r.db.GetContext(ctx, &id, query, name)
	return
}

func (r *castsRepository) IsProfessionWithNameExists(ctx context.Context, name string) (exist bool, id int32, err error) {
	query := fmt.Sprintf("SELECT id FROM %s WHERE name=$1", professionsTableName)
	err = r.db.GetContext(ctx, &id, query, name)
	if err != nil {
		r.handleError(ctx, &err, "IsProfessionWithNameExists")
		if models.Code(err) == models.NotFound {
			return false, 0, nil
		}
		return
	}

	return true, id, nil
}

func (r *castsRepository) UpdateProfession(ctx context.Context, id int32, name string) (err error) {
	defer r.handleError(ctx, &err, "UpdateProfession")

	query := fmt.Sprintf("UPDATE %s SET name=$1 WHERE id=$2", professionsTableName)
	_, err = r.db.ExecContext(ctx, query, name, id)
	return
}

func (r *castsRepository) GetAllProfessions(ctx context.Context) (professions []models.Profession, err error) {
	defer r.handleError(ctx, &err, "GetAllProfessions")

	query := fmt.Sprintf("SELECT * FROM %s", professionsTableName)
	err = r.db.SelectContext(ctx, &professions, query)
	return
}

func (r *castsRepository) IsProfessionExists(ctx context.Context, id int32) (exist bool, err error) {
	query := fmt.Sprintf("SELECT id FROM %s WHERE id=$1", professionsTableName)
	var findedid int32
	err = r.db.GetContext(ctx, &findedid, query, id)
	if err != nil {
		r.handleError(ctx, &err, "IsProfessionExists")
		if models.Code(err) == models.NotFound {
			return false, nil
		}
		return
	}

	return true, nil
}

func (r *castsRepository) handleError(ctx context.Context, err *error, functionName string) {
	if ctx.Err() != nil {
		var code models.ErrorCode
		switch {
		case errors.Is(ctx.Err(), context.Canceled):
			code = models.Canceled
		case errors.Is(ctx.Err(), context.DeadlineExceeded):
			code = models.DeadlineExceeded
		}
		*err = models.Error(code, ctx.Err().Error())
		r.logError(*err, functionName)
		return
	}

	if err == nil || *err == nil {
		return
	}

	r.logError(*err, functionName)
	var repoErr = &models.ServiceError{}
	if !errors.As(*err, &repoErr) {
		var code models.ErrorCode
		switch {
		case errors.Is(*err, sql.ErrNoRows):
			code = models.NotFound
			*err = models.Error(code, "cast not found")
		case *err != nil:
			code = models.Internal
			*err = models.Error(code, "repository internal error")
		}
	}
}

func (r *castsRepository) logError(err error, functionName string) {
	if err == nil {
		return
	}

	var repoErr = &models.ServiceError{}
	if errors.As(err, &repoErr) {
		r.logger.WithFields(
			logrus.Fields{
				"error.function.name": functionName,
				"error.msg":           repoErr.Msg,
				"error.code":          repoErr.Code,
			},
		).Error("casts repository error occurred")
	} else {
		r.logger.WithFields(
			logrus.Fields{
				"error.function.name": functionName,
				"error.msg":           err.Error(),
			},
		).Error("casts repository error occurred")
	}
}
