package repository

import "main/internal"

type ProductRepositoryMock struct {
	FuncGet     func() (p []internal.Product, err error)
	FuncGetById func(id int) (p internal.Product, err error)
	FuncSave    func(product *internal.Product) (err error)
}

func NewRepositoryMock() *ProductRepositoryMock {
	return &ProductRepositoryMock{}
}

func (s *ProductRepositoryMock) GetAll() (p []internal.Product, err error) {
	p, err = s.FuncGet()
	return
}

func (s *ProductRepositoryMock) Save(product *internal.Product) (err error) {
	err = s.FuncSave(product)
	return
}

func (s *ProductRepositoryMock) GetByID(id int) (p internal.Product, err error) {
	p, err = s.FuncGetById(id)
	return
}

func (s *ProductRepositoryMock) GetByPriceGT(price float64) (p []internal.Product, err error) {
	p, err = s.FuncGet()
	return
}

func (s *ProductRepositoryMock) Update(product *internal.Product) (err error) {
	err = s.FuncSave(product)
	return
}
