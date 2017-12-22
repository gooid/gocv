package video

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

type DenseOpticalFlow struct {
	*Algorithm
}

func NewDenseOpticalFlow(addr int64) (rcvr *DenseOpticalFlow) {
	rcvr = &DenseOpticalFlow{}
	rcvr.Algorithm = NewAlgorithm(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func (rcvr *DenseOpticalFlow) Calc(I0 *Mat, I1 *Mat, flow *Mat) {
	DenseOpticalFlowNative_calc_0(rcvr.GetNativeObjAddr(), I0.GetNativeObjAddr(), I1.GetNativeObjAddr(), flow.GetNativeObjAddr())
	return
}
func (rcvr *DenseOpticalFlow) CollectGarbage() {
	DenseOpticalFlowNative_collectGarbage_0(rcvr.GetNativeObjAddr())
	return
}
func (rcvr *DenseOpticalFlow) finalize() {
	DenseOpticalFlowNative_delete(rcvr.GetNativeObjAddr())
}
