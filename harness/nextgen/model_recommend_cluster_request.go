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

type RecommendClusterRequest struct {
	AllowBurst    bool     `json:"allowBurst,omitempty"`
	AllowOlderGen bool     `json:"allowOlderGen,omitempty"`
	Category      []string `json:"category,omitempty"`
	Excludes      []string `json:"excludes,omitempty"`
	Includes      []string `json:"includes,omitempty"`
	MaxNodes      int64    `json:"maxNodes,omitempty"`
	MinNodes      int64    `json:"minNodes,omitempty"`
	NetworkPerf   []string `json:"networkPerf,omitempty"`
	OnDemandPct   int64    `json:"onDemandPct,omitempty"`
	SameSize      bool     `json:"sameSize,omitempty"`
	SumCpu        float64  `json:"sumCpu,omitempty"`
	SumGpu        int64    `json:"sumGpu,omitempty"`
	SumMem        float64  `json:"sumMem,omitempty"`
	Zone          string   `json:"zone,omitempty"`
}
