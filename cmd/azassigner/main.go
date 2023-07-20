package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/devOwlish/az-devops-assigner/internal/azdevops"
	"github.com/jxskiss/mcli"
)

type args struct {
	Name                 string `cli:"#R -n, --name, User email or a group name you want give permissions to (test@example.com, TestGroup, etc.)"`
	Role                 string `cli:"#R -r, --role, Role you want to assign to the user or group (Administrator, Reader, etc.)"`
	VariableGroupPattern string `cli:"#R -p, --pattern, Case-insensitive part of a variable group name to filter by (VarGroup, var, vaRG matches 'var' pattern))"`
}

/*
The idea is to use the `azuredevops` library as much as possible to avoid "raw" API calls.
But the library is not complete and some endpoints are missing.

 1. Figure out, whether we have a user or a group passed as input
 2. Get the identityID of the user or group
    > Different endpoints and API versions
 3. Get all projects
 4. Get all variable groups in each project
 5. Filter variable groups by name
 6. Set role assignment for each variable group

What can be improved:
- Unit tests, but currently it's not worth investing time into mocking Azure DevOps API
- In case of more complex logic, the main() function can be split into more granular functions
- Error handling can be improved, but it's not worth it for a simple CLI tool
- The CLI tool can be extended to support more arguments, like --action=SET/DELETE to even further manage permissions
- More authentication options can be added, like AAD, principal, etc.
- With more complex logic, the CLI tool can be extended to support a config file with a list of users/groups to assign permissions to,
	so we'll get something like a declarative approach which allows us to neglect configuration drift.
*/

func main() {
	var (
		arguments args
		identity  string
		err       error
	)

	azdevops.InitCredentials()

	_, err = mcli.Parse(&arguments)
	if err != nil {
		log.Fatalf("Error parsing arguments: %s\n", err)
	}

	log.Printf("Running with username: %s, desired role: %s, variable group pattern: %s",
		arguments.Name, arguments.Role, arguments.VariableGroupPattern,
	)

	if strings.Contains(arguments.Name, "@") {
		identity, err = azdevops.GetUserIDByEmail(arguments.Name)
		log.Printf("Assigning user with ID: %s\n", identity)
	} else {
		identity, err = azdevops.GetGroupIDByName(arguments.Name)
		log.Printf("Assigning group with ID: %s\n", identity)
	}

	if err != nil {
		log.Fatalf("Error getting user or group ID: %s\n", err)
	}

	projects, err := azdevops.GetProjectIDs()
	if err != nil {
		log.Fatalf("Error getting projects: %s\n", err)
		return
	}

	for _, project := range projects {
		log.Printf("Project: %s\n", project)

		vgs, err := azdevops.GetVaribleGroupsIDByPattern(project, arguments.VariableGroupPattern)
		if err != nil {
			fmt.Printf("Error getting security roles: %s\n", err)
			return
		}

		for _, vg := range vgs {
			log.Printf("Project: %s, Variable Group: %d", project, vg)

			err = azdevops.SetRoleAssignment(project, vg, identity, arguments.Role)
			if err != nil {
				log.Fatalf("Error setting role assignment: %s\n", err)
			}
		}

	}

}
