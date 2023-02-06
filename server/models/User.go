package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"uname" gorm:"not null;primaryKey"`
	Pass     string `json:"pass" gorm:"not null;"`
	Articles uint   `json:"articles"`
	CharHit  uint   `json:"charhit"`
	CharMiss uint   `json:"charmiss"`
}

//	PassWord string `json:"PassWord"`
//	Id       int    `json:"Id"`
//	Articles int    `json:"Articles"`
//	CharHit  int    `json:"CharHit"`
//	CharMiss int    `json:"CharMiss"`

//}
//    "UserName":"jon",
//    "PassWord":"pswd",
//    "Id":"1",
//    "Articles":"1",
//    "CharHit":1,
//    "CharMiss":1
//}
