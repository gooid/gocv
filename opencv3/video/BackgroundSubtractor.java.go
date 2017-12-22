package video

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

type BackgroundSubtractor struct {
	*Algorithm
}

func NewBackgroundSubtractor(addr int64) (rcvr *BackgroundSubtractor) {
	rcvr = &BackgroundSubtractor{}
	rcvr.Algorithm = NewAlgorithm(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func (rcvr *BackgroundSubtractor) Apply(image *Mat, fgmask *Mat, learningRate float64) {
	BackgroundSubtractorNative_apply_0(rcvr.GetNativeObjAddr(), image.GetNativeObjAddr(), fgmask.GetNativeObjAddr(), learningRate)
	return
}
func (rcvr *BackgroundSubtractor) Apply2(image *Mat, fgmask *Mat) {
	BackgroundSubtractorNative_apply_1(rcvr.GetNativeObjAddr(), image.GetNativeObjAddr(), fgmask.GetNativeObjAddr())
	return
}
func (rcvr *BackgroundSubtractor) finalize() {
	BackgroundSubtractorNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *BackgroundSubtractor) GetBackgroundImage(backgroundImage *Mat) {
	BackgroundSubtractorNative_getBackgroundImage_0(rcvr.GetNativeObjAddr(), backgroundImage.GetNativeObjAddr())
	return
}
