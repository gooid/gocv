package core

/*

#include "jni.h"

extern void Java_org_opencv_photo_AlignExposures_delete(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_photo_AlignExposures_process_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong);
extern jdoubleArray Java_org_opencv_photo_AlignMTB_calculateShift_10(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_photo_AlignMTB_computeBitmaps_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_photo_AlignMTB_delete(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_photo_AlignMTB_getCut_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_photo_AlignMTB_getExcludeRange_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_photo_AlignMTB_getMaxBits_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_photo_AlignMTB_process_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_photo_AlignMTB_process_11(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_photo_AlignMTB_setCut_10(JNIEnv*, jclass, jlong, jboolean);
extern void Java_org_opencv_photo_AlignMTB_setExcludeRange_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_photo_AlignMTB_setMaxBits_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_photo_AlignMTB_shiftMat_10(JNIEnv*, jclass, jlong, jlong, jlong, jdouble, jdouble);
extern void Java_org_opencv_photo_CalibrateCRF_delete(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_photo_CalibrateCRF_process_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_photo_CalibrateDebevec_delete(JNIEnv*, jclass, jlong);
extern jfloat Java_org_opencv_photo_CalibrateDebevec_getLambda_10(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_photo_CalibrateDebevec_getRandom_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_photo_CalibrateDebevec_getSamples_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_photo_CalibrateDebevec_setLambda_10(JNIEnv*, jclass, jlong, jfloat);
extern void Java_org_opencv_photo_CalibrateDebevec_setRandom_10(JNIEnv*, jclass, jlong, jboolean);
extern void Java_org_opencv_photo_CalibrateDebevec_setSamples_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_photo_CalibrateRobertson_delete(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_photo_CalibrateRobertson_getMaxIter_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_photo_CalibrateRobertson_getRadiance_10(JNIEnv*, jclass, jlong);
extern jfloat Java_org_opencv_photo_CalibrateRobertson_getThreshold_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_photo_CalibrateRobertson_setMaxIter_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_photo_CalibrateRobertson_setThreshold_10(JNIEnv*, jclass, jlong, jfloat);
extern void Java_org_opencv_photo_MergeDebevec_delete(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_photo_MergeDebevec_process_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_photo_MergeDebevec_process_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_photo_MergeExposures_delete(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_photo_MergeExposures_process_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_photo_MergeMertens_delete(JNIEnv*, jclass, jlong);
extern jfloat Java_org_opencv_photo_MergeMertens_getContrastWeight_10(JNIEnv*, jclass, jlong);
extern jfloat Java_org_opencv_photo_MergeMertens_getExposureWeight_10(JNIEnv*, jclass, jlong);
extern jfloat Java_org_opencv_photo_MergeMertens_getSaturationWeight_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_photo_MergeMertens_process_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_photo_MergeMertens_process_11(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_photo_MergeMertens_setContrastWeight_10(JNIEnv*, jclass, jlong, jfloat);
extern void Java_org_opencv_photo_MergeMertens_setExposureWeight_10(JNIEnv*, jclass, jlong, jfloat);
extern void Java_org_opencv_photo_MergeMertens_setSaturationWeight_10(JNIEnv*, jclass, jlong, jfloat);
extern void Java_org_opencv_photo_MergeRobertson_delete(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_photo_MergeRobertson_process_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_photo_MergeRobertson_process_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_photo_Photo_colorChange_10(JNIEnv*, jclass, jlong, jlong, jlong, jfloat, jfloat, jfloat);
extern void Java_org_opencv_photo_Photo_colorChange_11(JNIEnv*, jclass, jlong, jlong, jlong);
extern jlong Java_org_opencv_photo_Photo_createAlignMTB_10(JNIEnv*, jclass, jint, jint, jboolean);
extern jlong Java_org_opencv_photo_Photo_createAlignMTB_11(JNIEnv*, jclass);
extern jlong Java_org_opencv_photo_Photo_createCalibrateDebevec_10(JNIEnv*, jclass, jint, jfloat, jboolean);
extern jlong Java_org_opencv_photo_Photo_createCalibrateDebevec_11(JNIEnv*, jclass);
extern jlong Java_org_opencv_photo_Photo_createCalibrateRobertson_10(JNIEnv*, jclass, jint, jfloat);
extern jlong Java_org_opencv_photo_Photo_createCalibrateRobertson_11(JNIEnv*, jclass);
extern jlong Java_org_opencv_photo_Photo_createMergeDebevec_10(JNIEnv*, jclass);
extern jlong Java_org_opencv_photo_Photo_createMergeMertens_10(JNIEnv*, jclass, jfloat, jfloat, jfloat);
extern jlong Java_org_opencv_photo_Photo_createMergeMertens_11(JNIEnv*, jclass);
extern jlong Java_org_opencv_photo_Photo_createMergeRobertson_10(JNIEnv*, jclass);
extern jlong Java_org_opencv_photo_Photo_createTonemapDrago_10(JNIEnv*, jclass, jfloat, jfloat, jfloat);
extern jlong Java_org_opencv_photo_Photo_createTonemapDrago_11(JNIEnv*, jclass);
extern jlong Java_org_opencv_photo_Photo_createTonemapDurand_10(JNIEnv*, jclass, jfloat, jfloat, jfloat, jfloat, jfloat);
extern jlong Java_org_opencv_photo_Photo_createTonemapDurand_11(JNIEnv*, jclass);
extern jlong Java_org_opencv_photo_Photo_createTonemapMantiuk_10(JNIEnv*, jclass, jfloat, jfloat, jfloat);
extern jlong Java_org_opencv_photo_Photo_createTonemapMantiuk_11(JNIEnv*, jclass);
extern jlong Java_org_opencv_photo_Photo_createTonemapReinhard_10(JNIEnv*, jclass, jfloat, jfloat, jfloat, jfloat);
extern jlong Java_org_opencv_photo_Photo_createTonemapReinhard_11(JNIEnv*, jclass);
extern jlong Java_org_opencv_photo_Photo_createTonemap_10(JNIEnv*, jclass, jfloat);
extern jlong Java_org_opencv_photo_Photo_createTonemap_11(JNIEnv*, jclass);
extern void Java_org_opencv_photo_Photo_decolor_10(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_photo_Photo_denoise_1TVL1_10(JNIEnv*, jclass, jlong, jlong, jdouble, jint);
extern void Java_org_opencv_photo_Photo_denoise_1TVL1_11(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_photo_Photo_detailEnhance_10(JNIEnv*, jclass, jlong, jlong, jfloat, jfloat);
extern void Java_org_opencv_photo_Photo_detailEnhance_11(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_photo_Photo_edgePreservingFilter_10(JNIEnv*, jclass, jlong, jlong, jint, jfloat, jfloat);
extern void Java_org_opencv_photo_Photo_edgePreservingFilter_11(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_photo_Photo_fastNlMeansDenoisingColoredMulti_10(JNIEnv*, jclass, jlong, jlong, jint, jint, jfloat, jfloat, jint, jint);
extern void Java_org_opencv_photo_Photo_fastNlMeansDenoisingColoredMulti_11(JNIEnv*, jclass, jlong, jlong, jint, jint);
extern void Java_org_opencv_photo_Photo_fastNlMeansDenoisingColored_10(JNIEnv*, jclass, jlong, jlong, jfloat, jfloat, jint, jint);
extern void Java_org_opencv_photo_Photo_fastNlMeansDenoisingColored_11(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_photo_Photo_fastNlMeansDenoisingMulti_10(JNIEnv*, jclass, jlong, jlong, jint, jint, jfloat, jint, jint);
extern void Java_org_opencv_photo_Photo_fastNlMeansDenoisingMulti_11(JNIEnv*, jclass, jlong, jlong, jint, jint);
extern void Java_org_opencv_photo_Photo_fastNlMeansDenoisingMulti_12(JNIEnv*, jclass, jlong, jlong, jint, jint, jlong, jint, jint, jint);
extern void Java_org_opencv_photo_Photo_fastNlMeansDenoisingMulti_13(JNIEnv*, jclass, jlong, jlong, jint, jint, jlong);
extern void Java_org_opencv_photo_Photo_fastNlMeansDenoising_10(JNIEnv*, jclass, jlong, jlong, jfloat, jint, jint);
extern void Java_org_opencv_photo_Photo_fastNlMeansDenoising_11(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_photo_Photo_fastNlMeansDenoising_12(JNIEnv*, jclass, jlong, jlong, jlong, jint, jint, jint);
extern void Java_org_opencv_photo_Photo_fastNlMeansDenoising_13(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_photo_Photo_illuminationChange_10(JNIEnv*, jclass, jlong, jlong, jlong, jfloat, jfloat);
extern void Java_org_opencv_photo_Photo_illuminationChange_11(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_photo_Photo_inpaint_10(JNIEnv*, jclass, jlong, jlong, jlong, jdouble, jint);
extern void Java_org_opencv_photo_Photo_pencilSketch_10(JNIEnv*, jclass, jlong, jlong, jlong, jfloat, jfloat, jfloat);
extern void Java_org_opencv_photo_Photo_pencilSketch_11(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_photo_Photo_seamlessClone_10(JNIEnv*, jclass, jlong, jlong, jlong, jdouble, jdouble, jlong, jint);
extern void Java_org_opencv_photo_Photo_stylization_10(JNIEnv*, jclass, jlong, jlong, jfloat, jfloat);
extern void Java_org_opencv_photo_Photo_stylization_11(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_photo_Photo_textureFlattening_10(JNIEnv*, jclass, jlong, jlong, jlong, jfloat, jfloat, jint);
extern void Java_org_opencv_photo_Photo_textureFlattening_11(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_photo_Tonemap_delete(JNIEnv*, jclass, jlong);
extern jfloat Java_org_opencv_photo_Tonemap_getGamma_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_photo_Tonemap_process_10(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_photo_Tonemap_setGamma_10(JNIEnv*, jclass, jlong, jfloat);
extern void Java_org_opencv_photo_TonemapDrago_delete(JNIEnv*, jclass, jlong);
extern jfloat Java_org_opencv_photo_TonemapDrago_getBias_10(JNIEnv*, jclass, jlong);
extern jfloat Java_org_opencv_photo_TonemapDrago_getSaturation_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_photo_TonemapDrago_setBias_10(JNIEnv*, jclass, jlong, jfloat);
extern void Java_org_opencv_photo_TonemapDrago_setSaturation_10(JNIEnv*, jclass, jlong, jfloat);
extern void Java_org_opencv_photo_TonemapDurand_delete(JNIEnv*, jclass, jlong);
extern jfloat Java_org_opencv_photo_TonemapDurand_getContrast_10(JNIEnv*, jclass, jlong);
extern jfloat Java_org_opencv_photo_TonemapDurand_getSaturation_10(JNIEnv*, jclass, jlong);
extern jfloat Java_org_opencv_photo_TonemapDurand_getSigmaColor_10(JNIEnv*, jclass, jlong);
extern jfloat Java_org_opencv_photo_TonemapDurand_getSigmaSpace_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_photo_TonemapDurand_setContrast_10(JNIEnv*, jclass, jlong, jfloat);
extern void Java_org_opencv_photo_TonemapDurand_setSaturation_10(JNIEnv*, jclass, jlong, jfloat);
extern void Java_org_opencv_photo_TonemapDurand_setSigmaColor_10(JNIEnv*, jclass, jlong, jfloat);
extern void Java_org_opencv_photo_TonemapDurand_setSigmaSpace_10(JNIEnv*, jclass, jlong, jfloat);
extern void Java_org_opencv_photo_TonemapMantiuk_delete(JNIEnv*, jclass, jlong);
extern jfloat Java_org_opencv_photo_TonemapMantiuk_getSaturation_10(JNIEnv*, jclass, jlong);
extern jfloat Java_org_opencv_photo_TonemapMantiuk_getScale_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_photo_TonemapMantiuk_setSaturation_10(JNIEnv*, jclass, jlong, jfloat);
extern void Java_org_opencv_photo_TonemapMantiuk_setScale_10(JNIEnv*, jclass, jlong, jfloat);
extern void Java_org_opencv_photo_TonemapReinhard_delete(JNIEnv*, jclass, jlong);
extern jfloat Java_org_opencv_photo_TonemapReinhard_getColorAdaptation_10(JNIEnv*, jclass, jlong);
extern jfloat Java_org_opencv_photo_TonemapReinhard_getIntensity_10(JNIEnv*, jclass, jlong);
extern jfloat Java_org_opencv_photo_TonemapReinhard_getLightAdaptation_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_photo_TonemapReinhard_setColorAdaptation_10(JNIEnv*, jclass, jlong, jfloat);
extern void Java_org_opencv_photo_TonemapReinhard_setIntensity_10(JNIEnv*, jclass, jlong, jfloat);
extern void Java_org_opencv_photo_TonemapReinhard_setLightAdaptation_10(JNIEnv*, jclass, jlong, jfloat);

*/
import "C"

func AlignExposuresNative_delete(nativeObj int64) {
	C.Java_org_opencv_photo_AlignExposures_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}
func AlignExposuresNative_process_0(nativeObj int64,src_mat_nativeObj int64,dst_mat_nativeObj int64,times_nativeObj int64,response_nativeObj int64) {
	C.Java_org_opencv_photo_AlignExposures_process_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(src_mat_nativeObj),(C.jlong)(dst_mat_nativeObj),(C.jlong)(times_nativeObj),(C.jlong)(response_nativeObj))
}
func AlignMTBNative_calculateShift_0(nativeObj int64,img0_nativeObj int64,img1_nativeObj int64) []float64 {
	return togofloat64Array(C.Java_org_opencv_photo_AlignMTB_calculateShift_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(img0_nativeObj),(C.jlong)(img1_nativeObj)))
}

func AlignMTBNative_computeBitmaps_0(nativeObj int64,img_nativeObj int64,tb_nativeObj int64,eb_nativeObj int64) {
	C.Java_org_opencv_photo_AlignMTB_computeBitmaps_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(img_nativeObj),(C.jlong)(tb_nativeObj),(C.jlong)(eb_nativeObj))
}

func AlignMTBNative_delete(nativeObj int64) {
	C.Java_org_opencv_photo_AlignMTB_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func AlignMTBNative_getCut_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_photo_AlignMTB_getCut_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func AlignMTBNative_getExcludeRange_0(nativeObj int64) int {
	return int(C.Java_org_opencv_photo_AlignMTB_getExcludeRange_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func AlignMTBNative_getMaxBits_0(nativeObj int64) int {
	return int(C.Java_org_opencv_photo_AlignMTB_getMaxBits_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func AlignMTBNative_process_0(nativeObj int64,src_mat_nativeObj int64,dst_mat_nativeObj int64,times_nativeObj int64,response_nativeObj int64) {
	C.Java_org_opencv_photo_AlignMTB_process_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(src_mat_nativeObj),(C.jlong)(dst_mat_nativeObj),(C.jlong)(times_nativeObj),(C.jlong)(response_nativeObj))
}

func AlignMTBNative_process_1(nativeObj int64,src_mat_nativeObj int64,dst_mat_nativeObj int64) {
	C.Java_org_opencv_photo_AlignMTB_process_11(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(src_mat_nativeObj),(C.jlong)(dst_mat_nativeObj))
}

func AlignMTBNative_setCut_0(nativeObj int64,value bool) {
	C.Java_org_opencv_photo_AlignMTB_setCut_10(clzEnv,clzObj,(C.jlong)(nativeObj),tojboolean(value))
}

func AlignMTBNative_setExcludeRange_0(nativeObj int64,exclude_range int) {
	C.Java_org_opencv_photo_AlignMTB_setExcludeRange_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(exclude_range))
}

func AlignMTBNative_setMaxBits_0(nativeObj int64,max_bits int) {
	C.Java_org_opencv_photo_AlignMTB_setMaxBits_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(max_bits))
}

func AlignMTBNative_shiftMat_0(nativeObj int64,src_nativeObj int64,dst_nativeObj int64,shift_x float64,shift_y float64) {
	C.Java_org_opencv_photo_AlignMTB_shiftMat_10(clzEnv,clzObj,(C.jlong)(
		nativeObj),(C.jlong)(src_nativeObj),(C.jlong)(dst_nativeObj),(C.jdouble)(shift_x),(C.jdouble)(shift_y))
}

func CalibrateCRFNative_delete(nativeObj int64) {
	C.Java_org_opencv_photo_CalibrateCRF_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func CalibrateCRFNative_process_0(nativeObj int64,src_mat_nativeObj int64,dst_nativeObj int64,times_nativeObj int64) {
	C.Java_org_opencv_photo_CalibrateCRF_process_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(src_mat_nativeObj),(C.jlong)(dst_nativeObj),(C.jlong)(times_nativeObj))
}

func CalibrateDebevecNative_delete(nativeObj int64) {
	C.Java_org_opencv_photo_CalibrateDebevec_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func CalibrateDebevecNative_getLambda_0(nativeObj int64) float32 {
	return float32(C.Java_org_opencv_photo_CalibrateDebevec_getLambda_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func CalibrateDebevecNative_getRandom_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_photo_CalibrateDebevec_getRandom_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func CalibrateDebevecNative_getSamples_0(nativeObj int64) int {
	return int(C.Java_org_opencv_photo_CalibrateDebevec_getSamples_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func CalibrateDebevecNative_setLambda_0(nativeObj int64,lambda float32) {
	C.Java_org_opencv_photo_CalibrateDebevec_setLambda_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jfloat)(lambda))
}

func CalibrateDebevecNative_setRandom_0(nativeObj int64,random bool) {
	C.Java_org_opencv_photo_CalibrateDebevec_setRandom_10(clzEnv,clzObj,(C.jlong)(nativeObj),tojboolean(random))
}

func CalibrateDebevecNative_setSamples_0(nativeObj int64,samples int) {
	C.Java_org_opencv_photo_CalibrateDebevec_setSamples_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(samples))
}

func CalibrateRobertsonNative_delete(nativeObj int64) {
	C.Java_org_opencv_photo_CalibrateRobertson_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func CalibrateRobertsonNative_getMaxIter_0(nativeObj int64) int {
	return int(C.Java_org_opencv_photo_CalibrateRobertson_getMaxIter_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func CalibrateRobertsonNative_getRadiance_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_photo_CalibrateRobertson_getRadiance_10(
			clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func CalibrateRobertsonNative_getThreshold_0(nativeObj int64) float32 {
	return float32(C.Java_org_opencv_photo_CalibrateRobertson_getThreshold_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func CalibrateRobertsonNative_setMaxIter_0(nativeObj int64,max_iter int) {
	C.Java_org_opencv_photo_CalibrateRobertson_setMaxIter_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(max_iter))
}

func CalibrateRobertsonNative_setThreshold_0(nativeObj int64,threshold float32) {
	C.Java_org_opencv_photo_CalibrateRobertson_setThreshold_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jfloat)(threshold))
}

func MergeDebevecNative_delete(nativeObj int64) {
	C.Java_org_opencv_photo_MergeDebevec_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func MergeDebevecNative_process_0(nativeObj int64,src_mat_nativeObj int64,dst_nativeObj int64,times_nativeObj int64,response_nativeObj int64) {
	C.Java_org_opencv_photo_MergeDebevec_process_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(src_mat_nativeObj),(C.jlong)(dst_nativeObj),(C.jlong)(times_nativeObj),(C.jlong)(response_nativeObj))
}

func MergeDebevecNative_process_1(nativeObj int64,src_mat_nativeObj int64,dst_nativeObj int64,times_nativeObj int64) {
	C.Java_org_opencv_photo_MergeDebevec_process_11(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(src_mat_nativeObj),(C.jlong)(dst_nativeObj),(C.jlong)(times_nativeObj))
}

func MergeExposuresNative_delete(nativeObj int64) {
	C.Java_org_opencv_photo_MergeExposures_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func MergeExposuresNative_process_0(nativeObj int64,src_mat_nativeObj int64,dst_nativeObj int64,times_nativeObj int64,response_nativeObj int64) {
	C.Java_org_opencv_photo_MergeExposures_process_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(src_mat_nativeObj),(C.jlong)(dst_nativeObj),(C.jlong)(times_nativeObj),(C.jlong)(response_nativeObj))
}

func MergeMertensNative_delete(nativeObj int64) {
	C.Java_org_opencv_photo_MergeMertens_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func MergeMertensNative_getContrastWeight_0(nativeObj int64) float32 {
	return float32(C.Java_org_opencv_photo_MergeMertens_getContrastWeight_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func MergeMertensNative_getExposureWeight_0(nativeObj int64) float32 {
	return float32(C.Java_org_opencv_photo_MergeMertens_getExposureWeight_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func MergeMertensNative_getSaturationWeight_0(nativeObj int64) float32 {
	return float32(C.Java_org_opencv_photo_MergeMertens_getSaturationWeight_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func MergeMertensNative_process_0(nativeObj int64,src_mat_nativeObj int64,dst_nativeObj int64,times_nativeObj int64,response_nativeObj int64) {
	C.Java_org_opencv_photo_MergeMertens_process_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(src_mat_nativeObj),(C.jlong)(dst_nativeObj),(C.jlong)(times_nativeObj),(C.jlong)(response_nativeObj))
}

func MergeMertensNative_process_1(nativeObj int64,src_mat_nativeObj int64,dst_nativeObj int64) {
	C.Java_org_opencv_photo_MergeMertens_process_11(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(src_mat_nativeObj),(C.jlong)(dst_nativeObj))
}

func MergeMertensNative_setContrastWeight_0(nativeObj int64,contrast_weiht float32) {
	C.Java_org_opencv_photo_MergeMertens_setContrastWeight_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jfloat)(contrast_weiht))
}

func MergeMertensNative_setExposureWeight_0(nativeObj int64,exposure_weight float32) {
	C.Java_org_opencv_photo_MergeMertens_setExposureWeight_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jfloat)(exposure_weight))
}

func MergeMertensNative_setSaturationWeight_0(nativeObj int64,saturation_weight float32) {
	C.Java_org_opencv_photo_MergeMertens_setSaturationWeight_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jfloat)(saturation_weight))
}

func MergeRobertsonNative_delete(nativeObj int64) {
	C.Java_org_opencv_photo_MergeRobertson_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func MergeRobertsonNative_process_0(nativeObj int64,src_mat_nativeObj int64,dst_nativeObj int64,times_nativeObj int64,response_nativeObj int64) {
	C.Java_org_opencv_photo_MergeRobertson_process_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(src_mat_nativeObj),(C.jlong)(dst_nativeObj),(C.jlong)(times_nativeObj),(C.jlong)(response_nativeObj))
}

func MergeRobertsonNative_process_1(nativeObj int64,src_mat_nativeObj int64,dst_nativeObj int64,times_nativeObj int64) {
	C.Java_org_opencv_photo_MergeRobertson_process_11(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(src_mat_nativeObj),(C.jlong)(dst_nativeObj),(C.jlong)(times_nativeObj))
}

func PhotoNative_colorChange_0(src_nativeObj int64,mask_nativeObj int64,dst_nativeObj int64,red_mul float32,green_mul float32,blue_mul float32) {
	C.Java_org_opencv_photo_Photo_colorChange_10(clzEnv,clzObj,(C.jlong)(
		src_nativeObj),(C.jlong)(mask_nativeObj),(C.jlong)(dst_nativeObj),(C.jfloat)(red_mul),(C.jfloat)(green_mul),(C.jfloat)(blue_mul))
}

func PhotoNative_colorChange_1(src_nativeObj int64,mask_nativeObj int64,dst_nativeObj int64) {
	C.Java_org_opencv_photo_Photo_colorChange_11(clzEnv,clzObj,(C.jlong)(
		src_nativeObj),(C.jlong)(mask_nativeObj),(C.jlong)(dst_nativeObj))
}

func PhotoNative_createAlignMTB_0(max_bits int,exclude_range int,cut bool) int64 {
	return int64(C.Java_org_opencv_photo_Photo_createAlignMTB_10(clzEnv,clzObj,(C.jint)(max_bits),(C.jint)(exclude_range),tojboolean(cut)))
}

func PhotoNative_createAlignMTB_1() int64 {
	return int64(C.Java_org_opencv_photo_Photo_createAlignMTB_11(clzEnv,clzObj))
}

func PhotoNative_createCalibrateDebevec_0(samples int,lambda float32,random bool) int64 {
	return int64(C.Java_org_opencv_photo_Photo_createCalibrateDebevec_10(clzEnv,clzObj,(C.jint)(samples),(C.jfloat)(lambda),tojboolean(random)))
}

func PhotoNative_createCalibrateDebevec_1() int64 {
	return int64(C.Java_org_opencv_photo_Photo_createCalibrateDebevec_11(clzEnv,clzObj))
}

func PhotoNative_createCalibrateRobertson_0(max_iter int,threshold float32) int64 {
	return int64(C.Java_org_opencv_photo_Photo_createCalibrateRobertson_10(
			clzEnv,clzObj,(C.jint)(max_iter),(C.jfloat)(threshold)))
}

func PhotoNative_createCalibrateRobertson_1() int64 {
	return int64(C.Java_org_opencv_photo_Photo_createCalibrateRobertson_11(
			clzEnv,clzObj))
}

func PhotoNative_createMergeDebevec_0() int64 {
	return int64(C.Java_org_opencv_photo_Photo_createMergeDebevec_10(clzEnv,clzObj))
}

func PhotoNative_createMergeMertens_0(contrast_weight float32,saturation_weight float32,exposure_weight float32) int64 {
	return int64(C.Java_org_opencv_photo_Photo_createMergeMertens_10(clzEnv,clzObj,(C.jfloat)(contrast_weight),(C.jfloat)(saturation_weight),(C.jfloat)(exposure_weight)))
}

func PhotoNative_createMergeMertens_1() int64 {
	return int64(C.Java_org_opencv_photo_Photo_createMergeMertens_11(clzEnv,clzObj))
}

func PhotoNative_createMergeRobertson_0() int64 {
	return int64(C.Java_org_opencv_photo_Photo_createMergeRobertson_10(clzEnv,clzObj))
}

func PhotoNative_createTonemapDrago_0(gamma float32,saturation float32,bias float32) int64 {
	return int64(C.Java_org_opencv_photo_Photo_createTonemapDrago_10(clzEnv,clzObj,(C.jfloat)(gamma),(C.jfloat)(saturation),(C.jfloat)(bias)))
}

func PhotoNative_createTonemapDrago_1() int64 {
	return int64(C.Java_org_opencv_photo_Photo_createTonemapDrago_11(clzEnv,clzObj))
}

func PhotoNative_createTonemapDurand_0(gamma float32,contrast float32,saturation float32,sigma_space float32,sigma_color float32) int64 {
	return int64(C.Java_org_opencv_photo_Photo_createTonemapDurand_10(clzEnv,clzObj,(C.jfloat)(gamma),(C.jfloat)(contrast),(C.jfloat)(saturation),(C.jfloat)(sigma_space),(C.jfloat)(sigma_color)))
}

func PhotoNative_createTonemapDurand_1() int64 {
	return int64(C.Java_org_opencv_photo_Photo_createTonemapDurand_11(clzEnv,clzObj))
}

func PhotoNative_createTonemapMantiuk_0(gamma float32,scale float32,saturation float32) int64 {
	return int64(C.Java_org_opencv_photo_Photo_createTonemapMantiuk_10(clzEnv,clzObj,(C.jfloat)(gamma),(C.jfloat)(scale),(C.jfloat)(saturation)))
}

func PhotoNative_createTonemapMantiuk_1() int64 {
	return int64(C.Java_org_opencv_photo_Photo_createTonemapMantiuk_11(clzEnv,clzObj))
}

func PhotoNative_createTonemapReinhard_0(gamma float32,intensity float32,light_adapt float32,color_adapt float32) int64 {
	return int64(C.Java_org_opencv_photo_Photo_createTonemapReinhard_10(clzEnv,clzObj,(C.jfloat)(gamma),(C.jfloat)(intensity),(C.jfloat)(light_adapt),(C.jfloat)(color_adapt)))
}

func PhotoNative_createTonemapReinhard_1() int64 {
	return int64(C.Java_org_opencv_photo_Photo_createTonemapReinhard_11(clzEnv,clzObj))
}

func PhotoNative_createTonemap_0(gamma float32) int64 {
	return int64(C.Java_org_opencv_photo_Photo_createTonemap_10(clzEnv,clzObj,(C.jfloat)(gamma)))
}

func PhotoNative_createTonemap_1() int64 {
	return int64(C.Java_org_opencv_photo_Photo_createTonemap_11(clzEnv,clzObj))
}

func PhotoNative_decolor_0(src_nativeObj int64,grayscale_nativeObj int64,color_boost_nativeObj int64) {
	C.Java_org_opencv_photo_Photo_decolor_10(clzEnv,clzObj,(C.jlong)(src_nativeObj),(C.jlong)(grayscale_nativeObj),(C.jlong)(color_boost_nativeObj))
}

func PhotoNative_denoise_TVL1_0(observations_mat_nativeObj int64,result_nativeObj int64,lambda float64,niters int) {
	C.Java_org_opencv_photo_Photo_denoise_1TVL1_10(clzEnv,clzObj,(C.jlong)(observations_mat_nativeObj),(C.jlong)(result_nativeObj),(C.jdouble)(lambda),(C.jint)(niters))
}

func PhotoNative_denoise_TVL1_1(observations_mat_nativeObj int64,result_nativeObj int64) {
	C.Java_org_opencv_photo_Photo_denoise_1TVL1_11(clzEnv,clzObj,(C.jlong)(observations_mat_nativeObj),(C.jlong)(result_nativeObj))
}

func PhotoNative_detailEnhance_0(src_nativeObj int64,dst_nativeObj int64,sigma_s float32,sigma_r float32) {
	C.Java_org_opencv_photo_Photo_detailEnhance_10(clzEnv,clzObj,(C.jlong)(src_nativeObj),(C.jlong)(dst_nativeObj),(C.jfloat)(sigma_s),(C.jfloat)(sigma_r))
}

func PhotoNative_detailEnhance_1(src_nativeObj int64,dst_nativeObj int64) {
	C.Java_org_opencv_photo_Photo_detailEnhance_11(clzEnv,clzObj,(C.jlong)(src_nativeObj),(C.jlong)(dst_nativeObj))
}

func PhotoNative_edgePreservingFilter_0(src_nativeObj int64,dst_nativeObj int64,flags int,sigma_s float32,sigma_r float32) {
	C.Java_org_opencv_photo_Photo_edgePreservingFilter_10(clzEnv,clzObj,(C.jlong)(src_nativeObj),(C.jlong)(dst_nativeObj),(C.jint)(flags),(C.jfloat)(sigma_s),(C.jfloat)(sigma_r))
}

func PhotoNative_edgePreservingFilter_1(src_nativeObj int64,dst_nativeObj int64) {
	C.Java_org_opencv_photo_Photo_edgePreservingFilter_11(clzEnv,clzObj,(C.jlong)(src_nativeObj),(C.jlong)(dst_nativeObj))
}

func PhotoNative_fastNlMeansDenoisingColoredMulti_0(srcImgs_mat_nativeObj int64,dst_nativeObj int64,imgToDenoiseIndex int,temporalWindowSize int,h float32,hColor float32,templateWindowSize int,searchWindowSize int) {
	C.Java_org_opencv_photo_Photo_fastNlMeansDenoisingColoredMulti_10(clzEnv,clzObj,(C.jlong)(srcImgs_mat_nativeObj),(C.jlong)(dst_nativeObj),(C.jint)(imgToDenoiseIndex),(C.jint)(temporalWindowSize),(C.jfloat)(h),(C.jfloat)(hColor),(C.jint)(templateWindowSize),(C.jint)(searchWindowSize))
}

func PhotoNative_fastNlMeansDenoisingColoredMulti_1(srcImgs_mat_nativeObj int64,dst_nativeObj int64,imgToDenoiseIndex int,temporalWindowSize int) {
	C.Java_org_opencv_photo_Photo_fastNlMeansDenoisingColoredMulti_11(clzEnv,clzObj,(C.jlong)(srcImgs_mat_nativeObj),(C.jlong)(dst_nativeObj),(C.jint)(imgToDenoiseIndex),(C.jint)(temporalWindowSize))
}

func PhotoNative_fastNlMeansDenoisingColored_0(src_nativeObj int64,dst_nativeObj int64,h float32,hColor float32,templateWindowSize int,searchWindowSize int) {
	C.Java_org_opencv_photo_Photo_fastNlMeansDenoisingColored_10(clzEnv,clzObj,(C.jlong)(src_nativeObj),(C.jlong)(dst_nativeObj),(C.jfloat)(h),(C.jfloat)(hColor),(C.jint)(templateWindowSize),(C.jint)(searchWindowSize))
}

func PhotoNative_fastNlMeansDenoisingColored_1(src_nativeObj int64,dst_nativeObj int64) {
	C.Java_org_opencv_photo_Photo_fastNlMeansDenoisingColored_11(clzEnv,clzObj,(C.jlong)(src_nativeObj),(C.jlong)(dst_nativeObj))
}

func PhotoNative_fastNlMeansDenoisingMulti_0(srcImgs_mat_nativeObj int64,dst_nativeObj int64,imgToDenoiseIndex int,temporalWindowSize int,h float32,templateWindowSize int,searchWindowSize int) {
	C.Java_org_opencv_photo_Photo_fastNlMeansDenoisingMulti_10(clzEnv,clzObj,(C.jlong)(srcImgs_mat_nativeObj),(C.jlong)(dst_nativeObj),(C.jint)(imgToDenoiseIndex),(C.jint)(temporalWindowSize),(C.jfloat)(h),(C.jint)(templateWindowSize),(C.jint)(searchWindowSize))
}

func PhotoNative_fastNlMeansDenoisingMulti_1(srcImgs_mat_nativeObj int64,dst_nativeObj int64,imgToDenoiseIndex int,temporalWindowSize int) {
	C.Java_org_opencv_photo_Photo_fastNlMeansDenoisingMulti_11(clzEnv,clzObj,(C.jlong)(srcImgs_mat_nativeObj),(C.jlong)(dst_nativeObj),(C.jint)(imgToDenoiseIndex),(C.jint)(temporalWindowSize))
}

func PhotoNative_fastNlMeansDenoisingMulti_2(srcImgs_mat_nativeObj int64,dst_nativeObj int64,imgToDenoiseIndex int,temporalWindowSize int,h_mat_nativeObj int64,templateWindowSize int,searchWindowSize int,normType int) {
	C.Java_org_opencv_photo_Photo_fastNlMeansDenoisingMulti_12(clzEnv,clzObj,(C.jlong)(srcImgs_mat_nativeObj),(C.jlong)(dst_nativeObj),(C.jint)(imgToDenoiseIndex),(C.jint)(temporalWindowSize),(C.jlong)(h_mat_nativeObj),(C.jint)(templateWindowSize),(C.jint)(searchWindowSize),(C.jint)(normType))
}

func PhotoNative_fastNlMeansDenoisingMulti_3(srcImgs_mat_nativeObj int64,dst_nativeObj int64,imgToDenoiseIndex int,temporalWindowSize int,h_mat_nativeObj int64) {
	C.Java_org_opencv_photo_Photo_fastNlMeansDenoisingMulti_13(clzEnv,clzObj,(C.jlong)(srcImgs_mat_nativeObj),(C.jlong)(dst_nativeObj),(C.jint)(imgToDenoiseIndex),(C.jint)(temporalWindowSize),(C.jlong)(h_mat_nativeObj))
}

func PhotoNative_fastNlMeansDenoising_0(src_nativeObj int64,dst_nativeObj int64,h float32,templateWindowSize int,searchWindowSize int) {
	C.Java_org_opencv_photo_Photo_fastNlMeansDenoising_10(clzEnv,clzObj,(C.jlong)(src_nativeObj),(C.jlong)(dst_nativeObj),(C.jfloat)(h),(C.jint)(templateWindowSize),(C.jint)(searchWindowSize))
}

func PhotoNative_fastNlMeansDenoising_1(src_nativeObj int64,dst_nativeObj int64) {
	C.Java_org_opencv_photo_Photo_fastNlMeansDenoising_11(clzEnv,clzObj,(C.jlong)(src_nativeObj),(C.jlong)(dst_nativeObj))
}

func PhotoNative_fastNlMeansDenoising_2(src_nativeObj int64,dst_nativeObj int64,h_mat_nativeObj int64,templateWindowSize int,searchWindowSize int,normType int) {
	C.Java_org_opencv_photo_Photo_fastNlMeansDenoising_12(clzEnv,clzObj,(C.jlong)(src_nativeObj),(C.jlong)(dst_nativeObj),(C.jlong)(h_mat_nativeObj),(C.jint)(templateWindowSize),(C.jint)(searchWindowSize),(C.jint)(normType))
}

func PhotoNative_fastNlMeansDenoising_3(src_nativeObj int64,dst_nativeObj int64,h_mat_nativeObj int64) {
	C.Java_org_opencv_photo_Photo_fastNlMeansDenoising_13(clzEnv,clzObj,(C.jlong)(src_nativeObj),(C.jlong)(dst_nativeObj),(C.jlong)(h_mat_nativeObj))
}

func PhotoNative_illuminationChange_0(src_nativeObj int64,mask_nativeObj int64,dst_nativeObj int64,alpha float32,beta float32) {
	C.Java_org_opencv_photo_Photo_illuminationChange_10(clzEnv,clzObj,(C.jlong)(src_nativeObj),(C.jlong)(mask_nativeObj),(C.jlong)(dst_nativeObj),(C.jfloat)(alpha),(C.jfloat)(beta))
}

func PhotoNative_illuminationChange_1(src_nativeObj int64,mask_nativeObj int64,dst_nativeObj int64) {
	C.Java_org_opencv_photo_Photo_illuminationChange_11(clzEnv,clzObj,(C.jlong)(src_nativeObj),(C.jlong)(mask_nativeObj),(C.jlong)(dst_nativeObj))
}

func PhotoNative_inpaint_0(src_nativeObj int64,inpaintMask_nativeObj int64,dst_nativeObj int64,inpaintRadius float64,flags int) {
	C.Java_org_opencv_photo_Photo_inpaint_10(clzEnv,clzObj,(C.jlong)(src_nativeObj),(C.jlong)(inpaintMask_nativeObj),(C.jlong)(dst_nativeObj),(C.jdouble)(inpaintRadius),(C.jint)(flags))
}

func PhotoNative_pencilSketch_0(src_nativeObj int64,dst1_nativeObj int64,dst2_nativeObj int64,sigma_s float32,sigma_r float32,shade_factor float32) {
	C.Java_org_opencv_photo_Photo_pencilSketch_10(clzEnv,clzObj,(C.jlong)(src_nativeObj),(C.jlong)(dst1_nativeObj),(C.jlong)(dst2_nativeObj),(C.jfloat)(sigma_s),(C.jfloat)(sigma_r),(C.jfloat)(shade_factor))
}

func PhotoNative_pencilSketch_1(src_nativeObj int64,dst1_nativeObj int64,dst2_nativeObj int64) {
	C.Java_org_opencv_photo_Photo_pencilSketch_11(clzEnv,clzObj,(C.jlong)(src_nativeObj),(C.jlong)(dst1_nativeObj),(C.jlong)(dst2_nativeObj))
}

func PhotoNative_seamlessClone_0(src_nativeObj int64,dst_nativeObj int64,mask_nativeObj int64,p_x float64,p_y float64,blend_nativeObj int64,flags int) {
	C.Java_org_opencv_photo_Photo_seamlessClone_10(clzEnv,clzObj,(C.jlong)(src_nativeObj),(C.jlong)(dst_nativeObj),(C.jlong)(mask_nativeObj),(C.jdouble)(p_x),(C.jdouble)(p_y),(C.jlong)(blend_nativeObj),(C.jint)(flags))
}

func PhotoNative_stylization_0(src_nativeObj int64,dst_nativeObj int64,sigma_s float32,sigma_r float32) {
	C.Java_org_opencv_photo_Photo_stylization_10(clzEnv,clzObj,(C.jlong)(
		src_nativeObj),(C.jlong)(dst_nativeObj),(C.jfloat)(sigma_s),(C.jfloat)(sigma_r))
}

func PhotoNative_stylization_1(src_nativeObj int64,dst_nativeObj int64) {
	C.Java_org_opencv_photo_Photo_stylization_11(clzEnv,clzObj,(C.jlong)(
		src_nativeObj),(C.jlong)(dst_nativeObj))
}

func PhotoNative_textureFlattening_0(src_nativeObj int64,mask_nativeObj int64,dst_nativeObj int64,low_threshold float32,high_threshold float32,kernel_size int) {
	C.Java_org_opencv_photo_Photo_textureFlattening_10(clzEnv,clzObj,(C.jlong)(src_nativeObj),(C.jlong)(mask_nativeObj),(C.jlong)(dst_nativeObj),(C.jfloat)(low_threshold),(C.jfloat)(high_threshold),(C.jint)(kernel_size))
}

func PhotoNative_textureFlattening_1(src_nativeObj int64,mask_nativeObj int64,dst_nativeObj int64) {
	C.Java_org_opencv_photo_Photo_textureFlattening_11(clzEnv,clzObj,(C.jlong)(src_nativeObj),(C.jlong)(mask_nativeObj),(C.jlong)(dst_nativeObj))
}

func TonemapNative_delete(nativeObj int64) {
	C.Java_org_opencv_photo_Tonemap_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func TonemapNative_getGamma_0(nativeObj int64) float32 {
	return float32(C.Java_org_opencv_photo_Tonemap_getGamma_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TonemapNative_process_0(nativeObj int64,src_nativeObj int64,dst_nativeObj int64) {
	C.Java_org_opencv_photo_Tonemap_process_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(src_nativeObj),(C.jlong)(dst_nativeObj))
}

func TonemapNative_setGamma_0(nativeObj int64,gamma float32) {
	C.Java_org_opencv_photo_Tonemap_setGamma_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jfloat)(gamma))
}

func TonemapDragoNative_delete(nativeObj int64) {
	C.Java_org_opencv_photo_TonemapDrago_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func TonemapDragoNative_getBias_0(nativeObj int64) float32 {
	return float32(C.Java_org_opencv_photo_TonemapDrago_getBias_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TonemapDragoNative_getSaturation_0(nativeObj int64) float32 {
	return float32(C.Java_org_opencv_photo_TonemapDrago_getSaturation_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TonemapDragoNative_setBias_0(nativeObj int64,bias float32) {
	C.Java_org_opencv_photo_TonemapDrago_setBias_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jfloat)(bias))
}

func TonemapDragoNative_setSaturation_0(nativeObj int64,saturation float32) {
	C.Java_org_opencv_photo_TonemapDrago_setSaturation_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jfloat)(saturation))
}

func TonemapDurandNative_delete(nativeObj int64) {
	C.Java_org_opencv_photo_TonemapDurand_delete(clzEnv,clzObj,(C.jlong)(
		nativeObj))
}

func TonemapDurandNative_getContrast_0(nativeObj int64) float32 {
	return float32(C.Java_org_opencv_photo_TonemapDurand_getContrast_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TonemapDurandNative_getSaturation_0(nativeObj int64) float32 {
	return float32(C.Java_org_opencv_photo_TonemapDurand_getSaturation_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TonemapDurandNative_getSigmaColor_0(nativeObj int64) float32 {
	return float32(C.Java_org_opencv_photo_TonemapDurand_getSigmaColor_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TonemapDurandNative_getSigmaSpace_0(nativeObj int64) float32 {
	return float32(C.Java_org_opencv_photo_TonemapDurand_getSigmaSpace_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TonemapDurandNative_setContrast_0(nativeObj int64,contrast float32) {
	C.Java_org_opencv_photo_TonemapDurand_setContrast_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jfloat)(contrast))
}

func TonemapDurandNative_setSaturation_0(nativeObj int64,saturation float32) {
	C.Java_org_opencv_photo_TonemapDurand_setSaturation_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jfloat)(saturation))
}

func TonemapDurandNative_setSigmaColor_0(nativeObj int64,sigma_color float32) {
	C.Java_org_opencv_photo_TonemapDurand_setSigmaColor_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jfloat)(sigma_color))
}

func TonemapDurandNative_setSigmaSpace_0(nativeObj int64,sigma_space float32) {
	C.Java_org_opencv_photo_TonemapDurand_setSigmaSpace_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jfloat)(sigma_space))
}

func TonemapMantiukNative_delete(nativeObj int64) {
	C.Java_org_opencv_photo_TonemapMantiuk_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func TonemapMantiukNative_getSaturation_0(nativeObj int64) float32 {
	return float32(C.Java_org_opencv_photo_TonemapMantiuk_getSaturation_10(
				clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TonemapMantiukNative_getScale_0(nativeObj int64) float32 {
	return float32(C.Java_org_opencv_photo_TonemapMantiuk_getScale_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TonemapMantiukNative_setSaturation_0(nativeObj int64,saturation float32) {
	C.Java_org_opencv_photo_TonemapMantiuk_setSaturation_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jfloat)(saturation))
}

func TonemapMantiukNative_setScale_0(nativeObj int64,scale float32) {
	C.Java_org_opencv_photo_TonemapMantiuk_setScale_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jfloat)(scale))
}

func TonemapReinhardNative_delete(nativeObj int64) {
	C.Java_org_opencv_photo_TonemapReinhard_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func TonemapReinhardNative_getColorAdaptation_0(nativeObj int64) float32 {
	return float32(C.Java_org_opencv_photo_TonemapReinhard_getColorAdaptation_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TonemapReinhardNative_getIntensity_0(nativeObj int64) float32 {
	return float32(C.Java_org_opencv_photo_TonemapReinhard_getIntensity_10(
				clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TonemapReinhardNative_getLightAdaptation_0(nativeObj int64) float32 {
	return float32(C.Java_org_opencv_photo_TonemapReinhard_getLightAdaptation_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TonemapReinhardNative_setColorAdaptation_0(nativeObj int64,color_adapt float32) {
	C.Java_org_opencv_photo_TonemapReinhard_setColorAdaptation_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jfloat)(color_adapt))
}

func TonemapReinhardNative_setIntensity_0(nativeObj int64,intensity float32) {
	C.Java_org_opencv_photo_TonemapReinhard_setIntensity_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jfloat)(intensity))
}

func TonemapReinhardNative_setLightAdaptation_0(nativeObj int64,light_adapt float32) {
	C.Java_org_opencv_photo_TonemapReinhard_setLightAdaptation_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jfloat)(light_adapt))
}
