package core

import "fmt"

const TermCriteriaCOUNT = 1
const TermCriteriaEPS = 2

var TermCriteriaMAX_ITER = TermCriteriaCOUNT

type TermCriteria struct {
	Type     int
	MaxCount int
	Epsilon  float64
}

func NewTermCriteria(rtype int, maxCount int, epsilon float64) (rcvr *TermCriteria) {
	rcvr = &TermCriteria{}
	rcvr.Type = rtype
	rcvr.MaxCount = maxCount
	rcvr.Epsilon = epsilon
	return
}
func NewTermCriteria2() (rcvr *TermCriteria) {
	rcvr = NewTermCriteria(0, 0, 0.0)
	return
}
func NewTermCriteria3(vals []float64) (rcvr *TermCriteria) {
	rcvr = &TermCriteria{}
	rcvr.Set(vals)
	return
}
func (rcvr *TermCriteria) Clone() *TermCriteria {
	return NewTermCriteria(rcvr.Type, rcvr.MaxCount, rcvr.Epsilon)
}
func (rcvr *TermCriteria) Equals(obj interface{}) bool {
	if rcvr == obj {
		return true
	}
	it, ok := obj.(*TermCriteria)
	if !ok {
		return false
	}
	return rcvr.Type == it.Type && rcvr.MaxCount == it.MaxCount && rcvr.Epsilon == it.Epsilon
}
func (rcvr *TermCriteria) Set(vals []float64) {
	if vals != nil {
		rcvr.Type = func() int {
			if len(vals) > 0 {
				return int(vals[0])
			} else {
				return 0
			}
		}()
		rcvr.MaxCount = func() int {
			if len(vals) > 1 {
				return int(vals[1])
			} else {
				return 0
			}
		}()
		rcvr.Epsilon = func() float64 {
			if len(vals) > 2 {
				return vals[2]
			} else {
				return 0
			}
		}()
	} else {
		rcvr.Type = 0
		rcvr.MaxCount = 0
		rcvr.Epsilon = 0
	}
}
func (rcvr *TermCriteria) String() string {
	return fmt.Sprintf("%v%v%v%v%v%v%v", "{ type: ", rcvr.Type, ", maxCount: ", rcvr.MaxCount, ", epsilon: ", rcvr.Epsilon, "}")
}
