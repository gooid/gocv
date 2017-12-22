package video

import (
	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

const cV_LKFLOW_INITIAL_GUESSESVideo = 4
const cV_LKFLOW_GET_MIN_EIGENVALSVideo = 8
const OPTFLOW_USE_INITIAL_FLOW = 4
const OPTFLOW_LK_GET_MIN_EIGENVALS = 8
const OPTFLOW_FARNEBACK_GAUSSIAN = 256
const MOTION_TRANSLATION = 0
const MOTION_EUCLIDEAN = 1
const MOTION_AFFINE = 2
const MOTION_HOMOGRAPHY = 3

func CamShift(probImage *Mat, window *Rect, criteria *TermCriteria) *RotatedRect {
	window_out := make([]float64, 4)
	retVal := NewRotatedRect3(VideoNative_CamShift_0(probImage.GetNativeObjAddr(), window.X, window.Y, window.Width, window.Height, window_out, criteria.Type, criteria.MaxCount, criteria.Epsilon))
	if window != nil {
		window.X = int(window_out[0])
		window.Y = int(window_out[1])
		window.Width = int(window_out[2])
		window.Height = int(window_out[3])
	}
	return retVal
}
func BuildOpticalFlowPyramid(img *Mat, winSize *Size, maxLevel int, withDerivatives bool, pyrBorder int, derivBorder int, tryReuseInputImage bool) ([]*Mat, int) {
	pyramid_mat := NewMat2()
	retVal := VideoNative_buildOpticalFlowPyramid_0(img.GetNativeObjAddr(), pyramid_mat.GetNativeObjAddr(), winSize.Width, winSize.Height, maxLevel, withDerivatives, pyrBorder, derivBorder, tryReuseInputImage)
	pyramid := ConvertersToVectorMat(pyramid_mat)
	pyramid_mat.Release()
	return pyramid, retVal
}
func BuildOpticalFlowPyramid2(img *Mat, winSize *Size, maxLevel int) ([]*Mat, int) {
	pyramid_mat := NewMat2()
	retVal := VideoNative_buildOpticalFlowPyramid_1(img.GetNativeObjAddr(), pyramid_mat.GetNativeObjAddr(), winSize.Width, winSize.Height, maxLevel)
	pyramid := ConvertersToVectorMat(pyramid_mat)
	pyramid_mat.Release()
	return pyramid, retVal
}
func CalcOpticalFlowFarneback(prev *Mat, next *Mat, flow *Mat, pyr_scale float64, levels int, winsize int, iterations int, poly_n int, poly_sigma float64, flags int) {
	VideoNative_calcOpticalFlowFarneback_0(prev.GetNativeObjAddr(), next.GetNativeObjAddr(), flow.GetNativeObjAddr(), pyr_scale, levels, winsize, iterations, poly_n, poly_sigma, flags)
	return
}
func CalcOpticalFlowPyrLK(prevImg *Mat, nextImg *Mat, prevPts *MatOfPoint2f, nextPts *MatOfPoint2f, status *MatOfByte, err *MatOfFloat, winSize *Size, maxLevel int, criteria *TermCriteria, flags int, minEigThreshold float64) {
	prevPts_mat := prevPts
	nextPts_mat := nextPts
	status_mat := status
	err_mat := err
	VideoNative_calcOpticalFlowPyrLK_0(prevImg.GetNativeObjAddr(), nextImg.GetNativeObjAddr(), prevPts_mat.GetNativeObjAddr(), nextPts_mat.GetNativeObjAddr(), status_mat.GetNativeObjAddr(), err_mat.GetNativeObjAddr(), winSize.Width, winSize.Height, maxLevel, criteria.Type, criteria.MaxCount, criteria.Epsilon, flags, minEigThreshold)
	return
}
func CalcOpticalFlowPyrLK2(prevImg *Mat, nextImg *Mat, prevPts *MatOfPoint2f, nextPts *MatOfPoint2f, status *MatOfByte, err *MatOfFloat, winSize *Size, maxLevel int) {
	prevPts_mat := prevPts
	nextPts_mat := nextPts
	status_mat := status
	err_mat := err
	VideoNative_calcOpticalFlowPyrLK_1(prevImg.GetNativeObjAddr(), nextImg.GetNativeObjAddr(), prevPts_mat.GetNativeObjAddr(), nextPts_mat.GetNativeObjAddr(), status_mat.GetNativeObjAddr(), err_mat.GetNativeObjAddr(), winSize.Width, winSize.Height, maxLevel)
	return
}
func CalcOpticalFlowPyrLK3(prevImg *Mat, nextImg *Mat, prevPts *MatOfPoint2f, nextPts *MatOfPoint2f, status *MatOfByte, err *MatOfFloat) {
	prevPts_mat := prevPts
	nextPts_mat := nextPts
	status_mat := status
	err_mat := err
	VideoNative_calcOpticalFlowPyrLK_2(prevImg.GetNativeObjAddr(), nextImg.GetNativeObjAddr(), prevPts_mat.GetNativeObjAddr(), nextPts_mat.GetNativeObjAddr(), status_mat.GetNativeObjAddr(), err_mat.GetNativeObjAddr())
	return
}
func CreateBackgroundSubtractorKNN(history int, dist2Threshold float64, detectShadows bool) *BackgroundSubtractorKNN {
	retVal := NewBackgroundSubtractorKNN(VideoNative_createBackgroundSubtractorKNN_0(history, dist2Threshold, detectShadows))
	return retVal
}
func CreateBackgroundSubtractorKNN2() *BackgroundSubtractorKNN {
	retVal := NewBackgroundSubtractorKNN(VideoNative_createBackgroundSubtractorKNN_1())
	return retVal
}
func CreateBackgroundSubtractorMOG2(history int, varThreshold float64, detectShadows bool) *BackgroundSubtractorMOG2 {
	retVal := NewBackgroundSubtractorMOG2(VideoNative_createBackgroundSubtractorMOG2_0(history, varThreshold, detectShadows))
	return retVal
}
func CreateBackgroundSubtractorMOG22() *BackgroundSubtractorMOG2 {
	retVal := NewBackgroundSubtractorMOG2(VideoNative_createBackgroundSubtractorMOG2_1())
	return retVal
}
func CreateOptFlow_DualTVL1() *DualTVL1OpticalFlow {
	retVal := NewDualTVL1OpticalFlow(VideoNative_createOptFlow_DualTVL1_0())
	return retVal
}
func EstimateRigidTransform(src *Mat, dst *Mat, fullAffine bool) *Mat {
	retVal := NewMat(VideoNative_estimateRigidTransform_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), fullAffine))
	return retVal
}
func FindTransformECC(templateImage *Mat, inputImage *Mat, warpMatrix *Mat, motionType int, criteria *TermCriteria, inputMask *Mat) float64 {
	retVal := VideoNative_findTransformECC_0(templateImage.GetNativeObjAddr(), inputImage.GetNativeObjAddr(), warpMatrix.GetNativeObjAddr(), motionType, criteria.Type, criteria.MaxCount, criteria.Epsilon, inputMask.GetNativeObjAddr())
	return retVal
}
func FindTransformECC2(templateImage *Mat, inputImage *Mat, warpMatrix *Mat, motionType int) float64 {
	retVal := VideoNative_findTransformECC_1(templateImage.GetNativeObjAddr(), inputImage.GetNativeObjAddr(), warpMatrix.GetNativeObjAddr(), motionType)
	return retVal
}
func FindTransformECC3(templateImage *Mat, inputImage *Mat, warpMatrix *Mat) float64 {
	retVal := VideoNative_findTransformECC_2(templateImage.GetNativeObjAddr(), inputImage.GetNativeObjAddr(), warpMatrix.GetNativeObjAddr())
	return retVal
}
func MeanShift(probImage *Mat, window *Rect, criteria *TermCriteria) int {
	window_out := make([]float64, 4)
	retVal := VideoNative_meanShift_0(probImage.GetNativeObjAddr(), window.X, window.Y, window.Width, window.Height, window_out, criteria.Type, criteria.MaxCount, criteria.Epsilon)
	if window != nil {
		window.X = int(window_out[0])
		window.Y = int(window_out[1])
		window.Width = int(window_out[2])
		window.Height = int(window_out[3])
	}
	return retVal
}
