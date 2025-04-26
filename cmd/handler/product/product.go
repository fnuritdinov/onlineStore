package product

import (
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"onlineStore/cmd/internal/product"
	"strconv"
)

func (h *handler) AddProduct(c fiber.Ctx) error {
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

func (h *handler) UpdateProduct(c fiber.Ctx) error {
	var p product.Product

	err := json.Unmarshal(c.Body(), &p)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	p.ID, err = strconv.Atoi(c.Params("product_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid product id"})
	}

	pr, err := h.service.UpdateProduct(p)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(pr)
}

func (h *handler) GetProducts(c fiber.Ctx) error {
	products, err := h.service.GetProducts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(products)
}
