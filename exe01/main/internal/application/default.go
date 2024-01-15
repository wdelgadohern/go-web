package application

import (
	"main/internal"
	"main/internal/auth"
	"main/internal/handler"
	"main/internal/middleware"
	"main/internal/repository"
	"main/internal/service"
	"net/http"

	"github.com/go-chi/chi"
)

func NewDefaultHTTP(addr, token string) *DefaultHTTP {
	return &DefaultHTTP{
		addr:  addr,
		token: token,
	}
}

type DefaultHTTP struct {
	// addr is the address of the http server
	addr  string
	token string
}

func (h *DefaultHTTP) Run() (err error) {

	au := auth.NewAuthBasic(h.token)

	auth := middleware.NewAuthenticator(h.token)
	logger := middleware.NewLogger()

	rp := repository.NewProductStore(make([]internal.Product, 0), 0)
	sv := service.NewProductDefault(rp)
	hd := handler.NewDefaultProduct(sv, au)
	rt := chi.NewRouter()
	rt.Use(logger.Log, auth.Auth)
	rt.Post("/products", hd.Create())
	rt.Get("/products", hd.GetAll())
	rt.Get("/products/{id}", hd.GetByID())
	rt.Get("/products/search", hd.GetByPriceGT())
	rt.Put("/products/{id}", hd.Update())
	rt.Delete("/products/{id}", hd.Delete())
	err = http.ListenAndServe(h.addr, rt)
	return
}
