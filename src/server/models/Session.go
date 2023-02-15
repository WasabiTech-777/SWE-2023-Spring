package models

import "gorm.io/gorm"

type Session struct {
	gorm.Model
	SessionID uint `gorm:"not null;unique_index"`
	ArticleID uint `gorm:"not null;"`
	UserID    uint `gorm:"not null;"`
	CharHit   uint
	CharMiss  uint
	Time      uint
}

//`json:"ArticleID"`
//`json:"UserID"`
//`json:"CharHit"`
//`json:"CharMiss"`
//`json:"Time"`
