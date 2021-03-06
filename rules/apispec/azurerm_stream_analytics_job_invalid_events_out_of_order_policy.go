// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermStreamAnalyticsJobInvalidEventsOutOfOrderPolicyRule checks the pattern is valid
type AzurermStreamAnalyticsJobInvalidEventsOutOfOrderPolicyRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermStreamAnalyticsJobInvalidEventsOutOfOrderPolicyRule returns new rule with default attributes
func NewAzurermStreamAnalyticsJobInvalidEventsOutOfOrderPolicyRule() *AzurermStreamAnalyticsJobInvalidEventsOutOfOrderPolicyRule {
	return &AzurermStreamAnalyticsJobInvalidEventsOutOfOrderPolicyRule{
		resourceType:  "azurerm_stream_analytics_job",
		attributeName: "events_out_of_order_policy",
		enum: []string{
			"Adjust",
			"Drop",
		},
	}
}

// Name returns the rule name
func (r *AzurermStreamAnalyticsJobInvalidEventsOutOfOrderPolicyRule) Name() string {
	return "azurerm_stream_analytics_job_invalid_events_out_of_order_policy"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermStreamAnalyticsJobInvalidEventsOutOfOrderPolicyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermStreamAnalyticsJobInvalidEventsOutOfOrderPolicyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermStreamAnalyticsJobInvalidEventsOutOfOrderPolicyRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermStreamAnalyticsJobInvalidEventsOutOfOrderPolicyRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as events_out_of_order_policy`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
