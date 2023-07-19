package azdevops

import (
	"encoding/base64"
	"fmt"
	"log"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	apiBase        = "https://dev.azure.com/%s"
	apiProject     = "https://dev.azure.com/%s/%s"
	apiUser        = "https://vsaex.dev.azure.com/%s"
	contextTimeout = 10 * time.Second
	apiVersion     = "7.1-preview.1"
)

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
		api = "5.0-preview.2"
	default:
		url = fmt.Sprintf(apiBase, credentials.Organization)
	}

	url = url + "/_apis/" + route

	log.Printf(">> [%s] %s -> %s", api, method, url)

	client := resty.New()

	// Encode PAT
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
