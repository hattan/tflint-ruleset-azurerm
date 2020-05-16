<!--- This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT --->

# azurerm_express_route_circuit_peering_invalid_peering_type

Warns about values that appear to be invalid based on [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs).

Allowed values are:
- AzurePublicPeering
- AzurePrivatePeering
- MicrosoftPeering

## Example

```hcl
resource "azurerm_express_route_circuit_peering" "foo" {
  peering_type = ... // invalid value
}
```

```
$ tflint
1 issue(s) found:

Error: "..." is an invalid value as peering_type (azurerm_express_route_circuit_peering_invalid_peering_type)

  on template.tf line 2:
  2:   peering_type = ... // invalid value

Reference: https://github.com/terraform-linters/tflint-ruleset-azurerm/blob/v0.1.0/docs/rules/azurerm_express_route_circuit_peering_invalid_peering_type.md

```

## Why

Requests containing invalid values will return an error when calling the API by `terraform apply`.

## How to Fix

Replace the warned value with a valid value.

## Source

This rule is automatically generated from [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs). If you are uncertain about the warning, check the following API schema referenced by this rule.

https://github.com/Azure/azure-rest-api-specs/tree/master/specification/network/resource-manager/Microsoft.Network/stable/2020-03-01/expressRouteCircuit.json