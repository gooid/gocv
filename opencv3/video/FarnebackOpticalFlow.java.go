package video

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

type FarnebackOpticalFlow struct {
	*DenseOpticalFlow
}

func NewFarnebackOpticalFlow(addr int64) (rcvr *FarnebackOpticalFlow) {
	rcvr = &FarnebackOpticalFlow{}
	rcvr.DenseOpticalFlow = NewDenseOpticalFlow(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func FarnebackOpticalFlowCreate(numLevels int, pyrScale float64, fastPyramids bool, winSize int, numIters int, polyN int, polySigma float64, flags int) *FarnebackOpticalFlow {
	retVal := NewFarnebackOpticalFlow(FarnebackOpticalFlowNative_create_0(numLevels, pyrScale, fastPyramids, winSize, numIters, polyN, polySigma, flags))
	return retVal
}
func FarnebackOpticalFlowCreate2() *FarnebackOpticalFlow {
	retVal := NewFarnebackOpticalFlow(FarnebackOpticalFlowNative_create_1())
	return retVal
}
func (rcvr *FarnebackOpticalFlow) finalize() {
	FarnebackOpticalFlowNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *FarnebackOpticalFlow) GetFastPyramids() bool {
	retVal := FarnebackOpticalFlowNative_getFastPyramids_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *FarnebackOpticalFlow) GetFlags() int {
	retVal := FarnebackOpticalFlowNative_getFlags_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *FarnebackOpticalFlow) GetNumIters() int {
	retVal := FarnebackOpticalFlowNative_getNumIters_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *FarnebackOpticalFlow) GetNumLevels() int {
	retVal := FarnebackOpticalFlowNative_getNumLevels_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *FarnebackOpticalFlow) GetPolyN() int {
	retVal := FarnebackOpticalFlowNative_getPolyN_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *FarnebackOpticalFlow) GetPolySigma() float64 {
	retVal := FarnebackOpticalFlowNative_getPolySigma_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *FarnebackOpticalFlow) GetPyrScale() float64 {
	retVal := FarnebackOpticalFlowNative_getPyrScale_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *FarnebackOpticalFlow) GetWinSize() int {
	retVal := FarnebackOpticalFlowNative_getWinSize_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *FarnebackOpticalFlow) SetFastPyramids(fastPyramids bool) {
	FarnebackOpticalFlowNative_setFastPyramids_0(rcvr.GetNativeObjAddr(), fastPyramids)
	return
}
func (rcvr *FarnebackOpticalFlow) SetFlags(flags int) {
	FarnebackOpticalFlowNative_setFlags_0(rcvr.GetNativeObjAddr(), flags)
	return
}
func (rcvr *FarnebackOpticalFlow) SetNumIters(numIters int) {
	FarnebackOpticalFlowNative_setNumIters_0(rcvr.GetNativeObjAddr(), numIters)
	return
}
func (rcvr *FarnebackOpticalFlow) SetNumLevels(numLevels int) {
	FarnebackOpticalFlowNative_setNumLevels_0(rcvr.GetNativeObjAddr(), numLevels)
	return
}
func (rcvr *FarnebackOpticalFlow) SetPolyN(polyN int) {
	FarnebackOpticalFlowNative_setPolyN_0(rcvr.GetNativeObjAddr(), polyN)
	return
}
func (rcvr *FarnebackOpticalFlow) SetPolySigma(polySigma float64) {
	FarnebackOpticalFlowNative_setPolySigma_0(rcvr.GetNativeObjAddr(), polySigma)
	return
}
func (rcvr *FarnebackOpticalFlow) SetPyrScale(pyrScale float64) {
	FarnebackOpticalFlowNative_setPyrScale_0(rcvr.GetNativeObjAddr(), pyrScale)
	return
}
func (rcvr *FarnebackOpticalFlow) SetWinSize(winSize int) {
	FarnebackOpticalFlowNative_setWinSize_0(rcvr.GetNativeObjAddr(), winSize)
	return
}
