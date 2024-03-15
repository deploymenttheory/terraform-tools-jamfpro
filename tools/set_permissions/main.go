package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/deploymenttheory/terraform-tools-jamfpro/tools/models"
)

func main() {
	// Load your Terraform plan JSON file
	planFilePath := "path/to/your/terraform/plan.json"
	planJson, err := os.ReadFile(planFilePath)
	if err != nil {
		log.Fatalf("Failed to read Terraform plan file: %v", err)
	}

	// Unmarshal the JSON into your TerraformPlan struct
	var plan models.TerraformPlan
	err = json.Unmarshal(planJson, &plan)
	if err != nil {
		log.Fatalf("Failed to unmarshal Terraform plan: %v", err)
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
