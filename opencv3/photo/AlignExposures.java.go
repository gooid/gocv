package photo

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

type AlignExposures struct {
	*Algorithm
}

func NewAlignExposures(addr int64) (rcvr *AlignExposures) {
	rcvr = &AlignExposures{}
	rcvr.Algorithm = NewAlgorithm(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func (rcvr *AlignExposures) finalize() {
	AlignExposuresNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *AlignExposures) Process(src []*Mat, dst []*Mat, times *Mat, response *Mat) {
	src_mat := ConvertersVectorMat(src)
	dst_mat := ConvertersVectorMat(dst)
	AlignExposuresNative_process_0(rcvr.GetNativeObjAddr(), src_mat.GetNativeObjAddr(), dst_mat.GetNativeObjAddr(), times.GetNativeObjAddr(), response.GetNativeObjAddr())
	return
}
