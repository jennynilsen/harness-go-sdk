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

type UnknownFieldSet struct {
	Initialized                bool             `json:"initialized,omitempty"`
	SerializedSizeAsMessageSet int32            `json:"serializedSizeAsMessageSet,omitempty"`
	SerializedSize             int32            `json:"serializedSize,omitempty"`
	ParserForType              *Parser          `json:"parserForType,omitempty"`
	DefaultInstanceForType     *UnknownFieldSet `json:"defaultInstanceForType,omitempty"`
}