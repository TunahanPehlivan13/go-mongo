package mongo

import (
	"context"
	"github.com/TunahanPehlivan13/go-mongo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type Record struct {
	ID        primitive.ObjectID `bson:"_id"`
	Counts    []int              `bson:"counts"`
	CreatedAt time.Time          `bson:"createdAt"`
	Key       string             `bson:"key"`
	Value     string             `bson:"value"`
}

type RecordRepository struct {
	db *mongo.Collection
}

func NewRecordRepository(db *mongo.Database, collection string) *RecordRepository {
	return &RecordRepository{
		db: db.Collection(collection),
	}
}

func (recordRepo RecordRepository) GetRecords(ctx context.Context, startDate time.Time, endDate time.Time, minCount int, maxCount int) ([]*models.Record, error) {
	query := bson.M{
		"createdAt": bson.M{
			"$gt": startDate,
			"$lt": endDate,
		},
	}
	cur, err := recordRepo.db.Find(ctx, query)
	if err != nil {
		log.Printf("Error occurred while finding records with msg -> (%s)", err)
		return nil, err
	}
	defer cur.Close(ctx)

	return mapToRecords(ctx, cur, minCount, maxCount)
}

func mapToRecords(ctx context.Context, cur *mongo.Cursor, minCount int, maxCount int) ([]*models.Record, error) {
	out := make([]*Record, 0)

	for cur.Next(ctx) {
		record := new(Record)
		err := cur.Decode(record)
		if err != nil {
			return nil, err
		}

		out = append(out, record)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return filterAndConvert(out, minCount, maxCount), nil
}

func filterAndConvert(records []*Record, minCount int, maxCount int) []*models.Record {
	out := make([]*models.Record, 0)

	for _, r := range records {
		sumOfCounts := sum(r.Counts)
		if sumOfCounts < maxCount && sumOfCounts > minCount {
			out = append(out, toRecord(r))
		}
	}

	return out
}

func toRecord(r *Record) *models.Record {
	return &models.Record{
		Key:        r.Key,
		CreatedAt:  r.CreatedAt,
		TotalCount: sum(r.Counts),
	}
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
