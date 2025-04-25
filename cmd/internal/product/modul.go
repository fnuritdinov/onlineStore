package product

type Product struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Currency  string  `json:"currency"`
	Count     int     `json:"count"`
	Weight    string  `json:"weight"`
	Category  string  `json:"category"`
	CreatedAt string  `json:"created_at"`
}

type Service interface {
	AddProduct(p Product) ([]Product, error)
	GetProducts() ([]Product, error)
}

type service struct {
	products []Product
	nextID   int
}

func NewService() Service {
	return &service{products: []Product{}, nextID: 1}
}
