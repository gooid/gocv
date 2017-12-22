package core

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

const _channelsMatOfPoint3 = 3

var _depthMatOfPoint3 = CvTypeCV_32S

type MatOfPoint3 struct {
	*Mat
}

func NewMatOfPoint31() (rcvr *MatOfPoint3) {
	rcvr = &MatOfPoint3{}
	rcvr.Mat = NewMat2()
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func NewMatOfPoint32(addr int64) (rcvr *MatOfPoint3) {
	rcvr = &MatOfPoint3{}
	rcvr.Mat = NewMat(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	if !rcvr.Empty() && rcvr.CheckVector2(_channelsMatOfPoint3, _depthMatOfPoint3) < 0 {
		Throw(NewIllegalArgumentException("Incompatible Mat"))
	}
	return
}
func NewMatOfPoint33(m *Mat) (rcvr *MatOfPoint3) {
	rcvr = &MatOfPoint3{}
	rcvr.Mat = NewMat8(m, RangeAll())
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	if !rcvr.Empty() && rcvr.CheckVector2(_channelsMatOfPoint3, _depthMatOfPoint3) < 0 {
		Throw(NewIllegalArgumentException("Incompatible Mat"))
	}
	return
}
func (rcvr *MatOfPoint3) Alloc(elemNumber int) {
	if elemNumber > 0 {
		rcvr.Create(elemNumber, 1, CvTypeMakeType(_depthMatOfPoint3, _channelsMatOfPoint3))
	}
}
func (rcvr *MatOfPoint3) FromArray(a []*Point3) {
	if a == nil || len(a) == 0 {
		return
	}
	num := len(a)
	rcvr.Alloc(num)
	buff := make([]int32, num*_channelsMatOfPoint3)
	for i := (0); i < num; i++ {
		p := a[i]
		buff[_channelsMatOfPoint3*i+0] = int32(p.X)
		buff[_channelsMatOfPoint3*i+1] = int32(p.Y)
		buff[_channelsMatOfPoint3*i+2] = int32(p.Z)
	}
	rcvr.PutI(0, 0, buff)
}
func MatOfPoint3FromNativeAddr(addr int64) *MatOfPoint3 {
	return NewMatOfPoint32(addr)
}
func (rcvr *MatOfPoint3) ToArray() []*Point3 {
	num := rcvr.Total()
	ap := make([]*Point3, num)
	if num == 0 {
		return ap
	}
	buff := make([]int32, num*_channelsMatOfPoint3)
	rcvr.GetI(0, 0, buff)
	for i := int64(0); i < num; i++ {
		ap[i] = NewPoint31(float64(buff[i*_channelsMatOfPoint3]), float64(buff[i*_channelsMatOfPoint3+1]), float64(buff[i*_channelsMatOfPoint3+2]))
	}
	return ap
}
