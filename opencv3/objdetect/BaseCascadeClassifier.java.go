package objdetect

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

type BaseCascadeClassifier struct {
	*Algorithm
}

func NewBaseCascadeClassifier(addr int64) (rcvr *BaseCascadeClassifier) {
	rcvr = &BaseCascadeClassifier{}
	rcvr.Algorithm = NewAlgorithm(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func (rcvr *BaseCascadeClassifier) finalize() {
	BaseCascadeClassifierNative_delete(rcvr.GetNativeObjAddr())
}
