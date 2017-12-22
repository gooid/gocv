package objdetect

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

type CascadeClassifier struct {
	nativeObj int64
}

func NewCascadeClassifier(addr int64) (rcvr *CascadeClassifier) {
	rcvr = &CascadeClassifier{}
	rcvr.nativeObj = addr
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func NewCascadeClassifier2(filename string) (rcvr *CascadeClassifier) {
	rcvr = &CascadeClassifier{}
	rcvr.nativeObj = CascadeClassifierNative_CascadeClassifier_0(filename)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
	return
}
func NewCascadeClassifier3() (rcvr *CascadeClassifier) {
	rcvr = &CascadeClassifier{}
	rcvr.nativeObj = CascadeClassifierNative_CascadeClassifier_1()
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
	return
}
func Convert(oldcascade string, newcascade string) bool {
	retVal := CascadeClassifierNative_convert_0(oldcascade, newcascade)
	return retVal
}
func (rcvr *CascadeClassifier) DetectMultiScale(image *Mat, objects *MatOfRect, scaleFactor float64, minNeighbors int, flags int, minSize *Size, maxSize *Size) {
	objects_mat := objects
	CascadeClassifierNative_detectMultiScale_0(rcvr.GetNativeObjAddr(), image.GetNativeObjAddr(), objects_mat.GetNativeObjAddr(), scaleFactor, minNeighbors, flags, minSize.Width, minSize.Height, maxSize.Width, maxSize.Height)
	return
}
func (rcvr *CascadeClassifier) DetectMultiScale2(image *Mat, objects *MatOfRect) {
	objects_mat := objects
	CascadeClassifierNative_detectMultiScale_1(rcvr.GetNativeObjAddr(), image.GetNativeObjAddr(), objects_mat.GetNativeObjAddr())
	return
}
func (rcvr *CascadeClassifier) DetectMultiScale21(image *Mat, objects *MatOfRect, numDetections *MatOfInt, scaleFactor float64, minNeighbors int, flags int, minSize *Size, maxSize *Size) {
	objects_mat := objects
	numDetections_mat := numDetections
	CascadeClassifierNative_detectMultiScale2_0(rcvr.GetNativeObjAddr(), image.GetNativeObjAddr(), objects_mat.GetNativeObjAddr(), numDetections_mat.GetNativeObjAddr(), scaleFactor, minNeighbors, flags, minSize.Width, minSize.Height, maxSize.Width, maxSize.Height)
	return
}
func (rcvr *CascadeClassifier) DetectMultiScale22(image *Mat, objects *MatOfRect, numDetections *MatOfInt) {
	objects_mat := objects
	numDetections_mat := numDetections
	CascadeClassifierNative_detectMultiScale2_1(rcvr.GetNativeObjAddr(), image.GetNativeObjAddr(), objects_mat.GetNativeObjAddr(), numDetections_mat.GetNativeObjAddr())
	return
}
func (rcvr *CascadeClassifier) DetectMultiScale3(image *Mat, objects *MatOfRect, rejectLevels *MatOfInt, levelWeights *MatOfDouble, scaleFactor float64, minNeighbors int, flags int, minSize *Size, maxSize *Size, outputRejectLevels bool) {
	objects_mat := objects
	rejectLevels_mat := rejectLevels
	levelWeights_mat := levelWeights
	CascadeClassifierNative_detectMultiScale3_0(rcvr.GetNativeObjAddr(), image.GetNativeObjAddr(), objects_mat.GetNativeObjAddr(), rejectLevels_mat.GetNativeObjAddr(), levelWeights_mat.GetNativeObjAddr(), scaleFactor, minNeighbors, flags, minSize.Width, minSize.Height, maxSize.Width, maxSize.Height, outputRejectLevels)
	return
}
func (rcvr *CascadeClassifier) DetectMultiScale32(image *Mat, objects *MatOfRect, rejectLevels *MatOfInt, levelWeights *MatOfDouble) {
	objects_mat := objects
	rejectLevels_mat := rejectLevels
	levelWeights_mat := levelWeights
	CascadeClassifierNative_detectMultiScale3_1(rcvr.GetNativeObjAddr(), image.GetNativeObjAddr(), objects_mat.GetNativeObjAddr(), rejectLevels_mat.GetNativeObjAddr(), levelWeights_mat.GetNativeObjAddr())
	return
}
func (rcvr *CascadeClassifier) Empty() bool {
	retVal := CascadeClassifierNative_empty_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *CascadeClassifier) finalize() {
	CascadeClassifierNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *CascadeClassifier) GetFeatureType() int {
	retVal := CascadeClassifierNative_getFeatureType_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *CascadeClassifier) GetNativeObjAddr() int64 {
	return rcvr.nativeObj
}
func (rcvr *CascadeClassifier) GetOriginalWindowSize() *Size {
	retVal := NewSize4(CascadeClassifierNative_getOriginalWindowSize_0(rcvr.GetNativeObjAddr()))
	return retVal
}
func (rcvr *CascadeClassifier) IsOldFormatCascade() bool {
	retVal := CascadeClassifierNative_isOldFormatCascade_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *CascadeClassifier) Load(filename string) bool {
	retVal := CascadeClassifierNative_load_0(rcvr.GetNativeObjAddr(), filename)
	return retVal
}
