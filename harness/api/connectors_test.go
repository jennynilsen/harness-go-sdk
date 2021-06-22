package api

import (
	"fmt"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

func createGitConnector(name string) (*graphql.GitConnector, error) {
	secret, err := createEncryptedTextSecret(name, "foo")
	if err != nil {
		return nil, err
	}

	input := &graphql.GitConnectorInput{
		Name:             name,
		Url:              "https://github.com/micahlmartin/harness-demo",
		UrlType:          graphql.GitUrlTypes.Repo,
		PasswordSecretId: secret.Id,
		UserName:         "testuser",
	}

	client := getClient()
	conn, err := client.Connectors().CreateGitConnector(input)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func TestGetGitConnectorById(t *testing.T) {

	// Setup
	client := getClient()
	expectedName := fmt.Sprintf("%s-%s", t.Name(), utils.RandStringBytes(12))
	conn, err := createGitConnector(expectedName)
	require.NoError(t, err)

	connLookup, err := client.Connectors().GetGitConnectorById(conn.Id)

	// Verify
	require.NoError(t, err)
	require.Equal(t, conn.Id, connLookup.Id)
	require.Equal(t, conn.Name, connLookup.Name)

	// Cleanup
	err = client.Connectors().DeleteConnector(conn.Id)
	require.NoError(t, err)
}

func TestCreateGitConnector(t *testing.T) {

	expectedName := fmt.Sprintf("%s-%s", t.Name(), utils.RandStringBytes(12))

	secret, err := createEncryptedTextSecret(expectedName, "foo")
	require.NoError(t, err)

	// Create application
	client := getClient()
	name := fmt.Sprintf("%s-%s", t.Name(), utils.RandStringBytes(12))
	input := &graphql.GitConnectorInput{
		Name:             name,
		Url:              "https://github.com/micahlmartin/harness-demo",
		UrlType:          graphql.GitUrlTypes.Repo,
		PasswordSecretId: secret.Id,
		UserName:         "testuser",
	}

	conn, err := client.Connectors().CreateGitConnector(input)

	// Verify
	require.NoError(t, err)
	require.NotEmpty(t, conn.Id)
	require.Equal(t, name, conn.Name)

	// Cleanup
	err = client.Connectors().DeleteConnector(conn.Id)
	require.NoError(t, err)
}

func TestCreateGitConnector_ssh(t *testing.T) {
	t.Skip("Not yet implemented")
}

func TestUpdateGitConnector(t *testing.T) {

	// setup
	client := getClient()
	expectedName := fmt.Sprintf("%s-%s", t.Name(), utils.RandStringBytes(12))
	updatedName := fmt.Sprintf("%s_updated", expectedName)
	conn, err := createGitConnector(expectedName)
	require.NoError(t, err)

	updateInput := &graphql.GitConnectorInput{
		Name: updatedName,
	}

	// Verify
	updatedConn, err := client.Connectors().UpdateGitConnector(conn.Id, updateInput)
	require.NoError(t, err)
	require.Equal(t, updatedName, updatedConn.Name)
	require.Equal(t, conn.Id, updatedConn.Id)

	// Cleanup
	err = client.Connectors().DeleteConnector(conn.Id)
	require.NoError(t, err)
}