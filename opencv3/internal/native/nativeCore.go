package core

/*
#include "jni.h"

extern jstring Java_org_opencv_core_Mat_nDump(JNIEnv*, jclass, jlong);
extern jdoubleArray Java_org_opencv_core_Mat_nGet(JNIEnv*, jclass, jlong, jint, jint);
extern jint Java_org_opencv_core_Mat_nGetB(JNIEnv*, jclass, jlong, jint, jint, jint, jbyteArray);
extern jint Java_org_opencv_core_Mat_nGetD(JNIEnv*, jclass, jlong, jint, jint, jint, jdoubleArray);
extern jint Java_org_opencv_core_Mat_nGetF(JNIEnv*, jclass, jlong, jint, jint, jint, jfloatArray);
extern jint Java_org_opencv_core_Mat_nGetI(JNIEnv*, jclass, jlong, jint, jint, jint, jintArray);
extern jint Java_org_opencv_core_Mat_nGetS(JNIEnv*, jclass, jlong, jint, jint, jint, jshortArray);
extern jint Java_org_opencv_core_Mat_nPutB(JNIEnv*, jclass, jlong, jint, jint, jint, jbyteArray);
extern jint Java_org_opencv_core_Mat_nPutD(JNIEnv*, jclass, jlong, jint, jint, jint, jdoubleArray);
extern jint Java_org_opencv_core_Mat_nPutF(JNIEnv*, jclass, jlong, jint, jint, jint, jfloatArray);
extern jint Java_org_opencv_core_Mat_nPutI(JNIEnv*, jclass, jlong, jint, jint, jint, jintArray);
extern jint Java_org_opencv_core_Mat_nPutS(JNIEnv*, jclass, jlong, jint, jint, jint, jshortArray);
extern jlong Java_org_opencv_core_Mat_n_1Mat__(JNIEnv*, jclass);
extern jlong Java_org_opencv_core_Mat_n_1Mat__III(JNIEnv*, jclass, jint, jint, jint);
extern jlong Java_org_opencv_core_Mat_n_1Mat__DDI(JNIEnv*, jclass, jdouble, jdouble, jint);
extern jlong Java_org_opencv_core_Mat_n_1Mat__IIIDDDD(JNIEnv*, jclass, jint, jint, jint, jdouble, jdouble, jdouble, jdouble);
extern jlong Java_org_opencv_core_Mat_n_1Mat__DDIDDDD(JNIEnv*, jclass, jdouble, jdouble, jint, jdouble, jdouble, jdouble, jdouble);
extern jlong Java_org_opencv_core_Mat_n_1Mat__JIIII(JNIEnv*, jclass, jlong, jint, jint, jint, jint);
extern jlong Java_org_opencv_core_Mat_n_1Mat__JII(JNIEnv*, jclass, jlong, jint, jint);
extern jlong Java_org_opencv_core_Mat_n_1adjustROI(JNIEnv*, jclass, jlong, jint, jint, jint, jint);
extern void Java_org_opencv_core_Mat_n_1assignTo__JJI(JNIEnv*, jclass, jlong, jlong, jint);
extern void Java_org_opencv_core_Mat_n_1assignTo__JJ(JNIEnv*, jclass, jlong, jlong);
extern jint Java_org_opencv_core_Mat_n_1channels(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_core_Mat_n_1checkVector__JIIZ(JNIEnv*, jclass, jlong, jint, jint, jboolean);
extern jint Java_org_opencv_core_Mat_n_1checkVector__JII(JNIEnv*, jclass, jlong, jint, jint);
extern jint Java_org_opencv_core_Mat_n_1checkVector__JI(JNIEnv*, jclass, jlong, jint);
extern jlong Java_org_opencv_core_Mat_n_1clone(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_core_Mat_n_1col(JNIEnv*, jclass, jlong, jint);
extern jlong Java_org_opencv_core_Mat_n_1colRange(JNIEnv*, jclass, jlong, jint, jint);
extern jint Java_org_opencv_core_Mat_n_1cols(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_core_Mat_n_1convertTo__JJIDD(JNIEnv*, jclass, jlong, jlong, jint, jdouble, jdouble);
extern void Java_org_opencv_core_Mat_n_1convertTo__JJID(JNIEnv*, jclass, jlong, jlong, jint, jdouble);
extern void Java_org_opencv_core_Mat_n_1convertTo__JJI(JNIEnv*, jclass, jlong, jlong, jint);
extern void Java_org_opencv_core_Mat_n_1copyTo__JJ(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_core_Mat_n_1copyTo__JJJ(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_core_Mat_n_1create__JIII(JNIEnv*, jclass, jlong, jint, jint, jint);
extern void Java_org_opencv_core_Mat_n_1create__JDDI(JNIEnv*, jclass, jlong, jdouble, jdouble, jint);
extern jlong Java_org_opencv_core_Mat_n_1cross(JNIEnv*, jclass, jlong, jlong);
extern jlong Java_org_opencv_core_Mat_n_1dataAddr(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_core_Mat_n_1delete(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_core_Mat_n_1depth(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_core_Mat_n_1diag__JI(JNIEnv*, jclass, jlong, jint);
extern jlong Java_org_opencv_core_Mat_n_1diag__J(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_core_Mat_n_1dims(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_core_Mat_n_1dot(JNIEnv*, jclass, jlong, jlong);
extern jlong Java_org_opencv_core_Mat_n_1elemSize(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_core_Mat_n_1elemSize1(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_core_Mat_n_1empty(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_core_Mat_n_1eye__III(JNIEnv*, jclass, jint, jint, jint);
extern jlong Java_org_opencv_core_Mat_n_1eye__DDI(JNIEnv*, jclass, jdouble, jdouble, jint);
extern jlong Java_org_opencv_core_Mat_n_1inv__JI(JNIEnv*, jclass, jlong, jint);
extern jlong Java_org_opencv_core_Mat_n_1inv__J(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_core_Mat_n_1isContinuous(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_core_Mat_n_1isSubmatrix(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_core_Mat_n_1mul__JJD(JNIEnv*, jclass, jlong, jlong, jdouble);
extern jlong Java_org_opencv_core_Mat_n_1mul__JJ(JNIEnv*, jclass, jlong, jlong);
extern jlong Java_org_opencv_core_Mat_n_1ones__III(JNIEnv*, jclass, jint, jint, jint);
extern jlong Java_org_opencv_core_Mat_n_1ones__DDI(JNIEnv*, jclass, jdouble, jdouble, jint);
extern void Java_org_opencv_core_Mat_n_1push_1back(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_core_Mat_n_1release(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_core_Mat_n_1reshape__JII(JNIEnv*, jclass, jlong, jint, jint);
extern jlong Java_org_opencv_core_Mat_n_1reshape__JI(JNIEnv*, jclass, jlong, jint);
extern jlong Java_org_opencv_core_Mat_n_1row(JNIEnv*, jclass, jlong, jint);
extern jlong Java_org_opencv_core_Mat_n_1rowRange(JNIEnv*, jclass, jlong, jint, jint);
extern jint Java_org_opencv_core_Mat_n_1rows(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_core_Mat_n_1setTo__JDDDD(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble);
extern jlong Java_org_opencv_core_Mat_n_1setTo__JDDDDJ(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jlong);
extern jlong Java_org_opencv_core_Mat_n_1setTo__JJJ(JNIEnv*, jclass, jlong, jlong, jlong);
extern jlong Java_org_opencv_core_Mat_n_1setTo__JJ(JNIEnv*, jclass, jlong, jlong);
extern jdoubleArray Java_org_opencv_core_Mat_n_1size(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_core_Mat_n_1step1__JI(JNIEnv*, jclass, jlong, jint);
extern jlong Java_org_opencv_core_Mat_n_1step1__J(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_core_Mat_n_1submat(JNIEnv*, jclass, jlong, jint, jint, jint, jint);
extern jlong Java_org_opencv_core_Mat_n_1submat_1rr(JNIEnv*, jclass, jlong, jint, jint, jint, jint);
extern jlong Java_org_opencv_core_Mat_n_1t(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_core_Mat_n_1total(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_core_Mat_n_1type(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_core_Mat_n_1zeros__III(JNIEnv*, jclass, jint, jint, jint);
extern jlong Java_org_opencv_core_Mat_n_1zeros__DDI(JNIEnv*, jclass, jdouble, jdouble, jint);
extern void Java_org_opencv_core_Core_LUT_10(JNIEnv*, jclass, jlong, jlong, jlong);
extern jdouble Java_org_opencv_core_Core_Mahalanobis_10(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_core_Core_PCABackProject_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_core_Core_PCACompute_10(JNIEnv*, jclass, jlong, jlong, jlong, jdouble);
extern void Java_org_opencv_core_Core_PCACompute_11(JNIEnv*, jclass, jlong, jlong, jlong, jint);
extern void Java_org_opencv_core_Core_PCACompute_12(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_core_Core_PCAProject_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern jdouble Java_org_opencv_core_Core_PSNR_10(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_core_Core_SVBackSubst_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_core_Core_SVDecomp_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jint);
extern void Java_org_opencv_core_Core_SVDecomp_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_core_Core_absdiff_10(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_core_Core_absdiff_11(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jlong);
extern void Java_org_opencv_core_Core_addWeighted_10(JNIEnv*, jclass, jlong, jdouble, jlong, jdouble, jdouble, jlong, jint);
extern void Java_org_opencv_core_Core_addWeighted_11(JNIEnv*, jclass, jlong, jdouble, jlong, jdouble, jdouble, jlong);
extern void Java_org_opencv_core_Core_add_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jint);
extern void Java_org_opencv_core_Core_add_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_core_Core_add_12(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_core_Core_add_13(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jlong, jlong, jint);
extern void Java_org_opencv_core_Core_add_14(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jlong, jlong);
extern void Java_org_opencv_core_Core_add_15(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jlong);
extern void Java_org_opencv_core_Core_batchDistance_10(JNIEnv*, jclass, jlong, jlong, jlong, jint, jlong, jint, jint, jlong, jint, jboolean);
extern void Java_org_opencv_core_Core_batchDistance_11(JNIEnv*, jclass, jlong, jlong, jlong, jint, jlong, jint, jint);
extern void Java_org_opencv_core_Core_batchDistance_12(JNIEnv*, jclass, jlong, jlong, jlong, jint, jlong);
extern void Java_org_opencv_core_Core_bitwise_1and_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_core_Core_bitwise_1and_11(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_core_Core_bitwise_1not_10(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_core_Core_bitwise_1not_11(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_core_Core_bitwise_1or_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_core_Core_bitwise_1or_11(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_core_Core_bitwise_1xor_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_core_Core_bitwise_1xor_11(JNIEnv*, jclass, jlong, jlong, jlong);
extern jint Java_org_opencv_core_Core_borderInterpolate_10(JNIEnv*, jclass, jint, jint, jint);
extern void Java_org_opencv_core_Core_calcCovarMatrix_10(JNIEnv*, jclass, jlong, jlong, jlong, jint, jint);
extern void Java_org_opencv_core_Core_calcCovarMatrix_11(JNIEnv*, jclass, jlong, jlong, jlong, jint);
extern void Java_org_opencv_core_Core_cartToPolar_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jboolean);
extern void Java_org_opencv_core_Core_cartToPolar_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern jboolean Java_org_opencv_core_Core_checkRange_10(JNIEnv*, jclass, jlong, jboolean, jdouble, jdouble);
extern jboolean Java_org_opencv_core_Core_checkRange_11(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_core_Core_compare_10(JNIEnv*, jclass, jlong, jlong, jlong, jint);
extern void Java_org_opencv_core_Core_compare_11(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jlong, jint);
extern void Java_org_opencv_core_Core_completeSymm_10(JNIEnv*, jclass, jlong, jboolean);
extern void Java_org_opencv_core_Core_completeSymm_11(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_core_Core_convertFp16_10(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_core_Core_convertScaleAbs_10(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble);
extern void Java_org_opencv_core_Core_convertScaleAbs_11(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_core_Core_copyMakeBorder_10(JNIEnv*, jclass, jlong, jlong, jint, jint, jint, jint, jint, jdouble, jdouble, jdouble, jdouble);
extern void Java_org_opencv_core_Core_copyMakeBorder_11(JNIEnv*, jclass, jlong, jlong, jint, jint, jint, jint, jint);
extern jint Java_org_opencv_core_Core_countNonZero_10(JNIEnv*, jclass, jlong);
extern jfloat Java_org_opencv_core_Core_cubeRoot_10(JNIEnv*, jclass, jfloat);
extern void Java_org_opencv_core_Core_dct_10(JNIEnv*, jclass, jlong, jlong, jint);
extern void Java_org_opencv_core_Core_dct_11(JNIEnv*, jclass, jlong, jlong);
extern jdouble Java_org_opencv_core_Core_determinant_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_core_Core_dft_10(JNIEnv*, jclass, jlong, jlong, jint, jint);
extern void Java_org_opencv_core_Core_dft_11(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_core_Core_divide_10(JNIEnv*, jclass, jlong, jlong, jlong, jdouble, jint);
extern void Java_org_opencv_core_Core_divide_11(JNIEnv*, jclass, jlong, jlong, jlong, jdouble);
extern void Java_org_opencv_core_Core_divide_12(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_core_Core_divide_13(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jlong, jdouble, jint);
extern void Java_org_opencv_core_Core_divide_14(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jlong, jdouble);
extern void Java_org_opencv_core_Core_divide_15(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jlong);
extern void Java_org_opencv_core_Core_divide_16(JNIEnv*, jclass, jdouble, jlong, jlong, jint);
extern void Java_org_opencv_core_Core_divide_17(JNIEnv*, jclass, jdouble, jlong, jlong);
extern void Java_org_opencv_core_Core_eigenNonSymmetric_10(JNIEnv*, jclass, jlong, jlong, jlong);
extern jboolean Java_org_opencv_core_Core_eigen_10(JNIEnv*, jclass, jlong, jlong, jlong);
extern jboolean Java_org_opencv_core_Core_eigen_11(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_core_Core_exp_10(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_core_Core_extractChannel_10(JNIEnv*, jclass, jlong, jlong, jint);
extern jfloat Java_org_opencv_core_Core_fastAtan2_10(JNIEnv*, jclass, jfloat, jfloat);
extern void Java_org_opencv_core_Core_findNonZero_10(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_core_Core_flip_10(JNIEnv*, jclass, jlong, jlong, jint);
extern void Java_org_opencv_core_Core_gemm_10(JNIEnv*, jclass, jlong, jlong, jdouble, jlong, jdouble, jlong, jint);
extern void Java_org_opencv_core_Core_gemm_11(JNIEnv*, jclass, jlong, jlong, jdouble, jlong, jdouble, jlong);
extern jstring Java_org_opencv_core_Core_getBuildInformation_10(JNIEnv*, jclass);
extern jlong Java_org_opencv_core_Core_getCPUTickCount_10(JNIEnv*, jclass);
extern jstring Java_org_opencv_core_Core_getIppVersion_10(JNIEnv*, jclass);
extern jint Java_org_opencv_core_Core_getNumThreads_10(JNIEnv*, jclass);
extern jint Java_org_opencv_core_Core_getNumberOfCPUs_10(JNIEnv*, jclass);
extern jint Java_org_opencv_core_Core_getOptimalDFTSize_10(JNIEnv*, jclass, jint);
extern jint Java_org_opencv_core_Core_getThreadNum_10(JNIEnv*, jclass);
extern jlong Java_org_opencv_core_Core_getTickCount_10(JNIEnv*, jclass);
extern jdouble Java_org_opencv_core_Core_getTickFrequency_10(JNIEnv*, jclass);
extern void Java_org_opencv_core_Core_hconcat_10(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_core_Core_idct_10(JNIEnv*, jclass, jlong, jlong, jint);
extern void Java_org_opencv_core_Core_idct_11(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_core_Core_idft_10(JNIEnv*, jclass, jlong, jlong, jint, jint);
extern void Java_org_opencv_core_Core_idft_11(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_core_Core_inRange_10(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jlong);
extern void Java_org_opencv_core_Core_insertChannel_10(JNIEnv*, jclass, jlong, jlong, jint);
extern jdouble Java_org_opencv_core_Core_invert_10(JNIEnv*, jclass, jlong, jlong, jint);
extern jdouble Java_org_opencv_core_Core_invert_11(JNIEnv*, jclass, jlong, jlong);
extern jdouble Java_org_opencv_core_Core_kmeans_10(JNIEnv*, jclass, jlong, jint, jlong, jint, jint, jdouble, jint, jint, jlong);
extern jdouble Java_org_opencv_core_Core_kmeans_11(JNIEnv*, jclass, jlong, jint, jlong, jint, jint, jdouble, jint, jint);
extern void Java_org_opencv_core_Core_log_10(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_core_Core_magnitude_10(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_core_Core_max_10(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_core_Core_max_11(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jlong);
extern void Java_org_opencv_core_Core_meanStdDev_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_core_Core_meanStdDev_11(JNIEnv*, jclass, jlong, jlong, jlong);
extern jdoubleArray Java_org_opencv_core_Core_mean_10(JNIEnv*, jclass, jlong, jlong);
extern jdoubleArray Java_org_opencv_core_Core_mean_11(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_core_Core_merge_10(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_core_Core_min_10(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_core_Core_min_11(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jlong);
extern void Java_org_opencv_core_Core_mixChannels_10(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_core_Core_mulSpectrums_10(JNIEnv*, jclass, jlong, jlong, jlong, jint, jboolean);
extern void Java_org_opencv_core_Core_mulSpectrums_11(JNIEnv*, jclass, jlong, jlong, jlong, jint);
extern void Java_org_opencv_core_Core_mulTransposed_10(JNIEnv*, jclass, jlong, jlong, jboolean, jlong, jdouble, jint);
extern void Java_org_opencv_core_Core_mulTransposed_11(JNIEnv*, jclass, jlong, jlong, jboolean, jlong, jdouble);
extern void Java_org_opencv_core_Core_mulTransposed_12(JNIEnv*, jclass, jlong, jlong, jboolean);
extern void Java_org_opencv_core_Core_multiply_10(JNIEnv*, jclass, jlong, jlong, jlong, jdouble, jint);
extern void Java_org_opencv_core_Core_multiply_11(JNIEnv*, jclass, jlong, jlong, jlong, jdouble);
extern void Java_org_opencv_core_Core_multiply_12(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_core_Core_multiply_13(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jlong, jdouble, jint);
extern void Java_org_opencv_core_Core_multiply_14(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jlong, jdouble);
extern void Java_org_opencv_core_Core_multiply_15(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jlong);
extern jdoubleArray Java_org_opencv_core_Core_n_1minMaxLocManual(JNIEnv*, jclass, jlong, jlong);
extern jdouble Java_org_opencv_core_Core_norm_10(JNIEnv*, jclass, jlong, jlong, jint, jlong);
extern jdouble Java_org_opencv_core_Core_norm_11(JNIEnv*, jclass, jlong, jlong, jint);
extern jdouble Java_org_opencv_core_Core_norm_12(JNIEnv*, jclass, jlong, jlong);
extern jdouble Java_org_opencv_core_Core_norm_13(JNIEnv*, jclass, jlong, jint, jlong);
extern jdouble Java_org_opencv_core_Core_norm_14(JNIEnv*, jclass, jlong, jint);
extern jdouble Java_org_opencv_core_Core_norm_15(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_core_Core_normalize_10(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jint, jint, jlong);
extern void Java_org_opencv_core_Core_normalize_11(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jint, jint);
extern void Java_org_opencv_core_Core_normalize_12(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jint);
extern void Java_org_opencv_core_Core_normalize_13(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_core_Core_patchNaNs_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_core_Core_patchNaNs_11(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_core_Core_perspectiveTransform_10(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_core_Core_phase_10(JNIEnv*, jclass, jlong, jlong, jlong, jboolean);
extern void Java_org_opencv_core_Core_phase_11(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_core_Core_polarToCart_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jboolean);
extern void Java_org_opencv_core_Core_polarToCart_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_core_Core_pow_10(JNIEnv*, jclass, jlong, jdouble, jlong);
extern void Java_org_opencv_core_Core_randShuffle_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_core_Core_randShuffle_11(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_core_Core_randn_10(JNIEnv*, jclass, jlong, jdouble, jdouble);
extern void Java_org_opencv_core_Core_randu_10(JNIEnv*, jclass, jlong, jdouble, jdouble);
extern void Java_org_opencv_core_Core_reduce_10(JNIEnv*, jclass, jlong, jlong, jint, jint, jint);
extern void Java_org_opencv_core_Core_reduce_11(JNIEnv*, jclass, jlong, jlong, jint, jint);
extern void Java_org_opencv_core_Core_repeat_10(JNIEnv*, jclass, jlong, jint, jint, jlong);
extern void Java_org_opencv_core_Core_rotate_10(JNIEnv*, jclass, jlong, jlong, jint);
extern void Java_org_opencv_core_Core_scaleAdd_10(JNIEnv*, jclass, jlong, jdouble, jlong, jlong);
extern void Java_org_opencv_core_Core_setErrorVerbosity_10(JNIEnv*, jclass, jboolean);
extern void Java_org_opencv_core_Core_setIdentity_10(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble);
extern void Java_org_opencv_core_Core_setIdentity_11(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_core_Core_setNumThreads_10(JNIEnv*, jclass, jint);
extern void Java_org_opencv_core_Core_setRNGSeed_10(JNIEnv*, jclass, jint);
extern void Java_org_opencv_core_Core_setUseIPP_10(JNIEnv*, jclass, jboolean);
extern void Java_org_opencv_core_Core_setUseIPP_1NE_10(JNIEnv*, jclass, jboolean);
extern jint Java_org_opencv_core_Core_solveCubic_10(JNIEnv*, jclass, jlong, jlong);
extern jdouble Java_org_opencv_core_Core_solvePoly_10(JNIEnv*, jclass, jlong, jlong, jint);
extern jdouble Java_org_opencv_core_Core_solvePoly_11(JNIEnv*, jclass, jlong, jlong);
extern jboolean Java_org_opencv_core_Core_solve_10(JNIEnv*, jclass, jlong, jlong, jlong, jint);
extern jboolean Java_org_opencv_core_Core_solve_11(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_core_Core_sortIdx_10(JNIEnv*, jclass, jlong, jlong, jint);
extern void Java_org_opencv_core_Core_sort_10(JNIEnv*, jclass, jlong, jlong, jint);
extern void Java_org_opencv_core_Core_split_10(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_core_Core_sqrt_10(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_core_Core_subtract_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jint);
extern void Java_org_opencv_core_Core_subtract_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_core_Core_subtract_12(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_core_Core_subtract_13(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jlong, jlong, jint);
extern void Java_org_opencv_core_Core_subtract_14(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jlong, jlong);
extern void Java_org_opencv_core_Core_subtract_15(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jlong);
extern jdoubleArray Java_org_opencv_core_Core_sumElems_10(JNIEnv*, jclass, jlong);
extern jdoubleArray Java_org_opencv_core_Core_trace_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_core_Core_transform_10(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_core_Core_transpose_10(JNIEnv*, jclass, jlong, jlong);
extern jboolean Java_org_opencv_core_Core_useIPP_10(JNIEnv*, jclass);
extern jboolean Java_org_opencv_core_Core_useIPP_1NE_10(JNIEnv*, jclass);
extern void Java_org_opencv_core_Core_vconcat_10(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_core_Algorithm_clear_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_core_Algorithm_delete(JNIEnv*, jclass, jlong);
extern jstring Java_org_opencv_core_Algorithm_getDefaultName_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_core_Algorithm_save_10(JNIEnv*, jclass, jlong, jstring);
extern jlong Java_org_opencv_core_TickMeter_TickMeter_10(JNIEnv*, jclass);
extern void Java_org_opencv_core_TickMeter_delete(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_core_TickMeter_getCounter_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_core_TickMeter_getTimeMicro_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_core_TickMeter_getTimeMilli_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_core_TickMeter_getTimeSec_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_core_TickMeter_getTimeTicks_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_core_TickMeter_reset_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_core_TickMeter_start_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_core_TickMeter_stop_10(JNIEnv*, jclass, jlong);

*/
import "C"

func MatNative_nDump(self int64) string {
	return togostring(C.Java_org_opencv_core_Mat_nDump(clzEnv, clzObj, (C.jlong)(self)))
}

func MatNative_nGet(self int64, row int, col int) []float64 {
	return togofloat64Array(C.Java_org_opencv_core_Mat_nGet(clzEnv, clzObj, (C.jlong)(self), (C.jint)(row), (C.jint)(col)))
}

func MatNative_nGetB(self int64, row int, col int, count int, vals []byte) int {
	defer ungointerface(vals)
	return int(C.Java_org_opencv_core_Mat_nGetB(clzEnv, clzObj, (C.jlong)(self), (C.jint)(row), (C.jint)(col), (C.jint)(count), tojbyteArray(vals)))
}

func MatNative_nGetD(self int64, row int, col int, count int, vals []float64) int {
	defer ungointerface(vals)
	return int(C.Java_org_opencv_core_Mat_nGetD(clzEnv, clzObj, (C.jlong)(self), (C.jint)(row), (C.jint)(col), (C.jint)(count), tojdoubleArray(vals)))
}

func MatNative_nGetF(self int64, row int, col int, count int, vals []float32) int {
	defer ungointerface(vals)
	return int(C.Java_org_opencv_core_Mat_nGetF(clzEnv, clzObj, (C.jlong)(self), (C.jint)(row), (C.jint)(col), (C.jint)(count), tojfloatArray(vals)))
}

func MatNative_nGetI(self int64, row int, col int, count int, vals []int32) int {
	defer ungointerface(vals)
	return int(C.Java_org_opencv_core_Mat_nGetI(clzEnv, clzObj, (C.jlong)(self), (C.jint)(row), (C.jint)(col), (C.jint)(count), tojintArray(vals)))
}

func MatNative_nGetS(self int64, row int, col int, count int, vals []int16) int {
	defer ungointerface(vals)
	return int(C.Java_org_opencv_core_Mat_nGetS(clzEnv, clzObj, (C.jlong)(self), (C.jint)(row), (C.jint)(col), (C.jint)(count), tojshortArray(vals)))
}

func MatNative_nPutB(self int64, row int, col int, count int, data []byte) int {
	defer ungointerface(data)
	return int(C.Java_org_opencv_core_Mat_nPutB(clzEnv, clzObj, (C.jlong)(self), (C.jint)(row), (C.jint)(col), (C.jint)(count), tojbyteArray(data)))
}

func MatNative_nPutD(self int64, row int, col int, count int, data []float64) int {
	defer ungointerface(data)
	return int(C.Java_org_opencv_core_Mat_nPutD(clzEnv, clzObj, (C.jlong)(self), (C.jint)(row), (C.jint)(col), (C.jint)(count), tojdoubleArray(data)))
}

func MatNative_nPutF(self int64, row int, col int, count int, data []float32) int {
	defer ungointerface(data)
	return int(C.Java_org_opencv_core_Mat_nPutF(clzEnv, clzObj, (C.jlong)(self), (C.jint)(row), (C.jint)(col), (C.jint)(count), tojfloatArray(data)))
}

func MatNative_nPutI(self int64, row int, col int, count int, data []int32) int {
	defer ungointerface(data)
	return int(C.Java_org_opencv_core_Mat_nPutI(clzEnv, clzObj, (C.jlong)(self), (C.jint)(row), (C.jint)(col), (C.jint)(count), tojintArray(data)))
}

func MatNative_nPutS(self int64, row int, col int, count int, data []int16) int {
	defer ungointerface(data)
	return int(C.Java_org_opencv_core_Mat_nPutS(clzEnv, clzObj, (C.jlong)(self), (C.jint)(row), (C.jint)(col), (C.jint)(count), tojshortArray(data)))
}

func MatNative_n_Mat() int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1Mat__(clzEnv, clzObj))
}

func MatNative_n_Mat2(rows int, cols int, rtype int) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1Mat__III(clzEnv, clzObj, (C.jint)(rows), (C.jint)(cols), (C.jint)(rtype)))
}

func MatNative_n_Mat3(size_width float64, size_height float64, rtype int) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1Mat__DDI(clzEnv, clzObj, (C.jdouble)(size_width), (C.jdouble)(size_height), (C.jint)(rtype)))
}

func MatNative_n_Mat4(rows int, cols int, rtype int, s_val0 float64, s_val1 float64, s_val2 float64, s_val3 float64) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1Mat__IIIDDDD(clzEnv, clzObj, (C.jint)(rows), (C.jint)(cols), (C.jint)(rtype), (C.jdouble)(s_val0), (C.jdouble)(s_val1), (C.jdouble)(s_val2), (C.jdouble)(s_val3)))
}

func MatNative_n_Mat5(size_width float64, size_height float64, rtype int, s_val0 float64, s_val1 float64, s_val2 float64, s_val3 float64) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1Mat__DDIDDDD(clzEnv, clzObj, (C.jdouble)(size_width), (C.jdouble)(size_height), (C.jint)(rtype), (C.jdouble)(s_val0), (C.jdouble)(s_val1), (C.jdouble)(s_val2), (C.jdouble)(s_val3)))
}

func MatNative_n_Mat6(m_nativeObj int64, rowRange_start int, rowRange_end int, colRange_start int, colRange_end int) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1Mat__JIIII(clzEnv, clzObj, (C.jlong)(m_nativeObj), (C.jint)(rowRange_start), (C.jint)(rowRange_end), (C.jint)(colRange_start), (C.jint)(colRange_end)))
}

func MatNative_n_Mat7(m_nativeObj int64, rowRange_start int, rowRange_end int) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1Mat__JII(clzEnv, clzObj, (C.jlong)(m_nativeObj), (C.jint)(rowRange_start), (C.jint)(rowRange_end)))
}

func MatNative_n_adjustROI(nativeObj int64, dtop int, dbottom int, dleft int, dright int) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1adjustROI(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(dtop), (C.jint)(dbottom), (C.jint)(dleft), (C.jint)(dright)))
}

func MatNative_n_assignTo(nativeObj int64, m_nativeObj int64, rtype int) {
	C.Java_org_opencv_core_Mat_n_1assignTo__JJI(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(m_nativeObj), (C.jint)(rtype))
}

func MatNative_n_assignTo2(nativeObj int64, m_nativeObj int64) {
	C.Java_org_opencv_core_Mat_n_1assignTo__JJ(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(m_nativeObj))
}

func MatNative_n_channels(nativeObj int64) int {
	return int(C.Java_org_opencv_core_Mat_n_1channels(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func MatNative_n_checkVector(nativeObj int64, elemChannels int, depth int, requireContinuous bool) int {
	return int(C.Java_org_opencv_core_Mat_n_1checkVector__JIIZ(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(elemChannels), (C.jint)(depth), tojboolean(requireContinuous)))
}

func MatNative_n_checkVector2(nativeObj int64, elemChannels int, depth int) int {
	return int(C.Java_org_opencv_core_Mat_n_1checkVector__JII(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(elemChannels), (C.jint)(depth)))
}

func MatNative_n_checkVector3(nativeObj int64, elemChannels int) int {
	return int(C.Java_org_opencv_core_Mat_n_1checkVector__JI(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(elemChannels)))
}

func MatNative_n_clone(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1clone(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func MatNative_n_col(nativeObj int64, x int) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1col(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(x)))
}

func MatNative_n_colRange(nativeObj int64, startcol int, endcol int) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1colRange(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(startcol), (C.jint)(endcol)))
}

func MatNative_n_cols(nativeObj int64) int {
	return int(C.Java_org_opencv_core_Mat_n_1cols(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func MatNative_n_convertTo(nativeObj int64, m_nativeObj int64, rtype int, alpha float64, beta float64) {
	C.Java_org_opencv_core_Mat_n_1convertTo__JJIDD(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(m_nativeObj), (C.jint)(rtype), (C.jdouble)(alpha), (C.jdouble)(beta))
}

func MatNative_n_convertTo2(nativeObj int64, m_nativeObj int64, rtype int, alpha float64) {
	C.Java_org_opencv_core_Mat_n_1convertTo__JJID(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(m_nativeObj), (C.jint)(rtype), (C.jdouble)(alpha))
}

func MatNative_n_convertTo3(nativeObj int64, m_nativeObj int64, rtype int) {
	C.Java_org_opencv_core_Mat_n_1convertTo__JJI(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(m_nativeObj), (C.jint)(rtype))
}

func MatNative_n_copyTo(nativeObj int64, m_nativeObj int64) {
	C.Java_org_opencv_core_Mat_n_1copyTo__JJ(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(m_nativeObj))
}

func MatNative_n_copyTo2(nativeObj int64, m_nativeObj int64, mask_nativeObj int64) {
	C.Java_org_opencv_core_Mat_n_1copyTo__JJJ(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(m_nativeObj), (C.jlong)(mask_nativeObj))
}

func MatNative_n_create(nativeObj int64, rows int, cols int, rtype int) {
	C.Java_org_opencv_core_Mat_n_1create__JIII(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(rows), (C.jint)(cols), (C.jint)(rtype))
}

func MatNative_n_create2(nativeObj int64, size_width float64, size_height float64, rtype int) {
	C.Java_org_opencv_core_Mat_n_1create__JDDI(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jdouble)(size_width), (C.jdouble)(size_height), (C.jint)(rtype))
}

func MatNative_n_cross(nativeObj int64, m_nativeObj int64) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1cross(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(m_nativeObj)))
}

func MatNative_n_dataAddr(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1dataAddr(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func MatNative_n_delete(nativeObj int64) {
	C.Java_org_opencv_core_Mat_n_1delete(clzEnv, clzObj, (C.jlong)(nativeObj))
}

func MatNative_n_depth(nativeObj int64) int {
	return int(C.Java_org_opencv_core_Mat_n_1depth(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func MatNative_n_diag(nativeObj int64, d int) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1diag__JI(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(d)))
}

func MatNative_n_diag2(d_nativeObj int64) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1diag__J(clzEnv, clzObj, (C.jlong)(d_nativeObj)))
}

func MatNative_n_dims(nativeObj int64) int {
	return int(C.Java_org_opencv_core_Mat_n_1dims(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func MatNative_n_dot(nativeObj int64, m_nativeObj int64) float64 {
	return float64(C.Java_org_opencv_core_Mat_n_1dot(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(m_nativeObj)))
}

func MatNative_n_elemSize(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1elemSize(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func MatNative_n_elemSize1(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1elemSize1(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func MatNative_n_empty(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_core_Mat_n_1empty(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func MatNative_n_eye(rows int, cols int, rtype int) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1eye__III(clzEnv, clzObj, (C.jint)(rows), (C.jint)(cols), (C.jint)(rtype)))
}

func MatNative_n_eye2(size_width float64, size_height float64, rtype int) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1eye__DDI(clzEnv, clzObj, (C.jdouble)(size_width), (C.jdouble)(size_height), (C.jint)(rtype)))
}

func MatNative_n_inv(nativeObj int64, method int) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1inv__JI(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(method)))
}

func MatNative_n_inv2(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1inv__J(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func MatNative_n_isContinuous(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_core_Mat_n_1isContinuous(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func MatNative_n_isSubmatrix(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_core_Mat_n_1isSubmatrix(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func MatNative_n_mul(nativeObj int64, m_nativeObj int64, scale float64) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1mul__JJD(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(m_nativeObj), (C.jdouble)(scale)))
}

func MatNative_n_mul2(nativeObj int64, m_nativeObj int64) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1mul__JJ(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(m_nativeObj)))
}

func MatNative_n_ones(rows int, cols int, rtype int) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1ones__III(clzEnv, clzObj, (C.jint)(rows), (C.jint)(cols), (C.jint)(rtype)))
}

func MatNative_n_ones2(size_width float64, size_height float64, rtype int) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1ones__DDI(clzEnv, clzObj, (C.jdouble)(size_width), (C.jdouble)(size_height), (C.jint)(rtype)))
}

func MatNative_n_push_back(nativeObj int64, m_nativeObj int64) {
	C.Java_org_opencv_core_Mat_n_1push_1back(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(m_nativeObj))
}

func MatNative_n_release(nativeObj int64) {
	C.Java_org_opencv_core_Mat_n_1release(clzEnv, clzObj, (C.jlong)(nativeObj))
}

func MatNative_n_reshape(nativeObj int64, cn int, rows int) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1reshape__JII(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(cn), (C.jint)(rows)))
}

func MatNative_n_reshape2(nativeObj int64, cn int) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1reshape__JI(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(cn)))
}

func MatNative_n_row(nativeObj int64, y int) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1row(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(y)))
}

func MatNative_n_rowRange(nativeObj int64, startrow int, endrow int) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1rowRange(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(startrow), (C.jint)(endrow)))
}

func MatNative_n_rows(nativeObj int64) int {
	return int(C.Java_org_opencv_core_Mat_n_1rows(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func MatNative_n_setTo(nativeObj int64, s_val0 float64, s_val1 float64, s_val2 float64, s_val3 float64) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1setTo__JDDDD(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jdouble)(s_val0), (C.jdouble)(s_val1), (C.jdouble)(s_val2), (C.jdouble)(s_val3)))
}

func MatNative_n_setTo2(nativeObj int64, s_val0 float64, s_val1 float64, s_val2 float64, s_val3 float64, mask_nativeObj int64) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1setTo__JDDDDJ(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jdouble)(s_val0), (C.jdouble)(s_val1), (C.jdouble)(s_val2), (C.jdouble)(s_val3), (C.jlong)(mask_nativeObj)))
}

func MatNative_n_setTo3(nativeObj int64, value_nativeObj int64, mask_nativeObj int64) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1setTo__JJJ(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(value_nativeObj), (C.jlong)(mask_nativeObj)))
}

func MatNative_n_setTo4(nativeObj int64, value_nativeObj int64) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1setTo__JJ(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(value_nativeObj)))
}

func MatNative_n_size(nativeObj int64) []float64 {
	return togofloat64Array(C.Java_org_opencv_core_Mat_n_1size(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func MatNative_n_step1(nativeObj int64, i int) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1step1__JI(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(i)))
}

func MatNative_n_step12(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1step1__J(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func MatNative_n_submat(nativeObj int64, roi_x int, roi_y int, roi_width int, roi_height int) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1submat(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(roi_x), (C.jint)(roi_y), (C.jint)(roi_width), (C.jint)(roi_height)))
}

func MatNative_n_submat_rr(nativeObj int64, rowRange_start int, rowRange_end int, colRange_start int, colRange_end int) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1submat_1rr(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(rowRange_start), (C.jint)(rowRange_end), (C.jint)(colRange_start), (C.jint)(colRange_end)))
}

func MatNative_n_t(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1t(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func MatNative_n_total(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1total(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func MatNative_n_type(nativeObj int64) int {
	return int(C.Java_org_opencv_core_Mat_n_1type(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func MatNative_n_zeros(rows int, cols int, rtype int) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1zeros__III(clzEnv, clzObj, (C.jint)(rows), (C.jint)(cols), (C.jint)(rtype)))
}

func MatNative_n_zeros2(size_width float64, size_height float64, rtype int) int64 {
	return int64(C.Java_org_opencv_core_Mat_n_1zeros__DDI(clzEnv, clzObj, (C.jdouble)(size_width), (C.jdouble)(size_height), (C.jint)(rtype)))
}

func CoreNative_LUT_0(src_nativeObj int64, lut_nativeObj int64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_LUT_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(lut_nativeObj), (C.jlong)(dst_nativeObj))
}

func CoreNative_Mahalanobis_0(v1_nativeObj int64, v2_nativeObj int64, icovar_nativeObj int64) float64 {
	return float64(C.Java_org_opencv_core_Core_Mahalanobis_10(clzEnv, clzObj, (C.jlong)(v1_nativeObj), (C.jlong)(v2_nativeObj), (C.jlong)(icovar_nativeObj)))
}

func CoreNative_PCABackProject_0(data_nativeObj int64, mean_nativeObj int64, eigenvectors_nativeObj int64, result_nativeObj int64) {
	C.Java_org_opencv_core_Core_PCABackProject_10(clzEnv, clzObj, (C.jlong)(data_nativeObj), (C.jlong)(mean_nativeObj), (C.jlong)(eigenvectors_nativeObj), (C.jlong)(result_nativeObj))
}

func CoreNative_PCACompute_0(data_nativeObj int64, mean_nativeObj int64, eigenvectors_nativeObj int64, retainedVariance float64) {
	C.Java_org_opencv_core_Core_PCACompute_10(clzEnv, clzObj, (C.jlong)(data_nativeObj), (C.jlong)(mean_nativeObj), (C.jlong)(eigenvectors_nativeObj), (C.jdouble)(retainedVariance))
}

func CoreNative_PCACompute_1(data_nativeObj int64, mean_nativeObj int64, eigenvectors_nativeObj int64, maxComponents int) {
	C.Java_org_opencv_core_Core_PCACompute_11(clzEnv, clzObj, (C.jlong)(data_nativeObj), (C.jlong)(mean_nativeObj), (C.jlong)(eigenvectors_nativeObj), (C.jint)(maxComponents))
}

func CoreNative_PCACompute_2(data_nativeObj int64, mean_nativeObj int64, eigenvectors_nativeObj int64) {
	C.Java_org_opencv_core_Core_PCACompute_12(clzEnv, clzObj, (C.jlong)(data_nativeObj), (C.jlong)(mean_nativeObj), (C.jlong)(eigenvectors_nativeObj))
}

func CoreNative_PCAProject_0(data_nativeObj int64, mean_nativeObj int64, eigenvectors_nativeObj int64, result_nativeObj int64) {
	C.Java_org_opencv_core_Core_PCAProject_10(clzEnv, clzObj, (C.jlong)(data_nativeObj), (C.jlong)(mean_nativeObj), (C.jlong)(eigenvectors_nativeObj), (C.jlong)(result_nativeObj))
}

func CoreNative_PSNR_0(src1_nativeObj int64, src2_nativeObj int64) float64 {
	return float64(C.Java_org_opencv_core_Core_PSNR_10(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jlong)(src2_nativeObj)))
}

func CoreNative_SVBackSubst_0(w_nativeObj int64, u_nativeObj int64, vt_nativeObj int64, rhs_nativeObj int64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_SVBackSubst_10(clzEnv, clzObj, (C.jlong)(w_nativeObj), (C.jlong)(u_nativeObj), (C.jlong)(vt_nativeObj), (C.jlong)(rhs_nativeObj), (C.jlong)(dst_nativeObj))
}

func CoreNative_SVDecomp_0(src_nativeObj int64, w_nativeObj int64, u_nativeObj int64, vt_nativeObj int64, flags int) {
	C.Java_org_opencv_core_Core_SVDecomp_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(w_nativeObj), (C.jlong)(u_nativeObj), (C.jlong)(vt_nativeObj), (C.jint)(flags))
}

func CoreNative_SVDecomp_1(src_nativeObj int64, w_nativeObj int64, u_nativeObj int64, vt_nativeObj int64) {
	C.Java_org_opencv_core_Core_SVDecomp_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(w_nativeObj), (C.jlong)(u_nativeObj), (C.jlong)(vt_nativeObj))
}

func CoreNative_absdiff_0(src1_nativeObj int64, src2_nativeObj int64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_absdiff_10(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jlong)(src2_nativeObj), (C.jlong)(dst_nativeObj))
}

func CoreNative_absdiff_1(src1_nativeObj int64, src2_val0 float64, src2_val1 float64, src2_val2 float64, src2_val3 float64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_absdiff_11(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jdouble)(src2_val0), (C.jdouble)(src2_val1), (C.jdouble)(src2_val2), (C.jdouble)(src2_val3), (C.jlong)(dst_nativeObj))
}

func CoreNative_addWeighted_0(src1_nativeObj int64, alpha float64, src2_nativeObj int64, beta float64, gamma float64, dst_nativeObj int64, dtype int) {
	C.Java_org_opencv_core_Core_addWeighted_10(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jdouble)(alpha), (C.jlong)(src2_nativeObj), (C.jdouble)(beta), (C.jdouble)(gamma), (C.jlong)(dst_nativeObj), (C.jint)(dtype))
}

func CoreNative_addWeighted_1(src1_nativeObj int64, alpha float64, src2_nativeObj int64, beta float64, gamma float64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_addWeighted_11(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jdouble)(alpha), (C.jlong)(src2_nativeObj), (C.jdouble)(beta), (C.jdouble)(gamma), (C.jlong)(dst_nativeObj))
}

func CoreNative_add_0(src1_nativeObj int64, src2_nativeObj int64, dst_nativeObj int64, mask_nativeObj int64, dtype int) {
	C.Java_org_opencv_core_Core_add_10(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jlong)(src2_nativeObj), (C.jlong)(dst_nativeObj), (C.jlong)(mask_nativeObj), (C.jint)(dtype))
}

func CoreNative_add_1(src1_nativeObj int64, src2_nativeObj int64, dst_nativeObj int64, mask_nativeObj int64) {
	C.Java_org_opencv_core_Core_add_11(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jlong)(src2_nativeObj), (C.jlong)(dst_nativeObj), (C.jlong)(mask_nativeObj))
}

func CoreNative_add_2(src1_nativeObj int64, src2_nativeObj int64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_add_12(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jlong)(src2_nativeObj), (C.jlong)(dst_nativeObj))
}

func CoreNative_add_3(src1_nativeObj int64, src2_val0 float64, src2_val1 float64, src2_val2 float64, src2_val3 float64, dst_nativeObj int64, mask_nativeObj int64, dtype int) {
	C.Java_org_opencv_core_Core_add_13(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jdouble)(src2_val0), (C.jdouble)(src2_val1), (C.jdouble)(src2_val2), (C.jdouble)(src2_val3), (C.jlong)(dst_nativeObj), (C.jlong)(mask_nativeObj), (C.jint)(dtype))
}

func CoreNative_add_4(src1_nativeObj int64, src2_val0 float64, src2_val1 float64, src2_val2 float64, src2_val3 float64, dst_nativeObj int64, mask_nativeObj int64) {
	C.Java_org_opencv_core_Core_add_14(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jdouble)(src2_val0), (C.jdouble)(src2_val1), (C.jdouble)(src2_val2), (C.jdouble)(src2_val3), (C.jlong)(dst_nativeObj), (C.jlong)(mask_nativeObj))
}

func CoreNative_add_5(src1_nativeObj int64, src2_val0 float64, src2_val1 float64, src2_val2 float64, src2_val3 float64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_add_15(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jdouble)(src2_val0), (C.jdouble)(src2_val1), (C.jdouble)(src2_val2), (C.jdouble)(src2_val3), (C.jlong)(dst_nativeObj))
}

func CoreNative_batchDistance_0(src1_nativeObj int64, src2_nativeObj int64, dist_nativeObj int64, dtype int, nidx_nativeObj int64, normType int, K int, mask_nativeObj int64, update int, crosscheck bool) {
	C.Java_org_opencv_core_Core_batchDistance_10(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jlong)(src2_nativeObj), (C.jlong)(dist_nativeObj), (C.jint)(dtype), (C.jlong)(nidx_nativeObj), (C.jint)(normType), (C.jint)(K), (C.jlong)(mask_nativeObj), (C.jint)(update), tojboolean(crosscheck))
}

func CoreNative_batchDistance_1(src1_nativeObj int64, src2_nativeObj int64, dist_nativeObj int64, dtype int, nidx_nativeObj int64, normType int, K int) {
	C.Java_org_opencv_core_Core_batchDistance_11(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jlong)(src2_nativeObj), (C.jlong)(dist_nativeObj), (C.jint)(dtype), (C.jlong)(nidx_nativeObj), (C.jint)(normType), (C.jint)(K))
}

func CoreNative_batchDistance_2(src1_nativeObj int64, src2_nativeObj int64, dist_nativeObj int64, dtype int, nidx_nativeObj int64) {
	C.Java_org_opencv_core_Core_batchDistance_12(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jlong)(src2_nativeObj), (C.jlong)(dist_nativeObj), (C.jint)(dtype), (C.jlong)(nidx_nativeObj))
}

func CoreNative_bitwise_and_0(src1_nativeObj int64, src2_nativeObj int64, dst_nativeObj int64, mask_nativeObj int64) {
	C.Java_org_opencv_core_Core_bitwise_1and_10(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jlong)(src2_nativeObj), (C.jlong)(dst_nativeObj), (C.jlong)(mask_nativeObj))
}

func CoreNative_bitwise_and_1(src1_nativeObj int64, src2_nativeObj int64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_bitwise_1and_11(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jlong)(src2_nativeObj), (C.jlong)(dst_nativeObj))
}

func CoreNative_bitwise_not_0(src_nativeObj int64, dst_nativeObj int64, mask_nativeObj int64) {
	C.Java_org_opencv_core_Core_bitwise_1not_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jlong)(mask_nativeObj))
}

func CoreNative_bitwise_not_1(src_nativeObj int64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_bitwise_1not_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj))
}

func CoreNative_bitwise_or_0(src1_nativeObj int64, src2_nativeObj int64, dst_nativeObj int64, mask_nativeObj int64) {
	C.Java_org_opencv_core_Core_bitwise_1or_10(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jlong)(src2_nativeObj), (C.jlong)(dst_nativeObj), (C.jlong)(mask_nativeObj))
}

func CoreNative_bitwise_or_1(src1_nativeObj int64, src2_nativeObj int64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_bitwise_1or_11(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jlong)(src2_nativeObj), (C.jlong)(dst_nativeObj))
}

func CoreNative_bitwise_xor_0(src1_nativeObj int64, src2_nativeObj int64, dst_nativeObj int64, mask_nativeObj int64) {
	C.Java_org_opencv_core_Core_bitwise_1xor_10(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jlong)(src2_nativeObj), (C.jlong)(dst_nativeObj), (C.jlong)(mask_nativeObj))
}

func CoreNative_bitwise_xor_1(src1_nativeObj int64, src2_nativeObj int64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_bitwise_1xor_11(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jlong)(src2_nativeObj), (C.jlong)(dst_nativeObj))
}

func CoreNative_borderInterpolate_0(p int, length int, borderType int) int {
	return int(C.Java_org_opencv_core_Core_borderInterpolate_10(clzEnv, clzObj, (C.jint)(p), (C.jint)(length), (C.jint)(borderType)))
}

func CoreNative_calcCovarMatrix_0(samples_nativeObj int64, covar_nativeObj int64, mean_nativeObj int64, flags int, ctype int) {
	C.Java_org_opencv_core_Core_calcCovarMatrix_10(clzEnv, clzObj, (C.jlong)(samples_nativeObj), (C.jlong)(covar_nativeObj), (C.jlong)(mean_nativeObj), (C.jint)(flags), (C.jint)(ctype))
}

func CoreNative_calcCovarMatrix_1(samples_nativeObj int64, covar_nativeObj int64, mean_nativeObj int64, flags int) {
	C.Java_org_opencv_core_Core_calcCovarMatrix_11(clzEnv, clzObj, (C.jlong)(samples_nativeObj), (C.jlong)(covar_nativeObj), (C.jlong)(mean_nativeObj), (C.jint)(flags))
}

func CoreNative_cartToPolar_0(x_nativeObj int64, y_nativeObj int64, magnitude_nativeObj int64, angle_nativeObj int64, angleInDegrees bool) {
	C.Java_org_opencv_core_Core_cartToPolar_10(clzEnv, clzObj, (C.jlong)(x_nativeObj), (C.jlong)(y_nativeObj), (C.jlong)(magnitude_nativeObj), (C.jlong)(angle_nativeObj), tojboolean(angleInDegrees))
}

func CoreNative_cartToPolar_1(x_nativeObj int64, y_nativeObj int64, magnitude_nativeObj int64, angle_nativeObj int64) {
	C.Java_org_opencv_core_Core_cartToPolar_11(clzEnv, clzObj, (C.jlong)(x_nativeObj), (C.jlong)(y_nativeObj), (C.jlong)(magnitude_nativeObj), (C.jlong)(angle_nativeObj))
}

func CoreNative_checkRange_0(a_nativeObj int64, quiet bool, minVal float64, maxVal float64) bool {
	return togobool(C.Java_org_opencv_core_Core_checkRange_10(clzEnv, clzObj, (C.jlong)(a_nativeObj), tojboolean(quiet), (C.jdouble)(minVal), (C.jdouble)(maxVal)))
}

func CoreNative_checkRange_1(a_nativeObj int64) bool {
	return togobool(C.Java_org_opencv_core_Core_checkRange_11(clzEnv, clzObj, (C.jlong)(a_nativeObj)))
}

func CoreNative_compare_0(src1_nativeObj int64, src2_nativeObj int64, dst_nativeObj int64, cmpop int) {
	C.Java_org_opencv_core_Core_compare_10(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jlong)(src2_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(cmpop))
}

func CoreNative_compare_1(src1_nativeObj int64, src2_val0 float64, src2_val1 float64, src2_val2 float64, src2_val3 float64, dst_nativeObj int64, cmpop int) {
	C.Java_org_opencv_core_Core_compare_11(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jdouble)(src2_val0), (C.jdouble)(src2_val1), (C.jdouble)(src2_val2), (C.jdouble)(src2_val3), (C.jlong)(dst_nativeObj), (C.jint)(cmpop))
}

func CoreNative_completeSymm_0(mtx_nativeObj int64, lowerToUpper bool) {
	C.Java_org_opencv_core_Core_completeSymm_10(clzEnv, clzObj, (C.jlong)(mtx_nativeObj), tojboolean(lowerToUpper))
}

func CoreNative_completeSymm_1(mtx_nativeObj int64) {
	C.Java_org_opencv_core_Core_completeSymm_11(clzEnv, clzObj, (C.jlong)(mtx_nativeObj))
}

func CoreNative_convertFp16_0(src_nativeObj int64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_convertFp16_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj))
}

func CoreNative_convertScaleAbs_0(src_nativeObj int64, dst_nativeObj int64, alpha float64, beta float64) {
	C.Java_org_opencv_core_Core_convertScaleAbs_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jdouble)(alpha), (C.jdouble)(beta))
}

func CoreNative_convertScaleAbs_1(src_nativeObj int64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_convertScaleAbs_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj))
}

func CoreNative_copyMakeBorder_0(src_nativeObj int64, dst_nativeObj int64, top int, bottom int, left int, right int, borderType int, value_val0 float64, value_val1 float64, value_val2 float64, value_val3 float64) {
	C.Java_org_opencv_core_Core_copyMakeBorder_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(top), (C.jint)(bottom), (C.jint)(left), (C.jint)(right), (C.jint)(borderType), (C.jdouble)(value_val0), (C.jdouble)(value_val1), (C.jdouble)(value_val2), (C.jdouble)(value_val3))
}

func CoreNative_copyMakeBorder_1(src_nativeObj int64, dst_nativeObj int64, top int, bottom int, left int, right int, borderType int) {
	C.Java_org_opencv_core_Core_copyMakeBorder_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(top), (C.jint)(bottom), (C.jint)(left), (C.jint)(right), (C.jint)(borderType))
}

func CoreNative_countNonZero_0(src_nativeObj int64) int {
	return int(C.Java_org_opencv_core_Core_countNonZero_10(clzEnv, clzObj, (C.jlong)(src_nativeObj)))
}

func CoreNative_cubeRoot_0(val float32) float32 {
	return float32(C.Java_org_opencv_core_Core_cubeRoot_10(clzEnv, clzObj, (C.jfloat)(val)))
}

func CoreNative_dct_0(src_nativeObj int64, dst_nativeObj int64, flags int) {
	C.Java_org_opencv_core_Core_dct_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(flags))
}

func CoreNative_dct_1(src_nativeObj int64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_dct_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj))
}

func CoreNative_determinant_0(mtx_nativeObj int64) float64 {
	return float64(C.Java_org_opencv_core_Core_determinant_10(clzEnv, clzObj, (C.jlong)(mtx_nativeObj)))
}

func CoreNative_dft_0(src_nativeObj int64, dst_nativeObj int64, flags int, nonzeroRows int) {
	C.Java_org_opencv_core_Core_dft_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(flags), (C.jint)(nonzeroRows))
}

func CoreNative_dft_1(src_nativeObj int64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_dft_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj))
}

func CoreNative_divide_0(src1_nativeObj int64, src2_nativeObj int64, dst_nativeObj int64, scale float64, dtype int) {
	C.Java_org_opencv_core_Core_divide_10(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jlong)(src2_nativeObj), (C.jlong)(dst_nativeObj), (C.jdouble)(scale), (C.jint)(dtype))
}

func CoreNative_divide_1(src1_nativeObj int64, src2_nativeObj int64, dst_nativeObj int64, scale float64) {
	C.Java_org_opencv_core_Core_divide_11(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jlong)(src2_nativeObj), (C.jlong)(dst_nativeObj), (C.jdouble)(scale))
}

func CoreNative_divide_2(src1_nativeObj int64, src2_nativeObj int64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_divide_12(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jlong)(src2_nativeObj), (C.jlong)(dst_nativeObj))
}

func CoreNative_divide_3(src1_nativeObj int64, src2_val0 float64, src2_val1 float64, src2_val2 float64, src2_val3 float64, dst_nativeObj int64, scale float64, dtype int) {
	C.Java_org_opencv_core_Core_divide_13(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jdouble)(src2_val0), (C.jdouble)(src2_val1), (C.jdouble)(src2_val2), (C.jdouble)(src2_val3), (C.jlong)(dst_nativeObj), (C.jdouble)(scale), (C.jint)(dtype))
}

func CoreNative_divide_4(src1_nativeObj int64, src2_val0 float64, src2_val1 float64, src2_val2 float64, src2_val3 float64, dst_nativeObj int64, scale float64) {
	C.Java_org_opencv_core_Core_divide_14(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jdouble)(src2_val0), (C.jdouble)(src2_val1), (C.jdouble)(src2_val2), (C.jdouble)(src2_val3), (C.jlong)(dst_nativeObj), (C.jdouble)(scale))
}

func CoreNative_divide_5(src1_nativeObj int64, src2_val0 float64, src2_val1 float64, src2_val2 float64, src2_val3 float64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_divide_15(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jdouble)(src2_val0), (C.jdouble)(src2_val1), (C.jdouble)(src2_val2), (C.jdouble)(src2_val3), (C.jlong)(dst_nativeObj))
}

func CoreNative_divide_6(scale float64, src2_nativeObj int64, dst_nativeObj int64, dtype int) {
	C.Java_org_opencv_core_Core_divide_16(clzEnv, clzObj, (C.jdouble)(scale), (C.jlong)(src2_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(dtype))
}

func CoreNative_divide_7(scale float64, src2_nativeObj int64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_divide_17(clzEnv, clzObj, (C.jdouble)(scale), (C.jlong)(src2_nativeObj), (C.jlong)(dst_nativeObj))
}

func CoreNative_eigenNonSymmetric_0(src_nativeObj int64, eigenvalues_nativeObj int64, eigenvectors_nativeObj int64) {
	C.Java_org_opencv_core_Core_eigenNonSymmetric_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(eigenvalues_nativeObj), (C.jlong)(eigenvectors_nativeObj))
}

func CoreNative_eigen_0(src_nativeObj int64, eigenvalues_nativeObj int64, eigenvectors_nativeObj int64) bool {
	return togobool(C.Java_org_opencv_core_Core_eigen_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(eigenvalues_nativeObj), (C.jlong)(eigenvectors_nativeObj)))
}

func CoreNative_eigen_1(src_nativeObj int64, eigenvalues_nativeObj int64) bool {
	return togobool(C.Java_org_opencv_core_Core_eigen_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(eigenvalues_nativeObj)))
}

func CoreNative_exp_0(src_nativeObj int64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_exp_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj))
}

func CoreNative_extractChannel_0(src_nativeObj int64, dst_nativeObj int64, coi int) {
	C.Java_org_opencv_core_Core_extractChannel_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(coi))
}

func CoreNative_fastAtan2_0(y float32, x float32) float32 {
	return float32(C.Java_org_opencv_core_Core_fastAtan2_10(clzEnv, clzObj, (C.jfloat)(y), (C.jfloat)(x)))
}

func CoreNative_findNonZero_0(src_nativeObj int64, idx_nativeObj int64) {
	C.Java_org_opencv_core_Core_findNonZero_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(idx_nativeObj))
}

func CoreNative_flip_0(src_nativeObj int64, dst_nativeObj int64, flipCode int) {
	C.Java_org_opencv_core_Core_flip_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(flipCode))
}

func CoreNative_gemm_0(src1_nativeObj int64, src2_nativeObj int64, alpha float64, src3_nativeObj int64, beta float64, dst_nativeObj int64, flags int) {
	C.Java_org_opencv_core_Core_gemm_10(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jlong)(src2_nativeObj), (C.jdouble)(alpha), (C.jlong)(src3_nativeObj), (C.jdouble)(beta), (C.jlong)(dst_nativeObj), (C.jint)(flags))
}

func CoreNative_gemm_1(src1_nativeObj int64, src2_nativeObj int64, alpha float64, src3_nativeObj int64, beta float64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_gemm_11(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jlong)(src2_nativeObj), (C.jdouble)(alpha), (C.jlong)(src3_nativeObj), (C.jdouble)(beta), (C.jlong)(dst_nativeObj))
}

func CoreNative_getBuildInformation_0() string {
	return togostring(C.Java_org_opencv_core_Core_getBuildInformation_10(clzEnv, clzObj))
}

func CoreNative_getCPUTickCount_0() int64 {
	return int64(C.Java_org_opencv_core_Core_getCPUTickCount_10(clzEnv, clzObj))
}

func CoreNative_getIppVersion_0() string {
	return togostring(C.Java_org_opencv_core_Core_getIppVersion_10(clzEnv, clzObj))
}

func CoreNative_getNumThreads_0() int {
	return int(C.Java_org_opencv_core_Core_getNumThreads_10(clzEnv, clzObj))
}

func CoreNative_getNumberOfCPUs_0() int {
	return int(C.Java_org_opencv_core_Core_getNumberOfCPUs_10(clzEnv, clzObj))
}

func CoreNative_getOptimalDFTSize_0(vecsize int) int {
	return int(C.Java_org_opencv_core_Core_getOptimalDFTSize_10(clzEnv, clzObj, (C.jint)(vecsize)))
}

func CoreNative_getThreadNum_0() int {
	return int(C.Java_org_opencv_core_Core_getThreadNum_10(clzEnv, clzObj))
}

func CoreNative_getTickCount_0() int64 {
	return int64(C.Java_org_opencv_core_Core_getTickCount_10(clzEnv, clzObj))
}

func CoreNative_getTickFrequency_0() float64 {
	return float64(C.Java_org_opencv_core_Core_getTickFrequency_10(clzEnv, clzObj))
}

func CoreNative_hconcat_0(src_mat_nativeObj int64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_hconcat_10(clzEnv, clzObj, (C.jlong)(src_mat_nativeObj), (C.jlong)(dst_nativeObj))
}

func CoreNative_idct_0(src_nativeObj int64, dst_nativeObj int64, flags int) {
	C.Java_org_opencv_core_Core_idct_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(flags))
}

func CoreNative_idct_1(src_nativeObj int64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_idct_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj))
}

func CoreNative_idft_0(src_nativeObj int64, dst_nativeObj int64, flags int, nonzeroRows int) {
	C.Java_org_opencv_core_Core_idft_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(flags), (C.jint)(nonzeroRows))
}

func CoreNative_idft_1(src_nativeObj int64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_idft_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj))
}

func CoreNative_inRange_0(src_nativeObj int64, lowerb_val0 float64, lowerb_val1 float64, lowerb_val2 float64, lowerb_val3 float64, upperb_val0 float64, upperb_val1 float64, upperb_val2 float64, upperb_val3 float64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_inRange_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jdouble)(lowerb_val0), (C.jdouble)(lowerb_val1), (C.jdouble)(lowerb_val2), (C.jdouble)(lowerb_val3), (C.jdouble)(upperb_val0), (C.jdouble)(upperb_val1), (C.jdouble)(upperb_val2), (C.jdouble)(upperb_val3), (C.jlong)(dst_nativeObj))
}

func CoreNative_insertChannel_0(src_nativeObj int64, dst_nativeObj int64, coi int) {
	C.Java_org_opencv_core_Core_insertChannel_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(coi))
}

func CoreNative_invert_0(src_nativeObj int64, dst_nativeObj int64, flags int) float64 {
	return float64(C.Java_org_opencv_core_Core_invert_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(flags)))
}

func CoreNative_invert_1(src_nativeObj int64, dst_nativeObj int64) float64 {
	return float64(C.Java_org_opencv_core_Core_invert_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj)))
}

func CoreNative_kmeans_0(data_nativeObj int64, K int, bestLabels_nativeObj int64, criteria_type int, criteria_maxCount int, criteria_epsilon float64, attempts int, flags int, centers_nativeObj int64) float64 {
	return float64(C.Java_org_opencv_core_Core_kmeans_10(clzEnv, clzObj, (C.jlong)(data_nativeObj), (C.jint)(K), (C.jlong)(bestLabels_nativeObj), (C.jint)(criteria_type), (C.jint)(criteria_maxCount), (C.jdouble)(criteria_epsilon), (C.jint)(attempts), (C.jint)(flags), (C.jlong)(centers_nativeObj)))
}

func CoreNative_kmeans_1(data_nativeObj int64, K int, bestLabels_nativeObj int64, criteria_type int, criteria_maxCount int, criteria_epsilon float64, attempts int, flags int) float64 {
	return float64(C.Java_org_opencv_core_Core_kmeans_11(clzEnv, clzObj, (C.jlong)(data_nativeObj), (C.jint)(K), (C.jlong)(bestLabels_nativeObj), (C.jint)(criteria_type), (C.jint)(criteria_maxCount), (C.jdouble)(criteria_epsilon), (C.jint)(attempts), (C.jint)(flags)))
}

func CoreNative_log_0(src_nativeObj int64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_log_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj))
}

func CoreNative_magnitude_0(x_nativeObj int64, y_nativeObj int64, magnitude_nativeObj int64) {
	C.Java_org_opencv_core_Core_magnitude_10(clzEnv, clzObj, (C.jlong)(x_nativeObj), (C.jlong)(y_nativeObj), (C.jlong)(magnitude_nativeObj))
}

func CoreNative_max_0(src1_nativeObj int64, src2_nativeObj int64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_max_10(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jlong)(src2_nativeObj), (C.jlong)(dst_nativeObj))
}

func CoreNative_max_1(src1_nativeObj int64, src2_val0 float64, src2_val1 float64, src2_val2 float64, src2_val3 float64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_max_11(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jdouble)(src2_val0), (C.jdouble)(src2_val1), (C.jdouble)(src2_val2), (C.jdouble)(src2_val3), (C.jlong)(dst_nativeObj))
}

func CoreNative_meanStdDev_0(src_nativeObj int64, mean_mat_nativeObj int64, stddev_mat_nativeObj int64, mask_nativeObj int64) {
	C.Java_org_opencv_core_Core_meanStdDev_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(mean_mat_nativeObj), (C.jlong)(stddev_mat_nativeObj), (C.jlong)(mask_nativeObj))
}

func CoreNative_meanStdDev_1(src_nativeObj int64, mean_mat_nativeObj int64, stddev_mat_nativeObj int64) {
	C.Java_org_opencv_core_Core_meanStdDev_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(mean_mat_nativeObj), (C.jlong)(stddev_mat_nativeObj))
}

func CoreNative_mean_0(src_nativeObj int64, mask_nativeObj int64) []float64 {
	return togofloat64Array(C.Java_org_opencv_core_Core_mean_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(mask_nativeObj)))
}

func CoreNative_mean_1(src_nativeObj int64) []float64 {
	return togofloat64Array(C.Java_org_opencv_core_Core_mean_11(clzEnv, clzObj, (C.jlong)(src_nativeObj)))
}

func CoreNative_merge_0(mv_mat_nativeObj int64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_merge_10(clzEnv, clzObj, (C.jlong)(mv_mat_nativeObj), (C.jlong)(dst_nativeObj))
}

func CoreNative_min_0(src1_nativeObj int64, src2_nativeObj int64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_min_10(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jlong)(src2_nativeObj), (C.jlong)(dst_nativeObj))
}

func CoreNative_min_1(src1_nativeObj int64, src2_val0 float64, src2_val1 float64, src2_val2 float64, src2_val3 float64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_min_11(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jdouble)(src2_val0), (C.jdouble)(src2_val1), (C.jdouble)(src2_val2), (C.jdouble)(src2_val3), (C.jlong)(dst_nativeObj))
}

func CoreNative_mixChannels_0(src_mat_nativeObj int64, dst_mat_nativeObj int64, fromTo_mat_nativeObj int64) {
	C.Java_org_opencv_core_Core_mixChannels_10(clzEnv, clzObj, (C.jlong)(src_mat_nativeObj), (C.jlong)(dst_mat_nativeObj), (C.jlong)(fromTo_mat_nativeObj))
}

func CoreNative_mulSpectrums_0(a_nativeObj int64, b_nativeObj int64, c_nativeObj int64, flags int, conjB bool) {
	C.Java_org_opencv_core_Core_mulSpectrums_10(clzEnv, clzObj, (C.jlong)(a_nativeObj), (C.jlong)(b_nativeObj), (C.jlong)(c_nativeObj), (C.jint)(flags), tojboolean(conjB))
}

func CoreNative_mulSpectrums_1(a_nativeObj int64, b_nativeObj int64, c_nativeObj int64, flags int) {
	C.Java_org_opencv_core_Core_mulSpectrums_11(clzEnv, clzObj, (C.jlong)(a_nativeObj), (C.jlong)(b_nativeObj), (C.jlong)(c_nativeObj), (C.jint)(flags))
}

func CoreNative_mulTransposed_0(src_nativeObj int64, dst_nativeObj int64, aTa bool, delta_nativeObj int64, scale float64, dtype int) {
	C.Java_org_opencv_core_Core_mulTransposed_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), tojboolean(aTa), (C.jlong)(delta_nativeObj), (C.jdouble)(scale), (C.jint)(dtype))
}

func CoreNative_mulTransposed_1(src_nativeObj int64, dst_nativeObj int64, aTa bool, delta_nativeObj int64, scale float64) {
	C.Java_org_opencv_core_Core_mulTransposed_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), tojboolean(aTa), (C.jlong)(delta_nativeObj), (C.jdouble)(scale))
}

func CoreNative_mulTransposed_2(src_nativeObj int64, dst_nativeObj int64, aTa bool) {
	C.Java_org_opencv_core_Core_mulTransposed_12(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), tojboolean(aTa))
}

func CoreNative_multiply_0(src1_nativeObj int64, src2_nativeObj int64, dst_nativeObj int64, scale float64, dtype int) {
	C.Java_org_opencv_core_Core_multiply_10(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jlong)(src2_nativeObj), (C.jlong)(dst_nativeObj), (C.jdouble)(scale), (C.jint)(dtype))
}

func CoreNative_multiply_1(src1_nativeObj int64, src2_nativeObj int64, dst_nativeObj int64, scale float64) {
	C.Java_org_opencv_core_Core_multiply_11(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jlong)(src2_nativeObj), (C.jlong)(dst_nativeObj), (C.jdouble)(scale))
}

func CoreNative_multiply_2(src1_nativeObj int64, src2_nativeObj int64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_multiply_12(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jlong)(src2_nativeObj), (C.jlong)(dst_nativeObj))
}

func CoreNative_multiply_3(src1_nativeObj int64, src2_val0 float64, src2_val1 float64, src2_val2 float64, src2_val3 float64, dst_nativeObj int64, scale float64, dtype int) {
	C.Java_org_opencv_core_Core_multiply_13(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jdouble)(src2_val0), (C.jdouble)(src2_val1), (C.jdouble)(src2_val2), (C.jdouble)(src2_val3), (C.jlong)(dst_nativeObj), (C.jdouble)(scale), (C.jint)(dtype))
}

func CoreNative_multiply_4(src1_nativeObj int64, src2_val0 float64, src2_val1 float64, src2_val2 float64, src2_val3 float64, dst_nativeObj int64, scale float64) {
	C.Java_org_opencv_core_Core_multiply_14(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jdouble)(src2_val0), (C.jdouble)(src2_val1), (C.jdouble)(src2_val2), (C.jdouble)(src2_val3), (C.jlong)(dst_nativeObj), (C.jdouble)(scale))
}

func CoreNative_multiply_5(src1_nativeObj int64, src2_val0 float64, src2_val1 float64, src2_val2 float64, src2_val3 float64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_multiply_15(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jdouble)(src2_val0), (C.jdouble)(src2_val1), (C.jdouble)(src2_val2), (C.jdouble)(src2_val3), (C.jlong)(dst_nativeObj))
}

func CoreNative_n_minMaxLocManual(src_nativeObj int64, mask_nativeObj int64) []float64 {
	return togofloat64Array(C.Java_org_opencv_core_Core_n_1minMaxLocManual(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(mask_nativeObj)))
}

func CoreNative_norm_0(src1_nativeObj int64, src2_nativeObj int64, normType int, mask_nativeObj int64) float64 {
	return float64(C.Java_org_opencv_core_Core_norm_10(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jlong)(src2_nativeObj), (C.jint)(normType), (C.jlong)(mask_nativeObj)))
}

func CoreNative_norm_1(src1_nativeObj int64, src2_nativeObj int64, normType int) float64 {
	return float64(C.Java_org_opencv_core_Core_norm_11(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jlong)(src2_nativeObj), (C.jint)(normType)))
}

func CoreNative_norm_2(src1_nativeObj int64, src2_nativeObj int64) float64 {
	return float64(C.Java_org_opencv_core_Core_norm_12(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jlong)(src2_nativeObj)))
}

func CoreNative_norm_3(src1_nativeObj int64, normType int, mask_nativeObj int64) float64 {
	return float64(C.Java_org_opencv_core_Core_norm_13(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jint)(normType), (C.jlong)(mask_nativeObj)))
}

func CoreNative_norm_4(src1_nativeObj int64, normType int) float64 {
	return float64(C.Java_org_opencv_core_Core_norm_14(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jint)(normType)))
}

func CoreNative_norm_5(src1_nativeObj int64) float64 {
	return float64(C.Java_org_opencv_core_Core_norm_15(clzEnv, clzObj, (C.jlong)(src1_nativeObj)))
}

func CoreNative_normalize_0(src_nativeObj int64, dst_nativeObj int64, alpha float64, beta float64, norm_type int, dtype int, mask_nativeObj int64) {
	C.Java_org_opencv_core_Core_normalize_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jdouble)(alpha), (C.jdouble)(beta), (C.jint)(norm_type), (C.jint)(dtype), (C.jlong)(mask_nativeObj))
}

func CoreNative_normalize_1(src_nativeObj int64, dst_nativeObj int64, alpha float64, beta float64, norm_type int, dtype int) {
	C.Java_org_opencv_core_Core_normalize_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jdouble)(alpha), (C.jdouble)(beta), (C.jint)(norm_type), (C.jint)(dtype))
}

func CoreNative_normalize_2(src_nativeObj int64, dst_nativeObj int64, alpha float64, beta float64, norm_type int) {
	C.Java_org_opencv_core_Core_normalize_12(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jdouble)(alpha), (C.jdouble)(beta), (C.jint)(norm_type))
}

func CoreNative_normalize_3(src_nativeObj int64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_normalize_13(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj))
}

func CoreNative_patchNaNs_0(a_nativeObj int64, val float64) {
	C.Java_org_opencv_core_Core_patchNaNs_10(clzEnv, clzObj, (C.jlong)(a_nativeObj), (C.jdouble)(val))
}

func CoreNative_patchNaNs_1(a_nativeObj int64) {
	C.Java_org_opencv_core_Core_patchNaNs_11(clzEnv, clzObj, (C.jlong)(a_nativeObj))
}

func CoreNative_perspectiveTransform_0(src_nativeObj int64, dst_nativeObj int64, m_nativeObj int64) {
	C.Java_org_opencv_core_Core_perspectiveTransform_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jlong)(m_nativeObj))
}

func CoreNative_phase_0(x_nativeObj int64, y_nativeObj int64, angle_nativeObj int64, angleInDegrees bool) {
	C.Java_org_opencv_core_Core_phase_10(clzEnv, clzObj, (C.jlong)(x_nativeObj), (C.jlong)(y_nativeObj), (C.jlong)(angle_nativeObj), tojboolean(angleInDegrees))
}

func CoreNative_phase_1(x_nativeObj int64, y_nativeObj int64, angle_nativeObj int64) {
	C.Java_org_opencv_core_Core_phase_11(clzEnv, clzObj, (C.jlong)(x_nativeObj), (C.jlong)(y_nativeObj), (C.jlong)(angle_nativeObj))
}

func CoreNative_polarToCart_0(magnitude_nativeObj int64, angle_nativeObj int64, x_nativeObj int64, y_nativeObj int64, angleInDegrees bool) {
	C.Java_org_opencv_core_Core_polarToCart_10(clzEnv, clzObj, (C.jlong)(magnitude_nativeObj), (C.jlong)(angle_nativeObj), (C.jlong)(x_nativeObj), (C.jlong)(y_nativeObj), tojboolean(angleInDegrees))
}

func CoreNative_polarToCart_1(magnitude_nativeObj int64, angle_nativeObj int64, x_nativeObj int64, y_nativeObj int64) {
	C.Java_org_opencv_core_Core_polarToCart_11(clzEnv, clzObj, (C.jlong)(magnitude_nativeObj), (C.jlong)(angle_nativeObj), (C.jlong)(x_nativeObj), (C.jlong)(y_nativeObj))
}

func CoreNative_pow_0(src_nativeObj int64, power float64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_pow_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jdouble)(power), (C.jlong)(dst_nativeObj))
}

func CoreNative_randShuffle_0(dst_nativeObj int64, iterFactor float64) {
	C.Java_org_opencv_core_Core_randShuffle_10(clzEnv, clzObj, (C.jlong)(dst_nativeObj), (C.jdouble)(iterFactor))
}

func CoreNative_randShuffle_1(dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_randShuffle_11(clzEnv, clzObj, (C.jlong)(dst_nativeObj))
}

func CoreNative_randn_0(dst_nativeObj int64, mean float64, stddev float64) {
	C.Java_org_opencv_core_Core_randn_10(clzEnv, clzObj, (C.jlong)(dst_nativeObj), (C.jdouble)(mean), (C.jdouble)(stddev))
}

func CoreNative_randu_0(dst_nativeObj int64, low float64, high float64) {
	C.Java_org_opencv_core_Core_randu_10(clzEnv, clzObj, (C.jlong)(dst_nativeObj), (C.jdouble)(low), (C.jdouble)(high))
}

func CoreNative_reduce_0(src_nativeObj int64, dst_nativeObj int64, dim int, rtype int, dtype int) {
	C.Java_org_opencv_core_Core_reduce_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(dim), (C.jint)(rtype), (C.jint)(dtype))
}

func CoreNative_reduce_1(src_nativeObj int64, dst_nativeObj int64, dim int, rtype int) {
	C.Java_org_opencv_core_Core_reduce_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(dim), (C.jint)(rtype))
}

func CoreNative_repeat_0(src_nativeObj int64, ny int, nx int, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_repeat_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jint)(ny), (C.jint)(nx), (C.jlong)(dst_nativeObj))
}

func CoreNative_rotate_0(src_nativeObj int64, dst_nativeObj int64, rotateCode int) {
	C.Java_org_opencv_core_Core_rotate_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(rotateCode))
}

func CoreNative_scaleAdd_0(src1_nativeObj int64, alpha float64, src2_nativeObj int64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_scaleAdd_10(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jdouble)(alpha), (C.jlong)(src2_nativeObj), (C.jlong)(dst_nativeObj))
}

func CoreNative_setErrorVerbosity_0(verbose bool) {
	C.Java_org_opencv_core_Core_setErrorVerbosity_10(clzEnv, clzObj, tojboolean(verbose))
}

func CoreNative_setIdentity_0(mtx_nativeObj int64, s_val0 float64, s_val1 float64, s_val2 float64, s_val3 float64) {
	C.Java_org_opencv_core_Core_setIdentity_10(clzEnv, clzObj, (C.jlong)(mtx_nativeObj), (C.jdouble)(s_val0), (C.jdouble)(s_val1), (C.jdouble)(s_val2), (C.jdouble)(s_val3))
}

func CoreNative_setIdentity_1(mtx_nativeObj int64) {
	C.Java_org_opencv_core_Core_setIdentity_11(clzEnv, clzObj, (C.jlong)(mtx_nativeObj))
}

func CoreNative_setNumThreads_0(nthreads int) {
	C.Java_org_opencv_core_Core_setNumThreads_10(clzEnv, clzObj, (C.jint)(nthreads))
}

func CoreNative_setRNGSeed_0(seed int) {
	C.Java_org_opencv_core_Core_setRNGSeed_10(clzEnv, clzObj, (C.jint)(seed))
}

func CoreNative_setUseIPP_0(flag bool) {
	C.Java_org_opencv_core_Core_setUseIPP_10(clzEnv, clzObj, tojboolean(flag))
}

func CoreNative_setUseIPP_NE_0(flag bool) {
	C.Java_org_opencv_core_Core_setUseIPP_1NE_10(clzEnv, clzObj, tojboolean(flag))
}

func CoreNative_solveCubic_0(coeffs_nativeObj int64, roots_nativeObj int64) int {
	return int(C.Java_org_opencv_core_Core_solveCubic_10(clzEnv, clzObj, (C.jlong)(coeffs_nativeObj), (C.jlong)(roots_nativeObj)))
}

func CoreNative_solvePoly_0(coeffs_nativeObj int64, roots_nativeObj int64, maxIters int) float64 {
	return float64(C.Java_org_opencv_core_Core_solvePoly_10(clzEnv, clzObj, (C.jlong)(coeffs_nativeObj), (C.jlong)(roots_nativeObj), (C.jint)(maxIters)))
}

func CoreNative_solvePoly_1(coeffs_nativeObj int64, roots_nativeObj int64) float64 {
	return float64(C.Java_org_opencv_core_Core_solvePoly_11(clzEnv, clzObj, (C.jlong)(coeffs_nativeObj), (C.jlong)(roots_nativeObj)))
}

func CoreNative_solve_0(src1_nativeObj int64, src2_nativeObj int64, dst_nativeObj int64, flags int) bool {
	return togobool(C.Java_org_opencv_core_Core_solve_10(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jlong)(src2_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(flags)))
}

func CoreNative_solve_1(src1_nativeObj int64, src2_nativeObj int64, dst_nativeObj int64) bool {
	return togobool(C.Java_org_opencv_core_Core_solve_11(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jlong)(src2_nativeObj), (C.jlong)(dst_nativeObj)))
}

func CoreNative_sortIdx_0(src_nativeObj int64, dst_nativeObj int64, flags int) {
	C.Java_org_opencv_core_Core_sortIdx_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(flags))
}

func CoreNative_sort_0(src_nativeObj int64, dst_nativeObj int64, flags int) {
	C.Java_org_opencv_core_Core_sort_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(flags))
}

func CoreNative_split_0(m_nativeObj int64, mv_mat_nativeObj int64) {
	C.Java_org_opencv_core_Core_split_10(clzEnv, clzObj, (C.jlong)(m_nativeObj), (C.jlong)(mv_mat_nativeObj))
}

func CoreNative_sqrt_0(src_nativeObj int64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_sqrt_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj))
}

func CoreNative_subtract_0(src1_nativeObj int64, src2_nativeObj int64, dst_nativeObj int64, mask_nativeObj int64, dtype int) {
	C.Java_org_opencv_core_Core_subtract_10(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jlong)(src2_nativeObj), (C.jlong)(dst_nativeObj), (C.jlong)(mask_nativeObj), (C.jint)(dtype))
}

func CoreNative_subtract_1(src1_nativeObj int64, src2_nativeObj int64, dst_nativeObj int64, mask_nativeObj int64) {
	C.Java_org_opencv_core_Core_subtract_11(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jlong)(src2_nativeObj), (C.jlong)(dst_nativeObj), (C.jlong)(mask_nativeObj))
}

func CoreNative_subtract_2(src1_nativeObj int64, src2_nativeObj int64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_subtract_12(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jlong)(src2_nativeObj), (C.jlong)(dst_nativeObj))
}

func CoreNative_subtract_3(src1_nativeObj int64, src2_val0 float64, src2_val1 float64, src2_val2 float64, src2_val3 float64, dst_nativeObj int64, mask_nativeObj int64, dtype int) {
	C.Java_org_opencv_core_Core_subtract_13(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jdouble)(src2_val0), (C.jdouble)(src2_val1), (C.jdouble)(src2_val2), (C.jdouble)(src2_val3), (C.jlong)(dst_nativeObj), (C.jlong)(mask_nativeObj), (C.jint)(dtype))
}

func CoreNative_subtract_4(src1_nativeObj int64, src2_val0 float64, src2_val1 float64, src2_val2 float64, src2_val3 float64, dst_nativeObj int64, mask_nativeObj int64) {
	C.Java_org_opencv_core_Core_subtract_14(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jdouble)(src2_val0), (C.jdouble)(src2_val1), (C.jdouble)(src2_val2), (C.jdouble)(src2_val3), (C.jlong)(dst_nativeObj), (C.jlong)(mask_nativeObj))
}

func CoreNative_subtract_5(src1_nativeObj int64, src2_val0 float64, src2_val1 float64, src2_val2 float64, src2_val3 float64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_subtract_15(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jdouble)(src2_val0), (C.jdouble)(src2_val1), (C.jdouble)(src2_val2), (C.jdouble)(src2_val3), (C.jlong)(dst_nativeObj))
}

func CoreNative_sumElems_0(src_nativeObj int64) []float64 {
	return togofloat64Array(C.Java_org_opencv_core_Core_sumElems_10(clzEnv, clzObj, (C.jlong)(src_nativeObj)))
}

func CoreNative_trace_0(mtx_nativeObj int64) []float64 {
	return togofloat64Array(C.Java_org_opencv_core_Core_trace_10(clzEnv, clzObj, (C.jlong)(mtx_nativeObj)))
}

func CoreNative_transform_0(src_nativeObj int64, dst_nativeObj int64, m_nativeObj int64) {
	C.Java_org_opencv_core_Core_transform_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jlong)(m_nativeObj))
}

func CoreNative_transpose_0(src_nativeObj int64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_transpose_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj))
}

func CoreNative_useIPP_0() bool {
	return togobool(C.Java_org_opencv_core_Core_useIPP_10(clzEnv, clzObj))
}

func CoreNative_useIPP_NE_0() bool {
	return togobool(C.Java_org_opencv_core_Core_useIPP_1NE_10(clzEnv, clzObj))
}

func CoreNative_vconcat_0(src_mat_nativeObj int64, dst_nativeObj int64) {
	C.Java_org_opencv_core_Core_vconcat_10(clzEnv, clzObj, (C.jlong)(src_mat_nativeObj), (C.jlong)(dst_nativeObj))
}

func AlgorithmNative_clear_0(nativeObj int64) {
	C.Java_org_opencv_core_Algorithm_clear_10(clzEnv, clzObj, (C.jlong)(nativeObj))
}

func AlgorithmNative_delete(nativeObj int64) {
	C.Java_org_opencv_core_Algorithm_delete(clzEnv, clzObj, (C.jlong)(nativeObj))
}

func AlgorithmNative_getDefaultName_0(nativeObj int64) string {
	return togostring(C.Java_org_opencv_core_Algorithm_getDefaultName_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func AlgorithmNative_save_0(nativeObj int64, filename string) {
	defer ungointerface(filename)
	C.Java_org_opencv_core_Algorithm_save_10(clzEnv, clzObj, (C.jlong)(nativeObj), tojstring(filename))
}

func TickMeterNative_TickMeter_0() int64 {
	return int64(C.Java_org_opencv_core_TickMeter_TickMeter_10(clzEnv, clzObj))
}

func TickMeterNative_delete(nativeObj int64) {
	C.Java_org_opencv_core_TickMeter_delete(clzEnv, clzObj, (C.jlong)(nativeObj))
}

func TickMeterNative_getCounter_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_core_TickMeter_getCounter_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func TickMeterNative_getTimeMicro_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_core_TickMeter_getTimeMicro_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func TickMeterNative_getTimeMilli_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_core_TickMeter_getTimeMilli_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func TickMeterNative_getTimeSec_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_core_TickMeter_getTimeSec_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func TickMeterNative_getTimeTicks_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_core_TickMeter_getTimeTicks_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func TickMeterNative_reset_0(nativeObj int64) {
	C.Java_org_opencv_core_TickMeter_reset_10(clzEnv, clzObj, (C.jlong)(nativeObj))
}

func TickMeterNative_start_0(nativeObj int64) {
	C.Java_org_opencv_core_TickMeter_start_10(clzEnv, clzObj, (C.jlong)(nativeObj))
}

func TickMeterNative_stop_0(nativeObj int64) {
	C.Java_org_opencv_core_TickMeter_stop_10(clzEnv, clzObj, (C.jlong)(nativeObj))
}
