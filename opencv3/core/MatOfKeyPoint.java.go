package core

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

const _channelsMatOfKeyPoint = 7

var _depthMatOfKeyPoint = CvTypeCV_32F

type MatOfKeyPoint struct {
	*Mat
}

func NewMatOfKeyPoint() (rcvr *MatOfKeyPoint) {
	rcvr = &MatOfKeyPoint{}
	rcvr.Mat = NewMat2()
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func NewMatOfKeyPoint2(addr int64) (rcvr *MatOfKeyPoint) {
	rcvr = &MatOfKeyPoint{}
	rcvr.Mat = NewMat(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	if !rcvr.Empty() && rcvr.CheckVector2(_channelsMatOfKeyPoint, _depthMatOfKeyPoint) < 0 {
		Throw(NewIllegalArgumentException("Incompatible Mat"))
	}
	return
}
func NewMatOfKeyPoint3(m *Mat) (rcvr *MatOfKeyPoint) {
	rcvr = &MatOfKeyPoint{}
	rcvr.Mat = NewMat8(m, RangeAll())
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	if !rcvr.Empty() && rcvr.CheckVector2(_channelsMatOfKeyPoint, _depthMatOfKeyPoint) < 0 {
		Throw(NewIllegalArgumentException("Incompatible Mat"))
	}
	return
}
func NewMatOfKeyPoint4(a []*KeyPoint) (rcvr *MatOfKeyPoint) {
	rcvr = &MatOfKeyPoint{}
	rcvr.Mat = NewMat2()
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	rcvr.FromArray(a)
	return
}
func (rcvr *MatOfKeyPoint) Alloc(elemNumber int) {
	if elemNumber > 0 {
		rcvr.Create(elemNumber, 1, CvTypeMakeType(_depthMatOfKeyPoint, _channelsMatOfKeyPoint))
	}
}
func (rcvr *MatOfKeyPoint) FromArray(a []*KeyPoint) {
	if a == nil || len(a) == 0 {
		return
	}
	num := len(a)
	rcvr.Alloc(num)
	buff := make([]float32, num*_channelsMatOfKeyPoint)
	for i := (0); i < num; i++ {
		kp := a[i]
		buff[_channelsMatOfKeyPoint*i+0] = float32(kp.Pt.X)
		buff[_channelsMatOfKeyPoint*i+1] = float32(kp.Pt.Y)
		buff[_channelsMatOfKeyPoint*i+2] = float32(kp.Size)
		buff[_channelsMatOfKeyPoint*i+3] = float32(kp.Angle)
		buff[_channelsMatOfKeyPoint*i+4] = float32(kp.Response)
		buff[_channelsMatOfKeyPoint*i+5] = float32(kp.Octave)
		buff[_channelsMatOfKeyPoint*i+6] = float32(kp.Class_id)
	}
	rcvr.PutF(0, 0, buff)
}
func FromNativeAddr(addr int64) *MatOfKeyPoint {
	return NewMatOfKeyPoint2(addr)
}
func (rcvr *MatOfKeyPoint) ToArray() []*KeyPoint {
	num := rcvr.Total()
	a := make([]*KeyPoint, num)
	if num == 0 {
		return a
	}
	buff := make([]float32, num*_channelsMatOfKeyPoint)
	rcvr.GetF(0, 0, buff)
	for i := int64(0); i < num; i++ {
		a[i] = NewKeyPoint(buff[_channelsMatOfKeyPoint*i+0], buff[_channelsMatOfKeyPoint*i+1], buff[_channelsMatOfKeyPoint*i+2], buff[_channelsMatOfKeyPoint*i+3], buff[_channelsMatOfKeyPoint*i+4], int(buff[_channelsMatOfKeyPoint*i+5]), int(buff[_channelsMatOfKeyPoint*i+6]))
	}
	return a
}
