package product

import (
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"onlineStore/cmd/internal/product"
)

func (h *Handler) AddProduct(c fiber.Ctx) error {
	var p product.Product
	err := json.Unmarshal(c.Body(), &p)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	products, err := h.service.AddProduct(p)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(products)
}

func (h *Handler) GetProducts(c fiber.Ctx) error {
	products, err := h.service.GetProducts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(products)
}
