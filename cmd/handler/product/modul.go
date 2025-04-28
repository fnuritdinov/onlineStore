package product

import (
	"github.com/gofiber/fiber/v3"
	"onlineStore/cmd/internal/product"
)

type Handler interface {
	AddProduct(c fiber.Ctx) error
	UpdateProduct(c fiber.Ctx) error
	GetProducts(c fiber.Ctx) error
	GetProductByID(c fiber.Ctx) error
	DeleteProduct(c fiber.Ctx) error
}

type handler struct {
	service product.Service
}

func New(service product.Service) Handler {
	return &handler{service: service}
}
