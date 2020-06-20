package repository

import (
	"log"
	"time"

	db "github.com/fredoliveira-ca/products-golang-java/user-service/app/config"
	"github.com/fredoliveira-ca/products-golang-java/user-service/domain"
)

const (
	selectOneUserByID = "SELECT * FROM \"user\" WHERE user_id=$1"
)

// FindOne has the responsibility to connect with the database
// 	and return a user based on the informed identifier.
func FindOne(id string) domain.User {
	db := db.ConnectDataBase()

	selection, err := db.Query(selectOneUserByID, id)
	if err != nil {
		log.Println(err)
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
