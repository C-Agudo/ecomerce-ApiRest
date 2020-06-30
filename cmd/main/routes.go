package main

import (
	"encoding/json"
	"net/http"

	"github.com/C-Agudo/ecomerce-ApiRest/products/web"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes(
	product *web.ProductHandler,
) *chi.Mux {
	mux := chi.NewMux()

	// globals middleware
	mux.Use(
		middleware.Logger,    //log every http request
		middleware.Recoverer, // recover if a panic occurs
	)

	mux.Get("/hello", helloHandler)
	// mux.Post("/product", product.SaveProductHandler)
	// mux.Delete("/product/{productId}", product.DeleteProductHandler)

	mux.Route("/product", func(r chi.Router) {
		r.Get("/", product.GetProductsHandler)
		r.Post("/", product.SaveProductHandler)
		r.Delete("/{productId:[0-9]+}", product.DeleteProductHandler)
		r.Patch("/{productId:[0-9]+}", product.UpdateProductHandler)

		// r.Patch("/{cityID:[0-9]+}", city.UpdateCityHandler)
	})

	return mux
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res := map[string]interface{}{"message": "hello world"}

	_ = json.NewEncoder(w).Encode(res)
}
