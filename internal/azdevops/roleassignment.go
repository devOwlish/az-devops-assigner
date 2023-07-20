package azdevops

import (
	"encoding/json"
	"fmt"
	"log"
)

// API Endpoint: https://dev.azure.com/{organization}/_apis/securityroles/scopes/{scopeId}/roleassignments/resources/{resourceId}
// API Reference: https://learn.microsoft.com/en-us/rest/api/azure/devops/securityroles/roleassignments/list?view=azure-devops-rest-7.0

// GetVariableGroupsRolesAssignments returns a list of role assignments for a variable group
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

	fmt.Println(response.String())

	return nil
}

// SetRoleAssignment sets a role assignment for a variable group.
// Role must be set to a valid role name, i.e. "Administrator", "Contributor", "Reader".
// UserID is a UUID of the user or group to assign the role to.
// NOTE: OriginID/InternalID/ID from graph.Users() doesn't work, only the UUID.
func SetRoleAssignment(project string, id int, userID, role string) error {
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

	// If the response doesn't contain a count of modified resources or it's equal to zero,
	// we can assume that the request has failed, even if it returns a 200 OK code.
	// Otherwise, the response contains created assignment resource.
	var s struct {
		Count int `json:"count"`
	}

	err = json.Unmarshal(response.Body(), &s)
	if err != nil {
		return fmt.Errorf("failed to unmarshal role assignment response: %w", err)
	}

	if s.Count < 1 {
		return fmt.Errorf("failed to set role assignment, empty response %s", response.String())
	}

	log.Printf("Affected resourced: %d", s.Count)

	return nil
}
