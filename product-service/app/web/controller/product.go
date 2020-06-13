package controller

import (
	"encoding/json"
	"net/http"

	service "github.com/fredoliveira-ca/products-golang-java/product-service/app/web/service"
)

// ListProducts is ..
func ListProducts(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()["user"]

	products := service.FetchAll(keys)

	jsonList, err := json.Marshal(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonList)
}
