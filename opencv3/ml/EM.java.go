package ml

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

const EMCOV_MAT_SPHERICAL = 0
const EMCOV_MAT_DIAGONAL = 1
const EMCOV_MAT_GENERIC = 2
const EMDEFAULT_NCLUSTERS = 5
const EMDEFAULT_MAX_ITERS = 100
const EMSTART_E_STEP = 1
const EMSTART_M_STEP = 2
const EMSTART_AUTO_STEP = 0

var EMCOV_MAT_DEFAULT = EMCOV_MAT_DIAGONAL

type EM struct {
	*StatModel
}

func NewEM(addr int64) (rcvr *EM) {
	rcvr = &EM{}
	rcvr.StatModel = NewStatModel(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func EMCreate() *EM {
	retVal := NewEM(EMNative_create_0())
	return retVal
}
func (rcvr *EM) finalize() {
	EMNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *EM) GetClustersNumber() int {
	retVal := EMNative_getClustersNumber_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *EM) GetCovarianceMatrixType() int {
	retVal := EMNative_getCovarianceMatrixType_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *EM) GetCovs() (covs []*Mat) {
	covs_mat := NewMat2()
	EMNative_getCovs_0(rcvr.GetNativeObjAddr(), covs_mat.GetNativeObjAddr())
	covs = ConvertersToVectorMat(covs_mat)
	covs_mat.Release()
	return
}
func (rcvr *EM) GetMeans() *Mat {
	retVal := NewMat(EMNative_getMeans_0(rcvr.GetNativeObjAddr()))
	return retVal
}
func (rcvr *EM) GetTermCriteria() *TermCriteria {
	retVal := NewTermCriteria3(EMNative_getTermCriteria_0(rcvr.GetNativeObjAddr()))
	return retVal
}
func (rcvr *EM) GetWeights() *Mat {
	retVal := NewMat(EMNative_getWeights_0(rcvr.GetNativeObjAddr()))
	return retVal
}
func EMLoad(filepath string, nodeName string) *EM {
	retVal := NewEM(EMNative_load_0(filepath, nodeName))
	return retVal
}
func EMLoad2(filepath string) *EM {
	retVal := NewEM(EMNative_load_1(filepath))
	return retVal
}
func (rcvr *EM) Predict(samples *Mat, results *Mat, flags int) float32 {
	retVal := EMNative_predict_0(rcvr.GetNativeObjAddr(), samples.GetNativeObjAddr(), results.GetNativeObjAddr(), flags)
	return retVal
}
func (rcvr *EM) Predict2(samples *Mat) float32 {
	retVal := EMNative_predict_1(rcvr.GetNativeObjAddr(), samples.GetNativeObjAddr())
	return retVal
}
func (rcvr *EM) Predict21(sample *Mat, probs *Mat) []float64 {
	retVal := EMNative_predict2_0(rcvr.GetNativeObjAddr(), sample.GetNativeObjAddr(), probs.GetNativeObjAddr())
	return retVal
}
func (rcvr *EM) SetClustersNumber(val int) {
	EMNative_setClustersNumber_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *EM) SetCovarianceMatrixType(val int) {
	EMNative_setCovarianceMatrixType_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *EM) SetTermCriteria(val *TermCriteria) {
	EMNative_setTermCriteria_0(rcvr.GetNativeObjAddr(), val.Type, val.MaxCount, val.Epsilon)
	return
}
func (rcvr *EM) TrainE(samples *Mat, means0 *Mat, covs0 *Mat, weights0 *Mat, logLikelihoods *Mat, labels *Mat, probs *Mat) bool {
	retVal := EMNative_trainE_0(rcvr.GetNativeObjAddr(), samples.GetNativeObjAddr(), means0.GetNativeObjAddr(), covs0.GetNativeObjAddr(), weights0.GetNativeObjAddr(), logLikelihoods.GetNativeObjAddr(), labels.GetNativeObjAddr(), probs.GetNativeObjAddr())
	return retVal
}
func (rcvr *EM) TrainE2(samples *Mat, means0 *Mat) bool {
	retVal := EMNative_trainE_1(rcvr.GetNativeObjAddr(), samples.GetNativeObjAddr(), means0.GetNativeObjAddr())
	return retVal
}
func (rcvr *EM) TrainEM(samples *Mat, logLikelihoods *Mat, labels *Mat, probs *Mat) bool {
	retVal := EMNative_trainEM_0(rcvr.GetNativeObjAddr(), samples.GetNativeObjAddr(), logLikelihoods.GetNativeObjAddr(), labels.GetNativeObjAddr(), probs.GetNativeObjAddr())
	return retVal
}
func (rcvr *EM) TrainEM2(samples *Mat) bool {
	retVal := EMNative_trainEM_1(rcvr.GetNativeObjAddr(), samples.GetNativeObjAddr())
	return retVal
}
func (rcvr *EM) TrainM(samples *Mat, probs0 *Mat, logLikelihoods *Mat, labels *Mat, probs *Mat) bool {
	retVal := EMNative_trainM_0(rcvr.GetNativeObjAddr(), samples.GetNativeObjAddr(), probs0.GetNativeObjAddr(), logLikelihoods.GetNativeObjAddr(), labels.GetNativeObjAddr(), probs.GetNativeObjAddr())
	return retVal
}
func (rcvr *EM) TrainM2(samples *Mat, probs0 *Mat) bool {
	retVal := EMNative_trainM_1(rcvr.GetNativeObjAddr(), samples.GetNativeObjAddr(), probs0.GetNativeObjAddr())
	return retVal
}
