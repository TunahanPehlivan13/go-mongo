package memdb

import (
	itemErr "github.com/TunahanPehlivan13/go-mongo/item"
	"github.com/TunahanPehlivan13/go-mongo/models"
	"github.com/nedscode/memdb"
)

type ItemRepository struct {
	db *memdb.Store
}

func NewItemRepository(db *memdb.Store) *ItemRepository {
	return &ItemRepository{
		db: db,
	}
}

func (itemRepo ItemRepository) Get(key string) (*models.Item, error) {
	item, ok := itemRepo.db.In("key").One(key).(*models.Item)

	if !ok {
		return nil, itemErr.ErrItemNotFound
	}
	return item, nil
}

func (itemRepo ItemRepository) Persist(key string, value string) error {
	_, err := itemRepo.db.Put(&models.Item{
		Key:   key,
		Value: value,
	})

	return err
}
