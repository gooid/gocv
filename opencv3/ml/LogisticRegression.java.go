package ml

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

const LogisticRegressionREG_L1 = 0
const LogisticRegressionREG_L2 = 1
const LogisticRegressionBATCH = 0
const LogisticRegressionMINI_BATCH = 1

var LogisticRegressionREG_DISABLE = -1

type LogisticRegression struct {
	*StatModel
}

func NewLogisticRegression(addr int64) (rcvr *LogisticRegression) {
	rcvr = &LogisticRegression{}
	rcvr.StatModel = NewStatModel(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func LogisticRegressionCreate() *LogisticRegression {
	retVal := NewLogisticRegression(LogisticRegressionNative_create_0())
	return retVal
}
func (rcvr *LogisticRegression) finalize() {
	LogisticRegressionNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *LogisticRegression) GetIterations() int {
	retVal := LogisticRegressionNative_getIterations_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *LogisticRegression) GetLearningRate() float64 {
	retVal := LogisticRegressionNative_getLearningRate_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *LogisticRegression) GetMiniBatchSize() int {
	retVal := LogisticRegressionNative_getMiniBatchSize_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *LogisticRegression) GetRegularization() int {
	retVal := LogisticRegressionNative_getRegularization_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *LogisticRegression) GetTermCriteria() *TermCriteria {
	retVal := NewTermCriteria3(LogisticRegressionNative_getTermCriteria_0(rcvr.GetNativeObjAddr()))
	return retVal
}
func (rcvr *LogisticRegression) GetTrainMethod() int {
	retVal := LogisticRegressionNative_getTrainMethod_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *LogisticRegression) Get_learnt_thetas() *Mat {
	retVal := NewMat(LogisticRegressionNative_get_learnt_thetas_0(rcvr.GetNativeObjAddr()))
	return retVal
}
func LogisticRegressionLoad(filepath string, nodeName string) *LogisticRegression {
	retVal := NewLogisticRegression(LogisticRegressionNative_load_0(filepath, nodeName))
	return retVal
}
func LogisticRegressionLoad2(filepath string) *LogisticRegression {
	retVal := NewLogisticRegression(LogisticRegressionNative_load_1(filepath))
	return retVal
}
func (rcvr *LogisticRegression) Predict(samples *Mat, results *Mat, flags int) float32 {
	retVal := LogisticRegressionNative_predict_0(rcvr.GetNativeObjAddr(), samples.GetNativeObjAddr(), results.GetNativeObjAddr(), flags)
	return retVal
}
func (rcvr *LogisticRegression) Predict2(samples *Mat) float32 {
	retVal := LogisticRegressionNative_predict_1(rcvr.GetNativeObjAddr(), samples.GetNativeObjAddr())
	return retVal
}
func (rcvr *LogisticRegression) SetIterations(val int) {
	LogisticRegressionNative_setIterations_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *LogisticRegression) SetLearningRate(val float64) {
	LogisticRegressionNative_setLearningRate_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *LogisticRegression) SetMiniBatchSize(val int) {
	LogisticRegressionNative_setMiniBatchSize_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *LogisticRegression) SetRegularization(val int) {
	LogisticRegressionNative_setRegularization_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *LogisticRegression) SetTermCriteria(val *TermCriteria) {
	LogisticRegressionNative_setTermCriteria_0(rcvr.GetNativeObjAddr(), val.Type, val.MaxCount, val.Epsilon)
	return
}
func (rcvr *LogisticRegression) SetTrainMethod(val int) {
	LogisticRegressionNative_setTrainMethod_0(rcvr.GetNativeObjAddr(), val)
	return
}
