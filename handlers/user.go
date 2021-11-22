package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"mocerize-api/model"
	"mocerize-api/service"
)

type UserHandler struct {
	UserService service.UserService
}

func NewUserHandler(userRoute fiber.Router, us service.UserService) {
	handler := &UserHandler{UserService: us}

	log.Println("User Routes initialized")

	userRoute.Get("", handler.getUser)
	userRoute.Put("", handler.update)
	// for middleware
	// userRoute.Use()
}

func (uh UserHandler) getUser(c *fiber.Ctx) error {

	user, err := uh.UserService.FindByUID(fmt.Sprint(c.Locals("UID")))

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

func (uh UserHandler) update(c *fiber.Ctx) error {

	var requestBody model.User

	err := c.BodyParser(&requestBody)
	if err != nil {
		log.Println("user PUT: " + err.Error())

		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status": "error",
			"data":   err.Error(),
		})
	}

	result, err := uh.UserService.Update(fmt.Sprint(c.Locals("UID")), &requestBody)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status": "error",
			"data":   err.Error(),
		})
	}

	// Return results.
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   result,
	})
}
