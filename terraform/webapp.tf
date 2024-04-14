
# App service plan - API
resource "azurerm_service_plan" "app_service_plan_webapp" {
  name                = format("svc-%s-%s-example-%s01", local.location_short, local.product_name, var.environment)
  location            = var.location
  resource_group_name = azurerm_resource_group.group.name
  os_type             = "Linux"
  sku_name            = var.app_service_sku
}

# App service - API
# The `azurerm_app_service` resource has been superseded by the `azurerm_linux_web_app` and `azurerm_windows_web_app`
resource "azurerm_linux_web_app" "webapp" {
  name                = format("svc-%s-%s-example-%s02", local.location_short, local.product_name, var.environment)
  location            = var.location
  resource_group_name = azurerm_resource_group.group.name
  service_plan_id     = azurerm_service_plan.app_service_plan_webapp.id

  site_config {}
}
