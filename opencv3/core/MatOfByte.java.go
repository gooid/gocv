package core

import (
	"fmt"
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

const _channelsMatOfByte = 1

var _depthMatOfByte = CvTypeCV_8U

type MatOfByte struct {
	*Mat
}

func NewMatOfByte() (rcvr *MatOfByte) {
	rcvr = &MatOfByte{}
	rcvr.Mat = NewMat2()
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func NewMatOfByte2(addr int64) (rcvr *MatOfByte) {
	rcvr = &MatOfByte{}
	rcvr.Mat = NewMat(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	if !rcvr.Empty() && rcvr.CheckVector2(_channelsMatOfByte, _depthMatOfByte) < 0 {
		Throw(NewIllegalArgumentException("Incompatible Mat"))
	}
	return
}
func NewMatOfByte3(m *Mat) (rcvr *MatOfByte) {
	rcvr = &MatOfByte{}
	rcvr.Mat = NewMat8(m, RangeAll())
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	if !rcvr.Empty() && rcvr.CheckVector2(_channelsMatOfByte, _depthMatOfByte) < 0 {
		Throw(NewIllegalArgumentException("Incompatible Mat"))
	}
	return
}
func NewMatOfByte4(a ...byte) (rcvr *MatOfByte) {
	rcvr = &MatOfByte{}
	rcvr.Mat = NewMat2()
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	rcvr.FromArray(a)
	return
}
func (rcvr *MatOfByte) Alloc(elemNumber int) {
	if elemNumber > 0 {
		rcvr.Create(elemNumber, 1, CvTypeMakeType(_depthMatOfByte, _channelsMatOfByte))
	}
}
func (rcvr *MatOfByte) FromArray(a []byte) {
	if a == nil || len(a) == 0 {
		return
	}
	num := len(a) / _channelsMatOfByte
	rcvr.Alloc(num)
	rcvr.PutB(0, 0, a)
}
func MatOfByteFromNativeAddr(addr int64) *MatOfByte {
	return NewMatOfByte2(addr)
}
func (rcvr *MatOfByte) ToArray() []byte {
	num := rcvr.CheckVector2(_channelsMatOfByte, _depthMatOfByte)
	if num < 0 {
		Throw(NewRuntimeException(fmt.Sprintf("%v%v", "Native Mat has unexpected type or size: ", rcvr)))
	}
	a := make([]byte, num*_channelsMatOfByte)
	if num == 0 {
		return a
	}
	rcvr.GetB(0, 0, a)
	return a
}
