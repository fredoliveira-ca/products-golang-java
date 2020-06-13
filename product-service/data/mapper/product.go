package mapper

import (
	"github.com/fredoliveira-ca/products-golang-java/product-service/app/web/dto"
	"github.com/fredoliveira-ca/products-golang-java/product-service/domain"
)

// Of is ...
func Of(product domain.Product) dto.Product {
	return dto.Product{
		ID:           product.ID,
		PriceInCents: product.PriceInCents,
		Description:  product.Description,
		Title:        product.Title,
	}
}
