package font

import (
	"io"
	"net/http"
	"os"
)

// DownloadFrom: downloads file from 'url' and saves it as 'dest`
func DownloadFrom(url string, dest string) error {
	response, resErr := http.Get(url)

	if resErr != nil {
		return resErr
	}

	defer response.Body.Close()

	contents, readErr := io.ReadAll(response.Body)

	if readErr != nil {
		return readErr
	}

	return os.WriteFile(dest, contents, 0777)
}
