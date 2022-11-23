package api

import (
	"encoding/json"
	"fmt"
	"fontman/client/pkg/model"
	"io/ioutil"
	"net/http"
	"os"
)

// TODO: load this from a remotes config variable, so: remotes: [ registry.fontman.io, http://196.1668... ]
var BASE_URL string = "http://127.0.0.1:8080"

func DownloadFrom(url string, dest string) error {
	// response, err := http.Get("https://github.com/google/fonts/raw/main/ofl/anonymouspro/AnonymousPro-Bold.ttf")
	response, err := http.Get(url)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	contents, readErr := ioutil.ReadAll(response.Body)

	if readErr != nil {
		return readErr
	}

	return os.WriteFile(dest, contents, 0777)
}

func GetFontDetails(id string) (*model.RemoteFontFamily, error) {
	url := fmt.Sprintf("%s/api/font/%s", BASE_URL, id)
	response, getErr := http.Get(url)

	if getErr != nil {
		return nil, getErr
	}

	defer response.Body.Close()

	// read response to byte array
	body, bodyErr := ioutil.ReadAll(response.Body)

	if bodyErr != nil {
		return nil, bodyErr
	}

	// convert downloaded JSON to internal struct
	var font model.RemoteFontFamily
	parseErr := json.Unmarshal(body, &font)

	if parseErr != nil {
		return nil, parseErr
	}

	// download each style to a file with name: <family>-<style>.<format>
	for _, style := range font.Styles {
		dest := fmt.Sprintf("%s-%s.%s", font.Name, style.Type, "ttf")

		if err := DownloadFrom(fmt.Sprintf("http://%s", style.Url), dest); err != nil {
			return nil, err
		}
	}

	return &font, nil
}
