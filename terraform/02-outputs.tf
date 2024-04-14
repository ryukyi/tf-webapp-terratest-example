# Service plan 
output "azurerm_service_plan_api_id" {
  value = azurerm_service_plan.app_service_plan_webapp.id
}

output "azurerm_service_plan_api_name" {
  value = azurerm_service_plan.app_service_plan_webapp.name
}

output "azurerm_linux_web_app_webapp_id" {
  value     = azurerm_linux_web_app.webapp.id
  sensitive = true
}

output "azurerm_linux_web_app_webapp_default_hostname" {
  value     = azurerm_linux_web_app.webapp.default_hostname
  sensitive = true
}

