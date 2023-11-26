package main

import (
	"backend1/database"
	"backend1/middleware"
	"backend1/route"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//mesin utama 
    app := fiber.New()
	database.Connect()
	//route 
	app.Post("/login",route.Login )
	app.Post("/register", route.Register)
	app.Get("/getUser", middleware.JwtProtect() ,route.GetUSer)
	// Router      ("/" = endpoint, fungsi logic )
    app.Get("/", func (c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

	app.Get("/getAll", route.GetAlluser)
	app.Put("/update", middleware.JwtProtect(),route.Update)

	//akan berjalan di port mana 
	//aplikasi ini akan berjalan di port 3000
    log.Fatal(app.Listen(":3000"))
}