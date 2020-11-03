
# Mapping Design

1. Clone AzureRM Repo
2. Get Tags for AzureRm Major version (2.x) forwards.
3. For each mapping template
    For each azurerm version (2.x.x - x.x.x)
      Get imports for azure go sdk
      Update importMappings, bump max version or create new entry in list with rest api and min version
4. Generate rules
      Using Mapping Template (.hcl) with import paths and min/max versions, generate rules in go

## Mapping Template

```go
mapping "azurerm_analysis_services_server" {
  name = CheckServerNameAvailabilityParameters.name,
  querypool_connection_mode = AnalysisServicesServerMutableProperties.querypoolConnectionMode,
  resource_group_name = ResourceGroupNameParameter
  
  importMappings[]
}
```

## Mapping Output from step 3

```go
mapping "azurerm_analysis_services_server" {
  name = CheckServerNameAvailabilityParameters.name,
  querypool_connection_mode = AnalysisServicesServerMutableProperties.querypoolConnectionMode,
  resource_group_name = ResourceGroupNameParameter
  spec_file = "analysisservices.json""

  importMappings[
    {
      minVersion: "v2.0.0",
      maxVersion: "v2.10.0",
      import_path: "azure-rest-api-specs/specification/analysisservices/resource-manager/Microsoft.AnalysisServices/stable/2017-08-01/analysisservices.json"
    },
    {
      minVersion: "v2.11.0",
      maxVersion: "v2.20.0",
      import_path: "azure-rest-api-specs/specification/analysisservices/resource-manager/Microsoft.AnalysisServices/stable/2018-08-01/analysisservices.json"},
    },
    {
      minVersion: "v2.21.0",
      maxVersion: "",
      import_path: "azure-rest-api-specs/specification/analysisservices/resource-manager/Microsoft.AnalysisServices/stable/2019-08-01/analysisservices.json"},
    },
  ]
}
```

## File names

`azurerm_analysis_services_server_invalid_name.go` becomes `azurerm_analysis_services_server_invalid_name_minVersion-maxVersion.go`

Using the example above you would have:

- `azurerm_analysis_services_server_invalid_name_v2.0.0-v2.10.0.go`
- `azurerm_analysis_services_server_invalid_name_v2.11.0-v2.20.0.go`
- `azurerm_analysis_services_server_invalid_name_v2.21.0-curent.go`
