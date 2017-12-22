package video

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

type SparsePyrLKOpticalFlow struct {
	*SparseOpticalFlow
}

func NewSparsePyrLKOpticalFlow(addr int64) (rcvr *SparsePyrLKOpticalFlow) {
	rcvr = &SparsePyrLKOpticalFlow{}
	rcvr.SparseOpticalFlow = NewSparseOpticalFlow(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func SparsePyrLKOpticalFlowCreate(winSize *Size, maxLevel int, crit *TermCriteria, flags int, minEigThreshold float64) *SparsePyrLKOpticalFlow {
	retVal := NewSparsePyrLKOpticalFlow(SparsePyrLKOpticalFlowNative_create_0(winSize.Width, winSize.Height, maxLevel, crit.Type, crit.MaxCount, crit.Epsilon, flags, minEigThreshold))
	return retVal
}
func SparsePyrLKOpticalFlowCreate2() *SparsePyrLKOpticalFlow {
	retVal := NewSparsePyrLKOpticalFlow(SparsePyrLKOpticalFlowNative_create_1())
	return retVal
}
func (rcvr *SparsePyrLKOpticalFlow) finalize() {
	SparsePyrLKOpticalFlowNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *SparsePyrLKOpticalFlow) GetFlags() int {
	retVal := SparsePyrLKOpticalFlowNative_getFlags_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *SparsePyrLKOpticalFlow) GetMaxLevel() int {
	retVal := SparsePyrLKOpticalFlowNative_getMaxLevel_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *SparsePyrLKOpticalFlow) GetMinEigThreshold() float64 {
	retVal := SparsePyrLKOpticalFlowNative_getMinEigThreshold_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *SparsePyrLKOpticalFlow) GetTermCriteria() *TermCriteria {
	retVal := NewTermCriteria3(SparsePyrLKOpticalFlowNative_getTermCriteria_0(rcvr.GetNativeObjAddr()))
	return retVal
}
func (rcvr *SparsePyrLKOpticalFlow) GetWinSize() *Size {
	retVal := NewSize4(SparsePyrLKOpticalFlowNative_getWinSize_0(rcvr.GetNativeObjAddr()))
	return retVal
}
func (rcvr *SparsePyrLKOpticalFlow) SetFlags(flags int) {
	SparsePyrLKOpticalFlowNative_setFlags_0(rcvr.GetNativeObjAddr(), flags)
	return
}
func (rcvr *SparsePyrLKOpticalFlow) SetMaxLevel(maxLevel int) {
	SparsePyrLKOpticalFlowNative_setMaxLevel_0(rcvr.GetNativeObjAddr(), maxLevel)
	return
}
func (rcvr *SparsePyrLKOpticalFlow) SetMinEigThreshold(minEigThreshold float64) {
	SparsePyrLKOpticalFlowNative_setMinEigThreshold_0(rcvr.GetNativeObjAddr(), minEigThreshold)
	return
}
func (rcvr *SparsePyrLKOpticalFlow) SetTermCriteria(crit *TermCriteria) {
	SparsePyrLKOpticalFlowNative_setTermCriteria_0(rcvr.GetNativeObjAddr(), crit.Type, crit.MaxCount, crit.Epsilon)
	return
}
func (rcvr *SparsePyrLKOpticalFlow) SetWinSize(winSize *Size) {
	SparsePyrLKOpticalFlowNative_setWinSize_0(rcvr.GetNativeObjAddr(), winSize.Width, winSize.Height)
	return
}
