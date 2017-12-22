package photo

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

type CalibrateCRF struct {
	*Algorithm
}

func NewCalibrateCRF(addr int64) (rcvr *CalibrateCRF) {
	rcvr = &CalibrateCRF{}
	rcvr.Algorithm = NewAlgorithm(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func (rcvr *CalibrateCRF) finalize() {
	CalibrateCRFNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *CalibrateCRF) Process(src []*Mat, dst *Mat, times *Mat) {
	src_mat := ConvertersVectorMat(src)
	CalibrateCRFNative_process_0(rcvr.GetNativeObjAddr(), src_mat.GetNativeObjAddr(), dst.GetNativeObjAddr(), times.GetNativeObjAddr())
	return
}
