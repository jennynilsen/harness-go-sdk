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

// This has API Key details and metadata.
type ApiKeyAggregate struct {
	ApiKey *ApiKey `json:"apiKey"`
	// This is the time at which API Key was created.
	CreatedAt int64 `json:"createdAt"`
	// This is the time at which API Key was last modified.
	LastModifiedAt int64 `json:"lastModifiedAt"`
	// The number of tokens within an API Key.
	TokensCount int32 `json:"tokensCount,omitempty"`
}
