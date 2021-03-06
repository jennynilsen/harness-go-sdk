package cd

import (
	"fmt"

	"github.com/harness/harness-go-sdk/harness/cd/graphql"
)

func (c *CloudProviderClient) GetGcpCloudProviderById(id string) (*graphql.GcpCloudProvider, error) {
	cp := &graphql.GcpCloudProvider{}
	err := c.getCloudProviderById(id, getGcpCloudProviderFields(), &cp)
	if err != nil {
		return nil, err
	}

	return cp, nil
}

func (c *CloudProviderClient) GetGcpCloudProviderByName(name string) (*graphql.GcpCloudProvider, error) {
	cp := &graphql.GcpCloudProvider{}
	err := c.getCloudProviderByName(name, getGcpCloudProviderFields(), &cp)
	if err != nil {
		return nil, err
	}

	return cp, nil
}

func (c *CloudProviderClient) CreateGcpCloudProvider(provider *graphql.GcpCloudProvider) (*graphql.GcpCloudProvider, error) {
	input := &graphql.CreateCloudProviderInput{
		CloudProviderType: graphql.CloudProviderTypes.Gcp,
		GCPCloudProvider:  provider,
	}

	resp := &graphql.GcpCloudProvider{}
	err := c.createCloudProvider(input, getGcpCloudProviderFields(), resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CloudProviderClient) UpdateGcpCloudProvider(id string, cp *graphql.UpdateGcpCloudProviderInput) (*graphql.GcpCloudProvider, error) {
	input := &graphql.UpdateCloudProvider{
		GcpCloudProvider:  cp,
		CloudProviderType: &graphql.CloudProviderTypes.Gcp,
		CloudProviderId:   id,
	}

	resp := &graphql.GcpCloudProvider{}
	err := c.updateCloudProvider(input, getGcpCloudProviderFields(), resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func getGcpCloudProviderFields() string {
	return fmt.Sprintf(`
		%[1]s
		... on GcpCloudProvider {
			delegateSelectors
		}
	`, commonCloudProviderFields)
}
