#include "jni.h"

#ifdef DYNLOAD

#include <stdlib.h>
#if defined(__APPLE__)
#include <dlfcn.h>
#elif defined(_WIN32)
#define WIN32_LEAN_AND_MEAN 1
#include <windows.h>
#else
#include <dlfcn.h>
#endif

#ifdef _WIN32
static HMODULE hmod = NULL;
const char* libpath = "opencv_java331.dll";
#elif !defined __APPLE__
static void* plib = NULL;
const char* libpath = "libopencv_java3.so";
#endif
void setLibPath(const char* path) {
#if !defined __APPLE__
libpath = path;
#endif
}

static int failCnt = 0;
// Helps bind function pointers to c functions.
static void* procLoadSub(const char* name) {
#ifdef __APPLE__
	return dlsym(RTLD_DEFAULT, name);
#elif _WIN32
	if(hmod == NULL) {
		hmod = LoadLibraryA(libpath);
	}
	return GetProcAddress(hmod, (LPCSTR)name);
#else
	if(plib == NULL) {
		plib = dlopen(libpath, RTLD_LAZY);
		if(plib == NULL) return NULL;
	}
	return dlsym(plib, name);
#endif
}

static void* procLoad(const char* name) {
	void* p = procLoadSub(name);
	if (p == NULL) { failCnt++;	}
	return p;
}

static jint (*proc_JNI_OnLoad)(JavaVM*, void*);
jint JNI_OnLoad(JavaVM* a0, void* a1) {
	return proc_JNI_OnLoad(a0, a1);
}

static jlong (*proc_features2d_FastFeatureDetector_create_10)(JNIEnv*,  jclass,  jint,  jboolean,  jint);
jlong Java_org_opencv_features2d_FastFeatureDetector_create_10(JNIEnv* a0, jclass a1, jint a2, jboolean a3, jint a4) {
    return proc_features2d_FastFeatureDetector_create_10 (a0, a1, a2, a3, a4);
}

static jlong (*proc_ml_LogisticRegression_create_10)(JNIEnv*,  jclass);
jlong Java_org_opencv_ml_LogisticRegression_create_10(JNIEnv* a0, jclass a1) {
    return proc_ml_LogisticRegression_create_10 (a0, a1);
}

static void (*proc_calib3d_Calib3d_undistortImage_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jdouble,  jdouble);
void Java_org_opencv_calib3d_Calib3d_undistortImage_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jdouble a7, jdouble a8) {
    proc_calib3d_Calib3d_undistortImage_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static void (*proc_core_Core_add_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jint);
void Java_org_opencv_core_Core_add_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jint a6) {
    proc_core_Core_add_10 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_imgproc_Imgproc_demosaicing_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint);
void Java_org_opencv_imgproc_Imgproc_demosaicing_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5) {
    proc_imgproc_Imgproc_demosaicing_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_photo_Photo_stylization_10)(JNIEnv*,  jclass,  jlong,  jlong,  jfloat,  jfloat);
void Java_org_opencv_photo_Photo_stylization_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jfloat a4, jfloat a5) {
    proc_photo_Photo_stylization_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_video_FarnebackOpticalFlow_setPyrScale_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_video_FarnebackOpticalFlow_setPyrScale_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_video_FarnebackOpticalFlow_setPyrScale_10 (a0, a1, a2, a3);
}

static jlong (*proc_features2d_FastFeatureDetector_create_11)(JNIEnv*,  jclass);
jlong Java_org_opencv_features2d_FastFeatureDetector_create_11(JNIEnv* a0, jclass a1) {
    return proc_features2d_FastFeatureDetector_create_11 (a0, a1);
}

static void (*proc_calib3d_Calib3d_undistortImage_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_calib3d_Calib3d_undistortImage_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_calib3d_Calib3d_undistortImage_11 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_core_Core_add_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_core_Core_add_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_core_Core_add_11 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_imgproc_Imgproc_demosaicing_11)(JNIEnv*,  jclass,  jlong,  jlong,  jint);
void Java_org_opencv_imgproc_Imgproc_demosaicing_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4) {
    proc_imgproc_Imgproc_demosaicing_11 (a0, a1, a2, a3, a4);
}

static void (*proc_ml_LogisticRegression_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_ml_LogisticRegression_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_ml_LogisticRegression_delete (a0, a1, a2);
}

static void (*proc_photo_Photo_stylization_11)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_photo_Photo_stylization_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_photo_Photo_stylization_11 (a0, a1, a2, a3);
}

static void (*proc_video_FarnebackOpticalFlow_setWinSize_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_video_FarnebackOpticalFlow_setWinSize_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_video_FarnebackOpticalFlow_setWinSize_10 (a0, a1, a2, a3);
}

static jint (*proc_ml_LogisticRegression_getIterations_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_ml_LogisticRegression_getIterations_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_LogisticRegression_getIterations_10 (a0, a1, a2);
}

static jlong (*proc_video_KalmanFilter_KalmanFilter_10)(JNIEnv*,  jclass,  jint,  jint,  jint,  jint);
jlong Java_org_opencv_video_KalmanFilter_KalmanFilter_10(JNIEnv* a0, jclass a1, jint a2, jint a3, jint a4, jint a5) {
    return proc_video_KalmanFilter_KalmanFilter_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_calib3d_Calib3d_undistortPoints_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_calib3d_Calib3d_undistortPoints_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7) {
    proc_calib3d_Calib3d_undistortPoints_10 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_core_Core_add_12)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_core_Core_add_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_core_Core_add_12 (a0, a1, a2, a3, a4);
}

static void (*proc_features2d_FastFeatureDetector_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_features2d_FastFeatureDetector_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_features2d_FastFeatureDetector_delete (a0, a1, a2);
}

static void (*proc_imgproc_Imgproc_dilate_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jdouble,  jdouble,  jint,  jint,  jdouble,  jdouble,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_dilate_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jdouble a5, jdouble a6, jint a7, jint a8, jdouble a9, jdouble a10, jdouble a11, jdouble a12) {
    proc_imgproc_Imgproc_dilate_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12);
}

static void (*proc_photo_Photo_textureFlattening_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jfloat,  jfloat,  jint);
void Java_org_opencv_photo_Photo_textureFlattening_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jfloat a5, jfloat a6, jint a7) {
    proc_photo_Photo_textureFlattening_10 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static jdouble (*proc_ml_LogisticRegression_getLearningRate_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_ml_LogisticRegression_getLearningRate_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_LogisticRegression_getLearningRate_10 (a0, a1, a2);
}

static jlong (*proc_video_KalmanFilter_KalmanFilter_11)(JNIEnv*,  jclass,  jint,  jint);
jlong Java_org_opencv_video_KalmanFilter_KalmanFilter_11(JNIEnv* a0, jclass a1, jint a2, jint a3) {
    return proc_video_KalmanFilter_KalmanFilter_11 (a0, a1, a2, a3);
}

static jstring (*proc_features2d_FastFeatureDetector_getDefaultName_10)(JNIEnv*,  jclass,  jlong);
jstring Java_org_opencv_features2d_FastFeatureDetector_getDefaultName_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_FastFeatureDetector_getDefaultName_10 (a0, a1, a2);
}

static void (*proc_calib3d_Calib3d_undistortPoints_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_calib3d_Calib3d_undistortPoints_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_calib3d_Calib3d_undistortPoints_11 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_core_Core_add_13)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jlong,  jlong,  jint);
void Java_org_opencv_core_Core_add_13(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jlong a7, jlong a8, jint a9) {
    proc_core_Core_add_13 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static void (*proc_imgproc_Imgproc_dilate_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jdouble,  jdouble,  jint);
void Java_org_opencv_imgproc_Imgproc_dilate_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jdouble a5, jdouble a6, jint a7) {
    proc_imgproc_Imgproc_dilate_11 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_photo_Photo_textureFlattening_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_photo_Photo_textureFlattening_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_photo_Photo_textureFlattening_11 (a0, a1, a2, a3, a4);
}

static jboolean (*proc_features2d_FastFeatureDetector_getNonmaxSuppression_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_features2d_FastFeatureDetector_getNonmaxSuppression_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_FastFeatureDetector_getNonmaxSuppression_10 (a0, a1, a2);
}

static jint (*proc_ml_LogisticRegression_getMiniBatchSize_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_ml_LogisticRegression_getMiniBatchSize_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_LogisticRegression_getMiniBatchSize_10 (a0, a1, a2);
}

static jlong (*proc_video_KalmanFilter_KalmanFilter_12)(JNIEnv*,  jclass);
jlong Java_org_opencv_video_KalmanFilter_KalmanFilter_12(JNIEnv* a0, jclass a1) {
    return proc_video_KalmanFilter_KalmanFilter_12 (a0, a1);
}

static void (*proc_calib3d_Calib3d_validateDisparity_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint,  jint);
void Java_org_opencv_calib3d_Calib3d_validateDisparity_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5, jint a6) {
    proc_calib3d_Calib3d_validateDisparity_10 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_core_Core_add_14)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jlong,  jlong);
void Java_org_opencv_core_Core_add_14(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jlong a7, jlong a8) {
    proc_core_Core_add_14 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static void (*proc_imgproc_Imgproc_dilate_12)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_imgproc_Imgproc_dilate_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_imgproc_Imgproc_dilate_12 (a0, a1, a2, a3, a4);
}

static void (*proc_photo_Tonemap_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_photo_Tonemap_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_photo_Tonemap_delete (a0, a1, a2);
}

static jfloat (*proc_photo_Tonemap_getGamma_10)(JNIEnv*,  jclass,  jlong);
jfloat Java_org_opencv_photo_Tonemap_getGamma_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_photo_Tonemap_getGamma_10 (a0, a1, a2);
}

static jint (*proc_features2d_FastFeatureDetector_getThreshold_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_features2d_FastFeatureDetector_getThreshold_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_FastFeatureDetector_getThreshold_10 (a0, a1, a2);
}

static jint (*proc_ml_LogisticRegression_getRegularization_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_ml_LogisticRegression_getRegularization_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_LogisticRegression_getRegularization_10 (a0, a1, a2);
}

static jlong (*proc_video_KalmanFilter_correct_10)(JNIEnv*,  jclass,  jlong,  jlong);
jlong Java_org_opencv_video_KalmanFilter_correct_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_video_KalmanFilter_correct_10 (a0, a1, a2, a3);
}

static void (*proc_calib3d_Calib3d_validateDisparity_11)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint);
void Java_org_opencv_calib3d_Calib3d_validateDisparity_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5) {
    proc_calib3d_Calib3d_validateDisparity_11 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_core_Core_add_15)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jlong);
void Java_org_opencv_core_Core_add_15(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jlong a7) {
    proc_core_Core_add_15 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_imgproc_Imgproc_distanceTransformWithLabels_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jint,  jint,  jint);
void Java_org_opencv_imgproc_Imgproc_distanceTransformWithLabels_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jint a5, jint a6, jint a7) {
    proc_imgproc_Imgproc_distanceTransformWithLabels_10 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static jdoubleArray (*proc_ml_LogisticRegression_getTermCriteria_10)(JNIEnv*,  jclass,  jlong);
jdoubleArray Java_org_opencv_ml_LogisticRegression_getTermCriteria_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_LogisticRegression_getTermCriteria_10 (a0, a1, a2);
}

static jint (*proc_features2d_FastFeatureDetector_getType_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_features2d_FastFeatureDetector_getType_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_FastFeatureDetector_getType_10 (a0, a1, a2);
}

static jlong (*proc_calib3d_StereoBM_create_10)(JNIEnv*,  jclass,  jint,  jint);
jlong Java_org_opencv_calib3d_StereoBM_create_10(JNIEnv* a0, jclass a1, jint a2, jint a3) {
    return proc_calib3d_StereoBM_create_10 (a0, a1, a2, a3);
}

static void (*proc_core_Core_batchDistance_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jint,  jlong,  jint,  jint,  jlong,  jint,  jboolean);
void Java_org_opencv_core_Core_batchDistance_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jint a5, jlong a6, jint a7, jint a8, jlong a9, jint a10, jboolean a11) {
    proc_core_Core_batchDistance_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11);
}

static void (*proc_imgproc_Imgproc_distanceTransformWithLabels_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jint,  jint);
void Java_org_opencv_imgproc_Imgproc_distanceTransformWithLabels_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jint a5, jint a6) {
    proc_imgproc_Imgproc_distanceTransformWithLabels_11 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_photo_Tonemap_process_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_photo_Tonemap_process_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_photo_Tonemap_process_10 (a0, a1, a2, a3, a4);
}

static void (*proc_video_KalmanFilter_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_video_KalmanFilter_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_video_KalmanFilter_delete (a0, a1, a2);
}

static jint (*proc_ml_LogisticRegression_getTrainMethod_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_ml_LogisticRegression_getTrainMethod_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_LogisticRegression_getTrainMethod_10 (a0, a1, a2);
}

static jlong (*proc_calib3d_StereoBM_create_11)(JNIEnv*,  jclass);
jlong Java_org_opencv_calib3d_StereoBM_create_11(JNIEnv* a0, jclass a1) {
    return proc_calib3d_StereoBM_create_11 (a0, a1);
}

static jlong (*proc_video_KalmanFilter_get_1controlMatrix_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_video_KalmanFilter_get_1controlMatrix_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_KalmanFilter_get_1controlMatrix_10 (a0, a1, a2);
}

static void (*proc_core_Core_batchDistance_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jint,  jlong,  jint,  jint);
void Java_org_opencv_core_Core_batchDistance_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jint a5, jlong a6, jint a7, jint a8) {
    proc_core_Core_batchDistance_11 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static void (*proc_features2d_FastFeatureDetector_setNonmaxSuppression_10)(JNIEnv*,  jclass,  jlong,  jboolean);
void Java_org_opencv_features2d_FastFeatureDetector_setNonmaxSuppression_10(JNIEnv* a0, jclass a1, jlong a2, jboolean a3) {
    proc_features2d_FastFeatureDetector_setNonmaxSuppression_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_distanceTransform_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint,  jint);
void Java_org_opencv_imgproc_Imgproc_distanceTransform_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5, jint a6) {
    proc_imgproc_Imgproc_distanceTransform_10 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_photo_Tonemap_setGamma_10)(JNIEnv*,  jclass,  jlong,  jfloat);
void Java_org_opencv_photo_Tonemap_setGamma_10(JNIEnv* a0, jclass a1, jlong a2, jfloat a3) {
    proc_photo_Tonemap_setGamma_10 (a0, a1, a2, a3);
}

static jlong (*proc_ml_LogisticRegression_get_1learnt_1thetas_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_ml_LogisticRegression_get_1learnt_1thetas_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_LogisticRegression_get_1learnt_1thetas_10 (a0, a1, a2);
}

static jlong (*proc_video_KalmanFilter_get_1errorCovPost_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_video_KalmanFilter_get_1errorCovPost_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_KalmanFilter_get_1errorCovPost_10 (a0, a1, a2);
}

static void (*proc_calib3d_StereoBM_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_calib3d_StereoBM_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_calib3d_StereoBM_delete (a0, a1, a2);
}

static void (*proc_core_Core_batchDistance_12)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jint,  jlong);
void Java_org_opencv_core_Core_batchDistance_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jint a5, jlong a6) {
    proc_core_Core_batchDistance_12 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_features2d_FastFeatureDetector_setThreshold_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_features2d_FastFeatureDetector_setThreshold_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_features2d_FastFeatureDetector_setThreshold_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_distanceTransform_11)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint);
void Java_org_opencv_imgproc_Imgproc_distanceTransform_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5) {
    proc_imgproc_Imgproc_distanceTransform_11 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_photo_TonemapDrago_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_photo_TonemapDrago_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_photo_TonemapDrago_delete (a0, a1, a2);
}

static jfloat (*proc_photo_TonemapDrago_getBias_10)(JNIEnv*,  jclass,  jlong);
jfloat Java_org_opencv_photo_TonemapDrago_getBias_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_photo_TonemapDrago_getBias_10 (a0, a1, a2);
}

static jint (*proc_calib3d_StereoBM_getPreFilterCap_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_calib3d_StereoBM_getPreFilterCap_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_calib3d_StereoBM_getPreFilterCap_10 (a0, a1, a2);
}

static jlong (*proc_ml_LogisticRegression_load_10)(JNIEnv*,  jclass,  jstring,  jstring);
jlong Java_org_opencv_ml_LogisticRegression_load_10(JNIEnv* a0, jclass a1, jstring a2, jstring a3) {
    return proc_ml_LogisticRegression_load_10 (a0, a1, a2, a3);
}

static jlong (*proc_video_KalmanFilter_get_1errorCovPre_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_video_KalmanFilter_get_1errorCovPre_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_KalmanFilter_get_1errorCovPre_10 (a0, a1, a2);
}

static void (*proc_core_Core_bitwise_1and_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_core_Core_bitwise_1and_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_core_Core_bitwise_1and_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_features2d_FastFeatureDetector_setType_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_features2d_FastFeatureDetector_setType_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_features2d_FastFeatureDetector_setType_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_drawContours_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jdouble,  jdouble,  jdouble,  jdouble,  jint,  jint,  jlong,  jint,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_drawContours_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8, jint a9, jint a10, jlong a11, jint a12, jdouble a13, jdouble a14) {
    proc_imgproc_Imgproc_drawContours_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14);
}

static jdouble (*proc_calib3d_Calib3d_calibrateCameraExtended_10)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jint,  jint,  jint,  jdouble);
jdouble Java_org_opencv_calib3d_Calib3d_calibrateCameraExtended_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jlong a6, jlong a7, jlong a8, jlong a9, jlong a10, jlong a11, jlong a12, jint a13, jint a14, jint a15, jdouble a16) {
    return proc_calib3d_Calib3d_calibrateCameraExtended_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16);
}

static jdouble (*proc_imgproc_CLAHE_getClipLimit_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_imgproc_CLAHE_getClipLimit_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_imgproc_CLAHE_getClipLimit_10 (a0, a1, a2);
}

static jdouble (*proc_ml_ANN_1MLP_getBackpropWeightScale_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_ml_ANN_1MLP_getBackpropWeightScale_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_ANN_1MLP_getBackpropWeightScale_10 (a0, a1, a2);
}

static jint (*proc_core_Mat_nGetF)(JNIEnv*,  jclass,  jlong,  jint,  jint,  jint,  jfloatArray);
jint Java_org_opencv_core_Mat_nGetF(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4, jint a5, jfloatArray a6) {
    return proc_core_Mat_nGetF (a0, a1, a2, a3, a4, a5, a6);
}

static jlong (*proc_imgcodecs_Imgcodecs_imread_10)(JNIEnv*,  jclass,  jstring,  jint);
jlong Java_org_opencv_imgcodecs_Imgcodecs_imread_10(JNIEnv* a0, jclass a1, jstring a2, jint a3) {
    return proc_imgcodecs_Imgcodecs_imread_10 (a0, a1, a2, a3);
}

static jlong (*proc_objdetect_HOGDescriptor_HOGDescriptor_11)(JNIEnv*,  jclass,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jint);
jlong Java_org_opencv_objdetect_HOGDescriptor_HOGDescriptor_11(JNIEnv* a0, jclass a1, jdouble a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8, jdouble a9, jint a10) {
    return proc_objdetect_HOGDescriptor_HOGDescriptor_11 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10);
}

static jlong (*proc_videoio_VideoCapture_VideoCapture_13)(JNIEnv*,  jclass);
jlong Java_org_opencv_videoio_VideoCapture_VideoCapture_13(JNIEnv* a0, jclass a1) {
    return proc_videoio_VideoCapture_VideoCapture_13 (a0, a1);
}

static jstring (*proc_features2d_AKAZE_getDefaultName_10)(JNIEnv*,  jclass,  jlong);
jstring Java_org_opencv_features2d_AKAZE_getDefaultName_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_AKAZE_getDefaultName_10 (a0, a1, a2);
}

static void (*proc_dnn_DictValue_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_dnn_DictValue_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_dnn_DictValue_delete (a0, a1, a2);
}

static void (*proc_photo_AlignMTB_computeBitmaps_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_photo_AlignMTB_computeBitmaps_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_photo_AlignMTB_computeBitmaps_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_video_BackgroundSubtractor_getBackgroundImage_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_video_BackgroundSubtractor_getBackgroundImage_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_video_BackgroundSubtractor_getBackgroundImage_10 (a0, a1, a2, a3);
}

static jfloat (*proc_photo_TonemapDrago_getSaturation_10)(JNIEnv*,  jclass,  jlong);
jfloat Java_org_opencv_photo_TonemapDrago_getSaturation_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_photo_TonemapDrago_getSaturation_10 (a0, a1, a2);
}

static jint (*proc_calib3d_StereoBM_getPreFilterSize_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_calib3d_StereoBM_getPreFilterSize_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_calib3d_StereoBM_getPreFilterSize_10 (a0, a1, a2);
}

static jlong (*proc_ml_LogisticRegression_load_11)(JNIEnv*,  jclass,  jstring);
jlong Java_org_opencv_ml_LogisticRegression_load_11(JNIEnv* a0, jclass a1, jstring a2) {
    return proc_ml_LogisticRegression_load_11 (a0, a1, a2);
}

static jlong (*proc_video_KalmanFilter_get_1gain_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_video_KalmanFilter_get_1gain_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_KalmanFilter_get_1gain_10 (a0, a1, a2);
}

static void (*proc_core_Core_bitwise_1and_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_core_Core_bitwise_1and_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_core_Core_bitwise_1and_11 (a0, a1, a2, a3, a4);
}

static void (*proc_features2d_Feature2D_compute_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_features2d_Feature2D_compute_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_features2d_Feature2D_compute_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_imgproc_Imgproc_drawContours_11)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jdouble,  jdouble,  jdouble,  jdouble,  jint);
void Java_org_opencv_imgproc_Imgproc_drawContours_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8, jint a9) {
    proc_imgproc_Imgproc_drawContours_11 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static jfloat (*proc_ml_LogisticRegression_predict_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jint);
jfloat Java_org_opencv_ml_LogisticRegression_predict_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jint a5) {
    return proc_ml_LogisticRegression_predict_10 (a0, a1, a2, a3, a4, a5);
}

static jint (*proc_calib3d_StereoBM_getPreFilterType_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_calib3d_StereoBM_getPreFilterType_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_calib3d_StereoBM_getPreFilterType_10 (a0, a1, a2);
}

static jlong (*proc_video_KalmanFilter_get_1measurementMatrix_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_video_KalmanFilter_get_1measurementMatrix_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_KalmanFilter_get_1measurementMatrix_10 (a0, a1, a2);
}

static void (*proc_core_Core_bitwise_1not_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_core_Core_bitwise_1not_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_core_Core_bitwise_1not_10 (a0, a1, a2, a3, a4);
}

static void (*proc_features2d_Feature2D_compute_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_features2d_Feature2D_compute_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_features2d_Feature2D_compute_11 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_imgproc_Imgproc_drawContours_12)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jdouble,  jdouble,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_drawContours_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8) {
    proc_imgproc_Imgproc_drawContours_12 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static void (*proc_photo_TonemapDrago_setBias_10)(JNIEnv*,  jclass,  jlong,  jfloat);
void Java_org_opencv_photo_TonemapDrago_setBias_10(JNIEnv* a0, jclass a1, jlong a2, jfloat a3) {
    proc_photo_TonemapDrago_setBias_10 (a0, a1, a2, a3);
}

static jdoubleArray (*proc_calib3d_StereoBM_getROI1_10)(JNIEnv*,  jclass,  jlong);
jdoubleArray Java_org_opencv_calib3d_StereoBM_getROI1_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_calib3d_StereoBM_getROI1_10 (a0, a1, a2);
}

static jfloat (*proc_ml_LogisticRegression_predict_11)(JNIEnv*,  jclass,  jlong,  jlong);
jfloat Java_org_opencv_ml_LogisticRegression_predict_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_ml_LogisticRegression_predict_11 (a0, a1, a2, a3);
}

static jint (*proc_features2d_Feature2D_defaultNorm_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_features2d_Feature2D_defaultNorm_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_Feature2D_defaultNorm_10 (a0, a1, a2);
}

static jlong (*proc_video_KalmanFilter_get_1measurementNoiseCov_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_video_KalmanFilter_get_1measurementNoiseCov_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_KalmanFilter_get_1measurementNoiseCov_10 (a0, a1, a2);
}

static void (*proc_core_Core_bitwise_1not_11)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_core_Core_bitwise_1not_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_core_Core_bitwise_1not_11 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_drawMarker_10)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jint,  jint,  jint,  jint);
void Java_org_opencv_imgproc_Imgproc_drawMarker_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8, jint a9, jint a10, jint a11, jint a12) {
    proc_imgproc_Imgproc_drawMarker_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12);
}

static void (*proc_photo_TonemapDrago_setSaturation_10)(JNIEnv*,  jclass,  jlong,  jfloat);
void Java_org_opencv_photo_TonemapDrago_setSaturation_10(JNIEnv* a0, jclass a1, jlong a2, jfloat a3) {
    proc_photo_TonemapDrago_setSaturation_10 (a0, a1, a2, a3);
}

static jdoubleArray (*proc_calib3d_StereoBM_getROI2_10)(JNIEnv*,  jclass,  jlong);
jdoubleArray Java_org_opencv_calib3d_StereoBM_getROI2_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_calib3d_StereoBM_getROI2_10 (a0, a1, a2);
}

static jlong (*proc_video_KalmanFilter_get_1processNoiseCov_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_video_KalmanFilter_get_1processNoiseCov_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_KalmanFilter_get_1processNoiseCov_10 (a0, a1, a2);
}

static void (*proc_core_Core_bitwise_1or_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_core_Core_bitwise_1or_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_core_Core_bitwise_1or_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_features2d_Feature2D_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_features2d_Feature2D_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_features2d_Feature2D_delete (a0, a1, a2);
}

static void (*proc_imgproc_Imgproc_drawMarker_11)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_drawMarker_11(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8) {
    proc_imgproc_Imgproc_drawMarker_11 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static void (*proc_ml_LogisticRegression_setIterations_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_ml_LogisticRegression_setIterations_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_ml_LogisticRegression_setIterations_10 (a0, a1, a2, a3);
}

static void (*proc_photo_TonemapDurand_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_photo_TonemapDurand_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_photo_TonemapDurand_delete (a0, a1, a2);
}

static jfloat (*proc_photo_TonemapDurand_getContrast_10)(JNIEnv*,  jclass,  jlong);
jfloat Java_org_opencv_photo_TonemapDurand_getContrast_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_photo_TonemapDurand_getContrast_10 (a0, a1, a2);
}

static jint (*proc_calib3d_StereoBM_getSmallerBlockSize_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_calib3d_StereoBM_getSmallerBlockSize_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_calib3d_StereoBM_getSmallerBlockSize_10 (a0, a1, a2);
}

static jint (*proc_features2d_Feature2D_descriptorSize_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_features2d_Feature2D_descriptorSize_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_Feature2D_descriptorSize_10 (a0, a1, a2);
}

static jlong (*proc_video_KalmanFilter_get_1statePost_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_video_KalmanFilter_get_1statePost_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_KalmanFilter_get_1statePost_10 (a0, a1, a2);
}

static void (*proc_core_Core_bitwise_1or_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_core_Core_bitwise_1or_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_core_Core_bitwise_1or_11 (a0, a1, a2, a3, a4);
}

static void (*proc_imgproc_Imgproc_ellipse2Poly_10)(JNIEnv*,  jclass,  jdouble,  jdouble,  jdouble,  jdouble,  jint,  jint,  jint,  jint,  jlong);
void Java_org_opencv_imgproc_Imgproc_ellipse2Poly_10(JNIEnv* a0, jclass a1, jdouble a2, jdouble a3, jdouble a4, jdouble a5, jint a6, jint a7, jint a8, jint a9, jlong a10) {
    proc_imgproc_Imgproc_ellipse2Poly_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10);
}

static void (*proc_ml_LogisticRegression_setLearningRate_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_ml_LogisticRegression_setLearningRate_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_ml_LogisticRegression_setLearningRate_10 (a0, a1, a2, a3);
}

static jfloat (*proc_photo_TonemapDurand_getSaturation_10)(JNIEnv*,  jclass,  jlong);
jfloat Java_org_opencv_photo_TonemapDurand_getSaturation_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_photo_TonemapDurand_getSaturation_10 (a0, a1, a2);
}

static jint (*proc_calib3d_StereoBM_getTextureThreshold_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_calib3d_StereoBM_getTextureThreshold_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_calib3d_StereoBM_getTextureThreshold_10 (a0, a1, a2);
}

static jint (*proc_features2d_Feature2D_descriptorType_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_features2d_Feature2D_descriptorType_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_Feature2D_descriptorType_10 (a0, a1, a2);
}

static jlong (*proc_video_KalmanFilter_get_1statePre_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_video_KalmanFilter_get_1statePre_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_KalmanFilter_get_1statePre_10 (a0, a1, a2);
}

static void (*proc_core_Core_bitwise_1xor_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_core_Core_bitwise_1xor_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_core_Core_bitwise_1xor_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_imgproc_Imgproc_ellipse_10)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jint,  jint,  jint);
void Java_org_opencv_imgproc_Imgproc_ellipse_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8, jdouble a9, jdouble a10, jdouble a11, jdouble a12, jdouble a13, jint a14, jint a15, jint a16) {
    proc_imgproc_Imgproc_ellipse_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16);
}

static void (*proc_ml_LogisticRegression_setMiniBatchSize_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_ml_LogisticRegression_setMiniBatchSize_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_ml_LogisticRegression_setMiniBatchSize_10 (a0, a1, a2, a3);
}

static jfloat (*proc_photo_TonemapDurand_getSigmaColor_10)(JNIEnv*,  jclass,  jlong);
jfloat Java_org_opencv_photo_TonemapDurand_getSigmaColor_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_photo_TonemapDurand_getSigmaColor_10 (a0, a1, a2);
}

static jint (*proc_calib3d_StereoBM_getUniquenessRatio_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_calib3d_StereoBM_getUniquenessRatio_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_calib3d_StereoBM_getUniquenessRatio_10 (a0, a1, a2);
}

static jlong (*proc_video_KalmanFilter_get_1transitionMatrix_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_video_KalmanFilter_get_1transitionMatrix_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_KalmanFilter_get_1transitionMatrix_10 (a0, a1, a2);
}

static void (*proc_core_Core_bitwise_1xor_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_core_Core_bitwise_1xor_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_core_Core_bitwise_1xor_11 (a0, a1, a2, a3, a4);
}

static void (*proc_features2d_Feature2D_detectAndCompute_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jboolean);
void Java_org_opencv_features2d_Feature2D_detectAndCompute_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jboolean a7) {
    proc_features2d_Feature2D_detectAndCompute_10 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_imgproc_Imgproc_ellipse_11)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jint);
void Java_org_opencv_imgproc_Imgproc_ellipse_11(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8, jdouble a9, jdouble a10, jdouble a11, jdouble a12, jdouble a13, jint a14) {
    proc_imgproc_Imgproc_ellipse_11 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14);
}

static void (*proc_ml_LogisticRegression_setRegularization_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_ml_LogisticRegression_setRegularization_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_ml_LogisticRegression_setRegularization_10 (a0, a1, a2, a3);
}

static jfloat (*proc_photo_TonemapDurand_getSigmaSpace_10)(JNIEnv*,  jclass,  jlong);
jfloat Java_org_opencv_photo_TonemapDurand_getSigmaSpace_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_photo_TonemapDurand_getSigmaSpace_10 (a0, a1, a2);
}

static jint (*proc_core_Core_borderInterpolate_10)(JNIEnv*,  jclass,  jint,  jint,  jint);
jint Java_org_opencv_core_Core_borderInterpolate_10(JNIEnv* a0, jclass a1, jint a2, jint a3, jint a4) {
    return proc_core_Core_borderInterpolate_10 (a0, a1, a2, a3, a4);
}

static jlong (*proc_video_KalmanFilter_predict_10)(JNIEnv*,  jclass,  jlong,  jlong);
jlong Java_org_opencv_video_KalmanFilter_predict_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_video_KalmanFilter_predict_10 (a0, a1, a2, a3);
}

static void (*proc_calib3d_StereoBM_setPreFilterCap_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_calib3d_StereoBM_setPreFilterCap_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_calib3d_StereoBM_setPreFilterCap_10 (a0, a1, a2, a3);
}

static void (*proc_features2d_Feature2D_detectAndCompute_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_features2d_Feature2D_detectAndCompute_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6) {
    proc_features2d_Feature2D_detectAndCompute_11 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_imgproc_Imgproc_ellipse_12)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_ellipse_12(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8, jdouble a9, jdouble a10, jdouble a11, jdouble a12, jdouble a13) {
    proc_imgproc_Imgproc_ellipse_12 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13);
}

static void (*proc_ml_LogisticRegression_setTermCriteria_10)(JNIEnv*,  jclass,  jlong,  jint,  jint,  jdouble);
void Java_org_opencv_ml_LogisticRegression_setTermCriteria_10(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4, jdouble a5) {
    proc_ml_LogisticRegression_setTermCriteria_10 (a0, a1, a2, a3, a4, a5);
}

static jlong (*proc_video_KalmanFilter_predict_11)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_video_KalmanFilter_predict_11(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_KalmanFilter_predict_11 (a0, a1, a2);
}

static void (*proc_calib3d_StereoBM_setPreFilterSize_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_calib3d_StereoBM_setPreFilterSize_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_calib3d_StereoBM_setPreFilterSize_10 (a0, a1, a2, a3);
}

static void (*proc_core_Core_calcCovarMatrix_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jint,  jint);
void Java_org_opencv_core_Core_calcCovarMatrix_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jint a5, jint a6) {
    proc_core_Core_calcCovarMatrix_10 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_features2d_Feature2D_detect_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_features2d_Feature2D_detect_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_features2d_Feature2D_detect_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_imgproc_Imgproc_ellipse_13)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jint,  jint);
void Java_org_opencv_imgproc_Imgproc_ellipse_13(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8, jdouble a9, jdouble a10, jdouble a11, jint a12, jint a13) {
    proc_imgproc_Imgproc_ellipse_13 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13);
}

static void (*proc_ml_LogisticRegression_setTrainMethod_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_ml_LogisticRegression_setTrainMethod_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_ml_LogisticRegression_setTrainMethod_10 (a0, a1, a2, a3);
}

static void (*proc_photo_TonemapDurand_setContrast_10)(JNIEnv*,  jclass,  jlong,  jfloat);
void Java_org_opencv_photo_TonemapDurand_setContrast_10(JNIEnv* a0, jclass a1, jlong a2, jfloat a3) {
    proc_photo_TonemapDurand_setContrast_10 (a0, a1, a2, a3);
}

static jlong (*proc_ml_NormalBayesClassifier_create_10)(JNIEnv*,  jclass);
jlong Java_org_opencv_ml_NormalBayesClassifier_create_10(JNIEnv* a0, jclass a1) {
    return proc_ml_NormalBayesClassifier_create_10 (a0, a1);
}

static void (*proc_calib3d_StereoBM_setPreFilterType_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_calib3d_StereoBM_setPreFilterType_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_calib3d_StereoBM_setPreFilterType_10 (a0, a1, a2, a3);
}

static void (*proc_core_Core_calcCovarMatrix_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jint);
void Java_org_opencv_core_Core_calcCovarMatrix_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jint a5) {
    proc_core_Core_calcCovarMatrix_11 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_features2d_Feature2D_detect_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_features2d_Feature2D_detect_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_features2d_Feature2D_detect_11 (a0, a1, a2, a3, a4);
}

static void (*proc_imgproc_Imgproc_ellipse_14)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jint);
void Java_org_opencv_imgproc_Imgproc_ellipse_14(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8, jdouble a9, jdouble a10, jdouble a11, jint a12) {
    proc_imgproc_Imgproc_ellipse_14 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12);
}

static void (*proc_photo_TonemapDurand_setSaturation_10)(JNIEnv*,  jclass,  jlong,  jfloat);
void Java_org_opencv_photo_TonemapDurand_setSaturation_10(JNIEnv* a0, jclass a1, jlong a2, jfloat a3) {
    proc_photo_TonemapDurand_setSaturation_10 (a0, a1, a2, a3);
}

static void (*proc_video_KalmanFilter_set_1controlMatrix_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_video_KalmanFilter_set_1controlMatrix_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_video_KalmanFilter_set_1controlMatrix_10 (a0, a1, a2, a3);
}

static jdouble (*proc_calib3d_Calib3d_calibrateCameraExtended_11)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jint);
jdouble Java_org_opencv_calib3d_Calib3d_calibrateCameraExtended_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jlong a6, jlong a7, jlong a8, jlong a9, jlong a10, jlong a11, jlong a12, jint a13) {
    return proc_calib3d_Calib3d_calibrateCameraExtended_11 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13);
}

static jdoubleArray (*proc_imgproc_CLAHE_getTilesGridSize_10)(JNIEnv*,  jclass,  jlong);
jdoubleArray Java_org_opencv_imgproc_CLAHE_getTilesGridSize_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_imgproc_CLAHE_getTilesGridSize_10 (a0, a1, a2);
}

static jint (*proc_core_Mat_nGetI)(JNIEnv*,  jclass,  jlong,  jint,  jint,  jint,  jintArray);
jint Java_org_opencv_core_Mat_nGetI(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4, jint a5, jintArray a6) {
    return proc_core_Mat_nGetI (a0, a1, a2, a3, a4, a5, a6);
}

static jint (*proc_dnn_DictValue_getIntValue_10)(JNIEnv*,  jclass,  jlong,  jint);
jint Java_org_opencv_dnn_DictValue_getIntValue_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    return proc_dnn_DictValue_getIntValue_10 (a0, a1, a2, a3);
}

static jint (*proc_features2d_AKAZE_getDescriptorChannels_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_features2d_AKAZE_getDescriptorChannels_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_AKAZE_getDescriptorChannels_10 (a0, a1, a2);
}

static jlong (*proc_imgcodecs_Imgcodecs_imread_11)(JNIEnv*,  jclass,  jstring);
jlong Java_org_opencv_imgcodecs_Imgcodecs_imread_11(JNIEnv* a0, jclass a1, jstring a2) {
    return proc_imgcodecs_Imgcodecs_imread_11 (a0, a1, a2);
}

static jlong (*proc_ml_ANN_1MLP_getLayerSizes_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_ml_ANN_1MLP_getLayerSizes_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_ANN_1MLP_getLayerSizes_10 (a0, a1, a2);
}

static jlong (*proc_objdetect_HOGDescriptor_HOGDescriptor_12)(JNIEnv*,  jclass,  jstring);
jlong Java_org_opencv_objdetect_HOGDescriptor_HOGDescriptor_12(JNIEnv* a0, jclass a1, jstring a2) {
    return proc_objdetect_HOGDescriptor_HOGDescriptor_12 (a0, a1, a2);
}

static void (*proc_photo_AlignMTB_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_photo_AlignMTB_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_photo_AlignMTB_delete (a0, a1, a2);
}

static void (*proc_video_BackgroundSubtractorKNN_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_video_BackgroundSubtractorKNN_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_video_BackgroundSubtractorKNN_delete (a0, a1, a2);
}

static void (*proc_videoio_VideoCapture_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_videoio_VideoCapture_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_videoio_VideoCapture_delete (a0, a1, a2);
}

static void (*proc_calib3d_StereoBM_setROI1_10)(JNIEnv*,  jclass,  jlong,  jint,  jint,  jint,  jint);
void Java_org_opencv_calib3d_StereoBM_setROI1_10(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4, jint a5, jint a6) {
    proc_calib3d_StereoBM_setROI1_10 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_core_Core_cartToPolar_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jboolean);
void Java_org_opencv_core_Core_cartToPolar_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jboolean a6) {
    proc_core_Core_cartToPolar_10 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_features2d_Feature2D_detect_12)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_features2d_Feature2D_detect_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_features2d_Feature2D_detect_12 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_imgproc_Imgproc_ellipse_15)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_ellipse_15(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8, jdouble a9, jdouble a10, jdouble a11) {
    proc_imgproc_Imgproc_ellipse_15 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11);
}

static void (*proc_ml_NormalBayesClassifier_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_ml_NormalBayesClassifier_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_ml_NormalBayesClassifier_delete (a0, a1, a2);
}

static void (*proc_photo_TonemapDurand_setSigmaColor_10)(JNIEnv*,  jclass,  jlong,  jfloat);
void Java_org_opencv_photo_TonemapDurand_setSigmaColor_10(JNIEnv* a0, jclass a1, jlong a2, jfloat a3) {
    proc_photo_TonemapDurand_setSigmaColor_10 (a0, a1, a2, a3);
}

static void (*proc_video_KalmanFilter_set_1errorCovPost_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_video_KalmanFilter_set_1errorCovPost_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_video_KalmanFilter_set_1errorCovPost_10 (a0, a1, a2, a3);
}

static jlong (*proc_ml_NormalBayesClassifier_load_10)(JNIEnv*,  jclass,  jstring,  jstring);
jlong Java_org_opencv_ml_NormalBayesClassifier_load_10(JNIEnv* a0, jclass a1, jstring a2, jstring a3) {
    return proc_ml_NormalBayesClassifier_load_10 (a0, a1, a2, a3);
}

static void (*proc_calib3d_StereoBM_setROI2_10)(JNIEnv*,  jclass,  jlong,  jint,  jint,  jint,  jint);
void Java_org_opencv_calib3d_StereoBM_setROI2_10(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4, jint a5, jint a6) {
    proc_calib3d_StereoBM_setROI2_10 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_core_Core_cartToPolar_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_core_Core_cartToPolar_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_core_Core_cartToPolar_11 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_features2d_Feature2D_detect_13)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_features2d_Feature2D_detect_13(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_features2d_Feature2D_detect_13 (a0, a1, a2, a3, a4);
}

static void (*proc_imgproc_Imgproc_equalizeHist_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_imgproc_Imgproc_equalizeHist_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_imgproc_Imgproc_equalizeHist_10 (a0, a1, a2, a3);
}

static void (*proc_photo_TonemapDurand_setSigmaSpace_10)(JNIEnv*,  jclass,  jlong,  jfloat);
void Java_org_opencv_photo_TonemapDurand_setSigmaSpace_10(JNIEnv* a0, jclass a1, jlong a2, jfloat a3) {
    proc_photo_TonemapDurand_setSigmaSpace_10 (a0, a1, a2, a3);
}

static void (*proc_video_KalmanFilter_set_1errorCovPre_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_video_KalmanFilter_set_1errorCovPre_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_video_KalmanFilter_set_1errorCovPre_10 (a0, a1, a2, a3);
}

static jboolean (*proc_core_Core_checkRange_10)(JNIEnv*,  jclass,  jlong,  jboolean,  jdouble,  jdouble);
jboolean Java_org_opencv_core_Core_checkRange_10(JNIEnv* a0, jclass a1, jlong a2, jboolean a3, jdouble a4, jdouble a5) {
    return proc_core_Core_checkRange_10 (a0, a1, a2, a3, a4, a5);
}

static jboolean (*proc_features2d_Feature2D_empty_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_features2d_Feature2D_empty_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_Feature2D_empty_10 (a0, a1, a2);
}

static jlong (*proc_ml_NormalBayesClassifier_load_11)(JNIEnv*,  jclass,  jstring);
jlong Java_org_opencv_ml_NormalBayesClassifier_load_11(JNIEnv* a0, jclass a1, jstring a2) {
    return proc_ml_NormalBayesClassifier_load_11 (a0, a1, a2);
}

static void (*proc_calib3d_StereoBM_setSmallerBlockSize_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_calib3d_StereoBM_setSmallerBlockSize_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_calib3d_StereoBM_setSmallerBlockSize_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_erode_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jdouble,  jdouble,  jint,  jint,  jdouble,  jdouble,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_erode_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jdouble a5, jdouble a6, jint a7, jint a8, jdouble a9, jdouble a10, jdouble a11, jdouble a12) {
    proc_imgproc_Imgproc_erode_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12);
}

static void (*proc_photo_TonemapMantiuk_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_photo_TonemapMantiuk_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_photo_TonemapMantiuk_delete (a0, a1, a2);
}

static void (*proc_video_KalmanFilter_set_1gain_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_video_KalmanFilter_set_1gain_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_video_KalmanFilter_set_1gain_10 (a0, a1, a2, a3);
}

static jboolean (*proc_core_Core_checkRange_11)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_core_Core_checkRange_11(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_core_Core_checkRange_11 (a0, a1, a2);
}

static jfloat (*proc_ml_NormalBayesClassifier_predictProb_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jint);
jfloat Java_org_opencv_ml_NormalBayesClassifier_predictProb_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jint a6) {
    return proc_ml_NormalBayesClassifier_predictProb_10 (a0, a1, a2, a3, a4, a5, a6);
}

static jfloat (*proc_photo_TonemapMantiuk_getSaturation_10)(JNIEnv*,  jclass,  jlong);
jfloat Java_org_opencv_photo_TonemapMantiuk_getSaturation_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_photo_TonemapMantiuk_getSaturation_10 (a0, a1, a2);
}

static jstring (*proc_features2d_Feature2D_getDefaultName_10)(JNIEnv*,  jclass,  jlong);
jstring Java_org_opencv_features2d_Feature2D_getDefaultName_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_Feature2D_getDefaultName_10 (a0, a1, a2);
}

static void (*proc_calib3d_StereoBM_setTextureThreshold_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_calib3d_StereoBM_setTextureThreshold_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_calib3d_StereoBM_setTextureThreshold_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_erode_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jdouble,  jdouble,  jint);
void Java_org_opencv_imgproc_Imgproc_erode_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jdouble a5, jdouble a6, jint a7) {
    proc_imgproc_Imgproc_erode_11 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_video_KalmanFilter_set_1measurementMatrix_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_video_KalmanFilter_set_1measurementMatrix_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_video_KalmanFilter_set_1measurementMatrix_10 (a0, a1, a2, a3);
}

static jfloat (*proc_ml_NormalBayesClassifier_predictProb_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
jfloat Java_org_opencv_ml_NormalBayesClassifier_predictProb_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    return proc_ml_NormalBayesClassifier_predictProb_11 (a0, a1, a2, a3, a4, a5);
}

static jfloat (*proc_photo_TonemapMantiuk_getScale_10)(JNIEnv*,  jclass,  jlong);
jfloat Java_org_opencv_photo_TonemapMantiuk_getScale_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_photo_TonemapMantiuk_getScale_10 (a0, a1, a2);
}

static void (*proc_calib3d_StereoBM_setUniquenessRatio_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_calib3d_StereoBM_setUniquenessRatio_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_calib3d_StereoBM_setUniquenessRatio_10 (a0, a1, a2, a3);
}

static void (*proc_core_Core_compare_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jint);
void Java_org_opencv_core_Core_compare_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jint a5) {
    proc_core_Core_compare_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_features2d_Feature2D_read_10)(JNIEnv*,  jclass,  jlong,  jstring);
void Java_org_opencv_features2d_Feature2D_read_10(JNIEnv* a0, jclass a1, jlong a2, jstring a3) {
    proc_features2d_Feature2D_read_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_erode_12)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_imgproc_Imgproc_erode_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_imgproc_Imgproc_erode_12 (a0, a1, a2, a3, a4);
}

static void (*proc_video_KalmanFilter_set_1measurementNoiseCov_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_video_KalmanFilter_set_1measurementNoiseCov_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_video_KalmanFilter_set_1measurementNoiseCov_10 (a0, a1, a2, a3);
}

static jlong (*proc_ml_ParamGrid_create_10)(JNIEnv*,  jclass,  jdouble,  jdouble,  jdouble);
jlong Java_org_opencv_ml_ParamGrid_create_10(JNIEnv* a0, jclass a1, jdouble a2, jdouble a3, jdouble a4) {
    return proc_ml_ParamGrid_create_10 (a0, a1, a2, a3, a4);
}

static void (*proc_calib3d_StereoMatcher_compute_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_calib3d_StereoMatcher_compute_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_calib3d_StereoMatcher_compute_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_core_Core_compare_11)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jlong,  jint);
void Java_org_opencv_core_Core_compare_11(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jlong a7, jint a8) {
    proc_core_Core_compare_11 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static void (*proc_features2d_Feature2D_write_10)(JNIEnv*,  jclass,  jlong,  jstring);
void Java_org_opencv_features2d_Feature2D_write_10(JNIEnv* a0, jclass a1, jlong a2, jstring a3) {
    proc_features2d_Feature2D_write_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_fillConvexPoly_10)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jint,  jint);
void Java_org_opencv_imgproc_Imgproc_fillConvexPoly_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7, jint a8, jint a9) {
    proc_imgproc_Imgproc_fillConvexPoly_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static void (*proc_photo_TonemapMantiuk_setSaturation_10)(JNIEnv*,  jclass,  jlong,  jfloat);
void Java_org_opencv_photo_TonemapMantiuk_setSaturation_10(JNIEnv* a0, jclass a1, jlong a2, jfloat a3) {
    proc_photo_TonemapMantiuk_setSaturation_10 (a0, a1, a2, a3);
}

static void (*proc_video_KalmanFilter_set_1processNoiseCov_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_video_KalmanFilter_set_1processNoiseCov_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_video_KalmanFilter_set_1processNoiseCov_10 (a0, a1, a2, a3);
}

static jlong (*proc_features2d_FeatureDetector_create_10)(JNIEnv*,  jclass,  jint);
jlong Java_org_opencv_features2d_FeatureDetector_create_10(JNIEnv* a0, jclass a1, jint a2) {
    return proc_features2d_FeatureDetector_create_10 (a0, a1, a2);
}

static jlong (*proc_ml_ParamGrid_create_11)(JNIEnv*,  jclass);
jlong Java_org_opencv_ml_ParamGrid_create_11(JNIEnv* a0, jclass a1) {
    return proc_ml_ParamGrid_create_11 (a0, a1);
}

static void (*proc_calib3d_StereoMatcher_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_calib3d_StereoMatcher_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_calib3d_StereoMatcher_delete (a0, a1, a2);
}

static void (*proc_core_Core_completeSymm_10)(JNIEnv*,  jclass,  jlong,  jboolean);
void Java_org_opencv_core_Core_completeSymm_10(JNIEnv* a0, jclass a1, jlong a2, jboolean a3) {
    proc_core_Core_completeSymm_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_fillConvexPoly_11)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_fillConvexPoly_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7) {
    proc_imgproc_Imgproc_fillConvexPoly_11 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_photo_TonemapMantiuk_setScale_10)(JNIEnv*,  jclass,  jlong,  jfloat);
void Java_org_opencv_photo_TonemapMantiuk_setScale_10(JNIEnv* a0, jclass a1, jlong a2, jfloat a3) {
    proc_photo_TonemapMantiuk_setScale_10 (a0, a1, a2, a3);
}

static void (*proc_video_KalmanFilter_set_1statePost_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_video_KalmanFilter_set_1statePost_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_video_KalmanFilter_set_1statePost_10 (a0, a1, a2, a3);
}

static jint (*proc_calib3d_StereoMatcher_getBlockSize_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_calib3d_StereoMatcher_getBlockSize_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_calib3d_StereoMatcher_getBlockSize_10 (a0, a1, a2);
}

static void (*proc_core_Core_completeSymm_11)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_core_Core_completeSymm_11(JNIEnv* a0, jclass a1, jlong a2) {
    proc_core_Core_completeSymm_11 (a0, a1, a2);
}

static void (*proc_features2d_FeatureDetector_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_features2d_FeatureDetector_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_features2d_FeatureDetector_delete (a0, a1, a2);
}

static void (*proc_imgproc_Imgproc_fillPoly_10)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jint,  jint,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_fillPoly_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7, jint a8, jint a9, jdouble a10, jdouble a11) {
    proc_imgproc_Imgproc_fillPoly_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11);
}

static void (*proc_ml_ParamGrid_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_ml_ParamGrid_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_ml_ParamGrid_delete (a0, a1, a2);
}

static void (*proc_photo_TonemapReinhard_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_photo_TonemapReinhard_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_photo_TonemapReinhard_delete (a0, a1, a2);
}

static void (*proc_video_KalmanFilter_set_1statePre_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_video_KalmanFilter_set_1statePre_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_video_KalmanFilter_set_1statePre_10 (a0, a1, a2, a3);
}

static jdouble (*proc_ml_ParamGrid_get_1logStep_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_ml_ParamGrid_get_1logStep_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_ParamGrid_get_1logStep_10 (a0, a1, a2);
}

static jfloat (*proc_photo_TonemapReinhard_getColorAdaptation_10)(JNIEnv*,  jclass,  jlong);
jfloat Java_org_opencv_photo_TonemapReinhard_getColorAdaptation_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_photo_TonemapReinhard_getColorAdaptation_10 (a0, a1, a2);
}

static jint (*proc_calib3d_StereoMatcher_getDisp12MaxDiff_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_calib3d_StereoMatcher_getDisp12MaxDiff_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_calib3d_StereoMatcher_getDisp12MaxDiff_10 (a0, a1, a2);
}

static void (*proc_core_Core_convertFp16_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_core_Core_convertFp16_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_core_Core_convertFp16_10 (a0, a1, a2, a3);
}

static void (*proc_features2d_FeatureDetector_detect_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_features2d_FeatureDetector_detect_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_features2d_FeatureDetector_detect_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_imgproc_Imgproc_fillPoly_11)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_fillPoly_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7) {
    proc_imgproc_Imgproc_fillPoly_11 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_video_KalmanFilter_set_1transitionMatrix_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_video_KalmanFilter_set_1transitionMatrix_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_video_KalmanFilter_set_1transitionMatrix_10 (a0, a1, a2, a3);
}

static jdouble (*proc_ml_ParamGrid_get_1maxVal_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_ml_ParamGrid_get_1maxVal_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_ParamGrid_get_1maxVal_10 (a0, a1, a2);
}

static jfloat (*proc_photo_TonemapReinhard_getIntensity_10)(JNIEnv*,  jclass,  jlong);
jfloat Java_org_opencv_photo_TonemapReinhard_getIntensity_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_photo_TonemapReinhard_getIntensity_10 (a0, a1, a2);
}

static jint (*proc_calib3d_StereoMatcher_getMinDisparity_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_calib3d_StereoMatcher_getMinDisparity_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_calib3d_StereoMatcher_getMinDisparity_10 (a0, a1, a2);
}

static void (*proc_core_Core_convertScaleAbs_10)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble);
void Java_org_opencv_core_Core_convertScaleAbs_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5) {
    proc_core_Core_convertScaleAbs_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_features2d_FeatureDetector_detect_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_features2d_FeatureDetector_detect_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_features2d_FeatureDetector_detect_11 (a0, a1, a2, a3, a4);
}

static void (*proc_imgproc_Imgproc_filter2D_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jlong,  jdouble,  jdouble,  jdouble,  jint);
void Java_org_opencv_imgproc_Imgproc_filter2D_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jlong a5, jdouble a6, jdouble a7, jdouble a8, jint a9) {
    proc_imgproc_Imgproc_filter2D_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static void (*proc_video_SparseOpticalFlow_calc_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_video_SparseOpticalFlow_calc_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7, jlong a8) {
    proc_video_SparseOpticalFlow_calc_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static jboolean (*proc_imgcodecs_Imgcodecs_imreadmulti_10)(JNIEnv*,  jclass,  jstring,  jlong,  jint);
jboolean Java_org_opencv_imgcodecs_Imgcodecs_imreadmulti_10(JNIEnv* a0, jclass a1, jstring a2, jlong a3, jint a4) {
    return proc_imgcodecs_Imgcodecs_imreadmulti_10 (a0, a1, a2, a3, a4);
}

static jboolean (*proc_photo_AlignMTB_getCut_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_photo_AlignMTB_getCut_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_photo_AlignMTB_getCut_10 (a0, a1, a2);
}

static jboolean (*proc_video_BackgroundSubtractorKNN_getDetectShadows_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_video_BackgroundSubtractorKNN_getDetectShadows_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_BackgroundSubtractorKNN_getDetectShadows_10 (a0, a1, a2);
}

static jdouble (*proc_calib3d_Calib3d_calibrateCameraExtended_12)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong);
jdouble Java_org_opencv_calib3d_Calib3d_calibrateCameraExtended_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jlong a6, jlong a7, jlong a8, jlong a9, jlong a10, jlong a11, jlong a12) {
    return proc_calib3d_Calib3d_calibrateCameraExtended_12 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12);
}

static jdouble (*proc_ml_ANN_1MLP_getRpropDW0_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_ml_ANN_1MLP_getRpropDW0_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_ANN_1MLP_getRpropDW0_10 (a0, a1, a2);
}

static jdouble (*proc_videoio_VideoCapture_get_10)(JNIEnv*,  jclass,  jlong,  jint);
jdouble Java_org_opencv_videoio_VideoCapture_get_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    return proc_videoio_VideoCapture_get_10 (a0, a1, a2, a3);
}

static jint (*proc_core_Mat_nGetS)(JNIEnv*,  jclass,  jlong,  jint,  jint,  jint,  jshortArray);
jint Java_org_opencv_core_Mat_nGetS(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4, jint a5, jshortArray a6) {
    return proc_core_Mat_nGetS (a0, a1, a2, a3, a4, a5, a6);
}

static jint (*proc_dnn_DictValue_getIntValue_11)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_dnn_DictValue_getIntValue_11(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_dnn_DictValue_getIntValue_11 (a0, a1, a2);
}

static jint (*proc_features2d_AKAZE_getDescriptorSize_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_features2d_AKAZE_getDescriptorSize_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_AKAZE_getDescriptorSize_10 (a0, a1, a2);
}

static jlong (*proc_objdetect_HOGDescriptor_HOGDescriptor_13)(JNIEnv*,  jclass);
jlong Java_org_opencv_objdetect_HOGDescriptor_HOGDescriptor_13(JNIEnv* a0, jclass a1) {
    return proc_objdetect_HOGDescriptor_HOGDescriptor_13 (a0, a1);
}

static void (*proc_imgproc_CLAHE_setClipLimit_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_imgproc_CLAHE_setClipLimit_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_imgproc_CLAHE_setClipLimit_10 (a0, a1, a2, a3);
}

static jdouble (*proc_ml_ParamGrid_get_1minVal_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_ml_ParamGrid_get_1minVal_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_ParamGrid_get_1minVal_10 (a0, a1, a2);
}

static jfloat (*proc_photo_TonemapReinhard_getLightAdaptation_10)(JNIEnv*,  jclass,  jlong);
jfloat Java_org_opencv_photo_TonemapReinhard_getLightAdaptation_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_photo_TonemapReinhard_getLightAdaptation_10 (a0, a1, a2);
}

static jint (*proc_calib3d_StereoMatcher_getNumDisparities_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_calib3d_StereoMatcher_getNumDisparities_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_calib3d_StereoMatcher_getNumDisparities_10 (a0, a1, a2);
}

static void (*proc_core_Core_convertScaleAbs_11)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_core_Core_convertScaleAbs_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_core_Core_convertScaleAbs_11 (a0, a1, a2, a3);
}

static void (*proc_features2d_FeatureDetector_detect_12)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_features2d_FeatureDetector_detect_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_features2d_FeatureDetector_detect_12 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_imgproc_Imgproc_filter2D_11)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jlong,  jdouble,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_filter2D_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jlong a5, jdouble a6, jdouble a7, jdouble a8) {
    proc_imgproc_Imgproc_filter2D_11 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static void (*proc_video_SparseOpticalFlow_calc_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_video_SparseOpticalFlow_calc_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7) {
    proc_video_SparseOpticalFlow_calc_11 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static jint (*proc_calib3d_StereoMatcher_getSpeckleRange_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_calib3d_StereoMatcher_getSpeckleRange_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_calib3d_StereoMatcher_getSpeckleRange_10 (a0, a1, a2);
}

static void (*proc_core_Core_copyMakeBorder_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint,  jint,  jint,  jint,  jdouble,  jdouble,  jdouble,  jdouble);
void Java_org_opencv_core_Core_copyMakeBorder_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5, jint a6, jint a7, jint a8, jdouble a9, jdouble a10, jdouble a11, jdouble a12) {
    proc_core_Core_copyMakeBorder_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12);
}

static void (*proc_features2d_FeatureDetector_detect_13)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_features2d_FeatureDetector_detect_13(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_features2d_FeatureDetector_detect_13 (a0, a1, a2, a3, a4);
}

static void (*proc_imgproc_Imgproc_filter2D_12)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jlong);
void Java_org_opencv_imgproc_Imgproc_filter2D_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jlong a5) {
    proc_imgproc_Imgproc_filter2D_12 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_ml_ParamGrid_set_1logStep_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_ml_ParamGrid_set_1logStep_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_ml_ParamGrid_set_1logStep_10 (a0, a1, a2, a3);
}

static void (*proc_photo_TonemapReinhard_setColorAdaptation_10)(JNIEnv*,  jclass,  jlong,  jfloat);
void Java_org_opencv_photo_TonemapReinhard_setColorAdaptation_10(JNIEnv* a0, jclass a1, jlong a2, jfloat a3) {
    proc_photo_TonemapReinhard_setColorAdaptation_10 (a0, a1, a2, a3);
}

static void (*proc_video_SparseOpticalFlow_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_video_SparseOpticalFlow_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_video_SparseOpticalFlow_delete (a0, a1, a2);
}

static jboolean (*proc_features2d_FeatureDetector_empty_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_features2d_FeatureDetector_empty_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_FeatureDetector_empty_10 (a0, a1, a2);
}

static jint (*proc_calib3d_StereoMatcher_getSpeckleWindowSize_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_calib3d_StereoMatcher_getSpeckleWindowSize_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_calib3d_StereoMatcher_getSpeckleWindowSize_10 (a0, a1, a2);
}

static jlong (*proc_video_SparsePyrLKOpticalFlow_create_10)(JNIEnv*,  jclass,  jdouble,  jdouble,  jint,  jint,  jint,  jdouble,  jint,  jdouble);
jlong Java_org_opencv_video_SparsePyrLKOpticalFlow_create_10(JNIEnv* a0, jclass a1, jdouble a2, jdouble a3, jint a4, jint a5, jint a6, jdouble a7, jint a8, jdouble a9) {
    return proc_video_SparsePyrLKOpticalFlow_create_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static void (*proc_core_Core_copyMakeBorder_11)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint,  jint,  jint,  jint);
void Java_org_opencv_core_Core_copyMakeBorder_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5, jint a6, jint a7, jint a8) {
    proc_core_Core_copyMakeBorder_11 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static void (*proc_imgproc_Imgproc_findContours_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jint,  jint,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_findContours_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jint a5, jint a6, jdouble a7, jdouble a8) {
    proc_imgproc_Imgproc_findContours_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static void (*proc_ml_ParamGrid_set_1maxVal_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_ml_ParamGrid_set_1maxVal_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_ml_ParamGrid_set_1maxVal_10 (a0, a1, a2, a3);
}

static void (*proc_photo_TonemapReinhard_setIntensity_10)(JNIEnv*,  jclass,  jlong,  jfloat);
void Java_org_opencv_photo_TonemapReinhard_setIntensity_10(JNIEnv* a0, jclass a1, jlong a2, jfloat a3) {
    proc_photo_TonemapReinhard_setIntensity_10 (a0, a1, a2, a3);
}

static jint (*proc_core_Core_countNonZero_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_core_Core_countNonZero_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_core_Core_countNonZero_10 (a0, a1, a2);
}

static jlong (*proc_video_SparsePyrLKOpticalFlow_create_11)(JNIEnv*,  jclass);
jlong Java_org_opencv_video_SparsePyrLKOpticalFlow_create_11(JNIEnv* a0, jclass a1) {
    return proc_video_SparsePyrLKOpticalFlow_create_11 (a0, a1);
}

static void (*proc_calib3d_StereoMatcher_setBlockSize_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_calib3d_StereoMatcher_setBlockSize_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_calib3d_StereoMatcher_setBlockSize_10 (a0, a1, a2, a3);
}

static void (*proc_features2d_FeatureDetector_read_10)(JNIEnv*,  jclass,  jlong,  jstring);
void Java_org_opencv_features2d_FeatureDetector_read_10(JNIEnv* a0, jclass a1, jlong a2, jstring a3) {
    proc_features2d_FeatureDetector_read_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_findContours_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jint,  jint);
void Java_org_opencv_imgproc_Imgproc_findContours_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jint a5, jint a6) {
    proc_imgproc_Imgproc_findContours_11 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_ml_ParamGrid_set_1minVal_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_ml_ParamGrid_set_1minVal_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_ml_ParamGrid_set_1minVal_10 (a0, a1, a2, a3);
}

static void (*proc_photo_TonemapReinhard_setLightAdaptation_10)(JNIEnv*,  jclass,  jlong,  jfloat);
void Java_org_opencv_photo_TonemapReinhard_setLightAdaptation_10(JNIEnv* a0, jclass a1, jlong a2, jfloat a3) {
    proc_photo_TonemapReinhard_setLightAdaptation_10 (a0, a1, a2, a3);
}

static jdoubleArray (*proc_imgproc_Imgproc_fitEllipseAMS_10)(JNIEnv*,  jclass,  jlong);
jdoubleArray Java_org_opencv_imgproc_Imgproc_fitEllipseAMS_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_imgproc_Imgproc_fitEllipseAMS_10 (a0, a1, a2);
}

static jfloat (*proc_core_Core_cubeRoot_10)(JNIEnv*,  jclass,  jfloat);
jfloat Java_org_opencv_core_Core_cubeRoot_10(JNIEnv* a0, jclass a1, jfloat a2) {
    return proc_core_Core_cubeRoot_10 (a0, a1, a2);
}

static jlong (*proc_ml_RTrees_create_10)(JNIEnv*,  jclass);
jlong Java_org_opencv_ml_RTrees_create_10(JNIEnv* a0, jclass a1) {
    return proc_ml_RTrees_create_10 (a0, a1);
}

static void (*proc_calib3d_StereoMatcher_setDisp12MaxDiff_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_calib3d_StereoMatcher_setDisp12MaxDiff_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_calib3d_StereoMatcher_setDisp12MaxDiff_10 (a0, a1, a2, a3);
}

static void (*proc_features2d_FeatureDetector_write_10)(JNIEnv*,  jclass,  jlong,  jstring);
void Java_org_opencv_features2d_FeatureDetector_write_10(JNIEnv* a0, jclass a1, jlong a2, jstring a3) {
    proc_features2d_FeatureDetector_write_10 (a0, a1, a2, a3);
}

static void (*proc_video_SparsePyrLKOpticalFlow_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_video_SparsePyrLKOpticalFlow_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_video_SparsePyrLKOpticalFlow_delete (a0, a1, a2);
}

static jdoubleArray (*proc_imgproc_Imgproc_fitEllipseDirect_10)(JNIEnv*,  jclass,  jlong);
jdoubleArray Java_org_opencv_imgproc_Imgproc_fitEllipseDirect_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_imgproc_Imgproc_fitEllipseDirect_10 (a0, a1, a2);
}

static jint (*proc_video_SparsePyrLKOpticalFlow_getFlags_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_video_SparsePyrLKOpticalFlow_getFlags_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_SparsePyrLKOpticalFlow_getFlags_10 (a0, a1, a2);
}

static void (*proc_calib3d_StereoMatcher_setMinDisparity_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_calib3d_StereoMatcher_setMinDisparity_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_calib3d_StereoMatcher_setMinDisparity_10 (a0, a1, a2, a3);
}

static void (*proc_core_Core_dct_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint);
void Java_org_opencv_core_Core_dct_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4) {
    proc_core_Core_dct_10 (a0, a1, a2, a3, a4);
}

static void (*proc_features2d_Features2d_drawKeypoints_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jint);
void Java_org_opencv_features2d_Features2d_drawKeypoints_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8, jint a9) {
    proc_features2d_Features2d_drawKeypoints_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static void (*proc_ml_RTrees_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_ml_RTrees_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_ml_RTrees_delete (a0, a1, a2);
}

static jdoubleArray (*proc_imgproc_Imgproc_fitEllipse_10)(JNIEnv*,  jclass,  jlong);
jdoubleArray Java_org_opencv_imgproc_Imgproc_fitEllipse_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_imgproc_Imgproc_fitEllipse_10 (a0, a1, a2);
}

static jint (*proc_ml_RTrees_getActiveVarCount_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_ml_RTrees_getActiveVarCount_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_RTrees_getActiveVarCount_10 (a0, a1, a2);
}

static jint (*proc_video_SparsePyrLKOpticalFlow_getMaxLevel_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_video_SparsePyrLKOpticalFlow_getMaxLevel_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_SparsePyrLKOpticalFlow_getMaxLevel_10 (a0, a1, a2);
}

static void (*proc_calib3d_StereoMatcher_setNumDisparities_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_calib3d_StereoMatcher_setNumDisparities_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_calib3d_StereoMatcher_setNumDisparities_10 (a0, a1, a2, a3);
}

static void (*proc_core_Core_dct_11)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_core_Core_dct_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_core_Core_dct_11 (a0, a1, a2, a3);
}

static void (*proc_features2d_Features2d_drawKeypoints_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_features2d_Features2d_drawKeypoints_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_features2d_Features2d_drawKeypoints_11 (a0, a1, a2, a3, a4);
}

static jboolean (*proc_ml_RTrees_getCalculateVarImportance_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_ml_RTrees_getCalculateVarImportance_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_RTrees_getCalculateVarImportance_10 (a0, a1, a2);
}

static jdouble (*proc_core_Core_determinant_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_core_Core_determinant_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_core_Core_determinant_10 (a0, a1, a2);
}

static jdouble (*proc_video_SparsePyrLKOpticalFlow_getMinEigThreshold_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_video_SparsePyrLKOpticalFlow_getMinEigThreshold_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_SparsePyrLKOpticalFlow_getMinEigThreshold_10 (a0, a1, a2);
}

static void (*proc_calib3d_StereoMatcher_setSpeckleRange_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_calib3d_StereoMatcher_setSpeckleRange_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_calib3d_StereoMatcher_setSpeckleRange_10 (a0, a1, a2, a3);
}

static void (*proc_features2d_Features2d_drawMatches2_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jlong,  jint);
void Java_org_opencv_features2d_Features2d_drawMatches2_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7, jdouble a8, jdouble a9, jdouble a10, jdouble a11, jdouble a12, jdouble a13, jdouble a14, jdouble a15, jlong a16, jint a17) {
    proc_features2d_Features2d_drawMatches2_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17);
}

static void (*proc_imgproc_Imgproc_fitLine_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jdouble,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_fitLine_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jdouble a5, jdouble a6, jdouble a7) {
    proc_imgproc_Imgproc_fitLine_10 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static jdoubleArray (*proc_ml_RTrees_getTermCriteria_10)(JNIEnv*,  jclass,  jlong);
jdoubleArray Java_org_opencv_ml_RTrees_getTermCriteria_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_RTrees_getTermCriteria_10 (a0, a1, a2);
}

static jdoubleArray (*proc_video_SparsePyrLKOpticalFlow_getTermCriteria_10)(JNIEnv*,  jclass,  jlong);
jdoubleArray Java_org_opencv_video_SparsePyrLKOpticalFlow_getTermCriteria_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_SparsePyrLKOpticalFlow_getTermCriteria_10 (a0, a1, a2);
}

static jint (*proc_imgproc_Imgproc_floodFill_10)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdoubleArray,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jint);
jint Java_org_opencv_imgproc_Imgproc_floodFill_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8, jdouble a9, jdoubleArray a10, jdouble a11, jdouble a12, jdouble a13, jdouble a14, jdouble a15, jdouble a16, jdouble a17, jdouble a18, jint a19) {
    return proc_imgproc_Imgproc_floodFill_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19);
}

static void (*proc_calib3d_StereoMatcher_setSpeckleWindowSize_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_calib3d_StereoMatcher_setSpeckleWindowSize_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_calib3d_StereoMatcher_setSpeckleWindowSize_10 (a0, a1, a2, a3);
}

static void (*proc_core_Core_dft_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint);
void Java_org_opencv_core_Core_dft_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5) {
    proc_core_Core_dft_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_features2d_Features2d_drawMatches2_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_features2d_Features2d_drawMatches2_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7) {
    proc_features2d_Features2d_drawMatches2_11 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static jdoubleArray (*proc_video_SparsePyrLKOpticalFlow_getWinSize_10)(JNIEnv*,  jclass,  jlong);
jdoubleArray Java_org_opencv_video_SparsePyrLKOpticalFlow_getWinSize_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_SparsePyrLKOpticalFlow_getWinSize_10 (a0, a1, a2);
}

static jint (*proc_imgproc_Imgproc_floodFill_11)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble);
jint Java_org_opencv_imgproc_Imgproc_floodFill_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8, jdouble a9) {
    return proc_imgproc_Imgproc_floodFill_11 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static jlong (*proc_calib3d_StereoSGBM_create_10)(JNIEnv*,  jclass,  jint,  jint,  jint,  jint,  jint,  jint,  jint,  jint,  jint,  jint,  jint);
jlong Java_org_opencv_calib3d_StereoSGBM_create_10(JNIEnv* a0, jclass a1, jint a2, jint a3, jint a4, jint a5, jint a6, jint a7, jint a8, jint a9, jint a10, jint a11, jint a12) {
    return proc_calib3d_StereoSGBM_create_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12);
}

static jlong (*proc_ml_RTrees_getVarImportance_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_ml_RTrees_getVarImportance_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_RTrees_getVarImportance_10 (a0, a1, a2);
}

static void (*proc_core_Core_dft_11)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_core_Core_dft_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_core_Core_dft_11 (a0, a1, a2, a3);
}

static void (*proc_features2d_Features2d_drawMatchesKnn_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jlong,  jint);
void Java_org_opencv_features2d_Features2d_drawMatchesKnn_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7, jdouble a8, jdouble a9, jdouble a10, jdouble a11, jdouble a12, jdouble a13, jdouble a14, jdouble a15, jlong a16, jint a17) {
    proc_features2d_Features2d_drawMatchesKnn_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17);
}

static jboolean (*proc_imgcodecs_Imgcodecs_imreadmulti_11)(JNIEnv*,  jclass,  jstring,  jlong);
jboolean Java_org_opencv_imgcodecs_Imgcodecs_imreadmulti_11(JNIEnv* a0, jclass a1, jstring a2, jlong a3) {
    return proc_imgcodecs_Imgcodecs_imreadmulti_11 (a0, a1, a2, a3);
}

static jboolean (*proc_objdetect_HOGDescriptor_checkDetectorSize_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_objdetect_HOGDescriptor_checkDetectorSize_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_objdetect_HOGDescriptor_checkDetectorSize_10 (a0, a1, a2);
}

static jboolean (*proc_videoio_VideoCapture_grab_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_videoio_VideoCapture_grab_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_videoio_VideoCapture_grab_10 (a0, a1, a2);
}

static jdouble (*proc_calib3d_Calib3d_calibrateCamera_10)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jlong,  jlong,  jlong,  jlong,  jint,  jint,  jint,  jdouble);
jdouble Java_org_opencv_calib3d_Calib3d_calibrateCamera_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jlong a6, jlong a7, jlong a8, jlong a9, jint a10, jint a11, jint a12, jdouble a13) {
    return proc_calib3d_Calib3d_calibrateCamera_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13);
}

static jdouble (*proc_dnn_DictValue_getRealValue_10)(JNIEnv*,  jclass,  jlong,  jint);
jdouble Java_org_opencv_dnn_DictValue_getRealValue_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    return proc_dnn_DictValue_getRealValue_10 (a0, a1, a2, a3);
}

static jdouble (*proc_ml_ANN_1MLP_getRpropDWMax_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_ml_ANN_1MLP_getRpropDWMax_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_ANN_1MLP_getRpropDWMax_10 (a0, a1, a2);
}

static jdouble (*proc_video_BackgroundSubtractorKNN_getDist2Threshold_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_video_BackgroundSubtractorKNN_getDist2Threshold_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_BackgroundSubtractorKNN_getDist2Threshold_10 (a0, a1, a2);
}

static jint (*proc_core_Mat_nPutB)(JNIEnv*,  jclass,  jlong,  jint,  jint,  jint,  jbyteArray);
jint Java_org_opencv_core_Mat_nPutB(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4, jint a5, jbyteArray a6) {
    return proc_core_Mat_nPutB (a0, a1, a2, a3, a4, a5, a6);
}

static jint (*proc_features2d_AKAZE_getDescriptorType_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_features2d_AKAZE_getDescriptorType_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_AKAZE_getDescriptorType_10 (a0, a1, a2);
}

static jint (*proc_photo_AlignMTB_getExcludeRange_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_photo_AlignMTB_getExcludeRange_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_photo_AlignMTB_getExcludeRange_10 (a0, a1, a2);
}

static void (*proc_imgproc_CLAHE_setTilesGridSize_10)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble);
void Java_org_opencv_imgproc_CLAHE_setTilesGridSize_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4) {
    proc_imgproc_CLAHE_setTilesGridSize_10 (a0, a1, a2, a3, a4);
}

static jlong (*proc_calib3d_StereoSGBM_create_11)(JNIEnv*,  jclass);
jlong Java_org_opencv_calib3d_StereoSGBM_create_11(JNIEnv* a0, jclass a1) {
    return proc_calib3d_StereoSGBM_create_11 (a0, a1);
}

static jlong (*proc_imgproc_Imgproc_getAffineTransform_10)(JNIEnv*,  jclass,  jlong,  jlong);
jlong Java_org_opencv_imgproc_Imgproc_getAffineTransform_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_imgproc_Imgproc_getAffineTransform_10 (a0, a1, a2, a3);
}

static void (*proc_core_Core_divide_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jdouble,  jint);
void Java_org_opencv_core_Core_divide_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jdouble a5, jint a6) {
    proc_core_Core_divide_10 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_features2d_Features2d_drawMatchesKnn_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_features2d_Features2d_drawMatchesKnn_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7) {
    proc_features2d_Features2d_drawMatchesKnn_11 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_ml_RTrees_getVotes_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jint);
void Java_org_opencv_ml_RTrees_getVotes_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jint a5) {
    proc_ml_RTrees_getVotes_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_video_SparsePyrLKOpticalFlow_setFlags_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_video_SparsePyrLKOpticalFlow_setFlags_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_video_SparsePyrLKOpticalFlow_setFlags_10 (a0, a1, a2, a3);
}

static jlong (*proc_imgproc_Imgproc_getDefaultNewCameraMatrix_10)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jboolean);
jlong Java_org_opencv_imgproc_Imgproc_getDefaultNewCameraMatrix_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jboolean a5) {
    return proc_imgproc_Imgproc_getDefaultNewCameraMatrix_10 (a0, a1, a2, a3, a4, a5);
}

static jlong (*proc_ml_RTrees_load_10)(JNIEnv*,  jclass,  jstring,  jstring);
jlong Java_org_opencv_ml_RTrees_load_10(JNIEnv* a0, jclass a1, jstring a2, jstring a3) {
    return proc_ml_RTrees_load_10 (a0, a1, a2, a3);
}

static void (*proc_calib3d_StereoSGBM_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_calib3d_StereoSGBM_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_calib3d_StereoSGBM_delete (a0, a1, a2);
}

static void (*proc_core_Core_divide_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jdouble);
void Java_org_opencv_core_Core_divide_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jdouble a5) {
    proc_core_Core_divide_11 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_features2d_Features2d_drawMatches_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jlong,  jint);
void Java_org_opencv_features2d_Features2d_drawMatches_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7, jdouble a8, jdouble a9, jdouble a10, jdouble a11, jdouble a12, jdouble a13, jdouble a14, jdouble a15, jlong a16, jint a17) {
    proc_features2d_Features2d_drawMatches_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17);
}

static void (*proc_video_SparsePyrLKOpticalFlow_setMaxLevel_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_video_SparsePyrLKOpticalFlow_setMaxLevel_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_video_SparsePyrLKOpticalFlow_setMaxLevel_10 (a0, a1, a2, a3);
}

static jint (*proc_calib3d_StereoSGBM_getMode_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_calib3d_StereoSGBM_getMode_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_calib3d_StereoSGBM_getMode_10 (a0, a1, a2);
}

static jlong (*proc_imgproc_Imgproc_getDefaultNewCameraMatrix_11)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_imgproc_Imgproc_getDefaultNewCameraMatrix_11(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_imgproc_Imgproc_getDefaultNewCameraMatrix_11 (a0, a1, a2);
}

static jlong (*proc_ml_RTrees_load_11)(JNIEnv*,  jclass,  jstring);
jlong Java_org_opencv_ml_RTrees_load_11(JNIEnv* a0, jclass a1, jstring a2) {
    return proc_ml_RTrees_load_11 (a0, a1, a2);
}

static void (*proc_core_Core_divide_12)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_core_Core_divide_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_core_Core_divide_12 (a0, a1, a2, a3, a4);
}

static void (*proc_features2d_Features2d_drawMatches_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_features2d_Features2d_drawMatches_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7) {
    proc_features2d_Features2d_drawMatches_11 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_video_SparsePyrLKOpticalFlow_setMinEigThreshold_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_video_SparsePyrLKOpticalFlow_setMinEigThreshold_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_video_SparsePyrLKOpticalFlow_setMinEigThreshold_10 (a0, a1, a2, a3);
}

static jint (*proc_calib3d_StereoSGBM_getP1_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_calib3d_StereoSGBM_getP1_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_calib3d_StereoSGBM_getP1_10 (a0, a1, a2);
}

static jlong (*proc_features2d_FlannBasedMatcher_FlannBasedMatcher_10)(JNIEnv*,  jclass);
jlong Java_org_opencv_features2d_FlannBasedMatcher_FlannBasedMatcher_10(JNIEnv* a0, jclass a1) {
    return proc_features2d_FlannBasedMatcher_FlannBasedMatcher_10 (a0, a1);
}

static void (*proc_core_Core_divide_13)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jlong,  jdouble,  jint);
void Java_org_opencv_core_Core_divide_13(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jlong a7, jdouble a8, jint a9) {
    proc_core_Core_divide_13 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static void (*proc_imgproc_Imgproc_getDerivKernels_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint,  jint,  jboolean,  jint);
void Java_org_opencv_imgproc_Imgproc_getDerivKernels_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5, jint a6, jboolean a7, jint a8) {
    proc_imgproc_Imgproc_getDerivKernels_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static void (*proc_ml_RTrees_setActiveVarCount_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_ml_RTrees_setActiveVarCount_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_ml_RTrees_setActiveVarCount_10 (a0, a1, a2, a3);
}

static void (*proc_video_SparsePyrLKOpticalFlow_setTermCriteria_10)(JNIEnv*,  jclass,  jlong,  jint,  jint,  jdouble);
void Java_org_opencv_video_SparsePyrLKOpticalFlow_setTermCriteria_10(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4, jdouble a5) {
    proc_video_SparsePyrLKOpticalFlow_setTermCriteria_10 (a0, a1, a2, a3, a4, a5);
}

static jint (*proc_calib3d_StereoSGBM_getP2_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_calib3d_StereoSGBM_getP2_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_calib3d_StereoSGBM_getP2_10 (a0, a1, a2);
}

static jlong (*proc_features2d_FlannBasedMatcher_create_10)(JNIEnv*,  jclass);
jlong Java_org_opencv_features2d_FlannBasedMatcher_create_10(JNIEnv* a0, jclass a1) {
    return proc_features2d_FlannBasedMatcher_create_10 (a0, a1);
}

static void (*proc_core_Core_divide_14)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jlong,  jdouble);
void Java_org_opencv_core_Core_divide_14(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jlong a7, jdouble a8) {
    proc_core_Core_divide_14 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static void (*proc_imgproc_Imgproc_getDerivKernels_11)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint,  jint);
void Java_org_opencv_imgproc_Imgproc_getDerivKernels_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5, jint a6) {
    proc_imgproc_Imgproc_getDerivKernels_11 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_ml_RTrees_setCalculateVarImportance_10)(JNIEnv*,  jclass,  jlong,  jboolean);
void Java_org_opencv_ml_RTrees_setCalculateVarImportance_10(JNIEnv* a0, jclass a1, jlong a2, jboolean a3) {
    proc_ml_RTrees_setCalculateVarImportance_10 (a0, a1, a2, a3);
}

static void (*proc_video_SparsePyrLKOpticalFlow_setWinSize_10)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble);
void Java_org_opencv_video_SparsePyrLKOpticalFlow_setWinSize_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4) {
    proc_video_SparsePyrLKOpticalFlow_setWinSize_10 (a0, a1, a2, a3, a4);
}

static jdoubleArray (*proc_video_Video_CamShift_10)(JNIEnv*,  jclass,  jlong,  jint,  jint,  jint,  jint,  jdoubleArray,  jint,  jint,  jdouble);
jdoubleArray Java_org_opencv_video_Video_CamShift_10(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4, jint a5, jint a6, jdoubleArray a7, jint a8, jint a9, jdouble a10) {
    return proc_video_Video_CamShift_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10);
}

static jint (*proc_calib3d_StereoSGBM_getPreFilterCap_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_calib3d_StereoSGBM_getPreFilterCap_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_calib3d_StereoSGBM_getPreFilterCap_10 (a0, a1, a2);
}

static jlong (*proc_imgproc_Imgproc_getGaborKernel_10)(JNIEnv*,  jclass,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jint);
jlong Java_org_opencv_imgproc_Imgproc_getGaborKernel_10(JNIEnv* a0, jclass a1, jdouble a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8, jint a9) {
    return proc_imgproc_Imgproc_getGaborKernel_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static void (*proc_core_Core_divide_15)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jlong);
void Java_org_opencv_core_Core_divide_15(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jlong a7) {
    proc_core_Core_divide_15 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_features2d_FlannBasedMatcher_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_features2d_FlannBasedMatcher_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_features2d_FlannBasedMatcher_delete (a0, a1, a2);
}

static void (*proc_ml_RTrees_setTermCriteria_10)(JNIEnv*,  jclass,  jlong,  jint,  jint,  jdouble);
void Java_org_opencv_ml_RTrees_setTermCriteria_10(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4, jdouble a5) {
    proc_ml_RTrees_setTermCriteria_10 (a0, a1, a2, a3, a4, a5);
}

static jint (*proc_calib3d_StereoSGBM_getUniquenessRatio_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_calib3d_StereoSGBM_getUniquenessRatio_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_calib3d_StereoSGBM_getUniquenessRatio_10 (a0, a1, a2);
}

static jint (*proc_video_Video_buildOpticalFlowPyramid_10)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jint,  jboolean,  jint,  jint,  jboolean);
jint Java_org_opencv_video_Video_buildOpticalFlowPyramid_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jint a6, jboolean a7, jint a8, jint a9, jboolean a10) {
    return proc_video_Video_buildOpticalFlowPyramid_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10);
}

static jlong (*proc_features2d_GFTTDetector_create_10)(JNIEnv*,  jclass,  jint,  jdouble,  jdouble,  jint,  jint,  jboolean,  jdouble);
jlong Java_org_opencv_features2d_GFTTDetector_create_10(JNIEnv* a0, jclass a1, jint a2, jdouble a3, jdouble a4, jint a5, jint a6, jboolean a7, jdouble a8) {
    return proc_features2d_GFTTDetector_create_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static jlong (*proc_imgproc_Imgproc_getGaborKernel_11)(JNIEnv*,  jclass,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble);
jlong Java_org_opencv_imgproc_Imgproc_getGaborKernel_11(JNIEnv* a0, jclass a1, jdouble a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7) {
    return proc_imgproc_Imgproc_getGaborKernel_11 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static jlong (*proc_ml_SVM_create_10)(JNIEnv*,  jclass);
jlong Java_org_opencv_ml_SVM_create_10(JNIEnv* a0, jclass a1) {
    return proc_ml_SVM_create_10 (a0, a1);
}

static void (*proc_core_Core_divide_16)(JNIEnv*,  jclass,  jdouble,  jlong,  jlong,  jint);
void Java_org_opencv_core_Core_divide_16(JNIEnv* a0, jclass a1, jdouble a2, jlong a3, jlong a4, jint a5) {
    proc_core_Core_divide_16 (a0, a1, a2, a3, a4, a5);
}

static jint (*proc_video_Video_buildOpticalFlowPyramid_11)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jint);
jint Java_org_opencv_video_Video_buildOpticalFlowPyramid_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jint a6) {
    return proc_video_Video_buildOpticalFlowPyramid_11 (a0, a1, a2, a3, a4, a5, a6);
}

static jlong (*proc_features2d_GFTTDetector_create_11)(JNIEnv*,  jclass,  jint,  jdouble,  jdouble,  jint,  jint);
jlong Java_org_opencv_features2d_GFTTDetector_create_11(JNIEnv* a0, jclass a1, jint a2, jdouble a3, jdouble a4, jint a5, jint a6) {
    return proc_features2d_GFTTDetector_create_11 (a0, a1, a2, a3, a4, a5, a6);
}

static jlong (*proc_imgproc_Imgproc_getGaussianKernel_10)(JNIEnv*,  jclass,  jint,  jdouble,  jint);
jlong Java_org_opencv_imgproc_Imgproc_getGaussianKernel_10(JNIEnv* a0, jclass a1, jint a2, jdouble a3, jint a4) {
    return proc_imgproc_Imgproc_getGaussianKernel_10 (a0, a1, a2, a3, a4);
}

static void (*proc_calib3d_StereoSGBM_setMode_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_calib3d_StereoSGBM_setMode_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_calib3d_StereoSGBM_setMode_10 (a0, a1, a2, a3);
}

static void (*proc_core_Core_divide_17)(JNIEnv*,  jclass,  jdouble,  jlong,  jlong);
void Java_org_opencv_core_Core_divide_17(JNIEnv* a0, jclass a1, jdouble a2, jlong a3, jlong a4) {
    proc_core_Core_divide_17 (a0, a1, a2, a3, a4);
}

static void (*proc_ml_SVM_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_ml_SVM_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_ml_SVM_delete (a0, a1, a2);
}

static jdouble (*proc_ml_SVM_getC_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_ml_SVM_getC_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_SVM_getC_10 (a0, a1, a2);
}

static jlong (*proc_features2d_GFTTDetector_create_12)(JNIEnv*,  jclass,  jint,  jdouble,  jdouble,  jint,  jboolean,  jdouble);
jlong Java_org_opencv_features2d_GFTTDetector_create_12(JNIEnv* a0, jclass a1, jint a2, jdouble a3, jdouble a4, jint a5, jboolean a6, jdouble a7) {
    return proc_features2d_GFTTDetector_create_12 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static jlong (*proc_imgproc_Imgproc_getGaussianKernel_11)(JNIEnv*,  jclass,  jint,  jdouble);
jlong Java_org_opencv_imgproc_Imgproc_getGaussianKernel_11(JNIEnv* a0, jclass a1, jint a2, jdouble a3) {
    return proc_imgproc_Imgproc_getGaussianKernel_11 (a0, a1, a2, a3);
}

static void (*proc_calib3d_StereoSGBM_setP1_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_calib3d_StereoSGBM_setP1_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_calib3d_StereoSGBM_setP1_10 (a0, a1, a2, a3);
}

static void (*proc_core_Core_eigenNonSymmetric_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_core_Core_eigenNonSymmetric_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_core_Core_eigenNonSymmetric_10 (a0, a1, a2, a3, a4);
}

static void (*proc_video_Video_calcOpticalFlowFarneback_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jdouble,  jint,  jint,  jint,  jint,  jdouble,  jint);
void Java_org_opencv_video_Video_calcOpticalFlowFarneback_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jdouble a5, jint a6, jint a7, jint a8, jint a9, jdouble a10, jint a11) {
    proc_video_Video_calcOpticalFlowFarneback_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11);
}

static jboolean (*proc_core_Core_eigen_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
jboolean Java_org_opencv_core_Core_eigen_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    return proc_core_Core_eigen_10 (a0, a1, a2, a3, a4);
}

static jlong (*proc_features2d_GFTTDetector_create_13)(JNIEnv*,  jclass);
jlong Java_org_opencv_features2d_GFTTDetector_create_13(JNIEnv* a0, jclass a1) {
    return proc_features2d_GFTTDetector_create_13 (a0, a1);
}

static jlong (*proc_imgproc_Imgproc_getPerspectiveTransform_10)(JNIEnv*,  jclass,  jlong,  jlong);
jlong Java_org_opencv_imgproc_Imgproc_getPerspectiveTransform_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_imgproc_Imgproc_getPerspectiveTransform_10 (a0, a1, a2, a3);
}

static jlong (*proc_ml_SVM_getClassWeights_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_ml_SVM_getClassWeights_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_SVM_getClassWeights_10 (a0, a1, a2);
}

static void (*proc_calib3d_StereoSGBM_setP2_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_calib3d_StereoSGBM_setP2_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_calib3d_StereoSGBM_setP2_10 (a0, a1, a2, a3);
}

static void (*proc_video_Video_calcOpticalFlowPyrLK_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jdouble,  jdouble,  jint,  jint,  jint,  jdouble,  jint,  jdouble);
void Java_org_opencv_video_Video_calcOpticalFlowPyrLK_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7, jdouble a8, jdouble a9, jint a10, jint a11, jint a12, jdouble a13, jint a14, jdouble a15) {
    proc_video_Video_calcOpticalFlowPyrLK_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15);
}

static jboolean (*proc_imgcodecs_Imgcodecs_imwrite_10)(JNIEnv*,  jclass,  jstring,  jlong,  jlong);
jboolean Java_org_opencv_imgcodecs_Imgcodecs_imwrite_10(JNIEnv* a0, jclass a1, jstring a2, jlong a3, jlong a4) {
    return proc_imgcodecs_Imgcodecs_imwrite_10 (a0, a1, a2, a3, a4);
}

static jboolean (*proc_videoio_VideoCapture_isOpened_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_videoio_VideoCapture_isOpened_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_videoio_VideoCapture_isOpened_10 (a0, a1, a2);
}

static jdouble (*proc_calib3d_Calib3d_calibrateCamera_11)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jlong,  jlong,  jlong,  jlong,  jint);
jdouble Java_org_opencv_calib3d_Calib3d_calibrateCamera_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jlong a6, jlong a7, jlong a8, jlong a9, jint a10) {
    return proc_calib3d_Calib3d_calibrateCamera_11 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10);
}

static jdouble (*proc_dnn_DictValue_getRealValue_11)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_dnn_DictValue_getRealValue_11(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_dnn_DictValue_getRealValue_11 (a0, a1, a2);
}

static jdouble (*proc_ml_ANN_1MLP_getRpropDWMin_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_ml_ANN_1MLP_getRpropDWMin_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_ANN_1MLP_getRpropDWMin_10 (a0, a1, a2);
}

static jint (*proc_core_Mat_nPutD)(JNIEnv*,  jclass,  jlong,  jint,  jint,  jint,  jdoubleArray);
jint Java_org_opencv_core_Mat_nPutD(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4, jint a5, jdoubleArray a6) {
    return proc_core_Mat_nPutD (a0, a1, a2, a3, a4, a5, a6);
}

static jint (*proc_features2d_AKAZE_getDiffusivity_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_features2d_AKAZE_getDiffusivity_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_AKAZE_getDiffusivity_10 (a0, a1, a2);
}

static jint (*proc_photo_AlignMTB_getMaxBits_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_photo_AlignMTB_getMaxBits_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_photo_AlignMTB_getMaxBits_10 (a0, a1, a2);
}

static jint (*proc_video_BackgroundSubtractorKNN_getHistory_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_video_BackgroundSubtractorKNN_getHistory_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_BackgroundSubtractorKNN_getHistory_10 (a0, a1, a2);
}

static void (*proc_imgproc_Imgproc_Canny_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jdouble,  jdouble,  jboolean);
void Java_org_opencv_imgproc_Imgproc_Canny_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jdouble a5, jdouble a6, jboolean a7) {
    proc_imgproc_Imgproc_Canny_10 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_objdetect_HOGDescriptor_computeGradient_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jdouble,  jdouble,  jdouble,  jdouble);
void Java_org_opencv_objdetect_HOGDescriptor_computeGradient_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jdouble a6, jdouble a7, jdouble a8, jdouble a9) {
    proc_objdetect_HOGDescriptor_computeGradient_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static jboolean (*proc_core_Core_eigen_11)(JNIEnv*,  jclass,  jlong,  jlong);
jboolean Java_org_opencv_core_Core_eigen_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_core_Core_eigen_11 (a0, a1, a2, a3);
}

static jdouble (*proc_ml_SVM_getCoef0_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_ml_SVM_getCoef0_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_SVM_getCoef0_10 (a0, a1, a2);
}

static void (*proc_calib3d_StereoSGBM_setPreFilterCap_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_calib3d_StereoSGBM_setPreFilterCap_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_calib3d_StereoSGBM_setPreFilterCap_10 (a0, a1, a2, a3);
}

static void (*proc_features2d_GFTTDetector_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_features2d_GFTTDetector_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_features2d_GFTTDetector_delete (a0, a1, a2);
}

static void (*proc_imgproc_Imgproc_getRectSubPix_10)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jlong,  jint);
void Java_org_opencv_imgproc_Imgproc_getRectSubPix_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jlong a7, jint a8) {
    proc_imgproc_Imgproc_getRectSubPix_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static void (*proc_video_Video_calcOpticalFlowPyrLK_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jdouble,  jdouble,  jint);
void Java_org_opencv_video_Video_calcOpticalFlowPyrLK_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7, jdouble a8, jdouble a9, jint a10) {
    proc_video_Video_calcOpticalFlowPyrLK_11 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10);
}

static jdouble (*proc_ml_SVM_getDecisionFunction_10)(JNIEnv*,  jclass,  jlong,  jint,  jlong,  jlong);
jdouble Java_org_opencv_ml_SVM_getDecisionFunction_10(JNIEnv* a0, jclass a1, jlong a2, jint a3, jlong a4, jlong a5) {
    return proc_ml_SVM_getDecisionFunction_10 (a0, a1, a2, a3, a4, a5);
}

static jint (*proc_features2d_GFTTDetector_getBlockSize_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_features2d_GFTTDetector_getBlockSize_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_GFTTDetector_getBlockSize_10 (a0, a1, a2);
}

static void (*proc_calib3d_StereoSGBM_setUniquenessRatio_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_calib3d_StereoSGBM_setUniquenessRatio_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_calib3d_StereoSGBM_setUniquenessRatio_10 (a0, a1, a2, a3);
}

static void (*proc_core_Core_exp_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_core_Core_exp_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_core_Core_exp_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_getRectSubPix_11)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jlong);
void Java_org_opencv_imgproc_Imgproc_getRectSubPix_11(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jlong a7) {
    proc_imgproc_Imgproc_getRectSubPix_11 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_video_Video_calcOpticalFlowPyrLK_12)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_video_Video_calcOpticalFlowPyrLK_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7) {
    proc_video_Video_calcOpticalFlowPyrLK_12 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static jlong (*proc_imgproc_Imgproc_getRotationMatrix2D_10)(JNIEnv*,  jclass,  jdouble,  jdouble,  jdouble,  jdouble);
jlong Java_org_opencv_imgproc_Imgproc_getRotationMatrix2D_10(JNIEnv* a0, jclass a1, jdouble a2, jdouble a3, jdouble a4, jdouble a5) {
    return proc_imgproc_Imgproc_getRotationMatrix2D_10 (a0, a1, a2, a3, a4, a5);
}

static jlong (*proc_ml_SVM_getDefaultGridPtr_10)(JNIEnv*,  jclass,  jint);
jlong Java_org_opencv_ml_SVM_getDefaultGridPtr_10(JNIEnv* a0, jclass a1, jint a2) {
    return proc_ml_SVM_getDefaultGridPtr_10 (a0, a1, a2);
}

static jlong (*proc_video_Video_createBackgroundSubtractorKNN_10)(JNIEnv*,  jclass,  jint,  jdouble,  jboolean);
jlong Java_org_opencv_video_Video_createBackgroundSubtractorKNN_10(JNIEnv* a0, jclass a1, jint a2, jdouble a3, jboolean a4) {
    return proc_video_Video_createBackgroundSubtractorKNN_10 (a0, a1, a2, a3, a4);
}

static jstring (*proc_features2d_GFTTDetector_getDefaultName_10)(JNIEnv*,  jclass,  jlong);
jstring Java_org_opencv_features2d_GFTTDetector_getDefaultName_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_GFTTDetector_getDefaultName_10 (a0, a1, a2);
}

static void (*proc_core_Core_extractChannel_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint);
void Java_org_opencv_core_Core_extractChannel_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4) {
    proc_core_Core_extractChannel_10 (a0, a1, a2, a3, a4);
}

static jboolean (*proc_features2d_GFTTDetector_getHarrisDetector_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_features2d_GFTTDetector_getHarrisDetector_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_GFTTDetector_getHarrisDetector_10 (a0, a1, a2);
}

static jdouble (*proc_ml_SVM_getDegree_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_ml_SVM_getDegree_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_SVM_getDegree_10 (a0, a1, a2);
}

static jfloat (*proc_core_Core_fastAtan2_10)(JNIEnv*,  jclass,  jfloat,  jfloat);
jfloat Java_org_opencv_core_Core_fastAtan2_10(JNIEnv* a0, jclass a1, jfloat a2, jfloat a3) {
    return proc_core_Core_fastAtan2_10 (a0, a1, a2, a3);
}

static jlong (*proc_imgproc_Imgproc_getStructuringElement_10)(JNIEnv*,  jclass,  jint,  jdouble,  jdouble,  jdouble,  jdouble);
jlong Java_org_opencv_imgproc_Imgproc_getStructuringElement_10(JNIEnv* a0, jclass a1, jint a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6) {
    return proc_imgproc_Imgproc_getStructuringElement_10 (a0, a1, a2, a3, a4, a5, a6);
}

static jlong (*proc_video_Video_createBackgroundSubtractorKNN_11)(JNIEnv*,  jclass);
jlong Java_org_opencv_video_Video_createBackgroundSubtractorKNN_11(JNIEnv* a0, jclass a1) {
    return proc_video_Video_createBackgroundSubtractorKNN_11 (a0, a1);
}

static jdouble (*proc_features2d_GFTTDetector_getK_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_features2d_GFTTDetector_getK_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_GFTTDetector_getK_10 (a0, a1, a2);
}

static jdouble (*proc_ml_SVM_getGamma_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_ml_SVM_getGamma_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_SVM_getGamma_10 (a0, a1, a2);
}

static jlong (*proc_imgproc_Imgproc_getStructuringElement_11)(JNIEnv*,  jclass,  jint,  jdouble,  jdouble);
jlong Java_org_opencv_imgproc_Imgproc_getStructuringElement_11(JNIEnv* a0, jclass a1, jint a2, jdouble a3, jdouble a4) {
    return proc_imgproc_Imgproc_getStructuringElement_11 (a0, a1, a2, a3, a4);
}

static jlong (*proc_video_Video_createBackgroundSubtractorMOG2_10)(JNIEnv*,  jclass,  jint,  jdouble,  jboolean);
jlong Java_org_opencv_video_Video_createBackgroundSubtractorMOG2_10(JNIEnv* a0, jclass a1, jint a2, jdouble a3, jboolean a4) {
    return proc_video_Video_createBackgroundSubtractorMOG2_10 (a0, a1, a2, a3, a4);
}

static void (*proc_core_Core_findNonZero_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_core_Core_findNonZero_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_core_Core_findNonZero_10 (a0, a1, a2, a3);
}

static jint (*proc_features2d_GFTTDetector_getMaxFeatures_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_features2d_GFTTDetector_getMaxFeatures_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_GFTTDetector_getMaxFeatures_10 (a0, a1, a2);
}

static jint (*proc_ml_SVM_getKernelType_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_ml_SVM_getKernelType_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_SVM_getKernelType_10 (a0, a1, a2);
}

static jlong (*proc_video_Video_createBackgroundSubtractorMOG2_11)(JNIEnv*,  jclass);
jlong Java_org_opencv_video_Video_createBackgroundSubtractorMOG2_11(JNIEnv* a0, jclass a1) {
    return proc_video_Video_createBackgroundSubtractorMOG2_11 (a0, a1);
}

static void (*proc_core_Core_flip_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint);
void Java_org_opencv_core_Core_flip_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4) {
    proc_core_Core_flip_10 (a0, a1, a2, a3, a4);
}

static void (*proc_imgproc_Imgproc_goodFeaturesToTrack_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jdouble,  jdouble,  jlong,  jint,  jint,  jboolean,  jdouble);
void Java_org_opencv_imgproc_Imgproc_goodFeaturesToTrack_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jdouble a5, jdouble a6, jlong a7, jint a8, jint a9, jboolean a10, jdouble a11) {
    proc_imgproc_Imgproc_goodFeaturesToTrack_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11);
}

static jdouble (*proc_features2d_GFTTDetector_getMinDistance_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_features2d_GFTTDetector_getMinDistance_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_GFTTDetector_getMinDistance_10 (a0, a1, a2);
}

static jdouble (*proc_ml_SVM_getNu_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_ml_SVM_getNu_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_SVM_getNu_10 (a0, a1, a2);
}

static jlong (*proc_video_Video_createOptFlow_1DualTVL1_10)(JNIEnv*,  jclass);
jlong Java_org_opencv_video_Video_createOptFlow_1DualTVL1_10(JNIEnv* a0, jclass a1) {
    return proc_video_Video_createOptFlow_1DualTVL1_10 (a0, a1);
}

static void (*proc_core_Core_gemm_10)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jlong,  jdouble,  jlong,  jint);
void Java_org_opencv_core_Core_gemm_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jlong a5, jdouble a6, jlong a7, jint a8) {
    proc_core_Core_gemm_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static void (*proc_imgproc_Imgproc_goodFeaturesToTrack_11)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jdouble,  jdouble,  jlong,  jint,  jint);
void Java_org_opencv_imgproc_Imgproc_goodFeaturesToTrack_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jdouble a5, jdouble a6, jlong a7, jint a8, jint a9) {
    proc_imgproc_Imgproc_goodFeaturesToTrack_11 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static jdouble (*proc_features2d_GFTTDetector_getQualityLevel_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_features2d_GFTTDetector_getQualityLevel_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_GFTTDetector_getQualityLevel_10 (a0, a1, a2);
}

static jdouble (*proc_ml_SVM_getP_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_ml_SVM_getP_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_SVM_getP_10 (a0, a1, a2);
}

static jlong (*proc_video_Video_estimateRigidTransform_10)(JNIEnv*,  jclass,  jlong,  jlong,  jboolean);
jlong Java_org_opencv_video_Video_estimateRigidTransform_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jboolean a4) {
    return proc_video_Video_estimateRigidTransform_10 (a0, a1, a2, a3, a4);
}

static void (*proc_core_Core_gemm_11)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jlong,  jdouble,  jlong);
void Java_org_opencv_core_Core_gemm_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jlong a5, jdouble a6, jlong a7) {
    proc_core_Core_gemm_11 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_imgproc_Imgproc_goodFeaturesToTrack_12)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jdouble,  jdouble,  jlong,  jint,  jboolean,  jdouble);
void Java_org_opencv_imgproc_Imgproc_goodFeaturesToTrack_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jdouble a5, jdouble a6, jlong a7, jint a8, jboolean a9, jdouble a10) {
    proc_imgproc_Imgproc_goodFeaturesToTrack_12 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10);
}

static jdouble (*proc_video_Video_findTransformECC_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jint,  jint,  jint,  jdouble,  jlong);
jdouble Java_org_opencv_video_Video_findTransformECC_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jint a5, jint a6, jint a7, jdouble a8, jlong a9) {
    return proc_video_Video_findTransformECC_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static jlong (*proc_ml_SVM_getSupportVectors_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_ml_SVM_getSupportVectors_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_SVM_getSupportVectors_10 (a0, a1, a2);
}

static jstring (*proc_core_Core_getBuildInformation_10)(JNIEnv*,  jclass);
jstring Java_org_opencv_core_Core_getBuildInformation_10(JNIEnv* a0, jclass a1) {
    return proc_core_Core_getBuildInformation_10 (a0, a1);
}

static void (*proc_features2d_GFTTDetector_setBlockSize_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_features2d_GFTTDetector_setBlockSize_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_features2d_GFTTDetector_setBlockSize_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_goodFeaturesToTrack_13)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_goodFeaturesToTrack_13(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jdouble a5, jdouble a6) {
    proc_imgproc_Imgproc_goodFeaturesToTrack_13 (a0, a1, a2, a3, a4, a5, a6);
}

static jdouble (*proc_video_Video_findTransformECC_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jint);
jdouble Java_org_opencv_video_Video_findTransformECC_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jint a5) {
    return proc_video_Video_findTransformECC_11 (a0, a1, a2, a3, a4, a5);
}

static jdoubleArray (*proc_ml_SVM_getTermCriteria_10)(JNIEnv*,  jclass,  jlong);
jdoubleArray Java_org_opencv_ml_SVM_getTermCriteria_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_SVM_getTermCriteria_10 (a0, a1, a2);
}

static jlong (*proc_core_Core_getCPUTickCount_10)(JNIEnv*,  jclass);
jlong Java_org_opencv_core_Core_getCPUTickCount_10(JNIEnv* a0, jclass a1) {
    return proc_core_Core_getCPUTickCount_10 (a0, a1);
}

static void (*proc_features2d_GFTTDetector_setHarrisDetector_10)(JNIEnv*,  jclass,  jlong,  jboolean);
void Java_org_opencv_features2d_GFTTDetector_setHarrisDetector_10(JNIEnv* a0, jclass a1, jlong a2, jboolean a3) {
    proc_features2d_GFTTDetector_setHarrisDetector_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_grabCut_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint,  jint,  jint,  jlong,  jlong,  jint,  jint);
void Java_org_opencv_imgproc_Imgproc_grabCut_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5, jint a6, jint a7, jlong a8, jlong a9, jint a10, jint a11) {
    proc_imgproc_Imgproc_grabCut_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11);
}

static jboolean (*proc_imgcodecs_Imgcodecs_imwrite_11)(JNIEnv*,  jclass,  jstring,  jlong);
jboolean Java_org_opencv_imgcodecs_Imgcodecs_imwrite_11(JNIEnv* a0, jclass a1, jstring a2, jlong a3) {
    return proc_imgcodecs_Imgcodecs_imwrite_11 (a0, a1, a2, a3);
}

static jboolean (*proc_videoio_VideoCapture_open_10)(JNIEnv*,  jclass,  jlong,  jstring,  jint);
jboolean Java_org_opencv_videoio_VideoCapture_open_10(JNIEnv* a0, jclass a1, jlong a2, jstring a3, jint a4) {
    return proc_videoio_VideoCapture_open_10 (a0, a1, a2, a3, a4);
}

static jdouble (*proc_calib3d_Calib3d_calibrateCamera_12)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jlong,  jlong,  jlong,  jlong);
jdouble Java_org_opencv_calib3d_Calib3d_calibrateCamera_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jlong a6, jlong a7, jlong a8, jlong a9) {
    return proc_calib3d_Calib3d_calibrateCamera_12 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static jdouble (*proc_ml_ANN_1MLP_getRpropDWMinus_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_ml_ANN_1MLP_getRpropDWMinus_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_ANN_1MLP_getRpropDWMinus_10 (a0, a1, a2);
}

static jint (*proc_core_Mat_nPutF)(JNIEnv*,  jclass,  jlong,  jint,  jint,  jint,  jfloatArray);
jint Java_org_opencv_core_Mat_nPutF(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4, jint a5, jfloatArray a6) {
    return proc_core_Mat_nPutF (a0, a1, a2, a3, a4, a5, a6);
}

static jint (*proc_features2d_AKAZE_getNOctaveLayers_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_features2d_AKAZE_getNOctaveLayers_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_AKAZE_getNOctaveLayers_10 (a0, a1, a2);
}

static jint (*proc_video_BackgroundSubtractorKNN_getNSamples_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_video_BackgroundSubtractorKNN_getNSamples_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_BackgroundSubtractorKNN_getNSamples_10 (a0, a1, a2);
}

static jstring (*proc_dnn_DictValue_getStringValue_10)(JNIEnv*,  jclass,  jlong,  jint);
jstring Java_org_opencv_dnn_DictValue_getStringValue_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    return proc_dnn_DictValue_getStringValue_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_Canny_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_Canny_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jdouble a5, jdouble a6) {
    proc_imgproc_Imgproc_Canny_11 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_objdetect_HOGDescriptor_computeGradient_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_objdetect_HOGDescriptor_computeGradient_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_objdetect_HOGDescriptor_computeGradient_11 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_photo_AlignMTB_process_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_photo_AlignMTB_process_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6) {
    proc_photo_AlignMTB_process_10 (a0, a1, a2, a3, a4, a5, a6);
}

static jdouble (*proc_video_Video_findTransformECC_12)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
jdouble Java_org_opencv_video_Video_findTransformECC_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    return proc_video_Video_findTransformECC_12 (a0, a1, a2, a3, a4);
}

static jint (*proc_ml_SVM_getType_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_ml_SVM_getType_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_SVM_getType_10 (a0, a1, a2);
}

static jstring (*proc_core_Core_getIppVersion_10)(JNIEnv*,  jclass);
jstring Java_org_opencv_core_Core_getIppVersion_10(JNIEnv* a0, jclass a1) {
    return proc_core_Core_getIppVersion_10 (a0, a1);
}

static void (*proc_features2d_GFTTDetector_setK_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_features2d_GFTTDetector_setK_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_features2d_GFTTDetector_setK_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_grabCut_11)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint,  jint,  jint,  jlong,  jlong,  jint);
void Java_org_opencv_imgproc_Imgproc_grabCut_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5, jint a6, jint a7, jlong a8, jlong a9, jint a10) {
    proc_imgproc_Imgproc_grabCut_11 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10);
}

static jint (*proc_core_Core_getNumThreads_10)(JNIEnv*,  jclass);
jint Java_org_opencv_core_Core_getNumThreads_10(JNIEnv* a0, jclass a1) {
    return proc_core_Core_getNumThreads_10 (a0, a1);
}

static jint (*proc_video_Video_meanShift_10)(JNIEnv*,  jclass,  jlong,  jint,  jint,  jint,  jint,  jdoubleArray,  jint,  jint,  jdouble);
jint Java_org_opencv_video_Video_meanShift_10(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4, jint a5, jint a6, jdoubleArray a7, jint a8, jint a9, jdouble a10) {
    return proc_video_Video_meanShift_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10);
}

static jlong (*proc_ml_SVM_getUncompressedSupportVectors_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_ml_SVM_getUncompressedSupportVectors_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_SVM_getUncompressedSupportVectors_10 (a0, a1, a2);
}

static void (*proc_features2d_GFTTDetector_setMaxFeatures_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_features2d_GFTTDetector_setMaxFeatures_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_features2d_GFTTDetector_setMaxFeatures_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_initUndistortRectifyMap_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jdouble,  jdouble,  jint,  jlong,  jlong);
void Java_org_opencv_imgproc_Imgproc_initUndistortRectifyMap_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jdouble a6, jdouble a7, jint a8, jlong a9, jlong a10) {
    proc_imgproc_Imgproc_initUndistortRectifyMap_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10);
}

static jfloat (*proc_imgproc_Imgproc_initWideAngleProjMap_10)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jint,  jint,  jlong,  jlong,  jint,  jdouble);
jfloat Java_org_opencv_imgproc_Imgproc_initWideAngleProjMap_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jint a6, jint a7, jlong a8, jlong a9, jint a10, jdouble a11) {
    return proc_imgproc_Imgproc_initWideAngleProjMap_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11);
}

static jint (*proc_core_Core_getNumberOfCPUs_10)(JNIEnv*,  jclass);
jint Java_org_opencv_core_Core_getNumberOfCPUs_10(JNIEnv* a0, jclass a1) {
    return proc_core_Core_getNumberOfCPUs_10 (a0, a1);
}

static jlong (*proc_ml_SVM_load_10)(JNIEnv*,  jclass,  jstring);
jlong Java_org_opencv_ml_SVM_load_10(JNIEnv* a0, jclass a1, jstring a2) {
    return proc_ml_SVM_load_10 (a0, a1, a2);
}

static void (*proc_features2d_GFTTDetector_setMinDistance_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_features2d_GFTTDetector_setMinDistance_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_features2d_GFTTDetector_setMinDistance_10 (a0, a1, a2, a3);
}

static jfloat (*proc_imgproc_Imgproc_initWideAngleProjMap_11)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jint,  jint,  jlong,  jlong);
jfloat Java_org_opencv_imgproc_Imgproc_initWideAngleProjMap_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jint a6, jint a7, jlong a8, jlong a9) {
    return proc_imgproc_Imgproc_initWideAngleProjMap_11 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static jint (*proc_core_Core_getOptimalDFTSize_10)(JNIEnv*,  jclass,  jint);
jint Java_org_opencv_core_Core_getOptimalDFTSize_10(JNIEnv* a0, jclass a1, jint a2) {
    return proc_core_Core_getOptimalDFTSize_10 (a0, a1, a2);
}

static void (*proc_features2d_GFTTDetector_setQualityLevel_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_features2d_GFTTDetector_setQualityLevel_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_features2d_GFTTDetector_setQualityLevel_10 (a0, a1, a2, a3);
}

static void (*proc_ml_SVM_setC_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_ml_SVM_setC_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_ml_SVM_setC_10 (a0, a1, a2, a3);
}

static jint (*proc_core_Core_getThreadNum_10)(JNIEnv*,  jclass);
jint Java_org_opencv_core_Core_getThreadNum_10(JNIEnv* a0, jclass a1) {
    return proc_core_Core_getThreadNum_10 (a0, a1);
}

static jlong (*proc_features2d_KAZE_create_10)(JNIEnv*,  jclass,  jboolean,  jboolean,  jfloat,  jint,  jint,  jint);
jlong Java_org_opencv_features2d_KAZE_create_10(JNIEnv* a0, jclass a1, jboolean a2, jboolean a3, jfloat a4, jint a5, jint a6, jint a7) {
    return proc_features2d_KAZE_create_10 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_imgproc_Imgproc_integral2_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jint,  jint);
void Java_org_opencv_imgproc_Imgproc_integral2_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jint a5, jint a6) {
    proc_imgproc_Imgproc_integral2_10 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_ml_SVM_setClassWeights_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_ml_SVM_setClassWeights_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_ml_SVM_setClassWeights_10 (a0, a1, a2, a3);
}

static jlong (*proc_core_Core_getTickCount_10)(JNIEnv*,  jclass);
jlong Java_org_opencv_core_Core_getTickCount_10(JNIEnv* a0, jclass a1) {
    return proc_core_Core_getTickCount_10 (a0, a1);
}

static jlong (*proc_features2d_KAZE_create_11)(JNIEnv*,  jclass);
jlong Java_org_opencv_features2d_KAZE_create_11(JNIEnv* a0, jclass a1) {
    return proc_features2d_KAZE_create_11 (a0, a1);
}

static void (*proc_imgproc_Imgproc_integral2_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_imgproc_Imgproc_integral2_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_imgproc_Imgproc_integral2_11 (a0, a1, a2, a3, a4);
}

static void (*proc_ml_SVM_setCoef0_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_ml_SVM_setCoef0_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_ml_SVM_setCoef0_10 (a0, a1, a2, a3);
}

static jdouble (*proc_core_Core_getTickFrequency_10)(JNIEnv*,  jclass);
jdouble Java_org_opencv_core_Core_getTickFrequency_10(JNIEnv* a0, jclass a1) {
    return proc_core_Core_getTickFrequency_10 (a0, a1);
}

static void (*proc_features2d_KAZE_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_features2d_KAZE_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_features2d_KAZE_delete (a0, a1, a2);
}

static void (*proc_imgproc_Imgproc_integral3_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jint,  jint);
void Java_org_opencv_imgproc_Imgproc_integral3_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jint a6, jint a7) {
    proc_imgproc_Imgproc_integral3_10 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_ml_SVM_setDegree_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_ml_SVM_setDegree_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_ml_SVM_setDegree_10 (a0, a1, a2, a3);
}

static jstring (*proc_features2d_KAZE_getDefaultName_10)(JNIEnv*,  jclass,  jlong);
jstring Java_org_opencv_features2d_KAZE_getDefaultName_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_KAZE_getDefaultName_10 (a0, a1, a2);
}

static void (*proc_core_Core_hconcat_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_core_Core_hconcat_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_core_Core_hconcat_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_integral3_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_imgproc_Imgproc_integral3_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_imgproc_Imgproc_integral3_11 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_ml_SVM_setGamma_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_ml_SVM_setGamma_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_ml_SVM_setGamma_10 (a0, a1, a2, a3);
}

static jint (*proc_features2d_KAZE_getDiffusivity_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_features2d_KAZE_getDiffusivity_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_KAZE_getDiffusivity_10 (a0, a1, a2);
}

static void (*proc_core_Core_idct_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint);
void Java_org_opencv_core_Core_idct_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4) {
    proc_core_Core_idct_10 (a0, a1, a2, a3, a4);
}

static void (*proc_imgproc_Imgproc_integral_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint);
void Java_org_opencv_imgproc_Imgproc_integral_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4) {
    proc_imgproc_Imgproc_integral_10 (a0, a1, a2, a3, a4);
}

static void (*proc_ml_SVM_setKernel_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_ml_SVM_setKernel_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_ml_SVM_setKernel_10 (a0, a1, a2, a3);
}

static jboolean (*proc_features2d_KAZE_getExtended_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_features2d_KAZE_getExtended_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_KAZE_getExtended_10 (a0, a1, a2);
}

static void (*proc_core_Core_idct_11)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_core_Core_idct_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_core_Core_idct_11 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_integral_11)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_imgproc_Imgproc_integral_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_imgproc_Imgproc_integral_11 (a0, a1, a2, a3);
}

static void (*proc_ml_SVM_setNu_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_ml_SVM_setNu_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_ml_SVM_setNu_10 (a0, a1, a2, a3);
}

static jboolean (*proc_videoio_VideoCapture_open_11)(JNIEnv*,  jclass,  jlong,  jstring);
jboolean Java_org_opencv_videoio_VideoCapture_open_11(JNIEnv* a0, jclass a1, jlong a2, jstring a3) {
    return proc_videoio_VideoCapture_open_11 (a0, a1, a2, a3);
}

static jdouble (*proc_calib3d_Calib3d_calibrate_10)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jlong,  jlong,  jlong,  jlong,  jint,  jint,  jint,  jdouble);
jdouble Java_org_opencv_calib3d_Calib3d_calibrate_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jlong a6, jlong a7, jlong a8, jlong a9, jint a10, jint a11, jint a12, jdouble a13) {
    return proc_calib3d_Calib3d_calibrate_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13);
}

static jdouble (*proc_ml_ANN_1MLP_getRpropDWPlus_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_ml_ANN_1MLP_getRpropDWPlus_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_ANN_1MLP_getRpropDWPlus_10 (a0, a1, a2);
}

static jdouble (*proc_video_BackgroundSubtractorKNN_getShadowThreshold_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_video_BackgroundSubtractorKNN_getShadowThreshold_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_BackgroundSubtractorKNN_getShadowThreshold_10 (a0, a1, a2);
}

static jint (*proc_core_Mat_nPutI)(JNIEnv*,  jclass,  jlong,  jint,  jint,  jint,  jintArray);
jint Java_org_opencv_core_Mat_nPutI(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4, jint a5, jintArray a6) {
    return proc_core_Mat_nPutI (a0, a1, a2, a3, a4, a5, a6);
}

static jint (*proc_features2d_AKAZE_getNOctaves_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_features2d_AKAZE_getNOctaves_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_AKAZE_getNOctaves_10 (a0, a1, a2);
}

static jstring (*proc_dnn_DictValue_getStringValue_11)(JNIEnv*,  jclass,  jlong);
jstring Java_org_opencv_dnn_DictValue_getStringValue_11(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_dnn_DictValue_getStringValue_11 (a0, a1, a2);
}

static void (*proc_imgproc_Imgproc_Canny_12)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jint,  jboolean);
void Java_org_opencv_imgproc_Imgproc_Canny_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jint a6, jboolean a7) {
    proc_imgproc_Imgproc_Canny_12 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_objdetect_HOGDescriptor_compute_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jlong);
void Java_org_opencv_objdetect_HOGDescriptor_compute_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8, jlong a9) {
    proc_objdetect_HOGDescriptor_compute_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static void (*proc_photo_AlignMTB_process_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_photo_AlignMTB_process_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_photo_AlignMTB_process_11 (a0, a1, a2, a3, a4);
}

static jfloat (*proc_imgproc_Imgproc_intersectConvexConvex_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jboolean);
jfloat Java_org_opencv_imgproc_Imgproc_intersectConvexConvex_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jboolean a5) {
    return proc_imgproc_Imgproc_intersectConvexConvex_10 (a0, a1, a2, a3, a4, a5);
}

static jint (*proc_features2d_KAZE_getNOctaveLayers_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_features2d_KAZE_getNOctaveLayers_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_KAZE_getNOctaveLayers_10 (a0, a1, a2);
}

static void (*proc_core_Core_idft_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint);
void Java_org_opencv_core_Core_idft_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5) {
    proc_core_Core_idft_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_ml_SVM_setP_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_ml_SVM_setP_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_ml_SVM_setP_10 (a0, a1, a2, a3);
}

static jfloat (*proc_imgproc_Imgproc_intersectConvexConvex_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
jfloat Java_org_opencv_imgproc_Imgproc_intersectConvexConvex_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    return proc_imgproc_Imgproc_intersectConvexConvex_11 (a0, a1, a2, a3, a4);
}

static jint (*proc_features2d_KAZE_getNOctaves_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_features2d_KAZE_getNOctaves_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_KAZE_getNOctaves_10 (a0, a1, a2);
}

static void (*proc_core_Core_idft_11)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_core_Core_idft_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_core_Core_idft_11 (a0, a1, a2, a3);
}

static void (*proc_ml_SVM_setTermCriteria_10)(JNIEnv*,  jclass,  jlong,  jint,  jint,  jdouble);
void Java_org_opencv_ml_SVM_setTermCriteria_10(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4, jdouble a5) {
    proc_ml_SVM_setTermCriteria_10 (a0, a1, a2, a3, a4, a5);
}

static jdouble (*proc_features2d_KAZE_getThreshold_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_features2d_KAZE_getThreshold_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_KAZE_getThreshold_10 (a0, a1, a2);
}

static void (*proc_core_Core_inRange_10)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jlong);
void Java_org_opencv_core_Core_inRange_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8, jdouble a9, jdouble a10, jlong a11) {
    proc_core_Core_inRange_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11);
}

static void (*proc_imgproc_Imgproc_invertAffineTransform_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_imgproc_Imgproc_invertAffineTransform_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_imgproc_Imgproc_invertAffineTransform_10 (a0, a1, a2, a3);
}

static void (*proc_ml_SVM_setType_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_ml_SVM_setType_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_ml_SVM_setType_10 (a0, a1, a2, a3);
}

static jboolean (*proc_features2d_KAZE_getUpright_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_features2d_KAZE_getUpright_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_KAZE_getUpright_10 (a0, a1, a2);
}

static jboolean (*proc_imgproc_Imgproc_isContourConvex_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_imgproc_Imgproc_isContourConvex_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_imgproc_Imgproc_isContourConvex_10 (a0, a1, a2);
}

static jboolean (*proc_ml_SVM_trainAuto_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jlong,  jint,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jboolean);
jboolean Java_org_opencv_ml_SVM_trainAuto_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jlong a5, jint a6, jlong a7, jlong a8, jlong a9, jlong a10, jlong a11, jlong a12, jboolean a13) {
    return proc_ml_SVM_trainAuto_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13);
}

static void (*proc_core_Core_insertChannel_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint);
void Java_org_opencv_core_Core_insertChannel_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4) {
    proc_core_Core_insertChannel_10 (a0, a1, a2, a3, a4);
}

static jboolean (*proc_ml_SVM_trainAuto_11)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jlong);
jboolean Java_org_opencv_ml_SVM_trainAuto_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jlong a5) {
    return proc_ml_SVM_trainAuto_11 (a0, a1, a2, a3, a4, a5);
}

static jdouble (*proc_core_Core_invert_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint);
jdouble Java_org_opencv_core_Core_invert_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4) {
    return proc_core_Core_invert_10 (a0, a1, a2, a3, a4);
}

static void (*proc_features2d_KAZE_setDiffusivity_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_features2d_KAZE_setDiffusivity_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_features2d_KAZE_setDiffusivity_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_line_10)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jint,  jint,  jint);
void Java_org_opencv_imgproc_Imgproc_line_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8, jdouble a9, jdouble a10, jint a11, jint a12, jint a13) {
    proc_imgproc_Imgproc_line_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13);
}

static jdouble (*proc_core_Core_invert_11)(JNIEnv*,  jclass,  jlong,  jlong);
jdouble Java_org_opencv_core_Core_invert_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_core_Core_invert_11 (a0, a1, a2, a3);
}

static jlong (*proc_ml_SVMSGD_create_10)(JNIEnv*,  jclass);
jlong Java_org_opencv_ml_SVMSGD_create_10(JNIEnv* a0, jclass a1) {
    return proc_ml_SVMSGD_create_10 (a0, a1);
}

static void (*proc_features2d_KAZE_setExtended_10)(JNIEnv*,  jclass,  jlong,  jboolean);
void Java_org_opencv_features2d_KAZE_setExtended_10(JNIEnv* a0, jclass a1, jlong a2, jboolean a3) {
    proc_features2d_KAZE_setExtended_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_line_11)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jint);
void Java_org_opencv_imgproc_Imgproc_line_11(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8, jdouble a9, jdouble a10, jint a11) {
    proc_imgproc_Imgproc_line_11 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11);
}

static jdouble (*proc_core_Core_kmeans_10)(JNIEnv*,  jclass,  jlong,  jint,  jlong,  jint,  jint,  jdouble,  jint,  jint,  jlong);
jdouble Java_org_opencv_core_Core_kmeans_10(JNIEnv* a0, jclass a1, jlong a2, jint a3, jlong a4, jint a5, jint a6, jdouble a7, jint a8, jint a9, jlong a10) {
    return proc_core_Core_kmeans_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10);
}

static void (*proc_features2d_KAZE_setNOctaveLayers_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_features2d_KAZE_setNOctaveLayers_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_features2d_KAZE_setNOctaveLayers_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_line_12)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_line_12(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8, jdouble a9, jdouble a10) {
    proc_imgproc_Imgproc_line_12 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10);
}

static void (*proc_ml_SVMSGD_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_ml_SVMSGD_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_ml_SVMSGD_delete (a0, a1, a2);
}

static jdouble (*proc_core_Core_kmeans_11)(JNIEnv*,  jclass,  jlong,  jint,  jlong,  jint,  jint,  jdouble,  jint,  jint);
jdouble Java_org_opencv_core_Core_kmeans_11(JNIEnv* a0, jclass a1, jlong a2, jint a3, jlong a4, jint a5, jint a6, jdouble a7, jint a8, jint a9) {
    return proc_core_Core_kmeans_11 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static jfloat (*proc_ml_SVMSGD_getInitialStepSize_10)(JNIEnv*,  jclass,  jlong);
jfloat Java_org_opencv_ml_SVMSGD_getInitialStepSize_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_SVMSGD_getInitialStepSize_10 (a0, a1, a2);
}

static void (*proc_features2d_KAZE_setNOctaves_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_features2d_KAZE_setNOctaves_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_features2d_KAZE_setNOctaves_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_linearPolar_10)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jdouble,  jint);
void Java_org_opencv_imgproc_Imgproc_linearPolar_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jdouble a6, jint a7) {
    proc_imgproc_Imgproc_linearPolar_10 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static jfloat (*proc_ml_SVMSGD_getMarginRegularization_10)(JNIEnv*,  jclass,  jlong);
jfloat Java_org_opencv_ml_SVMSGD_getMarginRegularization_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_SVMSGD_getMarginRegularization_10 (a0, a1, a2);
}

static void (*proc_core_Core_log_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_core_Core_log_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_core_Core_log_10 (a0, a1, a2, a3);
}

static void (*proc_features2d_KAZE_setThreshold_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_features2d_KAZE_setThreshold_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_features2d_KAZE_setThreshold_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_logPolar_10)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jdouble,  jint);
void Java_org_opencv_imgproc_Imgproc_logPolar_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jdouble a6, jint a7) {
    proc_imgproc_Imgproc_logPolar_10 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static jdouble (*proc_imgproc_Imgproc_matchShapes_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jdouble);
jdouble Java_org_opencv_imgproc_Imgproc_matchShapes_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jdouble a5) {
    return proc_imgproc_Imgproc_matchShapes_10 (a0, a1, a2, a3, a4, a5);
}

static jint (*proc_ml_SVMSGD_getMarginType_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_ml_SVMSGD_getMarginType_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_SVMSGD_getMarginType_10 (a0, a1, a2);
}

static void (*proc_core_Core_magnitude_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_core_Core_magnitude_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_core_Core_magnitude_10 (a0, a1, a2, a3, a4);
}

static void (*proc_features2d_KAZE_setUpright_10)(JNIEnv*,  jclass,  jlong,  jboolean);
void Java_org_opencv_features2d_KAZE_setUpright_10(JNIEnv* a0, jclass a1, jlong a2, jboolean a3) {
    proc_features2d_KAZE_setUpright_10 (a0, a1, a2, a3);
}

static jboolean (*proc_dnn_DictValue_isInt_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_dnn_DictValue_isInt_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_dnn_DictValue_isInt_10 (a0, a1, a2);
}

static jboolean (*proc_videoio_VideoCapture_open_12)(JNIEnv*,  jclass,  jlong,  jint,  jint);
jboolean Java_org_opencv_videoio_VideoCapture_open_12(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4) {
    return proc_videoio_VideoCapture_open_12 (a0, a1, a2, a3, a4);
}

static jdouble (*proc_calib3d_Calib3d_calibrate_11)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jlong,  jlong,  jlong,  jlong,  jint);
jdouble Java_org_opencv_calib3d_Calib3d_calibrate_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jlong a6, jlong a7, jlong a8, jlong a9, jint a10) {
    return proc_calib3d_Calib3d_calibrate_11 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10);
}

static jdouble (*proc_features2d_AKAZE_getThreshold_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_features2d_AKAZE_getThreshold_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_AKAZE_getThreshold_10 (a0, a1, a2);
}

static jdoubleArray (*proc_ml_ANN_1MLP_getTermCriteria_10)(JNIEnv*,  jclass,  jlong);
jdoubleArray Java_org_opencv_ml_ANN_1MLP_getTermCriteria_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_ANN_1MLP_getTermCriteria_10 (a0, a1, a2);
}

static jint (*proc_core_Mat_nPutS)(JNIEnv*,  jclass,  jlong,  jint,  jint,  jint,  jshortArray);
jint Java_org_opencv_core_Mat_nPutS(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4, jint a5, jshortArray a6) {
    return proc_core_Mat_nPutS (a0, a1, a2, a3, a4, a5, a6);
}

static jint (*proc_video_BackgroundSubtractorKNN_getShadowValue_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_video_BackgroundSubtractorKNN_getShadowValue_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_BackgroundSubtractorKNN_getShadowValue_10 (a0, a1, a2);
}

static void (*proc_imgproc_Imgproc_Canny_13)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_Canny_13(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5) {
    proc_imgproc_Imgproc_Canny_13 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_objdetect_HOGDescriptor_compute_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_objdetect_HOGDescriptor_compute_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_objdetect_HOGDescriptor_compute_11 (a0, a1, a2, a3, a4);
}

static void (*proc_photo_AlignMTB_setCut_10)(JNIEnv*,  jclass,  jlong,  jboolean);
void Java_org_opencv_photo_AlignMTB_setCut_10(JNIEnv* a0, jclass a1, jlong a2, jboolean a3) {
    proc_photo_AlignMTB_setCut_10 (a0, a1, a2, a3);
}

static jfloat (*proc_ml_SVMSGD_getShift_10)(JNIEnv*,  jclass,  jlong);
jfloat Java_org_opencv_ml_SVMSGD_getShift_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_SVMSGD_getShift_10 (a0, a1, a2);
}

static jlong (*proc_features2d_MSER_create_10)(JNIEnv*,  jclass,  jint,  jint,  jint,  jdouble,  jdouble,  jint,  jdouble,  jdouble,  jint);
jlong Java_org_opencv_features2d_MSER_create_10(JNIEnv* a0, jclass a1, jint a2, jint a3, jint a4, jdouble a5, jdouble a6, jint a7, jdouble a8, jdouble a9, jint a10) {
    return proc_features2d_MSER_create_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10);
}

static void (*proc_core_Core_max_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_core_Core_max_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_core_Core_max_10 (a0, a1, a2, a3, a4);
}

static void (*proc_imgproc_Imgproc_matchTemplate_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jint,  jlong);
void Java_org_opencv_imgproc_Imgproc_matchTemplate_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jint a5, jlong a6) {
    proc_imgproc_Imgproc_matchTemplate_10 (a0, a1, a2, a3, a4, a5, a6);
}

static jfloat (*proc_ml_SVMSGD_getStepDecreasingPower_10)(JNIEnv*,  jclass,  jlong);
jfloat Java_org_opencv_ml_SVMSGD_getStepDecreasingPower_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_SVMSGD_getStepDecreasingPower_10 (a0, a1, a2);
}

static jlong (*proc_features2d_MSER_create_11)(JNIEnv*,  jclass);
jlong Java_org_opencv_features2d_MSER_create_11(JNIEnv* a0, jclass a1) {
    return proc_features2d_MSER_create_11 (a0, a1);
}

static void (*proc_core_Core_max_11)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jlong);
void Java_org_opencv_core_Core_max_11(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jlong a7) {
    proc_core_Core_max_11 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_imgproc_Imgproc_matchTemplate_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jint);
void Java_org_opencv_imgproc_Imgproc_matchTemplate_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jint a5) {
    proc_imgproc_Imgproc_matchTemplate_11 (a0, a1, a2, a3, a4, a5);
}

static jint (*proc_ml_SVMSGD_getSvmsgdType_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_ml_SVMSGD_getSvmsgdType_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_SVMSGD_getSvmsgdType_10 (a0, a1, a2);
}

static void (*proc_core_Core_meanStdDev_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_core_Core_meanStdDev_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_core_Core_meanStdDev_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_features2d_MSER_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_features2d_MSER_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_features2d_MSER_delete (a0, a1, a2);
}

static void (*proc_imgproc_Imgproc_medianBlur_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint);
void Java_org_opencv_imgproc_Imgproc_medianBlur_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4) {
    proc_imgproc_Imgproc_medianBlur_10 (a0, a1, a2, a3, a4);
}

static jdoubleArray (*proc_imgproc_Imgproc_minAreaRect_10)(JNIEnv*,  jclass,  jlong);
jdoubleArray Java_org_opencv_imgproc_Imgproc_minAreaRect_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_imgproc_Imgproc_minAreaRect_10 (a0, a1, a2);
}

static jdoubleArray (*proc_ml_SVMSGD_getTermCriteria_10)(JNIEnv*,  jclass,  jlong);
jdoubleArray Java_org_opencv_ml_SVMSGD_getTermCriteria_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_SVMSGD_getTermCriteria_10 (a0, a1, a2);
}

static void (*proc_core_Core_meanStdDev_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_core_Core_meanStdDev_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_core_Core_meanStdDev_11 (a0, a1, a2, a3, a4);
}

static void (*proc_features2d_MSER_detectRegions_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_features2d_MSER_detectRegions_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_features2d_MSER_detectRegions_10 (a0, a1, a2, a3, a4, a5);
}

static jdoubleArray (*proc_core_Core_mean_10)(JNIEnv*,  jclass,  jlong,  jlong);
jdoubleArray Java_org_opencv_core_Core_mean_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_core_Core_mean_10 (a0, a1, a2, a3);
}

static jlong (*proc_ml_SVMSGD_getWeights_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_ml_SVMSGD_getWeights_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_SVMSGD_getWeights_10 (a0, a1, a2);
}

static jstring (*proc_features2d_MSER_getDefaultName_10)(JNIEnv*,  jclass,  jlong);
jstring Java_org_opencv_features2d_MSER_getDefaultName_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_MSER_getDefaultName_10 (a0, a1, a2);
}

static void (*proc_imgproc_Imgproc_minEnclosingCircle_10)(JNIEnv*,  jclass,  jlong,  jdoubleArray,  jdoubleArray);
void Java_org_opencv_imgproc_Imgproc_minEnclosingCircle_10(JNIEnv* a0, jclass a1, jlong a2, jdoubleArray a3, jdoubleArray a4) {
    proc_imgproc_Imgproc_minEnclosingCircle_10 (a0, a1, a2, a3, a4);
}

static jdouble (*proc_imgproc_Imgproc_minEnclosingTriangle_10)(JNIEnv*,  jclass,  jlong,  jlong);
jdouble Java_org_opencv_imgproc_Imgproc_minEnclosingTriangle_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_imgproc_Imgproc_minEnclosingTriangle_10 (a0, a1, a2, a3);
}

static jdoubleArray (*proc_core_Core_mean_11)(JNIEnv*,  jclass,  jlong);
jdoubleArray Java_org_opencv_core_Core_mean_11(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_core_Core_mean_11 (a0, a1, a2);
}

static jint (*proc_features2d_MSER_getDelta_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_features2d_MSER_getDelta_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_MSER_getDelta_10 (a0, a1, a2);
}

static jlong (*proc_ml_SVMSGD_load_10)(JNIEnv*,  jclass,  jstring,  jstring);
jlong Java_org_opencv_ml_SVMSGD_load_10(JNIEnv* a0, jclass a1, jstring a2, jstring a3) {
    return proc_ml_SVMSGD_load_10 (a0, a1, a2, a3);
}

static jdoubleArray (*proc_imgproc_Imgproc_moments_10)(JNIEnv*,  jclass,  jlong,  jboolean);
jdoubleArray Java_org_opencv_imgproc_Imgproc_moments_10(JNIEnv* a0, jclass a1, jlong a2, jboolean a3) {
    return proc_imgproc_Imgproc_moments_10 (a0, a1, a2, a3);
}

static jint (*proc_features2d_MSER_getMaxArea_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_features2d_MSER_getMaxArea_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_MSER_getMaxArea_10 (a0, a1, a2);
}

static jlong (*proc_ml_SVMSGD_load_11)(JNIEnv*,  jclass,  jstring);
jlong Java_org_opencv_ml_SVMSGD_load_11(JNIEnv* a0, jclass a1, jstring a2) {
    return proc_ml_SVMSGD_load_11 (a0, a1, a2);
}

static void (*proc_core_Core_merge_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_core_Core_merge_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_core_Core_merge_10 (a0, a1, a2, a3);
}

static jdoubleArray (*proc_imgproc_Imgproc_moments_11)(JNIEnv*,  jclass,  jlong);
jdoubleArray Java_org_opencv_imgproc_Imgproc_moments_11(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_imgproc_Imgproc_moments_11 (a0, a1, a2);
}

static jint (*proc_features2d_MSER_getMinArea_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_features2d_MSER_getMinArea_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_MSER_getMinArea_10 (a0, a1, a2);
}

static void (*proc_core_Core_min_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_core_Core_min_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_core_Core_min_10 (a0, a1, a2, a3, a4);
}

static void (*proc_ml_SVMSGD_setInitialStepSize_10)(JNIEnv*,  jclass,  jlong,  jfloat);
void Java_org_opencv_ml_SVMSGD_setInitialStepSize_10(JNIEnv* a0, jclass a1, jlong a2, jfloat a3) {
    proc_ml_SVMSGD_setInitialStepSize_10 (a0, a1, a2, a3);
}

static jboolean (*proc_features2d_MSER_getPass2Only_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_features2d_MSER_getPass2Only_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_MSER_getPass2Only_10 (a0, a1, a2);
}

static void (*proc_core_Core_min_11)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jlong);
void Java_org_opencv_core_Core_min_11(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jlong a7) {
    proc_core_Core_min_11 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_imgproc_Imgproc_morphologyEx_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jlong,  jdouble,  jdouble,  jint,  jint,  jdouble,  jdouble,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_morphologyEx_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jlong a5, jdouble a6, jdouble a7, jint a8, jint a9, jdouble a10, jdouble a11, jdouble a12, jdouble a13) {
    proc_imgproc_Imgproc_morphologyEx_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13);
}

static void (*proc_ml_SVMSGD_setMarginRegularization_10)(JNIEnv*,  jclass,  jlong,  jfloat);
void Java_org_opencv_ml_SVMSGD_setMarginRegularization_10(JNIEnv* a0, jclass a1, jlong a2, jfloat a3) {
    proc_ml_SVMSGD_setMarginRegularization_10 (a0, a1, a2, a3);
}

static void (*proc_core_Core_mixChannels_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_core_Core_mixChannels_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_core_Core_mixChannels_10 (a0, a1, a2, a3, a4);
}

static void (*proc_features2d_MSER_setDelta_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_features2d_MSER_setDelta_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_features2d_MSER_setDelta_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_morphologyEx_11)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jlong,  jdouble,  jdouble,  jint);
void Java_org_opencv_imgproc_Imgproc_morphologyEx_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jlong a5, jdouble a6, jdouble a7, jint a8) {
    proc_imgproc_Imgproc_morphologyEx_11 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static void (*proc_ml_SVMSGD_setMarginType_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_ml_SVMSGD_setMarginType_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_ml_SVMSGD_setMarginType_10 (a0, a1, a2, a3);
}

static jboolean (*proc_dnn_DictValue_isReal_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_dnn_DictValue_isReal_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_dnn_DictValue_isReal_10 (a0, a1, a2);
}

static jboolean (*proc_videoio_VideoCapture_open_13)(JNIEnv*,  jclass,  jlong,  jint);
jboolean Java_org_opencv_videoio_VideoCapture_open_13(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    return proc_videoio_VideoCapture_open_13 (a0, a1, a2, a3);
}

static jdouble (*proc_calib3d_Calib3d_calibrate_12)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jlong,  jlong,  jlong,  jlong);
jdouble Java_org_opencv_calib3d_Calib3d_calibrate_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jlong a6, jlong a7, jlong a8, jlong a9) {
    return proc_calib3d_Calib3d_calibrate_12 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static jfloat (*proc_imgproc_Imgproc_EMD_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jlong,  jlong);
jfloat Java_org_opencv_imgproc_Imgproc_EMD_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jlong a5, jlong a6) {
    return proc_imgproc_Imgproc_EMD_10 (a0, a1, a2, a3, a4, a5, a6);
}

static jint (*proc_ml_ANN_1MLP_getTrainMethod_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_ml_ANN_1MLP_getTrainMethod_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_ANN_1MLP_getTrainMethod_10 (a0, a1, a2);
}

static jint (*proc_video_BackgroundSubtractorKNN_getkNNSamples_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_video_BackgroundSubtractorKNN_getkNNSamples_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_BackgroundSubtractorKNN_getkNNSamples_10 (a0, a1, a2);
}

static jlong (*proc_core_Mat_n_1Mat__)(JNIEnv*,  jclass);
jlong Java_org_opencv_core_Mat_n_1Mat__(JNIEnv* a0, jclass a1) {
    return proc_core_Mat_n_1Mat__ (a0, a1);
}

static void (*proc_features2d_AKAZE_setDescriptorChannels_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_features2d_AKAZE_setDescriptorChannels_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_features2d_AKAZE_setDescriptorChannels_10 (a0, a1, a2, a3);
}

static void (*proc_objdetect_HOGDescriptor_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_objdetect_HOGDescriptor_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_objdetect_HOGDescriptor_delete (a0, a1, a2);
}

static void (*proc_photo_AlignMTB_setExcludeRange_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_photo_AlignMTB_setExcludeRange_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_photo_AlignMTB_setExcludeRange_10 (a0, a1, a2, a3);
}

static void (*proc_core_Core_mulSpectrums_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jint,  jboolean);
void Java_org_opencv_core_Core_mulSpectrums_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jint a5, jboolean a6) {
    proc_core_Core_mulSpectrums_10 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_features2d_MSER_setMaxArea_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_features2d_MSER_setMaxArea_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_features2d_MSER_setMaxArea_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_morphologyEx_12)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jlong);
void Java_org_opencv_imgproc_Imgproc_morphologyEx_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jlong a5) {
    proc_imgproc_Imgproc_morphologyEx_12 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_ml_SVMSGD_setOptimalParameters_10)(JNIEnv*,  jclass,  jlong,  jint,  jint);
void Java_org_opencv_ml_SVMSGD_setOptimalParameters_10(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4) {
    proc_ml_SVMSGD_setOptimalParameters_10 (a0, a1, a2, a3, a4);
}

static jdoubleArray (*proc_imgproc_Imgproc_n_1getTextSize)(JNIEnv*,  jclass,  jstring,  jint,  jdouble,  jint,  jintArray);
jdoubleArray Java_org_opencv_imgproc_Imgproc_n_1getTextSize(JNIEnv* a0, jclass a1, jstring a2, jint a3, jdouble a4, jint a5, jintArray a6) {
    return proc_imgproc_Imgproc_n_1getTextSize (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_core_Core_mulSpectrums_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jint);
void Java_org_opencv_core_Core_mulSpectrums_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jint a5) {
    proc_core_Core_mulSpectrums_11 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_features2d_MSER_setMinArea_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_features2d_MSER_setMinArea_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_features2d_MSER_setMinArea_10 (a0, a1, a2, a3);
}

static void (*proc_ml_SVMSGD_setOptimalParameters_11)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_ml_SVMSGD_setOptimalParameters_11(JNIEnv* a0, jclass a1, jlong a2) {
    proc_ml_SVMSGD_setOptimalParameters_11 (a0, a1, a2);
}

static jdoubleArray (*proc_imgproc_Imgproc_phaseCorrelate_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jdoubleArray);
jdoubleArray Java_org_opencv_imgproc_Imgproc_phaseCorrelate_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jdoubleArray a5) {
    return proc_imgproc_Imgproc_phaseCorrelate_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_core_Core_mulTransposed_10)(JNIEnv*,  jclass,  jlong,  jlong,  jboolean,  jlong,  jdouble,  jint);
void Java_org_opencv_core_Core_mulTransposed_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jboolean a4, jlong a5, jdouble a6, jint a7) {
    proc_core_Core_mulTransposed_10 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_features2d_MSER_setPass2Only_10)(JNIEnv*,  jclass,  jlong,  jboolean);
void Java_org_opencv_features2d_MSER_setPass2Only_10(JNIEnv* a0, jclass a1, jlong a2, jboolean a3) {
    proc_features2d_MSER_setPass2Only_10 (a0, a1, a2, a3);
}

static void (*proc_ml_SVMSGD_setStepDecreasingPower_10)(JNIEnv*,  jclass,  jlong,  jfloat);
void Java_org_opencv_ml_SVMSGD_setStepDecreasingPower_10(JNIEnv* a0, jclass a1, jlong a2, jfloat a3) {
    proc_ml_SVMSGD_setStepDecreasingPower_10 (a0, a1, a2, a3);
}

static jdoubleArray (*proc_imgproc_Imgproc_phaseCorrelate_11)(JNIEnv*,  jclass,  jlong,  jlong);
jdoubleArray Java_org_opencv_imgproc_Imgproc_phaseCorrelate_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_imgproc_Imgproc_phaseCorrelate_11 (a0, a1, a2, a3);
}

static jlong (*proc_features2d_ORB_create_10)(JNIEnv*,  jclass,  jint,  jfloat,  jint,  jint,  jint,  jint,  jint,  jint,  jint);
jlong Java_org_opencv_features2d_ORB_create_10(JNIEnv* a0, jclass a1, jint a2, jfloat a3, jint a4, jint a5, jint a6, jint a7, jint a8, jint a9, jint a10) {
    return proc_features2d_ORB_create_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10);
}

static void (*proc_core_Core_mulTransposed_11)(JNIEnv*,  jclass,  jlong,  jlong,  jboolean,  jlong,  jdouble);
void Java_org_opencv_core_Core_mulTransposed_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jboolean a4, jlong a5, jdouble a6) {
    proc_core_Core_mulTransposed_11 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_ml_SVMSGD_setSvmsgdType_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_ml_SVMSGD_setSvmsgdType_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_ml_SVMSGD_setSvmsgdType_10 (a0, a1, a2, a3);
}

static jdouble (*proc_imgproc_Imgproc_pointPolygonTest_10)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jboolean);
jdouble Java_org_opencv_imgproc_Imgproc_pointPolygonTest_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jboolean a5) {
    return proc_imgproc_Imgproc_pointPolygonTest_10 (a0, a1, a2, a3, a4, a5);
}

static jlong (*proc_features2d_ORB_create_11)(JNIEnv*,  jclass);
jlong Java_org_opencv_features2d_ORB_create_11(JNIEnv* a0, jclass a1) {
    return proc_features2d_ORB_create_11 (a0, a1);
}

static void (*proc_core_Core_mulTransposed_12)(JNIEnv*,  jclass,  jlong,  jlong,  jboolean);
void Java_org_opencv_core_Core_mulTransposed_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jboolean a4) {
    proc_core_Core_mulTransposed_12 (a0, a1, a2, a3, a4);
}

static void (*proc_ml_SVMSGD_setTermCriteria_10)(JNIEnv*,  jclass,  jlong,  jint,  jint,  jdouble);
void Java_org_opencv_ml_SVMSGD_setTermCriteria_10(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4, jdouble a5) {
    proc_ml_SVMSGD_setTermCriteria_10 (a0, a1, a2, a3, a4, a5);
}

static jfloat (*proc_ml_StatModel_calcError_10)(JNIEnv*,  jclass,  jlong,  jlong,  jboolean,  jlong);
jfloat Java_org_opencv_ml_StatModel_calcError_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jboolean a4, jlong a5) {
    return proc_ml_StatModel_calcError_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_core_Core_multiply_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jdouble,  jint);
void Java_org_opencv_core_Core_multiply_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jdouble a5, jint a6) {
    proc_core_Core_multiply_10 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_features2d_ORB_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_features2d_ORB_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_features2d_ORB_delete (a0, a1, a2);
}

static void (*proc_imgproc_Imgproc_polylines_10)(JNIEnv*,  jclass,  jlong,  jlong,  jboolean,  jdouble,  jdouble,  jdouble,  jdouble,  jint,  jint,  jint);
void Java_org_opencv_imgproc_Imgproc_polylines_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jboolean a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8, jint a9, jint a10, jint a11) {
    proc_imgproc_Imgproc_polylines_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11);
}

static jstring (*proc_features2d_ORB_getDefaultName_10)(JNIEnv*,  jclass,  jlong);
jstring Java_org_opencv_features2d_ORB_getDefaultName_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_ORB_getDefaultName_10 (a0, a1, a2);
}

static void (*proc_core_Core_multiply_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jdouble);
void Java_org_opencv_core_Core_multiply_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jdouble a5) {
    proc_core_Core_multiply_11 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_imgproc_Imgproc_polylines_11)(JNIEnv*,  jclass,  jlong,  jlong,  jboolean,  jdouble,  jdouble,  jdouble,  jdouble,  jint);
void Java_org_opencv_imgproc_Imgproc_polylines_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jboolean a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8, jint a9) {
    proc_imgproc_Imgproc_polylines_11 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static void (*proc_ml_StatModel_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_ml_StatModel_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_ml_StatModel_delete (a0, a1, a2);
}

static jboolean (*proc_ml_StatModel_empty_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_ml_StatModel_empty_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_StatModel_empty_10 (a0, a1, a2);
}

static jint (*proc_features2d_ORB_getEdgeThreshold_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_features2d_ORB_getEdgeThreshold_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_ORB_getEdgeThreshold_10 (a0, a1, a2);
}

static void (*proc_core_Core_multiply_12)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_core_Core_multiply_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_core_Core_multiply_12 (a0, a1, a2, a3, a4);
}

static void (*proc_imgproc_Imgproc_polylines_12)(JNIEnv*,  jclass,  jlong,  jlong,  jboolean,  jdouble,  jdouble,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_polylines_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jboolean a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8) {
    proc_imgproc_Imgproc_polylines_12 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static jint (*proc_features2d_ORB_getFastThreshold_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_features2d_ORB_getFastThreshold_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_ORB_getFastThreshold_10 (a0, a1, a2);
}

static jint (*proc_ml_StatModel_getVarCount_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_ml_StatModel_getVarCount_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_StatModel_getVarCount_10 (a0, a1, a2);
}

static void (*proc_core_Core_multiply_13)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jlong,  jdouble,  jint);
void Java_org_opencv_core_Core_multiply_13(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jlong a7, jdouble a8, jint a9) {
    proc_core_Core_multiply_13 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static void (*proc_imgproc_Imgproc_preCornerDetect_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint);
void Java_org_opencv_imgproc_Imgproc_preCornerDetect_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5) {
    proc_imgproc_Imgproc_preCornerDetect_10 (a0, a1, a2, a3, a4, a5);
}

static jboolean (*proc_ml_StatModel_isClassifier_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_ml_StatModel_isClassifier_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_StatModel_isClassifier_10 (a0, a1, a2);
}

static jint (*proc_features2d_ORB_getFirstLevel_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_features2d_ORB_getFirstLevel_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_ORB_getFirstLevel_10 (a0, a1, a2);
}

static void (*proc_core_Core_multiply_14)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jlong,  jdouble);
void Java_org_opencv_core_Core_multiply_14(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jlong a7, jdouble a8) {
    proc_core_Core_multiply_14 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static void (*proc_imgproc_Imgproc_preCornerDetect_11)(JNIEnv*,  jclass,  jlong,  jlong,  jint);
void Java_org_opencv_imgproc_Imgproc_preCornerDetect_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4) {
    proc_imgproc_Imgproc_preCornerDetect_11 (a0, a1, a2, a3, a4);
}

static jboolean (*proc_dnn_DictValue_isString_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_dnn_DictValue_isString_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_dnn_DictValue_isString_10 (a0, a1, a2);
}

static jboolean (*proc_videoio_VideoCapture_read_10)(JNIEnv*,  jclass,  jlong,  jlong);
jboolean Java_org_opencv_videoio_VideoCapture_read_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_videoio_VideoCapture_read_10 (a0, a1, a2, a3);
}

static jfloat (*proc_imgproc_Imgproc_EMD_11)(JNIEnv*,  jclass,  jlong,  jlong,  jint);
jfloat Java_org_opencv_imgproc_Imgproc_EMD_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4) {
    return proc_imgproc_Imgproc_EMD_11 (a0, a1, a2, a3, a4);
}

static jlong (*proc_core_Mat_n_1Mat__III)(JNIEnv*,  jclass,  jint,  jint,  jint);
jlong Java_org_opencv_core_Mat_n_1Mat__III(JNIEnv* a0, jclass a1, jint a2, jint a3, jint a4) {
    return proc_core_Mat_n_1Mat__III (a0, a1, a2, a3, a4);
}

static jlong (*proc_ml_ANN_1MLP_getWeights_10)(JNIEnv*,  jclass,  jlong,  jint);
jlong Java_org_opencv_ml_ANN_1MLP_getWeights_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    return proc_ml_ANN_1MLP_getWeights_10 (a0, a1, a2, a3);
}

static void (*proc_calib3d_Calib3d_calibrationMatrixValues_10)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jdoubleArray,  jdoubleArray,  jdoubleArray,  jdoubleArray,  jdoubleArray);
void Java_org_opencv_calib3d_Calib3d_calibrationMatrixValues_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jdoubleArray a7, jdoubleArray a8, jdoubleArray a9, jdoubleArray a10, jdoubleArray a11) {
    proc_calib3d_Calib3d_calibrationMatrixValues_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11);
}

static void (*proc_features2d_AKAZE_setDescriptorSize_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_features2d_AKAZE_setDescriptorSize_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_features2d_AKAZE_setDescriptorSize_10 (a0, a1, a2, a3);
}

static void (*proc_objdetect_HOGDescriptor_detectMultiScale_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jboolean);
void Java_org_opencv_objdetect_HOGDescriptor_detectMultiScale_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jdouble a6, jdouble a7, jdouble a8, jdouble a9, jdouble a10, jdouble a11, jdouble a12, jboolean a13) {
    proc_objdetect_HOGDescriptor_detectMultiScale_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13);
}

static void (*proc_photo_AlignMTB_setMaxBits_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_photo_AlignMTB_setMaxBits_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_photo_AlignMTB_setMaxBits_10 (a0, a1, a2, a3);
}

static void (*proc_video_BackgroundSubtractorKNN_setDetectShadows_10)(JNIEnv*,  jclass,  jlong,  jboolean);
void Java_org_opencv_video_BackgroundSubtractorKNN_setDetectShadows_10(JNIEnv* a0, jclass a1, jlong a2, jboolean a3) {
    proc_video_BackgroundSubtractorKNN_setDetectShadows_10 (a0, a1, a2, a3);
}

static jboolean (*proc_ml_StatModel_isTrained_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_ml_StatModel_isTrained_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_StatModel_isTrained_10 (a0, a1, a2);
}

static jint (*proc_features2d_ORB_getMaxFeatures_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_features2d_ORB_getMaxFeatures_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_ORB_getMaxFeatures_10 (a0, a1, a2);
}

static void (*proc_core_Core_multiply_15)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jlong);
void Java_org_opencv_core_Core_multiply_15(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jlong a7) {
    proc_core_Core_multiply_15 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_imgproc_Imgproc_putText_10)(JNIEnv*,  jclass,  jlong,  jstring,  jdouble,  jdouble,  jint,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jint,  jint,  jboolean);
void Java_org_opencv_imgproc_Imgproc_putText_10(JNIEnv* a0, jclass a1, jlong a2, jstring a3, jdouble a4, jdouble a5, jint a6, jdouble a7, jdouble a8, jdouble a9, jdouble a10, jdouble a11, jint a12, jint a13, jboolean a14) {
    proc_imgproc_Imgproc_putText_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14);
}

static jdoubleArray (*proc_core_Core_n_1minMaxLocManual)(JNIEnv*,  jclass,  jlong,  jlong);
jdoubleArray Java_org_opencv_core_Core_n_1minMaxLocManual(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_core_Core_n_1minMaxLocManual (a0, a1, a2, a3);
}

static jfloat (*proc_ml_StatModel_predict_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jint);
jfloat Java_org_opencv_ml_StatModel_predict_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jint a5) {
    return proc_ml_StatModel_predict_10 (a0, a1, a2, a3, a4, a5);
}

static jint (*proc_features2d_ORB_getNLevels_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_features2d_ORB_getNLevels_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_ORB_getNLevels_10 (a0, a1, a2);
}

static void (*proc_imgproc_Imgproc_putText_11)(JNIEnv*,  jclass,  jlong,  jstring,  jdouble,  jdouble,  jint,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jint);
void Java_org_opencv_imgproc_Imgproc_putText_11(JNIEnv* a0, jclass a1, jlong a2, jstring a3, jdouble a4, jdouble a5, jint a6, jdouble a7, jdouble a8, jdouble a9, jdouble a10, jdouble a11, jint a12) {
    proc_imgproc_Imgproc_putText_11 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12);
}

static jdouble (*proc_core_Core_norm_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jlong);
jdouble Java_org_opencv_core_Core_norm_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jlong a5) {
    return proc_core_Core_norm_10 (a0, a1, a2, a3, a4, a5);
}

static jfloat (*proc_ml_StatModel_predict_11)(JNIEnv*,  jclass,  jlong,  jlong);
jfloat Java_org_opencv_ml_StatModel_predict_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_ml_StatModel_predict_11 (a0, a1, a2, a3);
}

static jint (*proc_features2d_ORB_getPatchSize_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_features2d_ORB_getPatchSize_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_ORB_getPatchSize_10 (a0, a1, a2);
}

static void (*proc_imgproc_Imgproc_putText_12)(JNIEnv*,  jclass,  jlong,  jstring,  jdouble,  jdouble,  jint,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_putText_12(JNIEnv* a0, jclass a1, jlong a2, jstring a3, jdouble a4, jdouble a5, jint a6, jdouble a7, jdouble a8, jdouble a9, jdouble a10, jdouble a11) {
    proc_imgproc_Imgproc_putText_12 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11);
}

static jboolean (*proc_ml_StatModel_train_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jlong);
jboolean Java_org_opencv_ml_StatModel_train_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jlong a5) {
    return proc_ml_StatModel_train_10 (a0, a1, a2, a3, a4, a5);
}

static jdouble (*proc_core_Core_norm_11)(JNIEnv*,  jclass,  jlong,  jlong,  jint);
jdouble Java_org_opencv_core_Core_norm_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4) {
    return proc_core_Core_norm_11 (a0, a1, a2, a3, a4);
}

static jdouble (*proc_features2d_ORB_getScaleFactor_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_features2d_ORB_getScaleFactor_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_ORB_getScaleFactor_10 (a0, a1, a2);
}

static void (*proc_imgproc_Imgproc_pyrDown_10)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jint);
void Java_org_opencv_imgproc_Imgproc_pyrDown_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jint a6) {
    proc_imgproc_Imgproc_pyrDown_10 (a0, a1, a2, a3, a4, a5, a6);
}

static jboolean (*proc_ml_StatModel_train_11)(JNIEnv*,  jclass,  jlong,  jlong,  jint);
jboolean Java_org_opencv_ml_StatModel_train_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4) {
    return proc_ml_StatModel_train_11 (a0, a1, a2, a3, a4);
}

static jdouble (*proc_core_Core_norm_12)(JNIEnv*,  jclass,  jlong,  jlong);
jdouble Java_org_opencv_core_Core_norm_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_core_Core_norm_12 (a0, a1, a2, a3);
}

static jint (*proc_features2d_ORB_getScoreType_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_features2d_ORB_getScoreType_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_ORB_getScoreType_10 (a0, a1, a2);
}

static void (*proc_imgproc_Imgproc_pyrDown_11)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_pyrDown_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5) {
    proc_imgproc_Imgproc_pyrDown_11 (a0, a1, a2, a3, a4, a5);
}

static jboolean (*proc_ml_StatModel_train_12)(JNIEnv*,  jclass,  jlong,  jlong);
jboolean Java_org_opencv_ml_StatModel_train_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_ml_StatModel_train_12 (a0, a1, a2, a3);
}

static jdouble (*proc_core_Core_norm_13)(JNIEnv*,  jclass,  jlong,  jint,  jlong);
jdouble Java_org_opencv_core_Core_norm_13(JNIEnv* a0, jclass a1, jlong a2, jint a3, jlong a4) {
    return proc_core_Core_norm_13 (a0, a1, a2, a3, a4);
}

static jint (*proc_features2d_ORB_getWTA_1K_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_features2d_ORB_getWTA_1K_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_ORB_getWTA_1K_10 (a0, a1, a2);
}

static void (*proc_imgproc_Imgproc_pyrDown_12)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_imgproc_Imgproc_pyrDown_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_imgproc_Imgproc_pyrDown_12 (a0, a1, a2, a3);
}

static jdouble (*proc_core_Core_norm_14)(JNIEnv*,  jclass,  jlong,  jint);
jdouble Java_org_opencv_core_Core_norm_14(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    return proc_core_Core_norm_14 (a0, a1, a2, a3);
}

static jlong (*proc_ml_TrainData_create_10)(JNIEnv*,  jclass,  jlong,  jint,  jlong,  jlong,  jlong,  jlong,  jlong);
jlong Java_org_opencv_ml_TrainData_create_10(JNIEnv* a0, jclass a1, jlong a2, jint a3, jlong a4, jlong a5, jlong a6, jlong a7, jlong a8) {
    return proc_ml_TrainData_create_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static void (*proc_features2d_ORB_setEdgeThreshold_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_features2d_ORB_setEdgeThreshold_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_features2d_ORB_setEdgeThreshold_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_pyrMeanShiftFiltering_10)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jint,  jint,  jint,  jdouble);
void Java_org_opencv_imgproc_Imgproc_pyrMeanShiftFiltering_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jint a6, jint a7, jint a8, jdouble a9) {
    proc_imgproc_Imgproc_pyrMeanShiftFiltering_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static jdouble (*proc_core_Core_norm_15)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_core_Core_norm_15(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_core_Core_norm_15 (a0, a1, a2);
}

static jlong (*proc_ml_TrainData_create_11)(JNIEnv*,  jclass,  jlong,  jint,  jlong);
jlong Java_org_opencv_ml_TrainData_create_11(JNIEnv* a0, jclass a1, jlong a2, jint a3, jlong a4) {
    return proc_ml_TrainData_create_11 (a0, a1, a2, a3, a4);
}

static void (*proc_features2d_ORB_setFastThreshold_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_features2d_ORB_setFastThreshold_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_features2d_ORB_setFastThreshold_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_pyrMeanShiftFiltering_11)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_pyrMeanShiftFiltering_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5) {
    proc_imgproc_Imgproc_pyrMeanShiftFiltering_11 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_core_Core_normalize_10)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jint,  jint,  jlong);
void Java_org_opencv_core_Core_normalize_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jint a6, jint a7, jlong a8) {
    proc_core_Core_normalize_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static void (*proc_features2d_ORB_setFirstLevel_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_features2d_ORB_setFirstLevel_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_features2d_ORB_setFirstLevel_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_pyrUp_10)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jint);
void Java_org_opencv_imgproc_Imgproc_pyrUp_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jint a6) {
    proc_imgproc_Imgproc_pyrUp_10 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_ml_TrainData_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_ml_TrainData_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_ml_TrainData_delete (a0, a1, a2);
}

static jint (*proc_ml_TrainData_getCatCount_10)(JNIEnv*,  jclass,  jlong,  jint);
jint Java_org_opencv_ml_TrainData_getCatCount_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    return proc_ml_TrainData_getCatCount_10 (a0, a1, a2, a3);
}

static void (*proc_core_Core_normalize_11)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jint,  jint);
void Java_org_opencv_core_Core_normalize_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jint a6, jint a7) {
    proc_core_Core_normalize_11 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_features2d_ORB_setMaxFeatures_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_features2d_ORB_setMaxFeatures_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_features2d_ORB_setMaxFeatures_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_pyrUp_11)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_pyrUp_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5) {
    proc_imgproc_Imgproc_pyrUp_11 (a0, a1, a2, a3, a4, a5);
}

static jlong (*proc_core_Mat_n_1Mat__DDI)(JNIEnv*,  jclass,  jdouble,  jdouble,  jint);
jlong Java_org_opencv_core_Mat_n_1Mat__DDI(JNIEnv* a0, jclass a1, jdouble a2, jdouble a3, jint a4) {
    return proc_core_Mat_n_1Mat__DDI (a0, a1, a2, a3, a4);
}

static jlong (*proc_dnn_Dnn_blobFromImage_10)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jboolean,  jboolean);
jlong Java_org_opencv_dnn_Dnn_blobFromImage_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8, jdouble a9, jboolean a10, jboolean a11) {
    return proc_dnn_Dnn_blobFromImage_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11);
}

static jlong (*proc_ml_ANN_1MLP_load_10)(JNIEnv*,  jclass,  jstring);
jlong Java_org_opencv_ml_ANN_1MLP_load_10(JNIEnv* a0, jclass a1, jstring a2) {
    return proc_ml_ANN_1MLP_load_10 (a0, a1, a2);
}

static void (*proc_calib3d_Calib3d_composeRT_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_calib3d_Calib3d_composeRT_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7, jlong a8, jlong a9, jlong a10, jlong a11, jlong a12, jlong a13, jlong a14, jlong a15) {
    proc_calib3d_Calib3d_composeRT_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15);
}

static void (*proc_features2d_AKAZE_setDescriptorType_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_features2d_AKAZE_setDescriptorType_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_features2d_AKAZE_setDescriptorType_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_GaussianBlur_10)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jint);
void Java_org_opencv_imgproc_Imgproc_GaussianBlur_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7, jint a8) {
    proc_imgproc_Imgproc_GaussianBlur_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static void (*proc_objdetect_HOGDescriptor_detectMultiScale_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_objdetect_HOGDescriptor_detectMultiScale_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_objdetect_HOGDescriptor_detectMultiScale_11 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_photo_AlignMTB_shiftMat_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jdouble,  jdouble);
void Java_org_opencv_photo_AlignMTB_shiftMat_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jdouble a5, jdouble a6) {
    proc_photo_AlignMTB_shiftMat_10 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_video_BackgroundSubtractorKNN_setDist2Threshold_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_video_BackgroundSubtractorKNN_setDist2Threshold_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_video_BackgroundSubtractorKNN_setDist2Threshold_10 (a0, a1, a2, a3);
}

static void (*proc_videoio_VideoCapture_release_10)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_videoio_VideoCapture_release_10(JNIEnv* a0, jclass a1, jlong a2) {
    proc_videoio_VideoCapture_release_10 (a0, a1, a2);
}

static jlong (*proc_ml_TrainData_getCatMap_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_ml_TrainData_getCatMap_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_TrainData_getCatMap_10 (a0, a1, a2);
}

static void (*proc_core_Core_normalize_12)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jint);
void Java_org_opencv_core_Core_normalize_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jint a6) {
    proc_core_Core_normalize_12 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_features2d_ORB_setNLevels_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_features2d_ORB_setNLevels_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_features2d_ORB_setNLevels_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_pyrUp_12)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_imgproc_Imgproc_pyrUp_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_imgproc_Imgproc_pyrUp_12 (a0, a1, a2, a3);
}

static jlong (*proc_ml_TrainData_getCatOfs_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_ml_TrainData_getCatOfs_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_TrainData_getCatOfs_10 (a0, a1, a2);
}

static void (*proc_core_Core_normalize_13)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_core_Core_normalize_13(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_core_Core_normalize_13 (a0, a1, a2, a3);
}

static void (*proc_features2d_ORB_setPatchSize_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_features2d_ORB_setPatchSize_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_features2d_ORB_setPatchSize_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_rectangle_10)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jint,  jint,  jint);
void Java_org_opencv_imgproc_Imgproc_rectangle_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8, jdouble a9, jdouble a10, jint a11, jint a12, jint a13) {
    proc_imgproc_Imgproc_rectangle_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13);
}

static jlong (*proc_ml_TrainData_getClassLabels_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_ml_TrainData_getClassLabels_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_TrainData_getClassLabels_10 (a0, a1, a2);
}

static void (*proc_core_Core_patchNaNs_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_core_Core_patchNaNs_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_core_Core_patchNaNs_10 (a0, a1, a2, a3);
}

static void (*proc_features2d_ORB_setScaleFactor_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_features2d_ORB_setScaleFactor_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_features2d_ORB_setScaleFactor_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_rectangle_11)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jint);
void Java_org_opencv_imgproc_Imgproc_rectangle_11(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8, jdouble a9, jdouble a10, jint a11) {
    proc_imgproc_Imgproc_rectangle_11 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11);
}

static jlong (*proc_ml_TrainData_getDefaultSubstValues_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_ml_TrainData_getDefaultSubstValues_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_TrainData_getDefaultSubstValues_10 (a0, a1, a2);
}

static void (*proc_core_Core_patchNaNs_11)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_core_Core_patchNaNs_11(JNIEnv* a0, jclass a1, jlong a2) {
    proc_core_Core_patchNaNs_11 (a0, a1, a2);
}

static void (*proc_features2d_ORB_setScoreType_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_features2d_ORB_setScoreType_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_features2d_ORB_setScoreType_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_rectangle_12)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_rectangle_12(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8, jdouble a9, jdouble a10) {
    proc_imgproc_Imgproc_rectangle_12 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10);
}

static jint (*proc_ml_TrainData_getLayout_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_ml_TrainData_getLayout_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_TrainData_getLayout_10 (a0, a1, a2);
}

static void (*proc_core_Core_perspectiveTransform_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_core_Core_perspectiveTransform_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_core_Core_perspectiveTransform_10 (a0, a1, a2, a3, a4);
}

static void (*proc_features2d_ORB_setWTA_1K_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_features2d_ORB_setWTA_1K_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_features2d_ORB_setWTA_1K_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_remap_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jint,  jint,  jdouble,  jdouble,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_remap_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jint a6, jint a7, jdouble a8, jdouble a9, jdouble a10, jdouble a11) {
    proc_imgproc_Imgproc_remap_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11);
}

static jlong (*proc_features2d_Params_Params_10)(JNIEnv*,  jclass);
jlong Java_org_opencv_features2d_Params_Params_10(JNIEnv* a0, jclass a1) {
    return proc_features2d_Params_Params_10 (a0, a1);
}

static jlong (*proc_ml_TrainData_getMissing_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_ml_TrainData_getMissing_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_TrainData_getMissing_10 (a0, a1, a2);
}

static void (*proc_core_Core_phase_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jboolean);
void Java_org_opencv_core_Core_phase_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jboolean a5) {
    proc_core_Core_phase_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_imgproc_Imgproc_remap_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jint);
void Java_org_opencv_imgproc_Imgproc_remap_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jint a6) {
    proc_imgproc_Imgproc_remap_11 (a0, a1, a2, a3, a4, a5, a6);
}

static jint (*proc_ml_TrainData_getNAllVars_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_ml_TrainData_getNAllVars_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_TrainData_getNAllVars_10 (a0, a1, a2);
}

static void (*proc_core_Core_phase_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_core_Core_phase_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_core_Core_phase_11 (a0, a1, a2, a3, a4);
}

static void (*proc_features2d_Params_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_features2d_Params_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_features2d_Params_delete (a0, a1, a2);
}

static void (*proc_imgproc_Imgproc_resize_10)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jint);
void Java_org_opencv_imgproc_Imgproc_resize_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7, jint a8) {
    proc_imgproc_Imgproc_resize_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static jboolean (*proc_features2d_Params_get_1filterByArea_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_features2d_Params_get_1filterByArea_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_Params_get_1filterByArea_10 (a0, a1, a2);
}

static jint (*proc_ml_TrainData_getNSamples_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_ml_TrainData_getNSamples_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_TrainData_getNSamples_10 (a0, a1, a2);
}

static void (*proc_core_Core_polarToCart_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jboolean);
void Java_org_opencv_core_Core_polarToCart_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jboolean a6) {
    proc_core_Core_polarToCart_10 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_imgproc_Imgproc_resize_11)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_resize_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5) {
    proc_imgproc_Imgproc_resize_11 (a0, a1, a2, a3, a4, a5);
}

static jboolean (*proc_features2d_Params_get_1filterByCircularity_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_features2d_Params_get_1filterByCircularity_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_Params_get_1filterByCircularity_10 (a0, a1, a2);
}

static jint (*proc_imgproc_Imgproc_rotatedRectangleIntersection_10)(JNIEnv*,  jclass,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jlong);
jint Java_org_opencv_imgproc_Imgproc_rotatedRectangleIntersection_10(JNIEnv* a0, jclass a1, jdouble a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8, jdouble a9, jdouble a10, jdouble a11, jlong a12) {
    return proc_imgproc_Imgproc_rotatedRectangleIntersection_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12);
}

static jint (*proc_ml_TrainData_getNTestSamples_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_ml_TrainData_getNTestSamples_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_TrainData_getNTestSamples_10 (a0, a1, a2);
}

static void (*proc_core_Core_polarToCart_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_core_Core_polarToCart_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_core_Core_polarToCart_11 (a0, a1, a2, a3, a4, a5);
}

static jboolean (*proc_features2d_Params_get_1filterByColor_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_features2d_Params_get_1filterByColor_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_Params_get_1filterByColor_10 (a0, a1, a2);
}

static jint (*proc_ml_TrainData_getNTrainSamples_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_ml_TrainData_getNTrainSamples_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_TrainData_getNTrainSamples_10 (a0, a1, a2);
}

static void (*proc_core_Core_pow_10)(JNIEnv*,  jclass,  jlong,  jdouble,  jlong);
void Java_org_opencv_core_Core_pow_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jlong a4) {
    proc_core_Core_pow_10 (a0, a1, a2, a3, a4);
}

static void (*proc_imgproc_Imgproc_sepFilter2D_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jlong,  jlong,  jdouble,  jdouble,  jdouble,  jint);
void Java_org_opencv_imgproc_Imgproc_sepFilter2D_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jlong a5, jlong a6, jdouble a7, jdouble a8, jdouble a9, jint a10) {
    proc_imgproc_Imgproc_sepFilter2D_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10);
}

static jboolean (*proc_videoio_VideoCapture_retrieve_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint);
jboolean Java_org_opencv_videoio_VideoCapture_retrieve_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4) {
    return proc_videoio_VideoCapture_retrieve_10 (a0, a1, a2, a3, a4);
}

static jlong (*proc_core_Mat_n_1Mat__IIIDDDD)(JNIEnv*,  jclass,  jint,  jint,  jint,  jdouble,  jdouble,  jdouble,  jdouble);
jlong Java_org_opencv_core_Mat_n_1Mat__IIIDDDD(JNIEnv* a0, jclass a1, jint a2, jint a3, jint a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8) {
    return proc_core_Mat_n_1Mat__IIIDDDD (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static jlong (*proc_dnn_Dnn_blobFromImage_11)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_dnn_Dnn_blobFromImage_11(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_dnn_Dnn_blobFromImage_11 (a0, a1, a2);
}

static void (*proc_calib3d_Calib3d_composeRT_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_calib3d_Calib3d_composeRT_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7) {
    proc_calib3d_Calib3d_composeRT_11 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_features2d_AKAZE_setDiffusivity_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_features2d_AKAZE_setDiffusivity_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_features2d_AKAZE_setDiffusivity_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_GaussianBlur_11)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_GaussianBlur_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7) {
    proc_imgproc_Imgproc_GaussianBlur_11 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_ml_ANN_1MLP_setActivationFunction_10)(JNIEnv*,  jclass,  jlong,  jint,  jdouble,  jdouble);
void Java_org_opencv_ml_ANN_1MLP_setActivationFunction_10(JNIEnv* a0, jclass a1, jlong a2, jint a3, jdouble a4, jdouble a5) {
    proc_ml_ANN_1MLP_setActivationFunction_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_objdetect_HOGDescriptor_detect_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jlong);
void Java_org_opencv_objdetect_HOGDescriptor_detect_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jdouble a6, jdouble a7, jdouble a8, jdouble a9, jdouble a10, jlong a11) {
    proc_objdetect_HOGDescriptor_detect_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11);
}

static void (*proc_photo_CalibrateCRF_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_photo_CalibrateCRF_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_photo_CalibrateCRF_delete (a0, a1, a2);
}

static void (*proc_video_BackgroundSubtractorKNN_setHistory_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_video_BackgroundSubtractorKNN_setHistory_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_video_BackgroundSubtractorKNN_setHistory_10 (a0, a1, a2, a3);
}

static jboolean (*proc_features2d_Params_get_1filterByConvexity_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_features2d_Params_get_1filterByConvexity_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_Params_get_1filterByConvexity_10 (a0, a1, a2);
}

static jint (*proc_ml_TrainData_getNVars_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_ml_TrainData_getNVars_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_TrainData_getNVars_10 (a0, a1, a2);
}

static void (*proc_core_Core_randShuffle_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_core_Core_randShuffle_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_core_Core_randShuffle_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_sepFilter2D_11)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jlong,  jlong,  jdouble,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_sepFilter2D_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jlong a5, jlong a6, jdouble a7, jdouble a8, jdouble a9) {
    proc_imgproc_Imgproc_sepFilter2D_11 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static jboolean (*proc_features2d_Params_get_1filterByInertia_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_features2d_Params_get_1filterByInertia_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_Params_get_1filterByInertia_10 (a0, a1, a2);
}

static void (*proc_core_Core_randShuffle_11)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_core_Core_randShuffle_11(JNIEnv* a0, jclass a1, jlong a2) {
    proc_core_Core_randShuffle_11 (a0, a1, a2);
}

static void (*proc_imgproc_Imgproc_sepFilter2D_12)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jlong,  jlong);
void Java_org_opencv_imgproc_Imgproc_sepFilter2D_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jlong a5, jlong a6) {
    proc_imgproc_Imgproc_sepFilter2D_12 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_ml_TrainData_getNames_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_ml_TrainData_getNames_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_ml_TrainData_getNames_10 (a0, a1, a2, a3);
}

static jfloat (*proc_features2d_Params_get_1maxArea_10)(JNIEnv*,  jclass,  jlong);
jfloat Java_org_opencv_features2d_Params_get_1maxArea_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_Params_get_1maxArea_10 (a0, a1, a2);
}

static jlong (*proc_ml_TrainData_getNormCatResponses_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_ml_TrainData_getNormCatResponses_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_TrainData_getNormCatResponses_10 (a0, a1, a2);
}

static void (*proc_core_Core_randn_10)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble);
void Java_org_opencv_core_Core_randn_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4) {
    proc_core_Core_randn_10 (a0, a1, a2, a3, a4);
}

static void (*proc_imgproc_Imgproc_spatialGradient_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jint,  jint);
void Java_org_opencv_imgproc_Imgproc_spatialGradient_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jint a5, jint a6) {
    proc_imgproc_Imgproc_spatialGradient_10 (a0, a1, a2, a3, a4, a5, a6);
}

static jfloat (*proc_features2d_Params_get_1maxCircularity_10)(JNIEnv*,  jclass,  jlong);
jfloat Java_org_opencv_features2d_Params_get_1maxCircularity_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_Params_get_1maxCircularity_10 (a0, a1, a2);
}

static jint (*proc_ml_TrainData_getResponseType_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_ml_TrainData_getResponseType_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_TrainData_getResponseType_10 (a0, a1, a2);
}

static void (*proc_core_Core_randu_10)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble);
void Java_org_opencv_core_Core_randu_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4) {
    proc_core_Core_randu_10 (a0, a1, a2, a3, a4);
}

static void (*proc_imgproc_Imgproc_spatialGradient_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jint);
void Java_org_opencv_imgproc_Imgproc_spatialGradient_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jint a5) {
    proc_imgproc_Imgproc_spatialGradient_11 (a0, a1, a2, a3, a4, a5);
}

static jfloat (*proc_features2d_Params_get_1maxConvexity_10)(JNIEnv*,  jclass,  jlong);
jfloat Java_org_opencv_features2d_Params_get_1maxConvexity_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_Params_get_1maxConvexity_10 (a0, a1, a2);
}

static jlong (*proc_ml_TrainData_getResponses_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_ml_TrainData_getResponses_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_TrainData_getResponses_10 (a0, a1, a2);
}

static void (*proc_core_Core_reduce_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint,  jint);
void Java_org_opencv_core_Core_reduce_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5, jint a6) {
    proc_core_Core_reduce_10 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_imgproc_Imgproc_spatialGradient_12)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_imgproc_Imgproc_spatialGradient_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_imgproc_Imgproc_spatialGradient_12 (a0, a1, a2, a3, a4);
}

static jfloat (*proc_features2d_Params_get_1maxInertiaRatio_10)(JNIEnv*,  jclass,  jlong);
jfloat Java_org_opencv_features2d_Params_get_1maxInertiaRatio_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_Params_get_1maxInertiaRatio_10 (a0, a1, a2);
}

static jlong (*proc_ml_TrainData_getSampleWeights_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_ml_TrainData_getSampleWeights_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_TrainData_getSampleWeights_10 (a0, a1, a2);
}

static void (*proc_core_Core_reduce_11)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint);
void Java_org_opencv_core_Core_reduce_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5) {
    proc_core_Core_reduce_11 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_imgproc_Imgproc_sqrBoxFilter_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jdouble,  jdouble,  jdouble,  jdouble,  jboolean,  jint);
void Java_org_opencv_imgproc_Imgproc_sqrBoxFilter_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8, jboolean a9, jint a10) {
    proc_imgproc_Imgproc_sqrBoxFilter_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10);
}

static jfloat (*proc_features2d_Params_get_1maxThreshold_10)(JNIEnv*,  jclass,  jlong);
jfloat Java_org_opencv_features2d_Params_get_1maxThreshold_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_Params_get_1maxThreshold_10 (a0, a1, a2);
}

static void (*proc_core_Core_repeat_10)(JNIEnv*,  jclass,  jlong,  jint,  jint,  jlong);
void Java_org_opencv_core_Core_repeat_10(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4, jlong a5) {
    proc_core_Core_repeat_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_imgproc_Imgproc_sqrBoxFilter_11)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jdouble,  jdouble,  jdouble,  jdouble,  jboolean);
void Java_org_opencv_imgproc_Imgproc_sqrBoxFilter_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8, jboolean a9) {
    proc_imgproc_Imgproc_sqrBoxFilter_11 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static void (*proc_ml_TrainData_getSample_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jfloat);
void Java_org_opencv_ml_TrainData_getSample_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jfloat a5) {
    proc_ml_TrainData_getSample_10 (a0, a1, a2, a3, a4, a5);
}

static jfloat (*proc_features2d_Params_get_1minArea_10)(JNIEnv*,  jclass,  jlong);
jfloat Java_org_opencv_features2d_Params_get_1minArea_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_Params_get_1minArea_10 (a0, a1, a2);
}

static jlong (*proc_ml_TrainData_getSamples_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_ml_TrainData_getSamples_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_TrainData_getSamples_10 (a0, a1, a2);
}

static void (*proc_core_Core_rotate_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint);
void Java_org_opencv_core_Core_rotate_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4) {
    proc_core_Core_rotate_10 (a0, a1, a2, a3, a4);
}

static void (*proc_imgproc_Imgproc_sqrBoxFilter_12)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_sqrBoxFilter_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jdouble a5, jdouble a6) {
    proc_imgproc_Imgproc_sqrBoxFilter_12 (a0, a1, a2, a3, a4, a5, a6);
}

static jdouble (*proc_imgproc_Imgproc_threshold_10)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jint);
jdouble Java_org_opencv_imgproc_Imgproc_threshold_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jint a6) {
    return proc_imgproc_Imgproc_threshold_10 (a0, a1, a2, a3, a4, a5, a6);
}

static jfloat (*proc_features2d_Params_get_1minCircularity_10)(JNIEnv*,  jclass,  jlong);
jfloat Java_org_opencv_features2d_Params_get_1minCircularity_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_Params_get_1minCircularity_10 (a0, a1, a2);
}

static jlong (*proc_ml_TrainData_getSubVector_10)(JNIEnv*,  jclass,  jlong,  jlong);
jlong Java_org_opencv_ml_TrainData_getSubVector_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_ml_TrainData_getSubVector_10 (a0, a1, a2, a3);
}

static void (*proc_core_Core_scaleAdd_10)(JNIEnv*,  jclass,  jlong,  jdouble,  jlong,  jlong);
void Java_org_opencv_core_Core_scaleAdd_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jlong a4, jlong a5) {
    proc_core_Core_scaleAdd_10 (a0, a1, a2, a3, a4, a5);
}

static jfloat (*proc_features2d_Params_get_1minConvexity_10)(JNIEnv*,  jclass,  jlong);
jfloat Java_org_opencv_features2d_Params_get_1minConvexity_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_Params_get_1minConvexity_10 (a0, a1, a2);
}

static jlong (*proc_ml_TrainData_getTestNormCatResponses_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_ml_TrainData_getTestNormCatResponses_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_TrainData_getTestNormCatResponses_10 (a0, a1, a2);
}

static void (*proc_core_Core_setErrorVerbosity_10)(JNIEnv*,  jclass,  jboolean);
void Java_org_opencv_core_Core_setErrorVerbosity_10(JNIEnv* a0, jclass a1, jboolean a2) {
    proc_core_Core_setErrorVerbosity_10 (a0, a1, a2);
}

static void (*proc_imgproc_Imgproc_undistortPointsIter_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jint,  jint,  jdouble);
void Java_org_opencv_imgproc_Imgproc_undistortPointsIter_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7, jint a8, jint a9, jdouble a10) {
    proc_imgproc_Imgproc_undistortPointsIter_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10);
}

static jboolean (*proc_videoio_VideoCapture_retrieve_11)(JNIEnv*,  jclass,  jlong,  jlong);
jboolean Java_org_opencv_videoio_VideoCapture_retrieve_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_videoio_VideoCapture_retrieve_11 (a0, a1, a2, a3);
}

static jlong (*proc_core_Mat_n_1Mat__DDIDDDD)(JNIEnv*,  jclass,  jdouble,  jdouble,  jint,  jdouble,  jdouble,  jdouble,  jdouble);
jlong Java_org_opencv_core_Mat_n_1Mat__DDIDDDD(JNIEnv* a0, jclass a1, jdouble a2, jdouble a3, jint a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8) {
    return proc_core_Mat_n_1Mat__DDIDDDD (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static jlong (*proc_dnn_Dnn_blobFromImages_10)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jboolean,  jboolean);
jlong Java_org_opencv_dnn_Dnn_blobFromImages_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8, jdouble a9, jboolean a10, jboolean a11) {
    return proc_dnn_Dnn_blobFromImages_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11);
}

static void (*proc_calib3d_Calib3d_computeCorrespondEpilines_10)(JNIEnv*,  jclass,  jlong,  jint,  jlong,  jlong);
void Java_org_opencv_calib3d_Calib3d_computeCorrespondEpilines_10(JNIEnv* a0, jclass a1, jlong a2, jint a3, jlong a4, jlong a5) {
    proc_calib3d_Calib3d_computeCorrespondEpilines_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_features2d_AKAZE_setNOctaveLayers_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_features2d_AKAZE_setNOctaveLayers_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_features2d_AKAZE_setNOctaveLayers_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_GaussianBlur_12)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_GaussianBlur_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jdouble a6) {
    proc_imgproc_Imgproc_GaussianBlur_12 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_ml_ANN_1MLP_setActivationFunction_11)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_ml_ANN_1MLP_setActivationFunction_11(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_ml_ANN_1MLP_setActivationFunction_11 (a0, a1, a2, a3);
}

static void (*proc_objdetect_HOGDescriptor_detect_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_objdetect_HOGDescriptor_detect_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_objdetect_HOGDescriptor_detect_11 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_photo_CalibrateCRF_process_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_photo_CalibrateCRF_process_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_photo_CalibrateCRF_process_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_video_BackgroundSubtractorKNN_setNSamples_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_video_BackgroundSubtractorKNN_setNSamples_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_video_BackgroundSubtractorKNN_setNSamples_10 (a0, a1, a2, a3);
}

static jfloat (*proc_features2d_Params_get_1minDistBetweenBlobs_10)(JNIEnv*,  jclass,  jlong);
jfloat Java_org_opencv_features2d_Params_get_1minDistBetweenBlobs_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_Params_get_1minDistBetweenBlobs_10 (a0, a1, a2);
}

static jlong (*proc_ml_TrainData_getTestResponses_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_ml_TrainData_getTestResponses_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_TrainData_getTestResponses_10 (a0, a1, a2);
}

static void (*proc_core_Core_setIdentity_10)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble);
void Java_org_opencv_core_Core_setIdentity_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6) {
    proc_core_Core_setIdentity_10 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_imgproc_Imgproc_undistortPoints_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_imgproc_Imgproc_undistortPoints_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7) {
    proc_imgproc_Imgproc_undistortPoints_10 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static jfloat (*proc_features2d_Params_get_1minInertiaRatio_10)(JNIEnv*,  jclass,  jlong);
jfloat Java_org_opencv_features2d_Params_get_1minInertiaRatio_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_Params_get_1minInertiaRatio_10 (a0, a1, a2);
}

static jlong (*proc_ml_TrainData_getTestSampleIdx_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_ml_TrainData_getTestSampleIdx_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_TrainData_getTestSampleIdx_10 (a0, a1, a2);
}

static void (*proc_core_Core_setIdentity_11)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_core_Core_setIdentity_11(JNIEnv* a0, jclass a1, jlong a2) {
    proc_core_Core_setIdentity_11 (a0, a1, a2);
}

static void (*proc_imgproc_Imgproc_undistortPoints_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_imgproc_Imgproc_undistortPoints_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_imgproc_Imgproc_undistortPoints_11 (a0, a1, a2, a3, a4, a5);
}

static jlong (*proc_features2d_Params_get_1minRepeatability_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_features2d_Params_get_1minRepeatability_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_Params_get_1minRepeatability_10 (a0, a1, a2);
}

static jlong (*proc_ml_TrainData_getTestSampleWeights_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_ml_TrainData_getTestSampleWeights_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_TrainData_getTestSampleWeights_10 (a0, a1, a2);
}

static void (*proc_core_Core_setNumThreads_10)(JNIEnv*,  jclass,  jint);
void Java_org_opencv_core_Core_setNumThreads_10(JNIEnv* a0, jclass a1, jint a2) {
    proc_core_Core_setNumThreads_10 (a0, a1, a2);
}

static void (*proc_imgproc_Imgproc_undistort_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_imgproc_Imgproc_undistort_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6) {
    proc_imgproc_Imgproc_undistort_10 (a0, a1, a2, a3, a4, a5, a6);
}

static jfloat (*proc_features2d_Params_get_1minThreshold_10)(JNIEnv*,  jclass,  jlong);
jfloat Java_org_opencv_features2d_Params_get_1minThreshold_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_Params_get_1minThreshold_10 (a0, a1, a2);
}

static jlong (*proc_ml_TrainData_getTestSamples_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_ml_TrainData_getTestSamples_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_TrainData_getTestSamples_10 (a0, a1, a2);
}

static void (*proc_core_Core_setRNGSeed_10)(JNIEnv*,  jclass,  jint);
void Java_org_opencv_core_Core_setRNGSeed_10(JNIEnv* a0, jclass a1, jint a2) {
    proc_core_Core_setRNGSeed_10 (a0, a1, a2);
}

static void (*proc_imgproc_Imgproc_undistort_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_imgproc_Imgproc_undistort_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_imgproc_Imgproc_undistort_11 (a0, a1, a2, a3, a4, a5);
}

static jfloat (*proc_features2d_Params_get_1thresholdStep_10)(JNIEnv*,  jclass,  jlong);
jfloat Java_org_opencv_features2d_Params_get_1thresholdStep_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_Params_get_1thresholdStep_10 (a0, a1, a2);
}

static jlong (*proc_ml_TrainData_getTrainNormCatResponses_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_ml_TrainData_getTrainNormCatResponses_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_TrainData_getTrainNormCatResponses_10 (a0, a1, a2);
}

static void (*proc_core_Core_setUseIPP_10)(JNIEnv*,  jclass,  jboolean);
void Java_org_opencv_core_Core_setUseIPP_10(JNIEnv* a0, jclass a1, jboolean a2) {
    proc_core_Core_setUseIPP_10 (a0, a1, a2);
}

static void (*proc_imgproc_Imgproc_warpAffine_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jdouble,  jdouble,  jint,  jint,  jdouble,  jdouble,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_warpAffine_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jdouble a5, jdouble a6, jint a7, jint a8, jdouble a9, jdouble a10, jdouble a11, jdouble a12) {
    proc_imgproc_Imgproc_warpAffine_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12);
}

static jlong (*proc_ml_TrainData_getTrainResponses_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_ml_TrainData_getTrainResponses_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_TrainData_getTrainResponses_10 (a0, a1, a2);
}

static void (*proc_core_Core_setUseIPP_1NE_10)(JNIEnv*,  jclass,  jboolean);
void Java_org_opencv_core_Core_setUseIPP_1NE_10(JNIEnv* a0, jclass a1, jboolean a2) {
    proc_core_Core_setUseIPP_1NE_10 (a0, a1, a2);
}

static void (*proc_features2d_Params_set_1filterByArea_10)(JNIEnv*,  jclass,  jlong,  jboolean);
void Java_org_opencv_features2d_Params_set_1filterByArea_10(JNIEnv* a0, jclass a1, jlong a2, jboolean a3) {
    proc_features2d_Params_set_1filterByArea_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_warpAffine_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jdouble,  jdouble,  jint);
void Java_org_opencv_imgproc_Imgproc_warpAffine_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jdouble a5, jdouble a6, jint a7) {
    proc_imgproc_Imgproc_warpAffine_11 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static jint (*proc_core_Core_solveCubic_10)(JNIEnv*,  jclass,  jlong,  jlong);
jint Java_org_opencv_core_Core_solveCubic_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_core_Core_solveCubic_10 (a0, a1, a2, a3);
}

static jlong (*proc_ml_TrainData_getTrainSampleIdx_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_ml_TrainData_getTrainSampleIdx_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_TrainData_getTrainSampleIdx_10 (a0, a1, a2);
}

static void (*proc_features2d_Params_set_1filterByCircularity_10)(JNIEnv*,  jclass,  jlong,  jboolean);
void Java_org_opencv_features2d_Params_set_1filterByCircularity_10(JNIEnv* a0, jclass a1, jlong a2, jboolean a3) {
    proc_features2d_Params_set_1filterByCircularity_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_warpAffine_12)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_warpAffine_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jdouble a5, jdouble a6) {
    proc_imgproc_Imgproc_warpAffine_12 (a0, a1, a2, a3, a4, a5, a6);
}

static jdouble (*proc_core_Core_solvePoly_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint);
jdouble Java_org_opencv_core_Core_solvePoly_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4) {
    return proc_core_Core_solvePoly_10 (a0, a1, a2, a3, a4);
}

static jlong (*proc_ml_TrainData_getTrainSampleWeights_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_ml_TrainData_getTrainSampleWeights_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_TrainData_getTrainSampleWeights_10 (a0, a1, a2);
}

static void (*proc_features2d_Params_set_1filterByColor_10)(JNIEnv*,  jclass,  jlong,  jboolean);
void Java_org_opencv_features2d_Params_set_1filterByColor_10(JNIEnv* a0, jclass a1, jlong a2, jboolean a3) {
    proc_features2d_Params_set_1filterByColor_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_warpPerspective_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jdouble,  jdouble,  jint,  jint,  jdouble,  jdouble,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_warpPerspective_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jdouble a5, jdouble a6, jint a7, jint a8, jdouble a9, jdouble a10, jdouble a11, jdouble a12) {
    proc_imgproc_Imgproc_warpPerspective_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12);
}

static jdouble (*proc_core_Core_solvePoly_11)(JNIEnv*,  jclass,  jlong,  jlong);
jdouble Java_org_opencv_core_Core_solvePoly_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_core_Core_solvePoly_11 (a0, a1, a2, a3);
}

static jlong (*proc_ml_TrainData_getTrainSamples_10)(JNIEnv*,  jclass,  jlong,  jint,  jboolean,  jboolean);
jlong Java_org_opencv_ml_TrainData_getTrainSamples_10(JNIEnv* a0, jclass a1, jlong a2, jint a3, jboolean a4, jboolean a5) {
    return proc_ml_TrainData_getTrainSamples_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_features2d_Params_set_1filterByConvexity_10)(JNIEnv*,  jclass,  jlong,  jboolean);
void Java_org_opencv_features2d_Params_set_1filterByConvexity_10(JNIEnv* a0, jclass a1, jlong a2, jboolean a3) {
    proc_features2d_Params_set_1filterByConvexity_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_warpPerspective_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jdouble,  jdouble,  jint);
void Java_org_opencv_imgproc_Imgproc_warpPerspective_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jdouble a5, jdouble a6, jint a7) {
    proc_imgproc_Imgproc_warpPerspective_11 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static jboolean (*proc_core_Core_solve_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jint);
jboolean Java_org_opencv_core_Core_solve_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jint a5) {
    return proc_core_Core_solve_10 (a0, a1, a2, a3, a4, a5);
}

static jlong (*proc_ml_TrainData_getTrainSamples_11)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_ml_TrainData_getTrainSamples_11(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_TrainData_getTrainSamples_11 (a0, a1, a2);
}

static void (*proc_features2d_Params_set_1filterByInertia_10)(JNIEnv*,  jclass,  jlong,  jboolean);
void Java_org_opencv_features2d_Params_set_1filterByInertia_10(JNIEnv* a0, jclass a1, jlong a2, jboolean a3) {
    proc_features2d_Params_set_1filterByInertia_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_warpPerspective_12)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_warpPerspective_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jdouble a5, jdouble a6) {
    proc_imgproc_Imgproc_warpPerspective_12 (a0, a1, a2, a3, a4, a5, a6);
}

static jboolean (*proc_videoio_VideoCapture_set_10)(JNIEnv*,  jclass,  jlong,  jint,  jdouble);
jboolean Java_org_opencv_videoio_VideoCapture_set_10(JNIEnv* a0, jclass a1, jlong a2, jint a3, jdouble a4) {
    return proc_videoio_VideoCapture_set_10 (a0, a1, a2, a3, a4);
}

static jlong (*proc_core_Mat_n_1Mat__JIIII)(JNIEnv*,  jclass,  jlong,  jint,  jint,  jint,  jint);
jlong Java_org_opencv_core_Mat_n_1Mat__JIIII(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4, jint a5, jint a6) {
    return proc_core_Mat_n_1Mat__JIIII (a0, a1, a2, a3, a4, a5, a6);
}

static jlong (*proc_dnn_Dnn_blobFromImages_11)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_dnn_Dnn_blobFromImages_11(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_dnn_Dnn_blobFromImages_11 (a0, a1, a2);
}

static jlong (*proc_objdetect_HOGDescriptor_getDaimlerPeopleDetector_10)(JNIEnv*,  jclass);
jlong Java_org_opencv_objdetect_HOGDescriptor_getDaimlerPeopleDetector_10(JNIEnv* a0, jclass a1) {
    return proc_objdetect_HOGDescriptor_getDaimlerPeopleDetector_10 (a0, a1);
}

static void (*proc_calib3d_Calib3d_convertPointsFromHomogeneous_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_calib3d_Calib3d_convertPointsFromHomogeneous_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_calib3d_Calib3d_convertPointsFromHomogeneous_10 (a0, a1, a2, a3);
}

static void (*proc_features2d_AKAZE_setNOctaves_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_features2d_AKAZE_setNOctaves_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_features2d_AKAZE_setNOctaves_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_HoughCircles_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jdouble,  jdouble,  jdouble,  jdouble,  jint,  jint);
void Java_org_opencv_imgproc_Imgproc_HoughCircles_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8, jint a9, jint a10) {
    proc_imgproc_Imgproc_HoughCircles_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10);
}

static void (*proc_ml_ANN_1MLP_setBackpropMomentumScale_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_ml_ANN_1MLP_setBackpropMomentumScale_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_ml_ANN_1MLP_setBackpropMomentumScale_10 (a0, a1, a2, a3);
}

static void (*proc_photo_CalibrateDebevec_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_photo_CalibrateDebevec_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_photo_CalibrateDebevec_delete (a0, a1, a2);
}

static void (*proc_video_BackgroundSubtractorKNN_setShadowThreshold_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_video_BackgroundSubtractorKNN_setShadowThreshold_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_video_BackgroundSubtractorKNN_setShadowThreshold_10 (a0, a1, a2, a3);
}

static jboolean (*proc_core_Core_solve_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
jboolean Java_org_opencv_core_Core_solve_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    return proc_core_Core_solve_11 (a0, a1, a2, a3, a4);
}

static void (*proc_features2d_Params_set_1maxArea_10)(JNIEnv*,  jclass,  jlong,  jfloat);
void Java_org_opencv_features2d_Params_set_1maxArea_10(JNIEnv* a0, jclass a1, jlong a2, jfloat a3) {
    proc_features2d_Params_set_1maxArea_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_watershed_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_imgproc_Imgproc_watershed_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_imgproc_Imgproc_watershed_10 (a0, a1, a2, a3);
}

static void (*proc_ml_TrainData_getValues_10)(JNIEnv*,  jclass,  jlong,  jint,  jlong,  jfloat);
void Java_org_opencv_ml_TrainData_getValues_10(JNIEnv* a0, jclass a1, jlong a2, jint a3, jlong a4, jfloat a5) {
    proc_ml_TrainData_getValues_10 (a0, a1, a2, a3, a4, a5);
}

static jint (*proc_imgproc_LineSegmentDetector_compareSegments_10)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jlong,  jlong,  jlong);
jint Java_org_opencv_imgproc_LineSegmentDetector_compareSegments_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jlong a5, jlong a6, jlong a7) {
    return proc_imgproc_LineSegmentDetector_compareSegments_10 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static jlong (*proc_ml_TrainData_getVarIdx_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_ml_TrainData_getVarIdx_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_TrainData_getVarIdx_10 (a0, a1, a2);
}

static void (*proc_core_Core_sortIdx_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint);
void Java_org_opencv_core_Core_sortIdx_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4) {
    proc_core_Core_sortIdx_10 (a0, a1, a2, a3, a4);
}

static void (*proc_features2d_Params_set_1maxCircularity_10)(JNIEnv*,  jclass,  jlong,  jfloat);
void Java_org_opencv_features2d_Params_set_1maxCircularity_10(JNIEnv* a0, jclass a1, jlong a2, jfloat a3) {
    proc_features2d_Params_set_1maxCircularity_10 (a0, a1, a2, a3);
}

static jint (*proc_imgproc_LineSegmentDetector_compareSegments_11)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jlong,  jlong);
jint Java_org_opencv_imgproc_LineSegmentDetector_compareSegments_11(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jlong a5, jlong a6) {
    return proc_imgproc_LineSegmentDetector_compareSegments_11 (a0, a1, a2, a3, a4, a5, a6);
}

static jlong (*proc_ml_TrainData_getVarSymbolFlags_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_ml_TrainData_getVarSymbolFlags_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_TrainData_getVarSymbolFlags_10 (a0, a1, a2);
}

static void (*proc_core_Core_sort_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint);
void Java_org_opencv_core_Core_sort_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4) {
    proc_core_Core_sort_10 (a0, a1, a2, a3, a4);
}

static void (*proc_features2d_Params_set_1maxConvexity_10)(JNIEnv*,  jclass,  jlong,  jfloat);
void Java_org_opencv_features2d_Params_set_1maxConvexity_10(JNIEnv* a0, jclass a1, jlong a2, jfloat a3) {
    proc_features2d_Params_set_1maxConvexity_10 (a0, a1, a2, a3);
}

static jlong (*proc_ml_TrainData_getVarType_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_ml_TrainData_getVarType_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_TrainData_getVarType_10 (a0, a1, a2);
}

static void (*proc_core_Core_split_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_core_Core_split_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_core_Core_split_10 (a0, a1, a2, a3);
}

static void (*proc_features2d_Params_set_1maxInertiaRatio_10)(JNIEnv*,  jclass,  jlong,  jfloat);
void Java_org_opencv_features2d_Params_set_1maxInertiaRatio_10(JNIEnv* a0, jclass a1, jlong a2, jfloat a3) {
    proc_features2d_Params_set_1maxInertiaRatio_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_LineSegmentDetector_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_imgproc_LineSegmentDetector_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_imgproc_LineSegmentDetector_delete (a0, a1, a2);
}

static void (*proc_core_Core_sqrt_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_core_Core_sqrt_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_core_Core_sqrt_10 (a0, a1, a2, a3);
}

static void (*proc_features2d_Params_set_1maxThreshold_10)(JNIEnv*,  jclass,  jlong,  jfloat);
void Java_org_opencv_features2d_Params_set_1maxThreshold_10(JNIEnv* a0, jclass a1, jlong a2, jfloat a3) {
    proc_features2d_Params_set_1maxThreshold_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_LineSegmentDetector_detect_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_imgproc_LineSegmentDetector_detect_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7) {
    proc_imgproc_LineSegmentDetector_detect_10 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_ml_TrainData_setTrainTestSplitRatio_10)(JNIEnv*,  jclass,  jlong,  jdouble,  jboolean);
void Java_org_opencv_ml_TrainData_setTrainTestSplitRatio_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jboolean a4) {
    proc_ml_TrainData_setTrainTestSplitRatio_10 (a0, a1, a2, a3, a4);
}

static void (*proc_core_Core_subtract_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jint);
void Java_org_opencv_core_Core_subtract_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jint a6) {
    proc_core_Core_subtract_10 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_features2d_Params_set_1minArea_10)(JNIEnv*,  jclass,  jlong,  jfloat);
void Java_org_opencv_features2d_Params_set_1minArea_10(JNIEnv* a0, jclass a1, jlong a2, jfloat a3) {
    proc_features2d_Params_set_1minArea_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_LineSegmentDetector_detect_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_imgproc_LineSegmentDetector_detect_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_imgproc_LineSegmentDetector_detect_11 (a0, a1, a2, a3, a4);
}

static void (*proc_ml_TrainData_setTrainTestSplitRatio_11)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_ml_TrainData_setTrainTestSplitRatio_11(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_ml_TrainData_setTrainTestSplitRatio_11 (a0, a1, a2, a3);
}

static void (*proc_core_Core_subtract_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_core_Core_subtract_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_core_Core_subtract_11 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_features2d_Params_set_1minCircularity_10)(JNIEnv*,  jclass,  jlong,  jfloat);
void Java_org_opencv_features2d_Params_set_1minCircularity_10(JNIEnv* a0, jclass a1, jlong a2, jfloat a3) {
    proc_features2d_Params_set_1minCircularity_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_LineSegmentDetector_drawSegments_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_imgproc_LineSegmentDetector_drawSegments_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_imgproc_LineSegmentDetector_drawSegments_10 (a0, a1, a2, a3, a4);
}

static void (*proc_ml_TrainData_setTrainTestSplit_10)(JNIEnv*,  jclass,  jlong,  jint,  jboolean);
void Java_org_opencv_ml_TrainData_setTrainTestSplit_10(JNIEnv* a0, jclass a1, jlong a2, jint a3, jboolean a4) {
    proc_ml_TrainData_setTrainTestSplit_10 (a0, a1, a2, a3, a4);
}

static jlong (*proc_imgproc_Subdiv2D_Subdiv2D_10)(JNIEnv*,  jclass,  jint,  jint,  jint,  jint);
jlong Java_org_opencv_imgproc_Subdiv2D_Subdiv2D_10(JNIEnv* a0, jclass a1, jint a2, jint a3, jint a4, jint a5) {
    return proc_imgproc_Subdiv2D_Subdiv2D_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_core_Core_subtract_12)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_core_Core_subtract_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_core_Core_subtract_12 (a0, a1, a2, a3, a4);
}

static void (*proc_features2d_Params_set_1minConvexity_10)(JNIEnv*,  jclass,  jlong,  jfloat);
void Java_org_opencv_features2d_Params_set_1minConvexity_10(JNIEnv* a0, jclass a1, jlong a2, jfloat a3) {
    proc_features2d_Params_set_1minConvexity_10 (a0, a1, a2, a3);
}

static void (*proc_ml_TrainData_setTrainTestSplit_11)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_ml_TrainData_setTrainTestSplit_11(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_ml_TrainData_setTrainTestSplit_11 (a0, a1, a2, a3);
}

static jlong (*proc_imgproc_Subdiv2D_Subdiv2D_11)(JNIEnv*,  jclass);
jlong Java_org_opencv_imgproc_Subdiv2D_Subdiv2D_11(JNIEnv* a0, jclass a1) {
    return proc_imgproc_Subdiv2D_Subdiv2D_11 (a0, a1);
}

static void (*proc_core_Core_subtract_13)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jlong,  jlong,  jint);
void Java_org_opencv_core_Core_subtract_13(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jlong a7, jlong a8, jint a9) {
    proc_core_Core_subtract_13 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static void (*proc_features2d_Params_set_1minDistBetweenBlobs_10)(JNIEnv*,  jclass,  jlong,  jfloat);
void Java_org_opencv_features2d_Params_set_1minDistBetweenBlobs_10(JNIEnv* a0, jclass a1, jlong a2, jfloat a3) {
    proc_features2d_Params_set_1minDistBetweenBlobs_10 (a0, a1, a2, a3);
}

static void (*proc_ml_TrainData_shuffleTrainTest_10)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_ml_TrainData_shuffleTrainTest_10(JNIEnv* a0, jclass a1, jlong a2) {
    proc_ml_TrainData_shuffleTrainTest_10 (a0, a1, a2);
}

static void (*proc_core_Core_subtract_14)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jlong,  jlong);
void Java_org_opencv_core_Core_subtract_14(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jlong a7, jlong a8) {
    proc_core_Core_subtract_14 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static void (*proc_features2d_Params_set_1minInertiaRatio_10)(JNIEnv*,  jclass,  jlong,  jfloat);
void Java_org_opencv_features2d_Params_set_1minInertiaRatio_10(JNIEnv* a0, jclass a1, jlong a2, jfloat a3) {
    proc_features2d_Params_set_1minInertiaRatio_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Subdiv2D_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_imgproc_Subdiv2D_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_imgproc_Subdiv2D_delete (a0, a1, a2);
}

static jfloat (*proc_photo_CalibrateDebevec_getLambda_10)(JNIEnv*,  jclass,  jlong);
jfloat Java_org_opencv_photo_CalibrateDebevec_getLambda_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_photo_CalibrateDebevec_getLambda_10 (a0, a1, a2);
}

static jlong (*proc_core_Mat_n_1Mat__JII)(JNIEnv*,  jclass,  jlong,  jint,  jint);
jlong Java_org_opencv_core_Mat_n_1Mat__JII(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4) {
    return proc_core_Mat_n_1Mat__JII (a0, a1, a2, a3, a4);
}

static jlong (*proc_dnn_Dnn_createCaffeImporter_10)(JNIEnv*,  jclass,  jstring,  jstring);
jlong Java_org_opencv_dnn_Dnn_createCaffeImporter_10(JNIEnv* a0, jclass a1, jstring a2, jstring a3) {
    return proc_dnn_Dnn_createCaffeImporter_10 (a0, a1, a2, a3);
}

static jlong (*proc_objdetect_HOGDescriptor_getDefaultPeopleDetector_10)(JNIEnv*,  jclass);
jlong Java_org_opencv_objdetect_HOGDescriptor_getDefaultPeopleDetector_10(JNIEnv* a0, jclass a1) {
    return proc_objdetect_HOGDescriptor_getDefaultPeopleDetector_10 (a0, a1);
}

static jlong (*proc_videoio_VideoWriter_VideoWriter_10)(JNIEnv*,  jclass,  jstring,  jint,  jint,  jdouble,  jdouble,  jdouble,  jboolean);
jlong Java_org_opencv_videoio_VideoWriter_VideoWriter_10(JNIEnv* a0, jclass a1, jstring a2, jint a3, jint a4, jdouble a5, jdouble a6, jdouble a7, jboolean a8) {
    return proc_videoio_VideoWriter_VideoWriter_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static void (*proc_calib3d_Calib3d_convertPointsToHomogeneous_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_calib3d_Calib3d_convertPointsToHomogeneous_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_calib3d_Calib3d_convertPointsToHomogeneous_10 (a0, a1, a2, a3);
}

static void (*proc_features2d_AKAZE_setThreshold_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_features2d_AKAZE_setThreshold_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_features2d_AKAZE_setThreshold_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_HoughCircles_11)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_HoughCircles_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jdouble a5, jdouble a6) {
    proc_imgproc_Imgproc_HoughCircles_11 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_ml_ANN_1MLP_setBackpropWeightScale_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_ml_ANN_1MLP_setBackpropWeightScale_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_ml_ANN_1MLP_setBackpropWeightScale_10 (a0, a1, a2, a3);
}

static void (*proc_video_BackgroundSubtractorKNN_setShadowValue_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_video_BackgroundSubtractorKNN_setShadowValue_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_video_BackgroundSubtractorKNN_setShadowValue_10 (a0, a1, a2, a3);
}

static jint (*proc_imgproc_Subdiv2D_edgeDst_10)(JNIEnv*,  jclass,  jlong,  jint,  jdoubleArray);
jint Java_org_opencv_imgproc_Subdiv2D_edgeDst_10(JNIEnv* a0, jclass a1, jlong a2, jint a3, jdoubleArray a4) {
    return proc_imgproc_Subdiv2D_edgeDst_10 (a0, a1, a2, a3, a4);
}

static void (*proc_core_Core_subtract_15)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jlong);
void Java_org_opencv_core_Core_subtract_15(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jlong a7) {
    proc_core_Core_subtract_15 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_features2d_Params_set_1minRepeatability_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_features2d_Params_set_1minRepeatability_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_features2d_Params_set_1minRepeatability_10 (a0, a1, a2, a3);
}

static jdoubleArray (*proc_core_Core_sumElems_10)(JNIEnv*,  jclass,  jlong);
jdoubleArray Java_org_opencv_core_Core_sumElems_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_core_Core_sumElems_10 (a0, a1, a2);
}

static jint (*proc_imgproc_Subdiv2D_edgeDst_11)(JNIEnv*,  jclass,  jlong,  jint);
jint Java_org_opencv_imgproc_Subdiv2D_edgeDst_11(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    return proc_imgproc_Subdiv2D_edgeDst_11 (a0, a1, a2, a3);
}

static void (*proc_features2d_Params_set_1minThreshold_10)(JNIEnv*,  jclass,  jlong,  jfloat);
void Java_org_opencv_features2d_Params_set_1minThreshold_10(JNIEnv* a0, jclass a1, jlong a2, jfloat a3) {
    proc_features2d_Params_set_1minThreshold_10 (a0, a1, a2, a3);
}

static jdoubleArray (*proc_core_Core_trace_10)(JNIEnv*,  jclass,  jlong);
jdoubleArray Java_org_opencv_core_Core_trace_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_core_Core_trace_10 (a0, a1, a2);
}

static jint (*proc_imgproc_Subdiv2D_edgeOrg_10)(JNIEnv*,  jclass,  jlong,  jint,  jdoubleArray);
jint Java_org_opencv_imgproc_Subdiv2D_edgeOrg_10(JNIEnv* a0, jclass a1, jlong a2, jint a3, jdoubleArray a4) {
    return proc_imgproc_Subdiv2D_edgeOrg_10 (a0, a1, a2, a3, a4);
}

static void (*proc_features2d_Params_set_1thresholdStep_10)(JNIEnv*,  jclass,  jlong,  jfloat);
void Java_org_opencv_features2d_Params_set_1thresholdStep_10(JNIEnv* a0, jclass a1, jlong a2, jfloat a3) {
    proc_features2d_Params_set_1thresholdStep_10 (a0, a1, a2, a3);
}

static jint (*proc_imgproc_Subdiv2D_edgeOrg_11)(JNIEnv*,  jclass,  jlong,  jint);
jint Java_org_opencv_imgproc_Subdiv2D_edgeOrg_11(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    return proc_imgproc_Subdiv2D_edgeOrg_11 (a0, a1, a2, a3);
}

static void (*proc_core_Core_transform_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_core_Core_transform_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_core_Core_transform_10 (a0, a1, a2, a3, a4);
}

static jint (*proc_imgproc_Subdiv2D_findNearest_10)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdoubleArray);
jint Java_org_opencv_imgproc_Subdiv2D_findNearest_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdoubleArray a5) {
    return proc_imgproc_Subdiv2D_findNearest_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_core_Core_transpose_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_core_Core_transpose_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_core_Core_transpose_10 (a0, a1, a2, a3);
}

static jboolean (*proc_core_Core_useIPP_10)(JNIEnv*,  jclass);
jboolean Java_org_opencv_core_Core_useIPP_10(JNIEnv* a0, jclass a1) {
    return proc_core_Core_useIPP_10 (a0, a1);
}

static jint (*proc_imgproc_Subdiv2D_findNearest_11)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble);
jint Java_org_opencv_imgproc_Subdiv2D_findNearest_11(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4) {
    return proc_imgproc_Subdiv2D_findNearest_11 (a0, a1, a2, a3, a4);
}

static jboolean (*proc_core_Core_useIPP_1NE_10)(JNIEnv*,  jclass);
jboolean Java_org_opencv_core_Core_useIPP_1NE_10(JNIEnv* a0, jclass a1) {
    return proc_core_Core_useIPP_1NE_10 (a0, a1);
}

static void (*proc_imgproc_Subdiv2D_getEdgeList_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_imgproc_Subdiv2D_getEdgeList_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_imgproc_Subdiv2D_getEdgeList_10 (a0, a1, a2, a3);
}

static jint (*proc_imgproc_Subdiv2D_getEdge_10)(JNIEnv*,  jclass,  jlong,  jint,  jint);
jint Java_org_opencv_imgproc_Subdiv2D_getEdge_10(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4) {
    return proc_imgproc_Subdiv2D_getEdge_10 (a0, a1, a2, a3, a4);
}

static void (*proc_core_Core_vconcat_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_core_Core_vconcat_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_core_Core_vconcat_10 (a0, a1, a2, a3);
}

static void (*proc_core_Algorithm_clear_10)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_core_Algorithm_clear_10(JNIEnv* a0, jclass a1, jlong a2) {
    proc_core_Algorithm_clear_10 (a0, a1, a2);
}

static void (*proc_imgproc_Subdiv2D_getLeadingEdgeList_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_imgproc_Subdiv2D_getLeadingEdgeList_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_imgproc_Subdiv2D_getLeadingEdgeList_10 (a0, a1, a2, a3);
}

static void (*proc_core_Algorithm_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_core_Algorithm_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_core_Algorithm_delete (a0, a1, a2);
}

static void (*proc_imgproc_Subdiv2D_getTriangleList_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_imgproc_Subdiv2D_getTriangleList_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_imgproc_Subdiv2D_getTriangleList_10 (a0, a1, a2, a3);
}

static jboolean (*proc_photo_CalibrateDebevec_getRandom_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_photo_CalibrateDebevec_getRandom_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_photo_CalibrateDebevec_getRandom_10 (a0, a1, a2);
}

static jlong (*proc_core_Mat_n_1adjustROI)(JNIEnv*,  jclass,  jlong,  jint,  jint,  jint,  jint);
jlong Java_org_opencv_core_Mat_n_1adjustROI(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4, jint a5, jint a6) {
    return proc_core_Mat_n_1adjustROI (a0, a1, a2, a3, a4, a5, a6);
}

static jlong (*proc_dnn_Dnn_createCaffeImporter_11)(JNIEnv*,  jclass,  jstring);
jlong Java_org_opencv_dnn_Dnn_createCaffeImporter_11(JNIEnv* a0, jclass a1, jstring a2) {
    return proc_dnn_Dnn_createCaffeImporter_11 (a0, a1, a2);
}

static jlong (*proc_features2d_AgastFeatureDetector_create_10)(JNIEnv*,  jclass,  jint,  jboolean,  jint);
jlong Java_org_opencv_features2d_AgastFeatureDetector_create_10(JNIEnv* a0, jclass a1, jint a2, jboolean a3, jint a4) {
    return proc_features2d_AgastFeatureDetector_create_10 (a0, a1, a2, a3, a4);
}

static jlong (*proc_objdetect_HOGDescriptor_getDescriptorSize_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_objdetect_HOGDescriptor_getDescriptorSize_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_objdetect_HOGDescriptor_getDescriptorSize_10 (a0, a1, a2);
}

static jlong (*proc_videoio_VideoWriter_VideoWriter_11)(JNIEnv*,  jclass,  jstring,  jint,  jint,  jdouble,  jdouble,  jdouble);
jlong Java_org_opencv_videoio_VideoWriter_VideoWriter_11(JNIEnv* a0, jclass a1, jstring a2, jint a3, jint a4, jdouble a5, jdouble a6, jdouble a7) {
    return proc_videoio_VideoWriter_VideoWriter_11 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_calib3d_Calib3d_correctMatches_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_calib3d_Calib3d_correctMatches_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6) {
    proc_calib3d_Calib3d_correctMatches_10 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_imgproc_Imgproc_HoughLinesP_10)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jint,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_HoughLinesP_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jint a6, jdouble a7, jdouble a8) {
    proc_imgproc_Imgproc_HoughLinesP_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static void (*proc_ml_ANN_1MLP_setLayerSizes_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_ml_ANN_1MLP_setLayerSizes_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_ml_ANN_1MLP_setLayerSizes_10 (a0, a1, a2, a3);
}

static void (*proc_video_BackgroundSubtractorKNN_setkNNSamples_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_video_BackgroundSubtractorKNN_setkNNSamples_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_video_BackgroundSubtractorKNN_setkNNSamples_10 (a0, a1, a2, a3);
}

static jdoubleArray (*proc_imgproc_Subdiv2D_getVertex_10)(JNIEnv*,  jclass,  jlong,  jint,  jdoubleArray);
jdoubleArray Java_org_opencv_imgproc_Subdiv2D_getVertex_10(JNIEnv* a0, jclass a1, jlong a2, jint a3, jdoubleArray a4) {
    return proc_imgproc_Subdiv2D_getVertex_10 (a0, a1, a2, a3, a4);
}

static jstring (*proc_core_Algorithm_getDefaultName_10)(JNIEnv*,  jclass,  jlong);
jstring Java_org_opencv_core_Algorithm_getDefaultName_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_core_Algorithm_getDefaultName_10 (a0, a1, a2);
}

static jdoubleArray (*proc_imgproc_Subdiv2D_getVertex_11)(JNIEnv*,  jclass,  jlong,  jint);
jdoubleArray Java_org_opencv_imgproc_Subdiv2D_getVertex_11(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    return proc_imgproc_Subdiv2D_getVertex_11 (a0, a1, a2, a3);
}

static void (*proc_core_Algorithm_save_10)(JNIEnv*,  jclass,  jlong,  jstring);
void Java_org_opencv_core_Algorithm_save_10(JNIEnv* a0, jclass a1, jlong a2, jstring a3) {
    proc_core_Algorithm_save_10 (a0, a1, a2, a3);
}

static jlong (*proc_core_TickMeter_TickMeter_10)(JNIEnv*,  jclass);
jlong Java_org_opencv_core_TickMeter_TickMeter_10(JNIEnv* a0, jclass a1) {
    return proc_core_TickMeter_TickMeter_10 (a0, a1);
}

static void (*proc_imgproc_Subdiv2D_getVoronoiFacetList_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_imgproc_Subdiv2D_getVoronoiFacetList_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_imgproc_Subdiv2D_getVoronoiFacetList_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_core_TickMeter_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_core_TickMeter_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_core_TickMeter_delete (a0, a1, a2);
}

static void (*proc_imgproc_Subdiv2D_initDelaunay_10)(JNIEnv*,  jclass,  jlong,  jint,  jint,  jint,  jint);
void Java_org_opencv_imgproc_Subdiv2D_initDelaunay_10(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4, jint a5, jint a6) {
    proc_imgproc_Subdiv2D_initDelaunay_10 (a0, a1, a2, a3, a4, a5, a6);
}

static jint (*proc_imgproc_Subdiv2D_insert_10)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble);
jint Java_org_opencv_imgproc_Subdiv2D_insert_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4) {
    return proc_imgproc_Subdiv2D_insert_10 (a0, a1, a2, a3, a4);
}

static jlong (*proc_core_TickMeter_getCounter_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_core_TickMeter_getCounter_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_core_TickMeter_getCounter_10 (a0, a1, a2);
}

static jdouble (*proc_core_TickMeter_getTimeMicro_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_core_TickMeter_getTimeMicro_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_core_TickMeter_getTimeMicro_10 (a0, a1, a2);
}

static void (*proc_imgproc_Subdiv2D_insert_11)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_imgproc_Subdiv2D_insert_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_imgproc_Subdiv2D_insert_11 (a0, a1, a2, a3);
}

static jdouble (*proc_core_TickMeter_getTimeMilli_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_core_TickMeter_getTimeMilli_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_core_TickMeter_getTimeMilli_10 (a0, a1, a2);
}

static jint (*proc_imgproc_Subdiv2D_locate_10)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdoubleArray,  jdoubleArray);
jint Java_org_opencv_imgproc_Subdiv2D_locate_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdoubleArray a5, jdoubleArray a6) {
    return proc_imgproc_Subdiv2D_locate_10 (a0, a1, a2, a3, a4, a5, a6);
}

static jdouble (*proc_core_TickMeter_getTimeSec_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_core_TickMeter_getTimeSec_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_core_TickMeter_getTimeSec_10 (a0, a1, a2);
}

static jint (*proc_imgproc_Subdiv2D_nextEdge_10)(JNIEnv*,  jclass,  jlong,  jint);
jint Java_org_opencv_imgproc_Subdiv2D_nextEdge_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    return proc_imgproc_Subdiv2D_nextEdge_10 (a0, a1, a2, a3);
}

static jint (*proc_imgproc_Subdiv2D_rotateEdge_10)(JNIEnv*,  jclass,  jlong,  jint,  jint);
jint Java_org_opencv_imgproc_Subdiv2D_rotateEdge_10(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4) {
    return proc_imgproc_Subdiv2D_rotateEdge_10 (a0, a1, a2, a3, a4);
}

static jlong (*proc_core_TickMeter_getTimeTicks_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_core_TickMeter_getTimeTicks_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_core_TickMeter_getTimeTicks_10 (a0, a1, a2);
}

static jint (*proc_imgproc_Subdiv2D_symEdge_10)(JNIEnv*,  jclass,  jlong,  jint);
jint Java_org_opencv_imgproc_Subdiv2D_symEdge_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    return proc_imgproc_Subdiv2D_symEdge_10 (a0, a1, a2, a3);
}

static void (*proc_core_TickMeter_reset_10)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_core_TickMeter_reset_10(JNIEnv* a0, jclass a1, jlong a2) {
    proc_core_TickMeter_reset_10 (a0, a1, a2);
}

static jdouble (*proc_objdetect_HOGDescriptor_getWinSigma_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_objdetect_HOGDescriptor_getWinSigma_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_objdetect_HOGDescriptor_getWinSigma_10 (a0, a1, a2);
}

static jint (*proc_photo_CalibrateDebevec_getSamples_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_photo_CalibrateDebevec_getSamples_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_photo_CalibrateDebevec_getSamples_10 (a0, a1, a2);
}

static jlong (*proc_dnn_Dnn_createTensorflowImporter_10)(JNIEnv*,  jclass,  jstring);
jlong Java_org_opencv_dnn_Dnn_createTensorflowImporter_10(JNIEnv* a0, jclass a1, jstring a2) {
    return proc_dnn_Dnn_createTensorflowImporter_10 (a0, a1, a2);
}

static jlong (*proc_features2d_AgastFeatureDetector_create_11)(JNIEnv*,  jclass);
jlong Java_org_opencv_features2d_AgastFeatureDetector_create_11(JNIEnv* a0, jclass a1) {
    return proc_features2d_AgastFeatureDetector_create_11 (a0, a1);
}

static jlong (*proc_videoio_VideoWriter_VideoWriter_12)(JNIEnv*,  jclass,  jstring,  jint,  jdouble,  jdouble,  jdouble,  jboolean);
jlong Java_org_opencv_videoio_VideoWriter_VideoWriter_12(JNIEnv* a0, jclass a1, jstring a2, jint a3, jdouble a4, jdouble a5, jdouble a6, jboolean a7) {
    return proc_videoio_VideoWriter_VideoWriter_12 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_calib3d_Calib3d_decomposeEssentialMat_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_calib3d_Calib3d_decomposeEssentialMat_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_calib3d_Calib3d_decomposeEssentialMat_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_core_Mat_n_1assignTo__JJI)(JNIEnv*,  jclass,  jlong,  jlong,  jint);
void Java_org_opencv_core_Mat_n_1assignTo__JJI(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4) {
    proc_core_Mat_n_1assignTo__JJI (a0, a1, a2, a3, a4);
}

static void (*proc_imgproc_Imgproc_HoughLinesP_11)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jint);
void Java_org_opencv_imgproc_Imgproc_HoughLinesP_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jint a6) {
    proc_imgproc_Imgproc_HoughLinesP_11 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_ml_ANN_1MLP_setRpropDW0_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_ml_ANN_1MLP_setRpropDW0_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_ml_ANN_1MLP_setRpropDW0_10 (a0, a1, a2, a3);
}

static void (*proc_video_BackgroundSubtractorMOG2_apply_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jdouble);
void Java_org_opencv_video_BackgroundSubtractorMOG2_apply_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jdouble a5) {
    proc_video_BackgroundSubtractorMOG2_apply_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_core_TickMeter_start_10)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_core_TickMeter_start_10(JNIEnv* a0, jclass a1, jlong a2) {
    proc_core_TickMeter_start_10 (a0, a1, a2);
}

static void (*proc_core_TickMeter_stop_10)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_core_TickMeter_stop_10(JNIEnv* a0, jclass a1, jlong a2) {
    proc_core_TickMeter_stop_10 (a0, a1, a2);
}

static jdouble (*proc_objdetect_HOGDescriptor_get_1L2HysThreshold_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_objdetect_HOGDescriptor_get_1L2HysThreshold_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_objdetect_HOGDescriptor_get_1L2HysThreshold_10 (a0, a1, a2);
}

static jint (*proc_calib3d_Calib3d_decomposeHomographyMat_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong);
jint Java_org_opencv_calib3d_Calib3d_decomposeHomographyMat_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6) {
    return proc_calib3d_Calib3d_decomposeHomographyMat_10 (a0, a1, a2, a3, a4, a5, a6);
}

static jlong (*proc_dnn_Dnn_createTorchImporter_10)(JNIEnv*,  jclass,  jstring,  jboolean);
jlong Java_org_opencv_dnn_Dnn_createTorchImporter_10(JNIEnv* a0, jclass a1, jstring a2, jboolean a3) {
    return proc_dnn_Dnn_createTorchImporter_10 (a0, a1, a2, a3);
}

static jlong (*proc_videoio_VideoWriter_VideoWriter_13)(JNIEnv*,  jclass,  jstring,  jint,  jdouble,  jdouble,  jdouble);
jlong Java_org_opencv_videoio_VideoWriter_VideoWriter_13(JNIEnv* a0, jclass a1, jstring a2, jint a3, jdouble a4, jdouble a5, jdouble a6) {
    return proc_videoio_VideoWriter_VideoWriter_13 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_core_Mat_n_1assignTo__JJ)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_core_Mat_n_1assignTo__JJ(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_core_Mat_n_1assignTo__JJ (a0, a1, a2, a3);
}

static void (*proc_features2d_AgastFeatureDetector_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_features2d_AgastFeatureDetector_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_features2d_AgastFeatureDetector_delete (a0, a1, a2);
}

static void (*proc_imgproc_Imgproc_HoughLines_10)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jint,  jdouble,  jdouble,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_HoughLines_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jint a6, jdouble a7, jdouble a8, jdouble a9, jdouble a10) {
    proc_imgproc_Imgproc_HoughLines_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10);
}

static void (*proc_ml_ANN_1MLP_setRpropDWMax_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_ml_ANN_1MLP_setRpropDWMax_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_ml_ANN_1MLP_setRpropDWMax_10 (a0, a1, a2, a3);
}

static void (*proc_photo_CalibrateDebevec_setLambda_10)(JNIEnv*,  jclass,  jlong,  jfloat);
void Java_org_opencv_photo_CalibrateDebevec_setLambda_10(JNIEnv* a0, jclass a1, jlong a2, jfloat a3) {
    proc_photo_CalibrateDebevec_setLambda_10 (a0, a1, a2, a3);
}

static void (*proc_video_BackgroundSubtractorMOG2_apply_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_video_BackgroundSubtractorMOG2_apply_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_video_BackgroundSubtractorMOG2_apply_11 (a0, a1, a2, a3, a4);
}

static jdoubleArray (*proc_objdetect_HOGDescriptor_get_1blockSize_10)(JNIEnv*,  jclass,  jlong);
jdoubleArray Java_org_opencv_objdetect_HOGDescriptor_get_1blockSize_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_objdetect_HOGDescriptor_get_1blockSize_10 (a0, a1, a2);
}

static jint (*proc_core_Mat_n_1channels)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_core_Mat_n_1channels(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_core_Mat_n_1channels (a0, a1, a2);
}

static jlong (*proc_dnn_Dnn_createTorchImporter_11)(JNIEnv*,  jclass,  jstring);
jlong Java_org_opencv_dnn_Dnn_createTorchImporter_11(JNIEnv* a0, jclass a1, jstring a2) {
    return proc_dnn_Dnn_createTorchImporter_11 (a0, a1, a2);
}

static jlong (*proc_videoio_VideoWriter_VideoWriter_14)(JNIEnv*,  jclass);
jlong Java_org_opencv_videoio_VideoWriter_VideoWriter_14(JNIEnv* a0, jclass a1) {
    return proc_videoio_VideoWriter_VideoWriter_14 (a0, a1);
}

static jstring (*proc_features2d_AgastFeatureDetector_getDefaultName_10)(JNIEnv*,  jclass,  jlong);
jstring Java_org_opencv_features2d_AgastFeatureDetector_getDefaultName_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_AgastFeatureDetector_getDefaultName_10 (a0, a1, a2);
}

static void (*proc_calib3d_Calib3d_decomposeProjectionMatrix_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_calib3d_Calib3d_decomposeProjectionMatrix_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7, jlong a8, jlong a9) {
    proc_calib3d_Calib3d_decomposeProjectionMatrix_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static void (*proc_imgproc_Imgproc_HoughLines_11)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jint);
void Java_org_opencv_imgproc_Imgproc_HoughLines_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jint a6) {
    proc_imgproc_Imgproc_HoughLines_11 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_ml_ANN_1MLP_setRpropDWMin_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_ml_ANN_1MLP_setRpropDWMin_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_ml_ANN_1MLP_setRpropDWMin_10 (a0, a1, a2, a3);
}

static void (*proc_photo_CalibrateDebevec_setRandom_10)(JNIEnv*,  jclass,  jlong,  jboolean);
void Java_org_opencv_photo_CalibrateDebevec_setRandom_10(JNIEnv* a0, jclass a1, jlong a2, jboolean a3) {
    proc_photo_CalibrateDebevec_setRandom_10 (a0, a1, a2, a3);
}

static void (*proc_video_BackgroundSubtractorMOG2_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_video_BackgroundSubtractorMOG2_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_video_BackgroundSubtractorMOG2_delete (a0, a1, a2);
}

static jboolean (*proc_features2d_AgastFeatureDetector_getNonmaxSuppression_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_features2d_AgastFeatureDetector_getNonmaxSuppression_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_AgastFeatureDetector_getNonmaxSuppression_10 (a0, a1, a2);
}

static jdouble (*proc_video_BackgroundSubtractorMOG2_getBackgroundRatio_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_video_BackgroundSubtractorMOG2_getBackgroundRatio_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_BackgroundSubtractorMOG2_getBackgroundRatio_10 (a0, a1, a2);
}

static jdoubleArray (*proc_objdetect_HOGDescriptor_get_1blockStride_10)(JNIEnv*,  jclass,  jlong);
jdoubleArray Java_org_opencv_objdetect_HOGDescriptor_get_1blockStride_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_objdetect_HOGDescriptor_get_1blockStride_10 (a0, a1, a2);
}

static jint (*proc_core_Mat_n_1checkVector__JIIZ)(JNIEnv*,  jclass,  jlong,  jint,  jint,  jboolean);
jint Java_org_opencv_core_Mat_n_1checkVector__JIIZ(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4, jboolean a5) {
    return proc_core_Mat_n_1checkVector__JIIZ (a0, a1, a2, a3, a4, a5);
}

static jlong (*proc_dnn_Dnn_readNetFromCaffe_10)(JNIEnv*,  jclass,  jstring,  jstring);
jlong Java_org_opencv_dnn_Dnn_readNetFromCaffe_10(JNIEnv* a0, jclass a1, jstring a2, jstring a3) {
    return proc_dnn_Dnn_readNetFromCaffe_10 (a0, a1, a2, a3);
}

static void (*proc_calib3d_Calib3d_decomposeProjectionMatrix_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_calib3d_Calib3d_decomposeProjectionMatrix_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_calib3d_Calib3d_decomposeProjectionMatrix_11 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_imgproc_Imgproc_HuMoments_10)(JNIEnv*,  jclass,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jlong);
void Java_org_opencv_imgproc_Imgproc_HuMoments_10(JNIEnv* a0, jclass a1, jdouble a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8, jdouble a9, jdouble a10, jdouble a11, jlong a12) {
    proc_imgproc_Imgproc_HuMoments_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12);
}

static void (*proc_ml_ANN_1MLP_setRpropDWMinus_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_ml_ANN_1MLP_setRpropDWMinus_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_ml_ANN_1MLP_setRpropDWMinus_10 (a0, a1, a2, a3);
}

static void (*proc_photo_CalibrateDebevec_setSamples_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_photo_CalibrateDebevec_setSamples_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_photo_CalibrateDebevec_setSamples_10 (a0, a1, a2, a3);
}

static void (*proc_videoio_VideoWriter_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_videoio_VideoWriter_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_videoio_VideoWriter_delete (a0, a1, a2);
}

static jdouble (*proc_video_BackgroundSubtractorMOG2_getComplexityReductionThreshold_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_video_BackgroundSubtractorMOG2_getComplexityReductionThreshold_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_BackgroundSubtractorMOG2_getComplexityReductionThreshold_10 (a0, a1, a2);
}

static jdoubleArray (*proc_objdetect_HOGDescriptor_get_1cellSize_10)(JNIEnv*,  jclass,  jlong);
jdoubleArray Java_org_opencv_objdetect_HOGDescriptor_get_1cellSize_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_objdetect_HOGDescriptor_get_1cellSize_10 (a0, a1, a2);
}

static jint (*proc_core_Mat_n_1checkVector__JII)(JNIEnv*,  jclass,  jlong,  jint,  jint);
jint Java_org_opencv_core_Mat_n_1checkVector__JII(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4) {
    return proc_core_Mat_n_1checkVector__JII (a0, a1, a2, a3, a4);
}

static jint (*proc_features2d_AgastFeatureDetector_getThreshold_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_features2d_AgastFeatureDetector_getThreshold_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_AgastFeatureDetector_getThreshold_10 (a0, a1, a2);
}

static jint (*proc_videoio_VideoWriter_fourcc_10)(JNIEnv*,  jclass,  jshort,  jshort,  jshort,  jshort);
jint Java_org_opencv_videoio_VideoWriter_fourcc_10(JNIEnv* a0, jclass a1, jshort a2, jshort a3, jshort a4, jshort a5) {
    return proc_videoio_VideoWriter_fourcc_10 (a0, a1, a2, a3, a4, a5);
}

static jlong (*proc_dnn_Dnn_readNetFromCaffe_11)(JNIEnv*,  jclass,  jstring);
jlong Java_org_opencv_dnn_Dnn_readNetFromCaffe_11(JNIEnv* a0, jclass a1, jstring a2) {
    return proc_dnn_Dnn_readNetFromCaffe_11 (a0, a1, a2);
}

static void (*proc_calib3d_Calib3d_distortPoints_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jdouble);
void Java_org_opencv_calib3d_Calib3d_distortPoints_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jdouble a6) {
    proc_calib3d_Calib3d_distortPoints_10 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_imgproc_Imgproc_Laplacian_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint,  jdouble,  jdouble,  jint);
void Java_org_opencv_imgproc_Imgproc_Laplacian_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5, jdouble a6, jdouble a7, jint a8) {
    proc_imgproc_Imgproc_Laplacian_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static void (*proc_ml_ANN_1MLP_setRpropDWPlus_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_ml_ANN_1MLP_setRpropDWPlus_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_ml_ANN_1MLP_setRpropDWPlus_10 (a0, a1, a2, a3);
}

static void (*proc_photo_CalibrateRobertson_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_photo_CalibrateRobertson_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_photo_CalibrateRobertson_delete (a0, a1, a2);
}

static jboolean (*proc_video_BackgroundSubtractorMOG2_getDetectShadows_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_video_BackgroundSubtractorMOG2_getDetectShadows_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_BackgroundSubtractorMOG2_getDetectShadows_10 (a0, a1, a2);
}

static jdouble (*proc_videoio_VideoWriter_get_10)(JNIEnv*,  jclass,  jlong,  jint);
jdouble Java_org_opencv_videoio_VideoWriter_get_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    return proc_videoio_VideoWriter_get_10 (a0, a1, a2, a3);
}

static jint (*proc_core_Mat_n_1checkVector__JI)(JNIEnv*,  jclass,  jlong,  jint);
jint Java_org_opencv_core_Mat_n_1checkVector__JI(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    return proc_core_Mat_n_1checkVector__JI (a0, a1, a2, a3);
}

static jint (*proc_features2d_AgastFeatureDetector_getType_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_features2d_AgastFeatureDetector_getType_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_AgastFeatureDetector_getType_10 (a0, a1, a2);
}

static jint (*proc_objdetect_HOGDescriptor_get_1derivAperture_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_objdetect_HOGDescriptor_get_1derivAperture_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_objdetect_HOGDescriptor_get_1derivAperture_10 (a0, a1, a2);
}

static jint (*proc_photo_CalibrateRobertson_getMaxIter_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_photo_CalibrateRobertson_getMaxIter_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_photo_CalibrateRobertson_getMaxIter_10 (a0, a1, a2);
}

static jlong (*proc_dnn_Dnn_readNetFromDarknet_10)(JNIEnv*,  jclass,  jstring,  jstring);
jlong Java_org_opencv_dnn_Dnn_readNetFromDarknet_10(JNIEnv* a0, jclass a1, jstring a2, jstring a3) {
    return proc_dnn_Dnn_readNetFromDarknet_10 (a0, a1, a2, a3);
}

static void (*proc_calib3d_Calib3d_distortPoints_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_calib3d_Calib3d_distortPoints_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_calib3d_Calib3d_distortPoints_11 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_imgproc_Imgproc_Laplacian_11)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_Laplacian_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5, jdouble a6, jdouble a7) {
    proc_imgproc_Imgproc_Laplacian_11 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_ml_ANN_1MLP_setTermCriteria_10)(JNIEnv*,  jclass,  jlong,  jint,  jint,  jdouble);
void Java_org_opencv_ml_ANN_1MLP_setTermCriteria_10(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4, jdouble a5) {
    proc_ml_ANN_1MLP_setTermCriteria_10 (a0, a1, a2, a3, a4, a5);
}

static jboolean (*proc_objdetect_HOGDescriptor_get_1gammaCorrection_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_objdetect_HOGDescriptor_get_1gammaCorrection_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_objdetect_HOGDescriptor_get_1gammaCorrection_10 (a0, a1, a2);
}

static jboolean (*proc_videoio_VideoWriter_isOpened_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_videoio_VideoWriter_isOpened_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_videoio_VideoWriter_isOpened_10 (a0, a1, a2);
}

static jint (*proc_video_BackgroundSubtractorMOG2_getHistory_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_video_BackgroundSubtractorMOG2_getHistory_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_BackgroundSubtractorMOG2_getHistory_10 (a0, a1, a2);
}

static jlong (*proc_core_Mat_n_1clone)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_core_Mat_n_1clone(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_core_Mat_n_1clone (a0, a1, a2);
}

static jlong (*proc_dnn_Dnn_readNetFromDarknet_11)(JNIEnv*,  jclass,  jstring);
jlong Java_org_opencv_dnn_Dnn_readNetFromDarknet_11(JNIEnv* a0, jclass a1, jstring a2) {
    return proc_dnn_Dnn_readNetFromDarknet_11 (a0, a1, a2);
}

static jlong (*proc_photo_CalibrateRobertson_getRadiance_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_photo_CalibrateRobertson_getRadiance_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_photo_CalibrateRobertson_getRadiance_10 (a0, a1, a2);
}

static void (*proc_calib3d_Calib3d_drawChessboardCorners_10)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jlong,  jboolean);
void Java_org_opencv_calib3d_Calib3d_drawChessboardCorners_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jlong a5, jboolean a6) {
    proc_calib3d_Calib3d_drawChessboardCorners_10 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_features2d_AgastFeatureDetector_setNonmaxSuppression_10)(JNIEnv*,  jclass,  jlong,  jboolean);
void Java_org_opencv_features2d_AgastFeatureDetector_setNonmaxSuppression_10(JNIEnv* a0, jclass a1, jlong a2, jboolean a3) {
    proc_features2d_AgastFeatureDetector_setNonmaxSuppression_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_Laplacian_12)(JNIEnv*,  jclass,  jlong,  jlong,  jint);
void Java_org_opencv_imgproc_Imgproc_Laplacian_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4) {
    proc_imgproc_Imgproc_Laplacian_12 (a0, a1, a2, a3, a4);
}

static void (*proc_ml_ANN_1MLP_setTrainMethod_10)(JNIEnv*,  jclass,  jlong,  jint,  jdouble,  jdouble);
void Java_org_opencv_ml_ANN_1MLP_setTrainMethod_10(JNIEnv* a0, jclass a1, jlong a2, jint a3, jdouble a4, jdouble a5) {
    proc_ml_ANN_1MLP_setTrainMethod_10 (a0, a1, a2, a3, a4, a5);
}

static jboolean (*proc_videoio_VideoWriter_open_10)(JNIEnv*,  jclass,  jlong,  jstring,  jint,  jint,  jdouble,  jdouble,  jdouble,  jboolean);
jboolean Java_org_opencv_videoio_VideoWriter_open_10(JNIEnv* a0, jclass a1, jlong a2, jstring a3, jint a4, jint a5, jdouble a6, jdouble a7, jdouble a8, jboolean a9) {
    return proc_videoio_VideoWriter_open_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static jfloat (*proc_photo_CalibrateRobertson_getThreshold_10)(JNIEnv*,  jclass,  jlong);
jfloat Java_org_opencv_photo_CalibrateRobertson_getThreshold_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_photo_CalibrateRobertson_getThreshold_10 (a0, a1, a2);
}

static jint (*proc_objdetect_HOGDescriptor_get_1histogramNormType_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_objdetect_HOGDescriptor_get_1histogramNormType_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_objdetect_HOGDescriptor_get_1histogramNormType_10 (a0, a1, a2);
}

static jint (*proc_video_BackgroundSubtractorMOG2_getNMixtures_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_video_BackgroundSubtractorMOG2_getNMixtures_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_BackgroundSubtractorMOG2_getNMixtures_10 (a0, a1, a2);
}

static jlong (*proc_calib3d_Calib3d_estimateAffine2D_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jint,  jdouble,  jlong,  jdouble,  jlong);
jlong Java_org_opencv_calib3d_Calib3d_estimateAffine2D_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jint a5, jdouble a6, jlong a7, jdouble a8, jlong a9) {
    return proc_calib3d_Calib3d_estimateAffine2D_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static jlong (*proc_core_Mat_n_1col)(JNIEnv*,  jclass,  jlong,  jint);
jlong Java_org_opencv_core_Mat_n_1col(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    return proc_core_Mat_n_1col (a0, a1, a2, a3);
}

static jlong (*proc_dnn_Dnn_readNetFromTensorflow_10)(JNIEnv*,  jclass,  jstring,  jstring);
jlong Java_org_opencv_dnn_Dnn_readNetFromTensorflow_10(JNIEnv* a0, jclass a1, jstring a2, jstring a3) {
    return proc_dnn_Dnn_readNetFromTensorflow_10 (a0, a1, a2, a3);
}

static void (*proc_features2d_AgastFeatureDetector_setThreshold_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_features2d_AgastFeatureDetector_setThreshold_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_features2d_AgastFeatureDetector_setThreshold_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_Scharr_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint,  jint,  jdouble,  jdouble,  jint);
void Java_org_opencv_imgproc_Imgproc_Scharr_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5, jint a6, jdouble a7, jdouble a8, jint a9) {
    proc_imgproc_Imgproc_Scharr_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static void (*proc_ml_ANN_1MLP_setTrainMethod_11)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_ml_ANN_1MLP_setTrainMethod_11(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_ml_ANN_1MLP_setTrainMethod_11 (a0, a1, a2, a3);
}

static jboolean (*proc_videoio_VideoWriter_open_11)(JNIEnv*,  jclass,  jlong,  jstring,  jint,  jint,  jdouble,  jdouble,  jdouble);
jboolean Java_org_opencv_videoio_VideoWriter_open_11(JNIEnv* a0, jclass a1, jlong a2, jstring a3, jint a4, jint a5, jdouble a6, jdouble a7, jdouble a8) {
    return proc_videoio_VideoWriter_open_11 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static jdouble (*proc_video_BackgroundSubtractorMOG2_getShadowThreshold_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_video_BackgroundSubtractorMOG2_getShadowThreshold_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_BackgroundSubtractorMOG2_getShadowThreshold_10 (a0, a1, a2);
}

static jint (*proc_objdetect_HOGDescriptor_get_1nbins_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_objdetect_HOGDescriptor_get_1nbins_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_objdetect_HOGDescriptor_get_1nbins_10 (a0, a1, a2);
}

static jlong (*proc_calib3d_Calib3d_estimateAffine2D_11)(JNIEnv*,  jclass,  jlong,  jlong);
jlong Java_org_opencv_calib3d_Calib3d_estimateAffine2D_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_calib3d_Calib3d_estimateAffine2D_11 (a0, a1, a2, a3);
}

static jlong (*proc_core_Mat_n_1colRange)(JNIEnv*,  jclass,  jlong,  jint,  jint);
jlong Java_org_opencv_core_Mat_n_1colRange(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4) {
    return proc_core_Mat_n_1colRange (a0, a1, a2, a3, a4);
}

static jlong (*proc_dnn_Dnn_readNetFromTensorflow_11)(JNIEnv*,  jclass,  jstring);
jlong Java_org_opencv_dnn_Dnn_readNetFromTensorflow_11(JNIEnv* a0, jclass a1, jstring a2) {
    return proc_dnn_Dnn_readNetFromTensorflow_11 (a0, a1, a2);
}

static jlong (*proc_ml_Boost_create_10)(JNIEnv*,  jclass);
jlong Java_org_opencv_ml_Boost_create_10(JNIEnv* a0, jclass a1) {
    return proc_ml_Boost_create_10 (a0, a1);
}

static void (*proc_features2d_AgastFeatureDetector_setType_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_features2d_AgastFeatureDetector_setType_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_features2d_AgastFeatureDetector_setType_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_Scharr_11)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint,  jint,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_Scharr_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5, jint a6, jdouble a7, jdouble a8) {
    proc_imgproc_Imgproc_Scharr_11 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static void (*proc_photo_CalibrateRobertson_setMaxIter_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_photo_CalibrateRobertson_setMaxIter_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_photo_CalibrateRobertson_setMaxIter_10 (a0, a1, a2, a3);
}

static jboolean (*proc_videoio_VideoWriter_open_12)(JNIEnv*,  jclass,  jlong,  jstring,  jint,  jdouble,  jdouble,  jdouble,  jboolean);
jboolean Java_org_opencv_videoio_VideoWriter_open_12(JNIEnv* a0, jclass a1, jlong a2, jstring a3, jint a4, jdouble a5, jdouble a6, jdouble a7, jboolean a8) {
    return proc_videoio_VideoWriter_open_12 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static jint (*proc_calib3d_Calib3d_estimateAffine3D_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jdouble,  jdouble);
jint Java_org_opencv_calib3d_Calib3d_estimateAffine3D_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jdouble a6, jdouble a7) {
    return proc_calib3d_Calib3d_estimateAffine3D_10 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static jint (*proc_core_Mat_n_1cols)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_core_Mat_n_1cols(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_core_Mat_n_1cols (a0, a1, a2);
}

static jint (*proc_objdetect_HOGDescriptor_get_1nlevels_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_objdetect_HOGDescriptor_get_1nlevels_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_objdetect_HOGDescriptor_get_1nlevels_10 (a0, a1, a2);
}

static jint (*proc_video_BackgroundSubtractorMOG2_getShadowValue_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_video_BackgroundSubtractorMOG2_getShadowValue_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_BackgroundSubtractorMOG2_getShadowValue_10 (a0, a1, a2);
}

static jlong (*proc_dnn_Dnn_readNetFromTorch_10)(JNIEnv*,  jclass,  jstring,  jboolean);
jlong Java_org_opencv_dnn_Dnn_readNetFromTorch_10(JNIEnv* a0, jclass a1, jstring a2, jboolean a3) {
    return proc_dnn_Dnn_readNetFromTorch_10 (a0, a1, a2, a3);
}

static jlong (*proc_features2d_BFMatcher_BFMatcher_10)(JNIEnv*,  jclass,  jint,  jboolean);
jlong Java_org_opencv_features2d_BFMatcher_BFMatcher_10(JNIEnv* a0, jclass a1, jint a2, jboolean a3) {
    return proc_features2d_BFMatcher_BFMatcher_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_Scharr_12)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint,  jint);
void Java_org_opencv_imgproc_Imgproc_Scharr_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5, jint a6) {
    proc_imgproc_Imgproc_Scharr_12 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_ml_Boost_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_ml_Boost_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_ml_Boost_delete (a0, a1, a2);
}

static void (*proc_photo_CalibrateRobertson_setThreshold_10)(JNIEnv*,  jclass,  jlong,  jfloat);
void Java_org_opencv_photo_CalibrateRobertson_setThreshold_10(JNIEnv* a0, jclass a1, jlong a2, jfloat a3) {
    proc_photo_CalibrateRobertson_setThreshold_10 (a0, a1, a2, a3);
}

static jboolean (*proc_objdetect_HOGDescriptor_get_1signedGradient_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_objdetect_HOGDescriptor_get_1signedGradient_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_objdetect_HOGDescriptor_get_1signedGradient_10 (a0, a1, a2);
}

static jboolean (*proc_videoio_VideoWriter_open_13)(JNIEnv*,  jclass,  jlong,  jstring,  jint,  jdouble,  jdouble,  jdouble);
jboolean Java_org_opencv_videoio_VideoWriter_open_13(JNIEnv* a0, jclass a1, jlong a2, jstring a3, jint a4, jdouble a5, jdouble a6, jdouble a7) {
    return proc_videoio_VideoWriter_open_13 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static jdouble (*proc_video_BackgroundSubtractorMOG2_getVarInit_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_video_BackgroundSubtractorMOG2_getVarInit_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_BackgroundSubtractorMOG2_getVarInit_10 (a0, a1, a2);
}

static jint (*proc_calib3d_Calib3d_estimateAffine3D_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
jint Java_org_opencv_calib3d_Calib3d_estimateAffine3D_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    return proc_calib3d_Calib3d_estimateAffine3D_11 (a0, a1, a2, a3, a4, a5);
}

static jint (*proc_ml_Boost_getBoostType_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_ml_Boost_getBoostType_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_Boost_getBoostType_10 (a0, a1, a2);
}

static jlong (*proc_dnn_Dnn_readNetFromTorch_11)(JNIEnv*,  jclass,  jstring);
jlong Java_org_opencv_dnn_Dnn_readNetFromTorch_11(JNIEnv* a0, jclass a1, jstring a2) {
    return proc_dnn_Dnn_readNetFromTorch_11 (a0, a1, a2);
}

static jlong (*proc_features2d_BFMatcher_BFMatcher_11)(JNIEnv*,  jclass);
jlong Java_org_opencv_features2d_BFMatcher_BFMatcher_11(JNIEnv* a0, jclass a1) {
    return proc_features2d_BFMatcher_BFMatcher_11 (a0, a1);
}

static void (*proc_core_Mat_n_1convertTo__JJIDD)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jdouble,  jdouble);
void Java_org_opencv_core_Mat_n_1convertTo__JJIDD(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jdouble a5, jdouble a6) {
    proc_core_Mat_n_1convertTo__JJIDD (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_imgproc_Imgproc_Sobel_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint,  jint,  jint,  jdouble,  jdouble,  jint);
void Java_org_opencv_imgproc_Imgproc_Sobel_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5, jint a6, jint a7, jdouble a8, jdouble a9, jint a10) {
    proc_imgproc_Imgproc_Sobel_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10);
}

static void (*proc_photo_MergeDebevec_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_photo_MergeDebevec_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_photo_MergeDebevec_delete (a0, a1, a2);
}

static jdouble (*proc_video_BackgroundSubtractorMOG2_getVarMax_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_video_BackgroundSubtractorMOG2_getVarMax_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_BackgroundSubtractorMOG2_getVarMax_10 (a0, a1, a2);
}

static jint (*proc_ml_Boost_getWeakCount_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_ml_Boost_getWeakCount_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_Boost_getWeakCount_10 (a0, a1, a2);
}

static jlong (*proc_calib3d_Calib3d_estimateAffinePartial2D_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jint,  jdouble,  jlong,  jdouble,  jlong);
jlong Java_org_opencv_calib3d_Calib3d_estimateAffinePartial2D_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jint a5, jdouble a6, jlong a7, jdouble a8, jlong a9) {
    return proc_calib3d_Calib3d_estimateAffinePartial2D_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static jlong (*proc_dnn_Dnn_readTorchBlob_10)(JNIEnv*,  jclass,  jstring,  jboolean);
jlong Java_org_opencv_dnn_Dnn_readTorchBlob_10(JNIEnv* a0, jclass a1, jstring a2, jboolean a3) {
    return proc_dnn_Dnn_readTorchBlob_10 (a0, a1, a2, a3);
}

static jlong (*proc_features2d_BFMatcher_create_10)(JNIEnv*,  jclass,  jint,  jboolean);
jlong Java_org_opencv_features2d_BFMatcher_create_10(JNIEnv* a0, jclass a1, jint a2, jboolean a3) {
    return proc_features2d_BFMatcher_create_10 (a0, a1, a2, a3);
}

static jlong (*proc_objdetect_HOGDescriptor_get_1svmDetector_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_objdetect_HOGDescriptor_get_1svmDetector_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_objdetect_HOGDescriptor_get_1svmDetector_10 (a0, a1, a2);
}

static void (*proc_core_Mat_n_1convertTo__JJID)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jdouble);
void Java_org_opencv_core_Mat_n_1convertTo__JJID(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jdouble a5) {
    proc_core_Mat_n_1convertTo__JJID (a0, a1, a2, a3, a4, a5);
}

static void (*proc_imgproc_Imgproc_Sobel_11)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint,  jint,  jint,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_Sobel_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5, jint a6, jint a7, jdouble a8, jdouble a9) {
    proc_imgproc_Imgproc_Sobel_11 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static void (*proc_photo_MergeDebevec_process_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_photo_MergeDebevec_process_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6) {
    proc_photo_MergeDebevec_process_10 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_videoio_VideoWriter_release_10)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_videoio_VideoWriter_release_10(JNIEnv* a0, jclass a1, jlong a2) {
    proc_videoio_VideoWriter_release_10 (a0, a1, a2);
}

static jboolean (*proc_videoio_VideoWriter_set_10)(JNIEnv*,  jclass,  jlong,  jint,  jdouble);
jboolean Java_org_opencv_videoio_VideoWriter_set_10(JNIEnv* a0, jclass a1, jlong a2, jint a3, jdouble a4) {
    return proc_videoio_VideoWriter_set_10 (a0, a1, a2, a3, a4);
}

static jdouble (*proc_ml_Boost_getWeightTrimRate_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_ml_Boost_getWeightTrimRate_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_Boost_getWeightTrimRate_10 (a0, a1, a2);
}

static jdouble (*proc_objdetect_HOGDescriptor_get_1winSigma_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_objdetect_HOGDescriptor_get_1winSigma_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_objdetect_HOGDescriptor_get_1winSigma_10 (a0, a1, a2);
}

static jdouble (*proc_video_BackgroundSubtractorMOG2_getVarMin_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_video_BackgroundSubtractorMOG2_getVarMin_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_BackgroundSubtractorMOG2_getVarMin_10 (a0, a1, a2);
}

static jlong (*proc_calib3d_Calib3d_estimateAffinePartial2D_11)(JNIEnv*,  jclass,  jlong,  jlong);
jlong Java_org_opencv_calib3d_Calib3d_estimateAffinePartial2D_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_calib3d_Calib3d_estimateAffinePartial2D_11 (a0, a1, a2, a3);
}

static jlong (*proc_dnn_Dnn_readTorchBlob_11)(JNIEnv*,  jclass,  jstring);
jlong Java_org_opencv_dnn_Dnn_readTorchBlob_11(JNIEnv* a0, jclass a1, jstring a2) {
    return proc_dnn_Dnn_readTorchBlob_11 (a0, a1, a2);
}

static jlong (*proc_features2d_BFMatcher_create_11)(JNIEnv*,  jclass);
jlong Java_org_opencv_features2d_BFMatcher_create_11(JNIEnv* a0, jclass a1) {
    return proc_features2d_BFMatcher_create_11 (a0, a1);
}

static void (*proc_core_Mat_n_1convertTo__JJI)(JNIEnv*,  jclass,  jlong,  jlong,  jint);
void Java_org_opencv_core_Mat_n_1convertTo__JJI(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4) {
    proc_core_Mat_n_1convertTo__JJI (a0, a1, a2, a3, a4);
}

static void (*proc_imgproc_Imgproc_Sobel_12)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint,  jint);
void Java_org_opencv_imgproc_Imgproc_Sobel_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5, jint a6) {
    proc_imgproc_Imgproc_Sobel_12 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_photo_MergeDebevec_process_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_photo_MergeDebevec_process_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_photo_MergeDebevec_process_11 (a0, a1, a2, a3, a4, a5);
}

static jdouble (*proc_video_BackgroundSubtractorMOG2_getVarThresholdGen_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_video_BackgroundSubtractorMOG2_getVarThresholdGen_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_BackgroundSubtractorMOG2_getVarThresholdGen_10 (a0, a1, a2);
}

static jdoubleArray (*proc_objdetect_HOGDescriptor_get_1winSize_10)(JNIEnv*,  jclass,  jlong);
jdoubleArray Java_org_opencv_objdetect_HOGDescriptor_get_1winSize_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_objdetect_HOGDescriptor_get_1winSize_10 (a0, a1, a2);
}

static jlong (*proc_ml_Boost_load_10)(JNIEnv*,  jclass,  jstring,  jstring);
jlong Java_org_opencv_ml_Boost_load_10(JNIEnv* a0, jclass a1, jstring a2, jstring a3) {
    return proc_ml_Boost_load_10 (a0, a1, a2, a3);
}

static void (*proc_calib3d_Calib3d_estimateNewCameraMatrixForUndistortRectify_10)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jlong,  jlong,  jdouble,  jdouble,  jdouble,  jdouble);
void Java_org_opencv_calib3d_Calib3d_estimateNewCameraMatrixForUndistortRectify_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jlong a6, jlong a7, jdouble a8, jdouble a9, jdouble a10, jdouble a11) {
    proc_calib3d_Calib3d_estimateNewCameraMatrixForUndistortRectify_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11);
}

static void (*proc_core_Mat_n_1copyTo__JJ)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_core_Mat_n_1copyTo__JJ(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_core_Mat_n_1copyTo__JJ (a0, a1, a2, a3);
}

static void (*proc_dnn_Dnn_shrinkCaffeModel_10)(JNIEnv*,  jclass,  jstring,  jstring);
void Java_org_opencv_dnn_Dnn_shrinkCaffeModel_10(JNIEnv* a0, jclass a1, jstring a2, jstring a3) {
    proc_dnn_Dnn_shrinkCaffeModel_10 (a0, a1, a2, a3);
}

static void (*proc_features2d_BFMatcher_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_features2d_BFMatcher_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_features2d_BFMatcher_delete (a0, a1, a2);
}

static void (*proc_imgproc_Imgproc_accumulateProduct_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_imgproc_Imgproc_accumulateProduct_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_imgproc_Imgproc_accumulateProduct_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_photo_MergeExposures_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_photo_MergeExposures_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_photo_MergeExposures_delete (a0, a1, a2);
}

static void (*proc_videoio_VideoWriter_write_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_videoio_VideoWriter_write_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_videoio_VideoWriter_write_10 (a0, a1, a2, a3);
}

static jboolean (*proc_objdetect_HOGDescriptor_load_10)(JNIEnv*,  jclass,  jlong,  jstring,  jstring);
jboolean Java_org_opencv_objdetect_HOGDescriptor_load_10(JNIEnv* a0, jclass a1, jlong a2, jstring a3, jstring a4) {
    return proc_objdetect_HOGDescriptor_load_10 (a0, a1, a2, a3, a4);
}

static jdouble (*proc_video_BackgroundSubtractorMOG2_getVarThreshold_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_video_BackgroundSubtractorMOG2_getVarThreshold_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_BackgroundSubtractorMOG2_getVarThreshold_10 (a0, a1, a2);
}

static jlong (*proc_ml_Boost_load_11)(JNIEnv*,  jclass,  jstring);
jlong Java_org_opencv_ml_Boost_load_11(JNIEnv* a0, jclass a1, jstring a2) {
    return proc_ml_Boost_load_11 (a0, a1, a2);
}

static void (*proc_calib3d_Calib3d_estimateNewCameraMatrixForUndistortRectify_11)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jlong,  jlong);
void Java_org_opencv_calib3d_Calib3d_estimateNewCameraMatrixForUndistortRectify_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jlong a6, jlong a7) {
    proc_calib3d_Calib3d_estimateNewCameraMatrixForUndistortRectify_11 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_core_Mat_n_1copyTo__JJJ)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_core_Mat_n_1copyTo__JJJ(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_core_Mat_n_1copyTo__JJJ (a0, a1, a2, a3, a4);
}

static void (*proc_dnn_Importer_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_dnn_Importer_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_dnn_Importer_delete (a0, a1, a2);
}

static void (*proc_features2d_BOWImgDescriptorExtractor_compute_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_features2d_BOWImgDescriptorExtractor_compute_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_features2d_BOWImgDescriptorExtractor_compute_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_imgproc_Imgproc_accumulateProduct_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_imgproc_Imgproc_accumulateProduct_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_imgproc_Imgproc_accumulateProduct_11 (a0, a1, a2, a3, a4);
}

static void (*proc_photo_MergeExposures_process_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_photo_MergeExposures_process_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6) {
    proc_photo_MergeExposures_process_10 (a0, a1, a2, a3, a4, a5, a6);
}

static jboolean (*proc_objdetect_HOGDescriptor_load_11)(JNIEnv*,  jclass,  jlong,  jstring);
jboolean Java_org_opencv_objdetect_HOGDescriptor_load_11(JNIEnv* a0, jclass a1, jlong a2, jstring a3) {
    return proc_objdetect_HOGDescriptor_load_11 (a0, a1, a2, a3);
}

static void (*proc_calib3d_Calib3d_filterSpeckles_10)(JNIEnv*,  jclass,  jlong,  jdouble,  jint,  jdouble,  jlong);
void Java_org_opencv_calib3d_Calib3d_filterSpeckles_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jint a4, jdouble a5, jlong a6) {
    proc_calib3d_Calib3d_filterSpeckles_10 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_core_Mat_n_1create__JIII)(JNIEnv*,  jclass,  jlong,  jint,  jint,  jint);
void Java_org_opencv_core_Mat_n_1create__JIII(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4, jint a5) {
    proc_core_Mat_n_1create__JIII (a0, a1, a2, a3, a4, a5);
}

static void (*proc_dnn_Importer_populateNet_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_dnn_Importer_populateNet_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_dnn_Importer_populateNet_10 (a0, a1, a2, a3);
}

static void (*proc_features2d_BOWImgDescriptorExtractor_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_features2d_BOWImgDescriptorExtractor_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_features2d_BOWImgDescriptorExtractor_delete (a0, a1, a2);
}

static void (*proc_imgproc_Imgproc_accumulateSquare_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_imgproc_Imgproc_accumulateSquare_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_imgproc_Imgproc_accumulateSquare_10 (a0, a1, a2, a3, a4);
}

static void (*proc_ml_Boost_setBoostType_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_ml_Boost_setBoostType_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_ml_Boost_setBoostType_10 (a0, a1, a2, a3);
}

static void (*proc_photo_MergeMertens_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_photo_MergeMertens_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_photo_MergeMertens_delete (a0, a1, a2);
}

static void (*proc_video_BackgroundSubtractorMOG2_setBackgroundRatio_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_video_BackgroundSubtractorMOG2_setBackgroundRatio_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_video_BackgroundSubtractorMOG2_setBackgroundRatio_10 (a0, a1, a2, a3);
}

static jfloat (*proc_photo_MergeMertens_getContrastWeight_10)(JNIEnv*,  jclass,  jlong);
jfloat Java_org_opencv_photo_MergeMertens_getContrastWeight_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_photo_MergeMertens_getContrastWeight_10 (a0, a1, a2);
}

static jint (*proc_features2d_BOWImgDescriptorExtractor_descriptorSize_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_features2d_BOWImgDescriptorExtractor_descriptorSize_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_BOWImgDescriptorExtractor_descriptorSize_10 (a0, a1, a2);
}

static void (*proc_calib3d_Calib3d_filterSpeckles_11)(JNIEnv*,  jclass,  jlong,  jdouble,  jint,  jdouble);
void Java_org_opencv_calib3d_Calib3d_filterSpeckles_11(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jint a4, jdouble a5) {
    proc_calib3d_Calib3d_filterSpeckles_11 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_core_Mat_n_1create__JDDI)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jint);
void Java_org_opencv_core_Mat_n_1create__JDDI(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jint a5) {
    proc_core_Mat_n_1create__JDDI (a0, a1, a2, a3, a4, a5);
}

static void (*proc_dnn_Layer_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_dnn_Layer_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_dnn_Layer_delete (a0, a1, a2);
}

static void (*proc_imgproc_Imgproc_accumulateSquare_11)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_imgproc_Imgproc_accumulateSquare_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_imgproc_Imgproc_accumulateSquare_11 (a0, a1, a2, a3);
}

static void (*proc_ml_Boost_setWeakCount_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_ml_Boost_setWeakCount_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_ml_Boost_setWeakCount_10 (a0, a1, a2, a3);
}

static void (*proc_objdetect_HOGDescriptor_save_10)(JNIEnv*,  jclass,  jlong,  jstring,  jstring);
void Java_org_opencv_objdetect_HOGDescriptor_save_10(JNIEnv* a0, jclass a1, jlong a2, jstring a3, jstring a4) {
    proc_objdetect_HOGDescriptor_save_10 (a0, a1, a2, a3, a4);
}

static void (*proc_video_BackgroundSubtractorMOG2_setComplexityReductionThreshold_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_video_BackgroundSubtractorMOG2_setComplexityReductionThreshold_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_video_BackgroundSubtractorMOG2_setComplexityReductionThreshold_10 (a0, a1, a2, a3);
}

static jboolean (*proc_calib3d_Calib3d_findChessboardCorners_10)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jlong,  jint);
jboolean Java_org_opencv_calib3d_Calib3d_findChessboardCorners_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jlong a5, jint a6) {
    return proc_calib3d_Calib3d_findChessboardCorners_10 (a0, a1, a2, a3, a4, a5, a6);
}

static jfloat (*proc_photo_MergeMertens_getExposureWeight_10)(JNIEnv*,  jclass,  jlong);
jfloat Java_org_opencv_photo_MergeMertens_getExposureWeight_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_photo_MergeMertens_getExposureWeight_10 (a0, a1, a2);
}

static jint (*proc_features2d_BOWImgDescriptorExtractor_descriptorType_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_features2d_BOWImgDescriptorExtractor_descriptorType_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_BOWImgDescriptorExtractor_descriptorType_10 (a0, a1, a2);
}

static jlong (*proc_core_Mat_n_1cross)(JNIEnv*,  jclass,  jlong,  jlong);
jlong Java_org_opencv_core_Mat_n_1cross(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_core_Mat_n_1cross (a0, a1, a2, a3);
}

static jlong (*proc_dnn_Layer_finalize_10)(JNIEnv*,  jclass,  jlong,  jlong);
jlong Java_org_opencv_dnn_Layer_finalize_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_dnn_Layer_finalize_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_accumulateWeighted_10)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jlong);
void Java_org_opencv_imgproc_Imgproc_accumulateWeighted_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jlong a5) {
    proc_imgproc_Imgproc_accumulateWeighted_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_ml_Boost_setWeightTrimRate_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_ml_Boost_setWeightTrimRate_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_ml_Boost_setWeightTrimRate_10 (a0, a1, a2, a3);
}

static void (*proc_objdetect_HOGDescriptor_save_11)(JNIEnv*,  jclass,  jlong,  jstring);
void Java_org_opencv_objdetect_HOGDescriptor_save_11(JNIEnv* a0, jclass a1, jlong a2, jstring a3) {
    proc_objdetect_HOGDescriptor_save_11 (a0, a1, a2, a3);
}

static void (*proc_video_BackgroundSubtractorMOG2_setDetectShadows_10)(JNIEnv*,  jclass,  jlong,  jboolean);
void Java_org_opencv_video_BackgroundSubtractorMOG2_setDetectShadows_10(JNIEnv* a0, jclass a1, jlong a2, jboolean a3) {
    proc_video_BackgroundSubtractorMOG2_setDetectShadows_10 (a0, a1, a2, a3);
}

static jboolean (*proc_calib3d_Calib3d_findChessboardCorners_11)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jlong);
jboolean Java_org_opencv_calib3d_Calib3d_findChessboardCorners_11(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jlong a5) {
    return proc_calib3d_Calib3d_findChessboardCorners_11 (a0, a1, a2, a3, a4, a5);
}

static jfloat (*proc_photo_MergeMertens_getSaturationWeight_10)(JNIEnv*,  jclass,  jlong);
jfloat Java_org_opencv_photo_MergeMertens_getSaturationWeight_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_photo_MergeMertens_getSaturationWeight_10 (a0, a1, a2);
}

static jlong (*proc_core_Mat_n_1dataAddr)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_core_Mat_n_1dataAddr(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_core_Mat_n_1dataAddr (a0, a1, a2);
}

static jlong (*proc_features2d_BOWImgDescriptorExtractor_getVocabulary_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_features2d_BOWImgDescriptorExtractor_getVocabulary_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_BOWImgDescriptorExtractor_getVocabulary_10 (a0, a1, a2);
}

static jlong (*proc_ml_DTrees_create_10)(JNIEnv*,  jclass);
jlong Java_org_opencv_ml_DTrees_create_10(JNIEnv* a0, jclass a1) {
    return proc_ml_DTrees_create_10 (a0, a1);
}

static void (*proc_dnn_Layer_finalize_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_dnn_Layer_finalize_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_dnn_Layer_finalize_11 (a0, a1, a2, a3, a4);
}

static void (*proc_imgproc_Imgproc_accumulateWeighted_11)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble);
void Java_org_opencv_imgproc_Imgproc_accumulateWeighted_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4) {
    proc_imgproc_Imgproc_accumulateWeighted_11 (a0, a1, a2, a3, a4);
}

static void (*proc_objdetect_HOGDescriptor_setSVMDetector_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_objdetect_HOGDescriptor_setSVMDetector_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_objdetect_HOGDescriptor_setSVMDetector_10 (a0, a1, a2, a3);
}

static void (*proc_video_BackgroundSubtractorMOG2_setHistory_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_video_BackgroundSubtractorMOG2_setHistory_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_video_BackgroundSubtractorMOG2_setHistory_10 (a0, a1, a2, a3);
}

static jboolean (*proc_calib3d_Calib3d_findCirclesGrid_10)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jlong,  jint);
jboolean Java_org_opencv_calib3d_Calib3d_findCirclesGrid_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jlong a5, jint a6) {
    return proc_calib3d_Calib3d_findCirclesGrid_10 (a0, a1, a2, a3, a4, a5, a6);
}

static jlong (*proc_objdetect_CascadeClassifier_CascadeClassifier_10)(JNIEnv*,  jclass,  jstring);
jlong Java_org_opencv_objdetect_CascadeClassifier_CascadeClassifier_10(JNIEnv* a0, jclass a1, jstring a2) {
    return proc_objdetect_CascadeClassifier_CascadeClassifier_10 (a0, a1, a2);
}

static void (*proc_core_Mat_n_1delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_core_Mat_n_1delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_core_Mat_n_1delete (a0, a1, a2);
}

static void (*proc_dnn_Layer_forward_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_dnn_Layer_forward_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_dnn_Layer_forward_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_features2d_BOWImgDescriptorExtractor_setVocabulary_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_features2d_BOWImgDescriptorExtractor_setVocabulary_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_features2d_BOWImgDescriptorExtractor_setVocabulary_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_accumulate_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_imgproc_Imgproc_accumulate_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_imgproc_Imgproc_accumulate_10 (a0, a1, a2, a3, a4);
}

static void (*proc_ml_DTrees_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_ml_DTrees_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_ml_DTrees_delete (a0, a1, a2);
}

static void (*proc_photo_MergeMertens_process_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_photo_MergeMertens_process_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6) {
    proc_photo_MergeMertens_process_10 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_video_BackgroundSubtractorMOG2_setNMixtures_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_video_BackgroundSubtractorMOG2_setNMixtures_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_video_BackgroundSubtractorMOG2_setNMixtures_10 (a0, a1, a2, a3);
}

static jboolean (*proc_calib3d_Calib3d_findCirclesGrid_11)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jlong);
jboolean Java_org_opencv_calib3d_Calib3d_findCirclesGrid_11(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jlong a5) {
    return proc_calib3d_Calib3d_findCirclesGrid_11 (a0, a1, a2, a3, a4, a5);
}

static jint (*proc_core_Mat_n_1depth)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_core_Mat_n_1depth(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_core_Mat_n_1depth (a0, a1, a2);
}

static jint (*proc_ml_DTrees_getCVFolds_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_ml_DTrees_getCVFolds_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_DTrees_getCVFolds_10 (a0, a1, a2);
}

static jlong (*proc_dnn_Layer_get_1blobs_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_dnn_Layer_get_1blobs_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_dnn_Layer_get_1blobs_10 (a0, a1, a2);
}

static jlong (*proc_features2d_BOWKMeansTrainer_BOWKMeansTrainer_10)(JNIEnv*,  jclass,  jint,  jint,  jint,  jdouble,  jint,  jint);
jlong Java_org_opencv_features2d_BOWKMeansTrainer_BOWKMeansTrainer_10(JNIEnv* a0, jclass a1, jint a2, jint a3, jint a4, jdouble a5, jint a6, jint a7) {
    return proc_features2d_BOWKMeansTrainer_BOWKMeansTrainer_10 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static jlong (*proc_objdetect_CascadeClassifier_CascadeClassifier_11)(JNIEnv*,  jclass);
jlong Java_org_opencv_objdetect_CascadeClassifier_CascadeClassifier_11(JNIEnv* a0, jclass a1) {
    return proc_objdetect_CascadeClassifier_CascadeClassifier_11 (a0, a1);
}

static void (*proc_imgproc_Imgproc_accumulate_11)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_imgproc_Imgproc_accumulate_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_imgproc_Imgproc_accumulate_11 (a0, a1, a2, a3);
}

static void (*proc_photo_MergeMertens_process_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_photo_MergeMertens_process_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_photo_MergeMertens_process_11 (a0, a1, a2, a3, a4);
}

static void (*proc_video_BackgroundSubtractorMOG2_setShadowThreshold_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_video_BackgroundSubtractorMOG2_setShadowThreshold_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_video_BackgroundSubtractorMOG2_setShadowThreshold_10 (a0, a1, a2, a3);
}

static jboolean (*proc_objdetect_CascadeClassifier_convert_10)(JNIEnv*,  jclass,  jstring,  jstring);
jboolean Java_org_opencv_objdetect_CascadeClassifier_convert_10(JNIEnv* a0, jclass a1, jstring a2, jstring a3) {
    return proc_objdetect_CascadeClassifier_convert_10 (a0, a1, a2, a3);
}

static jint (*proc_ml_DTrees_getMaxCategories_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_ml_DTrees_getMaxCategories_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_DTrees_getMaxCategories_10 (a0, a1, a2);
}

static jlong (*proc_calib3d_Calib3d_findEssentialMat_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jint,  jdouble,  jdouble,  jlong);
jlong Java_org_opencv_calib3d_Calib3d_findEssentialMat_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jint a5, jdouble a6, jdouble a7, jlong a8) {
    return proc_calib3d_Calib3d_findEssentialMat_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static jlong (*proc_core_Mat_n_1diag__JI)(JNIEnv*,  jclass,  jlong,  jint);
jlong Java_org_opencv_core_Mat_n_1diag__JI(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    return proc_core_Mat_n_1diag__JI (a0, a1, a2, a3);
}

static jlong (*proc_features2d_BOWKMeansTrainer_BOWKMeansTrainer_11)(JNIEnv*,  jclass,  jint);
jlong Java_org_opencv_features2d_BOWKMeansTrainer_BOWKMeansTrainer_11(JNIEnv* a0, jclass a1, jint a2) {
    return proc_features2d_BOWKMeansTrainer_BOWKMeansTrainer_11 (a0, a1, a2);
}

static jstring (*proc_dnn_Layer_get_1name_10)(JNIEnv*,  jclass,  jlong);
jstring Java_org_opencv_dnn_Layer_get_1name_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_dnn_Layer_get_1name_10 (a0, a1, a2);
}

static void (*proc_imgproc_Imgproc_adaptiveThreshold_10)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jint,  jint,  jint,  jdouble);
void Java_org_opencv_imgproc_Imgproc_adaptiveThreshold_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jint a5, jint a6, jint a7, jdouble a8) {
    proc_imgproc_Imgproc_adaptiveThreshold_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static void (*proc_photo_MergeMertens_setContrastWeight_10)(JNIEnv*,  jclass,  jlong,  jfloat);
void Java_org_opencv_photo_MergeMertens_setContrastWeight_10(JNIEnv* a0, jclass a1, jlong a2, jfloat a3) {
    proc_photo_MergeMertens_setContrastWeight_10 (a0, a1, a2, a3);
}

static void (*proc_video_BackgroundSubtractorMOG2_setShadowValue_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_video_BackgroundSubtractorMOG2_setShadowValue_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_video_BackgroundSubtractorMOG2_setShadowValue_10 (a0, a1, a2, a3);
}

static jint (*proc_dnn_Layer_get_1preferableTarget_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_dnn_Layer_get_1preferableTarget_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_dnn_Layer_get_1preferableTarget_10 (a0, a1, a2);
}

static jint (*proc_ml_DTrees_getMaxDepth_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_ml_DTrees_getMaxDepth_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_DTrees_getMaxDepth_10 (a0, a1, a2);
}

static jlong (*proc_calib3d_Calib3d_findEssentialMat_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jint,  jdouble,  jdouble);
jlong Java_org_opencv_calib3d_Calib3d_findEssentialMat_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jint a5, jdouble a6, jdouble a7) {
    return proc_calib3d_Calib3d_findEssentialMat_11 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static jlong (*proc_core_Mat_n_1diag__J)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_core_Mat_n_1diag__J(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_core_Mat_n_1diag__J (a0, a1, a2);
}

static jlong (*proc_features2d_BOWKMeansTrainer_cluster_10)(JNIEnv*,  jclass,  jlong,  jlong);
jlong Java_org_opencv_features2d_BOWKMeansTrainer_cluster_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_features2d_BOWKMeansTrainer_cluster_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_applyColorMap_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_imgproc_Imgproc_applyColorMap_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_imgproc_Imgproc_applyColorMap_10 (a0, a1, a2, a3, a4);
}

static void (*proc_objdetect_CascadeClassifier_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_objdetect_CascadeClassifier_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_objdetect_CascadeClassifier_delete (a0, a1, a2);
}

static void (*proc_photo_MergeMertens_setExposureWeight_10)(JNIEnv*,  jclass,  jlong,  jfloat);
void Java_org_opencv_photo_MergeMertens_setExposureWeight_10(JNIEnv* a0, jclass a1, jlong a2, jfloat a3) {
    proc_photo_MergeMertens_setExposureWeight_10 (a0, a1, a2, a3);
}

static void (*proc_video_BackgroundSubtractorMOG2_setVarInit_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_video_BackgroundSubtractorMOG2_setVarInit_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_video_BackgroundSubtractorMOG2_setVarInit_10 (a0, a1, a2, a3);
}

static jint (*proc_core_Mat_n_1dims)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_core_Mat_n_1dims(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_core_Mat_n_1dims (a0, a1, a2);
}

static jint (*proc_ml_DTrees_getMinSampleCount_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_ml_DTrees_getMinSampleCount_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_DTrees_getMinSampleCount_10 (a0, a1, a2);
}

static jlong (*proc_calib3d_Calib3d_findEssentialMat_12)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
jlong Java_org_opencv_calib3d_Calib3d_findEssentialMat_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    return proc_calib3d_Calib3d_findEssentialMat_12 (a0, a1, a2, a3, a4);
}

static jlong (*proc_features2d_BOWKMeansTrainer_cluster_11)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_features2d_BOWKMeansTrainer_cluster_11(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_BOWKMeansTrainer_cluster_11 (a0, a1, a2);
}

static jstring (*proc_dnn_Layer_get_1type_10)(JNIEnv*,  jclass,  jlong);
jstring Java_org_opencv_dnn_Layer_get_1type_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_dnn_Layer_get_1type_10 (a0, a1, a2);
}

static void (*proc_imgproc_Imgproc_applyColorMap_11)(JNIEnv*,  jclass,  jlong,  jlong,  jint);
void Java_org_opencv_imgproc_Imgproc_applyColorMap_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4) {
    proc_imgproc_Imgproc_applyColorMap_11 (a0, a1, a2, a3, a4);
}

static void (*proc_objdetect_CascadeClassifier_detectMultiScale2_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jdouble,  jint,  jint,  jdouble,  jdouble,  jdouble,  jdouble);
void Java_org_opencv_objdetect_CascadeClassifier_detectMultiScale2_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jdouble a6, jint a7, jint a8, jdouble a9, jdouble a10, jdouble a11, jdouble a12) {
    proc_objdetect_CascadeClassifier_detectMultiScale2_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12);
}

static void (*proc_photo_MergeMertens_setSaturationWeight_10)(JNIEnv*,  jclass,  jlong,  jfloat);
void Java_org_opencv_photo_MergeMertens_setSaturationWeight_10(JNIEnv* a0, jclass a1, jlong a2, jfloat a3) {
    proc_photo_MergeMertens_setSaturationWeight_10 (a0, a1, a2, a3);
}

static void (*proc_video_BackgroundSubtractorMOG2_setVarMax_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_video_BackgroundSubtractorMOG2_setVarMax_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_video_BackgroundSubtractorMOG2_setVarMax_10 (a0, a1, a2, a3);
}

static jdouble (*proc_core_Mat_n_1dot)(JNIEnv*,  jclass,  jlong,  jlong);
jdouble Java_org_opencv_core_Mat_n_1dot(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_core_Mat_n_1dot (a0, a1, a2, a3);
}

static jlong (*proc_calib3d_Calib3d_findEssentialMat_13)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jdouble,  jint,  jdouble,  jdouble,  jlong);
jlong Java_org_opencv_calib3d_Calib3d_findEssentialMat_13(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jdouble a6, jint a7, jdouble a8, jdouble a9, jlong a10) {
    return proc_calib3d_Calib3d_findEssentialMat_13 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10);
}

static jlong (*proc_ml_DTrees_getPriors_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_ml_DTrees_getPriors_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_DTrees_getPriors_10 (a0, a1, a2);
}

static void (*proc_dnn_Layer_run_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_dnn_Layer_run_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_dnn_Layer_run_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_features2d_BOWKMeansTrainer_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_features2d_BOWKMeansTrainer_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_features2d_BOWKMeansTrainer_delete (a0, a1, a2);
}

static void (*proc_imgproc_Imgproc_approxPolyDP_10)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jboolean);
void Java_org_opencv_imgproc_Imgproc_approxPolyDP_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jboolean a5) {
    proc_imgproc_Imgproc_approxPolyDP_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_objdetect_CascadeClassifier_detectMultiScale2_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_objdetect_CascadeClassifier_detectMultiScale2_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_objdetect_CascadeClassifier_detectMultiScale2_11 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_photo_MergeRobertson_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_photo_MergeRobertson_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_photo_MergeRobertson_delete (a0, a1, a2);
}

static void (*proc_video_BackgroundSubtractorMOG2_setVarMin_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_video_BackgroundSubtractorMOG2_setVarMin_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_video_BackgroundSubtractorMOG2_setVarMin_10 (a0, a1, a2, a3);
}

static jdouble (*proc_imgproc_Imgproc_arcLength_10)(JNIEnv*,  jclass,  jlong,  jboolean);
jdouble Java_org_opencv_imgproc_Imgproc_arcLength_10(JNIEnv* a0, jclass a1, jlong a2, jboolean a3) {
    return proc_imgproc_Imgproc_arcLength_10 (a0, a1, a2, a3);
}

static jfloat (*proc_ml_DTrees_getRegressionAccuracy_10)(JNIEnv*,  jclass,  jlong);
jfloat Java_org_opencv_ml_DTrees_getRegressionAccuracy_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_DTrees_getRegressionAccuracy_10 (a0, a1, a2);
}

static jlong (*proc_calib3d_Calib3d_findEssentialMat_14)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jdouble,  jint,  jdouble,  jdouble);
jlong Java_org_opencv_calib3d_Calib3d_findEssentialMat_14(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jdouble a6, jint a7, jdouble a8, jdouble a9) {
    return proc_calib3d_Calib3d_findEssentialMat_14 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static jlong (*proc_core_Mat_n_1elemSize)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_core_Mat_n_1elemSize(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_core_Mat_n_1elemSize (a0, a1, a2);
}

static void (*proc_dnn_Layer_set_1blobs_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_dnn_Layer_set_1blobs_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_dnn_Layer_set_1blobs_10 (a0, a1, a2, a3);
}

static void (*proc_features2d_BOWTrainer_add_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_features2d_BOWTrainer_add_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_features2d_BOWTrainer_add_10 (a0, a1, a2, a3);
}

static void (*proc_objdetect_CascadeClassifier_detectMultiScale3_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jdouble,  jint,  jint,  jdouble,  jdouble,  jdouble,  jdouble,  jboolean);
void Java_org_opencv_objdetect_CascadeClassifier_detectMultiScale3_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jdouble a7, jint a8, jint a9, jdouble a10, jdouble a11, jdouble a12, jdouble a13, jboolean a14) {
    proc_objdetect_CascadeClassifier_detectMultiScale3_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14);
}

static void (*proc_photo_MergeRobertson_process_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_photo_MergeRobertson_process_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6) {
    proc_photo_MergeRobertson_process_10 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_video_BackgroundSubtractorMOG2_setVarThresholdGen_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_video_BackgroundSubtractorMOG2_setVarThresholdGen_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_video_BackgroundSubtractorMOG2_setVarThresholdGen_10 (a0, a1, a2, a3);
}

static jboolean (*proc_ml_DTrees_getTruncatePrunedTree_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_ml_DTrees_getTruncatePrunedTree_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_DTrees_getTruncatePrunedTree_10 (a0, a1, a2);
}

static jlong (*proc_calib3d_Calib3d_findEssentialMat_15)(JNIEnv*,  jclass,  jlong,  jlong);
jlong Java_org_opencv_calib3d_Calib3d_findEssentialMat_15(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_calib3d_Calib3d_findEssentialMat_15 (a0, a1, a2, a3);
}

static jlong (*proc_core_Mat_n_1elemSize1)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_core_Mat_n_1elemSize1(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_core_Mat_n_1elemSize1 (a0, a1, a2);
}

static jlong (*proc_dnn_Net_Net_10)(JNIEnv*,  jclass);
jlong Java_org_opencv_dnn_Net_Net_10(JNIEnv* a0, jclass a1) {
    return proc_dnn_Net_Net_10 (a0, a1);
}

static void (*proc_features2d_BOWTrainer_clear_10)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_features2d_BOWTrainer_clear_10(JNIEnv* a0, jclass a1, jlong a2) {
    proc_features2d_BOWTrainer_clear_10 (a0, a1, a2);
}

static void (*proc_imgproc_Imgproc_arrowedLine_10)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jint,  jint,  jint,  jdouble);
void Java_org_opencv_imgproc_Imgproc_arrowedLine_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8, jdouble a9, jdouble a10, jint a11, jint a12, jint a13, jdouble a14) {
    proc_imgproc_Imgproc_arrowedLine_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14);
}

static void (*proc_objdetect_CascadeClassifier_detectMultiScale3_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_objdetect_CascadeClassifier_detectMultiScale3_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6) {
    proc_objdetect_CascadeClassifier_detectMultiScale3_11 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_photo_MergeRobertson_process_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_photo_MergeRobertson_process_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_photo_MergeRobertson_process_11 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_video_BackgroundSubtractorMOG2_setVarThreshold_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_video_BackgroundSubtractorMOG2_setVarThreshold_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_video_BackgroundSubtractorMOG2_setVarThreshold_10 (a0, a1, a2, a3);
}

static jboolean (*proc_core_Mat_n_1empty)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_core_Mat_n_1empty(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_core_Mat_n_1empty (a0, a1, a2);
}

static jboolean (*proc_ml_DTrees_getUse1SERule_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_ml_DTrees_getUse1SERule_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_DTrees_getUse1SERule_10 (a0, a1, a2);
}

static jlong (*proc_calib3d_Calib3d_findFundamentalMat_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jdouble,  jdouble,  jlong);
jlong Java_org_opencv_calib3d_Calib3d_findFundamentalMat_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jdouble a5, jdouble a6, jlong a7) {
    return proc_calib3d_Calib3d_findFundamentalMat_10 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static jlong (*proc_features2d_BOWTrainer_cluster_10)(JNIEnv*,  jclass,  jlong,  jlong);
jlong Java_org_opencv_features2d_BOWTrainer_cluster_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_features2d_BOWTrainer_cluster_10 (a0, a1, a2, a3);
}

static void (*proc_dnn_Net_connect_10)(JNIEnv*,  jclass,  jlong,  jstring,  jstring);
void Java_org_opencv_dnn_Net_connect_10(JNIEnv* a0, jclass a1, jlong a2, jstring a3, jstring a4) {
    proc_dnn_Net_connect_10 (a0, a1, a2, a3, a4);
}

static void (*proc_imgproc_Imgproc_arrowedLine_11)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_arrowedLine_11(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8, jdouble a9, jdouble a10) {
    proc_imgproc_Imgproc_arrowedLine_11 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10);
}

static void (*proc_objdetect_CascadeClassifier_detectMultiScale_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jdouble,  jint,  jint,  jdouble,  jdouble,  jdouble,  jdouble);
void Java_org_opencv_objdetect_CascadeClassifier_detectMultiScale_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jdouble a5, jint a6, jint a7, jdouble a8, jdouble a9, jdouble a10, jdouble a11) {
    proc_objdetect_CascadeClassifier_detectMultiScale_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11);
}

static void (*proc_photo_Photo_colorChange_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jfloat,  jfloat,  jfloat);
void Java_org_opencv_photo_Photo_colorChange_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jfloat a5, jfloat a6, jfloat a7) {
    proc_photo_Photo_colorChange_10 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_video_DenseOpticalFlow_calc_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_video_DenseOpticalFlow_calc_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_video_DenseOpticalFlow_calc_10 (a0, a1, a2, a3, a4, a5);
}

static jboolean (*proc_ml_DTrees_getUseSurrogates_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_ml_DTrees_getUseSurrogates_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_DTrees_getUseSurrogates_10 (a0, a1, a2);
}

static jlong (*proc_calib3d_Calib3d_findFundamentalMat_11)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jdouble,  jdouble);
jlong Java_org_opencv_calib3d_Calib3d_findFundamentalMat_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jdouble a5, jdouble a6) {
    return proc_calib3d_Calib3d_findFundamentalMat_11 (a0, a1, a2, a3, a4, a5, a6);
}

static jlong (*proc_core_Mat_n_1eye__III)(JNIEnv*,  jclass,  jint,  jint,  jint);
jlong Java_org_opencv_core_Mat_n_1eye__III(JNIEnv* a0, jclass a1, jint a2, jint a3, jint a4) {
    return proc_core_Mat_n_1eye__III (a0, a1, a2, a3, a4);
}

static jlong (*proc_features2d_BOWTrainer_cluster_11)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_features2d_BOWTrainer_cluster_11(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_BOWTrainer_cluster_11 (a0, a1, a2);
}

static void (*proc_dnn_Net_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_dnn_Net_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_dnn_Net_delete (a0, a1, a2);
}

static void (*proc_imgproc_Imgproc_bilateralFilter_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jdouble,  jdouble,  jint);
void Java_org_opencv_imgproc_Imgproc_bilateralFilter_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jdouble a5, jdouble a6, jint a7) {
    proc_imgproc_Imgproc_bilateralFilter_10 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_objdetect_CascadeClassifier_detectMultiScale_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_objdetect_CascadeClassifier_detectMultiScale_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_objdetect_CascadeClassifier_detectMultiScale_11 (a0, a1, a2, a3, a4);
}

static void (*proc_photo_Photo_colorChange_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_photo_Photo_colorChange_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_photo_Photo_colorChange_11 (a0, a1, a2, a3, a4);
}

static void (*proc_video_DenseOpticalFlow_collectGarbage_10)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_video_DenseOpticalFlow_collectGarbage_10(JNIEnv* a0, jclass a1, jlong a2) {
    proc_video_DenseOpticalFlow_collectGarbage_10 (a0, a1, a2);
}

static jboolean (*proc_objdetect_CascadeClassifier_empty_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_objdetect_CascadeClassifier_empty_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_objdetect_CascadeClassifier_empty_10 (a0, a1, a2);
}

static jlong (*proc_calib3d_Calib3d_findFundamentalMat_12)(JNIEnv*,  jclass,  jlong,  jlong);
jlong Java_org_opencv_calib3d_Calib3d_findFundamentalMat_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_calib3d_Calib3d_findFundamentalMat_12 (a0, a1, a2, a3);
}

static jlong (*proc_core_Mat_n_1eye__DDI)(JNIEnv*,  jclass,  jdouble,  jdouble,  jint);
jlong Java_org_opencv_core_Mat_n_1eye__DDI(JNIEnv* a0, jclass a1, jdouble a2, jdouble a3, jint a4) {
    return proc_core_Mat_n_1eye__DDI (a0, a1, a2, a3, a4);
}

static jlong (*proc_ml_DTrees_load_10)(JNIEnv*,  jclass,  jstring,  jstring);
jlong Java_org_opencv_ml_DTrees_load_10(JNIEnv* a0, jclass a1, jstring a2, jstring a3) {
    return proc_ml_DTrees_load_10 (a0, a1, a2, a3);
}

static jlong (*proc_photo_Photo_createAlignMTB_10)(JNIEnv*,  jclass,  jint,  jint,  jboolean);
jlong Java_org_opencv_photo_Photo_createAlignMTB_10(JNIEnv* a0, jclass a1, jint a2, jint a3, jboolean a4) {
    return proc_photo_Photo_createAlignMTB_10 (a0, a1, a2, a3, a4);
}

static void (*proc_dnn_Net_deleteLayer_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_dnn_Net_deleteLayer_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_dnn_Net_deleteLayer_10 (a0, a1, a2, a3);
}

static void (*proc_features2d_BOWTrainer_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_features2d_BOWTrainer_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_features2d_BOWTrainer_delete (a0, a1, a2);
}

static void (*proc_imgproc_Imgproc_bilateralFilter_11)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_bilateralFilter_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jdouble a5, jdouble a6) {
    proc_imgproc_Imgproc_bilateralFilter_11 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_video_DenseOpticalFlow_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_video_DenseOpticalFlow_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_video_DenseOpticalFlow_delete (a0, a1, a2);
}

static jboolean (*proc_dnn_Net_empty_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_dnn_Net_empty_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_dnn_Net_empty_10 (a0, a1, a2);
}

static jint (*proc_features2d_BOWTrainer_descriptorsCount_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_features2d_BOWTrainer_descriptorsCount_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_BOWTrainer_descriptorsCount_10 (a0, a1, a2);
}

static jint (*proc_objdetect_CascadeClassifier_getFeatureType_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_objdetect_CascadeClassifier_getFeatureType_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_objdetect_CascadeClassifier_getFeatureType_10 (a0, a1, a2);
}

static jlong (*proc_calib3d_Calib3d_findHomography_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jdouble,  jlong,  jint,  jdouble);
jlong Java_org_opencv_calib3d_Calib3d_findHomography_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jdouble a5, jlong a6, jint a7, jdouble a8) {
    return proc_calib3d_Calib3d_findHomography_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static jlong (*proc_core_Mat_n_1inv__JI)(JNIEnv*,  jclass,  jlong,  jint);
jlong Java_org_opencv_core_Mat_n_1inv__JI(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    return proc_core_Mat_n_1inv__JI (a0, a1, a2, a3);
}

static jlong (*proc_ml_DTrees_load_11)(JNIEnv*,  jclass,  jstring);
jlong Java_org_opencv_ml_DTrees_load_11(JNIEnv* a0, jclass a1, jstring a2) {
    return proc_ml_DTrees_load_11 (a0, a1, a2);
}

static jlong (*proc_photo_Photo_createAlignMTB_11)(JNIEnv*,  jclass);
jlong Java_org_opencv_photo_Photo_createAlignMTB_11(JNIEnv* a0, jclass a1) {
    return proc_photo_Photo_createAlignMTB_11 (a0, a1);
}

static jlong (*proc_video_DualTVL1OpticalFlow_create_10)(JNIEnv*,  jclass,  jdouble,  jdouble,  jdouble,  jint,  jint,  jdouble,  jint,  jint,  jdouble,  jdouble,  jint,  jboolean);
jlong Java_org_opencv_video_DualTVL1OpticalFlow_create_10(JNIEnv* a0, jclass a1, jdouble a2, jdouble a3, jdouble a4, jint a5, jint a6, jdouble a7, jint a8, jint a9, jdouble a10, jdouble a11, jint a12, jboolean a13) {
    return proc_video_DualTVL1OpticalFlow_create_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13);
}

static void (*proc_imgproc_Imgproc_blur_10)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jint);
void Java_org_opencv_imgproc_Imgproc_blur_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7, jint a8) {
    proc_imgproc_Imgproc_blur_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static jdoubleArray (*proc_objdetect_CascadeClassifier_getOriginalWindowSize_10)(JNIEnv*,  jclass,  jlong);
jdoubleArray Java_org_opencv_objdetect_CascadeClassifier_getOriginalWindowSize_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_objdetect_CascadeClassifier_getOriginalWindowSize_10 (a0, a1, a2);
}

static jlong (*proc_calib3d_Calib3d_findHomography_11)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jdouble);
jlong Java_org_opencv_calib3d_Calib3d_findHomography_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jdouble a5) {
    return proc_calib3d_Calib3d_findHomography_11 (a0, a1, a2, a3, a4, a5);
}

static jlong (*proc_core_Mat_n_1inv__J)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_core_Mat_n_1inv__J(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_core_Mat_n_1inv__J (a0, a1, a2);
}

static jlong (*proc_features2d_BOWTrainer_getDescriptors_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_features2d_BOWTrainer_getDescriptors_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_BOWTrainer_getDescriptors_10 (a0, a1, a2);
}

static jlong (*proc_photo_Photo_createCalibrateDebevec_10)(JNIEnv*,  jclass,  jint,  jfloat,  jboolean);
jlong Java_org_opencv_photo_Photo_createCalibrateDebevec_10(JNIEnv* a0, jclass a1, jint a2, jfloat a3, jboolean a4) {
    return proc_photo_Photo_createCalibrateDebevec_10 (a0, a1, a2, a3, a4);
}

static jlong (*proc_video_DualTVL1OpticalFlow_create_11)(JNIEnv*,  jclass);
jlong Java_org_opencv_video_DualTVL1OpticalFlow_create_11(JNIEnv* a0, jclass a1) {
    return proc_video_DualTVL1OpticalFlow_create_11 (a0, a1);
}

static void (*proc_dnn_Net_enableFusion_10)(JNIEnv*,  jclass,  jlong,  jboolean);
void Java_org_opencv_dnn_Net_enableFusion_10(JNIEnv* a0, jclass a1, jlong a2, jboolean a3) {
    proc_dnn_Net_enableFusion_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_blur_11)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_blur_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7) {
    proc_imgproc_Imgproc_blur_11 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_ml_DTrees_setCVFolds_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_ml_DTrees_setCVFolds_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_ml_DTrees_setCVFolds_10 (a0, a1, a2, a3);
}

static jboolean (*proc_core_Mat_n_1isContinuous)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_core_Mat_n_1isContinuous(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_core_Mat_n_1isContinuous (a0, a1, a2);
}

static jboolean (*proc_objdetect_CascadeClassifier_isOldFormatCascade_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_objdetect_CascadeClassifier_isOldFormatCascade_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_objdetect_CascadeClassifier_isOldFormatCascade_10 (a0, a1, a2);
}

static jlong (*proc_calib3d_Calib3d_findHomography_12)(JNIEnv*,  jclass,  jlong,  jlong);
jlong Java_org_opencv_calib3d_Calib3d_findHomography_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_calib3d_Calib3d_findHomography_12 (a0, a1, a2, a3);
}

static jlong (*proc_dnn_Net_forward_10)(JNIEnv*,  jclass,  jlong,  jstring);
jlong Java_org_opencv_dnn_Net_forward_10(JNIEnv* a0, jclass a1, jlong a2, jstring a3) {
    return proc_dnn_Net_forward_10 (a0, a1, a2, a3);
}

static jlong (*proc_features2d_BRISK_create_10)(JNIEnv*,  jclass,  jint,  jint,  jlong,  jlong,  jfloat,  jfloat,  jlong);
jlong Java_org_opencv_features2d_BRISK_create_10(JNIEnv* a0, jclass a1, jint a2, jint a3, jlong a4, jlong a5, jfloat a6, jfloat a7, jlong a8) {
    return proc_features2d_BRISK_create_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static jlong (*proc_photo_Photo_createCalibrateDebevec_11)(JNIEnv*,  jclass);
jlong Java_org_opencv_photo_Photo_createCalibrateDebevec_11(JNIEnv* a0, jclass a1) {
    return proc_photo_Photo_createCalibrateDebevec_11 (a0, a1);
}

static void (*proc_imgproc_Imgproc_blur_12)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_blur_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5) {
    proc_imgproc_Imgproc_blur_12 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_ml_DTrees_setMaxCategories_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_ml_DTrees_setMaxCategories_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_ml_DTrees_setMaxCategories_10 (a0, a1, a2, a3);
}

static void (*proc_video_DualTVL1OpticalFlow_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_video_DualTVL1OpticalFlow_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_video_DualTVL1OpticalFlow_delete (a0, a1, a2);
}

static jboolean (*proc_core_Mat_n_1isSubmatrix)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_core_Mat_n_1isSubmatrix(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_core_Mat_n_1isSubmatrix (a0, a1, a2);
}

static jboolean (*proc_objdetect_CascadeClassifier_load_10)(JNIEnv*,  jclass,  jlong,  jstring);
jboolean Java_org_opencv_objdetect_CascadeClassifier_load_10(JNIEnv* a0, jclass a1, jlong a2, jstring a3) {
    return proc_objdetect_CascadeClassifier_load_10 (a0, a1, a2, a3);
}

static jdouble (*proc_video_DualTVL1OpticalFlow_getEpsilon_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_video_DualTVL1OpticalFlow_getEpsilon_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_DualTVL1OpticalFlow_getEpsilon_10 (a0, a1, a2);
}

static jdoubleArray (*proc_imgproc_Imgproc_boundingRect_10)(JNIEnv*,  jclass,  jlong);
jdoubleArray Java_org_opencv_imgproc_Imgproc_boundingRect_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_imgproc_Imgproc_boundingRect_10 (a0, a1, a2);
}

static jlong (*proc_calib3d_Calib3d_getOptimalNewCameraMatrix_10)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdoubleArray,  jboolean);
jlong Java_org_opencv_calib3d_Calib3d_getOptimalNewCameraMatrix_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8, jdoubleArray a9, jboolean a10) {
    return proc_calib3d_Calib3d_getOptimalNewCameraMatrix_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10);
}

static jlong (*proc_dnn_Net_forward_11)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_dnn_Net_forward_11(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_dnn_Net_forward_11 (a0, a1, a2);
}

static jlong (*proc_features2d_BRISK_create_11)(JNIEnv*,  jclass,  jint,  jint,  jlong,  jlong);
jlong Java_org_opencv_features2d_BRISK_create_11(JNIEnv* a0, jclass a1, jint a2, jint a3, jlong a4, jlong a5) {
    return proc_features2d_BRISK_create_11 (a0, a1, a2, a3, a4, a5);
}

static jlong (*proc_photo_Photo_createCalibrateRobertson_10)(JNIEnv*,  jclass,  jint,  jfloat);
jlong Java_org_opencv_photo_Photo_createCalibrateRobertson_10(JNIEnv* a0, jclass a1, jint a2, jfloat a3) {
    return proc_photo_Photo_createCalibrateRobertson_10 (a0, a1, a2, a3);
}

static void (*proc_ml_DTrees_setMaxDepth_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_ml_DTrees_setMaxDepth_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_ml_DTrees_setMaxDepth_10 (a0, a1, a2, a3);
}

static jdouble (*proc_video_DualTVL1OpticalFlow_getGamma_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_video_DualTVL1OpticalFlow_getGamma_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_DualTVL1OpticalFlow_getGamma_10 (a0, a1, a2);
}

static jlong (*proc_calib3d_Calib3d_getOptimalNewCameraMatrix_11)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jdouble);
jlong Java_org_opencv_calib3d_Calib3d_getOptimalNewCameraMatrix_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jdouble a6) {
    return proc_calib3d_Calib3d_getOptimalNewCameraMatrix_11 (a0, a1, a2, a3, a4, a5, a6);
}

static jlong (*proc_core_Mat_n_1mul__JJD)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble);
jlong Java_org_opencv_core_Mat_n_1mul__JJD(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4) {
    return proc_core_Mat_n_1mul__JJD (a0, a1, a2, a3, a4);
}

static jlong (*proc_features2d_BRISK_create_12)(JNIEnv*,  jclass,  jint,  jint,  jfloat);
jlong Java_org_opencv_features2d_BRISK_create_12(JNIEnv* a0, jclass a1, jint a2, jint a3, jfloat a4) {
    return proc_features2d_BRISK_create_12 (a0, a1, a2, a3, a4);
}

static jlong (*proc_photo_Photo_createCalibrateRobertson_11)(JNIEnv*,  jclass);
jlong Java_org_opencv_photo_Photo_createCalibrateRobertson_11(JNIEnv* a0, jclass a1) {
    return proc_photo_Photo_createCalibrateRobertson_11 (a0, a1);
}

static void (*proc_dnn_Net_forward_12)(JNIEnv*,  jclass,  jlong,  jlong,  jstring);
void Java_org_opencv_dnn_Net_forward_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jstring a4) {
    proc_dnn_Net_forward_12 (a0, a1, a2, a3, a4);
}

static void (*proc_imgproc_Imgproc_boxFilter_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jdouble,  jdouble,  jdouble,  jdouble,  jboolean,  jint);
void Java_org_opencv_imgproc_Imgproc_boxFilter_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8, jboolean a9, jint a10) {
    proc_imgproc_Imgproc_boxFilter_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10);
}

static void (*proc_ml_DTrees_setMinSampleCount_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_ml_DTrees_setMinSampleCount_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_ml_DTrees_setMinSampleCount_10 (a0, a1, a2, a3);
}

static void (*proc_objdetect_BaseCascadeClassifier_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_objdetect_BaseCascadeClassifier_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_objdetect_BaseCascadeClassifier_delete (a0, a1, a2);
}

static jdoubleArray (*proc_calib3d_Calib3d_getValidDisparityROI_10)(JNIEnv*,  jclass,  jint,  jint,  jint,  jint,  jint,  jint,  jint,  jint,  jint,  jint,  jint);
jdoubleArray Java_org_opencv_calib3d_Calib3d_getValidDisparityROI_10(JNIEnv* a0, jclass a1, jint a2, jint a3, jint a4, jint a5, jint a6, jint a7, jint a8, jint a9, jint a10, jint a11, jint a12) {
    return proc_calib3d_Calib3d_getValidDisparityROI_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12);
}

static jint (*proc_video_DualTVL1OpticalFlow_getInnerIterations_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_video_DualTVL1OpticalFlow_getInnerIterations_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_DualTVL1OpticalFlow_getInnerIterations_10 (a0, a1, a2);
}

static jlong (*proc_core_Mat_n_1mul__JJ)(JNIEnv*,  jclass,  jlong,  jlong);
jlong Java_org_opencv_core_Mat_n_1mul__JJ(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_core_Mat_n_1mul__JJ (a0, a1, a2, a3);
}

static jlong (*proc_features2d_BRISK_create_13)(JNIEnv*,  jclass);
jlong Java_org_opencv_features2d_BRISK_create_13(JNIEnv* a0, jclass a1) {
    return proc_features2d_BRISK_create_13 (a0, a1);
}

static jlong (*proc_photo_Photo_createMergeDebevec_10)(JNIEnv*,  jclass);
jlong Java_org_opencv_photo_Photo_createMergeDebevec_10(JNIEnv* a0, jclass a1) {
    return proc_photo_Photo_createMergeDebevec_10 (a0, a1);
}

static void (*proc_dnn_Net_forward_13)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_dnn_Net_forward_13(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_dnn_Net_forward_13 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_boxFilter_11)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jdouble,  jdouble,  jdouble,  jdouble,  jboolean);
void Java_org_opencv_imgproc_Imgproc_boxFilter_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8, jboolean a9) {
    proc_imgproc_Imgproc_boxFilter_11 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static void (*proc_ml_DTrees_setPriors_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_ml_DTrees_setPriors_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_ml_DTrees_setPriors_10 (a0, a1, a2, a3);
}

static jdouble (*proc_video_DualTVL1OpticalFlow_getLambda_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_video_DualTVL1OpticalFlow_getLambda_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_DualTVL1OpticalFlow_getLambda_10 (a0, a1, a2);
}

static jlong (*proc_calib3d_Calib3d_initCameraMatrix2D_10)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jdouble);
jlong Java_org_opencv_calib3d_Calib3d_initCameraMatrix2D_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jdouble a6) {
    return proc_calib3d_Calib3d_initCameraMatrix2D_10 (a0, a1, a2, a3, a4, a5, a6);
}

static jlong (*proc_core_Mat_n_1ones__III)(JNIEnv*,  jclass,  jint,  jint,  jint);
jlong Java_org_opencv_core_Mat_n_1ones__III(JNIEnv* a0, jclass a1, jint a2, jint a3, jint a4) {
    return proc_core_Mat_n_1ones__III (a0, a1, a2, a3, a4);
}

static jlong (*proc_features2d_BRISK_create_14)(JNIEnv*,  jclass,  jlong,  jlong,  jfloat,  jfloat,  jlong);
jlong Java_org_opencv_features2d_BRISK_create_14(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jfloat a4, jfloat a5, jlong a6) {
    return proc_features2d_BRISK_create_14 (a0, a1, a2, a3, a4, a5, a6);
}

static jlong (*proc_photo_Photo_createMergeMertens_10)(JNIEnv*,  jclass,  jfloat,  jfloat,  jfloat);
jlong Java_org_opencv_photo_Photo_createMergeMertens_10(JNIEnv* a0, jclass a1, jfloat a2, jfloat a3, jfloat a4) {
    return proc_photo_Photo_createMergeMertens_10 (a0, a1, a2, a3, a4);
}

static void (*proc_dnn_Net_forward_14)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_dnn_Net_forward_14(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_dnn_Net_forward_14 (a0, a1, a2, a3, a4);
}

static void (*proc_imgproc_Imgproc_boxFilter_12)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_boxFilter_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jdouble a5, jdouble a6) {
    proc_imgproc_Imgproc_boxFilter_12 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_ml_DTrees_setRegressionAccuracy_10)(JNIEnv*,  jclass,  jlong,  jfloat);
void Java_org_opencv_ml_DTrees_setRegressionAccuracy_10(JNIEnv* a0, jclass a1, jlong a2, jfloat a3) {
    proc_ml_DTrees_setRegressionAccuracy_10 (a0, a1, a2, a3);
}

static jint (*proc_video_DualTVL1OpticalFlow_getMedianFiltering_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_video_DualTVL1OpticalFlow_getMedianFiltering_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_DualTVL1OpticalFlow_getMedianFiltering_10 (a0, a1, a2);
}

static jlong (*proc_calib3d_Calib3d_initCameraMatrix2D_11)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble);
jlong Java_org_opencv_calib3d_Calib3d_initCameraMatrix2D_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5) {
    return proc_calib3d_Calib3d_initCameraMatrix2D_11 (a0, a1, a2, a3, a4, a5);
}

static jlong (*proc_core_Mat_n_1ones__DDI)(JNIEnv*,  jclass,  jdouble,  jdouble,  jint);
jlong Java_org_opencv_core_Mat_n_1ones__DDI(JNIEnv* a0, jclass a1, jdouble a2, jdouble a3, jint a4) {
    return proc_core_Mat_n_1ones__DDI (a0, a1, a2, a3, a4);
}

static jlong (*proc_dnn_Net_getFLOPS_10)(JNIEnv*,  jclass,  jlong,  jlong);
jlong Java_org_opencv_dnn_Net_getFLOPS_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_dnn_Net_getFLOPS_10 (a0, a1, a2, a3);
}

static jlong (*proc_features2d_BRISK_create_15)(JNIEnv*,  jclass,  jlong,  jlong);
jlong Java_org_opencv_features2d_BRISK_create_15(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_features2d_BRISK_create_15 (a0, a1, a2, a3);
}

static jlong (*proc_photo_Photo_createMergeMertens_11)(JNIEnv*,  jclass);
jlong Java_org_opencv_photo_Photo_createMergeMertens_11(JNIEnv* a0, jclass a1) {
    return proc_photo_Photo_createMergeMertens_11 (a0, a1);
}

static void (*proc_imgproc_Imgproc_boxPoints_10)(JNIEnv*,  jclass,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jlong);
void Java_org_opencv_imgproc_Imgproc_boxPoints_10(JNIEnv* a0, jclass a1, jdouble a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jlong a7) {
    proc_imgproc_Imgproc_boxPoints_10 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_ml_DTrees_setTruncatePrunedTree_10)(JNIEnv*,  jclass,  jlong,  jboolean);
void Java_org_opencv_ml_DTrees_setTruncatePrunedTree_10(JNIEnv* a0, jclass a1, jlong a2, jboolean a3) {
    proc_ml_DTrees_setTruncatePrunedTree_10 (a0, a1, a2, a3);
}

static jint (*proc_video_DualTVL1OpticalFlow_getOuterIterations_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_video_DualTVL1OpticalFlow_getOuterIterations_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_DualTVL1OpticalFlow_getOuterIterations_10 (a0, a1, a2);
}

static jlong (*proc_dnn_Net_getFLOPS_11)(JNIEnv*,  jclass,  jlong,  jint,  jlong);
jlong Java_org_opencv_dnn_Net_getFLOPS_11(JNIEnv* a0, jclass a1, jlong a2, jint a3, jlong a4) {
    return proc_dnn_Net_getFLOPS_11 (a0, a1, a2, a3, a4);
}

static jlong (*proc_photo_Photo_createMergeRobertson_10)(JNIEnv*,  jclass);
jlong Java_org_opencv_photo_Photo_createMergeRobertson_10(JNIEnv* a0, jclass a1) {
    return proc_photo_Photo_createMergeRobertson_10 (a0, a1);
}

static void (*proc_calib3d_Calib3d_initUndistortRectifyMap_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jdouble,  jdouble,  jint,  jlong,  jlong);
void Java_org_opencv_calib3d_Calib3d_initUndistortRectifyMap_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jdouble a6, jdouble a7, jint a8, jlong a9, jlong a10) {
    proc_calib3d_Calib3d_initUndistortRectifyMap_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10);
}

static void (*proc_core_Mat_n_1push_1back)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_core_Mat_n_1push_1back(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_core_Mat_n_1push_1back (a0, a1, a2, a3);
}

static void (*proc_features2d_BRISK_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_features2d_BRISK_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_features2d_BRISK_delete (a0, a1, a2);
}

static void (*proc_imgproc_Imgproc_calcBackProject_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jdouble);
void Java_org_opencv_imgproc_Imgproc_calcBackProject_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jdouble a7) {
    proc_imgproc_Imgproc_calcBackProject_10 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_ml_DTrees_setUse1SERule_10)(JNIEnv*,  jclass,  jlong,  jboolean);
void Java_org_opencv_ml_DTrees_setUse1SERule_10(JNIEnv* a0, jclass a1, jlong a2, jboolean a3) {
    proc_ml_DTrees_setUse1SERule_10 (a0, a1, a2, a3);
}

static jdouble (*proc_video_DualTVL1OpticalFlow_getScaleStep_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_video_DualTVL1OpticalFlow_getScaleStep_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_DualTVL1OpticalFlow_getScaleStep_10 (a0, a1, a2);
}

static jlong (*proc_dnn_Net_getFLOPS_12)(JNIEnv*,  jclass,  jlong,  jint,  jlong);
jlong Java_org_opencv_dnn_Net_getFLOPS_12(JNIEnv* a0, jclass a1, jlong a2, jint a3, jlong a4) {
    return proc_dnn_Net_getFLOPS_12 (a0, a1, a2, a3, a4);
}

static jlong (*proc_photo_Photo_createTonemapDrago_10)(JNIEnv*,  jclass,  jfloat,  jfloat,  jfloat);
jlong Java_org_opencv_photo_Photo_createTonemapDrago_10(JNIEnv* a0, jclass a1, jfloat a2, jfloat a3, jfloat a4) {
    return proc_photo_Photo_createTonemapDrago_10 (a0, a1, a2, a3, a4);
}

static jstring (*proc_features2d_BRISK_getDefaultName_10)(JNIEnv*,  jclass,  jlong);
jstring Java_org_opencv_features2d_BRISK_getDefaultName_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_BRISK_getDefaultName_10 (a0, a1, a2);
}

static void (*proc_calib3d_Calib3d_matMulDeriv_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_calib3d_Calib3d_matMulDeriv_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_calib3d_Calib3d_matMulDeriv_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_core_Mat_n_1release)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_core_Mat_n_1release(JNIEnv* a0, jclass a1, jlong a2) {
    proc_core_Mat_n_1release (a0, a1, a2);
}

static void (*proc_imgproc_Imgproc_calcHist_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jboolean);
void Java_org_opencv_imgproc_Imgproc_calcHist_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7, jboolean a8) {
    proc_imgproc_Imgproc_calcHist_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static void (*proc_ml_DTrees_setUseSurrogates_10)(JNIEnv*,  jclass,  jlong,  jboolean);
void Java_org_opencv_ml_DTrees_setUseSurrogates_10(JNIEnv* a0, jclass a1, jlong a2, jboolean a3) {
    proc_ml_DTrees_setUseSurrogates_10 (a0, a1, a2, a3);
}

static jint (*proc_video_DualTVL1OpticalFlow_getScalesNumber_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_video_DualTVL1OpticalFlow_getScalesNumber_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_DualTVL1OpticalFlow_getScalesNumber_10 (a0, a1, a2);
}

static jlong (*proc_core_Mat_n_1reshape__JII)(JNIEnv*,  jclass,  jlong,  jint,  jint);
jlong Java_org_opencv_core_Mat_n_1reshape__JII(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4) {
    return proc_core_Mat_n_1reshape__JII (a0, a1, a2, a3, a4);
}

static jlong (*proc_dnn_Net_getFLOPS_13)(JNIEnv*,  jclass,  jlong,  jlong);
jlong Java_org_opencv_dnn_Net_getFLOPS_13(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_dnn_Net_getFLOPS_13 (a0, a1, a2, a3);
}

static jlong (*proc_ml_EM_create_10)(JNIEnv*,  jclass);
jlong Java_org_opencv_ml_EM_create_10(JNIEnv* a0, jclass a1) {
    return proc_ml_EM_create_10 (a0, a1);
}

static jlong (*proc_photo_Photo_createTonemapDrago_11)(JNIEnv*,  jclass);
jlong Java_org_opencv_photo_Photo_createTonemapDrago_11(JNIEnv* a0, jclass a1) {
    return proc_photo_Photo_createTonemapDrago_11 (a0, a1);
}

static void (*proc_calib3d_Calib3d_projectPoints_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jdouble);
void Java_org_opencv_calib3d_Calib3d_projectPoints_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7, jlong a8, jdouble a9) {
    proc_calib3d_Calib3d_projectPoints_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static void (*proc_features2d_DescriptorExtractor_compute_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_features2d_DescriptorExtractor_compute_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_features2d_DescriptorExtractor_compute_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_imgproc_Imgproc_calcHist_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_imgproc_Imgproc_calcHist_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7) {
    proc_imgproc_Imgproc_calcHist_11 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static jdouble (*proc_video_DualTVL1OpticalFlow_getTau_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_video_DualTVL1OpticalFlow_getTau_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_DualTVL1OpticalFlow_getTau_10 (a0, a1, a2);
}

static jint (*proc_dnn_Net_getLayerId_10)(JNIEnv*,  jclass,  jlong,  jstring);
jint Java_org_opencv_dnn_Net_getLayerId_10(JNIEnv* a0, jclass a1, jlong a2, jstring a3) {
    return proc_dnn_Net_getLayerId_10 (a0, a1, a2, a3);
}

static jlong (*proc_core_Mat_n_1reshape__JI)(JNIEnv*,  jclass,  jlong,  jint);
jlong Java_org_opencv_core_Mat_n_1reshape__JI(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    return proc_core_Mat_n_1reshape__JI (a0, a1, a2, a3);
}

static jlong (*proc_photo_Photo_createTonemapDurand_10)(JNIEnv*,  jclass,  jfloat,  jfloat,  jfloat,  jfloat,  jfloat);
jlong Java_org_opencv_photo_Photo_createTonemapDurand_10(JNIEnv* a0, jclass a1, jfloat a2, jfloat a3, jfloat a4, jfloat a5, jfloat a6) {
    return proc_photo_Photo_createTonemapDurand_10 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_calib3d_Calib3d_projectPoints_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_calib3d_Calib3d_projectPoints_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7) {
    proc_calib3d_Calib3d_projectPoints_11 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_features2d_DescriptorExtractor_compute_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_features2d_DescriptorExtractor_compute_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_features2d_DescriptorExtractor_compute_11 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_imgproc_Imgproc_circle_10)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jint,  jdouble,  jdouble,  jdouble,  jdouble,  jint,  jint,  jint);
void Java_org_opencv_imgproc_Imgproc_circle_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jint a5, jdouble a6, jdouble a7, jdouble a8, jdouble a9, jint a10, jint a11, jint a12) {
    proc_imgproc_Imgproc_circle_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12);
}

static void (*proc_ml_EM_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_ml_EM_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_ml_EM_delete (a0, a1, a2);
}

static jdouble (*proc_video_DualTVL1OpticalFlow_getTheta_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_video_DualTVL1OpticalFlow_getTheta_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_DualTVL1OpticalFlow_getTheta_10 (a0, a1, a2);
}

static jint (*proc_ml_EM_getClustersNumber_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_ml_EM_getClustersNumber_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_EM_getClustersNumber_10 (a0, a1, a2);
}

static jlong (*proc_core_Mat_n_1row)(JNIEnv*,  jclass,  jlong,  jint);
jlong Java_org_opencv_core_Mat_n_1row(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    return proc_core_Mat_n_1row (a0, a1, a2, a3);
}

static jlong (*proc_dnn_Net_getLayerNames_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_dnn_Net_getLayerNames_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_dnn_Net_getLayerNames_10 (a0, a1, a2);
}

static jlong (*proc_features2d_DescriptorExtractor_create_10)(JNIEnv*,  jclass,  jint);
jlong Java_org_opencv_features2d_DescriptorExtractor_create_10(JNIEnv* a0, jclass a1, jint a2) {
    return proc_features2d_DescriptorExtractor_create_10 (a0, a1, a2);
}

static jlong (*proc_photo_Photo_createTonemapDurand_11)(JNIEnv*,  jclass);
jlong Java_org_opencv_photo_Photo_createTonemapDurand_11(JNIEnv* a0, jclass a1) {
    return proc_photo_Photo_createTonemapDurand_11 (a0, a1);
}

static void (*proc_calib3d_Calib3d_projectPoints_12)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jdouble,  jlong);
void Java_org_opencv_calib3d_Calib3d_projectPoints_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7, jdouble a8, jlong a9) {
    proc_calib3d_Calib3d_projectPoints_12 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static void (*proc_imgproc_Imgproc_circle_11)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jint,  jdouble,  jdouble,  jdouble,  jdouble,  jint);
void Java_org_opencv_imgproc_Imgproc_circle_11(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jint a5, jdouble a6, jdouble a7, jdouble a8, jdouble a9, jint a10) {
    proc_imgproc_Imgproc_circle_11 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10);
}

static jboolean (*proc_video_DualTVL1OpticalFlow_getUseInitialFlow_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_video_DualTVL1OpticalFlow_getUseInitialFlow_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_DualTVL1OpticalFlow_getUseInitialFlow_10 (a0, a1, a2);
}

static jint (*proc_ml_EM_getCovarianceMatrixType_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_ml_EM_getCovarianceMatrixType_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_EM_getCovarianceMatrixType_10 (a0, a1, a2);
}

static jlong (*proc_core_Mat_n_1rowRange)(JNIEnv*,  jclass,  jlong,  jint,  jint);
jlong Java_org_opencv_core_Mat_n_1rowRange(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4) {
    return proc_core_Mat_n_1rowRange (a0, a1, a2, a3, a4);
}

static jlong (*proc_photo_Photo_createTonemapMantiuk_10)(JNIEnv*,  jclass,  jfloat,  jfloat,  jfloat);
jlong Java_org_opencv_photo_Photo_createTonemapMantiuk_10(JNIEnv* a0, jclass a1, jfloat a2, jfloat a3, jfloat a4) {
    return proc_photo_Photo_createTonemapMantiuk_10 (a0, a1, a2, a3, a4);
}

static void (*proc_calib3d_Calib3d_projectPoints_13)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_calib3d_Calib3d_projectPoints_13(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7) {
    proc_calib3d_Calib3d_projectPoints_13 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_dnn_Net_getLayerTypes_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_dnn_Net_getLayerTypes_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_dnn_Net_getLayerTypes_10 (a0, a1, a2, a3);
}

static void (*proc_features2d_DescriptorExtractor_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_features2d_DescriptorExtractor_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_features2d_DescriptorExtractor_delete (a0, a1, a2);
}

static void (*proc_imgproc_Imgproc_circle_12)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jint,  jdouble,  jdouble,  jdouble,  jdouble);
void Java_org_opencv_imgproc_Imgproc_circle_12(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jint a5, jdouble a6, jdouble a7, jdouble a8, jdouble a9) {
    proc_imgproc_Imgproc_circle_12 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static jdoubleArray (*proc_calib3d_Calib3d_RQDecomp3x3_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong);
jdoubleArray Java_org_opencv_calib3d_Calib3d_RQDecomp3x3_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7) {
    return proc_calib3d_Calib3d_RQDecomp3x3_10 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static jstring (*proc_core_Mat_nDump)(JNIEnv*,  jclass,  jlong);
jstring Java_org_opencv_core_Mat_nDump(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_core_Mat_nDump (a0, a1, a2);
}

static jboolean (*proc_imgproc_Imgproc_clipLine_10)(JNIEnv*,  jclass,  jint,  jint,  jint,  jint,  jdouble,  jdouble,  jdoubleArray,  jdouble,  jdouble,  jdoubleArray);
jboolean Java_org_opencv_imgproc_Imgproc_clipLine_10(JNIEnv* a0, jclass a1, jint a2, jint a3, jint a4, jint a5, jdouble a6, jdouble a7, jdoubleArray a8, jdouble a9, jdouble a10, jdoubleArray a11) {
    return proc_imgproc_Imgproc_clipLine_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11);
}

static jint (*proc_calib3d_Calib3d_recoverPose_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jdouble,  jdouble,  jdouble,  jlong);
jint Java_org_opencv_calib3d_Calib3d_recoverPose_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jdouble a7, jdouble a8, jdouble a9, jlong a10) {
    return proc_calib3d_Calib3d_recoverPose_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10);
}

static jint (*proc_core_Mat_n_1rows)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_core_Mat_n_1rows(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_core_Mat_n_1rows (a0, a1, a2);
}

static jint (*proc_features2d_DescriptorExtractor_descriptorSize_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_features2d_DescriptorExtractor_descriptorSize_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_DescriptorExtractor_descriptorSize_10 (a0, a1, a2);
}

static jint (*proc_video_DualTVL1OpticalFlow_getWarpingsNumber_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_video_DualTVL1OpticalFlow_getWarpingsNumber_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_DualTVL1OpticalFlow_getWarpingsNumber_10 (a0, a1, a2);
}

static jlong (*proc_dnn_Net_getLayer_10)(JNIEnv*,  jclass,  jlong,  jlong);
jlong Java_org_opencv_dnn_Net_getLayer_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_dnn_Net_getLayer_10 (a0, a1, a2, a3);
}

static jlong (*proc_photo_Photo_createTonemapMantiuk_11)(JNIEnv*,  jclass);
jlong Java_org_opencv_photo_Photo_createTonemapMantiuk_11(JNIEnv* a0, jclass a1) {
    return proc_photo_Photo_createTonemapMantiuk_11 (a0, a1);
}

static void (*proc_ml_EM_getCovs_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_ml_EM_getCovs_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_ml_EM_getCovs_10 (a0, a1, a2, a3);
}

static jdouble (*proc_imgproc_Imgproc_compareHist_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint);
jdouble Java_org_opencv_imgproc_Imgproc_compareHist_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4) {
    return proc_imgproc_Imgproc_compareHist_10 (a0, a1, a2, a3, a4);
}

static jint (*proc_calib3d_Calib3d_recoverPose_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jdouble,  jdouble,  jdouble);
jint Java_org_opencv_calib3d_Calib3d_recoverPose_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jdouble a7, jdouble a8, jdouble a9) {
    return proc_calib3d_Calib3d_recoverPose_11 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static jint (*proc_dnn_Net_getLayersCount_10)(JNIEnv*,  jclass,  jlong,  jstring);
jint Java_org_opencv_dnn_Net_getLayersCount_10(JNIEnv* a0, jclass a1, jlong a2, jstring a3) {
    return proc_dnn_Net_getLayersCount_10 (a0, a1, a2, a3);
}

static jint (*proc_features2d_DescriptorExtractor_descriptorType_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_features2d_DescriptorExtractor_descriptorType_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_DescriptorExtractor_descriptorType_10 (a0, a1, a2);
}

static jlong (*proc_core_Mat_n_1setTo__JDDDD)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble);
jlong Java_org_opencv_core_Mat_n_1setTo__JDDDD(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6) {
    return proc_core_Mat_n_1setTo__JDDDD (a0, a1, a2, a3, a4, a5, a6);
}

static jlong (*proc_ml_EM_getMeans_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_ml_EM_getMeans_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_EM_getMeans_10 (a0, a1, a2);
}

static jlong (*proc_photo_Photo_createTonemapReinhard_10)(JNIEnv*,  jclass,  jfloat,  jfloat,  jfloat,  jfloat);
jlong Java_org_opencv_photo_Photo_createTonemapReinhard_10(JNIEnv* a0, jclass a1, jfloat a2, jfloat a3, jfloat a4, jfloat a5) {
    return proc_photo_Photo_createTonemapReinhard_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_video_DualTVL1OpticalFlow_setEpsilon_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_video_DualTVL1OpticalFlow_setEpsilon_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_video_DualTVL1OpticalFlow_setEpsilon_10 (a0, a1, a2, a3);
}

static jboolean (*proc_features2d_DescriptorExtractor_empty_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_features2d_DescriptorExtractor_empty_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_DescriptorExtractor_empty_10 (a0, a1, a2);
}

static jdoubleArray (*proc_ml_EM_getTermCriteria_10)(JNIEnv*,  jclass,  jlong);
jdoubleArray Java_org_opencv_ml_EM_getTermCriteria_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_EM_getTermCriteria_10 (a0, a1, a2);
}

static jint (*proc_calib3d_Calib3d_recoverPose_12)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong);
jint Java_org_opencv_calib3d_Calib3d_recoverPose_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6) {
    return proc_calib3d_Calib3d_recoverPose_12 (a0, a1, a2, a3, a4, a5, a6);
}

static jint (*proc_imgproc_Imgproc_connectedComponentsWithAlgorithm_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint,  jint);
jint Java_org_opencv_imgproc_Imgproc_connectedComponentsWithAlgorithm_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5, jint a6) {
    return proc_imgproc_Imgproc_connectedComponentsWithAlgorithm_10 (a0, a1, a2, a3, a4, a5, a6);
}

static jlong (*proc_core_Mat_n_1setTo__JDDDDJ)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jlong);
jlong Java_org_opencv_core_Mat_n_1setTo__JDDDDJ(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jlong a7) {
    return proc_core_Mat_n_1setTo__JDDDDJ (a0, a1, a2, a3, a4, a5, a6, a7);
}

static jlong (*proc_photo_Photo_createTonemapReinhard_11)(JNIEnv*,  jclass);
jlong Java_org_opencv_photo_Photo_createTonemapReinhard_11(JNIEnv* a0, jclass a1) {
    return proc_photo_Photo_createTonemapReinhard_11 (a0, a1);
}

static void (*proc_dnn_Net_getMemoryConsumption_10)(JNIEnv*,  jclass,  jlong,  jlong,  jdoubleArray,  jdoubleArray);
void Java_org_opencv_dnn_Net_getMemoryConsumption_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdoubleArray a4, jdoubleArray a5) {
    proc_dnn_Net_getMemoryConsumption_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_video_DualTVL1OpticalFlow_setGamma_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_video_DualTVL1OpticalFlow_setGamma_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_video_DualTVL1OpticalFlow_setGamma_10 (a0, a1, a2, a3);
}

static jint (*proc_calib3d_Calib3d_recoverPose_13)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong);
jint Java_org_opencv_calib3d_Calib3d_recoverPose_13(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7, jlong a8) {
    return proc_calib3d_Calib3d_recoverPose_13 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static jint (*proc_imgproc_Imgproc_connectedComponentsWithStatsWithAlgorithm_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jint,  jint,  jint);
jint Java_org_opencv_imgproc_Imgproc_connectedComponentsWithStatsWithAlgorithm_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jint a6, jint a7, jint a8) {
    return proc_imgproc_Imgproc_connectedComponentsWithStatsWithAlgorithm_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static jlong (*proc_core_Mat_n_1setTo__JJJ)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
jlong Java_org_opencv_core_Mat_n_1setTo__JJJ(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    return proc_core_Mat_n_1setTo__JJJ (a0, a1, a2, a3, a4);
}

static jlong (*proc_ml_EM_getWeights_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_ml_EM_getWeights_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_EM_getWeights_10 (a0, a1, a2);
}

static jlong (*proc_photo_Photo_createTonemap_10)(JNIEnv*,  jclass,  jfloat);
jlong Java_org_opencv_photo_Photo_createTonemap_10(JNIEnv* a0, jclass a1, jfloat a2) {
    return proc_photo_Photo_createTonemap_10 (a0, a1, a2);
}

static void (*proc_dnn_Net_getMemoryConsumption_11)(JNIEnv*,  jclass,  jlong,  jint,  jlong,  jdoubleArray,  jdoubleArray);
void Java_org_opencv_dnn_Net_getMemoryConsumption_11(JNIEnv* a0, jclass a1, jlong a2, jint a3, jlong a4, jdoubleArray a5, jdoubleArray a6) {
    proc_dnn_Net_getMemoryConsumption_11 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_features2d_DescriptorExtractor_read_10)(JNIEnv*,  jclass,  jlong,  jstring);
void Java_org_opencv_features2d_DescriptorExtractor_read_10(JNIEnv* a0, jclass a1, jlong a2, jstring a3) {
    proc_features2d_DescriptorExtractor_read_10 (a0, a1, a2, a3);
}

static void (*proc_video_DualTVL1OpticalFlow_setInnerIterations_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_video_DualTVL1OpticalFlow_setInnerIterations_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_video_DualTVL1OpticalFlow_setInnerIterations_10 (a0, a1, a2, a3);
}

static jint (*proc_calib3d_Calib3d_recoverPose_14)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong);
jint Java_org_opencv_calib3d_Calib3d_recoverPose_14(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7) {
    return proc_calib3d_Calib3d_recoverPose_14 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static jint (*proc_imgproc_Imgproc_connectedComponentsWithStats_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jint,  jint);
jint Java_org_opencv_imgproc_Imgproc_connectedComponentsWithStats_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jint a6, jint a7) {
    return proc_imgproc_Imgproc_connectedComponentsWithStats_10 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static jlong (*proc_core_Mat_n_1setTo__JJ)(JNIEnv*,  jclass,  jlong,  jlong);
jlong Java_org_opencv_core_Mat_n_1setTo__JJ(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_core_Mat_n_1setTo__JJ (a0, a1, a2, a3);
}

static jlong (*proc_ml_EM_load_10)(JNIEnv*,  jclass,  jstring,  jstring);
jlong Java_org_opencv_ml_EM_load_10(JNIEnv* a0, jclass a1, jstring a2, jstring a3) {
    return proc_ml_EM_load_10 (a0, a1, a2, a3);
}

static jlong (*proc_photo_Photo_createTonemap_11)(JNIEnv*,  jclass);
jlong Java_org_opencv_photo_Photo_createTonemap_11(JNIEnv* a0, jclass a1) {
    return proc_photo_Photo_createTonemap_11 (a0, a1);
}

static void (*proc_dnn_Net_getMemoryConsumption_12)(JNIEnv*,  jclass,  jlong,  jint,  jlong,  jdoubleArray,  jdoubleArray);
void Java_org_opencv_dnn_Net_getMemoryConsumption_12(JNIEnv* a0, jclass a1, jlong a2, jint a3, jlong a4, jdoubleArray a5, jdoubleArray a6) {
    proc_dnn_Net_getMemoryConsumption_12 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_features2d_DescriptorExtractor_write_10)(JNIEnv*,  jclass,  jlong,  jstring);
void Java_org_opencv_features2d_DescriptorExtractor_write_10(JNIEnv* a0, jclass a1, jlong a2, jstring a3) {
    proc_features2d_DescriptorExtractor_write_10 (a0, a1, a2, a3);
}

static void (*proc_video_DualTVL1OpticalFlow_setLambda_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_video_DualTVL1OpticalFlow_setLambda_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_video_DualTVL1OpticalFlow_setLambda_10 (a0, a1, a2, a3);
}

static jdoubleArray (*proc_core_Mat_n_1size)(JNIEnv*,  jclass,  jlong);
jdoubleArray Java_org_opencv_core_Mat_n_1size(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_core_Mat_n_1size (a0, a1, a2);
}

static jint (*proc_calib3d_Calib3d_recoverPose_15)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jdouble,  jlong,  jlong);
jint Java_org_opencv_calib3d_Calib3d_recoverPose_15(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7, jdouble a8, jlong a9, jlong a10) {
    return proc_calib3d_Calib3d_recoverPose_15 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10);
}

static jint (*proc_imgproc_Imgproc_connectedComponentsWithStats_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
jint Java_org_opencv_imgproc_Imgproc_connectedComponentsWithStats_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    return proc_imgproc_Imgproc_connectedComponentsWithStats_11 (a0, a1, a2, a3, a4, a5);
}

static jlong (*proc_dnn_Net_getParam_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint);
jlong Java_org_opencv_dnn_Net_getParam_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4) {
    return proc_dnn_Net_getParam_10 (a0, a1, a2, a3, a4);
}

static jlong (*proc_ml_EM_load_11)(JNIEnv*,  jclass,  jstring);
jlong Java_org_opencv_ml_EM_load_11(JNIEnv* a0, jclass a1, jstring a2) {
    return proc_ml_EM_load_11 (a0, a1, a2);
}

static void (*proc_features2d_DescriptorMatcher_add_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_features2d_DescriptorMatcher_add_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_features2d_DescriptorMatcher_add_10 (a0, a1, a2, a3);
}

static void (*proc_photo_Photo_decolor_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_photo_Photo_decolor_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_photo_Photo_decolor_10 (a0, a1, a2, a3, a4);
}

static void (*proc_video_DualTVL1OpticalFlow_setMedianFiltering_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_video_DualTVL1OpticalFlow_setMedianFiltering_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_video_DualTVL1OpticalFlow_setMedianFiltering_10 (a0, a1, a2, a3);
}

static jdoubleArray (*proc_ml_EM_predict2_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
jdoubleArray Java_org_opencv_ml_EM_predict2_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    return proc_ml_EM_predict2_10 (a0, a1, a2, a3, a4);
}

static jint (*proc_calib3d_Calib3d_recoverPose_16)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jdouble);
jint Java_org_opencv_calib3d_Calib3d_recoverPose_16(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7, jdouble a8) {
    return proc_calib3d_Calib3d_recoverPose_16 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static jint (*proc_imgproc_Imgproc_connectedComponents_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint);
jint Java_org_opencv_imgproc_Imgproc_connectedComponents_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5) {
    return proc_imgproc_Imgproc_connectedComponents_10 (a0, a1, a2, a3, a4, a5);
}

static jlong (*proc_core_Mat_n_1step1__JI)(JNIEnv*,  jclass,  jlong,  jint);
jlong Java_org_opencv_core_Mat_n_1step1__JI(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    return proc_core_Mat_n_1step1__JI (a0, a1, a2, a3);
}

static jlong (*proc_dnn_Net_getParam_11)(JNIEnv*,  jclass,  jlong,  jlong);
jlong Java_org_opencv_dnn_Net_getParam_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_dnn_Net_getParam_11 (a0, a1, a2, a3);
}

static void (*proc_features2d_DescriptorMatcher_clear_10)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_features2d_DescriptorMatcher_clear_10(JNIEnv* a0, jclass a1, jlong a2) {
    proc_features2d_DescriptorMatcher_clear_10 (a0, a1, a2);
}

static void (*proc_photo_Photo_denoise_1TVL1_10)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jint);
void Java_org_opencv_photo_Photo_denoise_1TVL1_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jint a5) {
    proc_photo_Photo_denoise_1TVL1_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_video_DualTVL1OpticalFlow_setOuterIterations_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_video_DualTVL1OpticalFlow_setOuterIterations_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_video_DualTVL1OpticalFlow_setOuterIterations_10 (a0, a1, a2, a3);
}

static jfloat (*proc_calib3d_Calib3d_rectify3Collinear_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jdouble,  jdouble,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jdouble,  jdouble,  jdouble,  jdoubleArray,  jdoubleArray,  jint);
jfloat Java_org_opencv_calib3d_Calib3d_rectify3Collinear_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7, jlong a8, jlong a9, jdouble a10, jdouble a11, jlong a12, jlong a13, jlong a14, jlong a15, jlong a16, jlong a17, jlong a18, jlong a19, jlong a20, jlong a21, jlong a22, jdouble a23, jdouble a24, jdouble a25, jdoubleArray a26, jdoubleArray a27, jint a28) {
    return proc_calib3d_Calib3d_rectify3Collinear_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20, a21, a22, a23, a24, a25, a26, a27, a28);
}

static jfloat (*proc_ml_EM_predict_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jint);
jfloat Java_org_opencv_ml_EM_predict_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jint a5) {
    return proc_ml_EM_predict_10 (a0, a1, a2, a3, a4, a5);
}

static jint (*proc_imgproc_Imgproc_connectedComponents_11)(JNIEnv*,  jclass,  jlong,  jlong);
jint Java_org_opencv_imgproc_Imgproc_connectedComponents_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_imgproc_Imgproc_connectedComponents_11 (a0, a1, a2, a3);
}

static jlong (*proc_core_Mat_n_1step1__J)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_core_Mat_n_1step1__J(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_core_Mat_n_1step1__J (a0, a1, a2);
}

static jlong (*proc_dnn_Net_getPerfProfile_10)(JNIEnv*,  jclass,  jlong,  jlong);
jlong Java_org_opencv_dnn_Net_getPerfProfile_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_dnn_Net_getPerfProfile_10 (a0, a1, a2, a3);
}

static jlong (*proc_features2d_DescriptorMatcher_clone_10)(JNIEnv*,  jclass,  jlong,  jboolean);
jlong Java_org_opencv_features2d_DescriptorMatcher_clone_10(JNIEnv* a0, jclass a1, jlong a2, jboolean a3) {
    return proc_features2d_DescriptorMatcher_clone_10 (a0, a1, a2, a3);
}

static void (*proc_photo_Photo_denoise_1TVL1_11)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_photo_Photo_denoise_1TVL1_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_photo_Photo_denoise_1TVL1_11 (a0, a1, a2, a3);
}

static void (*proc_video_DualTVL1OpticalFlow_setScaleStep_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_video_DualTVL1OpticalFlow_setScaleStep_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_video_DualTVL1OpticalFlow_setScaleStep_10 (a0, a1, a2, a3);
}

static jdouble (*proc_imgproc_Imgproc_contourArea_10)(JNIEnv*,  jclass,  jlong,  jboolean);
jdouble Java_org_opencv_imgproc_Imgproc_contourArea_10(JNIEnv* a0, jclass a1, jlong a2, jboolean a3) {
    return proc_imgproc_Imgproc_contourArea_10 (a0, a1, a2, a3);
}

static jfloat (*proc_ml_EM_predict_11)(JNIEnv*,  jclass,  jlong,  jlong);
jfloat Java_org_opencv_ml_EM_predict_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_ml_EM_predict_11 (a0, a1, a2, a3);
}

static jlong (*proc_core_Mat_n_1submat)(JNIEnv*,  jclass,  jlong,  jint,  jint,  jint,  jint);
jlong Java_org_opencv_core_Mat_n_1submat(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4, jint a5, jint a6) {
    return proc_core_Mat_n_1submat (a0, a1, a2, a3, a4, a5, a6);
}

static jlong (*proc_dnn_Net_getUnconnectedOutLayers_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_dnn_Net_getUnconnectedOutLayers_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_dnn_Net_getUnconnectedOutLayers_10 (a0, a1, a2);
}

static jlong (*proc_features2d_DescriptorMatcher_clone_11)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_features2d_DescriptorMatcher_clone_11(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_DescriptorMatcher_clone_11 (a0, a1, a2);
}

static void (*proc_calib3d_Calib3d_reprojectImageTo3D_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jboolean,  jint);
void Java_org_opencv_calib3d_Calib3d_reprojectImageTo3D_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jboolean a5, jint a6) {
    proc_calib3d_Calib3d_reprojectImageTo3D_10 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_photo_Photo_detailEnhance_10)(JNIEnv*,  jclass,  jlong,  jlong,  jfloat,  jfloat);
void Java_org_opencv_photo_Photo_detailEnhance_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jfloat a4, jfloat a5) {
    proc_photo_Photo_detailEnhance_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_video_DualTVL1OpticalFlow_setScalesNumber_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_video_DualTVL1OpticalFlow_setScalesNumber_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_video_DualTVL1OpticalFlow_setScalesNumber_10 (a0, a1, a2, a3);
}

static jdouble (*proc_imgproc_Imgproc_contourArea_11)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_imgproc_Imgproc_contourArea_11(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_imgproc_Imgproc_contourArea_11 (a0, a1, a2);
}

static jlong (*proc_core_Mat_n_1submat_1rr)(JNIEnv*,  jclass,  jlong,  jint,  jint,  jint,  jint);
jlong Java_org_opencv_core_Mat_n_1submat_1rr(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4, jint a5, jint a6) {
    return proc_core_Mat_n_1submat_1rr (a0, a1, a2, a3, a4, a5, a6);
}

static jlong (*proc_features2d_DescriptorMatcher_create_10)(JNIEnv*,  jclass,  jstring);
jlong Java_org_opencv_features2d_DescriptorMatcher_create_10(JNIEnv* a0, jclass a1, jstring a2) {
    return proc_features2d_DescriptorMatcher_create_10 (a0, a1, a2);
}

static void (*proc_calib3d_Calib3d_reprojectImageTo3D_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jboolean);
void Java_org_opencv_calib3d_Calib3d_reprojectImageTo3D_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jboolean a5) {
    proc_calib3d_Calib3d_reprojectImageTo3D_11 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_dnn_Net_setHalideScheduler_10)(JNIEnv*,  jclass,  jlong,  jstring);
void Java_org_opencv_dnn_Net_setHalideScheduler_10(JNIEnv* a0, jclass a1, jlong a2, jstring a3) {
    proc_dnn_Net_setHalideScheduler_10 (a0, a1, a2, a3);
}

static void (*proc_ml_EM_setClustersNumber_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_ml_EM_setClustersNumber_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_ml_EM_setClustersNumber_10 (a0, a1, a2, a3);
}

static void (*proc_photo_Photo_detailEnhance_11)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_photo_Photo_detailEnhance_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_photo_Photo_detailEnhance_11 (a0, a1, a2, a3);
}

static void (*proc_video_DualTVL1OpticalFlow_setTau_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_video_DualTVL1OpticalFlow_setTau_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_video_DualTVL1OpticalFlow_setTau_10 (a0, a1, a2, a3);
}

static jdoubleArray (*proc_calib3d_Calib3d_RQDecomp3x3_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
jdoubleArray Java_org_opencv_calib3d_Calib3d_RQDecomp3x3_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    return proc_calib3d_Calib3d_RQDecomp3x3_11 (a0, a1, a2, a3, a4);
}

static jdoubleArray (*proc_core_Mat_nGet)(JNIEnv*,  jclass,  jlong,  jint,  jint);
jdoubleArray Java_org_opencv_core_Mat_nGet(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4) {
    return proc_core_Mat_nGet (a0, a1, a2, a3, a4);
}

static jlong (*proc_dnn_DictValue_DictValue_10)(JNIEnv*,  jclass,  jstring);
jlong Java_org_opencv_dnn_DictValue_DictValue_10(JNIEnv* a0, jclass a1, jstring a2) {
    return proc_dnn_DictValue_DictValue_10 (a0, a1, a2);
}

static jlong (*proc_features2d_AKAZE_create_10)(JNIEnv*,  jclass,  jint,  jint,  jint,  jfloat,  jint,  jint,  jint);
jlong Java_org_opencv_features2d_AKAZE_create_10(JNIEnv* a0, jclass a1, jint a2, jint a3, jint a4, jfloat a5, jint a6, jint a7, jint a8) {
    return proc_features2d_AKAZE_create_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static jlong (*proc_imgcodecs_Imgcodecs_imdecode_10)(JNIEnv*,  jclass,  jlong,  jint);
jlong Java_org_opencv_imgcodecs_Imgcodecs_imdecode_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    return proc_imgcodecs_Imgcodecs_imdecode_10 (a0, a1, a2, a3);
}

static jlong (*proc_ml_ANN_1MLP_create_10)(JNIEnv*,  jclass);
jlong Java_org_opencv_ml_ANN_1MLP_create_10(JNIEnv* a0, jclass a1) {
    return proc_ml_ANN_1MLP_create_10 (a0, a1);
}

static jlong (*proc_videoio_VideoCapture_VideoCapture_10)(JNIEnv*,  jclass,  jstring,  jint);
jlong Java_org_opencv_videoio_VideoCapture_VideoCapture_10(JNIEnv* a0, jclass a1, jstring a2, jint a3) {
    return proc_videoio_VideoCapture_VideoCapture_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_CLAHE_apply_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_imgproc_CLAHE_apply_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_imgproc_CLAHE_apply_10 (a0, a1, a2, a3, a4);
}

static void (*proc_objdetect_Objdetect_groupRectangles_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jdouble);
void Java_org_opencv_objdetect_Objdetect_groupRectangles_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jdouble a5) {
    proc_objdetect_Objdetect_groupRectangles_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_photo_AlignExposures_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_photo_AlignExposures_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_photo_AlignExposures_delete (a0, a1, a2);
}

static void (*proc_video_BackgroundSubtractor_apply_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jdouble);
void Java_org_opencv_video_BackgroundSubtractor_apply_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jdouble a5) {
    proc_video_BackgroundSubtractor_apply_10 (a0, a1, a2, a3, a4, a5);
}

static jlong (*proc_core_Mat_n_1t)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_core_Mat_n_1t(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_core_Mat_n_1t (a0, a1, a2);
}

static jlong (*proc_features2d_DescriptorMatcher_create_11)(JNIEnv*,  jclass,  jint);
jlong Java_org_opencv_features2d_DescriptorMatcher_create_11(JNIEnv* a0, jclass a1, jint a2) {
    return proc_features2d_DescriptorMatcher_create_11 (a0, a1, a2);
}

static void (*proc_calib3d_Calib3d_reprojectImageTo3D_12)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_calib3d_Calib3d_reprojectImageTo3D_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_calib3d_Calib3d_reprojectImageTo3D_12 (a0, a1, a2, a3, a4);
}

static void (*proc_dnn_Net_setInput_10)(JNIEnv*,  jclass,  jlong,  jlong,  jstring);
void Java_org_opencv_dnn_Net_setInput_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jstring a4) {
    proc_dnn_Net_setInput_10 (a0, a1, a2, a3, a4);
}

static void (*proc_imgproc_Imgproc_convertMaps_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jint,  jboolean);
void Java_org_opencv_imgproc_Imgproc_convertMaps_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jint a6, jboolean a7) {
    proc_imgproc_Imgproc_convertMaps_10 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_ml_EM_setCovarianceMatrixType_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_ml_EM_setCovarianceMatrixType_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_ml_EM_setCovarianceMatrixType_10 (a0, a1, a2, a3);
}

static void (*proc_photo_Photo_edgePreservingFilter_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jfloat,  jfloat);
void Java_org_opencv_photo_Photo_edgePreservingFilter_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jfloat a5, jfloat a6) {
    proc_photo_Photo_edgePreservingFilter_10 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_video_DualTVL1OpticalFlow_setTheta_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_video_DualTVL1OpticalFlow_setTheta_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_video_DualTVL1OpticalFlow_setTheta_10 (a0, a1, a2, a3);
}

static jdouble (*proc_calib3d_Calib3d_sampsonDistance_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
jdouble Java_org_opencv_calib3d_Calib3d_sampsonDistance_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    return proc_calib3d_Calib3d_sampsonDistance_10 (a0, a1, a2, a3, a4);
}

static jlong (*proc_core_Mat_n_1total)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_core_Mat_n_1total(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_core_Mat_n_1total (a0, a1, a2);
}

static void (*proc_dnn_Net_setInput_11)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_dnn_Net_setInput_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_dnn_Net_setInput_11 (a0, a1, a2, a3);
}

static void (*proc_features2d_DescriptorMatcher_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_features2d_DescriptorMatcher_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_features2d_DescriptorMatcher_delete (a0, a1, a2);
}

static void (*proc_imgproc_Imgproc_convertMaps_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jint);
void Java_org_opencv_imgproc_Imgproc_convertMaps_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jint a6) {
    proc_imgproc_Imgproc_convertMaps_11 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_ml_EM_setTermCriteria_10)(JNIEnv*,  jclass,  jlong,  jint,  jint,  jdouble);
void Java_org_opencv_ml_EM_setTermCriteria_10(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4, jdouble a5) {
    proc_ml_EM_setTermCriteria_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_photo_Photo_edgePreservingFilter_11)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_photo_Photo_edgePreservingFilter_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_photo_Photo_edgePreservingFilter_11 (a0, a1, a2, a3);
}

static void (*proc_video_DualTVL1OpticalFlow_setUseInitialFlow_10)(JNIEnv*,  jclass,  jlong,  jboolean);
void Java_org_opencv_video_DualTVL1OpticalFlow_setUseInitialFlow_10(JNIEnv* a0, jclass a1, jlong a2, jboolean a3) {
    proc_video_DualTVL1OpticalFlow_setUseInitialFlow_10 (a0, a1, a2, a3);
}

static jboolean (*proc_features2d_DescriptorMatcher_empty_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_features2d_DescriptorMatcher_empty_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_DescriptorMatcher_empty_10 (a0, a1, a2);
}

static jboolean (*proc_ml_EM_trainEM_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong);
jboolean Java_org_opencv_ml_EM_trainEM_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6) {
    return proc_ml_EM_trainEM_10 (a0, a1, a2, a3, a4, a5, a6);
}

static jint (*proc_calib3d_Calib3d_solveP3P_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jint);
jint Java_org_opencv_calib3d_Calib3d_solveP3P_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7, jint a8) {
    return proc_calib3d_Calib3d_solveP3P_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static jint (*proc_core_Mat_n_1type)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_core_Mat_n_1type(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_core_Mat_n_1type (a0, a1, a2);
}

static void (*proc_dnn_Net_setInputsNames_10)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_dnn_Net_setInputsNames_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_dnn_Net_setInputsNames_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_convexHull_10)(JNIEnv*,  jclass,  jlong,  jlong,  jboolean);
void Java_org_opencv_imgproc_Imgproc_convexHull_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jboolean a4) {
    proc_imgproc_Imgproc_convexHull_10 (a0, a1, a2, a3, a4);
}

static void (*proc_photo_Photo_fastNlMeansDenoisingColoredMulti_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint,  jfloat,  jfloat,  jint,  jint);
void Java_org_opencv_photo_Photo_fastNlMeansDenoisingColoredMulti_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5, jfloat a6, jfloat a7, jint a8, jint a9) {
    proc_photo_Photo_fastNlMeansDenoisingColoredMulti_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static void (*proc_video_DualTVL1OpticalFlow_setWarpingsNumber_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_video_DualTVL1OpticalFlow_setWarpingsNumber_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_video_DualTVL1OpticalFlow_setWarpingsNumber_10 (a0, a1, a2, a3);
}

static jboolean (*proc_calib3d_Calib3d_solvePnPRansac_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jboolean,  jint,  jfloat,  jdouble,  jlong,  jint);
jboolean Java_org_opencv_calib3d_Calib3d_solvePnPRansac_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7, jboolean a8, jint a9, jfloat a10, jdouble a11, jlong a12, jint a13) {
    return proc_calib3d_Calib3d_solvePnPRansac_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13);
}

static jboolean (*proc_ml_EM_trainEM_11)(JNIEnv*,  jclass,  jlong,  jlong);
jboolean Java_org_opencv_ml_EM_trainEM_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_ml_EM_trainEM_11 (a0, a1, a2, a3);
}

static jlong (*proc_core_Mat_n_1zeros__III)(JNIEnv*,  jclass,  jint,  jint,  jint);
jlong Java_org_opencv_core_Mat_n_1zeros__III(JNIEnv* a0, jclass a1, jint a2, jint a3, jint a4) {
    return proc_core_Mat_n_1zeros__III (a0, a1, a2, a3, a4);
}

static jlong (*proc_features2d_DescriptorMatcher_getTrainDescriptors_10)(JNIEnv*,  jclass,  jlong);
jlong Java_org_opencv_features2d_DescriptorMatcher_getTrainDescriptors_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_DescriptorMatcher_getTrainDescriptors_10 (a0, a1, a2);
}

static jlong (*proc_video_FarnebackOpticalFlow_create_10)(JNIEnv*,  jclass,  jint,  jdouble,  jboolean,  jint,  jint,  jint,  jdouble,  jint);
jlong Java_org_opencv_video_FarnebackOpticalFlow_create_10(JNIEnv* a0, jclass a1, jint a2, jdouble a3, jboolean a4, jint a5, jint a6, jint a7, jdouble a8, jint a9) {
    return proc_video_FarnebackOpticalFlow_create_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static void (*proc_dnn_Net_setParam_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jlong);
void Java_org_opencv_dnn_Net_setParam_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jlong a5) {
    proc_dnn_Net_setParam_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_imgproc_Imgproc_convexHull_11)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_imgproc_Imgproc_convexHull_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_imgproc_Imgproc_convexHull_11 (a0, a1, a2, a3);
}

static void (*proc_photo_Photo_fastNlMeansDenoisingColoredMulti_11)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint);
void Java_org_opencv_photo_Photo_fastNlMeansDenoisingColoredMulti_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5) {
    proc_photo_Photo_fastNlMeansDenoisingColoredMulti_11 (a0, a1, a2, a3, a4, a5);
}

static jboolean (*proc_calib3d_Calib3d_solvePnPRansac_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong);
jboolean Java_org_opencv_calib3d_Calib3d_solvePnPRansac_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7) {
    return proc_calib3d_Calib3d_solvePnPRansac_11 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static jboolean (*proc_features2d_DescriptorMatcher_isMaskSupported_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_features2d_DescriptorMatcher_isMaskSupported_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_features2d_DescriptorMatcher_isMaskSupported_10 (a0, a1, a2);
}

static jboolean (*proc_ml_EM_trainE_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong);
jboolean Java_org_opencv_ml_EM_trainE_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7, jlong a8, jlong a9) {
    return proc_ml_EM_trainE_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static jlong (*proc_core_Mat_n_1zeros__DDI)(JNIEnv*,  jclass,  jdouble,  jdouble,  jint);
jlong Java_org_opencv_core_Mat_n_1zeros__DDI(JNIEnv* a0, jclass a1, jdouble a2, jdouble a3, jint a4) {
    return proc_core_Mat_n_1zeros__DDI (a0, a1, a2, a3, a4);
}

static jlong (*proc_video_FarnebackOpticalFlow_create_11)(JNIEnv*,  jclass);
jlong Java_org_opencv_video_FarnebackOpticalFlow_create_11(JNIEnv* a0, jclass a1) {
    return proc_video_FarnebackOpticalFlow_create_11 (a0, a1);
}

static void (*proc_dnn_Net_setPreferableBackend_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_dnn_Net_setPreferableBackend_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_dnn_Net_setPreferableBackend_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_convexityDefects_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_imgproc_Imgproc_convexityDefects_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_imgproc_Imgproc_convexityDefects_10 (a0, a1, a2, a3, a4);
}

static void (*proc_photo_Photo_fastNlMeansDenoisingColored_10)(JNIEnv*,  jclass,  jlong,  jlong,  jfloat,  jfloat,  jint,  jint);
void Java_org_opencv_photo_Photo_fastNlMeansDenoisingColored_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jfloat a4, jfloat a5, jint a6, jint a7) {
    proc_photo_Photo_fastNlMeansDenoisingColored_10 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static jboolean (*proc_calib3d_Calib3d_solvePnP_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jboolean,  jint);
jboolean Java_org_opencv_calib3d_Calib3d_solvePnP_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7, jboolean a8, jint a9) {
    return proc_calib3d_Calib3d_solvePnP_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static jboolean (*proc_ml_EM_trainE_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
jboolean Java_org_opencv_ml_EM_trainE_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    return proc_ml_EM_trainE_11 (a0, a1, a2, a3, a4);
}

static void (*proc_core_Core_LUT_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_core_Core_LUT_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_core_Core_LUT_10 (a0, a1, a2, a3, a4);
}

static void (*proc_dnn_Net_setPreferableTarget_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_dnn_Net_setPreferableTarget_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_dnn_Net_setPreferableTarget_10 (a0, a1, a2, a3);
}

static void (*proc_features2d_DescriptorMatcher_knnMatch_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jint,  jlong,  jboolean);
void Java_org_opencv_features2d_DescriptorMatcher_knnMatch_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jint a6, jlong a7, jboolean a8) {
    proc_features2d_DescriptorMatcher_knnMatch_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static void (*proc_imgproc_Imgproc_cornerEigenValsAndVecs_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint,  jint);
void Java_org_opencv_imgproc_Imgproc_cornerEigenValsAndVecs_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5, jint a6) {
    proc_imgproc_Imgproc_cornerEigenValsAndVecs_10 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_photo_Photo_fastNlMeansDenoisingColored_11)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_photo_Photo_fastNlMeansDenoisingColored_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_photo_Photo_fastNlMeansDenoisingColored_11 (a0, a1, a2, a3);
}

static void (*proc_video_FarnebackOpticalFlow_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_video_FarnebackOpticalFlow_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_video_FarnebackOpticalFlow_delete (a0, a1, a2);
}

static jboolean (*proc_calib3d_Calib3d_solvePnP_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong);
jboolean Java_org_opencv_calib3d_Calib3d_solvePnP_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7) {
    return proc_calib3d_Calib3d_solvePnP_11 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static jboolean (*proc_ml_EM_trainM_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong);
jboolean Java_org_opencv_ml_EM_trainM_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7) {
    return proc_ml_EM_trainM_10 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static jboolean (*proc_video_FarnebackOpticalFlow_getFastPyramids_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_video_FarnebackOpticalFlow_getFastPyramids_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_FarnebackOpticalFlow_getFastPyramids_10 (a0, a1, a2);
}

static jdouble (*proc_core_Core_Mahalanobis_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
jdouble Java_org_opencv_core_Core_Mahalanobis_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    return proc_core_Core_Mahalanobis_10 (a0, a1, a2, a3, a4);
}

static void (*proc_features2d_DescriptorMatcher_knnMatch_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jint);
void Java_org_opencv_features2d_DescriptorMatcher_knnMatch_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jint a6) {
    proc_features2d_DescriptorMatcher_knnMatch_11 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_imgproc_Imgproc_cornerEigenValsAndVecs_11)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint);
void Java_org_opencv_imgproc_Imgproc_cornerEigenValsAndVecs_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5) {
    proc_imgproc_Imgproc_cornerEigenValsAndVecs_11 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_photo_Photo_fastNlMeansDenoisingMulti_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint,  jfloat,  jint,  jint);
void Java_org_opencv_photo_Photo_fastNlMeansDenoisingMulti_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5, jfloat a6, jint a7, jint a8) {
    proc_photo_Photo_fastNlMeansDenoisingMulti_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static jboolean (*proc_ml_EM_trainM_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
jboolean Java_org_opencv_ml_EM_trainM_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    return proc_ml_EM_trainM_11 (a0, a1, a2, a3, a4);
}

static jdouble (*proc_calib3d_Calib3d_stereoCalibrate_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jdouble,  jdouble,  jlong,  jlong,  jlong,  jlong,  jint,  jint,  jint,  jdouble);
jdouble Java_org_opencv_calib3d_Calib3d_stereoCalibrate_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7, jlong a8, jdouble a9, jdouble a10, jlong a11, jlong a12, jlong a13, jlong a14, jint a15, jint a16, jint a17, jdouble a18) {
    return proc_calib3d_Calib3d_stereoCalibrate_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18);
}

static jint (*proc_video_FarnebackOpticalFlow_getFlags_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_video_FarnebackOpticalFlow_getFlags_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_FarnebackOpticalFlow_getFlags_10 (a0, a1, a2);
}

static void (*proc_core_Core_PCABackProject_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_core_Core_PCABackProject_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_core_Core_PCABackProject_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_features2d_DescriptorMatcher_knnMatch_12)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jint,  jlong,  jboolean);
void Java_org_opencv_features2d_DescriptorMatcher_knnMatch_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jint a5, jlong a6, jboolean a7) {
    proc_features2d_DescriptorMatcher_knnMatch_12 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_imgproc_Imgproc_cornerHarris_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint,  jdouble,  jint);
void Java_org_opencv_imgproc_Imgproc_cornerHarris_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5, jdouble a6, jint a7) {
    proc_imgproc_Imgproc_cornerHarris_10 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_photo_Photo_fastNlMeansDenoisingMulti_11)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint);
void Java_org_opencv_photo_Photo_fastNlMeansDenoisingMulti_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5) {
    proc_photo_Photo_fastNlMeansDenoisingMulti_11 (a0, a1, a2, a3, a4, a5);
}

static jdouble (*proc_calib3d_Calib3d_stereoCalibrate_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jdouble,  jdouble,  jlong,  jlong,  jlong,  jlong,  jint);
jdouble Java_org_opencv_calib3d_Calib3d_stereoCalibrate_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7, jlong a8, jdouble a9, jdouble a10, jlong a11, jlong a12, jlong a13, jlong a14, jint a15) {
    return proc_calib3d_Calib3d_stereoCalibrate_11 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15);
}

static jint (*proc_video_FarnebackOpticalFlow_getNumIters_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_video_FarnebackOpticalFlow_getNumIters_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_FarnebackOpticalFlow_getNumIters_10 (a0, a1, a2);
}

static jlong (*proc_ml_KNearest_create_10)(JNIEnv*,  jclass);
jlong Java_org_opencv_ml_KNearest_create_10(JNIEnv* a0, jclass a1) {
    return proc_ml_KNearest_create_10 (a0, a1);
}

static void (*proc_core_Core_PCACompute_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jdouble);
void Java_org_opencv_core_Core_PCACompute_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jdouble a5) {
    proc_core_Core_PCACompute_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_features2d_DescriptorMatcher_knnMatch_13)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jint);
void Java_org_opencv_features2d_DescriptorMatcher_knnMatch_13(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jint a5) {
    proc_features2d_DescriptorMatcher_knnMatch_13 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_imgproc_Imgproc_cornerHarris_11)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint,  jdouble);
void Java_org_opencv_imgproc_Imgproc_cornerHarris_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5, jdouble a6) {
    proc_imgproc_Imgproc_cornerHarris_11 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_photo_Photo_fastNlMeansDenoisingMulti_12)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint,  jlong,  jint,  jint,  jint);
void Java_org_opencv_photo_Photo_fastNlMeansDenoisingMulti_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5, jlong a6, jint a7, jint a8, jint a9) {
    proc_photo_Photo_fastNlMeansDenoisingMulti_12 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static jdouble (*proc_calib3d_Calib3d_stereoCalibrate_12)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jdouble,  jdouble,  jlong,  jlong,  jlong,  jlong);
jdouble Java_org_opencv_calib3d_Calib3d_stereoCalibrate_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7, jlong a8, jdouble a9, jdouble a10, jlong a11, jlong a12, jlong a13, jlong a14) {
    return proc_calib3d_Calib3d_stereoCalibrate_12 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14);
}

static jint (*proc_video_FarnebackOpticalFlow_getNumLevels_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_video_FarnebackOpticalFlow_getNumLevels_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_FarnebackOpticalFlow_getNumLevels_10 (a0, a1, a2);
}

static void (*proc_core_Core_PCACompute_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jint);
void Java_org_opencv_core_Core_PCACompute_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jint a5) {
    proc_core_Core_PCACompute_11 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_features2d_DescriptorMatcher_match_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_features2d_DescriptorMatcher_match_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6) {
    proc_features2d_DescriptorMatcher_match_10 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_imgproc_Imgproc_cornerMinEigenVal_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint,  jint);
void Java_org_opencv_imgproc_Imgproc_cornerMinEigenVal_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5, jint a6) {
    proc_imgproc_Imgproc_cornerMinEigenVal_10 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_ml_KNearest_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_ml_KNearest_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_ml_KNearest_delete (a0, a1, a2);
}

static void (*proc_photo_Photo_fastNlMeansDenoisingMulti_13)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint,  jlong);
void Java_org_opencv_photo_Photo_fastNlMeansDenoisingMulti_13(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5, jlong a6) {
    proc_photo_Photo_fastNlMeansDenoisingMulti_13 (a0, a1, a2, a3, a4, a5, a6);
}

static jboolean (*proc_imgcodecs_Imgcodecs_imencode_10)(JNIEnv*,  jclass,  jstring,  jlong,  jlong,  jlong);
jboolean Java_org_opencv_imgcodecs_Imgcodecs_imencode_10(JNIEnv* a0, jclass a1, jstring a2, jlong a3, jlong a4, jlong a5) {
    return proc_imgcodecs_Imgcodecs_imencode_10 (a0, a1, a2, a3, a4, a5);
}

static jint (*proc_core_Mat_nGetB)(JNIEnv*,  jclass,  jlong,  jint,  jint,  jint,  jbyteArray);
jint Java_org_opencv_core_Mat_nGetB(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4, jint a5, jbyteArray a6) {
    return proc_core_Mat_nGetB (a0, a1, a2, a3, a4, a5, a6);
}

static jlong (*proc_dnn_DictValue_DictValue_11)(JNIEnv*,  jclass,  jdouble);
jlong Java_org_opencv_dnn_DictValue_DictValue_11(JNIEnv* a0, jclass a1, jdouble a2) {
    return proc_dnn_DictValue_DictValue_11 (a0, a1, a2);
}

static jlong (*proc_features2d_AKAZE_create_11)(JNIEnv*,  jclass);
jlong Java_org_opencv_features2d_AKAZE_create_11(JNIEnv* a0, jclass a1) {
    return proc_features2d_AKAZE_create_11 (a0, a1);
}

static jlong (*proc_videoio_VideoCapture_VideoCapture_11)(JNIEnv*,  jclass,  jstring);
jlong Java_org_opencv_videoio_VideoCapture_VideoCapture_11(JNIEnv* a0, jclass a1, jstring a2) {
    return proc_videoio_VideoCapture_VideoCapture_11 (a0, a1, a2);
}

static void (*proc_calib3d_Calib3d_Rodrigues_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_calib3d_Calib3d_Rodrigues_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_calib3d_Calib3d_Rodrigues_10 (a0, a1, a2, a3, a4);
}

static void (*proc_imgproc_CLAHE_collectGarbage_10)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_imgproc_CLAHE_collectGarbage_10(JNIEnv* a0, jclass a1, jlong a2) {
    proc_imgproc_CLAHE_collectGarbage_10 (a0, a1, a2);
}

static void (*proc_ml_ANN_1MLP_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_ml_ANN_1MLP_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_ml_ANN_1MLP_delete (a0, a1, a2);
}

static void (*proc_objdetect_Objdetect_groupRectangles_11)(JNIEnv*,  jclass,  jlong,  jlong,  jint);
void Java_org_opencv_objdetect_Objdetect_groupRectangles_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4) {
    proc_objdetect_Objdetect_groupRectangles_11 (a0, a1, a2, a3, a4);
}

static void (*proc_photo_AlignExposures_process_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_photo_AlignExposures_process_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6) {
    proc_photo_AlignExposures_process_10 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_video_BackgroundSubtractor_apply_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_video_BackgroundSubtractor_apply_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_video_BackgroundSubtractor_apply_11 (a0, a1, a2, a3, a4);
}

static jdouble (*proc_calib3d_Calib3d_stereoCalibrate_13)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jdouble,  jdouble,  jlong,  jlong,  jint,  jint,  jint,  jdouble);
jdouble Java_org_opencv_calib3d_Calib3d_stereoCalibrate_13(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7, jlong a8, jdouble a9, jdouble a10, jlong a11, jlong a12, jint a13, jint a14, jint a15, jdouble a16) {
    return proc_calib3d_Calib3d_stereoCalibrate_13 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16);
}

static jfloat (*proc_ml_KNearest_findNearest_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jlong,  jlong,  jlong);
jfloat Java_org_opencv_ml_KNearest_findNearest_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jlong a5, jlong a6, jlong a7) {
    return proc_ml_KNearest_findNearest_10 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static jint (*proc_video_FarnebackOpticalFlow_getPolyN_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_video_FarnebackOpticalFlow_getPolyN_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_FarnebackOpticalFlow_getPolyN_10 (a0, a1, a2);
}

static void (*proc_core_Core_PCACompute_12)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_core_Core_PCACompute_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_core_Core_PCACompute_12 (a0, a1, a2, a3, a4);
}

static void (*proc_features2d_DescriptorMatcher_match_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_features2d_DescriptorMatcher_match_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_features2d_DescriptorMatcher_match_11 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_imgproc_Imgproc_cornerMinEigenVal_11)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint);
void Java_org_opencv_imgproc_Imgproc_cornerMinEigenVal_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5) {
    proc_imgproc_Imgproc_cornerMinEigenVal_11 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_photo_Photo_fastNlMeansDenoising_10)(JNIEnv*,  jclass,  jlong,  jlong,  jfloat,  jint,  jint);
void Java_org_opencv_photo_Photo_fastNlMeansDenoising_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jfloat a4, jint a5, jint a6) {
    proc_photo_Photo_fastNlMeansDenoising_10 (a0, a1, a2, a3, a4, a5, a6);
}

static jdouble (*proc_calib3d_Calib3d_stereoCalibrate_14)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jdouble,  jdouble,  jlong,  jlong,  jint);
jdouble Java_org_opencv_calib3d_Calib3d_stereoCalibrate_14(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7, jlong a8, jdouble a9, jdouble a10, jlong a11, jlong a12, jint a13) {
    return proc_calib3d_Calib3d_stereoCalibrate_14 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13);
}

static jdouble (*proc_video_FarnebackOpticalFlow_getPolySigma_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_video_FarnebackOpticalFlow_getPolySigma_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_FarnebackOpticalFlow_getPolySigma_10 (a0, a1, a2);
}

static jfloat (*proc_ml_KNearest_findNearest_11)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jlong);
jfloat Java_org_opencv_ml_KNearest_findNearest_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jlong a5) {
    return proc_ml_KNearest_findNearest_11 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_core_Core_PCAProject_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_core_Core_PCAProject_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_core_Core_PCAProject_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_features2d_DescriptorMatcher_match_12)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_features2d_DescriptorMatcher_match_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_features2d_DescriptorMatcher_match_12 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_imgproc_Imgproc_cornerMinEigenVal_12)(JNIEnv*,  jclass,  jlong,  jlong,  jint);
void Java_org_opencv_imgproc_Imgproc_cornerMinEigenVal_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4) {
    proc_imgproc_Imgproc_cornerMinEigenVal_12 (a0, a1, a2, a3, a4);
}

static void (*proc_photo_Photo_fastNlMeansDenoising_11)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_photo_Photo_fastNlMeansDenoising_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_photo_Photo_fastNlMeansDenoising_11 (a0, a1, a2, a3);
}

static jdouble (*proc_calib3d_Calib3d_stereoCalibrate_15)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jdouble,  jdouble,  jlong,  jlong);
jdouble Java_org_opencv_calib3d_Calib3d_stereoCalibrate_15(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6, jlong a7, jlong a8, jdouble a9, jdouble a10, jlong a11, jlong a12) {
    return proc_calib3d_Calib3d_stereoCalibrate_15 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12);
}

static jdouble (*proc_core_Core_PSNR_10)(JNIEnv*,  jclass,  jlong,  jlong);
jdouble Java_org_opencv_core_Core_PSNR_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    return proc_core_Core_PSNR_10 (a0, a1, a2, a3);
}

static jdouble (*proc_video_FarnebackOpticalFlow_getPyrScale_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_video_FarnebackOpticalFlow_getPyrScale_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_FarnebackOpticalFlow_getPyrScale_10 (a0, a1, a2);
}

static jint (*proc_ml_KNearest_getAlgorithmType_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_ml_KNearest_getAlgorithmType_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_KNearest_getAlgorithmType_10 (a0, a1, a2);
}

static void (*proc_features2d_DescriptorMatcher_match_13)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_features2d_DescriptorMatcher_match_13(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_features2d_DescriptorMatcher_match_13 (a0, a1, a2, a3, a4);
}

static void (*proc_imgproc_Imgproc_cornerSubPix_10)(JNIEnv*,  jclass,  jlong,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jint,  jint,  jdouble);
void Java_org_opencv_imgproc_Imgproc_cornerSubPix_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7, jint a8, jint a9, jdouble a10) {
    proc_imgproc_Imgproc_cornerSubPix_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10);
}

static void (*proc_photo_Photo_fastNlMeansDenoising_12)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jint,  jint,  jint);
void Java_org_opencv_photo_Photo_fastNlMeansDenoising_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jint a5, jint a6, jint a7) {
    proc_photo_Photo_fastNlMeansDenoising_12 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static jboolean (*proc_calib3d_Calib3d_stereoRectifyUncalibrated_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jdouble,  jdouble,  jlong,  jlong,  jdouble);
jboolean Java_org_opencv_calib3d_Calib3d_stereoRectifyUncalibrated_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jdouble a5, jdouble a6, jlong a7, jlong a8, jdouble a9) {
    return proc_calib3d_Calib3d_stereoRectifyUncalibrated_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static jint (*proc_ml_KNearest_getDefaultK_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_ml_KNearest_getDefaultK_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_KNearest_getDefaultK_10 (a0, a1, a2);
}

static jint (*proc_video_FarnebackOpticalFlow_getWinSize_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_video_FarnebackOpticalFlow_getWinSize_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_video_FarnebackOpticalFlow_getWinSize_10 (a0, a1, a2);
}

static jlong (*proc_imgproc_Imgproc_createCLAHE_10)(JNIEnv*,  jclass,  jdouble,  jdouble,  jdouble);
jlong Java_org_opencv_imgproc_Imgproc_createCLAHE_10(JNIEnv* a0, jclass a1, jdouble a2, jdouble a3, jdouble a4) {
    return proc_imgproc_Imgproc_createCLAHE_10 (a0, a1, a2, a3, a4);
}

static void (*proc_core_Core_SVBackSubst_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_core_Core_SVBackSubst_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6) {
    proc_core_Core_SVBackSubst_10 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_features2d_DescriptorMatcher_radiusMatch_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jfloat,  jlong,  jboolean);
void Java_org_opencv_features2d_DescriptorMatcher_radiusMatch_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jfloat a6, jlong a7, jboolean a8) {
    proc_features2d_DescriptorMatcher_radiusMatch_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static void (*proc_photo_Photo_fastNlMeansDenoising_13)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_photo_Photo_fastNlMeansDenoising_13(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_photo_Photo_fastNlMeansDenoising_13 (a0, a1, a2, a3, a4);
}

static jboolean (*proc_calib3d_Calib3d_stereoRectifyUncalibrated_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jdouble,  jdouble,  jlong,  jlong);
jboolean Java_org_opencv_calib3d_Calib3d_stereoRectifyUncalibrated_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jdouble a5, jdouble a6, jlong a7, jlong a8) {
    return proc_calib3d_Calib3d_stereoRectifyUncalibrated_11 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static jint (*proc_ml_KNearest_getEmax_10)(JNIEnv*,  jclass,  jlong);
jint Java_org_opencv_ml_KNearest_getEmax_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_KNearest_getEmax_10 (a0, a1, a2);
}

static jlong (*proc_imgproc_Imgproc_createCLAHE_11)(JNIEnv*,  jclass);
jlong Java_org_opencv_imgproc_Imgproc_createCLAHE_11(JNIEnv* a0, jclass a1) {
    return proc_imgproc_Imgproc_createCLAHE_11 (a0, a1);
}

static void (*proc_core_Core_SVDecomp_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jint);
void Java_org_opencv_core_Core_SVDecomp_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jint a6) {
    proc_core_Core_SVDecomp_10 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_features2d_DescriptorMatcher_radiusMatch_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jfloat);
void Java_org_opencv_features2d_DescriptorMatcher_radiusMatch_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jfloat a6) {
    proc_features2d_DescriptorMatcher_radiusMatch_11 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_photo_Photo_illuminationChange_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jfloat,  jfloat);
void Java_org_opencv_photo_Photo_illuminationChange_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jfloat a5, jfloat a6) {
    proc_photo_Photo_illuminationChange_10 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_video_FarnebackOpticalFlow_setFastPyramids_10)(JNIEnv*,  jclass,  jlong,  jboolean);
void Java_org_opencv_video_FarnebackOpticalFlow_setFastPyramids_10(JNIEnv* a0, jclass a1, jlong a2, jboolean a3) {
    proc_video_FarnebackOpticalFlow_setFastPyramids_10 (a0, a1, a2, a3);
}

static jboolean (*proc_ml_KNearest_getIsClassifier_10)(JNIEnv*,  jclass,  jlong);
jboolean Java_org_opencv_ml_KNearest_getIsClassifier_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_KNearest_getIsClassifier_10 (a0, a1, a2);
}

static void (*proc_calib3d_Calib3d_stereoRectify_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jdouble,  jdouble,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jint,  jdouble,  jdouble,  jdouble,  jdoubleArray,  jdoubleArray);
void Java_org_opencv_calib3d_Calib3d_stereoRectify_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jdouble a6, jdouble a7, jlong a8, jlong a9, jlong a10, jlong a11, jlong a12, jlong a13, jlong a14, jint a15, jdouble a16, jdouble a17, jdouble a18, jdoubleArray a19, jdoubleArray a20) {
    proc_calib3d_Calib3d_stereoRectify_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19, a20);
}

static void (*proc_core_Core_SVDecomp_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_core_Core_SVDecomp_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5) {
    proc_core_Core_SVDecomp_11 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_features2d_DescriptorMatcher_radiusMatch_12)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jfloat,  jlong,  jboolean);
void Java_org_opencv_features2d_DescriptorMatcher_radiusMatch_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jfloat a5, jlong a6, jboolean a7) {
    proc_features2d_DescriptorMatcher_radiusMatch_12 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_imgproc_Imgproc_createHanningWindow_10)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jint);
void Java_org_opencv_imgproc_Imgproc_createHanningWindow_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jint a5) {
    proc_imgproc_Imgproc_createHanningWindow_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_photo_Photo_illuminationChange_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_photo_Photo_illuminationChange_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_photo_Photo_illuminationChange_11 (a0, a1, a2, a3, a4);
}

static void (*proc_video_FarnebackOpticalFlow_setFlags_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_video_FarnebackOpticalFlow_setFlags_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_video_FarnebackOpticalFlow_setFlags_10 (a0, a1, a2, a3);
}

static jlong (*proc_imgproc_Imgproc_createLineSegmentDetector_10)(JNIEnv*,  jclass,  jint,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jint);
jlong Java_org_opencv_imgproc_Imgproc_createLineSegmentDetector_10(JNIEnv* a0, jclass a1, jint a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8, jint a9) {
    return proc_imgproc_Imgproc_createLineSegmentDetector_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9);
}

static void (*proc_calib3d_Calib3d_stereoRectify_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jdouble,  jdouble,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_calib3d_Calib3d_stereoRectify_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jdouble a6, jdouble a7, jlong a8, jlong a9, jlong a10, jlong a11, jlong a12, jlong a13, jlong a14) {
    proc_calib3d_Calib3d_stereoRectify_11 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14);
}

static void (*proc_core_Core_absdiff_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_core_Core_absdiff_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_core_Core_absdiff_10 (a0, a1, a2, a3, a4);
}

static void (*proc_features2d_DescriptorMatcher_radiusMatch_13)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jfloat);
void Java_org_opencv_features2d_DescriptorMatcher_radiusMatch_13(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jfloat a5) {
    proc_features2d_DescriptorMatcher_radiusMatch_13 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_ml_KNearest_setAlgorithmType_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_ml_KNearest_setAlgorithmType_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_ml_KNearest_setAlgorithmType_10 (a0, a1, a2, a3);
}

static void (*proc_photo_Photo_inpaint_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jdouble,  jint);
void Java_org_opencv_photo_Photo_inpaint_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jdouble a5, jint a6) {
    proc_photo_Photo_inpaint_10 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_video_FarnebackOpticalFlow_setNumIters_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_video_FarnebackOpticalFlow_setNumIters_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_video_FarnebackOpticalFlow_setNumIters_10 (a0, a1, a2, a3);
}

static jlong (*proc_imgproc_Imgproc_createLineSegmentDetector_11)(JNIEnv*,  jclass);
jlong Java_org_opencv_imgproc_Imgproc_createLineSegmentDetector_11(JNIEnv* a0, jclass a1) {
    return proc_imgproc_Imgproc_createLineSegmentDetector_11 (a0, a1);
}

static void (*proc_calib3d_Calib3d_stereoRectify_12)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jdouble,  jdouble,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jint,  jdouble,  jdouble,  jdouble,  jdouble);
void Java_org_opencv_calib3d_Calib3d_stereoRectify_12(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jdouble a6, jdouble a7, jlong a8, jlong a9, jlong a10, jlong a11, jlong a12, jlong a13, jlong a14, jint a15, jdouble a16, jdouble a17, jdouble a18, jdouble a19) {
    proc_calib3d_Calib3d_stereoRectify_12 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18, a19);
}

static void (*proc_core_Core_absdiff_11)(JNIEnv*,  jclass,  jlong,  jdouble,  jdouble,  jdouble,  jdouble,  jlong);
void Java_org_opencv_core_Core_absdiff_11(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jlong a7) {
    proc_core_Core_absdiff_11 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_features2d_DescriptorMatcher_read_10)(JNIEnv*,  jclass,  jlong,  jstring);
void Java_org_opencv_features2d_DescriptorMatcher_read_10(JNIEnv* a0, jclass a1, jlong a2, jstring a3) {
    proc_features2d_DescriptorMatcher_read_10 (a0, a1, a2, a3);
}

static void (*proc_ml_KNearest_setDefaultK_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_ml_KNearest_setDefaultK_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_ml_KNearest_setDefaultK_10 (a0, a1, a2, a3);
}

static void (*proc_photo_Photo_pencilSketch_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jfloat,  jfloat,  jfloat);
void Java_org_opencv_photo_Photo_pencilSketch_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jfloat a5, jfloat a6, jfloat a7) {
    proc_photo_Photo_pencilSketch_10 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_video_FarnebackOpticalFlow_setNumLevels_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_video_FarnebackOpticalFlow_setNumLevels_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_video_FarnebackOpticalFlow_setNumLevels_10 (a0, a1, a2, a3);
}

static void (*proc_calib3d_Calib3d_stereoRectify_13)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jdouble,  jdouble,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jlong,  jint);
void Java_org_opencv_calib3d_Calib3d_stereoRectify_13(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jdouble a6, jdouble a7, jlong a8, jlong a9, jlong a10, jlong a11, jlong a12, jlong a13, jlong a14, jint a15) {
    proc_calib3d_Calib3d_stereoRectify_13 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15);
}

static void (*proc_core_Core_addWeighted_10)(JNIEnv*,  jclass,  jlong,  jdouble,  jlong,  jdouble,  jdouble,  jlong,  jint);
void Java_org_opencv_core_Core_addWeighted_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jlong a4, jdouble a5, jdouble a6, jlong a7, jint a8) {
    proc_core_Core_addWeighted_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static void (*proc_features2d_DescriptorMatcher_train_10)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_features2d_DescriptorMatcher_train_10(JNIEnv* a0, jclass a1, jlong a2) {
    proc_features2d_DescriptorMatcher_train_10 (a0, a1, a2);
}

static void (*proc_imgproc_Imgproc_cvtColor_10)(JNIEnv*,  jclass,  jlong,  jlong,  jint,  jint);
void Java_org_opencv_imgproc_Imgproc_cvtColor_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4, jint a5) {
    proc_imgproc_Imgproc_cvtColor_10 (a0, a1, a2, a3, a4, a5);
}

static void (*proc_ml_KNearest_setEmax_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_ml_KNearest_setEmax_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_ml_KNearest_setEmax_10 (a0, a1, a2, a3);
}

static void (*proc_photo_Photo_pencilSketch_11)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
void Java_org_opencv_photo_Photo_pencilSketch_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    proc_photo_Photo_pencilSketch_11 (a0, a1, a2, a3, a4);
}

static void (*proc_video_FarnebackOpticalFlow_setPolyN_10)(JNIEnv*,  jclass,  jlong,  jint);
void Java_org_opencv_video_FarnebackOpticalFlow_setPolyN_10(JNIEnv* a0, jclass a1, jlong a2, jint a3) {
    proc_video_FarnebackOpticalFlow_setPolyN_10 (a0, a1, a2, a3);
}

static void (*proc_calib3d_Calib3d_triangulatePoints_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jlong,  jlong);
void Java_org_opencv_calib3d_Calib3d_triangulatePoints_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jlong a5, jlong a6) {
    proc_calib3d_Calib3d_triangulatePoints_10 (a0, a1, a2, a3, a4, a5, a6);
}

static void (*proc_core_Core_addWeighted_11)(JNIEnv*,  jclass,  jlong,  jdouble,  jlong,  jdouble,  jdouble,  jlong);
void Java_org_opencv_core_Core_addWeighted_11(JNIEnv* a0, jclass a1, jlong a2, jdouble a3, jlong a4, jdouble a5, jdouble a6, jlong a7) {
    proc_core_Core_addWeighted_11 (a0, a1, a2, a3, a4, a5, a6, a7);
}

static void (*proc_features2d_DescriptorMatcher_write_10)(JNIEnv*,  jclass,  jlong,  jstring);
void Java_org_opencv_features2d_DescriptorMatcher_write_10(JNIEnv* a0, jclass a1, jlong a2, jstring a3) {
    proc_features2d_DescriptorMatcher_write_10 (a0, a1, a2, a3);
}

static void (*proc_imgproc_Imgproc_cvtColor_11)(JNIEnv*,  jclass,  jlong,  jlong,  jint);
void Java_org_opencv_imgproc_Imgproc_cvtColor_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jint a4) {
    proc_imgproc_Imgproc_cvtColor_11 (a0, a1, a2, a3, a4);
}

static void (*proc_ml_KNearest_setIsClassifier_10)(JNIEnv*,  jclass,  jlong,  jboolean);
void Java_org_opencv_ml_KNearest_setIsClassifier_10(JNIEnv* a0, jclass a1, jlong a2, jboolean a3) {
    proc_ml_KNearest_setIsClassifier_10 (a0, a1, a2, a3);
}

static void (*proc_photo_Photo_seamlessClone_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong,  jdouble,  jdouble,  jlong,  jint);
void Java_org_opencv_photo_Photo_seamlessClone_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4, jdouble a5, jdouble a6, jlong a7, jint a8) {
    proc_photo_Photo_seamlessClone_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8);
}

static void (*proc_video_FarnebackOpticalFlow_setPolySigma_10)(JNIEnv*,  jclass,  jlong,  jdouble);
void Java_org_opencv_video_FarnebackOpticalFlow_setPolySigma_10(JNIEnv* a0, jclass a1, jlong a2, jdouble a3) {
    proc_video_FarnebackOpticalFlow_setPolySigma_10 (a0, a1, a2, a3);
}

static jboolean (*proc_imgcodecs_Imgcodecs_imencode_11)(JNIEnv*,  jclass,  jstring,  jlong,  jlong);
jboolean Java_org_opencv_imgcodecs_Imgcodecs_imencode_11(JNIEnv* a0, jclass a1, jstring a2, jlong a3, jlong a4) {
    return proc_imgcodecs_Imgcodecs_imencode_11 (a0, a1, a2, a3, a4);
}

static jdouble (*proc_ml_ANN_1MLP_getBackpropMomentumScale_10)(JNIEnv*,  jclass,  jlong);
jdouble Java_org_opencv_ml_ANN_1MLP_getBackpropMomentumScale_10(JNIEnv* a0, jclass a1, jlong a2) {
    return proc_ml_ANN_1MLP_getBackpropMomentumScale_10 (a0, a1, a2);
}

static jdoubleArray (*proc_photo_AlignMTB_calculateShift_10)(JNIEnv*,  jclass,  jlong,  jlong,  jlong);
jdoubleArray Java_org_opencv_photo_AlignMTB_calculateShift_10(JNIEnv* a0, jclass a1, jlong a2, jlong a3, jlong a4) {
    return proc_photo_AlignMTB_calculateShift_10 (a0, a1, a2, a3, a4);
}

static jint (*proc_core_Mat_nGetD)(JNIEnv*,  jclass,  jlong,  jint,  jint,  jint,  jdoubleArray);
jint Java_org_opencv_core_Mat_nGetD(JNIEnv* a0, jclass a1, jlong a2, jint a3, jint a4, jint a5, jdoubleArray a6) {
    return proc_core_Mat_nGetD (a0, a1, a2, a3, a4, a5, a6);
}

static jlong (*proc_dnn_DictValue_DictValue_12)(JNIEnv*,  jclass,  jint);
jlong Java_org_opencv_dnn_DictValue_DictValue_12(JNIEnv* a0, jclass a1, jint a2) {
    return proc_dnn_DictValue_DictValue_12 (a0, a1, a2);
}

static jlong (*proc_objdetect_HOGDescriptor_HOGDescriptor_10)(JNIEnv*,  jclass,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jdouble,  jint,  jint,  jdouble,  jint,  jdouble,  jboolean,  jint,  jboolean);
jlong Java_org_opencv_objdetect_HOGDescriptor_HOGDescriptor_10(JNIEnv* a0, jclass a1, jdouble a2, jdouble a3, jdouble a4, jdouble a5, jdouble a6, jdouble a7, jdouble a8, jdouble a9, jint a10, jint a11, jdouble a12, jint a13, jdouble a14, jboolean a15, jint a16, jboolean a17) {
    return proc_objdetect_HOGDescriptor_HOGDescriptor_10 (a0, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17);
}

static jlong (*proc_videoio_VideoCapture_VideoCapture_12)(JNIEnv*,  jclass,  jint);
jlong Java_org_opencv_videoio_VideoCapture_VideoCapture_12(JNIEnv* a0, jclass a1, jint a2) {
    return proc_videoio_VideoCapture_VideoCapture_12 (a0, a1, a2);
}

static void (*proc_calib3d_Calib3d_Rodrigues_11)(JNIEnv*,  jclass,  jlong,  jlong);
void Java_org_opencv_calib3d_Calib3d_Rodrigues_11(JNIEnv* a0, jclass a1, jlong a2, jlong a3) {
    proc_calib3d_Calib3d_Rodrigues_11 (a0, a1, a2, a3);
}

static void (*proc_features2d_AKAZE_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_features2d_AKAZE_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_features2d_AKAZE_delete (a0, a1, a2);
}

static void (*proc_imgproc_CLAHE_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_imgproc_CLAHE_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_imgproc_CLAHE_delete (a0, a1, a2);
}

static void (*proc_video_BackgroundSubtractor_delete)(JNIEnv*,  jclass,  jlong);
void Java_org_opencv_video_BackgroundSubtractor_delete(JNIEnv* a0, jclass a1, jlong a2) {
    proc_video_BackgroundSubtractor_delete (a0, a1, a2);
}


int InitProcs() {
	failCnt = 0;
	proc_JNI_OnLoad = procLoad("JNI_OnLoad");
	if (proc_JNI_OnLoad == NULL) {
		return failCnt;
	}

    proc_features2d_FastFeatureDetector_create_10 = procLoad("Java_org_opencv_features2d_FastFeatureDetector_create_10");
    proc_ml_LogisticRegression_create_10 = procLoad("Java_org_opencv_ml_LogisticRegression_create_10");
    proc_calib3d_Calib3d_undistortImage_10 = procLoad("Java_org_opencv_calib3d_Calib3d_undistortImage_10");
    proc_core_Core_add_10 = procLoad("Java_org_opencv_core_Core_add_10");
    proc_imgproc_Imgproc_demosaicing_10 = procLoad("Java_org_opencv_imgproc_Imgproc_demosaicing_10");
    proc_photo_Photo_stylization_10 = procLoad("Java_org_opencv_photo_Photo_stylization_10");
    proc_video_FarnebackOpticalFlow_setPyrScale_10 = procLoad("Java_org_opencv_video_FarnebackOpticalFlow_setPyrScale_10");
    proc_features2d_FastFeatureDetector_create_11 = procLoad("Java_org_opencv_features2d_FastFeatureDetector_create_11");
    proc_calib3d_Calib3d_undistortImage_11 = procLoad("Java_org_opencv_calib3d_Calib3d_undistortImage_11");
    proc_core_Core_add_11 = procLoad("Java_org_opencv_core_Core_add_11");
    proc_imgproc_Imgproc_demosaicing_11 = procLoad("Java_org_opencv_imgproc_Imgproc_demosaicing_11");
    proc_ml_LogisticRegression_delete = procLoad("Java_org_opencv_ml_LogisticRegression_delete");
    proc_photo_Photo_stylization_11 = procLoad("Java_org_opencv_photo_Photo_stylization_11");
    proc_video_FarnebackOpticalFlow_setWinSize_10 = procLoad("Java_org_opencv_video_FarnebackOpticalFlow_setWinSize_10");
    proc_ml_LogisticRegression_getIterations_10 = procLoad("Java_org_opencv_ml_LogisticRegression_getIterations_10");
    proc_video_KalmanFilter_KalmanFilter_10 = procLoad("Java_org_opencv_video_KalmanFilter_KalmanFilter_10");
    proc_calib3d_Calib3d_undistortPoints_10 = procLoad("Java_org_opencv_calib3d_Calib3d_undistortPoints_10");
    proc_core_Core_add_12 = procLoad("Java_org_opencv_core_Core_add_12");
    proc_features2d_FastFeatureDetector_delete = procLoad("Java_org_opencv_features2d_FastFeatureDetector_delete");
    proc_imgproc_Imgproc_dilate_10 = procLoad("Java_org_opencv_imgproc_Imgproc_dilate_10");
    proc_photo_Photo_textureFlattening_10 = procLoad("Java_org_opencv_photo_Photo_textureFlattening_10");
    proc_ml_LogisticRegression_getLearningRate_10 = procLoad("Java_org_opencv_ml_LogisticRegression_getLearningRate_10");
    proc_video_KalmanFilter_KalmanFilter_11 = procLoad("Java_org_opencv_video_KalmanFilter_KalmanFilter_11");
    proc_features2d_FastFeatureDetector_getDefaultName_10 = procLoad("Java_org_opencv_features2d_FastFeatureDetector_getDefaultName_10");
    proc_calib3d_Calib3d_undistortPoints_11 = procLoad("Java_org_opencv_calib3d_Calib3d_undistortPoints_11");
    proc_core_Core_add_13 = procLoad("Java_org_opencv_core_Core_add_13");
    proc_imgproc_Imgproc_dilate_11 = procLoad("Java_org_opencv_imgproc_Imgproc_dilate_11");
    proc_photo_Photo_textureFlattening_11 = procLoad("Java_org_opencv_photo_Photo_textureFlattening_11");
    proc_features2d_FastFeatureDetector_getNonmaxSuppression_10 = procLoad("Java_org_opencv_features2d_FastFeatureDetector_getNonmaxSuppression_10");
    proc_ml_LogisticRegression_getMiniBatchSize_10 = procLoad("Java_org_opencv_ml_LogisticRegression_getMiniBatchSize_10");
    proc_video_KalmanFilter_KalmanFilter_12 = procLoad("Java_org_opencv_video_KalmanFilter_KalmanFilter_12");
    proc_calib3d_Calib3d_validateDisparity_10 = procLoad("Java_org_opencv_calib3d_Calib3d_validateDisparity_10");
    proc_core_Core_add_14 = procLoad("Java_org_opencv_core_Core_add_14");
    proc_imgproc_Imgproc_dilate_12 = procLoad("Java_org_opencv_imgproc_Imgproc_dilate_12");
    proc_photo_Tonemap_delete = procLoad("Java_org_opencv_photo_Tonemap_delete");
    proc_photo_Tonemap_getGamma_10 = procLoad("Java_org_opencv_photo_Tonemap_getGamma_10");
    proc_features2d_FastFeatureDetector_getThreshold_10 = procLoad("Java_org_opencv_features2d_FastFeatureDetector_getThreshold_10");
    proc_ml_LogisticRegression_getRegularization_10 = procLoad("Java_org_opencv_ml_LogisticRegression_getRegularization_10");
    proc_video_KalmanFilter_correct_10 = procLoad("Java_org_opencv_video_KalmanFilter_correct_10");
    proc_calib3d_Calib3d_validateDisparity_11 = procLoad("Java_org_opencv_calib3d_Calib3d_validateDisparity_11");
    proc_core_Core_add_15 = procLoad("Java_org_opencv_core_Core_add_15");
    proc_imgproc_Imgproc_distanceTransformWithLabels_10 = procLoad("Java_org_opencv_imgproc_Imgproc_distanceTransformWithLabels_10");
    proc_ml_LogisticRegression_getTermCriteria_10 = procLoad("Java_org_opencv_ml_LogisticRegression_getTermCriteria_10");
    proc_features2d_FastFeatureDetector_getType_10 = procLoad("Java_org_opencv_features2d_FastFeatureDetector_getType_10");
    proc_calib3d_StereoBM_create_10 = procLoad("Java_org_opencv_calib3d_StereoBM_create_10");
    proc_core_Core_batchDistance_10 = procLoad("Java_org_opencv_core_Core_batchDistance_10");
    proc_imgproc_Imgproc_distanceTransformWithLabels_11 = procLoad("Java_org_opencv_imgproc_Imgproc_distanceTransformWithLabels_11");
    proc_photo_Tonemap_process_10 = procLoad("Java_org_opencv_photo_Tonemap_process_10");
    proc_video_KalmanFilter_delete = procLoad("Java_org_opencv_video_KalmanFilter_delete");
    proc_ml_LogisticRegression_getTrainMethod_10 = procLoad("Java_org_opencv_ml_LogisticRegression_getTrainMethod_10");
    proc_calib3d_StereoBM_create_11 = procLoad("Java_org_opencv_calib3d_StereoBM_create_11");
    proc_video_KalmanFilter_get_1controlMatrix_10 = procLoad("Java_org_opencv_video_KalmanFilter_get_1controlMatrix_10");
    proc_core_Core_batchDistance_11 = procLoad("Java_org_opencv_core_Core_batchDistance_11");
    proc_features2d_FastFeatureDetector_setNonmaxSuppression_10 = procLoad("Java_org_opencv_features2d_FastFeatureDetector_setNonmaxSuppression_10");
    proc_imgproc_Imgproc_distanceTransform_10 = procLoad("Java_org_opencv_imgproc_Imgproc_distanceTransform_10");
    proc_photo_Tonemap_setGamma_10 = procLoad("Java_org_opencv_photo_Tonemap_setGamma_10");
    proc_ml_LogisticRegression_get_1learnt_1thetas_10 = procLoad("Java_org_opencv_ml_LogisticRegression_get_1learnt_1thetas_10");
    proc_video_KalmanFilter_get_1errorCovPost_10 = procLoad("Java_org_opencv_video_KalmanFilter_get_1errorCovPost_10");
    proc_calib3d_StereoBM_delete = procLoad("Java_org_opencv_calib3d_StereoBM_delete");
    proc_core_Core_batchDistance_12 = procLoad("Java_org_opencv_core_Core_batchDistance_12");
    proc_features2d_FastFeatureDetector_setThreshold_10 = procLoad("Java_org_opencv_features2d_FastFeatureDetector_setThreshold_10");
    proc_imgproc_Imgproc_distanceTransform_11 = procLoad("Java_org_opencv_imgproc_Imgproc_distanceTransform_11");
    proc_photo_TonemapDrago_delete = procLoad("Java_org_opencv_photo_TonemapDrago_delete");
    proc_photo_TonemapDrago_getBias_10 = procLoad("Java_org_opencv_photo_TonemapDrago_getBias_10");
    proc_calib3d_StereoBM_getPreFilterCap_10 = procLoad("Java_org_opencv_calib3d_StereoBM_getPreFilterCap_10");
    proc_ml_LogisticRegression_load_10 = procLoad("Java_org_opencv_ml_LogisticRegression_load_10");
    proc_video_KalmanFilter_get_1errorCovPre_10 = procLoad("Java_org_opencv_video_KalmanFilter_get_1errorCovPre_10");
    proc_core_Core_bitwise_1and_10 = procLoad("Java_org_opencv_core_Core_bitwise_1and_10");
    proc_features2d_FastFeatureDetector_setType_10 = procLoad("Java_org_opencv_features2d_FastFeatureDetector_setType_10");
    proc_imgproc_Imgproc_drawContours_10 = procLoad("Java_org_opencv_imgproc_Imgproc_drawContours_10");
    proc_calib3d_Calib3d_calibrateCameraExtended_10 = procLoad("Java_org_opencv_calib3d_Calib3d_calibrateCameraExtended_10");
    proc_imgproc_CLAHE_getClipLimit_10 = procLoad("Java_org_opencv_imgproc_CLAHE_getClipLimit_10");
    proc_ml_ANN_1MLP_getBackpropWeightScale_10 = procLoad("Java_org_opencv_ml_ANN_1MLP_getBackpropWeightScale_10");
    proc_core_Mat_nGetF = procLoad("Java_org_opencv_core_Mat_nGetF");
    proc_imgcodecs_Imgcodecs_imread_10 = procLoad("Java_org_opencv_imgcodecs_Imgcodecs_imread_10");
    proc_objdetect_HOGDescriptor_HOGDescriptor_11 = procLoad("Java_org_opencv_objdetect_HOGDescriptor_HOGDescriptor_11");
    proc_videoio_VideoCapture_VideoCapture_13 = procLoad("Java_org_opencv_videoio_VideoCapture_VideoCapture_13");
    proc_features2d_AKAZE_getDefaultName_10 = procLoad("Java_org_opencv_features2d_AKAZE_getDefaultName_10");
    proc_dnn_DictValue_delete = procLoad("Java_org_opencv_dnn_DictValue_delete");
    proc_photo_AlignMTB_computeBitmaps_10 = procLoad("Java_org_opencv_photo_AlignMTB_computeBitmaps_10");
    proc_video_BackgroundSubtractor_getBackgroundImage_10 = procLoad("Java_org_opencv_video_BackgroundSubtractor_getBackgroundImage_10");
    proc_photo_TonemapDrago_getSaturation_10 = procLoad("Java_org_opencv_photo_TonemapDrago_getSaturation_10");
    proc_calib3d_StereoBM_getPreFilterSize_10 = procLoad("Java_org_opencv_calib3d_StereoBM_getPreFilterSize_10");
    proc_ml_LogisticRegression_load_11 = procLoad("Java_org_opencv_ml_LogisticRegression_load_11");
    proc_video_KalmanFilter_get_1gain_10 = procLoad("Java_org_opencv_video_KalmanFilter_get_1gain_10");
    proc_core_Core_bitwise_1and_11 = procLoad("Java_org_opencv_core_Core_bitwise_1and_11");
    proc_features2d_Feature2D_compute_10 = procLoad("Java_org_opencv_features2d_Feature2D_compute_10");
    proc_imgproc_Imgproc_drawContours_11 = procLoad("Java_org_opencv_imgproc_Imgproc_drawContours_11");
    proc_ml_LogisticRegression_predict_10 = procLoad("Java_org_opencv_ml_LogisticRegression_predict_10");
    proc_calib3d_StereoBM_getPreFilterType_10 = procLoad("Java_org_opencv_calib3d_StereoBM_getPreFilterType_10");
    proc_video_KalmanFilter_get_1measurementMatrix_10 = procLoad("Java_org_opencv_video_KalmanFilter_get_1measurementMatrix_10");
    proc_core_Core_bitwise_1not_10 = procLoad("Java_org_opencv_core_Core_bitwise_1not_10");
    proc_features2d_Feature2D_compute_11 = procLoad("Java_org_opencv_features2d_Feature2D_compute_11");
    proc_imgproc_Imgproc_drawContours_12 = procLoad("Java_org_opencv_imgproc_Imgproc_drawContours_12");
    proc_photo_TonemapDrago_setBias_10 = procLoad("Java_org_opencv_photo_TonemapDrago_setBias_10");
    proc_calib3d_StereoBM_getROI1_10 = procLoad("Java_org_opencv_calib3d_StereoBM_getROI1_10");
    proc_ml_LogisticRegression_predict_11 = procLoad("Java_org_opencv_ml_LogisticRegression_predict_11");
    proc_features2d_Feature2D_defaultNorm_10 = procLoad("Java_org_opencv_features2d_Feature2D_defaultNorm_10");
    proc_video_KalmanFilter_get_1measurementNoiseCov_10 = procLoad("Java_org_opencv_video_KalmanFilter_get_1measurementNoiseCov_10");
    proc_core_Core_bitwise_1not_11 = procLoad("Java_org_opencv_core_Core_bitwise_1not_11");
    proc_imgproc_Imgproc_drawMarker_10 = procLoad("Java_org_opencv_imgproc_Imgproc_drawMarker_10");
    proc_photo_TonemapDrago_setSaturation_10 = procLoad("Java_org_opencv_photo_TonemapDrago_setSaturation_10");
    proc_calib3d_StereoBM_getROI2_10 = procLoad("Java_org_opencv_calib3d_StereoBM_getROI2_10");
    proc_video_KalmanFilter_get_1processNoiseCov_10 = procLoad("Java_org_opencv_video_KalmanFilter_get_1processNoiseCov_10");
    proc_core_Core_bitwise_1or_10 = procLoad("Java_org_opencv_core_Core_bitwise_1or_10");
    proc_features2d_Feature2D_delete = procLoad("Java_org_opencv_features2d_Feature2D_delete");
    proc_imgproc_Imgproc_drawMarker_11 = procLoad("Java_org_opencv_imgproc_Imgproc_drawMarker_11");
    proc_ml_LogisticRegression_setIterations_10 = procLoad("Java_org_opencv_ml_LogisticRegression_setIterations_10");
    proc_photo_TonemapDurand_delete = procLoad("Java_org_opencv_photo_TonemapDurand_delete");
    proc_photo_TonemapDurand_getContrast_10 = procLoad("Java_org_opencv_photo_TonemapDurand_getContrast_10");
    proc_calib3d_StereoBM_getSmallerBlockSize_10 = procLoad("Java_org_opencv_calib3d_StereoBM_getSmallerBlockSize_10");
    proc_features2d_Feature2D_descriptorSize_10 = procLoad("Java_org_opencv_features2d_Feature2D_descriptorSize_10");
    proc_video_KalmanFilter_get_1statePost_10 = procLoad("Java_org_opencv_video_KalmanFilter_get_1statePost_10");
    proc_core_Core_bitwise_1or_11 = procLoad("Java_org_opencv_core_Core_bitwise_1or_11");
    proc_imgproc_Imgproc_ellipse2Poly_10 = procLoad("Java_org_opencv_imgproc_Imgproc_ellipse2Poly_10");
    proc_ml_LogisticRegression_setLearningRate_10 = procLoad("Java_org_opencv_ml_LogisticRegression_setLearningRate_10");
    proc_photo_TonemapDurand_getSaturation_10 = procLoad("Java_org_opencv_photo_TonemapDurand_getSaturation_10");
    proc_calib3d_StereoBM_getTextureThreshold_10 = procLoad("Java_org_opencv_calib3d_StereoBM_getTextureThreshold_10");
    proc_features2d_Feature2D_descriptorType_10 = procLoad("Java_org_opencv_features2d_Feature2D_descriptorType_10");
    proc_video_KalmanFilter_get_1statePre_10 = procLoad("Java_org_opencv_video_KalmanFilter_get_1statePre_10");
    proc_core_Core_bitwise_1xor_10 = procLoad("Java_org_opencv_core_Core_bitwise_1xor_10");
    proc_imgproc_Imgproc_ellipse_10 = procLoad("Java_org_opencv_imgproc_Imgproc_ellipse_10");
    proc_ml_LogisticRegression_setMiniBatchSize_10 = procLoad("Java_org_opencv_ml_LogisticRegression_setMiniBatchSize_10");
    proc_photo_TonemapDurand_getSigmaColor_10 = procLoad("Java_org_opencv_photo_TonemapDurand_getSigmaColor_10");
    proc_calib3d_StereoBM_getUniquenessRatio_10 = procLoad("Java_org_opencv_calib3d_StereoBM_getUniquenessRatio_10");
    proc_video_KalmanFilter_get_1transitionMatrix_10 = procLoad("Java_org_opencv_video_KalmanFilter_get_1transitionMatrix_10");
    proc_core_Core_bitwise_1xor_11 = procLoad("Java_org_opencv_core_Core_bitwise_1xor_11");
    proc_features2d_Feature2D_detectAndCompute_10 = procLoad("Java_org_opencv_features2d_Feature2D_detectAndCompute_10");
    proc_imgproc_Imgproc_ellipse_11 = procLoad("Java_org_opencv_imgproc_Imgproc_ellipse_11");
    proc_ml_LogisticRegression_setRegularization_10 = procLoad("Java_org_opencv_ml_LogisticRegression_setRegularization_10");
    proc_photo_TonemapDurand_getSigmaSpace_10 = procLoad("Java_org_opencv_photo_TonemapDurand_getSigmaSpace_10");
    proc_core_Core_borderInterpolate_10 = procLoad("Java_org_opencv_core_Core_borderInterpolate_10");
    proc_video_KalmanFilter_predict_10 = procLoad("Java_org_opencv_video_KalmanFilter_predict_10");
    proc_calib3d_StereoBM_setPreFilterCap_10 = procLoad("Java_org_opencv_calib3d_StereoBM_setPreFilterCap_10");
    proc_features2d_Feature2D_detectAndCompute_11 = procLoad("Java_org_opencv_features2d_Feature2D_detectAndCompute_11");
    proc_imgproc_Imgproc_ellipse_12 = procLoad("Java_org_opencv_imgproc_Imgproc_ellipse_12");
    proc_ml_LogisticRegression_setTermCriteria_10 = procLoad("Java_org_opencv_ml_LogisticRegression_setTermCriteria_10");
    proc_video_KalmanFilter_predict_11 = procLoad("Java_org_opencv_video_KalmanFilter_predict_11");
    proc_calib3d_StereoBM_setPreFilterSize_10 = procLoad("Java_org_opencv_calib3d_StereoBM_setPreFilterSize_10");
    proc_core_Core_calcCovarMatrix_10 = procLoad("Java_org_opencv_core_Core_calcCovarMatrix_10");
    proc_features2d_Feature2D_detect_10 = procLoad("Java_org_opencv_features2d_Feature2D_detect_10");
    proc_imgproc_Imgproc_ellipse_13 = procLoad("Java_org_opencv_imgproc_Imgproc_ellipse_13");
    proc_ml_LogisticRegression_setTrainMethod_10 = procLoad("Java_org_opencv_ml_LogisticRegression_setTrainMethod_10");
    proc_photo_TonemapDurand_setContrast_10 = procLoad("Java_org_opencv_photo_TonemapDurand_setContrast_10");
    proc_ml_NormalBayesClassifier_create_10 = procLoad("Java_org_opencv_ml_NormalBayesClassifier_create_10");
    proc_calib3d_StereoBM_setPreFilterType_10 = procLoad("Java_org_opencv_calib3d_StereoBM_setPreFilterType_10");
    proc_core_Core_calcCovarMatrix_11 = procLoad("Java_org_opencv_core_Core_calcCovarMatrix_11");
    proc_features2d_Feature2D_detect_11 = procLoad("Java_org_opencv_features2d_Feature2D_detect_11");
    proc_imgproc_Imgproc_ellipse_14 = procLoad("Java_org_opencv_imgproc_Imgproc_ellipse_14");
    proc_photo_TonemapDurand_setSaturation_10 = procLoad("Java_org_opencv_photo_TonemapDurand_setSaturation_10");
    proc_video_KalmanFilter_set_1controlMatrix_10 = procLoad("Java_org_opencv_video_KalmanFilter_set_1controlMatrix_10");
    proc_calib3d_Calib3d_calibrateCameraExtended_11 = procLoad("Java_org_opencv_calib3d_Calib3d_calibrateCameraExtended_11");
    proc_imgproc_CLAHE_getTilesGridSize_10 = procLoad("Java_org_opencv_imgproc_CLAHE_getTilesGridSize_10");
    proc_core_Mat_nGetI = procLoad("Java_org_opencv_core_Mat_nGetI");
    proc_dnn_DictValue_getIntValue_10 = procLoad("Java_org_opencv_dnn_DictValue_getIntValue_10");
    proc_features2d_AKAZE_getDescriptorChannels_10 = procLoad("Java_org_opencv_features2d_AKAZE_getDescriptorChannels_10");
    proc_imgcodecs_Imgcodecs_imread_11 = procLoad("Java_org_opencv_imgcodecs_Imgcodecs_imread_11");
    proc_ml_ANN_1MLP_getLayerSizes_10 = procLoad("Java_org_opencv_ml_ANN_1MLP_getLayerSizes_10");
    proc_objdetect_HOGDescriptor_HOGDescriptor_12 = procLoad("Java_org_opencv_objdetect_HOGDescriptor_HOGDescriptor_12");
    proc_photo_AlignMTB_delete = procLoad("Java_org_opencv_photo_AlignMTB_delete");
    proc_video_BackgroundSubtractorKNN_delete = procLoad("Java_org_opencv_video_BackgroundSubtractorKNN_delete");
    proc_videoio_VideoCapture_delete = procLoad("Java_org_opencv_videoio_VideoCapture_delete");
    proc_calib3d_StereoBM_setROI1_10 = procLoad("Java_org_opencv_calib3d_StereoBM_setROI1_10");
    proc_core_Core_cartToPolar_10 = procLoad("Java_org_opencv_core_Core_cartToPolar_10");
    proc_features2d_Feature2D_detect_12 = procLoad("Java_org_opencv_features2d_Feature2D_detect_12");
    proc_imgproc_Imgproc_ellipse_15 = procLoad("Java_org_opencv_imgproc_Imgproc_ellipse_15");
    proc_ml_NormalBayesClassifier_delete = procLoad("Java_org_opencv_ml_NormalBayesClassifier_delete");
    proc_photo_TonemapDurand_setSigmaColor_10 = procLoad("Java_org_opencv_photo_TonemapDurand_setSigmaColor_10");
    proc_video_KalmanFilter_set_1errorCovPost_10 = procLoad("Java_org_opencv_video_KalmanFilter_set_1errorCovPost_10");
    proc_ml_NormalBayesClassifier_load_10 = procLoad("Java_org_opencv_ml_NormalBayesClassifier_load_10");
    proc_calib3d_StereoBM_setROI2_10 = procLoad("Java_org_opencv_calib3d_StereoBM_setROI2_10");
    proc_core_Core_cartToPolar_11 = procLoad("Java_org_opencv_core_Core_cartToPolar_11");
    proc_features2d_Feature2D_detect_13 = procLoad("Java_org_opencv_features2d_Feature2D_detect_13");
    proc_imgproc_Imgproc_equalizeHist_10 = procLoad("Java_org_opencv_imgproc_Imgproc_equalizeHist_10");
    proc_photo_TonemapDurand_setSigmaSpace_10 = procLoad("Java_org_opencv_photo_TonemapDurand_setSigmaSpace_10");
    proc_video_KalmanFilter_set_1errorCovPre_10 = procLoad("Java_org_opencv_video_KalmanFilter_set_1errorCovPre_10");
    proc_core_Core_checkRange_10 = procLoad("Java_org_opencv_core_Core_checkRange_10");
    proc_features2d_Feature2D_empty_10 = procLoad("Java_org_opencv_features2d_Feature2D_empty_10");
    proc_ml_NormalBayesClassifier_load_11 = procLoad("Java_org_opencv_ml_NormalBayesClassifier_load_11");
    proc_calib3d_StereoBM_setSmallerBlockSize_10 = procLoad("Java_org_opencv_calib3d_StereoBM_setSmallerBlockSize_10");
    proc_imgproc_Imgproc_erode_10 = procLoad("Java_org_opencv_imgproc_Imgproc_erode_10");
    proc_photo_TonemapMantiuk_delete = procLoad("Java_org_opencv_photo_TonemapMantiuk_delete");
    proc_video_KalmanFilter_set_1gain_10 = procLoad("Java_org_opencv_video_KalmanFilter_set_1gain_10");
    proc_core_Core_checkRange_11 = procLoad("Java_org_opencv_core_Core_checkRange_11");
    proc_ml_NormalBayesClassifier_predictProb_10 = procLoad("Java_org_opencv_ml_NormalBayesClassifier_predictProb_10");
    proc_photo_TonemapMantiuk_getSaturation_10 = procLoad("Java_org_opencv_photo_TonemapMantiuk_getSaturation_10");
    proc_features2d_Feature2D_getDefaultName_10 = procLoad("Java_org_opencv_features2d_Feature2D_getDefaultName_10");
    proc_calib3d_StereoBM_setTextureThreshold_10 = procLoad("Java_org_opencv_calib3d_StereoBM_setTextureThreshold_10");
    proc_imgproc_Imgproc_erode_11 = procLoad("Java_org_opencv_imgproc_Imgproc_erode_11");
    proc_video_KalmanFilter_set_1measurementMatrix_10 = procLoad("Java_org_opencv_video_KalmanFilter_set_1measurementMatrix_10");
    proc_ml_NormalBayesClassifier_predictProb_11 = procLoad("Java_org_opencv_ml_NormalBayesClassifier_predictProb_11");
    proc_photo_TonemapMantiuk_getScale_10 = procLoad("Java_org_opencv_photo_TonemapMantiuk_getScale_10");
    proc_calib3d_StereoBM_setUniquenessRatio_10 = procLoad("Java_org_opencv_calib3d_StereoBM_setUniquenessRatio_10");
    proc_core_Core_compare_10 = procLoad("Java_org_opencv_core_Core_compare_10");
    proc_features2d_Feature2D_read_10 = procLoad("Java_org_opencv_features2d_Feature2D_read_10");
    proc_imgproc_Imgproc_erode_12 = procLoad("Java_org_opencv_imgproc_Imgproc_erode_12");
    proc_video_KalmanFilter_set_1measurementNoiseCov_10 = procLoad("Java_org_opencv_video_KalmanFilter_set_1measurementNoiseCov_10");
    proc_ml_ParamGrid_create_10 = procLoad("Java_org_opencv_ml_ParamGrid_create_10");
    proc_calib3d_StereoMatcher_compute_10 = procLoad("Java_org_opencv_calib3d_StereoMatcher_compute_10");
    proc_core_Core_compare_11 = procLoad("Java_org_opencv_core_Core_compare_11");
    proc_features2d_Feature2D_write_10 = procLoad("Java_org_opencv_features2d_Feature2D_write_10");
    proc_imgproc_Imgproc_fillConvexPoly_10 = procLoad("Java_org_opencv_imgproc_Imgproc_fillConvexPoly_10");
    proc_photo_TonemapMantiuk_setSaturation_10 = procLoad("Java_org_opencv_photo_TonemapMantiuk_setSaturation_10");
    proc_video_KalmanFilter_set_1processNoiseCov_10 = procLoad("Java_org_opencv_video_KalmanFilter_set_1processNoiseCov_10");
    proc_features2d_FeatureDetector_create_10 = procLoad("Java_org_opencv_features2d_FeatureDetector_create_10");
    proc_ml_ParamGrid_create_11 = procLoad("Java_org_opencv_ml_ParamGrid_create_11");
    proc_calib3d_StereoMatcher_delete = procLoad("Java_org_opencv_calib3d_StereoMatcher_delete");
    proc_core_Core_completeSymm_10 = procLoad("Java_org_opencv_core_Core_completeSymm_10");
    proc_imgproc_Imgproc_fillConvexPoly_11 = procLoad("Java_org_opencv_imgproc_Imgproc_fillConvexPoly_11");
    proc_photo_TonemapMantiuk_setScale_10 = procLoad("Java_org_opencv_photo_TonemapMantiuk_setScale_10");
    proc_video_KalmanFilter_set_1statePost_10 = procLoad("Java_org_opencv_video_KalmanFilter_set_1statePost_10");
    proc_calib3d_StereoMatcher_getBlockSize_10 = procLoad("Java_org_opencv_calib3d_StereoMatcher_getBlockSize_10");
    proc_core_Core_completeSymm_11 = procLoad("Java_org_opencv_core_Core_completeSymm_11");
    proc_features2d_FeatureDetector_delete = procLoad("Java_org_opencv_features2d_FeatureDetector_delete");
    proc_imgproc_Imgproc_fillPoly_10 = procLoad("Java_org_opencv_imgproc_Imgproc_fillPoly_10");
    proc_ml_ParamGrid_delete = procLoad("Java_org_opencv_ml_ParamGrid_delete");
    proc_photo_TonemapReinhard_delete = procLoad("Java_org_opencv_photo_TonemapReinhard_delete");
    proc_video_KalmanFilter_set_1statePre_10 = procLoad("Java_org_opencv_video_KalmanFilter_set_1statePre_10");
    proc_ml_ParamGrid_get_1logStep_10 = procLoad("Java_org_opencv_ml_ParamGrid_get_1logStep_10");
    proc_photo_TonemapReinhard_getColorAdaptation_10 = procLoad("Java_org_opencv_photo_TonemapReinhard_getColorAdaptation_10");
    proc_calib3d_StereoMatcher_getDisp12MaxDiff_10 = procLoad("Java_org_opencv_calib3d_StereoMatcher_getDisp12MaxDiff_10");
    proc_core_Core_convertFp16_10 = procLoad("Java_org_opencv_core_Core_convertFp16_10");
    proc_features2d_FeatureDetector_detect_10 = procLoad("Java_org_opencv_features2d_FeatureDetector_detect_10");
    proc_imgproc_Imgproc_fillPoly_11 = procLoad("Java_org_opencv_imgproc_Imgproc_fillPoly_11");
    proc_video_KalmanFilter_set_1transitionMatrix_10 = procLoad("Java_org_opencv_video_KalmanFilter_set_1transitionMatrix_10");
    proc_ml_ParamGrid_get_1maxVal_10 = procLoad("Java_org_opencv_ml_ParamGrid_get_1maxVal_10");
    proc_photo_TonemapReinhard_getIntensity_10 = procLoad("Java_org_opencv_photo_TonemapReinhard_getIntensity_10");
    proc_calib3d_StereoMatcher_getMinDisparity_10 = procLoad("Java_org_opencv_calib3d_StereoMatcher_getMinDisparity_10");
    proc_core_Core_convertScaleAbs_10 = procLoad("Java_org_opencv_core_Core_convertScaleAbs_10");
    proc_features2d_FeatureDetector_detect_11 = procLoad("Java_org_opencv_features2d_FeatureDetector_detect_11");
    proc_imgproc_Imgproc_filter2D_10 = procLoad("Java_org_opencv_imgproc_Imgproc_filter2D_10");
    proc_video_SparseOpticalFlow_calc_10 = procLoad("Java_org_opencv_video_SparseOpticalFlow_calc_10");
    proc_imgcodecs_Imgcodecs_imreadmulti_10 = procLoad("Java_org_opencv_imgcodecs_Imgcodecs_imreadmulti_10");
    proc_photo_AlignMTB_getCut_10 = procLoad("Java_org_opencv_photo_AlignMTB_getCut_10");
    proc_video_BackgroundSubtractorKNN_getDetectShadows_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorKNN_getDetectShadows_10");
    proc_calib3d_Calib3d_calibrateCameraExtended_12 = procLoad("Java_org_opencv_calib3d_Calib3d_calibrateCameraExtended_12");
    proc_ml_ANN_1MLP_getRpropDW0_10 = procLoad("Java_org_opencv_ml_ANN_1MLP_getRpropDW0_10");
    proc_videoio_VideoCapture_get_10 = procLoad("Java_org_opencv_videoio_VideoCapture_get_10");
    proc_core_Mat_nGetS = procLoad("Java_org_opencv_core_Mat_nGetS");
    proc_dnn_DictValue_getIntValue_11 = procLoad("Java_org_opencv_dnn_DictValue_getIntValue_11");
    proc_features2d_AKAZE_getDescriptorSize_10 = procLoad("Java_org_opencv_features2d_AKAZE_getDescriptorSize_10");
    proc_objdetect_HOGDescriptor_HOGDescriptor_13 = procLoad("Java_org_opencv_objdetect_HOGDescriptor_HOGDescriptor_13");
    proc_imgproc_CLAHE_setClipLimit_10 = procLoad("Java_org_opencv_imgproc_CLAHE_setClipLimit_10");
    proc_ml_ParamGrid_get_1minVal_10 = procLoad("Java_org_opencv_ml_ParamGrid_get_1minVal_10");
    proc_photo_TonemapReinhard_getLightAdaptation_10 = procLoad("Java_org_opencv_photo_TonemapReinhard_getLightAdaptation_10");
    proc_calib3d_StereoMatcher_getNumDisparities_10 = procLoad("Java_org_opencv_calib3d_StereoMatcher_getNumDisparities_10");
    proc_core_Core_convertScaleAbs_11 = procLoad("Java_org_opencv_core_Core_convertScaleAbs_11");
    proc_features2d_FeatureDetector_detect_12 = procLoad("Java_org_opencv_features2d_FeatureDetector_detect_12");
    proc_imgproc_Imgproc_filter2D_11 = procLoad("Java_org_opencv_imgproc_Imgproc_filter2D_11");
    proc_video_SparseOpticalFlow_calc_11 = procLoad("Java_org_opencv_video_SparseOpticalFlow_calc_11");
    proc_calib3d_StereoMatcher_getSpeckleRange_10 = procLoad("Java_org_opencv_calib3d_StereoMatcher_getSpeckleRange_10");
    proc_core_Core_copyMakeBorder_10 = procLoad("Java_org_opencv_core_Core_copyMakeBorder_10");
    proc_features2d_FeatureDetector_detect_13 = procLoad("Java_org_opencv_features2d_FeatureDetector_detect_13");
    proc_imgproc_Imgproc_filter2D_12 = procLoad("Java_org_opencv_imgproc_Imgproc_filter2D_12");
    proc_ml_ParamGrid_set_1logStep_10 = procLoad("Java_org_opencv_ml_ParamGrid_set_1logStep_10");
    proc_photo_TonemapReinhard_setColorAdaptation_10 = procLoad("Java_org_opencv_photo_TonemapReinhard_setColorAdaptation_10");
    proc_video_SparseOpticalFlow_delete = procLoad("Java_org_opencv_video_SparseOpticalFlow_delete");
    proc_features2d_FeatureDetector_empty_10 = procLoad("Java_org_opencv_features2d_FeatureDetector_empty_10");
    proc_calib3d_StereoMatcher_getSpeckleWindowSize_10 = procLoad("Java_org_opencv_calib3d_StereoMatcher_getSpeckleWindowSize_10");
    proc_video_SparsePyrLKOpticalFlow_create_10 = procLoad("Java_org_opencv_video_SparsePyrLKOpticalFlow_create_10");
    proc_core_Core_copyMakeBorder_11 = procLoad("Java_org_opencv_core_Core_copyMakeBorder_11");
    proc_imgproc_Imgproc_findContours_10 = procLoad("Java_org_opencv_imgproc_Imgproc_findContours_10");
    proc_ml_ParamGrid_set_1maxVal_10 = procLoad("Java_org_opencv_ml_ParamGrid_set_1maxVal_10");
    proc_photo_TonemapReinhard_setIntensity_10 = procLoad("Java_org_opencv_photo_TonemapReinhard_setIntensity_10");
    proc_core_Core_countNonZero_10 = procLoad("Java_org_opencv_core_Core_countNonZero_10");
    proc_video_SparsePyrLKOpticalFlow_create_11 = procLoad("Java_org_opencv_video_SparsePyrLKOpticalFlow_create_11");
    proc_calib3d_StereoMatcher_setBlockSize_10 = procLoad("Java_org_opencv_calib3d_StereoMatcher_setBlockSize_10");
    proc_features2d_FeatureDetector_read_10 = procLoad("Java_org_opencv_features2d_FeatureDetector_read_10");
    proc_imgproc_Imgproc_findContours_11 = procLoad("Java_org_opencv_imgproc_Imgproc_findContours_11");
    proc_ml_ParamGrid_set_1minVal_10 = procLoad("Java_org_opencv_ml_ParamGrid_set_1minVal_10");
    proc_photo_TonemapReinhard_setLightAdaptation_10 = procLoad("Java_org_opencv_photo_TonemapReinhard_setLightAdaptation_10");
    proc_imgproc_Imgproc_fitEllipseAMS_10 = procLoad("Java_org_opencv_imgproc_Imgproc_fitEllipseAMS_10");
    proc_core_Core_cubeRoot_10 = procLoad("Java_org_opencv_core_Core_cubeRoot_10");
    proc_ml_RTrees_create_10 = procLoad("Java_org_opencv_ml_RTrees_create_10");
    proc_calib3d_StereoMatcher_setDisp12MaxDiff_10 = procLoad("Java_org_opencv_calib3d_StereoMatcher_setDisp12MaxDiff_10");
    proc_features2d_FeatureDetector_write_10 = procLoad("Java_org_opencv_features2d_FeatureDetector_write_10");
    proc_video_SparsePyrLKOpticalFlow_delete = procLoad("Java_org_opencv_video_SparsePyrLKOpticalFlow_delete");
    proc_imgproc_Imgproc_fitEllipseDirect_10 = procLoad("Java_org_opencv_imgproc_Imgproc_fitEllipseDirect_10");
    proc_video_SparsePyrLKOpticalFlow_getFlags_10 = procLoad("Java_org_opencv_video_SparsePyrLKOpticalFlow_getFlags_10");
    proc_calib3d_StereoMatcher_setMinDisparity_10 = procLoad("Java_org_opencv_calib3d_StereoMatcher_setMinDisparity_10");
    proc_core_Core_dct_10 = procLoad("Java_org_opencv_core_Core_dct_10");
    proc_features2d_Features2d_drawKeypoints_10 = procLoad("Java_org_opencv_features2d_Features2d_drawKeypoints_10");
    proc_ml_RTrees_delete = procLoad("Java_org_opencv_ml_RTrees_delete");
    proc_imgproc_Imgproc_fitEllipse_10 = procLoad("Java_org_opencv_imgproc_Imgproc_fitEllipse_10");
    proc_ml_RTrees_getActiveVarCount_10 = procLoad("Java_org_opencv_ml_RTrees_getActiveVarCount_10");
    proc_video_SparsePyrLKOpticalFlow_getMaxLevel_10 = procLoad("Java_org_opencv_video_SparsePyrLKOpticalFlow_getMaxLevel_10");
    proc_calib3d_StereoMatcher_setNumDisparities_10 = procLoad("Java_org_opencv_calib3d_StereoMatcher_setNumDisparities_10");
    proc_core_Core_dct_11 = procLoad("Java_org_opencv_core_Core_dct_11");
    proc_features2d_Features2d_drawKeypoints_11 = procLoad("Java_org_opencv_features2d_Features2d_drawKeypoints_11");
    proc_ml_RTrees_getCalculateVarImportance_10 = procLoad("Java_org_opencv_ml_RTrees_getCalculateVarImportance_10");
    proc_core_Core_determinant_10 = procLoad("Java_org_opencv_core_Core_determinant_10");
    proc_video_SparsePyrLKOpticalFlow_getMinEigThreshold_10 = procLoad("Java_org_opencv_video_SparsePyrLKOpticalFlow_getMinEigThreshold_10");
    proc_calib3d_StereoMatcher_setSpeckleRange_10 = procLoad("Java_org_opencv_calib3d_StereoMatcher_setSpeckleRange_10");
    proc_features2d_Features2d_drawMatches2_10 = procLoad("Java_org_opencv_features2d_Features2d_drawMatches2_10");
    proc_imgproc_Imgproc_fitLine_10 = procLoad("Java_org_opencv_imgproc_Imgproc_fitLine_10");
    proc_ml_RTrees_getTermCriteria_10 = procLoad("Java_org_opencv_ml_RTrees_getTermCriteria_10");
    proc_video_SparsePyrLKOpticalFlow_getTermCriteria_10 = procLoad("Java_org_opencv_video_SparsePyrLKOpticalFlow_getTermCriteria_10");
    proc_imgproc_Imgproc_floodFill_10 = procLoad("Java_org_opencv_imgproc_Imgproc_floodFill_10");
    proc_calib3d_StereoMatcher_setSpeckleWindowSize_10 = procLoad("Java_org_opencv_calib3d_StereoMatcher_setSpeckleWindowSize_10");
    proc_core_Core_dft_10 = procLoad("Java_org_opencv_core_Core_dft_10");
    proc_features2d_Features2d_drawMatches2_11 = procLoad("Java_org_opencv_features2d_Features2d_drawMatches2_11");
    proc_video_SparsePyrLKOpticalFlow_getWinSize_10 = procLoad("Java_org_opencv_video_SparsePyrLKOpticalFlow_getWinSize_10");
    proc_imgproc_Imgproc_floodFill_11 = procLoad("Java_org_opencv_imgproc_Imgproc_floodFill_11");
    proc_calib3d_StereoSGBM_create_10 = procLoad("Java_org_opencv_calib3d_StereoSGBM_create_10");
    proc_ml_RTrees_getVarImportance_10 = procLoad("Java_org_opencv_ml_RTrees_getVarImportance_10");
    proc_core_Core_dft_11 = procLoad("Java_org_opencv_core_Core_dft_11");
    proc_features2d_Features2d_drawMatchesKnn_10 = procLoad("Java_org_opencv_features2d_Features2d_drawMatchesKnn_10");
    proc_imgcodecs_Imgcodecs_imreadmulti_11 = procLoad("Java_org_opencv_imgcodecs_Imgcodecs_imreadmulti_11");
    proc_objdetect_HOGDescriptor_checkDetectorSize_10 = procLoad("Java_org_opencv_objdetect_HOGDescriptor_checkDetectorSize_10");
    proc_videoio_VideoCapture_grab_10 = procLoad("Java_org_opencv_videoio_VideoCapture_grab_10");
    proc_calib3d_Calib3d_calibrateCamera_10 = procLoad("Java_org_opencv_calib3d_Calib3d_calibrateCamera_10");
    proc_dnn_DictValue_getRealValue_10 = procLoad("Java_org_opencv_dnn_DictValue_getRealValue_10");
    proc_ml_ANN_1MLP_getRpropDWMax_10 = procLoad("Java_org_opencv_ml_ANN_1MLP_getRpropDWMax_10");
    proc_video_BackgroundSubtractorKNN_getDist2Threshold_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorKNN_getDist2Threshold_10");
    proc_core_Mat_nPutB = procLoad("Java_org_opencv_core_Mat_nPutB");
    proc_features2d_AKAZE_getDescriptorType_10 = procLoad("Java_org_opencv_features2d_AKAZE_getDescriptorType_10");
    proc_photo_AlignMTB_getExcludeRange_10 = procLoad("Java_org_opencv_photo_AlignMTB_getExcludeRange_10");
    proc_imgproc_CLAHE_setTilesGridSize_10 = procLoad("Java_org_opencv_imgproc_CLAHE_setTilesGridSize_10");
    proc_calib3d_StereoSGBM_create_11 = procLoad("Java_org_opencv_calib3d_StereoSGBM_create_11");
    proc_imgproc_Imgproc_getAffineTransform_10 = procLoad("Java_org_opencv_imgproc_Imgproc_getAffineTransform_10");
    proc_core_Core_divide_10 = procLoad("Java_org_opencv_core_Core_divide_10");
    proc_features2d_Features2d_drawMatchesKnn_11 = procLoad("Java_org_opencv_features2d_Features2d_drawMatchesKnn_11");
    proc_ml_RTrees_getVotes_10 = procLoad("Java_org_opencv_ml_RTrees_getVotes_10");
    proc_video_SparsePyrLKOpticalFlow_setFlags_10 = procLoad("Java_org_opencv_video_SparsePyrLKOpticalFlow_setFlags_10");
    proc_imgproc_Imgproc_getDefaultNewCameraMatrix_10 = procLoad("Java_org_opencv_imgproc_Imgproc_getDefaultNewCameraMatrix_10");
    proc_ml_RTrees_load_10 = procLoad("Java_org_opencv_ml_RTrees_load_10");
    proc_calib3d_StereoSGBM_delete = procLoad("Java_org_opencv_calib3d_StereoSGBM_delete");
    proc_core_Core_divide_11 = procLoad("Java_org_opencv_core_Core_divide_11");
    proc_features2d_Features2d_drawMatches_10 = procLoad("Java_org_opencv_features2d_Features2d_drawMatches_10");
    proc_video_SparsePyrLKOpticalFlow_setMaxLevel_10 = procLoad("Java_org_opencv_video_SparsePyrLKOpticalFlow_setMaxLevel_10");
    proc_calib3d_StereoSGBM_getMode_10 = procLoad("Java_org_opencv_calib3d_StereoSGBM_getMode_10");
    proc_imgproc_Imgproc_getDefaultNewCameraMatrix_11 = procLoad("Java_org_opencv_imgproc_Imgproc_getDefaultNewCameraMatrix_11");
    proc_ml_RTrees_load_11 = procLoad("Java_org_opencv_ml_RTrees_load_11");
    proc_core_Core_divide_12 = procLoad("Java_org_opencv_core_Core_divide_12");
    proc_features2d_Features2d_drawMatches_11 = procLoad("Java_org_opencv_features2d_Features2d_drawMatches_11");
    proc_video_SparsePyrLKOpticalFlow_setMinEigThreshold_10 = procLoad("Java_org_opencv_video_SparsePyrLKOpticalFlow_setMinEigThreshold_10");
    proc_calib3d_StereoSGBM_getP1_10 = procLoad("Java_org_opencv_calib3d_StereoSGBM_getP1_10");
    proc_features2d_FlannBasedMatcher_FlannBasedMatcher_10 = procLoad("Java_org_opencv_features2d_FlannBasedMatcher_FlannBasedMatcher_10");
    proc_core_Core_divide_13 = procLoad("Java_org_opencv_core_Core_divide_13");
    proc_imgproc_Imgproc_getDerivKernels_10 = procLoad("Java_org_opencv_imgproc_Imgproc_getDerivKernels_10");
    proc_ml_RTrees_setActiveVarCount_10 = procLoad("Java_org_opencv_ml_RTrees_setActiveVarCount_10");
    proc_video_SparsePyrLKOpticalFlow_setTermCriteria_10 = procLoad("Java_org_opencv_video_SparsePyrLKOpticalFlow_setTermCriteria_10");
    proc_calib3d_StereoSGBM_getP2_10 = procLoad("Java_org_opencv_calib3d_StereoSGBM_getP2_10");
    proc_features2d_FlannBasedMatcher_create_10 = procLoad("Java_org_opencv_features2d_FlannBasedMatcher_create_10");
    proc_core_Core_divide_14 = procLoad("Java_org_opencv_core_Core_divide_14");
    proc_imgproc_Imgproc_getDerivKernels_11 = procLoad("Java_org_opencv_imgproc_Imgproc_getDerivKernels_11");
    proc_ml_RTrees_setCalculateVarImportance_10 = procLoad("Java_org_opencv_ml_RTrees_setCalculateVarImportance_10");
    proc_video_SparsePyrLKOpticalFlow_setWinSize_10 = procLoad("Java_org_opencv_video_SparsePyrLKOpticalFlow_setWinSize_10");
    proc_video_Video_CamShift_10 = procLoad("Java_org_opencv_video_Video_CamShift_10");
    proc_calib3d_StereoSGBM_getPreFilterCap_10 = procLoad("Java_org_opencv_calib3d_StereoSGBM_getPreFilterCap_10");
    proc_imgproc_Imgproc_getGaborKernel_10 = procLoad("Java_org_opencv_imgproc_Imgproc_getGaborKernel_10");
    proc_core_Core_divide_15 = procLoad("Java_org_opencv_core_Core_divide_15");
    proc_features2d_FlannBasedMatcher_delete = procLoad("Java_org_opencv_features2d_FlannBasedMatcher_delete");
    proc_ml_RTrees_setTermCriteria_10 = procLoad("Java_org_opencv_ml_RTrees_setTermCriteria_10");
    proc_calib3d_StereoSGBM_getUniquenessRatio_10 = procLoad("Java_org_opencv_calib3d_StereoSGBM_getUniquenessRatio_10");
    proc_video_Video_buildOpticalFlowPyramid_10 = procLoad("Java_org_opencv_video_Video_buildOpticalFlowPyramid_10");
    proc_features2d_GFTTDetector_create_10 = procLoad("Java_org_opencv_features2d_GFTTDetector_create_10");
    proc_imgproc_Imgproc_getGaborKernel_11 = procLoad("Java_org_opencv_imgproc_Imgproc_getGaborKernel_11");
    proc_ml_SVM_create_10 = procLoad("Java_org_opencv_ml_SVM_create_10");
    proc_core_Core_divide_16 = procLoad("Java_org_opencv_core_Core_divide_16");
    proc_video_Video_buildOpticalFlowPyramid_11 = procLoad("Java_org_opencv_video_Video_buildOpticalFlowPyramid_11");
    proc_features2d_GFTTDetector_create_11 = procLoad("Java_org_opencv_features2d_GFTTDetector_create_11");
    proc_imgproc_Imgproc_getGaussianKernel_10 = procLoad("Java_org_opencv_imgproc_Imgproc_getGaussianKernel_10");
    proc_calib3d_StereoSGBM_setMode_10 = procLoad("Java_org_opencv_calib3d_StereoSGBM_setMode_10");
    proc_core_Core_divide_17 = procLoad("Java_org_opencv_core_Core_divide_17");
    proc_ml_SVM_delete = procLoad("Java_org_opencv_ml_SVM_delete");
    proc_ml_SVM_getC_10 = procLoad("Java_org_opencv_ml_SVM_getC_10");
    proc_features2d_GFTTDetector_create_12 = procLoad("Java_org_opencv_features2d_GFTTDetector_create_12");
    proc_imgproc_Imgproc_getGaussianKernel_11 = procLoad("Java_org_opencv_imgproc_Imgproc_getGaussianKernel_11");
    proc_calib3d_StereoSGBM_setP1_10 = procLoad("Java_org_opencv_calib3d_StereoSGBM_setP1_10");
    proc_core_Core_eigenNonSymmetric_10 = procLoad("Java_org_opencv_core_Core_eigenNonSymmetric_10");
    proc_video_Video_calcOpticalFlowFarneback_10 = procLoad("Java_org_opencv_video_Video_calcOpticalFlowFarneback_10");
    proc_core_Core_eigen_10 = procLoad("Java_org_opencv_core_Core_eigen_10");
    proc_features2d_GFTTDetector_create_13 = procLoad("Java_org_opencv_features2d_GFTTDetector_create_13");
    proc_imgproc_Imgproc_getPerspectiveTransform_10 = procLoad("Java_org_opencv_imgproc_Imgproc_getPerspectiveTransform_10");
    proc_ml_SVM_getClassWeights_10 = procLoad("Java_org_opencv_ml_SVM_getClassWeights_10");
    proc_calib3d_StereoSGBM_setP2_10 = procLoad("Java_org_opencv_calib3d_StereoSGBM_setP2_10");
    proc_video_Video_calcOpticalFlowPyrLK_10 = procLoad("Java_org_opencv_video_Video_calcOpticalFlowPyrLK_10");
    proc_imgcodecs_Imgcodecs_imwrite_10 = procLoad("Java_org_opencv_imgcodecs_Imgcodecs_imwrite_10");
    proc_videoio_VideoCapture_isOpened_10 = procLoad("Java_org_opencv_videoio_VideoCapture_isOpened_10");
    proc_calib3d_Calib3d_calibrateCamera_11 = procLoad("Java_org_opencv_calib3d_Calib3d_calibrateCamera_11");
    proc_dnn_DictValue_getRealValue_11 = procLoad("Java_org_opencv_dnn_DictValue_getRealValue_11");
    proc_ml_ANN_1MLP_getRpropDWMin_10 = procLoad("Java_org_opencv_ml_ANN_1MLP_getRpropDWMin_10");
    proc_core_Mat_nPutD = procLoad("Java_org_opencv_core_Mat_nPutD");
    proc_features2d_AKAZE_getDiffusivity_10 = procLoad("Java_org_opencv_features2d_AKAZE_getDiffusivity_10");
    proc_photo_AlignMTB_getMaxBits_10 = procLoad("Java_org_opencv_photo_AlignMTB_getMaxBits_10");
    proc_video_BackgroundSubtractorKNN_getHistory_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorKNN_getHistory_10");
    proc_imgproc_Imgproc_Canny_10 = procLoad("Java_org_opencv_imgproc_Imgproc_Canny_10");
    proc_objdetect_HOGDescriptor_computeGradient_10 = procLoad("Java_org_opencv_objdetect_HOGDescriptor_computeGradient_10");
    proc_core_Core_eigen_11 = procLoad("Java_org_opencv_core_Core_eigen_11");
    proc_ml_SVM_getCoef0_10 = procLoad("Java_org_opencv_ml_SVM_getCoef0_10");
    proc_calib3d_StereoSGBM_setPreFilterCap_10 = procLoad("Java_org_opencv_calib3d_StereoSGBM_setPreFilterCap_10");
    proc_features2d_GFTTDetector_delete = procLoad("Java_org_opencv_features2d_GFTTDetector_delete");
    proc_imgproc_Imgproc_getRectSubPix_10 = procLoad("Java_org_opencv_imgproc_Imgproc_getRectSubPix_10");
    proc_video_Video_calcOpticalFlowPyrLK_11 = procLoad("Java_org_opencv_video_Video_calcOpticalFlowPyrLK_11");
    proc_ml_SVM_getDecisionFunction_10 = procLoad("Java_org_opencv_ml_SVM_getDecisionFunction_10");
    proc_features2d_GFTTDetector_getBlockSize_10 = procLoad("Java_org_opencv_features2d_GFTTDetector_getBlockSize_10");
    proc_calib3d_StereoSGBM_setUniquenessRatio_10 = procLoad("Java_org_opencv_calib3d_StereoSGBM_setUniquenessRatio_10");
    proc_core_Core_exp_10 = procLoad("Java_org_opencv_core_Core_exp_10");
    proc_imgproc_Imgproc_getRectSubPix_11 = procLoad("Java_org_opencv_imgproc_Imgproc_getRectSubPix_11");
    proc_video_Video_calcOpticalFlowPyrLK_12 = procLoad("Java_org_opencv_video_Video_calcOpticalFlowPyrLK_12");
    proc_imgproc_Imgproc_getRotationMatrix2D_10 = procLoad("Java_org_opencv_imgproc_Imgproc_getRotationMatrix2D_10");
    proc_ml_SVM_getDefaultGridPtr_10 = procLoad("Java_org_opencv_ml_SVM_getDefaultGridPtr_10");
    proc_video_Video_createBackgroundSubtractorKNN_10 = procLoad("Java_org_opencv_video_Video_createBackgroundSubtractorKNN_10");
    proc_features2d_GFTTDetector_getDefaultName_10 = procLoad("Java_org_opencv_features2d_GFTTDetector_getDefaultName_10");
    proc_core_Core_extractChannel_10 = procLoad("Java_org_opencv_core_Core_extractChannel_10");
    proc_features2d_GFTTDetector_getHarrisDetector_10 = procLoad("Java_org_opencv_features2d_GFTTDetector_getHarrisDetector_10");
    proc_ml_SVM_getDegree_10 = procLoad("Java_org_opencv_ml_SVM_getDegree_10");
    proc_core_Core_fastAtan2_10 = procLoad("Java_org_opencv_core_Core_fastAtan2_10");
    proc_imgproc_Imgproc_getStructuringElement_10 = procLoad("Java_org_opencv_imgproc_Imgproc_getStructuringElement_10");
    proc_video_Video_createBackgroundSubtractorKNN_11 = procLoad("Java_org_opencv_video_Video_createBackgroundSubtractorKNN_11");
    proc_features2d_GFTTDetector_getK_10 = procLoad("Java_org_opencv_features2d_GFTTDetector_getK_10");
    proc_ml_SVM_getGamma_10 = procLoad("Java_org_opencv_ml_SVM_getGamma_10");
    proc_imgproc_Imgproc_getStructuringElement_11 = procLoad("Java_org_opencv_imgproc_Imgproc_getStructuringElement_11");
    proc_video_Video_createBackgroundSubtractorMOG2_10 = procLoad("Java_org_opencv_video_Video_createBackgroundSubtractorMOG2_10");
    proc_core_Core_findNonZero_10 = procLoad("Java_org_opencv_core_Core_findNonZero_10");
    proc_features2d_GFTTDetector_getMaxFeatures_10 = procLoad("Java_org_opencv_features2d_GFTTDetector_getMaxFeatures_10");
    proc_ml_SVM_getKernelType_10 = procLoad("Java_org_opencv_ml_SVM_getKernelType_10");
    proc_video_Video_createBackgroundSubtractorMOG2_11 = procLoad("Java_org_opencv_video_Video_createBackgroundSubtractorMOG2_11");
    proc_core_Core_flip_10 = procLoad("Java_org_opencv_core_Core_flip_10");
    proc_imgproc_Imgproc_goodFeaturesToTrack_10 = procLoad("Java_org_opencv_imgproc_Imgproc_goodFeaturesToTrack_10");
    proc_features2d_GFTTDetector_getMinDistance_10 = procLoad("Java_org_opencv_features2d_GFTTDetector_getMinDistance_10");
    proc_ml_SVM_getNu_10 = procLoad("Java_org_opencv_ml_SVM_getNu_10");
    proc_video_Video_createOptFlow_1DualTVL1_10 = procLoad("Java_org_opencv_video_Video_createOptFlow_1DualTVL1_10");
    proc_core_Core_gemm_10 = procLoad("Java_org_opencv_core_Core_gemm_10");
    proc_imgproc_Imgproc_goodFeaturesToTrack_11 = procLoad("Java_org_opencv_imgproc_Imgproc_goodFeaturesToTrack_11");
    proc_features2d_GFTTDetector_getQualityLevel_10 = procLoad("Java_org_opencv_features2d_GFTTDetector_getQualityLevel_10");
    proc_ml_SVM_getP_10 = procLoad("Java_org_opencv_ml_SVM_getP_10");
    proc_video_Video_estimateRigidTransform_10 = procLoad("Java_org_opencv_video_Video_estimateRigidTransform_10");
    proc_core_Core_gemm_11 = procLoad("Java_org_opencv_core_Core_gemm_11");
    proc_imgproc_Imgproc_goodFeaturesToTrack_12 = procLoad("Java_org_opencv_imgproc_Imgproc_goodFeaturesToTrack_12");
    proc_video_Video_findTransformECC_10 = procLoad("Java_org_opencv_video_Video_findTransformECC_10");
    proc_ml_SVM_getSupportVectors_10 = procLoad("Java_org_opencv_ml_SVM_getSupportVectors_10");
    proc_core_Core_getBuildInformation_10 = procLoad("Java_org_opencv_core_Core_getBuildInformation_10");
    proc_features2d_GFTTDetector_setBlockSize_10 = procLoad("Java_org_opencv_features2d_GFTTDetector_setBlockSize_10");
    proc_imgproc_Imgproc_goodFeaturesToTrack_13 = procLoad("Java_org_opencv_imgproc_Imgproc_goodFeaturesToTrack_13");
    proc_video_Video_findTransformECC_11 = procLoad("Java_org_opencv_video_Video_findTransformECC_11");
    proc_ml_SVM_getTermCriteria_10 = procLoad("Java_org_opencv_ml_SVM_getTermCriteria_10");
    proc_core_Core_getCPUTickCount_10 = procLoad("Java_org_opencv_core_Core_getCPUTickCount_10");
    proc_features2d_GFTTDetector_setHarrisDetector_10 = procLoad("Java_org_opencv_features2d_GFTTDetector_setHarrisDetector_10");
    proc_imgproc_Imgproc_grabCut_10 = procLoad("Java_org_opencv_imgproc_Imgproc_grabCut_10");
    proc_imgcodecs_Imgcodecs_imwrite_11 = procLoad("Java_org_opencv_imgcodecs_Imgcodecs_imwrite_11");
    proc_videoio_VideoCapture_open_10 = procLoad("Java_org_opencv_videoio_VideoCapture_open_10");
    proc_calib3d_Calib3d_calibrateCamera_12 = procLoad("Java_org_opencv_calib3d_Calib3d_calibrateCamera_12");
    proc_ml_ANN_1MLP_getRpropDWMinus_10 = procLoad("Java_org_opencv_ml_ANN_1MLP_getRpropDWMinus_10");
    proc_core_Mat_nPutF = procLoad("Java_org_opencv_core_Mat_nPutF");
    proc_features2d_AKAZE_getNOctaveLayers_10 = procLoad("Java_org_opencv_features2d_AKAZE_getNOctaveLayers_10");
    proc_video_BackgroundSubtractorKNN_getNSamples_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorKNN_getNSamples_10");
    proc_dnn_DictValue_getStringValue_10 = procLoad("Java_org_opencv_dnn_DictValue_getStringValue_10");
    proc_imgproc_Imgproc_Canny_11 = procLoad("Java_org_opencv_imgproc_Imgproc_Canny_11");
    proc_objdetect_HOGDescriptor_computeGradient_11 = procLoad("Java_org_opencv_objdetect_HOGDescriptor_computeGradient_11");
    proc_photo_AlignMTB_process_10 = procLoad("Java_org_opencv_photo_AlignMTB_process_10");
    proc_video_Video_findTransformECC_12 = procLoad("Java_org_opencv_video_Video_findTransformECC_12");
    proc_ml_SVM_getType_10 = procLoad("Java_org_opencv_ml_SVM_getType_10");
    proc_core_Core_getIppVersion_10 = procLoad("Java_org_opencv_core_Core_getIppVersion_10");
    proc_features2d_GFTTDetector_setK_10 = procLoad("Java_org_opencv_features2d_GFTTDetector_setK_10");
    proc_imgproc_Imgproc_grabCut_11 = procLoad("Java_org_opencv_imgproc_Imgproc_grabCut_11");
    proc_core_Core_getNumThreads_10 = procLoad("Java_org_opencv_core_Core_getNumThreads_10");
    proc_video_Video_meanShift_10 = procLoad("Java_org_opencv_video_Video_meanShift_10");
    proc_ml_SVM_getUncompressedSupportVectors_10 = procLoad("Java_org_opencv_ml_SVM_getUncompressedSupportVectors_10");
    proc_features2d_GFTTDetector_setMaxFeatures_10 = procLoad("Java_org_opencv_features2d_GFTTDetector_setMaxFeatures_10");
    proc_imgproc_Imgproc_initUndistortRectifyMap_10 = procLoad("Java_org_opencv_imgproc_Imgproc_initUndistortRectifyMap_10");
    proc_imgproc_Imgproc_initWideAngleProjMap_10 = procLoad("Java_org_opencv_imgproc_Imgproc_initWideAngleProjMap_10");
    proc_core_Core_getNumberOfCPUs_10 = procLoad("Java_org_opencv_core_Core_getNumberOfCPUs_10");
    proc_ml_SVM_load_10 = procLoad("Java_org_opencv_ml_SVM_load_10");
    proc_features2d_GFTTDetector_setMinDistance_10 = procLoad("Java_org_opencv_features2d_GFTTDetector_setMinDistance_10");
    proc_imgproc_Imgproc_initWideAngleProjMap_11 = procLoad("Java_org_opencv_imgproc_Imgproc_initWideAngleProjMap_11");
    proc_core_Core_getOptimalDFTSize_10 = procLoad("Java_org_opencv_core_Core_getOptimalDFTSize_10");
    proc_features2d_GFTTDetector_setQualityLevel_10 = procLoad("Java_org_opencv_features2d_GFTTDetector_setQualityLevel_10");
    proc_ml_SVM_setC_10 = procLoad("Java_org_opencv_ml_SVM_setC_10");
    proc_core_Core_getThreadNum_10 = procLoad("Java_org_opencv_core_Core_getThreadNum_10");
    proc_features2d_KAZE_create_10 = procLoad("Java_org_opencv_features2d_KAZE_create_10");
    proc_imgproc_Imgproc_integral2_10 = procLoad("Java_org_opencv_imgproc_Imgproc_integral2_10");
    proc_ml_SVM_setClassWeights_10 = procLoad("Java_org_opencv_ml_SVM_setClassWeights_10");
    proc_core_Core_getTickCount_10 = procLoad("Java_org_opencv_core_Core_getTickCount_10");
    proc_features2d_KAZE_create_11 = procLoad("Java_org_opencv_features2d_KAZE_create_11");
    proc_imgproc_Imgproc_integral2_11 = procLoad("Java_org_opencv_imgproc_Imgproc_integral2_11");
    proc_ml_SVM_setCoef0_10 = procLoad("Java_org_opencv_ml_SVM_setCoef0_10");
    proc_core_Core_getTickFrequency_10 = procLoad("Java_org_opencv_core_Core_getTickFrequency_10");
    proc_features2d_KAZE_delete = procLoad("Java_org_opencv_features2d_KAZE_delete");
    proc_imgproc_Imgproc_integral3_10 = procLoad("Java_org_opencv_imgproc_Imgproc_integral3_10");
    proc_ml_SVM_setDegree_10 = procLoad("Java_org_opencv_ml_SVM_setDegree_10");
    proc_features2d_KAZE_getDefaultName_10 = procLoad("Java_org_opencv_features2d_KAZE_getDefaultName_10");
    proc_core_Core_hconcat_10 = procLoad("Java_org_opencv_core_Core_hconcat_10");
    proc_imgproc_Imgproc_integral3_11 = procLoad("Java_org_opencv_imgproc_Imgproc_integral3_11");
    proc_ml_SVM_setGamma_10 = procLoad("Java_org_opencv_ml_SVM_setGamma_10");
    proc_features2d_KAZE_getDiffusivity_10 = procLoad("Java_org_opencv_features2d_KAZE_getDiffusivity_10");
    proc_core_Core_idct_10 = procLoad("Java_org_opencv_core_Core_idct_10");
    proc_imgproc_Imgproc_integral_10 = procLoad("Java_org_opencv_imgproc_Imgproc_integral_10");
    proc_ml_SVM_setKernel_10 = procLoad("Java_org_opencv_ml_SVM_setKernel_10");
    proc_features2d_KAZE_getExtended_10 = procLoad("Java_org_opencv_features2d_KAZE_getExtended_10");
    proc_core_Core_idct_11 = procLoad("Java_org_opencv_core_Core_idct_11");
    proc_imgproc_Imgproc_integral_11 = procLoad("Java_org_opencv_imgproc_Imgproc_integral_11");
    proc_ml_SVM_setNu_10 = procLoad("Java_org_opencv_ml_SVM_setNu_10");
    proc_videoio_VideoCapture_open_11 = procLoad("Java_org_opencv_videoio_VideoCapture_open_11");
    proc_calib3d_Calib3d_calibrate_10 = procLoad("Java_org_opencv_calib3d_Calib3d_calibrate_10");
    proc_ml_ANN_1MLP_getRpropDWPlus_10 = procLoad("Java_org_opencv_ml_ANN_1MLP_getRpropDWPlus_10");
    proc_video_BackgroundSubtractorKNN_getShadowThreshold_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorKNN_getShadowThreshold_10");
    proc_core_Mat_nPutI = procLoad("Java_org_opencv_core_Mat_nPutI");
    proc_features2d_AKAZE_getNOctaves_10 = procLoad("Java_org_opencv_features2d_AKAZE_getNOctaves_10");
    proc_dnn_DictValue_getStringValue_11 = procLoad("Java_org_opencv_dnn_DictValue_getStringValue_11");
    proc_imgproc_Imgproc_Canny_12 = procLoad("Java_org_opencv_imgproc_Imgproc_Canny_12");
    proc_objdetect_HOGDescriptor_compute_10 = procLoad("Java_org_opencv_objdetect_HOGDescriptor_compute_10");
    proc_photo_AlignMTB_process_11 = procLoad("Java_org_opencv_photo_AlignMTB_process_11");
    proc_imgproc_Imgproc_intersectConvexConvex_10 = procLoad("Java_org_opencv_imgproc_Imgproc_intersectConvexConvex_10");
    proc_features2d_KAZE_getNOctaveLayers_10 = procLoad("Java_org_opencv_features2d_KAZE_getNOctaveLayers_10");
    proc_core_Core_idft_10 = procLoad("Java_org_opencv_core_Core_idft_10");
    proc_ml_SVM_setP_10 = procLoad("Java_org_opencv_ml_SVM_setP_10");
    proc_imgproc_Imgproc_intersectConvexConvex_11 = procLoad("Java_org_opencv_imgproc_Imgproc_intersectConvexConvex_11");
    proc_features2d_KAZE_getNOctaves_10 = procLoad("Java_org_opencv_features2d_KAZE_getNOctaves_10");
    proc_core_Core_idft_11 = procLoad("Java_org_opencv_core_Core_idft_11");
    proc_ml_SVM_setTermCriteria_10 = procLoad("Java_org_opencv_ml_SVM_setTermCriteria_10");
    proc_features2d_KAZE_getThreshold_10 = procLoad("Java_org_opencv_features2d_KAZE_getThreshold_10");
    proc_core_Core_inRange_10 = procLoad("Java_org_opencv_core_Core_inRange_10");
    proc_imgproc_Imgproc_invertAffineTransform_10 = procLoad("Java_org_opencv_imgproc_Imgproc_invertAffineTransform_10");
    proc_ml_SVM_setType_10 = procLoad("Java_org_opencv_ml_SVM_setType_10");
    proc_features2d_KAZE_getUpright_10 = procLoad("Java_org_opencv_features2d_KAZE_getUpright_10");
    proc_imgproc_Imgproc_isContourConvex_10 = procLoad("Java_org_opencv_imgproc_Imgproc_isContourConvex_10");
    proc_ml_SVM_trainAuto_10 = procLoad("Java_org_opencv_ml_SVM_trainAuto_10");
    proc_core_Core_insertChannel_10 = procLoad("Java_org_opencv_core_Core_insertChannel_10");
    proc_ml_SVM_trainAuto_11 = procLoad("Java_org_opencv_ml_SVM_trainAuto_11");
    proc_core_Core_invert_10 = procLoad("Java_org_opencv_core_Core_invert_10");
    proc_features2d_KAZE_setDiffusivity_10 = procLoad("Java_org_opencv_features2d_KAZE_setDiffusivity_10");
    proc_imgproc_Imgproc_line_10 = procLoad("Java_org_opencv_imgproc_Imgproc_line_10");
    proc_core_Core_invert_11 = procLoad("Java_org_opencv_core_Core_invert_11");
    proc_ml_SVMSGD_create_10 = procLoad("Java_org_opencv_ml_SVMSGD_create_10");
    proc_features2d_KAZE_setExtended_10 = procLoad("Java_org_opencv_features2d_KAZE_setExtended_10");
    proc_imgproc_Imgproc_line_11 = procLoad("Java_org_opencv_imgproc_Imgproc_line_11");
    proc_core_Core_kmeans_10 = procLoad("Java_org_opencv_core_Core_kmeans_10");
    proc_features2d_KAZE_setNOctaveLayers_10 = procLoad("Java_org_opencv_features2d_KAZE_setNOctaveLayers_10");
    proc_imgproc_Imgproc_line_12 = procLoad("Java_org_opencv_imgproc_Imgproc_line_12");
    proc_ml_SVMSGD_delete = procLoad("Java_org_opencv_ml_SVMSGD_delete");
    proc_core_Core_kmeans_11 = procLoad("Java_org_opencv_core_Core_kmeans_11");
    proc_ml_SVMSGD_getInitialStepSize_10 = procLoad("Java_org_opencv_ml_SVMSGD_getInitialStepSize_10");
    proc_features2d_KAZE_setNOctaves_10 = procLoad("Java_org_opencv_features2d_KAZE_setNOctaves_10");
    proc_imgproc_Imgproc_linearPolar_10 = procLoad("Java_org_opencv_imgproc_Imgproc_linearPolar_10");
    proc_ml_SVMSGD_getMarginRegularization_10 = procLoad("Java_org_opencv_ml_SVMSGD_getMarginRegularization_10");
    proc_core_Core_log_10 = procLoad("Java_org_opencv_core_Core_log_10");
    proc_features2d_KAZE_setThreshold_10 = procLoad("Java_org_opencv_features2d_KAZE_setThreshold_10");
    proc_imgproc_Imgproc_logPolar_10 = procLoad("Java_org_opencv_imgproc_Imgproc_logPolar_10");
    proc_imgproc_Imgproc_matchShapes_10 = procLoad("Java_org_opencv_imgproc_Imgproc_matchShapes_10");
    proc_ml_SVMSGD_getMarginType_10 = procLoad("Java_org_opencv_ml_SVMSGD_getMarginType_10");
    proc_core_Core_magnitude_10 = procLoad("Java_org_opencv_core_Core_magnitude_10");
    proc_features2d_KAZE_setUpright_10 = procLoad("Java_org_opencv_features2d_KAZE_setUpright_10");
    proc_dnn_DictValue_isInt_10 = procLoad("Java_org_opencv_dnn_DictValue_isInt_10");
    proc_videoio_VideoCapture_open_12 = procLoad("Java_org_opencv_videoio_VideoCapture_open_12");
    proc_calib3d_Calib3d_calibrate_11 = procLoad("Java_org_opencv_calib3d_Calib3d_calibrate_11");
    proc_features2d_AKAZE_getThreshold_10 = procLoad("Java_org_opencv_features2d_AKAZE_getThreshold_10");
    proc_ml_ANN_1MLP_getTermCriteria_10 = procLoad("Java_org_opencv_ml_ANN_1MLP_getTermCriteria_10");
    proc_core_Mat_nPutS = procLoad("Java_org_opencv_core_Mat_nPutS");
    proc_video_BackgroundSubtractorKNN_getShadowValue_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorKNN_getShadowValue_10");
    proc_imgproc_Imgproc_Canny_13 = procLoad("Java_org_opencv_imgproc_Imgproc_Canny_13");
    proc_objdetect_HOGDescriptor_compute_11 = procLoad("Java_org_opencv_objdetect_HOGDescriptor_compute_11");
    proc_photo_AlignMTB_setCut_10 = procLoad("Java_org_opencv_photo_AlignMTB_setCut_10");
    proc_ml_SVMSGD_getShift_10 = procLoad("Java_org_opencv_ml_SVMSGD_getShift_10");
    proc_features2d_MSER_create_10 = procLoad("Java_org_opencv_features2d_MSER_create_10");
    proc_core_Core_max_10 = procLoad("Java_org_opencv_core_Core_max_10");
    proc_imgproc_Imgproc_matchTemplate_10 = procLoad("Java_org_opencv_imgproc_Imgproc_matchTemplate_10");
    proc_ml_SVMSGD_getStepDecreasingPower_10 = procLoad("Java_org_opencv_ml_SVMSGD_getStepDecreasingPower_10");
    proc_features2d_MSER_create_11 = procLoad("Java_org_opencv_features2d_MSER_create_11");
    proc_core_Core_max_11 = procLoad("Java_org_opencv_core_Core_max_11");
    proc_imgproc_Imgproc_matchTemplate_11 = procLoad("Java_org_opencv_imgproc_Imgproc_matchTemplate_11");
    proc_ml_SVMSGD_getSvmsgdType_10 = procLoad("Java_org_opencv_ml_SVMSGD_getSvmsgdType_10");
    proc_core_Core_meanStdDev_10 = procLoad("Java_org_opencv_core_Core_meanStdDev_10");
    proc_features2d_MSER_delete = procLoad("Java_org_opencv_features2d_MSER_delete");
    proc_imgproc_Imgproc_medianBlur_10 = procLoad("Java_org_opencv_imgproc_Imgproc_medianBlur_10");
    proc_imgproc_Imgproc_minAreaRect_10 = procLoad("Java_org_opencv_imgproc_Imgproc_minAreaRect_10");
    proc_ml_SVMSGD_getTermCriteria_10 = procLoad("Java_org_opencv_ml_SVMSGD_getTermCriteria_10");
    proc_core_Core_meanStdDev_11 = procLoad("Java_org_opencv_core_Core_meanStdDev_11");
    proc_features2d_MSER_detectRegions_10 = procLoad("Java_org_opencv_features2d_MSER_detectRegions_10");
    proc_core_Core_mean_10 = procLoad("Java_org_opencv_core_Core_mean_10");
    proc_ml_SVMSGD_getWeights_10 = procLoad("Java_org_opencv_ml_SVMSGD_getWeights_10");
    proc_features2d_MSER_getDefaultName_10 = procLoad("Java_org_opencv_features2d_MSER_getDefaultName_10");
    proc_imgproc_Imgproc_minEnclosingCircle_10 = procLoad("Java_org_opencv_imgproc_Imgproc_minEnclosingCircle_10");
    proc_imgproc_Imgproc_minEnclosingTriangle_10 = procLoad("Java_org_opencv_imgproc_Imgproc_minEnclosingTriangle_10");
    proc_core_Core_mean_11 = procLoad("Java_org_opencv_core_Core_mean_11");
    proc_features2d_MSER_getDelta_10 = procLoad("Java_org_opencv_features2d_MSER_getDelta_10");
    proc_ml_SVMSGD_load_10 = procLoad("Java_org_opencv_ml_SVMSGD_load_10");
    proc_imgproc_Imgproc_moments_10 = procLoad("Java_org_opencv_imgproc_Imgproc_moments_10");
    proc_features2d_MSER_getMaxArea_10 = procLoad("Java_org_opencv_features2d_MSER_getMaxArea_10");
    proc_ml_SVMSGD_load_11 = procLoad("Java_org_opencv_ml_SVMSGD_load_11");
    proc_core_Core_merge_10 = procLoad("Java_org_opencv_core_Core_merge_10");
    proc_imgproc_Imgproc_moments_11 = procLoad("Java_org_opencv_imgproc_Imgproc_moments_11");
    proc_features2d_MSER_getMinArea_10 = procLoad("Java_org_opencv_features2d_MSER_getMinArea_10");
    proc_core_Core_min_10 = procLoad("Java_org_opencv_core_Core_min_10");
    proc_ml_SVMSGD_setInitialStepSize_10 = procLoad("Java_org_opencv_ml_SVMSGD_setInitialStepSize_10");
    proc_features2d_MSER_getPass2Only_10 = procLoad("Java_org_opencv_features2d_MSER_getPass2Only_10");
    proc_core_Core_min_11 = procLoad("Java_org_opencv_core_Core_min_11");
    proc_imgproc_Imgproc_morphologyEx_10 = procLoad("Java_org_opencv_imgproc_Imgproc_morphologyEx_10");
    proc_ml_SVMSGD_setMarginRegularization_10 = procLoad("Java_org_opencv_ml_SVMSGD_setMarginRegularization_10");
    proc_core_Core_mixChannels_10 = procLoad("Java_org_opencv_core_Core_mixChannels_10");
    proc_features2d_MSER_setDelta_10 = procLoad("Java_org_opencv_features2d_MSER_setDelta_10");
    proc_imgproc_Imgproc_morphologyEx_11 = procLoad("Java_org_opencv_imgproc_Imgproc_morphologyEx_11");
    proc_ml_SVMSGD_setMarginType_10 = procLoad("Java_org_opencv_ml_SVMSGD_setMarginType_10");
    proc_dnn_DictValue_isReal_10 = procLoad("Java_org_opencv_dnn_DictValue_isReal_10");
    proc_videoio_VideoCapture_open_13 = procLoad("Java_org_opencv_videoio_VideoCapture_open_13");
    proc_calib3d_Calib3d_calibrate_12 = procLoad("Java_org_opencv_calib3d_Calib3d_calibrate_12");
    proc_imgproc_Imgproc_EMD_10 = procLoad("Java_org_opencv_imgproc_Imgproc_EMD_10");
    proc_ml_ANN_1MLP_getTrainMethod_10 = procLoad("Java_org_opencv_ml_ANN_1MLP_getTrainMethod_10");
    proc_video_BackgroundSubtractorKNN_getkNNSamples_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorKNN_getkNNSamples_10");
    proc_core_Mat_n_1Mat__ = procLoad("Java_org_opencv_core_Mat_n_1Mat__");
    proc_features2d_AKAZE_setDescriptorChannels_10 = procLoad("Java_org_opencv_features2d_AKAZE_setDescriptorChannels_10");
    proc_objdetect_HOGDescriptor_delete = procLoad("Java_org_opencv_objdetect_HOGDescriptor_delete");
    proc_photo_AlignMTB_setExcludeRange_10 = procLoad("Java_org_opencv_photo_AlignMTB_setExcludeRange_10");
    proc_core_Core_mulSpectrums_10 = procLoad("Java_org_opencv_core_Core_mulSpectrums_10");
    proc_features2d_MSER_setMaxArea_10 = procLoad("Java_org_opencv_features2d_MSER_setMaxArea_10");
    proc_imgproc_Imgproc_morphologyEx_12 = procLoad("Java_org_opencv_imgproc_Imgproc_morphologyEx_12");
    proc_ml_SVMSGD_setOptimalParameters_10 = procLoad("Java_org_opencv_ml_SVMSGD_setOptimalParameters_10");
    proc_imgproc_Imgproc_n_1getTextSize = procLoad("Java_org_opencv_imgproc_Imgproc_n_1getTextSize");
    proc_core_Core_mulSpectrums_11 = procLoad("Java_org_opencv_core_Core_mulSpectrums_11");
    proc_features2d_MSER_setMinArea_10 = procLoad("Java_org_opencv_features2d_MSER_setMinArea_10");
    proc_ml_SVMSGD_setOptimalParameters_11 = procLoad("Java_org_opencv_ml_SVMSGD_setOptimalParameters_11");
    proc_imgproc_Imgproc_phaseCorrelate_10 = procLoad("Java_org_opencv_imgproc_Imgproc_phaseCorrelate_10");
    proc_core_Core_mulTransposed_10 = procLoad("Java_org_opencv_core_Core_mulTransposed_10");
    proc_features2d_MSER_setPass2Only_10 = procLoad("Java_org_opencv_features2d_MSER_setPass2Only_10");
    proc_ml_SVMSGD_setStepDecreasingPower_10 = procLoad("Java_org_opencv_ml_SVMSGD_setStepDecreasingPower_10");
    proc_imgproc_Imgproc_phaseCorrelate_11 = procLoad("Java_org_opencv_imgproc_Imgproc_phaseCorrelate_11");
    proc_features2d_ORB_create_10 = procLoad("Java_org_opencv_features2d_ORB_create_10");
    proc_core_Core_mulTransposed_11 = procLoad("Java_org_opencv_core_Core_mulTransposed_11");
    proc_ml_SVMSGD_setSvmsgdType_10 = procLoad("Java_org_opencv_ml_SVMSGD_setSvmsgdType_10");
    proc_imgproc_Imgproc_pointPolygonTest_10 = procLoad("Java_org_opencv_imgproc_Imgproc_pointPolygonTest_10");
    proc_features2d_ORB_create_11 = procLoad("Java_org_opencv_features2d_ORB_create_11");
    proc_core_Core_mulTransposed_12 = procLoad("Java_org_opencv_core_Core_mulTransposed_12");
    proc_ml_SVMSGD_setTermCriteria_10 = procLoad("Java_org_opencv_ml_SVMSGD_setTermCriteria_10");
    proc_ml_StatModel_calcError_10 = procLoad("Java_org_opencv_ml_StatModel_calcError_10");
    proc_core_Core_multiply_10 = procLoad("Java_org_opencv_core_Core_multiply_10");
    proc_features2d_ORB_delete = procLoad("Java_org_opencv_features2d_ORB_delete");
    proc_imgproc_Imgproc_polylines_10 = procLoad("Java_org_opencv_imgproc_Imgproc_polylines_10");
    proc_features2d_ORB_getDefaultName_10 = procLoad("Java_org_opencv_features2d_ORB_getDefaultName_10");
    proc_core_Core_multiply_11 = procLoad("Java_org_opencv_core_Core_multiply_11");
    proc_imgproc_Imgproc_polylines_11 = procLoad("Java_org_opencv_imgproc_Imgproc_polylines_11");
    proc_ml_StatModel_delete = procLoad("Java_org_opencv_ml_StatModel_delete");
    proc_ml_StatModel_empty_10 = procLoad("Java_org_opencv_ml_StatModel_empty_10");
    proc_features2d_ORB_getEdgeThreshold_10 = procLoad("Java_org_opencv_features2d_ORB_getEdgeThreshold_10");
    proc_core_Core_multiply_12 = procLoad("Java_org_opencv_core_Core_multiply_12");
    proc_imgproc_Imgproc_polylines_12 = procLoad("Java_org_opencv_imgproc_Imgproc_polylines_12");
    proc_features2d_ORB_getFastThreshold_10 = procLoad("Java_org_opencv_features2d_ORB_getFastThreshold_10");
    proc_ml_StatModel_getVarCount_10 = procLoad("Java_org_opencv_ml_StatModel_getVarCount_10");
    proc_core_Core_multiply_13 = procLoad("Java_org_opencv_core_Core_multiply_13");
    proc_imgproc_Imgproc_preCornerDetect_10 = procLoad("Java_org_opencv_imgproc_Imgproc_preCornerDetect_10");
    proc_ml_StatModel_isClassifier_10 = procLoad("Java_org_opencv_ml_StatModel_isClassifier_10");
    proc_features2d_ORB_getFirstLevel_10 = procLoad("Java_org_opencv_features2d_ORB_getFirstLevel_10");
    proc_core_Core_multiply_14 = procLoad("Java_org_opencv_core_Core_multiply_14");
    proc_imgproc_Imgproc_preCornerDetect_11 = procLoad("Java_org_opencv_imgproc_Imgproc_preCornerDetect_11");
    proc_dnn_DictValue_isString_10 = procLoad("Java_org_opencv_dnn_DictValue_isString_10");
    proc_videoio_VideoCapture_read_10 = procLoad("Java_org_opencv_videoio_VideoCapture_read_10");
    proc_imgproc_Imgproc_EMD_11 = procLoad("Java_org_opencv_imgproc_Imgproc_EMD_11");
    proc_core_Mat_n_1Mat__III = procLoad("Java_org_opencv_core_Mat_n_1Mat__III");
    proc_ml_ANN_1MLP_getWeights_10 = procLoad("Java_org_opencv_ml_ANN_1MLP_getWeights_10");
    proc_calib3d_Calib3d_calibrationMatrixValues_10 = procLoad("Java_org_opencv_calib3d_Calib3d_calibrationMatrixValues_10");
    proc_features2d_AKAZE_setDescriptorSize_10 = procLoad("Java_org_opencv_features2d_AKAZE_setDescriptorSize_10");
    proc_objdetect_HOGDescriptor_detectMultiScale_10 = procLoad("Java_org_opencv_objdetect_HOGDescriptor_detectMultiScale_10");
    proc_photo_AlignMTB_setMaxBits_10 = procLoad("Java_org_opencv_photo_AlignMTB_setMaxBits_10");
    proc_video_BackgroundSubtractorKNN_setDetectShadows_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorKNN_setDetectShadows_10");
    proc_ml_StatModel_isTrained_10 = procLoad("Java_org_opencv_ml_StatModel_isTrained_10");
    proc_features2d_ORB_getMaxFeatures_10 = procLoad("Java_org_opencv_features2d_ORB_getMaxFeatures_10");
    proc_core_Core_multiply_15 = procLoad("Java_org_opencv_core_Core_multiply_15");
    proc_imgproc_Imgproc_putText_10 = procLoad("Java_org_opencv_imgproc_Imgproc_putText_10");
    proc_core_Core_n_1minMaxLocManual = procLoad("Java_org_opencv_core_Core_n_1minMaxLocManual");
    proc_ml_StatModel_predict_10 = procLoad("Java_org_opencv_ml_StatModel_predict_10");
    proc_features2d_ORB_getNLevels_10 = procLoad("Java_org_opencv_features2d_ORB_getNLevels_10");
    proc_imgproc_Imgproc_putText_11 = procLoad("Java_org_opencv_imgproc_Imgproc_putText_11");
    proc_core_Core_norm_10 = procLoad("Java_org_opencv_core_Core_norm_10");
    proc_ml_StatModel_predict_11 = procLoad("Java_org_opencv_ml_StatModel_predict_11");
    proc_features2d_ORB_getPatchSize_10 = procLoad("Java_org_opencv_features2d_ORB_getPatchSize_10");
    proc_imgproc_Imgproc_putText_12 = procLoad("Java_org_opencv_imgproc_Imgproc_putText_12");
    proc_ml_StatModel_train_10 = procLoad("Java_org_opencv_ml_StatModel_train_10");
    proc_core_Core_norm_11 = procLoad("Java_org_opencv_core_Core_norm_11");
    proc_features2d_ORB_getScaleFactor_10 = procLoad("Java_org_opencv_features2d_ORB_getScaleFactor_10");
    proc_imgproc_Imgproc_pyrDown_10 = procLoad("Java_org_opencv_imgproc_Imgproc_pyrDown_10");
    proc_ml_StatModel_train_11 = procLoad("Java_org_opencv_ml_StatModel_train_11");
    proc_core_Core_norm_12 = procLoad("Java_org_opencv_core_Core_norm_12");
    proc_features2d_ORB_getScoreType_10 = procLoad("Java_org_opencv_features2d_ORB_getScoreType_10");
    proc_imgproc_Imgproc_pyrDown_11 = procLoad("Java_org_opencv_imgproc_Imgproc_pyrDown_11");
    proc_ml_StatModel_train_12 = procLoad("Java_org_opencv_ml_StatModel_train_12");
    proc_core_Core_norm_13 = procLoad("Java_org_opencv_core_Core_norm_13");
    proc_features2d_ORB_getWTA_1K_10 = procLoad("Java_org_opencv_features2d_ORB_getWTA_1K_10");
    proc_imgproc_Imgproc_pyrDown_12 = procLoad("Java_org_opencv_imgproc_Imgproc_pyrDown_12");
    proc_core_Core_norm_14 = procLoad("Java_org_opencv_core_Core_norm_14");
    proc_ml_TrainData_create_10 = procLoad("Java_org_opencv_ml_TrainData_create_10");
    proc_features2d_ORB_setEdgeThreshold_10 = procLoad("Java_org_opencv_features2d_ORB_setEdgeThreshold_10");
    proc_imgproc_Imgproc_pyrMeanShiftFiltering_10 = procLoad("Java_org_opencv_imgproc_Imgproc_pyrMeanShiftFiltering_10");
    proc_core_Core_norm_15 = procLoad("Java_org_opencv_core_Core_norm_15");
    proc_ml_TrainData_create_11 = procLoad("Java_org_opencv_ml_TrainData_create_11");
    proc_features2d_ORB_setFastThreshold_10 = procLoad("Java_org_opencv_features2d_ORB_setFastThreshold_10");
    proc_imgproc_Imgproc_pyrMeanShiftFiltering_11 = procLoad("Java_org_opencv_imgproc_Imgproc_pyrMeanShiftFiltering_11");
    proc_core_Core_normalize_10 = procLoad("Java_org_opencv_core_Core_normalize_10");
    proc_features2d_ORB_setFirstLevel_10 = procLoad("Java_org_opencv_features2d_ORB_setFirstLevel_10");
    proc_imgproc_Imgproc_pyrUp_10 = procLoad("Java_org_opencv_imgproc_Imgproc_pyrUp_10");
    proc_ml_TrainData_delete = procLoad("Java_org_opencv_ml_TrainData_delete");
    proc_ml_TrainData_getCatCount_10 = procLoad("Java_org_opencv_ml_TrainData_getCatCount_10");
    proc_core_Core_normalize_11 = procLoad("Java_org_opencv_core_Core_normalize_11");
    proc_features2d_ORB_setMaxFeatures_10 = procLoad("Java_org_opencv_features2d_ORB_setMaxFeatures_10");
    proc_imgproc_Imgproc_pyrUp_11 = procLoad("Java_org_opencv_imgproc_Imgproc_pyrUp_11");
    proc_core_Mat_n_1Mat__DDI = procLoad("Java_org_opencv_core_Mat_n_1Mat__DDI");
    proc_dnn_Dnn_blobFromImage_10 = procLoad("Java_org_opencv_dnn_Dnn_blobFromImage_10");
    proc_ml_ANN_1MLP_load_10 = procLoad("Java_org_opencv_ml_ANN_1MLP_load_10");
    proc_calib3d_Calib3d_composeRT_10 = procLoad("Java_org_opencv_calib3d_Calib3d_composeRT_10");
    proc_features2d_AKAZE_setDescriptorType_10 = procLoad("Java_org_opencv_features2d_AKAZE_setDescriptorType_10");
    proc_imgproc_Imgproc_GaussianBlur_10 = procLoad("Java_org_opencv_imgproc_Imgproc_GaussianBlur_10");
    proc_objdetect_HOGDescriptor_detectMultiScale_11 = procLoad("Java_org_opencv_objdetect_HOGDescriptor_detectMultiScale_11");
    proc_photo_AlignMTB_shiftMat_10 = procLoad("Java_org_opencv_photo_AlignMTB_shiftMat_10");
    proc_video_BackgroundSubtractorKNN_setDist2Threshold_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorKNN_setDist2Threshold_10");
    proc_videoio_VideoCapture_release_10 = procLoad("Java_org_opencv_videoio_VideoCapture_release_10");
    proc_ml_TrainData_getCatMap_10 = procLoad("Java_org_opencv_ml_TrainData_getCatMap_10");
    proc_core_Core_normalize_12 = procLoad("Java_org_opencv_core_Core_normalize_12");
    proc_features2d_ORB_setNLevels_10 = procLoad("Java_org_opencv_features2d_ORB_setNLevels_10");
    proc_imgproc_Imgproc_pyrUp_12 = procLoad("Java_org_opencv_imgproc_Imgproc_pyrUp_12");
    proc_ml_TrainData_getCatOfs_10 = procLoad("Java_org_opencv_ml_TrainData_getCatOfs_10");
    proc_core_Core_normalize_13 = procLoad("Java_org_opencv_core_Core_normalize_13");
    proc_features2d_ORB_setPatchSize_10 = procLoad("Java_org_opencv_features2d_ORB_setPatchSize_10");
    proc_imgproc_Imgproc_rectangle_10 = procLoad("Java_org_opencv_imgproc_Imgproc_rectangle_10");
    proc_ml_TrainData_getClassLabels_10 = procLoad("Java_org_opencv_ml_TrainData_getClassLabels_10");
    proc_core_Core_patchNaNs_10 = procLoad("Java_org_opencv_core_Core_patchNaNs_10");
    proc_features2d_ORB_setScaleFactor_10 = procLoad("Java_org_opencv_features2d_ORB_setScaleFactor_10");
    proc_imgproc_Imgproc_rectangle_11 = procLoad("Java_org_opencv_imgproc_Imgproc_rectangle_11");
    proc_ml_TrainData_getDefaultSubstValues_10 = procLoad("Java_org_opencv_ml_TrainData_getDefaultSubstValues_10");
    proc_core_Core_patchNaNs_11 = procLoad("Java_org_opencv_core_Core_patchNaNs_11");
    proc_features2d_ORB_setScoreType_10 = procLoad("Java_org_opencv_features2d_ORB_setScoreType_10");
    proc_imgproc_Imgproc_rectangle_12 = procLoad("Java_org_opencv_imgproc_Imgproc_rectangle_12");
    proc_ml_TrainData_getLayout_10 = procLoad("Java_org_opencv_ml_TrainData_getLayout_10");
    proc_core_Core_perspectiveTransform_10 = procLoad("Java_org_opencv_core_Core_perspectiveTransform_10");
    proc_features2d_ORB_setWTA_1K_10 = procLoad("Java_org_opencv_features2d_ORB_setWTA_1K_10");
    proc_imgproc_Imgproc_remap_10 = procLoad("Java_org_opencv_imgproc_Imgproc_remap_10");
    proc_features2d_Params_Params_10 = procLoad("Java_org_opencv_features2d_Params_Params_10");
    proc_ml_TrainData_getMissing_10 = procLoad("Java_org_opencv_ml_TrainData_getMissing_10");
    proc_core_Core_phase_10 = procLoad("Java_org_opencv_core_Core_phase_10");
    proc_imgproc_Imgproc_remap_11 = procLoad("Java_org_opencv_imgproc_Imgproc_remap_11");
    proc_ml_TrainData_getNAllVars_10 = procLoad("Java_org_opencv_ml_TrainData_getNAllVars_10");
    proc_core_Core_phase_11 = procLoad("Java_org_opencv_core_Core_phase_11");
    proc_features2d_Params_delete = procLoad("Java_org_opencv_features2d_Params_delete");
    proc_imgproc_Imgproc_resize_10 = procLoad("Java_org_opencv_imgproc_Imgproc_resize_10");
    proc_features2d_Params_get_1filterByArea_10 = procLoad("Java_org_opencv_features2d_Params_get_1filterByArea_10");
    proc_ml_TrainData_getNSamples_10 = procLoad("Java_org_opencv_ml_TrainData_getNSamples_10");
    proc_core_Core_polarToCart_10 = procLoad("Java_org_opencv_core_Core_polarToCart_10");
    proc_imgproc_Imgproc_resize_11 = procLoad("Java_org_opencv_imgproc_Imgproc_resize_11");
    proc_features2d_Params_get_1filterByCircularity_10 = procLoad("Java_org_opencv_features2d_Params_get_1filterByCircularity_10");
    proc_imgproc_Imgproc_rotatedRectangleIntersection_10 = procLoad("Java_org_opencv_imgproc_Imgproc_rotatedRectangleIntersection_10");
    proc_ml_TrainData_getNTestSamples_10 = procLoad("Java_org_opencv_ml_TrainData_getNTestSamples_10");
    proc_core_Core_polarToCart_11 = procLoad("Java_org_opencv_core_Core_polarToCart_11");
    proc_features2d_Params_get_1filterByColor_10 = procLoad("Java_org_opencv_features2d_Params_get_1filterByColor_10");
    proc_ml_TrainData_getNTrainSamples_10 = procLoad("Java_org_opencv_ml_TrainData_getNTrainSamples_10");
    proc_core_Core_pow_10 = procLoad("Java_org_opencv_core_Core_pow_10");
    proc_imgproc_Imgproc_sepFilter2D_10 = procLoad("Java_org_opencv_imgproc_Imgproc_sepFilter2D_10");
    proc_videoio_VideoCapture_retrieve_10 = procLoad("Java_org_opencv_videoio_VideoCapture_retrieve_10");
    proc_core_Mat_n_1Mat__IIIDDDD = procLoad("Java_org_opencv_core_Mat_n_1Mat__IIIDDDD");
    proc_dnn_Dnn_blobFromImage_11 = procLoad("Java_org_opencv_dnn_Dnn_blobFromImage_11");
    proc_calib3d_Calib3d_composeRT_11 = procLoad("Java_org_opencv_calib3d_Calib3d_composeRT_11");
    proc_features2d_AKAZE_setDiffusivity_10 = procLoad("Java_org_opencv_features2d_AKAZE_setDiffusivity_10");
    proc_imgproc_Imgproc_GaussianBlur_11 = procLoad("Java_org_opencv_imgproc_Imgproc_GaussianBlur_11");
    proc_ml_ANN_1MLP_setActivationFunction_10 = procLoad("Java_org_opencv_ml_ANN_1MLP_setActivationFunction_10");
    proc_objdetect_HOGDescriptor_detect_10 = procLoad("Java_org_opencv_objdetect_HOGDescriptor_detect_10");
    proc_photo_CalibrateCRF_delete = procLoad("Java_org_opencv_photo_CalibrateCRF_delete");
    proc_video_BackgroundSubtractorKNN_setHistory_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorKNN_setHistory_10");
    proc_features2d_Params_get_1filterByConvexity_10 = procLoad("Java_org_opencv_features2d_Params_get_1filterByConvexity_10");
    proc_ml_TrainData_getNVars_10 = procLoad("Java_org_opencv_ml_TrainData_getNVars_10");
    proc_core_Core_randShuffle_10 = procLoad("Java_org_opencv_core_Core_randShuffle_10");
    proc_imgproc_Imgproc_sepFilter2D_11 = procLoad("Java_org_opencv_imgproc_Imgproc_sepFilter2D_11");
    proc_features2d_Params_get_1filterByInertia_10 = procLoad("Java_org_opencv_features2d_Params_get_1filterByInertia_10");
    proc_core_Core_randShuffle_11 = procLoad("Java_org_opencv_core_Core_randShuffle_11");
    proc_imgproc_Imgproc_sepFilter2D_12 = procLoad("Java_org_opencv_imgproc_Imgproc_sepFilter2D_12");
    proc_ml_TrainData_getNames_10 = procLoad("Java_org_opencv_ml_TrainData_getNames_10");
    proc_features2d_Params_get_1maxArea_10 = procLoad("Java_org_opencv_features2d_Params_get_1maxArea_10");
    proc_ml_TrainData_getNormCatResponses_10 = procLoad("Java_org_opencv_ml_TrainData_getNormCatResponses_10");
    proc_core_Core_randn_10 = procLoad("Java_org_opencv_core_Core_randn_10");
    proc_imgproc_Imgproc_spatialGradient_10 = procLoad("Java_org_opencv_imgproc_Imgproc_spatialGradient_10");
    proc_features2d_Params_get_1maxCircularity_10 = procLoad("Java_org_opencv_features2d_Params_get_1maxCircularity_10");
    proc_ml_TrainData_getResponseType_10 = procLoad("Java_org_opencv_ml_TrainData_getResponseType_10");
    proc_core_Core_randu_10 = procLoad("Java_org_opencv_core_Core_randu_10");
    proc_imgproc_Imgproc_spatialGradient_11 = procLoad("Java_org_opencv_imgproc_Imgproc_spatialGradient_11");
    proc_features2d_Params_get_1maxConvexity_10 = procLoad("Java_org_opencv_features2d_Params_get_1maxConvexity_10");
    proc_ml_TrainData_getResponses_10 = procLoad("Java_org_opencv_ml_TrainData_getResponses_10");
    proc_core_Core_reduce_10 = procLoad("Java_org_opencv_core_Core_reduce_10");
    proc_imgproc_Imgproc_spatialGradient_12 = procLoad("Java_org_opencv_imgproc_Imgproc_spatialGradient_12");
    proc_features2d_Params_get_1maxInertiaRatio_10 = procLoad("Java_org_opencv_features2d_Params_get_1maxInertiaRatio_10");
    proc_ml_TrainData_getSampleWeights_10 = procLoad("Java_org_opencv_ml_TrainData_getSampleWeights_10");
    proc_core_Core_reduce_11 = procLoad("Java_org_opencv_core_Core_reduce_11");
    proc_imgproc_Imgproc_sqrBoxFilter_10 = procLoad("Java_org_opencv_imgproc_Imgproc_sqrBoxFilter_10");
    proc_features2d_Params_get_1maxThreshold_10 = procLoad("Java_org_opencv_features2d_Params_get_1maxThreshold_10");
    proc_core_Core_repeat_10 = procLoad("Java_org_opencv_core_Core_repeat_10");
    proc_imgproc_Imgproc_sqrBoxFilter_11 = procLoad("Java_org_opencv_imgproc_Imgproc_sqrBoxFilter_11");
    proc_ml_TrainData_getSample_10 = procLoad("Java_org_opencv_ml_TrainData_getSample_10");
    proc_features2d_Params_get_1minArea_10 = procLoad("Java_org_opencv_features2d_Params_get_1minArea_10");
    proc_ml_TrainData_getSamples_10 = procLoad("Java_org_opencv_ml_TrainData_getSamples_10");
    proc_core_Core_rotate_10 = procLoad("Java_org_opencv_core_Core_rotate_10");
    proc_imgproc_Imgproc_sqrBoxFilter_12 = procLoad("Java_org_opencv_imgproc_Imgproc_sqrBoxFilter_12");
    proc_imgproc_Imgproc_threshold_10 = procLoad("Java_org_opencv_imgproc_Imgproc_threshold_10");
    proc_features2d_Params_get_1minCircularity_10 = procLoad("Java_org_opencv_features2d_Params_get_1minCircularity_10");
    proc_ml_TrainData_getSubVector_10 = procLoad("Java_org_opencv_ml_TrainData_getSubVector_10");
    proc_core_Core_scaleAdd_10 = procLoad("Java_org_opencv_core_Core_scaleAdd_10");
    proc_features2d_Params_get_1minConvexity_10 = procLoad("Java_org_opencv_features2d_Params_get_1minConvexity_10");
    proc_ml_TrainData_getTestNormCatResponses_10 = procLoad("Java_org_opencv_ml_TrainData_getTestNormCatResponses_10");
    proc_core_Core_setErrorVerbosity_10 = procLoad("Java_org_opencv_core_Core_setErrorVerbosity_10");
    proc_imgproc_Imgproc_undistortPointsIter_10 = procLoad("Java_org_opencv_imgproc_Imgproc_undistortPointsIter_10");
    proc_videoio_VideoCapture_retrieve_11 = procLoad("Java_org_opencv_videoio_VideoCapture_retrieve_11");
    proc_core_Mat_n_1Mat__DDIDDDD = procLoad("Java_org_opencv_core_Mat_n_1Mat__DDIDDDD");
    proc_dnn_Dnn_blobFromImages_10 = procLoad("Java_org_opencv_dnn_Dnn_blobFromImages_10");
    proc_calib3d_Calib3d_computeCorrespondEpilines_10 = procLoad("Java_org_opencv_calib3d_Calib3d_computeCorrespondEpilines_10");
    proc_features2d_AKAZE_setNOctaveLayers_10 = procLoad("Java_org_opencv_features2d_AKAZE_setNOctaveLayers_10");
    proc_imgproc_Imgproc_GaussianBlur_12 = procLoad("Java_org_opencv_imgproc_Imgproc_GaussianBlur_12");
    proc_ml_ANN_1MLP_setActivationFunction_11 = procLoad("Java_org_opencv_ml_ANN_1MLP_setActivationFunction_11");
    proc_objdetect_HOGDescriptor_detect_11 = procLoad("Java_org_opencv_objdetect_HOGDescriptor_detect_11");
    proc_photo_CalibrateCRF_process_10 = procLoad("Java_org_opencv_photo_CalibrateCRF_process_10");
    proc_video_BackgroundSubtractorKNN_setNSamples_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorKNN_setNSamples_10");
    proc_features2d_Params_get_1minDistBetweenBlobs_10 = procLoad("Java_org_opencv_features2d_Params_get_1minDistBetweenBlobs_10");
    proc_ml_TrainData_getTestResponses_10 = procLoad("Java_org_opencv_ml_TrainData_getTestResponses_10");
    proc_core_Core_setIdentity_10 = procLoad("Java_org_opencv_core_Core_setIdentity_10");
    proc_imgproc_Imgproc_undistortPoints_10 = procLoad("Java_org_opencv_imgproc_Imgproc_undistortPoints_10");
    proc_features2d_Params_get_1minInertiaRatio_10 = procLoad("Java_org_opencv_features2d_Params_get_1minInertiaRatio_10");
    proc_ml_TrainData_getTestSampleIdx_10 = procLoad("Java_org_opencv_ml_TrainData_getTestSampleIdx_10");
    proc_core_Core_setIdentity_11 = procLoad("Java_org_opencv_core_Core_setIdentity_11");
    proc_imgproc_Imgproc_undistortPoints_11 = procLoad("Java_org_opencv_imgproc_Imgproc_undistortPoints_11");
    proc_features2d_Params_get_1minRepeatability_10 = procLoad("Java_org_opencv_features2d_Params_get_1minRepeatability_10");
    proc_ml_TrainData_getTestSampleWeights_10 = procLoad("Java_org_opencv_ml_TrainData_getTestSampleWeights_10");
    proc_core_Core_setNumThreads_10 = procLoad("Java_org_opencv_core_Core_setNumThreads_10");
    proc_imgproc_Imgproc_undistort_10 = procLoad("Java_org_opencv_imgproc_Imgproc_undistort_10");
    proc_features2d_Params_get_1minThreshold_10 = procLoad("Java_org_opencv_features2d_Params_get_1minThreshold_10");
    proc_ml_TrainData_getTestSamples_10 = procLoad("Java_org_opencv_ml_TrainData_getTestSamples_10");
    proc_core_Core_setRNGSeed_10 = procLoad("Java_org_opencv_core_Core_setRNGSeed_10");
    proc_imgproc_Imgproc_undistort_11 = procLoad("Java_org_opencv_imgproc_Imgproc_undistort_11");
    proc_features2d_Params_get_1thresholdStep_10 = procLoad("Java_org_opencv_features2d_Params_get_1thresholdStep_10");
    proc_ml_TrainData_getTrainNormCatResponses_10 = procLoad("Java_org_opencv_ml_TrainData_getTrainNormCatResponses_10");
    proc_core_Core_setUseIPP_10 = procLoad("Java_org_opencv_core_Core_setUseIPP_10");
    proc_imgproc_Imgproc_warpAffine_10 = procLoad("Java_org_opencv_imgproc_Imgproc_warpAffine_10");
    proc_ml_TrainData_getTrainResponses_10 = procLoad("Java_org_opencv_ml_TrainData_getTrainResponses_10");
    proc_core_Core_setUseIPP_1NE_10 = procLoad("Java_org_opencv_core_Core_setUseIPP_1NE_10");
    proc_features2d_Params_set_1filterByArea_10 = procLoad("Java_org_opencv_features2d_Params_set_1filterByArea_10");
    proc_imgproc_Imgproc_warpAffine_11 = procLoad("Java_org_opencv_imgproc_Imgproc_warpAffine_11");
    proc_core_Core_solveCubic_10 = procLoad("Java_org_opencv_core_Core_solveCubic_10");
    proc_ml_TrainData_getTrainSampleIdx_10 = procLoad("Java_org_opencv_ml_TrainData_getTrainSampleIdx_10");
    proc_features2d_Params_set_1filterByCircularity_10 = procLoad("Java_org_opencv_features2d_Params_set_1filterByCircularity_10");
    proc_imgproc_Imgproc_warpAffine_12 = procLoad("Java_org_opencv_imgproc_Imgproc_warpAffine_12");
    proc_core_Core_solvePoly_10 = procLoad("Java_org_opencv_core_Core_solvePoly_10");
    proc_ml_TrainData_getTrainSampleWeights_10 = procLoad("Java_org_opencv_ml_TrainData_getTrainSampleWeights_10");
    proc_features2d_Params_set_1filterByColor_10 = procLoad("Java_org_opencv_features2d_Params_set_1filterByColor_10");
    proc_imgproc_Imgproc_warpPerspective_10 = procLoad("Java_org_opencv_imgproc_Imgproc_warpPerspective_10");
    proc_core_Core_solvePoly_11 = procLoad("Java_org_opencv_core_Core_solvePoly_11");
    proc_ml_TrainData_getTrainSamples_10 = procLoad("Java_org_opencv_ml_TrainData_getTrainSamples_10");
    proc_features2d_Params_set_1filterByConvexity_10 = procLoad("Java_org_opencv_features2d_Params_set_1filterByConvexity_10");
    proc_imgproc_Imgproc_warpPerspective_11 = procLoad("Java_org_opencv_imgproc_Imgproc_warpPerspective_11");
    proc_core_Core_solve_10 = procLoad("Java_org_opencv_core_Core_solve_10");
    proc_ml_TrainData_getTrainSamples_11 = procLoad("Java_org_opencv_ml_TrainData_getTrainSamples_11");
    proc_features2d_Params_set_1filterByInertia_10 = procLoad("Java_org_opencv_features2d_Params_set_1filterByInertia_10");
    proc_imgproc_Imgproc_warpPerspective_12 = procLoad("Java_org_opencv_imgproc_Imgproc_warpPerspective_12");
    proc_videoio_VideoCapture_set_10 = procLoad("Java_org_opencv_videoio_VideoCapture_set_10");
    proc_core_Mat_n_1Mat__JIIII = procLoad("Java_org_opencv_core_Mat_n_1Mat__JIIII");
    proc_dnn_Dnn_blobFromImages_11 = procLoad("Java_org_opencv_dnn_Dnn_blobFromImages_11");
    proc_objdetect_HOGDescriptor_getDaimlerPeopleDetector_10 = procLoad("Java_org_opencv_objdetect_HOGDescriptor_getDaimlerPeopleDetector_10");
    proc_calib3d_Calib3d_convertPointsFromHomogeneous_10 = procLoad("Java_org_opencv_calib3d_Calib3d_convertPointsFromHomogeneous_10");
    proc_features2d_AKAZE_setNOctaves_10 = procLoad("Java_org_opencv_features2d_AKAZE_setNOctaves_10");
    proc_imgproc_Imgproc_HoughCircles_10 = procLoad("Java_org_opencv_imgproc_Imgproc_HoughCircles_10");
    proc_ml_ANN_1MLP_setBackpropMomentumScale_10 = procLoad("Java_org_opencv_ml_ANN_1MLP_setBackpropMomentumScale_10");
    proc_photo_CalibrateDebevec_delete = procLoad("Java_org_opencv_photo_CalibrateDebevec_delete");
    proc_video_BackgroundSubtractorKNN_setShadowThreshold_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorKNN_setShadowThreshold_10");
    proc_core_Core_solve_11 = procLoad("Java_org_opencv_core_Core_solve_11");
    proc_features2d_Params_set_1maxArea_10 = procLoad("Java_org_opencv_features2d_Params_set_1maxArea_10");
    proc_imgproc_Imgproc_watershed_10 = procLoad("Java_org_opencv_imgproc_Imgproc_watershed_10");
    proc_ml_TrainData_getValues_10 = procLoad("Java_org_opencv_ml_TrainData_getValues_10");
    proc_imgproc_LineSegmentDetector_compareSegments_10 = procLoad("Java_org_opencv_imgproc_LineSegmentDetector_compareSegments_10");
    proc_ml_TrainData_getVarIdx_10 = procLoad("Java_org_opencv_ml_TrainData_getVarIdx_10");
    proc_core_Core_sortIdx_10 = procLoad("Java_org_opencv_core_Core_sortIdx_10");
    proc_features2d_Params_set_1maxCircularity_10 = procLoad("Java_org_opencv_features2d_Params_set_1maxCircularity_10");
    proc_imgproc_LineSegmentDetector_compareSegments_11 = procLoad("Java_org_opencv_imgproc_LineSegmentDetector_compareSegments_11");
    proc_ml_TrainData_getVarSymbolFlags_10 = procLoad("Java_org_opencv_ml_TrainData_getVarSymbolFlags_10");
    proc_core_Core_sort_10 = procLoad("Java_org_opencv_core_Core_sort_10");
    proc_features2d_Params_set_1maxConvexity_10 = procLoad("Java_org_opencv_features2d_Params_set_1maxConvexity_10");
    proc_ml_TrainData_getVarType_10 = procLoad("Java_org_opencv_ml_TrainData_getVarType_10");
    proc_core_Core_split_10 = procLoad("Java_org_opencv_core_Core_split_10");
    proc_features2d_Params_set_1maxInertiaRatio_10 = procLoad("Java_org_opencv_features2d_Params_set_1maxInertiaRatio_10");
    proc_imgproc_LineSegmentDetector_delete = procLoad("Java_org_opencv_imgproc_LineSegmentDetector_delete");
    proc_core_Core_sqrt_10 = procLoad("Java_org_opencv_core_Core_sqrt_10");
    proc_features2d_Params_set_1maxThreshold_10 = procLoad("Java_org_opencv_features2d_Params_set_1maxThreshold_10");
    proc_imgproc_LineSegmentDetector_detect_10 = procLoad("Java_org_opencv_imgproc_LineSegmentDetector_detect_10");
    proc_ml_TrainData_setTrainTestSplitRatio_10 = procLoad("Java_org_opencv_ml_TrainData_setTrainTestSplitRatio_10");
    proc_core_Core_subtract_10 = procLoad("Java_org_opencv_core_Core_subtract_10");
    proc_features2d_Params_set_1minArea_10 = procLoad("Java_org_opencv_features2d_Params_set_1minArea_10");
    proc_imgproc_LineSegmentDetector_detect_11 = procLoad("Java_org_opencv_imgproc_LineSegmentDetector_detect_11");
    proc_ml_TrainData_setTrainTestSplitRatio_11 = procLoad("Java_org_opencv_ml_TrainData_setTrainTestSplitRatio_11");
    proc_core_Core_subtract_11 = procLoad("Java_org_opencv_core_Core_subtract_11");
    proc_features2d_Params_set_1minCircularity_10 = procLoad("Java_org_opencv_features2d_Params_set_1minCircularity_10");
    proc_imgproc_LineSegmentDetector_drawSegments_10 = procLoad("Java_org_opencv_imgproc_LineSegmentDetector_drawSegments_10");
    proc_ml_TrainData_setTrainTestSplit_10 = procLoad("Java_org_opencv_ml_TrainData_setTrainTestSplit_10");
    proc_imgproc_Subdiv2D_Subdiv2D_10 = procLoad("Java_org_opencv_imgproc_Subdiv2D_Subdiv2D_10");
    proc_core_Core_subtract_12 = procLoad("Java_org_opencv_core_Core_subtract_12");
    proc_features2d_Params_set_1minConvexity_10 = procLoad("Java_org_opencv_features2d_Params_set_1minConvexity_10");
    proc_ml_TrainData_setTrainTestSplit_11 = procLoad("Java_org_opencv_ml_TrainData_setTrainTestSplit_11");
    proc_imgproc_Subdiv2D_Subdiv2D_11 = procLoad("Java_org_opencv_imgproc_Subdiv2D_Subdiv2D_11");
    proc_core_Core_subtract_13 = procLoad("Java_org_opencv_core_Core_subtract_13");
    proc_features2d_Params_set_1minDistBetweenBlobs_10 = procLoad("Java_org_opencv_features2d_Params_set_1minDistBetweenBlobs_10");
    proc_ml_TrainData_shuffleTrainTest_10 = procLoad("Java_org_opencv_ml_TrainData_shuffleTrainTest_10");
    proc_core_Core_subtract_14 = procLoad("Java_org_opencv_core_Core_subtract_14");
    proc_features2d_Params_set_1minInertiaRatio_10 = procLoad("Java_org_opencv_features2d_Params_set_1minInertiaRatio_10");
    proc_imgproc_Subdiv2D_delete = procLoad("Java_org_opencv_imgproc_Subdiv2D_delete");
    proc_photo_CalibrateDebevec_getLambda_10 = procLoad("Java_org_opencv_photo_CalibrateDebevec_getLambda_10");
    proc_core_Mat_n_1Mat__JII = procLoad("Java_org_opencv_core_Mat_n_1Mat__JII");
    proc_dnn_Dnn_createCaffeImporter_10 = procLoad("Java_org_opencv_dnn_Dnn_createCaffeImporter_10");
    proc_objdetect_HOGDescriptor_getDefaultPeopleDetector_10 = procLoad("Java_org_opencv_objdetect_HOGDescriptor_getDefaultPeopleDetector_10");
    proc_videoio_VideoWriter_VideoWriter_10 = procLoad("Java_org_opencv_videoio_VideoWriter_VideoWriter_10");
    proc_calib3d_Calib3d_convertPointsToHomogeneous_10 = procLoad("Java_org_opencv_calib3d_Calib3d_convertPointsToHomogeneous_10");
    proc_features2d_AKAZE_setThreshold_10 = procLoad("Java_org_opencv_features2d_AKAZE_setThreshold_10");
    proc_imgproc_Imgproc_HoughCircles_11 = procLoad("Java_org_opencv_imgproc_Imgproc_HoughCircles_11");
    proc_ml_ANN_1MLP_setBackpropWeightScale_10 = procLoad("Java_org_opencv_ml_ANN_1MLP_setBackpropWeightScale_10");
    proc_video_BackgroundSubtractorKNN_setShadowValue_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorKNN_setShadowValue_10");
    proc_imgproc_Subdiv2D_edgeDst_10 = procLoad("Java_org_opencv_imgproc_Subdiv2D_edgeDst_10");
    proc_core_Core_subtract_15 = procLoad("Java_org_opencv_core_Core_subtract_15");
    proc_features2d_Params_set_1minRepeatability_10 = procLoad("Java_org_opencv_features2d_Params_set_1minRepeatability_10");
    proc_core_Core_sumElems_10 = procLoad("Java_org_opencv_core_Core_sumElems_10");
    proc_imgproc_Subdiv2D_edgeDst_11 = procLoad("Java_org_opencv_imgproc_Subdiv2D_edgeDst_11");
    proc_features2d_Params_set_1minThreshold_10 = procLoad("Java_org_opencv_features2d_Params_set_1minThreshold_10");
    proc_core_Core_trace_10 = procLoad("Java_org_opencv_core_Core_trace_10");
    proc_imgproc_Subdiv2D_edgeOrg_10 = procLoad("Java_org_opencv_imgproc_Subdiv2D_edgeOrg_10");
    proc_features2d_Params_set_1thresholdStep_10 = procLoad("Java_org_opencv_features2d_Params_set_1thresholdStep_10");
    proc_imgproc_Subdiv2D_edgeOrg_11 = procLoad("Java_org_opencv_imgproc_Subdiv2D_edgeOrg_11");
    proc_core_Core_transform_10 = procLoad("Java_org_opencv_core_Core_transform_10");
    proc_imgproc_Subdiv2D_findNearest_10 = procLoad("Java_org_opencv_imgproc_Subdiv2D_findNearest_10");
    proc_core_Core_transpose_10 = procLoad("Java_org_opencv_core_Core_transpose_10");
    proc_core_Core_useIPP_10 = procLoad("Java_org_opencv_core_Core_useIPP_10");
    proc_imgproc_Subdiv2D_findNearest_11 = procLoad("Java_org_opencv_imgproc_Subdiv2D_findNearest_11");
    proc_core_Core_useIPP_1NE_10 = procLoad("Java_org_opencv_core_Core_useIPP_1NE_10");
    proc_imgproc_Subdiv2D_getEdgeList_10 = procLoad("Java_org_opencv_imgproc_Subdiv2D_getEdgeList_10");
    proc_imgproc_Subdiv2D_getEdge_10 = procLoad("Java_org_opencv_imgproc_Subdiv2D_getEdge_10");
    proc_core_Core_vconcat_10 = procLoad("Java_org_opencv_core_Core_vconcat_10");
    proc_core_Algorithm_clear_10 = procLoad("Java_org_opencv_core_Algorithm_clear_10");
    proc_imgproc_Subdiv2D_getLeadingEdgeList_10 = procLoad("Java_org_opencv_imgproc_Subdiv2D_getLeadingEdgeList_10");
    proc_core_Algorithm_delete = procLoad("Java_org_opencv_core_Algorithm_delete");
    proc_imgproc_Subdiv2D_getTriangleList_10 = procLoad("Java_org_opencv_imgproc_Subdiv2D_getTriangleList_10");
    proc_photo_CalibrateDebevec_getRandom_10 = procLoad("Java_org_opencv_photo_CalibrateDebevec_getRandom_10");
    proc_core_Mat_n_1adjustROI = procLoad("Java_org_opencv_core_Mat_n_1adjustROI");
    proc_dnn_Dnn_createCaffeImporter_11 = procLoad("Java_org_opencv_dnn_Dnn_createCaffeImporter_11");
    proc_features2d_AgastFeatureDetector_create_10 = procLoad("Java_org_opencv_features2d_AgastFeatureDetector_create_10");
    proc_objdetect_HOGDescriptor_getDescriptorSize_10 = procLoad("Java_org_opencv_objdetect_HOGDescriptor_getDescriptorSize_10");
    proc_videoio_VideoWriter_VideoWriter_11 = procLoad("Java_org_opencv_videoio_VideoWriter_VideoWriter_11");
    proc_calib3d_Calib3d_correctMatches_10 = procLoad("Java_org_opencv_calib3d_Calib3d_correctMatches_10");
    proc_imgproc_Imgproc_HoughLinesP_10 = procLoad("Java_org_opencv_imgproc_Imgproc_HoughLinesP_10");
    proc_ml_ANN_1MLP_setLayerSizes_10 = procLoad("Java_org_opencv_ml_ANN_1MLP_setLayerSizes_10");
    proc_video_BackgroundSubtractorKNN_setkNNSamples_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorKNN_setkNNSamples_10");
    proc_imgproc_Subdiv2D_getVertex_10 = procLoad("Java_org_opencv_imgproc_Subdiv2D_getVertex_10");
    proc_core_Algorithm_getDefaultName_10 = procLoad("Java_org_opencv_core_Algorithm_getDefaultName_10");
    proc_imgproc_Subdiv2D_getVertex_11 = procLoad("Java_org_opencv_imgproc_Subdiv2D_getVertex_11");
    proc_core_Algorithm_save_10 = procLoad("Java_org_opencv_core_Algorithm_save_10");
    proc_core_TickMeter_TickMeter_10 = procLoad("Java_org_opencv_core_TickMeter_TickMeter_10");
    proc_imgproc_Subdiv2D_getVoronoiFacetList_10 = procLoad("Java_org_opencv_imgproc_Subdiv2D_getVoronoiFacetList_10");
    proc_core_TickMeter_delete = procLoad("Java_org_opencv_core_TickMeter_delete");
    proc_imgproc_Subdiv2D_initDelaunay_10 = procLoad("Java_org_opencv_imgproc_Subdiv2D_initDelaunay_10");
    proc_imgproc_Subdiv2D_insert_10 = procLoad("Java_org_opencv_imgproc_Subdiv2D_insert_10");
    proc_core_TickMeter_getCounter_10 = procLoad("Java_org_opencv_core_TickMeter_getCounter_10");
    proc_core_TickMeter_getTimeMicro_10 = procLoad("Java_org_opencv_core_TickMeter_getTimeMicro_10");
    proc_imgproc_Subdiv2D_insert_11 = procLoad("Java_org_opencv_imgproc_Subdiv2D_insert_11");
    proc_core_TickMeter_getTimeMilli_10 = procLoad("Java_org_opencv_core_TickMeter_getTimeMilli_10");
    proc_imgproc_Subdiv2D_locate_10 = procLoad("Java_org_opencv_imgproc_Subdiv2D_locate_10");
    proc_core_TickMeter_getTimeSec_10 = procLoad("Java_org_opencv_core_TickMeter_getTimeSec_10");
    proc_imgproc_Subdiv2D_nextEdge_10 = procLoad("Java_org_opencv_imgproc_Subdiv2D_nextEdge_10");
    proc_imgproc_Subdiv2D_rotateEdge_10 = procLoad("Java_org_opencv_imgproc_Subdiv2D_rotateEdge_10");
    proc_core_TickMeter_getTimeTicks_10 = procLoad("Java_org_opencv_core_TickMeter_getTimeTicks_10");
    proc_imgproc_Subdiv2D_symEdge_10 = procLoad("Java_org_opencv_imgproc_Subdiv2D_symEdge_10");
    proc_core_TickMeter_reset_10 = procLoad("Java_org_opencv_core_TickMeter_reset_10");
    proc_objdetect_HOGDescriptor_getWinSigma_10 = procLoad("Java_org_opencv_objdetect_HOGDescriptor_getWinSigma_10");
    proc_photo_CalibrateDebevec_getSamples_10 = procLoad("Java_org_opencv_photo_CalibrateDebevec_getSamples_10");
    proc_dnn_Dnn_createTensorflowImporter_10 = procLoad("Java_org_opencv_dnn_Dnn_createTensorflowImporter_10");
    proc_features2d_AgastFeatureDetector_create_11 = procLoad("Java_org_opencv_features2d_AgastFeatureDetector_create_11");
    proc_videoio_VideoWriter_VideoWriter_12 = procLoad("Java_org_opencv_videoio_VideoWriter_VideoWriter_12");
    proc_calib3d_Calib3d_decomposeEssentialMat_10 = procLoad("Java_org_opencv_calib3d_Calib3d_decomposeEssentialMat_10");
    proc_core_Mat_n_1assignTo__JJI = procLoad("Java_org_opencv_core_Mat_n_1assignTo__JJI");
    proc_imgproc_Imgproc_HoughLinesP_11 = procLoad("Java_org_opencv_imgproc_Imgproc_HoughLinesP_11");
    proc_ml_ANN_1MLP_setRpropDW0_10 = procLoad("Java_org_opencv_ml_ANN_1MLP_setRpropDW0_10");
    proc_video_BackgroundSubtractorMOG2_apply_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorMOG2_apply_10");
    proc_core_TickMeter_start_10 = procLoad("Java_org_opencv_core_TickMeter_start_10");
    proc_core_TickMeter_stop_10 = procLoad("Java_org_opencv_core_TickMeter_stop_10");
    proc_objdetect_HOGDescriptor_get_1L2HysThreshold_10 = procLoad("Java_org_opencv_objdetect_HOGDescriptor_get_1L2HysThreshold_10");
    proc_calib3d_Calib3d_decomposeHomographyMat_10 = procLoad("Java_org_opencv_calib3d_Calib3d_decomposeHomographyMat_10");
    proc_dnn_Dnn_createTorchImporter_10 = procLoad("Java_org_opencv_dnn_Dnn_createTorchImporter_10");
    proc_videoio_VideoWriter_VideoWriter_13 = procLoad("Java_org_opencv_videoio_VideoWriter_VideoWriter_13");
    proc_core_Mat_n_1assignTo__JJ = procLoad("Java_org_opencv_core_Mat_n_1assignTo__JJ");
    proc_features2d_AgastFeatureDetector_delete = procLoad("Java_org_opencv_features2d_AgastFeatureDetector_delete");
    proc_imgproc_Imgproc_HoughLines_10 = procLoad("Java_org_opencv_imgproc_Imgproc_HoughLines_10");
    proc_ml_ANN_1MLP_setRpropDWMax_10 = procLoad("Java_org_opencv_ml_ANN_1MLP_setRpropDWMax_10");
    proc_photo_CalibrateDebevec_setLambda_10 = procLoad("Java_org_opencv_photo_CalibrateDebevec_setLambda_10");
    proc_video_BackgroundSubtractorMOG2_apply_11 = procLoad("Java_org_opencv_video_BackgroundSubtractorMOG2_apply_11");
    proc_objdetect_HOGDescriptor_get_1blockSize_10 = procLoad("Java_org_opencv_objdetect_HOGDescriptor_get_1blockSize_10");
    proc_core_Mat_n_1channels = procLoad("Java_org_opencv_core_Mat_n_1channels");
    proc_dnn_Dnn_createTorchImporter_11 = procLoad("Java_org_opencv_dnn_Dnn_createTorchImporter_11");
    proc_videoio_VideoWriter_VideoWriter_14 = procLoad("Java_org_opencv_videoio_VideoWriter_VideoWriter_14");
    proc_features2d_AgastFeatureDetector_getDefaultName_10 = procLoad("Java_org_opencv_features2d_AgastFeatureDetector_getDefaultName_10");
    proc_calib3d_Calib3d_decomposeProjectionMatrix_10 = procLoad("Java_org_opencv_calib3d_Calib3d_decomposeProjectionMatrix_10");
    proc_imgproc_Imgproc_HoughLines_11 = procLoad("Java_org_opencv_imgproc_Imgproc_HoughLines_11");
    proc_ml_ANN_1MLP_setRpropDWMin_10 = procLoad("Java_org_opencv_ml_ANN_1MLP_setRpropDWMin_10");
    proc_photo_CalibrateDebevec_setRandom_10 = procLoad("Java_org_opencv_photo_CalibrateDebevec_setRandom_10");
    proc_video_BackgroundSubtractorMOG2_delete = procLoad("Java_org_opencv_video_BackgroundSubtractorMOG2_delete");
    proc_features2d_AgastFeatureDetector_getNonmaxSuppression_10 = procLoad("Java_org_opencv_features2d_AgastFeatureDetector_getNonmaxSuppression_10");
    proc_video_BackgroundSubtractorMOG2_getBackgroundRatio_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorMOG2_getBackgroundRatio_10");
    proc_objdetect_HOGDescriptor_get_1blockStride_10 = procLoad("Java_org_opencv_objdetect_HOGDescriptor_get_1blockStride_10");
    proc_core_Mat_n_1checkVector__JIIZ = procLoad("Java_org_opencv_core_Mat_n_1checkVector__JIIZ");
    proc_dnn_Dnn_readNetFromCaffe_10 = procLoad("Java_org_opencv_dnn_Dnn_readNetFromCaffe_10");
    proc_calib3d_Calib3d_decomposeProjectionMatrix_11 = procLoad("Java_org_opencv_calib3d_Calib3d_decomposeProjectionMatrix_11");
    proc_imgproc_Imgproc_HuMoments_10 = procLoad("Java_org_opencv_imgproc_Imgproc_HuMoments_10");
    proc_ml_ANN_1MLP_setRpropDWMinus_10 = procLoad("Java_org_opencv_ml_ANN_1MLP_setRpropDWMinus_10");
    proc_photo_CalibrateDebevec_setSamples_10 = procLoad("Java_org_opencv_photo_CalibrateDebevec_setSamples_10");
    proc_videoio_VideoWriter_delete = procLoad("Java_org_opencv_videoio_VideoWriter_delete");
    proc_video_BackgroundSubtractorMOG2_getComplexityReductionThreshold_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorMOG2_getComplexityReductionThreshold_10");
    proc_objdetect_HOGDescriptor_get_1cellSize_10 = procLoad("Java_org_opencv_objdetect_HOGDescriptor_get_1cellSize_10");
    proc_core_Mat_n_1checkVector__JII = procLoad("Java_org_opencv_core_Mat_n_1checkVector__JII");
    proc_features2d_AgastFeatureDetector_getThreshold_10 = procLoad("Java_org_opencv_features2d_AgastFeatureDetector_getThreshold_10");
    proc_videoio_VideoWriter_fourcc_10 = procLoad("Java_org_opencv_videoio_VideoWriter_fourcc_10");
    proc_dnn_Dnn_readNetFromCaffe_11 = procLoad("Java_org_opencv_dnn_Dnn_readNetFromCaffe_11");
    proc_calib3d_Calib3d_distortPoints_10 = procLoad("Java_org_opencv_calib3d_Calib3d_distortPoints_10");
    proc_imgproc_Imgproc_Laplacian_10 = procLoad("Java_org_opencv_imgproc_Imgproc_Laplacian_10");
    proc_ml_ANN_1MLP_setRpropDWPlus_10 = procLoad("Java_org_opencv_ml_ANN_1MLP_setRpropDWPlus_10");
    proc_photo_CalibrateRobertson_delete = procLoad("Java_org_opencv_photo_CalibrateRobertson_delete");
    proc_video_BackgroundSubtractorMOG2_getDetectShadows_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorMOG2_getDetectShadows_10");
    proc_videoio_VideoWriter_get_10 = procLoad("Java_org_opencv_videoio_VideoWriter_get_10");
    proc_core_Mat_n_1checkVector__JI = procLoad("Java_org_opencv_core_Mat_n_1checkVector__JI");
    proc_features2d_AgastFeatureDetector_getType_10 = procLoad("Java_org_opencv_features2d_AgastFeatureDetector_getType_10");
    proc_objdetect_HOGDescriptor_get_1derivAperture_10 = procLoad("Java_org_opencv_objdetect_HOGDescriptor_get_1derivAperture_10");
    proc_photo_CalibrateRobertson_getMaxIter_10 = procLoad("Java_org_opencv_photo_CalibrateRobertson_getMaxIter_10");
    proc_dnn_Dnn_readNetFromDarknet_10 = procLoad("Java_org_opencv_dnn_Dnn_readNetFromDarknet_10");
    proc_calib3d_Calib3d_distortPoints_11 = procLoad("Java_org_opencv_calib3d_Calib3d_distortPoints_11");
    proc_imgproc_Imgproc_Laplacian_11 = procLoad("Java_org_opencv_imgproc_Imgproc_Laplacian_11");
    proc_ml_ANN_1MLP_setTermCriteria_10 = procLoad("Java_org_opencv_ml_ANN_1MLP_setTermCriteria_10");
    proc_objdetect_HOGDescriptor_get_1gammaCorrection_10 = procLoad("Java_org_opencv_objdetect_HOGDescriptor_get_1gammaCorrection_10");
    proc_videoio_VideoWriter_isOpened_10 = procLoad("Java_org_opencv_videoio_VideoWriter_isOpened_10");
    proc_video_BackgroundSubtractorMOG2_getHistory_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorMOG2_getHistory_10");
    proc_core_Mat_n_1clone = procLoad("Java_org_opencv_core_Mat_n_1clone");
    proc_dnn_Dnn_readNetFromDarknet_11 = procLoad("Java_org_opencv_dnn_Dnn_readNetFromDarknet_11");
    proc_photo_CalibrateRobertson_getRadiance_10 = procLoad("Java_org_opencv_photo_CalibrateRobertson_getRadiance_10");
    proc_calib3d_Calib3d_drawChessboardCorners_10 = procLoad("Java_org_opencv_calib3d_Calib3d_drawChessboardCorners_10");
    proc_features2d_AgastFeatureDetector_setNonmaxSuppression_10 = procLoad("Java_org_opencv_features2d_AgastFeatureDetector_setNonmaxSuppression_10");
    proc_imgproc_Imgproc_Laplacian_12 = procLoad("Java_org_opencv_imgproc_Imgproc_Laplacian_12");
    proc_ml_ANN_1MLP_setTrainMethod_10 = procLoad("Java_org_opencv_ml_ANN_1MLP_setTrainMethod_10");
    proc_videoio_VideoWriter_open_10 = procLoad("Java_org_opencv_videoio_VideoWriter_open_10");
    proc_photo_CalibrateRobertson_getThreshold_10 = procLoad("Java_org_opencv_photo_CalibrateRobertson_getThreshold_10");
    proc_objdetect_HOGDescriptor_get_1histogramNormType_10 = procLoad("Java_org_opencv_objdetect_HOGDescriptor_get_1histogramNormType_10");
    proc_video_BackgroundSubtractorMOG2_getNMixtures_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorMOG2_getNMixtures_10");
    proc_calib3d_Calib3d_estimateAffine2D_10 = procLoad("Java_org_opencv_calib3d_Calib3d_estimateAffine2D_10");
    proc_core_Mat_n_1col = procLoad("Java_org_opencv_core_Mat_n_1col");
    proc_dnn_Dnn_readNetFromTensorflow_10 = procLoad("Java_org_opencv_dnn_Dnn_readNetFromTensorflow_10");
    proc_features2d_AgastFeatureDetector_setThreshold_10 = procLoad("Java_org_opencv_features2d_AgastFeatureDetector_setThreshold_10");
    proc_imgproc_Imgproc_Scharr_10 = procLoad("Java_org_opencv_imgproc_Imgproc_Scharr_10");
    proc_ml_ANN_1MLP_setTrainMethod_11 = procLoad("Java_org_opencv_ml_ANN_1MLP_setTrainMethod_11");
    proc_videoio_VideoWriter_open_11 = procLoad("Java_org_opencv_videoio_VideoWriter_open_11");
    proc_video_BackgroundSubtractorMOG2_getShadowThreshold_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorMOG2_getShadowThreshold_10");
    proc_objdetect_HOGDescriptor_get_1nbins_10 = procLoad("Java_org_opencv_objdetect_HOGDescriptor_get_1nbins_10");
    proc_calib3d_Calib3d_estimateAffine2D_11 = procLoad("Java_org_opencv_calib3d_Calib3d_estimateAffine2D_11");
    proc_core_Mat_n_1colRange = procLoad("Java_org_opencv_core_Mat_n_1colRange");
    proc_dnn_Dnn_readNetFromTensorflow_11 = procLoad("Java_org_opencv_dnn_Dnn_readNetFromTensorflow_11");
    proc_ml_Boost_create_10 = procLoad("Java_org_opencv_ml_Boost_create_10");
    proc_features2d_AgastFeatureDetector_setType_10 = procLoad("Java_org_opencv_features2d_AgastFeatureDetector_setType_10");
    proc_imgproc_Imgproc_Scharr_11 = procLoad("Java_org_opencv_imgproc_Imgproc_Scharr_11");
    proc_photo_CalibrateRobertson_setMaxIter_10 = procLoad("Java_org_opencv_photo_CalibrateRobertson_setMaxIter_10");
    proc_videoio_VideoWriter_open_12 = procLoad("Java_org_opencv_videoio_VideoWriter_open_12");
    proc_calib3d_Calib3d_estimateAffine3D_10 = procLoad("Java_org_opencv_calib3d_Calib3d_estimateAffine3D_10");
    proc_core_Mat_n_1cols = procLoad("Java_org_opencv_core_Mat_n_1cols");
    proc_objdetect_HOGDescriptor_get_1nlevels_10 = procLoad("Java_org_opencv_objdetect_HOGDescriptor_get_1nlevels_10");
    proc_video_BackgroundSubtractorMOG2_getShadowValue_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorMOG2_getShadowValue_10");
    proc_dnn_Dnn_readNetFromTorch_10 = procLoad("Java_org_opencv_dnn_Dnn_readNetFromTorch_10");
    proc_features2d_BFMatcher_BFMatcher_10 = procLoad("Java_org_opencv_features2d_BFMatcher_BFMatcher_10");
    proc_imgproc_Imgproc_Scharr_12 = procLoad("Java_org_opencv_imgproc_Imgproc_Scharr_12");
    proc_ml_Boost_delete = procLoad("Java_org_opencv_ml_Boost_delete");
    proc_photo_CalibrateRobertson_setThreshold_10 = procLoad("Java_org_opencv_photo_CalibrateRobertson_setThreshold_10");
    proc_objdetect_HOGDescriptor_get_1signedGradient_10 = procLoad("Java_org_opencv_objdetect_HOGDescriptor_get_1signedGradient_10");
    proc_videoio_VideoWriter_open_13 = procLoad("Java_org_opencv_videoio_VideoWriter_open_13");
    proc_video_BackgroundSubtractorMOG2_getVarInit_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorMOG2_getVarInit_10");
    proc_calib3d_Calib3d_estimateAffine3D_11 = procLoad("Java_org_opencv_calib3d_Calib3d_estimateAffine3D_11");
    proc_ml_Boost_getBoostType_10 = procLoad("Java_org_opencv_ml_Boost_getBoostType_10");
    proc_dnn_Dnn_readNetFromTorch_11 = procLoad("Java_org_opencv_dnn_Dnn_readNetFromTorch_11");
    proc_features2d_BFMatcher_BFMatcher_11 = procLoad("Java_org_opencv_features2d_BFMatcher_BFMatcher_11");
    proc_core_Mat_n_1convertTo__JJIDD = procLoad("Java_org_opencv_core_Mat_n_1convertTo__JJIDD");
    proc_imgproc_Imgproc_Sobel_10 = procLoad("Java_org_opencv_imgproc_Imgproc_Sobel_10");
    proc_photo_MergeDebevec_delete = procLoad("Java_org_opencv_photo_MergeDebevec_delete");
    proc_video_BackgroundSubtractorMOG2_getVarMax_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorMOG2_getVarMax_10");
    proc_ml_Boost_getWeakCount_10 = procLoad("Java_org_opencv_ml_Boost_getWeakCount_10");
    proc_calib3d_Calib3d_estimateAffinePartial2D_10 = procLoad("Java_org_opencv_calib3d_Calib3d_estimateAffinePartial2D_10");
    proc_dnn_Dnn_readTorchBlob_10 = procLoad("Java_org_opencv_dnn_Dnn_readTorchBlob_10");
    proc_features2d_BFMatcher_create_10 = procLoad("Java_org_opencv_features2d_BFMatcher_create_10");
    proc_objdetect_HOGDescriptor_get_1svmDetector_10 = procLoad("Java_org_opencv_objdetect_HOGDescriptor_get_1svmDetector_10");
    proc_core_Mat_n_1convertTo__JJID = procLoad("Java_org_opencv_core_Mat_n_1convertTo__JJID");
    proc_imgproc_Imgproc_Sobel_11 = procLoad("Java_org_opencv_imgproc_Imgproc_Sobel_11");
    proc_photo_MergeDebevec_process_10 = procLoad("Java_org_opencv_photo_MergeDebevec_process_10");
    proc_videoio_VideoWriter_release_10 = procLoad("Java_org_opencv_videoio_VideoWriter_release_10");
    proc_videoio_VideoWriter_set_10 = procLoad("Java_org_opencv_videoio_VideoWriter_set_10");
    proc_ml_Boost_getWeightTrimRate_10 = procLoad("Java_org_opencv_ml_Boost_getWeightTrimRate_10");
    proc_objdetect_HOGDescriptor_get_1winSigma_10 = procLoad("Java_org_opencv_objdetect_HOGDescriptor_get_1winSigma_10");
    proc_video_BackgroundSubtractorMOG2_getVarMin_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorMOG2_getVarMin_10");
    proc_calib3d_Calib3d_estimateAffinePartial2D_11 = procLoad("Java_org_opencv_calib3d_Calib3d_estimateAffinePartial2D_11");
    proc_dnn_Dnn_readTorchBlob_11 = procLoad("Java_org_opencv_dnn_Dnn_readTorchBlob_11");
    proc_features2d_BFMatcher_create_11 = procLoad("Java_org_opencv_features2d_BFMatcher_create_11");
    proc_core_Mat_n_1convertTo__JJI = procLoad("Java_org_opencv_core_Mat_n_1convertTo__JJI");
    proc_imgproc_Imgproc_Sobel_12 = procLoad("Java_org_opencv_imgproc_Imgproc_Sobel_12");
    proc_photo_MergeDebevec_process_11 = procLoad("Java_org_opencv_photo_MergeDebevec_process_11");
    proc_video_BackgroundSubtractorMOG2_getVarThresholdGen_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorMOG2_getVarThresholdGen_10");
    proc_objdetect_HOGDescriptor_get_1winSize_10 = procLoad("Java_org_opencv_objdetect_HOGDescriptor_get_1winSize_10");
    proc_ml_Boost_load_10 = procLoad("Java_org_opencv_ml_Boost_load_10");
    proc_calib3d_Calib3d_estimateNewCameraMatrixForUndistortRectify_10 = procLoad("Java_org_opencv_calib3d_Calib3d_estimateNewCameraMatrixForUndistortRectify_10");
    proc_core_Mat_n_1copyTo__JJ = procLoad("Java_org_opencv_core_Mat_n_1copyTo__JJ");
    proc_dnn_Dnn_shrinkCaffeModel_10 = procLoad("Java_org_opencv_dnn_Dnn_shrinkCaffeModel_10");
    proc_features2d_BFMatcher_delete = procLoad("Java_org_opencv_features2d_BFMatcher_delete");
    proc_imgproc_Imgproc_accumulateProduct_10 = procLoad("Java_org_opencv_imgproc_Imgproc_accumulateProduct_10");
    proc_photo_MergeExposures_delete = procLoad("Java_org_opencv_photo_MergeExposures_delete");
    proc_videoio_VideoWriter_write_10 = procLoad("Java_org_opencv_videoio_VideoWriter_write_10");
    proc_objdetect_HOGDescriptor_load_10 = procLoad("Java_org_opencv_objdetect_HOGDescriptor_load_10");
    proc_video_BackgroundSubtractorMOG2_getVarThreshold_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorMOG2_getVarThreshold_10");
    proc_ml_Boost_load_11 = procLoad("Java_org_opencv_ml_Boost_load_11");
    proc_calib3d_Calib3d_estimateNewCameraMatrixForUndistortRectify_11 = procLoad("Java_org_opencv_calib3d_Calib3d_estimateNewCameraMatrixForUndistortRectify_11");
    proc_core_Mat_n_1copyTo__JJJ = procLoad("Java_org_opencv_core_Mat_n_1copyTo__JJJ");
    proc_dnn_Importer_delete = procLoad("Java_org_opencv_dnn_Importer_delete");
    proc_features2d_BOWImgDescriptorExtractor_compute_10 = procLoad("Java_org_opencv_features2d_BOWImgDescriptorExtractor_compute_10");
    proc_imgproc_Imgproc_accumulateProduct_11 = procLoad("Java_org_opencv_imgproc_Imgproc_accumulateProduct_11");
    proc_photo_MergeExposures_process_10 = procLoad("Java_org_opencv_photo_MergeExposures_process_10");
    proc_objdetect_HOGDescriptor_load_11 = procLoad("Java_org_opencv_objdetect_HOGDescriptor_load_11");
    proc_calib3d_Calib3d_filterSpeckles_10 = procLoad("Java_org_opencv_calib3d_Calib3d_filterSpeckles_10");
    proc_core_Mat_n_1create__JIII = procLoad("Java_org_opencv_core_Mat_n_1create__JIII");
    proc_dnn_Importer_populateNet_10 = procLoad("Java_org_opencv_dnn_Importer_populateNet_10");
    proc_features2d_BOWImgDescriptorExtractor_delete = procLoad("Java_org_opencv_features2d_BOWImgDescriptorExtractor_delete");
    proc_imgproc_Imgproc_accumulateSquare_10 = procLoad("Java_org_opencv_imgproc_Imgproc_accumulateSquare_10");
    proc_ml_Boost_setBoostType_10 = procLoad("Java_org_opencv_ml_Boost_setBoostType_10");
    proc_photo_MergeMertens_delete = procLoad("Java_org_opencv_photo_MergeMertens_delete");
    proc_video_BackgroundSubtractorMOG2_setBackgroundRatio_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorMOG2_setBackgroundRatio_10");
    proc_photo_MergeMertens_getContrastWeight_10 = procLoad("Java_org_opencv_photo_MergeMertens_getContrastWeight_10");
    proc_features2d_BOWImgDescriptorExtractor_descriptorSize_10 = procLoad("Java_org_opencv_features2d_BOWImgDescriptorExtractor_descriptorSize_10");
    proc_calib3d_Calib3d_filterSpeckles_11 = procLoad("Java_org_opencv_calib3d_Calib3d_filterSpeckles_11");
    proc_core_Mat_n_1create__JDDI = procLoad("Java_org_opencv_core_Mat_n_1create__JDDI");
    proc_dnn_Layer_delete = procLoad("Java_org_opencv_dnn_Layer_delete");
    proc_imgproc_Imgproc_accumulateSquare_11 = procLoad("Java_org_opencv_imgproc_Imgproc_accumulateSquare_11");
    proc_ml_Boost_setWeakCount_10 = procLoad("Java_org_opencv_ml_Boost_setWeakCount_10");
    proc_objdetect_HOGDescriptor_save_10 = procLoad("Java_org_opencv_objdetect_HOGDescriptor_save_10");
    proc_video_BackgroundSubtractorMOG2_setComplexityReductionThreshold_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorMOG2_setComplexityReductionThreshold_10");
    proc_calib3d_Calib3d_findChessboardCorners_10 = procLoad("Java_org_opencv_calib3d_Calib3d_findChessboardCorners_10");
    proc_photo_MergeMertens_getExposureWeight_10 = procLoad("Java_org_opencv_photo_MergeMertens_getExposureWeight_10");
    proc_features2d_BOWImgDescriptorExtractor_descriptorType_10 = procLoad("Java_org_opencv_features2d_BOWImgDescriptorExtractor_descriptorType_10");
    proc_core_Mat_n_1cross = procLoad("Java_org_opencv_core_Mat_n_1cross");
    proc_dnn_Layer_finalize_10 = procLoad("Java_org_opencv_dnn_Layer_finalize_10");
    proc_imgproc_Imgproc_accumulateWeighted_10 = procLoad("Java_org_opencv_imgproc_Imgproc_accumulateWeighted_10");
    proc_ml_Boost_setWeightTrimRate_10 = procLoad("Java_org_opencv_ml_Boost_setWeightTrimRate_10");
    proc_objdetect_HOGDescriptor_save_11 = procLoad("Java_org_opencv_objdetect_HOGDescriptor_save_11");
    proc_video_BackgroundSubtractorMOG2_setDetectShadows_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorMOG2_setDetectShadows_10");
    proc_calib3d_Calib3d_findChessboardCorners_11 = procLoad("Java_org_opencv_calib3d_Calib3d_findChessboardCorners_11");
    proc_photo_MergeMertens_getSaturationWeight_10 = procLoad("Java_org_opencv_photo_MergeMertens_getSaturationWeight_10");
    proc_core_Mat_n_1dataAddr = procLoad("Java_org_opencv_core_Mat_n_1dataAddr");
    proc_features2d_BOWImgDescriptorExtractor_getVocabulary_10 = procLoad("Java_org_opencv_features2d_BOWImgDescriptorExtractor_getVocabulary_10");
    proc_ml_DTrees_create_10 = procLoad("Java_org_opencv_ml_DTrees_create_10");
    proc_dnn_Layer_finalize_11 = procLoad("Java_org_opencv_dnn_Layer_finalize_11");
    proc_imgproc_Imgproc_accumulateWeighted_11 = procLoad("Java_org_opencv_imgproc_Imgproc_accumulateWeighted_11");
    proc_objdetect_HOGDescriptor_setSVMDetector_10 = procLoad("Java_org_opencv_objdetect_HOGDescriptor_setSVMDetector_10");
    proc_video_BackgroundSubtractorMOG2_setHistory_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorMOG2_setHistory_10");
    proc_calib3d_Calib3d_findCirclesGrid_10 = procLoad("Java_org_opencv_calib3d_Calib3d_findCirclesGrid_10");
    proc_objdetect_CascadeClassifier_CascadeClassifier_10 = procLoad("Java_org_opencv_objdetect_CascadeClassifier_CascadeClassifier_10");
    proc_core_Mat_n_1delete = procLoad("Java_org_opencv_core_Mat_n_1delete");
    proc_dnn_Layer_forward_10 = procLoad("Java_org_opencv_dnn_Layer_forward_10");
    proc_features2d_BOWImgDescriptorExtractor_setVocabulary_10 = procLoad("Java_org_opencv_features2d_BOWImgDescriptorExtractor_setVocabulary_10");
    proc_imgproc_Imgproc_accumulate_10 = procLoad("Java_org_opencv_imgproc_Imgproc_accumulate_10");
    proc_ml_DTrees_delete = procLoad("Java_org_opencv_ml_DTrees_delete");
    proc_photo_MergeMertens_process_10 = procLoad("Java_org_opencv_photo_MergeMertens_process_10");
    proc_video_BackgroundSubtractorMOG2_setNMixtures_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorMOG2_setNMixtures_10");
    proc_calib3d_Calib3d_findCirclesGrid_11 = procLoad("Java_org_opencv_calib3d_Calib3d_findCirclesGrid_11");
    proc_core_Mat_n_1depth = procLoad("Java_org_opencv_core_Mat_n_1depth");
    proc_ml_DTrees_getCVFolds_10 = procLoad("Java_org_opencv_ml_DTrees_getCVFolds_10");
    proc_dnn_Layer_get_1blobs_10 = procLoad("Java_org_opencv_dnn_Layer_get_1blobs_10");
    proc_features2d_BOWKMeansTrainer_BOWKMeansTrainer_10 = procLoad("Java_org_opencv_features2d_BOWKMeansTrainer_BOWKMeansTrainer_10");
    proc_objdetect_CascadeClassifier_CascadeClassifier_11 = procLoad("Java_org_opencv_objdetect_CascadeClassifier_CascadeClassifier_11");
    proc_imgproc_Imgproc_accumulate_11 = procLoad("Java_org_opencv_imgproc_Imgproc_accumulate_11");
    proc_photo_MergeMertens_process_11 = procLoad("Java_org_opencv_photo_MergeMertens_process_11");
    proc_video_BackgroundSubtractorMOG2_setShadowThreshold_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorMOG2_setShadowThreshold_10");
    proc_objdetect_CascadeClassifier_convert_10 = procLoad("Java_org_opencv_objdetect_CascadeClassifier_convert_10");
    proc_ml_DTrees_getMaxCategories_10 = procLoad("Java_org_opencv_ml_DTrees_getMaxCategories_10");
    proc_calib3d_Calib3d_findEssentialMat_10 = procLoad("Java_org_opencv_calib3d_Calib3d_findEssentialMat_10");
    proc_core_Mat_n_1diag__JI = procLoad("Java_org_opencv_core_Mat_n_1diag__JI");
    proc_features2d_BOWKMeansTrainer_BOWKMeansTrainer_11 = procLoad("Java_org_opencv_features2d_BOWKMeansTrainer_BOWKMeansTrainer_11");
    proc_dnn_Layer_get_1name_10 = procLoad("Java_org_opencv_dnn_Layer_get_1name_10");
    proc_imgproc_Imgproc_adaptiveThreshold_10 = procLoad("Java_org_opencv_imgproc_Imgproc_adaptiveThreshold_10");
    proc_photo_MergeMertens_setContrastWeight_10 = procLoad("Java_org_opencv_photo_MergeMertens_setContrastWeight_10");
    proc_video_BackgroundSubtractorMOG2_setShadowValue_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorMOG2_setShadowValue_10");
    proc_dnn_Layer_get_1preferableTarget_10 = procLoad("Java_org_opencv_dnn_Layer_get_1preferableTarget_10");
    proc_ml_DTrees_getMaxDepth_10 = procLoad("Java_org_opencv_ml_DTrees_getMaxDepth_10");
    proc_calib3d_Calib3d_findEssentialMat_11 = procLoad("Java_org_opencv_calib3d_Calib3d_findEssentialMat_11");
    proc_core_Mat_n_1diag__J = procLoad("Java_org_opencv_core_Mat_n_1diag__J");
    proc_features2d_BOWKMeansTrainer_cluster_10 = procLoad("Java_org_opencv_features2d_BOWKMeansTrainer_cluster_10");
    proc_imgproc_Imgproc_applyColorMap_10 = procLoad("Java_org_opencv_imgproc_Imgproc_applyColorMap_10");
    proc_objdetect_CascadeClassifier_delete = procLoad("Java_org_opencv_objdetect_CascadeClassifier_delete");
    proc_photo_MergeMertens_setExposureWeight_10 = procLoad("Java_org_opencv_photo_MergeMertens_setExposureWeight_10");
    proc_video_BackgroundSubtractorMOG2_setVarInit_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorMOG2_setVarInit_10");
    proc_core_Mat_n_1dims = procLoad("Java_org_opencv_core_Mat_n_1dims");
    proc_ml_DTrees_getMinSampleCount_10 = procLoad("Java_org_opencv_ml_DTrees_getMinSampleCount_10");
    proc_calib3d_Calib3d_findEssentialMat_12 = procLoad("Java_org_opencv_calib3d_Calib3d_findEssentialMat_12");
    proc_features2d_BOWKMeansTrainer_cluster_11 = procLoad("Java_org_opencv_features2d_BOWKMeansTrainer_cluster_11");
    proc_dnn_Layer_get_1type_10 = procLoad("Java_org_opencv_dnn_Layer_get_1type_10");
    proc_imgproc_Imgproc_applyColorMap_11 = procLoad("Java_org_opencv_imgproc_Imgproc_applyColorMap_11");
    proc_objdetect_CascadeClassifier_detectMultiScale2_10 = procLoad("Java_org_opencv_objdetect_CascadeClassifier_detectMultiScale2_10");
    proc_photo_MergeMertens_setSaturationWeight_10 = procLoad("Java_org_opencv_photo_MergeMertens_setSaturationWeight_10");
    proc_video_BackgroundSubtractorMOG2_setVarMax_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorMOG2_setVarMax_10");
    proc_core_Mat_n_1dot = procLoad("Java_org_opencv_core_Mat_n_1dot");
    proc_calib3d_Calib3d_findEssentialMat_13 = procLoad("Java_org_opencv_calib3d_Calib3d_findEssentialMat_13");
    proc_ml_DTrees_getPriors_10 = procLoad("Java_org_opencv_ml_DTrees_getPriors_10");
    proc_dnn_Layer_run_10 = procLoad("Java_org_opencv_dnn_Layer_run_10");
    proc_features2d_BOWKMeansTrainer_delete = procLoad("Java_org_opencv_features2d_BOWKMeansTrainer_delete");
    proc_imgproc_Imgproc_approxPolyDP_10 = procLoad("Java_org_opencv_imgproc_Imgproc_approxPolyDP_10");
    proc_objdetect_CascadeClassifier_detectMultiScale2_11 = procLoad("Java_org_opencv_objdetect_CascadeClassifier_detectMultiScale2_11");
    proc_photo_MergeRobertson_delete = procLoad("Java_org_opencv_photo_MergeRobertson_delete");
    proc_video_BackgroundSubtractorMOG2_setVarMin_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorMOG2_setVarMin_10");
    proc_imgproc_Imgproc_arcLength_10 = procLoad("Java_org_opencv_imgproc_Imgproc_arcLength_10");
    proc_ml_DTrees_getRegressionAccuracy_10 = procLoad("Java_org_opencv_ml_DTrees_getRegressionAccuracy_10");
    proc_calib3d_Calib3d_findEssentialMat_14 = procLoad("Java_org_opencv_calib3d_Calib3d_findEssentialMat_14");
    proc_core_Mat_n_1elemSize = procLoad("Java_org_opencv_core_Mat_n_1elemSize");
    proc_dnn_Layer_set_1blobs_10 = procLoad("Java_org_opencv_dnn_Layer_set_1blobs_10");
    proc_features2d_BOWTrainer_add_10 = procLoad("Java_org_opencv_features2d_BOWTrainer_add_10");
    proc_objdetect_CascadeClassifier_detectMultiScale3_10 = procLoad("Java_org_opencv_objdetect_CascadeClassifier_detectMultiScale3_10");
    proc_photo_MergeRobertson_process_10 = procLoad("Java_org_opencv_photo_MergeRobertson_process_10");
    proc_video_BackgroundSubtractorMOG2_setVarThresholdGen_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorMOG2_setVarThresholdGen_10");
    proc_ml_DTrees_getTruncatePrunedTree_10 = procLoad("Java_org_opencv_ml_DTrees_getTruncatePrunedTree_10");
    proc_calib3d_Calib3d_findEssentialMat_15 = procLoad("Java_org_opencv_calib3d_Calib3d_findEssentialMat_15");
    proc_core_Mat_n_1elemSize1 = procLoad("Java_org_opencv_core_Mat_n_1elemSize1");
    proc_dnn_Net_Net_10 = procLoad("Java_org_opencv_dnn_Net_Net_10");
    proc_features2d_BOWTrainer_clear_10 = procLoad("Java_org_opencv_features2d_BOWTrainer_clear_10");
    proc_imgproc_Imgproc_arrowedLine_10 = procLoad("Java_org_opencv_imgproc_Imgproc_arrowedLine_10");
    proc_objdetect_CascadeClassifier_detectMultiScale3_11 = procLoad("Java_org_opencv_objdetect_CascadeClassifier_detectMultiScale3_11");
    proc_photo_MergeRobertson_process_11 = procLoad("Java_org_opencv_photo_MergeRobertson_process_11");
    proc_video_BackgroundSubtractorMOG2_setVarThreshold_10 = procLoad("Java_org_opencv_video_BackgroundSubtractorMOG2_setVarThreshold_10");
    proc_core_Mat_n_1empty = procLoad("Java_org_opencv_core_Mat_n_1empty");
    proc_ml_DTrees_getUse1SERule_10 = procLoad("Java_org_opencv_ml_DTrees_getUse1SERule_10");
    proc_calib3d_Calib3d_findFundamentalMat_10 = procLoad("Java_org_opencv_calib3d_Calib3d_findFundamentalMat_10");
    proc_features2d_BOWTrainer_cluster_10 = procLoad("Java_org_opencv_features2d_BOWTrainer_cluster_10");
    proc_dnn_Net_connect_10 = procLoad("Java_org_opencv_dnn_Net_connect_10");
    proc_imgproc_Imgproc_arrowedLine_11 = procLoad("Java_org_opencv_imgproc_Imgproc_arrowedLine_11");
    proc_objdetect_CascadeClassifier_detectMultiScale_10 = procLoad("Java_org_opencv_objdetect_CascadeClassifier_detectMultiScale_10");
    proc_photo_Photo_colorChange_10 = procLoad("Java_org_opencv_photo_Photo_colorChange_10");
    proc_video_DenseOpticalFlow_calc_10 = procLoad("Java_org_opencv_video_DenseOpticalFlow_calc_10");
    proc_ml_DTrees_getUseSurrogates_10 = procLoad("Java_org_opencv_ml_DTrees_getUseSurrogates_10");
    proc_calib3d_Calib3d_findFundamentalMat_11 = procLoad("Java_org_opencv_calib3d_Calib3d_findFundamentalMat_11");
    proc_core_Mat_n_1eye__III = procLoad("Java_org_opencv_core_Mat_n_1eye__III");
    proc_features2d_BOWTrainer_cluster_11 = procLoad("Java_org_opencv_features2d_BOWTrainer_cluster_11");
    proc_dnn_Net_delete = procLoad("Java_org_opencv_dnn_Net_delete");
    proc_imgproc_Imgproc_bilateralFilter_10 = procLoad("Java_org_opencv_imgproc_Imgproc_bilateralFilter_10");
    proc_objdetect_CascadeClassifier_detectMultiScale_11 = procLoad("Java_org_opencv_objdetect_CascadeClassifier_detectMultiScale_11");
    proc_photo_Photo_colorChange_11 = procLoad("Java_org_opencv_photo_Photo_colorChange_11");
    proc_video_DenseOpticalFlow_collectGarbage_10 = procLoad("Java_org_opencv_video_DenseOpticalFlow_collectGarbage_10");
    proc_objdetect_CascadeClassifier_empty_10 = procLoad("Java_org_opencv_objdetect_CascadeClassifier_empty_10");
    proc_calib3d_Calib3d_findFundamentalMat_12 = procLoad("Java_org_opencv_calib3d_Calib3d_findFundamentalMat_12");
    proc_core_Mat_n_1eye__DDI = procLoad("Java_org_opencv_core_Mat_n_1eye__DDI");
    proc_ml_DTrees_load_10 = procLoad("Java_org_opencv_ml_DTrees_load_10");
    proc_photo_Photo_createAlignMTB_10 = procLoad("Java_org_opencv_photo_Photo_createAlignMTB_10");
    proc_dnn_Net_deleteLayer_10 = procLoad("Java_org_opencv_dnn_Net_deleteLayer_10");
    proc_features2d_BOWTrainer_delete = procLoad("Java_org_opencv_features2d_BOWTrainer_delete");
    proc_imgproc_Imgproc_bilateralFilter_11 = procLoad("Java_org_opencv_imgproc_Imgproc_bilateralFilter_11");
    proc_video_DenseOpticalFlow_delete = procLoad("Java_org_opencv_video_DenseOpticalFlow_delete");
    proc_dnn_Net_empty_10 = procLoad("Java_org_opencv_dnn_Net_empty_10");
    proc_features2d_BOWTrainer_descriptorsCount_10 = procLoad("Java_org_opencv_features2d_BOWTrainer_descriptorsCount_10");
    proc_objdetect_CascadeClassifier_getFeatureType_10 = procLoad("Java_org_opencv_objdetect_CascadeClassifier_getFeatureType_10");
    proc_calib3d_Calib3d_findHomography_10 = procLoad("Java_org_opencv_calib3d_Calib3d_findHomography_10");
    proc_core_Mat_n_1inv__JI = procLoad("Java_org_opencv_core_Mat_n_1inv__JI");
    proc_ml_DTrees_load_11 = procLoad("Java_org_opencv_ml_DTrees_load_11");
    proc_photo_Photo_createAlignMTB_11 = procLoad("Java_org_opencv_photo_Photo_createAlignMTB_11");
    proc_video_DualTVL1OpticalFlow_create_10 = procLoad("Java_org_opencv_video_DualTVL1OpticalFlow_create_10");
    proc_imgproc_Imgproc_blur_10 = procLoad("Java_org_opencv_imgproc_Imgproc_blur_10");
    proc_objdetect_CascadeClassifier_getOriginalWindowSize_10 = procLoad("Java_org_opencv_objdetect_CascadeClassifier_getOriginalWindowSize_10");
    proc_calib3d_Calib3d_findHomography_11 = procLoad("Java_org_opencv_calib3d_Calib3d_findHomography_11");
    proc_core_Mat_n_1inv__J = procLoad("Java_org_opencv_core_Mat_n_1inv__J");
    proc_features2d_BOWTrainer_getDescriptors_10 = procLoad("Java_org_opencv_features2d_BOWTrainer_getDescriptors_10");
    proc_photo_Photo_createCalibrateDebevec_10 = procLoad("Java_org_opencv_photo_Photo_createCalibrateDebevec_10");
    proc_video_DualTVL1OpticalFlow_create_11 = procLoad("Java_org_opencv_video_DualTVL1OpticalFlow_create_11");
    proc_dnn_Net_enableFusion_10 = procLoad("Java_org_opencv_dnn_Net_enableFusion_10");
    proc_imgproc_Imgproc_blur_11 = procLoad("Java_org_opencv_imgproc_Imgproc_blur_11");
    proc_ml_DTrees_setCVFolds_10 = procLoad("Java_org_opencv_ml_DTrees_setCVFolds_10");
    proc_core_Mat_n_1isContinuous = procLoad("Java_org_opencv_core_Mat_n_1isContinuous");
    proc_objdetect_CascadeClassifier_isOldFormatCascade_10 = procLoad("Java_org_opencv_objdetect_CascadeClassifier_isOldFormatCascade_10");
    proc_calib3d_Calib3d_findHomography_12 = procLoad("Java_org_opencv_calib3d_Calib3d_findHomography_12");
    proc_dnn_Net_forward_10 = procLoad("Java_org_opencv_dnn_Net_forward_10");
    proc_features2d_BRISK_create_10 = procLoad("Java_org_opencv_features2d_BRISK_create_10");
    proc_photo_Photo_createCalibrateDebevec_11 = procLoad("Java_org_opencv_photo_Photo_createCalibrateDebevec_11");
    proc_imgproc_Imgproc_blur_12 = procLoad("Java_org_opencv_imgproc_Imgproc_blur_12");
    proc_ml_DTrees_setMaxCategories_10 = procLoad("Java_org_opencv_ml_DTrees_setMaxCategories_10");
    proc_video_DualTVL1OpticalFlow_delete = procLoad("Java_org_opencv_video_DualTVL1OpticalFlow_delete");
    proc_core_Mat_n_1isSubmatrix = procLoad("Java_org_opencv_core_Mat_n_1isSubmatrix");
    proc_objdetect_CascadeClassifier_load_10 = procLoad("Java_org_opencv_objdetect_CascadeClassifier_load_10");
    proc_video_DualTVL1OpticalFlow_getEpsilon_10 = procLoad("Java_org_opencv_video_DualTVL1OpticalFlow_getEpsilon_10");
    proc_imgproc_Imgproc_boundingRect_10 = procLoad("Java_org_opencv_imgproc_Imgproc_boundingRect_10");
    proc_calib3d_Calib3d_getOptimalNewCameraMatrix_10 = procLoad("Java_org_opencv_calib3d_Calib3d_getOptimalNewCameraMatrix_10");
    proc_dnn_Net_forward_11 = procLoad("Java_org_opencv_dnn_Net_forward_11");
    proc_features2d_BRISK_create_11 = procLoad("Java_org_opencv_features2d_BRISK_create_11");
    proc_photo_Photo_createCalibrateRobertson_10 = procLoad("Java_org_opencv_photo_Photo_createCalibrateRobertson_10");
    proc_ml_DTrees_setMaxDepth_10 = procLoad("Java_org_opencv_ml_DTrees_setMaxDepth_10");
    proc_video_DualTVL1OpticalFlow_getGamma_10 = procLoad("Java_org_opencv_video_DualTVL1OpticalFlow_getGamma_10");
    proc_calib3d_Calib3d_getOptimalNewCameraMatrix_11 = procLoad("Java_org_opencv_calib3d_Calib3d_getOptimalNewCameraMatrix_11");
    proc_core_Mat_n_1mul__JJD = procLoad("Java_org_opencv_core_Mat_n_1mul__JJD");
    proc_features2d_BRISK_create_12 = procLoad("Java_org_opencv_features2d_BRISK_create_12");
    proc_photo_Photo_createCalibrateRobertson_11 = procLoad("Java_org_opencv_photo_Photo_createCalibrateRobertson_11");
    proc_dnn_Net_forward_12 = procLoad("Java_org_opencv_dnn_Net_forward_12");
    proc_imgproc_Imgproc_boxFilter_10 = procLoad("Java_org_opencv_imgproc_Imgproc_boxFilter_10");
    proc_ml_DTrees_setMinSampleCount_10 = procLoad("Java_org_opencv_ml_DTrees_setMinSampleCount_10");
    proc_objdetect_BaseCascadeClassifier_delete = procLoad("Java_org_opencv_objdetect_BaseCascadeClassifier_delete");
    proc_calib3d_Calib3d_getValidDisparityROI_10 = procLoad("Java_org_opencv_calib3d_Calib3d_getValidDisparityROI_10");
    proc_video_DualTVL1OpticalFlow_getInnerIterations_10 = procLoad("Java_org_opencv_video_DualTVL1OpticalFlow_getInnerIterations_10");
    proc_core_Mat_n_1mul__JJ = procLoad("Java_org_opencv_core_Mat_n_1mul__JJ");
    proc_features2d_BRISK_create_13 = procLoad("Java_org_opencv_features2d_BRISK_create_13");
    proc_photo_Photo_createMergeDebevec_10 = procLoad("Java_org_opencv_photo_Photo_createMergeDebevec_10");
    proc_dnn_Net_forward_13 = procLoad("Java_org_opencv_dnn_Net_forward_13");
    proc_imgproc_Imgproc_boxFilter_11 = procLoad("Java_org_opencv_imgproc_Imgproc_boxFilter_11");
    proc_ml_DTrees_setPriors_10 = procLoad("Java_org_opencv_ml_DTrees_setPriors_10");
    proc_video_DualTVL1OpticalFlow_getLambda_10 = procLoad("Java_org_opencv_video_DualTVL1OpticalFlow_getLambda_10");
    proc_calib3d_Calib3d_initCameraMatrix2D_10 = procLoad("Java_org_opencv_calib3d_Calib3d_initCameraMatrix2D_10");
    proc_core_Mat_n_1ones__III = procLoad("Java_org_opencv_core_Mat_n_1ones__III");
    proc_features2d_BRISK_create_14 = procLoad("Java_org_opencv_features2d_BRISK_create_14");
    proc_photo_Photo_createMergeMertens_10 = procLoad("Java_org_opencv_photo_Photo_createMergeMertens_10");
    proc_dnn_Net_forward_14 = procLoad("Java_org_opencv_dnn_Net_forward_14");
    proc_imgproc_Imgproc_boxFilter_12 = procLoad("Java_org_opencv_imgproc_Imgproc_boxFilter_12");
    proc_ml_DTrees_setRegressionAccuracy_10 = procLoad("Java_org_opencv_ml_DTrees_setRegressionAccuracy_10");
    proc_video_DualTVL1OpticalFlow_getMedianFiltering_10 = procLoad("Java_org_opencv_video_DualTVL1OpticalFlow_getMedianFiltering_10");
    proc_calib3d_Calib3d_initCameraMatrix2D_11 = procLoad("Java_org_opencv_calib3d_Calib3d_initCameraMatrix2D_11");
    proc_core_Mat_n_1ones__DDI = procLoad("Java_org_opencv_core_Mat_n_1ones__DDI");
    proc_dnn_Net_getFLOPS_10 = procLoad("Java_org_opencv_dnn_Net_getFLOPS_10");
    proc_features2d_BRISK_create_15 = procLoad("Java_org_opencv_features2d_BRISK_create_15");
    proc_photo_Photo_createMergeMertens_11 = procLoad("Java_org_opencv_photo_Photo_createMergeMertens_11");
    proc_imgproc_Imgproc_boxPoints_10 = procLoad("Java_org_opencv_imgproc_Imgproc_boxPoints_10");
    proc_ml_DTrees_setTruncatePrunedTree_10 = procLoad("Java_org_opencv_ml_DTrees_setTruncatePrunedTree_10");
    proc_video_DualTVL1OpticalFlow_getOuterIterations_10 = procLoad("Java_org_opencv_video_DualTVL1OpticalFlow_getOuterIterations_10");
    proc_dnn_Net_getFLOPS_11 = procLoad("Java_org_opencv_dnn_Net_getFLOPS_11");
    proc_photo_Photo_createMergeRobertson_10 = procLoad("Java_org_opencv_photo_Photo_createMergeRobertson_10");
    proc_calib3d_Calib3d_initUndistortRectifyMap_10 = procLoad("Java_org_opencv_calib3d_Calib3d_initUndistortRectifyMap_10");
    proc_core_Mat_n_1push_1back = procLoad("Java_org_opencv_core_Mat_n_1push_1back");
    proc_features2d_BRISK_delete = procLoad("Java_org_opencv_features2d_BRISK_delete");
    proc_imgproc_Imgproc_calcBackProject_10 = procLoad("Java_org_opencv_imgproc_Imgproc_calcBackProject_10");
    proc_ml_DTrees_setUse1SERule_10 = procLoad("Java_org_opencv_ml_DTrees_setUse1SERule_10");
    proc_video_DualTVL1OpticalFlow_getScaleStep_10 = procLoad("Java_org_opencv_video_DualTVL1OpticalFlow_getScaleStep_10");
    proc_dnn_Net_getFLOPS_12 = procLoad("Java_org_opencv_dnn_Net_getFLOPS_12");
    proc_photo_Photo_createTonemapDrago_10 = procLoad("Java_org_opencv_photo_Photo_createTonemapDrago_10");
    proc_features2d_BRISK_getDefaultName_10 = procLoad("Java_org_opencv_features2d_BRISK_getDefaultName_10");
    proc_calib3d_Calib3d_matMulDeriv_10 = procLoad("Java_org_opencv_calib3d_Calib3d_matMulDeriv_10");
    proc_core_Mat_n_1release = procLoad("Java_org_opencv_core_Mat_n_1release");
    proc_imgproc_Imgproc_calcHist_10 = procLoad("Java_org_opencv_imgproc_Imgproc_calcHist_10");
    proc_ml_DTrees_setUseSurrogates_10 = procLoad("Java_org_opencv_ml_DTrees_setUseSurrogates_10");
    proc_video_DualTVL1OpticalFlow_getScalesNumber_10 = procLoad("Java_org_opencv_video_DualTVL1OpticalFlow_getScalesNumber_10");
    proc_core_Mat_n_1reshape__JII = procLoad("Java_org_opencv_core_Mat_n_1reshape__JII");
    proc_dnn_Net_getFLOPS_13 = procLoad("Java_org_opencv_dnn_Net_getFLOPS_13");
    proc_ml_EM_create_10 = procLoad("Java_org_opencv_ml_EM_create_10");
    proc_photo_Photo_createTonemapDrago_11 = procLoad("Java_org_opencv_photo_Photo_createTonemapDrago_11");
    proc_calib3d_Calib3d_projectPoints_10 = procLoad("Java_org_opencv_calib3d_Calib3d_projectPoints_10");
    proc_features2d_DescriptorExtractor_compute_10 = procLoad("Java_org_opencv_features2d_DescriptorExtractor_compute_10");
    proc_imgproc_Imgproc_calcHist_11 = procLoad("Java_org_opencv_imgproc_Imgproc_calcHist_11");
    proc_video_DualTVL1OpticalFlow_getTau_10 = procLoad("Java_org_opencv_video_DualTVL1OpticalFlow_getTau_10");
    proc_dnn_Net_getLayerId_10 = procLoad("Java_org_opencv_dnn_Net_getLayerId_10");
    proc_core_Mat_n_1reshape__JI = procLoad("Java_org_opencv_core_Mat_n_1reshape__JI");
    proc_photo_Photo_createTonemapDurand_10 = procLoad("Java_org_opencv_photo_Photo_createTonemapDurand_10");
    proc_calib3d_Calib3d_projectPoints_11 = procLoad("Java_org_opencv_calib3d_Calib3d_projectPoints_11");
    proc_features2d_DescriptorExtractor_compute_11 = procLoad("Java_org_opencv_features2d_DescriptorExtractor_compute_11");
    proc_imgproc_Imgproc_circle_10 = procLoad("Java_org_opencv_imgproc_Imgproc_circle_10");
    proc_ml_EM_delete = procLoad("Java_org_opencv_ml_EM_delete");
    proc_video_DualTVL1OpticalFlow_getTheta_10 = procLoad("Java_org_opencv_video_DualTVL1OpticalFlow_getTheta_10");
    proc_ml_EM_getClustersNumber_10 = procLoad("Java_org_opencv_ml_EM_getClustersNumber_10");
    proc_core_Mat_n_1row = procLoad("Java_org_opencv_core_Mat_n_1row");
    proc_dnn_Net_getLayerNames_10 = procLoad("Java_org_opencv_dnn_Net_getLayerNames_10");
    proc_features2d_DescriptorExtractor_create_10 = procLoad("Java_org_opencv_features2d_DescriptorExtractor_create_10");
    proc_photo_Photo_createTonemapDurand_11 = procLoad("Java_org_opencv_photo_Photo_createTonemapDurand_11");
    proc_calib3d_Calib3d_projectPoints_12 = procLoad("Java_org_opencv_calib3d_Calib3d_projectPoints_12");
    proc_imgproc_Imgproc_circle_11 = procLoad("Java_org_opencv_imgproc_Imgproc_circle_11");
    proc_video_DualTVL1OpticalFlow_getUseInitialFlow_10 = procLoad("Java_org_opencv_video_DualTVL1OpticalFlow_getUseInitialFlow_10");
    proc_ml_EM_getCovarianceMatrixType_10 = procLoad("Java_org_opencv_ml_EM_getCovarianceMatrixType_10");
    proc_core_Mat_n_1rowRange = procLoad("Java_org_opencv_core_Mat_n_1rowRange");
    proc_photo_Photo_createTonemapMantiuk_10 = procLoad("Java_org_opencv_photo_Photo_createTonemapMantiuk_10");
    proc_calib3d_Calib3d_projectPoints_13 = procLoad("Java_org_opencv_calib3d_Calib3d_projectPoints_13");
    proc_dnn_Net_getLayerTypes_10 = procLoad("Java_org_opencv_dnn_Net_getLayerTypes_10");
    proc_features2d_DescriptorExtractor_delete = procLoad("Java_org_opencv_features2d_DescriptorExtractor_delete");
    proc_imgproc_Imgproc_circle_12 = procLoad("Java_org_opencv_imgproc_Imgproc_circle_12");
    proc_calib3d_Calib3d_RQDecomp3x3_10 = procLoad("Java_org_opencv_calib3d_Calib3d_RQDecomp3x3_10");
    proc_core_Mat_nDump = procLoad("Java_org_opencv_core_Mat_nDump");
    proc_imgproc_Imgproc_clipLine_10 = procLoad("Java_org_opencv_imgproc_Imgproc_clipLine_10");
    proc_calib3d_Calib3d_recoverPose_10 = procLoad("Java_org_opencv_calib3d_Calib3d_recoverPose_10");
    proc_core_Mat_n_1rows = procLoad("Java_org_opencv_core_Mat_n_1rows");
    proc_features2d_DescriptorExtractor_descriptorSize_10 = procLoad("Java_org_opencv_features2d_DescriptorExtractor_descriptorSize_10");
    proc_video_DualTVL1OpticalFlow_getWarpingsNumber_10 = procLoad("Java_org_opencv_video_DualTVL1OpticalFlow_getWarpingsNumber_10");
    proc_dnn_Net_getLayer_10 = procLoad("Java_org_opencv_dnn_Net_getLayer_10");
    proc_photo_Photo_createTonemapMantiuk_11 = procLoad("Java_org_opencv_photo_Photo_createTonemapMantiuk_11");
    proc_ml_EM_getCovs_10 = procLoad("Java_org_opencv_ml_EM_getCovs_10");
    proc_imgproc_Imgproc_compareHist_10 = procLoad("Java_org_opencv_imgproc_Imgproc_compareHist_10");
    proc_calib3d_Calib3d_recoverPose_11 = procLoad("Java_org_opencv_calib3d_Calib3d_recoverPose_11");
    proc_dnn_Net_getLayersCount_10 = procLoad("Java_org_opencv_dnn_Net_getLayersCount_10");
    proc_features2d_DescriptorExtractor_descriptorType_10 = procLoad("Java_org_opencv_features2d_DescriptorExtractor_descriptorType_10");
    proc_core_Mat_n_1setTo__JDDDD = procLoad("Java_org_opencv_core_Mat_n_1setTo__JDDDD");
    proc_ml_EM_getMeans_10 = procLoad("Java_org_opencv_ml_EM_getMeans_10");
    proc_photo_Photo_createTonemapReinhard_10 = procLoad("Java_org_opencv_photo_Photo_createTonemapReinhard_10");
    proc_video_DualTVL1OpticalFlow_setEpsilon_10 = procLoad("Java_org_opencv_video_DualTVL1OpticalFlow_setEpsilon_10");
    proc_features2d_DescriptorExtractor_empty_10 = procLoad("Java_org_opencv_features2d_DescriptorExtractor_empty_10");
    proc_ml_EM_getTermCriteria_10 = procLoad("Java_org_opencv_ml_EM_getTermCriteria_10");
    proc_calib3d_Calib3d_recoverPose_12 = procLoad("Java_org_opencv_calib3d_Calib3d_recoverPose_12");
    proc_imgproc_Imgproc_connectedComponentsWithAlgorithm_10 = procLoad("Java_org_opencv_imgproc_Imgproc_connectedComponentsWithAlgorithm_10");
    proc_core_Mat_n_1setTo__JDDDDJ = procLoad("Java_org_opencv_core_Mat_n_1setTo__JDDDDJ");
    proc_photo_Photo_createTonemapReinhard_11 = procLoad("Java_org_opencv_photo_Photo_createTonemapReinhard_11");
    proc_dnn_Net_getMemoryConsumption_10 = procLoad("Java_org_opencv_dnn_Net_getMemoryConsumption_10");
    proc_video_DualTVL1OpticalFlow_setGamma_10 = procLoad("Java_org_opencv_video_DualTVL1OpticalFlow_setGamma_10");
    proc_calib3d_Calib3d_recoverPose_13 = procLoad("Java_org_opencv_calib3d_Calib3d_recoverPose_13");
    proc_imgproc_Imgproc_connectedComponentsWithStatsWithAlgorithm_10 = procLoad("Java_org_opencv_imgproc_Imgproc_connectedComponentsWithStatsWithAlgorithm_10");
    proc_core_Mat_n_1setTo__JJJ = procLoad("Java_org_opencv_core_Mat_n_1setTo__JJJ");
    proc_ml_EM_getWeights_10 = procLoad("Java_org_opencv_ml_EM_getWeights_10");
    proc_photo_Photo_createTonemap_10 = procLoad("Java_org_opencv_photo_Photo_createTonemap_10");
    proc_dnn_Net_getMemoryConsumption_11 = procLoad("Java_org_opencv_dnn_Net_getMemoryConsumption_11");
    proc_features2d_DescriptorExtractor_read_10 = procLoad("Java_org_opencv_features2d_DescriptorExtractor_read_10");
    proc_video_DualTVL1OpticalFlow_setInnerIterations_10 = procLoad("Java_org_opencv_video_DualTVL1OpticalFlow_setInnerIterations_10");
    proc_calib3d_Calib3d_recoverPose_14 = procLoad("Java_org_opencv_calib3d_Calib3d_recoverPose_14");
    proc_imgproc_Imgproc_connectedComponentsWithStats_10 = procLoad("Java_org_opencv_imgproc_Imgproc_connectedComponentsWithStats_10");
    proc_core_Mat_n_1setTo__JJ = procLoad("Java_org_opencv_core_Mat_n_1setTo__JJ");
    proc_ml_EM_load_10 = procLoad("Java_org_opencv_ml_EM_load_10");
    proc_photo_Photo_createTonemap_11 = procLoad("Java_org_opencv_photo_Photo_createTonemap_11");
    proc_dnn_Net_getMemoryConsumption_12 = procLoad("Java_org_opencv_dnn_Net_getMemoryConsumption_12");
    proc_features2d_DescriptorExtractor_write_10 = procLoad("Java_org_opencv_features2d_DescriptorExtractor_write_10");
    proc_video_DualTVL1OpticalFlow_setLambda_10 = procLoad("Java_org_opencv_video_DualTVL1OpticalFlow_setLambda_10");
    proc_core_Mat_n_1size = procLoad("Java_org_opencv_core_Mat_n_1size");
    proc_calib3d_Calib3d_recoverPose_15 = procLoad("Java_org_opencv_calib3d_Calib3d_recoverPose_15");
    proc_imgproc_Imgproc_connectedComponentsWithStats_11 = procLoad("Java_org_opencv_imgproc_Imgproc_connectedComponentsWithStats_11");
    proc_dnn_Net_getParam_10 = procLoad("Java_org_opencv_dnn_Net_getParam_10");
    proc_ml_EM_load_11 = procLoad("Java_org_opencv_ml_EM_load_11");
    proc_features2d_DescriptorMatcher_add_10 = procLoad("Java_org_opencv_features2d_DescriptorMatcher_add_10");
    proc_photo_Photo_decolor_10 = procLoad("Java_org_opencv_photo_Photo_decolor_10");
    proc_video_DualTVL1OpticalFlow_setMedianFiltering_10 = procLoad("Java_org_opencv_video_DualTVL1OpticalFlow_setMedianFiltering_10");
    proc_ml_EM_predict2_10 = procLoad("Java_org_opencv_ml_EM_predict2_10");
    proc_calib3d_Calib3d_recoverPose_16 = procLoad("Java_org_opencv_calib3d_Calib3d_recoverPose_16");
    proc_imgproc_Imgproc_connectedComponents_10 = procLoad("Java_org_opencv_imgproc_Imgproc_connectedComponents_10");
    proc_core_Mat_n_1step1__JI = procLoad("Java_org_opencv_core_Mat_n_1step1__JI");
    proc_dnn_Net_getParam_11 = procLoad("Java_org_opencv_dnn_Net_getParam_11");
    proc_features2d_DescriptorMatcher_clear_10 = procLoad("Java_org_opencv_features2d_DescriptorMatcher_clear_10");
    proc_photo_Photo_denoise_1TVL1_10 = procLoad("Java_org_opencv_photo_Photo_denoise_1TVL1_10");
    proc_video_DualTVL1OpticalFlow_setOuterIterations_10 = procLoad("Java_org_opencv_video_DualTVL1OpticalFlow_setOuterIterations_10");
    proc_calib3d_Calib3d_rectify3Collinear_10 = procLoad("Java_org_opencv_calib3d_Calib3d_rectify3Collinear_10");
    proc_ml_EM_predict_10 = procLoad("Java_org_opencv_ml_EM_predict_10");
    proc_imgproc_Imgproc_connectedComponents_11 = procLoad("Java_org_opencv_imgproc_Imgproc_connectedComponents_11");
    proc_core_Mat_n_1step1__J = procLoad("Java_org_opencv_core_Mat_n_1step1__J");
    proc_dnn_Net_getPerfProfile_10 = procLoad("Java_org_opencv_dnn_Net_getPerfProfile_10");
    proc_features2d_DescriptorMatcher_clone_10 = procLoad("Java_org_opencv_features2d_DescriptorMatcher_clone_10");
    proc_photo_Photo_denoise_1TVL1_11 = procLoad("Java_org_opencv_photo_Photo_denoise_1TVL1_11");
    proc_video_DualTVL1OpticalFlow_setScaleStep_10 = procLoad("Java_org_opencv_video_DualTVL1OpticalFlow_setScaleStep_10");
    proc_imgproc_Imgproc_contourArea_10 = procLoad("Java_org_opencv_imgproc_Imgproc_contourArea_10");
    proc_ml_EM_predict_11 = procLoad("Java_org_opencv_ml_EM_predict_11");
    proc_core_Mat_n_1submat = procLoad("Java_org_opencv_core_Mat_n_1submat");
    proc_dnn_Net_getUnconnectedOutLayers_10 = procLoad("Java_org_opencv_dnn_Net_getUnconnectedOutLayers_10");
    proc_features2d_DescriptorMatcher_clone_11 = procLoad("Java_org_opencv_features2d_DescriptorMatcher_clone_11");
    proc_calib3d_Calib3d_reprojectImageTo3D_10 = procLoad("Java_org_opencv_calib3d_Calib3d_reprojectImageTo3D_10");
    proc_photo_Photo_detailEnhance_10 = procLoad("Java_org_opencv_photo_Photo_detailEnhance_10");
    proc_video_DualTVL1OpticalFlow_setScalesNumber_10 = procLoad("Java_org_opencv_video_DualTVL1OpticalFlow_setScalesNumber_10");
    proc_imgproc_Imgproc_contourArea_11 = procLoad("Java_org_opencv_imgproc_Imgproc_contourArea_11");
    proc_core_Mat_n_1submat_1rr = procLoad("Java_org_opencv_core_Mat_n_1submat_1rr");
    proc_features2d_DescriptorMatcher_create_10 = procLoad("Java_org_opencv_features2d_DescriptorMatcher_create_10");
    proc_calib3d_Calib3d_reprojectImageTo3D_11 = procLoad("Java_org_opencv_calib3d_Calib3d_reprojectImageTo3D_11");
    proc_dnn_Net_setHalideScheduler_10 = procLoad("Java_org_opencv_dnn_Net_setHalideScheduler_10");
    proc_ml_EM_setClustersNumber_10 = procLoad("Java_org_opencv_ml_EM_setClustersNumber_10");
    proc_photo_Photo_detailEnhance_11 = procLoad("Java_org_opencv_photo_Photo_detailEnhance_11");
    proc_video_DualTVL1OpticalFlow_setTau_10 = procLoad("Java_org_opencv_video_DualTVL1OpticalFlow_setTau_10");
    proc_calib3d_Calib3d_RQDecomp3x3_11 = procLoad("Java_org_opencv_calib3d_Calib3d_RQDecomp3x3_11");
    proc_core_Mat_nGet = procLoad("Java_org_opencv_core_Mat_nGet");
    proc_dnn_DictValue_DictValue_10 = procLoad("Java_org_opencv_dnn_DictValue_DictValue_10");
    proc_features2d_AKAZE_create_10 = procLoad("Java_org_opencv_features2d_AKAZE_create_10");
    proc_imgcodecs_Imgcodecs_imdecode_10 = procLoad("Java_org_opencv_imgcodecs_Imgcodecs_imdecode_10");
    proc_ml_ANN_1MLP_create_10 = procLoad("Java_org_opencv_ml_ANN_1MLP_create_10");
    proc_videoio_VideoCapture_VideoCapture_10 = procLoad("Java_org_opencv_videoio_VideoCapture_VideoCapture_10");
    proc_imgproc_CLAHE_apply_10 = procLoad("Java_org_opencv_imgproc_CLAHE_apply_10");
    proc_objdetect_Objdetect_groupRectangles_10 = procLoad("Java_org_opencv_objdetect_Objdetect_groupRectangles_10");
    proc_photo_AlignExposures_delete = procLoad("Java_org_opencv_photo_AlignExposures_delete");
    proc_video_BackgroundSubtractor_apply_10 = procLoad("Java_org_opencv_video_BackgroundSubtractor_apply_10");
    proc_core_Mat_n_1t = procLoad("Java_org_opencv_core_Mat_n_1t");
    proc_features2d_DescriptorMatcher_create_11 = procLoad("Java_org_opencv_features2d_DescriptorMatcher_create_11");
    proc_calib3d_Calib3d_reprojectImageTo3D_12 = procLoad("Java_org_opencv_calib3d_Calib3d_reprojectImageTo3D_12");
    proc_dnn_Net_setInput_10 = procLoad("Java_org_opencv_dnn_Net_setInput_10");
    proc_imgproc_Imgproc_convertMaps_10 = procLoad("Java_org_opencv_imgproc_Imgproc_convertMaps_10");
    proc_ml_EM_setCovarianceMatrixType_10 = procLoad("Java_org_opencv_ml_EM_setCovarianceMatrixType_10");
    proc_photo_Photo_edgePreservingFilter_10 = procLoad("Java_org_opencv_photo_Photo_edgePreservingFilter_10");
    proc_video_DualTVL1OpticalFlow_setTheta_10 = procLoad("Java_org_opencv_video_DualTVL1OpticalFlow_setTheta_10");
    proc_calib3d_Calib3d_sampsonDistance_10 = procLoad("Java_org_opencv_calib3d_Calib3d_sampsonDistance_10");
    proc_core_Mat_n_1total = procLoad("Java_org_opencv_core_Mat_n_1total");
    proc_dnn_Net_setInput_11 = procLoad("Java_org_opencv_dnn_Net_setInput_11");
    proc_features2d_DescriptorMatcher_delete = procLoad("Java_org_opencv_features2d_DescriptorMatcher_delete");
    proc_imgproc_Imgproc_convertMaps_11 = procLoad("Java_org_opencv_imgproc_Imgproc_convertMaps_11");
    proc_ml_EM_setTermCriteria_10 = procLoad("Java_org_opencv_ml_EM_setTermCriteria_10");
    proc_photo_Photo_edgePreservingFilter_11 = procLoad("Java_org_opencv_photo_Photo_edgePreservingFilter_11");
    proc_video_DualTVL1OpticalFlow_setUseInitialFlow_10 = procLoad("Java_org_opencv_video_DualTVL1OpticalFlow_setUseInitialFlow_10");
    proc_features2d_DescriptorMatcher_empty_10 = procLoad("Java_org_opencv_features2d_DescriptorMatcher_empty_10");
    proc_ml_EM_trainEM_10 = procLoad("Java_org_opencv_ml_EM_trainEM_10");
    proc_calib3d_Calib3d_solveP3P_10 = procLoad("Java_org_opencv_calib3d_Calib3d_solveP3P_10");
    proc_core_Mat_n_1type = procLoad("Java_org_opencv_core_Mat_n_1type");
    proc_dnn_Net_setInputsNames_10 = procLoad("Java_org_opencv_dnn_Net_setInputsNames_10");
    proc_imgproc_Imgproc_convexHull_10 = procLoad("Java_org_opencv_imgproc_Imgproc_convexHull_10");
    proc_photo_Photo_fastNlMeansDenoisingColoredMulti_10 = procLoad("Java_org_opencv_photo_Photo_fastNlMeansDenoisingColoredMulti_10");
    proc_video_DualTVL1OpticalFlow_setWarpingsNumber_10 = procLoad("Java_org_opencv_video_DualTVL1OpticalFlow_setWarpingsNumber_10");
    proc_calib3d_Calib3d_solvePnPRansac_10 = procLoad("Java_org_opencv_calib3d_Calib3d_solvePnPRansac_10");
    proc_ml_EM_trainEM_11 = procLoad("Java_org_opencv_ml_EM_trainEM_11");
    proc_core_Mat_n_1zeros__III = procLoad("Java_org_opencv_core_Mat_n_1zeros__III");
    proc_features2d_DescriptorMatcher_getTrainDescriptors_10 = procLoad("Java_org_opencv_features2d_DescriptorMatcher_getTrainDescriptors_10");
    proc_video_FarnebackOpticalFlow_create_10 = procLoad("Java_org_opencv_video_FarnebackOpticalFlow_create_10");
    proc_dnn_Net_setParam_10 = procLoad("Java_org_opencv_dnn_Net_setParam_10");
    proc_imgproc_Imgproc_convexHull_11 = procLoad("Java_org_opencv_imgproc_Imgproc_convexHull_11");
    proc_photo_Photo_fastNlMeansDenoisingColoredMulti_11 = procLoad("Java_org_opencv_photo_Photo_fastNlMeansDenoisingColoredMulti_11");
    proc_calib3d_Calib3d_solvePnPRansac_11 = procLoad("Java_org_opencv_calib3d_Calib3d_solvePnPRansac_11");
    proc_features2d_DescriptorMatcher_isMaskSupported_10 = procLoad("Java_org_opencv_features2d_DescriptorMatcher_isMaskSupported_10");
    proc_ml_EM_trainE_10 = procLoad("Java_org_opencv_ml_EM_trainE_10");
    proc_core_Mat_n_1zeros__DDI = procLoad("Java_org_opencv_core_Mat_n_1zeros__DDI");
    proc_video_FarnebackOpticalFlow_create_11 = procLoad("Java_org_opencv_video_FarnebackOpticalFlow_create_11");
    proc_dnn_Net_setPreferableBackend_10 = procLoad("Java_org_opencv_dnn_Net_setPreferableBackend_10");
    proc_imgproc_Imgproc_convexityDefects_10 = procLoad("Java_org_opencv_imgproc_Imgproc_convexityDefects_10");
    proc_photo_Photo_fastNlMeansDenoisingColored_10 = procLoad("Java_org_opencv_photo_Photo_fastNlMeansDenoisingColored_10");
    proc_calib3d_Calib3d_solvePnP_10 = procLoad("Java_org_opencv_calib3d_Calib3d_solvePnP_10");
    proc_ml_EM_trainE_11 = procLoad("Java_org_opencv_ml_EM_trainE_11");
    proc_core_Core_LUT_10 = procLoad("Java_org_opencv_core_Core_LUT_10");
    proc_dnn_Net_setPreferableTarget_10 = procLoad("Java_org_opencv_dnn_Net_setPreferableTarget_10");
    proc_features2d_DescriptorMatcher_knnMatch_10 = procLoad("Java_org_opencv_features2d_DescriptorMatcher_knnMatch_10");
    proc_imgproc_Imgproc_cornerEigenValsAndVecs_10 = procLoad("Java_org_opencv_imgproc_Imgproc_cornerEigenValsAndVecs_10");
    proc_photo_Photo_fastNlMeansDenoisingColored_11 = procLoad("Java_org_opencv_photo_Photo_fastNlMeansDenoisingColored_11");
    proc_video_FarnebackOpticalFlow_delete = procLoad("Java_org_opencv_video_FarnebackOpticalFlow_delete");
    proc_calib3d_Calib3d_solvePnP_11 = procLoad("Java_org_opencv_calib3d_Calib3d_solvePnP_11");
    proc_ml_EM_trainM_10 = procLoad("Java_org_opencv_ml_EM_trainM_10");
    proc_video_FarnebackOpticalFlow_getFastPyramids_10 = procLoad("Java_org_opencv_video_FarnebackOpticalFlow_getFastPyramids_10");
    proc_core_Core_Mahalanobis_10 = procLoad("Java_org_opencv_core_Core_Mahalanobis_10");
    proc_features2d_DescriptorMatcher_knnMatch_11 = procLoad("Java_org_opencv_features2d_DescriptorMatcher_knnMatch_11");
    proc_imgproc_Imgproc_cornerEigenValsAndVecs_11 = procLoad("Java_org_opencv_imgproc_Imgproc_cornerEigenValsAndVecs_11");
    proc_photo_Photo_fastNlMeansDenoisingMulti_10 = procLoad("Java_org_opencv_photo_Photo_fastNlMeansDenoisingMulti_10");
    proc_ml_EM_trainM_11 = procLoad("Java_org_opencv_ml_EM_trainM_11");
    proc_calib3d_Calib3d_stereoCalibrate_10 = procLoad("Java_org_opencv_calib3d_Calib3d_stereoCalibrate_10");
    proc_video_FarnebackOpticalFlow_getFlags_10 = procLoad("Java_org_opencv_video_FarnebackOpticalFlow_getFlags_10");
    proc_core_Core_PCABackProject_10 = procLoad("Java_org_opencv_core_Core_PCABackProject_10");
    proc_features2d_DescriptorMatcher_knnMatch_12 = procLoad("Java_org_opencv_features2d_DescriptorMatcher_knnMatch_12");
    proc_imgproc_Imgproc_cornerHarris_10 = procLoad("Java_org_opencv_imgproc_Imgproc_cornerHarris_10");
    proc_photo_Photo_fastNlMeansDenoisingMulti_11 = procLoad("Java_org_opencv_photo_Photo_fastNlMeansDenoisingMulti_11");
    proc_calib3d_Calib3d_stereoCalibrate_11 = procLoad("Java_org_opencv_calib3d_Calib3d_stereoCalibrate_11");
    proc_video_FarnebackOpticalFlow_getNumIters_10 = procLoad("Java_org_opencv_video_FarnebackOpticalFlow_getNumIters_10");
    proc_ml_KNearest_create_10 = procLoad("Java_org_opencv_ml_KNearest_create_10");
    proc_core_Core_PCACompute_10 = procLoad("Java_org_opencv_core_Core_PCACompute_10");
    proc_features2d_DescriptorMatcher_knnMatch_13 = procLoad("Java_org_opencv_features2d_DescriptorMatcher_knnMatch_13");
    proc_imgproc_Imgproc_cornerHarris_11 = procLoad("Java_org_opencv_imgproc_Imgproc_cornerHarris_11");
    proc_photo_Photo_fastNlMeansDenoisingMulti_12 = procLoad("Java_org_opencv_photo_Photo_fastNlMeansDenoisingMulti_12");
    proc_calib3d_Calib3d_stereoCalibrate_12 = procLoad("Java_org_opencv_calib3d_Calib3d_stereoCalibrate_12");
    proc_video_FarnebackOpticalFlow_getNumLevels_10 = procLoad("Java_org_opencv_video_FarnebackOpticalFlow_getNumLevels_10");
    proc_core_Core_PCACompute_11 = procLoad("Java_org_opencv_core_Core_PCACompute_11");
    proc_features2d_DescriptorMatcher_match_10 = procLoad("Java_org_opencv_features2d_DescriptorMatcher_match_10");
    proc_imgproc_Imgproc_cornerMinEigenVal_10 = procLoad("Java_org_opencv_imgproc_Imgproc_cornerMinEigenVal_10");
    proc_ml_KNearest_delete = procLoad("Java_org_opencv_ml_KNearest_delete");
    proc_photo_Photo_fastNlMeansDenoisingMulti_13 = procLoad("Java_org_opencv_photo_Photo_fastNlMeansDenoisingMulti_13");
    proc_imgcodecs_Imgcodecs_imencode_10 = procLoad("Java_org_opencv_imgcodecs_Imgcodecs_imencode_10");
    proc_core_Mat_nGetB = procLoad("Java_org_opencv_core_Mat_nGetB");
    proc_dnn_DictValue_DictValue_11 = procLoad("Java_org_opencv_dnn_DictValue_DictValue_11");
    proc_features2d_AKAZE_create_11 = procLoad("Java_org_opencv_features2d_AKAZE_create_11");
    proc_videoio_VideoCapture_VideoCapture_11 = procLoad("Java_org_opencv_videoio_VideoCapture_VideoCapture_11");
    proc_calib3d_Calib3d_Rodrigues_10 = procLoad("Java_org_opencv_calib3d_Calib3d_Rodrigues_10");
    proc_imgproc_CLAHE_collectGarbage_10 = procLoad("Java_org_opencv_imgproc_CLAHE_collectGarbage_10");
    proc_ml_ANN_1MLP_delete = procLoad("Java_org_opencv_ml_ANN_1MLP_delete");
    proc_objdetect_Objdetect_groupRectangles_11 = procLoad("Java_org_opencv_objdetect_Objdetect_groupRectangles_11");
    proc_photo_AlignExposures_process_10 = procLoad("Java_org_opencv_photo_AlignExposures_process_10");
    proc_video_BackgroundSubtractor_apply_11 = procLoad("Java_org_opencv_video_BackgroundSubtractor_apply_11");
    proc_calib3d_Calib3d_stereoCalibrate_13 = procLoad("Java_org_opencv_calib3d_Calib3d_stereoCalibrate_13");
    proc_ml_KNearest_findNearest_10 = procLoad("Java_org_opencv_ml_KNearest_findNearest_10");
    proc_video_FarnebackOpticalFlow_getPolyN_10 = procLoad("Java_org_opencv_video_FarnebackOpticalFlow_getPolyN_10");
    proc_core_Core_PCACompute_12 = procLoad("Java_org_opencv_core_Core_PCACompute_12");
    proc_features2d_DescriptorMatcher_match_11 = procLoad("Java_org_opencv_features2d_DescriptorMatcher_match_11");
    proc_imgproc_Imgproc_cornerMinEigenVal_11 = procLoad("Java_org_opencv_imgproc_Imgproc_cornerMinEigenVal_11");
    proc_photo_Photo_fastNlMeansDenoising_10 = procLoad("Java_org_opencv_photo_Photo_fastNlMeansDenoising_10");
    proc_calib3d_Calib3d_stereoCalibrate_14 = procLoad("Java_org_opencv_calib3d_Calib3d_stereoCalibrate_14");
    proc_video_FarnebackOpticalFlow_getPolySigma_10 = procLoad("Java_org_opencv_video_FarnebackOpticalFlow_getPolySigma_10");
    proc_ml_KNearest_findNearest_11 = procLoad("Java_org_opencv_ml_KNearest_findNearest_11");
    proc_core_Core_PCAProject_10 = procLoad("Java_org_opencv_core_Core_PCAProject_10");
    proc_features2d_DescriptorMatcher_match_12 = procLoad("Java_org_opencv_features2d_DescriptorMatcher_match_12");
    proc_imgproc_Imgproc_cornerMinEigenVal_12 = procLoad("Java_org_opencv_imgproc_Imgproc_cornerMinEigenVal_12");
    proc_photo_Photo_fastNlMeansDenoising_11 = procLoad("Java_org_opencv_photo_Photo_fastNlMeansDenoising_11");
    proc_calib3d_Calib3d_stereoCalibrate_15 = procLoad("Java_org_opencv_calib3d_Calib3d_stereoCalibrate_15");
    proc_core_Core_PSNR_10 = procLoad("Java_org_opencv_core_Core_PSNR_10");
    proc_video_FarnebackOpticalFlow_getPyrScale_10 = procLoad("Java_org_opencv_video_FarnebackOpticalFlow_getPyrScale_10");
    proc_ml_KNearest_getAlgorithmType_10 = procLoad("Java_org_opencv_ml_KNearest_getAlgorithmType_10");
    proc_features2d_DescriptorMatcher_match_13 = procLoad("Java_org_opencv_features2d_DescriptorMatcher_match_13");
    proc_imgproc_Imgproc_cornerSubPix_10 = procLoad("Java_org_opencv_imgproc_Imgproc_cornerSubPix_10");
    proc_photo_Photo_fastNlMeansDenoising_12 = procLoad("Java_org_opencv_photo_Photo_fastNlMeansDenoising_12");
    proc_calib3d_Calib3d_stereoRectifyUncalibrated_10 = procLoad("Java_org_opencv_calib3d_Calib3d_stereoRectifyUncalibrated_10");
    proc_ml_KNearest_getDefaultK_10 = procLoad("Java_org_opencv_ml_KNearest_getDefaultK_10");
    proc_video_FarnebackOpticalFlow_getWinSize_10 = procLoad("Java_org_opencv_video_FarnebackOpticalFlow_getWinSize_10");
    proc_imgproc_Imgproc_createCLAHE_10 = procLoad("Java_org_opencv_imgproc_Imgproc_createCLAHE_10");
    proc_core_Core_SVBackSubst_10 = procLoad("Java_org_opencv_core_Core_SVBackSubst_10");
    proc_features2d_DescriptorMatcher_radiusMatch_10 = procLoad("Java_org_opencv_features2d_DescriptorMatcher_radiusMatch_10");
    proc_photo_Photo_fastNlMeansDenoising_13 = procLoad("Java_org_opencv_photo_Photo_fastNlMeansDenoising_13");
    proc_calib3d_Calib3d_stereoRectifyUncalibrated_11 = procLoad("Java_org_opencv_calib3d_Calib3d_stereoRectifyUncalibrated_11");
    proc_ml_KNearest_getEmax_10 = procLoad("Java_org_opencv_ml_KNearest_getEmax_10");
    proc_imgproc_Imgproc_createCLAHE_11 = procLoad("Java_org_opencv_imgproc_Imgproc_createCLAHE_11");
    proc_core_Core_SVDecomp_10 = procLoad("Java_org_opencv_core_Core_SVDecomp_10");
    proc_features2d_DescriptorMatcher_radiusMatch_11 = procLoad("Java_org_opencv_features2d_DescriptorMatcher_radiusMatch_11");
    proc_photo_Photo_illuminationChange_10 = procLoad("Java_org_opencv_photo_Photo_illuminationChange_10");
    proc_video_FarnebackOpticalFlow_setFastPyramids_10 = procLoad("Java_org_opencv_video_FarnebackOpticalFlow_setFastPyramids_10");
    proc_ml_KNearest_getIsClassifier_10 = procLoad("Java_org_opencv_ml_KNearest_getIsClassifier_10");
    proc_calib3d_Calib3d_stereoRectify_10 = procLoad("Java_org_opencv_calib3d_Calib3d_stereoRectify_10");
    proc_core_Core_SVDecomp_11 = procLoad("Java_org_opencv_core_Core_SVDecomp_11");
    proc_features2d_DescriptorMatcher_radiusMatch_12 = procLoad("Java_org_opencv_features2d_DescriptorMatcher_radiusMatch_12");
    proc_imgproc_Imgproc_createHanningWindow_10 = procLoad("Java_org_opencv_imgproc_Imgproc_createHanningWindow_10");
    proc_photo_Photo_illuminationChange_11 = procLoad("Java_org_opencv_photo_Photo_illuminationChange_11");
    proc_video_FarnebackOpticalFlow_setFlags_10 = procLoad("Java_org_opencv_video_FarnebackOpticalFlow_setFlags_10");
    proc_imgproc_Imgproc_createLineSegmentDetector_10 = procLoad("Java_org_opencv_imgproc_Imgproc_createLineSegmentDetector_10");
    proc_calib3d_Calib3d_stereoRectify_11 = procLoad("Java_org_opencv_calib3d_Calib3d_stereoRectify_11");
    proc_core_Core_absdiff_10 = procLoad("Java_org_opencv_core_Core_absdiff_10");
    proc_features2d_DescriptorMatcher_radiusMatch_13 = procLoad("Java_org_opencv_features2d_DescriptorMatcher_radiusMatch_13");
    proc_ml_KNearest_setAlgorithmType_10 = procLoad("Java_org_opencv_ml_KNearest_setAlgorithmType_10");
    proc_photo_Photo_inpaint_10 = procLoad("Java_org_opencv_photo_Photo_inpaint_10");
    proc_video_FarnebackOpticalFlow_setNumIters_10 = procLoad("Java_org_opencv_video_FarnebackOpticalFlow_setNumIters_10");
    proc_imgproc_Imgproc_createLineSegmentDetector_11 = procLoad("Java_org_opencv_imgproc_Imgproc_createLineSegmentDetector_11");
    proc_calib3d_Calib3d_stereoRectify_12 = procLoad("Java_org_opencv_calib3d_Calib3d_stereoRectify_12");
    proc_core_Core_absdiff_11 = procLoad("Java_org_opencv_core_Core_absdiff_11");
    proc_features2d_DescriptorMatcher_read_10 = procLoad("Java_org_opencv_features2d_DescriptorMatcher_read_10");
    proc_ml_KNearest_setDefaultK_10 = procLoad("Java_org_opencv_ml_KNearest_setDefaultK_10");
    proc_photo_Photo_pencilSketch_10 = procLoad("Java_org_opencv_photo_Photo_pencilSketch_10");
    proc_video_FarnebackOpticalFlow_setNumLevels_10 = procLoad("Java_org_opencv_video_FarnebackOpticalFlow_setNumLevels_10");
    proc_calib3d_Calib3d_stereoRectify_13 = procLoad("Java_org_opencv_calib3d_Calib3d_stereoRectify_13");
    proc_core_Core_addWeighted_10 = procLoad("Java_org_opencv_core_Core_addWeighted_10");
    proc_features2d_DescriptorMatcher_train_10 = procLoad("Java_org_opencv_features2d_DescriptorMatcher_train_10");
    proc_imgproc_Imgproc_cvtColor_10 = procLoad("Java_org_opencv_imgproc_Imgproc_cvtColor_10");
    proc_ml_KNearest_setEmax_10 = procLoad("Java_org_opencv_ml_KNearest_setEmax_10");
    proc_photo_Photo_pencilSketch_11 = procLoad("Java_org_opencv_photo_Photo_pencilSketch_11");
    proc_video_FarnebackOpticalFlow_setPolyN_10 = procLoad("Java_org_opencv_video_FarnebackOpticalFlow_setPolyN_10");
    proc_calib3d_Calib3d_triangulatePoints_10 = procLoad("Java_org_opencv_calib3d_Calib3d_triangulatePoints_10");
    proc_core_Core_addWeighted_11 = procLoad("Java_org_opencv_core_Core_addWeighted_11");
    proc_features2d_DescriptorMatcher_write_10 = procLoad("Java_org_opencv_features2d_DescriptorMatcher_write_10");
    proc_imgproc_Imgproc_cvtColor_11 = procLoad("Java_org_opencv_imgproc_Imgproc_cvtColor_11");
    proc_ml_KNearest_setIsClassifier_10 = procLoad("Java_org_opencv_ml_KNearest_setIsClassifier_10");
    proc_photo_Photo_seamlessClone_10 = procLoad("Java_org_opencv_photo_Photo_seamlessClone_10");
    proc_video_FarnebackOpticalFlow_setPolySigma_10 = procLoad("Java_org_opencv_video_FarnebackOpticalFlow_setPolySigma_10");
    proc_imgcodecs_Imgcodecs_imencode_11 = procLoad("Java_org_opencv_imgcodecs_Imgcodecs_imencode_11");
    proc_ml_ANN_1MLP_getBackpropMomentumScale_10 = procLoad("Java_org_opencv_ml_ANN_1MLP_getBackpropMomentumScale_10");
    proc_photo_AlignMTB_calculateShift_10 = procLoad("Java_org_opencv_photo_AlignMTB_calculateShift_10");
    proc_core_Mat_nGetD = procLoad("Java_org_opencv_core_Mat_nGetD");
    proc_dnn_DictValue_DictValue_12 = procLoad("Java_org_opencv_dnn_DictValue_DictValue_12");
    proc_objdetect_HOGDescriptor_HOGDescriptor_10 = procLoad("Java_org_opencv_objdetect_HOGDescriptor_HOGDescriptor_10");
    proc_videoio_VideoCapture_VideoCapture_12 = procLoad("Java_org_opencv_videoio_VideoCapture_VideoCapture_12");
    proc_calib3d_Calib3d_Rodrigues_11 = procLoad("Java_org_opencv_calib3d_Calib3d_Rodrigues_11");
    proc_features2d_AKAZE_delete = procLoad("Java_org_opencv_features2d_AKAZE_delete");
    proc_imgproc_CLAHE_delete = procLoad("Java_org_opencv_imgproc_CLAHE_delete");
    proc_video_BackgroundSubtractor_delete = procLoad("Java_org_opencv_video_BackgroundSubtractor_delete");
	return failCnt;
}

#else

void setLibPath(const char* path) {}
int  InitProcs() { return 0; }

#endif
