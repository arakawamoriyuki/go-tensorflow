package inception

import (
	"io/ioutil"

	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

var (
	graph *tf.Graph
	session *tf.Session
)

func startSession() error {
	model, err := ioutil.ReadFile(modelFile.GetModelFilePath())
	if err != nil {
		return err
	}

	graph = tf.NewGraph()
	if err := graph.Import(model, ""); err != nil {
		return err
	}

	session, err = tf.NewSession(graph, nil)
	if err != nil {
		return err
	}

	return nil
}