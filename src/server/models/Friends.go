package models

import "gorm.io/gorm"

type Friends struct {
	gorm.Model
	UserName1 string `gorm:"not null;"`
	UserName2 string `gorm:"not null;"`
}

//`json:"UserName1"`
//`json:"UserName2"`
