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

func (repo *Scele) GetCourse(courseID uint, model entity.CoursesModel) (err error) {
	query := repo.db.Table("courses").
		Where("course_id = ?", courseID).
		First(&model)

	err = query.Error
	return
}
