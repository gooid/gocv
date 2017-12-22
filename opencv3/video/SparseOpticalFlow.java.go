package video

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

type SparseOpticalFlow struct {
	*Algorithm
}

func NewSparseOpticalFlow(addr int64) (rcvr *SparseOpticalFlow) {
	rcvr = &SparseOpticalFlow{}
	rcvr.Algorithm = NewAlgorithm(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func (rcvr *SparseOpticalFlow) Calc(prevImg *Mat, nextImg *Mat, prevPts *Mat, nextPts *Mat, status *Mat, err *Mat) {
	SparseOpticalFlowNative_calc_0(rcvr.GetNativeObjAddr(), prevImg.GetNativeObjAddr(), nextImg.GetNativeObjAddr(), prevPts.GetNativeObjAddr(), nextPts.GetNativeObjAddr(), status.GetNativeObjAddr(), err.GetNativeObjAddr())
	return
}
func (rcvr *SparseOpticalFlow) Calc2(prevImg *Mat, nextImg *Mat, prevPts *Mat, nextPts *Mat, status *Mat) {
	SparseOpticalFlowNative_calc_1(rcvr.GetNativeObjAddr(), prevImg.GetNativeObjAddr(), nextImg.GetNativeObjAddr(), prevPts.GetNativeObjAddr(), nextPts.GetNativeObjAddr(), status.GetNativeObjAddr())
	return
}
func (rcvr *SparseOpticalFlow) finalize() {
	SparseOpticalFlowNative_delete(rcvr.GetNativeObjAddr())
}
