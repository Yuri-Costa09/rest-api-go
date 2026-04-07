package service

import (
	"errors"

	"github.com/Yuri-Costa09/rest-api-go/internal/repository"
)

type ProductService struct {
	repo repository.ProductRepository
}

var (
	ErrNotFound = errors.New("product not found")
)

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{repo}
}

func (p *ProductService) GetAll() error {

}
