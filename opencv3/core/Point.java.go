package core

import "fmt"

type Point struct {
	X float64
	Y float64
}

func NewPoint(x float64, y float64) (rcvr *Point) {
	rcvr = &Point{}
	rcvr.X = x
	rcvr.Y = y
	return
}
func NewPoint2() (rcvr *Point) {
	rcvr = NewPoint(0, 0)
	return
}
func NewPoint3(vals []float64) (rcvr *Point) {
	rcvr = NewPoint2()
	rcvr.Set(vals)
	return
}
func (rcvr *Point) Clone() *Point {
	return NewPoint(rcvr.X, rcvr.Y)
}
func (rcvr *Point) Dot(p *Point) float64 {
	return rcvr.X*p.X + rcvr.Y*p.Y
}
func (rcvr *Point) Equals(obj interface{}) bool {
	if rcvr == obj {
		return true
	}
	it, ok := obj.(*Point)
	if !ok {
		return false
	}
	return rcvr.X == it.X && rcvr.Y == it.Y
}
func (rcvr *Point) Inside(r *Rect) bool {
	return r.Contains(rcvr)
}
func (rcvr *Point) Set(vals []float64) {
	if vals != nil {
		rcvr.X = func() float64 {
			if len(vals) > 0 {
				return vals[0]
			} else {
				return 0
			}
		}()
		rcvr.Y = func() float64 {
			if len(vals) > 1 {
				return vals[1]
			} else {
				return 0
			}
		}()
	} else {
		rcvr.X = 0
		rcvr.Y = 0
	}
}
func (rcvr *Point) String() string {
	return fmt.Sprintf("%v%v%v%v%v", "{", rcvr.X, ", ", rcvr.Y, "}")
}
