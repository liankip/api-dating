package main

import (
	"api-dating/delivery/http"
	"api-dating/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userHandler *http.UserHandler, swipeHandler *http.SwipeHandler) {

	api := app.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			auth := v1.Group("/auth")
			{
				auth.Post("/signup", userHandler.Signup)
				auth.Post("/signin", userHandler.Signin)
			}

			profile := v1.Group("/profile", middleware.JWTAuthMiddleware)
			{
				profile.Post("/swipe", swipeHandler.Swipe)

				premium := profile.Group("/premium")
				{
					premium.Post("/active", userHandler.PremiumActive)
				}
			}
		}
	}
}
