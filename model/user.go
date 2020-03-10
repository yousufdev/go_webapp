package model

import "github.com/jinzhu/gorm"

// User ...
type User struct {
	gorm.Model

	Name     string `sql:"type:varchar(256);not null"`
	Email    string `sql:"unique;not null"`
	Password string `sql:"not null"`

	ClientID uint `sql:"default:null"`
	RoleID   uint `sql:"default:null"`
	Role     Role
}
