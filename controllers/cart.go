package controllers

import (
	"strconv"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"

	"ilmudata/Project2/database"
	"ilmudata/Project2/models"
)

// type ProductForm struct {
// 	Email string `form:"email" validate:"required"`
// 	Address string `form:"address" validate:"required"`
// }

type CartController struct {
	// declare variables
	store *session.Store
	Db *gorm.DB
}
func InitCartController(s *session.Store) *CartController {
	db := database.InitDb()
	// gorm
	db.AutoMigrate(&models.Cart{})
	return &CartController{Db: db, store: s}
}

func (controller *CartController) IndexCart(c *fiber.Ctx) error {
	// params := c.AllParams()
	// id, _ := strconv.Atoi(params["id"])

	var cart []models.Cart
	err := models.ReadCart(controller.Db, &cart)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.JSON(cart)
}

func (controller *CartController) AddPostedCart(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	iduser := claims["id"].(float64)
	var idi int = int(iduser)
	
	sess, errr := controller.store.Get(c)
	if errr != nil {
		panic(errr)
	}
	val := sess.Get("username")
	vall := sess.Get("userID")

	var product models.Product
	err := models.ReadProductById(controller.Db, &product, idn)
	if err != nil {
		return c.JSON(fiber.Map{
			"Message": "Kesalahan data", 
		})
	}

	var myform models.Cart

	if err := c.BodyParser(&myform); err != nil {
		return c.SendStatus(400)
	}
	myform.IdUser = idi
	myform.IdProduck = product.Id
	myform.Name = product.Name
	myform.Gambar = product.Gambar
	myform.Price = product.Price
	myform.Status = "Belum Bayar"
	// save product
	errrr := models.CreateCart(controller.Db, &myform)
	if errrr != nil {
		return c.JSON(fiber.Map{
			"Message": "Kesalahan data", 
		})
	}
	// if succeed
	return c.JSON(fiber.Map{
		"Cart": myform,
		"Akun": val,
		"IDAkun": vall,
	})
}

func (controller *CartController) AddPostedBayar(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	iduser := claims["id"].(float64)
	var idi int = int(iduser)

	var cart models.Cart
	err := models.ReadCartById(controller.Db, &cart, idn)
	if err != nil {
		return c.JSON(fiber.Map{
			"Message": "Kesalahan data", 
		})
	}

	var myform models.Transaksi

	if err := c.BodyParser(&myform); err != nil {
		return c.SendStatus(400)
	}
	myform.IdUser = idi
	myform.IdProduck = cart.Id
	myform.Name = cart.Name
	myform.Gambar = cart.Gambar
	myform.Price = cart.Price
	myform.Status = "Sudah Bayar"
	models.DeleteCartById(controller.Db, &cart, idn)
	// save product
	errrr := models.CreateTransaksi(controller.Db, &myform)
	if errrr != nil {
		return c.JSON(fiber.Map{
			"Message": "Kesalahan data", 
		})
	}
	// if succeed
	return c.JSON(fiber.Map{
		"Cart": myform,
	})
}

// Untuk Delete berdasarkan id yang di terima
func (controller *CartController) DeleteCartById(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	var cart models.Cart
	models.DeleteCartById(controller.Db, &cart, idn)
	return c.JSON(fiber.Map{
		"message": "data was deleted",
	})
}