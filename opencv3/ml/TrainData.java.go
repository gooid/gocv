package ml

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

type TrainData struct {
	nativeObj int64
}

func NewTrainData(addr int64) (rcvr *TrainData) {
	rcvr = &TrainData{}
	rcvr.nativeObj = addr
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func TrainDataCreate(samples *Mat, layout int, responses *Mat, varIdx *Mat, sampleIdx *Mat, sampleWeights *Mat, varType *Mat) *TrainData {
	retVal := NewTrainData(TrainDataNative_create_0(samples.GetNativeObjAddr(), layout, responses.GetNativeObjAddr(), varIdx.GetNativeObjAddr(), sampleIdx.GetNativeObjAddr(), sampleWeights.GetNativeObjAddr(), varType.GetNativeObjAddr()))
	return retVal
}
func TrainDataCreate2(samples *Mat, layout int, responses *Mat) *TrainData {
	retVal := NewTrainData(TrainDataNative_create_1(samples.GetNativeObjAddr(), layout, responses.GetNativeObjAddr()))
	return retVal
}
func (rcvr *TrainData) finalize() {
	TrainDataNative_delete(rcvr.nativeObj)
}
func (rcvr *TrainData) GetCatCount(vi int) int {
	retVal := TrainDataNative_getCatCount_0(rcvr.nativeObj, vi)
	return retVal
}
func (rcvr *TrainData) GetCatMap() *Mat {
	retVal := NewMat(TrainDataNative_getCatMap_0(rcvr.nativeObj))
	return retVal
}
func (rcvr *TrainData) GetCatOfs() *Mat {
	retVal := NewMat(TrainDataNative_getCatOfs_0(rcvr.nativeObj))
	return retVal
}
func (rcvr *TrainData) GetClassLabels() *Mat {
	retVal := NewMat(TrainDataNative_getClassLabels_0(rcvr.nativeObj))
	return retVal
}
func (rcvr *TrainData) GetDefaultSubstValues() *Mat {
	retVal := NewMat(TrainDataNative_getDefaultSubstValues_0(rcvr.nativeObj))
	return retVal
}
func (rcvr *TrainData) GetLayout() int {
	retVal := TrainDataNative_getLayout_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *TrainData) GetMissing() *Mat {
	retVal := NewMat(TrainDataNative_getMissing_0(rcvr.nativeObj))
	return retVal
}
func (rcvr *TrainData) GetNAllVars() int {
	retVal := TrainDataNative_getNAllVars_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *TrainData) GetNSamples() int {
	retVal := TrainDataNative_getNSamples_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *TrainData) GetNTestSamples() int {
	retVal := TrainDataNative_getNTestSamples_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *TrainData) GetNTrainSamples() int {
	retVal := TrainDataNative_getNTrainSamples_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *TrainData) GetNVars() int {
	retVal := TrainDataNative_getNVars_0(rcvr.nativeObj)
	return retVal
}

//func (rcvr *TrainData) GetNames(names []*Mat) {
//	TrainDataNative_getNames_0(rcvr.nativeObj, names)
//	return
//}
func (rcvr *TrainData) GetNativeObjAddr() int64 {
	return rcvr.nativeObj
}
func (rcvr *TrainData) GetNormCatResponses() *Mat {
	retVal := NewMat(TrainDataNative_getNormCatResponses_0(rcvr.nativeObj))
	return retVal
}
func (rcvr *TrainData) GetResponseType() int {
	retVal := TrainDataNative_getResponseType_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *TrainData) GetResponses() *Mat {
	retVal := NewMat(TrainDataNative_getResponses_0(rcvr.nativeObj))
	return retVal
}
func (rcvr *TrainData) GetSample(varIdx *Mat, sidx int, buf float32) {
	TrainDataNative_getSample_0(rcvr.nativeObj, varIdx.GetNativeObjAddr(), sidx, buf)
	return
}
func (rcvr *TrainData) GetSampleWeights() *Mat {
	retVal := NewMat(TrainDataNative_getSampleWeights_0(rcvr.nativeObj))
	return retVal
}
func (rcvr *TrainData) GetSamples() *Mat {
	retVal := NewMat(TrainDataNative_getSamples_0(rcvr.nativeObj))
	return retVal
}
func GetSubVector(vec *Mat, idx *Mat) *Mat {
	retVal := NewMat(TrainDataNative_getSubVector_0(vec.GetNativeObjAddr(), idx.GetNativeObjAddr()))
	return retVal
}
func (rcvr *TrainData) GetTestNormCatResponses() *Mat {
	retVal := NewMat(TrainDataNative_getTestNormCatResponses_0(rcvr.nativeObj))
	return retVal
}
func (rcvr *TrainData) GetTestResponses() *Mat {
	retVal := NewMat(TrainDataNative_getTestResponses_0(rcvr.nativeObj))
	return retVal
}
func (rcvr *TrainData) GetTestSampleIdx() *Mat {
	retVal := NewMat(TrainDataNative_getTestSampleIdx_0(rcvr.nativeObj))
	return retVal
}
func (rcvr *TrainData) GetTestSampleWeights() *Mat {
	retVal := NewMat(TrainDataNative_getTestSampleWeights_0(rcvr.nativeObj))
	return retVal
}
func (rcvr *TrainData) GetTestSamples() *Mat {
	retVal := NewMat(TrainDataNative_getTestSamples_0(rcvr.nativeObj))
	return retVal
}
func (rcvr *TrainData) GetTrainNormCatResponses() *Mat {
	retVal := NewMat(TrainDataNative_getTrainNormCatResponses_0(rcvr.nativeObj))
	return retVal
}
func (rcvr *TrainData) GetTrainResponses() *Mat {
	retVal := NewMat(TrainDataNative_getTrainResponses_0(rcvr.nativeObj))
	return retVal
}
func (rcvr *TrainData) GetTrainSampleIdx() *Mat {
	retVal := NewMat(TrainDataNative_getTrainSampleIdx_0(rcvr.nativeObj))
	return retVal
}
func (rcvr *TrainData) GetTrainSampleWeights() *Mat {
	retVal := NewMat(TrainDataNative_getTrainSampleWeights_0(rcvr.nativeObj))
	return retVal
}
func (rcvr *TrainData) GetTrainSamples(layout int, compressSamples bool, compressVars bool) *Mat {
	retVal := NewMat(TrainDataNative_getTrainSamples_0(rcvr.nativeObj, layout, compressSamples, compressVars))
	return retVal
}
func (rcvr *TrainData) GetTrainSamples2() *Mat {
	retVal := NewMat(TrainDataNative_getTrainSamples_1(rcvr.nativeObj))
	return retVal
}
func (rcvr *TrainData) GetValues(vi int, sidx *Mat, values float32) {
	TrainDataNative_getValues_0(rcvr.nativeObj, vi, sidx.GetNativeObjAddr(), values)
	return
}
func (rcvr *TrainData) GetVarIdx() *Mat {
	retVal := NewMat(TrainDataNative_getVarIdx_0(rcvr.nativeObj))
	return retVal
}
func (rcvr *TrainData) GetVarSymbolFlags() *Mat {
	retVal := NewMat(TrainDataNative_getVarSymbolFlags_0(rcvr.nativeObj))
	return retVal
}
func (rcvr *TrainData) GetVarType() *Mat {
	retVal := NewMat(TrainDataNative_getVarType_0(rcvr.nativeObj))
	return retVal
}
func (rcvr *TrainData) SetTrainTestSplit(count int, shuffle bool) {
	TrainDataNative_setTrainTestSplit_0(rcvr.nativeObj, count, shuffle)
	return
}
func (rcvr *TrainData) SetTrainTestSplit2(count int) {
	TrainDataNative_setTrainTestSplit_1(rcvr.nativeObj, count)
	return
}
func (rcvr *TrainData) SetTrainTestSplitRatio(ratio float64, shuffle bool) {
	TrainDataNative_setTrainTestSplitRatio_0(rcvr.nativeObj, ratio, shuffle)
	return
}
func (rcvr *TrainData) SetTrainTestSplitRatio2(ratio float64) {
	TrainDataNative_setTrainTestSplitRatio_1(rcvr.nativeObj, ratio)
	return
}
func (rcvr *TrainData) ShuffleTrainTest() {
	TrainDataNative_shuffleTrainTest_0(rcvr.nativeObj)
	return
}
