package repository

import (
	"call-billing/model"

	"go.mongodb.org/mongo-driver/bson"
)

type (
	repo interface {
		Name() (string, string)
	}

	Repository interface {
		Record() RecordRepository
	}

	RecordRepository interface {
		repo
		FindAll(query *bson.M, sorts []string, offset int64) (*[]model.Record, error)
		Insert(document *model.Record) error
	}
)
