package controllers

import (
	"blog/admin/helpers"
	"blog/admin/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
)

func AddCategory(c *fiber.Ctx) error {
	//Guard process
	_, err := UserValidation(c)
	if err != nil {
		helpers.SetFlash(c, "Please Login")
		return c.Redirect("/admin/login")
	}

	categoryName := c.FormValue("category-name")
	categorySlug := slug.Make(categoryName)

	if categoryName == "" {
		helpers.SetFlash(c, "Category couldn't been saved.")
		return c.Redirect("/admin/categories")
	}
	category := models.Category{
		Name: categoryName,
		Slug: categorySlug,
	}
	category.Add()

	helpers.SetFlash(c, "Category is saved successfully!")
	return c.Redirect("/admin/categories")
}

func DeleteCategory(c *fiber.Ctx) error {
	//Guard Process
	_, err := UserValidation(c)
	if err != nil {
		helpers.SetFlash(c, "Please Login")
		return c.Redirect("/admin/login")
	}

	deleteId, err := helpers.ToInt(c.Params("id"))
	if err != nil {

	}
	category := models.Category{}.Get(deleteId)
	category.Delete()
	helpers.SetFlash(c, "Category is deleted successfully!")
	return c.Redirect("/admin/categories")
}
