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

type InterruptConfig struct {
	UnknownFields                 *UnknownFieldSet               `json:"unknownFields,omitempty"`
	Initialized                   bool                           `json:"initialized,omitempty"`
	RetryInterruptConfigOrBuilder *RetryInterruptConfigOrBuilder `json:"retryInterruptConfigOrBuilder,omitempty"`
	RetryInterruptConfig          *RetryInterruptConfig          `json:"retryInterruptConfig,omitempty"`
	IssuedBy                      *IssuedBy                      `json:"issuedBy,omitempty"`
	ConfigCase                    string                         `json:"configCase,omitempty"`
	SerializedSize                int32                          `json:"serializedSize,omitempty"`
	ParserForType                 *ParserInterruptConfig         `json:"parserForType,omitempty"`
	DefaultInstanceForType        *InterruptConfig               `json:"defaultInstanceForType,omitempty"`
	IssuedByOrBuilder             *IssuedByOrBuilder             `json:"issuedByOrBuilder,omitempty"`
	AllFields                     map[string]interface{}         `json:"allFields,omitempty"`
	InitializationErrorString     string                         `json:"initializationErrorString,omitempty"`
	DescriptorForType             *Descriptor                    `json:"descriptorForType,omitempty"`
	MemoizedSerializedSize        int32                          `json:"memoizedSerializedSize,omitempty"`
}