package utils

import "github.com/gofiber/fiber/v2"

// ParseBody is helper function for parsing the body.
// Is any error occurs it will panic.
// Its just a helper function to avoid writing if condition again n again.
func ParseBody(ctx *fiber.Ctx, body interface{}) error {
	if err := ctx.BodyParser(body); err != nil {
		return NewError(fiber.StatusBadRequest, err.Error())
	}
	return nil
}

// ParseBodyAndValidate is helper function for parsing the body.
// Is any error occurs it will panic.
// Its just a helper function to avoid writing if condition again n again.
func ParseBodyAndValidate(ctx *fiber.Ctx, body interface{}) error {
	if err := ParseBody(ctx, body); err != nil {
		return err
	}

	return Validate(body)
}

// GetUserID is helper function for getting authenticated user's id
func GetUserID(c *fiber.Ctx) string {
	id, _ := c.Locals("UserId").(string)
	return id
}

// GetUsername is helper function for getting authenticated user's id
func GetUsername(c *fiber.Ctx) string {
	id, _ := c.Locals("User").(string)
	return id
}
