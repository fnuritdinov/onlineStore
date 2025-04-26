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
	UpdatedAt string  `json:"updated_at"`
}

type Service interface {
	AddProduct(p Product) ([]Product, error)
	UpdateProduct(p Product) (*Product, error)
	GetProducts() ([]Product, error)
}

type service struct {
}

func NewService() Service {
	return &service{}
}
