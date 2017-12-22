package features2d

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

type BOWTrainer struct {
	nativeObj int64
}

func NewBOWTrainer(addr int64) (rcvr *BOWTrainer) {
	rcvr = &BOWTrainer{}
	rcvr.nativeObj = addr
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func (rcvr *BOWTrainer) Add(descriptors *Mat) {
	BOWTrainerNative_add_0(rcvr.nativeObj, descriptors.GetNativeObjAddr())
	return
}
func (rcvr *BOWTrainer) Clear() {
	BOWTrainerNative_clear_0(rcvr.nativeObj)
	return
}
func (rcvr *BOWTrainer) Cluster(descriptors *Mat) *Mat {
	retVal := NewMat(BOWTrainerNative_cluster_0(rcvr.nativeObj, descriptors.GetNativeObjAddr()))
	return retVal
}
func (rcvr *BOWTrainer) Cluster2() *Mat {
	retVal := NewMat(BOWTrainerNative_cluster_1(rcvr.nativeObj))
	return retVal
}
func (rcvr *BOWTrainer) DescriptorsCount() int {
	retVal := BOWTrainerNative_descriptorsCount_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *BOWTrainer) finalize() {
	BOWTrainerNative_delete(rcvr.nativeObj)
}
func (rcvr *BOWTrainer) GetDescriptors() []*Mat {
	retValMat := NewMat(BOWTrainerNative_getDescriptors_0(rcvr.nativeObj))
	return ConvertersToVectorMat(retValMat)
}
func (rcvr *BOWTrainer) GetNativeObjAddr() int64 {
	return rcvr.nativeObj
}
