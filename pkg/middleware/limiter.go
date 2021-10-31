package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"mocerize-api/pkg/config"
	"strconv"
)

func Limiter() fiber.Handler {

	defaultLimit := 100

	if environmentLimit, err := strconv.Atoi(config.Get("REQUEST_LIMIT")); err == nil {
		defaultLimit = environmentLimit
	}

	return limiter.New(limiter.Config{
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(&fiber.Map{
				"status":  "fail",
				"message": "You have requested too many in a single time-frame! Please wait another minute!",
			})
		},
		Max: defaultLimit,
	})
}
