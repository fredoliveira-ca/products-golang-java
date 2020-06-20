package repository

import (
	db "github.com/fredoliveira-ca/products-golang-java/product-service/app/config"
	grpc "github.com/fredoliveira-ca/products-golang-java/product-service/app/grpc/client"
	domain "github.com/fredoliveira-ca/products-golang-java/product-service/domain"
)

const (
	selectAllProducts    = "SELECT * FROM product"
	selectOneProductByID = "SELECT * FROM product WHERE product_id=$1"
)

// FindAll is ...
func FindAll(userID string) []domain.Product {
	db := db.ConnectDataBase()
	defer db.Close()

	records, err := db.Query(selectAllProducts)
	if err != nil {
		panic(err)
	}

	product := domain.Product{}
	products := []domain.Product{}

	for records.Next() {
		var id, title, description string
		var price int64

		err = records.Scan(&id, &price, &title, &description)
		if err != nil {
			panic(err.Error())
		}

		product.ID = id
		product.Title = title
		product.Description = description
		product.PriceInCents = price

		discount := grpc.CalculateDiscount(product.ID, userID)
		product.Discount = domain.Discount{Pct: 0, ValueInCents: 0}
		if discount != nil {
			product.Discount = domain.Discount{
				Pct:          discount.Pct,
				ValueInCents: discount.ValueInCents,
			}
		}

		products = append(products, product)
	}

	return products
}

// FindOne is ...
func FindOne(id string) domain.Product {
	db := db.ConnectDataBase()
	records, err := db.Query(selectOneProductByID, id)

	if err != nil {
		panic(err.Error())
	}

	product := domain.Product{}
	for records.Next() {
		var id, title, description string
		var price int64

		err = records.Scan(&id, &price, &title, &description)
		if err != nil {
			panic(err.Error())
		}

		product.ID = id
		product.Title = title
		product.Description = description
		product.PriceInCents = price
	}

	defer db.Close()
	return product
}
