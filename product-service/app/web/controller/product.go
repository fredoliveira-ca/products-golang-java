package controller

import (
	"encoding/json"
	"log"
	"net/http"

	repository "github.com/fredoliveira-ca/products-golang-java/product-service/data/repository"
)

// ListProducts is ..
func ListProducts(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()["user"]

	var userID string
	if len(params) > 0 {
		userID = params[0]
		log.Println("Listing products for the user: " + string(userID))
	}

	products := repository.FindAll(userID)

	jsonList, err := json.Marshal(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonList)
}
