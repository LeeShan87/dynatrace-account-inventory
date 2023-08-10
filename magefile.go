//go:build mage
// +build mage

package main

import (
	"github.com/joho/godotenv"
	"github.com/magefile/mage/sh"
	"os"
	//"path/filepath"
)

var (
	accountSwagger string
)

func init() {
	_ = godotenv.Load(".env")
	accountSwagger = os.Getenv("ACCOUNT_SWAGGER")
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
