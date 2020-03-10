package model

import "github.com/jinzhu/gorm"

// Role ...
type Role struct {
	gorm.Model

	Name string `sql:"unique;not null"`

	Users           []User
	RolePermissions []RolePermission
}

var (
	RoleAdmin   = Role{Name: "admin"}
	RoleAnalyst = Role{Name: "analyst"}
)
