package core

/*
#include "jni.h"

extern jdoubleArray Java_org_opencv_calib3d_Calib3d_RQDecomp3x3_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong);
extern jdoubleArray Java_org_opencv_calib3d_Calib3d_RQDecomp3x3_11(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_calib3d_Calib3d_Rodrigues_10(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_calib3d_Calib3d_Rodrigues_11(JNIEnv*, jclass, jlong, jlong);
extern jdouble Java_org_opencv_calib3d_Calib3d_calibrateCameraExtended_10(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jlong, jlong, jlong, jlong, jlong, jlong, jlong, jint, jint, jint, jdouble);
extern jdouble Java_org_opencv_calib3d_Calib3d_calibrateCameraExtended_11(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jlong, jlong, jlong, jlong, jlong, jlong, jlong, jint);
extern jdouble Java_org_opencv_calib3d_Calib3d_calibrateCameraExtended_12(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jlong, jlong, jlong, jlong, jlong, jlong, jlong);
extern jdouble Java_org_opencv_calib3d_Calib3d_calibrateCamera_10(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jlong, jlong, jlong, jlong, jint, jint, jint, jdouble);
extern jdouble Java_org_opencv_calib3d_Calib3d_calibrateCamera_11(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jlong, jlong, jlong, jlong, jint);
extern jdouble Java_org_opencv_calib3d_Calib3d_calibrateCamera_12(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jlong, jlong, jlong, jlong);
extern jdouble Java_org_opencv_calib3d_Calib3d_calibrate_10(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jlong, jlong, jlong, jlong, jint, jint, jint, jdouble);
extern jdouble Java_org_opencv_calib3d_Calib3d_calibrate_11(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jlong, jlong, jlong, jlong, jint);
extern jdouble Java_org_opencv_calib3d_Calib3d_calibrate_12(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_calib3d_Calib3d_calibrationMatrixValues_10(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jdoubleArray, jdoubleArray, jdoubleArray, jdoubleArray, jdoubleArray);
extern void Java_org_opencv_calib3d_Calib3d_composeRT_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong, jlong, jlong, jlong, jlong, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_calib3d_Calib3d_composeRT_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_calib3d_Calib3d_computeCorrespondEpilines_10(JNIEnv*, jclass, jlong, jint, jlong, jlong);
extern void Java_org_opencv_calib3d_Calib3d_convertPointsFromHomogeneous_10(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_calib3d_Calib3d_convertPointsToHomogeneous_10(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_calib3d_Calib3d_correctMatches_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_calib3d_Calib3d_decomposeEssentialMat_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern jint Java_org_opencv_calib3d_Calib3d_decomposeHomographyMat_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_calib3d_Calib3d_decomposeProjectionMatrix_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_calib3d_Calib3d_decomposeProjectionMatrix_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_calib3d_Calib3d_distortPoints_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jdouble);
extern void Java_org_opencv_calib3d_Calib3d_distortPoints_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_calib3d_Calib3d_drawChessboardCorners_10(JNIEnv*, jclass, jlong, jdouble, jdouble, jlong, jboolean);
extern jlong Java_org_opencv_calib3d_Calib3d_estimateAffine2D_10(JNIEnv*, jclass, jlong, jlong, jlong, jint, jdouble, jlong, jdouble, jlong);
extern jlong Java_org_opencv_calib3d_Calib3d_estimateAffine2D_11(JNIEnv*, jclass, jlong, jlong);
extern jint Java_org_opencv_calib3d_Calib3d_estimateAffine3D_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jdouble, jdouble);
extern jint Java_org_opencv_calib3d_Calib3d_estimateAffine3D_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern jlong Java_org_opencv_calib3d_Calib3d_estimateAffinePartial2D_10(JNIEnv*, jclass, jlong, jlong, jlong, jint, jdouble, jlong, jdouble, jlong);
extern jlong Java_org_opencv_calib3d_Calib3d_estimateAffinePartial2D_11(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_calib3d_Calib3d_estimateNewCameraMatrixForUndistortRectify_10(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jlong, jlong, jdouble, jdouble, jdouble, jdouble);
extern void Java_org_opencv_calib3d_Calib3d_estimateNewCameraMatrixForUndistortRectify_11(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jlong, jlong);
extern void Java_org_opencv_calib3d_Calib3d_filterSpeckles_10(JNIEnv*, jclass, jlong, jdouble, jint, jdouble, jlong);
extern void Java_org_opencv_calib3d_Calib3d_filterSpeckles_11(JNIEnv*, jclass, jlong, jdouble, jint, jdouble);
extern jboolean Java_org_opencv_calib3d_Calib3d_findChessboardCorners_10(JNIEnv*, jclass, jlong, jdouble, jdouble, jlong, jint);
extern jboolean Java_org_opencv_calib3d_Calib3d_findChessboardCorners_11(JNIEnv*, jclass, jlong, jdouble, jdouble, jlong);
extern jboolean Java_org_opencv_calib3d_Calib3d_findCirclesGrid_10(JNIEnv*, jclass, jlong, jdouble, jdouble, jlong, jint);
extern jboolean Java_org_opencv_calib3d_Calib3d_findCirclesGrid_11(JNIEnv*, jclass, jlong, jdouble, jdouble, jlong);
extern jlong Java_org_opencv_calib3d_Calib3d_findEssentialMat_10(JNIEnv*, jclass, jlong, jlong, jlong, jint, jdouble, jdouble, jlong);
extern jlong Java_org_opencv_calib3d_Calib3d_findEssentialMat_11(JNIEnv*, jclass, jlong, jlong, jlong, jint, jdouble, jdouble);
extern jlong Java_org_opencv_calib3d_Calib3d_findEssentialMat_12(JNIEnv*, jclass, jlong, jlong, jlong);
extern jlong Java_org_opencv_calib3d_Calib3d_findEssentialMat_13(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jdouble, jint, jdouble, jdouble, jlong);
extern jlong Java_org_opencv_calib3d_Calib3d_findEssentialMat_14(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jdouble, jint, jdouble, jdouble);
extern jlong Java_org_opencv_calib3d_Calib3d_findEssentialMat_15(JNIEnv*, jclass, jlong, jlong);
extern jlong Java_org_opencv_calib3d_Calib3d_findFundamentalMat_10(JNIEnv*, jclass, jlong, jlong, jint, jdouble, jdouble, jlong);
extern jlong Java_org_opencv_calib3d_Calib3d_findFundamentalMat_11(JNIEnv*, jclass, jlong, jlong, jint, jdouble, jdouble);
extern jlong Java_org_opencv_calib3d_Calib3d_findFundamentalMat_12(JNIEnv*, jclass, jlong, jlong);
extern jlong Java_org_opencv_calib3d_Calib3d_findHomography_10(JNIEnv*, jclass, jlong, jlong, jint, jdouble, jlong, jint, jdouble);
extern jlong Java_org_opencv_calib3d_Calib3d_findHomography_11(JNIEnv*, jclass, jlong, jlong, jint, jdouble);
extern jlong Java_org_opencv_calib3d_Calib3d_findHomography_12(JNIEnv*, jclass, jlong, jlong);
extern jlong Java_org_opencv_calib3d_Calib3d_getOptimalNewCameraMatrix_10(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jdouble, jdouble, jdouble, jdoubleArray, jboolean);
extern jlong Java_org_opencv_calib3d_Calib3d_getOptimalNewCameraMatrix_11(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jdouble);
extern jdoubleArray Java_org_opencv_calib3d_Calib3d_getValidDisparityROI_10(JNIEnv*, jclass, jint, jint, jint, jint, jint, jint, jint, jint, jint, jint, jint);
extern jlong Java_org_opencv_calib3d_Calib3d_initCameraMatrix2D_10(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble, jdouble);
extern jlong Java_org_opencv_calib3d_Calib3d_initCameraMatrix2D_11(JNIEnv*, jclass, jlong, jlong, jdouble, jdouble);
extern void Java_org_opencv_calib3d_Calib3d_initUndistortRectifyMap_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jdouble, jdouble, jint, jlong, jlong);
extern void Java_org_opencv_calib3d_Calib3d_matMulDeriv_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_calib3d_Calib3d_projectPoints_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong, jlong, jdouble);
extern void Java_org_opencv_calib3d_Calib3d_projectPoints_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_calib3d_Calib3d_projectPoints_12(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong, jdouble, jlong);
extern void Java_org_opencv_calib3d_Calib3d_projectPoints_13(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong);
extern jint Java_org_opencv_calib3d_Calib3d_recoverPose_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jdouble, jdouble, jdouble, jlong);
extern jint Java_org_opencv_calib3d_Calib3d_recoverPose_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jdouble, jdouble, jdouble);
extern jint Java_org_opencv_calib3d_Calib3d_recoverPose_12(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong);
extern jint Java_org_opencv_calib3d_Calib3d_recoverPose_13(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong, jlong);
extern jint Java_org_opencv_calib3d_Calib3d_recoverPose_14(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong);
extern jint Java_org_opencv_calib3d_Calib3d_recoverPose_15(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong, jdouble, jlong, jlong);
extern jint Java_org_opencv_calib3d_Calib3d_recoverPose_16(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong, jdouble);
extern jfloat Java_org_opencv_calib3d_Calib3d_rectify3Collinear_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong, jlong, jlong, jdouble, jdouble, jlong, jlong, jlong, jlong, jlong, jlong, jlong, jlong, jlong, jlong, jlong, jdouble, jdouble, jdouble, jdoubleArray, jdoubleArray, jint);
extern void Java_org_opencv_calib3d_Calib3d_reprojectImageTo3D_10(JNIEnv*, jclass, jlong, jlong, jlong, jboolean, jint);
extern void Java_org_opencv_calib3d_Calib3d_reprojectImageTo3D_11(JNIEnv*, jclass, jlong, jlong, jlong, jboolean);
extern void Java_org_opencv_calib3d_Calib3d_reprojectImageTo3D_12(JNIEnv*, jclass, jlong, jlong, jlong);
extern jdouble Java_org_opencv_calib3d_Calib3d_sampsonDistance_10(JNIEnv*, jclass, jlong, jlong, jlong);
extern jint Java_org_opencv_calib3d_Calib3d_solveP3P_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong, jint);
extern jboolean Java_org_opencv_calib3d_Calib3d_solvePnPRansac_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong, jboolean, jint, jfloat, jdouble, jlong, jint);
extern jboolean Java_org_opencv_calib3d_Calib3d_solvePnPRansac_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong);
extern jboolean Java_org_opencv_calib3d_Calib3d_solvePnP_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong, jboolean, jint);
extern jboolean Java_org_opencv_calib3d_Calib3d_solvePnP_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong);
extern jdouble Java_org_opencv_calib3d_Calib3d_stereoCalibrate_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong, jlong, jdouble, jdouble, jlong, jlong, jlong, jlong, jint, jint, jint, jdouble);
extern jdouble Java_org_opencv_calib3d_Calib3d_stereoCalibrate_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong, jlong, jdouble, jdouble, jlong, jlong, jlong, jlong, jint);
extern jdouble Java_org_opencv_calib3d_Calib3d_stereoCalibrate_12(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong, jlong, jdouble, jdouble, jlong, jlong, jlong, jlong);
extern jdouble Java_org_opencv_calib3d_Calib3d_stereoCalibrate_13(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong, jlong, jdouble, jdouble, jlong, jlong, jint, jint, jint, jdouble);
extern jdouble Java_org_opencv_calib3d_Calib3d_stereoCalibrate_14(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong, jlong, jdouble, jdouble, jlong, jlong, jint);
extern jdouble Java_org_opencv_calib3d_Calib3d_stereoCalibrate_15(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong, jlong, jdouble, jdouble, jlong, jlong);
extern jboolean Java_org_opencv_calib3d_Calib3d_stereoRectifyUncalibrated_10(JNIEnv*, jclass, jlong, jlong, jlong, jdouble, jdouble, jlong, jlong, jdouble);
extern jboolean Java_org_opencv_calib3d_Calib3d_stereoRectifyUncalibrated_11(JNIEnv*, jclass, jlong, jlong, jlong, jdouble, jdouble, jlong, jlong);
extern void Java_org_opencv_calib3d_Calib3d_stereoRectify_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jdouble, jdouble, jlong, jlong, jlong, jlong, jlong, jlong, jlong, jint, jdouble, jdouble, jdouble, jdoubleArray, jdoubleArray);
extern void Java_org_opencv_calib3d_Calib3d_stereoRectify_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jdouble, jdouble, jlong, jlong, jlong, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_calib3d_Calib3d_stereoRectify_12(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jdouble, jdouble, jlong, jlong, jlong, jlong, jlong, jlong, jlong, jint, jdouble, jdouble, jdouble, jdouble);
extern void Java_org_opencv_calib3d_Calib3d_stereoRectify_13(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jdouble, jdouble, jlong, jlong, jlong, jlong, jlong, jlong, jlong, jint);
extern void Java_org_opencv_calib3d_Calib3d_triangulatePoints_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_calib3d_Calib3d_undistortImage_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jdouble, jdouble);
extern void Java_org_opencv_calib3d_Calib3d_undistortImage_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_calib3d_Calib3d_undistortPoints_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_calib3d_Calib3d_undistortPoints_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_calib3d_Calib3d_validateDisparity_10(JNIEnv*, jclass, jlong, jlong, jint, jint, jint);
extern void Java_org_opencv_calib3d_Calib3d_validateDisparity_11(JNIEnv*, jclass, jlong, jlong, jint, jint);
extern jlong Java_org_opencv_calib3d_StereoBM_create_10(JNIEnv*, jclass, jint, jint);
extern jlong Java_org_opencv_calib3d_StereoBM_create_11(JNIEnv*, jclass);
extern void Java_org_opencv_calib3d_StereoBM_delete(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_calib3d_StereoBM_getPreFilterCap_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_calib3d_StereoBM_getPreFilterSize_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_calib3d_StereoBM_getPreFilterType_10(JNIEnv*, jclass, jlong);
extern jdoubleArray Java_org_opencv_calib3d_StereoBM_getROI1_10(JNIEnv*, jclass, jlong);
extern jdoubleArray Java_org_opencv_calib3d_StereoBM_getROI2_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_calib3d_StereoBM_getSmallerBlockSize_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_calib3d_StereoBM_getTextureThreshold_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_calib3d_StereoBM_getUniquenessRatio_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_calib3d_StereoBM_setPreFilterCap_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_calib3d_StereoBM_setPreFilterSize_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_calib3d_StereoBM_setPreFilterType_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_calib3d_StereoBM_setROI1_10(JNIEnv*, jclass, jlong, jint, jint, jint, jint);
extern void Java_org_opencv_calib3d_StereoBM_setROI2_10(JNIEnv*, jclass, jlong, jint, jint, jint, jint);
extern void Java_org_opencv_calib3d_StereoBM_setSmallerBlockSize_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_calib3d_StereoBM_setTextureThreshold_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_calib3d_StereoBM_setUniquenessRatio_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_calib3d_StereoMatcher_compute_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_calib3d_StereoMatcher_delete(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_calib3d_StereoMatcher_getBlockSize_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_calib3d_StereoMatcher_getDisp12MaxDiff_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_calib3d_StereoMatcher_getMinDisparity_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_calib3d_StereoMatcher_getNumDisparities_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_calib3d_StereoMatcher_getSpeckleRange_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_calib3d_StereoMatcher_getSpeckleWindowSize_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_calib3d_StereoMatcher_setBlockSize_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_calib3d_StereoMatcher_setDisp12MaxDiff_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_calib3d_StereoMatcher_setMinDisparity_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_calib3d_StereoMatcher_setNumDisparities_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_calib3d_StereoMatcher_setSpeckleRange_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_calib3d_StereoMatcher_setSpeckleWindowSize_10(JNIEnv*, jclass, jlong, jint);
extern jlong Java_org_opencv_calib3d_StereoSGBM_create_10(JNIEnv*, jclass, jint, jint, jint, jint, jint, jint, jint, jint, jint, jint, jint);
extern jlong Java_org_opencv_calib3d_StereoSGBM_create_11(JNIEnv*, jclass);
extern void Java_org_opencv_calib3d_StereoSGBM_delete(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_calib3d_StereoSGBM_getMode_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_calib3d_StereoSGBM_getP1_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_calib3d_StereoSGBM_getP2_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_calib3d_StereoSGBM_getPreFilterCap_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_calib3d_StereoSGBM_getUniquenessRatio_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_calib3d_StereoSGBM_setMode_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_calib3d_StereoSGBM_setP1_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_calib3d_StereoSGBM_setP2_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_calib3d_StereoSGBM_setPreFilterCap_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_calib3d_StereoSGBM_setUniquenessRatio_10(JNIEnv*, jclass, jlong, jint);

*/
import "C"

func Calib3dNative_RQDecomp3x3_0(src_nativeObj int64,mtxR_nativeObj int64,mtxQ_nativeObj int64,Qx_nativeObj int64,Qy_nativeObj int64,Qz_nativeObj int64) []float64 {
	return togofloat64Array(C.Java_org_opencv_calib3d_Calib3d_RQDecomp3x3_10(clzEnv,clzObj,(C.jlong)(src_nativeObj),(C.jlong)(mtxR_nativeObj),(C.jlong)(mtxQ_nativeObj),(C.jlong)(Qx_nativeObj),(C.jlong)(Qy_nativeObj),(C.jlong)(Qz_nativeObj)))
}
func Calib3dNative_RQDecomp3x3_1(src_nativeObj int64,mtxR_nativeObj int64,mtxQ_nativeObj int64) []float64 {
	return togofloat64Array(C.Java_org_opencv_calib3d_Calib3d_RQDecomp3x3_11(clzEnv,clzObj,(C.jlong)(src_nativeObj),(C.jlong)(mtxR_nativeObj),(C.jlong)(mtxQ_nativeObj)))
}

func Calib3dNative_Rodrigues_0(src_nativeObj int64,dst_nativeObj int64,jacobian_nativeObj int64) {
	C.Java_org_opencv_calib3d_Calib3d_Rodrigues_10(clzEnv,clzObj,(C.jlong)(src_nativeObj),(C.jlong)(dst_nativeObj),(C.jlong)(jacobian_nativeObj))
}

func Calib3dNative_Rodrigues_1(src_nativeObj int64,dst_nativeObj int64) {
	C.Java_org_opencv_calib3d_Calib3d_Rodrigues_11(clzEnv,clzObj,(C.jlong)(src_nativeObj),(C.jlong)(dst_nativeObj))
}

func Calib3dNative_calibrateCameraExtended_0(objectPoints_mat_nativeObj int64,imagePoints_mat_nativeObj int64,imageSize_width float64,imageSize_height float64,cameraMatrix_nativeObj int64,distCoeffs_nativeObj int64,rvecs_mat_nativeObj int64,tvecs_mat_nativeObj int64,stdDeviationsIntrinsics_nativeObj int64,stdDeviationsExtrinsics_nativeObj int64,perViewErrors_nativeObj int64,flags int,criteria_type int,criteria_maxCount int,criteria_epsilon float64) float64 {
	return float64(C.Java_org_opencv_calib3d_Calib3d_calibrateCameraExtended_10(clzEnv,clzObj,(C.jlong)(objectPoints_mat_nativeObj),(C.jlong)(imagePoints_mat_nativeObj),(C.jdouble)(imageSize_width),(C.jdouble)(imageSize_height),(C.jlong)(cameraMatrix_nativeObj),(C.jlong)(distCoeffs_nativeObj),(C.jlong)(rvecs_mat_nativeObj),(C.jlong)(tvecs_mat_nativeObj),(C.jlong)(stdDeviationsIntrinsics_nativeObj),(C.jlong)(stdDeviationsExtrinsics_nativeObj),(C.jlong)(perViewErrors_nativeObj),(C.jint)(flags),(C.jint)(criteria_type),(C.jint)(criteria_maxCount),(C.jdouble)(criteria_epsilon)))
}
func Calib3dNative_calibrateCameraExtended_1(objectPoints_mat_nativeObj int64,imagePoints_mat_nativeObj int64,imageSize_width float64,imageSize_height float64,cameraMatrix_nativeObj int64,distCoeffs_nativeObj int64,rvecs_mat_nativeObj int64,tvecs_mat_nativeObj int64,stdDeviationsIntrinsics_nativeObj int64,stdDeviationsExtrinsics_nativeObj int64,perViewErrors_nativeObj int64,flags int) float64 {
	return float64(C.Java_org_opencv_calib3d_Calib3d_calibrateCameraExtended_11(clzEnv,clzObj,(C.jlong)(objectPoints_mat_nativeObj),(C.jlong)(imagePoints_mat_nativeObj),(C.jdouble)(imageSize_width),(C.jdouble)(imageSize_height),(C.jlong)(cameraMatrix_nativeObj),(C.jlong)(distCoeffs_nativeObj),(C.jlong)(rvecs_mat_nativeObj),(C.jlong)(tvecs_mat_nativeObj),(C.jlong)(stdDeviationsIntrinsics_nativeObj),(C.jlong)(stdDeviationsExtrinsics_nativeObj),(C.jlong)(perViewErrors_nativeObj),(C.jint)(flags)))
}

func Calib3dNative_calibrateCameraExtended_2(objectPoints_mat_nativeObj int64,imagePoints_mat_nativeObj int64,imageSize_width float64,imageSize_height float64,cameraMatrix_nativeObj int64,distCoeffs_nativeObj int64,rvecs_mat_nativeObj int64,tvecs_mat_nativeObj int64,stdDeviationsIntrinsics_nativeObj int64,stdDeviationsExtrinsics_nativeObj int64,perViewErrors_nativeObj int64) float64 {
	return float64(C.Java_org_opencv_calib3d_Calib3d_calibrateCameraExtended_12(clzEnv,clzObj,(C.jlong)(objectPoints_mat_nativeObj),(C.jlong)(imagePoints_mat_nativeObj),(C.jdouble)(imageSize_width),(C.jdouble)(imageSize_height),(C.jlong)(cameraMatrix_nativeObj),(C.jlong)(distCoeffs_nativeObj),(C.jlong)(rvecs_mat_nativeObj),(C.jlong)(tvecs_mat_nativeObj),(C.jlong)(stdDeviationsIntrinsics_nativeObj),(C.jlong)(stdDeviationsExtrinsics_nativeObj),(C.jlong)(perViewErrors_nativeObj)))
}

func Calib3dNative_calibrateCamera_0(objectPoints_mat_nativeObj int64,imagePoints_mat_nativeObj int64,imageSize_width float64,imageSize_height float64,cameraMatrix_nativeObj int64,distCoeffs_nativeObj int64,rvecs_mat_nativeObj int64,tvecs_mat_nativeObj int64,flags int,criteria_type int,criteria_maxCount int,criteria_epsilon float64) float64 {
	return float64(C.Java_org_opencv_calib3d_Calib3d_calibrateCamera_10(clzEnv,clzObj,(C.jlong)(objectPoints_mat_nativeObj),(C.jlong)(imagePoints_mat_nativeObj),(C.jdouble)(imageSize_width),(C.jdouble)(imageSize_height),(C.jlong)(cameraMatrix_nativeObj),(C.jlong)(distCoeffs_nativeObj),(C.jlong)(rvecs_mat_nativeObj),(C.jlong)(tvecs_mat_nativeObj),(C.jint)(flags),(C.jint)(criteria_type),(C.jint)(criteria_maxCount),(C.jdouble)(criteria_epsilon)))
}

func Calib3dNative_calibrateCamera_1(objectPoints_mat_nativeObj int64,imagePoints_mat_nativeObj int64,imageSize_width float64,imageSize_height float64,cameraMatrix_nativeObj int64,distCoeffs_nativeObj int64,rvecs_mat_nativeObj int64,tvecs_mat_nativeObj int64,flags int) float64 {
	return float64(C.Java_org_opencv_calib3d_Calib3d_calibrateCamera_11(clzEnv,clzObj,(C.jlong)(objectPoints_mat_nativeObj),(C.jlong)(imagePoints_mat_nativeObj),(C.jdouble)(imageSize_width),(C.jdouble)(imageSize_height),(C.jlong)(cameraMatrix_nativeObj),(C.jlong)(distCoeffs_nativeObj),(C.jlong)(rvecs_mat_nativeObj),(C.jlong)(tvecs_mat_nativeObj),(C.jint)(flags)))
}

func Calib3dNative_calibrateCamera_2(objectPoints_mat_nativeObj int64,imagePoints_mat_nativeObj int64,imageSize_width float64,imageSize_height float64,cameraMatrix_nativeObj int64,distCoeffs_nativeObj int64,rvecs_mat_nativeObj int64,tvecs_mat_nativeObj int64) float64 {
	return float64(C.Java_org_opencv_calib3d_Calib3d_calibrateCamera_12(clzEnv,clzObj,(C.jlong)(objectPoints_mat_nativeObj),(C.jlong)(imagePoints_mat_nativeObj),(C.jdouble)(imageSize_width),(C.jdouble)(imageSize_height),(C.jlong)(cameraMatrix_nativeObj),(C.jlong)(distCoeffs_nativeObj),(C.jlong)(rvecs_mat_nativeObj),(C.jlong)(tvecs_mat_nativeObj)))
}

func Calib3dNative_calibrate_0(objectPoints_mat_nativeObj int64,imagePoints_mat_nativeObj int64,image_size_width float64,image_size_height float64,K_nativeObj int64,D_nativeObj int64,rvecs_mat_nativeObj int64,tvecs_mat_nativeObj int64,flags int,criteria_type int,criteria_maxCount int,criteria_epsilon float64) float64 {
	return float64(C.Java_org_opencv_calib3d_Calib3d_calibrate_10(clzEnv,clzObj,(C.jlong)(objectPoints_mat_nativeObj),(C.jlong)(imagePoints_mat_nativeObj),(C.jdouble)(image_size_width),(C.jdouble)(image_size_height),(C.jlong)(K_nativeObj),(C.jlong)(D_nativeObj),(C.jlong)(rvecs_mat_nativeObj),(C.jlong)(tvecs_mat_nativeObj),(C.jint)(flags),(C.jint)(criteria_type),(C.jint)(criteria_maxCount),(C.jdouble)(criteria_epsilon)))
}

func Calib3dNative_calibrate_1(objectPoints_mat_nativeObj int64,imagePoints_mat_nativeObj int64,image_size_width float64,image_size_height float64,K_nativeObj int64,D_nativeObj int64,rvecs_mat_nativeObj int64,tvecs_mat_nativeObj int64,flags int) float64 {
	return float64(C.Java_org_opencv_calib3d_Calib3d_calibrate_11(clzEnv,clzObj,(C.jlong)(objectPoints_mat_nativeObj),(C.jlong)(imagePoints_mat_nativeObj),(C.jdouble)(image_size_width),(C.jdouble)(image_size_height),(C.jlong)(K_nativeObj),(C.jlong)(D_nativeObj),(C.jlong)(rvecs_mat_nativeObj),(C.jlong)(tvecs_mat_nativeObj),(C.jint)(flags)))
}

func Calib3dNative_calibrate_2(objectPoints_mat_nativeObj int64,imagePoints_mat_nativeObj int64,image_size_width float64,image_size_height float64,K_nativeObj int64,D_nativeObj int64,rvecs_mat_nativeObj int64,tvecs_mat_nativeObj int64) float64 {
	return float64(C.Java_org_opencv_calib3d_Calib3d_calibrate_12(clzEnv,clzObj,(C.jlong)(objectPoints_mat_nativeObj),(C.jlong)(imagePoints_mat_nativeObj),(C.jdouble)(image_size_width),(C.jdouble)(image_size_height),(C.jlong)(K_nativeObj),(C.jlong)(D_nativeObj),(C.jlong)(rvecs_mat_nativeObj),(C.jlong)(tvecs_mat_nativeObj)))
}

func Calib3dNative_calibrationMatrixValues_0(cameraMatrix_nativeObj int64,imageSize_width float64,imageSize_height float64,apertureWidth float64,apertureHeight float64,fovx_out []float64,fovy_out []float64,focalLength_out []float64,principalPoint_out []float64,aspectRatio_out []float64) {
	C.Java_org_opencv_calib3d_Calib3d_calibrationMatrixValues_10(clzEnv,clzObj,(C.jlong)(cameraMatrix_nativeObj),(C.jdouble)(imageSize_width),(C.jdouble)(imageSize_height),(C.jdouble)(apertureWidth),(C.jdouble)(apertureHeight),tojdoubleArray(fovx_out),tojdoubleArray(fovy_out),tojdoubleArray(focalLength_out),tojdoubleArray(principalPoint_out),tojdoubleArray(aspectRatio_out))
}

func Calib3dNative_composeRT_0(rvec1_nativeObj int64,tvec1_nativeObj int64,rvec2_nativeObj int64,tvec2_nativeObj int64,rvec3_nativeObj int64,tvec3_nativeObj int64,dr3dr1_nativeObj int64,dr3dt1_nativeObj int64,dr3dr2_nativeObj int64,dr3dt2_nativeObj int64,dt3dr1_nativeObj int64,dt3dt1_nativeObj int64,dt3dr2_nativeObj int64,dt3dt2_nativeObj int64) {
	C.Java_org_opencv_calib3d_Calib3d_composeRT_10(clzEnv,clzObj,(C.jlong)(rvec1_nativeObj),(C.jlong)(tvec1_nativeObj),(C.jlong)(rvec2_nativeObj),(C.jlong)(tvec2_nativeObj),(C.jlong)(rvec3_nativeObj),(C.jlong)(tvec3_nativeObj),(C.jlong)(dr3dr1_nativeObj),(C.jlong)(dr3dt1_nativeObj),(C.jlong)(dr3dr2_nativeObj),(C.jlong)(dr3dt2_nativeObj),(C.jlong)(dt3dr1_nativeObj),(C.jlong)(dt3dt1_nativeObj),(C.jlong)(dt3dr2_nativeObj),(C.jlong)(dt3dt2_nativeObj))
}

func Calib3dNative_composeRT_1(rvec1_nativeObj int64,tvec1_nativeObj int64,rvec2_nativeObj int64,tvec2_nativeObj int64,rvec3_nativeObj int64,tvec3_nativeObj int64) {
	C.Java_org_opencv_calib3d_Calib3d_composeRT_11(clzEnv,clzObj,(C.jlong)(rvec1_nativeObj),(C.jlong)(tvec1_nativeObj),(C.jlong)(rvec2_nativeObj),(C.jlong)(tvec2_nativeObj),(C.jlong)(rvec3_nativeObj),(C.jlong)(tvec3_nativeObj))
}

func Calib3dNative_computeCorrespondEpilines_0(points_nativeObj int64,whichImage int,F_nativeObj int64,lines_nativeObj int64) {
	C.Java_org_opencv_calib3d_Calib3d_computeCorrespondEpilines_10(clzEnv,clzObj,(C.jlong)(points_nativeObj),(C.jint)(whichImage),(C.jlong)(F_nativeObj),(C.jlong)(lines_nativeObj))
}

func Calib3dNative_convertPointsFromHomogeneous_0(src_nativeObj int64,dst_nativeObj int64) {
	C.Java_org_opencv_calib3d_Calib3d_convertPointsFromHomogeneous_10(clzEnv,clzObj,(C.jlong)(src_nativeObj),(C.jlong)(dst_nativeObj))
}

func Calib3dNative_convertPointsToHomogeneous_0(src_nativeObj int64,dst_nativeObj int64) {
	C.Java_org_opencv_calib3d_Calib3d_convertPointsToHomogeneous_10(clzEnv,clzObj,(C.jlong)(src_nativeObj),(C.jlong)(dst_nativeObj))
}

func Calib3dNative_correctMatches_0(F_nativeObj int64,points1_nativeObj int64,points2_nativeObj int64,newPoints1_nativeObj int64,newPoints2_nativeObj int64) {
	C.Java_org_opencv_calib3d_Calib3d_correctMatches_10(clzEnv,clzObj,(C.jlong)(F_nativeObj),(C.jlong)(points1_nativeObj),(C.jlong)(points2_nativeObj),(C.jlong)(newPoints1_nativeObj),(C.jlong)(newPoints2_nativeObj))
}

func Calib3dNative_decomposeEssentialMat_0(E_nativeObj int64,R1_nativeObj int64,R2_nativeObj int64,t_nativeObj int64) {
	C.Java_org_opencv_calib3d_Calib3d_decomposeEssentialMat_10(clzEnv,clzObj,(C.jlong)(E_nativeObj),(C.jlong)(R1_nativeObj),(C.jlong)(R2_nativeObj),(C.jlong)(t_nativeObj))
}

func Calib3dNative_decomposeHomographyMat_0(H_nativeObj int64,K_nativeObj int64,rotations_mat_nativeObj int64,translations_mat_nativeObj int64,normals_mat_nativeObj int64) int {
	return int(C.Java_org_opencv_calib3d_Calib3d_decomposeHomographyMat_10(clzEnv,clzObj,(C.jlong)(H_nativeObj),(C.jlong)(K_nativeObj),(C.jlong)(rotations_mat_nativeObj),(C.jlong)(translations_mat_nativeObj),(C.jlong)(normals_mat_nativeObj)))
}

func Calib3dNative_decomposeProjectionMatrix_0(projMatrix_nativeObj int64,cameraMatrix_nativeObj int64,rotMatrix_nativeObj int64,transVect_nativeObj int64,rotMatrixX_nativeObj int64,rotMatrixY_nativeObj int64,rotMatrixZ_nativeObj int64,eulerAngles_nativeObj int64) {
	C.Java_org_opencv_calib3d_Calib3d_decomposeProjectionMatrix_10(clzEnv,clzObj,(C.jlong)(projMatrix_nativeObj),(C.jlong)(cameraMatrix_nativeObj),(C.jlong)(rotMatrix_nativeObj),(C.jlong)(transVect_nativeObj),(C.jlong)(rotMatrixX_nativeObj),(C.jlong)(rotMatrixY_nativeObj),(C.jlong)(rotMatrixZ_nativeObj),(C.jlong)(eulerAngles_nativeObj))
}

func Calib3dNative_decomposeProjectionMatrix_1(projMatrix_nativeObj int64,cameraMatrix_nativeObj int64,rotMatrix_nativeObj int64,transVect_nativeObj int64) {
	C.Java_org_opencv_calib3d_Calib3d_decomposeProjectionMatrix_11(clzEnv,clzObj,(C.jlong)(projMatrix_nativeObj),(C.jlong)(cameraMatrix_nativeObj),(C.jlong)(rotMatrix_nativeObj),(C.jlong)(transVect_nativeObj))
}

func Calib3dNative_distortPoints_0(undistorted_nativeObj int64,distorted_nativeObj int64,K_nativeObj int64,D_nativeObj int64,alpha float64) {
	C.Java_org_opencv_calib3d_Calib3d_distortPoints_10(clzEnv,clzObj,(C.jlong)(undistorted_nativeObj),(C.jlong)(distorted_nativeObj),(C.jlong)(K_nativeObj),(C.jlong)(D_nativeObj),(C.jdouble)(alpha))
}

func Calib3dNative_distortPoints_1(undistorted_nativeObj int64,distorted_nativeObj int64,K_nativeObj int64,D_nativeObj int64) {
	C.Java_org_opencv_calib3d_Calib3d_distortPoints_11(clzEnv,clzObj,(C.jlong)(undistorted_nativeObj),(C.jlong)(distorted_nativeObj),(C.jlong)(K_nativeObj),(C.jlong)(D_nativeObj))
}

func Calib3dNative_drawChessboardCorners_0(image_nativeObj int64,patternSize_width float64,patternSize_height float64,corners_mat_nativeObj int64,patternWasFound bool) {
	C.Java_org_opencv_calib3d_Calib3d_drawChessboardCorners_10(clzEnv,clzObj,(C.jlong)(image_nativeObj),(C.jdouble)(patternSize_width),(C.jdouble)(patternSize_height),(C.jlong)(corners_mat_nativeObj),tojboolean(patternWasFound))
}

func Calib3dNative_estimateAffine2D_0(from_nativeObj int64,to_nativeObj int64,inliers_nativeObj int64,method int,ransacReprojThreshold float64,maxIters int64,confidence float64,refineIters int64) int64 {
	return int64(C.Java_org_opencv_calib3d_Calib3d_estimateAffine2D_10(clzEnv,clzObj,(C.jlong)(from_nativeObj),(C.jlong)(to_nativeObj),(C.jlong)(inliers_nativeObj),(C.jint)(method),(C.jdouble)(ransacReprojThreshold),(C.jlong)(maxIters),(C.jdouble)(confidence),(C.jlong)(refineIters)))
}

func Calib3dNative_estimateAffine2D_1(from_nativeObj int64,to_nativeObj int64) int64 {
	return int64(C.Java_org_opencv_calib3d_Calib3d_estimateAffine2D_11(clzEnv,clzObj,(C.jlong)(from_nativeObj),(C.jlong)(to_nativeObj)))
}

func Calib3dNative_estimateAffine3D_0(src_nativeObj int64,dst_nativeObj int64,out_nativeObj int64,inliers_nativeObj int64,ransacThreshold float64,confidence float64) int {
	return int(C.Java_org_opencv_calib3d_Calib3d_estimateAffine3D_10(clzEnv,clzObj,(C.jlong)(src_nativeObj),(C.jlong)(dst_nativeObj),(C.jlong)(out_nativeObj),(C.jlong)(inliers_nativeObj),(C.jdouble)(ransacThreshold),(C.jdouble)(confidence)))
}

func Calib3dNative_estimateAffine3D_1(src_nativeObj int64,dst_nativeObj int64,out_nativeObj int64,inliers_nativeObj int64) int {
	return int(C.Java_org_opencv_calib3d_Calib3d_estimateAffine3D_11(clzEnv,clzObj,(C.jlong)(src_nativeObj),(C.jlong)(dst_nativeObj),(C.jlong)(out_nativeObj),(C.jlong)(inliers_nativeObj)))
}

func Calib3dNative_estimateAffinePartial2D_0(from_nativeObj int64,to_nativeObj int64,inliers_nativeObj int64,method int,ransacReprojThreshold float64,maxIters int64,confidence float64,refineIters int64) int64 {
	return int64(C.Java_org_opencv_calib3d_Calib3d_estimateAffinePartial2D_10(clzEnv,clzObj,(C.jlong)(from_nativeObj),(C.jlong)(to_nativeObj),(C.jlong)(inliers_nativeObj),(C.jint)(method),(C.jdouble)(ransacReprojThreshold),(C.jlong)(maxIters),(C.jdouble)(confidence),(C.jlong)(refineIters)))
}

func Calib3dNative_estimateAffinePartial2D_1(from_nativeObj int64,to_nativeObj int64) int64 {
	return int64(C.Java_org_opencv_calib3d_Calib3d_estimateAffinePartial2D_11(clzEnv,clzObj,(C.jlong)(from_nativeObj),(C.jlong)(to_nativeObj)))
}

func Calib3dNative_estimateNewCameraMatrixForUndistortRectify_0(K_nativeObj int64,D_nativeObj int64,image_size_width float64,image_size_height float64,R_nativeObj int64,P_nativeObj int64,balance float64,new_size_width float64,new_size_height float64,fov_scale float64) {
	C.Java_org_opencv_calib3d_Calib3d_estimateNewCameraMatrixForUndistortRectify_10(clzEnv,clzObj,(C.jlong)(K_nativeObj),(C.jlong)(D_nativeObj),(C.jdouble)(image_size_width),(C.jdouble)(image_size_height),(C.jlong)(R_nativeObj),(C.jlong)(P_nativeObj),(C.jdouble)(balance),(C.jdouble)(new_size_width),(C.jdouble)(new_size_height),(C.jdouble)(fov_scale))
}

func Calib3dNative_estimateNewCameraMatrixForUndistortRectify_1(K_nativeObj int64,D_nativeObj int64,image_size_width float64,image_size_height float64,R_nativeObj int64,P_nativeObj int64) {
	C.Java_org_opencv_calib3d_Calib3d_estimateNewCameraMatrixForUndistortRectify_11(clzEnv,clzObj,(C.jlong)(K_nativeObj),(C.jlong)(D_nativeObj),(C.jdouble)(image_size_width),(C.jdouble)(image_size_height),(C.jlong)(R_nativeObj),(C.jlong)(P_nativeObj))
}

func Calib3dNative_filterSpeckles_0(img_nativeObj int64,newVal float64,maxSpeckleSize int,maxDiff float64,buf_nativeObj int64) {
	C.Java_org_opencv_calib3d_Calib3d_filterSpeckles_10(clzEnv,clzObj,(C.jlong)(img_nativeObj),(C.jdouble)(newVal),(C.jint)(maxSpeckleSize),(C.jdouble)(maxDiff),(C.jlong)(buf_nativeObj))
}

func Calib3dNative_filterSpeckles_1(img_nativeObj int64,newVal float64,maxSpeckleSize int,maxDiff float64) {
	C.Java_org_opencv_calib3d_Calib3d_filterSpeckles_11(clzEnv,clzObj,(C.jlong)(img_nativeObj),(C.jdouble)(newVal),(C.jint)(maxSpeckleSize),(C.jdouble)(maxDiff))
}

func Calib3dNative_findChessboardCorners_0(image_nativeObj int64,patternSize_width float64,patternSize_height float64,corners_mat_nativeObj int64,flags int) bool {
	return togobool(C.Java_org_opencv_calib3d_Calib3d_findChessboardCorners_10(clzEnv,clzObj,(C.jlong)(image_nativeObj),(C.jdouble)(patternSize_width),(C.jdouble)(patternSize_height),(C.jlong)(corners_mat_nativeObj),(C.jint)(flags)))
}

func Calib3dNative_findChessboardCorners_1(image_nativeObj int64,patternSize_width float64,patternSize_height float64,corners_mat_nativeObj int64) bool {
	return togobool(C.Java_org_opencv_calib3d_Calib3d_findChessboardCorners_11(clzEnv,clzObj,(C.jlong)(image_nativeObj),(C.jdouble)(patternSize_width),(C.jdouble)(patternSize_height),(C.jlong)(corners_mat_nativeObj)))
}

func Calib3dNative_findCirclesGrid_0(image_nativeObj int64,patternSize_width float64,patternSize_height float64,centers_nativeObj int64,flags int) bool {
	return togobool(C.Java_org_opencv_calib3d_Calib3d_findCirclesGrid_10(clzEnv,clzObj,(C.jlong)(image_nativeObj),(C.jdouble)(patternSize_width),(C.jdouble)(patternSize_height),(C.jlong)(centers_nativeObj),(C.jint)(flags)))
}

func Calib3dNative_findCirclesGrid_1(image_nativeObj int64,patternSize_width float64,patternSize_height float64,centers_nativeObj int64) bool {
	return togobool(C.Java_org_opencv_calib3d_Calib3d_findCirclesGrid_11(clzEnv,clzObj,(C.jlong)(image_nativeObj),(C.jdouble)(patternSize_width),(C.jdouble)(patternSize_height),(C.jlong)(centers_nativeObj)))
}

func Calib3dNative_findEssentialMat_0(points1_nativeObj int64,points2_nativeObj int64,cameraMatrix_nativeObj int64,method int,prob float64,threshold float64,mask_nativeObj int64) int64 {
	return int64(C.Java_org_opencv_calib3d_Calib3d_findEssentialMat_10(clzEnv,clzObj,(C.jlong)(points1_nativeObj),(C.jlong)(points2_nativeObj),(C.jlong)(cameraMatrix_nativeObj),(C.jint)(method),(C.jdouble)(prob),(C.jdouble)(threshold),(C.jlong)(mask_nativeObj)))
}

func Calib3dNative_findEssentialMat_1(points1_nativeObj int64,points2_nativeObj int64,cameraMatrix_nativeObj int64,method int,prob float64,threshold float64) int64 {
	return int64(C.Java_org_opencv_calib3d_Calib3d_findEssentialMat_11(clzEnv,clzObj,(C.jlong)(points1_nativeObj),(C.jlong)(points2_nativeObj),(C.jlong)(cameraMatrix_nativeObj),(C.jint)(method),(C.jdouble)(prob),(C.jdouble)(threshold)))
}

func Calib3dNative_findEssentialMat_2(points1_nativeObj int64,points2_nativeObj int64,cameraMatrix_nativeObj int64) int64 {
	return int64(C.Java_org_opencv_calib3d_Calib3d_findEssentialMat_12(clzEnv,clzObj,(C.jlong)(points1_nativeObj),(C.jlong)(points2_nativeObj),(C.jlong)(cameraMatrix_nativeObj)))
}

func Calib3dNative_findEssentialMat_3(points1_nativeObj int64,points2_nativeObj int64,focal float64,pp_x float64,pp_y float64,method int,prob float64,threshold float64,mask_nativeObj int64) int64 {
	return int64(C.Java_org_opencv_calib3d_Calib3d_findEssentialMat_13(clzEnv,clzObj,(C.jlong)(points1_nativeObj),(C.jlong)(points2_nativeObj),(C.jdouble)(focal),(C.jdouble)(pp_x),(C.jdouble)(pp_y),(C.jint)(method),(C.jdouble)(prob),(C.jdouble)(threshold),(C.jlong)(mask_nativeObj)))
}

func Calib3dNative_findEssentialMat_4(points1_nativeObj int64,points2_nativeObj int64,focal float64,pp_x float64,pp_y float64,method int,prob float64,threshold float64) int64 {
	return int64(C.Java_org_opencv_calib3d_Calib3d_findEssentialMat_14(clzEnv,clzObj,(C.jlong)(points1_nativeObj),(C.jlong)(points2_nativeObj),(C.jdouble)(focal),(C.jdouble)(pp_x),(C.jdouble)(pp_y),(C.jint)(method),(C.jdouble)(prob),(C.jdouble)(threshold)))
}

func Calib3dNative_findEssentialMat_5(points1_nativeObj int64,points2_nativeObj int64) int64 {
	return int64(C.Java_org_opencv_calib3d_Calib3d_findEssentialMat_15(clzEnv,clzObj,(C.jlong)(points1_nativeObj),(C.jlong)(points2_nativeObj)))
}

func Calib3dNative_findFundamentalMat_0(points1_mat_nativeObj int64,points2_mat_nativeObj int64,method int,param1 float64,param2 float64,mask_nativeObj int64) int64 {
	return int64(C.Java_org_opencv_calib3d_Calib3d_findFundamentalMat_10(clzEnv,clzObj,(C.jlong)(points1_mat_nativeObj),(C.jlong)(points2_mat_nativeObj),(C.jint)(method),(C.jdouble)(param1),(C.jdouble)(param2),(C.jlong)(mask_nativeObj)))
}

func Calib3dNative_findFundamentalMat_1(points1_mat_nativeObj int64,points2_mat_nativeObj int64,method int,param1 float64,param2 float64) int64 {
	return int64(C.Java_org_opencv_calib3d_Calib3d_findFundamentalMat_11(clzEnv,clzObj,(C.jlong)(points1_mat_nativeObj),(C.jlong)(points2_mat_nativeObj),(C.jint)(method),(C.jdouble)(param1),(C.jdouble)(param2)))
}

func Calib3dNative_findFundamentalMat_2(points1_mat_nativeObj int64,points2_mat_nativeObj int64) int64 {
	return int64(C.Java_org_opencv_calib3d_Calib3d_findFundamentalMat_12(clzEnv,clzObj,(C.jlong)(points1_mat_nativeObj),(C.jlong)(points2_mat_nativeObj)))
}

func Calib3dNative_findHomography_0(srcPoints_mat_nativeObj int64,dstPoints_mat_nativeObj int64,method int,ransacReprojThreshold float64,mask_nativeObj int64,maxIters int,confidence float64) int64 {
	return int64(C.Java_org_opencv_calib3d_Calib3d_findHomography_10(clzEnv,clzObj,(C.jlong)(srcPoints_mat_nativeObj),(C.jlong)(dstPoints_mat_nativeObj),(C.jint)(method),(C.jdouble)(ransacReprojThreshold),(C.jlong)(mask_nativeObj),(C.jint)(maxIters),(C.jdouble)(confidence)))
}

func Calib3dNative_findHomography_1(srcPoints_mat_nativeObj int64,dstPoints_mat_nativeObj int64,method int,ransacReprojThreshold float64) int64 {
	return int64(C.Java_org_opencv_calib3d_Calib3d_findHomography_11(clzEnv,clzObj,(C.jlong)(srcPoints_mat_nativeObj),(C.jlong)(dstPoints_mat_nativeObj),(C.jint)(method),(C.jdouble)(ransacReprojThreshold)))
}

func Calib3dNative_findHomography_2(srcPoints_mat_nativeObj int64,dstPoints_mat_nativeObj int64) int64 {
	return int64(C.Java_org_opencv_calib3d_Calib3d_findHomography_12(clzEnv,clzObj,(C.jlong)(srcPoints_mat_nativeObj),(C.jlong)(dstPoints_mat_nativeObj)))
}

func Calib3dNative_getOptimalNewCameraMatrix_0(cameraMatrix_nativeObj int64,distCoeffs_nativeObj int64,imageSize_width float64,imageSize_height float64,alpha float64,newImgSize_width float64,newImgSize_height float64,validPixROI_out []float64,centerPrincipalPoint bool) int64 {
	return int64(C.Java_org_opencv_calib3d_Calib3d_getOptimalNewCameraMatrix_10(clzEnv,clzObj,(C.jlong)(cameraMatrix_nativeObj),(C.jlong)(distCoeffs_nativeObj),(C.jdouble)(imageSize_width),(C.jdouble)(imageSize_height),(C.jdouble)(alpha),(C.jdouble)(newImgSize_width),(C.jdouble)(newImgSize_height),tojdoubleArray(validPixROI_out),tojboolean(centerPrincipalPoint)))
}

func Calib3dNative_getOptimalNewCameraMatrix_1(cameraMatrix_nativeObj int64,distCoeffs_nativeObj int64,imageSize_width float64,imageSize_height float64,alpha float64) int64 {
	return int64(C.Java_org_opencv_calib3d_Calib3d_getOptimalNewCameraMatrix_11(clzEnv,clzObj,(C.jlong)(cameraMatrix_nativeObj),(C.jlong)(distCoeffs_nativeObj),(C.jdouble)(imageSize_width),(C.jdouble)(imageSize_height),(C.jdouble)(alpha)))
}

func Calib3dNative_getValidDisparityROI_0(roi1_x int,roi1_y int,roi1_width int,roi1_height int,roi2_x int,roi2_y int,roi2_width int,roi2_height int,minDisparity int,numberOfDisparities int,SADWindowSize int) []float64 {
	return togofloat64Array(C.Java_org_opencv_calib3d_Calib3d_getValidDisparityROI_10(clzEnv,clzObj,(C.jint)(roi1_x),(C.jint)(roi1_y),(C.jint)(roi1_width),(C.jint)(roi1_height),(C.jint)(roi2_x),(C.jint)(roi2_y),(C.jint)(roi2_width),(C.jint)(roi2_height),(C.jint)(minDisparity),(C.jint)(numberOfDisparities),(C.jint)(SADWindowSize)))
}

func Calib3dNative_initCameraMatrix2D_0(objectPoints_mat_nativeObj int64,imagePoints_mat_nativeObj int64,imageSize_width float64,imageSize_height float64,aspectRatio float64) int64 {
	return int64(C.Java_org_opencv_calib3d_Calib3d_initCameraMatrix2D_10(clzEnv,clzObj,(C.jlong)(objectPoints_mat_nativeObj),(C.jlong)(imagePoints_mat_nativeObj),(C.jdouble)(imageSize_width),(C.jdouble)(imageSize_height),(C.jdouble)(aspectRatio)))
}

func Calib3dNative_initCameraMatrix2D_1(objectPoints_mat_nativeObj int64,imagePoints_mat_nativeObj int64,imageSize_width float64,imageSize_height float64) int64 {
	return int64(C.Java_org_opencv_calib3d_Calib3d_initCameraMatrix2D_11(clzEnv,clzObj,(C.jlong)(objectPoints_mat_nativeObj),(C.jlong)(imagePoints_mat_nativeObj),(C.jdouble)(imageSize_width),(C.jdouble)(imageSize_height)))
}

func Calib3dNative_initUndistortRectifyMap_0(K_nativeObj int64,D_nativeObj int64,R_nativeObj int64,P_nativeObj int64,size_width float64,size_height float64,m1type int,map1_nativeObj int64,map2_nativeObj int64) {
	C.Java_org_opencv_calib3d_Calib3d_initUndistortRectifyMap_10(clzEnv,clzObj,(C.jlong)(K_nativeObj),(C.jlong)(D_nativeObj),(C.jlong)(R_nativeObj),(C.jlong)(P_nativeObj),(C.jdouble)(size_width),(C.jdouble)(size_height),(C.jint)(m1type),(C.jlong)(map1_nativeObj),(C.jlong)(map2_nativeObj))
}

func Calib3dNative_matMulDeriv_0(A_nativeObj int64,B_nativeObj int64,dABdA_nativeObj int64,dABdB_nativeObj int64) {
	C.Java_org_opencv_calib3d_Calib3d_matMulDeriv_10(clzEnv,clzObj,(C.jlong)(A_nativeObj),(C.jlong)(B_nativeObj),(C.jlong)(dABdA_nativeObj),(C.jlong)(dABdB_nativeObj))
}

func Calib3dNative_projectPoints_0(objectPoints_mat_nativeObj int64,rvec_nativeObj int64,tvec_nativeObj int64,cameraMatrix_nativeObj int64,distCoeffs_mat_nativeObj int64,imagePoints_mat_nativeObj int64,jacobian_nativeObj int64,aspectRatio float64) {
	C.Java_org_opencv_calib3d_Calib3d_projectPoints_10(clzEnv,clzObj,(C.jlong)(objectPoints_mat_nativeObj),(C.jlong)(rvec_nativeObj),(C.jlong)(tvec_nativeObj),(C.jlong)(cameraMatrix_nativeObj),(C.jlong)(distCoeffs_mat_nativeObj),(C.jlong)(imagePoints_mat_nativeObj),(C.jlong)(jacobian_nativeObj),(C.jdouble)(aspectRatio))
}

func Calib3dNative_projectPoints_1(objectPoints_mat_nativeObj int64,rvec_nativeObj int64,tvec_nativeObj int64,cameraMatrix_nativeObj int64,distCoeffs_mat_nativeObj int64,imagePoints_mat_nativeObj int64) {
	C.Java_org_opencv_calib3d_Calib3d_projectPoints_11(clzEnv,clzObj,(C.jlong)(objectPoints_mat_nativeObj),(C.jlong)(rvec_nativeObj),(C.jlong)(tvec_nativeObj),(C.jlong)(cameraMatrix_nativeObj),(C.jlong)(distCoeffs_mat_nativeObj),(C.jlong)(imagePoints_mat_nativeObj))
}

func Calib3dNative_projectPoints_2(objectPoints_mat_nativeObj int64,imagePoints_mat_nativeObj int64,rvec_nativeObj int64,tvec_nativeObj int64,K_nativeObj int64,D_nativeObj int64,alpha float64,jacobian_nativeObj int64) {
	C.Java_org_opencv_calib3d_Calib3d_projectPoints_12(clzEnv,clzObj,(C.jlong)(objectPoints_mat_nativeObj),(C.jlong)(imagePoints_mat_nativeObj),(C.jlong)(rvec_nativeObj),(C.jlong)(tvec_nativeObj),(C.jlong)(K_nativeObj),(C.jlong)(D_nativeObj),(C.jdouble)(alpha),(C.jlong)(jacobian_nativeObj))
}

func Calib3dNative_projectPoints_3(objectPoints_mat_nativeObj int64,imagePoints_mat_nativeObj int64,rvec_nativeObj int64,tvec_nativeObj int64,K_nativeObj int64,D_nativeObj int64) {
	C.Java_org_opencv_calib3d_Calib3d_projectPoints_13(clzEnv,clzObj,(C.jlong)(objectPoints_mat_nativeObj),(C.jlong)(imagePoints_mat_nativeObj),(C.jlong)(rvec_nativeObj),(C.jlong)(tvec_nativeObj),(C.jlong)(K_nativeObj),(C.jlong)(D_nativeObj))
}

func Calib3dNative_recoverPose_0(E_nativeObj int64,points1_nativeObj int64,points2_nativeObj int64,R_nativeObj int64,t_nativeObj int64,focal float64,pp_x float64,pp_y float64,mask_nativeObj int64) int {
	return int(C.Java_org_opencv_calib3d_Calib3d_recoverPose_10(clzEnv,clzObj,(C.jlong)(E_nativeObj),(C.jlong)(points1_nativeObj),(C.jlong)(points2_nativeObj),(C.jlong)(R_nativeObj),(C.jlong)(t_nativeObj),(C.jdouble)(focal),(C.jdouble)(pp_x),(C.jdouble)(pp_y),(C.jlong)(mask_nativeObj)))
}

func Calib3dNative_recoverPose_1(E_nativeObj int64,points1_nativeObj int64,points2_nativeObj int64,R_nativeObj int64,t_nativeObj int64,focal float64,pp_x float64,pp_y float64) int {
	return int(C.Java_org_opencv_calib3d_Calib3d_recoverPose_11(clzEnv,clzObj,(C.jlong)(E_nativeObj),(C.jlong)(points1_nativeObj),(C.jlong)(points2_nativeObj),(C.jlong)(R_nativeObj),(C.jlong)(t_nativeObj),(C.jdouble)(focal),(C.jdouble)(pp_x),(C.jdouble)(pp_y)))
}

func Calib3dNative_recoverPose_2(E_nativeObj int64,points1_nativeObj int64,points2_nativeObj int64,R_nativeObj int64,t_nativeObj int64) int {
	return int(C.Java_org_opencv_calib3d_Calib3d_recoverPose_12(clzEnv,clzObj,(C.jlong)(E_nativeObj),(C.jlong)(points1_nativeObj),(C.jlong)(points2_nativeObj),(C.jlong)(R_nativeObj),(C.jlong)(t_nativeObj)))
}

func Calib3dNative_recoverPose_3(E_nativeObj int64,points1_nativeObj int64,points2_nativeObj int64,cameraMatrix_nativeObj int64,R_nativeObj int64,t_nativeObj int64,mask_nativeObj int64) int {
	return int(C.Java_org_opencv_calib3d_Calib3d_recoverPose_13(clzEnv,clzObj,(C.jlong)(E_nativeObj),(C.jlong)(points1_nativeObj),(C.jlong)(points2_nativeObj),(C.jlong)(cameraMatrix_nativeObj),(C.jlong)(R_nativeObj),(C.jlong)(t_nativeObj),(C.jlong)(mask_nativeObj)))
}

func Calib3dNative_recoverPose_4(E_nativeObj int64,points1_nativeObj int64,points2_nativeObj int64,cameraMatrix_nativeObj int64,R_nativeObj int64,t_nativeObj int64) int {
	return int(C.Java_org_opencv_calib3d_Calib3d_recoverPose_14(clzEnv,clzObj,(C.jlong)(E_nativeObj),(C.jlong)(points1_nativeObj),(C.jlong)(points2_nativeObj),(C.jlong)(cameraMatrix_nativeObj),(C.jlong)(R_nativeObj),(C.jlong)(t_nativeObj)))
}

func Calib3dNative_recoverPose_5(E_nativeObj int64,points1_nativeObj int64,points2_nativeObj int64,cameraMatrix_nativeObj int64,R_nativeObj int64,t_nativeObj int64,distanceThresh float64,mask_nativeObj int64,triangulatedPoints_nativeObj int64) int {
	return int(C.Java_org_opencv_calib3d_Calib3d_recoverPose_15(clzEnv,clzObj,(C.jlong)(E_nativeObj),(C.jlong)(points1_nativeObj),(C.jlong)(points2_nativeObj),(C.jlong)(cameraMatrix_nativeObj),(C.jlong)(R_nativeObj),(C.jlong)(t_nativeObj),(C.jdouble)(distanceThresh),(C.jlong)(mask_nativeObj),(C.jlong)(triangulatedPoints_nativeObj)))
}

func Calib3dNative_recoverPose_6(E_nativeObj int64,points1_nativeObj int64,points2_nativeObj int64,cameraMatrix_nativeObj int64,R_nativeObj int64,t_nativeObj int64,distanceThresh float64) int {
	return int(C.Java_org_opencv_calib3d_Calib3d_recoverPose_16(clzEnv,clzObj,(C.jlong)(E_nativeObj),(C.jlong)(points1_nativeObj),(C.jlong)(points2_nativeObj),(C.jlong)(cameraMatrix_nativeObj),(C.jlong)(R_nativeObj),(C.jlong)(t_nativeObj),(C.jdouble)(distanceThresh)))
}

func Calib3dNative_rectify3Collinear_0(cameraMatrix1_nativeObj int64,distCoeffs1_nativeObj int64,cameraMatrix2_nativeObj int64,distCoeffs2_nativeObj int64,cameraMatrix3_nativeObj int64,distCoeffs3_nativeObj int64,imgpt1_mat_nativeObj int64,imgpt3_mat_nativeObj int64,imageSize_width float64,imageSize_height float64,R12_nativeObj int64,T12_nativeObj int64,R13_nativeObj int64,T13_nativeObj int64,R1_nativeObj int64,R2_nativeObj int64,R3_nativeObj int64,P1_nativeObj int64,P2_nativeObj int64,P3_nativeObj int64,Q_nativeObj int64,alpha float64,newImgSize_width float64,newImgSize_height float64,roi1_out []float64,roi2_out []float64,flags int) float32 {
	return float32(C.Java_org_opencv_calib3d_Calib3d_rectify3Collinear_10(clzEnv,clzObj,(C.jlong)(cameraMatrix1_nativeObj),(C.jlong)(distCoeffs1_nativeObj),(C.jlong)(cameraMatrix2_nativeObj),(C.jlong)(distCoeffs2_nativeObj),(C.jlong)(cameraMatrix3_nativeObj),(C.jlong)(distCoeffs3_nativeObj),(C.jlong)(imgpt1_mat_nativeObj),(C.jlong)(imgpt3_mat_nativeObj),(C.jdouble)(imageSize_width),(C.jdouble)(imageSize_height),(C.jlong)(R12_nativeObj),(C.jlong)(T12_nativeObj),(C.jlong)(R13_nativeObj),(C.jlong)(T13_nativeObj),(C.jlong)(R1_nativeObj),(C.jlong)(R2_nativeObj),(C.jlong)(R3_nativeObj),(C.jlong)(P1_nativeObj),(C.jlong)(P2_nativeObj),(C.jlong)(P3_nativeObj),(C.jlong)(Q_nativeObj),(C.jdouble)(alpha),(C.jdouble)(newImgSize_width),(C.jdouble)(newImgSize_height),tojdoubleArray(roi1_out),tojdoubleArray(roi2_out),(C.jint)(flags)))
}

func Calib3dNative_reprojectImageTo3D_0(disparity_nativeObj int64,_3dImage_nativeObj int64,Q_nativeObj int64,handleMissingValues bool,ddepth int) {
	C.Java_org_opencv_calib3d_Calib3d_reprojectImageTo3D_10(clzEnv,clzObj,(C.jlong)(disparity_nativeObj),(C.jlong)(_3dImage_nativeObj),(C.jlong)(Q_nativeObj),tojboolean(handleMissingValues),(C.jint)(ddepth))
}

func Calib3dNative_reprojectImageTo3D_1(disparity_nativeObj int64,_3dImage_nativeObj int64,Q_nativeObj int64,handleMissingValues bool) {
	C.Java_org_opencv_calib3d_Calib3d_reprojectImageTo3D_11(clzEnv,clzObj,(C.jlong)(disparity_nativeObj),(C.jlong)(_3dImage_nativeObj),(C.jlong)(Q_nativeObj),tojboolean(handleMissingValues))
}

func Calib3dNative_reprojectImageTo3D_2(disparity_nativeObj int64,_3dImage_nativeObj int64,Q_nativeObj int64) {
	C.Java_org_opencv_calib3d_Calib3d_reprojectImageTo3D_12(clzEnv,clzObj,(C.jlong)(disparity_nativeObj),(C.jlong)(_3dImage_nativeObj),(C.jlong)(Q_nativeObj))
}

func Calib3dNative_sampsonDistance_0(pt1_nativeObj int64,pt2_nativeObj int64,F_nativeObj int64) float64 {
	return float64(C.Java_org_opencv_calib3d_Calib3d_sampsonDistance_10(clzEnv,clzObj,(C.jlong)(pt1_nativeObj),(C.jlong)(pt2_nativeObj),(C.jlong)(F_nativeObj)))
}

func Calib3dNative_solveP3P_0(objectPoints_nativeObj int64,imagePoints_nativeObj int64,cameraMatrix_nativeObj int64,distCoeffs_nativeObj int64,rvecs_mat_nativeObj int64,tvecs_mat_nativeObj int64,flags int) int {
	return int(C.Java_org_opencv_calib3d_Calib3d_solveP3P_10(clzEnv,clzObj,(C.jlong)(objectPoints_nativeObj),(C.jlong)(imagePoints_nativeObj),(C.jlong)(cameraMatrix_nativeObj),(C.jlong)(distCoeffs_nativeObj),(C.jlong)(rvecs_mat_nativeObj),(C.jlong)(tvecs_mat_nativeObj),(C.jint)(flags)))
}

func Calib3dNative_solvePnPRansac_0(objectPoints_mat_nativeObj int64,imagePoints_mat_nativeObj int64,cameraMatrix_nativeObj int64,distCoeffs_mat_nativeObj int64,rvec_nativeObj int64,tvec_nativeObj int64,useExtrinsicGuess bool,iterationsCount int,reprojectionError float32,confidence float64,inliers_nativeObj int64,flags int) bool {
	return togobool(C.Java_org_opencv_calib3d_Calib3d_solvePnPRansac_10(clzEnv,clzObj,(C.jlong)(objectPoints_mat_nativeObj),(C.jlong)(imagePoints_mat_nativeObj),(C.jlong)(cameraMatrix_nativeObj),(C.jlong)(distCoeffs_mat_nativeObj),(C.jlong)(rvec_nativeObj),(C.jlong)(tvec_nativeObj),tojboolean(useExtrinsicGuess),(C.jint)(iterationsCount),(C.jfloat)(reprojectionError),(C.jdouble)(confidence),(C.jlong)(inliers_nativeObj),(C.jint)(flags)))
}

func Calib3dNative_solvePnPRansac_1(objectPoints_mat_nativeObj int64,imagePoints_mat_nativeObj int64,cameraMatrix_nativeObj int64,distCoeffs_mat_nativeObj int64,rvec_nativeObj int64,tvec_nativeObj int64) bool {
	return togobool(C.Java_org_opencv_calib3d_Calib3d_solvePnPRansac_11(clzEnv,clzObj,(C.jlong)(objectPoints_mat_nativeObj),(C.jlong)(imagePoints_mat_nativeObj),(C.jlong)(cameraMatrix_nativeObj),(C.jlong)(distCoeffs_mat_nativeObj),(C.jlong)(rvec_nativeObj),(C.jlong)(tvec_nativeObj)))
}

func Calib3dNative_solvePnP_0(objectPoints_mat_nativeObj int64,imagePoints_mat_nativeObj int64,cameraMatrix_nativeObj int64,distCoeffs_mat_nativeObj int64,rvec_nativeObj int64,tvec_nativeObj int64,useExtrinsicGuess bool,flags int) bool {
	return togobool(C.Java_org_opencv_calib3d_Calib3d_solvePnP_10(clzEnv,clzObj,(C.jlong)(objectPoints_mat_nativeObj),(C.jlong)(imagePoints_mat_nativeObj),(C.jlong)(cameraMatrix_nativeObj),(C.jlong)(distCoeffs_mat_nativeObj),(C.jlong)(rvec_nativeObj),(C.jlong)(tvec_nativeObj),tojboolean(useExtrinsicGuess),(C.jint)(flags)))
}

func Calib3dNative_solvePnP_1(objectPoints_mat_nativeObj int64,imagePoints_mat_nativeObj int64,cameraMatrix_nativeObj int64,distCoeffs_mat_nativeObj int64,rvec_nativeObj int64,tvec_nativeObj int64) bool {
	return togobool(C.Java_org_opencv_calib3d_Calib3d_solvePnP_11(clzEnv,clzObj,(C.jlong)(objectPoints_mat_nativeObj),(C.jlong)(imagePoints_mat_nativeObj),(C.jlong)(cameraMatrix_nativeObj),(C.jlong)(distCoeffs_mat_nativeObj),(C.jlong)(rvec_nativeObj),(C.jlong)(tvec_nativeObj)))
}

func Calib3dNative_stereoCalibrate_0(objectPoints_mat_nativeObj int64,imagePoints1_mat_nativeObj int64,imagePoints2_mat_nativeObj int64,cameraMatrix1_nativeObj int64,distCoeffs1_nativeObj int64,cameraMatrix2_nativeObj int64,distCoeffs2_nativeObj int64,imageSize_width float64,imageSize_height float64,R_nativeObj int64,T_nativeObj int64,E_nativeObj int64,F_nativeObj int64,flags int,criteria_type int,criteria_maxCount int,criteria_epsilon float64) float64 {
	return float64(C.Java_org_opencv_calib3d_Calib3d_stereoCalibrate_10(clzEnv,clzObj,(C.jlong)(objectPoints_mat_nativeObj),(C.jlong)(imagePoints1_mat_nativeObj),(C.jlong)(imagePoints2_mat_nativeObj),(C.jlong)(cameraMatrix1_nativeObj),(C.jlong)(distCoeffs1_nativeObj),(C.jlong)(cameraMatrix2_nativeObj),(C.jlong)(distCoeffs2_nativeObj),(C.jdouble)(imageSize_width),(C.jdouble)(imageSize_height),(C.jlong)(R_nativeObj),(C.jlong)(T_nativeObj),(C.jlong)(E_nativeObj),(C.jlong)(F_nativeObj),(C.jint)(flags),(C.jint)(criteria_type),(C.jint)(criteria_maxCount),(C.jdouble)(criteria_epsilon)))
}

func Calib3dNative_stereoCalibrate_1(objectPoints_mat_nativeObj int64,imagePoints1_mat_nativeObj int64,imagePoints2_mat_nativeObj int64,cameraMatrix1_nativeObj int64,distCoeffs1_nativeObj int64,cameraMatrix2_nativeObj int64,distCoeffs2_nativeObj int64,imageSize_width float64,imageSize_height float64,R_nativeObj int64,T_nativeObj int64,E_nativeObj int64,F_nativeObj int64,flags int) float64 {
	return float64(C.Java_org_opencv_calib3d_Calib3d_stereoCalibrate_11(clzEnv,clzObj,(C.jlong)(objectPoints_mat_nativeObj),(C.jlong)(imagePoints1_mat_nativeObj),(C.jlong)(imagePoints2_mat_nativeObj),(C.jlong)(cameraMatrix1_nativeObj),(C.jlong)(distCoeffs1_nativeObj),(C.jlong)(cameraMatrix2_nativeObj),(C.jlong)(distCoeffs2_nativeObj),(C.jdouble)(imageSize_width),(C.jdouble)(imageSize_height),(C.jlong)(R_nativeObj),(C.jlong)(T_nativeObj),(C.jlong)(E_nativeObj),(C.jlong)(F_nativeObj),(C.jint)(flags)))
}

func Calib3dNative_stereoCalibrate_2(objectPoints_mat_nativeObj int64,imagePoints1_mat_nativeObj int64,imagePoints2_mat_nativeObj int64,cameraMatrix1_nativeObj int64,distCoeffs1_nativeObj int64,cameraMatrix2_nativeObj int64,distCoeffs2_nativeObj int64,imageSize_width float64,imageSize_height float64,R_nativeObj int64,T_nativeObj int64,E_nativeObj int64,F_nativeObj int64) float64 {
	return float64(C.Java_org_opencv_calib3d_Calib3d_stereoCalibrate_12(clzEnv,clzObj,(C.jlong)(objectPoints_mat_nativeObj),(C.jlong)(imagePoints1_mat_nativeObj),(C.jlong)(imagePoints2_mat_nativeObj),(C.jlong)(cameraMatrix1_nativeObj),(C.jlong)(distCoeffs1_nativeObj),(C.jlong)(cameraMatrix2_nativeObj),(C.jlong)(distCoeffs2_nativeObj),(C.jdouble)(imageSize_width),(C.jdouble)(imageSize_height),(C.jlong)(R_nativeObj),(C.jlong)(T_nativeObj),(C.jlong)(E_nativeObj),(C.jlong)(F_nativeObj)))
}

func Calib3dNative_stereoCalibrate_3(objectPoints_mat_nativeObj int64,imagePoints1_mat_nativeObj int64,imagePoints2_mat_nativeObj int64,K1_nativeObj int64,D1_nativeObj int64,K2_nativeObj int64,D2_nativeObj int64,imageSize_width float64,imageSize_height float64,R_nativeObj int64,T_nativeObj int64,flags int,criteria_type int,criteria_maxCount int,criteria_epsilon float64) float64 {
	return float64(C.Java_org_opencv_calib3d_Calib3d_stereoCalibrate_13(clzEnv,clzObj,(C.jlong)(objectPoints_mat_nativeObj),(C.jlong)(imagePoints1_mat_nativeObj),(C.jlong)(imagePoints2_mat_nativeObj),(C.jlong)(K1_nativeObj),(C.jlong)(D1_nativeObj),(C.jlong)(K2_nativeObj),(C.jlong)(D2_nativeObj),(C.jdouble)(imageSize_width),(C.jdouble)(imageSize_height),(C.jlong)(R_nativeObj),(C.jlong)(T_nativeObj),(C.jint)(flags),(C.jint)(criteria_type),(C.jint)(criteria_maxCount),(C.jdouble)(criteria_epsilon)))
}

func Calib3dNative_stereoCalibrate_4(objectPoints_mat_nativeObj int64,imagePoints1_mat_nativeObj int64,imagePoints2_mat_nativeObj int64,K1_nativeObj int64,D1_nativeObj int64,K2_nativeObj int64,D2_nativeObj int64,imageSize_width float64,imageSize_height float64,R_nativeObj int64,T_nativeObj int64,flags int) float64 {
	return float64(C.Java_org_opencv_calib3d_Calib3d_stereoCalibrate_14(clzEnv,clzObj,(C.jlong)(objectPoints_mat_nativeObj),(C.jlong)(imagePoints1_mat_nativeObj),(C.jlong)(imagePoints2_mat_nativeObj),(C.jlong)(K1_nativeObj),(C.jlong)(D1_nativeObj),(C.jlong)(K2_nativeObj),(C.jlong)(D2_nativeObj),(C.jdouble)(imageSize_width),(C.jdouble)(imageSize_height),(C.jlong)(R_nativeObj),(C.jlong)(T_nativeObj),(C.jint)(flags)))
}

func Calib3dNative_stereoCalibrate_5(objectPoints_mat_nativeObj int64,imagePoints1_mat_nativeObj int64,imagePoints2_mat_nativeObj int64,K1_nativeObj int64,D1_nativeObj int64,K2_nativeObj int64,D2_nativeObj int64,imageSize_width float64,imageSize_height float64,R_nativeObj int64,T_nativeObj int64) float64 {
	return float64(C.Java_org_opencv_calib3d_Calib3d_stereoCalibrate_15(clzEnv,clzObj,(C.jlong)(objectPoints_mat_nativeObj),(C.jlong)(imagePoints1_mat_nativeObj),(C.jlong)(imagePoints2_mat_nativeObj),(C.jlong)(K1_nativeObj),(C.jlong)(D1_nativeObj),(C.jlong)(K2_nativeObj),(C.jlong)(D2_nativeObj),(C.jdouble)(imageSize_width),(C.jdouble)(imageSize_height),(C.jlong)(R_nativeObj),(C.jlong)(T_nativeObj)))
}

func Calib3dNative_stereoRectifyUncalibrated_0(points1_nativeObj int64,points2_nativeObj int64,F_nativeObj int64,imgSize_width float64,imgSize_height float64,H1_nativeObj int64,H2_nativeObj int64,threshold float64) bool {
	return togobool(C.Java_org_opencv_calib3d_Calib3d_stereoRectifyUncalibrated_10(clzEnv,clzObj,(C.jlong)(points1_nativeObj),(C.jlong)(points2_nativeObj),(C.jlong)(F_nativeObj),(C.jdouble)(
				imgSize_width),(C.jdouble)(imgSize_height),(C.jlong)(H1_nativeObj),(C.jlong)(H2_nativeObj),(C.jdouble)(threshold)))
}

func Calib3dNative_stereoRectifyUncalibrated_1(points1_nativeObj int64,points2_nativeObj int64,F_nativeObj int64,imgSize_width float64,imgSize_height float64,H1_nativeObj int64,H2_nativeObj int64) bool {
	return togobool(C.Java_org_opencv_calib3d_Calib3d_stereoRectifyUncalibrated_11(clzEnv,clzObj,(C.jlong)(points1_nativeObj),(C.jlong)(points2_nativeObj),(C.jlong)(F_nativeObj),(C.jdouble)(
				imgSize_width),(C.jdouble)(imgSize_height),(C.jlong)(H1_nativeObj),(C.jlong)(H2_nativeObj)))
}

func Calib3dNative_stereoRectify_0(cameraMatrix1_nativeObj int64,distCoeffs1_nativeObj int64,cameraMatrix2_nativeObj int64,distCoeffs2_nativeObj int64,imageSize_width float64,imageSize_height float64,R_nativeObj int64,T_nativeObj int64,R1_nativeObj int64,R2_nativeObj int64,P1_nativeObj int64,P2_nativeObj int64,Q_nativeObj int64,flags int,alpha float64,newImageSize_width float64,newImageSize_height float64,validPixROI1_out []float64,validPixROI2_out []float64) {
	C.Java_org_opencv_calib3d_Calib3d_stereoRectify_10(clzEnv,clzObj,(C.jlong)(cameraMatrix1_nativeObj),(C.jlong)(distCoeffs1_nativeObj),(C.jlong)(cameraMatrix2_nativeObj),(C.jlong)(distCoeffs2_nativeObj),(C.jdouble)(imageSize_width),(C.jdouble)(imageSize_height),(C.jlong)(R_nativeObj),(C.jlong)(T_nativeObj),(C.jlong)(R1_nativeObj),(C.jlong)(R2_nativeObj),(C.jlong)(P1_nativeObj),(C.jlong)(P2_nativeObj),(C.jlong)(Q_nativeObj),(C.jint)(flags),(C.jdouble)(alpha),(C.jdouble)(newImageSize_width),(C.jdouble)(newImageSize_height),tojdoubleArray(validPixROI1_out),tojdoubleArray(validPixROI2_out))
}

func Calib3dNative_stereoRectify_1(cameraMatrix1_nativeObj int64,distCoeffs1_nativeObj int64,cameraMatrix2_nativeObj int64,distCoeffs2_nativeObj int64,imageSize_width float64,imageSize_height float64,R_nativeObj int64,T_nativeObj int64,R1_nativeObj int64,R2_nativeObj int64,P1_nativeObj int64,P2_nativeObj int64,Q_nativeObj int64) {
	C.Java_org_opencv_calib3d_Calib3d_stereoRectify_11(clzEnv,clzObj,(C.jlong)(cameraMatrix1_nativeObj),(C.jlong)(distCoeffs1_nativeObj),(C.jlong)(cameraMatrix2_nativeObj),(C.jlong)(distCoeffs2_nativeObj),(C.jdouble)(imageSize_width),(C.jdouble)(imageSize_height),(C.jlong)(R_nativeObj),(C.jlong)(T_nativeObj),(C.jlong)(R1_nativeObj),(C.jlong)(R2_nativeObj),(C.jlong)(P1_nativeObj),(C.jlong)(P2_nativeObj),(C.jlong)(Q_nativeObj))
}

func Calib3dNative_stereoRectify_2(K1_nativeObj int64,D1_nativeObj int64,K2_nativeObj int64,D2_nativeObj int64,imageSize_width float64,imageSize_height float64,R_nativeObj int64,tvec_nativeObj int64,R1_nativeObj int64,R2_nativeObj int64,P1_nativeObj int64,P2_nativeObj int64,Q_nativeObj int64,flags int,newImageSize_width float64,newImageSize_height float64,balance float64,fov_scale float64) {
	C.Java_org_opencv_calib3d_Calib3d_stereoRectify_12(clzEnv,clzObj,(C.jlong)(K1_nativeObj),(C.jlong)(D1_nativeObj),(C.jlong)(K2_nativeObj),(C.jlong)(D2_nativeObj),(C.jdouble)(imageSize_width),(C.jdouble)(imageSize_height),(C.jlong)(R_nativeObj),(C.jlong)(tvec_nativeObj),(C.jlong)(R1_nativeObj),(C.jlong)(R2_nativeObj),(C.jlong)(P1_nativeObj),(C.jlong)(P2_nativeObj),(C.jlong)(Q_nativeObj),(C.jint)(flags),(C.jdouble)(newImageSize_width),(C.jdouble)(newImageSize_height),(C.jdouble)(balance),(C.jdouble)(fov_scale))
}

func Calib3dNative_stereoRectify_3(K1_nativeObj int64,D1_nativeObj int64,K2_nativeObj int64,D2_nativeObj int64,imageSize_width float64,imageSize_height float64,R_nativeObj int64,tvec_nativeObj int64,R1_nativeObj int64,R2_nativeObj int64,P1_nativeObj int64,P2_nativeObj int64,Q_nativeObj int64,flags int) {
	C.Java_org_opencv_calib3d_Calib3d_stereoRectify_13(clzEnv,clzObj,(C.jlong)(K1_nativeObj),(C.jlong)(D1_nativeObj),(C.jlong)(K2_nativeObj),(C.jlong)(D2_nativeObj),(C.jdouble)(imageSize_width),(C.jdouble)(imageSize_height),(C.jlong)(R_nativeObj),(C.jlong)(tvec_nativeObj),(C.jlong)(R1_nativeObj),(C.jlong)(R2_nativeObj),(C.jlong)(P1_nativeObj),(C.jlong)(P2_nativeObj),(C.jlong)(Q_nativeObj),(C.jint)(flags))
}

func Calib3dNative_triangulatePoints_0(projMatr1_nativeObj int64,projMatr2_nativeObj int64,projPoints1_nativeObj int64,projPoints2_nativeObj int64,points4D_nativeObj int64) {
	C.Java_org_opencv_calib3d_Calib3d_triangulatePoints_10(clzEnv,clzObj,(C.jlong)(projMatr1_nativeObj),(C.jlong)(projMatr2_nativeObj),(C.jlong)(projPoints1_nativeObj),(C.jlong)(projPoints2_nativeObj),(C.jlong)(points4D_nativeObj))
}

func Calib3dNative_undistortImage_0(distorted_nativeObj int64,undistorted_nativeObj int64,K_nativeObj int64,D_nativeObj int64,Knew_nativeObj int64,new_size_width float64,new_size_height float64) {
	C.Java_org_opencv_calib3d_Calib3d_undistortImage_10(clzEnv,clzObj,(C.jlong)(distorted_nativeObj),(C.jlong)(undistorted_nativeObj),(C.jlong)(K_nativeObj),(C.jlong)(D_nativeObj),(C.jlong)(Knew_nativeObj),(C.jdouble)(new_size_width),(C.jdouble)(new_size_height))
}

func Calib3dNative_undistortImage_1(distorted_nativeObj int64,undistorted_nativeObj int64,K_nativeObj int64,D_nativeObj int64) {
	C.Java_org_opencv_calib3d_Calib3d_undistortImage_11(clzEnv,clzObj,(C.jlong)(distorted_nativeObj),(C.jlong)(undistorted_nativeObj),(C.jlong)(K_nativeObj),(C.jlong)(D_nativeObj))
}

func Calib3dNative_undistortPoints_0(distorted_nativeObj int64,undistorted_nativeObj int64,K_nativeObj int64,D_nativeObj int64,R_nativeObj int64,P_nativeObj int64) {
	C.Java_org_opencv_calib3d_Calib3d_undistortPoints_10(clzEnv,clzObj,(C.jlong)(distorted_nativeObj),(C.jlong)(undistorted_nativeObj),(C.jlong)(K_nativeObj),(C.jlong)(D_nativeObj),(C.jlong)(R_nativeObj),(C.jlong)(P_nativeObj))
}

func Calib3dNative_undistortPoints_1(distorted_nativeObj int64,undistorted_nativeObj int64,K_nativeObj int64,D_nativeObj int64) {
	C.Java_org_opencv_calib3d_Calib3d_undistortPoints_11(clzEnv,clzObj,(C.jlong)(distorted_nativeObj),(C.jlong)(undistorted_nativeObj),(C.jlong)(K_nativeObj),(C.jlong)(D_nativeObj))
}

func Calib3dNative_validateDisparity_0(disparity_nativeObj int64,cost_nativeObj int64,minDisparity int,numberOfDisparities int,disp12MaxDisp int) {
	C.Java_org_opencv_calib3d_Calib3d_validateDisparity_10(clzEnv,clzObj,(C.jlong)(disparity_nativeObj),(C.jlong)(cost_nativeObj),(C.jint)(minDisparity),(C.jint)(numberOfDisparities),(C.jint)(disp12MaxDisp))
}

func Calib3dNative_validateDisparity_1(disparity_nativeObj int64,cost_nativeObj int64,minDisparity int,numberOfDisparities int) {
	C.Java_org_opencv_calib3d_Calib3d_validateDisparity_11(clzEnv,clzObj,(C.jlong)(disparity_nativeObj),(C.jlong)(cost_nativeObj),(C.jint)(minDisparity),(C.jint)(numberOfDisparities))
}

func StereoBMNative_create_0(numDisparities int,blockSize int) int64 {
	return int64(C.Java_org_opencv_calib3d_StereoBM_create_10(clzEnv,clzObj,(C.jint)(numDisparities),(C.jint)(blockSize)))
}

func StereoBMNative_create_1() int64 {
	return int64(C.Java_org_opencv_calib3d_StereoBM_create_11(clzEnv,clzObj))
}

func StereoBMNative_delete(nativeObj int64) {
	C.Java_org_opencv_calib3d_StereoBM_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func StereoBMNative_getPreFilterCap_0(nativeObj int64) int {
	return int(C.Java_org_opencv_calib3d_StereoBM_getPreFilterCap_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func StereoBMNative_getPreFilterSize_0(nativeObj int64) int {
	return int(C.Java_org_opencv_calib3d_StereoBM_getPreFilterSize_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func StereoBMNative_getPreFilterType_0(nativeObj int64) int {
	return int(C.Java_org_opencv_calib3d_StereoBM_getPreFilterType_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func StereoBMNative_getROI1_0(nativeObj int64) []float64 {
	return togofloat64Array(C.Java_org_opencv_calib3d_StereoBM_getROI1_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func StereoBMNative_getROI2_0(nativeObj int64) []float64 {
	return togofloat64Array(C.Java_org_opencv_calib3d_StereoBM_getROI2_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func StereoBMNative_getSmallerBlockSize_0(nativeObj int64) int {
	return int(C.Java_org_opencv_calib3d_StereoBM_getSmallerBlockSize_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func StereoBMNative_getTextureThreshold_0(nativeObj int64) int {
	return int(C.Java_org_opencv_calib3d_StereoBM_getTextureThreshold_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func StereoBMNative_getUniquenessRatio_0(nativeObj int64) int {
	return int(C.Java_org_opencv_calib3d_StereoBM_getUniquenessRatio_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func StereoBMNative_setPreFilterCap_0(nativeObj int64,preFilterCap int) {
	C.Java_org_opencv_calib3d_StereoBM_setPreFilterCap_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(preFilterCap))
}

func StereoBMNative_setPreFilterSize_0(nativeObj int64,preFilterSize int) {
	C.Java_org_opencv_calib3d_StereoBM_setPreFilterSize_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(preFilterSize))
}

func StereoBMNative_setPreFilterType_0(nativeObj int64,preFilterType int) {
	C.Java_org_opencv_calib3d_StereoBM_setPreFilterType_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(preFilterType))
}

func StereoBMNative_setROI1_0(nativeObj int64,roi1_x int,roi1_y int,roi1_width int,roi1_height int) {
	C.Java_org_opencv_calib3d_StereoBM_setROI1_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(roi1_x),(C.jint)(roi1_y),(C.jint)(roi1_width),(C.jint)(roi1_height))
}

func StereoBMNative_setROI2_0(nativeObj int64,roi2_x int,roi2_y int,roi2_width int,roi2_height int) {
	C.Java_org_opencv_calib3d_StereoBM_setROI2_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(roi2_x),(C.jint)(roi2_y),(C.jint)(roi2_width),(C.jint)(roi2_height))
}

func StereoBMNative_setSmallerBlockSize_0(nativeObj int64,blockSize int) {
	C.Java_org_opencv_calib3d_StereoBM_setSmallerBlockSize_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(blockSize))
}

func StereoBMNative_setTextureThreshold_0(nativeObj int64,textureThreshold int) {
	C.Java_org_opencv_calib3d_StereoBM_setTextureThreshold_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(textureThreshold))
}

func StereoBMNative_setUniquenessRatio_0(nativeObj int64,uniquenessRatio int) {
	C.Java_org_opencv_calib3d_StereoBM_setUniquenessRatio_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(uniquenessRatio))
}

func StereoMatcherNative_compute_0(nativeObj int64,left_nativeObj int64,right_nativeObj int64,disparity_nativeObj int64) {
	C.Java_org_opencv_calib3d_StereoMatcher_compute_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(left_nativeObj),(C.jlong)(right_nativeObj),(C.jlong)(disparity_nativeObj))
}

func StereoMatcherNative_delete(nativeObj int64) {
	C.Java_org_opencv_calib3d_StereoMatcher_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func StereoMatcherNative_getBlockSize_0(nativeObj int64) int {
	return int(C.Java_org_opencv_calib3d_StereoMatcher_getBlockSize_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func StereoMatcherNative_getDisp12MaxDiff_0(nativeObj int64) int {
	return int(C.Java_org_opencv_calib3d_StereoMatcher_getDisp12MaxDiff_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func StereoMatcherNative_getMinDisparity_0(nativeObj int64) int {
	return int(C.Java_org_opencv_calib3d_StereoMatcher_getMinDisparity_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func StereoMatcherNative_getNumDisparities_0(nativeObj int64) int {
	return int(C.Java_org_opencv_calib3d_StereoMatcher_getNumDisparities_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func StereoMatcherNative_getSpeckleRange_0(nativeObj int64) int {
	return int(C.Java_org_opencv_calib3d_StereoMatcher_getSpeckleRange_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func StereoMatcherNative_getSpeckleWindowSize_0(nativeObj int64) int {
	return int(C.Java_org_opencv_calib3d_StereoMatcher_getSpeckleWindowSize_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func StereoMatcherNative_setBlockSize_0(nativeObj int64,blockSize int) {
	C.Java_org_opencv_calib3d_StereoMatcher_setBlockSize_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(blockSize))
}

func StereoMatcherNative_setDisp12MaxDiff_0(nativeObj int64,disp12MaxDiff int) {
	C.Java_org_opencv_calib3d_StereoMatcher_setDisp12MaxDiff_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(disp12MaxDiff))
}

func StereoMatcherNative_setMinDisparity_0(nativeObj int64,minDisparity int) {
	C.Java_org_opencv_calib3d_StereoMatcher_setMinDisparity_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(minDisparity))
}

func StereoMatcherNative_setNumDisparities_0(nativeObj int64,numDisparities int) {
	C.Java_org_opencv_calib3d_StereoMatcher_setNumDisparities_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(numDisparities))
}

func StereoMatcherNative_setSpeckleRange_0(nativeObj int64,speckleRange int) {
	C.Java_org_opencv_calib3d_StereoMatcher_setSpeckleRange_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(speckleRange))
}

func StereoMatcherNative_setSpeckleWindowSize_0(nativeObj int64,speckleWindowSize int) {
	C.Java_org_opencv_calib3d_StereoMatcher_setSpeckleWindowSize_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(speckleWindowSize))
}

func StereoSGBMNative_create_0(minDisparity int,numDisparities int,blockSize int,P1 int,P2 int,disp12MaxDiff int,preFilterCap int,uniquenessRatio int,speckleWindowSize int,speckleRange int,mode int) int64 {
	return int64(C.Java_org_opencv_calib3d_StereoSGBM_create_10(clzEnv,clzObj,(C.jint)(minDisparity),(C.jint)(numDisparities),(C.jint)(blockSize),(C.jint)(P1),(C.jint)(P2),(C.jint)(disp12MaxDiff),(C.jint)(preFilterCap),(C.jint)(uniquenessRatio),(C.jint)(speckleWindowSize),(C.jint)(speckleRange),(C.jint)(mode)))
}

func StereoSGBMNative_create_1() int64 {
	return int64(C.Java_org_opencv_calib3d_StereoSGBM_create_11(clzEnv,clzObj))
}

func StereoSGBMNative_delete(nativeObj int64) {
	C.Java_org_opencv_calib3d_StereoSGBM_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func StereoSGBMNative_getMode_0(nativeObj int64) int {
	return int(C.Java_org_opencv_calib3d_StereoSGBM_getMode_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func StereoSGBMNative_getP1_0(nativeObj int64) int {
	return int(C.Java_org_opencv_calib3d_StereoSGBM_getP1_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func StereoSGBMNative_getP2_0(nativeObj int64) int {
	return int(C.Java_org_opencv_calib3d_StereoSGBM_getP2_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func StereoSGBMNative_getPreFilterCap_0(nativeObj int64) int {
	return int(C.Java_org_opencv_calib3d_StereoSGBM_getPreFilterCap_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func StereoSGBMNative_getUniquenessRatio_0(nativeObj int64) int {
	return int(C.Java_org_opencv_calib3d_StereoSGBM_getUniquenessRatio_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func StereoSGBMNative_setMode_0(nativeObj int64,mode int) {
	C.Java_org_opencv_calib3d_StereoSGBM_setMode_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(mode))
}

func StereoSGBMNative_setP1_0(nativeObj int64,P1 int) {
	C.Java_org_opencv_calib3d_StereoSGBM_setP1_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(P1))
}

func StereoSGBMNative_setP2_0(nativeObj int64,P2 int) {
	C.Java_org_opencv_calib3d_StereoSGBM_setP2_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(P2))
}

func StereoSGBMNative_setPreFilterCap_0(nativeObj int64,preFilterCap int) {
	C.Java_org_opencv_calib3d_StereoSGBM_setPreFilterCap_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(preFilterCap))
}

func StereoSGBMNative_setUniquenessRatio_0(nativeObj int64,uniquenessRatio int) {
	C.Java_org_opencv_calib3d_StereoSGBM_setUniquenessRatio_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(uniquenessRatio))
}
