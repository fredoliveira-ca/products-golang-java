package domain

import (
	"time"
)

//User is...
type User struct {
	ID          string
	FirstName   string
	LastName    string
	DateOfBirth time.Time
}
