package customer

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// CreateCustomerHandler godoc
// @Summary Create customer data
// @Description Create customer
// @Tags createcustomer
// @Accept  json
// @Produce  json
// @Param customer body CustomerInput true "customer data"
// @Success 200 {object} CustomerInput
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /create [post]
func (srv *Server) CreateCustomerHandler(c *fiber.Ctx) (err error) {
	fmt.Println("you are here")
	c.Status(fiber.StatusOK).JSON("OK")
	return
}

func (srv *Server) Test(c *fiber.Ctx) (err error) {
	return c.Status(fiber.StatusOK).JSON("OK")
}
