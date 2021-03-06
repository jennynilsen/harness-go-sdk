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

// Common filter for all Cloud Cost Recommendation APIs.
type K8sRecommendationFilter struct {
	Ids           []string `json:"ids,omitempty"`
	Names         []string `json:"names,omitempty"`
	Namespaces    []string `json:"namespaces,omitempty"`
	ClusterNames  []string `json:"clusterNames,omitempty"`
	ResourceTypes []string `json:"resourceTypes,omitempty"`
	// Get Recommendations for a perspective
	PerspectiveFilters []QlceViewFilterWrapper `json:"perspectiveFilters,omitempty"`
	MinSaving          float64                 `json:"minSaving,omitempty"`
	MinCost            float64                 `json:"minCost,omitempty"`
	Offset             int64                   `json:"offset,omitempty"`
	Limit              int64                   `json:"limit,omitempty"`
}
