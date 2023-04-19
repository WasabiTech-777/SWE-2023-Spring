package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	ID     uint   `gorm:"not null;unique_index"`
	Url    string `gorm:"not null;unique_index"`
	Length string
}

//`json:"Id"`
//`json:"Url"`
//`json:"Length"`
