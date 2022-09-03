package controllers

import (
	"blog/admin/helpers"
	"blog/admin/models"

	"github.com/gofiber/fiber/v2"
)

func AddSwiper(c *fiber.Ctx) error {
	//Guard operations
	_, err := UserValidation(c)
	if err != nil {
		helpers.SetFlash(c, "Please Login")
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
	helpers.SetFlash(c, "Swiper Slide Is Added Succesfully")
	return c.Redirect("/admin/swiperslide")
}

func RemoveSwiper(c *fiber.Ctx) error {
	//Guard operations
	_, err := UserValidation(c)
	if err != nil {
		helpers.SetFlash(c, "Please Login")
		return c.Redirect("/admin/login")
	}

	ssid, _ := helpers.ToInt(c.FormValue("post"))
	models.SwiperSlide{
		Postid: uint(ssid),
	}.Remove()

	helpers.SetFlash(c, "SwiperSlide Deleted Succesfully")
	return c.Redirect("/admin/swiperslide")
}
