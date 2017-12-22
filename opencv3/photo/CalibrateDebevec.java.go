package photo

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

type CalibrateDebevec struct {
	*CalibrateCRF
}

func NewCalibrateDebevec(addr int64) (rcvr *CalibrateDebevec) {
	rcvr = &CalibrateDebevec{}
	rcvr.CalibrateCRF = NewCalibrateCRF(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func (rcvr *CalibrateDebevec) finalize() {
	CalibrateDebevecNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *CalibrateDebevec) GetLambda() float32 {
	retVal := CalibrateDebevecNative_getLambda_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *CalibrateDebevec) GetRandom() bool {
	retVal := CalibrateDebevecNative_getRandom_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *CalibrateDebevec) GetSamples() int {
	retVal := CalibrateDebevecNative_getSamples_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *CalibrateDebevec) SetLambda(lambda float32) {
	CalibrateDebevecNative_setLambda_0(rcvr.GetNativeObjAddr(), lambda)
	return
}
func (rcvr *CalibrateDebevec) SetRandom(random bool) {
	CalibrateDebevecNative_setRandom_0(rcvr.GetNativeObjAddr(), random)
	return
}
func (rcvr *CalibrateDebevec) SetSamples(samples int) {
	CalibrateDebevecNative_setSamples_0(rcvr.GetNativeObjAddr(), samples)
	return
}
