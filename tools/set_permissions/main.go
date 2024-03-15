package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/deploymenttheory/terraform-tools-jamfpro/tools/models"
)

func main() {
	// Define a string flag for the Terraform plan file path
	tfPlanPath := flag.String("tfplan", "", "Path to the Terraform plan file in JSON format")

	// Parse the command-line flags
	flag.Parse()

	// Check if the tfplan flag has been set
	if *tfPlanPath == "" {
		fmt.Println("Usage: -tfplan <path to terraform plan json>")
		return
	}

	// Read the Terraform plan from the file using os.ReadFile
	planFile, err := os.ReadFile(*tfPlanPath)
	if err != nil {
		fmt.Printf("Error reading plan file: %v\n", err)
		return
	}

	// Unmarshal the JSON into the TerraformPlan struct
	var plan models.TerraformPlan
	err = json.Unmarshal(planFile, &plan)
	if err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
		return
	}

	// Iterate over ResourceChanges to identify CRUD operations
	for _, change := range plan.ResourceChanges {
		operation := getOperationType(change.Change.Actions)
		fmt.Printf("Resource: %s, Operation: %s\n", change.Address, operation)
	}
}

// getOperationType takes a slice of actions and determines the type of operation
func getOperationType(actions []string) string {
	if contains(actions, "create") {
		return "Create"
	} else if contains(actions, "update") {
		return "Update"
	} else if contains(actions, "delete") {
		return "Delete"
	} else if contains(actions, "read") {
		return "Read"
	} else {
		return "Unknown"
	}
}

// contains checks if a slice contains a specific string
func contains(slice []string, item string) bool {
	for _, sliceItem := range slice {
		if sliceItem == item {
			return true
		}
	}
	return false
}
