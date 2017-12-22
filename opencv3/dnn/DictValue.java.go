package dnn

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

type DictValue struct {
	nativeObj int64
}

func NewDictValue(addr int64) (rcvr *DictValue) {
	rcvr = &DictValue{}
	rcvr.nativeObj = addr
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func NewDictValue2(s string) (rcvr *DictValue) {
	rcvr = &DictValue{}
	rcvr.nativeObj = DictValueNative_DictValue_0(s)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
	return
}
func NewDictValue3(p float64) (rcvr *DictValue) {
	rcvr = &DictValue{}
	rcvr.nativeObj = DictValueNative_DictValue_1(p)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
	return
}
func NewDictValue4(i int) (rcvr *DictValue) {
	rcvr = &DictValue{}
	rcvr.nativeObj = DictValueNative_DictValue_2(i)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
	return
}
func (rcvr *DictValue) finalize() {
	DictValueNative_delete(rcvr.nativeObj)
}
func (rcvr *DictValue) GetIntValue(idx int) int {
	retVal := DictValueNative_getIntValue_0(rcvr.nativeObj, idx)
	return retVal
}
func (rcvr *DictValue) GetIntValue2() int {
	retVal := DictValueNative_getIntValue_1(rcvr.nativeObj)
	return retVal
}
func (rcvr *DictValue) GetNativeObjAddr() int64 {
	return rcvr.nativeObj
}
func (rcvr *DictValue) GetRealValue(idx int) float64 {
	retVal := DictValueNative_getRealValue_0(rcvr.nativeObj, idx)
	return retVal
}
func (rcvr *DictValue) GetRealValue2() float64 {
	retVal := DictValueNative_getRealValue_1(rcvr.nativeObj)
	return retVal
}
func (rcvr *DictValue) GetStringValue(idx int) string {
	retVal := DictValueNative_getStringValue_0(rcvr.nativeObj, idx)
	return retVal
}
func (rcvr *DictValue) GetStringValue2() string {
	retVal := DictValueNative_getStringValue_1(rcvr.nativeObj)
	return retVal
}
func (rcvr *DictValue) IsInt() bool {
	retVal := DictValueNative_isInt_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *DictValue) IsReal() bool {
	retVal := DictValueNative_isReal_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *DictValue) IsString() bool {
	retVal := DictValueNative_isString_0(rcvr.nativeObj)
	return retVal
}
