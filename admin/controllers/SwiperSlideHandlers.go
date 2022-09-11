package controllers

import (
	"blog/admin/helpers"
	"blog/admin/models"

	"github.com/gofiber/fiber/v2"
	"github.com/ogznglr/session"
)

func AddSwiper(c *fiber.Ctx) error {
	//Guard operations
	_, err := UserValidation(c)
	if err != nil {
		session.SetFlash(c, "Please Login")
		return c.Redirect("/admin/login")
	}

	postid, _ := helpers.ToInt(c.FormValue("post"))
	post := models.Post{}.Get(postid)
	ss := models.SwiperSlide{}.GetAll()

	for _, value := range ss {
		if value.Postid == post.ID {
			return c.Redirect("/admin/swiperslide")
		}
	}
	models.SwiperSlide{
		Postid: post.ID,
	}.Add()
	session.SetFlash(c, "Swipe Slide has added successfully!")
	return c.Redirect("/admin/swiperslide")
}

func RemoveSwiper(c *fiber.Ctx) error {
	//Guard operations
	_, err := UserValidation(c)
	if err != nil {
		session.SetFlash(c, "Please Login")
		return c.Redirect("/admin/login")
	}

	ssid, _ := helpers.ToInt(c.FormValue("post"))
	models.SwiperSlide{
		Postid: uint(ssid),
	}.Remove()

	session.SetFlash(c, "Swiper Slide deleted successfully!")
	return c.Redirect("/admin/swiperslide")
}
