package imgproc

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

type CLAHE struct {
	*Algorithm
}

func NewCLAHE(addr int64) (rcvr *CLAHE) {
	rcvr = &CLAHE{}
	rcvr.Algorithm = NewAlgorithm(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func (rcvr *CLAHE) Apply(src *Mat, dst *Mat) {
	CLAHENative_apply_0(rcvr.GetNativeObjAddr(), src.GetNativeObjAddr(), dst.GetNativeObjAddr())
	return
}

func (rcvr *CLAHE) CollectGarbage() {
	CLAHENative_collectGarbage_0(rcvr.GetNativeObjAddr())
	return
}

func (rcvr *CLAHE) finalize() {
	CLAHENative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *CLAHE) GetClipLimit() float64 {
	retVal := CLAHENative_getClipLimit_0(rcvr.GetNativeObjAddr())
	return retVal
}

func (rcvr *CLAHE) GetTilesGridSize() *Size {
	retVal := NewSize4(CLAHENative_getTilesGridSize_0(rcvr.GetNativeObjAddr()))
	return retVal
}

func (rcvr *CLAHE) SetClipLimit(clipLimit float64) {
	CLAHENative_setClipLimit_0(rcvr.GetNativeObjAddr(), clipLimit)
	return
}

func (rcvr *CLAHE) SetTilesGridSize(tileGridSize *Size) {
	CLAHENative_setTilesGridSize_0(rcvr.GetNativeObjAddr(), tileGridSize.Width, tileGridSize.Height)
	return
}
