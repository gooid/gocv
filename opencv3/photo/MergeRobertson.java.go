package photo

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

type MergeRobertson struct {
	*MergeExposures
}

func NewMergeRobertson(addr int64) (rcvr *MergeRobertson) {
	rcvr = &MergeRobertson{}
	rcvr.MergeExposures = NewMergeExposures(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func (rcvr *MergeRobertson) finalize() {
	MergeRobertsonNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *MergeRobertson) Process(src []*Mat, dst *Mat, times *Mat, response *Mat) {
	src_mat := ConvertersVectorMat(src)
	MergeRobertsonNative_process_0(rcvr.GetNativeObjAddr(), src_mat.GetNativeObjAddr(), dst.GetNativeObjAddr(), times.GetNativeObjAddr(), response.GetNativeObjAddr())
	return
}
func (rcvr *MergeRobertson) Process2(src []*Mat, dst *Mat, times *Mat) {
	src_mat := ConvertersVectorMat(src)
	MergeRobertsonNative_process_1(rcvr.GetNativeObjAddr(), src_mat.GetNativeObjAddr(), dst.GetNativeObjAddr(), times.GetNativeObjAddr())
	return
}
