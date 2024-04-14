terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "3.71"
    }
  }
}

provider "azurerm" {
  features {}
  subscription_id            = var.subscription_id
  skip_provider_registration = true
}

# Azure resource group
resource "azurerm_resource_group" "group" {
  name     = local.resource_group_name
  location = var.location
}

# Exposes information about the Azure CLI authenticator, e.g. the tenant
data "azurerm_client_config" "current" {}
