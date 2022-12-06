package api

import (
	"encoding/json"
	"fmt"
	"fontman/client/pkg/model"
	"io"
	"net/http"
)

// GetFontDetails: return details for a font with ID
func GetFontDetails(id string, baseUrl string) (*model.RemoteFontFamily, error) {
	url := fmt.Sprintf("%s/api/font/%s", baseUrl, id)
	response, getErr := http.Get(url)

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

func GetFontOptions(name string, baseUrl string) ([]model.RemoteFontFamily, error) {
	url := fmt.Sprintf("%s/api/font?name=%s", baseUrl, name)
	response, getErr := http.Get(url)

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
