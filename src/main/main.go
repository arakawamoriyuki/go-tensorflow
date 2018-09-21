package main

import (
	util "github.com/arakawamoriyuki/go-tensorflow/src/util"
	inception "github.com/arakawamoriyuki/go-tensorflow/src/inception"
)

func main() {
	const path string = "inception5h.zip"

	modelExists := util.FileExists(path)
	if modelExists == false {
		err := inception.DownloadModelZip(path)
		if err != nil {
			panic(err)
		}
	}
}

