package ml

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

const SVMSGDSGD = 0
const SVMSGDASGD = 1
const SVMSGDSOFT_MARGIN = 0
const SVMSGDHARD_MARGIN = 1

type SVMSGD struct {
	*StatModel
}

func NewSVMSGD(addr int64) (rcvr *SVMSGD) {
	rcvr = &SVMSGD{}
	rcvr.StatModel = NewStatModel(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func SVMSGDCreate() *SVMSGD {
	retVal := NewSVMSGD(SVMSGDNative_create_0())
	return retVal
}
func (rcvr *SVMSGD) finalize() {
	SVMSGDNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *SVMSGD) GetInitialStepSize() float32 {
	retVal := SVMSGDNative_getInitialStepSize_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *SVMSGD) GetMarginRegularization() float32 {
	retVal := SVMSGDNative_getMarginRegularization_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *SVMSGD) GetMarginType() int {
	retVal := SVMSGDNative_getMarginType_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *SVMSGD) GetShift() float32 {
	retVal := SVMSGDNative_getShift_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *SVMSGD) GetStepDecreasingPower() float32 {
	retVal := SVMSGDNative_getStepDecreasingPower_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *SVMSGD) GetSvmsgdType() int {
	retVal := SVMSGDNative_getSvmsgdType_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *SVMSGD) GetTermCriteria() *TermCriteria {
	retVal := NewTermCriteria3(SVMSGDNative_getTermCriteria_0(rcvr.GetNativeObjAddr()))
	return retVal
}
func (rcvr *SVMSGD) GetWeights() *Mat {
	retVal := NewMat(SVMSGDNative_getWeights_0(rcvr.GetNativeObjAddr()))
	return retVal
}
func SVMSGDLoad(filepath string, nodeName string) *SVMSGD {
	retVal := NewSVMSGD(SVMSGDNative_load_0(filepath, nodeName))
	return retVal
}
func SVMSGDLoad2(filepath string) *SVMSGD {
	retVal := NewSVMSGD(SVMSGDNative_load_1(filepath))
	return retVal
}
func (rcvr *SVMSGD) SetInitialStepSize(InitialStepSize float32) {
	SVMSGDNative_setInitialStepSize_0(rcvr.GetNativeObjAddr(), InitialStepSize)
	return
}
func (rcvr *SVMSGD) SetMarginRegularization(marginRegularization float32) {
	SVMSGDNative_setMarginRegularization_0(rcvr.GetNativeObjAddr(), marginRegularization)
	return
}
func (rcvr *SVMSGD) SetMarginType(marginType int) {
	SVMSGDNative_setMarginType_0(rcvr.GetNativeObjAddr(), marginType)
	return
}
func (rcvr *SVMSGD) SetOptimalParameters(svmsgdType int, marginType int) {
	SVMSGDNative_setOptimalParameters_0(rcvr.GetNativeObjAddr(), svmsgdType, marginType)
	return
}
func (rcvr *SVMSGD) SetOptimalParameters2() {
	SVMSGDNative_setOptimalParameters_1(rcvr.GetNativeObjAddr())
	return
}
func (rcvr *SVMSGD) SetStepDecreasingPower(stepDecreasingPower float32) {
	SVMSGDNative_setStepDecreasingPower_0(rcvr.GetNativeObjAddr(), stepDecreasingPower)
	return
}
func (rcvr *SVMSGD) SetSvmsgdType(svmsgdType int) {
	SVMSGDNative_setSvmsgdType_0(rcvr.GetNativeObjAddr(), svmsgdType)
	return
}
func (rcvr *SVMSGD) SetTermCriteria(val *TermCriteria) {
	SVMSGDNative_setTermCriteria_0(rcvr.GetNativeObjAddr(), val.Type, val.MaxCount, val.Epsilon)
	return
}
