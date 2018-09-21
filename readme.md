
# go-tensorflow

[Install TensorFlow for Go](https://www.tensorflow.org/install/install_go)
[TensorFlow in Go](https://github.com/tensorflow/tensorflow/tree/master/tensorflow/go)

## Warning

> Warning: The TensorFlow Go API is not covered by the TensorFlow API stability guarantees.

TensorFlow Go APIは動作保証してない！

## ENV

```
$ sw_vers
ProductName: Mac OS X
ProductVersion:	10.12.6
BuildVersion: 16G29

$ go version
go version go1.11 darwin/amd64
```

## Install TensorFlow for Go

```
# Goで利用する為にTensorFlow C ライブラリをインストール
$ TF_TYPE='cpu'
$ TARGET_DIRECTORY='/usr/local'
$ curl -L "https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-${TF_TYPE}-$(go env GOOS)-x86_64-1.10.1.tar.gz" | sudo tar -C $TARGET_DIRECTORY -xz

# Modulesを利用してv1.10.1をgo get
$ export GO111MODULE=on
$ go mod init
$ go get github.com/tensorflow/tensorflow/tensorflow/go@v1.10.1
$ go test github.com/tensorflow/tensorflow/tensorflow/go
ok github.com/tensorflow/tensorflow/tensorflow/go 0.345s

# 動作確認
$ go run src/hello_tensorflow/main.go
2018-09-19 21:31:21.055739: I tensorflow/core/platform/cpu_feature_guard.cc:141] Your CPU supports instructions that this TensorFlow binary was not compiled to use: SSE4.2 AVX AVX2 FMA
Hello from TensorFlow version 1.10.1
```


## Tensorflow Tutorials Image Recognition

学習済みモデルinception-v3を使って画像分類をGoで

[Tensorflow Tutorials Image Recognition](https://www.tensorflow.org/tutorials/images/image_recognition)

TODO:


### バイナリにしてみる

TODO:


### 画像仕分けしてみる

TODO:


## Dockerでやるなら

[github tinrab/go-tensorflow-image-recognition](https://github.com/tinrab/go-tensorflow-image-recognition)がお手軽

参考: [Build an Image Recognition API with Go and TensorFlow](https://outcrawl.com/image-recognition-api-go-tensorflow/)

- linux
- tensorflow v1.3.0
- go v1.10.3

```
$ git clone https://github.com/tinrab/go-tensorflow-image-recognition

# 環境がほしいなら
$ cd go-tensorflow-image-recognition/api
$ docker build -t tensorflow-go .
$ docker run -it --rm tensorflow-go /bin/bash

# Image Recognitionを試したいなら
$ cd go-tensorflow-image-recognition
$ docker-compose -f docker-compose.yaml up -d --build
$ curl localhost:8080/recognize -F 'image=@./cat.jpg'
```
