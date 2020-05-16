// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermIotcentralApplicationInvalidSkuRule checks the pattern is valid
type AzurermIotcentralApplicationInvalidSkuRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermIotcentralApplicationInvalidSkuRule returns new rule with default attributes
func NewAzurermIotcentralApplicationInvalidSkuRule() *AzurermIotcentralApplicationInvalidSkuRule {
	return &AzurermIotcentralApplicationInvalidSkuRule{
		resourceType:  "azurerm_iotcentral_application",
		attributeName: "sku",
		enum: []string{
			"F1",
			"S1",
			"ST0",
			"ST1",
			"ST2",
		},
	}
}

// Name returns the rule name
func (r *AzurermIotcentralApplicationInvalidSkuRule) Name() string {
	return "azurerm_iotcentral_application_invalid_sku"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermIotcentralApplicationInvalidSkuRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermIotcentralApplicationInvalidSkuRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermIotcentralApplicationInvalidSkuRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermIotcentralApplicationInvalidSkuRule) Check(runner tflint.Runner) error {
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