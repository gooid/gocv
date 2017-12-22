package core

import (
	"fmt"
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

const _channelsMatOfDMatch = 4

var _depthMatOfDMatch = CvTypeCV_32F

type MatOfDMatch struct {
	*Mat
}

func NewMatOfDMatch() (rcvr *MatOfDMatch) {
	rcvr = &MatOfDMatch{}
	rcvr.Mat = NewMat2()
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func NewMatOfDMatch2(addr int64) (rcvr *MatOfDMatch) {
	rcvr = &MatOfDMatch{}
	rcvr.Mat = NewMat(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	if !rcvr.Empty() && rcvr.CheckVector2(_channelsMatOfDMatch, _depthMatOfDMatch) < 0 {
		Throw(NewIllegalArgumentException(fmt.Sprintf("%v%v", "Incompatible Mat: ", rcvr)))
	}
	return
}
func NewMatOfDMatch3(m *Mat) (rcvr *MatOfDMatch) {
	rcvr = &MatOfDMatch{}
	rcvr.Mat = NewMat8(m, RangeAll())
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	if !rcvr.Empty() && rcvr.CheckVector2(_channelsMatOfDMatch, _depthMatOfDMatch) < 0 {
		Throw(NewIllegalArgumentException(fmt.Sprintf("%v%v", "Incompatible Mat: ", rcvr)))
	}
	return
}
func NewMatOfDMatch4(ap []*DMatch) (rcvr *MatOfDMatch) {
	rcvr = &MatOfDMatch{}
	rcvr.Mat = NewMat2()
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	rcvr.FromArray(ap)
	return
}
func (rcvr *MatOfDMatch) Alloc(elemNumber int) {
	if elemNumber > 0 {
		rcvr.Create(elemNumber, 1, CvTypeMakeType(_depthMatOfDMatch, _channelsMatOfDMatch))
	}
}
func (rcvr *MatOfDMatch) FromArray(a []*DMatch) {
	if a == nil || len(a) == 0 {
		return
	}
	num := len(a)
	rcvr.Alloc(num)
	buff := make([]float32, num*_channelsMatOfDMatch)
	for i := (0); i < num; i++ {
		m := a[i]
		buff[_channelsMatOfDMatch*i+0] = float32(m.QueryIdx)
		buff[_channelsMatOfDMatch*i+1] = float32(m.TrainIdx)
		buff[_channelsMatOfDMatch*i+2] = float32(m.ImgIdx)
		buff[_channelsMatOfDMatch*i+3] = float32(m.Distance)
	}
	rcvr.PutF(0, 0, buff)
}
func MatOfDMatchFromNativeAddr(addr int64) *MatOfDMatch {
	return NewMatOfDMatch2(addr)
}
func (rcvr *MatOfDMatch) ToArray() []*DMatch {
	num := rcvr.Total()
	a := make([]*DMatch, num)
	if num == 0 {
		return a
	}
	buff := make([]float32, num*_channelsMatOfDMatch)
	rcvr.GetF(0, 0, buff)
	for i := int64(0); i < num; i++ {
		a[i] = NewDMatch3(int(buff[_channelsMatOfDMatch*i+0]), int(buff[_channelsMatOfDMatch*i+1]), int(buff[_channelsMatOfDMatch*i+2]), buff[_channelsMatOfDMatch*i+3])
	}
	return a
}
