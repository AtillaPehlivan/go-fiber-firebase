package route

import (
	"github.com/gofiber/fiber/v2"
	"mocerize-api/pkg/firebase/storage"
)

func SetupApiV1(app *fiber.App) {

	// middleware.Auth
	apiRoute := app.Group("/api")

	apiV1Route := apiRoute.Group("/v1")

	apiV1Route.Get("upload-single", func(ctx *fiber.Ctx) error {
		a := storage.FirebaseStorage{}
		a.UploadSingle()
		return ctx.SendStatus(fiber.StatusOK)
	})

	userRoute := apiV1Route.Group("/user")
	userRoute.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("atilla user")
	})

}
