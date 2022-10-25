package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"

	"ilmudata/Project2/database"
	"ilmudata/Project2/models"
)

// type ProductForm struct {
// 	Email string `form:"email" validate:"required"`
// 	Address string `form:"address" validate:"required"`
// }

type TransaksiController struct {
	// declare variables
	store *session.Store
	Db *gorm.DB
}
func InitTransaksiController(s *session.Store) *TransaksiController {
	db := database.InitDb()
	// gorm
	db.AutoMigrate(&models.Transaksi{})
	return &TransaksiController{Db: db, store: s}
}

func (controller *TransaksiController) IndexTransaksi(c *fiber.Ctx) error {
	// params := c.AllParams()
	// id, _ := strconv.Atoi(params["id"])

	var transaksi []models.Transaksi
	err := models.ReadTransaksi(controller.Db, &transaksi)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.JSON(transaksi)
}

