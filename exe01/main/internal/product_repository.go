package internal

import "errors"

var (
	ErrProductNotFound      = errors.New("product not found")
	ErrProductAlreadyExists = errors.New("product already exists")
)

type ProductRepository interface {
	Save(product *Product) (error error)
	GetAll() (products []Product, err error)
	GetByID(id int) (product Product, err error)
	GetByPriceGT(price float64) (products []Product, err error)
}
