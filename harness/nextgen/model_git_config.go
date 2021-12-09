/*
 * CD NextGen API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

import "encoding/json"

// This contains details of the Generic Git connector
type GitConfig struct {
	Url               string                    `json:"url"`
	ValidationRepo    string                    `json:"validationRepo,omitempty"`
	BranchName        string                    `json:"branchName,omitempty"`
	DelegateSelectors []string                  `json:"delegateSelectors,omitempty"`
	Type_             GitAuthType               `json:"type"`
	Http              *GitHttpAuthenticationDto `json:"-"`
	Ssh               *GitSshAuthentication     `json:"-"`
	ConnectionType    string                    `json:"connectionType"`
	Spec              json.RawMessage           `json:"spec"`
}