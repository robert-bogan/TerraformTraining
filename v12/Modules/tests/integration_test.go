package test

import (
	"testing"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestApplySingleInstanceSingleRegion(t *testing.T) {
	t.Parallel()

	// Define variables
	locations := []string{"UK South"}

	// Enable retryable error
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{

		// The path to where the Terraform code is located
		TerraformDir: "../examples/",

		// Variables to pass to the Terraform code using -var options
		Vars: map[string]interface{}{
			"deployment": "SingleInstanceSingleRegion",
			"resource_instance_count": 1,
			"resource_location": locations,
		},
	})

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)
}

func TestApplyMultiInstanceSingleRegion(t *testing.T) {
	t.Parallel()

	// Define variables
	locations := []string{"UK South"}

	// Enable retryable error
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{

		// The path to where the Terraform code is located
		TerraformDir: "../examples/",

		// Variables to pass to the Terraform code using -var options
		Vars: map[string]interface{}{
			"deployment": "MultiInstanceSingleRegion",
			"resource_instance_count": 2,
			"resource_location": locations,
		},
	})

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)
}

func TestApplyMultiInstanceMultiRegion(t *testing.T) {
	t.Parallel()

	// Define variables
	locations := []string{"UK South", "North Europe"}

	// Enable retryable error
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{

		// The path to where the Terraform code is located
		TerraformDir: "../examples/",

		// Variables to pass to the Terraform code using -var options
		Vars: map[string]interface{}{
			"deployment": "MultiInstanceMultiRegion",
			"resource_instance_count": 4,
			"resource_location": locations,
		},
	})

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)
}