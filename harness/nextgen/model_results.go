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

// This result object shows details of how many times a variation has been evaluated
type Results struct {
	// The number of times this variation has been returned in a evaluation
	Count int32 `json:"count"`
	// The unique variation identifier
	VariationIdentifier string `json:"variationIdentifier"`
	// The user friendly variation name
	VariationName string `json:"variationName"`
}
