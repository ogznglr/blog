package main

import (
	"blog/admin/database"
	"blog/admin/models"
	"blog/routes"
	"crypto/tls"
	"fmt"
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

	//Certificate
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("oguzhanguler.dev"),
		Cache:      autocert.DirCache("certs"),
	}

	TLSConfig := &tls.Config{
		GetCertificate: certManager.GetCertificate,
	}
	TLSConfig.Certificates = append(TLSConfig.Certificates, certManager.TLSConfig().Certificates...)

	// listener, _ := net.Listen("tcp", ":8080")
	listener, _ := tls.Listen("tcp", ":443", TLSConfig)
	//-------------------------------------------------------------------

	routes.Setup(app)
	app.Listener(listener)

}

// // Let’s Encrypt has rate limits: https://letsencrypt.org/docs/rate-limits/
// 	// It's recommended to use it's staging environment to test the code:
// 	// https://letsencrypt.org/docs/staging-environment/

// 	// Certificate manager
// 	m := &autocert.Manager{
// 		Prompt: autocert.AcceptTOS,
// 		// Replace with your domain
// 		HostPolicy: autocert.HostWhitelist("oguzhanguler.dev"),
// 		// Folder to store the certificates
// 		Cache: autocert.DirCache("./certs"),
// 	}

// 	// TLS Config
// 	cfg := &tls.Config{
// 		// Get Certificate from Let's Encrypt
// 		GetCertificate: m.GetCertificate,
// 		// By default NextProtos contains the "h2"
// 		// This has to be removed since Fasthttp does not support HTTP/2
// 		// Or it will cause a flood of PRI method logs
// 		// http://webconcepts.info/concepts/http-method/PRI
// 		NextProtos: []string{
// 			"http/1.1", "acme-tls/1",
// 		},
// 	}
// 	ln, err := tls.Listen("tcp", ":443", cfg)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Start server
// 	log.Fatal(app.Listener(ln))
