package azdevops

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/taskagent"
)

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

func GetVaribleGroupsIDByPattern(project, pattern string) ([]int, error) {
	var ids []int

	client, err := newTaskAgentClient()
	if err != nil {
		return ids, fmt.Errorf("Failed to init coreClient: %w", err)
	}

	// ctx := context.TODO()
	groups, err := client.GetVariableGroups(context.Background(), taskagent.GetVariableGroupsArgs{
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

// // GET https://dev.azure.com/{organization}/{project}/_apis/distributedtask/variablegroups?api-version=7.0
// // https://learn.microsoft.com/en-us/rest/api/azure/devops/distributedtask/variablegroups/get-variable-groups?view=azure-devops-rest-7.0

// func GetVariableGroupsRaw(project string) ([]byte, error) {
// 	response, err := SendRequest(
// 		"distributedtask/variablegroups",
// 		"project",
// 		project,
// 		"GET",
// 		nil,
// 	)

// 	if err != nil {
// 		return []byte{}, err
// 	}

// 	// Process the response here
// 	return response.Body(), nil

// }

// func GetVariableGroupsCount(project string) (int, error) {
// 	response, err := GetVariableGroupsRaw(project)
// 	if err != nil {
// 		return -1, err
// 	}

// 	var tmp struct {
// 		Count int `json:"count"`
// 	}

// 	err = json.Unmarshal(response, &tmp)
// 	if err != nil {
// 		return -1, fmt.Errorf("failed to unmarshal variable group count response: %w", err)
// 	}

// 	return tmp.Count, nil
// }

// func GetVariableGroupIDsByName(project string) ([]int, error) {
// 	var ids []int

// 	response, err := GetVariableGroupsRaw(project)
// 	if err != nil {
// 		return ids, err
// 	}

// 	return ids, nil
// }
