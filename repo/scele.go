package repo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Scele struct {
	Db *mongo.Client
}

func NewSceleRepo(db *mongo.Client) *Scele {
	return &Scele{
		Db: db,
	}
}
