package controllers

import (
	"app_store_api/models"
	utils "app_store_api/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Products struct {
	AppController
}

func (c Products) InitProductController(r *chi.Mux) {
	r.Route("/api/v1/products", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Get("/", c.GetAllProducts)
			r.Post("/create", c.CreateProduct)

			r.Route("/{uuid}", func(r chi.Router) {
				r.Get("/detail", c.GetProduct)
				// r.Put("/update")
				// r.Delete("/delete")
			})
		})
	})
}

func (c Products) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	result, total, err := c.product.ReadAllProducts(utils.Paginator(query))
	price, _ := c.product.GetMinMaxPrice()

	if err != nil {
		return
	}

	data := &models.Products{Total: total, Results: result, MaxPrice: price.MaxPrice, MinPrice: price.MinPrice}

	c.toJson(w, data)
}

func (c Products) GetProduct(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")
	data, err := c.product.ReadOneProduct(uuid)

	if err != nil {
		log.Printf("Error: %s", err)
		c.handleErr(w, http.StatusNotFound, fmt.Sprintf("%s", err))
		return
	}

	c.toJson(w, data)
}

func (c Products) CreateProduct(w http.ResponseWriter, r *http.Request) {
	data := new(models.Product)

	if err := json.NewDecoder(r.Body).Decode(&data); err == nil {
		data, err := c.product.CreateProduct(data)

		if err != nil {
			log.Printf("Error: %s", err)
			c.handleErr(w, http.StatusNotFound, fmt.Sprintf("%s", err))
			return
		}

		c.toJson(w, data)
		return
	}
}
