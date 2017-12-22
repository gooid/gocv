package features2d

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

const FastFeatureDetectorTYPE_5_8 = 0
const FastFeatureDetectorTYPE_7_12 = 1
const FastFeatureDetectorTYPE_9_16 = 2
const FastFeatureDetectorTHRESHOLD = 10000
const FastFeatureDetectorNONMAX_SUPPRESSION = 10001
const FastFeatureDetectorFAST_N = 10002

type FastFeatureDetector struct {
	*Feature2D
}

func NewFastFeatureDetector(addr int64) (rcvr *FastFeatureDetector) {
	rcvr = &FastFeatureDetector{}
	rcvr.Feature2D = NewFeature2D(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func FastFeatureDetectorCreate(threshold int, nonmaxSuppression bool, rtype int) *FastFeatureDetector {
	retVal := NewFastFeatureDetector(FastFeatureDetectorNative_create_0(threshold, nonmaxSuppression, rtype))
	return retVal
}
func FastFeatureDetectorCreate2() *FastFeatureDetector {
	retVal := NewFastFeatureDetector(FastFeatureDetectorNative_create_1())
	return retVal
}
func (rcvr *FastFeatureDetector) finalize() {
	FastFeatureDetectorNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *FastFeatureDetector) GetDefaultName() string {
	retVal := FastFeatureDetectorNative_getDefaultName_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *FastFeatureDetector) GetNonmaxSuppression() bool {
	retVal := FastFeatureDetectorNative_getNonmaxSuppression_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *FastFeatureDetector) GetThreshold() int {
	retVal := FastFeatureDetectorNative_getThreshold_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *FastFeatureDetector) GetType() int {
	retVal := FastFeatureDetectorNative_getType_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *FastFeatureDetector) SetNonmaxSuppression(f bool) {
	FastFeatureDetectorNative_setNonmaxSuppression_0(rcvr.GetNativeObjAddr(), f)
	return
}
func (rcvr *FastFeatureDetector) SetThreshold(threshold int) {
	FastFeatureDetectorNative_setThreshold_0(rcvr.GetNativeObjAddr(), threshold)
	return
}
func (rcvr *FastFeatureDetector) SetType(rtype int) {
	FastFeatureDetectorNative_setType_0(rcvr.GetNativeObjAddr(), rtype)
	return
}
