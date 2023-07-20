package azdevops

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/taskagent"
)

// newTaskAgentClient initializes a new task agent client
func newTaskAgentClient() (taskagent.Client, error) {
	var client taskagent.Client

	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
	defer cancel()

	client, err := taskagent.NewClient(ctx, azuredevops.NewPatConnection(
		fmt.Sprintf(apiBase, credentials.Organization),
		credentials.Token,
	))
	if err != nil {
		return client, fmt.Errorf("error during client creation: %w", err)
	}

	return client, nil
}

// GetVaribleGroupsIDByPattern returns a list of variable group IDs that match a pattern;
// the pattern is case-insensitive and can be a part of the variable group name, thus
// "varG", "VaRGroup", "My Variable Group" are matched by a pattern "var"
func GetVaribleGroupsIDByPattern(project, pattern string) ([]int, error) {
	var ids []int

	client, err := newTaskAgentClient()
	if err != nil {
		return ids, fmt.Errorf("Failed to init coreClient: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
	defer cancel()

	groups, err := client.GetVariableGroups(ctx, taskagent.GetVariableGroupsArgs{
		Project: &project,
	})
	if err != nil {
		return ids, fmt.Errorf("failed to get variable groups: %w", err)
	}

	for _, group := range *groups {
		if group.Name != nil && strings.Contains(
			strings.ToLower(*group.Name), strings.ToLower(pattern)) {
			ids = append(ids, *group.Id)

			log.Printf("Found variable group %s with id %d", *group.Name, *group.Id)
		}
	}

	return ids, nil
}
