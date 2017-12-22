package core

import (
	"fmt"
)

type KeyPoint struct {
	Pt       Point
	Size     float32
	Angle    float32
	Response float32
	Octave   int
	Class_id int
}

func NewKeyPoint(x float32, y float32, _size float32, _angle float32, _response float32, _octave int, _class_id int) (rcvr *KeyPoint) {
	rcvr = &KeyPoint{}
	rcvr.Pt = *NewPoint(float64(x), float64(y))
	rcvr.Size = _size
	rcvr.Angle = _angle
	rcvr.Response = _response
	rcvr.Octave = _octave
	rcvr.Class_id = _class_id
	return
}
func NewKeyPoint2() (rcvr *KeyPoint) {
	rcvr = NewKeyPoint(0, 0, 0, -1, 0, 0, -1)
	return
}
func NewKeyPoint3(x float32, y float32, _size float32, _angle float32, _response float32, _octave int) (rcvr *KeyPoint) {
	rcvr = NewKeyPoint(x, y, _size, _angle, _response, _octave, -1)
	return
}
func NewKeyPoint4(x float32, y float32, _size float32, _angle float32, _response float32) (rcvr *KeyPoint) {
	rcvr = NewKeyPoint(x, y, _size, _angle, _response, 0, -1)
	return
}
func NewKeyPoint5(x float32, y float32, _size float32, _angle float32) (rcvr *KeyPoint) {
	rcvr = NewKeyPoint(x, y, _size, _angle, 0, 0, -1)
	return
}
func NewKeyPoint6(x float32, y float32, _size float32) (rcvr *KeyPoint) {
	rcvr = NewKeyPoint(x, y, _size, -1, 0, 0, -1)
	return
}
func (rcvr *KeyPoint) String() string {
	return fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v", "KeyPoint [pt=", rcvr.Pt, ", size=", rcvr.Size, ", angle=", rcvr.Angle, ", response=", rcvr.Response, ", octave=", rcvr.Octave, ", class_id=", rcvr.Class_id, "]")
}
