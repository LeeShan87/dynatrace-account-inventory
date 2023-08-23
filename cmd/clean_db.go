package main

import (
	"context"

	"github.com/LeeShan87/dynatrace-account-inventory/utils"
)

func main() {
	client, err := utils.ConnecToMongoDB()
	utils.CheckError(err)
	defer client.Disconnect(context.Background())
	client.Database("account_api").Drop(context.Background())
}
