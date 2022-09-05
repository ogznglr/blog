package controllers

import (
	"blog/admin/helpers"
	"blog/admin/models"
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
	"github.com/ogznglr/session"
)

func AddPost(c *fiber.Ctx) error {
	//Guard process
	_, err := UserValidation(c)
	if err != nil {
		session.SetFlash(c, "Please Login")
		return c.Redirect("/admin/login")
	}
	//if user is valid
	title := c.FormValue("post-title")
	desc := c.FormValue("post-desc")
	content := c.FormValue("post-content")
	categoryid, err := helpers.ToInt(c.FormValue("post-category"))
	slug := slug.Make(title)
	//upload picture
	picture, _ := c.FormFile("post-pic")
	if err != nil {
		return err
	}
	post := models.Post{
		Title:       title,
		Content:     content,
		Description: desc,
		CategoryId:  categoryid,
		Slug:        slug,
	}

	//if picture exists
	if picture == nil {
		return c.JSON(fiber.Map{
			"error": "Picture has not found",
		})
	}

	//save file
	err = c.SaveFile(picture, fmt.Sprintf("/uploads/%s", picture.Filename))
	if err != nil {
		return err
	}

	post.PictureURL = fmt.Sprintf("/uploads/%s", picture.Filename)

	//if everything is right for post
	if post.Content == "" {
		return c.JSON(fiber.Map{
			"error": "Content has not found",
		})
	}
	if post.Title == "" {
		return c.JSON(fiber.Map{
			"error": "Title has not found",
		})
	}
	if post.Description == "" {
		return c.JSON(fiber.Map{
			"error": "Description has not found",
		})
	}
	if categoryid == 0 {
		return c.JSON(fiber.Map{
			"error": "Category has not found",
		})
	}
	// Add to database
	post.Add()

	//Send A Flash message
	helpers.SetFlash(c, "Post published successfully!")

	return c.Redirect("/admin", http.StatusSeeOther)
	//TODO Alert
}

func DeletePost(c *fiber.Ctx) error {
	_, err := UserValidation(c)
	if err != nil {
		session.SetFlash(c, "Please Login")
		return c.Redirect("/admin/login")
	}
	//if user is valid
	deleteid, err := helpers.ToInt(c.Params("id"))
	if err != nil {
		return c.JSON(fiber.Map{
			"error": "yanlış veri girildi",
		})
	}

	post := models.Post{}.Get(deleteid)

	_ = os.Remove(helpers.DeleteFirstSlender(post.PictureURL))

	post.Delete()
	models.SwiperSlide{Postid: post.ID}.Remove()

	helpers.SetFlash(c, "Post is deleted succsessfully!")

	return c.Redirect("/admin")

}

func EditPost(c *fiber.Ctx) error {
	_, err := UserValidation(c)
	if err != nil {
		session.SetFlash(c, "Please Login")
		return c.Redirect("/admin/login")
	}
	//if user is valid
	id, _ := helpers.ToInt(c.FormValue("id"))
	title := c.FormValue("post-title")
	desc := c.FormValue("post-desc")
	content := c.FormValue("post-content")
	categoryid, _ := helpers.ToInt(c.FormValue("post-category"))
	slug := slug.Make(title)
	is_selected := c.FormValue("is-selected")

	refPost := models.Post{}.Get(id)

	//if picture upload has been made
	if is_selected == "1" {
		picture, err := c.FormFile("post-pic")

		//if we can't recieve picture return to admin
		if err != nil {
			return c.Redirect("/admin")
		}

		//first delete old picture
		_ = os.Remove(helpers.DeleteFirstSlender(refPost.PictureURL))

		_ = c.SaveFile(picture, fmt.Sprintf("/uploads/%s", picture.Filename))
		refPost.PictureURL = fmt.Sprintf("/uploads/%s", picture.Filename)
	}

	post := models.Post{
		Title:       title,
		Description: desc,
		Content:     content,
		CategoryId:  categoryid,
		Slug:        slug,
		PictureURL:  refPost.PictureURL,
	}
	refPost.Updates(post)

	helpers.SetFlash(c, "Post is edited succsessfully!")

	c.Redirect("/admin")
	return nil
}
