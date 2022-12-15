package api

import (
	"encoding/json"
	"fontman/client/pkg/model"
	"io"
	"net/http"
	"net/url"
)

// GetFontDetails returns details for a font with ID.
func GetFontDetails(id string, baseUrl string) (*model.RemoteFontFamily, error) {
	// build the URL, make sure all components are properly escaped
	registryUrl, err := url.Parse(baseUrl)

	if err != nil {
		return nil, err
	}

	registryUrl.Path += "api/font/" + id

	response, getErr := http.Get(registryUrl.String())

	if getErr != nil {
		return nil, getErr
	}

	defer response.Body.Close()

	// read response to byte array
	body, bodyErr := io.ReadAll(response.Body)

	if bodyErr != nil {
		return nil, bodyErr
	}

	// convert downloaded JSON to internal struct
	var font model.RemoteFontFamily
	parseErr := json.Unmarshal(body, &font)

	if parseErr != nil {
		return nil, parseErr
	}

	return &font, nil
}

// GetFontOptions will get possible font matches based on its name.
func GetFontOptions(name string, baseUrl string) ([]model.RemoteFontFamily, error) {
	// build the URL, make sure all components are properly escaped
	registryUrl, err := url.Parse(baseUrl)

	if err != nil {
		return nil, err
	}

	registryUrl.Path += "api/font"

	params := url.Values{}
	params.Add("name", name)

	registryUrl.RawQuery = params.Encode()

	response, getErr := http.Get(registryUrl.String())

	if getErr != nil {
		return nil, getErr
	}

	defer response.Body.Close()

	// read response to byte array
	body, bodyErr := io.ReadAll(response.Body)

	if bodyErr != nil {
		return nil, bodyErr
	}

	var fonts model.RemoteFontList
	parseErr := json.Unmarshal(body, &fonts)

	if parseErr != nil {
		return nil, parseErr
	}

	return fonts.Fonts, nil
}
