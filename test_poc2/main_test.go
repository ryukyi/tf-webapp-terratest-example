// https://github.com/gruntwork-io/terratest/blob/master/test/terraform_basic_example_test.go
package test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformWebApp(t *testing.T) {
	// Generate a unique timestamp ID for this test run
	unique_test_id := "-" + time.Now().Format("2006-01-02T15-04-05")

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../terraform",
		VarFiles:     []string{"vars/qa.tfvars"},
		Vars: map[string]interface{}{
			"unique_test_id": unique_test_id,
		},
	})
	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	// test the name
	output := terraform.Output(t, terraformOptions, "azurerm_service_plan_api_name")
	assert.Equal(t, "svc-AUE-tf-webapp-example-qa01", output)

	// test the url exists and returns response
	webAppDomain := terraform.Output(t, terraformOptions, "azurerm_linux_web_app_webapp_default_hostname")
	webAppURL := fmt.Sprintf("https://%s", webAppDomain)
	// Verify that the web app URL returns a 200 OK response
	response, err := http.Get(webAppURL)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}
