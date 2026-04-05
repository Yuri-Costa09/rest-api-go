package model

import (
	"errors"

	"github.com/google/uuid"
)

type Product struct {
	ID          string
	Name        string
	Price       float32
	Description *string // optional
}

var (
	ErrInvalidPrice = errors.New("price should be greater than 0")
	ErrInvalidName  = errors.New("name is invalid")
)

func NewProduct(price float32, name string, description *string) (*Product, error) {
	if price <= 0 {
		return nil, ErrInvalidPrice
	}
	if len(name) < 3 {
		return nil, ErrInvalidName
	}

	return &Product{
		ID:          uuid.NewString(),
		Name:        name,
		Price:       price,
		Description: description,
	}, nil
}
