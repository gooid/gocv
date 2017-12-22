package core

import (
	"fmt"
	"math"
)

type RotatedRect struct {
	Center Point
	Size   Size
	Angle  float64
}

func NewRotatedRect() (rcvr *RotatedRect) {
	rcvr = &RotatedRect{}
	rcvr.Center = *NewPoint2()
	rcvr.Size = *NewSize2()
	rcvr.Angle = 0
	return
}
func NewRotatedRect2(c *Point, s *Size, a float64) (rcvr *RotatedRect) {
	rcvr = &RotatedRect{}
	rcvr.Center = *c.Clone()
	rcvr.Size = *s.Clone()
	rcvr.Angle = a
	return
}
func NewRotatedRect3(vals []float64) (rcvr *RotatedRect) {
	rcvr = NewRotatedRect()
	rcvr.Set(vals)
	return
}
func (rcvr *RotatedRect) BoundingRect() *Rect {
	pt := make([]*Point, 4)
	rcvr.Points(pt)

	r := NewRect(int(math.Floor(math.Min(math.Min(math.Min(pt[0].X, pt[1].X), pt[2].X), pt[3].X))),
		int(math.Floor(math.Min(math.Min(math.Min(pt[0].Y, pt[1].Y), pt[2].Y), pt[3].Y))),
		int(math.Ceil(math.Max(math.Max(math.Max(pt[0].X, pt[1].X), pt[2].X), pt[3].X))),
		int(math.Ceil(math.Max(math.Max(math.Max(pt[0].Y, pt[1].Y), pt[2].Y), pt[3].Y))))

	r.Width -= r.X - 1
	r.Height -= r.Y - 1
	return r
}
func (rcvr *RotatedRect) Clone() *RotatedRect {
	return NewRotatedRect2(&rcvr.Center, &rcvr.Size, rcvr.Angle)
}
func (rcvr *RotatedRect) Equals(obj interface{}) bool {
	if rcvr == obj {
		return true
	}
	it, ok := obj.(*RotatedRect)
	if !ok {
		return false
	}
	return rcvr.Center.Equals(it.Center) && rcvr.Size.Equals(it.Size) && rcvr.Angle == it.Angle
}
func (rcvr *RotatedRect) Points(pt []*Point) {
	_angle := rcvr.Angle * math.Pi / 180.0
	b := math.Cos(_angle) * 0.5
	a := math.Sin(_angle) * 0.5
	pt[0] = NewPoint(rcvr.Center.X-a*rcvr.Size.Height-b*rcvr.Size.Width, rcvr.Center.Y+b*rcvr.Size.Height-a*rcvr.Size.Width)
	pt[1] = NewPoint(rcvr.Center.X+a*rcvr.Size.Height-b*rcvr.Size.Width, rcvr.Center.Y-b*rcvr.Size.Height-a*rcvr.Size.Width)
	pt[2] = NewPoint(2*rcvr.Center.X-pt[0].X, 2*rcvr.Center.Y-pt[0].Y)
	pt[3] = NewPoint(2*rcvr.Center.X-pt[0].X, 2*rcvr.Center.Y-pt[0].Y)
}
func (rcvr *RotatedRect) Set(vals []float64) {
	if vals != nil {
		rcvr.Center.X = func() float64 {
			if len(vals) > 0 {
				return vals[0]
			} else {
				return 0
			}
		}()
		rcvr.Center.Y = func() float64 {
			if len(vals) > 1 {
				return vals[1]
			} else {
				return 0
			}
		}()
		rcvr.Size.Width = func() float64 {
			if len(vals) > 2 {
				return vals[2]
			} else {
				return 0
			}
		}()
		rcvr.Size.Height = func() float64 {
			if len(vals) > 3 {
				return vals[3]
			} else {
				return 0
			}
		}()
		rcvr.Angle = func() float64 {
			if len(vals) > 4 {
				return vals[4]
			} else {
				return 0
			}
		}()
	} else {
		rcvr.Center.X = 0
		rcvr.Center.X = 0
		rcvr.Size.Width = 0
		rcvr.Size.Height = 0
		rcvr.Angle = 0
	}
}
func (rcvr *RotatedRect) String() string {
	return fmt.Sprintf("%v%v%v%v%v%v%v", "{ ", rcvr.Center, " ", rcvr.Size, " * ", rcvr.Angle, " }")
}
