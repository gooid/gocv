package core

import (
	"fmt"
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

const _channelsMatOfInt4 = 4

var _depthMatOfInt4 = CvTypeCV_32S

type MatOfInt4 struct {
	*Mat
}

func NewMatOfInt41() (rcvr *MatOfInt4) {
	rcvr = &MatOfInt4{}
	rcvr.Mat = NewMat2()
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func NewMatOfInt42(addr int64) (rcvr *MatOfInt4) {
	rcvr = &MatOfInt4{}
	rcvr.Mat = NewMat(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	if !rcvr.Empty() && rcvr.CheckVector2(_channelsMatOfInt4, _depthMatOfInt4) < 0 {
		Throw(NewIllegalArgumentException("Incompatible Mat"))
	}
	return
}
func NewMatOfInt43(m *Mat) (rcvr *MatOfInt4) {
	rcvr = &MatOfInt4{}
	rcvr.Mat = NewMat8(m, RangeAll())
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	if !rcvr.Empty() && rcvr.CheckVector2(_channelsMatOfInt4, _depthMatOfInt4) < 0 {
		Throw(NewIllegalArgumentException("Incompatible Mat"))
	}
	return
}
func NewMatOfInt44(a ...int32) (rcvr *MatOfInt4) {
	rcvr = &MatOfInt4{}
	rcvr.Mat = NewMat2()
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	rcvr.FromArray(a)
	return
}
func (rcvr *MatOfInt4) Alloc(elemNumber int) {
	if elemNumber > 0 {
		rcvr.Create(elemNumber, 1, CvTypeMakeType(_depthMatOfInt4, _channelsMatOfInt4))
	}
}
func (rcvr *MatOfInt4) FromArray(a []int32) {
	if a == nil || len(a) == 0 {
		return
	}
	num := len(a) / _channelsMatOfInt4
	rcvr.Alloc(num)
	rcvr.PutI(0, 0, a)
}
func MatOfInt4FromNativeAddr(addr int64) *MatOfInt4 {
	return NewMatOfInt42(addr)
}
func (rcvr *MatOfInt4) ToArray() []int32 {
	num := rcvr.CheckVector2(_channelsMatOfInt4, _depthMatOfInt4)
	if num < 0 {
		Throw(NewRuntimeException(fmt.Sprintf("%v%v", "Native Mat has unexpected type or size: ", rcvr)))
	}
	a := make([]int32, num*_channelsMatOfInt4)
	if num == 0 {
		return a
	}
	rcvr.GetI(0, 0, a)
	return a
}
