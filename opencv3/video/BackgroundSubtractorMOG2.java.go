package video

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

type BackgroundSubtractorMOG2 struct {
	*BackgroundSubtractor
}

func NewBackgroundSubtractorMOG2(addr int64) (rcvr *BackgroundSubtractorMOG2) {
	rcvr = &BackgroundSubtractorMOG2{}
	rcvr.BackgroundSubtractor = NewBackgroundSubtractor(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func (rcvr *BackgroundSubtractorMOG2) Apply(image *Mat, fgmask *Mat, learningRate float64) {
	BackgroundSubtractorMOG2Native_apply_0(rcvr.GetNativeObjAddr(), image.GetNativeObjAddr(), fgmask.GetNativeObjAddr(), learningRate)
	return
}
func (rcvr *BackgroundSubtractorMOG2) Apply2(image *Mat, fgmask *Mat) {
	BackgroundSubtractorMOG2Native_apply_1(rcvr.GetNativeObjAddr(), image.GetNativeObjAddr(), fgmask.GetNativeObjAddr())
	return
}
func (rcvr *BackgroundSubtractorMOG2) finalize() {
	BackgroundSubtractorMOG2Native_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *BackgroundSubtractorMOG2) GetBackgroundRatio() float64 {
	retVal := BackgroundSubtractorMOG2Native_getBackgroundRatio_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *BackgroundSubtractorMOG2) GetComplexityReductionThreshold() float64 {
	retVal := BackgroundSubtractorMOG2Native_getComplexityReductionThreshold_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *BackgroundSubtractorMOG2) GetDetectShadows() bool {
	retVal := BackgroundSubtractorMOG2Native_getDetectShadows_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *BackgroundSubtractorMOG2) GetHistory() int {
	retVal := BackgroundSubtractorMOG2Native_getHistory_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *BackgroundSubtractorMOG2) GetNMixtures() int {
	retVal := BackgroundSubtractorMOG2Native_getNMixtures_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *BackgroundSubtractorMOG2) GetShadowThreshold() float64 {
	retVal := BackgroundSubtractorMOG2Native_getShadowThreshold_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *BackgroundSubtractorMOG2) GetShadowValue() int {
	retVal := BackgroundSubtractorMOG2Native_getShadowValue_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *BackgroundSubtractorMOG2) GetVarInit() float64 {
	retVal := BackgroundSubtractorMOG2Native_getVarInit_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *BackgroundSubtractorMOG2) GetVarMax() float64 {
	retVal := BackgroundSubtractorMOG2Native_getVarMax_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *BackgroundSubtractorMOG2) GetVarMin() float64 {
	retVal := BackgroundSubtractorMOG2Native_getVarMin_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *BackgroundSubtractorMOG2) GetVarThreshold() float64 {
	retVal := BackgroundSubtractorMOG2Native_getVarThreshold_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *BackgroundSubtractorMOG2) GetVarThresholdGen() float64 {
	retVal := BackgroundSubtractorMOG2Native_getVarThresholdGen_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *BackgroundSubtractorMOG2) SetBackgroundRatio(ratio float64) {
	BackgroundSubtractorMOG2Native_setBackgroundRatio_0(rcvr.GetNativeObjAddr(), ratio)
	return
}
func (rcvr *BackgroundSubtractorMOG2) SetComplexityReductionThreshold(ct float64) {
	BackgroundSubtractorMOG2Native_setComplexityReductionThreshold_0(rcvr.GetNativeObjAddr(), ct)
	return
}
func (rcvr *BackgroundSubtractorMOG2) SetDetectShadows(detectShadows bool) {
	BackgroundSubtractorMOG2Native_setDetectShadows_0(rcvr.GetNativeObjAddr(), detectShadows)
	return
}
func (rcvr *BackgroundSubtractorMOG2) SetHistory(history int) {
	BackgroundSubtractorMOG2Native_setHistory_0(rcvr.GetNativeObjAddr(), history)
	return
}
func (rcvr *BackgroundSubtractorMOG2) SetNMixtures(nmixtures int) {
	BackgroundSubtractorMOG2Native_setNMixtures_0(rcvr.GetNativeObjAddr(), nmixtures)
	return
}
func (rcvr *BackgroundSubtractorMOG2) SetShadowThreshold(threshold float64) {
	BackgroundSubtractorMOG2Native_setShadowThreshold_0(rcvr.GetNativeObjAddr(), threshold)
	return
}
func (rcvr *BackgroundSubtractorMOG2) SetShadowValue(value int) {
	BackgroundSubtractorMOG2Native_setShadowValue_0(rcvr.GetNativeObjAddr(), value)
	return
}
func (rcvr *BackgroundSubtractorMOG2) SetVarInit(varInit float64) {
	BackgroundSubtractorMOG2Native_setVarInit_0(rcvr.GetNativeObjAddr(), varInit)
	return
}
func (rcvr *BackgroundSubtractorMOG2) SetVarMax(varMax float64) {
	BackgroundSubtractorMOG2Native_setVarMax_0(rcvr.GetNativeObjAddr(), varMax)
	return
}
func (rcvr *BackgroundSubtractorMOG2) SetVarMin(varMin float64) {
	BackgroundSubtractorMOG2Native_setVarMin_0(rcvr.GetNativeObjAddr(), varMin)
	return
}
func (rcvr *BackgroundSubtractorMOG2) SetVarThreshold(varThreshold float64) {
	BackgroundSubtractorMOG2Native_setVarThreshold_0(rcvr.GetNativeObjAddr(), varThreshold)
	return
}
func (rcvr *BackgroundSubtractorMOG2) SetVarThresholdGen(varThresholdGen float64) {
	BackgroundSubtractorMOG2Native_setVarThresholdGen_0(rcvr.GetNativeObjAddr(), varThresholdGen)
	return
}
