package azdevops

import (
	"context"
	"fmt"

	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
)

func newCoreClient() (core.Client, error) {
	var client core.Client

	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)

	defer cancel()

	client, err := core.NewClient(ctx, azuredevops.NewPatConnection(
		fmt.Sprintf(apiBase, credentials.Organization),
		credentials.Token,
	))
	if err != nil {
		return client, fmt.Errorf("error during client creation: %w", err)
	}

	return client, nil
}

func GetProjectIDs() ([]string, error) {
	var projects []string

	client, err := newCoreClient()
	if err != nil {
		return projects, fmt.Errorf("Failed to init coreClient: %w", err)
	}

	ctx := context.TODO()

	responseValue, err := client.GetProjects(ctx, core.GetProjectsArgs{})
	if err != nil {
		return projects, fmt.Errorf("failed to get projects: %w", err)
	}

	for responseValue != nil {
		for _, teamProjectReference := range responseValue.Value {
			projects = append(projects, teamProjectReference.Id.String())
		}

		if responseValue.ContinuationToken != "" {
			projectArgs := core.GetProjectsArgs{
				ContinuationToken: &responseValue.ContinuationToken,
			}

			responseValue, err = client.GetProjects(ctx, projectArgs)
			if err != nil {
				return projects, fmt.Errorf("failed to get projects: %w", err)
			}
		} else {
			responseValue = nil
		}
	}

	return projects, nil
}
