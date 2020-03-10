package model

import "github.com/jinzhu/gorm"

// Target ...
type Target struct {
	gorm.Model

	Name string `sql:"unique;not null"`

	ClientTargets []ClientTarget
}
