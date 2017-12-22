package core

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

const _channelsMatOfRect = 4

var _depthMatOfRect = CvTypeCV_32S

type MatOfRect struct {
	*Mat
}

func NewMatOfRect() (rcvr *MatOfRect) {
	rcvr = &MatOfRect{}
	rcvr.Mat = NewMat2()
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func NewMatOfRect2(addr int64) (rcvr *MatOfRect) {
	rcvr = &MatOfRect{}
	rcvr.Mat = NewMat(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	if !rcvr.Empty() && rcvr.CheckVector2(_channelsMatOfRect, _depthMatOfRect) < 0 {
		Throw(NewIllegalArgumentException("Incompatible Mat"))
	}
	return
}
func NewMatOfRect3(m *Mat) (rcvr *MatOfRect) {
	rcvr = &MatOfRect{}
	rcvr.Mat = NewMat8(m, RangeAll())
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	if !rcvr.Empty() && rcvr.CheckVector2(_channelsMatOfRect, _depthMatOfRect) < 0 {
		Throw(NewIllegalArgumentException("Incompatible Mat"))
	}
	return
}
func (rcvr *MatOfRect) Alloc(elemNumber int) {
	if elemNumber > 0 {
		rcvr.Create(elemNumber, 1, CvTypeMakeType(_depthMatOfRect, _channelsMatOfRect))
	}
}
func (rcvr *MatOfRect) FromArray(a []*Rect) {
	if a == nil || len(a) == 0 {
		return
	}
	num := len(a)
	rcvr.Alloc(num)
	buff := make([]int32, num*_channelsMatOfRect)
	for i := 0; i < num; i++ {
		r := a[i]
		buff[_channelsMatOfRect*i+0] = int32(r.X)
		buff[_channelsMatOfRect*i+1] = int32(r.Y)
		buff[_channelsMatOfRect*i+2] = int32(r.Width)
		buff[_channelsMatOfRect*i+3] = int32(r.Height)
	}
	rcvr.PutI(0, 0, buff)
}
func MatOfRectFromNativeAddr(addr int64) *MatOfRect {
	return NewMatOfRect2(addr)
}
func (rcvr *MatOfRect) ToArray() []*Rect {
	num := rcvr.Total()
	a := make([]*Rect, num)
	if num == 0 {
		return a
	}
	buff := make([]int32, num*_channelsMatOfRect)
	rcvr.GetI(0, 0, buff)
	for i := int64(0); i < num; i++ {
		a[i] = NewRect(int(buff[i*_channelsMatOfRect]), int(buff[i*_channelsMatOfRect+1]), int(buff[i*_channelsMatOfRect+2]), int(buff[i*_channelsMatOfRect+3]))
	}
	return a
}
