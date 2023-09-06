package application

import (
	"errors"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Status string

const (
	Disabled Status = "disabled"
	Enabled  Status = "enabled"
)

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetId() string
	GetName() string
	GetStatus() Status
	GetPrice() float64
}

type Product struct {
	Id     string  `valid:"uuidv4"`
	Name   string  `valid:"required"`
	Status string  `valid:"required"`
	Price  float64 `valid:"float,optional"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewProduct() *Product {
	product := Product{
		Id:     uuid.NewV4().String(),
		Status: string(Enabled),
	}

	return &product
}

func (p *Product) IsValid() (bool, error) {
	if p.Price < 0 {
		return false, errors.New("the price must be greather than zero")
	}

	_, err := govalidator.ValidateStruct(p)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = string(Enabled)
		return nil
	}

	return errors.New("the price must be greather than zero to enable the product")
}

func (p *Product) Disable() error {
	if p.Price == 0 {
		p.Status = string(Disabled)
		return nil
	}

	return errors.New("the price must be zero to disable the product")
}

func (p *Product) GetId() string {
	return p.Id
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() Status {
	return Status(p.Status)
}

func (p *Product) GetPrice() float64 {
	return p.Price
}
