package scele

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/mustafasegf/scelefeed/entity"
)

type Repo struct {
	db *pgxpool.Pool
}

func NewRepo(db *pgxpool.Pool) *Repo {
	return &Repo{
		db: db,
	}
}

func (repo *Repo) CreateUser(data entity.UsersModel) (err error) {
	sql, args, err := sq.Insert("users").Columns("token", "scele_id").
		Values(data.Token, data.SceleID).ToSql()
	if err != nil {
		err = fmt.Errorf("generate query: %w", err)
		return
	}

	ctx := context.Background()
	tx, err := repo.db.Begin(ctx)
	if err != nil {
		err = fmt.Errorf("begin tx: %w", err)
		return
	}

	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, sql, args)
	if err != nil {
		err = fmt.Errorf("executing query: %w", err)
		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		err = fmt.Errorf("commit tx: %w", err)
	}

	return
}

func (repo *Repo) GetLineIDFromCourse(courseId int) (model []entity.UsersModel, err error) {
	sql, args, err := sq.Select("line_id").From("users").
		Where("scele_id in (select user_id from user_subscribe where course_id=?)", courseId).ToSql()
	if err != nil {
		err = fmt.Errorf("generate query: %w", err)
		return
	}
	ctx := context.Background()
	err = pgxscan.Select(ctx, repo.db, &model, sql, args)
	return
}

func (repo *Repo) CreateCourse(data entity.CoursesModel) (err error) {
	sql, args, err := sq.Insert("courses").Columns("course_id", "short_name", "long_name", "user_token").
		Values(data.CourseID, data.ShortName, data.LongName, data.UserToken).ToSql()
	if err != nil {
		err = fmt.Errorf("generate query: %w", err)
		return
	}

	ctx := context.Background()
	tx, err := repo.db.Begin(ctx)
	if err != nil {
		err = fmt.Errorf("begin tx: %w", err)
		return
	}

	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, sql, args)
	if err != nil {
		err = fmt.Errorf("executing query: %w", err)
		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		err = fmt.Errorf("commit tx: %w", err)
	}

	return
}

func (repo *Repo) UpdateCourseResource(data entity.CoursesModel) (err error) {
	sql, args, err := sq.Update("courses").Where("course_id", data.CourseID).Set("resource", data.Resource).ToSql()
	if err != nil {
		err = fmt.Errorf("generate query: %w", err)
		return
	}

	ctx := context.Background()
	tx, err := repo.db.Begin(ctx)
	if err != nil {
		err = fmt.Errorf("begin tx: %w", err)
		return
	}

	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, sql, args)
	if err != nil {
		err = fmt.Errorf("executing query: %w", err)
		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		err = fmt.Errorf("commit tx: %w", err)
	}

	return
}

func (repo *Repo) GetCourseByID(courseID int) (model entity.CoursesModel, err error) {
	sql, args, err := sq.Select("*").From("courses").Where("course_id = ?", courseID).ToSql()
	if err != nil {
		err = fmt.Errorf("generate query: %w", err)
		return
	}
	ctx := context.Background()
	err = pgxscan.Select(ctx, repo.db, &model, sql, args)
	return
}

func (repo *Repo) GetCoursesNameByToken(token string) (model []entity.Course, err error) {
	sql, args, err := sq.Select("short_name").From("courses").Where("user_token = ?", token).ToSql()
	if err != nil {
		err = fmt.Errorf("generate query: %w", err)
		return
	}
	ctx := context.Background()
	err = pgxscan.Select(ctx, repo.db, &model, sql, args)
	return
}

func (repo *Repo) GetAllCourse(model *[]entity.CoursesModel) (err error) {
	sql, args, err := sq.Select("course_id, long_name, user_token, resource").
		From("courses").OrderBy("courde_id desc").ToSql()
	if err != nil {
		err = fmt.Errorf("generate query: %w", err)
		return
	}
	ctx := context.Background()
	err = pgxscan.Select(ctx, repo.db, &model, sql, args)
	return
}

func (repo *Repo) DeleteUserSubscribe(sceleID int) (err error) {
	sql, args, err := sq.Delete("user_subscribe").Where("scele_id = ?", sceleID).ToSql()
	if err != nil {
		err = fmt.Errorf("generate query: %w", err)
		return
	}

	ctx := context.Background()
	tx, err := repo.db.Begin(ctx)
	if err != nil {
		err = fmt.Errorf("begin tx: %w", err)
		return
	}

	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, sql, args)
	if err != nil {
		err = fmt.Errorf("executing query: %w", err)
		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		err = fmt.Errorf("commit tx: %w", err)
	}

	return
}

func (repo *Repo) CreateTokenCourse(data entity.TokenCourseModel) (err error) {
	sql, args, err := sq.Insert("token_course").Columns("course_id", "token").
		Values(data.CourseID, data.Token).ToSql()
	if err != nil {
		err = fmt.Errorf("generate query: %w", err)
		return
	}

	ctx := context.Background()
	tx, err := repo.db.Begin(ctx)
	if err != nil {
		err = fmt.Errorf("begin tx: %w", err)
		return
	}

	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, sql, args)
	if err != nil {
		err = fmt.Errorf("executing query: %w", err)
		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		err = fmt.Errorf("commit tx: %w", err)
	}

	return
}

func (repo *Repo) CreateUserSubscribe(data entity.UserSubscribeModel) (err error) {
	sql, args, err := sq.Insert("user_subscribe").Columns("scele_id", "type_id", "course_id").
		Values(data.SceleID, data.TypeID, data.CourseID).ToSql()
	if err != nil {
		err = fmt.Errorf("generate query: %w", err)
		return
	}

	ctx := context.Background()
	tx, err := repo.db.Begin(ctx)
	if err != nil {
		err = fmt.Errorf("begin tx: %w", err)
		return
	}

	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, sql, args)
	if err != nil {
		err = fmt.Errorf("executing query: %w", err)
		return
	}

	err = tx.Commit(ctx)
	if err != nil {
		err = fmt.Errorf("commit tx: %w", err)
	}

	return
}
