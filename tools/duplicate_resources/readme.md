# Terraform Jamf Pro Duplicate Resource Name Checker

This Go application provides a tool for detecting duplicate resource names in Jamf Pro configurations defined within a Terraform plan. It is specifically designed to help ensure unique naming conventions are maintained in Jamf Pro resources managed via Terraform, thus avoiding potential conflicts or issues during infrastructure provisioning.

## Features

- **Duplicate Detection**: Scans a Terraform plan for duplicate Jamf Pro resource names across a predefined set of resource types.
- **Customizable Resource Types**: Includes a pre-configured set of Jamf Pro resource types for duplicate checking, with easy customization to include additional types.
- **Color-Coded Output**: Utilizes ANSI color codes to highlight duplicate errors (in red) and success messages (in green) for easy visibility.

## Usage

To use this tool, you will need a Terraform plan file in JSON format. You can generate this file by running `terraform plan -out=tfplan.binary` followed by `terraform show -json tfplan.binary > tfplan.json`.

### Command Line Flags

- `-tfplan`: Specifies the path to the Terraform plan file in JSON format.



### Example

```bash
go run main.go -tfplan /path/to/your/terraform/plan.json
```

This command parses the specified Terraform plan file and checks for duplicate Jamf Pro resource names within the supported resource types.

Installation
To run this tool, you must have Go installed on your machine. Follow these steps to compile and run the application:

Clone the repository to your local machine:

```bash
git clone https://github.com/deploymenttheory/terraform-tools-jamfpro.git
```
Navigate to the cloned directory:

```bash
cd terraform-tools-jamfpro
```

Run the application with Go:

```bash
go run main.go -tfplan /path/to/your/terraform/plan.json
```

Supported Resource Types
The tool checks for duplicates in the following Jamf Pro resource types by default:

```bash
jamfpro_account
jamfpro_account_group
jamfpro_advanced_computer_search
jamfpro_advanced_mobile_device_search
jamfpro_advanced_user_search
jamfpro_allowed_file_extension
jamfpro_api_integration
jamfpro_api_role
jamfpro_building
jamfpro_category
jamfpro_computer_checkin
jamfpro_computer_extension_attribute
jamfpro_computer_group
jamfpro_computer_prestage
jamfpro_department
jamfpro_disk_encryption_configuration
jamfpro_dock_item
jamfpro_file_share_distribution_point
jamfpro_site
jamfpro_script
jamfpro_network_segment
jamfpro_package
jamfpro_policy
jamfpro_printer
```

You can customize this list by modifying the interestedResourceTypes map in the main.go file.

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request if you have suggestions for improvements or additions.

## License

This project is licensed under the MIT License - see the LICENSE file for details.