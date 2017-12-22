package calib3d

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

const MODE_SGBM = 0
const MODE_HH = 1
const MODE_SGBM_3WAY = 2
const MODE_HH4 = 3

type StereoSGBM struct {
	*StereoMatcher
}

func NewStereoSGBM(addr int64) (rcvr *StereoSGBM) {
	rcvr = &StereoSGBM{}
	rcvr.StereoMatcher = NewStereoMatcher(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func StereoSGBMCreate(minDisparity int, numDisparities int, blockSize int, P1 int, P2 int, disp12MaxDiff int, preFilterCap int, uniquenessRatio int, speckleWindowSize int, speckleRange int, mode int) *StereoSGBM {
	retVal := NewStereoSGBM(StereoSGBMNative_create_0(minDisparity, numDisparities, blockSize, P1, P2, disp12MaxDiff, preFilterCap, uniquenessRatio, speckleWindowSize, speckleRange, mode))
	return retVal
}
func StereoSGBMCreate2() *StereoSGBM {
	retVal := NewStereoSGBM(StereoSGBMNative_create_1())
	return retVal
}
func (rcvr *StereoSGBM) finalize() {
	StereoSGBMNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *StereoSGBM) GetMode() int {
	retVal := StereoSGBMNative_getMode_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *StereoSGBM) GetP1() int {
	retVal := StereoSGBMNative_getP1_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *StereoSGBM) GetP2() int {
	retVal := StereoSGBMNative_getP2_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *StereoSGBM) GetPreFilterCap() int {
	retVal := StereoSGBMNative_getPreFilterCap_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *StereoSGBM) GetUniquenessRatio() int {
	retVal := StereoSGBMNative_getUniquenessRatio_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *StereoSGBM) SetMode(mode int) {
	StereoSGBMNative_setMode_0(rcvr.GetNativeObjAddr(), mode)
	return
}
func (rcvr *StereoSGBM) SetP1(P1 int) {
	StereoSGBMNative_setP1_0(rcvr.GetNativeObjAddr(), P1)
	return
}
func (rcvr *StereoSGBM) SetP2(P2 int) {
	StereoSGBMNative_setP2_0(rcvr.GetNativeObjAddr(), P2)
	return
}
func (rcvr *StereoSGBM) SetPreFilterCap(preFilterCap int) {
	StereoSGBMNative_setPreFilterCap_0(rcvr.GetNativeObjAddr(), preFilterCap)
	return
}
func (rcvr *StereoSGBM) SetUniquenessRatio(uniquenessRatio int) {
	StereoSGBMNative_setUniquenessRatio_0(rcvr.GetNativeObjAddr(), uniquenessRatio)
	return
}
