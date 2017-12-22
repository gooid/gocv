package core

import (
	"fmt"
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

const _channelsMatOfFloat = 1

var _depthMatOfFloat = CvTypeCV_32F

type MatOfFloat struct {
	*Mat
}

func NewMatOfFloat() (rcvr *MatOfFloat) {
	rcvr = &MatOfFloat{}
	rcvr.Mat = NewMat2()
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func NewMatOfFloat2(addr int64) (rcvr *MatOfFloat) {
	rcvr = &MatOfFloat{}
	rcvr.Mat = NewMat(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	if !rcvr.Empty() && rcvr.CheckVector2(_channelsMatOfFloat, _depthMatOfFloat) < 0 {
		Throw(NewIllegalArgumentException("Incompatible Mat"))
	}
	return
}
func NewMatOfFloat3(m *Mat) (rcvr *MatOfFloat) {
	rcvr = &MatOfFloat{}
	rcvr.Mat = NewMat8(m, RangeAll())
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	if !rcvr.Empty() && rcvr.CheckVector2(_channelsMatOfFloat, _depthMatOfFloat) < 0 {
		Throw(NewIllegalArgumentException("Incompatible Mat"))
	}
	return
}
func NewMatOfFloat4(a ...float32) (rcvr *MatOfFloat) {
	rcvr = &MatOfFloat{}
	rcvr.Mat = NewMat2()
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	rcvr.FromArray(a)
	return
}
func (rcvr *MatOfFloat) Alloc(elemNumber int) {
	if elemNumber > 0 {
		rcvr.Create(elemNumber, 1, CvTypeMakeType(_depthMatOfFloat, _channelsMatOfFloat))
	}
}
func (rcvr *MatOfFloat) FromArray(a []float32) {
	if a == nil || len(a) == 0 {
		return
	}
	num := len(a) / _channelsMatOfFloat
	rcvr.Alloc(num)
	rcvr.PutF(0, 0, a)
}
func MatOfFloatFromNativeAddr(addr int64) *MatOfFloat {
	return NewMatOfFloat2(addr)
}
func (rcvr *MatOfFloat) ToArray() []float32 {
	num := rcvr.CheckVector2(_channelsMatOfFloat, _depthMatOfFloat)
	if num < 0 {
		Throw(NewRuntimeException(fmt.Sprintf("%v%v", "Native Mat has unexpected type or size: ", rcvr)))
	}
	a := make([]float32, num*_channelsMatOfFloat)
	if num == 0 {
		return a
	}
	rcvr.GetF(0, 0, a)
	return a
}
