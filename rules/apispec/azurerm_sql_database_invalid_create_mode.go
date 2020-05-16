// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermSQLDatabaseInvalidCreateModeRule checks the pattern is valid
type AzurermSQLDatabaseInvalidCreateModeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermSQLDatabaseInvalidCreateModeRule returns new rule with default attributes
func NewAzurermSQLDatabaseInvalidCreateModeRule() *AzurermSQLDatabaseInvalidCreateModeRule {
	return &AzurermSQLDatabaseInvalidCreateModeRule{
		resourceType:  "azurerm_sql_database",
		attributeName: "create_mode",
		enum: []string{
			"Copy",
			"Default",
			"NonReadableSecondary",
			"OnlineSecondary",
			"PointInTimeRestore",
			"Recovery",
			"Restore",
			"RestoreLongTermRetentionBackup",
		},
	}
}

// Name returns the rule name
func (r *AzurermSQLDatabaseInvalidCreateModeRule) Name() string {
	return "azurerm_sql_database_invalid_create_mode"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermSQLDatabaseInvalidCreateModeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermSQLDatabaseInvalidCreateModeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermSQLDatabaseInvalidCreateModeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermSQLDatabaseInvalidCreateModeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as create_mode`, truncateLongMessage(val)),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}
