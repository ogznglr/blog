package helpers

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

const name = "FlashName"

func SetFlash(c *fiber.Ctx, value string) {
	c.Cookie(&fiber.Cookie{
		Name:     name,
		Value:    value,
		HTTPOnly: true,
		Expires:  time.Now().Add(1 * time.Minute),
	})
}

func GetFlash(c *fiber.Ctx) string {
	theMessage := c.Cookies(name)

	//have an emty flash message to delete the cookie in user's browser
	c.Cookie(&fiber.Cookie{
		Name:     name,
		HTTPOnly: true,
		Expires:  time.Now().Add(-10 * time.Minute),
	})
	return theMessage
}
