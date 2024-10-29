package test

import (
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
)

type Options map[string]interface{}

func BackendConfig() Options {
	return Options{
		"storage_account_name": os.Getenv("STORAGE_ACCOUNT_NAME"),
		"container_name":       os.Getenv("CONTAINER_NAME"),
		"key":                  os.Getenv("KEY"),
	}
}

func DefaultOptions() Options {
	return Options{
		"docker_image_tag":    os.Getenv("DOCKER_IMAGE_TAG"),
		"psql_admin_username": os.Getenv("PSQL_ADMIN_USERNAME"),
		"psql_admin_password": os.Getenv("PSQL_ADMIN_PASSWORD"),
	}
}

func Setup(t *testing.T, e string, backendConfig Options, opts Options) *terraform.Options {
	tempFolder := test_structure.CopyTerraformFolderToTemp(t, "../", e)
	return &terraform.Options{
		BackendConfig: backendConfig,
		TerraformDir:  tempFolder,
		Vars:          opts,
	}
}
