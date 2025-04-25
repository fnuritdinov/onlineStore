package product

import "time"

func (s *service) AddProduct(p Product) ([]Product, error) {
	p.ID = s.nextID
	p.CreatedAt = time.Now().Format("02-01-2006")
	s.nextID++
	s.products = append(s.products, p)
	return s.products, nil
}

func (s *service) GetProducts() ([]Product, error) {
	return s.products, nil
}
