package repository

import (
	"time"

	db "github.com/fredoliveira-ca/products-golang-java/user-service/app/config"
	"github.com/fredoliveira-ca/products-golang-java/user-service/domain"
)

//FindOne is...
func FindOne(id string) domain.User {
	db := db.ConnectDataBase()
	sql := "SELECT * FROM \"user\" WHERE user_id=$1;"
	selection, err := db.Query(sql, id)
	if err != nil {
		panic(err.Error())
	}

	user := domain.User{}
	for selection.Next() {
		var id, firstName, lastName string
		var birthday time.Time

		err = selection.Scan(&id, &firstName, &lastName, &birthday)
		if err != nil {
			panic(err.Error())
		}

		user.ID = id
		user.FirstName = firstName
		user.LastName = lastName
		user.DateOfBirth = birthday
	}
	defer db.Close()

	return user
}
