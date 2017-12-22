package dnn

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

type Importer struct {
	*Algorithm
}

func NewImporter(addr int64) (rcvr *Importer) {
	rcvr = &Importer{}
	rcvr.Algorithm = NewAlgorithm(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func (rcvr *Importer) finalize() {
	ImporterNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *Importer) PopulateNet(net *Net) {
	ImporterNative_populateNet_0(rcvr.GetNativeObjAddr(), net.nativeObj)
	return
}
