package features2d

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

const AgastFeatureDetectorAGAST_5_8 = 0
const AgastFeatureDetectorAGAST_7_12d = 1
const AgastFeatureDetectorAGAST_7_12s = 2
const AgastFeatureDetectorOAST_9_16 = 3
const AgastFeatureDetectorTHRESHOLD = 10000
const AgastFeatureDetectorNONMAX_SUPPRESSION = 10001

type AgastFeatureDetector struct {
	*Feature2D
}

func NewAgastFeatureDetector(addr int64) (rcvr *AgastFeatureDetector) {
	rcvr = &AgastFeatureDetector{}
	rcvr.Feature2D = NewFeature2D(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func AgastFeatureDetectorCreate(threshold int, nonmaxSuppression bool, rtype int) *AgastFeatureDetector {
	retVal := NewAgastFeatureDetector(AgastFeatureDetectorNative_create_0(threshold, nonmaxSuppression, rtype))
	return retVal
}
func AgastFeatureDetectorCreate2() *AgastFeatureDetector {
	retVal := NewAgastFeatureDetector(AgastFeatureDetectorNative_create_1())
	return retVal
}
func (rcvr *AgastFeatureDetector) finalize() {
	AgastFeatureDetectorNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *AgastFeatureDetector) GetDefaultName() string {
	retVal := AgastFeatureDetectorNative_getDefaultName_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *AgastFeatureDetector) GetNonmaxSuppression() bool {
	retVal := AgastFeatureDetectorNative_getNonmaxSuppression_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *AgastFeatureDetector) GetThreshold() int {
	retVal := AgastFeatureDetectorNative_getThreshold_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *AgastFeatureDetector) GetType() int {
	retVal := AgastFeatureDetectorNative_getType_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *AgastFeatureDetector) SetNonmaxSuppression(f bool) {
	AgastFeatureDetectorNative_setNonmaxSuppression_0(rcvr.GetNativeObjAddr(), f)
	return
}
func (rcvr *AgastFeatureDetector) SetThreshold(threshold int) {
	AgastFeatureDetectorNative_setThreshold_0(rcvr.GetNativeObjAddr(), threshold)
	return
}
func (rcvr *AgastFeatureDetector) SetType(rtype int) {
	AgastFeatureDetectorNative_setType_0(rcvr.GetNativeObjAddr(), rtype)
	return
}
