package dynatrace

import (
	"github.com/LeeShan87/dynatrace-account-inventory/generated/account"
)

type MongoUser struct {
	account.UsersDto `bson:",inline"`
	Id               string `bson:"_id" json:"uid"`
}

func (m *MongoUser) ToUserDto() account.UsersDto {
	return m.UsersDto
}

func ToMongoUser(u account.UsersDto) MongoUser {
	return MongoUser{
		UsersDto: u,
		Id:       u.Uid,
	}
}
