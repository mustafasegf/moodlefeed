package line

import (
	"github.com/mustafasegf/scelefeed/entity"
	"gorm.io/gorm"
)

type Repo struct {
  DB *gorm.DB
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{
		DB: db,
	}
}

func (repo *Repo) GetUserFromLineId(model *entity.UsersModel) (err error) {
	fields := []string{
		"scele_id",
		"token",
	}

	query := repo.DB.Table("users").
		Select(fields).
		Find(model)

	err = query.Error
	return
}
