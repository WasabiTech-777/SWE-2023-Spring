package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `json:"ID" gorm:"primary_key"`
	Name     string `json:"uname" gorm:"not null;unique"`
	Pass     string `json:"pass" gorm:"not null;"`
	Articles uint   `json:"articles"`
	CharHit  uint   `json:"charhit"`
	CharMiss uint   `json:"charmiss"`
}
