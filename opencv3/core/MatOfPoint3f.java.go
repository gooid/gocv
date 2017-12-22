package core

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

const _channelsMatOfPoint3f = 3

var _depthMatOfPoint3f = CvTypeCV_32F

type MatOfPoint3f struct {
	*Mat
}

func NewMatOfPoint3f() (rcvr *MatOfPoint3f) {
	rcvr = &MatOfPoint3f{}
	rcvr.Mat = NewMat2()
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func NewMatOfPoint3f2(addr int64) (rcvr *MatOfPoint3f) {
	rcvr = &MatOfPoint3f{}
	rcvr.Mat = NewMat(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	if !rcvr.Empty() && rcvr.CheckVector2(_channelsMatOfPoint3f, _depthMatOfPoint3f) < 0 {
		Throw(NewIllegalArgumentException("Incompatible Mat"))
	}
	return
}
func NewMatOfPoint3f3(m *Mat) (rcvr *MatOfPoint3f) {
	rcvr = &MatOfPoint3f{}
	rcvr.Mat = NewMat8(m, RangeAll())
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	if !rcvr.Empty() && rcvr.CheckVector2(_channelsMatOfPoint3f, _depthMatOfPoint3f) < 0 {
		Throw(NewIllegalArgumentException("Incompatible Mat"))
	}
	return
}
func (rcvr *MatOfPoint3f) Alloc(elemNumber int) {
	if elemNumber > 0 {
		rcvr.Create(elemNumber, 1, CvTypeMakeType(_depthMatOfPoint3f, _channelsMatOfPoint3f))
	}
}
func (rcvr *MatOfPoint3f) FromArray(a []*Point3) {
	if a == nil || len(a) == 0 {
		return
	}
	num := len(a)
	rcvr.Alloc(num)
	buff := make([]float32, num*_channelsMatOfPoint3f)
	for i := (0); i < num; i++ {
		p := a[i]
		buff[_channelsMatOfPoint3f*i+0] = float32(p.X)
		buff[_channelsMatOfPoint3f*i+1] = float32(p.Y)
		buff[_channelsMatOfPoint3f*i+2] = float32(p.Z)
	}
	rcvr.PutF(0, 0, buff)
}
func MatOfPoint3fFromNativeAddr(addr int64) *MatOfPoint3f {
	return NewMatOfPoint3f2(addr)
}
func (rcvr *MatOfPoint3f) ToArray() []*Point3 {
	num := rcvr.Total()
	ap := make([]*Point3, num)
	if num == 0 {
		return ap
	}
	buff := make([]float32, num*_channelsMatOfPoint3f)
	rcvr.GetF(0, 0, buff)
	for i := int64(0); i < num; i++ {
		ap[i] = NewPoint31(float64(buff[i*_channelsMatOfPoint3f]), float64(buff[i*_channelsMatOfPoint3f+1]), float64(buff[i*_channelsMatOfPoint3f+2]))
	}
	return ap
}
