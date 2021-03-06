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

type AuditEventDto struct {
	AuditId            string                      `json:"auditId,omitempty"`
	InsertId           string                      `json:"insertId"`
	ResourceScope      *ResourcegroupResourceScope `json:"resourceScope"`
	HttpRequestInfo    *HttpRequestInfo            `json:"httpRequestInfo,omitempty"`
	RequestMetadata    *RequestMetadata            `json:"requestMetadata,omitempty"`
	Timestamp          int64                       `json:"timestamp"`
	AuthenticationInfo *AuthenticationInfoDto      `json:"authenticationInfo"`
	Module             string                      `json:"module"`
	Environment        *Environment                `json:"environment,omitempty"`
	Resource           *Resource                   `json:"resource"`
	YamlDiffRecord     *YamlDiffRecordDto          `json:"yamlDiffRecord,omitempty"`
	Action             string                      `json:"action"`
	AuditEventData     *AuditEventData             `json:"auditEventData,omitempty"`
	InternalInfo       map[string]string           `json:"internalInfo,omitempty"`
}
