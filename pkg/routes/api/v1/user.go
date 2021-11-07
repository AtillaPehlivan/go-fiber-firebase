package route

import (
	"github.com/gofiber/fiber/v2"
	"mocerize-api/handlers"
	"mocerize-api/pkg/firebase/storage"
	"mocerize-api/service"
)

func SetupUserRoute(router fiber.Router, userService service.UserService) {

	router.Get("upload-single", func(ctx *fiber.Ctx) error {
		a := storage.FirebaseStorage{}
		a.UploadSingle()
		return ctx.SendStatus(fiber.StatusOK)
	})

	userRoute := router.Group("/user")

	handlers.NewUserHandler(userRoute, userService)

}
