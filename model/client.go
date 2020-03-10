package model

import "github.com/jinzhu/gorm"

// Client ...
type Client struct {
	gorm.Model

	Name string `sql:"type:varchar(256);not null"`

	Users         []User
	ClientTargets []ClientTarget
}
