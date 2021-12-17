/*
 * CD NextGen API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

type RoleAssignmentFilter struct {
	ResourceGroupFilter  []string    `json:"resourceGroupFilter,omitempty"`
	RoleFilter           []string    `json:"roleFilter,omitempty"`
	PrincipalTypeFilter  []string    `json:"principalTypeFilter,omitempty"`
	PrincipalFilter      []Principal `json:"principalFilter,omitempty"`
	HarnessManagedFilter []bool      `json:"harnessManagedFilter,omitempty"`
	DisabledFilter       []bool      `json:"disabledFilter,omitempty"`
}