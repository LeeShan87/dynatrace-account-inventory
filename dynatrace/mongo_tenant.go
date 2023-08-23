package dynatrace

import (
	"github.com/LeeShan87/dynatrace-account-inventory/generated/account"
)

type MongoTenant struct {
	account.TenantResourceDto `bson:",inline"`
	Id                        string `bson:"_id" json:id"`
}

func (m *MongoTenant) ToTenantResourceDto() account.TenantResourceDto {
	return m.TenantResourceDto
}

func ToMongoTenant(t account.TenantResourceDto) MongoTenant {
	return MongoTenant{
		TenantResourceDto: t,
		Id:                t.Id,
	}
}
