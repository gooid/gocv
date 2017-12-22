package core

import (
	"fmt"
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

const _channelsMatOfInt = 1

var _depthMatOfInt = CvTypeCV_32S

type MatOfInt struct {
	*Mat
}

func NewMatOfInt() (rcvr *MatOfInt) {
	rcvr = &MatOfInt{}
	rcvr.Mat = NewMat2()
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func NewMatOfInt2(addr int64) (rcvr *MatOfInt) {
	rcvr = &MatOfInt{}
	rcvr.Mat = NewMat(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	if !rcvr.Empty() && rcvr.CheckVector2(_channelsMatOfInt, _depthMatOfInt) < 0 {
		Throw(NewIllegalArgumentException("Incompatible Mat"))
	}
	return
}
func NewMatOfInt3(m *Mat) (rcvr *MatOfInt) {
	rcvr = &MatOfInt{}
	rcvr.Mat = NewMat8(m, RangeAll())
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	if !rcvr.Empty() && rcvr.CheckVector2(_channelsMatOfInt, _depthMatOfInt) < 0 {
		Throw(NewIllegalArgumentException("Incompatible Mat"))
	}
	return
}
func NewMatOfInt4(a ...int32) (rcvr *MatOfInt) {
	rcvr = &MatOfInt{}
	rcvr.Mat = NewMat2()
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	rcvr.FromArray(a)
	return
}
func (rcvr *MatOfInt) Alloc(elemNumber int) {
	if elemNumber > 0 {
		rcvr.Create(elemNumber, 1, CvTypeMakeType(_depthMatOfInt, _channelsMatOfInt))
	}
}
func (rcvr *MatOfInt) FromArray(a []int32) {
	if a == nil || len(a) == 0 {
		return
	}
	num := len(a) / _channelsMatOfInt
	rcvr.Alloc(num)
	rcvr.PutI(0, 0, a)
}
func MatOfIntFromNativeAddr(addr int64) *MatOfInt {
	return NewMatOfInt2(addr)
}
func (rcvr *MatOfInt) ToArray() []int32 {
	num := rcvr.CheckVector2(_channelsMatOfInt, _depthMatOfInt)
	if num < 0 {
		Throw(NewRuntimeException(fmt.Sprintf("%v%v", "Native Mat has unexpected type or size: ", rcvr)))
	}
	a := make([]int32, num*_channelsMatOfInt)
	if num == 0 {
		return a
	}
	rcvr.GetI(0, 0, a)
	return a
}
