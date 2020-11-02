// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermAnalysisServicesServerInvalidNameRule checks the pattern is valid
type AzurermAnalysisServicesServerInvalidNameRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAzurermAnalysisServicesServerInvalidNameRule returns new rule with default attributes
func NewAzurermAnalysisServicesServerInvalidNameRule() *AzurermAnalysisServicesServerInvalidNameRule {
	return &AzurermAnalysisServicesServerInvalidNameRule{
		resourceType:  "azurerm_analysis_services_server",
		attributeName: "name",
		pattern:       regexp.MustCompile(`^[a-z][a-z0-9]*$`),
	}
}

// Name returns the rule name
func (r *AzurermAnalysisServicesServerInvalidNameRule) Name() string {
	return "azurerm_analysis_services_server_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermAnalysisServicesServerInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermAnalysisServicesServerInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermAnalysisServicesServerInvalidNameRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermAnalysisServicesServerInvalidNameRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[a-z][a-z0-9]*$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
