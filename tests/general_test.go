package test

import (
	"crypto/tls"
	"fmt"
	"testing"
	"time"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestSmoke(t *testing.T) {
	t.Parallel()

	backend := BackendConfig()
	opts := DefaultOptions()

	terraformOptions := Setup(t, "iac/fibonacci_app_v1", backend, opts)

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	smokeTestText := "Enter your index:"
	smokeUrl := terraform.Output(t, terraformOptions, "ingress")
	smokeUrl = fmt.Sprintf("http://%s", smokeUrl)
	tlsConfig := tls.Config{}
	maxRetries := 30
	timeBetweenRetries := 5 * time.Second

	http_helper.HttpGetWithRetry(t, smokeUrl, &tlsConfig, 200, smokeTestText, maxRetries, timeBetweenRetries)
}
