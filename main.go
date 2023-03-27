package main

import (
	"blog/admin/database"
	"blog/admin/models"
	"blog/routes"
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"time"

	"github.com/caddyserver/certmagic"
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

	httpapp := fiber.New(fiber.Config{
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
	routes.RouteHandler{}.Setup(httpapp)

	cache := certmagic.NewCache(certmagic.CacheOptions{
		GetConfigForCert: func(cert certmagic.Certificate) (*certmagic.Config, error) {
			// do whatever you need to do to get the right
			// configuration for this certificate; keep in
			// mind that this config value is used as a
			// template, and will be completed with any
			// defaults that are set in the Default config
			return &certmagic.Config{
				// ...
			}, nil
		},
	})

	magic := certmagic.New(cache, certmagic.Config{
		// any customizations you need go here
	})

	myACME := certmagic.NewACMEIssuer(magic, certmagic.ACMEIssuer{
		CA:     certmagic.LetsEncryptStagingCA,
		Email:  "admin@oguzhanguler.dev",
		Agreed: true,
		// plus any other customizations you need
	})

	magic.Issuers = append(magic.Issuers, myACME)

	err := magic.ManageSync(context.TODO(), []string{"oguzhanguler.dev"})
	if err != nil {
		panic(err)
	}

	tlsConfig := magic.TLSConfig()
	tlsConfig.NextProtos = append([]string{"h2", "http/1.1"}, tlsConfig.NextProtos...)

	_, err = tls.Listen("tcp", ":443", tlsConfig)

	//start the server with given listener.

	go log.Fatal(httpapp.Listen(":8080"))

}
