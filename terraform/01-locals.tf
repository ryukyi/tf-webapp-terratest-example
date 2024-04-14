resource "time_static" "created_on" {}

locals {
  # Moata Information Architect
  product_name   = "tf-webapp"
  location_short = local.location_map[var.location]

  resource_group_name = upper(format("RG-%s-%s-%s01%s", local.location_short, local.product_name, var.environment, var.unique_test_id))

  standard_tags = {
    "Resource Name" = local.resource_group_name
    Environment     = var.environment
    Description     = "Testing proof of concept for terraform"
    "Created by"    = "Terraform https://github.com/mottmac-global/workspaces-tf-webapp-terratest-poc"
    "Created on"    = time_static.created_on.rfc3339
  }

  location_map = {
    "australiaeast" = "AUE",
    "eastus"        = "USE",
    "uksouth"       = "UKS",
    "ukwest"        = "UKW",
  }
}
