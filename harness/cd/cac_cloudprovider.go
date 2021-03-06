package cd

import (
	"errors"

	"github.com/harness/harness-go-sdk/harness/cd/cac"
	"github.com/harness/harness-go-sdk/harness/utils"
)

func (c *ConfigAsCodeClient) UpsertSpotInstCloudProvider(input *cac.SpotInstCloudProvider) (*cac.SpotInstCloudProvider, error) {
	c.ApiClient.Log.Debugf("Upsert Spot cloud provider %s", input.Name)
	out := &cac.SpotInstCloudProvider{}
	err := c.UpsertCloudProvider(input, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c *ConfigAsCodeClient) UpsertPcfCloudProvider(input *cac.PcfCloudProvider) (*cac.PcfCloudProvider, error) {
	c.ApiClient.Log.Debugf("Upsert PCF cloud provider %s", input.Name)
	out := &cac.PcfCloudProvider{}
	err := c.UpsertCloudProvider(input, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c *ConfigAsCodeClient) UpsertKubernetesCloudProvider(input *cac.KubernetesCloudProvider) (*cac.KubernetesCloudProvider, error) {
	c.ApiClient.Log.Debugf("Upsert Kubernetes cloud provider %s", input.Name)
	out := &cac.KubernetesCloudProvider{}
	err := c.UpsertCloudProvider(input, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c *ConfigAsCodeClient) UpsertAzureCloudProvider(input *cac.AzureCloudProvider) (*cac.AzureCloudProvider, error) {
	c.ApiClient.Log.Debugf("Upsert Azure cloud provider %s", input.Name)
	out := &cac.AzureCloudProvider{}
	err := c.UpsertCloudProvider(input, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c *ConfigAsCodeClient) UpsertGcpCloudProvider(input *cac.GcpCloudProvider) (*cac.GcpCloudProvider, error) {
	c.ApiClient.Log.Debugf("Upsert GCP cloud provider %s", input.Name)
	out := &cac.GcpCloudProvider{}
	err := c.UpsertCloudProvider(input, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c *ConfigAsCodeClient) UpsertPhysicalDataCenterCloudProvider(input *cac.PhysicalDatacenterCloudProvider) (*cac.PhysicalDatacenterCloudProvider, error) {
	c.ApiClient.Log.Debugf("Upsert Datacenter cloud provider %s", input.Name)
	out := &cac.PhysicalDatacenterCloudProvider{}
	err := c.UpsertCloudProvider(input, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c *ConfigAsCodeClient) UpsertAwsCloudProvider(input *cac.AwsCloudProvider) (*cac.AwsCloudProvider, error) {
	c.ApiClient.Log.Debugf("Upsert AWS cloud provider %s", input.Name)
	out := &cac.AwsCloudProvider{}
	err := c.UpsertCloudProvider(input, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c *ConfigAsCodeClient) UpsertCloudProvider(input interface{}, output interface{}) error {
	if input == nil {
		return errors.New("cannot create cloud provider. input is nil")
	}

	name, ok := utils.TryGetFieldValue(input, "Name")
	if !ok || name == "" {
		return errors.New("expected cloud provider to have Name field set")
	}

	filePath := cac.GetCloudProviderYamlPath(name.(string))
	err := c.ApiClient.ConfigAsCodeClient.UpsertObject(input, filePath, output)
	if err != nil {
		return err
	}

	return nil
}

func (c *ConfigAsCodeClient) GetCloudProviderById(providerId string, out interface{}) error {
	c.ApiClient.Log.Debugf("Get cloud provider by id %s", providerId)
	rootItem, err := c.GetDirectoryTree("")
	if err != nil {
		return err
	}

	i := FindConfigAsCodeItemByUUID(rootItem, providerId)
	if i == nil {
		c.ApiClient.Log.Debugf("cannot find cloud provider with id: " + providerId)
		return nil
	}

	return c.ParseObject(i, cac.YamlPath(i.DirectoryPath.Path), "", out)
}

func (c *ConfigAsCodeClient) GetCloudProviderByName(name string, obj interface{}) error {
	c.ApiClient.Log.Debugf("Get cloud provider by name %s", name)
	filePath := cac.GetCloudProviderYamlPath(name)
	return c.FindObjectByPath("", filePath, obj)
}

// func (c *ConfigAsCodeClient) DeleteCloudProvider(name string) error {
// 	filePath := cac.GetCloudProviderYamlPath(name)
// 	return c.DeleteEntity(filePath)
// }
