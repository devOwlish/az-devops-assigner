package azassigner

import (
	"log"
	"strings"

	"github.com/devOwlish/az-devops-assigner/internal/azdevops"
	"github.com/jxskiss/mcli"
)

type args struct {
	//nolint: lll // We still can't have multiline tags
	Name string `cli:"#R -n, --name, User email or a group name you want give permissions to (test@example.com, TestGroup, etc.)"`
	Role string `cli:"#R -r, --role, Role you want to assign to the user or group (Administrator, Reader, etc.)"`
	//nolint: lll // We still can't have multiline tags
	VariableGroupPattern string `cli:"#R -p, --pattern, Case-insensitive part of a variable group name to filter by (VarGroup, var, vaRG matches 'var' pattern))"`
}

func Process() {
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
	}

	for _, project := range projects {
		log.Printf("Project: %s\n", project)

		vgs, varErr := azdevops.GetVaribleGroupsIDByPattern(project, arguments.VariableGroupPattern)

		if varErr != nil {
			log.Fatalf("Error getting security roles: %s\n", err)
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
