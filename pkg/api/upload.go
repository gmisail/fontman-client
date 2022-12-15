package api

import (
	"bytes"
	"encoding/json"
	"fontman/client/pkg/model"
	"net/http"
	"net/url"
)

// UploadRegistryFile uploads a registry file to the remote registry.
func UploadRegistryFile(file model.RegistryFile, baseUrl string) error {
	contents, err := json.Marshal(file)

	if err != nil {
		return err
	}

	registryUrl, urlErr := url.Parse(baseUrl)
	registryUrl.Path += "api/font"

	if urlErr != nil {
		return urlErr
	}

	// don't need the response, we only need to check for an error!
	_, httpErr := http.Post(registryUrl.String(), "application/json", bytes.NewBuffer(contents))

	if httpErr != nil {
		return httpErr
	}

	return nil
}
