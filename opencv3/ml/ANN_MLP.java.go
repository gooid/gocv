package ml

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

const ANN_MLPBACKPROP = 0
const ANN_MLPRPROP = 1
const ANN_MLPIDENTITY = 0
const ANN_MLPSIGMOID_SYM = 1
const ANN_MLPGAUSSIAN = 2
const ANN_MLPUPDATE_WEIGHTS = 1
const ANN_MLPNO_INPUT_SCALE = 2
const ANN_MLPNO_OUTPUT_SCALE = 4

type ANN_MLP struct {
	*StatModel
}

func NewANN_MLP(addr int64) (rcvr *ANN_MLP) {
	rcvr = &ANN_MLP{}
	rcvr.StatModel = NewStatModel(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func ANN_MLPCreate() *ANN_MLP {
	retVal := NewANN_MLP(ANN_MLPNative_create_0())
	return retVal
}
func (rcvr *ANN_MLP) finalize() {
	ANN_MLPNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *ANN_MLP) GetBackpropMomentumScale() float64 {
	retVal := ANN_MLPNative_getBackpropMomentumScale_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *ANN_MLP) GetBackpropWeightScale() float64 {
	retVal := ANN_MLPNative_getBackpropWeightScale_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *ANN_MLP) GetLayerSizes() *Mat {
	retVal := NewMat(ANN_MLPNative_getLayerSizes_0(rcvr.GetNativeObjAddr()))
	return retVal
}
func (rcvr *ANN_MLP) GetRpropDW0() float64 {
	retVal := ANN_MLPNative_getRpropDW0_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *ANN_MLP) GetRpropDWMax() float64 {
	retVal := ANN_MLPNative_getRpropDWMax_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *ANN_MLP) GetRpropDWMin() float64 {
	retVal := ANN_MLPNative_getRpropDWMin_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *ANN_MLP) GetRpropDWMinus() float64 {
	retVal := ANN_MLPNative_getRpropDWMinus_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *ANN_MLP) GetRpropDWPlus() float64 {
	retVal := ANN_MLPNative_getRpropDWPlus_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *ANN_MLP) GetTermCriteria() *TermCriteria {
	retVal := NewTermCriteria3(ANN_MLPNative_getTermCriteria_0(rcvr.GetNativeObjAddr()))
	return retVal
}
func (rcvr *ANN_MLP) GetTrainMethod() int {
	retVal := ANN_MLPNative_getTrainMethod_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *ANN_MLP) GetWeights(layerIdx int) *Mat {
	retVal := NewMat(ANN_MLPNative_getWeights_0(rcvr.GetNativeObjAddr(), layerIdx))
	return retVal
}
func ANN_MLPLoad(filepath string) *ANN_MLP {
	retVal := NewANN_MLP(ANN_MLPNative_load_0(filepath))
	return retVal
}
func (rcvr *ANN_MLP) SetActivationFunction(rtype int, param1 float64, param2 float64) {
	ANN_MLPNative_setActivationFunction_0(rcvr.GetNativeObjAddr(), rtype, param1, param2)
	return
}
func (rcvr *ANN_MLP) SetActivationFunction2(rtype int) {
	ANN_MLPNative_setActivationFunction_1(rcvr.GetNativeObjAddr(), rtype)
	return
}
func (rcvr *ANN_MLP) SetBackpropMomentumScale(val float64) {
	ANN_MLPNative_setBackpropMomentumScale_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *ANN_MLP) SetBackpropWeightScale(val float64) {
	ANN_MLPNative_setBackpropWeightScale_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *ANN_MLP) SetLayerSizes(_layer_sizes *Mat) {
	ANN_MLPNative_setLayerSizes_0(rcvr.GetNativeObjAddr(), _layer_sizes.GetNativeObjAddr())
	return
}
func (rcvr *ANN_MLP) SetRpropDW0(val float64) {
	ANN_MLPNative_setRpropDW0_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *ANN_MLP) SetRpropDWMax(val float64) {
	ANN_MLPNative_setRpropDWMax_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *ANN_MLP) SetRpropDWMin(val float64) {
	ANN_MLPNative_setRpropDWMin_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *ANN_MLP) SetRpropDWMinus(val float64) {
	ANN_MLPNative_setRpropDWMinus_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *ANN_MLP) SetRpropDWPlus(val float64) {
	ANN_MLPNative_setRpropDWPlus_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *ANN_MLP) SetTermCriteria(val *TermCriteria) {
	ANN_MLPNative_setTermCriteria_0(rcvr.GetNativeObjAddr(), val.Type, val.MaxCount, val.Epsilon)
	return
}
func (rcvr *ANN_MLP) SetTrainMethod(method int, param1 float64, param2 float64) {
	ANN_MLPNative_setTrainMethod_0(rcvr.GetNativeObjAddr(), method, param1, param2)
	return
}
func (rcvr *ANN_MLP) SetTrainMethod2(method int) {
	ANN_MLPNative_setTrainMethod_1(rcvr.GetNativeObjAddr(), method)
	return
}
