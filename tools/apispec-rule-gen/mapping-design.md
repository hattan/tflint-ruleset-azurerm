
1. Clone AzureRM Repo.
2. Get Tags for Major version (2.x) forwards.
3. For each template loop through versions.
    Get List of imports.
    Query versions of Rest Api to build rule.
      


      Get List of Resources.
      Get Mapping Template
      Get list of Imports.
      Populate Mapping Template with 

```go
mapping "azurerm_analysis_services_server" {
  import_path = "azure-rest-api-specs/specification/analysisservices/resource-manager/Microsoft.AnalysisServices/stable/2017-08-01/analysisservices.json"

  name = CheckServerNameAvailabilityParameters.name  
  querypool_connection_mode = AnalysisServicesServerMutableProperties.querypoolConnectionMode
  resource_group_name = ResourceGroupNameParameter                    
}
```