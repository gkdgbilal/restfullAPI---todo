package middleware

import (
	"RestFullAPI-todo/api/utils"
	"RestFullAPI-todo/api/utils/jwt"
	"RestFullAPI-todo/pkg/auth"
	"github.com/gofiber/fiber/v2"
	"strings"
)

// Auth is the authentication middleware
func Auth(service auth.Service, optional ...bool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		h := c.Get("Authorization")

		if len(optional) > 0 {
			if h == "" {
				return c.Next()
			}

			// Split the header
			chunks := strings.Split(h, " ")

			// If header signature is not like `Bearer <token>`, then throw
			// This is also required, otherwise chunks[1] will throw out of bound error
			if len(chunks) < 2 {
				return c.Next()
			}

			// Verify the token which is in the chunks
			user, err := jwt.Verify(chunks[1])

			if err != nil {
				return c.Next()
			}

			if isActive := service.IsUserActiveByUsername(user.Username); !isActive {
				return c.Next()
			}

			c.Locals("UserId", user.ID)
			c.Locals("User", user.Username)

			return c.Next()
		}

		if h == "" {
			return utils.ErrUnauthorized
		}

		// Split the header
		chunks := strings.Split(h, " ")

		// If header signature is not like `Bearer <token>`, then throw
		// This is also required, otherwise chunks[1] will throw out of bound error
		if len(chunks) < 2 {
			return utils.ErrUnauthorized
		}

		// Verify the token which is in the chunks
		user, err := jwt.Verify(chunks[1])

		if err != nil {
			return utils.ErrUnauthorized
		}

		if isActive := service.IsUserActiveByUsername(user.Username); !isActive {
			return utils.ErrUnauthorized
		}

		c.Locals("UserId", user.ID)
		c.Locals("User", user.Username)

		return c.Next()
	}
}
