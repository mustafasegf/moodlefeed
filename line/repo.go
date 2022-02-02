package line

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/mustafasegf/scelefeed/entity"
)

type Repo struct {
	DB *pgxpool.Pool
}

func NewRepo(db *pgxpool.Pool) *Repo {
	return &Repo{
		DB: db,
	}
}

func (repo *Repo) GetUserFromLineId(model *entity.UsersModel) (err error) {
	// fields := []string{
	// 	"scele_id",
	// 	"token",
	// }

	// query := repo.DB.Table("users").
	// 	Select(fields).
	// 	Find(model)

	// err = query.Error
	return
}
