package video

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

type BackgroundSubtractorKNN struct {
	*BackgroundSubtractor
}

func NewBackgroundSubtractorKNN(addr int64) (rcvr *BackgroundSubtractorKNN) {
	rcvr = &BackgroundSubtractorKNN{}
	rcvr.BackgroundSubtractor = NewBackgroundSubtractor(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func (rcvr *BackgroundSubtractorKNN) finalize() {
	BackgroundSubtractorKNNNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *BackgroundSubtractorKNN) GetDetectShadows() bool {
	retVal := BackgroundSubtractorKNNNative_getDetectShadows_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *BackgroundSubtractorKNN) GetDist2Threshold() float64 {
	retVal := BackgroundSubtractorKNNNative_getDist2Threshold_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *BackgroundSubtractorKNN) GetHistory() int {
	retVal := BackgroundSubtractorKNNNative_getHistory_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *BackgroundSubtractorKNN) GetNSamples() int {
	retVal := BackgroundSubtractorKNNNative_getNSamples_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *BackgroundSubtractorKNN) GetShadowThreshold() float64 {
	retVal := BackgroundSubtractorKNNNative_getShadowThreshold_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *BackgroundSubtractorKNN) GetShadowValue() int {
	retVal := BackgroundSubtractorKNNNative_getShadowValue_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *BackgroundSubtractorKNN) GetkNNSamples() int {
	retVal := BackgroundSubtractorKNNNative_getkNNSamples_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *BackgroundSubtractorKNN) SetDetectShadows(detectShadows bool) {
	BackgroundSubtractorKNNNative_setDetectShadows_0(rcvr.GetNativeObjAddr(), detectShadows)
	return
}
func (rcvr *BackgroundSubtractorKNN) SetDist2Threshold(_dist2Threshold float64) {
	BackgroundSubtractorKNNNative_setDist2Threshold_0(rcvr.GetNativeObjAddr(), _dist2Threshold)
	return
}
func (rcvr *BackgroundSubtractorKNN) SetHistory(history int) {
	BackgroundSubtractorKNNNative_setHistory_0(rcvr.GetNativeObjAddr(), history)
	return
}
func (rcvr *BackgroundSubtractorKNN) SetNSamples(_nN int) {
	BackgroundSubtractorKNNNative_setNSamples_0(rcvr.GetNativeObjAddr(), _nN)
	return
}
func (rcvr *BackgroundSubtractorKNN) SetShadowThreshold(threshold float64) {
	BackgroundSubtractorKNNNative_setShadowThreshold_0(rcvr.GetNativeObjAddr(), threshold)
	return
}
func (rcvr *BackgroundSubtractorKNN) SetShadowValue(value int) {
	BackgroundSubtractorKNNNative_setShadowValue_0(rcvr.GetNativeObjAddr(), value)
	return
}
func (rcvr *BackgroundSubtractorKNN) SetkNNSamples(_nkNN int) {
	BackgroundSubtractorKNNNative_setkNNSamples_0(rcvr.GetNativeObjAddr(), _nkNN)
	return
}
