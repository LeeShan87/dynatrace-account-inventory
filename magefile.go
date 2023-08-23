//go:build mage
// +build mage

package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/magefile/mage/sh"
	//"path/filepath"
)

var (
	accountSwagger string
)

func init() {
	_ = godotenv.Load(".env")
	accountSwagger = os.Getenv("ACCOUNT_SWAGGER")
}

func PopulateDB() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	err = sh.Run("go", "run", cwd+"/cmd/fetch_management_zones.go")
	if err != nil {
		return err
	}
	err = sh.Run("go", "run", cwd+"/cmd/fetch_groups.go")
	if err != nil {
		return err
	}
	err = sh.Run("go", "run", cwd+"/cmd/fetch_users.go")
	if err != nil {
		return err
	}
	err = sh.Run("go", "run", cwd+"/cmd/fetch_group_permissions.go")
	if err != nil {
		return err
	}
	err = sh.Run("go", "run", cwd+"/cmd/fetch_relations.go")
	if err != nil {
		return err
	}
	return nil
}

func CleanDB() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	err = sh.Run("go", "run", cwd+"/cmd/clean_db.go")
	if err != nil {
		return err
	}
	return nil
}

func Clean() error {
	err := CleanDB()
	if err != nil {
		return err
	}
	return nil
}

func Generate() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	//outputDir := filepath.Join(cwd, "generated", "account")
	err = sh.Run("docker", "run", "--rm",
		"-v", cwd+":/local",
		"openapitools/openapi-generator-cli",
		"generate", "-i", accountSwagger, "-g", "go", "-o", "/local/generated/account",
		"--skip-validate-spec", "--additional-properties", "packageName=account",
	)
	return err
}
