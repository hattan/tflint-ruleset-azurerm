// This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT

package apispec

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-azurerm/project"
)

// AzurermVirtualNetworkGatewayConnectionInvalidConnectionProtocolRule checks the pattern is valid
type AzurermVirtualNetworkGatewayConnectionInvalidConnectionProtocolRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAzurermVirtualNetworkGatewayConnectionInvalidConnectionProtocolRule returns new rule with default attributes
func NewAzurermVirtualNetworkGatewayConnectionInvalidConnectionProtocolRule() *AzurermVirtualNetworkGatewayConnectionInvalidConnectionProtocolRule {
	return &AzurermVirtualNetworkGatewayConnectionInvalidConnectionProtocolRule{
		resourceType:  "azurerm_virtual_network_gateway_connection",
		attributeName: "connection_protocol",
		enum: []string{
			"IKEv2",
			"IKEv1",
		},
	}
}

// Name returns the rule name
func (r *AzurermVirtualNetworkGatewayConnectionInvalidConnectionProtocolRule) Name() string {
	return "azurerm_virtual_network_gateway_connection_invalid_connection_protocol"
}

// Enabled returns whether the rule is enabled by default
func (r *AzurermVirtualNetworkGatewayConnectionInvalidConnectionProtocolRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AzurermVirtualNetworkGatewayConnectionInvalidConnectionProtocolRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AzurermVirtualNetworkGatewayConnectionInvalidConnectionProtocolRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks the pattern is valid
func (r *AzurermVirtualNetworkGatewayConnectionInvalidConnectionProtocolRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as connection_protocol`, truncateLongMessage(val)),
					attribute.Expr.Range(),
					tflint.Metadata{Expr: attribute.Expr},
				)
			}
			return nil
		})
	})
}
