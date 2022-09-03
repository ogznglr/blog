package controllers

import (
	"blog/admin/helpers"
	"blog/admin/models"

	"github.com/gofiber/fiber/v2"
)

type Dashboard struct {
}

type NewPost struct {
}

type Edit struct {
}
type LoginPage struct {
}
type Categories struct {
}
type SwiperSlidePage struct {
}

func (Dashboard) Index(c *fiber.Ctx) error {
	_, err := UserValidation(c)
	if err != nil {
		helpers.SetFlash(c, "Please Login")
		return c.Redirect("/admin/login")
	}
	//Alert operations
	message := helpers.GetFlash(c)
	isalert := false
	if message != "" {
		isalert = true
	}

	alert := make(map[string]interface{})
	alert["is-alert"] = isalert
	alert["message"] = message

	posts := models.Post{}.GetAll()

	return c.Render("index", fiber.Map{
		"Posts": posts,
		"Alert": alert,
	})
}

func (NewPost) Index(c *fiber.Ctx) error {
	_, err := UserValidation(c)
	if err != nil {
		helpers.SetFlash(c, "Please Login")
		return c.Redirect("/admin/login")
	}
	categories := models.Category{}.GetAll()

	return c.Render("add_index", fiber.Map{
		"Categories": categories,
	})
}

func (Edit) Index(c *fiber.Ctx) error {
	_, err := UserValidation(c)
	if err != nil {
		helpers.SetFlash(c, "Please Login")
		return c.Redirect("/admin/login")
	}

	id, err := helpers.ToInt(c.Params("id"))

	if err != nil {
		return err
	}

	post := models.Post{}.Get(id)
	categories := models.Category{}.GetAll()

	c.Render("edit_index", fiber.Map{
		"Post":       post,
		"Categories": categories,
	})
	return nil
}
func (LoginPage) Index(c *fiber.Ctx) error {
	message := helpers.GetFlash(c)
	isalert := false
	if message != "" {
		isalert = true
	}
	//alert operations
	alert := make(map[string]interface{})
	alert["is-alert"] = isalert
	alert["message"] = message

	c.Render("login", fiber.Map{
		"Alert": alert})
	return nil
}
func (Categories) Index(c *fiber.Ctx) error {
	_, err := UserValidation(c)
	if err != nil {
		helpers.SetFlash(c, "Please Login")
		return c.Redirect("/admin/login")
	}
	//Flash message operations
	message := helpers.GetFlash(c)
	isalert := false
	if message != "" {
		isalert = true
	}

	alert := make(map[string]interface{})
	alert["is-alert"] = isalert
	alert["message"] = message

	//Get all the categories and send to client
	categories := models.Category{}.GetAll()

	return c.Render("categories_index", fiber.Map{
		"Categories": categories,
		"Alert":      alert,
	})
}
func (SwiperSlidePage) Index(c *fiber.Ctx) error {
	_, err := UserValidation(c)
	if err != nil {
		helpers.SetFlash(c, "Please Login")
		return c.Redirect("/admin/login")
	}
	//Flash message operations
	message := helpers.GetFlash(c)
	isalert := false
	if message != "" {
		isalert = true
	}

	alert := make(map[string]interface{})
	alert["is-alert"] = isalert
	alert["message"] = message

	//Get all the posts and send to client
	allposts := models.Post{}.GetAll()
	ss := models.SwiperSlide{}.GetAll()
	var posts []models.Post
	var ssPosts []models.Post
	isswiperslide := false

	for _, value := range allposts {
		for _, value2 := range ss {
			if value.ID == value2.Postid {
				isswiperslide = true
				break
			}
		}
		if isswiperslide == false {
			posts = append(posts, value)
		} else {
			ssPosts = append(ssPosts, value)
		}
		isswiperslide = false
	}

	return c.Render("swiperslide_index", fiber.Map{
		"Posts":  posts,
		"SPosts": ssPosts,
		"Alert":  alert,
	})
}
