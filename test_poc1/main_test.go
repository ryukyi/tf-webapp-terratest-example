package test

import (
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"strings"
)

const (
	fixtures = "../terraform"
)

type TestCondition int

const (
	TestConditionEquals    TestCondition = 0
	TestConditionNotEmpty  TestCondition = 1
	TestConditionContains  TestCondition = 2
	TestConditionHTTPGetOk TestCondition = 3
)

// output variables from terraform
type TerraformOutput struct {
	azurerm_service_plan_api_name                 string
	azurerm_linux_web_app_webapp_default_hostname string
}

// All test cases follow the same pattern
// For additional logic, methods can be created
type TestCase struct {
	TestName  string
	Result    interface{}
	Expected  interface{}
	Condition TestCondition
}

func terraformOptions() *terraform.Options {
	return &terraform.Options{
		TerraformDir: fixtures,
		VarFiles:     []string{"vars/qa.tfvars"},
	}
}

// This is the whole terratest pipeline
func Test_automation(t *testing.T) {
	t.Parallel()

	terraformOptions := terraformOptions()
	terraform.InitAndApplyAndIdempotent(t, terraformOptions)
	defer terraform.Destroy(t, terraformOptions)

	output := TerraformOutput{
		azurerm_service_plan_api_name:                 terraform.Output(t, terraformOptions, "azurerm_service_plan_api_name"),
		azurerm_linux_web_app_webapp_default_hostname: terraform.Output(t, terraformOptions, "azurerm_linux_web_app_webapp_default_hostname"),
	}

	t.Run("Output Validation", func(t *testing.T) {
		OutputValidation(t, output)
	})
}

// Framework for testing each assert condition
func OutputValidation(t *testing.T, output TerraformOutput) {
	testCases := []TestCase{
		{"web app name", output.azurerm_service_plan_api_name, "svc-AUE-tf-webapp-example-qa01", TestConditionEquals},
		{"web app http status", getHTTPSResponse(output.azurerm_linux_web_app_webapp_default_hostname).StatusCode, http.StatusOK, TestConditionHTTPGetOk},
	}
	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			switch tc.Condition {
			case TestConditionEquals:
				assert.Equal(t, tc.Result, tc.Expected)
			case TestConditionNotEmpty:
				assert.NotEmpty(t, tc.Result)
			case TestConditionContains:
				assert.Contains(t, tc.Result, tc.Expected)
			case TestConditionHTTPGetOk:
				assert.Equal(t, tc.Result, tc.Expected)
			}
		})
	}
}

func getHTTPSResponse(url string) *http.Response {
	if !strings.HasPrefix(url, "https://") {
		modifiedURL := fmt.Sprintf("https://%s", url)
		log.Printf("Modified URL: %s", modifiedURL)
		url = modifiedURL
	}

	response, err := http.Get(url)
	if err != nil {
		log.Printf("Error making HTTP request: %v", err)
	}

	return response
}
