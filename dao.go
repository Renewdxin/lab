package main

import (
	"gorm.io/gorm"
)

type UserDaoPort interface {
	Create(user User) error
	FindByID(id int) (User, error)
	Login(password, id string) error
	Update(user User) error
	Delete(id int) error
}

type UserDaoAdapter struct {
	db   *gorm.DB
	core UserCorePort
}

func NewUserDaoAdapter(db *gorm.DB, core UserCorePort) *UserDaoAdapter {
	return &UserDaoAdapter{
		db:   db,
		core: core,
	}
}

func (u UserDaoAdapter) Create(user User) error {
	return u.db.Table(u.core.TableName()).Create(&user).Error
}

func (u UserDaoAdapter) FindByID(id int) (User, error) {
	var user User
	err := u.db.Table(u.core.TableName()).Where("id = ?", id).First(&user).Error
	return user, err
}

func (u UserDaoAdapter) Login(password, id string) error {
	var user User
	return u.db.Table(u.core.TableName()).Where("account = ? AND password = ?", id, password).First(&user).Error
}

func (u UserDaoAdapter) Update(user User) error {
	return u.db.Table(u.core.TableName()).Save(&user).Error
}

func (u UserDaoAdapter) Delete(id int) error {
	return u.db.Table(u.core.TableName()).Where("id = ?", id).Delete(&User{}).Error
}
