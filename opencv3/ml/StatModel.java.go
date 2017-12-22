package ml

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

const StatModelUPDATE_MODEL = 1
const StatModelRAW_OUTPUT = 1
const StatModelCOMPRESSED_INPUT = 2
const StatModelPREPROCESSED_INPUT = 4

type StatModel struct {
	*Algorithm
}

func NewStatModel(addr int64) (rcvr *StatModel) {
	rcvr = &StatModel{}
	rcvr.Algorithm = NewAlgorithm(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func (rcvr *StatModel) CalcError(data *TrainData, test bool, resp *Mat) float32 {
	retVal := StatModelNative_calcError_0(rcvr.GetNativeObjAddr(), data.GetNativeObjAddr(), test, resp.GetNativeObjAddr())
	return retVal
}
func (rcvr *StatModel) Empty() bool {
	retVal := StatModelNative_empty_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *StatModel) finalize() {
	StatModelNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *StatModel) GetVarCount() int {
	retVal := StatModelNative_getVarCount_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *StatModel) IsClassifier() bool {
	retVal := StatModelNative_isClassifier_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *StatModel) IsTrained() bool {
	retVal := StatModelNative_isTrained_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *StatModel) Predict(samples *Mat, results *Mat, flags int) float32 {
	retVal := StatModelNative_predict_0(rcvr.GetNativeObjAddr(), samples.GetNativeObjAddr(), results.GetNativeObjAddr(), flags)
	return retVal
}
func (rcvr *StatModel) Predict2(samples *Mat) float32 {
	retVal := StatModelNative_predict_1(rcvr.GetNativeObjAddr(), samples.GetNativeObjAddr())
	return retVal
}
func (rcvr *StatModel) Train(samples *Mat, layout int, responses *Mat) bool {
	retVal := StatModelNative_train_0(rcvr.GetNativeObjAddr(), samples.GetNativeObjAddr(), layout, responses.GetNativeObjAddr())
	return retVal
}
func (rcvr *StatModel) Train2(trainData *TrainData, flags int) bool {
	retVal := StatModelNative_train_1(rcvr.GetNativeObjAddr(), trainData.GetNativeObjAddr(), flags)
	return retVal
}
func (rcvr *StatModel) Train3(trainData *TrainData) bool {
	retVal := StatModelNative_train_2(rcvr.GetNativeObjAddr(), trainData.GetNativeObjAddr())
	return retVal
}
