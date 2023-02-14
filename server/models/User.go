package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"uname" gorm:"not null;unique"`
	Pass     string `json:"pass" gorm:"not null;"`
	Articles uint   `json:"articles"`
	CharHit  uint   `json:"charhit"`
	CharMiss uint   `json:"charmiss"`
}
