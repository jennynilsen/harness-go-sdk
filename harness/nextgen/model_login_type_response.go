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

type LoginTypeResponse struct {
	AuthenticationMechanism string      `json:"authenticationMechanism,omitempty"`
	SSORequest              *SsoRequest `json:"SSORequest,omitempty"`
	IsOauthEnabled          bool        `json:"isOauthEnabled,omitempty"`
	ShowCaptcha             bool        `json:"showCaptcha,omitempty"`
	DefaultExperience       string      `json:"defaultExperience,omitempty"`
	OauthEnabled            bool        `json:"oauthEnabled,omitempty"`
	Ssorequest              *SsoRequest `json:"ssorequest,omitempty"`
}
