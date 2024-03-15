package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/deploymenttheory/terraform-tools-jamfpro/tools/models"
)

const (
	// Define ANSI color codes
	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
	colorReset = "\033[0m"
)

// main function parses command line arguments to locate the Terraform plan file, unmarshals it, and checks for duplicate resource names.
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

	var plan models.TerraformPlan
	err = json.Unmarshal(planFile, &plan)
	if err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
		return
	}

	// Specified the resource types to validate duplicates for
	interestedResourceTypes := map[string]bool{
		"jamfpro_account":                       true,
		"jamfpro_account_group":                 true,
		"jamfpro_advanced_computer_search":      true,
		"jamfpro_advanced_mobile_device_search": true,
		"jamfpro_advanced_user_search":          true,
		"jamfpro_allowed_file_extension":        true,
		"jamfpro_api_integration":               true,
		"jamfpro_api_role":                      true,
		"jamfpro_building":                      true,
		"jamfpro_category":                      true,
		"jamfpro_computer_checkin":              true,
		"jamfpro_computer_extension_attribute":  true,
		"jamfpro_computer_group":                true,
		"jamfpro_computer_prestage":             true,
		"jamfpro_department":                    true,
		"jamfpro_disk_encryption_configuration": true,
		"jamfpro_dock_item":                     true,
		"jamfpro_file_share_distribution_point": true,
		"jamfpro_site":                          true,
		"jamfpro_script":                        true,
		"jamfpro_network_segment":               true,
		"jamfpro_package":                       true,
		"jamfpro_policy":                        true,
		"jamfpro_printer":                       true,
	}

	// Store resource names and their occurrences
	resourceNames := make(map[string]int)

	// Iterate over resources in the plan
	for _, resource := range plan.PlannedValues.RootModule.Resources {
		if _, ok := interestedResourceTypes[resource.Type]; ok {
			// Attempt to extract the name from the Values map
			if name, exists := resource.Values["name"]; exists && name != nil {
				if nameStr, ok := name.(string); ok {
					resourceNames[nameStr]++
				}
			}
		}
	}

	// Check for duplicates
	foundDuplicates := false
	for name, count := range resourceNames {
		if count > 1 {
			errorMessage := fmt.Sprintf("Error: Duplicate Jamf Pro resource name found: %s, Count: %d", name, count)
			printColor(errorMessage, colorRed)
			foundDuplicates = true
		}
	}

	if !foundDuplicates {
		printColor("Check completed: No duplicate Jamf Pro resource names found within the specified Terraform plan.", colorGreen)
	}
}

// printColor prints a message with the specified color to the console.
func printColor(message string, colorCode string) {
	fmt.Println(colorCode, message, colorReset)
}
