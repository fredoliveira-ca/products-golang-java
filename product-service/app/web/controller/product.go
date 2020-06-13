package controller

import (
	"net/http"

	"github.com/fredoliveira-ca/products-golang-java/product-service/app/web/service"
)

// ListProducts is ..
func ListProducts(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()["user"]
	
	list, err := service.FetchAll(keys)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(list)
}
