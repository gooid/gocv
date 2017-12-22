package core

import "fmt"

type Rect2d struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func NewRect2d(x float64, y float64, width float64, height float64) (rcvr *Rect2d) {
	rcvr = &Rect2d{}
	rcvr.X = x
	rcvr.Y = y
	rcvr.Width = width
	rcvr.Height = height
	return
}
func NewRect2d2() (rcvr *Rect2d) {
	rcvr = NewRect2d(0, 0, 0, 0)
	return
}
func NewRect2d3(p1 *Point, p2 *Point) (rcvr *Rect2d) {
	rcvr = &Rect2d{}
	rcvr.X = func() float64 {
		if p1.X < p2.X {
			return p1.X
		} else {
			return p2.X
		}
	}()
	rcvr.Y = func() float64 {
		if p1.Y < p2.Y {
			return p1.Y
		} else {
			return p2.Y
		}
	}()
	rcvr.Width = func() float64 {
		if p1.X > p2.X {
			return p1.X
		} else {
			return p2.X
		}
	}() - rcvr.X
	rcvr.Height = func() float64 {
		if p1.Y > p2.Y {
			return p1.Y
		} else {
			return p2.Y
		}
	}() - rcvr.Y
	return
}
func NewRect2d4(p *Point, s *Size) (rcvr *Rect2d) {
	rcvr = NewRect2d(p.X, p.Y, s.Width, s.Height)
	return
}
func NewRect2d5(vals []float64) (rcvr *Rect2d) {
	rcvr = &Rect2d{}
	rcvr.Set(vals)
	return
}
func (rcvr *Rect2d) Area() float64 {
	return rcvr.Width * rcvr.Height
}
func (rcvr *Rect2d) Br() *Point {
	return NewPoint(rcvr.X+rcvr.Width, rcvr.Y+rcvr.Height)
}
func (rcvr *Rect2d) Clone() *Rect2d {
	return NewRect2d(rcvr.X, rcvr.Y, rcvr.Width, rcvr.Height)
}
func (rcvr *Rect2d) Contains(p *Point) bool {
	return rcvr.X <= p.X && p.X < rcvr.X+rcvr.Width && rcvr.Y <= p.Y && p.Y < rcvr.Y+rcvr.Height
}
func (rcvr *Rect2d) Empty() bool {
	return rcvr.Width <= 0 || rcvr.Height <= 0
}
func (rcvr *Rect2d) Equals(obj interface{}) bool {
	if rcvr == obj {
		return true
	}
	it, ok := obj.(*Rect2d)
	if !ok {
		return false
	}
	return rcvr.X == it.X && rcvr.Y == it.Y && rcvr.Width == it.Width && rcvr.Height == it.Height
}
func (rcvr *Rect2d) Set(vals []float64) {
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
		rcvr.Width = func() float64 {
			if len(vals) > 2 {
				return vals[2]
			} else {
				return 0
			}
		}()
		rcvr.Height = func() float64 {
			if len(vals) > 3 {
				return vals[3]
			} else {
				return 0
			}
		}()
	} else {
		rcvr.X = 0
		rcvr.Y = 0
		rcvr.Width = 0
		rcvr.Height = 0
	}
}
func (rcvr *Rect2d) Size() *Size {
	return NewSize(rcvr.Width, rcvr.Height)
}
func (rcvr *Rect2d) Tl() *Point {
	return NewPoint(rcvr.X, rcvr.Y)
}
func (rcvr *Rect2d) String() string {
	return fmt.Sprintf("%v%v%v%v%v%v%v%v%v", "{", rcvr.X, ", ", rcvr.Y, ", ", rcvr.Width, "x", rcvr.Height, "}")
}
