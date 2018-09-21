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

func ModelZipExists() bool {
	return FileExists(modelFile.UnzipPath + modelFile.ZipPath)
}

func ModelExists() bool {
	return FileExists(modelFile.UnzipPath + modelFile.LabelPath) &&
		FileExists(modelFile.Path)
}

func DownloadModelZip() error {
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
	modelExists := ModelExists()
	if modelExists == false {
		modelZipExists := ModelZipExists()
		if modelZipExists == false {
			err := DownloadModelZip()
			if err != nil {
				return err
			}
		}
		err := Unzip(modelFile.ZipPath, modelFile.UnzipPath)
		if err != nil {
			return err
		}
	}

	return nil
}
