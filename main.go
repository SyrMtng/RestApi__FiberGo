package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	jwtware "github.com/gofiber/jwt/v3"

	"ilmudata/Project2/controllers"
)

func main() {
	// session
	store := session.New()

	app := fiber.New()

	// static
	app.Static("/public","./public")

	// controllers
	prodController := controllers.InitProductController(store)
	authController := controllers.InitAuthController(store)
	cartController := controllers.InitCartController(store)
	tranController := controllers.InitTransaksiController(store)

	cekAkun := func(c *fiber.Ctx) error {
		sess,_ := store.Get(c)
		val := sess.Get("username")
		vall := sess.Get("userID")
		
		if val != nil && vall != nil{
			return c.Next()
		}

		return c.JSON(fiber.Map{
			"message": "Silahkan Login Dahulu",
		})

	}

	// Untuk Bagian Home and CRUD
	prod := app.Group("/products")
	prod.Get("/", prodController.IndexProduct)
	prod.Post("/create", cekAkun, prodController.AddPostedProduct)
	prod.Get("/detail/:id", prodController.GetDetailProduct)
	prod.Post("/editproduct/:id", cekAkun, prodController.EditlPostedProduct)
	prod.Delete("/deleteproduct/:id", cekAkun, prodController.DeleteProduct)

	//Bagian Login
	app.Post("/login",authController.LoginPosted)
	app.Post("/Registrasi",authController.RegistrasiPosted)
	app.Get("/logout",authController.Logout)
	app.Get("/profile",authController.Profile)
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("mysecretpassword"),
	}))

	//Bagian ShoopingCart
	car := app.Group("/shooping")
	car.Get("/", cekAkun, cartController.IndexCart)
	car.Post("/cart/:id", cekAkun, cartController.AddPostedCart)
	car.Delete("/deletecart/:id", cekAkun, cartController.DeleteCartById)
	car.Post("/bayar/:id", cekAkun, cartController.AddPostedBayar)

	//Bagian Transaksi
	app.Get("/Transaksi",cekAkun, tranController.IndexTransaksi)
	
	app.Listen(":3000")
}