package video

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

type KalmanFilter struct {
	nativeObj int64
}

func NewKalmanFilter(addr int64) (rcvr *KalmanFilter) {
	rcvr = &KalmanFilter{}
	rcvr.nativeObj = addr
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func NewKalmanFilter2(dynamParams int, measureParams int, controlParams int, rtype int) (rcvr *KalmanFilter) {
	rcvr = &KalmanFilter{}
	rcvr.nativeObj = KalmanFilterNative_KalmanFilter_0(dynamParams, measureParams, controlParams, rtype)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
	return
}
func NewKalmanFilter3(dynamParams int, measureParams int) (rcvr *KalmanFilter) {
	rcvr = &KalmanFilter{}
	rcvr.nativeObj = KalmanFilterNative_KalmanFilter_1(dynamParams, measureParams)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
	return
}
func NewKalmanFilter4() (rcvr *KalmanFilter) {
	rcvr = &KalmanFilter{}
	rcvr.nativeObj = KalmanFilterNative_KalmanFilter_2()
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
	return
}
func (rcvr *KalmanFilter) Correct(measurement *Mat) *Mat {
	retVal := NewMat(KalmanFilterNative_correct_0(rcvr.nativeObj, measurement.GetNativeObjAddr()))
	return retVal
}
func (rcvr *KalmanFilter) finalize() {
	KalmanFilterNative_delete(rcvr.nativeObj)
}
func (rcvr *KalmanFilter) GetNativeObjAddr() int64 {
	return rcvr.nativeObj
}
func (rcvr *KalmanFilter) Get_controlMatrix() *Mat {
	retVal := NewMat(KalmanFilterNative_get_controlMatrix_0(rcvr.nativeObj))
	return retVal
}
func (rcvr *KalmanFilter) Get_errorCovPost() *Mat {
	retVal := NewMat(KalmanFilterNative_get_errorCovPost_0(rcvr.nativeObj))
	return retVal
}
func (rcvr *KalmanFilter) Get_errorCovPre() *Mat {
	retVal := NewMat(KalmanFilterNative_get_errorCovPre_0(rcvr.nativeObj))
	return retVal
}
func (rcvr *KalmanFilter) Get_gain() *Mat {
	retVal := NewMat(KalmanFilterNative_get_gain_0(rcvr.nativeObj))
	return retVal
}
func (rcvr *KalmanFilter) Get_measurementMatrix() *Mat {
	retVal := NewMat(KalmanFilterNative_get_measurementMatrix_0(rcvr.nativeObj))
	return retVal
}
func (rcvr *KalmanFilter) Get_measurementNoiseCov() *Mat {
	retVal := NewMat(KalmanFilterNative_get_measurementNoiseCov_0(rcvr.nativeObj))
	return retVal
}
func (rcvr *KalmanFilter) Get_processNoiseCov() *Mat {
	retVal := NewMat(KalmanFilterNative_get_processNoiseCov_0(rcvr.nativeObj))
	return retVal
}
func (rcvr *KalmanFilter) Get_statePost() *Mat {
	retVal := NewMat(KalmanFilterNative_get_statePost_0(rcvr.nativeObj))
	return retVal
}
func (rcvr *KalmanFilter) Get_statePre() *Mat {
	retVal := NewMat(KalmanFilterNative_get_statePre_0(rcvr.nativeObj))
	return retVal
}
func (rcvr *KalmanFilter) Get_transitionMatrix() *Mat {
	retVal := NewMat(KalmanFilterNative_get_transitionMatrix_0(rcvr.nativeObj))
	return retVal
}
func (rcvr *KalmanFilter) Predict(control *Mat) *Mat {
	retVal := NewMat(KalmanFilterNative_predict_0(rcvr.nativeObj, control.GetNativeObjAddr()))
	return retVal
}
func (rcvr *KalmanFilter) Predict2() *Mat {
	retVal := NewMat(KalmanFilterNative_predict_1(rcvr.nativeObj))
	return retVal
}
func (rcvr *KalmanFilter) Set_controlMatrix(controlMatrix *Mat) {
	KalmanFilterNative_set_controlMatrix_0(rcvr.nativeObj, controlMatrix.GetNativeObjAddr())
	return
}
func (rcvr *KalmanFilter) Set_errorCovPost(errorCovPost *Mat) {
	KalmanFilterNative_set_errorCovPost_0(rcvr.nativeObj, errorCovPost.GetNativeObjAddr())
	return
}
func (rcvr *KalmanFilter) Set_errorCovPre(errorCovPre *Mat) {
	KalmanFilterNative_set_errorCovPre_0(rcvr.nativeObj, errorCovPre.GetNativeObjAddr())
	return
}
func (rcvr *KalmanFilter) Set_gain(gain *Mat) {
	KalmanFilterNative_set_gain_0(rcvr.nativeObj, gain.GetNativeObjAddr())
	return
}
func (rcvr *KalmanFilter) Set_measurementMatrix(measurementMatrix *Mat) {
	KalmanFilterNative_set_measurementMatrix_0(rcvr.nativeObj, measurementMatrix.GetNativeObjAddr())
	return
}
func (rcvr *KalmanFilter) Set_measurementNoiseCov(measurementNoiseCov *Mat) {
	KalmanFilterNative_set_measurementNoiseCov_0(rcvr.nativeObj, measurementNoiseCov.GetNativeObjAddr())
	return
}
func (rcvr *KalmanFilter) Set_processNoiseCov(processNoiseCov *Mat) {
	KalmanFilterNative_set_processNoiseCov_0(rcvr.nativeObj, processNoiseCov.GetNativeObjAddr())
	return
}
func (rcvr *KalmanFilter) Set_statePost(statePost *Mat) {
	KalmanFilterNative_set_statePost_0(rcvr.nativeObj, statePost.GetNativeObjAddr())
	return
}
func (rcvr *KalmanFilter) Set_statePre(statePre *Mat) {
	KalmanFilterNative_set_statePre_0(rcvr.nativeObj, statePre.GetNativeObjAddr())
	return
}
func (rcvr *KalmanFilter) Set_transitionMatrix(transitionMatrix *Mat) {
	KalmanFilterNative_set_transitionMatrix_0(rcvr.nativeObj, transitionMatrix.GetNativeObjAddr())
	return
}
