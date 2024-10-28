package test

import (
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
)

type Options map[string]any

func DefaultOptions() Options {
	return Options{
		"docker_image_tag":    os.Getenv("DOCKER_IMAGE_TAG"),
		"psql_admin_username": os.Getenv("PSQL_ADMIN_USERNAME"),
		"psql_admin_password": os.Getenv("PSQL_ADMIN_PASSWORD"),
	}
}

func Setup(t *testing.T, e string, opts Options) *terraform.Options {
	tempFolder := test_structure.CopyTerraformFolderToTemp(t, "../", e)
	return &terraform.Options{
		TerraformDir: tempFolder,
		Vars:         opts,
	}
}
