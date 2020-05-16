// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermBatchCertificateInvalidFormatRule checks the pattern is valid
type AzurermBatchCertificateInvalidFormatRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermBatchCertificateInvalidFormatRule returns new rule with default attributes
func NewAzurermBatchCertificateInvalidFormatRule() *AzurermBatchCertificateInvalidFormatRule {
	return &AzurermBatchCertificateInvalidFormatRule{
		resourceType:  "azurerm_batch_certificate",
		attributeName: "format",
		enum: []string{
			"Pfx",
			"Cer",
		},
	}
}

// Name returns the rule name
func (r *AzurermBatchCertificateInvalidFormatRule) Name() string {
	return "azurerm_batch_certificate_invalid_format"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermBatchCertificateInvalidFormatRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermBatchCertificateInvalidFormatRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermBatchCertificateInvalidFormatRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermBatchCertificateInvalidFormatRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as format`, truncateLongMessage(val)),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}
