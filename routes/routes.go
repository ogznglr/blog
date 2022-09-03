package routes

import (
	admin "blog/admin/controllers"
	site "blog/site/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) error {

	//site
	app.Get("/", site.Mainpage{}.Index)
	app.Get("/articles/:category/:slug", site.ContentPage{}.Index)
	app.Get("/categories/:category", site.CategoryIndex{}.Index)

	//admin
	//admin GET
	app.Get("/admin", admin.Dashboard{}.Index)
	app.Get("/admin/new_post", admin.NewPost{}.Index)
	app.Get("/admin/delete/:id", admin.DeletePost)
	app.Get("/admin/edit/:id", admin.Edit{}.Index)
	app.Get("/admin/login", admin.LoginPage{}.Index)
	app.Get("/admin/logout", admin.Logout)
	app.Get("/admin/categories", admin.Categories{}.Index)
	app.Get("/admin/categories/delete/:id", admin.DeleteCategory)
	app.Get("/admin/swiperslide", admin.SwiperSlidePage{}.Index)

	//admin POST
	app.Post("/admin/new_post", admin.AddPost)
	app.Post("/admin/edit/:id", admin.EditPost)
	app.Post("/admin/login", admin.Login)
	app.Post("/admin/categories/add", admin.AddCategory)
	app.Post("/admin/add_swiper", admin.AddSwiper)
	app.Post("/admin/remove_swiper", admin.RemoveSwiper)
	//serve static files
	app.Static("/admin/assets/", "./admin/assets/")
	app.Static("/site/assets/", "./site/assets/")
	app.Static("/uploads/", "./uploads/")
	return nil
}
