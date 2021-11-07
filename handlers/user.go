package handlers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"mocerize-api/service"
)

type UserHandler struct {
	UserService service.UserService
}

func NewUserHandler(userRoute fiber.Router, us service.UserService) {
	handler := &UserHandler{UserService: us}

	log.Println("userRoute := router.Group()")

	userRoute.Get("", handler.getUser)
	// for middleware
	// userRoute.Use()
}

func (uh UserHandler) getUser(c *fiber.Ctx) error {

	log.Println("aitlla getUser")

	user, err := uh.UserService.FindByToken(c.Params("UID"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Return results.
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   user,
	})
}
