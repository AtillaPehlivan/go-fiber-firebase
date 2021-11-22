package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"mocerize-api/model"
	"mocerize-api/service"
)

type MocksHandler struct {
	MockService service.MockService
}

func NewMocksHandler(mocksRoute fiber.Router, ms service.MockService) {
	handler := &MocksHandler{MockService: ms}

	log.Println("Mocks Routes initialized")

	mocksRoute.Get("", handler.index)
	mocksRoute.Get("/:uid", handler.find)
	mocksRoute.Post("", handler.create)
	mocksRoute.Put("/:uid", handler.update)
	mocksRoute.Delete("/:uid", handler.destroy)
	// for middleware
	// userRoute.Use()
}

func (mh MocksHandler) index(c *fiber.Ctx) error {

	mocks, err := mh.MockService.Index(fmt.Sprint(c.Locals("UID")))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Return results.
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   mocks,
	})
}

func (mh MocksHandler) find(c *fiber.Ctx) error {

	mock, err := mh.MockService.FindByUID(c.Params("uid"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "not found",
			"message": err.Error(),
		})
	}

	// Return results.
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   mock,
	})
}

func (mh MocksHandler) update(c *fiber.Ctx) error {

	var requestBody model.Mock

	err := c.BodyParser(&requestBody)

	if requestBody.UserUid == "" {
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"status": "error",
			"data":   "not found",
		})
	}

	if err != nil {
		log.Println("user PUT: " + err.Error())

		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status": "error",
			"data":   err.Error(),
		})
	}

	result, err := mh.MockService.Update(c.Params("uid"), &requestBody)

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

func (mh MocksHandler) create(c *fiber.Ctx) error {

	var requestBody model.Mock

	err := c.BodyParser(&requestBody)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	savedMock, err := mh.MockService.Create(fmt.Sprint(c.Locals("UID")), &requestBody)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"data":   savedMock,
	})
}

func (mh MocksHandler) destroy(c *fiber.Ctx) error {
	status, err := mh.MockService.Destroy(c.Params("uid"), fmt.Sprint(c.Locals("UID")))

	if err != nil || status == false {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
	})
}
