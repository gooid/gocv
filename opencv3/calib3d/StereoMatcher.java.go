package calib3d

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

const DISP_SHIFT = 4

const DISP_SCALE = 1 << DISP_SHIFT

type StereoMatcher struct {
	*Algorithm
}

func NewStereoMatcher(addr int64) (rcvr *StereoMatcher) {
	rcvr = &StereoMatcher{}
	rcvr.Algorithm = NewAlgorithm(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func (rcvr *StereoMatcher) Compute(left *Mat, right *Mat, disparity *Mat) {
	StereoMatcherNative_compute_0(rcvr.GetNativeObjAddr(), left.GetNativeObjAddr(), right.GetNativeObjAddr(), disparity.GetNativeObjAddr())
	return
}
func (rcvr *StereoMatcher) finalize() {
	StereoMatcherNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *StereoMatcher) GetBlockSize() int {
	retVal := StereoMatcherNative_getBlockSize_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *StereoMatcher) GetDisp12MaxDiff() int {
	retVal := StereoMatcherNative_getDisp12MaxDiff_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *StereoMatcher) GetMinDisparity() int {
	retVal := StereoMatcherNative_getMinDisparity_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *StereoMatcher) GetNumDisparities() int {
	retVal := StereoMatcherNative_getNumDisparities_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *StereoMatcher) GetSpeckleRange() int {
	retVal := StereoMatcherNative_getSpeckleRange_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *StereoMatcher) GetSpeckleWindowSize() int {
	retVal := StereoMatcherNative_getSpeckleWindowSize_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *StereoMatcher) SetBlockSize(blockSize int) {
	StereoMatcherNative_setBlockSize_0(rcvr.GetNativeObjAddr(), blockSize)
	return
}
func (rcvr *StereoMatcher) SetDisp12MaxDiff(disp12MaxDiff int) {
	StereoMatcherNative_setDisp12MaxDiff_0(rcvr.GetNativeObjAddr(), disp12MaxDiff)
	return
}
func (rcvr *StereoMatcher) SetMinDisparity(minDisparity int) {
	StereoMatcherNative_setMinDisparity_0(rcvr.GetNativeObjAddr(), minDisparity)
	return
}
func (rcvr *StereoMatcher) SetNumDisparities(numDisparities int) {
	StereoMatcherNative_setNumDisparities_0(rcvr.GetNativeObjAddr(), numDisparities)
	return
}
func (rcvr *StereoMatcher) SetSpeckleRange(speckleRange int) {
	StereoMatcherNative_setSpeckleRange_0(rcvr.GetNativeObjAddr(), speckleRange)
	return
}
func (rcvr *StereoMatcher) SetSpeckleWindowSize(speckleWindowSize int) {
	StereoMatcherNative_setSpeckleWindowSize_0(rcvr.GetNativeObjAddr(), speckleWindowSize)
	return
}
