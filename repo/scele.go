package repo

import (
	"github.com/mustafasegf/scelefeed/entity"
	"gorm.io/gorm"
)

type Scele struct {
	db *gorm.DB
}

func NewSceleRepo(db *gorm.DB) *Scele {
	return &Scele{
		db: db,
	}
}

func (repo *Scele) CreateUser(model entity.UsersModel) (err error) {
	query := repo.db.Table("users").Begin().
		Create(&model)
	if err = query.Error; err != nil {
		query.Rollback()
		return
	}
	err = query.Commit().Error
	return
}

func (repo *Scele) CreateCourse(model entity.CoursesModel) (err error) {
	query := repo.db.Table("courses").Begin().
		Create(&model)
	if err = query.Error; err != nil {
		query.Rollback()
		return
	}
	err = query.Commit().Error
	return
}

func (repo *Scele) CreateTokenCourse(model entity.TokenCourseModel) (err error) {
	query := repo.db.Table("token_course").Begin().
		Create(&model)
	if err = query.Error; err != nil {
		query.Rollback()
		return
	}
	err = query.Commit().Error
	return
}

func (repo *Scele) CreateUserSubscribe(model entity.UserSubscribeModel) (err error) {
	query := repo.db.Table("user_subscribe").Begin().
		Create(&model)
	if err = query.Error; err != nil {
		query.Rollback()
		return
	}
	err = query.Commit().Error
	return
}

func (repo *Scele) GetCourse(courseID uint, model entity.CoursesModel) (err error) {
	query := repo.db.Table("courses").
		Where("course_id = ?", courseID).
		First(&model)

	err = query.Error
	return
}

func (repo *Scele) GetAllCourse(model *[]entity.CoursesModel) (err error) {
	fields := []string{
		"course_id",
		"long_name",
		"user_token",
		"resource",
	}
	query := repo.db.Table("courses").
		Select(fields).
		Order("course_id desc").
		Find(model)

	err = query.Error
	return
}

func (repo *Scele) GetIdLineFromCourse(courseId uint, model *[]entity.UsersModel) (err error) {
	fields := []string{
		"line_id",
	}
	query := repo.db.Table("users").
		Select(fields).
		Where("scele_id in (select user_id from user_subscribe where course_id=?)", courseId).
		Find(model)

	err = query.Error
	return
}

func (repo *Scele) UpdateCourseResource(courseID uint, model entity.Resource) (err error) {
	query := repo.db.Table("courses").Begin().
		Where("course_id = ?", courseID).Update("resource", &model)
	if err = query.Error; err != nil {
		query.Rollback()
		return
	}
	err = query.Commit().Error
	return
}
