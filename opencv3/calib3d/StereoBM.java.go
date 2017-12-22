package calib3d

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

const BMPREFILTER_NORMALIZED_RESPONSE = 0
const BMPREFILTER_XSOBEL = 1

type StereoBM struct {
	*StereoMatcher
}

func NewStereoBM(addr int64) (rcvr *StereoBM) {
	rcvr = &StereoBM{}
	rcvr.StereoMatcher = NewStereoMatcher(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func StereoBMCreate(numDisparities int, blockSize int) *StereoBM {
	retVal := NewStereoBM(StereoBMNative_create_0(numDisparities, blockSize))
	return retVal
}
func StereoBMCreate2() *StereoBM {
	retVal := NewStereoBM(StereoBMNative_create_1())
	return retVal
}
func (rcvr *StereoBM) finalize() {
	StereoBMNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *StereoBM) GetPreFilterCap() int {
	retVal := StereoBMNative_getPreFilterCap_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *StereoBM) GetPreFilterSize() int {
	retVal := StereoBMNative_getPreFilterSize_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *StereoBM) GetPreFilterType() int {
	retVal := StereoBMNative_getPreFilterType_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *StereoBM) GetROI1() *Rect {
	retVal := NewRect5(StereoBMNative_getROI1_0(rcvr.GetNativeObjAddr()))
	return retVal
}
func (rcvr *StereoBM) GetROI2() *Rect {
	retVal := NewRect5(StereoBMNative_getROI2_0(rcvr.GetNativeObjAddr()))
	return retVal
}
func (rcvr *StereoBM) GetSmallerBlockSize() int {
	retVal := StereoBMNative_getSmallerBlockSize_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *StereoBM) GetTextureThreshold() int {
	retVal := StereoBMNative_getTextureThreshold_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *StereoBM) GetUniquenessRatio() int {
	retVal := StereoBMNative_getUniquenessRatio_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *StereoBM) SetPreFilterCap(preFilterCap int) {
	StereoBMNative_setPreFilterCap_0(rcvr.GetNativeObjAddr(), preFilterCap)
	return
}
func (rcvr *StereoBM) SetPreFilterSize(preFilterSize int) {
	StereoBMNative_setPreFilterSize_0(rcvr.GetNativeObjAddr(), preFilterSize)
	return
}
func (rcvr *StereoBM) SetPreFilterType(preFilterType int) {
	StereoBMNative_setPreFilterType_0(rcvr.GetNativeObjAddr(), preFilterType)
	return
}
func (rcvr *StereoBM) SetROI1(roi1 *Rect) {
	StereoBMNative_setROI1_0(rcvr.GetNativeObjAddr(), roi1.X, roi1.Y, roi1.Width, roi1.Height)
	return
}
func (rcvr *StereoBM) SetROI2(roi2 *Rect) {
	StereoBMNative_setROI2_0(rcvr.GetNativeObjAddr(), roi2.X, roi2.Y, roi2.Width, roi2.Height)
	return
}
func (rcvr *StereoBM) SetSmallerBlockSize(blockSize int) {
	StereoBMNative_setSmallerBlockSize_0(rcvr.GetNativeObjAddr(), blockSize)
	return
}
func (rcvr *StereoBM) SetTextureThreshold(textureThreshold int) {
	StereoBMNative_setTextureThreshold_0(rcvr.GetNativeObjAddr(), textureThreshold)
	return
}
func (rcvr *StereoBM) SetUniquenessRatio(uniquenessRatio int) {
	StereoBMNative_setUniquenessRatio_0(rcvr.GetNativeObjAddr(), uniquenessRatio)
	return
}
