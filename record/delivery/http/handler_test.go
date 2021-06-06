package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/TunahanPehlivan13/go-mongo/models"
	"github.com/TunahanPehlivan13/go-mongo/record/usecase"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestPost(t *testing.T) {
	r := gin.Default()
	mock := new(usecase.RecordUseCaseMock)

	RegisterHTTPEndpoints(r, mock)

	createdAt, _ := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")

	startDate := time.Time{}
	endDate := time.Time{}

	records := make([]*models.Record, 0)
	records = append(records, &models.Record{
		Key:        "key",
		CreatedAt:  createdAt,
		TotalCount: 123,
	})
	mock.On("GetRecords", startDate, endDate, 1, 1000).Return(records)

	w := httptest.NewRecorder()

	inp, _ := json.Marshal(&filterInput{
		MaxCount:  1000,
		MinCount:  1,
		StartDate: startDate,
		EndDate:   endDate,
	})
	req, _ := http.NewRequest("POST", "/records", bytes.NewBuffer(inp))

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"code\":0,\"msg\":\"success\",\"records\":[{\"key\":\"key\",\"createdAt\":\"2012-11-01T22:08:41Z\",\"totalCount\":123}]}", w.Body.String())
}

func TestPostWhenErrorOccurred(t *testing.T) {
	r := gin.Default()
	mock := new(usecase.RecordUseCaseMock)

	RegisterHTTPEndpoints(r, mock)

	startDate := time.Time{}
	endDate := time.Time{}

	mock.On("GetRecords", startDate, endDate, 1, 1000).Return(nil, errors.New("failed"))

	w := httptest.NewRecorder()

	inp, _ := json.Marshal(&filterInput{
		MaxCount:  1000,
		MinCount:  1,
		StartDate: startDate,
		EndDate:   endDate,
	})
	req, _ := http.NewRequest("POST", "/records", bytes.NewBuffer(inp))

	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"code\":1,\"msg\":\"failed\",\"records\":[]}", w.Body.String())
}

func TestPostWhenNotFound(t *testing.T) {
	r := gin.Default()
	mock := new(usecase.RecordUseCaseMock)

	RegisterHTTPEndpoints(r, mock)

	startDate := time.Time{}
	endDate := time.Time{}

	records := make([]*models.Record, 0)
	mock.On("GetRecords", startDate, endDate, 1, 1000).Return(records)

	w := httptest.NewRecorder()

	inp, _ := json.Marshal(&filterInput{
		MaxCount:  1000,
		MinCount:  1,
		StartDate: startDate,
		EndDate:   endDate,
	})
	req, _ := http.NewRequest("POST", "/records", bytes.NewBuffer(inp))

	r.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	assert.Equal(t, "{\"code\":1,\"msg\":\"record not found\",\"records\":[]}", w.Body.String())
}
