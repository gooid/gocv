package photo

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

type MergeDebevec struct {
	*MergeExposures
}

func NewMergeDebevec(addr int64) (rcvr *MergeDebevec) {
	rcvr = &MergeDebevec{}
	rcvr.MergeExposures = NewMergeExposures(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func (rcvr *MergeDebevec) finalize() {
	MergeDebevecNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *MergeDebevec) Process(src []*Mat, dst *Mat, times *Mat, response *Mat) {
	src_mat := ConvertersVectorMat(src)
	MergeDebevecNative_process_0(rcvr.GetNativeObjAddr(), src_mat.GetNativeObjAddr(), dst.GetNativeObjAddr(), times.GetNativeObjAddr(), response.GetNativeObjAddr())
	return
}
func (rcvr *MergeDebevec) Process2(src []*Mat, dst *Mat, times *Mat) {
	src_mat := ConvertersVectorMat(src)
	MergeDebevecNative_process_1(rcvr.GetNativeObjAddr(), src_mat.GetNativeObjAddr(), dst.GetNativeObjAddr(), times.GetNativeObjAddr())
	return
}
