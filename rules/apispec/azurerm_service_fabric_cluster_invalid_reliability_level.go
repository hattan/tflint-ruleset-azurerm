// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermServiceFabricClusterInvalidReliabilityLevelRule checks the pattern is valid
type AzurermServiceFabricClusterInvalidReliabilityLevelRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermServiceFabricClusterInvalidReliabilityLevelRule returns new rule with default attributes
func NewAzurermServiceFabricClusterInvalidReliabilityLevelRule() *AzurermServiceFabricClusterInvalidReliabilityLevelRule {
	return &AzurermServiceFabricClusterInvalidReliabilityLevelRule{
		resourceType:  "azurerm_service_fabric_cluster",
		attributeName: "reliability_level",
		enum: []string{
			"None",
			"Bronze",
			"Silver",
			"Gold",
			"Platinum",
		},
	}
}

// Name returns the rule name
func (r *AzurermServiceFabricClusterInvalidReliabilityLevelRule) Name() string {
	return "azurerm_service_fabric_cluster_invalid_reliability_level"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermServiceFabricClusterInvalidReliabilityLevelRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermServiceFabricClusterInvalidReliabilityLevelRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermServiceFabricClusterInvalidReliabilityLevelRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermServiceFabricClusterInvalidReliabilityLevelRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as reliability_level`, truncateLongMessage(val)),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}
