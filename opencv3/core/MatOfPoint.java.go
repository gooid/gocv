package core

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

const _channelsMatOfPoint = 2

var _depthMatOfPoint = CvTypeCV_32S

type MatOfPoint struct {
	*Mat
}

func NewMatOfPoint() (rcvr *MatOfPoint) {
	rcvr = &MatOfPoint{}
	rcvr.Mat = NewMat2()
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func NewMatOfPoint2(addr int64) (rcvr *MatOfPoint) {
	rcvr = &MatOfPoint{}
	rcvr.Mat = NewMat(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	if !rcvr.Empty() && rcvr.CheckVector2(_channelsMatOfPoint, _depthMatOfPoint) < 0 {
		Throw(NewIllegalArgumentException("Incompatible Mat"))
	}
	return
}
func NewMatOfPoint3(m *Mat) (rcvr *MatOfPoint) {
	rcvr = &MatOfPoint{}
	rcvr.Mat = NewMat8(m, RangeAll())
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	if !rcvr.Empty() && rcvr.CheckVector2(_channelsMatOfPoint, _depthMatOfPoint) < 0 {
		Throw(NewIllegalArgumentException("Incompatible Mat"))
	}
	return
}
func NewMatOfPoint4(a []*Point) (rcvr *MatOfPoint) {
	rcvr = &MatOfPoint{}
	rcvr.Mat = NewMat2()
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	rcvr.FromArray(a)
	return
}
func (rcvr *MatOfPoint) Alloc(elemNumber int) {
	if elemNumber > 0 {
		rcvr.Create(elemNumber, 1, CvTypeMakeType(_depthMatOfPoint, _channelsMatOfPoint))
	}
}
func (rcvr *MatOfPoint) FromArray(a []*Point) {
	if a == nil || len(a) == 0 {
		return
	}
	num := len(a)
	rcvr.Alloc(num)
	buff := make([]int32, num*_channelsMatOfPoint)
	for i := (0); i < num; i++ {
		p := a[i]
		buff[_channelsMatOfPoint*i+0] = int32(p.X)
		buff[_channelsMatOfPoint*i+1] = int32(p.Y)
	}
	rcvr.PutI(0, 0, buff)
}
func MatOfPointFromNativeAddr(addr int64) *MatOfPoint {
	return NewMatOfPoint2(addr)
}
func (rcvr *MatOfPoint) ToArray() []*Point {
	num := rcvr.Total()
	ap := make([]*Point, num)
	if num == 0 {
		return ap
	}
	buff := make([]int32, num*_channelsMatOfPoint)
	rcvr.GetI(0, 0, buff)
	for i := int64(0); i < num; i++ {
		ap[i] = NewPoint(float64(buff[i*_channelsMatOfPoint]), float64(buff[i*_channelsMatOfPoint+1]))
	}
	return ap
}
