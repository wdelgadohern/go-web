package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

var products []Product

func main() {

	LoadProducts()

	r := chi.NewRouter()

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
		w.WriteHeader(http.StatusOK)
	})

	r.Get("/products", func(w http.ResponseWriter, r *http.Request) {
		productBytes, _ := json.Marshal(products)
		w.Write(productBytes)
	})

	r.Get("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		product, err := GetProductById(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		productBytes, _ := json.Marshal(product)
		w.Write(productBytes)
	})

	r.Get("/products/search", func(w http.ResponseWriter, r *http.Request) {
		price, _ := strconv.ParseFloat(r.URL.Query().Get("priceGT"), 64)
		myProducts, err := GetProductsByPrice(price)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		productBytes, _ := json.Marshal(myProducts)
		w.Write(productBytes)
	})

	r.Post("/products", func(w http.ResponseWriter, r *http.Request) {
		var product Product
		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		ValidateNewProduct(&product)
		w.WriteHeader(http.StatusCreated)
	})

	http.ListenAndServe(":8081", r)

}

func GetProductById(id int) (Product, error) {
	for _, product := range products {
		if product.ID == id {
			return product, nil
		}
	}
	return Product{}, errors.New("product not found")
}

func GetProductsByPrice(price float64) (myProducts []Product, err error) {
	for _, product := range products {
		if product.Price > price {
			myProducts = append(myProducts, product)
		}
	}
	if len(myProducts) == 0 {
		err = errors.New("product not found")
	}
	return
}

func LoadProducts() {
	file, err := os.ReadFile("products.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(file, &products)
}

func ValidateNewProduct(product *Product) (err error) {
	product.ID = (len(products) + 1)
	products = append(products, *product)

	for _, v := range products {
		if v.CodeValue == product.CodeValue {
			err = errors.New("product already exists")
		}
	}
	return

}
