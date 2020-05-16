// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermEventhubNamespaceInvalidCapacityRule checks the pattern is valid
type AzurermEventhubNamespaceInvalidCapacityRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAzurermEventhubNamespaceInvalidCapacityRule returns new rule with default attributes
func NewAzurermEventhubNamespaceInvalidCapacityRule() *AzurermEventhubNamespaceInvalidCapacityRule {
	return &AzurermEventhubNamespaceInvalidCapacityRule{
		resourceType:  "azurerm_eventhub_namespace",
		attributeName: "capacity",
		max:           20,
	}
}

// Name returns the rule name
func (r *AzurermEventhubNamespaceInvalidCapacityRule) Name() string {
	return "azurerm_eventhub_namespace_invalid_capacity"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermEventhubNamespaceInvalidCapacityRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermEventhubNamespaceInvalidCapacityRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermEventhubNamespaceInvalidCapacityRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermEventhubNamespaceInvalidCapacityRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val int
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if val > r.max {
				runner.EmitIssue(
					r,
					"capacity must be 20 or less",
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}
