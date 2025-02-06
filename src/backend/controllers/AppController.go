package controllers

import (
	"app_store_api/models"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type (
	AppController struct {
		product    *models.Product
		categories *models.Category
	}

	ErrorObject struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
)

func (c AppController) NotFound(r *chi.Mux) {
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		c.handleErr(w, http.StatusNotFound, "Route not found...")
	})
}

func (c AppController) toJson(w http.ResponseWriter, b interface{}) {
	jsonMarshal, _ := json.MarshalIndent(b, "", "  ")

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonMarshal))
}

func (c AppController) handleErr(w http.ResponseWriter, code int, message string) {
	errObj := ErrorObject{
		Code:    code,
		Message: message,
	}

	jsonMarshal, _ := json.MarshalIndent(errObj, "", "  ")

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(errObj.Code)

	w.Write([]byte(jsonMarshal))
}
