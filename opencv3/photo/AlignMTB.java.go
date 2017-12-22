package photo

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

type AlignMTB struct {
	*AlignExposures
}

func NewAlignMTB(addr int64) (rcvr *AlignMTB) {
	rcvr = &AlignMTB{}
	rcvr.AlignExposures = NewAlignExposures(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func (rcvr *AlignMTB) CalculateShift(img0 *Mat, img1 *Mat) *Point {
	retVal := NewPoint3(AlignMTBNative_calculateShift_0(rcvr.GetNativeObjAddr(), img0.GetNativeObjAddr(), img1.GetNativeObjAddr()))
	return retVal
}
func (rcvr *AlignMTB) ComputeBitmaps(img *Mat, tb *Mat, eb *Mat) {
	AlignMTBNative_computeBitmaps_0(rcvr.GetNativeObjAddr(), img.GetNativeObjAddr(), tb.GetNativeObjAddr(), eb.GetNativeObjAddr())
	return
}
func (rcvr *AlignMTB) finalize() {
	AlignMTBNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *AlignMTB) GetCut() bool {
	retVal := AlignMTBNative_getCut_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *AlignMTB) GetExcludeRange() int {
	retVal := AlignMTBNative_getExcludeRange_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *AlignMTB) GetMaxBits() int {
	retVal := AlignMTBNative_getMaxBits_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *AlignMTB) Process(src []*Mat, dst []*Mat, times *Mat, response *Mat) {
	src_mat := ConvertersVectorMat(src)
	dst_mat := ConvertersVectorMat(dst)
	AlignMTBNative_process_0(rcvr.GetNativeObjAddr(), src_mat.GetNativeObjAddr(), dst_mat.GetNativeObjAddr(), times.GetNativeObjAddr(), response.GetNativeObjAddr())
	return
}
func (rcvr *AlignMTB) Process2(src []*Mat, dst []*Mat) {
	src_mat := ConvertersVectorMat(src)
	dst_mat := ConvertersVectorMat(dst)
	AlignMTBNative_process_1(rcvr.GetNativeObjAddr(), src_mat.GetNativeObjAddr(), dst_mat.GetNativeObjAddr())
	return
}
func (rcvr *AlignMTB) SetCut(value bool) {
	AlignMTBNative_setCut_0(rcvr.GetNativeObjAddr(), value)
	return
}
func (rcvr *AlignMTB) SetExcludeRange(exclude_range int) {
	AlignMTBNative_setExcludeRange_0(rcvr.GetNativeObjAddr(), exclude_range)
	return
}
func (rcvr *AlignMTB) SetMaxBits(max_bits int) {
	AlignMTBNative_setMaxBits_0(rcvr.GetNativeObjAddr(), max_bits)
	return
}
func (rcvr *AlignMTB) ShiftMat(src *Mat, dst *Mat, shift *Point) {
	AlignMTBNative_shiftMat_0(rcvr.GetNativeObjAddr(), src.GetNativeObjAddr(), dst.GetNativeObjAddr(), shift.X, shift.Y)
	return
}
