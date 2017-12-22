package photo

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

type TonemapMantiuk struct {
	*Tonemap
}

func NewTonemapMantiuk(addr int64) (rcvr *TonemapMantiuk) {
	rcvr = &TonemapMantiuk{}
	rcvr.Tonemap = NewTonemap(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func (rcvr *TonemapMantiuk) finalize() {
	TonemapMantiukNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *TonemapMantiuk) GetSaturation() float32 {
	retVal := TonemapMantiukNative_getSaturation_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *TonemapMantiuk) GetScale() float32 {
	retVal := TonemapMantiukNative_getScale_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *TonemapMantiuk) SetSaturation(saturation float32) {
	TonemapMantiukNative_setSaturation_0(rcvr.GetNativeObjAddr(), saturation)
	return
}
func (rcvr *TonemapMantiuk) SetScale(scale float32) {
	TonemapMantiukNative_setScale_0(rcvr.GetNativeObjAddr(), scale)
	return
}
