package test

import (
	"testing"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestPlanSingleInstanceSingleRegion(t *testing.T) {
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

	// Run `terraform init` and `terraform plan`. Fail the test if there are any errors.
	terraform.InitAndPlan(t, terraformOptions)
}

func TestPlanMultiInstanceSingleRegion(t *testing.T) {
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

	// Run `terraform init` and `terraform plan`. Fail the test if there are any errors.
	terraform.InitAndPlan(t, terraformOptions)
}

func TestPlanMultiInstanceMultiRegion(t *testing.T) {
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

	// Run `terraform init` and `terraform plan`. Fail the test if there are any errors.
	terraform.InitAndPlan(t, terraformOptions)
}