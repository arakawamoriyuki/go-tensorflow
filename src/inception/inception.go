package inception

import (
	"os"
	"io"
	"net/http"
)

func DownloadModelZip(path string) error {
	const url string = "https://storage.googleapis.com/download.tensorflow.org/models/inception5h.zip"

	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	io.Copy(file, response.Body)
	return nil
}
