package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gooid/gocv/opencv3/core"
	"github.com/gooid/gocv/opencv3/dnn"
	"github.com/gooid/gocv/opencv3/imgcodecs"
	"github.com/gooid/gocv/opencv3/imgproc"
)

// inception5h.zip
//   download url: https://storage.googleapis.com/download.tensorflow.org/models/inception5h.zip

// : inception ..\testdata\tensorflow_inception_graph.pb ..\testdata\imagenet_comp_graph_label_strings.txt ..\testdata\baboon.jpg
func main() {
	flag.Parse()
	if flag.Arg(0) == "" || flag.Arg(1) == "" || flag.Arg(2) == "" {
		fmt.Println(filepath.Base(os.Args[0]), ".\\tensorflow_inception_graph.pb",
			".\\imagenet_comp_graph_label_strings.txt",
			"image_file")
		return
	}

	modelFile := flag.Arg(0)
	classNamesFile := flag.Arg(1)
	imageFile := flag.Arg(2)

	image := imgcodecs.Imread2(imageFile)
	imgproc.Resize2(image, image, core.NewSize(224, 224))
	fmt.Println("image:", image)

	inputBlob := dnn.BlobFromImage2(image)
	//core.Add6(inputBlob, core.NewScalar4(117), inputBlob)
	//fmt.Println("inputBlob:", inputBlob)

	net := dnn.ReadNetFromTensorflow2(modelFile)
	net.SetInput2(inputBlob)
	//net.SetInput(inputBlob, "")
	//net.SetPreferableBackend(dnn.DNN_BACKEND_HALIDE)

	result := net.Forward("softmax2")
	//result := net.Forward2()
	//fmt.Println("forward:", result)

	classNames := readClassNames(classNamesFile)
	//fmt.Println("classNames:", classNames[0], "...", len(classNames)-1, ":", classNames[len(classNames)-1])

	classId, classProb := getMaxClass(result)
	fmt.Println("result:", classId, classProb, classNames[classId])
}

func getMaxClass(probBlob *core.Mat) (classId int, classProb float64) {
	probMat := probBlob.Reshape(1, 1)
	_, _, maxv, max := core.MinMaxLoc2(probMat)
	//fmt.Printf("%s = %#v\n", probMat, ret)
	return int(max.X + max.Y), maxv
}

func readClassNames(filename string) []string {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(bs), "\n")
}
