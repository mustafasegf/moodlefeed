package service

import (
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
