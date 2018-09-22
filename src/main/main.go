package main

import (
	"flag"
	"fmt"

	inception "github.com/arakawamoriyuki/go-tensorflow/src/inception"
)

func main() {
	flag.Parse()
	args := flag.Args()
	imagePath := args[0]

	err := inception.SetUp()
	if err != nil {
		panic(err)
	}

	labelResults, err := inception.Classify(imagePath)
	if err != nil {
		panic(err)
	}

	fmt.Println(labelResults)
}

