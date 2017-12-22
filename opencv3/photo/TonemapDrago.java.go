package photo

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

type TonemapDrago struct {
	*Tonemap
}

func NewTonemapDrago(addr int64) (rcvr *TonemapDrago) {
	rcvr = &TonemapDrago{}
	rcvr.Tonemap = NewTonemap(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func (rcvr *TonemapDrago) finalize() {
	TonemapDragoNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *TonemapDrago) GetBias() float32 {
	retVal := TonemapDragoNative_getBias_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *TonemapDrago) GetSaturation() float32 {
	retVal := TonemapDragoNative_getSaturation_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *TonemapDrago) SetBias(bias float32) {
	TonemapDragoNative_setBias_0(rcvr.GetNativeObjAddr(), bias)
	return
}
func (rcvr *TonemapDrago) SetSaturation(saturation float32) {
	TonemapDragoNative_setSaturation_0(rcvr.GetNativeObjAddr(), saturation)
	return
}
