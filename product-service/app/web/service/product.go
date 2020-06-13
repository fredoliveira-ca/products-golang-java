package service

import (
	"log"

	"github.com/fredoliveira-ca/products-golang-java/product-service/app/web/dto"
	repository "github.com/fredoliveira-ca/products-golang-java/product-service/data/repository"
)

// FetchAll is ...
func FetchAll(params []string) []dto.Product {
	var userID string

	if len(params) > 0 {
		userID = params[0]
		log.Println("Listing products for the user: " + string(userID))
	}

	return repository.FindAll(userID)
}
