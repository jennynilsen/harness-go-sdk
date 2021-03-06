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

// This is stage level info in Retry Failed Pipeline
type RetryStageInfo struct {
	Name       string `json:"name,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	// This is the Execution Status of the entity
	Status    string `json:"status,omitempty"`
	CreatedAt int64  `json:"createdAt,omitempty"`
	ParentId  string `json:"parentId,omitempty"`
	NextId    string `json:"nextId,omitempty"`
}
