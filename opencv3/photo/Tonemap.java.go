package photo

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

type Tonemap struct {
	*Algorithm
}

func NewTonemap(addr int64) (rcvr *Tonemap) {
	rcvr = &Tonemap{}
	rcvr.Algorithm = NewAlgorithm(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func (rcvr *Tonemap) finalize() {
	TonemapNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *Tonemap) GetGamma() float32 {
	retVal := TonemapNative_getGamma_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *Tonemap) Process(src *Mat, dst *Mat) {
	TonemapNative_process_0(rcvr.GetNativeObjAddr(), src.GetNativeObjAddr(), dst.GetNativeObjAddr())
	return
}
func (rcvr *Tonemap) SetGamma(gamma float32) {
	TonemapNative_setGamma_0(rcvr.GetNativeObjAddr(), gamma)
	return
}
