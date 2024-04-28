package main

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID       string `json:"id" gorm:"column:id"`
	Name     string `json:"name" gorm:"column:name"`
	Account  string `json:"account" gorm:"column:account"`
	Password string `json:"password" gorm:"column:password"`
	Role     string `json:"role" gorm:"column:role"`
}

type LoginEvent struct {
	gorm.Model `gorm:"embedded"`
	UserID     string    `json:"user_id" gorm:"column:uid; primarykey"`
	LoginTime  time.Time `json:"login_time" gorm:"column:login_time"`
	IP         string    `json:"ip" gorm:"column:ip"`
}

type UserCorePort interface {
	TableName() string
}

type UserCoreAdapter struct {
}

func NewUserCoreAdapter() UserCorePort {
	return &UserCoreAdapter{}
}

func (u UserCoreAdapter) TableName() string {
	return "user"
}
