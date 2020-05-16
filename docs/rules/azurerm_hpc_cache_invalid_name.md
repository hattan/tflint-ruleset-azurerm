<!--- This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT --->

# azurerm_hpc_cache_invalid_name

Warns about values that appear to be invalid based on [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs).

In this rule, the string must match the regular expression `^[-0-9a-zA-Z_]{1,80}$``.

## Example

```hcl
resource "azurerm_hpc_cache" "foo" {
  name = ... // invalid value
}
```

```
$ tflint
1 issue(s) found:

Error: "..." does not match valid pattern ^[-0-9a-zA-Z_]{1,80}$ (azurerm_hpc_cache_invalid_name)

  on template.tf line 2:
  2:   name = ... // invalid value

Reference: https://github.com/terraform-linters/tflint-ruleset-azurerm/blob/v0.1.0/docs/rules/azurerm_hpc_cache_invalid_name.md

```

## Why

Requests containing invalid values will return an error when calling the API by `terraform apply`.

## How to Fix

Replace the warned value with a valid value.

## Source

This rule is automatically generated from [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs). If you are uncertain about the warning, check the following API schema referenced by this rule.

https://github.com/Azure/azure-rest-api-specs/tree/master/specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2020-03-01/storagecache.json