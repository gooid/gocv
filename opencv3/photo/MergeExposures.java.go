package photo

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

type MergeExposures struct {
	*Algorithm
}

func NewMergeExposures(addr int64) (rcvr *MergeExposures) {
	rcvr = &MergeExposures{}
	rcvr.Algorithm = NewAlgorithm(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func (rcvr *MergeExposures) finalize() {
	MergeExposuresNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *MergeExposures) Process(src []*Mat, dst *Mat, times *Mat, response *Mat) {
	src_mat := ConvertersVectorMat(src)
	MergeExposuresNative_process_0(rcvr.GetNativeObjAddr(), src_mat.GetNativeObjAddr(), dst.GetNativeObjAddr(), times.GetNativeObjAddr(), response.GetNativeObjAddr())
	return
}
