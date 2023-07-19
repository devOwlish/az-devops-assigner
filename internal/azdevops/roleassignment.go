package azdevops

import "fmt"

// https://dev.azure.com/{organization}/_apis/securityroles/scopes/{scopeId}/roleassignments/resources/{resourceId}
// https://learn.microsoft.com/en-us/rest/api/azure/devops/securityroles/roleassignments/list?view=azure-devops-rest-7.0

func GetVariableGroupsRolesAssignments(project string, id int) error {
	response, err := SendRequest(
		fmt.Sprintf("securityroles/scopes/distributedtask.variablegroup/roleassignments/resources/%s$%d", project, id),
		"base",
		"",
		"GET",
		nil,
	)

	if err != nil {
		return err
	}
	// Process the response here
	fmt.Println(response.String())

	return nil
}

func SetRoleAssignment(project string, id int, userID, role string) error {
	// PUT https://dev.azure.com/{organization}/_apis/securityroles/scopes/{scopeId}/roleassignments/resources/{resourceId}/{identityId}?api-version=7.0-preview.1
	response, err := SendRequest(
		fmt.Sprintf("securityroles/scopes/distributedtask.variablegroup/roleassignments/resources/%s$%d", project, id),
		"base",
		"",
		"PUT",
		[]map[string]string{
			{
				"roleName": role,
				"userId":   userID,
			},
		},
	)
	if err != nil {
		return fmt.Errorf("failed to set role assignment: %w", err)
	}

	//TODO: Handle empty response
	fmt.Println(response.String())

	return nil
}
