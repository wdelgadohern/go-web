package application

import (
	"main/internal"
	"main/internal/handler"
	"main/internal/repository"
	"main/internal/service"
	"net/http"

	"github.com/go-chi/chi"
)

func NewDefaultHTTP(addr string) *DefaultHTTP {
	return &DefaultHTTP{
		addr: addr,
	}
}

type DefaultHTTP struct {
	// addr is the address of the http server
	addr string
}

func (h *DefaultHTTP) Run() (err error) {

	rp := repository.NewProductStore(make([]internal.Product, 0), 0)
	sv := service.NewProductDefault(rp)
	hd := handler.NewDefaultProduct(sv)
	rt := chi.NewRouter()
	rt.Post("/products", hd.Create())
	rt.Get("/products", hd.GetAll())
	rt.Get("/products/{id}", hd.GetByID())
	rt.Get("/products/search", hd.GetByPriceGT())
	rt.Put("/products/{id}", hd.Update())
	err = http.ListenAndServe(h.addr, rt)
	return
}
