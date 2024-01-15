package handler_test

import (
	"context"
	"fmt"
	"main/internal"
	"main/internal/auth"
	"main/internal/handler"
	"main/internal/repository"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/require"
)

func TestHandlerProducts_Get_Handler(t *testing.T) {
	t.Run("Success to get a list of products", func(t *testing.T) {

		// Arrange
		// p.au.Auth("1234")
		at := auth.NewAuthTokenMock()
		at.FuncAuth = func(token string) (err error) {
			return
		}
		rp := repository.NewRepositoryMock()
		rp.FuncGet = func() (p []internal.Product, err error) {
			p = []internal.Product{
				*internal.NewProduct(1, "Product 1", 10, "1234", true, "2021-12-31", 10),
			}
			return
		}
		hd := handler.NewDefaultProduct(rp, at)
		hdFunc := hd.GetAll()

		// Act
		req := &http.Request{}
		res := httptest.NewRecorder()
		hdFunc(res, req)

		// Assert
		expectedCode := http.StatusOK
		expectedBody := `[{"id":1,"name":"Product 1","quantity":10,"code_value":"1234","is_published":true,"expiration":"2021-12-31","price":10}]`
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedCode, res.Code)

	})
}

func TestHandlerProducts_GetById_Handler(t *testing.T) {

	t.Run("Success to get a product by id", func(t *testing.T) {

		// Arrange
		at := auth.NewAuthTokenMock()
		at.FuncAuth = func(token string) (err error) {
			return
		}

		rp := repository.NewRepositoryMock()
		rp.FuncGetById = func(id int) (p internal.Product, err error) {
			p = internal.Product{
				ID:          1,
				Name:        "Product 1",
				Quantity:    10,
				CodeValue:   "1234",
				IsPublished: true,
				Expiration:  "2021-12-31",
				Price:       10,
			}
			return
		}
		hd := handler.NewDefaultProduct(rp, at)
		hdFunc := hd.GetByID()
		// Act
		req := &http.Request{}
		chiCtx := chi.NewRouteContext() // *chi.Context to handle params
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx)) // replace *http.Request with the new request having the updated context
		res := httptest.NewRecorder()
		hdFunc(res, req)

		// Assert
		expectedCode := http.StatusOK
		expectedBody := `{"id":1,"name":"Product 1","quantity":10,"code_value":"1234","is_published":true,"expiration":"2021-12-31","price":10}`

		fmt.Println("expectedBody: ", expectedBody)
		fmt.Println("res.Body.String(): ", res.Body.String())
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedCode, res.Code)

	})

}

func TestHandlerProducts_Post_Handler(t *testing.T) {

	t.Run("Success to post a product", func(t *testing.T) {
		// Arrange
		at := auth.NewAuthTokenMock()
		at.FuncAuth = func(token string) (err error) {
			return
		}

		rp := repository.NewRepositoryMock()
		rp.FuncSave = func(product *internal.Product) (err error) {
			product.ID = 1
			return
		}

		// Act
		hd := handler.NewDefaultProduct(rp, at)
		hdFunc := hd.Create()
		// req := &http.Request{
		// 	Body: io.NopCloser(strings.NewReader(`{"name":"Product 1","quantity":10,"code_value":"1234","is_published":true,"expiration":"2021-12-31","price":10}`)),
		// }
		req := httptest.NewRequest("POST", "/products", strings.NewReader(`{"name":"Product 1","quantity":10,"code_value":"1234","is_published":true,"expiration":"2021-12-31","price":10}`))
		req.Header.Set("Content-Type", "application/json") // Set the Content-Type header for JSON

		res := httptest.NewRecorder()
		hdFunc(res, req)

		// Assert
		expectedCode := http.StatusCreated
		expectedBody := `{"data":{"id":1,"name":"Product 1","quantity":10,"code_value":"1234","is_published":true,"expiration":"2021-12-31","price":10},"message":"product created successfully"}`

		fmt.Println("expectedBody: ", expectedBody)
		fmt.Println("res.Body.String(): ", res.Body.String())
		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())

	})

}
