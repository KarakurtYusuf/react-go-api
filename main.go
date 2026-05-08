package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("hello world")
	app := fiber.New()

	// 2. İlk Uç Noktamızı (Endpoint) yazıyoruz
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "Merhaba Dünya! İlk Go API sunucum calisiyor!",
		})
	})

	log.Fatal(app.Listen(":4000"))
}
