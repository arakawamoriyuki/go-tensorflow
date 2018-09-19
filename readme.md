
# go-tensorflow

[Install TensorFlow for Go](https://www.tensorflow.org/install/install_go)
[TensorFlow in Go](https://github.com/tensorflow/tensorflow/tree/master/tensorflow/go)

## Warning

> Warning: The TensorFlow Go API is not covered by the TensorFlow API stability guarantees.

TensorFlow Go APIは動作保証してない！

## 環境

全て最新では動かなかったorz

```
# 環境
$ sw_vers
ProductName: Mac OS X
ProductVersion:	10.12.6
BuildVersion: 16G29
$ go version
go version go1.11 darwin/amd64
$ python --version
Python 3.5.2 :: Anaconda 4.2.0 (x86_64)
$ conda list | grep tensorflow
tensorflow 1.10.0 py35_0 conda-forge

# Goで利用する為にTensorFlow C ライブラリをインストール
$ TF_TYPE='cpu'
$ TARGET_DIRECTORY='/usr/local'
$ curl -L "https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-${TF_TYPE}-$(go env GOOS)-x86_64-1.10.1.tar.gz" | sudo tar -C $TARGET_DIRECTORY -xz
$ ls $TARGET_DIRECTORY/lib | grep libtensorflow.so

# tensorflow/goパッケージダウンロード
$ go get github.com/tensorflow/tensorflow/tensorflow/go
$ go test github.com/tensorflow/tensorflow/tensorflow/go
...エラーもなにも吐かずgo testずっと実行中
```

結構古いけど辛くも動作した `python tensorflow v1.4.0` `go tensorflow v1.4.0`

```
# 環境
$ sw_vers
ProductName: Mac OS X
ProductVersion:	10.12.6
BuildVersion: 16G29
$ goenv install 1.11
$ goenv global 1.11
$ go version
go version go1.11 darwin/amd64
$ pyenv install anaconda3-4.2.0
$ pyenv global anaconda3-4.2.0
$ python --version
Python 3.5.2 :: Anaconda 4.2.0 (x86_64)
$ conda install -y -c conda-forge tensorflow==1.4.0
$ conda list | grep tensorflow
tensorflow 1.4.0 py35_0 conda-forge
$ python
>>> import tensorflow as tf
>>> tf.__version__
'1.4.0'

# Goで利用する為にTensorFlow C ライブラリをインストール
$ TF_TYPE='cpu'
$ TARGET_DIRECTORY='/usr/local'
$ curl -L "https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-${TF_TYPE}-$(go env GOOS)-x86_64-1.10.1.tar.gz" | sudo tar -C $TARGET_DIRECTORY -xz

# Modulesを利用してv1.4.0をgo get
$ export GO111MODULE=on
$ go mod init
$ go get github.com/tensorflow/tensorflow/tensorflow/go@v1.4.0
$ go test github.com/tensorflow/tensorflow/tensorflow/go
ok github.com/tensorflow/tensorflow/tensorflow/go 0.345s

# 動作確認
$ go run src/main/main.go
2018-09-19 21:31:21.055739: I tensorflow/core/platform/cpu_feature_guard.cc:141] Your CPU supports instructions that this TensorFlow binary was not compiled to use: SSE4.2 AVX AVX2 FMA
Hello from TensorFlow version 1.10.1
```


## Dockerでやるなら

[github tinrab/go-tensorflow-image-recognition](https://github.com/tinrab/go-tensorflow-image-recognition)がお手軽

参考: [Build an Image Recognition API with Go and TensorFlow](https://outcrawl.com/image-recognition-api-go-tensorflow/)

- linux
- tensorflow v1.3.0
- go v1.10.3

```
$ git clone https://github.com/tinrab/go-tensorflow-image-recognition
$ cd go-tensorflow-image-recognition/api
$ docker build -t tensorflow-go .
```