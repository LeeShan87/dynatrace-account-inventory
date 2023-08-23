package utils

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	ApiURL    = "https://api.dynatrace.com"
	AccountID string
)

func init() {
	godotenv.Load()
	AccountID = os.Getenv("ACCOUNT_ID")
}
