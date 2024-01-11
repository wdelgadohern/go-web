package service

import "main/internal"

func NewProductDefault(rp internal.ProductRepository) *ProductDefault {
	return &ProductDefault{
		rp: rp,
	}

}

type ProductDefault struct {
	rp internal.ProductRepository
}

func (p *ProductDefault) Save(product *internal.Product) (err error) {
	// Business logic
	// ...
	return p.rp.Save(product)
}

func (p *ProductDefault) GetAll() (products []internal.Product, err error) {
	return p.rp.GetAll()
}

func (p *ProductDefault) GetByID(id int) (product internal.Product, err error) {
	return p.rp.GetByID(id)
}

func (p *ProductDefault) GetByPriceGT(price float64) (products []internal.Product, err error) {
	return p.rp.GetByPriceGT(price)
}
