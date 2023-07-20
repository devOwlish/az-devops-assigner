package azdevops

import (
	"context"
	"fmt"
	"strings"

	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/identity"
)

// newIdentityClient initializes a new identity client.
func newIdentityClient() (identity.Client, error) {
	var client identity.Client

	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
	defer cancel()

	client, err := identity.NewClient(ctx, azuredevops.NewPatConnection(
		fmt.Sprintf(apiBase, credentials.Organization),
		credentials.Token,
	))
	if err != nil {
		return client, fmt.Errorf("error during client creation: %w", err)
	}

	return client, nil
}

// GetGroupIDByName lists all accessible groups and returns the ID of the group that
// matches a part of the provided name; i.e. [k8sowl]/TestGroup matches a TestGroup name.
func GetGroupIDByName(name string) (string, error) {
	client, err := newIdentityClient()
	if err != nil {
		return "", fmt.Errorf("failed to init coreClient: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
	defer cancel()

	responseValue, err := client.ListGroups(ctx, identity.ListGroupsArgs{})
	if err != nil {
		return "", fmt.Errorf("failed to get groups: %w", err)
	}

	for _, identity := range *responseValue {
		if strings.Contains(*identity.ProviderDisplayName, name) {
			return identity.Id.String(), nil
		}
	}

	return "", fmt.Errorf("failed to find group with name: %s", name)
}
