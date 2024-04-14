variable "subscription_id" {
  type        = string
  description = "The ID of the Azure subscription."
}

variable "location" {
  type        = string
  description = "The location of the resource group."
  validation {
    condition     = contains(["australiaeast", "uksouth", "eastus", "ukwest"], var.location)
    error_message = "The supplied environment name is not recognized (see local.location_map)"
  }
}

variable "environment" {
  type        = string
  description = "The environment name (e.g., dev, qa, prd)."

  validation {
    condition     = contains(["dev", "qa", "prd"], var.environment)
    error_message = "The environment must be one of: dev, qa, prd"
  }
}

variable "app_service_sku" {
  type        = string
  description = "Service tier for app service hosting plan"
  default     = "P1v2"
}

# SQL Vars
variable "sql_sku" {
  type        = string
  description = "Service tier for Azure SQL"
  default     = "S0"
}

# Test instance resource group suffix to make all test deployments unique
variable "unique_test_id" {
  type    = string
  default = ""
}
