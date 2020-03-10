package model

// RolePermission ...
type RolePermission struct {
	RoleID       uint `gorm:"primary_key;auto_increment:false"`
	PermissionID uint `gorm:"primary_key;auto_increment:false"`
}
