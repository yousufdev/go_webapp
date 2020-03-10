package model

// ClientTarget ...
type ClientTarget struct {
	ClientID uint `gorm:"primary_key;auto_increment:false"`
	TargetID uint `gorm:"primary_key;auto_increment:false"`
}
