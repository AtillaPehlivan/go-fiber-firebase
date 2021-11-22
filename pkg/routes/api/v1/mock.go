package route

import (
	"github.com/gofiber/fiber/v2"
	"mocerize-api/handlers"
	"mocerize-api/service"
)

func SetupMockRoute(router fiber.Router, mockService service.MockService) {

	mockRoute := router.Group("/mocks")

	handlers.NewMocksHandler(mockRoute, mockService)

}
