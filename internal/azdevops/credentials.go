package azdevops

import (
	"log"
	"os"
)

// Azure DevOps Personal Access Token Environment Variable name.
const azPAT = "AZ_DEVOPS_PAT"

// Azure DevOps Organization Environment Variable name.
const azOrganization = "AZ_DEVOPS_ORG"

// List of required environment variables.
//
//nolint:gochecknoglobals // This is basically a constant
var requiredEnvVariables = []string{azPAT, azOrganization}

// Structure that holds the credentials.
type AzCredentials struct {
	Token        string
	Organization string
}

// Module-wide singleton of AzCredentials.
//
//nolint:gochecknoglobals // We need a singleton here
var credentials *AzCredentials

// checks if the required environment variables are set and initializes the singleton instance.
func InitCredentials() {
	if credentials != nil {
		return
	}

	for _, variable := range requiredEnvVariables {
		if os.Getenv(variable) == "" {
			log.Fatalf("required variable %s is empty, exiting.", variable)
		}
	}

	credentials = &AzCredentials{
		Token:        os.Getenv(azPAT),
		Organization: os.Getenv(azOrganization),
	}
}
