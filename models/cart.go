package models

import (
	"gorm.io/gorm"
)


type Cart struct {
	gorm.Model
	Id			int     `form:"id" json:"id" validate:"required"`
	IdUser		int     `form:"iduser" json:"iduser" validate:"required"`
	IdProduck 	int     `form:"idproduck" json:"idproduck" validate:"required"`
	Name		string  `form:"name" json:"name" validate:"required"`
	Price		float32	`form:"price" json:"price" validate:"required"`
	Status    	string  `form:"status" json:"status" validate:"required"`
	Gambar		string	`form:"gambar" json: "gambar" validate:"required"`
}

// CRUD
func CreateCart(db *gorm.DB, newCart *Cart) (err error) {
	err = db.Create(newCart).Error
	if err != nil {
		return err
	}
	return nil
}
func ReadCart(db *gorm.DB, cart *[]Cart) (err error) {
	err = db.Find(cart).Error
	if err != nil {
		return err
	}
	return nil
}
func ReadCartById(db *gorm.DB, cart *Cart, id int) (err error) {
	err = db.Where("id=?", id).Find(cart).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteCartById(db *gorm.DB, cart *Cart, id int) (err error) {
	db.Where("id=?", id).Delete(cart)

	return nil
}