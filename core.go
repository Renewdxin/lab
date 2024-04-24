package main

type User struct {
	ID       string `json:"id" gorm:"column:id"`
	Name     string `json:"name" gorm:"column:name"`
	Account  string `json:"account" gorm:"column:account"`
	Password string `json:"password" gorm:"column:password"`
	Role     string `json:"role" gorm:"column:role"`
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
