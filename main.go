package main

import (
	"blog/admin/database"
	"blog/admin/models"
	"blog/routes"
	"crypto/tls"
	"fmt"
	"net"
	"time"

	"golang.org/x/crypto/acme/autocert"

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
	//Certificate
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("oguzhanguler.dev", "www.oguzhanguler.dev"),
		Cache:      autocert.DirCache("certs"),
	}

	TLSConfig := &tls.Config{
		GetCertificate: certManager.GetCertificate,
	}

	// listener, _ := net.Listen("tcp", ":8080")
	listener, _ := net.Listen("tcp", ":443")
	listener = tls.NewListener(listener, TLSConfig)

	//-------------------------------------------------------------------

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
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
	}))

	routes.Setup(app)
	app.ListenTLS(":443", "./go-server.crt", "./go-server.key")

}
