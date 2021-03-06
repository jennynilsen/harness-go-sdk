package cd

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/cd/graphql"
	"github.com/harness/harness-go-sdk/harness/helpers"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

func TestGetAwsCloudProviderById(t *testing.T) {
	c := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	cp, err := createAwsCloudProvider(expectedName)
	require.NoError(t, err)

	foundCP, err := c.CloudProviderClient.GetAwsCloudProviderById(cp.Id)
	require.NoError(t, err)
	require.NotNil(t, foundCP)
	require.Equal(t, cp.Id, foundCP.Id)

	err = c.CloudProviderClient.DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func TestGetAwsCloudProviderByName(t *testing.T) {
	c := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	cp, err := createAwsCloudProvider(expectedName)
	require.NoError(t, err)

	foundCP, err := c.CloudProviderClient.GetAwsCloudProviderByName(expectedName)
	require.NoError(t, err)
	require.NotNil(t, foundCP)
	require.Equal(t, expectedName, foundCP.Name)

	err = c.CloudProviderClient.DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func TestCreateAwsCloudProvider(t *testing.T) {
	c := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	cp, err := createAwsCloudProvider(expectedName)
	require.NoError(t, err)
	require.NotNil(t, cp)
	require.Equal(t, expectedName, cp.Name)

	err = c.CloudProviderClient.DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func TestUpdateAwsCloudProvider(t *testing.T) {
	c := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))
	updatedName := fmt.Sprintf("%s_updated", expectedName)

	cp, err := createAwsCloudProvider(expectedName)
	require.NoError(t, err)
	require.NotNil(t, cp)

	input := graphql.UpdateAwsCloudProviderInput{
		Name: updatedName,
	}
	updatedCP, err := c.CloudProviderClient.UpdateAwsCloudProvider(cp.Id, &input)
	require.NoError(t, err)
	require.NotNil(t, updatedCP)
	require.Equal(t, updatedName, updatedCP.Name)

	err = c.CloudProviderClient.DeleteCloudProvider(cp.Id)
	require.NoError(t, err)
}

func TestDeleteAwsCloudProvider(t *testing.T) {
	c := getClient()
	expectedName := fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(4))

	cp, err := createAwsCloudProvider(expectedName)
	require.NoError(t, err)
	require.NotNil(t, cp)

	err = c.CloudProviderClient.DeleteCloudProvider(cp.Id)
	require.NoError(t, err)

	foundCP, err := c.CloudProviderClient.GetAwsCloudProviderByName(expectedName)
	require.Error(t, err)
	require.Nil(t, foundCP)
}

func createAwsCloudProvider(name string) (*graphql.AwsCloudProvider, error) {

	c := getClient()
	expectedName := name

	secret, err := createEncryptedTextSecret(expectedName, helpers.TestEnvVars.AwsSecretAccessKey.Get())
	if err != nil {
		return nil, err
	}

	input := &graphql.AwsCloudProvider{}
	input.Name = expectedName
	input.CredentialsType = graphql.AwsCredentialsTypes.Manual
	input.ManualCredentials = &graphql.AwsManualCredentials{
		AccessKey:         helpers.TestEnvVars.AwsAccessKeyId.Get(),
		SecretKeySecretId: secret.Id,
	}

	return c.CloudProviderClient.CreateAwsCloudProvider(input)
}
