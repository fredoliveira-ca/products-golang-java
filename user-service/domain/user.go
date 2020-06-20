package domain

import (
	"time"
)

//User domain struct.
type User struct {
	ID          string
	FirstName   string
	LastName    string
	DateOfBirth time.Time
}
