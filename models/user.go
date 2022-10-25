package models

import (
	"gorm.io/gorm"
)


type User struct {
	gorm.Model
	Id	int	`form:"id" json:"id" validate:"required"`
	Email	string `form:"email" json:"email" validate:"required"`
	Username	string `form:"username" json:"username" validate:"required"`
	Password	string `form:"password" json:"password" validate:"required"`
}

// CRUD
func CreateUser(db *gorm.DB, newUser *User) (err error) {
	err = db.Create(newUser).Error
	if err != nil {
		return err
	}
	return nil
}
func FindByUsername(db *gorm.DB, user *User, username string) (err error) {
	err = db.Where("username=?", username).First(user).Error
	if err != nil {
		return err
	}
	return nil
}
func ReadProfile(db *gorm.DB, user *[]User)(err error) {
	err = db.Find(user).Error
	if err != nil {
		return err
	}
	return nil
}