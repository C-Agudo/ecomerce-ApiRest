package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/C-Agudo/ecomerce-ApiRest/internal/database"
	"github.com/C-Agudo/ecomerce-ApiRest/internal/logs"
	"github.com/C-Agudo/ecomerce-ApiRest/products/domain"
	"github.com/C-Agudo/ecomerce-ApiRest/products/gateway"
	"github.com/go-chi/chi"
)

type ProductHandler struct {
	gateway.ProductGateway
}

func (h *ProductHandler) SaveProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")

	cmd := parseRequest(r)
	res, err := h.Create(cmd)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		m := map[string]interface{}{"msg": "error in create product"}
		_ = json.NewEncoder(w).Encode(m)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(res)
}

func NewCreateProductHandler(db *database.MySqlClient) *ProductHandler {
	return &ProductHandler{gateway.NewProductCreateGateway(db)}
}

func parseRequest(r *http.Request) *domain.CreateProductCMD {
	body := r.Body
	defer body.Close()
	var cmd domain.CreateProductCMD
	_ = json.NewDecoder(body).Decode(&cmd)
	return &cmd
}

func (h *ProductHandler) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// productID, err := strconv.ParseInt(chi.URLParam(r, "productId"), 10, 32)
	productID := chi.URLParam(r, "productId")
	// var p = reflect.TypeOf(productID)
	fmt.Println(reflect.TypeOf(productID))
	productId, err := strconv.Atoi(productID)
	fmt.Println(reflect.TypeOf(productId))

	// var n int64 = productId

	// productIdInt, err := strconv.ParseInt(productID, 10, 64)
	productIdInt64 := int64(productId)
	fmt.Println(reflect.TypeOf(productIdInt64))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "cannot parse parameters"})
		return
	}

	res := h.Delete(productIdInt64)

	if res == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "cannot delete product"})
		return
	}

	json.NewEncoder(w).Encode(&res)
}

func (h *ProductHandler) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	productID, err := strconv.ParseInt(chi.URLParam(r, "productId"), 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "cannot parse parameters"})
		return
	}

	cmd, err := parseUpdateRequest(r, productID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "cannot parse parameters"})
		return
	}

	res := h.Update(cmd)

	json.NewEncoder(w).Encode(&res)
}

func parseUpdateRequest(r *http.Request, productID int64) (*domain.UpdateProductCMD, error) {
	body := r.Body
	defer body.Close()
	var cmd domain.UpdateProductCMD
	err := json.NewDecoder(body).Decode(&cmd)

	if err != nil {
		logs.Log().Error(err.Error())
		return nil, err
	}
	cmd.Id = productID
	return &cmd, nil
}

func (h *ProductHandler) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res := h.List()

	if res == nil || len(res) == 0 {
		res = []*domain.Product{}
	}

	json.NewEncoder(w).Encode(&res)

}
