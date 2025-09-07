package main

import (
	"fmt"
	"example/rest-api/internal/api"
	"example/rest-api/internal/config"
	"example/rest-api/internal/connection"
	"example/rest-api/internal/repository"
	"example/rest-api/internal/service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()
	dbConnection := connection.GetDatabase(cnf.Database)

	app := fiber.New()

	customerRepository := repository.NewCustomer(dbConnection)

	customerService := service.NewCustomer(customerRepository)

	api.NewCustomer(app, customerService)
	
	addr := cnf.Server.Host + ":" + cnf.Server.Port
    fmt.Println("🚀 Server running at http://" + addr)

    if err := app.Listen(addr); err != nil {
        panic(err)
    }
}
