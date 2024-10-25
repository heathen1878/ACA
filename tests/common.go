package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
)

type Options map[string]any

func DefaultOptions() Options {
	return Options{
		""
	}
}

func (o Options) With(with Options) Options {
	options := o
	for k, v := range with {
		options[k] = v
	}
	return options
}

func (o Options) Without(key string) Options {
	option := o
	delete(option, key)
	return option
}

func Setup(t *testing.T, e string, opts Options) *terraform.Options {
	tempFolder := test_structure.CopyTerraformFolderToTemp(t, "../", e)
	return &terraform.Options{
		TerraformDir: tempFolder,
		Vars:         opts,
	}
}

func GetTestConfig(t *testing.T) Options {
	t.Helper()

	return Options{}
}
