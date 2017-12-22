package core

import (
	"fmt"
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

const _channelsMatOfDouble = 1

var _depthMatOfDouble = CvTypeCV_64F

type MatOfDouble struct {
	*Mat
}

func NewMatOfDouble() (rcvr *MatOfDouble) {
	rcvr = &MatOfDouble{}
	rcvr.Mat = NewMat2()
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func NewMatOfDouble2(addr int64) (rcvr *MatOfDouble) {
	rcvr = &MatOfDouble{}
	rcvr.Mat = NewMat(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	if !rcvr.Empty() && rcvr.CheckVector2(_channelsMatOfDouble, _depthMatOfDouble) < 0 {
		Throw(NewIllegalArgumentException("Incompatible Mat"))
	}
	return
}
func NewMatOfDouble3(m *Mat) (rcvr *MatOfDouble) {
	rcvr = &MatOfDouble{}
	rcvr.Mat = NewMat8(m, RangeAll())
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	if !rcvr.Empty() && rcvr.CheckVector2(_channelsMatOfDouble, _depthMatOfDouble) < 0 {
		Throw(NewIllegalArgumentException("Incompatible Mat"))
	}
	return
}
func NewMatOfDouble4(a ...float64) (rcvr *MatOfDouble) {
	rcvr = &MatOfDouble{}
	rcvr.Mat = NewMat2()
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	rcvr.FromArray(a)
	return
}
func (rcvr *MatOfDouble) Alloc(elemNumber int) {
	if elemNumber > 0 {
		rcvr.Create(elemNumber, 1, CvTypeMakeType(_depthMatOfDouble, _channelsMatOfDouble))
	}
}
func (rcvr *MatOfDouble) FromArray(a []float64) {
	if a == nil || len(a) == 0 {
		return
	}
	num := len(a) / _channelsMatOfDouble
	rcvr.Alloc(num)
	rcvr.PutD(0, 0, a)
}
func MatOfDoubleFromNativeAddr(addr int64) *MatOfDouble {
	return NewMatOfDouble2(addr)
}
func (rcvr *MatOfDouble) ToArray() []float64 {
	num := rcvr.CheckVector2(_channelsMatOfDouble, _depthMatOfDouble)
	if num < 0 {
		Throw(NewRuntimeException(fmt.Sprintf("%v%v", "Native Mat has unexpected type or size: ", rcvr)))
	}
	a := make([]float64, num*_channelsMatOfDouble)
	if num == 0 {
		return a
	}
	rcvr.GetD(0, 0, a)
	return a
}
