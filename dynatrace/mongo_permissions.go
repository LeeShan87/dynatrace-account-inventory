package dynatrace

import (
	"strings"

	"github.com/LeeShan87/dynatrace-account-inventory/generated/account"
)

type MongoPermission struct {
	account.PermissionsDto `bson:",inline"`
	Tenant                 string
	GroupID                string
	ManagementZone         string
	// Id                 string `bson:"_id" json:id"`
}

func (m *MongoPermission) ToPermissionsDto() account.PermissionsDto {
	return m.PermissionsDto
}

func ToMongoPermission(p account.PermissionsDto) MongoPermission {
	perm := MongoPermission{
		PermissionsDto: p,
	}
	if p.ScopeType == "account" {
		return perm
	}
	if p.ScopeType == "tenant" {
		perm.Tenant = p.Scope
		return perm
	}
	parts := strings.Split(p.Scope, ":")
	tenant := parts[0]
	mz := parts[1]
	perm.Tenant = tenant
	perm.ManagementZone = mz
	return perm
}
