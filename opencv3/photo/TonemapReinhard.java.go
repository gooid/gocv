package photo

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

type TonemapReinhard struct {
	*Tonemap
}

func NewTonemapReinhard(addr int64) (rcvr *TonemapReinhard) {
	rcvr = &TonemapReinhard{}
	rcvr.Tonemap = NewTonemap(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func (rcvr *TonemapReinhard) finalize() {
	TonemapReinhardNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *TonemapReinhard) GetColorAdaptation() float32 {
	retVal := TonemapReinhardNative_getColorAdaptation_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *TonemapReinhard) GetIntensity() float32 {
	retVal := TonemapReinhardNative_getIntensity_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *TonemapReinhard) GetLightAdaptation() float32 {
	retVal := TonemapReinhardNative_getLightAdaptation_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *TonemapReinhard) SetColorAdaptation(color_adapt float32) {
	TonemapReinhardNative_setColorAdaptation_0(rcvr.GetNativeObjAddr(), color_adapt)
	return
}
func (rcvr *TonemapReinhard) SetIntensity(intensity float32) {
	TonemapReinhardNative_setIntensity_0(rcvr.GetNativeObjAddr(), intensity)
	return
}
func (rcvr *TonemapReinhard) SetLightAdaptation(light_adapt float32) {
	TonemapReinhardNative_setLightAdaptation_0(rcvr.GetNativeObjAddr(), light_adapt)
	return
}
