package repository

import (
	"encoding/json"
	"main/internal"
	"os"
)

func NewProductStore(db []internal.Product, lastId int) *ProductStore {
	product := &ProductStore{
		products: db,
		lastId:   lastId,
	}

	product.LoadProducts()
	product.lastId = len(product.products)
	return product
}

type ProductStore struct {
	products []internal.Product
	lastId   int
}

func (p *ProductStore) Save(product *internal.Product) (err error) {

	for _, v := range p.products {
		if v.CodeValue == product.CodeValue {
			return internal.ErrProductAlreadyExists
		}
	}
	(*p).lastId++
	(*product).ID = p.lastId + 1

	(*p).products = append(p.products, *product)

	return
}

func (p *ProductStore) LoadProducts() {

	file, err := os.ReadFile("products.json")
	if err != nil {
		panic(err)
	}

	json.Unmarshal(file, &(*p).products)
}

func (p *ProductStore) GetAll() (products []internal.Product, err error) {
	return p.products, nil
}

func (p *ProductStore) GetByID(id int) (product internal.Product, err error) {
	for _, product := range p.products {
		if product.ID == id {
			return product, nil
		}
	}
	return internal.Product{}, internal.ErrProductNotFound
}

func (p *ProductStore) GetByPriceGT(price float64) (products []internal.Product, err error) {
	for _, product := range p.products {
		if product.Price > price {
			products = append(products, product)
		}
	}
	if len(products) == 0 {
		err = internal.ErrProductNotFound
	}
	return
}

func (p *ProductStore) Update(product *internal.Product) (err error) {
	for i, v := range p.products {
		if v.ID == product.ID {
			p.products[i] = *product
			return nil
		}
	}
	return internal.ErrProductNotFound
}
