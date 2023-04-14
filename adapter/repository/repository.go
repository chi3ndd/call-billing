package repository

import "call-billing/model"

type repository struct {
	record RecordRepository
}

func New(config model.MongoConfig) (Repository, error) {
	handler := &repository{}
	recordRepo, err := newRecordRepository(config)
	if err != nil {
		return nil, err
	}
	handler.record = recordRepo
	// Success
	return handler, nil
}

func (inst *repository) Record() RecordRepository {
	// Success
	return inst.record
}
