package main

import (
	"blog/admin/database"
	"blog/admin/models"
	"blog/routes"
	"crypto/tls"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/template/html"
)

func init() {
	database.Connection()
	database.Migrate(&models.Post{})
	database.Migrate(&models.User{})
	database.Migrate(&models.Category{})
	database.Migrate(&models.SwiperSlide{})
}

func main() {

	//Engine processess
	engine := html.New("./views", ".html")
	engine.AddFuncMap(fiber.Map{
		"getCategoryName": func(categoryId int) string {
			return models.Category{}.Get(categoryId).Name
		},
		"getRange": func(postNumber int) int {
			if postNumber <= 3 {
				return 1
			}
			return postNumber / 3
		},
		"getDate": func(t time.Time) string {
			return fmt.Sprintf("%d %s %d", t.Day(), t.Month().String(), t.Year())
		},
	})

	engine.Reload(true) //updates files in each rendering

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	//------------------------------------------------------------------------

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
	}))

	routes.RouteHandler{}.Setup(app)

	startServerHttp(app)

}

func startServerHttps(app *fiber.App) {
	// Load SSL certificate and private key files
	cert, err := tls.LoadX509KeyPair("oguzhanguler.dev.crt", "oguzhanguler.dev_key.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Create a TLS config object
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	ln, err := tls.Listen("tcp", ":443", tlsConfig)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(app.Listener(ln))
}

func startServerHttp(app *fiber.App) {
	app.Listen(":80")
}
