package usecase

import (
	"github.com/TunahanPehlivan13/go-mongo/item"
	"github.com/TunahanPehlivan13/go-mongo/models"
)

type ItemUseCase struct {
	itemRepo item.Repository
}

func NewItemUseCase(itemRepo item.Repository) *ItemUseCase {
	return &ItemUseCase{
		itemRepo: itemRepo,
	}
}

func (itemUseCase ItemUseCase) Get(key string) (*models.Item, error) {
	return itemUseCase.itemRepo.Get(key)
}

func (itemUseCase ItemUseCase) Persist(key string, value string) error {
	return itemUseCase.itemRepo.Persist(key, value)
}
