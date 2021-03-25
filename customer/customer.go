package customer

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"

	fiberSwagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/bxcodec/faker/v3"
	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
	// "github.com/gofiber/fiber/v2/middleware/requestid"
	// "github.com/google/uuid"
)

type Server struct {
	App *fiber.App
}

var server Server

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

func CreateCustomer() (err error) {
	return
}

func createMockTestData() (err error) {
	// f, err := os.OpenFile("/Users/teeranuwatueaprasert/Me/workspace/unit-test/customer.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	f, err := os.Create("/Users/teeranuwatueaprasert/Me/workspace/unit-test/customer.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	//randomData()
	_json, err := json.Marshal(randomData())
	_, err = f.Write(_json)
	if err != nil {
		log.Fatal(err)
	}
	// flush in-memory copy
	err = f.Sync()
	if err != nil {
		log.Fatal(err)
	}
	return
}

func randomData() (customers []CustomerInput) {
	num := rand.Intn(10) + 1
	fmt.Println("random with :", num)
	for i := 0; i < num; i++ {
		c := CustomerInput{
			IDCard:    faker.CCNumber(),
			FirstName: faker.FirstName(),
			LastName:  faker.LastName(),
			Email:     faker.Email(),
			Tel:       faker.E164PhoneNumber(),
			BirthDate: faker.Date(),
		}
		customers = append(customers, c)
	}
	return
}

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3500
// @BasePath /api/v1
func StartServer() (server Server) {
	app := fiber.New()

	// app.Use(requestid.New(requestid.Config{
	// 	Header: fiber.HeaderXRequestID,
	// 	Generator: func() string {
	// 		return uuid.New().String()
	// 	},
	// }))

	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins: "http://localhost, https://acb.co.th",
	// 	AllowHeaders: "x-token , Content-Type, Accept, Origin, X-Request-ID",
	// }))

	g := app.Group("/api/v1")
	g.Post("/create", server.CreateCustomerHandler)
	g.Get("/accounts/:id", ShowAccount)
	g.Get("/test", server.Test)
	app.Get("/swagger/*any", fiberSwagger.Handler)
	server.App = app

	app.Listen(":3500")

	return
}




// ShowAccount godoc
// @Summary Show a account
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} Account
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /accounts/{id} [get]
func ShowAccount(c *fiber.Ctx) error {
	return c.JSON(Account{
		Id: c.Params("id"),
	})
}

type Account struct {
	Id string
}
