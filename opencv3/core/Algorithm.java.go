package core

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

type Algorithm struct {
	nativeObj int64
}

func NewAlgorithm(addr int64) (rcvr *Algorithm) {
	rcvr = &Algorithm{}
	rcvr.nativeObj = addr
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func (rcvr *Algorithm) Clear() {
	AlgorithmNative_clear_0(rcvr.nativeObj)
	return
}
func (rcvr *Algorithm) finalize() {
	AlgorithmNative_delete(rcvr.nativeObj)
}
func (rcvr *Algorithm) GetDefaultName() string {
	retVal := AlgorithmNative_getDefaultName_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *Algorithm) GetNativeObjAddr() int64 {
	return rcvr.nativeObj
}
func (rcvr *Algorithm) Save(filename string) {
	AlgorithmNative_save_0(rcvr.nativeObj, filename)
	return
}
