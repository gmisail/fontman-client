package api

import (
	"encoding/json"
	"fmt"
	"fontman/client/pkg/model"
	"io/ioutil"
	"net/http"
)

// TODO: load this from a remotes config variable, so: remotes: [ registry.fontman.io, http://196.1668... ]
var BASE_URL string = "http://127.0.0.1:8080"

func GetFontDetails(id string) (*model.RemoteFontFamily, error) {
	url := fmt.Sprintf("%s/api/font/%s", BASE_URL, id)
	response, getErr := http.Get(url)

	if getErr != nil {
		return nil, getErr
	}

	defer response.Body.Close()

	body, bodyErr := ioutil.ReadAll(response.Body)

	if bodyErr != nil {
		return nil, bodyErr
	}

	var font model.RemoteFontFamily
	parseErr := json.Unmarshal(body, &font)

	if parseErr != nil {
		return nil, parseErr
	}

	return &font, nil
}
