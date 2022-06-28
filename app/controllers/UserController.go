package controllers

import (
	"github.com/gowlfer/gowlfer/app/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gowlfer/gowlfer/internal/utils/globals"
)

func RegisterUser(ctx *fiber.Ctx) error {
	data := struct {
		Email    string `json:"email"`
		UserName string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	var user models.GowlferUser

	if got := user.DoesUserExist(data.Email); got {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.Status(400).SendString("Invalid Credentials")
	}

	err := user.CreateUsers(data.Email, data.UserName, data.Password)

	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		ctx.JSON(fiber.Map{
			"error": err,
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "User Created",
	})
}

func LoginUser(ctx *fiber.Ctx) error {

	data := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	var user models.GowlferUser

	// Check Email
	if got := user.DoesUserExist(data.Email); !got {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.Status(400).SendString("Invalid Credentials")
	}

	// Check password
	if got := user.CheckPassword(data.Email, data.Password); !got {
		return ctx.Status(400).SendString("Invalid Credentials")
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
	})

	_, err := claims.SignedString([]byte(globals.SecretKey))

	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return ctx.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Login Successfully",
	})
}

func GetUser(ctx *fiber.Ctx) error {
	var user models.GowlferUser

	user = user.GetUserByID(ctx.Query("id"))

	return ctx.JSON(fiber.Map{
		"message":  "User Found",
		"email":    user.UserEmail,
		"username": user.UserName,
	})
}
