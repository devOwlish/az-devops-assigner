package azdevops

import (
	"log"
	"os"
)

const azPAT = "AZ_DEVOPS_PAT"
const azOrganization = "AZ_DEVOPS_ORG"

var requiredEnvVariables = []string{azPAT, azOrganization}

type AzCredentials struct {
	Token        string
	Organization string
}

var credentials *AzCredentials

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
