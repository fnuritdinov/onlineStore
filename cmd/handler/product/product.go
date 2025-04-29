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
	var page = c.Query("page")
	var perPage = c.Query("per_page")

	products, err := h.service.GetProducts(page, perPage)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(products)
}

func (h *handler) GetProductByID(c fiber.Ctx) error { // 👈 новый метод
	idParam := c.Params("product_id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Ошибка": "неверный id продукта"})
	}

	product, err := h.service.GetProductByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"ошибка": "продукт не найден"})
	}

	return c.JSON(product)
}

func (h *handler) DeleteProduct(c fiber.Ctx) error {
	idParam := c.Params("product_id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"ошибка": "продукт не найден"})
	}

	products, err := h.service.DeleteProduct(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(products)
}
