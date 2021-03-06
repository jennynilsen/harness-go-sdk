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

// This is the Project Entity details defined in Harness
type Project struct {
	// Organization Identifier for the Entity
	OrgIdentifier string `json:"orgIdentifier,omitempty"`
	// Project Identifier for the Entity
	Identifier string `json:"identifier,omitempty"`
	// Project Name for the entity
	Name string `json:"name,omitempty"`
	// Color
	Color string `json:"color,omitempty"`
	// List of modules
	Modules []string `json:"modules,omitempty"`
	// Description
	Description string `json:"description,omitempty"`
	// Tags
	Tags map[string]string `json:"tags,omitempty"`
}
