package main

import (
	config "main/configs"
	"main/handlers"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	config.Connect()

	app.Get("/dogs", handlers.GetDogs)
	app.Get("/dogs/{id}", handlers.GetDog)
	app.Post("/dogs", handlers.AddDog)
	app.Put("/dogs/{id}", handlers.UpdateDog)
	app.Delete("/dogs/{id}", handlers.DeleteDog)

	app.Logger().SetLevel("debug")
	app.Listen(":4000")
}
