package core

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

const _channelsMatOfRect2d = 4

var _depthMatOfRect2d = CvTypeCV_64F

type MatOfRect2d struct {
	*Mat
}

func NewMatOfRect2d() (rcvr *MatOfRect2d) {
	rcvr = &MatOfRect2d{}
	rcvr.Mat = NewMat2()
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func NewMatOfRect2d2(addr int64) (rcvr *MatOfRect2d) {
	rcvr = &MatOfRect2d{}
	rcvr.Mat = NewMat(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	if !rcvr.Empty() && rcvr.CheckVector2(_channelsMatOfRect2d, _depthMatOfRect2d) < 0 {
		Throw(NewIllegalArgumentException("Incompatible Mat"))
	}
	return
}
func NewMatOfRect2d3(m *Mat) (rcvr *MatOfRect2d) {
	rcvr = &MatOfRect2d{}
	rcvr.Mat = NewMat8(m, RangeAll())
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	if !rcvr.Empty() && rcvr.CheckVector2(_channelsMatOfRect2d, _depthMatOfRect2d) < 0 {
		Throw(NewIllegalArgumentException("Incompatible Mat"))
	}
	return
}
func (rcvr *MatOfRect2d) Alloc(elemNumber int) {
	if elemNumber > 0 {
		rcvr.Create(elemNumber, 1, CvTypeMakeType(_depthMatOfRect2d, _channelsMatOfRect2d))
	}
}
func (rcvr *MatOfRect2d) FromArray(a []*Rect2d) {
	if a == nil || len(a) == 0 {
		return
	}
	num := len(a)
	rcvr.Alloc(num)
	buff := make([]float64, num*_channelsMatOfRect2d)
	for i := 0; i < num; i++ {
		r := a[i]
		buff[_channelsMatOfRect2d*i+0] = r.X
		buff[_channelsMatOfRect2d*i+1] = r.Y
		buff[_channelsMatOfRect2d*i+2] = r.Width
		buff[_channelsMatOfRect2d*i+3] = r.Height
	}
	rcvr.PutD(0, 0, buff)
}
func MatOfRect2dFromNativeAddr(addr int64) *MatOfRect2d {
	return NewMatOfRect2d2(addr)
}
func (rcvr *MatOfRect2d) ToArray() []*Rect2d {
	num := rcvr.Total()
	a := make([]*Rect2d, num)
	if num == 0 {
		return a
	}
	buff := make([]float64, num*_channelsMatOfRect2d)
	rcvr.GetD(0, 0, buff)
	for i := int64(0); i < num; i++ {
		a[i] = NewRect2d(buff[i*_channelsMatOfRect2d], buff[i*_channelsMatOfRect2d+1], buff[i*_channelsMatOfRect2d+2], buff[i*_channelsMatOfRect2d+3])
	}
	return a
}
