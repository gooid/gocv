package core

import (
	"fmt"
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

const _channelsMatOfFloat6 = 6

var _depthMatOfFloat6 = CvTypeCV_32F

type MatOfFloat6 struct {
	*Mat
}

func NewMatOfFloat6() (rcvr *MatOfFloat6) {
	rcvr = &MatOfFloat6{}
	rcvr.Mat = NewMat2()
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func NewMatOfFloat62(addr int64) (rcvr *MatOfFloat6) {
	rcvr = &MatOfFloat6{}
	rcvr.Mat = NewMat(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	if !rcvr.Empty() && rcvr.CheckVector2(_channelsMatOfFloat6, _depthMatOfFloat6) < 0 {
		Throw(NewIllegalArgumentException("Incompatible Mat"))
	}
	return
}
func NewMatOfFloat63(m *Mat) (rcvr *MatOfFloat6) {
	rcvr = &MatOfFloat6{}
	rcvr.Mat = NewMat8(m, RangeAll())
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	if !rcvr.Empty() && rcvr.CheckVector2(_channelsMatOfFloat6, _depthMatOfFloat6) < 0 {
		Throw(NewIllegalArgumentException("Incompatible Mat"))
	}
	return
}
func NewMatOfFloat64(a ...float32) (rcvr *MatOfFloat6) {
	rcvr = &MatOfFloat6{}
	rcvr.Mat = NewMat2()
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	rcvr.FromArray(a)
	return
}
func (rcvr *MatOfFloat6) Alloc(elemNumber int) {
	if elemNumber > 0 {
		rcvr.Create(elemNumber, 1, CvTypeMakeType(_depthMatOfFloat6, _channelsMatOfFloat6))
	}
}
func (rcvr *MatOfFloat6) FromArray(a []float32) {
	if a == nil || len(a) == 0 {
		return
	}
	num := len(a) / _channelsMatOfFloat6
	rcvr.Alloc(num)
	rcvr.PutF(0, 0, a)
}
func MatOfFloat6FromNativeAddr(addr int64) *MatOfFloat6 {
	return NewMatOfFloat62(addr)
}
func (rcvr *MatOfFloat6) ToArray() []float32 {
	num := rcvr.CheckVector2(_channelsMatOfFloat6, _depthMatOfFloat6)
	if num < 0 {
		Throw(NewRuntimeException(fmt.Sprintf("%v%v", "Native Mat has unexpected type or size: ", rcvr)))
	}
	a := make([]float32, num*_channelsMatOfFloat6)
	if num == 0 {
		return a
	}
	rcvr.GetF(0, 0, a)
	return a
}
