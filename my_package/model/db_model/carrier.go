package db_model

import (
	"time"

	"github.com/jinzhu/gorm"
)

/*
author : tiger
2017/5/27
*/

//User 用户表
type User struct {
	gorm.Model
	UserInfo
	Account  string `gorm:"type:varchar(20);unique_index"`
	Password string
	Num      int `gorm:"AUTO_INCREMENT"`
}

//UserInfo 用户信息表
type UserInfo struct {
	UserID   int `gorm:"AUTO_INCREMENT"`
	Age      int
	BirthDay time.Time
	RealName string `gorm:"size:255"`
	Phone    string `gorm:"type:varchar(20);unique_index"`
	NickName string `gorm:"type:varchar(20)"`
	Email    string `gorm:"type:varchar(100);unique_index"`
}
