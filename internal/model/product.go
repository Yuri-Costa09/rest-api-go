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

	if err := validateName(p.Name); err != nil {
		return err
	}

	if err := validatePrice(p.Price); err != nil {
		return err
	}

	return nil
}

func validateName(name string) error {
	if len(name) < 3 {
		return ErrInvalidName
	}

	return nil
}

func validatePrice(price float32) error {
	if price <= 0 {
		return ErrInvalidPrice
	}

	return nil
}

func (p *Product) UpdateName(name string) error {
	if err := validateName(name); err != nil {
		return err
	}

	p.Name = name

	return nil
}

func (p *Product) UpdatePrice(price float32) error {
	if err := validatePrice(price); err != nil {
		return err
	}

	p.Price = price

	return nil
}

func (p *Product) UpdateDesc(desc *string) error {

	p.Description = desc

	return nil
}
