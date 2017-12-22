package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gooid/gocv/opencv3/core"
	"github.com/gooid/gocv/opencv3/imgcodecs"
	"github.com/gooid/gocv/opencv3/imgproc"
	"github.com/gooid/gocv/opencv3/objdetect"
)

// : facedetect ..\testdata\lena.jpg ..\testdata\haarcascade_frontalface_alt2.xml ..\testdata\haarcascade_eye_tree_eyeglasses.xml
func main() {
	flag.Parse()

	if flag.Arg(0) == "" || flag.Arg(1) == "" {
		fmt.Println(filepath.Base(os.Args[0]), "image_file", "face_cascade", "[eye_cascade]")
		return
	}

	org := imgcodecs.Imread(flag.Arg(0), imgcodecs.IMREAD_ANYCOLOR)
	defer org.Release()
	fmt.Println("org:", org)

	cascade := objdetect.NewCascadeClassifier2(flag.Arg(1))

	var eyecascade *objdetect.CascadeClassifier
	if flag.Arg(2) != "" {
		eyecascade = objdetect.NewCascadeClassifier2(flag.Arg(2))
	}

	gray := core.NewMat2()
	imgproc.CvtColor2(org, gray, imgproc.COLOR_RGB2GRAY)

	rv := core.NewMatOfRect()
	cascade.DetectMultiScale2(gray, rv)

	for _, r := range rv.ToArray() {
		fmt.Println("\t", r.X, r.Y, r.Width, r.Height)
		imgproc.Rectangle3(org, r.Tl(), r.Br(), core.NewScalar4(0.9))

		if eyecascade != nil {
			ev := core.NewMatOfRect()
			egray := gray.Submat3(r)
			eyecascade.DetectMultiScale2(egray, ev)
			drawEyeResult(org, r, ev, core.NewScalar2(128.0, 125.0, 255.0))
		}
	}
	if eyecascade != nil {
		r := core.NewRect(0, 0, gray.Width(), gray.Height())
		ev := core.NewMatOfRect()
		eyecascade.DetectMultiScale2(gray, ev)
		drawEyeResult(org, r, ev, core.NewScalar2(255.0, 255.0, 255.0))
	}

	imgcodecs.Imwrite2("f_"+filepath.Base(flag.Arg(0)), org)
}

func drawEyeResult(org *core.Mat, gr *core.Rect, rs *core.MatOfRect, clr *core.Scalar) {
	for _, r := range rs.ToArray() {
		fmt.Println("\te:", r.X, r.Y, r.Width, r.Height)
		imgproc.Ellipse3(org, core.NewPoint(float64(gr.X+r.X+r.Width/2), float64(gr.Y+r.Y+r.Height/2)),
			core.NewSize(float64(r.Width/2), float64(r.Height/2)),
			360.0, 0.0, 360.0, clr)
	}
}
