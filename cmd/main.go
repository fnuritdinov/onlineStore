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
	app.Put("/products/:product_id", handler.UpdateProduct)
	app.Get("/products", handler.GetProducts)
	app.Get("/products/:product_id", handler.GetProductByID)
	app.Delete("/products/:product_id", handler.DeleteProduct)

	app.Listen(":3000")
}
