package ml

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

type NormalBayesClassifier struct {
	*StatModel
}

func NewNormalBayesClassifier(addr int64) (rcvr *NormalBayesClassifier) {
	rcvr = &NormalBayesClassifier{}
	rcvr.StatModel = NewStatModel(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func NormalBayesClassifierCreate() *NormalBayesClassifier {
	retVal := NewNormalBayesClassifier(NormalBayesClassifierNative_create_0())
	return retVal
}
func (rcvr *NormalBayesClassifier) finalize() {
	NormalBayesClassifierNative_delete(rcvr.GetNativeObjAddr())
}
func NormalBayesClassifierLoad(filepath string, nodeName string) *NormalBayesClassifier {
	retVal := NewNormalBayesClassifier(NormalBayesClassifierNative_load_0(filepath, nodeName))
	return retVal
}
func NormalBayesClassifierLoad2(filepath string) *NormalBayesClassifier {
	retVal := NewNormalBayesClassifier(NormalBayesClassifierNative_load_1(filepath))
	return retVal
}
func (rcvr *NormalBayesClassifier) PredictProb(inputs *Mat, outputs *Mat, outputProbs *Mat, flags int) float32 {
	retVal := NormalBayesClassifierNative_predictProb_0(rcvr.GetNativeObjAddr(), inputs.GetNativeObjAddr(), outputs.GetNativeObjAddr(), outputProbs.GetNativeObjAddr(), flags)
	return retVal
}
func (rcvr *NormalBayesClassifier) PredictProb2(inputs *Mat, outputs *Mat, outputProbs *Mat) float32 {
	retVal := NormalBayesClassifierNative_predictProb_1(rcvr.GetNativeObjAddr(), inputs.GetNativeObjAddr(), outputs.GetNativeObjAddr(), outputProbs.GetNativeObjAddr())
	return retVal
}
