package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/devOwlish/az-devops-assigner/internal/azdevops"
)

func main() {
	azdevops.InitCredentials()

	// TODO: CLI/Env
	matchPattern := "var"
	// userEmail := "test2@example.com"
	userEmail := "TestGroup"
	userRole := "Reader"
	var identity string
	var err error

	if strings.Contains(userEmail, "@") {
		identity, err = azdevops.GetUserIDByEmail(userEmail)
		log.Printf("Assigning user with ID: %s\n", identity)
	} else {
		identity, err = azdevops.GetGroupIDByName("TestGroup")
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

		vgs, err := azdevops.GetVaribleGroupsIDByPattern(project, matchPattern)
		if err != nil {
			fmt.Printf("Error getting security roles: %s\n", err)
			return
		}

		for _, vg := range vgs {
			log.Printf("Project: %s, Variable Group: %d", project, vg)

			err = azdevops.SetRoleAssignment(project, vg, identity, userRole)
			if err != nil {
				log.Fatalf("Error setting role assignment: %s\n", err)
			}
		}

	}

}
