package main

import (
	"fmt"
	"log"

	"github.com/devOwlish/az-devops-assigner/internal/azdevops"
)

func main() {
	azdevops.InitCredentials()

	// TODO: CLI/Env
	matchPattern := "var"
	userEmail := "test2@example.com"
	userRole := "Reader"

	userID, err := azdevops.GetUserIDByEmail(userEmail)
	if err != nil {
		log.Fatalf("Error getting user ID: %s\n", err)
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
			fmt.Println("Project: ", project, "Variable Group: ", vg)

			err = azdevops.SetRoleAssignment(project, vg, userID, userRole)
			if err != nil {
				log.Fatalf("Error setting role assignment: %s\n", err)
			}
		}

	}

}
