package core

import (
	"fmt"
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

const _channelsMatOfFloat4 = 4

var _depthMatOfFloat4 = CvTypeCV_32F

type MatOfFloat4 struct {
	*Mat
}

func NewMatOfFloat41() (rcvr *MatOfFloat4) {
	rcvr = &MatOfFloat4{}
	rcvr.Mat = NewMat2()
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func NewMatOfFloat42(addr int64) (rcvr *MatOfFloat4) {
	rcvr = &MatOfFloat4{}
	rcvr.Mat = NewMat(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	if !rcvr.Empty() && rcvr.CheckVector2(_channelsMatOfFloat4, _depthMatOfFloat4) < 0 {
		Throw(NewIllegalArgumentException("Incompatible Mat"))
	}
	return
}
func NewMatOfFloat43(m *Mat) (rcvr *MatOfFloat4) {
	rcvr = &MatOfFloat4{}
	rcvr.Mat = NewMat8(m, RangeAll())
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	if !rcvr.Empty() && rcvr.CheckVector2(_channelsMatOfFloat4, _depthMatOfFloat4) < 0 {
		Throw(NewIllegalArgumentException("Incompatible Mat"))
	}
	return
}
func NewMatOfFloat44(a ...float32) (rcvr *MatOfFloat4) {
	rcvr = &MatOfFloat4{}
	rcvr.Mat = NewMat2()
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	rcvr.FromArray(a)
	return
}
func (rcvr *MatOfFloat4) Alloc(elemNumber int) {
	if elemNumber > 0 {
		rcvr.Create(elemNumber, 1, CvTypeMakeType(_depthMatOfFloat4, _channelsMatOfFloat4))
	}
}
func (rcvr *MatOfFloat4) FromArray(a []float32) {
	if a == nil || len(a) == 0 {
		return
	}
	num := len(a) / _channelsMatOfFloat4
	rcvr.Alloc(num)
	rcvr.PutF(0, 0, a)
}
func MatOfFloat4FromNativeAddr(addr int64) *MatOfFloat4 {
	return NewMatOfFloat42(addr)
}
func (rcvr *MatOfFloat4) ToArray() []float32 {
	num := rcvr.CheckVector2(_channelsMatOfFloat4, _depthMatOfFloat4)
	if num < 0 {
		Throw(NewRuntimeException(fmt.Sprintf("%v%v", "Native Mat has unexpected type or size: ", rcvr)))
	}
	a := make([]float32, num*_channelsMatOfFloat4)
	if num == 0 {
		return a
	}
	rcvr.GetF(0, 0, a)
	return a
}
