package routers

import (
	"app_store_api/controllers"
	"app_store_api/middlewares"

	"github.com/go-chi/chi/v5"
)

var (
	m                   middlewares.Middleware
	appController       = &controllers.AppController{}
	productControllers  = &controllers.Products{}
	categoryControllers = &controllers.Categories{}
)

func InitRouter() (r *chi.Mux) {
	r = chi.NewRouter()

	m.UseMiddlewares(r)

	appController.NotFound(r)

	// @router /products
	productControllers.InitProductController(r)

	// @router /categories
	categoryControllers.InitCategoryController(r)

	return
}
