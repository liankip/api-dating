package main

import (
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	app, err := InitializeApplication()
	if err != nil {
		log.Fatal("Failed to initialize application: ", err)
	}

	println("Server running on :4000")
	if err := app.Listen(":4000"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
