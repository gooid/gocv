package core

import (
	"fmt"
)

type Point3 struct {
	X float64
	Y float64
	Z float64
}

func NewPoint31(x float64, y float64, z float64) (rcvr *Point3) {
	rcvr = &Point3{}
	rcvr.X = x
	rcvr.Y = y
	rcvr.Z = z
	return
}
func NewPoint32() (rcvr *Point3) {
	rcvr = NewPoint31(0, 0, 0)
	return
}
func NewPoint33(p *Point) (rcvr *Point3) {
	rcvr = &Point3{}
	rcvr.X = p.X
	rcvr.Y = p.Y
	rcvr.Z = 0
	return
}
func NewPoint34(vals []float64) (rcvr *Point3) {
	rcvr = NewPoint32()
	rcvr.Set(vals)
	return
}
func (rcvr *Point3) Clone() *Point3 {
	return NewPoint31(rcvr.X, rcvr.Y, rcvr.Z)
}
func (rcvr *Point3) Cross(p *Point3) *Point3 {
	return NewPoint31(rcvr.Y*p.Z-rcvr.Z*p.Y, rcvr.Z*p.X-rcvr.X*p.Z, rcvr.X*p.Y-rcvr.Y*p.X)
}
func (rcvr *Point3) Dot(p *Point3) float64 {
	return rcvr.X*p.X + rcvr.Y*p.Y + rcvr.Z*p.Z
}
func (rcvr *Point3) Equals(obj interface{}) bool {
	if rcvr == obj {
		return true
	}
	it, ok := obj.(*Point3)
	if !ok {
		return false
	}
	return rcvr.X == it.X && rcvr.Y == it.Y && rcvr.Z == it.Z
}
func (rcvr *Point3) Set(vals []float64) {
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
		rcvr.Z = func() float64 {
			if len(vals) > 2 {
				return vals[2]
			} else {
				return 0
			}
		}()
	} else {
		rcvr.X = 0
		rcvr.Y = 0
		rcvr.Z = 0
	}
}
func (rcvr *Point3) String() string {
	return fmt.Sprintf("%v%v%v%v%v%v%v", "{", rcvr.X, ", ", rcvr.Y, ", ", rcvr.Z, "}")
}
