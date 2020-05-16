// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermDevspaceControllerInvalidSkuNameRule checks the pattern is valid
type AzurermDevspaceControllerInvalidSkuNameRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermDevspaceControllerInvalidSkuNameRule returns new rule with default attributes
func NewAzurermDevspaceControllerInvalidSkuNameRule() *AzurermDevspaceControllerInvalidSkuNameRule {
	return &AzurermDevspaceControllerInvalidSkuNameRule{
		resourceType:  "azurerm_devspace_controller",
		attributeName: "sku_name",
		enum: []string{
			"S1",
		},
	}
}

// Name returns the rule name
func (r *AzurermDevspaceControllerInvalidSkuNameRule) Name() string {
	return "azurerm_devspace_controller_invalid_sku_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermDevspaceControllerInvalidSkuNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermDevspaceControllerInvalidSkuNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermDevspaceControllerInvalidSkuNameRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermDevspaceControllerInvalidSkuNameRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as sku_name`, truncateLongMessage(val)),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}
