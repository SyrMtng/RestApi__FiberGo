package controllers

import (
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"ilmudata/Project2/database"
	"ilmudata/Project2/models"
)

type LoginForm struct {
	// declare variables
	Username string `form:"username" validate:"required"`
	Password string `form:"password" validate:"required"`
}

type AuthController struct {
	// declare variables
	store *session.Store
	Db *gorm.DB
}
func InitAuthController(s *session.Store) *AuthController {
	db := database.InitDb()
	// gorm
	db.AutoMigrate(&models.User{})
	return &AuthController{Db: db, store: s}
}
		// post /login
		func (controller *AuthController) LoginPosted(c *fiber.Ctx) error {
			sess, err := controller.store.Get(c)
			if err != nil {
				panic(err)
			}

			var user models.User
			var myform LoginForm
			if err := c.BodyParser(&myform); err != nil {
				return c.JSON(fiber.Map{
					"message": "Kesalahan",
				})
			}

			er := models.FindByUsername(controller.Db, &user, myform.Username)
			if er != nil {
				return c.JSON(fiber.Map{
					"message": "Kesalahan username",
				})
			}

			// hardcode auth
			compare := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(myform.Password))
			if compare != nil {
				exp := time.Now().Add(time.Hour * 72)
				claims := jwt.MapClaims{
					"id":    user.Id,
					"ussername":  user.Username,
					"exp":   exp.Unix(),
				}
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
				
				t, err := token.SignedString([]byte("mysecretpassword"))
				if err != nil {
					return c.SendStatus(fiber.StatusInternalServerError)
				}

				sess.Set("username", user.Username)
				sess.Set("userID", user.Id)
				sess.Save()

				return c.JSON(fiber.Map{
					"message": "Berhasil Login",
					"token":   t,
					"expired": exp.Format("2006-01-02 15:04:05"),
				})
			}
			return c.JSON(fiber.Map{
				"message": "Kesalahan Login",
			})
		}

// post registrasi
func (controller *AuthController) RegistrasiPosted(c *fiber.Ctx) error {
	var myform models.User
	var convertpass LoginForm

	if err := c.BodyParser(&myform); err != nil {
		return c.JSON(fiber.Map{
			"message": "Berhasil",
		})
	}
	comvertpassword, _ := bcrypt.GenerateFromPassword([]byte(convertpass.Password), 10)
	sHash := string(comvertpassword)

	myform.Password = sHash

	// save product
	err := models.CreateUser(controller.Db, &myform)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "Kesalahan",
		})
	}
	// if succeed
	return c.JSON(fiber.Map{
		"message": "Berhasil",
	})
}

//profile
func (controller *AuthController) Profile(c *fiber.Ctx) error {
	var user []models.User
	err := models.ReadProfile(controller.Db, &user)
	if err!=nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	return c.JSON(fiber.Map{
		"AkunProfile": user,
	})
}
//logout
func (controller *AuthController) Logout(c *fiber.Ctx) error {
	
	sess,err := controller.store.Get(c)
	if err!=nil {
		panic(err)
	}
	sess.Destroy()
	return c.JSON(fiber.Map{
		"message": "Berhasil Logout",
	})
}