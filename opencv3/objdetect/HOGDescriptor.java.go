package objdetect

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

const HOGDescriptorL2Hys = 0
const HOGDescriptorDEFAULT_NLEVELS = 64

type HOGDescriptor struct {
	nativeObj int64
}

func NewHOGDescriptor(addr int64) (rcvr *HOGDescriptor) {
	rcvr = &HOGDescriptor{}
	rcvr.nativeObj = addr
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func NewHOGDescriptor2(_winSize *Size, _blockSize *Size, _blockStride *Size, _cellSize *Size, _nbins int, _derivAperture int, _winSigma float64, _histogramNormType int, _L2HysThreshold float64, _gammaCorrection bool, _nlevels int, _signedGradient bool) (rcvr *HOGDescriptor) {
	rcvr = &HOGDescriptor{}
	rcvr.nativeObj = HOGDescriptorNative_HOGDescriptor_0(_winSize.Width, _winSize.Height, _blockSize.Width, _blockSize.Height, _blockStride.Width, _blockStride.Height, _cellSize.Width, _cellSize.Height, _nbins, _derivAperture, _winSigma, _histogramNormType, _L2HysThreshold, _gammaCorrection, _nlevels, _signedGradient)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
	return
}
func NewHOGDescriptor3(_winSize *Size, _blockSize *Size, _blockStride *Size, _cellSize *Size, _nbins int) (rcvr *HOGDescriptor) {
	rcvr = &HOGDescriptor{}
	rcvr.nativeObj = HOGDescriptorNative_HOGDescriptor_1(_winSize.Width, _winSize.Height, _blockSize.Width, _blockSize.Height, _blockStride.Width, _blockStride.Height, _cellSize.Width, _cellSize.Height, _nbins)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
	return
}
func NewHOGDescriptor4(filename string) (rcvr *HOGDescriptor) {
	rcvr = &HOGDescriptor{}
	rcvr.nativeObj = HOGDescriptorNative_HOGDescriptor_2(filename)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
	return
}
func NewHOGDescriptor5() (rcvr *HOGDescriptor) {
	rcvr = &HOGDescriptor{}
	rcvr.nativeObj = HOGDescriptorNative_HOGDescriptor_3()
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
	return
}
func (rcvr *HOGDescriptor) CheckDetectorSize() bool {
	retVal := HOGDescriptorNative_checkDetectorSize_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *HOGDescriptor) Compute(img *Mat, descriptors *MatOfFloat, winStride *Size, padding *Size, locations *MatOfPoint) {
	descriptors_mat := descriptors
	locations_mat := locations
	HOGDescriptorNative_compute_0(rcvr.nativeObj, img.GetNativeObjAddr(), descriptors_mat.GetNativeObjAddr(), winStride.Width, winStride.Height, padding.Width, padding.Height, locations_mat.GetNativeObjAddr())
	return
}
func (rcvr *HOGDescriptor) Compute2(img *Mat, descriptors *MatOfFloat) {
	descriptors_mat := descriptors
	HOGDescriptorNative_compute_1(rcvr.nativeObj, img.GetNativeObjAddr(), descriptors_mat.GetNativeObjAddr())
	return
}
func (rcvr *HOGDescriptor) ComputeGradient(img *Mat, grad *Mat, angleOfs *Mat, paddingTL *Size, paddingBR *Size) {
	HOGDescriptorNative_computeGradient_0(rcvr.nativeObj, img.GetNativeObjAddr(), grad.GetNativeObjAddr(), angleOfs.GetNativeObjAddr(), paddingTL.Width, paddingTL.Height, paddingBR.Width, paddingBR.Height)
	return
}
func (rcvr *HOGDescriptor) ComputeGradient2(img *Mat, grad *Mat, angleOfs *Mat) {
	HOGDescriptorNative_computeGradient_1(rcvr.nativeObj, img.GetNativeObjAddr(), grad.GetNativeObjAddr(), angleOfs.GetNativeObjAddr())
	return
}
func (rcvr *HOGDescriptor) Detect(img *Mat, foundLocations *MatOfPoint, weights *MatOfDouble, hitThreshold float64, winStride *Size, padding *Size, searchLocations *MatOfPoint) {
	foundLocations_mat := foundLocations
	weights_mat := weights
	searchLocations_mat := searchLocations
	HOGDescriptorNative_detect_0(rcvr.nativeObj, img.GetNativeObjAddr(), foundLocations_mat.GetNativeObjAddr(), weights_mat.GetNativeObjAddr(), hitThreshold, winStride.Width, winStride.Height, padding.Width, padding.Height, searchLocations_mat.GetNativeObjAddr())
	return
}
func (rcvr *HOGDescriptor) Detect2(img *Mat, foundLocations *MatOfPoint, weights *MatOfDouble) {
	foundLocations_mat := foundLocations
	weights_mat := weights
	HOGDescriptorNative_detect_1(rcvr.nativeObj, img.GetNativeObjAddr(), foundLocations_mat.GetNativeObjAddr(), weights_mat.GetNativeObjAddr())
	return
}
func (rcvr *HOGDescriptor) DetectMultiScale(img *Mat, foundLocations *MatOfRect, foundWeights *MatOfDouble, hitThreshold float64, winStride *Size, padding *Size, scale float64, finalThreshold float64, useMeanshiftGrouping bool) {
	foundLocations_mat := foundLocations
	foundWeights_mat := foundWeights
	HOGDescriptorNative_detectMultiScale_0(rcvr.nativeObj, img.GetNativeObjAddr(), foundLocations_mat.GetNativeObjAddr(), foundWeights_mat.GetNativeObjAddr(), hitThreshold, winStride.Width, winStride.Height, padding.Width, padding.Height, scale, finalThreshold, useMeanshiftGrouping)
	return
}
func (rcvr *HOGDescriptor) DetectMultiScale2(img *Mat, foundLocations *MatOfRect, foundWeights *MatOfDouble) {
	foundLocations_mat := foundLocations
	foundWeights_mat := foundWeights
	HOGDescriptorNative_detectMultiScale_1(rcvr.nativeObj, img.GetNativeObjAddr(), foundLocations_mat.GetNativeObjAddr(), foundWeights_mat.GetNativeObjAddr())
	return
}
func (rcvr *HOGDescriptor) finalize() {
	HOGDescriptorNative_delete(rcvr.nativeObj)
}
func GetDaimlerPeopleDetector() *MatOfFloat {
	retVal := MatOfFloatFromNativeAddr(HOGDescriptorNative_getDaimlerPeopleDetector_0())
	return retVal
}
