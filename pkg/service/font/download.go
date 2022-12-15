package font

import (
	"io"
	"net/http"
	"os"
)

// DownloadFrom will download the file at URL and save it to `dest`.
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
