package middleware

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
	"mocerize-api/pkg/firebase/auth"
	"strings"
)

func Auth(c *fiber.Ctx) error {

	authHeader := c.Get("Authorization", "")
	if authHeader == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	headerToken := strings.ReplaceAll(authHeader, "Bearer ", "")

	token, err := auth.Client().VerifyIDToken(context.Background(), headerToken)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	c.Locals("uid", token.UID)

	log.Println(c.Locals("UID"))
	return c.Next()
}
