mapping "azurerm_analysis_services_server" {
  import_path = "azure-rest-api-specs/specification/analysisservices/resource-manager/Microsoft.AnalysisServices/stable/2017-08-01/analysisservices.json"

  name = CheckServerNameAvailabilityParameters.name  
  querypool_connection_mode = AnalysisServicesServerMutableProperties.querypoolConnectionMode
  resource_group_name = ResourceGroupNameParameter                    
}
