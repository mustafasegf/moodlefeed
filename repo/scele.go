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

func (repo *Scele) insert(model interface{}, table string) (err error) {
	query := repo.db.Table(table).Begin().
		Create(&model)
	if err = query.Error; err != nil {
		query.Rollback()
		return
	}
	err = query.Commit().Error
	return
}

func (repo *Scele) SaveToken(model entity.UserModel) (err error) {
	err = repo.insert(model, "users")
	return
}
