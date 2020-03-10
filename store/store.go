package store

import (
	"github.com/jinzhu/gorm"
	"github.com/yousufdev/go_webapp/model"
)

type store struct {
	DB *gorm.DB
}

func newStore(db *gorm.DB) *store {
	return &store{db}
}

// UserStore is the interface implemented
// by types that can perform CRUD operation
// on model.User type to/from a data source
type UserStore interface {
	Create(user *model.User) error
	FindOne(where ...interface{}) (*model.User, error)
	FindAll() ([]model.User, error)
}
