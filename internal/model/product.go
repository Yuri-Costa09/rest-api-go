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
	// if price <= 0 {
	// 	return nil, ErrInvalidPrice
	// }
	// if len(name) < 3 {
	// 	return nil, ErrInvalidName
	// }

	p := &Product{
		ID:          uuid.NewString(),
		Name:        name,
		Price:       price,
		Description: description,
	}

	if err := p.validate(); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Product) validate() error {
	if p.Price <= 0 {
		return ErrInvalidPrice
	}
	if len(p.Name) < 3 {
		return ErrInvalidName
	}

	return nil
}
