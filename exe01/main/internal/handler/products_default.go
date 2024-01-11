package handler

import (
	"encoding/json"
	"main/internal"
	"main/platform/web/request"
	"main/platform/web/response"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func NewDefaultProduct(sv internal.ProductService) *ProductDefault {
	return &ProductDefault{
		sv: sv,
	}

}

type ProductDefault struct {
	sv internal.ProductService
}

type ProductJSON struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type BodyRequestProductJSON struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func (p *ProductDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body BodyRequestProductJSON
		if err := request.JSON(r, &body); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
			return
		}

		product := internal.Product{
			Name:        body.Name,
			Quantity:    body.Quantity,
			CodeValue:   body.CodeValue,
			IsPublished: body.IsPublished,
			Expiration:  body.Expiration,
			Price:       body.Price,
		}

		if err := p.sv.Save(&product); err != nil {
			response.Text(w, http.StatusInternalServerError, "internal server error")
			return
		}

		data := ProductJSON{
			ID:          product.ID,
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration,
			Price:       product.Price,
		}
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "product created successfully",
			"data":    data,
		})

	}
}

func (p *ProductDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, _ := p.sv.GetAll()
		productBytes, _ := json.Marshal(products)
		w.Write(productBytes)
	}
}

func (p *ProductDefault) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		product, err := p.sv.GetByID(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		productBytes, _ := json.Marshal(product)
		w.Write(productBytes)
	}
}

func (p *ProductDefault) GetByPriceGT() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		price, _ := strconv.ParseFloat(r.URL.Query().Get("priceGT"), 64)
		myProducts, err := p.sv.GetByPriceGT(price)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		productBytes, _ := json.Marshal(myProducts)
		w.Write(productBytes)

	}
}

func (p *ProductDefault) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		product, err := p.sv.GetByID(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		var body BodyRequestProductJSON
		if err := request.JSON(r, &body); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
			return
		}
		product.Name = body.Name
		product.Quantity = body.Quantity
		product.CodeValue = body.CodeValue
		product.IsPublished = body.IsPublished
		product.Expiration = body.Expiration
		product.Price = body.Price

		if err := p.sv.Update(&product); err != nil {
			response.Text(w, http.StatusInternalServerError, "internal server error")
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "movie updated",
			"data":    product,
		})
	}
}
