package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/deploymenttheory/terraform-tools-jamfpro/tools/models"
)

func main() {
	tfPlanPath := flag.String("tfplan", "", "Path to the Terraform plan file in JSON format")
	flag.Parse()

	if *tfPlanPath == "" {
		fmt.Println("Usage: -tfplan <path to terraform plan json>")
		return
	}

	planFile, err := os.ReadFile(*tfPlanPath)
	if err != nil {
		fmt.Printf("Error reading plan file: %v\n", err)
		return
	}

	var plan models.TerraformPlan
	err = json.Unmarshal(planFile, &plan)
	if err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
		return
	}

	// Define your security-related conditions here
	securityResources := map[string]bool{
		"jamfpro_api_integration":               true,
		"jamfpro_disk_encryption_configuration": true,
		// Add more resources or properties that you consider security-related
	}

	securityChangesDetected := false

	for _, change := range plan.ResourceChanges {
		// Check if the resource type is one of the security related resources
		if _, ok := securityResources[change.Type]; ok {
			// Check the actions for create, update, or delete
			for _, action := range change.Change.Actions {
				if action == "create" || action == "update" || action == "delete" {
					securityChangesDetected = true
					fmt.Printf("Security-related change detected: %s action on %s\n", action, change.Address)
					break // Break out of the inner loop once a security-related change is found
				}
			}
			if securityChangesDetected {
				break // Break out of the outer loop once a security-related change is found
			}
		}
	}

	// Write to GITHUB_OUTPUT environment file
	if securityChangesDetected {
		outputFile := os.Getenv("GITHUB_OUTPUT")
		if outputFile == "" {
			fmt.Println("GITHUB_OUTPUT environment variable not set")
			return
		}

		file, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Error opening output file: %v\n", err)
			return
		}
		defer file.Close()

		if _, err := file.WriteString("approval_group=Security\n"); err != nil {
			fmt.Printf("Error writing to output file: %v\n", err)
			return
		}

		fmt.Println("Security-related changes detected. 'Security' group set for GitHub PR approval.")
	} else {
		fmt.Println("No security-related changes detected.")
	}
}
