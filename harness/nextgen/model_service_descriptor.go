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

type ServiceDescriptor struct {
	Index    int32                   `json:"index,omitempty"`
	Proto    *ServiceDescriptorProto `json:"proto,omitempty"`
	FullName string                  `json:"fullName,omitempty"`
	File     *FileDescriptor         `json:"file,omitempty"`
	Methods  []MethodDescriptor      `json:"methods,omitempty"`
	Options  *ServiceOptions         `json:"options,omitempty"`
	Name     string                  `json:"name,omitempty"`
}
