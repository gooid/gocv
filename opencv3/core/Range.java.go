package core

import (
	"fmt"
	"math"
)

type Range struct {
	Start int
	End   int
}

func NewRange(s int, e int) (rcvr *Range) {
	rcvr = &Range{}
	rcvr.Start = s
	rcvr.End = e
	return
}
func NewRange2() (rcvr *Range) {
	rcvr = NewRange(0, 0)
	return
}
func NewRange3(vals []float64) (rcvr *Range) {
	rcvr = &Range{}
	rcvr.Set(vals)
	return
}
func RangeAll() *Range {
	return NewRange(math.MinInt32, math.MaxInt32)
}
func (rcvr *Range) Clone() *Range {
	return NewRange(rcvr.Start, rcvr.End)
}
func (rcvr *Range) Empty() bool {
	return rcvr.End <= rcvr.Start
}
func (rcvr *Range) Equals(obj interface{}) bool {
	if rcvr == obj {
		return true
	}
	it, ok := obj.(*Range)
	if !ok {
		return false
	}
	return rcvr.Start == it.Start && rcvr.End == it.End
}
func (rcvr *Range) Intersection(r1 *Range) *Range {
	s := r1.Start
	if s < rcvr.Start {
		s = rcvr.Start
	}
	e := r1.End
	if e > rcvr.End {
		e = rcvr.End
	}
	r := NewRange(s, e)
	if r.End < s {
		r.End = s
	}
	return r
}
func (rcvr *Range) Set(vals []float64) {
	if vals != nil {
		rcvr.Start = func() int {
			if len(vals) > 0 {
				return int(vals[0])
			} else {
				return 0
			}
		}()
		rcvr.End = func() int {
			if len(vals) > 1 {
				return int(vals[1])
			} else {
				return 0
			}
		}()
	} else {
		rcvr.Start = 0
		rcvr.End = 0
	}
}
func (rcvr *Range) Shift(delta int) *Range {
	return NewRange(rcvr.Start+delta, rcvr.End+delta)
}
func (rcvr *Range) Size() int {
	return func() int {
		if rcvr.Empty() {
			return 0
		} else {
			return rcvr.End - rcvr.Start
		}
	}()
}
func (rcvr *Range) String() string {
	return fmt.Sprintf("%v%v%v%v%v", "[", rcvr.Start, ", ", rcvr.End, ")")
}
