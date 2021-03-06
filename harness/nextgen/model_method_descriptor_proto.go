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

type MethodDescriptorProto struct {
	UnknownFields             *UnknownFieldSet             `json:"unknownFields,omitempty"`
	ClientStreaming           bool                         `json:"clientStreaming,omitempty"`
	ServerStreaming           bool                         `json:"serverStreaming,omitempty"`
	Initialized               bool                         `json:"initialized,omitempty"`
	Options                   *MethodOptions               `json:"options,omitempty"`
	OptionsOrBuilder          *MethodOptionsOrBuilder      `json:"optionsOrBuilder,omitempty"`
	OutputType                string                       `json:"outputType,omitempty"`
	InputType                 string                       `json:"inputType,omitempty"`
	SerializedSize            int32                        `json:"serializedSize,omitempty"`
	ParserForType             *ParserMethodDescriptorProto `json:"parserForType,omitempty"`
	DefaultInstanceForType    *MethodDescriptorProto       `json:"defaultInstanceForType,omitempty"`
	NameBytes                 *ByteString                  `json:"nameBytes,omitempty"`
	InputTypeBytes            *ByteString                  `json:"inputTypeBytes,omitempty"`
	OutputTypeBytes           *ByteString                  `json:"outputTypeBytes,omitempty"`
	Name                      string                       `json:"name,omitempty"`
	AllFields                 map[string]interface{}       `json:"allFields,omitempty"`
	InitializationErrorString string                       `json:"initializationErrorString,omitempty"`
	DescriptorForType         *Descriptor                  `json:"descriptorForType,omitempty"`
	MemoizedSerializedSize    int32                        `json:"memoizedSerializedSize,omitempty"`
}
