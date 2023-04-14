package repository

import (
	"context"
	"errors"
	"strings"

	"call-billing/defs"
	"call-billing/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type recordRepository struct {
	con *mongo.Client
}

func newRecordRepository(config model.MongoConfig) (RecordRepository, error) {
	con, err := mongo.NewClient(
		options.Client().ApplyURI(config.Address),
	)
	if err != nil {
		return nil, err
	}
	if err = con.Connect(context.Background()); err != nil {
		return nil, err
	}
	// Success
	return &recordRepository{con: con}, nil
}

func (inst *recordRepository) Name() (string, string) {
	// Success
	return defs.DatabaseMobile, defs.CollectionRecord
}

func (inst *recordRepository) FindAll(query *bson.M, sorts []string, offset int64) (*[]model.Record, error) {
	database, collection := inst.Name()
	results := make([]model.Record, 0)
	opts := options.Find()
	if offset > 0 {
		opts.SetSkip(offset)
	}
	if sorts != nil && len(sorts) > 0 {
		s := bson.D{}
		for _, sort := range sorts {
			if strings.HasPrefix(sort, "-") {
				s = append(s, bson.E{Key: strings.TrimPrefix(sort, "-"), Value: -1})
			} else if strings.HasPrefix(sort, "+") {
				s = append(s, bson.E{Key: strings.TrimPrefix(sort, "+"), Value: 1})
			}
		}
		opts.SetSort(s)
	}
	cur, err := inst.con.Database(database).Collection(collection).Find(context.Background(), query, opts)
	if err != nil {
		return &results, err
	}
	if cur == nil {
		return &results, errors.New("no documents")
	}
	for cur.Next(context.Background()) {
		var doc model.Record
		if err = cur.Decode(&doc); err != nil {
			return &results, err
		}
		results = append(results, doc)
	}
	// Success
	return &results, nil
}

func (inst *recordRepository) Insert(document *model.Record) error {
	database, collection := inst.Name()
	opts := options.InsertOne()
	opts.SetBypassDocumentValidation(true)
	if _, err := inst.con.Database(database).Collection(collection).InsertOne(context.Background(), document, opts); err != nil {
		return err
	}
	// Success
	return nil
}
