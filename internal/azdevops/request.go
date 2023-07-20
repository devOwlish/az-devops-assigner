package azdevops

import (
	"encoding/base64"
	"fmt"
	"log"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	// Default base API URL
	apiBase = "https://dev.azure.com/%s"
	// Project-scoped base API URL
	apiProject = "https://dev.azure.com/%s/%s"
	// User-scoped base API URL
	apiUser = "https://vsaex.dev.azure.com/%s"
	// Default context timeout for API requests
	contextTimeout = 10 * time.Second
	// Default API version
	apiVersion = "7.1-preview.1"
)

// SendRequest constructs and sends a request to the Azure DevOps API;
// can be used in case of a missing endpoint implementation in the Azure DevOps Go SDK
func SendRequest(route, urlType, project, method string, payload interface{}) (*resty.Response, error) {
	var (
		url string
		api string = apiVersion
	)

	switch urlType {
	case "project":
		url = fmt.Sprintf(apiProject, credentials.Organization, project)
	case "group":
		url = fmt.Sprintf(apiUser, credentials.Organization)
	case "user":
		url = fmt.Sprintf(apiUser, credentials.Organization)
		// Endpoint is not available in API newer than 5.0-preview.2
		// TODO: There should be a another endpoint compatible with modern API versions
		api = "5.0-preview.2"
	default:
		url = fmt.Sprintf(apiBase, credentials.Organization)
	}

	// Construct URL
	url = url + "/_apis/" + route

	log.Printf(">> [%s] %s -> %s", api, method, url)

	client := resty.New()

	// PAT must be base64 econded.
	// ':' is required for basic auth, but PATs don't have a username
	auth := base64.StdEncoding.EncodeToString(
		[]byte(fmt.Sprintf(":%s", credentials.Token)))

	response, err := client.R().
		SetHeaders(map[string]string{
			"Authorization": "Basic " + auth,
			"Content-Type":  "application/json",
		}).
		SetBody(payload).
		SetQueryParam("api-version", api).
		Execute(method, url)

	if err != nil {
		return nil, err
	}

	return response, nil
}
