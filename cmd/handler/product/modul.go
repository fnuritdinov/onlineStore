package product

import "onlineStore/cmd/internal/product"

type Handler struct {
	service product.Service
}

func New(service product.Service) *Handler {
	return &Handler{service: service}
}
