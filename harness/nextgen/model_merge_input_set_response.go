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

// View of the Response of Merging of Input Sets of a Pipeline
type MergeInputSetResponse struct {
	PipelineYaml         string                      `json:"pipelineYaml,omitempty"`
	CompletePipelineYaml string                      `json:"completePipelineYaml,omitempty"`
	IsErrorResponse      bool                        `json:"isErrorResponse,omitempty"`
	InputSetErrorWrapper *InputSetErrorWrapperDtopms `json:"inputSetErrorWrapper,omitempty"`
	ErrorResponse        bool                        `json:"errorResponse,omitempty"`
}
