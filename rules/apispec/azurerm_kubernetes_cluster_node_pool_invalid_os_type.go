// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermKubernetesClusterNodePoolInvalidOSTypeRule checks the pattern is valid
type AzurermKubernetesClusterNodePoolInvalidOSTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermKubernetesClusterNodePoolInvalidOSTypeRule returns new rule with default attributes
func NewAzurermKubernetesClusterNodePoolInvalidOSTypeRule() *AzurermKubernetesClusterNodePoolInvalidOSTypeRule {
	return &AzurermKubernetesClusterNodePoolInvalidOSTypeRule{
		resourceType:  "azurerm_kubernetes_cluster_node_pool",
		attributeName: "os_type",
		enum: []string{
			"Linux",
			"Windows",
		},
	}
}

// Name returns the rule name
func (r *AzurermKubernetesClusterNodePoolInvalidOSTypeRule) Name() string {
	return "azurerm_kubernetes_cluster_node_pool_invalid_os_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermKubernetesClusterNodePoolInvalidOSTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermKubernetesClusterNodePoolInvalidOSTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermKubernetesClusterNodePoolInvalidOSTypeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermKubernetesClusterNodePoolInvalidOSTypeRule) Check(runner tflint.Runner) error {
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
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is an invalid value as os_type`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
