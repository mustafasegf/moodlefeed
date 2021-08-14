package entity

import "gorm.io/gorm"

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Token struct {
	Token string `json:"token"`
}

type UserModel struct {
	gorm.Model
	Token  string
	LineId string
}

type MatkulModel struct {
	gorm.Model
}
type MessageTypeModel struct {
	gorm.Model
	Name string
}

type TokenMatkulModel struct {
	gorm.Model
	MatkulID uint
	UserId   uint
}

type UserSubscribe struct {
	gorm.Model
	UserId   uint
	TypeId   uint
	MatkulId uint
}
