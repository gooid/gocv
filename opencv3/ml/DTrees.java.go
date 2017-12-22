package ml

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

const DTreesPREDICT_AUTO = 0

var DTreesPREDICT_SUM = 1 << 8
var DTreesPREDICT_MAX_VOTE = 2 << 8
var DTreesPREDICT_MASK = 3 << 8

type DTrees struct {
	*StatModel
}

func NewDTrees(addr int64) (rcvr *DTrees) {
	rcvr = &DTrees{}
	rcvr.StatModel = NewStatModel(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func DTreesCreate() *DTrees {
	retVal := NewDTrees(DTreesNative_create_0())
	return retVal
}
func (rcvr *DTrees) finalize() {
	DTreesNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *DTrees) GetCVFolds() int {
	retVal := DTreesNative_getCVFolds_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *DTrees) GetMaxCategories() int {
	retVal := DTreesNative_getMaxCategories_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *DTrees) GetMaxDepth() int {
	retVal := DTreesNative_getMaxDepth_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *DTrees) GetMinSampleCount() int {
	retVal := DTreesNative_getMinSampleCount_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *DTrees) GetPriors() *Mat {
	retVal := NewMat(DTreesNative_getPriors_0(rcvr.GetNativeObjAddr()))
	return retVal
}
func (rcvr *DTrees) GetRegressionAccuracy() float32 {
	retVal := DTreesNative_getRegressionAccuracy_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *DTrees) GetTruncatePrunedTree() bool {
	retVal := DTreesNative_getTruncatePrunedTree_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *DTrees) GetUse1SERule() bool {
	retVal := DTreesNative_getUse1SERule_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *DTrees) GetUseSurrogates() bool {
	retVal := DTreesNative_getUseSurrogates_0(rcvr.GetNativeObjAddr())
	return retVal
}
func DTreesLoad(filepath string, nodeName string) *DTrees {
	retVal := NewDTrees(DTreesNative_load_0(filepath, nodeName))
	return retVal
}
func DTreesLoad2(filepath string) *DTrees {
	retVal := NewDTrees(DTreesNative_load_1(filepath))
	return retVal
}
func (rcvr *DTrees) SetCVFolds(val int) {
	DTreesNative_setCVFolds_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *DTrees) SetMaxCategories(val int) {
	DTreesNative_setMaxCategories_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *DTrees) SetMaxDepth(val int) {
	DTreesNative_setMaxDepth_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *DTrees) SetMinSampleCount(val int) {
	DTreesNative_setMinSampleCount_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *DTrees) SetPriors(val *Mat) {
	DTreesNative_setPriors_0(rcvr.GetNativeObjAddr(), val.GetNativeObjAddr())
	return
}
func (rcvr *DTrees) SetRegressionAccuracy(val float32) {
	DTreesNative_setRegressionAccuracy_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *DTrees) SetTruncatePrunedTree(val bool) {
	DTreesNative_setTruncatePrunedTree_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *DTrees) SetUse1SERule(val bool) {
	DTreesNative_setUse1SERule_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *DTrees) SetUseSurrogates(val bool) {
	DTreesNative_setUseSurrogates_0(rcvr.GetNativeObjAddr(), val)
	return
}
