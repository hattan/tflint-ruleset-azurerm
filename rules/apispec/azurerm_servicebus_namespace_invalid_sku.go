// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermServicebusNamespaceInvalidSkuRule checks the pattern is valid
type AzurermServicebusNamespaceInvalidSkuRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermServicebusNamespaceInvalidSkuRule returns new rule with default attributes
func NewAzurermServicebusNamespaceInvalidSkuRule() *AzurermServicebusNamespaceInvalidSkuRule {
	return &AzurermServicebusNamespaceInvalidSkuRule{
		resourceType:  "azurerm_servicebus_namespace",
		attributeName: "sku",
		enum: []string{
			"Basic",
			"Standard",
			"Premium",
		},
	}
}

// Name returns the rule name
func (r *AzurermServicebusNamespaceInvalidSkuRule) Name() string {
	return "azurerm_servicebus_namespace_invalid_sku"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermServicebusNamespaceInvalidSkuRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermServicebusNamespaceInvalidSkuRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermServicebusNamespaceInvalidSkuRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermServicebusNamespaceInvalidSkuRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as sku`, truncateLongMessage(val)),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}
