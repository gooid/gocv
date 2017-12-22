package photo

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

type TonemapDurand struct {
	*Tonemap
}

func NewTonemapDurand(addr int64) (rcvr *TonemapDurand) {
	rcvr = &TonemapDurand{}
	rcvr.Tonemap = NewTonemap(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func (rcvr *TonemapDurand) finalize() {
	TonemapDurandNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *TonemapDurand) GetContrast() float32 {
	retVal := TonemapDurandNative_getContrast_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *TonemapDurand) GetSaturation() float32 {
	retVal := TonemapDurandNative_getSaturation_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *TonemapDurand) GetSigmaColor() float32 {
	retVal := TonemapDurandNative_getSigmaColor_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *TonemapDurand) GetSigmaSpace() float32 {
	retVal := TonemapDurandNative_getSigmaSpace_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *TonemapDurand) SetContrast(contrast float32) {
	TonemapDurandNative_setContrast_0(rcvr.GetNativeObjAddr(), contrast)
	return
}
func (rcvr *TonemapDurand) SetSaturation(saturation float32) {
	TonemapDurandNative_setSaturation_0(rcvr.GetNativeObjAddr(), saturation)
	return
}
func (rcvr *TonemapDurand) SetSigmaColor(sigma_color float32) {
	TonemapDurandNative_setSigmaColor_0(rcvr.GetNativeObjAddr(), sigma_color)
	return
}
func (rcvr *TonemapDurand) SetSigmaSpace(sigma_space float32) {
	TonemapDurandNative_setSigmaSpace_0(rcvr.GetNativeObjAddr(), sigma_space)
	return
}
