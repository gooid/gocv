package ml

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

const KNearestBRUTE_FORCE = 1
const KNearestKDTREE = 2

type KNearest struct {
	*StatModel
}

func NewKNearest(addr int64) (rcvr *KNearest) {
	rcvr = &KNearest{}
	rcvr.StatModel = NewStatModel(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func KNearestCreate() *KNearest {
	retVal := NewKNearest(KNearestNative_create_0())
	return retVal
}
func (rcvr *KNearest) finalize() {
	KNearestNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *KNearest) FindNearest(samples *Mat, k int, results *Mat, neighborResponses *Mat, dist *Mat) float32 {
	retVal := KNearestNative_findNearest_0(rcvr.GetNativeObjAddr(), samples.GetNativeObjAddr(), k, results.GetNativeObjAddr(), neighborResponses.GetNativeObjAddr(), dist.GetNativeObjAddr())
	return retVal
}
func (rcvr *KNearest) FindNearest2(samples *Mat, k int, results *Mat) float32 {
	retVal := KNearestNative_findNearest_1(rcvr.GetNativeObjAddr(), samples.GetNativeObjAddr(), k, results.GetNativeObjAddr())
	return retVal
}
func (rcvr *KNearest) GetAlgorithmType() int {
	retVal := KNearestNative_getAlgorithmType_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *KNearest) GetDefaultK() int {
	retVal := KNearestNative_getDefaultK_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *KNearest) GetEmax() int {
	retVal := KNearestNative_getEmax_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *KNearest) GetIsClassifier() bool {
	retVal := KNearestNative_getIsClassifier_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *KNearest) SetAlgorithmType(val int) {
	KNearestNative_setAlgorithmType_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *KNearest) SetDefaultK(val int) {
	KNearestNative_setDefaultK_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *KNearest) SetEmax(val int) {
	KNearestNative_setEmax_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *KNearest) SetIsClassifier(val bool) {
	KNearestNative_setIsClassifier_0(rcvr.GetNativeObjAddr(), val)
	return
}
