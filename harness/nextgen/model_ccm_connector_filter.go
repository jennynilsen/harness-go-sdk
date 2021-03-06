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

type CcmConnectorFilter struct {
	FeaturesEnabled     []string `json:"featuresEnabled,omitempty"`
	AwsAccountId        string   `json:"awsAccountId,omitempty"`
	AzureSubscriptionId string   `json:"azureSubscriptionId,omitempty"`
	AzureTenantId       string   `json:"azureTenantId,omitempty"`
	GcpProjectId        string   `json:"gcpProjectId,omitempty"`
	K8sConnectorRef     string   `json:"k8sConnectorRef,omitempty"`
}
