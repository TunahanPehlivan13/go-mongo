package record

import (
	"context"
	"github.com/TunahanPehlivan13/go-mongo/models"
	"time"
)

type UseCase interface {
	GetRecords(ctx context.Context, startDate time.Time, endDate time.Time, minCount int, maxCount int) ([]*models.Record, error)
}
