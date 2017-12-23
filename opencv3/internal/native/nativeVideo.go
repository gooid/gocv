package core

/*

#include "jni.h"

extern void Java_org_opencv_video_BackgroundSubtractor_apply_10(JNIEnv*, jclass, jlong, jlong, jlong, jdouble);
extern void Java_org_opencv_video_BackgroundSubtractor_apply_11(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_video_BackgroundSubtractor_delete(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_video_BackgroundSubtractor_getBackgroundImage_10(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_video_BackgroundSubtractorKNN_delete(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_video_BackgroundSubtractorKNN_getDetectShadows_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_video_BackgroundSubtractorKNN_getDist2Threshold_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_video_BackgroundSubtractorKNN_getHistory_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_video_BackgroundSubtractorKNN_getNSamples_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_video_BackgroundSubtractorKNN_getShadowThreshold_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_video_BackgroundSubtractorKNN_getShadowValue_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_video_BackgroundSubtractorKNN_getkNNSamples_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_video_BackgroundSubtractorKNN_setDetectShadows_10(JNIEnv*, jclass, jlong, jboolean);
extern void Java_org_opencv_video_BackgroundSubtractorKNN_setDist2Threshold_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_video_BackgroundSubtractorKNN_setHistory_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_video_BackgroundSubtractorKNN_setNSamples_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_video_BackgroundSubtractorKNN_setShadowThreshold_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_video_BackgroundSubtractorKNN_setShadowValue_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_video_BackgroundSubtractorKNN_setkNNSamples_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_video_BackgroundSubtractorMOG2_apply_10(JNIEnv*, jclass, jlong, jlong, jlong, jdouble);
extern void Java_org_opencv_video_BackgroundSubtractorMOG2_apply_11(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_video_BackgroundSubtractorMOG2_delete(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_video_BackgroundSubtractorMOG2_getBackgroundRatio_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_video_BackgroundSubtractorMOG2_getComplexityReductionThreshold_10(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_video_BackgroundSubtractorMOG2_getDetectShadows_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_video_BackgroundSubtractorMOG2_getHistory_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_video_BackgroundSubtractorMOG2_getNMixtures_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_video_BackgroundSubtractorMOG2_getShadowThreshold_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_video_BackgroundSubtractorMOG2_getShadowValue_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_video_BackgroundSubtractorMOG2_getVarInit_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_video_BackgroundSubtractorMOG2_getVarMax_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_video_BackgroundSubtractorMOG2_getVarMin_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_video_BackgroundSubtractorMOG2_getVarThresholdGen_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_video_BackgroundSubtractorMOG2_getVarThreshold_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_video_BackgroundSubtractorMOG2_setBackgroundRatio_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_video_BackgroundSubtractorMOG2_setComplexityReductionThreshold_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_video_BackgroundSubtractorMOG2_setDetectShadows_10(JNIEnv*, jclass, jlong, jboolean);
extern void Java_org_opencv_video_BackgroundSubtractorMOG2_setHistory_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_video_BackgroundSubtractorMOG2_setNMixtures_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_video_BackgroundSubtractorMOG2_setShadowThreshold_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_video_BackgroundSubtractorMOG2_setShadowValue_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_video_BackgroundSubtractorMOG2_setVarInit_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_video_BackgroundSubtractorMOG2_setVarMax_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_video_BackgroundSubtractorMOG2_setVarMin_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_video_BackgroundSubtractorMOG2_setVarThresholdGen_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_video_BackgroundSubtractorMOG2_setVarThreshold_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_video_DenseOpticalFlow_calc_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_video_DenseOpticalFlow_collectGarbage_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_video_DenseOpticalFlow_delete(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_video_DualTVL1OpticalFlow_create_10(JNIEnv*, jclass, jdouble, jdouble, jdouble, jint, jint, jdouble, jint, jint, jdouble, jdouble, jint, jboolean);
extern jlong Java_org_opencv_video_DualTVL1OpticalFlow_create_11(JNIEnv*, jclass);
extern void Java_org_opencv_video_DualTVL1OpticalFlow_delete(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_video_DualTVL1OpticalFlow_getEpsilon_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_video_DualTVL1OpticalFlow_getGamma_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_video_DualTVL1OpticalFlow_getInnerIterations_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_video_DualTVL1OpticalFlow_getLambda_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_video_DualTVL1OpticalFlow_getMedianFiltering_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_video_DualTVL1OpticalFlow_getOuterIterations_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_video_DualTVL1OpticalFlow_getScaleStep_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_video_DualTVL1OpticalFlow_getScalesNumber_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_video_DualTVL1OpticalFlow_getTau_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_video_DualTVL1OpticalFlow_getTheta_10(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_video_DualTVL1OpticalFlow_getUseInitialFlow_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_video_DualTVL1OpticalFlow_getWarpingsNumber_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_video_DualTVL1OpticalFlow_setEpsilon_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_video_DualTVL1OpticalFlow_setGamma_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_video_DualTVL1OpticalFlow_setInnerIterations_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_video_DualTVL1OpticalFlow_setLambda_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_video_DualTVL1OpticalFlow_setMedianFiltering_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_video_DualTVL1OpticalFlow_setOuterIterations_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_video_DualTVL1OpticalFlow_setScaleStep_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_video_DualTVL1OpticalFlow_setScalesNumber_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_video_DualTVL1OpticalFlow_setTau_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_video_DualTVL1OpticalFlow_setTheta_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_video_DualTVL1OpticalFlow_setUseInitialFlow_10(JNIEnv*, jclass, jlong, jboolean);
extern void Java_org_opencv_video_DualTVL1OpticalFlow_setWarpingsNumber_10(JNIEnv*, jclass, jlong, jint);
extern jlong Java_org_opencv_video_FarnebackOpticalFlow_create_10(JNIEnv*, jclass, jint, jdouble, jboolean, jint, jint, jint, jdouble, jint);
extern jlong Java_org_opencv_video_FarnebackOpticalFlow_create_11(JNIEnv*, jclass);
extern void Java_org_opencv_video_FarnebackOpticalFlow_delete(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_video_FarnebackOpticalFlow_getFastPyramids_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_video_FarnebackOpticalFlow_getFlags_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_video_FarnebackOpticalFlow_getNumIters_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_video_FarnebackOpticalFlow_getNumLevels_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_video_FarnebackOpticalFlow_getPolyN_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_video_FarnebackOpticalFlow_getPolySigma_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_video_FarnebackOpticalFlow_getPyrScale_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_video_FarnebackOpticalFlow_getWinSize_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_video_FarnebackOpticalFlow_setFastPyramids_10(JNIEnv*, jclass, jlong, jboolean);
extern void Java_org_opencv_video_FarnebackOpticalFlow_setFlags_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_video_FarnebackOpticalFlow_setNumIters_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_video_FarnebackOpticalFlow_setNumLevels_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_video_FarnebackOpticalFlow_setPolyN_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_video_FarnebackOpticalFlow_setPolySigma_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_video_FarnebackOpticalFlow_setPyrScale_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_video_FarnebackOpticalFlow_setWinSize_10(JNIEnv*, jclass, jlong, jint);
extern jlong Java_org_opencv_video_KalmanFilter_KalmanFilter_10(JNIEnv*, jclass, jint, jint, jint, jint);
extern jlong Java_org_opencv_video_KalmanFilter_KalmanFilter_11(JNIEnv*, jclass, jint, jint);
extern jlong Java_org_opencv_video_KalmanFilter_KalmanFilter_12(JNIEnv*, jclass);
extern jlong Java_org_opencv_video_KalmanFilter_correct_10(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_video_KalmanFilter_delete(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_video_KalmanFilter_get_1controlMatrix_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_video_KalmanFilter_get_1errorCovPost_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_video_KalmanFilter_get_1errorCovPre_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_video_KalmanFilter_get_1gain_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_video_KalmanFilter_get_1measurementMatrix_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_video_KalmanFilter_get_1measurementNoiseCov_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_video_KalmanFilter_get_1processNoiseCov_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_video_KalmanFilter_get_1statePost_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_video_KalmanFilter_get_1statePre_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_video_KalmanFilter_get_1transitionMatrix_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_video_KalmanFilter_predict_10(JNIEnv*, jclass, jlong, jlong);
extern jlong Java_org_opencv_video_KalmanFilter_predict_11(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_video_KalmanFilter_set_1controlMatrix_10(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_video_KalmanFilter_set_1errorCovPost_10(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_video_KalmanFilter_set_1errorCovPre_10(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_video_KalmanFilter_set_1gain_10(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_video_KalmanFilter_set_1measurementMatrix_10(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_video_KalmanFilter_set_1measurementNoiseCov_10(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_video_KalmanFilter_set_1processNoiseCov_10(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_video_KalmanFilter_set_1statePost_10(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_video_KalmanFilter_set_1statePre_10(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_video_KalmanFilter_set_1transitionMatrix_10(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_video_SparseOpticalFlow_calc_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_video_SparseOpticalFlow_calc_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_video_SparseOpticalFlow_delete(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_video_SparsePyrLKOpticalFlow_create_10(JNIEnv*, jclass, jdouble, jdouble, jint, jint, jint, jdouble, jint, jdouble);
extern jlong Java_org_opencv_video_SparsePyrLKOpticalFlow_create_11(JNIEnv*, jclass);
extern void Java_org_opencv_video_SparsePyrLKOpticalFlow_delete(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_video_SparsePyrLKOpticalFlow_getFlags_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_video_SparsePyrLKOpticalFlow_getMaxLevel_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_video_SparsePyrLKOpticalFlow_getMinEigThreshold_10(JNIEnv*, jclass, jlong);
extern jdoubleArray Java_org_opencv_video_SparsePyrLKOpticalFlow_getTermCriteria_10(JNIEnv*, jclass, jlong);
extern jdoubleArray Java_org_opencv_video_SparsePyrLKOpticalFlow_getWinSize_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_video_SparsePyrLKOpticalFlow_setFlags_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_video_SparsePyrLKOpticalFlow_setMaxLevel_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_video_SparsePyrLKOpticalFlow_setMinEigThreshold_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_video_SparsePyrLKOpticalFlow_setTermCriteria_10(JNIEnv*, jclass, jlong, jint, jint, jdouble);
extern void Java_org_opencv_video_SparsePyrLKOpticalFlow_setWinSize_10(JNIEnv*, jclass, jlong, jdouble, jdouble);
extern jdoubleArray Java_org_opencv_video_Video_CamShift_10(JNIEnv*, jclass, jlong, jint, jint, jint, jint, jdoubleArray, jint, jint, jdouble);
extern jint Java_org_opencv_video_Video_buildOpticalFlowPyramid_10(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jint, jboolean, jint, jint, jboolean);
extern jint Java_org_opencv_video_Video_buildOpticalFlowPyramid_11(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jint);
extern void Java_org_opencv_video_Video_calcOpticalFlowFarneback_10(JNIEnv*, jclass, jlong, jlong, jlong, jdouble, jint, jint, jint, jint, jdouble, jint);
extern void Java_org_opencv_video_Video_calcOpticalFlowPyrLK_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong, jdouble, jdouble, jint, jint, jint, jdouble, jint, jdouble);
extern void Java_org_opencv_video_Video_calcOpticalFlowPyrLK_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong, jdouble, jdouble, jint);
extern void Java_org_opencv_video_Video_calcOpticalFlowPyrLK_12(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong);
extern jlong Java_org_opencv_video_Video_createBackgroundSubtractorKNN_10(JNIEnv*, jclass, jint, jdouble, jboolean);
extern jlong Java_org_opencv_video_Video_createBackgroundSubtractorKNN_11(JNIEnv*, jclass);
extern jlong Java_org_opencv_video_Video_createBackgroundSubtractorMOG2_10(JNIEnv*, jclass, jint, jdouble, jboolean);
extern jlong Java_org_opencv_video_Video_createBackgroundSubtractorMOG2_11(JNIEnv*, jclass);
extern jlong Java_org_opencv_video_Video_createOptFlow_1DualTVL1_10(JNIEnv*, jclass);
extern jlong Java_org_opencv_video_Video_estimateRigidTransform_10(JNIEnv*, jclass, jlong, jlong, jboolean);
extern jdouble Java_org_opencv_video_Video_findTransformECC_10(JNIEnv*, jclass, jlong, jlong, jlong, jint, jint, jint, jdouble, jlong);
extern jdouble Java_org_opencv_video_Video_findTransformECC_11(JNIEnv*, jclass, jlong, jlong, jlong, jint);
extern jdouble Java_org_opencv_video_Video_findTransformECC_12(JNIEnv*, jclass, jlong, jlong, jlong);
extern jint Java_org_opencv_video_Video_meanShift_10(JNIEnv*, jclass, jlong, jint, jint, jint, jint, jdoubleArray, jint, jint, jdouble);

*/
import "C"

func BackgroundSubtractorNative_apply_0(nativeObj int64, image_nativeObj int64, fgmask_nativeObj int64, learningRate float64) {
	C.Java_org_opencv_video_BackgroundSubtractor_apply_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(image_nativeObj), (C.jlong)(fgmask_nativeObj), (C.jdouble)(learningRate))
}
func BackgroundSubtractorNative_apply_1(nativeObj int64, image_nativeObj int64, fgmask_nativeObj int64) {
	C.Java_org_opencv_video_BackgroundSubtractor_apply_11(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(image_nativeObj), (C.jlong)(fgmask_nativeObj))
}
func BackgroundSubtractorNative_delete(nativeObj int64) {
	C.Java_org_opencv_video_BackgroundSubtractor_delete(clzEnv, clzObj, (C.jlong)(nativeObj))
}

func BackgroundSubtractorNative_getBackgroundImage_0(nativeObj int64, backgroundImage_nativeObj int64) {
	C.Java_org_opencv_video_BackgroundSubtractor_getBackgroundImage_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(backgroundImage_nativeObj))
}

func BackgroundSubtractorKNNNative_delete(nativeObj int64) {
	C.Java_org_opencv_video_BackgroundSubtractorKNN_delete(clzEnv, clzObj, (C.jlong)(nativeObj))
}

func BackgroundSubtractorKNNNative_getDetectShadows_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_video_BackgroundSubtractorKNN_getDetectShadows_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func BackgroundSubtractorKNNNative_getDist2Threshold_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_video_BackgroundSubtractorKNN_getDist2Threshold_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func BackgroundSubtractorKNNNative_getHistory_0(nativeObj int64) int {
	return int(C.Java_org_opencv_video_BackgroundSubtractorKNN_getHistory_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func BackgroundSubtractorKNNNative_getNSamples_0(nativeObj int64) int {
	return int(C.Java_org_opencv_video_BackgroundSubtractorKNN_getNSamples_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func BackgroundSubtractorKNNNative_getShadowThreshold_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_video_BackgroundSubtractorKNN_getShadowThreshold_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func BackgroundSubtractorKNNNative_getShadowValue_0(nativeObj int64) int {
	return int(C.Java_org_opencv_video_BackgroundSubtractorKNN_getShadowValue_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func BackgroundSubtractorKNNNative_getkNNSamples_0(nativeObj int64) int {
	return int(C.Java_org_opencv_video_BackgroundSubtractorKNN_getkNNSamples_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func BackgroundSubtractorKNNNative_setDetectShadows_0(nativeObj int64, detectShadows bool) {
	C.Java_org_opencv_video_BackgroundSubtractorKNN_setDetectShadows_10(clzEnv, clzObj, (C.jlong)(nativeObj), tojboolean(detectShadows))
}

func BackgroundSubtractorKNNNative_setDist2Threshold_0(nativeObj int64, _dist2Threshold float64) {
	C.Java_org_opencv_video_BackgroundSubtractorKNN_setDist2Threshold_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jdouble)(_dist2Threshold))
}

func BackgroundSubtractorKNNNative_setHistory_0(nativeObj int64, history int) {
	C.Java_org_opencv_video_BackgroundSubtractorKNN_setHistory_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(history))
}

func BackgroundSubtractorKNNNative_setNSamples_0(nativeObj int64, _nN int) {
	C.Java_org_opencv_video_BackgroundSubtractorKNN_setNSamples_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(_nN))
}

func BackgroundSubtractorKNNNative_setShadowThreshold_0(nativeObj int64, threshold float64) {
	C.Java_org_opencv_video_BackgroundSubtractorKNN_setShadowThreshold_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jdouble)(threshold))
}

func BackgroundSubtractorKNNNative_setShadowValue_0(nativeObj int64, value int) {
	C.Java_org_opencv_video_BackgroundSubtractorKNN_setShadowValue_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(value))
}

func BackgroundSubtractorKNNNative_setkNNSamples_0(nativeObj int64, _nkNN int) {
	C.Java_org_opencv_video_BackgroundSubtractorKNN_setkNNSamples_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(_nkNN))
}

func BackgroundSubtractorMOG2Native_apply_0(nativeObj int64, image_nativeObj int64, fgmask_nativeObj int64, learningRate float64) {
	C.Java_org_opencv_video_BackgroundSubtractorMOG2_apply_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(image_nativeObj), (C.jlong)(fgmask_nativeObj), (C.jdouble)(learningRate))
}

func BackgroundSubtractorMOG2Native_apply_1(nativeObj int64, image_nativeObj int64, fgmask_nativeObj int64) {
	C.Java_org_opencv_video_BackgroundSubtractorMOG2_apply_11(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(image_nativeObj), (C.jlong)(fgmask_nativeObj))
}

func BackgroundSubtractorMOG2Native_delete(nativeObj int64) {
	C.Java_org_opencv_video_BackgroundSubtractorMOG2_delete(clzEnv, clzObj, (C.jlong)(nativeObj))
}

func BackgroundSubtractorMOG2Native_getBackgroundRatio_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_video_BackgroundSubtractorMOG2_getBackgroundRatio_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func BackgroundSubtractorMOG2Native_getComplexityReductionThreshold_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_video_BackgroundSubtractorMOG2_getComplexityReductionThreshold_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func BackgroundSubtractorMOG2Native_getDetectShadows_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_video_BackgroundSubtractorMOG2_getDetectShadows_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func BackgroundSubtractorMOG2Native_getHistory_0(nativeObj int64) int {
	return int(C.Java_org_opencv_video_BackgroundSubtractorMOG2_getHistory_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func BackgroundSubtractorMOG2Native_getNMixtures_0(nativeObj int64) int {
	return int(C.Java_org_opencv_video_BackgroundSubtractorMOG2_getNMixtures_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func BackgroundSubtractorMOG2Native_getShadowThreshold_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_video_BackgroundSubtractorMOG2_getShadowThreshold_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func BackgroundSubtractorMOG2Native_getShadowValue_0(nativeObj int64) int {
	return int(C.Java_org_opencv_video_BackgroundSubtractorMOG2_getShadowValue_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func BackgroundSubtractorMOG2Native_getVarInit_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_video_BackgroundSubtractorMOG2_getVarInit_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func BackgroundSubtractorMOG2Native_getVarMax_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_video_BackgroundSubtractorMOG2_getVarMax_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func BackgroundSubtractorMOG2Native_getVarMin_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_video_BackgroundSubtractorMOG2_getVarMin_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func BackgroundSubtractorMOG2Native_getVarThresholdGen_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_video_BackgroundSubtractorMOG2_getVarThresholdGen_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func BackgroundSubtractorMOG2Native_getVarThreshold_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_video_BackgroundSubtractorMOG2_getVarThreshold_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func BackgroundSubtractorMOG2Native_setBackgroundRatio_0(nativeObj int64, ratio float64) {
	C.Java_org_opencv_video_BackgroundSubtractorMOG2_setBackgroundRatio_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jdouble)(ratio))
}

func BackgroundSubtractorMOG2Native_setComplexityReductionThreshold_0(nativeObj int64, ct float64) {
	C.Java_org_opencv_video_BackgroundSubtractorMOG2_setComplexityReductionThreshold_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jdouble)(ct))
}

func BackgroundSubtractorMOG2Native_setDetectShadows_0(nativeObj int64, detectShadows bool) {
	C.Java_org_opencv_video_BackgroundSubtractorMOG2_setDetectShadows_10(clzEnv, clzObj, (C.jlong)(nativeObj), tojboolean(detectShadows))
}

func BackgroundSubtractorMOG2Native_setHistory_0(nativeObj int64, history int) {
	C.Java_org_opencv_video_BackgroundSubtractorMOG2_setHistory_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(history))
}

func BackgroundSubtractorMOG2Native_setNMixtures_0(nativeObj int64, nmixtures int) {
	C.Java_org_opencv_video_BackgroundSubtractorMOG2_setNMixtures_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(nmixtures))
}

func BackgroundSubtractorMOG2Native_setShadowThreshold_0(nativeObj int64, threshold float64) {
	C.Java_org_opencv_video_BackgroundSubtractorMOG2_setShadowThreshold_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jdouble)(threshold))
}

func BackgroundSubtractorMOG2Native_setShadowValue_0(nativeObj int64, value int) {
	C.Java_org_opencv_video_BackgroundSubtractorMOG2_setShadowValue_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(value))
}

func BackgroundSubtractorMOG2Native_setVarInit_0(nativeObj int64, varInit float64) {
	C.Java_org_opencv_video_BackgroundSubtractorMOG2_setVarInit_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jdouble)(varInit))
}

func BackgroundSubtractorMOG2Native_setVarMax_0(nativeObj int64, varMax float64) {
	C.Java_org_opencv_video_BackgroundSubtractorMOG2_setVarMax_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jdouble)(varMax))
}

func BackgroundSubtractorMOG2Native_setVarMin_0(nativeObj int64, varMin float64) {
	C.Java_org_opencv_video_BackgroundSubtractorMOG2_setVarMin_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jdouble)(varMin))
}

func BackgroundSubtractorMOG2Native_setVarThresholdGen_0(nativeObj int64, varThresholdGen float64) {
	C.Java_org_opencv_video_BackgroundSubtractorMOG2_setVarThresholdGen_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jdouble)(varThresholdGen))
}

func BackgroundSubtractorMOG2Native_setVarThreshold_0(nativeObj int64, varThreshold float64) {
	C.Java_org_opencv_video_BackgroundSubtractorMOG2_setVarThreshold_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jdouble)(varThreshold))
}

func DenseOpticalFlowNative_calc_0(nativeObj int64, I0_nativeObj int64, I1_nativeObj int64, flow_nativeObj int64) {
	C.Java_org_opencv_video_DenseOpticalFlow_calc_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(I0_nativeObj), (C.jlong)(I1_nativeObj), (C.jlong)(
		flow_nativeObj))
}

func DenseOpticalFlowNative_collectGarbage_0(nativeObj int64) {
	C.Java_org_opencv_video_DenseOpticalFlow_collectGarbage_10(clzEnv, clzObj, (C.jlong)(nativeObj))
}

func DenseOpticalFlowNative_delete(nativeObj int64) {
	C.Java_org_opencv_video_DenseOpticalFlow_delete(clzEnv, clzObj, (C.jlong)(nativeObj))
}

func DualTVL1OpticalFlowNative_create_0(tau float64, lambda float64, theta float64, nscales int, warps int, epsilon float64, innnerIterations int, outerIterations int, scaleStep float64, gamma float64, medianFiltering int, useInitialFlow bool) int64 {
	return int64(C.Java_org_opencv_video_DualTVL1OpticalFlow_create_10(clzEnv, clzObj, (C.jdouble)(tau), (C.jdouble)(lambda), (C.jdouble)(theta), (C.jint)(nscales), (C.jint)(warps), (C.jdouble)(epsilon), (C.jint)(innnerIterations), (C.jint)(outerIterations), (C.jdouble)(scaleStep), (C.jdouble)(gamma), (C.jint)(medianFiltering), tojboolean(useInitialFlow)))
}

func DualTVL1OpticalFlowNative_create_1() int64 {
	return int64(C.Java_org_opencv_video_DualTVL1OpticalFlow_create_11(clzEnv, clzObj))
}

func DualTVL1OpticalFlowNative_delete(nativeObj int64) {
	C.Java_org_opencv_video_DualTVL1OpticalFlow_delete(clzEnv, clzObj, (C.jlong)(nativeObj))
}

func DualTVL1OpticalFlowNative_getEpsilon_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_video_DualTVL1OpticalFlow_getEpsilon_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func DualTVL1OpticalFlowNative_getGamma_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_video_DualTVL1OpticalFlow_getGamma_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func DualTVL1OpticalFlowNative_getInnerIterations_0(nativeObj int64) int {
	return int(C.Java_org_opencv_video_DualTVL1OpticalFlow_getInnerIterations_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func DualTVL1OpticalFlowNative_getLambda_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_video_DualTVL1OpticalFlow_getLambda_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func DualTVL1OpticalFlowNative_getMedianFiltering_0(nativeObj int64) int {
	return int(C.Java_org_opencv_video_DualTVL1OpticalFlow_getMedianFiltering_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func DualTVL1OpticalFlowNative_getOuterIterations_0(nativeObj int64) int {
	return int(C.Java_org_opencv_video_DualTVL1OpticalFlow_getOuterIterations_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func DualTVL1OpticalFlowNative_getScaleStep_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_video_DualTVL1OpticalFlow_getScaleStep_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func DualTVL1OpticalFlowNative_getScalesNumber_0(nativeObj int64) int {
	return int(C.Java_org_opencv_video_DualTVL1OpticalFlow_getScalesNumber_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func DualTVL1OpticalFlowNative_getTau_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_video_DualTVL1OpticalFlow_getTau_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func DualTVL1OpticalFlowNative_getTheta_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_video_DualTVL1OpticalFlow_getTheta_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func DualTVL1OpticalFlowNative_getUseInitialFlow_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_video_DualTVL1OpticalFlow_getUseInitialFlow_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func DualTVL1OpticalFlowNative_getWarpingsNumber_0(nativeObj int64) int {
	return int(C.Java_org_opencv_video_DualTVL1OpticalFlow_getWarpingsNumber_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func DualTVL1OpticalFlowNative_setEpsilon_0(nativeObj int64, val float64) {
	C.Java_org_opencv_video_DualTVL1OpticalFlow_setEpsilon_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jdouble)(val))
}

func DualTVL1OpticalFlowNative_setGamma_0(nativeObj int64, val float64) {
	C.Java_org_opencv_video_DualTVL1OpticalFlow_setGamma_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jdouble)(val))
}

func DualTVL1OpticalFlowNative_setInnerIterations_0(nativeObj int64, val int) {
	C.Java_org_opencv_video_DualTVL1OpticalFlow_setInnerIterations_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(val))
}

func DualTVL1OpticalFlowNative_setLambda_0(nativeObj int64, val float64) {
	C.Java_org_opencv_video_DualTVL1OpticalFlow_setLambda_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jdouble)(val))
}

func DualTVL1OpticalFlowNative_setMedianFiltering_0(nativeObj int64, val int) {
	C.Java_org_opencv_video_DualTVL1OpticalFlow_setMedianFiltering_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(val))
}

func DualTVL1OpticalFlowNative_setOuterIterations_0(nativeObj int64, val int) {
	C.Java_org_opencv_video_DualTVL1OpticalFlow_setOuterIterations_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(val))
}

func DualTVL1OpticalFlowNative_setScaleStep_0(nativeObj int64, val float64) {
	C.Java_org_opencv_video_DualTVL1OpticalFlow_setScaleStep_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jdouble)(val))
}

func DualTVL1OpticalFlowNative_setScalesNumber_0(nativeObj int64, val int) {
	C.Java_org_opencv_video_DualTVL1OpticalFlow_setScalesNumber_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(val))
}

func DualTVL1OpticalFlowNative_setTau_0(nativeObj int64, val float64) {
	C.Java_org_opencv_video_DualTVL1OpticalFlow_setTau_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jdouble)(val))
}

func DualTVL1OpticalFlowNative_setTheta_0(nativeObj int64, val float64) {
	C.Java_org_opencv_video_DualTVL1OpticalFlow_setTheta_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jdouble)(val))
}

func DualTVL1OpticalFlowNative_setUseInitialFlow_0(nativeObj int64, val bool) {
	C.Java_org_opencv_video_DualTVL1OpticalFlow_setUseInitialFlow_10(clzEnv, clzObj, (C.jlong)(nativeObj), tojboolean(val))
}

func DualTVL1OpticalFlowNative_setWarpingsNumber_0(nativeObj int64, val int) {
	C.Java_org_opencv_video_DualTVL1OpticalFlow_setWarpingsNumber_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(val))
}

func FarnebackOpticalFlowNative_create_0(numLevels int, pyrScale float64, fastPyramids bool, winSize int, numIters int, polyN int, polySigma float64, flags int) int64 {
	return int64(C.Java_org_opencv_video_FarnebackOpticalFlow_create_10(clzEnv, clzObj, (C.jint)(numLevels), (C.jdouble)(pyrScale), tojboolean(fastPyramids), (C.jint)(winSize), (C.jint)(numIters), (C.jint)(polyN), (C.jdouble)(polySigma), (C.jint)(flags)))
}

func FarnebackOpticalFlowNative_create_1() int64 {
	return int64(C.Java_org_opencv_video_FarnebackOpticalFlow_create_11(clzEnv, clzObj))
}

func FarnebackOpticalFlowNative_delete(nativeObj int64) {
	C.Java_org_opencv_video_FarnebackOpticalFlow_delete(clzEnv, clzObj, (C.jlong)(nativeObj))
}

func FarnebackOpticalFlowNative_getFastPyramids_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_video_FarnebackOpticalFlow_getFastPyramids_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func FarnebackOpticalFlowNative_getFlags_0(nativeObj int64) int {
	return int(C.Java_org_opencv_video_FarnebackOpticalFlow_getFlags_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func FarnebackOpticalFlowNative_getNumIters_0(nativeObj int64) int {
	return int(C.Java_org_opencv_video_FarnebackOpticalFlow_getNumIters_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func FarnebackOpticalFlowNative_getNumLevels_0(nativeObj int64) int {
	return int(C.Java_org_opencv_video_FarnebackOpticalFlow_getNumLevels_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func FarnebackOpticalFlowNative_getPolyN_0(nativeObj int64) int {
	return int(C.Java_org_opencv_video_FarnebackOpticalFlow_getPolyN_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func FarnebackOpticalFlowNative_getPolySigma_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_video_FarnebackOpticalFlow_getPolySigma_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func FarnebackOpticalFlowNative_getPyrScale_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_video_FarnebackOpticalFlow_getPyrScale_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func FarnebackOpticalFlowNative_getWinSize_0(nativeObj int64) int {
	return int(C.Java_org_opencv_video_FarnebackOpticalFlow_getWinSize_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func FarnebackOpticalFlowNative_setFastPyramids_0(nativeObj int64, fastPyramids bool) {
	C.Java_org_opencv_video_FarnebackOpticalFlow_setFastPyramids_10(clzEnv, clzObj, (C.jlong)(nativeObj), tojboolean(fastPyramids))
}

func FarnebackOpticalFlowNative_setFlags_0(nativeObj int64, flags int) {
	C.Java_org_opencv_video_FarnebackOpticalFlow_setFlags_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(flags))
}

func FarnebackOpticalFlowNative_setNumIters_0(nativeObj int64, numIters int) {
	C.Java_org_opencv_video_FarnebackOpticalFlow_setNumIters_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(numIters))
}

func FarnebackOpticalFlowNative_setNumLevels_0(nativeObj int64, numLevels int) {
	C.Java_org_opencv_video_FarnebackOpticalFlow_setNumLevels_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(numLevels))
}

func FarnebackOpticalFlowNative_setPolyN_0(nativeObj int64, polyN int) {
	C.Java_org_opencv_video_FarnebackOpticalFlow_setPolyN_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(polyN))
}

func FarnebackOpticalFlowNative_setPolySigma_0(nativeObj int64, polySigma float64) {
	C.Java_org_opencv_video_FarnebackOpticalFlow_setPolySigma_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jdouble)(polySigma))
}

func FarnebackOpticalFlowNative_setPyrScale_0(nativeObj int64, pyrScale float64) {
	C.Java_org_opencv_video_FarnebackOpticalFlow_setPyrScale_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jdouble)(pyrScale))
}

func FarnebackOpticalFlowNative_setWinSize_0(nativeObj int64, winSize int) {
	C.Java_org_opencv_video_FarnebackOpticalFlow_setWinSize_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(winSize))
}

func KalmanFilterNative_KalmanFilter_0(dynamParams int, measureParams int, controlParams int, rtype int) int64 {
	return int64(C.Java_org_opencv_video_KalmanFilter_KalmanFilter_10(clzEnv, clzObj, (C.jint)(dynamParams), (C.jint)(measureParams), (C.jint)(controlParams), (C.jint)(rtype)))
}

func KalmanFilterNative_KalmanFilter_1(dynamParams int, measureParams int) int64 {
	return int64(C.Java_org_opencv_video_KalmanFilter_KalmanFilter_11(clzEnv, clzObj, (C.jint)(dynamParams), (C.jint)(measureParams)))
}

func KalmanFilterNative_KalmanFilter_2() int64 {
	return int64(C.Java_org_opencv_video_KalmanFilter_KalmanFilter_12(clzEnv, clzObj))
}

func KalmanFilterNative_correct_0(nativeObj int64, measurement_nativeObj int64) int64 {
	return int64(C.Java_org_opencv_video_KalmanFilter_correct_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(measurement_nativeObj)))
}

func KalmanFilterNative_delete(nativeObj int64) {
	C.Java_org_opencv_video_KalmanFilter_delete(clzEnv, clzObj, (C.jlong)(nativeObj))
}

func KalmanFilterNative_get_controlMatrix_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_video_KalmanFilter_get_1controlMatrix_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func KalmanFilterNative_get_errorCovPost_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_video_KalmanFilter_get_1errorCovPost_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func KalmanFilterNative_get_errorCovPre_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_video_KalmanFilter_get_1errorCovPre_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func KalmanFilterNative_get_gain_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_video_KalmanFilter_get_1gain_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func KalmanFilterNative_get_measurementMatrix_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_video_KalmanFilter_get_1measurementMatrix_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func KalmanFilterNative_get_measurementNoiseCov_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_video_KalmanFilter_get_1measurementNoiseCov_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func KalmanFilterNative_get_processNoiseCov_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_video_KalmanFilter_get_1processNoiseCov_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func KalmanFilterNative_get_statePost_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_video_KalmanFilter_get_1statePost_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func KalmanFilterNative_get_statePre_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_video_KalmanFilter_get_1statePre_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func KalmanFilterNative_get_transitionMatrix_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_video_KalmanFilter_get_1transitionMatrix_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func KalmanFilterNative_predict_0(nativeObj int64, control_nativeObj int64) int64 {
	return int64(C.Java_org_opencv_video_KalmanFilter_predict_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(control_nativeObj)))
}

func KalmanFilterNative_predict_1(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_video_KalmanFilter_predict_11(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func KalmanFilterNative_set_controlMatrix_0(nativeObj int64, controlMatrix_nativeObj int64) {
	C.Java_org_opencv_video_KalmanFilter_set_1controlMatrix_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(controlMatrix_nativeObj))
}

func KalmanFilterNative_set_errorCovPost_0(nativeObj int64, errorCovPost_nativeObj int64) {
	C.Java_org_opencv_video_KalmanFilter_set_1errorCovPost_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(errorCovPost_nativeObj))
}

func KalmanFilterNative_set_errorCovPre_0(nativeObj int64, errorCovPre_nativeObj int64) {
	C.Java_org_opencv_video_KalmanFilter_set_1errorCovPre_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(errorCovPre_nativeObj))
}

func KalmanFilterNative_set_gain_0(nativeObj int64, gain_nativeObj int64) {
	C.Java_org_opencv_video_KalmanFilter_set_1gain_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(gain_nativeObj))
}

func KalmanFilterNative_set_measurementMatrix_0(nativeObj int64, measurementMatrix_nativeObj int64) {
	C.Java_org_opencv_video_KalmanFilter_set_1measurementMatrix_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(measurementMatrix_nativeObj))
}

func KalmanFilterNative_set_measurementNoiseCov_0(nativeObj int64, measurementNoiseCov_nativeObj int64) {
	C.Java_org_opencv_video_KalmanFilter_set_1measurementNoiseCov_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(measurementNoiseCov_nativeObj))
}

func KalmanFilterNative_set_processNoiseCov_0(nativeObj int64, processNoiseCov_nativeObj int64) {
	C.Java_org_opencv_video_KalmanFilter_set_1processNoiseCov_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(processNoiseCov_nativeObj))
}

func KalmanFilterNative_set_statePost_0(nativeObj int64, statePost_nativeObj int64) {
	C.Java_org_opencv_video_KalmanFilter_set_1statePost_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(statePost_nativeObj))
}

func KalmanFilterNative_set_statePre_0(nativeObj int64, statePre_nativeObj int64) {
	C.Java_org_opencv_video_KalmanFilter_set_1statePre_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(statePre_nativeObj))
}

func KalmanFilterNative_set_transitionMatrix_0(nativeObj int64, transitionMatrix_nativeObj int64) {
	C.Java_org_opencv_video_KalmanFilter_set_1transitionMatrix_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(transitionMatrix_nativeObj))
}

func SparseOpticalFlowNative_calc_0(nativeObj int64, prevImg_nativeObj int64, nextImg_nativeObj int64, prevPts_nativeObj int64, nextPts_nativeObj int64, status_nativeObj int64, err_nativeObj int64) {
	C.Java_org_opencv_video_SparseOpticalFlow_calc_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(prevImg_nativeObj), (C.jlong)(nextImg_nativeObj), (C.jlong)(prevPts_nativeObj), (C.jlong)(nextPts_nativeObj), (C.jlong)(status_nativeObj), (C.jlong)(err_nativeObj))
}

func SparseOpticalFlowNative_calc_1(nativeObj int64, prevImg_nativeObj int64, nextImg_nativeObj int64, prevPts_nativeObj int64, nextPts_nativeObj int64, status_nativeObj int64) {
	C.Java_org_opencv_video_SparseOpticalFlow_calc_11(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(prevImg_nativeObj), (C.jlong)(nextImg_nativeObj), (C.jlong)(prevPts_nativeObj), (C.jlong)(nextPts_nativeObj), (C.jlong)(status_nativeObj))
}

func SparseOpticalFlowNative_delete(nativeObj int64) {
	C.Java_org_opencv_video_SparseOpticalFlow_delete(clzEnv, clzObj, (C.jlong)(nativeObj))
}

func SparsePyrLKOpticalFlowNative_create_0(winSize_width float64, winSize_height float64, maxLevel int, crit_type int, crit_maxCount int, crit_epsilon float64, flags int, minEigThreshold float64) int64 {
	return int64(C.Java_org_opencv_video_SparsePyrLKOpticalFlow_create_10(clzEnv, clzObj, (C.jdouble)(winSize_width), (C.jdouble)(winSize_height), (C.jint)(maxLevel), (C.jint)(crit_type), (C.jint)(crit_maxCount), (C.jdouble)(crit_epsilon), (C.jint)(flags), (C.jdouble)(minEigThreshold)))
}

func SparsePyrLKOpticalFlowNative_create_1() int64 {
	return int64(C.Java_org_opencv_video_SparsePyrLKOpticalFlow_create_11(clzEnv, clzObj))
}

func SparsePyrLKOpticalFlowNative_delete(nativeObj int64) {
	C.Java_org_opencv_video_SparsePyrLKOpticalFlow_delete(clzEnv, clzObj, (C.jlong)(nativeObj))
}

func SparsePyrLKOpticalFlowNative_getFlags_0(nativeObj int64) int {
	return int(C.Java_org_opencv_video_SparsePyrLKOpticalFlow_getFlags_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func SparsePyrLKOpticalFlowNative_getMaxLevel_0(nativeObj int64) int {
	return int(C.Java_org_opencv_video_SparsePyrLKOpticalFlow_getMaxLevel_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func SparsePyrLKOpticalFlowNative_getMinEigThreshold_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_video_SparsePyrLKOpticalFlow_getMinEigThreshold_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func SparsePyrLKOpticalFlowNative_getTermCriteria_0(nativeObj int64) []float64 {
	return togofloat64Array(C.Java_org_opencv_video_SparsePyrLKOpticalFlow_getTermCriteria_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func SparsePyrLKOpticalFlowNative_getWinSize_0(nativeObj int64) []float64 {
	return togofloat64Array(C.Java_org_opencv_video_SparsePyrLKOpticalFlow_getWinSize_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func SparsePyrLKOpticalFlowNative_setFlags_0(nativeObj int64, flags int) {
	C.Java_org_opencv_video_SparsePyrLKOpticalFlow_setFlags_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(flags))
}

func SparsePyrLKOpticalFlowNative_setMaxLevel_0(nativeObj int64, maxLevel int) {
	C.Java_org_opencv_video_SparsePyrLKOpticalFlow_setMaxLevel_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(maxLevel))
}

func SparsePyrLKOpticalFlowNative_setMinEigThreshold_0(nativeObj int64, minEigThreshold float64) {
	C.Java_org_opencv_video_SparsePyrLKOpticalFlow_setMinEigThreshold_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jdouble)(minEigThreshold))
}

func SparsePyrLKOpticalFlowNative_setTermCriteria_0(nativeObj int64, crit_type int, crit_maxCount int, crit_epsilon float64) {
	C.Java_org_opencv_video_SparsePyrLKOpticalFlow_setTermCriteria_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(crit_type), (C.jint)(crit_maxCount), (C.jdouble)(crit_epsilon))
}

func SparsePyrLKOpticalFlowNative_setWinSize_0(nativeObj int64, winSize_width float64, winSize_height float64) {
	C.Java_org_opencv_video_SparsePyrLKOpticalFlow_setWinSize_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jdouble)(winSize_width), (C.jdouble)(winSize_height))
}

func VideoNative_CamShift_0(probImage_nativeObj int64, window_x int, window_y int, window_width int, window_height int, window_out []float64, criteria_type int, criteria_maxCount int, criteria_epsilon float64) []float64 {
	defer ungointerface(window_out)
	return togofloat64Array(C.Java_org_opencv_video_Video_CamShift_10(clzEnv, clzObj, (C.jlong)(probImage_nativeObj), (C.jint)(window_x), (C.jint)(window_y), (C.jint)(window_width), (C.jint)(window_height), tojdoubleArray(window_out), (C.jint)(criteria_type), (C.jint)(criteria_maxCount), (C.jdouble)(criteria_epsilon)))
}

func VideoNative_buildOpticalFlowPyramid_0(img_nativeObj int64, pyramid_mat_nativeObj int64, winSize_width float64, winSize_height float64, maxLevel int, withDerivatives bool, pyrBorder int, derivBorder int, tryReuseInputImage bool) int {
	return int(C.Java_org_opencv_video_Video_buildOpticalFlowPyramid_10(clzEnv, clzObj, (C.jlong)(img_nativeObj), (C.jlong)(pyramid_mat_nativeObj), (C.jdouble)(winSize_width), (C.jdouble)(winSize_height), (C.jint)(maxLevel), tojboolean(withDerivatives), (C.jint)(pyrBorder), (C.jint)(derivBorder), tojboolean(tryReuseInputImage)))
}

func VideoNative_buildOpticalFlowPyramid_1(img_nativeObj int64, pyramid_mat_nativeObj int64, winSize_width float64, winSize_height float64, maxLevel int) int {
	return int(C.Java_org_opencv_video_Video_buildOpticalFlowPyramid_11(clzEnv, clzObj, (C.jlong)(img_nativeObj), (C.jlong)(pyramid_mat_nativeObj), (C.jdouble)(winSize_width), (C.jdouble)(winSize_height), (C.jint)(maxLevel)))
}

func VideoNative_calcOpticalFlowFarneback_0(prev_nativeObj int64, next_nativeObj int64, flow_nativeObj int64, pyr_scale float64, levels int, winsize int, iterations int, poly_n int, poly_sigma float64, flags int) {
	C.Java_org_opencv_video_Video_calcOpticalFlowFarneback_10(clzEnv, clzObj, (C.jlong)(prev_nativeObj), (C.jlong)(next_nativeObj), (C.jlong)(flow_nativeObj), (C.jdouble)(pyr_scale), (C.jint)(levels), (C.jint)(winsize), (C.jint)(iterations), (C.jint)(poly_n), (C.jdouble)(poly_sigma), (C.jint)(flags))
}

func VideoNative_calcOpticalFlowPyrLK_0(prevImg_nativeObj int64, nextImg_nativeObj int64, prevPts_mat_nativeObj int64, nextPts_mat_nativeObj int64, status_mat_nativeObj int64, err_mat_nativeObj int64, winSize_width float64, winSize_height float64, maxLevel int, criteria_type int, criteria_maxCount int, criteria_epsilon float64, flags int, minEigThreshold float64) {
	C.Java_org_opencv_video_Video_calcOpticalFlowPyrLK_10(clzEnv, clzObj, (C.jlong)(prevImg_nativeObj), (C.jlong)(nextImg_nativeObj), (C.jlong)(prevPts_mat_nativeObj), (C.jlong)(nextPts_mat_nativeObj), (C.jlong)(status_mat_nativeObj), (C.jlong)(err_mat_nativeObj), (C.jdouble)(winSize_width), (C.jdouble)(winSize_height), (C.jint)(maxLevel), (C.jint)(criteria_type), (C.jint)(criteria_maxCount), (C.jdouble)(criteria_epsilon), (C.jint)(flags), (C.jdouble)(minEigThreshold))
}

func VideoNative_calcOpticalFlowPyrLK_1(prevImg_nativeObj int64, nextImg_nativeObj int64, prevPts_mat_nativeObj int64, nextPts_mat_nativeObj int64, status_mat_nativeObj int64, err_mat_nativeObj int64, winSize_width float64, winSize_height float64, maxLevel int) {
	C.Java_org_opencv_video_Video_calcOpticalFlowPyrLK_11(clzEnv, clzObj, (C.jlong)(prevImg_nativeObj), (C.jlong)(nextImg_nativeObj), (C.jlong)(prevPts_mat_nativeObj), (C.jlong)(nextPts_mat_nativeObj), (C.jlong)(status_mat_nativeObj), (C.jlong)(err_mat_nativeObj), (C.jdouble)(winSize_width), (C.jdouble)(winSize_height), (C.jint)(maxLevel))
}

func VideoNative_calcOpticalFlowPyrLK_2(prevImg_nativeObj int64, nextImg_nativeObj int64, prevPts_mat_nativeObj int64, nextPts_mat_nativeObj int64, status_mat_nativeObj int64, err_mat_nativeObj int64) {
	C.Java_org_opencv_video_Video_calcOpticalFlowPyrLK_12(clzEnv, clzObj, (C.jlong)(prevImg_nativeObj), (C.jlong)(nextImg_nativeObj), (C.jlong)(prevPts_mat_nativeObj), (C.jlong)(nextPts_mat_nativeObj), (C.jlong)(status_mat_nativeObj), (C.jlong)(err_mat_nativeObj))
}

func VideoNative_createBackgroundSubtractorKNN_0(history int, dist2Threshold float64, detectShadows bool) int64 {
	return int64(C.Java_org_opencv_video_Video_createBackgroundSubtractorKNN_10(clzEnv, clzObj, (C.jint)(history), (C.jdouble)(dist2Threshold), tojboolean(detectShadows)))
}

func VideoNative_createBackgroundSubtractorKNN_1() int64 {
	return int64(C.Java_org_opencv_video_Video_createBackgroundSubtractorKNN_11(clzEnv, clzObj))
}

func VideoNative_createBackgroundSubtractorMOG2_0(history int, varThreshold float64, detectShadows bool) int64 {
	return int64(C.Java_org_opencv_video_Video_createBackgroundSubtractorMOG2_10(clzEnv, clzObj, (C.jint)(history), (C.jdouble)(varThreshold), tojboolean(detectShadows)))
}

func VideoNative_createBackgroundSubtractorMOG2_1() int64 {
	return int64(C.Java_org_opencv_video_Video_createBackgroundSubtractorMOG2_11(clzEnv, clzObj))
}

func VideoNative_createOptFlow_DualTVL1_0() int64 {
	return int64(C.Java_org_opencv_video_Video_createOptFlow_1DualTVL1_10(clzEnv, clzObj))
}

func VideoNative_estimateRigidTransform_0(src_nativeObj int64, dst_nativeObj int64, fullAffine bool) int64 {
	return int64(C.Java_org_opencv_video_Video_estimateRigidTransform_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), tojboolean(fullAffine)))
}

func VideoNative_findTransformECC_0(templateImage_nativeObj int64, inputImage_nativeObj int64, warpMatrix_nativeObj int64, motionType int, criteria_type int, criteria_maxCount int, criteria_epsilon float64, inputMask_nativeObj int64) float64 {
	return float64(C.Java_org_opencv_video_Video_findTransformECC_10(clzEnv, clzObj, (C.jlong)(templateImage_nativeObj), (C.jlong)(inputImage_nativeObj), (C.jlong)(warpMatrix_nativeObj), (C.jint)(motionType), (C.jint)(criteria_type), (C.jint)(criteria_maxCount), (C.jdouble)(criteria_epsilon), (C.jlong)(inputMask_nativeObj)))
}

func VideoNative_findTransformECC_1(templateImage_nativeObj int64, inputImage_nativeObj int64, warpMatrix_nativeObj int64, motionType int) float64 {
	return float64(C.Java_org_opencv_video_Video_findTransformECC_11(clzEnv, clzObj, (C.jlong)(templateImage_nativeObj), (C.jlong)(inputImage_nativeObj), (C.jlong)(warpMatrix_nativeObj), (C.jint)(motionType)))
}

func VideoNative_findTransformECC_2(templateImage_nativeObj int64, inputImage_nativeObj int64, warpMatrix_nativeObj int64) float64 {
	return float64(C.Java_org_opencv_video_Video_findTransformECC_12(clzEnv, clzObj, (C.jlong)(templateImage_nativeObj), (C.jlong)(inputImage_nativeObj), (C.jlong)(warpMatrix_nativeObj)))
}

func VideoNative_meanShift_0(probImage_nativeObj int64, window_x int, window_y int, window_width int, window_height int, window_out []float64, criteria_type int, criteria_maxCount int, criteria_epsilon float64) int {
	defer ungointerface(window_out)
	return int(C.Java_org_opencv_video_Video_meanShift_10(clzEnv, clzObj, (C.jlong)(probImage_nativeObj), (C.jint)(window_x), (C.jint)(window_y), (C.jint)(window_width), (C.jint)(window_height), tojdoubleArray(window_out), (C.jint)(criteria_type), (C.jint)(criteria_maxCount), (C.jdouble)(criteria_epsilon)))
}
