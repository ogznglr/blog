package controllers

import (
	"blog/admin/database"
	"blog/admin/helpers"
	"blog/admin/models"
	"fmt"
	"strconv"

	"crypto/sha256"

	"github.com/gofiber/fiber/v2"
	"github.com/ogznglr/session"
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
	s := session.New(24)
	s.Set(c, strconv.Itoa(int(user.ID)), secretKey)

	session.SetFlash(c, "Login Successfully!")
	return c.Redirect("/admin")
}

func UserValidation(c *fiber.Ctx) (string, error) {
	s := session.New()
	return s.Get(c, secretKey)
}

func Logout(c *fiber.Ctx) error {
	s := session.New()
	s.Delete(c)
	return c.Redirect("/admin")
}
