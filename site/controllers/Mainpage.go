package controllers

import (
	"blog/admin/helpers"
	"blog/admin/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Mainpage struct {
}
type ContentPage struct {
}
type CategoryIndex struct {
}

func (Mainpage) Index(c *fiber.Ctx) error {
	posts := models.Post{}.GetAll()
	categories := models.Category{}.GetAll()

	//Swiper Slide operations
	ss := models.SwiperSlide{}.GetAll()
	var ssPosts []models.Post
	isswiperslide := false
	for _, value := range posts {
		for _, value2 := range ss {
			if value.ID == value2.Postid {
				isswiperslide = true
				break
			}
		}
		if isswiperslide == true {
			ssPosts = append(ssPosts, value)
		}
		isswiperslide = false
	}

	return c.Render("mainpage", fiber.Map{
		"Posts":      posts,
		"Categories": categories,
		"ssPosts":    ssPosts,
	})
}

func (ContentPage) Index(c *fiber.Ctx) error {
	category := c.Params("category")
	slug := c.Params("slug")
	post := models.Post{}.GetFromSlug(slug)
	categories := models.Category{}.GetAll()
	name := models.Category{}.Get(post.CategoryId).Name
	if name != category {
		c.Redirect("/")
	}

	pagepath := fmt.Sprintf("/articles/%s/%s", category, slug)
	viewercount, _ := helpers.Ga4Request(pagepath)

	return c.Render("contentpage", fiber.Map{
		"Post":       post,
		"Categories": categories,
		"Viewcount":  viewercount,
	})
}

func (CategoryIndex) Index(c *fiber.Ctx) error {
	category := models.Category{}.GetByName(c.Params("category"))
	posts := models.Post{}.GetFromCategory(int(category.ID))
	categories := models.Category{}.GetAll()

	return c.Render("postsbycategoryindex", fiber.Map{
		"Posts":      posts,
		"Categories": categories,
	})
}
