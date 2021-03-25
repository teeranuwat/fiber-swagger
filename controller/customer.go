package controller

import "github.com/gofiber/fiber/v2"

type CustomerInput struct {
	IDCard    string `json:"idcard" example:"11129292929292" format:"id card"`
	FirstName string `json:"firstname" example:"mark" format:"first name"`
	LastName  string `json:"lastname" example:"woodsome" format:"last name"`
	Email     string `json:"email" example:"mark.woodsome@gmail.com" format:"email"`
	Tel       string `json:"tel" example:"0877766988" format:"mobile no."`
	BirthDate string `json:"birthdate" example:"2000-01-02" format:"date"`
}

type CustomerData struct {
	ID        string `json:"id"`
	IDCard    string `json:"idcard"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Tel       string `json:"tel"`
	BirthDate string `json:"birthdate"`
	Age       int    `json:"age"`
}

// CreateCustomer godoc
// @Summary Add a customer
// @Description add by customer data json
// @Tags customer
// @Accept  json
// @Produce  json
// @Param account body CustomerInput true "Add customer"
// @Success 200 {object} CustomerData
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /create [post]
func CreateCustomer(c *fiber.Ctx) (err error) {
	customer := CustomerData{}
	return c.Status(fiber.StatusOK).JSON(customer)
}
