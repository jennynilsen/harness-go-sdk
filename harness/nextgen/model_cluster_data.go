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

type ClusterData struct {
	Id                             string  `json:"id,omitempty"`
	Name                           string  `json:"name,omitempty"`
	Type_                          string  `json:"type,omitempty"`
	TotalCost                      float64 `json:"totalCost,omitempty"`
	IdleCost                       float64 `json:"idleCost,omitempty"`
	NetworkCost                    float64 `json:"networkCost,omitempty"`
	SystemCost                     float64 `json:"systemCost,omitempty"`
	CpuIdleCost                    float64 `json:"cpuIdleCost,omitempty"`
	CpuActualIdleCost              float64 `json:"cpuActualIdleCost,omitempty"`
	MemoryIdleCost                 float64 `json:"memoryIdleCost,omitempty"`
	MemoryActualIdleCost           float64 `json:"memoryActualIdleCost,omitempty"`
	CostTrend                      float64 `json:"costTrend,omitempty"`
	TrendType                      string  `json:"trendType,omitempty"`
	Region                         string  `json:"region,omitempty"`
	LaunchType                     string  `json:"launchType,omitempty"`
	CloudServiceName               string  `json:"cloudServiceName,omitempty"`
	TaskId                         string  `json:"taskId,omitempty"`
	WorkloadName                   string  `json:"workloadName,omitempty"`
	WorkloadType                   string  `json:"workloadType,omitempty"`
	Namespace                      string  `json:"namespace,omitempty"`
	ClusterType                    string  `json:"clusterType,omitempty"`
	ClusterId                      string  `json:"clusterId,omitempty"`
	InstanceId                     string  `json:"instanceId,omitempty"`
	InstanceName                   string  `json:"instanceName,omitempty"`
	InstanceType                   string  `json:"instanceType,omitempty"`
	Environment                    string  `json:"environment,omitempty"`
	CloudProvider                  string  `json:"cloudProvider,omitempty"`
	MaxCpuUtilization              float64 `json:"maxCpuUtilization,omitempty"`
	MaxMemoryUtilization           float64 `json:"maxMemoryUtilization,omitempty"`
	AvgCpuUtilization              float64 `json:"avgCpuUtilization,omitempty"`
	AvgMemoryUtilization           float64 `json:"avgMemoryUtilization,omitempty"`
	UnallocatedCost                float64 `json:"unallocatedCost,omitempty"`
	PrevBillingAmount              float64 `json:"prevBillingAmount,omitempty"`
	AppName                        string  `json:"appName,omitempty"`
	AppId                          string  `json:"appId,omitempty"`
	ServiceName                    string  `json:"serviceName,omitempty"`
	ServiceId                      string  `json:"serviceId,omitempty"`
	EnvId                          string  `json:"envId,omitempty"`
	EnvName                        string  `json:"envName,omitempty"`
	CloudProviderId                string  `json:"cloudProviderId,omitempty"`
	ClusterName                    string  `json:"clusterName,omitempty"`
	StorageCost                    float64 `json:"storageCost,omitempty"`
	MemoryBillingAmount            float64 `json:"memoryBillingAmount,omitempty"`
	CpuBillingAmount               float64 `json:"cpuBillingAmount,omitempty"`
	StorageUnallocatedCost         float64 `json:"storageUnallocatedCost,omitempty"`
	MemoryUnallocatedCost          float64 `json:"memoryUnallocatedCost,omitempty"`
	CpuUnallocatedCost             float64 `json:"cpuUnallocatedCost,omitempty"`
	StorageRequest                 float64 `json:"storageRequest,omitempty"`
	StorageUtilizationValue        float64 `json:"storageUtilizationValue,omitempty"`
	StorageActualIdleCost          float64 `json:"storageActualIdleCost,omitempty"`
	EfficiencyScore                int32   `json:"efficiencyScore,omitempty"`
	EfficiencyScoreTrendPercentage int32   `json:"efficiencyScoreTrendPercentage,omitempty"`
}
