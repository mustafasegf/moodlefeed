package service

import (
	"github.com/mustafasegf/scelefeed/entity"
	"github.com/mustafasegf/scelefeed/repo"
)

type Scele struct {
	repo *repo.Scele
}

func NewSceleService(repo *repo.Scele) *Scele {
	return &Scele{
		repo: repo,
	}
}

func (svc *Scele) SaveToken(token, lineID string) (err error) {
	user := entity.UserModel{
		Token: token,
		LineId: lineID,
	}
	err = svc.repo.SaveToken(user)
	return
}
