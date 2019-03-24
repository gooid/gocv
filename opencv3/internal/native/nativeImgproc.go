package core

/*

#include "jni.h"

extern void Java_org_opencv_imgproc_CLAHE_apply_10(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_imgproc_CLAHE_collectGarbage_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_imgproc_CLAHE_delete(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_imgproc_CLAHE_getClipLimit_10(JNIEnv*, jclass, jlong);
extern jdoubleArray Java_org_opencv_imgproc_CLAHE_getTilesGridSize_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_imgproc_CLAHE_setClipLimit_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_imgproc_CLAHE_setTilesGridSize_10(JNIEnv*, jclass, jlong, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_Canny_10(JNIEnv*, jclass, jlong, jlong, jlong, jdouble, jdouble, jboolean);
extern void Java_org_opencv_imgproc_Imgproc_Canny_11(JNIEnv*, jclass, jlong, jlong, jlong, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_Canny_12(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jint, jboolean);
extern void Java_org_opencv_imgproc_Imgproc_Canny_13(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble);
extern jfloat Java_org_opencv_imgproc_Imgproc_EMD_10(JNIEnv*, jclass, jlong, jlong, jint, jlong, jlong);
extern jfloat Java_org_opencv_imgproc_Imgproc_EMD_11(JNIEnv*, jclass, jlong, jlong, jint);
extern void Java_org_opencv_imgproc_Imgproc_GaussianBlur_10(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jdouble, jdouble, jint);
extern void Java_org_opencv_imgproc_Imgproc_GaussianBlur_11(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_GaussianBlur_12(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_HoughCircles_10(JNIEnv*, jclass, jlong, jlong, jint, jdouble, jdouble, jdouble, jdouble, jint, jint);
extern void Java_org_opencv_imgproc_Imgproc_HoughCircles_11(JNIEnv*, jclass, jlong, jlong, jint, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_HoughLinesP_10(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jint, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_HoughLinesP_11(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jint);
extern void Java_org_opencv_imgproc_Imgproc_HoughLines_10(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jint, jdouble, jdouble, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_HoughLines_11(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jint);
extern void Java_org_opencv_imgproc_Imgproc_HuMoments_10(JNIEnv*, jclass, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jlong);
extern void Java_org_opencv_imgproc_Imgproc_Laplacian_10(JNIEnv*, jclass, jlong, jlong, jint, jint, jdouble, jdouble, jint);
extern void Java_org_opencv_imgproc_Imgproc_Laplacian_11(JNIEnv*, jclass, jlong, jlong, jint, jint, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_Laplacian_12(JNIEnv*, jclass, jlong, jlong, jint);
extern void Java_org_opencv_imgproc_Imgproc_Scharr_10(JNIEnv*, jclass, jlong, jlong, jint, jint, jint, jdouble, jdouble, jint);
extern void Java_org_opencv_imgproc_Imgproc_Scharr_11(JNIEnv*, jclass, jlong, jlong, jint, jint, jint, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_Scharr_12(JNIEnv*, jclass, jlong, jlong, jint, jint, jint);
extern void Java_org_opencv_imgproc_Imgproc_Sobel_10(JNIEnv*, jclass, jlong, jlong, jint, jint, jint, jint, jdouble, jdouble, jint);
extern void Java_org_opencv_imgproc_Imgproc_Sobel_11(JNIEnv*, jclass, jlong, jlong, jint, jint, jint, jint, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_Sobel_12(JNIEnv*, jclass, jlong, jlong, jint, jint, jint);
extern void Java_org_opencv_imgproc_Imgproc_accumulateProduct_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_imgproc_Imgproc_accumulateProduct_11(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_imgproc_Imgproc_accumulateSquare_10(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_imgproc_Imgproc_accumulateSquare_11(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_imgproc_Imgproc_accumulateWeighted_10(JNIEnv*, jclass, jlong, jlong, jdouble, jlong);
extern void Java_org_opencv_imgproc_Imgproc_accumulateWeighted_11(JNIEnv*, jclass, jlong, jlong, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_accumulate_10(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_imgproc_Imgproc_accumulate_11(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_imgproc_Imgproc_adaptiveThreshold_10(JNIEnv*, jclass, jlong, jlong, jdouble, jint, jint, jint, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_applyColorMap_10(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_imgproc_Imgproc_applyColorMap_11(JNIEnv*, jclass, jlong, jlong, jint);
extern void Java_org_opencv_imgproc_Imgproc_approxPolyDP_10(JNIEnv*, jclass, jlong, jlong, jdouble, jboolean);
extern jdouble Java_org_opencv_imgproc_Imgproc_arcLength_10(JNIEnv*, jclass, jlong, jboolean);
extern void Java_org_opencv_imgproc_Imgproc_arrowedLine_10(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jint, jint, jint, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_arrowedLine_11(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_bilateralFilter_10(JNIEnv*, jclass, jlong, jlong, jint, jdouble, jdouble, jint);
extern void Java_org_opencv_imgproc_Imgproc_bilateralFilter_11(JNIEnv*, jclass, jlong, jlong, jint, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_blur_10(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jdouble, jdouble, jint);
extern void Java_org_opencv_imgproc_Imgproc_blur_11(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_blur_12(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble);
extern jdoubleArray Java_org_opencv_imgproc_Imgproc_boundingRect_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_imgproc_Imgproc_boxFilter_10(JNIEnv*, jclass, jlong, jlong, jint, jdouble, jdouble, jdouble, jdouble, jboolean, jint);
extern void Java_org_opencv_imgproc_Imgproc_boxFilter_11(JNIEnv*, jclass, jlong, jlong, jint, jdouble, jdouble, jdouble, jdouble, jboolean);
extern void Java_org_opencv_imgproc_Imgproc_boxFilter_12(JNIEnv*, jclass, jlong, jlong, jint, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_boxPoints_10(JNIEnv*, jclass, jdouble, jdouble, jdouble, jdouble, jdouble, jlong);
extern void Java_org_opencv_imgproc_Imgproc_calcBackProject_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_calcHist_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong, jboolean);
extern void Java_org_opencv_imgproc_Imgproc_calcHist_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_imgproc_Imgproc_circle_10(JNIEnv*, jclass, jlong, jdouble, jdouble, jint, jdouble, jdouble, jdouble, jdouble, jint, jint, jint);
extern void Java_org_opencv_imgproc_Imgproc_circle_11(JNIEnv*, jclass, jlong, jdouble, jdouble, jint, jdouble, jdouble, jdouble, jdouble, jint);
extern void Java_org_opencv_imgproc_Imgproc_circle_12(JNIEnv*, jclass, jlong, jdouble, jdouble, jint, jdouble, jdouble, jdouble, jdouble);
extern jboolean Java_org_opencv_imgproc_Imgproc_clipLine_10(JNIEnv*, jclass, jint, jint, jint, jint, jdouble, jdouble, jdoubleArray, jdouble, jdouble, jdoubleArray);
extern jdouble Java_org_opencv_imgproc_Imgproc_compareHist_10(JNIEnv*, jclass, jlong, jlong, jint);
extern jint Java_org_opencv_imgproc_Imgproc_connectedComponentsWithAlgorithm_10(JNIEnv*, jclass, jlong, jlong, jint, jint, jint);
extern jint Java_org_opencv_imgproc_Imgproc_connectedComponentsWithStatsWithAlgorithm_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jint, jint, jint);
extern jint Java_org_opencv_imgproc_Imgproc_connectedComponentsWithStats_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jint, jint);
extern jint Java_org_opencv_imgproc_Imgproc_connectedComponentsWithStats_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern jint Java_org_opencv_imgproc_Imgproc_connectedComponents_10(JNIEnv*, jclass, jlong, jlong, jint, jint);
extern jint Java_org_opencv_imgproc_Imgproc_connectedComponents_11(JNIEnv*, jclass, jlong, jlong);
extern jdouble Java_org_opencv_imgproc_Imgproc_contourArea_10(JNIEnv*, jclass, jlong, jboolean);
extern jdouble Java_org_opencv_imgproc_Imgproc_contourArea_11(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_imgproc_Imgproc_convertMaps_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jint, jboolean);
extern void Java_org_opencv_imgproc_Imgproc_convertMaps_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jint);
extern void Java_org_opencv_imgproc_Imgproc_convexHull_10(JNIEnv*, jclass, jlong, jlong, jboolean);
extern void Java_org_opencv_imgproc_Imgproc_convexHull_11(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_imgproc_Imgproc_convexityDefects_10(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_imgproc_Imgproc_cornerEigenValsAndVecs_10(JNIEnv*, jclass, jlong, jlong, jint, jint, jint);
extern void Java_org_opencv_imgproc_Imgproc_cornerEigenValsAndVecs_11(JNIEnv*, jclass, jlong, jlong, jint, jint);
extern void Java_org_opencv_imgproc_Imgproc_cornerHarris_10(JNIEnv*, jclass, jlong, jlong, jint, jint, jdouble, jint);
extern void Java_org_opencv_imgproc_Imgproc_cornerHarris_11(JNIEnv*, jclass, jlong, jlong, jint, jint, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_cornerMinEigenVal_10(JNIEnv*, jclass, jlong, jlong, jint, jint, jint);
extern void Java_org_opencv_imgproc_Imgproc_cornerMinEigenVal_11(JNIEnv*, jclass, jlong, jlong, jint, jint);
extern void Java_org_opencv_imgproc_Imgproc_cornerMinEigenVal_12(JNIEnv*, jclass, jlong, jlong, jint);
extern void Java_org_opencv_imgproc_Imgproc_cornerSubPix_10(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jdouble, jdouble, jint, jint, jdouble);
extern jlong Java_org_opencv_imgproc_Imgproc_createCLAHE_10(JNIEnv*, jclass, jdouble, jdouble, jdouble);
extern jlong Java_org_opencv_imgproc_Imgproc_createCLAHE_11(JNIEnv*, jclass);
extern void Java_org_opencv_imgproc_Imgproc_createHanningWindow_10(JNIEnv*, jclass, jlong, jdouble, jdouble, jint);
extern jlong Java_org_opencv_imgproc_Imgproc_createLineSegmentDetector_10(JNIEnv*, jclass, jint, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jint);
extern jlong Java_org_opencv_imgproc_Imgproc_createLineSegmentDetector_11(JNIEnv*, jclass);
extern void Java_org_opencv_imgproc_Imgproc_cvtColor_10(JNIEnv*, jclass, jlong, jlong, jint, jint);
extern void Java_org_opencv_imgproc_Imgproc_cvtColor_11(JNIEnv*, jclass, jlong, jlong, jint);
extern void Java_org_opencv_imgproc_Imgproc_demosaicing_10(JNIEnv*, jclass, jlong, jlong, jint, jint);
extern void Java_org_opencv_imgproc_Imgproc_demosaicing_11(JNIEnv*, jclass, jlong, jlong, jint);
extern void Java_org_opencv_imgproc_Imgproc_dilate_10(JNIEnv*, jclass, jlong, jlong, jlong, jdouble, jdouble, jint, jint, jdouble, jdouble, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_dilate_11(JNIEnv*, jclass, jlong, jlong, jlong, jdouble, jdouble, jint);
extern void Java_org_opencv_imgproc_Imgproc_dilate_12(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_imgproc_Imgproc_distanceTransformWithLabels_10(JNIEnv*, jclass, jlong, jlong, jlong, jint, jint, jint);
extern void Java_org_opencv_imgproc_Imgproc_distanceTransformWithLabels_11(JNIEnv*, jclass, jlong, jlong, jlong, jint, jint);
extern void Java_org_opencv_imgproc_Imgproc_distanceTransform_10(JNIEnv*, jclass, jlong, jlong, jint, jint, jint);
extern void Java_org_opencv_imgproc_Imgproc_distanceTransform_11(JNIEnv*, jclass, jlong, jlong, jint, jint);
extern void Java_org_opencv_imgproc_Imgproc_drawContours_10(JNIEnv*, jclass, jlong, jlong, jint, jdouble, jdouble, jdouble, jdouble, jint, jint, jlong, jint, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_drawContours_11(JNIEnv*, jclass, jlong, jlong, jint, jdouble, jdouble, jdouble, jdouble, jint);
extern void Java_org_opencv_imgproc_Imgproc_drawContours_12(JNIEnv*, jclass, jlong, jlong, jint, jdouble, jdouble, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_drawMarker_10(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jint, jint, jint, jint);
extern void Java_org_opencv_imgproc_Imgproc_drawMarker_11(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_ellipse2Poly_10(JNIEnv*, jclass, jdouble, jdouble, jdouble, jdouble, jint, jint, jint, jint, jlong);
extern void Java_org_opencv_imgproc_Imgproc_ellipse_10(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jint, jint, jint);
extern void Java_org_opencv_imgproc_Imgproc_ellipse_11(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jint);
extern void Java_org_opencv_imgproc_Imgproc_ellipse_12(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_ellipse_13(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jint, jint);
extern void Java_org_opencv_imgproc_Imgproc_ellipse_14(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jint);
extern void Java_org_opencv_imgproc_Imgproc_ellipse_15(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_equalizeHist_10(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_imgproc_Imgproc_erode_10(JNIEnv*, jclass, jlong, jlong, jlong, jdouble, jdouble, jint, jint, jdouble, jdouble, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_erode_11(JNIEnv*, jclass, jlong, jlong, jlong, jdouble, jdouble, jint);
extern void Java_org_opencv_imgproc_Imgproc_erode_12(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_imgproc_Imgproc_fillConvexPoly_10(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jdouble, jdouble, jint, jint);
extern void Java_org_opencv_imgproc_Imgproc_fillConvexPoly_11(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_fillPoly_10(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jdouble, jdouble, jint, jint, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_fillPoly_11(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_filter2D_10(JNIEnv*, jclass, jlong, jlong, jint, jlong, jdouble, jdouble, jdouble, jint);
extern void Java_org_opencv_imgproc_Imgproc_filter2D_11(JNIEnv*, jclass, jlong, jlong, jint, jlong, jdouble, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_filter2D_12(JNIEnv*, jclass, jlong, jlong, jint, jlong);
extern void Java_org_opencv_imgproc_Imgproc_findContours_10(JNIEnv*, jclass, jlong, jlong, jlong, jint, jint, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_findContours_11(JNIEnv*, jclass, jlong, jlong, jlong, jint, jint);
extern jdoubleArray Java_org_opencv_imgproc_Imgproc_fitEllipseAMS_10(JNIEnv*, jclass, jlong);
extern jdoubleArray Java_org_opencv_imgproc_Imgproc_fitEllipseDirect_10(JNIEnv*, jclass, jlong);
extern jdoubleArray Java_org_opencv_imgproc_Imgproc_fitEllipse_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_imgproc_Imgproc_fitLine_10(JNIEnv*, jclass, jlong, jlong, jint, jdouble, jdouble, jdouble);
extern jint Java_org_opencv_imgproc_Imgproc_floodFill_10(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdoubleArray, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jint);
extern jint Java_org_opencv_imgproc_Imgproc_floodFill_11(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble);
extern jlong Java_org_opencv_imgproc_Imgproc_getAffineTransform_10(JNIEnv*, jclass, jlong, jlong);
extern jlong Java_org_opencv_imgproc_Imgproc_getDefaultNewCameraMatrix_10(JNIEnv*, jclass, jlong, jdouble, jdouble, jboolean);
extern jlong Java_org_opencv_imgproc_Imgproc_getDefaultNewCameraMatrix_11(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_imgproc_Imgproc_getDerivKernels_10(JNIEnv*, jclass, jlong, jlong, jint, jint, jint, jboolean, jint);
extern void Java_org_opencv_imgproc_Imgproc_getDerivKernels_11(JNIEnv*, jclass, jlong, jlong, jint, jint, jint);
extern jlong Java_org_opencv_imgproc_Imgproc_getGaborKernel_10(JNIEnv*, jclass, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jint);
extern jlong Java_org_opencv_imgproc_Imgproc_getGaborKernel_11(JNIEnv*, jclass, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble);
extern jlong Java_org_opencv_imgproc_Imgproc_getGaussianKernel_10(JNIEnv*, jclass, jint, jdouble, jint);
extern jlong Java_org_opencv_imgproc_Imgproc_getGaussianKernel_11(JNIEnv*, jclass, jint, jdouble);
extern jlong Java_org_opencv_imgproc_Imgproc_getPerspectiveTransform_10(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_imgproc_Imgproc_getRectSubPix_10(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jlong, jint);
extern void Java_org_opencv_imgproc_Imgproc_getRectSubPix_11(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jlong);
extern jlong Java_org_opencv_imgproc_Imgproc_getRotationMatrix2D_10(JNIEnv*, jclass, jdouble, jdouble, jdouble, jdouble);
extern jlong Java_org_opencv_imgproc_Imgproc_getStructuringElement_10(JNIEnv*, jclass, jint, jdouble, jdouble, jdouble, jdouble);
extern jlong Java_org_opencv_imgproc_Imgproc_getStructuringElement_11(JNIEnv*, jclass, jint, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_goodFeaturesToTrack_10(JNIEnv*, jclass, jlong, jlong, jint, jdouble, jdouble, jlong, jint, jint, jboolean, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_goodFeaturesToTrack_11(JNIEnv*, jclass, jlong, jlong, jint, jdouble, jdouble, jlong, jint, jint);
extern void Java_org_opencv_imgproc_Imgproc_goodFeaturesToTrack_12(JNIEnv*, jclass, jlong, jlong, jint, jdouble, jdouble, jlong, jint, jboolean, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_goodFeaturesToTrack_13(JNIEnv*, jclass, jlong, jlong, jint, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_grabCut_10(JNIEnv*, jclass, jlong, jlong, jint, jint, jint, jint, jlong, jlong, jint, jint);
extern void Java_org_opencv_imgproc_Imgproc_grabCut_11(JNIEnv*, jclass, jlong, jlong, jint, jint, jint, jint, jlong, jlong, jint);
extern void Java_org_opencv_imgproc_Imgproc_initUndistortRectifyMap_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jdouble, jdouble, jint, jlong, jlong);
extern jfloat Java_org_opencv_imgproc_Imgproc_initWideAngleProjMap_10(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jint, jint, jlong, jlong, jint, jdouble);
extern jfloat Java_org_opencv_imgproc_Imgproc_initWideAngleProjMap_11(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jint, jint, jlong, jlong);
extern void Java_org_opencv_imgproc_Imgproc_integral2_10(JNIEnv*, jclass, jlong, jlong, jlong, jint, jint);
extern void Java_org_opencv_imgproc_Imgproc_integral2_11(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_imgproc_Imgproc_integral3_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jint, jint);
extern void Java_org_opencv_imgproc_Imgproc_integral3_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_imgproc_Imgproc_integral_10(JNIEnv*, jclass, jlong, jlong, jint);
extern void Java_org_opencv_imgproc_Imgproc_integral_11(JNIEnv*, jclass, jlong, jlong);
extern jfloat Java_org_opencv_imgproc_Imgproc_intersectConvexConvex_10(JNIEnv*, jclass, jlong, jlong, jlong, jboolean);
extern jfloat Java_org_opencv_imgproc_Imgproc_intersectConvexConvex_11(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_imgproc_Imgproc_invertAffineTransform_10(JNIEnv*, jclass, jlong, jlong);
extern jboolean Java_org_opencv_imgproc_Imgproc_isContourConvex_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_imgproc_Imgproc_line_10(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jint, jint, jint);
extern void Java_org_opencv_imgproc_Imgproc_line_11(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jint);
extern void Java_org_opencv_imgproc_Imgproc_line_12(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_linearPolar_10(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jdouble, jint);
extern void Java_org_opencv_imgproc_Imgproc_logPolar_10(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jdouble, jint);
extern jdouble Java_org_opencv_imgproc_Imgproc_matchShapes_10(JNIEnv*, jclass, jlong, jlong, jint, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_matchTemplate_10(JNIEnv*, jclass, jlong, jlong, jlong, jint, jlong);
extern void Java_org_opencv_imgproc_Imgproc_matchTemplate_11(JNIEnv*, jclass, jlong, jlong, jlong, jint);
extern void Java_org_opencv_imgproc_Imgproc_medianBlur_10(JNIEnv*, jclass, jlong, jlong, jint);
extern jdoubleArray Java_org_opencv_imgproc_Imgproc_minAreaRect_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_imgproc_Imgproc_minEnclosingCircle_10(JNIEnv*, jclass, jlong, jdoubleArray, jdoubleArray);
extern jdouble Java_org_opencv_imgproc_Imgproc_minEnclosingTriangle_10(JNIEnv*, jclass, jlong, jlong);
extern jdoubleArray Java_org_opencv_imgproc_Imgproc_moments_10(JNIEnv*, jclass, jlong, jboolean);
extern jdoubleArray Java_org_opencv_imgproc_Imgproc_moments_11(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_imgproc_Imgproc_morphologyEx_10(JNIEnv*, jclass, jlong, jlong, jint, jlong, jdouble, jdouble, jint, jint, jdouble, jdouble, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_morphologyEx_11(JNIEnv*, jclass, jlong, jlong, jint, jlong, jdouble, jdouble, jint);
extern void Java_org_opencv_imgproc_Imgproc_morphologyEx_12(JNIEnv*, jclass, jlong, jlong, jint, jlong);
extern jdoubleArray Java_org_opencv_imgproc_Imgproc_n_1getTextSize(JNIEnv*, jclass, jstring, jint, jdouble, jint, jintArray);
extern jdoubleArray Java_org_opencv_imgproc_Imgproc_phaseCorrelate_10(JNIEnv*, jclass, jlong, jlong, jlong, jdoubleArray);
extern jdoubleArray Java_org_opencv_imgproc_Imgproc_phaseCorrelate_11(JNIEnv*, jclass, jlong, jlong);
extern jdouble Java_org_opencv_imgproc_Imgproc_pointPolygonTest_10(JNIEnv*, jclass, jlong, jdouble, jdouble, jboolean);
extern void Java_org_opencv_imgproc_Imgproc_polylines_10(JNIEnv*, jclass, jlong, jlong, jboolean, jdouble, jdouble, jdouble, jdouble, jint, jint, jint);
extern void Java_org_opencv_imgproc_Imgproc_polylines_11(JNIEnv*, jclass, jlong, jlong, jboolean, jdouble, jdouble, jdouble, jdouble, jint);
extern void Java_org_opencv_imgproc_Imgproc_polylines_12(JNIEnv*, jclass, jlong, jlong, jboolean, jdouble, jdouble, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_preCornerDetect_10(JNIEnv*, jclass, jlong, jlong, jint, jint);
extern void Java_org_opencv_imgproc_Imgproc_preCornerDetect_11(JNIEnv*, jclass, jlong, jlong, jint);
extern void Java_org_opencv_imgproc_Imgproc_putText_10(JNIEnv*, jclass, jlong, jstring, jdouble, jdouble, jint, jdouble, jdouble, jdouble, jdouble, jdouble, jint, jint, jboolean);
extern void Java_org_opencv_imgproc_Imgproc_putText_11(JNIEnv*, jclass, jlong, jstring, jdouble, jdouble, jint, jdouble, jdouble, jdouble, jdouble, jdouble, jint);
extern void Java_org_opencv_imgproc_Imgproc_putText_12(JNIEnv*, jclass, jlong, jstring, jdouble, jdouble, jint, jdouble, jdouble, jdouble, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_pyrDown_10(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jint);
extern void Java_org_opencv_imgproc_Imgproc_pyrDown_11(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_pyrDown_12(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_imgproc_Imgproc_pyrMeanShiftFiltering_10(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jint, jint, jint, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_pyrMeanShiftFiltering_11(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_pyrUp_10(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jint);
extern void Java_org_opencv_imgproc_Imgproc_pyrUp_11(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_pyrUp_12(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_imgproc_Imgproc_rectangle_10(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jint, jint, jint);
extern void Java_org_opencv_imgproc_Imgproc_rectangle_11(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jint);
extern void Java_org_opencv_imgproc_Imgproc_rectangle_12(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_remap_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jint, jint, jdouble, jdouble, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_remap_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jint);
extern void Java_org_opencv_imgproc_Imgproc_resize_10(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jdouble, jdouble, jint);
extern void Java_org_opencv_imgproc_Imgproc_resize_11(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble);
extern jint Java_org_opencv_imgproc_Imgproc_rotatedRectangleIntersection_10(JNIEnv*, jclass, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jlong);
extern void Java_org_opencv_imgproc_Imgproc_sepFilter2D_10(JNIEnv*, jclass, jlong, jlong, jint, jlong, jlong, jdouble, jdouble, jdouble, jint);
extern void Java_org_opencv_imgproc_Imgproc_sepFilter2D_11(JNIEnv*, jclass, jlong, jlong, jint, jlong, jlong, jdouble, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_sepFilter2D_12(JNIEnv*, jclass, jlong, jlong, jint, jlong, jlong);
extern void Java_org_opencv_imgproc_Imgproc_spatialGradient_10(JNIEnv*, jclass, jlong, jlong, jlong, jint, jint);
extern void Java_org_opencv_imgproc_Imgproc_spatialGradient_11(JNIEnv*, jclass, jlong, jlong, jlong, jint);
extern void Java_org_opencv_imgproc_Imgproc_spatialGradient_12(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_imgproc_Imgproc_sqrBoxFilter_10(JNIEnv*, jclass, jlong, jlong, jint, jdouble, jdouble, jdouble, jdouble, jboolean, jint);
extern void Java_org_opencv_imgproc_Imgproc_sqrBoxFilter_11(JNIEnv*, jclass, jlong, jlong, jint, jdouble, jdouble, jdouble, jdouble, jboolean);
extern void Java_org_opencv_imgproc_Imgproc_sqrBoxFilter_12(JNIEnv*, jclass, jlong, jlong, jint, jdouble, jdouble);
extern jdouble Java_org_opencv_imgproc_Imgproc_threshold_10(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jint);
extern void Java_org_opencv_imgproc_Imgproc_undistortPointsIter_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong, jint, jint, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_undistortPoints_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_imgproc_Imgproc_undistortPoints_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_imgproc_Imgproc_undistort_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_imgproc_Imgproc_undistort_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_imgproc_Imgproc_warpAffine_10(JNIEnv*, jclass, jlong, jlong, jlong, jdouble, jdouble, jint, jint, jdouble, jdouble, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_warpAffine_11(JNIEnv*, jclass, jlong, jlong, jlong, jdouble, jdouble, jint);
extern void Java_org_opencv_imgproc_Imgproc_warpAffine_12(JNIEnv*, jclass, jlong, jlong, jlong, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_warpPerspective_10(JNIEnv*, jclass, jlong, jlong, jlong, jdouble, jdouble, jint, jint, jdouble, jdouble, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_warpPerspective_11(JNIEnv*, jclass, jlong, jlong, jlong, jdouble, jdouble, jint);
extern void Java_org_opencv_imgproc_Imgproc_warpPerspective_12(JNIEnv*, jclass, jlong, jlong, jlong, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Imgproc_watershed_10(JNIEnv*, jclass, jlong, jlong);
extern jint Java_org_opencv_imgproc_LineSegmentDetector_compareSegments_10(JNIEnv*, jclass, jlong, jdouble, jdouble, jlong, jlong, jlong);
extern jint Java_org_opencv_imgproc_LineSegmentDetector_compareSegments_11(JNIEnv*, jclass, jlong, jdouble, jdouble, jlong, jlong);
extern void Java_org_opencv_imgproc_LineSegmentDetector_delete(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_imgproc_LineSegmentDetector_detect_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_imgproc_LineSegmentDetector_detect_11(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_imgproc_LineSegmentDetector_drawSegments_10(JNIEnv*, jclass, jlong, jlong, jlong);
extern jlong Java_org_opencv_imgproc_Subdiv2D_Subdiv2D_10(JNIEnv*, jclass, jint, jint, jint, jint);
extern jlong Java_org_opencv_imgproc_Subdiv2D_Subdiv2D_11(JNIEnv*, jclass);
extern void Java_org_opencv_imgproc_Subdiv2D_delete(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_imgproc_Subdiv2D_edgeDst_10(JNIEnv*, jclass, jlong, jint, jdoubleArray);
extern jint Java_org_opencv_imgproc_Subdiv2D_edgeDst_11(JNIEnv*, jclass, jlong, jint);
extern jint Java_org_opencv_imgproc_Subdiv2D_edgeOrg_10(JNIEnv*, jclass, jlong, jint, jdoubleArray);
extern jint Java_org_opencv_imgproc_Subdiv2D_edgeOrg_11(JNIEnv*, jclass, jlong, jint);
extern jint Java_org_opencv_imgproc_Subdiv2D_findNearest_10(JNIEnv*, jclass, jlong, jdouble, jdouble, jdoubleArray);
extern jint Java_org_opencv_imgproc_Subdiv2D_findNearest_11(JNIEnv*, jclass, jlong, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Subdiv2D_getEdgeList_10(JNIEnv*, jclass, jlong, jlong);
extern jint Java_org_opencv_imgproc_Subdiv2D_getEdge_10(JNIEnv*, jclass, jlong, jint, jint);
extern void Java_org_opencv_imgproc_Subdiv2D_getLeadingEdgeList_10(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_imgproc_Subdiv2D_getTriangleList_10(JNIEnv*, jclass, jlong, jlong);
extern jdoubleArray Java_org_opencv_imgproc_Subdiv2D_getVertex_10(JNIEnv*, jclass, jlong, jint, jdoubleArray);
extern jdoubleArray Java_org_opencv_imgproc_Subdiv2D_getVertex_11(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_imgproc_Subdiv2D_getVoronoiFacetList_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_imgproc_Subdiv2D_initDelaunay_10(JNIEnv*, jclass, jlong, jint, jint, jint, jint);
extern jint Java_org_opencv_imgproc_Subdiv2D_insert_10(JNIEnv*, jclass, jlong, jdouble, jdouble);
extern void Java_org_opencv_imgproc_Subdiv2D_insert_11(JNIEnv*, jclass, jlong, jlong);
extern jint Java_org_opencv_imgproc_Subdiv2D_locate_10(JNIEnv*, jclass, jlong, jdouble, jdouble, jdoubleArray, jdoubleArray);
extern jint Java_org_opencv_imgproc_Subdiv2D_nextEdge_10(JNIEnv*, jclass, jlong, jint);
extern jint Java_org_opencv_imgproc_Subdiv2D_rotateEdge_10(JNIEnv*, jclass, jlong, jint, jint);
extern jint Java_org_opencv_imgproc_Subdiv2D_symEdge_10(JNIEnv*, jclass, jlong, jint);

*/
import "C"

func CLAHENative_apply_0(nativeObj int64, src_nativeObj int64, dst_nativeObj int64) {
	C.Java_org_opencv_imgproc_CLAHE_apply_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj))
}
func CLAHENative_collectGarbage_0(nativeObj int64) {
	C.Java_org_opencv_imgproc_CLAHE_collectGarbage_10(clzEnv, clzObj, (C.jlong)(nativeObj))
}

func CLAHENative_delete(nativeObj int64) {
	C.Java_org_opencv_imgproc_CLAHE_delete(clzEnv, clzObj, (C.jlong)(nativeObj))
}

func CLAHENative_getClipLimit_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_imgproc_CLAHE_getClipLimit_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func CLAHENative_getTilesGridSize_0(nativeObj int64) []float64 {
	return togofloat64Array(C.Java_org_opencv_imgproc_CLAHE_getTilesGridSize_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func CLAHENative_setClipLimit_0(nativeObj int64, clipLimit float64) {
	C.Java_org_opencv_imgproc_CLAHE_setClipLimit_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jdouble)(clipLimit))
}

func CLAHENative_setTilesGridSize_0(nativeObj int64, tileGridSize_width float64, tileGridSize_height float64) {
	C.Java_org_opencv_imgproc_CLAHE_setTilesGridSize_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jdouble)(tileGridSize_width), (C.jdouble)(tileGridSize_height))
}

func ImgprocNative_Canny_0(dx_nativeObj int64, dy_nativeObj int64, edges_nativeObj int64, threshold1 float64, threshold2 float64, L2gradient bool) {
	C.Java_org_opencv_imgproc_Imgproc_Canny_10(clzEnv, clzObj, (C.jlong)(dx_nativeObj), (C.jlong)(dy_nativeObj), (C.jlong)(edges_nativeObj), (C.jdouble)(threshold1), (C.jdouble)(threshold2), tojboolean(L2gradient))
}

func ImgprocNative_Canny_1(dx_nativeObj int64, dy_nativeObj int64, edges_nativeObj int64, threshold1 float64, threshold2 float64) {
	C.Java_org_opencv_imgproc_Imgproc_Canny_11(clzEnv, clzObj, (C.jlong)(dx_nativeObj), (C.jlong)(dy_nativeObj), (C.jlong)(edges_nativeObj), (C.jdouble)(threshold1), (C.jdouble)(threshold2))
}

func ImgprocNative_Canny_2(image_nativeObj int64, edges_nativeObj int64, threshold1 float64, threshold2 float64, apertureSize int, L2gradient bool) {
	C.Java_org_opencv_imgproc_Imgproc_Canny_12(clzEnv, clzObj, (C.jlong)(image_nativeObj), (C.jlong)(edges_nativeObj), (C.jdouble)(threshold1), (C.jdouble)(threshold2), (C.jint)(apertureSize), tojboolean(L2gradient))
}

func ImgprocNative_Canny_3(image_nativeObj int64, edges_nativeObj int64, threshold1 float64, threshold2 float64) {
	C.Java_org_opencv_imgproc_Imgproc_Canny_13(clzEnv, clzObj, (C.jlong)(image_nativeObj), (C.jlong)(edges_nativeObj), (C.jdouble)(threshold1), (C.jdouble)(threshold2))
}

func ImgprocNative_EMD_0(signature1_nativeObj int64, signature2_nativeObj int64, distType int, cost_nativeObj int64, flow_nativeObj int64) float32 {
	return float32(C.Java_org_opencv_imgproc_Imgproc_EMD_10(clzEnv, clzObj, (C.jlong)(signature1_nativeObj), (C.jlong)(signature2_nativeObj), (C.jint)(distType), (C.jlong)(cost_nativeObj), (C.jlong)(flow_nativeObj)))
}

func ImgprocNative_EMD_1(signature1_nativeObj int64, signature2_nativeObj int64, distType int) float32 {
	return float32(C.Java_org_opencv_imgproc_Imgproc_EMD_11(clzEnv, clzObj, (C.jlong)(signature1_nativeObj), (C.jlong)(signature2_nativeObj), (C.jint)(distType)))
}

func ImgprocNative_GaussianBlur_0(src_nativeObj int64, dst_nativeObj int64, ksize_width float64, ksize_height float64, sigmaX float64, sigmaY float64, borderType int) {
	C.Java_org_opencv_imgproc_Imgproc_GaussianBlur_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jdouble)(ksize_width), (C.jdouble)(ksize_height), (C.jdouble)(sigmaX), (C.jdouble)(sigmaY), (C.jint)(borderType))
}

func ImgprocNative_GaussianBlur_1(src_nativeObj int64, dst_nativeObj int64, ksize_width float64, ksize_height float64, sigmaX float64, sigmaY float64) {
	C.Java_org_opencv_imgproc_Imgproc_GaussianBlur_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jdouble)(ksize_width), (C.jdouble)(ksize_height), (C.jdouble)(sigmaX), (C.jdouble)(sigmaY))
}

func ImgprocNative_GaussianBlur_2(src_nativeObj int64, dst_nativeObj int64, ksize_width float64, ksize_height float64, sigmaX float64) {
	C.Java_org_opencv_imgproc_Imgproc_GaussianBlur_12(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jdouble)(ksize_width), (C.jdouble)(ksize_height), (C.jdouble)(sigmaX))
}

func ImgprocNative_HoughCircles_0(image_nativeObj int64, circles_nativeObj int64, method int, dp float64, minDist float64, param1 float64, param2 float64, minRadius int, maxRadius int) {
	C.Java_org_opencv_imgproc_Imgproc_HoughCircles_10(clzEnv, clzObj, (C.jlong)(image_nativeObj), (C.jlong)(circles_nativeObj), (C.jint)(method), (C.jdouble)(dp), (C.jdouble)(minDist), (C.jdouble)(param1), (C.jdouble)(param2), (C.jint)(minRadius), (C.jint)(maxRadius))
}

func ImgprocNative_HoughCircles_1(image_nativeObj int64, circles_nativeObj int64, method int, dp float64, minDist float64) {
	C.Java_org_opencv_imgproc_Imgproc_HoughCircles_11(clzEnv, clzObj, (C.jlong)(image_nativeObj), (C.jlong)(circles_nativeObj), (C.jint)(method), (C.jdouble)(dp), (C.jdouble)(minDist))
}

func ImgprocNative_HoughLinesP_0(image_nativeObj int64, lines_nativeObj int64, rho float64, theta float64, threshold int, minLineLength float64, maxLineGap float64) {
	C.Java_org_opencv_imgproc_Imgproc_HoughLinesP_10(clzEnv, clzObj, (C.jlong)(image_nativeObj), (C.jlong)(lines_nativeObj), (C.jdouble)(rho), (C.jdouble)(theta), (C.jint)(threshold), (C.jdouble)(minLineLength), (C.jdouble)(maxLineGap))
}

func ImgprocNative_HoughLinesP_1(image_nativeObj int64, lines_nativeObj int64, rho float64, theta float64, threshold int) {
	C.Java_org_opencv_imgproc_Imgproc_HoughLinesP_11(clzEnv, clzObj, (C.jlong)(image_nativeObj), (C.jlong)(lines_nativeObj), (C.jdouble)(rho), (C.jdouble)(theta), (C.jint)(threshold))
}

func ImgprocNative_HoughLines_0(image_nativeObj int64, lines_nativeObj int64, rho float64, theta float64, threshold int, srn float64, stn float64, min_theta float64, max_theta float64) {
	C.Java_org_opencv_imgproc_Imgproc_HoughLines_10(clzEnv, clzObj, (C.jlong)(image_nativeObj), (C.jlong)(lines_nativeObj), (C.jdouble)(rho), (C.jdouble)(theta), (C.jint)(threshold), (C.jdouble)(srn), (C.jdouble)(stn), (C.jdouble)(min_theta), (C.jdouble)(max_theta))
}

func ImgprocNative_HoughLines_1(image_nativeObj int64, lines_nativeObj int64, rho float64, theta float64, threshold int) {
	C.Java_org_opencv_imgproc_Imgproc_HoughLines_11(clzEnv, clzObj, (C.jlong)(image_nativeObj), (C.jlong)(lines_nativeObj), (C.jdouble)(rho), (C.jdouble)(theta), (C.jint)(threshold))
}

func ImgprocNative_HuMoments_0(m_m00 float64, m_m10 float64, m_m01 float64, m_m20 float64, m_m11 float64, m_m02 float64, m_m30 float64, m_m21 float64, m_m12 float64, m_m03 float64, hu_nativeObj int64) {
	C.Java_org_opencv_imgproc_Imgproc_HuMoments_10(clzEnv, clzObj, (C.jdouble)(m_m00), (C.jdouble)(m_m10), (C.jdouble)(m_m01), (C.jdouble)(m_m20), (C.jdouble)(m_m11), (C.jdouble)(m_m02), (C.jdouble)(m_m30), (C.jdouble)(m_m21), (C.jdouble)(m_m12), (C.jdouble)(m_m03), (C.jlong)(hu_nativeObj))
}

func ImgprocNative_Laplacian_0(src_nativeObj int64, dst_nativeObj int64, ddepth int, ksize int, scale float64, delta float64, borderType int) {
	C.Java_org_opencv_imgproc_Imgproc_Laplacian_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(ddepth), (C.jint)(ksize), (C.jdouble)(scale), (C.jdouble)(delta), (C.jint)(borderType))
}

func ImgprocNative_Laplacian_1(src_nativeObj int64, dst_nativeObj int64, ddepth int, ksize int, scale float64, delta float64) {
	C.Java_org_opencv_imgproc_Imgproc_Laplacian_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(ddepth), (C.jint)(ksize), (C.jdouble)(scale), (C.jdouble)(delta))
}

func ImgprocNative_Laplacian_2(src_nativeObj int64, dst_nativeObj int64, ddepth int) {
	C.Java_org_opencv_imgproc_Imgproc_Laplacian_12(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(ddepth))
}

func ImgprocNative_Scharr_0(src_nativeObj int64, dst_nativeObj int64, ddepth int, dx int, dy int, scale float64, delta float64, borderType int) {
	C.Java_org_opencv_imgproc_Imgproc_Scharr_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(ddepth), (C.jint)(dx), (C.jint)(dy), (C.jdouble)(scale), (C.jdouble)(delta), (C.jint)(borderType))
}

func ImgprocNative_Scharr_1(src_nativeObj int64, dst_nativeObj int64, ddepth int, dx int, dy int, scale float64, delta float64) {
	C.Java_org_opencv_imgproc_Imgproc_Scharr_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(ddepth), (C.jint)(dx), (C.jint)(dy), (C.jdouble)(scale), (C.jdouble)(delta))
}

func ImgprocNative_Scharr_2(src_nativeObj int64, dst_nativeObj int64, ddepth int, dx int, dy int) {
	C.Java_org_opencv_imgproc_Imgproc_Scharr_12(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(ddepth), (C.jint)(dx), (C.jint)(dy))
}

func ImgprocNative_Sobel_0(src_nativeObj int64, dst_nativeObj int64, ddepth int, dx int, dy int, ksize int, scale float64, delta float64, borderType int) {
	C.Java_org_opencv_imgproc_Imgproc_Sobel_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(ddepth), (C.jint)(dx), (C.jint)(dy), (C.jint)(ksize), (C.jdouble)(scale), (C.jdouble)(delta), (C.jint)(borderType))
}

func ImgprocNative_Sobel_1(src_nativeObj int64, dst_nativeObj int64, ddepth int, dx int, dy int, ksize int, scale float64, delta float64) {
	C.Java_org_opencv_imgproc_Imgproc_Sobel_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(ddepth), (C.jint)(dx), (C.jint)(dy), (C.jint)(ksize), (C.jdouble)(scale), (C.jdouble)(delta))
}

func ImgprocNative_Sobel_2(src_nativeObj int64, dst_nativeObj int64, ddepth int, dx int, dy int) {
	C.Java_org_opencv_imgproc_Imgproc_Sobel_12(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(ddepth), (C.jint)(dx), (C.jint)(dy))
}

func ImgprocNative_accumulateProduct_0(src1_nativeObj int64, src2_nativeObj int64, dst_nativeObj int64, mask_nativeObj int64) {
	C.Java_org_opencv_imgproc_Imgproc_accumulateProduct_10(clzEnv, clzObj, (C.jlong)(src1_nativeObj),
		(C.jlong)(src2_nativeObj), (C.jlong)(dst_nativeObj), (C.jlong)(mask_nativeObj))
}

func ImgprocNative_accumulateProduct_1(src1_nativeObj int64, src2_nativeObj int64, dst_nativeObj int64) {
	C.Java_org_opencv_imgproc_Imgproc_accumulateProduct_11(clzEnv, clzObj, (C.jlong)(src1_nativeObj),
		(C.jlong)(src2_nativeObj), (C.jlong)(dst_nativeObj))
}

func ImgprocNative_accumulateSquare_0(src_nativeObj int64, dst_nativeObj int64, mask_nativeObj int64) {
	C.Java_org_opencv_imgproc_Imgproc_accumulateSquare_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jlong)(mask_nativeObj))
}

func ImgprocNative_accumulateSquare_1(src_nativeObj int64, dst_nativeObj int64) {
	C.Java_org_opencv_imgproc_Imgproc_accumulateSquare_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj))
}

func ImgprocNative_accumulateWeighted_0(src_nativeObj int64, dst_nativeObj int64, alpha float64, mask_nativeObj int64) {
	C.Java_org_opencv_imgproc_Imgproc_accumulateWeighted_10(clzEnv, clzObj, (C.jlong)(src_nativeObj),
		(C.jlong)(dst_nativeObj), (C.jdouble)(alpha), (C.jlong)(mask_nativeObj))
}

func ImgprocNative_accumulateWeighted_1(src_nativeObj int64, dst_nativeObj int64, alpha float64) {
	C.Java_org_opencv_imgproc_Imgproc_accumulateWeighted_11(clzEnv, clzObj, (C.jlong)(src_nativeObj),
		(C.jlong)(dst_nativeObj), (C.jdouble)(alpha))
}

func ImgprocNative_accumulate_0(src_nativeObj int64, dst_nativeObj int64, mask_nativeObj int64) {
	C.Java_org_opencv_imgproc_Imgproc_accumulate_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jlong)(mask_nativeObj))
}

func ImgprocNative_accumulate_1(src_nativeObj int64, dst_nativeObj int64) {
	C.Java_org_opencv_imgproc_Imgproc_accumulate_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj))
}

func ImgprocNative_adaptiveThreshold_0(src_nativeObj int64, dst_nativeObj int64, maxValue float64, adaptiveMethod int, thresholdType int, blockSize int, C float64) {
	C.Java_org_opencv_imgproc_Imgproc_adaptiveThreshold_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jdouble)(maxValue), (C.jint)(adaptiveMethod), (C.jint)(thresholdType), (C.jint)(blockSize), (C.jdouble)(C))
}

func ImgprocNative_applyColorMap_0(src_nativeObj int64, dst_nativeObj int64, userColor_nativeObj int64) {
	C.Java_org_opencv_imgproc_Imgproc_applyColorMap_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jlong)(userColor_nativeObj))
}

func ImgprocNative_applyColorMap_1(src_nativeObj int64, dst_nativeObj int64, colormap int) {
	C.Java_org_opencv_imgproc_Imgproc_applyColorMap_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(colormap))
}

func ImgprocNative_approxPolyDP_0(curve_mat_nativeObj int64, approxCurve_mat_nativeObj int64, epsilon float64, closed bool) {
	C.Java_org_opencv_imgproc_Imgproc_approxPolyDP_10(clzEnv, clzObj, (C.jlong)(curve_mat_nativeObj),
		(C.jlong)(approxCurve_mat_nativeObj), (C.jdouble)(epsilon), tojboolean(closed))
}

func ImgprocNative_arcLength_0(curve_mat_nativeObj int64, closed bool) float64 {
	return float64(C.Java_org_opencv_imgproc_Imgproc_arcLength_10(clzEnv, clzObj, (C.jlong)(curve_mat_nativeObj), tojboolean(closed)))
}

func ImgprocNative_arrowedLine_0(img_nativeObj int64, pt1_x float64, pt1_y float64, pt2_x float64, pt2_y float64, color_val0 float64, color_val1 float64, color_val2 float64, color_val3 float64, thickness int, line_type int, shift int, tipLength float64) {
	C.Java_org_opencv_imgproc_Imgproc_arrowedLine_10(clzEnv, clzObj, (C.jlong)(img_nativeObj), (C.jdouble)(pt1_x), (C.jdouble)(pt1_y), (C.jdouble)(pt2_x),
		(C.jdouble)(pt2_y), (C.jdouble)(color_val0),
		(C.jdouble)(color_val1), (C.jdouble)(color_val2), (C.jdouble)(color_val3), (C.jint)(thickness), (C.jint)(line_type), (C.jint)(shift), (C.jdouble)(tipLength))
}

func ImgprocNative_arrowedLine_1(img_nativeObj int64, pt1_x float64, pt1_y float64, pt2_x float64, pt2_y float64, color_val0 float64, color_val1 float64, color_val2 float64, color_val3 float64) {
	C.Java_org_opencv_imgproc_Imgproc_arrowedLine_11(clzEnv, clzObj, (C.jlong)(img_nativeObj), (C.jdouble)(pt1_x), (C.jdouble)(pt1_y), (C.jdouble)(pt2_x),
		(C.jdouble)(pt2_y), (C.jdouble)(color_val0),
		(C.jdouble)(color_val1), (C.jdouble)(color_val2), (C.jdouble)(color_val3))
}

func ImgprocNative_bilateralFilter_0(src_nativeObj int64, dst_nativeObj int64, d int, sigmaColor float64, sigmaSpace float64, borderType int) {
	C.Java_org_opencv_imgproc_Imgproc_bilateralFilter_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(d), (C.jdouble)(sigmaColor), (C.jdouble)(sigmaSpace), (C.jint)(borderType))
}

func ImgprocNative_bilateralFilter_1(src_nativeObj int64, dst_nativeObj int64, d int, sigmaColor float64, sigmaSpace float64) {
	C.Java_org_opencv_imgproc_Imgproc_bilateralFilter_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(d), (C.jdouble)(sigmaColor), (C.jdouble)(sigmaSpace))
}

func ImgprocNative_blur_0(src_nativeObj int64, dst_nativeObj int64, ksize_width float64, ksize_height float64, anchor_x float64, anchor_y float64, borderType int) {
	C.Java_org_opencv_imgproc_Imgproc_blur_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jdouble)(ksize_width), (C.jdouble)(ksize_height), (C.jdouble)(anchor_x), (C.jdouble)(anchor_y), (C.jint)(borderType))
}

func ImgprocNative_blur_1(src_nativeObj int64, dst_nativeObj int64, ksize_width float64, ksize_height float64, anchor_x float64, anchor_y float64) {
	C.Java_org_opencv_imgproc_Imgproc_blur_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jdouble)(ksize_width), (C.jdouble)(ksize_height), (C.jdouble)(anchor_x), (C.jdouble)(anchor_y))
}

func ImgprocNative_blur_2(src_nativeObj int64, dst_nativeObj int64, ksize_width float64, ksize_height float64) {
	C.Java_org_opencv_imgproc_Imgproc_blur_12(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jdouble)(ksize_width), (C.jdouble)(ksize_height))
}

func ImgprocNative_boundingRect_0(points_mat_nativeObj int64) []float64 {
	return togofloat64Array(C.Java_org_opencv_imgproc_Imgproc_boundingRect_10(clzEnv, clzObj, (C.jlong)(points_mat_nativeObj)))
}

func ImgprocNative_boxFilter_0(src_nativeObj int64, dst_nativeObj int64, ddepth int, ksize_width float64, ksize_height float64, anchor_x float64, anchor_y float64, normalize bool, borderType int) {
	C.Java_org_opencv_imgproc_Imgproc_boxFilter_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(ddepth), (C.jdouble)(ksize_width), (C.jdouble)(ksize_height), (C.jdouble)(anchor_x), (C.jdouble)(anchor_y), tojboolean(normalize), (C.jint)(borderType))
}

func ImgprocNative_boxFilter_1(src_nativeObj int64, dst_nativeObj int64, ddepth int, ksize_width float64, ksize_height float64, anchor_x float64, anchor_y float64, normalize bool) {
	C.Java_org_opencv_imgproc_Imgproc_boxFilter_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(ddepth), (C.jdouble)(ksize_width), (C.jdouble)(ksize_height), (C.jdouble)(anchor_x), (C.jdouble)(anchor_y), tojboolean(normalize))
}

func ImgprocNative_boxFilter_2(src_nativeObj int64, dst_nativeObj int64, ddepth int, ksize_width float64, ksize_height float64) {
	C.Java_org_opencv_imgproc_Imgproc_boxFilter_12(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(ddepth), (C.jdouble)(ksize_width), (C.jdouble)(ksize_height))
}

func ImgprocNative_boxPoints_0(box_center_x float64, box_center_y float64, box_size_width float64, box_size_height float64, box_angle float64, points_nativeObj int64) {
	C.Java_org_opencv_imgproc_Imgproc_boxPoints_10(clzEnv, clzObj, (C.jdouble)(box_center_x), (C.jdouble)(box_center_y), (C.jdouble)(box_size_width), (C.jdouble)(box_size_height), (C.jdouble)(box_angle), (C.jlong)(points_nativeObj))
}

func ImgprocNative_calcBackProject_0(images_mat_nativeObj int64, channels_mat_nativeObj int64, hist_nativeObj int64, dst_nativeObj int64, ranges_mat_nativeObj int64, scale float64) {
	C.Java_org_opencv_imgproc_Imgproc_calcBackProject_10(clzEnv, clzObj, (C.jlong)(images_mat_nativeObj), (C.jlong)(channels_mat_nativeObj), (C.jlong)(hist_nativeObj), (C.jlong)(dst_nativeObj), (C.jlong)(ranges_mat_nativeObj), (C.jdouble)(scale))
}

func ImgprocNative_calcHist_0(images_mat_nativeObj int64, channels_mat_nativeObj int64, mask_nativeObj int64, hist_nativeObj int64, histSize_mat_nativeObj int64, ranges_mat_nativeObj int64, accumulate bool) {
	C.Java_org_opencv_imgproc_Imgproc_calcHist_10(clzEnv, clzObj, (C.jlong)(images_mat_nativeObj), (C.jlong)(channels_mat_nativeObj), (C.jlong)(mask_nativeObj), (C.jlong)(hist_nativeObj), (C.jlong)(histSize_mat_nativeObj), (C.jlong)(ranges_mat_nativeObj), tojboolean(accumulate))
}

func ImgprocNative_calcHist_1(images_mat_nativeObj int64, channels_mat_nativeObj int64, mask_nativeObj int64, hist_nativeObj int64, histSize_mat_nativeObj int64, ranges_mat_nativeObj int64) {
	C.Java_org_opencv_imgproc_Imgproc_calcHist_11(clzEnv, clzObj, (C.jlong)(images_mat_nativeObj), (C.jlong)(channels_mat_nativeObj), (C.jlong)(mask_nativeObj), (C.jlong)(hist_nativeObj), (C.jlong)(histSize_mat_nativeObj), (C.jlong)(ranges_mat_nativeObj))
}

func ImgprocNative_circle_0(img_nativeObj int64, center_x float64, center_y float64, radius int, color_val0 float64, color_val1 float64, color_val2 float64, color_val3 float64, thickness int, lineType int, shift int) {
	C.Java_org_opencv_imgproc_Imgproc_circle_10(clzEnv, clzObj, (C.jlong)(img_nativeObj), (C.jdouble)(center_x), (C.jdouble)(center_y), (C.jint)(radius), (C.jdouble)(color_val0), (C.jdouble)(color_val1), (C.jdouble)(color_val2), (C.jdouble)(color_val3), (C.jint)(thickness), (C.jint)(lineType), (C.jint)(shift))
}

func ImgprocNative_circle_1(img_nativeObj int64, center_x float64, center_y float64, radius int, color_val0 float64, color_val1 float64, color_val2 float64, color_val3 float64, thickness int) {
	C.Java_org_opencv_imgproc_Imgproc_circle_11(clzEnv, clzObj, (C.jlong)(img_nativeObj), (C.jdouble)(center_x), (C.jdouble)(center_y), (C.jint)(radius), (C.jdouble)(color_val0), (C.jdouble)(color_val1), (C.jdouble)(color_val2), (C.jdouble)(color_val3), (C.jint)(thickness))
}

func ImgprocNative_circle_2(img_nativeObj int64, center_x float64, center_y float64, radius int, color_val0 float64, color_val1 float64, color_val2 float64, color_val3 float64) {
	C.Java_org_opencv_imgproc_Imgproc_circle_12(clzEnv, clzObj, (C.jlong)(img_nativeObj), (C.jdouble)(center_x), (C.jdouble)(center_y), (C.jint)(radius), (C.jdouble)(color_val0), (C.jdouble)(color_val1), (C.jdouble)(color_val2), (C.jdouble)(color_val3))
}

func ImgprocNative_clipLine_0(imgRect_x int, imgRect_y int, imgRect_width int, imgRect_height int, pt1_x float64, pt1_y float64, pt1_out []float64, pt2_x float64, pt2_y float64, pt2_out []float64) bool {
	defer ungointerface(pt1_out)
	defer ungointerface(pt2_out)
	return togobool(C.Java_org_opencv_imgproc_Imgproc_clipLine_10(clzEnv, clzObj, (C.jint)(imgRect_x), (C.jint)(imgRect_y), (C.jint)(imgRect_width), (C.jint)(imgRect_height), (C.jdouble)(pt1_x), (C.jdouble)(pt1_y), tojdoubleArray(pt1_out), (C.jdouble)(pt2_x), (C.jdouble)(pt2_y), tojdoubleArray(pt2_out)))
}

func ImgprocNative_compareHist_0(H1_nativeObj int64, H2_nativeObj int64, method int) float64 {
	return float64(C.Java_org_opencv_imgproc_Imgproc_compareHist_10(clzEnv, clzObj, (C.jlong)(H1_nativeObj), (C.jlong)(H2_nativeObj), (C.jint)(method)))
}

func ImgprocNative_connectedComponentsWithAlgorithm_0(image_nativeObj int64, labels_nativeObj int64, connectivity int, ltype int, ccltype int) int {
	return int(C.Java_org_opencv_imgproc_Imgproc_connectedComponentsWithAlgorithm_10(clzEnv, clzObj, (C.jlong)(image_nativeObj), (C.jlong)(labels_nativeObj),
		(C.jint)(connectivity), (C.jint)(ltype), (C.jint)(ccltype)))
}

func ImgprocNative_connectedComponentsWithStatsWithAlgorithm_0(image_nativeObj int64, labels_nativeObj int64, stats_nativeObj int64, centroids_nativeObj int64, connectivity int, ltype int, ccltype int) int {
	return int(C.Java_org_opencv_imgproc_Imgproc_connectedComponentsWithStatsWithAlgorithm_10(clzEnv, clzObj, (C.jlong)(image_nativeObj), (C.jlong)(labels_nativeObj), (C.jlong)(stats_nativeObj), (C.jlong)(centroids_nativeObj), (C.jint)(connectivity), (C.jint)(ltype), (C.jint)(ccltype)))
}

func ImgprocNative_connectedComponentsWithStats_0(image_nativeObj int64, labels_nativeObj int64, stats_nativeObj int64, centroids_nativeObj int64, connectivity int, ltype int) int {
	return int(C.Java_org_opencv_imgproc_Imgproc_connectedComponentsWithStats_10(clzEnv, clzObj, (C.jlong)(image_nativeObj), (C.jlong)(labels_nativeObj), (C.jlong)(stats_nativeObj), (C.jlong)(centroids_nativeObj), (C.jint)(connectivity), (C.jint)(ltype)))
}

func ImgprocNative_connectedComponentsWithStats_1(image_nativeObj int64, labels_nativeObj int64, stats_nativeObj int64, centroids_nativeObj int64) int {
	return int(C.Java_org_opencv_imgproc_Imgproc_connectedComponentsWithStats_11(clzEnv, clzObj, (C.jlong)(image_nativeObj), (C.jlong)(labels_nativeObj), (C.jlong)(stats_nativeObj), (C.jlong)(centroids_nativeObj)))
}

func ImgprocNative_connectedComponents_0(image_nativeObj int64, labels_nativeObj int64, connectivity int, ltype int) int {
	return int(C.Java_org_opencv_imgproc_Imgproc_connectedComponents_10(clzEnv, clzObj, (C.jlong)(image_nativeObj), (C.jlong)(labels_nativeObj), (C.jint)(connectivity), (C.jint)(ltype)))
}

func ImgprocNative_connectedComponents_1(image_nativeObj int64, labels_nativeObj int64) int {
	return int(C.Java_org_opencv_imgproc_Imgproc_connectedComponents_11(clzEnv, clzObj, (C.jlong)(image_nativeObj), (C.jlong)(labels_nativeObj)))
}

func ImgprocNative_contourArea_0(contour_nativeObj int64, oriented bool) float64 {
	return float64(C.Java_org_opencv_imgproc_Imgproc_contourArea_10(clzEnv, clzObj, (C.jlong)(contour_nativeObj), tojboolean(oriented)))
}

func ImgprocNative_contourArea_1(contour_nativeObj int64) float64 {
	return float64(C.Java_org_opencv_imgproc_Imgproc_contourArea_11(clzEnv, clzObj, (C.jlong)(contour_nativeObj)))
}

func ImgprocNative_convertMaps_0(map1_nativeObj int64, map2_nativeObj int64, dstmap1_nativeObj int64, dstmap2_nativeObj int64, dstmap1type int, nninterpolation bool) {
	C.Java_org_opencv_imgproc_Imgproc_convertMaps_10(clzEnv, clzObj, (C.jlong)(map1_nativeObj), (C.jlong)(map2_nativeObj), (C.jlong)(dstmap1_nativeObj), (C.jlong)(dstmap2_nativeObj), (C.jint)(dstmap1type), tojboolean(nninterpolation))
}

func ImgprocNative_convertMaps_1(map1_nativeObj int64, map2_nativeObj int64, dstmap1_nativeObj int64, dstmap2_nativeObj int64, dstmap1type int) {
	C.Java_org_opencv_imgproc_Imgproc_convertMaps_11(clzEnv, clzObj, (C.jlong)(map1_nativeObj), (C.jlong)(map2_nativeObj), (C.jlong)(dstmap1_nativeObj), (C.jlong)(dstmap2_nativeObj), (C.jint)(dstmap1type))
}

func ImgprocNative_convexHull_0(points_mat_nativeObj int64, hull_mat_nativeObj int64, clockwise bool) {
	C.Java_org_opencv_imgproc_Imgproc_convexHull_10(clzEnv, clzObj, (C.jlong)(points_mat_nativeObj), (C.jlong)(hull_mat_nativeObj), tojboolean(clockwise))
}

func ImgprocNative_convexHull_1(points_mat_nativeObj int64, hull_mat_nativeObj int64) {
	C.Java_org_opencv_imgproc_Imgproc_convexHull_11(clzEnv, clzObj, (C.jlong)(points_mat_nativeObj), (C.jlong)(hull_mat_nativeObj))
}

func ImgprocNative_convexityDefects_0(contour_mat_nativeObj int64, convexhull_mat_nativeObj int64, convexityDefects_mat_nativeObj int64) {
	C.Java_org_opencv_imgproc_Imgproc_convexityDefects_10(clzEnv, clzObj, (C.jlong)(contour_mat_nativeObj), (C.jlong)(convexhull_mat_nativeObj), (C.jlong)(convexityDefects_mat_nativeObj))
}

func ImgprocNative_cornerEigenValsAndVecs_0(src_nativeObj int64, dst_nativeObj int64, blockSize int, ksize int, borderType int) {
	C.Java_org_opencv_imgproc_Imgproc_cornerEigenValsAndVecs_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(blockSize), (C.jint)(ksize), (C.jint)(borderType))
}

func ImgprocNative_cornerEigenValsAndVecs_1(src_nativeObj int64, dst_nativeObj int64, blockSize int, ksize int) {
	C.Java_org_opencv_imgproc_Imgproc_cornerEigenValsAndVecs_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(blockSize), (C.jint)(ksize))
}

func ImgprocNative_cornerHarris_0(src_nativeObj int64, dst_nativeObj int64, blockSize int, ksize int, k float64, borderType int) {
	C.Java_org_opencv_imgproc_Imgproc_cornerHarris_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(blockSize), (C.jint)(ksize), (C.jdouble)(k), (C.jint)(borderType))
}

func ImgprocNative_cornerHarris_1(src_nativeObj int64, dst_nativeObj int64, blockSize int, ksize int, k float64) {
	C.Java_org_opencv_imgproc_Imgproc_cornerHarris_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(blockSize), (C.jint)(ksize), (C.jdouble)(k))
}

func ImgprocNative_cornerMinEigenVal_0(src_nativeObj int64, dst_nativeObj int64, blockSize int, ksize int, borderType int) {
	C.Java_org_opencv_imgproc_Imgproc_cornerMinEigenVal_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(blockSize), (C.jint)(ksize), (C.jint)(borderType))
}

func ImgprocNative_cornerMinEigenVal_1(src_nativeObj int64, dst_nativeObj int64, blockSize int, ksize int) {
	C.Java_org_opencv_imgproc_Imgproc_cornerMinEigenVal_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(blockSize), (C.jint)(ksize))
}

func ImgprocNative_cornerMinEigenVal_2(src_nativeObj int64, dst_nativeObj int64, blockSize int) {
	C.Java_org_opencv_imgproc_Imgproc_cornerMinEigenVal_12(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(blockSize))
}

func ImgprocNative_cornerSubPix_0(image_nativeObj int64, corners_nativeObj int64, winSize_width float64, winSize_height float64, zeroZone_width float64, zeroZone_height float64, criteria_type int, criteria_maxCount int, criteria_epsilon float64) {
	C.Java_org_opencv_imgproc_Imgproc_cornerSubPix_10(clzEnv, clzObj, (C.jlong)(image_nativeObj), (C.jlong)(corners_nativeObj), (C.jdouble)(winSize_width), (C.jdouble)(winSize_height), (C.jdouble)(zeroZone_width), (C.jdouble)(zeroZone_height), (C.jint)(criteria_type), (C.jint)(criteria_maxCount), (C.jdouble)(criteria_epsilon))
}

func ImgprocNative_createCLAHE_0(clipLimit float64, tileGridSize_width float64, tileGridSize_height float64) int64 {
	return int64(C.Java_org_opencv_imgproc_Imgproc_createCLAHE_10(clzEnv, clzObj, (C.jdouble)(clipLimit), (C.jdouble)(tileGridSize_width), (C.jdouble)(tileGridSize_height)))
}

func ImgprocNative_createCLAHE_1() int64 {
	return int64(C.Java_org_opencv_imgproc_Imgproc_createCLAHE_11(clzEnv, clzObj))
}

func ImgprocNative_createHanningWindow_0(dst_nativeObj int64, winSize_width float64, winSize_height float64, rtype int) {
	C.Java_org_opencv_imgproc_Imgproc_createHanningWindow_10(clzEnv, clzObj, (C.jlong)(dst_nativeObj), (C.jdouble)(winSize_width), (C.jdouble)(winSize_height), (C.jint)(rtype))
}

func ImgprocNative_createLineSegmentDetector_0(_refine int, _scale float64, _sigma_scale float64, _quant float64, _ang_th float64, _log_eps float64, _density_th float64, _n_bins int) int64 {
	return int64(C.Java_org_opencv_imgproc_Imgproc_createLineSegmentDetector_10(clzEnv, clzObj, (C.jint)(_refine), (C.jdouble)(_scale), (C.jdouble)(_sigma_scale), (C.jdouble)(_quant), (C.jdouble)(_ang_th), (C.jdouble)(_log_eps), (C.jdouble)(_density_th), (C.jint)(_n_bins)))
}

func ImgprocNative_createLineSegmentDetector_1() int64 {
	return int64(C.Java_org_opencv_imgproc_Imgproc_createLineSegmentDetector_11(clzEnv, clzObj))
}

func ImgprocNative_cvtColor_0(src_nativeObj int64, dst_nativeObj int64, code int, dstCn int) {
	C.Java_org_opencv_imgproc_Imgproc_cvtColor_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(code), (C.jint)(dstCn))
}

func ImgprocNative_cvtColor_1(src_nativeObj int64, dst_nativeObj int64, code int) {
	C.Java_org_opencv_imgproc_Imgproc_cvtColor_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(code))
}

func ImgprocNative_demosaicing_0(_src_nativeObj int64, _dst_nativeObj int64, code int, dcn int) {
	C.Java_org_opencv_imgproc_Imgproc_demosaicing_10(clzEnv, clzObj, (C.jlong)(_src_nativeObj), (C.jlong)(_dst_nativeObj), (C.jint)(code), (C.jint)(dcn))
}

func ImgprocNative_demosaicing_1(_src_nativeObj int64, _dst_nativeObj int64, code int) {
	C.Java_org_opencv_imgproc_Imgproc_demosaicing_11(clzEnv, clzObj, (C.jlong)(_src_nativeObj), (C.jlong)(_dst_nativeObj), (C.jint)(code))
}

func ImgprocNative_dilate_0(src_nativeObj int64, dst_nativeObj int64, kernel_nativeObj int64, anchor_x float64, anchor_y float64, iterations int, borderType int, borderValue_val0 float64, borderValue_val1 float64, borderValue_val2 float64, borderValue_val3 float64) {
	C.Java_org_opencv_imgproc_Imgproc_dilate_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jlong)(kernel_nativeObj), (C.jdouble)(anchor_x), (C.jdouble)(anchor_y), (C.jint)(iterations), (C.jint)(borderType), (C.jdouble)(borderValue_val0), (C.jdouble)(borderValue_val1), (C.jdouble)(borderValue_val2),
		(C.jdouble)(borderValue_val3))
}

func ImgprocNative_dilate_1(src_nativeObj int64, dst_nativeObj int64, kernel_nativeObj int64, anchor_x float64, anchor_y float64, iterations int) {
	C.Java_org_opencv_imgproc_Imgproc_dilate_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jlong)(kernel_nativeObj), (C.jdouble)(anchor_x), (C.jdouble)(anchor_y), (C.jint)(iterations))
}

func ImgprocNative_dilate_2(src_nativeObj int64, dst_nativeObj int64, kernel_nativeObj int64) {
	C.Java_org_opencv_imgproc_Imgproc_dilate_12(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jlong)(kernel_nativeObj))
}

func ImgprocNative_distanceTransformWithLabels_0(src_nativeObj int64, dst_nativeObj int64, labels_nativeObj int64, distanceType int, maskSize int, labelType int) {
	C.Java_org_opencv_imgproc_Imgproc_distanceTransformWithLabels_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jlong)(labels_nativeObj), (C.jint)(distanceType), (C.jint)(maskSize), (C.jint)(labelType))
}

func ImgprocNative_distanceTransformWithLabels_1(src_nativeObj int64, dst_nativeObj int64, labels_nativeObj int64, distanceType int, maskSize int) {
	C.Java_org_opencv_imgproc_Imgproc_distanceTransformWithLabels_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jlong)(labels_nativeObj), (C.jint)(distanceType), (C.jint)(maskSize))
}

func ImgprocNative_distanceTransform_0(src_nativeObj int64, dst_nativeObj int64, distanceType int, maskSize int, dstType int) {
	C.Java_org_opencv_imgproc_Imgproc_distanceTransform_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(distanceType), (C.jint)(maskSize), (C.jint)(dstType))
}

func ImgprocNative_distanceTransform_1(src_nativeObj int64, dst_nativeObj int64, distanceType int, maskSize int) {
	C.Java_org_opencv_imgproc_Imgproc_distanceTransform_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(distanceType), (C.jint)(maskSize))
}

func ImgprocNative_drawContours_0(image_nativeObj int64, contours_mat_nativeObj int64, contourIdx int, color_val0 float64, color_val1 float64, color_val2 float64, color_val3 float64, thickness int, lineType int, hierarchy_nativeObj int64, maxLevel int, offset_x float64, offset_y float64) {
	C.Java_org_opencv_imgproc_Imgproc_drawContours_10(clzEnv, clzObj, (C.jlong)(image_nativeObj), (C.jlong)(contours_mat_nativeObj), (C.jint)(contourIdx),
		(C.jdouble)(color_val0), (C.jdouble)(color_val1), (C.jdouble)(color_val2), (C.jdouble)(color_val3), (C.jint)(thickness), (C.jint)(lineType), (C.jlong)(hierarchy_nativeObj), (C.jint)(maxLevel), (C.jdouble)(offset_x), (C.jdouble)(offset_y))
}

func ImgprocNative_drawContours_1(image_nativeObj int64, contours_mat_nativeObj int64, contourIdx int, color_val0 float64, color_val1 float64, color_val2 float64, color_val3 float64, thickness int) {
	C.Java_org_opencv_imgproc_Imgproc_drawContours_11(clzEnv, clzObj, (C.jlong)(image_nativeObj), (C.jlong)(contours_mat_nativeObj), (C.jint)(contourIdx),
		(C.jdouble)(color_val0), (C.jdouble)(color_val1), (C.jdouble)(color_val2), (C.jdouble)(color_val3), (C.jint)(thickness))
}

func ImgprocNative_drawContours_2(image_nativeObj int64, contours_mat_nativeObj int64, contourIdx int, color_val0 float64, color_val1 float64, color_val2 float64, color_val3 float64) {
	C.Java_org_opencv_imgproc_Imgproc_drawContours_12(clzEnv, clzObj, (C.jlong)(image_nativeObj), (C.jlong)(contours_mat_nativeObj), (C.jint)(contourIdx),
		(C.jdouble)(color_val0), (C.jdouble)(color_val1), (C.jdouble)(color_val2), (C.jdouble)(color_val3))
}

func ImgprocNative_drawMarker_0(img_nativeObj int64, position_x float64, position_y float64, color_val0 float64, color_val1 float64, color_val2 float64, color_val3 float64, markerType int, markerSize int, thickness int, line_type int) {
	C.Java_org_opencv_imgproc_Imgproc_drawMarker_10(clzEnv, clzObj, (C.jlong)(img_nativeObj), (C.jdouble)(position_x), (C.jdouble)(position_y), (C.jdouble)(color_val0), (C.jdouble)(color_val1), (C.jdouble)(color_val2), (C.jdouble)(color_val3), (C.jint)(markerType), (C.jint)(markerSize), (C.jint)(thickness), (C.jint)(line_type))
}

func ImgprocNative_drawMarker_1(img_nativeObj int64, position_x float64, position_y float64, color_val0 float64, color_val1 float64, color_val2 float64, color_val3 float64) {
	C.Java_org_opencv_imgproc_Imgproc_drawMarker_11(clzEnv, clzObj, (C.jlong)(img_nativeObj), (C.jdouble)(position_x), (C.jdouble)(position_y), (C.jdouble)(color_val0), (C.jdouble)(color_val1), (C.jdouble)(color_val2), (C.jdouble)(color_val3))
}

func ImgprocNative_ellipse2Poly_0(center_x float64, center_y float64, axes_width float64, axes_height float64, angle int, arcStart int, arcEnd int, delta int, pts_mat_nativeObj int64) {
	C.Java_org_opencv_imgproc_Imgproc_ellipse2Poly_10(clzEnv, clzObj, (C.jdouble)(center_x), (C.jdouble)(center_y), (C.jdouble)(axes_width), (C.jdouble)(axes_height), (C.jint)(angle), (C.jint)(arcStart), (C.jint)(arcEnd), (C.jint)(delta), (C.jlong)(pts_mat_nativeObj))
}

func ImgprocNative_ellipse_0(img_nativeObj int64, center_x float64, center_y float64, axes_width float64, axes_height float64, angle float64, startAngle float64, endAngle float64, color_val0 float64, color_val1 float64, color_val2 float64, color_val3 float64, thickness int, lineType int, shift int) {
	C.Java_org_opencv_imgproc_Imgproc_ellipse_10(clzEnv, clzObj, (C.jlong)(img_nativeObj), (C.jdouble)(center_x), (C.jdouble)(center_y), (C.jdouble)(axes_width), (C.jdouble)(axes_height), (C.jdouble)(angle), (C.jdouble)(startAngle), (C.jdouble)(endAngle), (C.jdouble)(color_val0), (C.jdouble)(color_val1), (C.jdouble)(color_val2), (C.jdouble)(color_val3), (C.jint)(thickness), (C.jint)(lineType), (C.jint)(shift))
}

func ImgprocNative_ellipse_1(img_nativeObj int64, center_x float64, center_y float64, axes_width float64, axes_height float64, angle float64, startAngle float64, endAngle float64, color_val0 float64, color_val1 float64, color_val2 float64, color_val3 float64, thickness int) {
	C.Java_org_opencv_imgproc_Imgproc_ellipse_11(clzEnv, clzObj, (C.jlong)(img_nativeObj), (C.jdouble)(center_x), (C.jdouble)(center_y), (C.jdouble)(axes_width), (C.jdouble)(axes_height), (C.jdouble)(angle), (C.jdouble)(startAngle), (C.jdouble)(endAngle), (C.jdouble)(color_val0), (C.jdouble)(color_val1), (C.jdouble)(color_val2), (C.jdouble)(color_val3), (C.jint)(thickness))
}

func ImgprocNative_ellipse_2(img_nativeObj int64, center_x float64, center_y float64, axes_width float64, axes_height float64, angle float64, startAngle float64, endAngle float64, color_val0 float64, color_val1 float64, color_val2 float64, color_val3 float64) {
	C.Java_org_opencv_imgproc_Imgproc_ellipse_12(clzEnv, clzObj, (C.jlong)(img_nativeObj), (C.jdouble)(center_x), (C.jdouble)(center_y), (C.jdouble)(axes_width), (C.jdouble)(axes_height), (C.jdouble)(angle), (C.jdouble)(startAngle), (C.jdouble)(endAngle), (C.jdouble)(color_val0), (C.jdouble)(color_val1), (C.jdouble)(color_val2), (C.jdouble)(color_val3))
}

func ImgprocNative_ellipse_3(img_nativeObj int64, box_center_x float64, box_center_y float64, box_size_width float64, box_size_height float64, box_angle float64, color_val0 float64, color_val1 float64, color_val2 float64, color_val3 float64, thickness int, lineType int) {
	C.Java_org_opencv_imgproc_Imgproc_ellipse_13(clzEnv, clzObj, (C.jlong)(img_nativeObj), (C.jdouble)(box_center_x), (C.jdouble)(box_center_y), (C.jdouble)(box_size_width), (C.jdouble)(box_size_height), (C.jdouble)(box_angle), (C.jdouble)(color_val0), (C.jdouble)(color_val1), (C.jdouble)(color_val2), (C.jdouble)(color_val3), (C.jint)(thickness), (C.jint)(lineType))
}

func ImgprocNative_ellipse_4(img_nativeObj int64, box_center_x float64, box_center_y float64, box_size_width float64, box_size_height float64, box_angle float64, color_val0 float64, color_val1 float64, color_val2 float64, color_val3 float64, thickness int) {
	C.Java_org_opencv_imgproc_Imgproc_ellipse_14(clzEnv, clzObj, (C.jlong)(img_nativeObj), (C.jdouble)(box_center_x), (C.jdouble)(box_center_y), (C.jdouble)(box_size_width), (C.jdouble)(box_size_height), (C.jdouble)(box_angle), (C.jdouble)(color_val0), (C.jdouble)(color_val1), (C.jdouble)(color_val2), (C.jdouble)(color_val3), (C.jint)(thickness))
}

func ImgprocNative_ellipse_5(img_nativeObj int64, box_center_x float64, box_center_y float64, box_size_width float64, box_size_height float64, box_angle float64, color_val0 float64, color_val1 float64, color_val2 float64, color_val3 float64) {
	C.Java_org_opencv_imgproc_Imgproc_ellipse_15(clzEnv, clzObj, (C.jlong)(img_nativeObj), (C.jdouble)(box_center_x), (C.jdouble)(box_center_y), (C.jdouble)(box_size_width), (C.jdouble)(box_size_height), (C.jdouble)(box_angle), (C.jdouble)(color_val0), (C.jdouble)(color_val1), (C.jdouble)(color_val2), (C.jdouble)(color_val3))
}

func ImgprocNative_equalizeHist_0(src_nativeObj int64, dst_nativeObj int64) {
	C.Java_org_opencv_imgproc_Imgproc_equalizeHist_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj))
}

func ImgprocNative_erode_0(src_nativeObj int64, dst_nativeObj int64, kernel_nativeObj int64, anchor_x float64, anchor_y float64, iterations int, borderType int, borderValue_val0 float64, borderValue_val1 float64, borderValue_val2 float64, borderValue_val3 float64) {
	C.Java_org_opencv_imgproc_Imgproc_erode_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jlong)(kernel_nativeObj), (C.jdouble)(anchor_x), (C.jdouble)(anchor_y), (C.jint)(iterations), (C.jint)(borderType), (C.jdouble)(borderValue_val0), (C.jdouble)(borderValue_val1), (C.jdouble)(borderValue_val2), (C.jdouble)(borderValue_val3))
}

func ImgprocNative_erode_1(src_nativeObj int64, dst_nativeObj int64, kernel_nativeObj int64, anchor_x float64, anchor_y float64, iterations int) {
	C.Java_org_opencv_imgproc_Imgproc_erode_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jlong)(kernel_nativeObj), (C.jdouble)(anchor_x), (C.jdouble)(anchor_y), (C.jint)(iterations))
}

func ImgprocNative_erode_2(src_nativeObj int64, dst_nativeObj int64, kernel_nativeObj int64) {
	C.Java_org_opencv_imgproc_Imgproc_erode_12(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jlong)(kernel_nativeObj))
}

func ImgprocNative_fillConvexPoly_0(img_nativeObj int64, points_mat_nativeObj int64, color_val0 float64, color_val1 float64, color_val2 float64, color_val3 float64, lineType int, shift int) {
	C.Java_org_opencv_imgproc_Imgproc_fillConvexPoly_10(clzEnv, clzObj, (C.jlong)(img_nativeObj), (C.jlong)(points_mat_nativeObj), (C.jdouble)(color_val0), (C.jdouble)(color_val1), (C.jdouble)(color_val2), (C.jdouble)(color_val3), (C.jint)(lineType), (C.jint)(shift))
}

func ImgprocNative_fillConvexPoly_1(img_nativeObj int64, points_mat_nativeObj int64, color_val0 float64, color_val1 float64, color_val2 float64, color_val3 float64) {
	C.Java_org_opencv_imgproc_Imgproc_fillConvexPoly_11(clzEnv, clzObj, (C.jlong)(img_nativeObj), (C.jlong)(points_mat_nativeObj), (C.jdouble)(color_val0), (C.jdouble)(color_val1), (C.jdouble)(color_val2), (C.jdouble)(color_val3))
}

func ImgprocNative_fillPoly_0(img_nativeObj int64, pts_mat_nativeObj int64, color_val0 float64, color_val1 float64, color_val2 float64, color_val3 float64, lineType int, shift int, offset_x float64, offset_y float64) {
	C.Java_org_opencv_imgproc_Imgproc_fillPoly_10(clzEnv, clzObj, (C.jlong)(img_nativeObj), (C.jlong)(pts_mat_nativeObj), (C.jdouble)(color_val0), (C.jdouble)(color_val1), (C.jdouble)(color_val2), (C.jdouble)(color_val3), (C.jint)(lineType), (C.jint)(shift), (C.jdouble)(offset_x), (C.jdouble)(offset_y))
}

func ImgprocNative_fillPoly_1(img_nativeObj int64, pts_mat_nativeObj int64, color_val0 float64, color_val1 float64, color_val2 float64, color_val3 float64) {
	C.Java_org_opencv_imgproc_Imgproc_fillPoly_11(clzEnv, clzObj, (C.jlong)(img_nativeObj), (C.jlong)(pts_mat_nativeObj), (C.jdouble)(color_val0), (C.jdouble)(color_val1), (C.jdouble)(color_val2), (C.jdouble)(color_val3))
}

func ImgprocNative_filter2D_0(src_nativeObj int64, dst_nativeObj int64, ddepth int, kernel_nativeObj int64, anchor_x float64, anchor_y float64, delta float64, borderType int) {
	C.Java_org_opencv_imgproc_Imgproc_filter2D_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(ddepth), (C.jlong)(kernel_nativeObj), (C.jdouble)(anchor_x), (C.jdouble)(anchor_y), (C.jdouble)(delta), (C.jint)(borderType))
}

func ImgprocNative_filter2D_1(src_nativeObj int64, dst_nativeObj int64, ddepth int, kernel_nativeObj int64, anchor_x float64, anchor_y float64, delta float64) {
	C.Java_org_opencv_imgproc_Imgproc_filter2D_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(ddepth), (C.jlong)(kernel_nativeObj), (C.jdouble)(anchor_x), (C.jdouble)(anchor_y), (C.jdouble)(delta))
}

func ImgprocNative_filter2D_2(src_nativeObj int64, dst_nativeObj int64, ddepth int, kernel_nativeObj int64) {
	C.Java_org_opencv_imgproc_Imgproc_filter2D_12(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(ddepth), (C.jlong)(kernel_nativeObj))
}

func ImgprocNative_findContours_0(image_nativeObj int64, contours_mat_nativeObj int64, hierarchy_nativeObj int64, mode int, method int, offset_x float64, offset_y float64) {
	C.Java_org_opencv_imgproc_Imgproc_findContours_10(clzEnv, clzObj, (C.jlong)(image_nativeObj), (C.jlong)(contours_mat_nativeObj), (C.jlong)(hierarchy_nativeObj), (C.jint)(mode), (C.jint)(method), (C.jdouble)(offset_x), (C.jdouble)(offset_y))
}

func ImgprocNative_findContours_1(image_nativeObj int64, contours_mat_nativeObj int64, hierarchy_nativeObj int64, mode int, method int) {
	C.Java_org_opencv_imgproc_Imgproc_findContours_11(clzEnv, clzObj, (C.jlong)(image_nativeObj), (C.jlong)(contours_mat_nativeObj), (C.jlong)(hierarchy_nativeObj), (C.jint)(mode), (C.jint)(method))
}

func ImgprocNative_fitEllipseAMS_0(points_nativeObj int64) []float64 {
	return togofloat64Array(C.Java_org_opencv_imgproc_Imgproc_fitEllipseAMS_10(clzEnv, clzObj, (C.jlong)(points_nativeObj)))
}

func ImgprocNative_fitEllipseDirect_0(points_nativeObj int64) []float64 {
	return togofloat64Array(C.Java_org_opencv_imgproc_Imgproc_fitEllipseDirect_10(clzEnv, clzObj, (C.jlong)(points_nativeObj)))
}

func ImgprocNative_fitEllipse_0(points_mat_nativeObj int64) []float64 {
	return togofloat64Array(C.Java_org_opencv_imgproc_Imgproc_fitEllipse_10(clzEnv, clzObj, (C.jlong)(points_mat_nativeObj)))
}

func ImgprocNative_fitLine_0(points_nativeObj int64, line_nativeObj int64, distType int, param float64, reps float64, aeps float64) {
	C.Java_org_opencv_imgproc_Imgproc_fitLine_10(clzEnv, clzObj, (C.jlong)(points_nativeObj), (C.jlong)(line_nativeObj), (C.jint)(distType), (C.jdouble)(param), (C.jdouble)(reps), (C.jdouble)(aeps))
}

func ImgprocNative_floodFill_0(image_nativeObj int64, mask_nativeObj int64, seedPoint_x float64, seedPoint_y float64, newVal_val0 float64, newVal_val1 float64, newVal_val2 float64, newVal_val3 float64, rect_out []float64, loDiff_val0 float64, loDiff_val1 float64, loDiff_val2 float64, loDiff_val3 float64, upDiff_val0 float64, upDiff_val1 float64, upDiff_val2 float64, upDiff_val3 float64, flags int) int {
	defer ungointerface(rect_out)
	return int(C.Java_org_opencv_imgproc_Imgproc_floodFill_10(clzEnv, clzObj, (C.jlong)(image_nativeObj), (C.jlong)(mask_nativeObj), (C.jdouble)(seedPoint_x), (C.jdouble)(seedPoint_y), (C.jdouble)(newVal_val0), (C.jdouble)(newVal_val1), (C.jdouble)(newVal_val2), (C.jdouble)(newVal_val3), tojdoubleArray(rect_out), (C.jdouble)(loDiff_val0), (C.jdouble)(loDiff_val1), (C.jdouble)(loDiff_val2), (C.jdouble)(loDiff_val3), (C.jdouble)(upDiff_val0), (C.jdouble)(upDiff_val1), (C.jdouble)(upDiff_val2), (C.jdouble)(upDiff_val3), (C.jint)(flags)))
}

func ImgprocNative_floodFill_1(image_nativeObj int64, mask_nativeObj int64, seedPoint_x float64, seedPoint_y float64, newVal_val0 float64, newVal_val1 float64, newVal_val2 float64, newVal_val3 float64) int {
	return int(C.Java_org_opencv_imgproc_Imgproc_floodFill_11(clzEnv, clzObj, (C.jlong)(image_nativeObj), (C.jlong)(mask_nativeObj), (C.jdouble)(seedPoint_x), (C.jdouble)(seedPoint_y), (C.jdouble)(newVal_val0), (C.jdouble)(newVal_val1), (C.jdouble)(newVal_val2), (C.jdouble)(newVal_val3)))
}

func ImgprocNative_getAffineTransform_0(src_mat_nativeObj int64, dst_mat_nativeObj int64) int64 {
	return int64(C.Java_org_opencv_imgproc_Imgproc_getAffineTransform_10(clzEnv, clzObj, (C.jlong)(src_mat_nativeObj), (C.jlong)(dst_mat_nativeObj)))
}

func ImgprocNative_getDefaultNewCameraMatrix_0(cameraMatrix_nativeObj int64, imgsize_width float64, imgsize_height float64, centerPrincipalPoint bool) int64 {
	return int64(C.Java_org_opencv_imgproc_Imgproc_getDefaultNewCameraMatrix_10(clzEnv, clzObj, (C.jlong)(cameraMatrix_nativeObj), (C.jdouble)(imgsize_width), (C.jdouble)(imgsize_height), tojboolean(centerPrincipalPoint)))
}

func ImgprocNative_getDefaultNewCameraMatrix_1(cameraMatrix_nativeObj int64) int64 {
	return int64(C.Java_org_opencv_imgproc_Imgproc_getDefaultNewCameraMatrix_11(clzEnv, clzObj, (C.jlong)(cameraMatrix_nativeObj)))
}

func ImgprocNative_getDerivKernels_0(kx_nativeObj int64, ky_nativeObj int64, dx int, dy int, ksize int, normalize bool, ktype int) {
	C.Java_org_opencv_imgproc_Imgproc_getDerivKernels_10(clzEnv, clzObj, (C.jlong)(kx_nativeObj), (C.jlong)(ky_nativeObj), (C.jint)(dx), (C.jint)(dy), (C.jint)(ksize), tojboolean(normalize), (C.jint)(ktype))
}

func ImgprocNative_getDerivKernels_1(kx_nativeObj int64, ky_nativeObj int64, dx int, dy int, ksize int) {
	C.Java_org_opencv_imgproc_Imgproc_getDerivKernels_11(clzEnv, clzObj, (C.jlong)(kx_nativeObj), (C.jlong)(ky_nativeObj), (C.jint)(dx), (C.jint)(dy), (C.jint)(ksize))
}

func ImgprocNative_getGaborKernel_0(ksize_width float64, ksize_height float64, sigma float64, theta float64, lambd float64, gamma float64, psi float64, ktype int) int64 {
	return int64(C.Java_org_opencv_imgproc_Imgproc_getGaborKernel_10(clzEnv, clzObj, (C.jdouble)(ksize_width), (C.jdouble)(ksize_height), (C.jdouble)(sigma), (C.jdouble)(theta), (C.jdouble)(lambd), (C.jdouble)(gamma), (C.jdouble)(psi), (C.jint)(ktype)))
}

func ImgprocNative_getGaborKernel_1(ksize_width float64, ksize_height float64, sigma float64, theta float64, lambd float64, gamma float64) int64 {
	return int64(C.Java_org_opencv_imgproc_Imgproc_getGaborKernel_11(clzEnv, clzObj, (C.jdouble)(ksize_width), (C.jdouble)(ksize_height), (C.jdouble)(sigma), (C.jdouble)(theta), (C.jdouble)(lambd), (C.jdouble)(gamma)))
}

func ImgprocNative_getGaussianKernel_0(ksize int, sigma float64, ktype int) int64 {
	return int64(C.Java_org_opencv_imgproc_Imgproc_getGaussianKernel_10(clzEnv, clzObj, (C.jint)(ksize), (C.jdouble)(sigma), (C.jint)(ktype)))
}

func ImgprocNative_getGaussianKernel_1(ksize int, sigma float64) int64 {
	return int64(C.Java_org_opencv_imgproc_Imgproc_getGaussianKernel_11(clzEnv, clzObj, (C.jint)(ksize), (C.jdouble)(sigma)))
}

func ImgprocNative_getPerspectiveTransform_0(src_nativeObj int64, dst_nativeObj int64) int64 {
	return int64(C.Java_org_opencv_imgproc_Imgproc_getPerspectiveTransform_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj)))
}

func ImgprocNative_getRectSubPix_0(image_nativeObj int64, patchSize_width float64, patchSize_height float64, center_x float64, center_y float64, patch_nativeObj int64, patchType int) {
	C.Java_org_opencv_imgproc_Imgproc_getRectSubPix_10(clzEnv, clzObj, (C.jlong)(image_nativeObj), (C.jdouble)(patchSize_width), (C.jdouble)(patchSize_height), (C.jdouble)(center_x), (C.jdouble)(center_y), (C.jlong)(patch_nativeObj), (C.jint)(patchType))
}

func ImgprocNative_getRectSubPix_1(image_nativeObj int64, patchSize_width float64, patchSize_height float64, center_x float64, center_y float64, patch_nativeObj int64) {
	C.Java_org_opencv_imgproc_Imgproc_getRectSubPix_11(clzEnv, clzObj, (C.jlong)(image_nativeObj), (C.jdouble)(patchSize_width), (C.jdouble)(patchSize_height), (C.jdouble)(center_x), (C.jdouble)(center_y), (C.jlong)(patch_nativeObj))
}

func ImgprocNative_getRotationMatrix2D_0(center_x float64, center_y float64, angle float64, scale float64) int64 {
	return int64(C.Java_org_opencv_imgproc_Imgproc_getRotationMatrix2D_10(clzEnv, clzObj, (C.jdouble)(center_x), (C.jdouble)(center_y), (C.jdouble)(angle), (C.jdouble)(scale)))
}

func ImgprocNative_getStructuringElement_0(shape int, ksize_width float64, ksize_height float64, anchor_x float64, anchor_y float64) int64 {
	return int64(C.Java_org_opencv_imgproc_Imgproc_getStructuringElement_10(clzEnv, clzObj, (C.jint)(shape), (C.jdouble)(ksize_width), (C.jdouble)(ksize_height), (C.jdouble)(anchor_x), (C.jdouble)(anchor_y)))
}

func ImgprocNative_getStructuringElement_1(shape int, ksize_width float64, ksize_height float64) int64 {
	return int64(C.Java_org_opencv_imgproc_Imgproc_getStructuringElement_11(clzEnv, clzObj, (C.jint)(shape), (C.jdouble)(ksize_width), (C.jdouble)(ksize_height)))
}

func ImgprocNative_goodFeaturesToTrack_0(image_nativeObj int64, corners_mat_nativeObj int64, maxCorners int, qualityLevel float64, minDistance float64, mask_nativeObj int64, blockSize int, gradientSize int, useHarrisDetector bool, k float64) {
	C.Java_org_opencv_imgproc_Imgproc_goodFeaturesToTrack_10(clzEnv, clzObj, (C.jlong)(image_nativeObj), (C.jlong)(corners_mat_nativeObj), (C.jint)(maxCorners), (C.jdouble)(qualityLevel), (C.jdouble)(minDistance), (C.jlong)(mask_nativeObj), (C.jint)(blockSize), (C.jint)(gradientSize), tojboolean(useHarrisDetector), (C.jdouble)(k))
}

func ImgprocNative_goodFeaturesToTrack_1(image_nativeObj int64, corners_mat_nativeObj int64, maxCorners int, qualityLevel float64, minDistance float64, mask_nativeObj int64, blockSize int, gradientSize int) {
	C.Java_org_opencv_imgproc_Imgproc_goodFeaturesToTrack_11(clzEnv, clzObj, (C.jlong)(image_nativeObj), (C.jlong)(corners_mat_nativeObj), (C.jint)(maxCorners), (C.jdouble)(qualityLevel), (C.jdouble)(minDistance), (C.jlong)(mask_nativeObj), (C.jint)(blockSize), (C.jint)(gradientSize))
}

func ImgprocNative_goodFeaturesToTrack_2(image_nativeObj int64, corners_mat_nativeObj int64, maxCorners int, qualityLevel float64, minDistance float64, mask_nativeObj int64, blockSize int, useHarrisDetector bool, k float64) {
	C.Java_org_opencv_imgproc_Imgproc_goodFeaturesToTrack_12(clzEnv, clzObj, (C.jlong)(image_nativeObj), (C.jlong)(corners_mat_nativeObj), (C.jint)(maxCorners), (C.jdouble)(qualityLevel), (C.jdouble)(minDistance), (C.jlong)(mask_nativeObj), (C.jint)(blockSize), tojboolean(useHarrisDetector), (C.jdouble)(k))
}

func ImgprocNative_goodFeaturesToTrack_3(image_nativeObj int64, corners_mat_nativeObj int64, maxCorners int, qualityLevel float64, minDistance float64) {
	C.Java_org_opencv_imgproc_Imgproc_goodFeaturesToTrack_13(clzEnv, clzObj, (C.jlong)(image_nativeObj), (C.jlong)(corners_mat_nativeObj), (C.jint)(maxCorners), (C.jdouble)(qualityLevel), (C.jdouble)(minDistance))
}

func ImgprocNative_grabCut_0(img_nativeObj int64, mask_nativeObj int64, rect_x int, rect_y int, rect_width int, rect_height int, bgdModel_nativeObj int64, fgdModel_nativeObj int64, iterCount int, mode int) {
	C.Java_org_opencv_imgproc_Imgproc_grabCut_10(clzEnv, clzObj, (C.jlong)(img_nativeObj), (C.jlong)(mask_nativeObj), (C.jint)(rect_x), (C.jint)(rect_y), (C.jint)(rect_width), (C.jint)(rect_height), (C.jlong)(bgdModel_nativeObj), (C.jlong)(fgdModel_nativeObj), (C.jint)(iterCount), (C.jint)(mode))
}

func ImgprocNative_grabCut_1(img_nativeObj int64, mask_nativeObj int64, rect_x int, rect_y int, rect_width int, rect_height int, bgdModel_nativeObj int64, fgdModel_nativeObj int64, iterCount int) {
	C.Java_org_opencv_imgproc_Imgproc_grabCut_11(clzEnv, clzObj, (C.jlong)(img_nativeObj), (C.jlong)(mask_nativeObj), (C.jint)(rect_x), (C.jint)(rect_y), (C.jint)(rect_width), (C.jint)(rect_height), (C.jlong)(bgdModel_nativeObj), (C.jlong)(fgdModel_nativeObj), (C.jint)(iterCount))
}

func ImgprocNative_initUndistortRectifyMap_0(cameraMatrix_nativeObj int64, distCoeffs_nativeObj int64, R_nativeObj int64, newCameraMatrix_nativeObj int64, size_width float64, size_height float64, m1type int, map1_nativeObj int64, map2_nativeObj int64) {
	C.Java_org_opencv_imgproc_Imgproc_initUndistortRectifyMap_10(clzEnv, clzObj, (C.jlong)(cameraMatrix_nativeObj), (C.jlong)(distCoeffs_nativeObj), (C.jlong)(R_nativeObj), (C.jlong)(newCameraMatrix_nativeObj), (C.jdouble)(size_width), (C.jdouble)(size_height), (C.jint)(m1type), (C.jlong)(map1_nativeObj), (C.jlong)(map2_nativeObj))
}

func ImgprocNative_initWideAngleProjMap_0(cameraMatrix_nativeObj int64, distCoeffs_nativeObj int64, imageSize_width float64, imageSize_height float64, destImageWidth int, m1type int, map1_nativeObj int64, map2_nativeObj int64, projType int, alpha float64) float32 {
	return float32(C.Java_org_opencv_imgproc_Imgproc_initWideAngleProjMap_10(clzEnv, clzObj, (C.jlong)(cameraMatrix_nativeObj), (C.jlong)(distCoeffs_nativeObj), (C.jdouble)(imageSize_width), (C.jdouble)(imageSize_height), (C.jint)(destImageWidth), (C.jint)(m1type), (C.jlong)(map1_nativeObj), (C.jlong)(map2_nativeObj), (C.jint)(projType), (C.jdouble)(alpha)))
}

func ImgprocNative_initWideAngleProjMap_1(cameraMatrix_nativeObj int64, distCoeffs_nativeObj int64, imageSize_width float64, imageSize_height float64, destImageWidth int, m1type int, map1_nativeObj int64, map2_nativeObj int64) float32 {
	return float32(C.Java_org_opencv_imgproc_Imgproc_initWideAngleProjMap_11(clzEnv, clzObj, (C.jlong)(cameraMatrix_nativeObj), (C.jlong)(distCoeffs_nativeObj), (C.jdouble)(imageSize_width), (C.jdouble)(imageSize_height), (C.jint)(destImageWidth), (C.jint)(m1type), (C.jlong)(map1_nativeObj), (C.jlong)(map2_nativeObj)))
}

func ImgprocNative_integral2_0(src_nativeObj int64, sum_nativeObj int64, sqsum_nativeObj int64, sdepth int, sqdepth int) {
	C.Java_org_opencv_imgproc_Imgproc_integral2_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(sum_nativeObj), (C.jlong)(sqsum_nativeObj), (C.jint)(sdepth), (C.jint)(sqdepth))
}

func ImgprocNative_integral2_1(src_nativeObj int64, sum_nativeObj int64, sqsum_nativeObj int64) {
	C.Java_org_opencv_imgproc_Imgproc_integral2_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(sum_nativeObj), (C.jlong)(sqsum_nativeObj))
}

func ImgprocNative_integral3_0(src_nativeObj int64, sum_nativeObj int64, sqsum_nativeObj int64, tilted_nativeObj int64, sdepth int, sqdepth int) {
	C.Java_org_opencv_imgproc_Imgproc_integral3_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(sum_nativeObj), (C.jlong)(sqsum_nativeObj), (C.jlong)(tilted_nativeObj), (C.jint)(sdepth), (C.jint)(sqdepth))
}

func ImgprocNative_integral3_1(src_nativeObj int64, sum_nativeObj int64, sqsum_nativeObj int64, tilted_nativeObj int64) {
	C.Java_org_opencv_imgproc_Imgproc_integral3_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(sum_nativeObj), (C.jlong)(sqsum_nativeObj), (C.jlong)(tilted_nativeObj))
}

func ImgprocNative_integral_0(src_nativeObj int64, sum_nativeObj int64, sdepth int) {
	C.Java_org_opencv_imgproc_Imgproc_integral_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(sum_nativeObj), (C.jint)(sdepth))
}

func ImgprocNative_integral_1(src_nativeObj int64, sum_nativeObj int64) {
	C.Java_org_opencv_imgproc_Imgproc_integral_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(sum_nativeObj))
}

func ImgprocNative_intersectConvexConvex_0(_p1_nativeObj int64, _p2_nativeObj int64, _p12_nativeObj int64, handleNested bool) float32 {
	return float32(C.Java_org_opencv_imgproc_Imgproc_intersectConvexConvex_10(clzEnv, clzObj, (C.jlong)(_p1_nativeObj), (C.jlong)(_p2_nativeObj), (C.jlong)(_p12_nativeObj), tojboolean(handleNested)))
}

func ImgprocNative_intersectConvexConvex_1(_p1_nativeObj int64, _p2_nativeObj int64, _p12_nativeObj int64) float32 {
	return float32(C.Java_org_opencv_imgproc_Imgproc_intersectConvexConvex_11(clzEnv, clzObj, (C.jlong)(_p1_nativeObj), (C.jlong)(_p2_nativeObj), (C.jlong)(_p12_nativeObj)))
}

func ImgprocNative_invertAffineTransform_0(M_nativeObj int64, iM_nativeObj int64) {
	C.Java_org_opencv_imgproc_Imgproc_invertAffineTransform_10(clzEnv, clzObj, (C.jlong)(M_nativeObj), (C.jlong)(iM_nativeObj))
}

func ImgprocNative_isContourConvex_0(contour_mat_nativeObj int64) bool {
	return togobool(C.Java_org_opencv_imgproc_Imgproc_isContourConvex_10(clzEnv, clzObj, (C.jlong)(contour_mat_nativeObj)))
}

func ImgprocNative_line_0(img_nativeObj int64, pt1_x float64, pt1_y float64, pt2_x float64, pt2_y float64, color_val0 float64, color_val1 float64, color_val2 float64, color_val3 float64, thickness int, lineType int, shift int) {
	C.Java_org_opencv_imgproc_Imgproc_line_10(clzEnv, clzObj, (C.jlong)(img_nativeObj), (C.jdouble)(pt1_x), (C.jdouble)(pt1_y), (C.jdouble)(pt2_x), (C.jdouble)(pt2_y), (C.jdouble)(color_val0), (C.jdouble)(color_val1), (C.jdouble)(color_val2), (C.jdouble)(color_val3), (C.jint)(thickness), (C.jint)(lineType), (C.jint)(shift))
}

func ImgprocNative_line_1(img_nativeObj int64, pt1_x float64, pt1_y float64, pt2_x float64, pt2_y float64, color_val0 float64, color_val1 float64, color_val2 float64, color_val3 float64, thickness int) {
	C.Java_org_opencv_imgproc_Imgproc_line_11(clzEnv, clzObj, (C.jlong)(img_nativeObj), (C.jdouble)(pt1_x), (C.jdouble)(pt1_y), (C.jdouble)(pt2_x), (C.jdouble)(pt2_y), (C.jdouble)(color_val0), (C.jdouble)(color_val1), (C.jdouble)(color_val2), (C.jdouble)(color_val3), (C.jint)(thickness))
}

func ImgprocNative_line_2(img_nativeObj int64, pt1_x float64, pt1_y float64, pt2_x float64, pt2_y float64, color_val0 float64, color_val1 float64, color_val2 float64, color_val3 float64) {
	C.Java_org_opencv_imgproc_Imgproc_line_12(clzEnv, clzObj, (C.jlong)(img_nativeObj), (C.jdouble)(pt1_x), (C.jdouble)(pt1_y), (C.jdouble)(pt2_x), (C.jdouble)(pt2_y), (C.jdouble)(color_val0), (C.jdouble)(color_val1), (C.jdouble)(color_val2), (C.jdouble)(color_val3))
}

func ImgprocNative_linearPolar_0(src_nativeObj int64, dst_nativeObj int64, center_x float64, center_y float64, maxRadius float64, flags int) {
	C.Java_org_opencv_imgproc_Imgproc_linearPolar_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jdouble)(center_x), (C.jdouble)(center_y), (C.jdouble)(maxRadius), (C.jint)(flags))
}

func ImgprocNative_logPolar_0(src_nativeObj int64, dst_nativeObj int64, center_x float64, center_y float64, M float64, flags int) {
	C.Java_org_opencv_imgproc_Imgproc_logPolar_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jdouble)(center_x), (C.jdouble)(center_y), (C.jdouble)(M), (C.jint)(flags))
}

func ImgprocNative_matchShapes_0(contour1_nativeObj int64, contour2_nativeObj int64, method int, parameter float64) float64 {
	return float64(C.Java_org_opencv_imgproc_Imgproc_matchShapes_10(clzEnv, clzObj, (C.jlong)(contour1_nativeObj), (C.jlong)(contour2_nativeObj), (C.jint)(method), (C.jdouble)(parameter)))
}

func ImgprocNative_matchTemplate_0(image_nativeObj int64, templ_nativeObj int64, result_nativeObj int64, method int, mask_nativeObj int64) {
	C.Java_org_opencv_imgproc_Imgproc_matchTemplate_10(clzEnv, clzObj, (C.jlong)(image_nativeObj), (C.jlong)(templ_nativeObj), (C.jlong)(result_nativeObj), (C.jint)(method), (C.jlong)(mask_nativeObj))
}

func ImgprocNative_matchTemplate_1(image_nativeObj int64, templ_nativeObj int64, result_nativeObj int64, method int) {
	C.Java_org_opencv_imgproc_Imgproc_matchTemplate_11(clzEnv, clzObj, (C.jlong)(image_nativeObj), (C.jlong)(templ_nativeObj), (C.jlong)(result_nativeObj), (C.jint)(method))
}

func ImgprocNative_medianBlur_0(src_nativeObj int64, dst_nativeObj int64, ksize int) {
	C.Java_org_opencv_imgproc_Imgproc_medianBlur_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(ksize))
}

func ImgprocNative_minAreaRect_0(points_mat_nativeObj int64) []float64 {
	return togofloat64Array(C.Java_org_opencv_imgproc_Imgproc_minAreaRect_10(clzEnv, clzObj, (C.jlong)(points_mat_nativeObj)))
}

func ImgprocNative_minEnclosingCircle_0(points_mat_nativeObj int64, center_out []float64, radius_out []float64) {
	defer ungointerface(center_out)
	defer ungointerface(radius_out)
	C.Java_org_opencv_imgproc_Imgproc_minEnclosingCircle_10(clzEnv, clzObj, (C.jlong)(points_mat_nativeObj), tojdoubleArray(center_out), tojdoubleArray(radius_out))
}

func ImgprocNative_minEnclosingTriangle_0(points_nativeObj int64, triangle_nativeObj int64) float64 {
	return float64(C.Java_org_opencv_imgproc_Imgproc_minEnclosingTriangle_10(clzEnv, clzObj, (C.jlong)(points_nativeObj), (C.jlong)(triangle_nativeObj)))
}

func ImgprocNative_moments_0(array_nativeObj int64, binaryImage bool) []float64 {
	return togofloat64Array(C.Java_org_opencv_imgproc_Imgproc_moments_10(clzEnv, clzObj, (C.jlong)(array_nativeObj), tojboolean(binaryImage)))
}

func ImgprocNative_moments_1(array_nativeObj int64) []float64 {
	return togofloat64Array(C.Java_org_opencv_imgproc_Imgproc_moments_11(clzEnv, clzObj, (C.jlong)(array_nativeObj)))
}

func ImgprocNative_morphologyEx_0(src_nativeObj int64, dst_nativeObj int64, op int, kernel_nativeObj int64, anchor_x float64, anchor_y float64, iterations int, borderType int, borderValue_val0 float64, borderValue_val1 float64, borderValue_val2 float64, borderValue_val3 float64) {
	C.Java_org_opencv_imgproc_Imgproc_morphologyEx_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(op), (C.jlong)(kernel_nativeObj), (C.jdouble)(anchor_x), (C.jdouble)(anchor_y), (C.jint)(iterations), (C.jint)(borderType), (C.jdouble)(borderValue_val0), (C.jdouble)(borderValue_val1), (C.jdouble)(borderValue_val2), (C.jdouble)(borderValue_val3))
}

func ImgprocNative_morphologyEx_1(src_nativeObj int64, dst_nativeObj int64, op int, kernel_nativeObj int64, anchor_x float64, anchor_y float64, iterations int) {
	C.Java_org_opencv_imgproc_Imgproc_morphologyEx_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(op), (C.jlong)(kernel_nativeObj), (C.jdouble)(anchor_x), (C.jdouble)(anchor_y), (C.jint)(iterations))
}

func ImgprocNative_morphologyEx_2(src_nativeObj int64, dst_nativeObj int64, op int, kernel_nativeObj int64) {
	C.Java_org_opencv_imgproc_Imgproc_morphologyEx_12(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(op), (C.jlong)(kernel_nativeObj))
}

func ImgprocNative_n_getTextSize(text string, fontFace int, fontScale float64, thickness int, baseLine []int32) []float64 {
	defer ungointerface(baseLine)
	return togofloat64Array(C.Java_org_opencv_imgproc_Imgproc_n_1getTextSize(clzEnv, clzObj, tojstring(text), (C.jint)(fontFace), (C.jdouble)(fontScale), (C.jint)(thickness), tojintArray(baseLine)))
}

func ImgprocNative_phaseCorrelate_0(src1_nativeObj int64, src2_nativeObj int64, window_nativeObj int64, response_out []float64) []float64 {
	defer ungointerface(response_out)
	return togofloat64Array(C.Java_org_opencv_imgproc_Imgproc_phaseCorrelate_10(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jlong)(src2_nativeObj), (C.jlong)(window_nativeObj), tojdoubleArray(response_out)))
}

func ImgprocNative_phaseCorrelate_1(src1_nativeObj int64, src2_nativeObj int64) []float64 {
	return togofloat64Array(C.Java_org_opencv_imgproc_Imgproc_phaseCorrelate_11(clzEnv, clzObj, (C.jlong)(src1_nativeObj), (C.jlong)(src2_nativeObj)))
}

func ImgprocNative_pointPolygonTest_0(contour_mat_nativeObj int64, pt_x float64, pt_y float64, measureDist bool) float64 {
	return float64(C.Java_org_opencv_imgproc_Imgproc_pointPolygonTest_10(clzEnv, clzObj, (C.jlong)(contour_mat_nativeObj), (C.jdouble)(pt_x), (C.jdouble)(pt_y), tojboolean(measureDist)))
}

func ImgprocNative_polylines_0(img_nativeObj int64, pts_mat_nativeObj int64, isClosed bool, color_val0 float64, color_val1 float64, color_val2 float64, color_val3 float64, thickness int, lineType int, shift int) {
	C.Java_org_opencv_imgproc_Imgproc_polylines_10(clzEnv, clzObj, (C.jlong)(img_nativeObj), (C.jlong)(pts_mat_nativeObj), tojboolean(isClosed), (C.jdouble)(color_val0), (C.jdouble)(color_val1), (C.jdouble)(color_val2), (C.jdouble)(color_val3), (C.jint)(thickness), (C.jint)(lineType), (C.jint)(shift))
}

func ImgprocNative_polylines_1(img_nativeObj int64, pts_mat_nativeObj int64, isClosed bool, color_val0 float64, color_val1 float64, color_val2 float64, color_val3 float64, thickness int) {
	C.Java_org_opencv_imgproc_Imgproc_polylines_11(clzEnv, clzObj, (C.jlong)(img_nativeObj), (C.jlong)(pts_mat_nativeObj), tojboolean(isClosed), (C.jdouble)(color_val0), (C.jdouble)(color_val1), (C.jdouble)(color_val2), (C.jdouble)(color_val3), (C.jint)(thickness))
}

func ImgprocNative_polylines_2(img_nativeObj int64, pts_mat_nativeObj int64, isClosed bool, color_val0 float64, color_val1 float64, color_val2 float64, color_val3 float64) {
	C.Java_org_opencv_imgproc_Imgproc_polylines_12(clzEnv, clzObj, (C.jlong)(img_nativeObj), (C.jlong)(pts_mat_nativeObj), tojboolean(isClosed), (C.jdouble)(color_val0), (C.jdouble)(color_val1), (C.jdouble)(color_val2), (C.jdouble)(color_val3))
}

func ImgprocNative_preCornerDetect_0(src_nativeObj int64, dst_nativeObj int64, ksize int, borderType int) {
	C.Java_org_opencv_imgproc_Imgproc_preCornerDetect_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(ksize), (C.jint)(borderType))
}

func ImgprocNative_preCornerDetect_1(src_nativeObj int64, dst_nativeObj int64, ksize int) {
	C.Java_org_opencv_imgproc_Imgproc_preCornerDetect_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(ksize))
}

func ImgprocNative_putText_0(img_nativeObj int64, text string, org_x float64, org_y float64, fontFace int, fontScale float64, color_val0 float64, color_val1 float64, color_val2 float64, color_val3 float64, thickness int, lineType int, bottomLeftOrigin bool) {
	defer ungointerface(text)
	C.Java_org_opencv_imgproc_Imgproc_putText_10(clzEnv, clzObj, (C.jlong)(img_nativeObj), tojstring(text), (C.jdouble)(org_x), (C.jdouble)(org_y), (C.jint)(fontFace), (C.jdouble)(fontScale), (C.jdouble)(color_val0), (C.jdouble)(color_val1), (C.jdouble)(color_val2), (C.jdouble)(color_val3), (C.jint)(thickness), (C.jint)(lineType), tojboolean(bottomLeftOrigin))
}

func ImgprocNative_putText_1(img_nativeObj int64, text string, org_x float64, org_y float64, fontFace int, fontScale float64, color_val0 float64, color_val1 float64, color_val2 float64, color_val3 float64, thickness int) {
	defer ungointerface(text)
	C.Java_org_opencv_imgproc_Imgproc_putText_11(clzEnv, clzObj, (C.jlong)(img_nativeObj), tojstring(text), (C.jdouble)(org_x), (C.jdouble)(org_y), (C.jint)(fontFace), (C.jdouble)(fontScale), (C.jdouble)(color_val0), (C.jdouble)(color_val1), (C.jdouble)(color_val2), (C.jdouble)(color_val3), (C.jint)(thickness))
}

func ImgprocNative_putText_2(img_nativeObj int64, text string, org_x float64, org_y float64, fontFace int, fontScale float64, color_val0 float64, color_val1 float64, color_val2 float64, color_val3 float64) {
	defer ungointerface(text)
	C.Java_org_opencv_imgproc_Imgproc_putText_12(clzEnv, clzObj, (C.jlong)(img_nativeObj), tojstring(text), (C.jdouble)(org_x), (C.jdouble)(org_y), (C.jint)(fontFace), (C.jdouble)(fontScale), (C.jdouble)(color_val0), (C.jdouble)(color_val1), (C.jdouble)(color_val2), (C.jdouble)(color_val3))
}

func ImgprocNative_pyrDown_0(src_nativeObj int64, dst_nativeObj int64, dstsize_width float64, dstsize_height float64, borderType int) {
	C.Java_org_opencv_imgproc_Imgproc_pyrDown_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jdouble)(dstsize_width), (C.jdouble)(dstsize_height), (C.jint)(borderType))
}

func ImgprocNative_pyrDown_1(src_nativeObj int64, dst_nativeObj int64, dstsize_width float64, dstsize_height float64) {
	C.Java_org_opencv_imgproc_Imgproc_pyrDown_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jdouble)(dstsize_width), (C.jdouble)(dstsize_height))
}

func ImgprocNative_pyrDown_2(src_nativeObj int64, dst_nativeObj int64) {
	C.Java_org_opencv_imgproc_Imgproc_pyrDown_12(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj))
}

func ImgprocNative_pyrMeanShiftFiltering_0(src_nativeObj int64, dst_nativeObj int64, sp float64, sr float64, maxLevel int, termcrit_type int, termcrit_maxCount int, termcrit_epsilon float64) {
	C.Java_org_opencv_imgproc_Imgproc_pyrMeanShiftFiltering_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jdouble)(sp), (C.jdouble)(sr), (C.jint)(maxLevel), (C.jint)(termcrit_type), (C.jint)(termcrit_maxCount), (C.jdouble)(termcrit_epsilon))
}

func ImgprocNative_pyrMeanShiftFiltering_1(src_nativeObj int64, dst_nativeObj int64, sp float64, sr float64) {
	C.Java_org_opencv_imgproc_Imgproc_pyrMeanShiftFiltering_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jdouble)(sp), (C.jdouble)(sr))
}

func ImgprocNative_pyrUp_0(src_nativeObj int64, dst_nativeObj int64, dstsize_width float64, dstsize_height float64, borderType int) {
	C.Java_org_opencv_imgproc_Imgproc_pyrUp_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jdouble)(dstsize_width), (C.jdouble)(dstsize_height), (C.jint)(borderType))
}

func ImgprocNative_pyrUp_1(src_nativeObj int64, dst_nativeObj int64, dstsize_width float64, dstsize_height float64) {
	C.Java_org_opencv_imgproc_Imgproc_pyrUp_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jdouble)(dstsize_width), (C.jdouble)(dstsize_height))
}

func ImgprocNative_pyrUp_2(src_nativeObj int64, dst_nativeObj int64) {
	C.Java_org_opencv_imgproc_Imgproc_pyrUp_12(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj))
}

func ImgprocNative_rectangle_0(img_nativeObj int64, pt1_x float64, pt1_y float64, pt2_x float64, pt2_y float64, color_val0 float64, color_val1 float64, color_val2 float64, color_val3 float64, thickness int, lineType int, shift int) {
	C.Java_org_opencv_imgproc_Imgproc_rectangle_10(clzEnv, clzObj, (C.jlong)(img_nativeObj), (C.jdouble)(pt1_x), (C.jdouble)(pt1_y), (C.jdouble)(pt2_x), (C.jdouble)(pt2_y), (C.jdouble)(color_val0), (C.jdouble)(color_val1), (C.jdouble)(color_val2), (C.jdouble)(color_val3), (C.jint)(thickness), (C.jint)(lineType), (C.jint)(shift))
}

func ImgprocNative_rectangle_1(img_nativeObj int64, pt1_x float64, pt1_y float64, pt2_x float64, pt2_y float64, color_val0 float64, color_val1 float64, color_val2 float64, color_val3 float64, thickness int) {
	C.Java_org_opencv_imgproc_Imgproc_rectangle_11(clzEnv, clzObj, (C.jlong)(img_nativeObj), (C.jdouble)(pt1_x), (C.jdouble)(pt1_y), (C.jdouble)(pt2_x), (C.jdouble)(pt2_y), (C.jdouble)(color_val0), (C.jdouble)(color_val1), (C.jdouble)(color_val2), (C.jdouble)(color_val3), (C.jint)(thickness))
}

func ImgprocNative_rectangle_2(img_nativeObj int64, pt1_x float64, pt1_y float64, pt2_x float64, pt2_y float64, color_val0 float64, color_val1 float64, color_val2 float64, color_val3 float64) {
	C.Java_org_opencv_imgproc_Imgproc_rectangle_12(clzEnv, clzObj, (C.jlong)(img_nativeObj), (C.jdouble)(pt1_x), (C.jdouble)(pt1_y), (C.jdouble)(pt2_x), (C.jdouble)(pt2_y), (C.jdouble)(color_val0), (C.jdouble)(color_val1), (C.jdouble)(color_val2), (C.jdouble)(color_val3))
}

func ImgprocNative_remap_0(src_nativeObj int64, dst_nativeObj int64, map1_nativeObj int64, map2_nativeObj int64, interpolation int, borderMode int, borderValue_val0 float64, borderValue_val1 float64, borderValue_val2 float64, borderValue_val3 float64) {
	C.Java_org_opencv_imgproc_Imgproc_remap_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jlong)(map1_nativeObj), (C.jlong)(map2_nativeObj), (C.jint)(interpolation), (C.jint)(borderMode), (C.jdouble)(borderValue_val0), (C.jdouble)(borderValue_val1), (C.jdouble)(borderValue_val2), (C.jdouble)(borderValue_val3))
}

func ImgprocNative_remap_1(src_nativeObj int64, dst_nativeObj int64, map1_nativeObj int64, map2_nativeObj int64, interpolation int) {
	C.Java_org_opencv_imgproc_Imgproc_remap_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jlong)(map1_nativeObj), (C.jlong)(map2_nativeObj), (C.jint)(interpolation))
}

func ImgprocNative_resize_0(src_nativeObj int64, dst_nativeObj int64, dsize_width float64, dsize_height float64, fx float64, fy float64, interpolation int) {
	C.Java_org_opencv_imgproc_Imgproc_resize_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jdouble)(dsize_width), (C.jdouble)(dsize_height), (C.jdouble)(fx), (C.jdouble)(fy), (C.jint)(interpolation))
}

func ImgprocNative_resize_1(src_nativeObj int64, dst_nativeObj int64, dsize_width float64, dsize_height float64) {
	C.Java_org_opencv_imgproc_Imgproc_resize_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jdouble)(dsize_width), (C.jdouble)(dsize_height))
}

func ImgprocNative_rotatedRectangleIntersection_0(rect1_center_x float64, rect1_center_y float64, rect1_size_width float64, rect1_size_height float64, rect1_angle float64, rect2_center_x float64, rect2_center_y float64, rect2_size_width float64, rect2_size_height float64, rect2_angle float64, intersectingRegion_nativeObj int64) int {
	return int(C.Java_org_opencv_imgproc_Imgproc_rotatedRectangleIntersection_10(clzEnv, clzObj, (C.jdouble)(rect1_center_x), (C.jdouble)(rect1_center_y),
		(C.jdouble)(rect1_size_width), (C.jdouble)(rect1_size_height), (C.jdouble)(rect1_angle), (C.jdouble)(rect2_center_x), (C.jdouble)(rect2_center_y), (C.jdouble)(rect2_size_width), (C.jdouble)(rect2_size_height), (C.jdouble)(rect2_angle), (C.jlong)(intersectingRegion_nativeObj)))
}

func ImgprocNative_sepFilter2D_0(src_nativeObj int64, dst_nativeObj int64, ddepth int, kernelX_nativeObj int64, kernelY_nativeObj int64, anchor_x float64, anchor_y float64, delta float64, borderType int) {
	C.Java_org_opencv_imgproc_Imgproc_sepFilter2D_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(ddepth), (C.jlong)(kernelX_nativeObj), (C.jlong)(kernelY_nativeObj), (C.jdouble)(anchor_x), (C.jdouble)(anchor_y), (C.jdouble)(delta), (C.jint)(borderType))
}

func ImgprocNative_sepFilter2D_1(src_nativeObj int64, dst_nativeObj int64, ddepth int, kernelX_nativeObj int64, kernelY_nativeObj int64, anchor_x float64, anchor_y float64, delta float64) {
	C.Java_org_opencv_imgproc_Imgproc_sepFilter2D_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(ddepth), (C.jlong)(kernelX_nativeObj), (C.jlong)(kernelY_nativeObj), (C.jdouble)(anchor_x), (C.jdouble)(anchor_y), (C.jdouble)(delta))
}

func ImgprocNative_sepFilter2D_2(src_nativeObj int64, dst_nativeObj int64, ddepth int, kernelX_nativeObj int64, kernelY_nativeObj int64) {
	C.Java_org_opencv_imgproc_Imgproc_sepFilter2D_12(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jint)(ddepth), (C.jlong)(kernelX_nativeObj), (C.jlong)(kernelY_nativeObj))
}

func ImgprocNative_spatialGradient_0(src_nativeObj int64, dx_nativeObj int64, dy_nativeObj int64, ksize int, borderType int) {
	C.Java_org_opencv_imgproc_Imgproc_spatialGradient_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dx_nativeObj), (C.jlong)(dy_nativeObj), (C.jint)(ksize), (C.jint)(borderType))
}

func ImgprocNative_spatialGradient_1(src_nativeObj int64, dx_nativeObj int64, dy_nativeObj int64, ksize int) {
	C.Java_org_opencv_imgproc_Imgproc_spatialGradient_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dx_nativeObj), (C.jlong)(dy_nativeObj), (C.jint)(ksize))
}

func ImgprocNative_spatialGradient_2(src_nativeObj int64, dx_nativeObj int64, dy_nativeObj int64) {
	C.Java_org_opencv_imgproc_Imgproc_spatialGradient_12(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dx_nativeObj), (C.jlong)(dy_nativeObj))
}

func ImgprocNative_sqrBoxFilter_0(_src_nativeObj int64, _dst_nativeObj int64, ddepth int, ksize_width float64, ksize_height float64, anchor_x float64, anchor_y float64, normalize bool, borderType int) {
	C.Java_org_opencv_imgproc_Imgproc_sqrBoxFilter_10(clzEnv, clzObj, (C.jlong)(_src_nativeObj), (C.jlong)(_dst_nativeObj), (C.jint)(ddepth), (C.jdouble)(ksize_width), (C.jdouble)(ksize_height), (C.jdouble)(anchor_x), (C.jdouble)(anchor_y), tojboolean(normalize), (C.jint)(borderType))
}

func ImgprocNative_sqrBoxFilter_1(_src_nativeObj int64, _dst_nativeObj int64, ddepth int, ksize_width float64, ksize_height float64, anchor_x float64, anchor_y float64, normalize bool) {
	C.Java_org_opencv_imgproc_Imgproc_sqrBoxFilter_11(clzEnv, clzObj, (C.jlong)(_src_nativeObj), (C.jlong)(_dst_nativeObj), (C.jint)(ddepth), (C.jdouble)(ksize_width), (C.jdouble)(ksize_height), (C.jdouble)(anchor_x), (C.jdouble)(anchor_y), tojboolean(normalize))
}

func ImgprocNative_sqrBoxFilter_2(_src_nativeObj int64, _dst_nativeObj int64, ddepth int, ksize_width float64, ksize_height float64) {
	C.Java_org_opencv_imgproc_Imgproc_sqrBoxFilter_12(clzEnv, clzObj, (C.jlong)(_src_nativeObj), (C.jlong)(_dst_nativeObj), (C.jint)(ddepth), (C.jdouble)(ksize_width), (C.jdouble)(ksize_height))
}

func ImgprocNative_threshold_0(src_nativeObj int64, dst_nativeObj int64, thresh float64, maxval float64, rtype int) float64 {
	return float64(C.Java_org_opencv_imgproc_Imgproc_threshold_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jdouble)(thresh), (C.jdouble)(maxval), (C.jint)(rtype)))
}

func ImgprocNative_undistortPointsIter_0(src_nativeObj int64, dst_nativeObj int64, cameraMatrix_nativeObj int64, distCoeffs_nativeObj int64, R_nativeObj int64, P_nativeObj int64, criteria_type int, criteria_maxCount int, criteria_epsilon float64) {
	C.Java_org_opencv_imgproc_Imgproc_undistortPointsIter_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jlong)(cameraMatrix_nativeObj), (C.jlong)(distCoeffs_nativeObj), (C.jlong)(R_nativeObj), (C.jlong)(P_nativeObj), (C.jint)(criteria_type), (C.jint)(criteria_maxCount), (C.jdouble)(criteria_epsilon))
}

func ImgprocNative_undistortPoints_0(src_nativeObj int64, dst_nativeObj int64, cameraMatrix_nativeObj int64, distCoeffs_nativeObj int64, R_nativeObj int64, P_nativeObj int64) {
	C.Java_org_opencv_imgproc_Imgproc_undistortPoints_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jlong)(cameraMatrix_nativeObj), (C.jlong)(distCoeffs_nativeObj), (C.jlong)(R_nativeObj), (C.jlong)(P_nativeObj))
}

func ImgprocNative_undistortPoints_1(src_nativeObj int64, dst_nativeObj int64, cameraMatrix_nativeObj int64, distCoeffs_nativeObj int64) {
	C.Java_org_opencv_imgproc_Imgproc_undistortPoints_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jlong)(cameraMatrix_nativeObj), (C.jlong)(distCoeffs_nativeObj))
}

func ImgprocNative_undistort_0(src_nativeObj int64, dst_nativeObj int64, cameraMatrix_nativeObj int64, distCoeffs_nativeObj int64, newCameraMatrix_nativeObj int64) {
	C.Java_org_opencv_imgproc_Imgproc_undistort_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jlong)(cameraMatrix_nativeObj), (C.jlong)(distCoeffs_nativeObj), (C.jlong)(newCameraMatrix_nativeObj))
}

func ImgprocNative_undistort_1(src_nativeObj int64, dst_nativeObj int64, cameraMatrix_nativeObj int64, distCoeffs_nativeObj int64) {
	C.Java_org_opencv_imgproc_Imgproc_undistort_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jlong)(cameraMatrix_nativeObj), (C.jlong)(distCoeffs_nativeObj))
}

func ImgprocNative_warpAffine_0(src_nativeObj int64, dst_nativeObj int64, M_nativeObj int64, dsize_width float64, dsize_height float64, flags int, borderMode int, borderValue_val0 float64, borderValue_val1 float64, borderValue_val2 float64, borderValue_val3 float64) {
	C.Java_org_opencv_imgproc_Imgproc_warpAffine_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jlong)(M_nativeObj), (C.jdouble)(dsize_width), (C.jdouble)(dsize_height), (C.jint)(flags), (C.jint)(borderMode), (C.jdouble)(borderValue_val0), (C.jdouble)(borderValue_val1), (C.jdouble)(borderValue_val2), (C.jdouble)(borderValue_val3))
}

func ImgprocNative_warpAffine_1(src_nativeObj int64, dst_nativeObj int64, M_nativeObj int64, dsize_width float64, dsize_height float64, flags int) {
	C.Java_org_opencv_imgproc_Imgproc_warpAffine_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jlong)(M_nativeObj), (C.jdouble)(dsize_width), (C.jdouble)(dsize_height), (C.jint)(flags))
}

func ImgprocNative_warpAffine_2(src_nativeObj int64, dst_nativeObj int64, M_nativeObj int64, dsize_width float64, dsize_height float64) {
	C.Java_org_opencv_imgproc_Imgproc_warpAffine_12(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jlong)(M_nativeObj), (C.jdouble)(dsize_width), (C.jdouble)(dsize_height))
}

func ImgprocNative_warpPerspective_0(src_nativeObj int64, dst_nativeObj int64, M_nativeObj int64, dsize_width float64, dsize_height float64, flags int, borderMode int, borderValue_val0 float64, borderValue_val1 float64, borderValue_val2 float64, borderValue_val3 float64) {
	C.Java_org_opencv_imgproc_Imgproc_warpPerspective_10(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jlong)(M_nativeObj), (C.jdouble)(dsize_width), (C.jdouble)(dsize_height), (C.jint)(flags), (C.jint)(borderMode), (C.jdouble)(borderValue_val0), (C.jdouble)(borderValue_val1), (C.jdouble)(borderValue_val2), (C.jdouble)(borderValue_val3))
}

func ImgprocNative_warpPerspective_1(src_nativeObj int64, dst_nativeObj int64, M_nativeObj int64, dsize_width float64, dsize_height float64, flags int) {
	C.Java_org_opencv_imgproc_Imgproc_warpPerspective_11(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jlong)(M_nativeObj), (C.jdouble)(dsize_width), (C.jdouble)(dsize_height), (C.jint)(flags))
}

func ImgprocNative_warpPerspective_2(src_nativeObj int64, dst_nativeObj int64, M_nativeObj int64, dsize_width float64, dsize_height float64) {
	C.Java_org_opencv_imgproc_Imgproc_warpPerspective_12(clzEnv, clzObj, (C.jlong)(src_nativeObj), (C.jlong)(dst_nativeObj), (C.jlong)(M_nativeObj), (C.jdouble)(dsize_width), (C.jdouble)(dsize_height))
}

func ImgprocNative_watershed_0(image_nativeObj int64, markers_nativeObj int64) {
	C.Java_org_opencv_imgproc_Imgproc_watershed_10(clzEnv, clzObj, (C.jlong)(image_nativeObj), (C.jlong)(markers_nativeObj))
}

func LineSegmentDetectorNative_compareSegments_0(nativeObj int64, size_width float64, size_height float64, lines1_nativeObj int64, lines2_nativeObj int64, _image_nativeObj int64) int {
	return int(C.Java_org_opencv_imgproc_LineSegmentDetector_compareSegments_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jdouble)(size_width), (C.jdouble)(size_height), (C.jlong)(lines1_nativeObj), (C.jlong)(lines2_nativeObj), (C.jlong)(_image_nativeObj)))
}

func LineSegmentDetectorNative_compareSegments_1(nativeObj int64, size_width float64, size_height float64, lines1_nativeObj int64, lines2_nativeObj int64) int {
	return int(C.Java_org_opencv_imgproc_LineSegmentDetector_compareSegments_11(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jdouble)(size_width), (C.jdouble)(size_height), (C.jlong)(lines1_nativeObj), (C.jlong)(lines2_nativeObj)))
}

func LineSegmentDetectorNative_delete(nativeObj int64) {
	C.Java_org_opencv_imgproc_LineSegmentDetector_delete(clzEnv, clzObj, (C.jlong)(nativeObj))
}

func LineSegmentDetectorNative_detect_0(nativeObj int64, _image_nativeObj int64, _lines_nativeObj int64, width_nativeObj int64, prec_nativeObj int64, nfa_nativeObj int64) {
	C.Java_org_opencv_imgproc_LineSegmentDetector_detect_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(_image_nativeObj), (C.jlong)(_lines_nativeObj), (C.jlong)(width_nativeObj), (C.jlong)(prec_nativeObj), (C.jlong)(nfa_nativeObj))
}

func LineSegmentDetectorNative_detect_1(nativeObj int64, _image_nativeObj int64, _lines_nativeObj int64) {
	C.Java_org_opencv_imgproc_LineSegmentDetector_detect_11(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(_image_nativeObj), (C.jlong)(_lines_nativeObj))
}

func LineSegmentDetectorNative_drawSegments_0(nativeObj int64, _image_nativeObj int64, lines_nativeObj int64) {
	C.Java_org_opencv_imgproc_LineSegmentDetector_drawSegments_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(_image_nativeObj), (C.jlong)(lines_nativeObj))
}

func Subdiv2DNative_Subdiv2D_0(rect_x int, rect_y int, rect_width int, rect_height int) int64 {
	return int64(C.Java_org_opencv_imgproc_Subdiv2D_Subdiv2D_10(clzEnv, clzObj, (C.jint)(rect_x), (C.jint)(rect_y), (C.jint)(rect_width), (C.jint)(rect_height)))
}

func Subdiv2DNative_Subdiv2D_1() int64 {
	return int64(C.Java_org_opencv_imgproc_Subdiv2D_Subdiv2D_11(clzEnv, clzObj))
}

func Subdiv2DNative_delete(nativeObj int64) {
	C.Java_org_opencv_imgproc_Subdiv2D_delete(clzEnv, clzObj, (C.jlong)(nativeObj))
}

func Subdiv2DNative_edgeDst_0(nativeObj int64, edge int, dstpt_out []float64) int {
	defer ungointerface(dstpt_out)
	return int(C.Java_org_opencv_imgproc_Subdiv2D_edgeDst_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(edge), tojdoubleArray(dstpt_out)))
}

func Subdiv2DNative_edgeDst_1(nativeObj int64, edge int) int {
	return int(C.Java_org_opencv_imgproc_Subdiv2D_edgeDst_11(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(edge)))
}

func Subdiv2DNative_edgeOrg_0(nativeObj int64, edge int, orgpt_out []float64) int {
	defer ungointerface(orgpt_out)
	return int(C.Java_org_opencv_imgproc_Subdiv2D_edgeOrg_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(edge), tojdoubleArray(orgpt_out)))
}

func Subdiv2DNative_edgeOrg_1(nativeObj int64, edge int) int {
	return int(C.Java_org_opencv_imgproc_Subdiv2D_edgeOrg_11(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(edge)))
}

func Subdiv2DNative_findNearest_0(nativeObj int64, pt_x float64, pt_y float64, nearestPt_out []float64) int {
	defer ungointerface(nearestPt_out)
	return int(C.Java_org_opencv_imgproc_Subdiv2D_findNearest_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jdouble)(pt_x), (C.jdouble)(pt_y), tojdoubleArray(nearestPt_out)))
}

func Subdiv2DNative_findNearest_1(nativeObj int64, pt_x float64, pt_y float64) int {
	return int(C.Java_org_opencv_imgproc_Subdiv2D_findNearest_11(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jdouble)(pt_x), (C.jdouble)(pt_y)))
}

func Subdiv2DNative_getEdgeList_0(nativeObj int64, edgeList_mat_nativeObj int64) {
	C.Java_org_opencv_imgproc_Subdiv2D_getEdgeList_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(edgeList_mat_nativeObj))
}

func Subdiv2DNative_getEdge_0(nativeObj int64, edge int, nextEdgeType int) int {
	return int(C.Java_org_opencv_imgproc_Subdiv2D_getEdge_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(edge), (C.jint)(nextEdgeType)))
}

func Subdiv2DNative_getLeadingEdgeList_0(nativeObj int64, leadingEdgeList_mat_nativeObj int64) {
	C.Java_org_opencv_imgproc_Subdiv2D_getLeadingEdgeList_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(leadingEdgeList_mat_nativeObj))
}

func Subdiv2DNative_getTriangleList_0(nativeObj int64, triangleList_mat_nativeObj int64) {
	C.Java_org_opencv_imgproc_Subdiv2D_getTriangleList_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(triangleList_mat_nativeObj))
}

func Subdiv2DNative_getVertex_0(nativeObj int64, vertex int, firstEdge_out []float64) []float64 {
	defer ungointerface(firstEdge_out)
	return togofloat64Array(C.Java_org_opencv_imgproc_Subdiv2D_getVertex_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(vertex), tojdoubleArray(firstEdge_out)))
}

func Subdiv2DNative_getVertex_1(nativeObj int64, vertex int) []float64 {
	return togofloat64Array(C.Java_org_opencv_imgproc_Subdiv2D_getVertex_11(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(vertex)))
}

func Subdiv2DNative_getVoronoiFacetList_0(nativeObj int64, idx_mat_nativeObj int64, facetList_mat_nativeObj int64, facetCenters_mat_nativeObj int64) {
	C.Java_org_opencv_imgproc_Subdiv2D_getVoronoiFacetList_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(idx_mat_nativeObj), (C.jlong)(facetList_mat_nativeObj), (C.jlong)(facetCenters_mat_nativeObj))
}

func Subdiv2DNative_initDelaunay_0(nativeObj int64, rect_x int, rect_y int, rect_width int, rect_height int) {
	C.Java_org_opencv_imgproc_Subdiv2D_initDelaunay_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(rect_x), (C.jint)(rect_y), (C.jint)(rect_width), (C.jint)(rect_height))
}

func Subdiv2DNative_insert_0(nativeObj int64, pt_x float64, pt_y float64) int {
	return int(C.Java_org_opencv_imgproc_Subdiv2D_insert_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jdouble)(pt_x), (C.jdouble)(pt_y)))
}

func Subdiv2DNative_insert_1(nativeObj int64, ptvec_mat_nativeObj int64) {
	C.Java_org_opencv_imgproc_Subdiv2D_insert_11(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(ptvec_mat_nativeObj))
}

func Subdiv2DNative_locate_0(nativeObj int64, pt_x float64, pt_y float64, edge_out []float64, vertex_out []float64) int {
	defer ungointerface(edge_out)
	defer ungointerface(vertex_out)
	return int(C.Java_org_opencv_imgproc_Subdiv2D_locate_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jdouble)(pt_x), (C.jdouble)(pt_y), tojdoubleArray(edge_out), tojdoubleArray(vertex_out)))
}

func Subdiv2DNative_nextEdge_0(nativeObj int64, edge int) int {
	return int(C.Java_org_opencv_imgproc_Subdiv2D_nextEdge_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(edge)))
}

func Subdiv2DNative_rotateEdge_0(nativeObj int64, edge int, rotate int) int {
	return int(C.Java_org_opencv_imgproc_Subdiv2D_rotateEdge_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(edge), (C.jint)(rotate)))
}

func Subdiv2DNative_symEdge_0(nativeObj int64, edge int) int {
	return int(C.Java_org_opencv_imgproc_Subdiv2D_symEdge_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(edge)))
}
