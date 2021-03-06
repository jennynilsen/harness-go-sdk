package cd

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/cd/graphql"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

func TestGetKubernetesCloudProviderById(t *testing.T) {
	c := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	cp, err := createKubernetesCloudProvider(expectedName)
	require.NoError(t, err)

	foundCP, err := c.CloudProviderClient.GetKubernetesCloudProviderById(cp.Id)
	require.NoError(t, err)
	require.NotNil(t, foundCP)
	require.Equal(t, cp.Id, foundCP.Id)

	err = c.CloudProviderClient.DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func TestGetKubernetesCloudProviderByName(t *testing.T) {
	c := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	cp, err := createKubernetesCloudProvider(expectedName)
	require.NoError(t, err)

	foundCP, err := c.CloudProviderClient.GetKubernetesCloudProviderByName(expectedName)
	require.NoError(t, err)
	require.NotNil(t, foundCP)
	require.Equal(t, expectedName, foundCP.Name)

	err = c.CloudProviderClient.DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func TestCreateKubernetesCloudProvider(t *testing.T) {
	c := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	cp, err := createKubernetesCloudProvider(expectedName)
	require.NoError(t, err)
	require.NotNil(t, cp)
	require.Equal(t, expectedName, cp.Name)

	err = c.CloudProviderClient.DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func TestUpdateKubernetesCloudProvider(t *testing.T) {
	c := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	updatedName := fmt.Sprintf("%s_updated", expectedName)
	cp, err := createKubernetesCloudProvider(expectedName)
	require.NoError(t, err)
	require.NotNil(t, cp)

	input := &graphql.UpdateKubernetesCloudProviderInput{
		Name:               updatedName,
		ClusterDetailsType: graphql.ClusterDetailsTypes.InheritClusterDetails,
		InheritClusterDetails: &graphql.InheritClusterDetails{
			DelegateSelectors: []string{"Primary"},
		},
		SkipValidation: true,
	}

	updatedCP, err := c.CloudProviderClient.UpdateKubernetesCloudProvider(cp.Id, input)
	require.NoError(t, err)
	require.NotNil(t, updatedCP)
	require.Equal(t, updatedName, updatedCP.Name)

	err = c.CloudProviderClient.DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func createKubernetesCloudProvider(name string) (*graphql.KubernetesCloudProvider, error) {
	c := getClient()

	input := &graphql.KubernetesCloudProvider{}
	input.Name = name
	input.ClusterDetailsType = graphql.ClusterDetailsTypes.InheritClusterDetails
	input.InheritClusterDetails = &graphql.InheritClusterDetails{
		DelegateSelectors: []string{"Primary"},
	}
	input.SkipValidation = true

	cp, err := c.CloudProviderClient.CreateKubernetesCloudProvider(input)
	if err != nil {
		return nil, err
	}

	return cp, nil
}
