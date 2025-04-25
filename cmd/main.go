package main

import (
	"github.com/gofiber/fiber/v3"
	productHandler "onlineStore/cmd/handler/product"
	productService "onlineStore/cmd/internal/product"
)

func main() {
	app := fiber.New()

	service := productService.NewService()
	handler := productHandler.New(service)

	app.Post("/products", handler.AddProduct)
	app.Get("/products", handler.GetProducts)

	app.Listen(":3000")
}
