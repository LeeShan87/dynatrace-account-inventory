package dynatrace

import (
	"github.com/LeeShan87/dynatrace-account-inventory/generated/account"
)

type MongoGroups struct {
	account.GetGroupDto `bson:",inline"`
	Id                  string `bson:"_id" json:"gid"`
}

func (m *MongoGroups) ToGroupsDto() account.GetGroupDto {
	return m.GetGroupDto
}

func ToMongoGroups(g account.GetGroupDto) MongoGroups {
	return MongoGroups{
		GetGroupDto: g,
		Id:          *g.Uuid,
	}
}
