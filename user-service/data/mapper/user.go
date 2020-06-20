package mapper

import (
	"github.com/fredoliveira-ca/products-golang-java/user-service/app/grpc/proto/userpb"
	"github.com/fredoliveira-ca/products-golang-java/user-service/domain"
)

// ToProto is a converter which parse domain class to proto.
func ToProto(user domain.User) *userpb.UserProto {
	return &userpb.UserProto{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		DateOfBirth: &userpb.Date{
			Year:  int32(user.DateOfBirth.Year()),
			Month: int32(user.DateOfBirth.Month()),
			Day:   int32(user.DateOfBirth.Day()),
		},
	}
}
