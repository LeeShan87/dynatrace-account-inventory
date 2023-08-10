package main

import (
	"context"
	"fmt"
	"os"

	auth "github.com/LeeShan87/dynatrace-account-inventory/dynatrace"
	"github.com/LeeShan87/dynatrace-account-inventory/generated/account"
	"github.com/LeeShan87/dynatrace-account-inventory/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	apiURL = "https://api.dynatrace.com"
)

func main() {
	accountID := os.Getenv("ACOUNT_ID")
	mongoURI := "mongodb://localhost:27017"
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	utils.CheckError(err)
	defer client.Disconnect(context.Background())

	db := client.Database("account_api")

	authClient := auth.NewDTAuthClient()
	apiClient, err := auth.GetApiClient(*authClient, apiURL)
	utils.CheckError(err)

	// fetch tenants and management zones
	environments, err := fetchEnvironments(apiClient, accountID)
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
		bsonTenant, _ := bson.Marshal(tenant)
		tenantBson[i] = bsonTenant
	}
	for i, managementZone := range environments.ManagementZoneResources {
		bsonManagementZone, _ := bson.Marshal(managementZone)
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
