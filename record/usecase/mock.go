package usecase

import (
	"context"
	"github.com/TunahanPehlivan13/go-mongo/models"
	"github.com/stretchr/testify/mock"
	"time"
)

type RecordUseCaseMock struct {
	mock.Mock
}

func (mock RecordUseCaseMock) GetRecords(ctx context.Context,
	startDate time.Time, endDate time.Time, minCount int, maxCount int) ([]*models.Record, error) {

	args := mock.Called(startDate, endDate, minCount, maxCount)

	records, ok := args.Get(0).([]*models.Record)
	if !ok {
		return records, args.Error(1)
	}
	return records, nil
}
