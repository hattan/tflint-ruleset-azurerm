// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermIothubEndpointStorageContainerInvalidEncodingRule checks the pattern is valid
type AzurermIothubEndpointStorageContainerInvalidEncodingRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermIothubEndpointStorageContainerInvalidEncodingRule returns new rule with default attributes
func NewAzurermIothubEndpointStorageContainerInvalidEncodingRule() *AzurermIothubEndpointStorageContainerInvalidEncodingRule {
	return &AzurermIothubEndpointStorageContainerInvalidEncodingRule{
		resourceType:  "azurerm_iothub_endpoint_storage_container",
		attributeName: "encoding",
		enum: []string{
			"Avro",
			"AvroDeflate",
			"JSON",
		},
	}
}

// Name returns the rule name
func (r *AzurermIothubEndpointStorageContainerInvalidEncodingRule) Name() string {
	return "azurerm_iothub_endpoint_storage_container_invalid_encoding"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermIothubEndpointStorageContainerInvalidEncodingRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermIothubEndpointStorageContainerInvalidEncodingRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermIothubEndpointStorageContainerInvalidEncodingRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermIothubEndpointStorageContainerInvalidEncodingRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as encoding`, truncateLongMessage(val)),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}
