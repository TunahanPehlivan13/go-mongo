package usecase

import (
	"github.com/TunahanPehlivan13/go-mongo/models"
	"github.com/stretchr/testify/mock"
)

type ItemUseCaseMock struct {
	mock.Mock
}

func (itemUseCaseMock ItemUseCaseMock) Get(key string) (*models.Item, error) {
	args := itemUseCaseMock.Called(key)
	item, ok := args.Get(0).(*models.Item)

	if !ok {
		return nil, args.Error(1)
	}
	return item, nil
}

func (itemUseCaseMock ItemUseCaseMock) Persist(key string, value string) error {
	args := itemUseCaseMock.Called(key, value)

	return args.Error(0)
}
