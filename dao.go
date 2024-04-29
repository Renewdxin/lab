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

type DaoAdapter struct {
	db   *gorm.DB
	core UserCorePort
}

func NewUserDaoAdapter(db *gorm.DB, core UserCorePort) *DaoAdapter {
	return &DaoAdapter{
		db:   db,
		core: core,
	}
}

func (u DaoAdapter) Create(user User) error {
	return u.db.Table(u.core.TableName()).Create(&user).Error
}

func (u DaoAdapter) FindByID(id int) (User, error) {
	var user User
	err := u.db.Table(u.core.TableName()).Where("id = ?", id).First(&user).Error
	return user, err
}

func (u DaoAdapter) Login(password, id string) error {
	var user User
	return u.db.Table(u.core.TableName()).Where("account = ? AND password = ?", id, password).First(&user).Error
}

func (u DaoAdapter) Update(user User) error {
	return u.db.Table(u.core.TableName()).Save(&user).Error
}

func (u DaoAdapter) Delete(id int) error {
	return u.db.Table(u.core.TableName()).Where("id = ?", id).Delete(&User{}).Error
}
