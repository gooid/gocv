package features2d

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

type BOWKMeansTrainer struct {
	*BOWTrainer
}

func NewBOWKMeansTrainer(addr int64) (rcvr *BOWKMeansTrainer) {
	rcvr = &BOWKMeansTrainer{}
	rcvr.BOWTrainer = NewBOWTrainer(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func NewBOWKMeansTrainer2(clusterCount int, termcrit *TermCriteria, attempts int, flags int) (rcvr *BOWKMeansTrainer) {
	rcvr = &BOWKMeansTrainer{}
	rcvr.BOWTrainer = NewBOWTrainer(BOWKMeansTrainerNative_BOWKMeansTrainer_0(clusterCount, termcrit.Type, termcrit.MaxCount, termcrit.Epsilon, attempts, flags))
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
	return
}
func NewBOWKMeansTrainer3(clusterCount int) (rcvr *BOWKMeansTrainer) {
	rcvr = &BOWKMeansTrainer{}
	rcvr.BOWTrainer = NewBOWTrainer(BOWKMeansTrainerNative_BOWKMeansTrainer_1(clusterCount))
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
	return
}
func (rcvr *BOWKMeansTrainer) Cluster(descriptors *Mat) *Mat {
	retVal := NewMat(BOWKMeansTrainerNative_cluster_0(rcvr.GetNativeObjAddr(), descriptors.GetNativeObjAddr()))
	return retVal
}
func (rcvr *BOWKMeansTrainer) Cluster2() *Mat {
	retVal := NewMat(BOWKMeansTrainerNative_cluster_1(rcvr.GetNativeObjAddr()))
	return retVal
}
func (rcvr *BOWKMeansTrainer) finalize() {
	BOWKMeansTrainerNative_delete(rcvr.GetNativeObjAddr())
}
