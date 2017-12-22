package photo

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

type CalibrateRobertson struct {
	*CalibrateCRF
}

func NewCalibrateRobertson(addr int64) (rcvr *CalibrateRobertson) {
	rcvr = &CalibrateRobertson{}
	rcvr.CalibrateCRF = NewCalibrateCRF(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func (rcvr *CalibrateRobertson) finalize() {
	CalibrateRobertsonNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *CalibrateRobertson) GetMaxIter() int {
	retVal := CalibrateRobertsonNative_getMaxIter_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *CalibrateRobertson) GetRadiance() *Mat {
	retVal := NewMat(CalibrateRobertsonNative_getRadiance_0(rcvr.GetNativeObjAddr()))
	return retVal
}
func (rcvr *CalibrateRobertson) GetThreshold() float32 {
	retVal := CalibrateRobertsonNative_getThreshold_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *CalibrateRobertson) SetMaxIter(max_iter int) {
	CalibrateRobertsonNative_setMaxIter_0(rcvr.GetNativeObjAddr(), max_iter)
	return
}
func (rcvr *CalibrateRobertson) SetThreshold(threshold float32) {
	CalibrateRobertsonNative_setThreshold_0(rcvr.GetNativeObjAddr(), threshold)
	return
}
