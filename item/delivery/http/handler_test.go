package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/TunahanPehlivan13/go-mongo/item"
	"github.com/TunahanPehlivan13/go-mongo/item/usecase"
	"github.com/TunahanPehlivan13/go-mongo/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGet(t *testing.T) {
	r := gin.Default()
	mock := new(usecase.ItemUseCaseMock)

	RegisterHTTPEndpoints(r, mock)

	item := &models.Item{Key: "key1", Value: "value"}

	mock.On("Get", "key1").Return(item, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/in-memory?key=key1", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"key\":\"key1\",\"value\":\"value\"}", w.Body.String())
}

func TestGetWhenNotFound(t *testing.T) {
	r := gin.Default()
	mock := new(usecase.ItemUseCaseMock)

	RegisterHTTPEndpoints(r, mock)

	mock.On("Get", "key1").Return(nil, item.ErrItemNotFound)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/in-memory?key=key1", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

func TestPost(t *testing.T) {
	r := gin.Default()
	mock := new(usecase.ItemUseCaseMock)

	RegisterHTTPEndpoints(r, mock)

	mock.On("Persist", "key1", "value1").Return(nil)

	w := httptest.NewRecorder()

	inp, _ := json.Marshal(&Item{Key: "key1", Value: "value1"})
	req, _ := http.NewRequest("POST", "/in-memory", bytes.NewBuffer(inp))

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"key\":\"key1\",\"value\":\"value1\"}", w.Body.String())
}

func TestPostWhenOccurredError(t *testing.T) {
	r := gin.Default()
	mock := new(usecase.ItemUseCaseMock)

	RegisterHTTPEndpoints(r, mock)

	mock.On("Persist", "key1", "value1").Return(errors.New("could not persisted"))

	w := httptest.NewRecorder()

	inp, _ := json.Marshal(&Item{Key: "key1", Value: "value1"})
	req, _ := http.NewRequest("POST", "/in-memory", bytes.NewBuffer(inp))

	r.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
}
