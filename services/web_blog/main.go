package main

import (
	"fmt"
	"log"

	"github.com/beego/beego/v2/client/orm"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/kazimovzaman2/Go-Microservices/web_blog.git/database"
	"github.com/kazimovzaman2/Go-Microservices/web_blog.git/models"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
	}
	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello Universe!")
	})
	orm.RegisterModel(new(models.Authors))
	database.ConnectDB()
	log.Fatal(app.Listen(":3069"))
}
