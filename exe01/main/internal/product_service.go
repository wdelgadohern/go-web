package internal

type ProductService interface {
	Save(product *Product) (err error)
	GetAll() (products []Product, err error)
	GetByID(id int) (product Product, err error)
	GetByPriceGT(price float64) (products []Product, err error)
	Update(product *Product) (err error)
	Delete(id int) (err error)
}
