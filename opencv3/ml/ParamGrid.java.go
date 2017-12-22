package ml

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

type ParamGrid struct {
	nativeObj int64
}

func NewParamGrid(addr int64) (rcvr *ParamGrid) {
	rcvr = &ParamGrid{}
	rcvr.nativeObj = addr
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func ParamGridCreate(minVal float64, maxVal float64, logstep float64) *ParamGrid {
	retVal := NewParamGrid(ParamGridNative_create_0(minVal, maxVal, logstep))
	return retVal
}
func ParamGridCreate2() *ParamGrid {
	retVal := NewParamGrid(ParamGridNative_create_1())
	return retVal
}
func (rcvr *ParamGrid) finalize() {
	ParamGridNative_delete(rcvr.nativeObj)
}
func (rcvr *ParamGrid) GetNativeObjAddr() int64 {
	return rcvr.nativeObj
}
func (rcvr *ParamGrid) Get_logStep() float64 {
	retVal := ParamGridNative_get_logStep_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *ParamGrid) Get_maxVal() float64 {
	retVal := ParamGridNative_get_maxVal_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *ParamGrid) Get_minVal() float64 {
	retVal := ParamGridNative_get_minVal_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *ParamGrid) Set_logStep(logStep float64) {
	ParamGridNative_set_logStep_0(rcvr.nativeObj, logStep)
	return
}
func (rcvr *ParamGrid) Set_maxVal(maxVal float64) {
	ParamGridNative_set_maxVal_0(rcvr.nativeObj, maxVal)
	return
}
func (rcvr *ParamGrid) Set_minVal(minVal float64) {
	ParamGridNative_set_minVal_0(rcvr.nativeObj, minVal)
	return
}
