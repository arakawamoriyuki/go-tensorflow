package inception

import (
	"os"
	"io"
	"bytes"
	"strings"

	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"github.com/tensorflow/tensorflow/tensorflow/go/op"
)

func SetUp() error {
	modelExists := modelExists()
	if modelExists == false {
		modelZipExists := modelZipExists()
		if modelZipExists == false {
			if err := downloadModelZip(); err != nil {
				return err
			}
		}
		if err := unzip(modelFile.ZipFilePath, modelFile.UnzipDestPath); err != nil {
			return err
		}
	}

	if err := loadLabelFile(); err != nil {
		return err
	}

	if err := startSession(); err != nil {
		return err
	}

	return nil
}

func Classify(imagePath string) (LabelResults, error) {
	imageFile, err := os.Open(imagePath)
	if err != nil {
		return nil, err
	}
	defer imageFile.Close()

	var imageBuffer bytes.Buffer
	io.Copy(&imageBuffer, imageFile)

	extPosition := strings.LastIndex(imagePath, ".") + 1
	imageFormat := imagePath[extPosition:]

	tensor, err := makeTensorFromImage(&imageBuffer, imageFormat)
	if err != nil {
		return nil, err
	}

	// Run inference
	output, err := session.Run(
		map[tf.Output]*tf.Tensor{
			graph.Operation("input").Output(0): tensor,
		},
		[]tf.Output{
			graph.Operation("output").Output(0),
		},
		nil)
	if err != nil {
		return nil, err
	}

	labelResults := findBestLabels(output[0].Value().([][]float32)[0])

	return labelResults, nil
}

func makeTensorFromImage(imageBuffer *bytes.Buffer, imageFormat string) (*tf.Tensor, error) {
	tensor, err := tf.NewTensor(imageBuffer.String())
	if err != nil {
		return nil, err
	}

	graph, input, output, err := makeTransformImageGraph(imageFormat)
	if err != nil {
		return nil, err
	}

	session, err := tf.NewSession(graph, nil)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	normalized, err := session.Run(
		map[tf.Output]*tf.Tensor{input: tensor},
		[]tf.Output{output},
		nil)
	if err != nil {
		return nil, err
	}

	return normalized[0], nil
}

func makeTransformImageGraph(imageFormat string) (graph *tf.Graph, input, output tf.Output, err error) {
	const (
		H, W  = 224, 224
		Mean  = float32(117)
		Scale = float32(1)
	)
	s := op.NewScope()
	input = op.Placeholder(s, tf.String)
	var decode tf.Output
	if imageFormat == "png" {
		decode = op.DecodePng(s, input, op.DecodePngChannels(3))
	} else {
		decode = op.DecodeJpeg(s, input, op.DecodeJpegChannels(3))
	}

	output = op.Div(s,
		op.Sub(s,
			op.ResizeBilinear(s,
				op.ExpandDims(s,
					op.Cast(s, decode, tf.Float),
					op.Const(s.SubScope("make_batch"), int32(0))),
				op.Const(s.SubScope("size"), []int32{H, W})),
			op.Const(s.SubScope("mean"), Mean)),
		op.Const(s.SubScope("scale"), Scale))
	graph, err = s.Finalize()

	return graph, input, output, err
}
