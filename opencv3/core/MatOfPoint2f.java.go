package core

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

const _channelsMatOfPoint2f = 2

var _depthMatOfPoint2f = CvTypeCV_32F

type MatOfPoint2f struct {
	*Mat
}

func NewMatOfPoint2f() (rcvr *MatOfPoint2f) {
	rcvr = &MatOfPoint2f{}
	rcvr.Mat = NewMat2()
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func NewMatOfPoint2f2(addr int64) (rcvr *MatOfPoint2f) {
	rcvr = &MatOfPoint2f{}
	rcvr.Mat = NewMat(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	if !rcvr.Empty() && rcvr.CheckVector2(_channelsMatOfPoint2f, _depthMatOfPoint2f) < 0 {
		Throw(NewIllegalArgumentException("Incompatible Mat"))
	}
	return
}
func NewMatOfPoint2f3(m *Mat) (rcvr *MatOfPoint2f) {
	rcvr = &MatOfPoint2f{}
	rcvr.Mat = NewMat8(m, RangeAll())
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	if !rcvr.Empty() && rcvr.CheckVector2(_channelsMatOfPoint2f, _depthMatOfPoint2f) < 0 {
		Throw(NewIllegalArgumentException("Incompatible Mat"))
	}
	return
}
func NewMatOfPoint2f4(a []*Point) (rcvr *MatOfPoint2f) {
	rcvr = &MatOfPoint2f{}
	rcvr.Mat = NewMat2()
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	rcvr.FromArray(a)
	return
}
func (rcvr *MatOfPoint2f) Alloc(elemNumber int) {
	if elemNumber > 0 {
		rcvr.Create(elemNumber, 1, CvTypeMakeType(_depthMatOfPoint2f, _channelsMatOfPoint2f))
	}
}
func (rcvr *MatOfPoint2f) FromArray(a []*Point) {
	if a == nil || len(a) == 0 {
		return
	}
	num := len(a)
	rcvr.Alloc(num)
	buff := make([]float32, num*_channelsMatOfPoint2f)
	for i := (0); i < num; i++ {
		p := a[i]
		buff[_channelsMatOfPoint2f*i+0] = float32(p.X)
		buff[_channelsMatOfPoint2f*i+1] = float32(p.Y)
	}
	rcvr.PutF(0, 0, buff)
}
func MatOfPoint2fFromNativeAddr(addr int64) *MatOfPoint2f {
	return NewMatOfPoint2f2(addr)
}
func (rcvr *MatOfPoint2f) ToArray() []*Point {
	num := rcvr.Total()
	ap := make([]*Point, num)
	if num == 0 {
		return ap
	}
	buff := make([]float32, num*_channelsMatOfPoint2f)
	rcvr.GetF(0, 0, buff)
	for i := int64(0); i < num; i++ {
		ap[i] = NewPoint(float64(buff[i*_channelsMatOfPoint2f]), float64(buff[i*_channelsMatOfPoint2f+1]))
	}
	return ap
}
