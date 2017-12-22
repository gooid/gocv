package ml

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

type RTrees struct {
	*DTrees
}

func NewRTrees(addr int64) (rcvr *RTrees) {
	rcvr = &RTrees{}
	rcvr.DTrees = NewDTrees(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func RTreesCreate() *RTrees {
	retVal := NewRTrees(RTreesNative_create_0())
	return retVal
}
func (rcvr *RTrees) finalize() {
	RTreesNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *RTrees) GetActiveVarCount() int {
	retVal := RTreesNative_getActiveVarCount_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *RTrees) GetCalculateVarImportance() bool {
	retVal := RTreesNative_getCalculateVarImportance_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *RTrees) GetTermCriteria() *TermCriteria {
	retVal := NewTermCriteria3(RTreesNative_getTermCriteria_0(rcvr.GetNativeObjAddr()))
	return retVal
}
func (rcvr *RTrees) GetVarImportance() *Mat {
	retVal := NewMat(RTreesNative_getVarImportance_0(rcvr.GetNativeObjAddr()))
	return retVal
}
func (rcvr *RTrees) GetVotes(samples *Mat, results *Mat, flags int) {
	RTreesNative_getVotes_0(rcvr.GetNativeObjAddr(), samples.GetNativeObjAddr(), results.GetNativeObjAddr(), flags)
	return
}
func RTreesLoad(filepath string, nodeName string) *RTrees {
	retVal := NewRTrees(RTreesNative_load_0(filepath, nodeName))
	return retVal
}
func RTreesLoad2(filepath string) *RTrees {
	retVal := NewRTrees(RTreesNative_load_1(filepath))
	return retVal
}
func (rcvr *RTrees) SetActiveVarCount(val int) {
	RTreesNative_setActiveVarCount_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *RTrees) SetCalculateVarImportance(val bool) {
	RTreesNative_setCalculateVarImportance_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *RTrees) SetTermCriteria(val *TermCriteria) {
	RTreesNative_setTermCriteria_0(rcvr.GetNativeObjAddr(), val.Type, val.MaxCount, val.Epsilon)
	return
}
