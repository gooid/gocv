package photo

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

type MergeMertens struct {
	*MergeExposures
}

func NewMergeMertens(addr int64) (rcvr *MergeMertens) {
	rcvr = &MergeMertens{}
	rcvr.MergeExposures = NewMergeExposures(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func (rcvr *MergeMertens) finalize() {
	MergeMertensNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *MergeMertens) GetContrastWeight() float32 {
	retVal := MergeMertensNative_getContrastWeight_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *MergeMertens) GetExposureWeight() float32 {
	retVal := MergeMertensNative_getExposureWeight_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *MergeMertens) GetSaturationWeight() float32 {
	retVal := MergeMertensNative_getSaturationWeight_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *MergeMertens) Process(src []*Mat, dst *Mat, times *Mat, response *Mat) {
	src_mat := ConvertersVectorMat(src)
	MergeMertensNative_process_0(rcvr.GetNativeObjAddr(), src_mat.GetNativeObjAddr(), dst.GetNativeObjAddr(), times.GetNativeObjAddr(), response.GetNativeObjAddr())
	return
}
func (rcvr *MergeMertens) Process2(src []*Mat, dst *Mat) {
	src_mat := ConvertersVectorMat(src)
	MergeMertensNative_process_1(rcvr.GetNativeObjAddr(), src_mat.GetNativeObjAddr(), dst.GetNativeObjAddr())
	return
}
func (rcvr *MergeMertens) SetContrastWeight(contrast_weiht float32) {
	MergeMertensNative_setContrastWeight_0(rcvr.GetNativeObjAddr(), contrast_weiht)
	return
}
func (rcvr *MergeMertens) SetExposureWeight(exposure_weight float32) {
	MergeMertensNative_setExposureWeight_0(rcvr.GetNativeObjAddr(), exposure_weight)
	return
}
func (rcvr *MergeMertens) SetSaturationWeight(saturation_weight float32) {
	MergeMertensNative_setSaturationWeight_0(rcvr.GetNativeObjAddr(), saturation_weight)
	return
}
