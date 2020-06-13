package repository

import (
	db "github.com/fredoliveira-ca/products-golang-java/product-service/app/config"
	grpc "github.com/fredoliveira-ca/products-golang-java/product-service/app/grpc/client"
	dto "github.com/fredoliveira-ca/products-golang-java/product-service/app/web/dto"
	mapper "github.com/fredoliveira-ca/products-golang-java/product-service/data/mapper"
	domain "github.com/fredoliveira-ca/products-golang-java/product-service/domain"
)

// FindAll is ...
func FindAll(userID string) []dto.Product {
	db := db.ConnectDataBase()
	defer db.Close()

	records, err := db.Query("SELECT * FROM product")
	if err != nil {
		panic(err)
	}

	product := domain.Product{}
	products := []dto.Product{}

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

		discount := grpc.CalculateDicount(product.ID, userID)
		productDTO := mapper.Of(product)
		productDTO.Discount = dto.Discount{Pct: 0, ValueInCents: 0}
		if discount != nil {
			productDTO.Discount = dto.Discount{
				Pct:          discount.Pct,
				ValueInCents: discount.ValueInCents,
			}
		}

		products = append(products, productDTO)
	}
	
	return products
}

// FindOne is ...
func FindOne(id string) dto.Product {
	db := db.ConnectDataBase()
	sql := "SELECT * FROM product WHERE product_id=$1;"
	records, err := db.Query(sql, id)

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
	return mapper.Of(product)
}
