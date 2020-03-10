package store

import (
	"github.com/jinzhu/gorm"
	"github.com/yousufdev/go_webapp/model"
)

// User ...
type User struct {
	*store
}

// NewUserStore ...
func NewUserStore(db *gorm.DB) *User {
	return &User{newStore(db)}
}

// Create ...
func (userStore *User) Create(user *model.User) error {
	return userStore.DB.Create(user).Error
}

// FindOne ...
func (userStore *User) FindOne(where ...interface{}) (*model.User, error) {
	user := model.User{}
	if err := userStore.DB.Preload("Role").First(&user, where...).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// FindAll ...
func (userStore *User) FindAll() ([]model.User, error) {
	users := []model.User{}
	if err := userStore.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// Update ...
func (userStore *User) Update(user *model.User) error {
	if err := userStore.DB.Save(user).Error; err != nil {
		return err
	}
	return nil
}

// Delete ...
func (userStore *User) Delete(userID uint) error {
	user := &model.User{
		Model: gorm.Model{
			ID: userID,
		},
	}
	if err := userStore.DB.Delete(user).Error; err != nil {
		return err
	}
	return nil
}
