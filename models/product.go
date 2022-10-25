package models

import (
	"gorm.io/gorm"
)


type Product struct {
	gorm.Model
	Id       int     `form:"id" json:"id" validate:"required"`
	Name	string  `form:"name" json:"name" validate:"required"`
	Deskripsi	string	`form:"deskripsi" json:"deskripsi" validate:"required"`
	Price	float32	`form:"price" json:"price" validate:"required"`
	Gambar	string	`form:"gambar" json: "gambar" validate:"required"`
}

// CRUD
func CreateProduct(db *gorm.DB, newProduct *Product) (err error) {
	err = db.Create(newProduct).Error
	if err != nil {
		return err
	}
	return nil
}
func ReadProduct(db *gorm.DB, product *[]Product)(err error) {
	err = db.Find(product).Error
	if err != nil {
		return err
	}
	return nil
}
func ReadProductById(db *gorm.DB, product *Product, id int)(err error) {
	err = db.Where("id=?", id).First(product).Error
	if err != nil {
		return err
	}
	return nil
}
func UpdateProduct(db *gorm.DB, product *Product)(err error) {
	db.Save(product)
	
	return nil
}
func DeleteProductById(db *gorm.DB, product *Product, id int)(err error) {
	db.Where("id=?", id).Delete(product)
	
	return nil
}