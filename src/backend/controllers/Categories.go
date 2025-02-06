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

type Categories struct {
	AppController
}

func (c Categories) InitCategoryController(r *chi.Mux) {
	r.Route("/api/v1/categories", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Get("/", c.GetAllCategories)
			r.Post("/create", c.CreateCategory)

			r.Route("/{uuid}", func(r chi.Router) {
				r.Get("/detail", c.GetCategory)
				// r.Put("/update")
				// r.Delete("/delete")
			})
		})
	})
}

func (c Categories) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	result, total, err := c.categories.ReadAllCategories(utils.Paginator(query))

	if err != nil {
		return
	}

	data := &models.Categories{Total: total, Results: result}

	c.toJson(w, data)
}

func (c Categories) GetCategory(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")
	data, err := c.categories.ReadOneCategory(uuid)

	if err != nil {
		log.Printf("Error: %s", err)
		c.handleErr(w, http.StatusNotFound, fmt.Sprintf("%s", err))
		return
	}

	c.toJson(w, data)
}

func (c Categories) CreateCategory(w http.ResponseWriter, r *http.Request) {
	data := new(models.Category)

	if err := json.NewDecoder(r.Body).Decode(&data); err == nil {
		data, err := c.categories.CreateCategory(data)

		if err != nil {
			log.Printf("Error: %s", err)
			c.handleErr(w, http.StatusNotFound, fmt.Sprintf("%s", err))
			return
		}

		c.toJson(w, data)
		return
	}
}
