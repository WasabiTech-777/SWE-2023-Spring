package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	PassWord string `json:"PassWord"`
	Id       int    `json:"Id"`
	Articles int    `json:"Articles"`
	CharHit  int    `json:"CharHit"`
	CharMiss int    `json:"CharMiss"`
}

//}
//    "UserName":"jon",
//    "PassWord":"pswd",
//    "Id":"1",
//    "Articles":"1",
//    "CharHit":1,
//    "CharMiss":1
//}

type Friends struct {
	gorm.Model
	UserName1 string `json:"UserName1"`
	UserName2 string `json:"UserName2"`
}

type Articles struct {
	gorm.Model
	Id     int    `json:"Id"`
	Url    string `json:"Url"`
	Length string `json:"Length"`
}

type TypingSessions struct {
	gorm.Model
	ArticleID int `json:"ArticleID"`
	UserID    int `json:"UserID"`
	CharHit   int `json:"CharHit"`
	CharMiss  int `json:"CharMiss"`
	Time      int `json:"Time"`
}
