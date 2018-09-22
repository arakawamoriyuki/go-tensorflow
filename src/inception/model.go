package inception

import (
	"os"
	"io"
	"net/http"
	"bufio"
	"sort"
	"fmt"
)

type LabelResult struct {
	Label string `json:"label"`
	Probability float32 `json:"probability"`
}

func (labelResult LabelResult) String() string {
	return fmt.Sprintf("%s) %f", labelResult.Label, labelResult.Probability)
}

type LabelResults []LabelResult

func (labelResults LabelResults) String() string {
	var text string
	for _, labelResult := range labelResults {
		text += fmt.Sprintln(labelResult.String())
	}
	return text
}

func (labelResults LabelResults) Len() int {
	return len(labelResults)
}

func (labelResults LabelResults) Swap(i, j int) {
	labelResults[i], labelResults[j] = labelResults[j], labelResults[i]
}

func (labelResults LabelResults) Less(i, j int) bool {
	return labelResults[i].Probability > labelResults[j].Probability
}

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

var labels []string

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

func loadLabelFile() error {
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

func findBestLabels(probabilities []float32) []LabelResult {
	var resultLabels []LabelResult
	for i, p := range probabilities {
		if i >= len(labels) {
			break
		}
		resultLabels = append(resultLabels, LabelResult{Label: labels[i], Probability: p})
	}
	sort.Sort(LabelResults(resultLabels))
	// Return top 5 labels
	return resultLabels[:5]
}
