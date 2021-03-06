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

type YamlOutputProperties struct {
	UnknownFields             *UnknownFieldSet            `json:"unknownFields,omitempty"`
	LocalName                 string                      `json:"localName,omitempty"`
	Initialized               bool                        `json:"initialized,omitempty"`
	FqnBytes                  *ByteString                 `json:"fqnBytes,omitempty"`
	LocalNameBytes            *ByteString                 `json:"localNameBytes,omitempty"`
	SerializedSize            int32                       `json:"serializedSize,omitempty"`
	Fqn                       string                      `json:"fqn,omitempty"`
	ParserForType             *ParserYamlOutputProperties `json:"parserForType,omitempty"`
	DefaultInstanceForType    *YamlOutputProperties       `json:"defaultInstanceForType,omitempty"`
	AllFields                 map[string]interface{}      `json:"allFields,omitempty"`
	InitializationErrorString string                      `json:"initializationErrorString,omitempty"`
	DescriptorForType         *Descriptor                 `json:"descriptorForType,omitempty"`
	MemoizedSerializedSize    int32                       `json:"memoizedSerializedSize,omitempty"`
}
