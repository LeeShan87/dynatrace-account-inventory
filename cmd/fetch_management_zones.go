package main

import (
	"context"
	"fmt"

	auth "github.com/LeeShan87/dynatrace-account-inventory/dynatrace"
	"github.com/LeeShan87/dynatrace-account-inventory/generated/account"
	"github.com/LeeShan87/dynatrace-account-inventory/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	client, err := utils.ConnecToMongoDB()
	utils.CheckError(err)
	defer client.Disconnect(context.Background())

	db := client.Database("account_api")

	authClient := auth.NewDTAuthClient()
	apiClient, err := auth.GetApiClient(*authClient, utils.ApiURL)
	utils.CheckError(err)

	// fetch tenants and management zones
	environments, err := fetchEnvironments(apiClient, utils.AccountID)
	utils.CheckError(err)
	tenants, MZs, err := saveEnvironmentsInMongo(db, environments)
	utils.CheckError(err)
	fmt.Println("Inserted ID:", len(tenants.InsertedIDs))
	fmt.Println("Inserted ID:", len(MZs.InsertedIDs))
}

func fetchEnvironments(client *account.APIClient, accountID string) (*account.EnvironmentResourceDto, error) {
	environments, _, err := client.EnvironmentManagementAPI.EnvironmentResourcesControllerGetEnvironmentResources(context.Background(), accountID).Execute()
	return environments, err
}

func saveEnvironmentsInMongo(db *mongo.Database, environments *account.EnvironmentResourceDto) (*mongo.InsertManyResult, *mongo.InsertManyResult, error) {
	tenantCollection := db.Collection("tenants")
	managementZoneCollection := db.Collection("management_zones")
	tenantBson := make([]interface{}, len(environments.TenantResources))
	managementZoneBson := make([]interface{}, len(environments.ManagementZoneResources))
	for i, tenant := range environments.TenantResources {
		mongoTenant := auth.ToMongoTenant(tenant)
		bsonTenant, _ := bson.Marshal(mongoTenant)
		tenantBson[i] = bsonTenant
	}
	for i, managementZone := range environments.ManagementZoneResources {
		mongoManagementZone := auth.ToMongoManagementZone(managementZone)
		bsonManagementZone, _ := bson.Marshal(mongoManagementZone)
		managementZoneBson[i] = bsonManagementZone
	}
	insertedTenant, err := tenantCollection.InsertMany(context.Background(), tenantBson)
	if err != nil {
		return nil, nil, err
	}
	insertedMZs, err := managementZoneCollection.InsertMany(context.Background(), managementZoneBson)
	if err != nil {
		return nil, nil, err
	}

	return insertedTenant, insertedMZs, nil
}
