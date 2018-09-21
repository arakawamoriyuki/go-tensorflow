package inception

import (
	"os"
	"io"
	"io/ioutil"
	"net/http"
	"bufio"

	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

type ModelFile struct {
	DownloadUrl string
	ZipFilePath string
	UnzipDestPath string
	LabelFileName string
	ModelFileName string
}

func (modelFile ModelFile) GetLabelFilePath() string {
    return modelFile.UnzipDestPath + modelFile.LabelFileName
}

func (modelFile ModelFile) GetModelFilePath() string {
    return modelFile.UnzipDestPath + modelFile.ModelFileName
}

var modelFile ModelFile = ModelFile {
	DownloadUrl: "https://storage.googleapis.com/download.tensorflow.org/models/inception5h.zip",
	ZipFilePath: "inception5h.zip",
	UnzipDestPath: "./",
	LabelFileName: "imagenet_comp_graph_label_strings.txt",
	ModelFileName: "tensorflow_inception_graph.pb",
}

var (
	graphModel   *tf.Graph
	sessionModel *tf.Session
	labels       []string
)

func modelZipExists() bool {
	return fileExists(modelFile.ZipFilePath)
}

func modelExists() bool {
	return fileExists(modelFile.GetLabelFilePath()) &&
		fileExists(modelFile.GetModelFilePath())
}

func downloadModelZip() error {
	response, err := http.Get(modelFile.DownloadUrl)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(modelFile.UnzipDestPath)
	if err != nil {
		return err
	}
	defer file.Close()

	io.Copy(file, response.Body)
	return nil
}

func loadModel() error {
	model, err := ioutil.ReadFile(modelFile.GetModelFilePath())
	if err != nil {
		return err
	}

	graphModel = tf.NewGraph()
	if err := graphModel.Import(model, ""); err != nil {
		return err
	}

	sessionModel, err = tf.NewSession(graphModel, nil)
	if err != nil {
		return err
	}

	labelsFile, err := os.Open(modelFile.GetLabelFilePath())
	if err != nil {
		return err
	}
	defer labelsFile.Close()

	scanner := bufio.NewScanner(labelsFile)
	for scanner.Scan() {
		labels = append(labels, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
