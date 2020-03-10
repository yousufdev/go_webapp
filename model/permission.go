package model

import "github.com/jinzhu/gorm"

// Permission ...
type Permission struct {
	gorm.Model

	Name string `sql:"unique;not null"`

	RolePermissions []RolePermission
}
