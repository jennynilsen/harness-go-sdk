package cd

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/cd/graphql"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

func TestGetGcpCloudProviderById(t *testing.T) {
	c := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	cp, err := createGcpCloudProvider(expectedName)
	require.NoError(t, err)

	foundCP, err := c.CloudProviderClient.GetGcpCloudProviderById(cp.Id)
	require.NoError(t, err)
	require.NotNil(t, foundCP)
	require.Equal(t, cp.Id, foundCP.Id)

	err = c.CloudProviderClient.DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func TestGetGcpCloudProviderByName(t *testing.T) {
	c := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	cp, err := createGcpCloudProvider(expectedName)
	require.NoError(t, err)

	foundCP, err := c.CloudProviderClient.GetGcpCloudProviderByName(expectedName)
	require.NoError(t, err)
	require.NotNil(t, foundCP)
	require.Equal(t, expectedName, foundCP.Name)

	err = c.CloudProviderClient.DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func TestCreateGcpCloudProvider(t *testing.T) {
	c := getClient()

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	cp, err := createGcpCloudProvider(expectedName)
	require.NoError(t, err)
	require.NotNil(t, cp)
	require.Equal(t, expectedName, cp.Name)

	err = c.CloudProviderClient.DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func TestUpdateGcpCloudProvider(t *testing.T) {
	c := getClient()

	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	updatedName := fmt.Sprintf("%s_updated", expectedName)

	cp, err := createGcpCloudProvider(expectedName)
	require.NoError(t, err)
	require.NotNil(t, cp)

	input := &graphql.UpdateGcpCloudProviderInput{
		Name: updatedName,
	}

	updatedCP, err := c.CloudProviderClient.UpdateGcpCloudProvider(cp.Id, input)
	require.NoError(t, err)
	require.NotNil(t, updatedCP)
	require.Equal(t, updatedName, updatedCP.Name)

	err = c.CloudProviderClient.DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func createGcpCloudProvider(name string) (*graphql.GcpCloudProvider, error) {
	c := getClient()

	input := &graphql.GcpCloudProvider{}
	input.Name = name
	input.DelegateSelectors = []string{"Primary"}
	input.UseDelegate = true
	input.SkipValidation = true

	cp, err := c.CloudProviderClient.CreateGcpCloudProvider(input)
	if err != nil {
		return nil, err
	}

	return cp, nil
}
