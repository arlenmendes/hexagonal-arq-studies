package application

import "errors"

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	ID     string
	Name   string
	Status string
	Price  float64
}

func (p *Product) IsValid() (bool, error) {
	if p.Name == "" {
		return false, errors.New("product name is required")
	}

	if p.Price <= 0 {
		return false, errors.New("product price is required")
	}

	return true, nil
}

func (p *Product) Enable() error {
	if p.Price <= 0 {
		return errors.New("The price must be greater than zero to enable the product")
	}
	p.Status = ENABLED
	return nil
}

func (p *Product) Disable() error {
	p.Status = DISABLED
	return nil
}

func (p *Product) GetID() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}
