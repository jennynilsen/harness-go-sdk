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

type UninterpretedOptionOrBuilder struct {
	StringValue               *ByteString            `json:"stringValue,omitempty"`
	NameList                  []NamePart             `json:"nameList,omitempty"`
	NameOrBuilderList         []NamePartOrBuilder    `json:"nameOrBuilderList,omitempty"`
	IdentifierValue           string                 `json:"identifierValue,omitempty"`
	IdentifierValueBytes      *ByteString            `json:"identifierValueBytes,omitempty"`
	PositiveIntValue          int64                  `json:"positiveIntValue,omitempty"`
	NegativeIntValue          int64                  `json:"negativeIntValue,omitempty"`
	AggregateValue            string                 `json:"aggregateValue,omitempty"`
	AggregateValueBytes       *ByteString            `json:"aggregateValueBytes,omitempty"`
	DoubleValue               float64                `json:"doubleValue,omitempty"`
	NameCount                 int32                  `json:"nameCount,omitempty"`
	AllFields                 map[string]interface{} `json:"allFields,omitempty"`
	InitializationErrorString string                 `json:"initializationErrorString,omitempty"`
	DescriptorForType         *Descriptor            `json:"descriptorForType,omitempty"`
	UnknownFields             *UnknownFieldSet       `json:"unknownFields,omitempty"`
	DefaultInstanceForType    *Message               `json:"defaultInstanceForType,omitempty"`
	Initialized               bool                   `json:"initialized,omitempty"`
}
