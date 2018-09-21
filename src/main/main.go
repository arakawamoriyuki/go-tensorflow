package main

import (
	inception "github.com/arakawamoriyuki/go-tensorflow/src/inception"
)

func main() {
	err := inception.SetUp()
	if err != nil {
		panic(err)
	}
}

