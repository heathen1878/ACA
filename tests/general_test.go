package test

import (
	"crypto/tls"
	"testing"
	"time"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestSmoke(t *testing.T) {
	t.Parallel()

	opts := DefaultOptions()

	terraformOptions := Setup(t, "iac/fibonacci_app_v1", opts)

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	smokeTestText := "Enter your index:"
	smokeUrl := terraform.Output(t, terraformOptions, "ingress")
	tlsConfig := tls.Config{}
	maxRetries := 30
	timeBetweenRetries := 5 * time.Second

	http_helper.HttpGetWithRetry(t, smokeUrl, &tlsConfig, 200, smokeTestText, maxRetries, timeBetweenRetries)
}
