package product

import "time"

var products []Product
var productID = 0

func (s *service) AddProduct(p Product) ([]Product, error) {
	p.ID = productID + 1
	p.CreatedAt = time.Now().Format("02-01-2006")
	products = append(products, p)
	return products, nil
}

func (s *service) GetProducts() ([]Product, error) {
	return products, nil
}
