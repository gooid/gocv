package main

import (
	"fmt"

	"github.com/gooid/gocv/opencv3/calib3d"
	"github.com/gooid/gocv/opencv3/core"
	"github.com/gooid/gocv/opencv3/dnn"
	"github.com/gooid/gocv/opencv3/features2d"
	"github.com/gooid/gocv/opencv3/imgcodecs"
	"github.com/gooid/gocv/opencv3/imgproc"
	"github.com/gooid/gocv/opencv3/ml"
	"github.com/gooid/gocv/opencv3/objdetect"
	"github.com/gooid/gocv/opencv3/photo"
	"github.com/gooid/gocv/opencv3/video"
	"github.com/gooid/gocv/opencv3/videoio"
)

func main() {
	err := core.LoadCv()
	if err != nil {
		fmt.Println(err)
		return
	}

	_ = calib3d.CALIB_USE_INTRINSIC_GUESS
	_ = core.CoreCMP_LE
	_ = dnn.DNN_BACKEND_DEFAULT
	_ = features2d.AgastFeatureDetectorAGAST_5_8
	_ = imgcodecs.CV_CVTIMG_FLIP
	_ = imgproc.CCL_GRANA
	_ = ml.EMCOV_MAT_SPHERICAL
	_ = objdetect.HOGDescriptorDEFAULT_NLEVELS
	_ = photo.INPAINT_TELEA
	_ = video.OPTFLOW_USE_INITIAL_FLOW
	_ = videoio.CV_CAP_ANY

	fmt.Println(core.GetBuildInformation())
}
