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

// This contains the feature restriction details object defined in Harness
type FeatureRestrictionDetails struct {
	Name            string       `json:"name,omitempty"`
	Description     string       `json:"description,omitempty"`
	ModuleType      string       `json:"moduleType,omitempty"`
	Allowed         bool         `json:"allowed,omitempty"`
	RestrictionType string       `json:"restrictionType,omitempty"`
	Restriction     *Restriction `json:"restriction,omitempty"`
}