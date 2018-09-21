package inception

import (
	"os"
	"io"
	"net/http"
)

type ModelFile struct {
	DownloadUrl string
	ZipPath string
	UnzipPath string
	LabelPath string
	Path string
}

var modelFile ModelFile = ModelFile {
	DownloadUrl: "https://storage.googleapis.com/download.tensorflow.org/models/inception5h.zip",
	ZipPath: "inception5h.zip",
	UnzipPath: "./",
	LabelPath: "imagenet_comp_graph_label_strings.txt",
	Path: "tensorflow_inception_graph.pb",
}

func modelZipExists() bool {
	return fileExists(modelFile.UnzipPath + modelFile.ZipPath)
}

func modelExists() bool {
	return fileExists(modelFile.UnzipPath + modelFile.LabelPath) &&
		fileExists(modelFile.Path)
}

func downloadModelZip() error {
	response, err := http.Get(modelFile.DownloadUrl)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(modelFile.UnzipPath)
	if err != nil {
		return err
	}
	defer file.Close()

	io.Copy(file, response.Body)
	return nil
}

func SetUp() error {
	modelExists := modelExists()
	if modelExists == false {
		modelZipExists := modelZipExists()
		if modelZipExists == false {
			err := downloadModelZip()
			if err != nil {
				return err
			}
		}
		err := unzip(modelFile.ZipPath, modelFile.UnzipPath)
		if err != nil {
			return err
		}
	}

	return nil
}
