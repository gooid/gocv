package core

/*

#include "jni.h"

extern void Java_org_opencv_objdetect_Objdetect_groupRectangles_10(JNIEnv*, jclass, jlong, jlong, jint, jdouble);
extern void Java_org_opencv_objdetect_Objdetect_groupRectangles_11(JNIEnv*, jclass, jlong, jlong, jint);
extern jlong Java_org_opencv_objdetect_HOGDescriptor_HOGDescriptor_10(JNIEnv*, jclass, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jint, jint, jdouble, jint, jdouble, jboolean, jint, jboolean);
extern jlong Java_org_opencv_objdetect_HOGDescriptor_HOGDescriptor_11(JNIEnv*, jclass, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jint);
extern jlong Java_org_opencv_objdetect_HOGDescriptor_HOGDescriptor_12(JNIEnv*, jclass, jstring);
extern jlong Java_org_opencv_objdetect_HOGDescriptor_HOGDescriptor_13(JNIEnv*, jclass);
extern jboolean Java_org_opencv_objdetect_HOGDescriptor_checkDetectorSize_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_objdetect_HOGDescriptor_computeGradient_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jdouble, jdouble, jdouble, jdouble);
extern void Java_org_opencv_objdetect_HOGDescriptor_computeGradient_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_objdetect_HOGDescriptor_compute_10(JNIEnv*, jclass, jlong, jlong, jlong, jdouble, jdouble, jdouble, jdouble, jlong);
extern void Java_org_opencv_objdetect_HOGDescriptor_compute_11(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_objdetect_HOGDescriptor_delete(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_objdetect_HOGDescriptor_detectMultiScale_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jboolean);
extern void Java_org_opencv_objdetect_HOGDescriptor_detectMultiScale_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_objdetect_HOGDescriptor_detect_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jdouble, jdouble, jdouble, jdouble, jdouble, jlong);
extern void Java_org_opencv_objdetect_HOGDescriptor_detect_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern jlong Java_org_opencv_objdetect_HOGDescriptor_getDaimlerPeopleDetector_10(JNIEnv*, jclass);
extern jlong Java_org_opencv_objdetect_HOGDescriptor_getDefaultPeopleDetector_10(JNIEnv*, jclass);
extern jlong Java_org_opencv_objdetect_HOGDescriptor_getDescriptorSize_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_objdetect_HOGDescriptor_getWinSigma_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_objdetect_HOGDescriptor_get_1L2HysThreshold_10(JNIEnv*, jclass, jlong);
extern jdoubleArray Java_org_opencv_objdetect_HOGDescriptor_get_1blockSize_10(JNIEnv*, jclass, jlong);
extern jdoubleArray Java_org_opencv_objdetect_HOGDescriptor_get_1blockStride_10(JNIEnv*, jclass, jlong);
extern jdoubleArray Java_org_opencv_objdetect_HOGDescriptor_get_1cellSize_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_objdetect_HOGDescriptor_get_1derivAperture_10(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_objdetect_HOGDescriptor_get_1gammaCorrection_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_objdetect_HOGDescriptor_get_1histogramNormType_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_objdetect_HOGDescriptor_get_1nbins_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_objdetect_HOGDescriptor_get_1nlevels_10(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_objdetect_HOGDescriptor_get_1signedGradient_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_objdetect_HOGDescriptor_get_1svmDetector_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_objdetect_HOGDescriptor_get_1winSigma_10(JNIEnv*, jclass, jlong);
extern jdoubleArray Java_org_opencv_objdetect_HOGDescriptor_get_1winSize_10(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_objdetect_HOGDescriptor_load_10(JNIEnv*, jclass, jlong, jstring, jstring);
extern jboolean Java_org_opencv_objdetect_HOGDescriptor_load_11(JNIEnv*, jclass, jlong, jstring);
extern void Java_org_opencv_objdetect_HOGDescriptor_save_10(JNIEnv*, jclass, jlong, jstring, jstring);
extern void Java_org_opencv_objdetect_HOGDescriptor_save_11(JNIEnv*, jclass, jlong, jstring);
extern void Java_org_opencv_objdetect_HOGDescriptor_setSVMDetector_10(JNIEnv*, jclass, jlong, jlong);
extern jlong Java_org_opencv_objdetect_CascadeClassifier_CascadeClassifier_10(JNIEnv*, jclass, jstring);
extern jlong Java_org_opencv_objdetect_CascadeClassifier_CascadeClassifier_11(JNIEnv*, jclass);
extern jboolean Java_org_opencv_objdetect_CascadeClassifier_convert_10(JNIEnv*, jclass, jstring, jstring);
extern void Java_org_opencv_objdetect_CascadeClassifier_delete(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_objdetect_CascadeClassifier_detectMultiScale2_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jdouble, jint, jint, jdouble, jdouble, jdouble, jdouble);
extern void Java_org_opencv_objdetect_CascadeClassifier_detectMultiScale2_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_objdetect_CascadeClassifier_detectMultiScale3_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jdouble, jint, jint, jdouble, jdouble, jdouble, jdouble, jboolean);
extern void Java_org_opencv_objdetect_CascadeClassifier_detectMultiScale3_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_objdetect_CascadeClassifier_detectMultiScale_10(JNIEnv*, jclass, jlong, jlong, jlong, jdouble, jint, jint, jdouble, jdouble, jdouble, jdouble);
extern void Java_org_opencv_objdetect_CascadeClassifier_detectMultiScale_11(JNIEnv*, jclass, jlong, jlong, jlong);
extern jboolean Java_org_opencv_objdetect_CascadeClassifier_empty_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_objdetect_CascadeClassifier_getFeatureType_10(JNIEnv*, jclass, jlong);
extern jdoubleArray Java_org_opencv_objdetect_CascadeClassifier_getOriginalWindowSize_10(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_objdetect_CascadeClassifier_isOldFormatCascade_10(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_objdetect_CascadeClassifier_load_10(JNIEnv*, jclass, jlong, jstring);
extern void Java_org_opencv_objdetect_BaseCascadeClassifier_delete(JNIEnv*, jclass, jlong);

*/
import "C"

func ObjdetectNative_groupRectangles_0(rectList_mat_nativeObj int64, weights_mat_nativeObj int64, groupThreshold int, eps float64) {
	C.Java_org_opencv_objdetect_Objdetect_groupRectangles_10(clzEnv, clzObj, (C.jlong)(rectList_mat_nativeObj), (C.jlong)(weights_mat_nativeObj), (C.jint)(groupThreshold), (C.jdouble)(eps))
}
func ObjdetectNative_groupRectangles_1(rectList_mat_nativeObj int64, weights_mat_nativeObj int64, groupThreshold int) {
	C.Java_org_opencv_objdetect_Objdetect_groupRectangles_11(clzEnv, clzObj, (C.jlong)(rectList_mat_nativeObj), (C.jlong)(weights_mat_nativeObj), (C.jint)(groupThreshold))
}

func HOGDescriptorNative_HOGDescriptor_0(_winSize_width float64, _winSize_height float64, _blockSize_width float64, _blockSize_height float64, _blockStride_width float64, _blockStride_height float64, _cellSize_width float64, _cellSize_height float64, _nbins int, _derivAperture int, _winSigma float64, _histogramNormType int, _L2HysThreshold float64, _gammaCorrection bool, _nlevels int, _signedGradient bool) int64 {
	return int64(C.Java_org_opencv_objdetect_HOGDescriptor_HOGDescriptor_10(clzEnv, clzObj, (C.jdouble)(_winSize_width), (C.jdouble)(_winSize_height),
		(C.jdouble)(_blockSize_width), (C.jdouble)(_blockSize_height), (C.jdouble)(_blockStride_width), (C.jdouble)(_blockStride_height), (C.jdouble)(_cellSize_width), (C.jdouble)(_cellSize_height), (C.jint)(_nbins), (C.jint)(_derivAperture), (C.jdouble)(_winSigma), (C.jint)(_histogramNormType), (C.jdouble)(_L2HysThreshold), tojboolean(_gammaCorrection), (C.jint)(_nlevels), tojboolean(_signedGradient)))
}

func HOGDescriptorNative_HOGDescriptor_1(_winSize_width float64, _winSize_height float64, _blockSize_width float64, _blockSize_height float64, _blockStride_width float64, _blockStride_height float64, _cellSize_width float64, _cellSize_height float64, _nbins int) int64 {
	return int64(C.Java_org_opencv_objdetect_HOGDescriptor_HOGDescriptor_11(clzEnv, clzObj, (C.jdouble)(_winSize_width), (C.jdouble)(_winSize_height),
		(C.jdouble)(_blockSize_width), (C.jdouble)(_blockSize_height), (C.jdouble)(_blockStride_width), (C.jdouble)(_blockStride_height), (C.jdouble)(_cellSize_width), (C.jdouble)(_cellSize_height), (C.jint)(_nbins)))
}

func HOGDescriptorNative_HOGDescriptor_2(filename string) int64 {
	defer ungointerface(filename)
	return int64(C.Java_org_opencv_objdetect_HOGDescriptor_HOGDescriptor_12(clzEnv, clzObj, tojstring(filename)))
}

func HOGDescriptorNative_HOGDescriptor_3() int64 {
	return int64(C.Java_org_opencv_objdetect_HOGDescriptor_HOGDescriptor_13(clzEnv, clzObj))
}

func HOGDescriptorNative_checkDetectorSize_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_objdetect_HOGDescriptor_checkDetectorSize_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func HOGDescriptorNative_computeGradient_0(nativeObj int64, img_nativeObj int64, grad_nativeObj int64, angleOfs_nativeObj int64, paddingTL_width float64, paddingTL_height float64, paddingBR_width float64, paddingBR_height float64) {
	C.Java_org_opencv_objdetect_HOGDescriptor_computeGradient_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(img_nativeObj), (C.jlong)(grad_nativeObj),
		(C.jlong)(angleOfs_nativeObj), (C.jdouble)(paddingTL_width), (C.jdouble)(paddingTL_height), (C.jdouble)(paddingBR_width), (C.jdouble)(paddingBR_height))
}

func HOGDescriptorNative_computeGradient_1(nativeObj int64, img_nativeObj int64, grad_nativeObj int64, angleOfs_nativeObj int64) {
	C.Java_org_opencv_objdetect_HOGDescriptor_computeGradient_11(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(img_nativeObj), (C.jlong)(grad_nativeObj), (C.jlong)(angleOfs_nativeObj))
}

func HOGDescriptorNative_compute_0(nativeObj int64, img_nativeObj int64, descriptors_mat_nativeObj int64, winStride_width float64, winStride_height float64, padding_width float64, padding_height float64, locations_mat_nativeObj int64) {
	C.Java_org_opencv_objdetect_HOGDescriptor_compute_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(img_nativeObj), (C.jlong)(descriptors_mat_nativeObj), (C.jdouble)(winStride_width), (C.jdouble)(winStride_height), (C.jdouble)(padding_width), (C.jdouble)(padding_height), (C.jlong)(locations_mat_nativeObj))
}

func HOGDescriptorNative_compute_1(nativeObj int64, img_nativeObj int64, descriptors_mat_nativeObj int64) {
	C.Java_org_opencv_objdetect_HOGDescriptor_compute_11(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(img_nativeObj), (C.jlong)(descriptors_mat_nativeObj))
}

func HOGDescriptorNative_delete(nativeObj int64) {
	C.Java_org_opencv_objdetect_HOGDescriptor_delete(clzEnv, clzObj, (C.jlong)(nativeObj))
}

func HOGDescriptorNative_detectMultiScale_0(nativeObj int64, img_nativeObj int64, foundLocations_mat_nativeObj int64, foundWeights_mat_nativeObj int64, hitThreshold float64, winStride_width float64, winStride_height float64, padding_width float64, padding_height float64, scale float64, finalThreshold float64, useMeanshiftGrouping bool) {
	C.Java_org_opencv_objdetect_HOGDescriptor_detectMultiScale_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(img_nativeObj), (C.jlong)(foundLocations_mat_nativeObj), (C.jlong)(foundWeights_mat_nativeObj), (C.jdouble)(hitThreshold), (C.jdouble)(winStride_width), (C.jdouble)(winStride_height), (C.jdouble)(padding_width), (C.jdouble)(padding_height), (C.jdouble)(scale), (C.jdouble)(finalThreshold), tojboolean(useMeanshiftGrouping))
}

func HOGDescriptorNative_detectMultiScale_1(nativeObj int64, img_nativeObj int64, foundLocations_mat_nativeObj int64, foundWeights_mat_nativeObj int64) {
	C.Java_org_opencv_objdetect_HOGDescriptor_detectMultiScale_11(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(img_nativeObj), (C.jlong)(foundLocations_mat_nativeObj), (C.jlong)(foundWeights_mat_nativeObj))
}

func HOGDescriptorNative_detect_0(nativeObj int64, img_nativeObj int64, foundLocations_mat_nativeObj int64, weights_mat_nativeObj int64, hitThreshold float64, winStride_width float64, winStride_height float64, padding_width float64, padding_height float64, searchLocations_mat_nativeObj int64) {
	C.Java_org_opencv_objdetect_HOGDescriptor_detect_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(img_nativeObj), (C.jlong)(foundLocations_mat_nativeObj), (C.jlong)(weights_mat_nativeObj), (C.jdouble)(hitThreshold), (C.jdouble)(winStride_width), (C.jdouble)(winStride_height), (C.jdouble)(padding_width), (C.jdouble)(padding_height), (C.jlong)(searchLocations_mat_nativeObj))
}

func HOGDescriptorNative_detect_1(nativeObj int64, img_nativeObj int64, foundLocations_mat_nativeObj int64, weights_mat_nativeObj int64) {
	C.Java_org_opencv_objdetect_HOGDescriptor_detect_11(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(img_nativeObj), (C.jlong)(foundLocations_mat_nativeObj), (C.jlong)(weights_mat_nativeObj))
}

func HOGDescriptorNative_getDaimlerPeopleDetector_0() int64 {
	return int64(C.Java_org_opencv_objdetect_HOGDescriptor_getDaimlerPeopleDetector_10(clzEnv, clzObj))
}

func HOGDescriptorNative_getDefaultPeopleDetector_0() int64 {
	return int64(C.Java_org_opencv_objdetect_HOGDescriptor_getDefaultPeopleDetector_10(clzEnv, clzObj))
}

func HOGDescriptorNative_getDescriptorSize_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_objdetect_HOGDescriptor_getDescriptorSize_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func HOGDescriptorNative_getWinSigma_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_objdetect_HOGDescriptor_getWinSigma_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func HOGDescriptorNative_get_L2HysThreshold_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_objdetect_HOGDescriptor_get_1L2HysThreshold_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func HOGDescriptorNative_get_blockSize_0(nativeObj int64) []float64 {
	return togofloat64Array(C.Java_org_opencv_objdetect_HOGDescriptor_get_1blockSize_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func HOGDescriptorNative_get_blockStride_0(nativeObj int64) []float64 {
	return togofloat64Array(C.Java_org_opencv_objdetect_HOGDescriptor_get_1blockStride_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func HOGDescriptorNative_get_cellSize_0(nativeObj int64) []float64 {
	return togofloat64Array(C.Java_org_opencv_objdetect_HOGDescriptor_get_1cellSize_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func HOGDescriptorNative_get_derivAperture_0(nativeObj int64) int {
	return int(C.Java_org_opencv_objdetect_HOGDescriptor_get_1derivAperture_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func HOGDescriptorNative_get_gammaCorrection_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_objdetect_HOGDescriptor_get_1gammaCorrection_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func HOGDescriptorNative_get_histogramNormType_0(nativeObj int64) int {
	return int(C.Java_org_opencv_objdetect_HOGDescriptor_get_1histogramNormType_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func HOGDescriptorNative_get_nbins_0(nativeObj int64) int {
	return int(C.Java_org_opencv_objdetect_HOGDescriptor_get_1nbins_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func HOGDescriptorNative_get_nlevels_0(nativeObj int64) int {
	return int(C.Java_org_opencv_objdetect_HOGDescriptor_get_1nlevels_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func HOGDescriptorNative_get_signedGradient_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_objdetect_HOGDescriptor_get_1signedGradient_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func HOGDescriptorNative_get_svmDetector_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_objdetect_HOGDescriptor_get_1svmDetector_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func HOGDescriptorNative_get_winSigma_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_objdetect_HOGDescriptor_get_1winSigma_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func HOGDescriptorNative_get_winSize_0(nativeObj int64) []float64 {
	return togofloat64Array(C.Java_org_opencv_objdetect_HOGDescriptor_get_1winSize_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func HOGDescriptorNative_load_0(nativeObj int64, filename string, objname string) bool {
	defer ungointerface(filename)
	defer ungointerface(objname)
	return togobool(C.Java_org_opencv_objdetect_HOGDescriptor_load_10(clzEnv, clzObj, (C.jlong)(nativeObj), tojstring(filename), tojstring(objname)))
}

func HOGDescriptorNative_load_1(nativeObj int64, filename string) bool {
	defer ungointerface(filename)
	return togobool(C.Java_org_opencv_objdetect_HOGDescriptor_load_11(clzEnv, clzObj, (C.jlong)(nativeObj), tojstring(filename)))
}

func HOGDescriptorNative_save_0(nativeObj int64, filename string, objname string) {
	defer ungointerface(filename)
	defer ungointerface(objname)
	C.Java_org_opencv_objdetect_HOGDescriptor_save_10(clzEnv, clzObj, (C.jlong)(nativeObj), tojstring(filename), tojstring(objname))
}

func HOGDescriptorNative_save_1(nativeObj int64, filename string) {
	defer ungointerface(filename)
	C.Java_org_opencv_objdetect_HOGDescriptor_save_11(clzEnv, clzObj, (C.jlong)(nativeObj), tojstring(filename))
}

func HOGDescriptorNative_setSVMDetector_0(nativeObj int64, _svmdetector_nativeObj int64) {
	C.Java_org_opencv_objdetect_HOGDescriptor_setSVMDetector_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(_svmdetector_nativeObj))
}

func CascadeClassifierNative_CascadeClassifier_0(filename string) int64 {
	defer ungointerface(filename)
	return int64(C.Java_org_opencv_objdetect_CascadeClassifier_CascadeClassifier_10(clzEnv, clzObj, tojstring(filename)))
}

func CascadeClassifierNative_CascadeClassifier_1() int64 {
	return int64(C.Java_org_opencv_objdetect_CascadeClassifier_CascadeClassifier_11(clzEnv, clzObj))
}

func CascadeClassifierNative_convert_0(oldcascade string, newcascade string) bool {
	defer ungointerface(oldcascade)
	defer ungointerface(newcascade)
	return togobool(C.Java_org_opencv_objdetect_CascadeClassifier_convert_10(clzEnv, clzObj, tojstring(oldcascade), tojstring(newcascade)))
}

func CascadeClassifierNative_delete(nativeObj int64) {
	C.Java_org_opencv_objdetect_CascadeClassifier_delete(clzEnv, clzObj, (C.jlong)(nativeObj))
}

func CascadeClassifierNative_detectMultiScale2_0(nativeObj int64, image_nativeObj int64, objects_mat_nativeObj int64, numDetections_mat_nativeObj int64, scaleFactor float64, minNeighbors int, flags int, minSize_width float64, minSize_height float64, maxSize_width float64, maxSize_height float64) {
	C.Java_org_opencv_objdetect_CascadeClassifier_detectMultiScale2_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(image_nativeObj), (C.jlong)(objects_mat_nativeObj), (C.jlong)(numDetections_mat_nativeObj), (C.jdouble)(scaleFactor), (C.jint)(minNeighbors), (C.jint)(flags),
		(C.jdouble)(minSize_width), (C.jdouble)(minSize_height), (C.jdouble)(maxSize_width), (C.jdouble)(maxSize_height))
}

func CascadeClassifierNative_detectMultiScale2_1(nativeObj int64, image_nativeObj int64, objects_mat_nativeObj int64, numDetections_mat_nativeObj int64) {
	C.Java_org_opencv_objdetect_CascadeClassifier_detectMultiScale2_11(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(image_nativeObj), (C.jlong)(objects_mat_nativeObj), (C.jlong)(numDetections_mat_nativeObj))
}

func CascadeClassifierNative_detectMultiScale3_0(nativeObj int64, image_nativeObj int64, objects_mat_nativeObj int64, rejectLevels_mat_nativeObj int64, levelWeights_mat_nativeObj int64, scaleFactor float64, minNeighbors int, flags int, minSize_width float64, minSize_height float64, maxSize_width float64, maxSize_height float64, outputRejectLevels bool) {
	C.Java_org_opencv_objdetect_CascadeClassifier_detectMultiScale3_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(image_nativeObj), (C.jlong)(objects_mat_nativeObj), (C.jlong)(rejectLevels_mat_nativeObj), (C.jlong)(levelWeights_mat_nativeObj), (C.jdouble)(scaleFactor), (C.jint)(minNeighbors), (C.jint)(flags), (C.jdouble)(minSize_width), (C.jdouble)(minSize_height), (C.jdouble)(maxSize_width), (C.jdouble)(maxSize_height), tojboolean(outputRejectLevels))
}

func CascadeClassifierNative_detectMultiScale3_1(nativeObj int64, image_nativeObj int64, objects_mat_nativeObj int64, rejectLevels_mat_nativeObj int64, levelWeights_mat_nativeObj int64) {
	C.Java_org_opencv_objdetect_CascadeClassifier_detectMultiScale3_11(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(image_nativeObj), (C.jlong)(objects_mat_nativeObj), (C.jlong)(rejectLevels_mat_nativeObj), (C.jlong)(levelWeights_mat_nativeObj))
}

func CascadeClassifierNative_detectMultiScale_0(nativeObj int64, image_nativeObj int64, objects_mat_nativeObj int64, scaleFactor float64, minNeighbors int, flags int, minSize_width float64, minSize_height float64, maxSize_width float64, maxSize_height float64) {
	C.Java_org_opencv_objdetect_CascadeClassifier_detectMultiScale_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(image_nativeObj), (C.jlong)(objects_mat_nativeObj), (C.jdouble)(scaleFactor), (C.jint)(minNeighbors), (C.jint)(flags), (C.jdouble)(minSize_width), (C.jdouble)(minSize_height), (C.jdouble)(maxSize_width), (C.jdouble)(maxSize_height))
}

func CascadeClassifierNative_detectMultiScale_1(nativeObj int64, image_nativeObj int64, objects_mat_nativeObj int64) {
	C.Java_org_opencv_objdetect_CascadeClassifier_detectMultiScale_11(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(image_nativeObj), (C.jlong)(objects_mat_nativeObj))
}

func CascadeClassifierNative_empty_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_objdetect_CascadeClassifier_empty_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func CascadeClassifierNative_getFeatureType_0(nativeObj int64) int {
	return int(C.Java_org_opencv_objdetect_CascadeClassifier_getFeatureType_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func CascadeClassifierNative_getOriginalWindowSize_0(nativeObj int64) []float64 {
	return togofloat64Array(C.Java_org_opencv_objdetect_CascadeClassifier_getOriginalWindowSize_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func CascadeClassifierNative_isOldFormatCascade_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_objdetect_CascadeClassifier_isOldFormatCascade_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func CascadeClassifierNative_load_0(nativeObj int64, filename string) bool {
	defer ungointerface(filename)
	return togobool(C.Java_org_opencv_objdetect_CascadeClassifier_load_10(clzEnv, clzObj, (C.jlong)(nativeObj), tojstring(filename)))
}

func BaseCascadeClassifierNative_delete(nativeObj int64) {
	C.Java_org_opencv_objdetect_BaseCascadeClassifier_delete(clzEnv, clzObj, (C.jlong)(nativeObj))
}
