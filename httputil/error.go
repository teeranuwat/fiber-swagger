package httputil

import (
	"github.com/gofiber/fiber/v2"
)

// NewError example
func NewError(ctx *fiber.Ctx, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.Status(er.Code).JSON(er)
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}