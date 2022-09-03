package controllers

import (
	"blog/admin/database"
	"blog/admin/helpers"
	"blog/admin/models"
	"fmt"
	"strconv"
	"time"

	"crypto/sha256"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
)

var secretKey = "SecretKey"

func Login(c *fiber.Ctx) error {
	var data = make(map[string]string) //kullanıcının girdiği user bilgileri
	data["username"] = c.FormValue("username")
	data["password"] = c.FormValue("password")
	var user models.User
	database.DB.Where("username = ?", data["username"]).First(&user) //find the user from database that we want
	//if user is not found
	if user.ID == 0 {
		helpers.SetFlash(c, "User not found!")
		return c.Redirect("/admin/login")
	}
	//if user is found confirm if password is true
	pw := fmt.Sprintf("%x", sha256.Sum256([]byte(data["password"])))
	if user.Password != pw {
		helpers.SetFlash(c, "Wrong Password!")
		return c.Redirect("/admin/login")
	}
	//Cookie and Session operations
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: jwt.NewTime(1900000000),
	})

	token, _ := claims.SignedString([]byte(secretKey))
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		Secure:   false,
	}
	c.Cookie(&cookie)

	helpers.SetFlash(c, "Login Successfully!")
	return c.Redirect("/admin")
}

func UserValidation(c *fiber.Ctx) (*jwt.StandardClaims, error) {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return token.Claims.(*jwt.StandardClaims), nil
}

func Logout(c *fiber.Ctx) error {
	logoutcookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&logoutcookie)
	return c.Redirect("/admin")
}
