package services

import (
	"api/initializers"
	"api/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis/v2"
	"golang.org/x/crypto/bcrypt"
)

var storage = redis.New()
var store = session.New(session.Config{
	KeyLookup: "cookie:my_reviews_session",
	Storage:   storage,
})

func Login(c *fiber.Ctx) error {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user := models.User{}
	initializers.DB.Where(&models.User{Email: body.Email}).First(&user)

	if user.ID == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}

	session, err := store.Get(c)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	session.Set("user_id", user.ID)

	if err := session.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}

func Register(c *fiber.Ctx) error {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	password := []byte(body.Password)
	hashedPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	user := models.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: string(hashedPassword),
	}

	initializers.DB.Create(&user)

	if user.ID == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Save user failed",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"user":    user,
	})
}
