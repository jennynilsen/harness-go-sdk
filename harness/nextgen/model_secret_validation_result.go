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

// This has validation details for the Secret defined in Harness.
type SecretValidationResult struct {
	// This has the validation status for a secret. It is Success, if validation is successful, else the status is Failed.
	Success bool `json:"success,omitempty"`
	// This is the error message when validation for secret fails.
	Message string `json:"message,omitempty"`
}
