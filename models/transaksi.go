package models

import (
	"gorm.io/gorm"
)


type Transaksi struct {
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
func ReadTransaksi(db *gorm.DB, transaksi *[]Transaksi) (err error) {
	err = db.Find(transaksi).Error
	if err != nil {
		return err
	}
	return nil
}
func CreateTransaksi(db *gorm.DB, newTransaksi *Transaksi) (err error) {
	err = db.Create(newTransaksi).Error
	if err != nil {
		return err
	}
	return nil
}