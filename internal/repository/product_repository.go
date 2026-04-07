package repository

import "github.com/Yuri-Costa09/rest-api-go/internal/model"

type ProductRepository interface {
	FindProductByID(id string) (p *model.Product, err error)
	FindAllProducts() (p []model.Product, err error)
	CreateProduct(p *model.Product) error
	UpdateProduct(p *model.Product) error
	DeleteProductById(id string) error
}
