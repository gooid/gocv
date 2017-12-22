package video

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

type DualTVL1OpticalFlow struct {
	*DenseOpticalFlow
}

func NewDualTVL1OpticalFlow(addr int64) (rcvr *DualTVL1OpticalFlow) {
	rcvr = &DualTVL1OpticalFlow{}
	rcvr.DenseOpticalFlow = NewDenseOpticalFlow(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func DualTVL1OpticalFlowCreate(tau float64, lambda float64, theta float64, nscales int, warps int, epsilon float64, innnerIterations int, outerIterations int, scaleStep float64, gamma float64, medianFiltering int, useInitialFlow bool) *DualTVL1OpticalFlow {
	retVal := NewDualTVL1OpticalFlow(DualTVL1OpticalFlowNative_create_0(tau, lambda, theta, nscales, warps, epsilon, innnerIterations, outerIterations, scaleStep, gamma, medianFiltering, useInitialFlow))
	return retVal
}
func DualTVL1OpticalFlowCreate2() *DualTVL1OpticalFlow {
	retVal := NewDualTVL1OpticalFlow(DualTVL1OpticalFlowNative_create_1())
	return retVal
}
func (rcvr *DualTVL1OpticalFlow) finalize() {
	DualTVL1OpticalFlowNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *DualTVL1OpticalFlow) GetEpsilon() float64 {
	retVal := DualTVL1OpticalFlowNative_getEpsilon_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *DualTVL1OpticalFlow) GetGamma() float64 {
	retVal := DualTVL1OpticalFlowNative_getGamma_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *DualTVL1OpticalFlow) GetInnerIterations() int {
	retVal := DualTVL1OpticalFlowNative_getInnerIterations_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *DualTVL1OpticalFlow) GetLambda() float64 {
	retVal := DualTVL1OpticalFlowNative_getLambda_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *DualTVL1OpticalFlow) GetMedianFiltering() int {
	retVal := DualTVL1OpticalFlowNative_getMedianFiltering_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *DualTVL1OpticalFlow) GetOuterIterations() int {
	retVal := DualTVL1OpticalFlowNative_getOuterIterations_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *DualTVL1OpticalFlow) GetScaleStep() float64 {
	retVal := DualTVL1OpticalFlowNative_getScaleStep_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *DualTVL1OpticalFlow) GetScalesNumber() int {
	retVal := DualTVL1OpticalFlowNative_getScalesNumber_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *DualTVL1OpticalFlow) GetTau() float64 {
	retVal := DualTVL1OpticalFlowNative_getTau_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *DualTVL1OpticalFlow) GetTheta() float64 {
	retVal := DualTVL1OpticalFlowNative_getTheta_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *DualTVL1OpticalFlow) GetUseInitialFlow() bool {
	retVal := DualTVL1OpticalFlowNative_getUseInitialFlow_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *DualTVL1OpticalFlow) GetWarpingsNumber() int {
	retVal := DualTVL1OpticalFlowNative_getWarpingsNumber_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *DualTVL1OpticalFlow) SetEpsilon(val float64) {
	DualTVL1OpticalFlowNative_setEpsilon_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *DualTVL1OpticalFlow) SetGamma(val float64) {
	DualTVL1OpticalFlowNative_setGamma_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *DualTVL1OpticalFlow) SetInnerIterations(val int) {
	DualTVL1OpticalFlowNative_setInnerIterations_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *DualTVL1OpticalFlow) SetLambda(val float64) {
	DualTVL1OpticalFlowNative_setLambda_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *DualTVL1OpticalFlow) SetMedianFiltering(val int) {
	DualTVL1OpticalFlowNative_setMedianFiltering_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *DualTVL1OpticalFlow) SetOuterIterations(val int) {
	DualTVL1OpticalFlowNative_setOuterIterations_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *DualTVL1OpticalFlow) SetScaleStep(val float64) {
	DualTVL1OpticalFlowNative_setScaleStep_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *DualTVL1OpticalFlow) SetScalesNumber(val int) {
	DualTVL1OpticalFlowNative_setScalesNumber_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *DualTVL1OpticalFlow) SetTau(val float64) {
	DualTVL1OpticalFlowNative_setTau_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *DualTVL1OpticalFlow) SetTheta(val float64) {
	DualTVL1OpticalFlowNative_setTheta_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *DualTVL1OpticalFlow) SetUseInitialFlow(val bool) {
	DualTVL1OpticalFlowNative_setUseInitialFlow_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *DualTVL1OpticalFlow) SetWarpingsNumber(val int) {
	DualTVL1OpticalFlowNative_setWarpingsNumber_0(rcvr.GetNativeObjAddr(), val)
	return
}
