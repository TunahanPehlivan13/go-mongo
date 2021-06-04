package usecase

import (
	"context"
	"github.com/TunahanPehlivan13/go-mongo/models"
	"github.com/TunahanPehlivan13/go-mongo/record"
	"time"
)

type RecordUseCase struct {
	recordRepo record.Repository
}

func NewRecordUseCase(recordRepo record.Repository) *RecordUseCase {
	return &RecordUseCase{
		recordRepo: recordRepo,
	}
}

func (recordUseCase RecordUseCase) GetRecords(ctx context.Context, startDate time.Time, endDate time.Time, minCount int, maxCount int) ([]*models.Record, error) {
	return recordUseCase.recordRepo.GetRecords(ctx, startDate, endDate, minCount, maxCount)
}
