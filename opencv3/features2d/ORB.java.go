package features2d

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

const ORBKBytes = 32
const ORBHARRIS_SCORE = 0
const ORBFAST_SCORE = 1

type ORB struct {
	*Feature2D
}

func NewORB(addr int64) (rcvr *ORB) {
	rcvr = &ORB{}
	rcvr.Feature2D = NewFeature2D(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func ORBCreate(nfeatures int, scaleFactor float32, nlevels int, edgeThreshold int, firstLevel int, WTA_K int, scoreType int, patchSize int, fastThreshold int) *ORB {
	retVal := NewORB(ORBNative_create_0(nfeatures, scaleFactor, nlevels, edgeThreshold, firstLevel, WTA_K, scoreType, patchSize, fastThreshold))
	return retVal
}
func ORBCreate2() *ORB {
	retVal := NewORB(ORBNative_create_1())
	return retVal
}
func (rcvr *ORB) finalize() {
	ORBNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *ORB) GetDefaultName() string {
	retVal := ORBNative_getDefaultName_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *ORB) GetEdgeThreshold() int {
	retVal := ORBNative_getEdgeThreshold_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *ORB) GetFastThreshold() int {
	retVal := ORBNative_getFastThreshold_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *ORB) GetFirstLevel() int {
	retVal := ORBNative_getFirstLevel_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *ORB) GetMaxFeatures() int {
	retVal := ORBNative_getMaxFeatures_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *ORB) GetNLevels() int {
	retVal := ORBNative_getNLevels_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *ORB) GetPatchSize() int {
	retVal := ORBNative_getPatchSize_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *ORB) GetScaleFactor() float64 {
	retVal := ORBNative_getScaleFactor_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *ORB) GetScoreType() int {
	retVal := ORBNative_getScoreType_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *ORB) GetWTA_K() int {
	retVal := ORBNative_getWTA_K_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *ORB) SetEdgeThreshold(edgeThreshold int) {
	ORBNative_setEdgeThreshold_0(rcvr.GetNativeObjAddr(), edgeThreshold)
	return
}
func (rcvr *ORB) SetFastThreshold(fastThreshold int) {
	ORBNative_setFastThreshold_0(rcvr.GetNativeObjAddr(), fastThreshold)
	return
}
func (rcvr *ORB) SetFirstLevel(firstLevel int) {
	ORBNative_setFirstLevel_0(rcvr.GetNativeObjAddr(), firstLevel)
	return
}
func (rcvr *ORB) SetMaxFeatures(maxFeatures int) {
	ORBNative_setMaxFeatures_0(rcvr.GetNativeObjAddr(), maxFeatures)
	return
}
func (rcvr *ORB) SetNLevels(nlevels int) {
	ORBNative_setNLevels_0(rcvr.GetNativeObjAddr(), nlevels)
	return
}
func (rcvr *ORB) SetPatchSize(patchSize int) {
	ORBNative_setPatchSize_0(rcvr.GetNativeObjAddr(), patchSize)
	return
}
func (rcvr *ORB) SetScaleFactor(scaleFactor float64) {
	ORBNative_setScaleFactor_0(rcvr.GetNativeObjAddr(), scaleFactor)
	return
}
func (rcvr *ORB) SetScoreType(scoreType int) {
	ORBNative_setScoreType_0(rcvr.GetNativeObjAddr(), scoreType)
	return
}
func (rcvr *ORB) SetWTA_K(wta_k int) {
	ORBNative_setWTA_K_0(rcvr.GetNativeObjAddr(), wta_k)
	return
}
