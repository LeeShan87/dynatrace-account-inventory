package dynatrace

import (
	"github.com/LeeShan87/dynatrace-account-inventory/generated/account"
)

type MongoManagementZone struct {
	account.ManagementZoneResourceDto `bson:",inline"`
	Id                                string `bson:"_id" json:id"`
}

func (m *MongoManagementZone) ToManagementZoneResourceDto() account.ManagementZoneResourceDto {
	return m.ManagementZoneResourceDto
}

func ToMongoManagementZone(mz account.ManagementZoneResourceDto) MongoManagementZone {
	return MongoManagementZone{
		ManagementZoneResourceDto: mz,
		Id:                        mz.Id,
	}
}
