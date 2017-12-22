package features2d

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

type GFTTDetector struct {
	*Feature2D
}

func NewGFTTDetector(addr int64) (rcvr *GFTTDetector) {
	rcvr = &GFTTDetector{}
	rcvr.Feature2D = NewFeature2D(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func GFTTDetectorCreate(maxCorners int, qualityLevel float64, minDistance float64, blockSize int, gradiantSize int, useHarrisDetector bool, k float64) *GFTTDetector {
	retVal := NewGFTTDetector(GFTTDetectorNative_create_0(maxCorners, qualityLevel, minDistance, blockSize, gradiantSize, useHarrisDetector, k))
	return retVal
}
func GFTTDetectorCreate2(maxCorners int, qualityLevel float64, minDistance float64, blockSize int, gradiantSize int) *GFTTDetector {
	retVal := NewGFTTDetector(GFTTDetectorNative_create_1(maxCorners, qualityLevel, minDistance, blockSize, gradiantSize))
	return retVal
}
func GFTTDetectorCreate3(maxCorners int, qualityLevel float64, minDistance float64, blockSize int, useHarrisDetector bool, k float64) *GFTTDetector {
	retVal := NewGFTTDetector(GFTTDetectorNative_create_2(maxCorners, qualityLevel, minDistance, blockSize, useHarrisDetector, k))
	return retVal
}
func GFTTDetectorCreate4() *GFTTDetector {
	retVal := NewGFTTDetector(GFTTDetectorNative_create_3())
	return retVal
}
func (rcvr *GFTTDetector) finalize() {
	GFTTDetectorNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *GFTTDetector) GetBlockSize() int {
	retVal := GFTTDetectorNative_getBlockSize_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *GFTTDetector) GetDefaultName() string {
	retVal := GFTTDetectorNative_getDefaultName_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *GFTTDetector) GetHarrisDetector() bool {
	retVal := GFTTDetectorNative_getHarrisDetector_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *GFTTDetector) GetK() float64 {
	retVal := GFTTDetectorNative_getK_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *GFTTDetector) GetMaxFeatures() int {
	retVal := GFTTDetectorNative_getMaxFeatures_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *GFTTDetector) GetMinDistance() float64 {
	retVal := GFTTDetectorNative_getMinDistance_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *GFTTDetector) GetQualityLevel() float64 {
	retVal := GFTTDetectorNative_getQualityLevel_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *GFTTDetector) SetBlockSize(blockSize int) {
	GFTTDetectorNative_setBlockSize_0(rcvr.GetNativeObjAddr(), blockSize)
	return
}
func (rcvr *GFTTDetector) SetHarrisDetector(val bool) {
	GFTTDetectorNative_setHarrisDetector_0(rcvr.GetNativeObjAddr(), val)
	return
}
func (rcvr *GFTTDetector) SetK(k float64) {
	GFTTDetectorNative_setK_0(rcvr.GetNativeObjAddr(), k)
	return
}
func (rcvr *GFTTDetector) SetMaxFeatures(maxFeatures int) {
	GFTTDetectorNative_setMaxFeatures_0(rcvr.GetNativeObjAddr(), maxFeatures)
	return
}
func (rcvr *GFTTDetector) SetMinDistance(minDistance float64) {
	GFTTDetectorNative_setMinDistance_0(rcvr.GetNativeObjAddr(), minDistance)
	return
}
func (rcvr *GFTTDetector) SetQualityLevel(qlevel float64) {
	GFTTDetectorNative_setQualityLevel_0(rcvr.GetNativeObjAddr(), qlevel)
	return
}
