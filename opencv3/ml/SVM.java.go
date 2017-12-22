package ml

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

const SVMC_SVC = 100
const SVMNU_SVC = 101
const SVMONE_CLASS = 102
const SVMEPS_SVR = 103
const SVMNU_SVR = 104
const SVMLINEAR = 0
const SVMPOLY = 1
const SVMRBF = 2
const SVMSIGMOID = 3
const SVMCHI2 = 4
const SVMINTER = 5
const SVMC = 0
const SVMGAMMA = 1
const SVMP = 2
const SVMNU = 3
const SVMCOEF = 4
const SVMDEGREE = 5

var SVMCUSTOM = -1

type SVM struct {
	*StatModel
}

func NewSVM(addr int64) (rcvr *SVM) {
	rcvr = &SVM{}
	rcvr.StatModel = NewStatModel(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func SVMCreate() *SVM {
	retVal := NewSVM(SVMNative_create_0())
	return retVal
}
func (rcvr *SVM) finalize() {
	SVMNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *SVM) GetC() float64 {
	retVal := SVMNative_getC_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *SVM) GetClassWeights() *Mat {
	retVal := NewMat(SVMNative_getClassWeights_0(rcvr.GetNativeObjAddr()))
	return retVal
}
func (rcvr *SVM) GetCoef0() float64 {
	retVal := SVMNative_getCoef0_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *SVM) GetDecisionFunction(i int, alpha *Mat, svidx *Mat) float64 {
	retVal := SVMNative_getDecisionFunction_0(rcvr.GetNativeObjAddr(), i, alpha.GetNativeObjAddr(), svidx.GetNativeObjAddr())
	return retVal
}
func GetDefaultGridPtr(param_id int) *ParamGrid {
	retVal := NewParamGrid(SVMNative_getDefaultGridPtr_0(param_id))
	return retVal
}
func (rcvr *SVM) GetDegree() float64 {
	retVal := SVMNative_getDegree_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *SVM) GetGamma() float64 {
	retVal := SVMNative_getGamma_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *SVM) GetKernelType() int {
	retVal := SVMNative_getKernelType_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *SVM) GetNu() float64 {
	retVal := SVMNative_getNu_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *SVM) GetP() float64 {
	retVal := SVMNative_getP_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *SVM) GetSupportVectors() *Mat {
	retVal := NewMat(SVMNative_getSupportVectors_0(rcvr.GetNativeObjAddr()))
	return retVal
}
func (rcvr *SVM) GetTermCriteria() *TermCriteria {
	retVal := NewTermCriteria3(SVMNative_getTermCriteria_0(rcvr.GetNativeObjAddr()))
	return retVal
}
func (rcvr *SVM) GetType() int {
	retVal := SVMNative_getType_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *SVM) GetUncompressedSupportVectors() *Mat {
	retVal := NewMat(SVMNative_getUncompressedSupportVectors_0(rcvr.GetNativeObjAddr()))
	return retVal
}
func SVMLoad(filepath string) *SVM {
	retVal := NewSVM(SVMNative_load_0(filepath))
	return retVal
}
func (rcvr *SVM) SetC(val float64) {
	SVMNative_setC_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *SVM) SetClassWeights(val *Mat) {
	SVMNative_setClassWeights_0(rcvr.GetNativeObjAddr(), val.GetNativeObjAddr())
	return
}
func (rcvr *SVM) SetCoef0(val float64) {
	SVMNative_setCoef0_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *SVM) SetDegree(val float64) {
	SVMNative_setDegree_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *SVM) SetGamma(val float64) {
	SVMNative_setGamma_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *SVM) SetKernel(kernelType int) {
	SVMNative_setKernel_0(rcvr.GetNativeObjAddr(), kernelType)
	return
}
func (rcvr *SVM) SetNu(val float64) {
	SVMNative_setNu_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *SVM) SetP(val float64) {
	SVMNative_setP_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *SVM) SetTermCriteria(val *TermCriteria) {
	SVMNative_setTermCriteria_0(rcvr.GetNativeObjAddr(), val.Type, val.MaxCount, val.Epsilon)
	return
}
func (rcvr *SVM) SetType(val int) {
	SVMNative_setType_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *SVM) TrainAuto(samples *Mat, layout int, responses *Mat, kFold int, Cgrid *ParamGrid, gammaGrid *ParamGrid, pGrid *ParamGrid, nuGrid *ParamGrid, coeffGrid *ParamGrid, degreeGrid *ParamGrid, balanced bool) bool {
	retVal := SVMNative_trainAuto_0(rcvr.GetNativeObjAddr(), samples.GetNativeObjAddr(), layout, responses.GetNativeObjAddr(), kFold, Cgrid.GetNativeObjAddr(), gammaGrid.GetNativeObjAddr(), pGrid.GetNativeObjAddr(), nuGrid.GetNativeObjAddr(), coeffGrid.GetNativeObjAddr(), degreeGrid.GetNativeObjAddr(), balanced)
	return retVal
}
func (rcvr *SVM) TrainAuto2(samples *Mat, layout int, responses *Mat) bool {
	retVal := SVMNative_trainAuto_1(rcvr.GetNativeObjAddr(), samples.GetNativeObjAddr(), layout, responses.GetNativeObjAddr())
	return retVal
}
