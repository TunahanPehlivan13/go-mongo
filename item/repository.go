package item

import (
	"github.com/TunahanPehlivan13/go-mongo/models"
)

type Repository interface {
	Get(key string) (*models.Item, error)
	Persist(key string, value string) error
}
