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

// This has information about the password strength policy in Harness.
type PasswordStrengthPolicy struct {
	// This value is true if the password strength policy is enabled. Otherwise, it is false.
	Enabled bool `json:"enabled,omitempty"`
	// Minimum number of characters required in a password.
	MinNumberOfCharacters int32 `json:"minNumberOfCharacters,omitempty"`
	// Minimum number of uppercase characters required in a password.
	MinNumberOfUppercaseCharacters int32 `json:"minNumberOfUppercaseCharacters,omitempty"`
	// Minimum number of lower characters required in a password.
	MinNumberOfLowercaseCharacters int32 `json:"minNumberOfLowercaseCharacters,omitempty"`
	// Minimum number of special characters required in a password.
	MinNumberOfSpecialCharacters int32 `json:"minNumberOfSpecialCharacters,omitempty"`
	// Minimum number of digits required in a password.
	MinNumberOfDigits int32 `json:"minNumberOfDigits,omitempty"`
}
