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

// This is the view for a particular Execution in Retry History
type ExecutionInfo struct {
	Uuid    string `json:"uuid,omitempty"`
	StartTs int64  `json:"startTs,omitempty"`
	EndTs   int64  `json:"endTs,omitempty"`
	// This is the Execution Status of the entity
	Status string `json:"status,omitempty"`
}
