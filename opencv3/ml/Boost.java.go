package ml

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

const BoostDISCRETE = 0
const BoostREAL = 1
const BoostLOGIT = 2
const BoostGENTLE = 3

type Boost struct {
	*DTrees
}

func NewBoost(addr int64) (rcvr *Boost) {
	rcvr = &Boost{}
	rcvr.DTrees = NewDTrees(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func BoostCreate() *Boost {
	retVal := NewBoost(BoostNative_create_0())
	return retVal
}
func (rcvr *Boost) finalize() {
	BoostNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *Boost) GetBoostType() int {
	retVal := BoostNative_getBoostType_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *Boost) GetWeakCount() int {
	retVal := BoostNative_getWeakCount_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *Boost) GetWeightTrimRate() float64 {
	retVal := BoostNative_getWeightTrimRate_0(rcvr.GetNativeObjAddr())
	return retVal
}
func BoostLoad(filepath string, nodeName string) *Boost {
	retVal := NewBoost(BoostNative_load_0(filepath, nodeName))
	return retVal
}
func BoostLoad2(filepath string) *Boost {
	retVal := NewBoost(BoostNative_load_1(filepath))
	return retVal
}
func (rcvr *Boost) SetBoostType(val int) {
	BoostNative_setBoostType_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *Boost) SetWeakCount(val int) {
	BoostNative_setWeakCount_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *Boost) SetWeightTrimRate(val float64) {
	BoostNative_setWeightTrimRate_0(rcvr.GetNativeObjAddr(), val)
	return
}
