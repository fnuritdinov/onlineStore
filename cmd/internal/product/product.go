package product

import (
	"github.com/gofiber/fiber/v3"
	"time"
)

var products []Product
var productID = 0

func (s *service) AddProduct(p Product) ([]Product, error) {
	productID++
	p.ID = productID
	p.CreatedAt = time.Now().Format("02-01-2006 15:04:05")
	p.UpdatedAt = time.Now().Format("02-01-2006 15:04:05")
	products = append(products, p)
	return products, nil
}

func (s *service) UpdateProduct(p Product) (*Product, error) {
	p.UpdatedAt = time.Now().Format("02-01-2006 15:04:05")
	for i, product := range products {
		if product.ID == p.ID {
			product.UpdatedAt = time.Now().Format("02-01-2006 15:04:05")
			product.Name = p.Name
			product.Price = p.Price
			product.Currency = p.Currency
			product.Weight = p.Weight
			product.Category = p.Category

			products[i] = product

			return &product, nil
		}
	}

	return nil, fiber.ErrNotFound
}

func (s *service) GetProducts() ([]Product, error) {
	return products, nil
}
