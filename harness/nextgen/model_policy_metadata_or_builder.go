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

type PolicyMetadataOrBuilder struct {
	Identifier                string                 `json:"identifier,omitempty"`
	Severity                  string                 `json:"severity,omitempty"`
	Error_                    string                 `json:"error,omitempty"`
	Updated                   int64                  `json:"updated,omitempty"`
	Created                   int64                  `json:"created,omitempty"`
	Status                    string                 `json:"status,omitempty"`
	PolicyName                string                 `json:"policyName,omitempty"`
	PolicyId                  string                 `json:"policyId,omitempty"`
	PolicyIdBytes             *ByteString            `json:"policyIdBytes,omitempty"`
	PolicyNameBytes           *ByteString            `json:"policyNameBytes,omitempty"`
	SeverityBytes             *ByteString            `json:"severityBytes,omitempty"`
	DenyMessagesList          []string               `json:"denyMessagesList,omitempty"`
	DenyMessagesCount         int32                  `json:"denyMessagesCount,omitempty"`
	StatusBytes               *ByteString            `json:"statusBytes,omitempty"`
	AccountId                 string                 `json:"accountId,omitempty"`
	AccountIdBytes            *ByteString            `json:"accountIdBytes,omitempty"`
	OrgIdBytes                *ByteString            `json:"orgIdBytes,omitempty"`
	ProjectId                 string                 `json:"projectId,omitempty"`
	ProjectIdBytes            *ByteString            `json:"projectIdBytes,omitempty"`
	ErrorBytes                *ByteString            `json:"errorBytes,omitempty"`
	OrgId                     string                 `json:"orgId,omitempty"`
	IdentifierBytes           *ByteString            `json:"identifierBytes,omitempty"`
	AllFields                 map[string]interface{} `json:"allFields,omitempty"`
	UnknownFields             *UnknownFieldSet       `json:"unknownFields,omitempty"`
	DefaultInstanceForType    *Message               `json:"defaultInstanceForType,omitempty"`
	InitializationErrorString string                 `json:"initializationErrorString,omitempty"`
	DescriptorForType         *Descriptor            `json:"descriptorForType,omitempty"`
	Initialized               bool                   `json:"initialized,omitempty"`
}