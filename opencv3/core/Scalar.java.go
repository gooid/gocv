package core

import "fmt"

type Scalar struct {
	Val []float64
}

func NewScalar(v0 float64, v1 float64, v2 float64, v3 float64) (rcvr *Scalar) {
	rcvr = &Scalar{}
	rcvr.Val = []float64{v0, v1, v2, v3}
	return
}
func NewScalar2(v0 float64, v1 float64, v2 float64) (rcvr *Scalar) {
	rcvr = &Scalar{}
	rcvr.Val = []float64{v0, v1, v2, 0}
	return
}
func NewScalar3(v0 float64, v1 float64) (rcvr *Scalar) {
	rcvr = &Scalar{}
	rcvr.Val = []float64{v0, v1, 0, 0}
	return
}
func NewScalar4(v0 float64) (rcvr *Scalar) {
	rcvr = &Scalar{}
	rcvr.Val = []float64{v0, 0, 0, 0}
	return
}
func NewScalar5(vals []float64) (rcvr *Scalar) {
	rcvr = &Scalar{}
	rcvr.Set(vals)
	return
}
func ScalarAll(v float64) *Scalar {
	return NewScalar(v, v, v, v)
}
func (rcvr *Scalar) Clone() *Scalar {
	return NewScalar5(rcvr.Val)
}
func (rcvr *Scalar) Conj() *Scalar {
	return NewScalar(rcvr.Val[0], -rcvr.Val[1], -rcvr.Val[2], -rcvr.Val[3])
}
func (rcvr *Scalar) Equals(obj interface{}) bool {
	if rcvr == obj {
		return true
	}
	it, ok := obj.(*Scalar)
	if !ok {
		return false
	}
	return rcvr.Val[0] == it.Val[0] && rcvr.Val[1] == it.Val[1] && rcvr.Val[2] == it.Val[2] && rcvr.Val[3] == it.Val[3]
}
func (rcvr *Scalar) IsReal() bool {
	return rcvr.Val[1] == 0 && rcvr.Val[2] == 0 && rcvr.Val[3] == 0
}
func (rcvr *Scalar) Mul(it *Scalar, scale float64) *Scalar {
	return NewScalar(rcvr.Val[0]*it.Val[0]*scale, rcvr.Val[1]*it.Val[1]*scale, rcvr.Val[2]*it.Val[2]*scale, rcvr.Val[3]*it.Val[3]*scale)
}
func (rcvr *Scalar) Mul2(it *Scalar) *Scalar {
	return rcvr.Mul(it, 1)
}
func (rcvr *Scalar) Set(vals []float64) {
	if vals != nil {
		rcvr.Val[0] = func() float64 {
			if len(vals) > 0 {
				return vals[0]
			} else {
				return 0
			}
		}()
		rcvr.Val[1] = func() float64 {
			if len(vals) > 1 {
				return vals[1]
			} else {
				return 0
			}
		}()
		rcvr.Val[2] = func() float64 {
			if len(vals) > 2 {
				return vals[2]
			} else {
				return 0
			}
		}()
		rcvr.Val[3] = func() float64 {
			if len(vals) > 3 {
				return vals[3]
			} else {
				return 0
			}
		}()
	} else {
		rcvr.Val[0], rcvr.Val[1], rcvr.Val[2], rcvr.Val[3] = 0, 0, 0, 0
	}
}
func (rcvr *Scalar) String() string {
	return fmt.Sprintf("%v%v%v%v%v%v%v%v%v", "[", rcvr.Val[0], ", ", rcvr.Val[1], ", ", rcvr.Val[2], ", ", rcvr.Val[3], "]")
}
