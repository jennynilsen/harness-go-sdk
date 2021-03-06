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

type UnitProgress struct {
	UnknownFields             *UnknownFieldSet       `json:"unknownFields,omitempty"`
	StartTime                 int64                  `json:"startTime,omitempty"`
	Initialized               bool                   `json:"initialized,omitempty"`
	EndTime                   int64                  `json:"endTime,omitempty"`
	Status                    string                 `json:"status,omitempty"`
	UnitName                  string                 `json:"unitName,omitempty"`
	UnitNameBytes             *ByteString            `json:"unitNameBytes,omitempty"`
	StatusValue               int32                  `json:"statusValue,omitempty"`
	SerializedSize            int32                  `json:"serializedSize,omitempty"`
	ParserForType             *ParserUnitProgress    `json:"parserForType,omitempty"`
	DefaultInstanceForType    *UnitProgress          `json:"defaultInstanceForType,omitempty"`
	AllFields                 map[string]interface{} `json:"allFields,omitempty"`
	InitializationErrorString string                 `json:"initializationErrorString,omitempty"`
	DescriptorForType         *Descriptor            `json:"descriptorForType,omitempty"`
	MemoizedSerializedSize    int32                  `json:"memoizedSerializedSize,omitempty"`
}
