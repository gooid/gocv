package core

/*

#include "jni.h"

extern jlong Java_org_opencv_features2d_AKAZE_create_10(JNIEnv*, jclass, jint, jint, jint, jfloat, jint, jint, jint);
extern jlong Java_org_opencv_features2d_AKAZE_create_11(JNIEnv*, jclass);
extern void Java_org_opencv_features2d_AKAZE_delete(JNIEnv*, jclass, jlong);
extern jstring Java_org_opencv_features2d_AKAZE_getDefaultName_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_features2d_AKAZE_getDescriptorChannels_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_features2d_AKAZE_getDescriptorSize_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_features2d_AKAZE_getDescriptorType_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_features2d_AKAZE_getDiffusivity_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_features2d_AKAZE_getNOctaveLayers_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_features2d_AKAZE_getNOctaves_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_features2d_AKAZE_getThreshold_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_features2d_AKAZE_setDescriptorChannels_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_features2d_AKAZE_setDescriptorSize_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_features2d_AKAZE_setDescriptorType_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_features2d_AKAZE_setDiffusivity_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_features2d_AKAZE_setNOctaveLayers_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_features2d_AKAZE_setNOctaves_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_features2d_AKAZE_setThreshold_10(JNIEnv*, jclass, jlong, jdouble);
extern jlong Java_org_opencv_features2d_AgastFeatureDetector_create_10(JNIEnv*, jclass, jint, jboolean, jint);
extern jlong Java_org_opencv_features2d_AgastFeatureDetector_create_11(JNIEnv*, jclass);
extern void Java_org_opencv_features2d_AgastFeatureDetector_delete(JNIEnv*, jclass, jlong);
extern jstring Java_org_opencv_features2d_AgastFeatureDetector_getDefaultName_10(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_features2d_AgastFeatureDetector_getNonmaxSuppression_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_features2d_AgastFeatureDetector_getThreshold_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_features2d_AgastFeatureDetector_getType_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_features2d_AgastFeatureDetector_setNonmaxSuppression_10(JNIEnv*, jclass, jlong, jboolean);
extern void Java_org_opencv_features2d_AgastFeatureDetector_setThreshold_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_features2d_AgastFeatureDetector_setType_10(JNIEnv*, jclass, jlong, jint);
extern jlong Java_org_opencv_features2d_BFMatcher_BFMatcher_10(JNIEnv*, jclass, jint, jboolean);
extern jlong Java_org_opencv_features2d_BFMatcher_BFMatcher_11(JNIEnv*, jclass);
extern jlong Java_org_opencv_features2d_BFMatcher_create_10(JNIEnv*, jclass, jint, jboolean);
extern jlong Java_org_opencv_features2d_BFMatcher_create_11(JNIEnv*, jclass);
extern void Java_org_opencv_features2d_BFMatcher_delete(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_features2d_BOWImgDescriptorExtractor_compute_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_features2d_BOWImgDescriptorExtractor_delete(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_features2d_BOWImgDescriptorExtractor_descriptorSize_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_features2d_BOWImgDescriptorExtractor_descriptorType_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_features2d_BOWImgDescriptorExtractor_getVocabulary_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_features2d_BOWImgDescriptorExtractor_setVocabulary_10(JNIEnv*, jclass, jlong, jlong);
extern jlong Java_org_opencv_features2d_BOWKMeansTrainer_BOWKMeansTrainer_10(JNIEnv*, jclass, jint, jint, jint, jdouble, jint, jint);
extern jlong Java_org_opencv_features2d_BOWKMeansTrainer_BOWKMeansTrainer_11(JNIEnv*, jclass, jint);
extern jlong Java_org_opencv_features2d_BOWKMeansTrainer_cluster_10(JNIEnv*, jclass, jlong, jlong);
extern jlong Java_org_opencv_features2d_BOWKMeansTrainer_cluster_11(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_features2d_BOWKMeansTrainer_delete(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_features2d_BOWTrainer_add_10(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_features2d_BOWTrainer_clear_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_features2d_BOWTrainer_cluster_10(JNIEnv*, jclass, jlong, jlong);
extern jlong Java_org_opencv_features2d_BOWTrainer_cluster_11(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_features2d_BOWTrainer_delete(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_features2d_BOWTrainer_descriptorsCount_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_features2d_BOWTrainer_getDescriptors_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_features2d_BRISK_create_10(JNIEnv*, jclass, jint, jint, jlong, jlong, jfloat, jfloat, jlong);
extern jlong Java_org_opencv_features2d_BRISK_create_11(JNIEnv*, jclass, jint, jint, jlong, jlong);
extern jlong Java_org_opencv_features2d_BRISK_create_12(JNIEnv*, jclass, jint, jint, jfloat);
extern jlong Java_org_opencv_features2d_BRISK_create_13(JNIEnv*, jclass);
extern jlong Java_org_opencv_features2d_BRISK_create_14(JNIEnv*, jclass, jlong, jlong, jfloat, jfloat, jlong);
extern jlong Java_org_opencv_features2d_BRISK_create_15(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_features2d_BRISK_delete(JNIEnv*, jclass, jlong);
extern jstring Java_org_opencv_features2d_BRISK_getDefaultName_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_features2d_DescriptorExtractor_compute_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_features2d_DescriptorExtractor_compute_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern jlong Java_org_opencv_features2d_DescriptorExtractor_create_10(JNIEnv*, jclass, jint);
extern void Java_org_opencv_features2d_DescriptorExtractor_delete(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_features2d_DescriptorExtractor_descriptorSize_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_features2d_DescriptorExtractor_descriptorType_10(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_features2d_DescriptorExtractor_empty_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_features2d_DescriptorExtractor_read_10(JNIEnv*, jclass, jlong, jstring);
extern void Java_org_opencv_features2d_DescriptorExtractor_write_10(JNIEnv*, jclass, jlong, jstring);
extern void Java_org_opencv_features2d_DescriptorMatcher_add_10(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_features2d_DescriptorMatcher_clear_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_features2d_DescriptorMatcher_clone_10(JNIEnv*, jclass, jlong, jboolean);
extern jlong Java_org_opencv_features2d_DescriptorMatcher_clone_11(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_features2d_DescriptorMatcher_create_10(JNIEnv*, jclass, jstring);
extern jlong Java_org_opencv_features2d_DescriptorMatcher_create_11(JNIEnv*, jclass, jint);
extern void Java_org_opencv_features2d_DescriptorMatcher_delete(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_features2d_DescriptorMatcher_empty_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_features2d_DescriptorMatcher_getTrainDescriptors_10(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_features2d_DescriptorMatcher_isMaskSupported_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_features2d_DescriptorMatcher_knnMatch_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jint, jlong, jboolean);
extern void Java_org_opencv_features2d_DescriptorMatcher_knnMatch_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jint);
extern void Java_org_opencv_features2d_DescriptorMatcher_knnMatch_12(JNIEnv*, jclass, jlong, jlong, jlong, jint, jlong, jboolean);
extern void Java_org_opencv_features2d_DescriptorMatcher_knnMatch_13(JNIEnv*, jclass, jlong, jlong, jlong, jint);
extern void Java_org_opencv_features2d_DescriptorMatcher_match_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_features2d_DescriptorMatcher_match_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_features2d_DescriptorMatcher_match_12(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_features2d_DescriptorMatcher_match_13(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_features2d_DescriptorMatcher_radiusMatch_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jfloat, jlong, jboolean);
extern void Java_org_opencv_features2d_DescriptorMatcher_radiusMatch_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jfloat);
extern void Java_org_opencv_features2d_DescriptorMatcher_radiusMatch_12(JNIEnv*, jclass, jlong, jlong, jlong, jfloat, jlong, jboolean);
extern void Java_org_opencv_features2d_DescriptorMatcher_radiusMatch_13(JNIEnv*, jclass, jlong, jlong, jlong, jfloat);
extern void Java_org_opencv_features2d_DescriptorMatcher_read_10(JNIEnv*, jclass, jlong, jstring);
extern void Java_org_opencv_features2d_DescriptorMatcher_train_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_features2d_DescriptorMatcher_write_10(JNIEnv*, jclass, jlong, jstring);
extern jlong Java_org_opencv_features2d_FastFeatureDetector_create_10(JNIEnv*, jclass, jint, jboolean, jint);
extern jlong Java_org_opencv_features2d_FastFeatureDetector_create_11(JNIEnv*, jclass);
extern void Java_org_opencv_features2d_FastFeatureDetector_delete(JNIEnv*, jclass, jlong);
extern jstring Java_org_opencv_features2d_FastFeatureDetector_getDefaultName_10(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_features2d_FastFeatureDetector_getNonmaxSuppression_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_features2d_FastFeatureDetector_getThreshold_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_features2d_FastFeatureDetector_getType_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_features2d_FastFeatureDetector_setNonmaxSuppression_10(JNIEnv*, jclass, jlong, jboolean);
extern void Java_org_opencv_features2d_FastFeatureDetector_setThreshold_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_features2d_FastFeatureDetector_setType_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_features2d_Feature2D_compute_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_features2d_Feature2D_compute_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern jint Java_org_opencv_features2d_Feature2D_defaultNorm_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_features2d_Feature2D_delete(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_features2d_Feature2D_descriptorSize_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_features2d_Feature2D_descriptorType_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_features2d_Feature2D_detectAndCompute_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jboolean);
extern void Java_org_opencv_features2d_Feature2D_detectAndCompute_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_features2d_Feature2D_detect_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_features2d_Feature2D_detect_11(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_features2d_Feature2D_detect_12(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_features2d_Feature2D_detect_13(JNIEnv*, jclass, jlong, jlong, jlong);
extern jboolean Java_org_opencv_features2d_Feature2D_empty_10(JNIEnv*, jclass, jlong);
extern jstring Java_org_opencv_features2d_Feature2D_getDefaultName_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_features2d_Feature2D_read_10(JNIEnv*, jclass, jlong, jstring);
extern void Java_org_opencv_features2d_Feature2D_write_10(JNIEnv*, jclass, jlong, jstring);
extern jlong Java_org_opencv_features2d_FeatureDetector_create_10(JNIEnv*, jclass, jint);
extern void Java_org_opencv_features2d_FeatureDetector_delete(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_features2d_FeatureDetector_detect_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_features2d_FeatureDetector_detect_11(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_features2d_FeatureDetector_detect_12(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_features2d_FeatureDetector_detect_13(JNIEnv*, jclass, jlong, jlong, jlong);
extern jboolean Java_org_opencv_features2d_FeatureDetector_empty_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_features2d_FeatureDetector_read_10(JNIEnv*, jclass, jlong, jstring);
extern void Java_org_opencv_features2d_FeatureDetector_write_10(JNIEnv*, jclass, jlong, jstring);
extern void Java_org_opencv_features2d_Features2d_drawKeypoints_10(JNIEnv*, jclass, jlong, jlong, jlong, jdouble, jdouble, jdouble, jdouble, jint);
extern void Java_org_opencv_features2d_Features2d_drawKeypoints_11(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_features2d_Features2d_drawMatches2_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jlong, jint);
extern void Java_org_opencv_features2d_Features2d_drawMatches2_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_features2d_Features2d_drawMatchesKnn_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jlong, jint);
extern void Java_org_opencv_features2d_Features2d_drawMatchesKnn_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_features2d_Features2d_drawMatches_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jlong, jint);
extern void Java_org_opencv_features2d_Features2d_drawMatches_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong);
extern jlong Java_org_opencv_features2d_FlannBasedMatcher_FlannBasedMatcher_10(JNIEnv*, jclass);
extern jlong Java_org_opencv_features2d_FlannBasedMatcher_create_10(JNIEnv*, jclass);
extern void Java_org_opencv_features2d_FlannBasedMatcher_delete(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_features2d_GFTTDetector_create_10(JNIEnv*, jclass, jint, jdouble, jdouble, jint, jint, jboolean, jdouble);
extern jlong Java_org_opencv_features2d_GFTTDetector_create_11(JNIEnv*, jclass, jint, jdouble, jdouble, jint, jint);
extern jlong Java_org_opencv_features2d_GFTTDetector_create_12(JNIEnv*, jclass, jint, jdouble, jdouble, jint, jboolean, jdouble);
extern jlong Java_org_opencv_features2d_GFTTDetector_create_13(JNIEnv*, jclass);
extern void Java_org_opencv_features2d_GFTTDetector_delete(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_features2d_GFTTDetector_getBlockSize_10(JNIEnv*, jclass, jlong);
extern jstring Java_org_opencv_features2d_GFTTDetector_getDefaultName_10(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_features2d_GFTTDetector_getHarrisDetector_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_features2d_GFTTDetector_getK_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_features2d_GFTTDetector_getMaxFeatures_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_features2d_GFTTDetector_getMinDistance_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_features2d_GFTTDetector_getQualityLevel_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_features2d_GFTTDetector_setBlockSize_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_features2d_GFTTDetector_setHarrisDetector_10(JNIEnv*, jclass, jlong, jboolean);
extern void Java_org_opencv_features2d_GFTTDetector_setK_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_features2d_GFTTDetector_setMaxFeatures_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_features2d_GFTTDetector_setMinDistance_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_features2d_GFTTDetector_setQualityLevel_10(JNIEnv*, jclass, jlong, jdouble);
extern jlong Java_org_opencv_features2d_KAZE_create_10(JNIEnv*, jclass, jboolean, jboolean, jfloat, jint, jint, jint);
extern jlong Java_org_opencv_features2d_KAZE_create_11(JNIEnv*, jclass);
extern void Java_org_opencv_features2d_KAZE_delete(JNIEnv*, jclass, jlong);
extern jstring Java_org_opencv_features2d_KAZE_getDefaultName_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_features2d_KAZE_getDiffusivity_10(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_features2d_KAZE_getExtended_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_features2d_KAZE_getNOctaveLayers_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_features2d_KAZE_getNOctaves_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_features2d_KAZE_getThreshold_10(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_features2d_KAZE_getUpright_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_features2d_KAZE_setDiffusivity_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_features2d_KAZE_setExtended_10(JNIEnv*, jclass, jlong, jboolean);
extern void Java_org_opencv_features2d_KAZE_setNOctaveLayers_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_features2d_KAZE_setNOctaves_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_features2d_KAZE_setThreshold_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_features2d_KAZE_setUpright_10(JNIEnv*, jclass, jlong, jboolean);
extern jlong Java_org_opencv_features2d_MSER_create_10(JNIEnv*, jclass, jint, jint, jint, jdouble, jdouble, jint, jdouble, jdouble, jint);
extern jlong Java_org_opencv_features2d_MSER_create_11(JNIEnv*, jclass);
extern void Java_org_opencv_features2d_MSER_delete(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_features2d_MSER_detectRegions_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern jstring Java_org_opencv_features2d_MSER_getDefaultName_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_features2d_MSER_getDelta_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_features2d_MSER_getMaxArea_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_features2d_MSER_getMinArea_10(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_features2d_MSER_getPass2Only_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_features2d_MSER_setDelta_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_features2d_MSER_setMaxArea_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_features2d_MSER_setMinArea_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_features2d_MSER_setPass2Only_10(JNIEnv*, jclass, jlong, jboolean);
extern jlong Java_org_opencv_features2d_ORB_create_10(JNIEnv*, jclass, jint, jfloat, jint, jint, jint, jint, jint, jint, jint);
extern jlong Java_org_opencv_features2d_ORB_create_11(JNIEnv*, jclass);
extern void Java_org_opencv_features2d_ORB_delete(JNIEnv*, jclass, jlong);
extern jstring Java_org_opencv_features2d_ORB_getDefaultName_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_features2d_ORB_getEdgeThreshold_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_features2d_ORB_getFastThreshold_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_features2d_ORB_getFirstLevel_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_features2d_ORB_getMaxFeatures_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_features2d_ORB_getNLevels_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_features2d_ORB_getPatchSize_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_features2d_ORB_getScaleFactor_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_features2d_ORB_getScoreType_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_features2d_ORB_getWTA_1K_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_features2d_ORB_setEdgeThreshold_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_features2d_ORB_setFastThreshold_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_features2d_ORB_setFirstLevel_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_features2d_ORB_setMaxFeatures_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_features2d_ORB_setNLevels_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_features2d_ORB_setPatchSize_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_features2d_ORB_setScaleFactor_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_features2d_ORB_setScoreType_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_features2d_ORB_setWTA_1K_10(JNIEnv*, jclass, jlong, jint);
extern jlong Java_org_opencv_features2d_Params_Params_10(JNIEnv*, jclass);
extern void Java_org_opencv_features2d_Params_delete(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_features2d_Params_get_1filterByArea_10(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_features2d_Params_get_1filterByCircularity_10(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_features2d_Params_get_1filterByColor_10(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_features2d_Params_get_1filterByConvexity_10(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_features2d_Params_get_1filterByInertia_10(JNIEnv*, jclass, jlong);
extern jfloat Java_org_opencv_features2d_Params_get_1maxArea_10(JNIEnv*, jclass, jlong);
extern jfloat Java_org_opencv_features2d_Params_get_1maxCircularity_10(JNIEnv*, jclass, jlong);
extern jfloat Java_org_opencv_features2d_Params_get_1maxConvexity_10(JNIEnv*, jclass, jlong);
extern jfloat Java_org_opencv_features2d_Params_get_1maxInertiaRatio_10(JNIEnv*, jclass, jlong);
extern jfloat Java_org_opencv_features2d_Params_get_1maxThreshold_10(JNIEnv*, jclass, jlong);
extern jfloat Java_org_opencv_features2d_Params_get_1minArea_10(JNIEnv*, jclass, jlong);
extern jfloat Java_org_opencv_features2d_Params_get_1minCircularity_10(JNIEnv*, jclass, jlong);
extern jfloat Java_org_opencv_features2d_Params_get_1minConvexity_10(JNIEnv*, jclass, jlong);
extern jfloat Java_org_opencv_features2d_Params_get_1minDistBetweenBlobs_10(JNIEnv*, jclass, jlong);
extern jfloat Java_org_opencv_features2d_Params_get_1minInertiaRatio_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_features2d_Params_get_1minRepeatability_10(JNIEnv*, jclass, jlong);
extern jfloat Java_org_opencv_features2d_Params_get_1minThreshold_10(JNIEnv*, jclass, jlong);
extern jfloat Java_org_opencv_features2d_Params_get_1thresholdStep_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_features2d_Params_set_1filterByArea_10(JNIEnv*, jclass, jlong, jboolean);
extern void Java_org_opencv_features2d_Params_set_1filterByCircularity_10(JNIEnv*, jclass, jlong, jboolean);
extern void Java_org_opencv_features2d_Params_set_1filterByColor_10(JNIEnv*, jclass, jlong, jboolean);
extern void Java_org_opencv_features2d_Params_set_1filterByConvexity_10(JNIEnv*, jclass, jlong, jboolean);
extern void Java_org_opencv_features2d_Params_set_1filterByInertia_10(JNIEnv*, jclass, jlong, jboolean);
extern void Java_org_opencv_features2d_Params_set_1maxArea_10(JNIEnv*, jclass, jlong, jfloat);
extern void Java_org_opencv_features2d_Params_set_1maxCircularity_10(JNIEnv*, jclass, jlong, jfloat);
extern void Java_org_opencv_features2d_Params_set_1maxConvexity_10(JNIEnv*, jclass, jlong, jfloat);
extern void Java_org_opencv_features2d_Params_set_1maxInertiaRatio_10(JNIEnv*, jclass, jlong, jfloat);
extern void Java_org_opencv_features2d_Params_set_1maxThreshold_10(JNIEnv*, jclass, jlong, jfloat);
extern void Java_org_opencv_features2d_Params_set_1minArea_10(JNIEnv*, jclass, jlong, jfloat);
extern void Java_org_opencv_features2d_Params_set_1minCircularity_10(JNIEnv*, jclass, jlong, jfloat);
extern void Java_org_opencv_features2d_Params_set_1minConvexity_10(JNIEnv*, jclass, jlong, jfloat);
extern void Java_org_opencv_features2d_Params_set_1minDistBetweenBlobs_10(JNIEnv*, jclass, jlong, jfloat);
extern void Java_org_opencv_features2d_Params_set_1minInertiaRatio_10(JNIEnv*, jclass, jlong, jfloat);
extern void Java_org_opencv_features2d_Params_set_1minRepeatability_10(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_features2d_Params_set_1minThreshold_10(JNIEnv*, jclass, jlong, jfloat);
extern void Java_org_opencv_features2d_Params_set_1thresholdStep_10(JNIEnv*, jclass, jlong, jfloat);

*/
import "C"

func AKAZENative_create_0(descriptor_type int,descriptor_size int,descriptor_channels int,threshold float32,nOctaves int,nOctaveLayers int,diffusivity int) int64 {
	return int64(C.Java_org_opencv_features2d_AKAZE_create_10(clzEnv,clzObj,(C.jint)(descriptor_type),(C.jint)(descriptor_size),(C.jint)(descriptor_channels),(C.jfloat)(threshold),(C.jint)(
			nOctaves),(C.jint)(nOctaveLayers),(C.jint)(diffusivity)))
}
func AKAZENative_create_1() int64 {
	return int64(C.Java_org_opencv_features2d_AKAZE_create_11(clzEnv,clzObj))
}

func AKAZENative_delete(nativeObj int64) {
	C.Java_org_opencv_features2d_AKAZE_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func AKAZENative_getDefaultName_0(nativeObj int64) string {
	return togostring(
			C.Java_org_opencv_features2d_AKAZE_getDefaultName_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func AKAZENative_getDescriptorChannels_0(nativeObj int64) int {
	return int(C.Java_org_opencv_features2d_AKAZE_getDescriptorChannels_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func AKAZENative_getDescriptorSize_0(nativeObj int64) int {
	return int(C.Java_org_opencv_features2d_AKAZE_getDescriptorSize_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func AKAZENative_getDescriptorType_0(nativeObj int64) int {
	return int(C.Java_org_opencv_features2d_AKAZE_getDescriptorType_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func AKAZENative_getDiffusivity_0(nativeObj int64) int {
	return int(C.Java_org_opencv_features2d_AKAZE_getDiffusivity_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func AKAZENative_getNOctaveLayers_0(nativeObj int64) int {
	return int(C.Java_org_opencv_features2d_AKAZE_getNOctaveLayers_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func AKAZENative_getNOctaves_0(nativeObj int64) int {
	return int(C.Java_org_opencv_features2d_AKAZE_getNOctaves_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func AKAZENative_getThreshold_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_features2d_AKAZE_getThreshold_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func AKAZENative_setDescriptorChannels_0(nativeObj int64,dch int) {
	C.Java_org_opencv_features2d_AKAZE_setDescriptorChannels_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(dch))
}

func AKAZENative_setDescriptorSize_0(nativeObj int64,dsize int) {
	C.Java_org_opencv_features2d_AKAZE_setDescriptorSize_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(dsize))
}

func AKAZENative_setDescriptorType_0(nativeObj int64,dtype int) {
	C.Java_org_opencv_features2d_AKAZE_setDescriptorType_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(dtype))
}

func AKAZENative_setDiffusivity_0(nativeObj int64,diff int) {
	C.Java_org_opencv_features2d_AKAZE_setDiffusivity_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(diff))
}

func AKAZENative_setNOctaveLayers_0(nativeObj int64,octaveLayers int) {
	C.Java_org_opencv_features2d_AKAZE_setNOctaveLayers_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(octaveLayers))
}

func AKAZENative_setNOctaves_0(nativeObj int64,octaves int) {
	C.Java_org_opencv_features2d_AKAZE_setNOctaves_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(octaves))
}

func AKAZENative_setThreshold_0(nativeObj int64,threshold float64) {
	C.Java_org_opencv_features2d_AKAZE_setThreshold_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jdouble)(threshold))
}

func AgastFeatureDetectorNative_create_0(threshold int,nonmaxSuppression bool,rtype int) int64 {
	return int64(C.Java_org_opencv_features2d_AgastFeatureDetector_create_10(clzEnv,clzObj,(C.jint)(threshold),tojboolean(nonmaxSuppression),(C.jint)(rtype)))
}

func AgastFeatureDetectorNative_create_1() int64 {
	return int64(C.Java_org_opencv_features2d_AgastFeatureDetector_create_11(clzEnv,clzObj))
}

func AgastFeatureDetectorNative_delete(nativeObj int64) {
	C.Java_org_opencv_features2d_AgastFeatureDetector_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func AgastFeatureDetectorNative_getDefaultName_0(nativeObj int64) string {
	return togostring(
			C.Java_org_opencv_features2d_AgastFeatureDetector_getDefaultName_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func AgastFeatureDetectorNative_getNonmaxSuppression_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_features2d_AgastFeatureDetector_getNonmaxSuppression_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func AgastFeatureDetectorNative_getThreshold_0(nativeObj int64) int {
	return int(C.Java_org_opencv_features2d_AgastFeatureDetector_getThreshold_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func AgastFeatureDetectorNative_getType_0(nativeObj int64) int {
	return int(C.Java_org_opencv_features2d_AgastFeatureDetector_getType_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func AgastFeatureDetectorNative_setNonmaxSuppression_0(nativeObj int64,f bool) {
	C.Java_org_opencv_features2d_AgastFeatureDetector_setNonmaxSuppression_10(clzEnv,clzObj,(C.jlong)(nativeObj),tojboolean(f))
}

func AgastFeatureDetectorNative_setThreshold_0(nativeObj int64,threshold int) {
	C.Java_org_opencv_features2d_AgastFeatureDetector_setThreshold_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(threshold))
}

func AgastFeatureDetectorNative_setType_0(nativeObj int64,rtype int) {
	C.Java_org_opencv_features2d_AgastFeatureDetector_setType_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(rtype))
}

func BFMatcherNative_BFMatcher_0(normType int,crossCheck bool) int64 {
	return int64(C.Java_org_opencv_features2d_BFMatcher_BFMatcher_10(clzEnv,clzObj,(C.jint)(normType),tojboolean(crossCheck)))
}

func BFMatcherNative_BFMatcher_1() int64 {
	return int64(C.Java_org_opencv_features2d_BFMatcher_BFMatcher_11(clzEnv,clzObj))
}

func BFMatcherNative_create_0(normType int,crossCheck bool) int64 {
	return int64(C.Java_org_opencv_features2d_BFMatcher_create_10(clzEnv,clzObj,(C.jint)(normType),tojboolean(crossCheck)))
}

func BFMatcherNative_create_1() int64 {
	return int64(C.Java_org_opencv_features2d_BFMatcher_create_11(clzEnv,clzObj))
}

func BFMatcherNative_delete(nativeObj int64) {
	C.Java_org_opencv_features2d_BFMatcher_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func BOWImgDescriptorExtractorNative_compute_0(nativeObj int64,image_nativeObj int64,keypoints_mat_nativeObj int64,imgDescriptor_nativeObj int64) {
	C.Java_org_opencv_features2d_BOWImgDescriptorExtractor_compute_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(image_nativeObj),(C.jlong)(keypoints_mat_nativeObj),(C.jlong)(imgDescriptor_nativeObj))
}

func BOWImgDescriptorExtractorNative_delete(nativeObj int64) {
	C.Java_org_opencv_features2d_BOWImgDescriptorExtractor_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func BOWImgDescriptorExtractorNative_descriptorSize_0(nativeObj int64) int {
	return int(C.Java_org_opencv_features2d_BOWImgDescriptorExtractor_descriptorSize_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func BOWImgDescriptorExtractorNative_descriptorType_0(nativeObj int64) int {
	return int(C.Java_org_opencv_features2d_BOWImgDescriptorExtractor_descriptorType_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func BOWImgDescriptorExtractorNative_getVocabulary_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_features2d_BOWImgDescriptorExtractor_getVocabulary_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func BOWImgDescriptorExtractorNative_setVocabulary_0(nativeObj int64,vocabulary_nativeObj int64) {
	C.Java_org_opencv_features2d_BOWImgDescriptorExtractor_setVocabulary_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(vocabulary_nativeObj))
}

func BOWKMeansTrainerNative_BOWKMeansTrainer_0(clusterCount int,termcrit_type int,termcrit_maxCount int,termcrit_epsilon float64,attempts int,flags int) int64 {
	return int64(C.Java_org_opencv_features2d_BOWKMeansTrainer_BOWKMeansTrainer_10(clzEnv,clzObj,(C.jint)(clusterCount),(C.jint)(termcrit_type),(C.jint)(termcrit_maxCount),(C.jdouble)(termcrit_epsilon),(C.jint)(attempts),(C.jint)(flags)))
}

func BOWKMeansTrainerNative_BOWKMeansTrainer_1(clusterCount int) int64 {
	return int64(C.Java_org_opencv_features2d_BOWKMeansTrainer_BOWKMeansTrainer_11(clzEnv,clzObj,(C.jint)(clusterCount)))
}

func BOWKMeansTrainerNative_cluster_0(nativeObj int64,descriptors_nativeObj int64) int64 {
	return int64(C.Java_org_opencv_features2d_BOWKMeansTrainer_cluster_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(descriptors_nativeObj)))
}

func BOWKMeansTrainerNative_cluster_1(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_features2d_BOWKMeansTrainer_cluster_11(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func BOWKMeansTrainerNative_delete(nativeObj int64) {
	C.Java_org_opencv_features2d_BOWKMeansTrainer_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func BOWTrainerNative_add_0(nativeObj int64,descriptors_nativeObj int64) {
	C.Java_org_opencv_features2d_BOWTrainer_add_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(descriptors_nativeObj))
}

func BOWTrainerNative_clear_0(nativeObj int64) {
	C.Java_org_opencv_features2d_BOWTrainer_clear_10(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func BOWTrainerNative_cluster_0(nativeObj int64,descriptors_nativeObj int64) int64 {
	return int64(C.Java_org_opencv_features2d_BOWTrainer_cluster_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(descriptors_nativeObj)))
}

func BOWTrainerNative_cluster_1(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_features2d_BOWTrainer_cluster_11(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func BOWTrainerNative_delete(nativeObj int64) {
	C.Java_org_opencv_features2d_BOWTrainer_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func BOWTrainerNative_descriptorsCount_0(nativeObj int64) int {
	return int(C.Java_org_opencv_features2d_BOWTrainer_descriptorsCount_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func BOWTrainerNative_getDescriptors_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_features2d_BOWTrainer_getDescriptors_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func BRISKNative_create_0(thresh int,octaves int,radiusList_mat_nativeObj int64,numberList_mat_nativeObj int64,dMax float32,dMin float32,indexChange_mat_nativeObj int64) int64 {
	return int64(C.Java_org_opencv_features2d_BRISK_create_10(clzEnv,clzObj,(C.jint)(thresh),(C.jint)(octaves),(C.jlong)(radiusList_mat_nativeObj),(C.jlong)(numberList_mat_nativeObj),(C.jfloat)(dMax),(C.jfloat)(dMin),(C.jlong)(indexChange_mat_nativeObj)))
}

func BRISKNative_create_1(thresh int,octaves int,radiusList_mat_nativeObj int64,numberList_mat_nativeObj int64) int64 {
	return int64(C.Java_org_opencv_features2d_BRISK_create_11(clzEnv,clzObj,(C.jint)(thresh),(C.jint)(octaves),(C.jlong)(radiusList_mat_nativeObj),(C.jlong)(numberList_mat_nativeObj)))
}

func BRISKNative_create_2(thresh int,octaves int,patternScale float32) int64 {
	return int64(C.Java_org_opencv_features2d_BRISK_create_12(clzEnv,clzObj,(C.jint)(thresh),(C.jint)(octaves),(C.jfloat)(patternScale)))
}

func BRISKNative_create_3() int64 {
	return int64(C.Java_org_opencv_features2d_BRISK_create_13(clzEnv,clzObj))
}

func BRISKNative_create_4(radiusList_mat_nativeObj int64,numberList_mat_nativeObj int64,dMax float32,dMin float32,indexChange_mat_nativeObj int64) int64 {
	return int64(C.Java_org_opencv_features2d_BRISK_create_14(clzEnv,clzObj,(C.jlong)(radiusList_mat_nativeObj),(C.jlong)(numberList_mat_nativeObj),(C.jfloat)(dMax),(C.jfloat)(dMin),(C.jlong)(indexChange_mat_nativeObj)))
}

func BRISKNative_create_5(radiusList_mat_nativeObj int64,numberList_mat_nativeObj int64) int64 {
	return int64(C.Java_org_opencv_features2d_BRISK_create_15(clzEnv,clzObj,(C.jlong)(radiusList_mat_nativeObj),(C.jlong)(numberList_mat_nativeObj)))
}

func BRISKNative_delete(nativeObj int64) {
	C.Java_org_opencv_features2d_BRISK_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func BRISKNative_getDefaultName_0(nativeObj int64) string {
	return togostring(
			C.Java_org_opencv_features2d_BRISK_getDefaultName_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func DescriptorExtractorNative_compute_0(nativeObj int64,image_nativeObj int64,keypoints_mat_nativeObj int64,descriptors_nativeObj int64) {
	C.Java_org_opencv_features2d_DescriptorExtractor_compute_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(image_nativeObj),(C.jlong)(keypoints_mat_nativeObj),(C.jlong)(descriptors_nativeObj))
}

func DescriptorExtractorNative_compute_1(nativeObj int64,images_mat_nativeObj int64,keypoints_mat_nativeObj int64,descriptors_mat_nativeObj int64) {
	C.Java_org_opencv_features2d_DescriptorExtractor_compute_11(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(images_mat_nativeObj),(C.jlong)(keypoints_mat_nativeObj),(C.jlong)(descriptors_mat_nativeObj))
}

func DescriptorExtractorNative_create_0(extractorType int) int64 {
	return int64(C.Java_org_opencv_features2d_DescriptorExtractor_create_10(clzEnv,clzObj,(C.jint)(extractorType)))
}

func DescriptorExtractorNative_delete(nativeObj int64) {
	C.Java_org_opencv_features2d_DescriptorExtractor_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func DescriptorExtractorNative_descriptorSize_0(nativeObj int64) int {
	return int(C.Java_org_opencv_features2d_DescriptorExtractor_descriptorSize_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func DescriptorExtractorNative_descriptorType_0(nativeObj int64) int {
	return int(C.Java_org_opencv_features2d_DescriptorExtractor_descriptorType_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func DescriptorExtractorNative_empty_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_features2d_DescriptorExtractor_empty_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func DescriptorExtractorNative_read_0(nativeObj int64,fileName string) {
	C.Java_org_opencv_features2d_DescriptorExtractor_read_10(clzEnv,clzObj,(C.jlong)(nativeObj),tojstring(fileName))
}

func DescriptorExtractorNative_write_0(nativeObj int64,fileName string) {
	C.Java_org_opencv_features2d_DescriptorExtractor_write_10(clzEnv,clzObj,(C.jlong)(nativeObj),tojstring(fileName))
}

func DescriptorMatcherNative_add_0(nativeObj int64,descriptors_mat_nativeObj int64) {
	C.Java_org_opencv_features2d_DescriptorMatcher_add_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(descriptors_mat_nativeObj))
}

func DescriptorMatcherNative_clear_0(nativeObj int64) {
	C.Java_org_opencv_features2d_DescriptorMatcher_clear_10(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func DescriptorMatcherNative_clone_0(nativeObj int64,emptyTrainData bool) int64 {
	return int64(C.Java_org_opencv_features2d_DescriptorMatcher_clone_10(clzEnv,clzObj,(C.jlong)(nativeObj),tojboolean(emptyTrainData)))
}

func DescriptorMatcherNative_clone_1(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_features2d_DescriptorMatcher_clone_11(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func DescriptorMatcherNative_create_0(descriptorMatcherType string) int64 {
	return int64(C.Java_org_opencv_features2d_DescriptorMatcher_create_10(clzEnv,clzObj,tojstring(descriptorMatcherType)))
}

func DescriptorMatcherNative_create_1(matcherType int) int64 {
	return int64(C.Java_org_opencv_features2d_DescriptorMatcher_create_11(clzEnv,clzObj,(C.jint)(matcherType)))
}

func DescriptorMatcherNative_delete(nativeObj int64) {
	C.Java_org_opencv_features2d_DescriptorMatcher_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func DescriptorMatcherNative_empty_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_features2d_DescriptorMatcher_empty_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func DescriptorMatcherNative_getTrainDescriptors_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_features2d_DescriptorMatcher_getTrainDescriptors_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func DescriptorMatcherNative_isMaskSupported_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_features2d_DescriptorMatcher_isMaskSupported_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func DescriptorMatcherNative_knnMatch_0(nativeObj int64,queryDescriptors_nativeObj int64,trainDescriptors_nativeObj int64,matches_mat_nativeObj int64,k int,mask_nativeObj int64,compactResult bool) {
	C.Java_org_opencv_features2d_DescriptorMatcher_knnMatch_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(queryDescriptors_nativeObj),(C.jlong)(trainDescriptors_nativeObj),(C.jlong)(matches_mat_nativeObj),(C.jint)(k),(C.jlong)(mask_nativeObj),tojboolean(compactResult))
}

func DescriptorMatcherNative_knnMatch_1(nativeObj int64,queryDescriptors_nativeObj int64,trainDescriptors_nativeObj int64,matches_mat_nativeObj int64,k int) {
	C.Java_org_opencv_features2d_DescriptorMatcher_knnMatch_11(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(queryDescriptors_nativeObj),(C.jlong)(trainDescriptors_nativeObj),(C.jlong)(matches_mat_nativeObj),(C.jint)(k))
}

func DescriptorMatcherNative_knnMatch_2(nativeObj int64,queryDescriptors_nativeObj int64,matches_mat_nativeObj int64,k int,masks_mat_nativeObj int64,compactResult bool) {
	C.Java_org_opencv_features2d_DescriptorMatcher_knnMatch_12(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(queryDescriptors_nativeObj),(C.jlong)(matches_mat_nativeObj),(C.jint)(k),(C.jlong)(masks_mat_nativeObj),tojboolean(compactResult))
}

func DescriptorMatcherNative_knnMatch_3(nativeObj int64,queryDescriptors_nativeObj int64,matches_mat_nativeObj int64,k int) {
	C.Java_org_opencv_features2d_DescriptorMatcher_knnMatch_13(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(queryDescriptors_nativeObj),(C.jlong)(matches_mat_nativeObj),(C.jint)(k))
}

func DescriptorMatcherNative_match_0(nativeObj int64,queryDescriptors_nativeObj int64,trainDescriptors_nativeObj int64,matches_mat_nativeObj int64,mask_nativeObj int64) {
	C.Java_org_opencv_features2d_DescriptorMatcher_match_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(queryDescriptors_nativeObj),(C.jlong)(trainDescriptors_nativeObj),(C.jlong)(matches_mat_nativeObj),(C.jlong)(mask_nativeObj))
}

func DescriptorMatcherNative_match_1(nativeObj int64,queryDescriptors_nativeObj int64,trainDescriptors_nativeObj int64,matches_mat_nativeObj int64) {
	C.Java_org_opencv_features2d_DescriptorMatcher_match_11(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(queryDescriptors_nativeObj),(C.jlong)(trainDescriptors_nativeObj),(C.jlong)(matches_mat_nativeObj))
}

func DescriptorMatcherNative_match_2(nativeObj int64,queryDescriptors_nativeObj int64,matches_mat_nativeObj int64,masks_mat_nativeObj int64) {
	C.Java_org_opencv_features2d_DescriptorMatcher_match_12(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(queryDescriptors_nativeObj),(C.jlong)(matches_mat_nativeObj),(C.jlong)(masks_mat_nativeObj))
}

func DescriptorMatcherNative_match_3(nativeObj int64,queryDescriptors_nativeObj int64,matches_mat_nativeObj int64) {
	C.Java_org_opencv_features2d_DescriptorMatcher_match_13(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(queryDescriptors_nativeObj),(C.jlong)(matches_mat_nativeObj))
}

func DescriptorMatcherNative_radiusMatch_0(nativeObj int64,queryDescriptors_nativeObj int64,trainDescriptors_nativeObj int64,matches_mat_nativeObj int64,maxDistance float32,mask_nativeObj int64,compactResult bool) {
	C.Java_org_opencv_features2d_DescriptorMatcher_radiusMatch_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(queryDescriptors_nativeObj),(C.jlong)(trainDescriptors_nativeObj),(C.jlong)(matches_mat_nativeObj),(C.jfloat)(maxDistance),(C.jlong)(mask_nativeObj),tojboolean(compactResult))
}

func DescriptorMatcherNative_radiusMatch_1(nativeObj int64,queryDescriptors_nativeObj int64,trainDescriptors_nativeObj int64,matches_mat_nativeObj int64,maxDistance float32) {
	C.Java_org_opencv_features2d_DescriptorMatcher_radiusMatch_11(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(queryDescriptors_nativeObj),(C.jlong)(trainDescriptors_nativeObj),(C.jlong)(matches_mat_nativeObj),(C.jfloat)(maxDistance))
}

func DescriptorMatcherNative_radiusMatch_2(nativeObj int64,queryDescriptors_nativeObj int64,matches_mat_nativeObj int64,maxDistance float32,masks_mat_nativeObj int64,compactResult bool) {
	C.Java_org_opencv_features2d_DescriptorMatcher_radiusMatch_12(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(queryDescriptors_nativeObj),(C.jlong)(matches_mat_nativeObj),(C.jfloat)(maxDistance),(C.jlong)(masks_mat_nativeObj),tojboolean(compactResult))
}

func DescriptorMatcherNative_radiusMatch_3(nativeObj int64,queryDescriptors_nativeObj int64,matches_mat_nativeObj int64,maxDistance float32) {
	C.Java_org_opencv_features2d_DescriptorMatcher_radiusMatch_13(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(queryDescriptors_nativeObj),(C.jlong)(matches_mat_nativeObj),(C.jfloat)(maxDistance))
}

func DescriptorMatcherNative_read_0(nativeObj int64,fileName string) {
	C.Java_org_opencv_features2d_DescriptorMatcher_read_10(clzEnv,clzObj,(C.jlong)(nativeObj),tojstring(fileName))
}

func DescriptorMatcherNative_train_0(nativeObj int64) {
	C.Java_org_opencv_features2d_DescriptorMatcher_train_10(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func DescriptorMatcherNative_write_0(nativeObj int64,fileName string) {
	C.Java_org_opencv_features2d_DescriptorMatcher_write_10(clzEnv,clzObj,(C.jlong)(nativeObj),tojstring(fileName))
}

func FastFeatureDetectorNative_create_0(threshold int,nonmaxSuppression bool,rtype int) int64 {
	return int64(C.Java_org_opencv_features2d_FastFeatureDetector_create_10(clzEnv,clzObj,(C.jint)(threshold),tojboolean(nonmaxSuppression),(C.jint)(rtype)))
}

func FastFeatureDetectorNative_create_1() int64 {
	return int64(C.Java_org_opencv_features2d_FastFeatureDetector_create_11(clzEnv,clzObj))
}

func FastFeatureDetectorNative_delete(nativeObj int64) {
	C.Java_org_opencv_features2d_FastFeatureDetector_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func FastFeatureDetectorNative_getDefaultName_0(nativeObj int64) string {
	return togostring(
			C.Java_org_opencv_features2d_FastFeatureDetector_getDefaultName_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func FastFeatureDetectorNative_getNonmaxSuppression_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_features2d_FastFeatureDetector_getNonmaxSuppression_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func FastFeatureDetectorNative_getThreshold_0(nativeObj int64) int {
	return int(C.Java_org_opencv_features2d_FastFeatureDetector_getThreshold_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func FastFeatureDetectorNative_getType_0(nativeObj int64) int {
	return int(C.Java_org_opencv_features2d_FastFeatureDetector_getType_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func FastFeatureDetectorNative_setNonmaxSuppression_0(nativeObj int64,f bool) {
	C.Java_org_opencv_features2d_FastFeatureDetector_setNonmaxSuppression_10(clzEnv,clzObj,(C.jlong)(nativeObj),tojboolean(f))
}

func FastFeatureDetectorNative_setThreshold_0(nativeObj int64,threshold int) {
	C.Java_org_opencv_features2d_FastFeatureDetector_setThreshold_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(threshold))
}

func FastFeatureDetectorNative_setType_0(nativeObj int64,rtype int) {
	C.Java_org_opencv_features2d_FastFeatureDetector_setType_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(rtype))
}

func Feature2DNative_compute_0(nativeObj int64,image_nativeObj int64,keypoints_mat_nativeObj int64,descriptors_nativeObj int64) {
	C.Java_org_opencv_features2d_Feature2D_compute_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(image_nativeObj),(C.jlong)(keypoints_mat_nativeObj),(C.jlong)(descriptors_nativeObj))
}

func Feature2DNative_compute_1(nativeObj int64,images_mat_nativeObj int64,keypoints_mat_nativeObj int64,descriptors_mat_nativeObj int64) {
	C.Java_org_opencv_features2d_Feature2D_compute_11(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(images_mat_nativeObj),(C.jlong)(keypoints_mat_nativeObj),(C.jlong)(descriptors_mat_nativeObj))
}

func Feature2DNative_defaultNorm_0(nativeObj int64) int {
	return int(C.Java_org_opencv_features2d_Feature2D_defaultNorm_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func Feature2DNative_delete(nativeObj int64) {
	C.Java_org_opencv_features2d_Feature2D_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func Feature2DNative_descriptorSize_0(nativeObj int64) int {
	return int(C.Java_org_opencv_features2d_Feature2D_descriptorSize_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func Feature2DNative_descriptorType_0(nativeObj int64) int {
	return int(C.Java_org_opencv_features2d_Feature2D_descriptorType_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func Feature2DNative_detectAndCompute_0(nativeObj int64,image_nativeObj int64,mask_nativeObj int64,keypoints_mat_nativeObj int64,descriptors_nativeObj int64,useProvidedKeypoints bool) {
	C.Java_org_opencv_features2d_Feature2D_detectAndCompute_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(image_nativeObj),(C.jlong)(mask_nativeObj),(C.jlong)(keypoints_mat_nativeObj),(C.jlong)(descriptors_nativeObj),tojboolean(useProvidedKeypoints))
}

func Feature2DNative_detectAndCompute_1(nativeObj int64,image_nativeObj int64,mask_nativeObj int64,keypoints_mat_nativeObj int64,descriptors_nativeObj int64) {
	C.Java_org_opencv_features2d_Feature2D_detectAndCompute_11(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(image_nativeObj),(C.jlong)(mask_nativeObj),(C.jlong)(keypoints_mat_nativeObj),(C.jlong)(descriptors_nativeObj))
}

func Feature2DNative_detect_0(nativeObj int64,image_nativeObj int64,keypoints_mat_nativeObj int64,mask_nativeObj int64) {
	C.Java_org_opencv_features2d_Feature2D_detect_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(image_nativeObj),(C.jlong)(keypoints_mat_nativeObj),(C.jlong)(mask_nativeObj))
}

func Feature2DNative_detect_1(nativeObj int64,image_nativeObj int64,keypoints_mat_nativeObj int64) {
	C.Java_org_opencv_features2d_Feature2D_detect_11(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(image_nativeObj),(C.jlong)(keypoints_mat_nativeObj))
}

func Feature2DNative_detect_2(nativeObj int64,images_mat_nativeObj int64,keypoints_mat_nativeObj int64,masks_mat_nativeObj int64) {
	C.Java_org_opencv_features2d_Feature2D_detect_12(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(images_mat_nativeObj),(C.jlong)(keypoints_mat_nativeObj),(C.jlong)(masks_mat_nativeObj))
}

func Feature2DNative_detect_3(nativeObj int64,images_mat_nativeObj int64,keypoints_mat_nativeObj int64) {
	C.Java_org_opencv_features2d_Feature2D_detect_13(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(images_mat_nativeObj),(C.jlong)(keypoints_mat_nativeObj))
}

func Feature2DNative_empty_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_features2d_Feature2D_empty_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func Feature2DNative_getDefaultName_0(nativeObj int64) string {
	return togostring(
			C.Java_org_opencv_features2d_Feature2D_getDefaultName_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func Feature2DNative_read_0(nativeObj int64,fileName string) {
	C.Java_org_opencv_features2d_Feature2D_read_10(clzEnv,clzObj,(C.jlong)(nativeObj),tojstring(fileName))
}

func Feature2DNative_write_0(nativeObj int64,fileName string) {
	C.Java_org_opencv_features2d_Feature2D_write_10(clzEnv,clzObj,(C.jlong)(nativeObj),tojstring(fileName))
}

func FeatureDetectorNative_create_0(detectorType int) int64 {
	return int64(C.Java_org_opencv_features2d_FeatureDetector_create_10(clzEnv,clzObj,(C.jint)(detectorType)))
}

func FeatureDetectorNative_delete(nativeObj int64) {
	C.Java_org_opencv_features2d_FeatureDetector_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func FeatureDetectorNative_detect_0(nativeObj int64,image_nativeObj int64,keypoints_mat_nativeObj int64,mask_nativeObj int64) {
	C.Java_org_opencv_features2d_FeatureDetector_detect_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(image_nativeObj),(C.jlong)(keypoints_mat_nativeObj),(C.jlong)(mask_nativeObj))
}

func FeatureDetectorNative_detect_1(nativeObj int64,image_nativeObj int64,keypoints_mat_nativeObj int64) {
	C.Java_org_opencv_features2d_FeatureDetector_detect_11(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(image_nativeObj),(C.jlong)(keypoints_mat_nativeObj))
}

func FeatureDetectorNative_detect_2(nativeObj int64,images_mat_nativeObj int64,keypoints_mat_nativeObj int64,masks_mat_nativeObj int64) {
	C.Java_org_opencv_features2d_FeatureDetector_detect_12(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(images_mat_nativeObj),(C.jlong)(keypoints_mat_nativeObj),(C.jlong)(masks_mat_nativeObj))
}

func FeatureDetectorNative_detect_3(nativeObj int64,images_mat_nativeObj int64,keypoints_mat_nativeObj int64) {
	C.Java_org_opencv_features2d_FeatureDetector_detect_13(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(images_mat_nativeObj),(C.jlong)(keypoints_mat_nativeObj))
}

func FeatureDetectorNative_empty_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_features2d_FeatureDetector_empty_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func FeatureDetectorNative_read_0(nativeObj int64,fileName string) {
	C.Java_org_opencv_features2d_FeatureDetector_read_10(clzEnv,clzObj,(C.jlong)(nativeObj),tojstring(fileName))
}

func FeatureDetectorNative_write_0(nativeObj int64,fileName string) {
	C.Java_org_opencv_features2d_FeatureDetector_write_10(clzEnv,clzObj,(C.jlong)(nativeObj),tojstring(fileName))
}

func Features2dNative_drawKeypoints_0(image_nativeObj int64,keypoints_mat_nativeObj int64,outImage_nativeObj int64,color_val0 float64,color_val1 float64,color_val2 float64,color_val3 float64,flags int) {
	C.Java_org_opencv_features2d_Features2d_drawKeypoints_10(clzEnv,clzObj,(C.jlong)(image_nativeObj),(C.jlong)(keypoints_mat_nativeObj),(C.jlong)(outImage_nativeObj),(C.jdouble)(color_val0),(C.jdouble)(color_val1),(C.jdouble)(
			color_val2),(C.jdouble)(color_val3),(C.jint)(flags))
}

func Features2dNative_drawKeypoints_1(image_nativeObj int64,keypoints_mat_nativeObj int64,outImage_nativeObj int64) {
	C.Java_org_opencv_features2d_Features2d_drawKeypoints_11(clzEnv,clzObj,(C.jlong)(image_nativeObj),(C.jlong)(keypoints_mat_nativeObj),(C.jlong)(outImage_nativeObj))
}

func Features2dNative_drawMatches2_0(img1_nativeObj int64,keypoints1_mat_nativeObj int64,img2_nativeObj int64,keypoints2_mat_nativeObj int64,matches1to2_mat_nativeObj int64,outImg_nativeObj int64,matchColor_val0 float64,matchColor_val1 float64,matchColor_val2 float64,matchColor_val3 float64,singlePointColor_val0 float64,singlePointColor_val1 float64,singlePointColor_val2 float64,singlePointColor_val3 float64,matchesMask_mat_nativeObj int64,flags int) {
	C.Java_org_opencv_features2d_Features2d_drawMatches2_10(clzEnv,clzObj,(C.jlong)(img1_nativeObj),(C.jlong)(keypoints1_mat_nativeObj),(C.jlong)(img2_nativeObj),(C.jlong)(keypoints2_mat_nativeObj),(C.jlong)(matches1to2_mat_nativeObj),(C.jlong)(outImg_nativeObj),(C.jdouble)(matchColor_val0),(C.jdouble)(matchColor_val1),(C.jdouble)(matchColor_val2),(C.jdouble)(matchColor_val3),(C.jdouble)(singlePointColor_val0),(C.jdouble)(singlePointColor_val1),(C.jdouble)(singlePointColor_val2),(C.jdouble)(singlePointColor_val3),(C.jlong)(matchesMask_mat_nativeObj),(C.jint)(flags))
}

func Features2dNative_drawMatches2_1(img1_nativeObj int64,keypoints1_mat_nativeObj int64,img2_nativeObj int64,keypoints2_mat_nativeObj int64,matches1to2_mat_nativeObj int64,outImg_nativeObj int64) {
	C.Java_org_opencv_features2d_Features2d_drawMatches2_11(clzEnv,clzObj,(C.jlong)(img1_nativeObj),(C.jlong)(keypoints1_mat_nativeObj),(C.jlong)(img2_nativeObj),(C.jlong)(keypoints2_mat_nativeObj),(C.jlong)(matches1to2_mat_nativeObj),(C.jlong)(outImg_nativeObj))
}

func Features2dNative_drawMatchesKnn_0(img1_nativeObj int64,keypoints1_mat_nativeObj int64,img2_nativeObj int64,keypoints2_mat_nativeObj int64,matches1to2_mat_nativeObj int64,outImg_nativeObj int64,matchColor_val0 float64,matchColor_val1 float64,matchColor_val2 float64,matchColor_val3 float64,singlePointColor_val0 float64,singlePointColor_val1 float64,singlePointColor_val2 float64,singlePointColor_val3 float64,matchesMask_mat_nativeObj int64,flags int) {
	C.Java_org_opencv_features2d_Features2d_drawMatchesKnn_10(clzEnv,clzObj,(C.jlong)(img1_nativeObj),(C.jlong)(keypoints1_mat_nativeObj),(C.jlong)(img2_nativeObj),(C.jlong)(keypoints2_mat_nativeObj),(C.jlong)(matches1to2_mat_nativeObj),(C.jlong)(outImg_nativeObj),(C.jdouble)(matchColor_val0),(C.jdouble)(matchColor_val1),(C.jdouble)(matchColor_val2),(C.jdouble)(matchColor_val3),(C.jdouble)(singlePointColor_val0),(C.jdouble)(singlePointColor_val1),(C.jdouble)(singlePointColor_val2),(C.jdouble)(singlePointColor_val3),(C.jlong)(matchesMask_mat_nativeObj),(C.jint)(flags))
}

func Features2dNative_drawMatchesKnn_1(img1_nativeObj int64,keypoints1_mat_nativeObj int64,img2_nativeObj int64,keypoints2_mat_nativeObj int64,matches1to2_mat_nativeObj int64,outImg_nativeObj int64) {
	C.Java_org_opencv_features2d_Features2d_drawMatchesKnn_11(clzEnv,clzObj,(C.jlong)(img1_nativeObj),(C.jlong)(keypoints1_mat_nativeObj),(C.jlong)(img2_nativeObj),(C.jlong)(keypoints2_mat_nativeObj),(C.jlong)(matches1to2_mat_nativeObj),(C.jlong)(outImg_nativeObj))
}

func Features2dNative_drawMatches_0(img1_nativeObj int64,keypoints1_mat_nativeObj int64,img2_nativeObj int64,keypoints2_mat_nativeObj int64,matches1to2_mat_nativeObj int64,outImg_nativeObj int64,matchColor_val0 float64,matchColor_val1 float64,matchColor_val2 float64,matchColor_val3 float64,singlePointColor_val0 float64,singlePointColor_val1 float64,singlePointColor_val2 float64,singlePointColor_val3 float64,matchesMask_mat_nativeObj int64,flags int) {
	C.Java_org_opencv_features2d_Features2d_drawMatches_10(clzEnv,clzObj,(C.jlong)(img1_nativeObj),(C.jlong)(keypoints1_mat_nativeObj),(C.jlong)(img2_nativeObj),(C.jlong)(keypoints2_mat_nativeObj),(C.jlong)(matches1to2_mat_nativeObj),(C.jlong)(outImg_nativeObj),(C.jdouble)(matchColor_val0),(C.jdouble)(matchColor_val1),(C.jdouble)(matchColor_val2),(C.jdouble)(matchColor_val3),(C.jdouble)(singlePointColor_val0),(C.jdouble)(singlePointColor_val1),(C.jdouble)(singlePointColor_val2),(C.jdouble)(singlePointColor_val3),(C.jlong)(matchesMask_mat_nativeObj),(C.jint)(flags))
}

func Features2dNative_drawMatches_1(img1_nativeObj int64,keypoints1_mat_nativeObj int64,img2_nativeObj int64,keypoints2_mat_nativeObj int64,matches1to2_mat_nativeObj int64,outImg_nativeObj int64) {
	C.Java_org_opencv_features2d_Features2d_drawMatches_11(clzEnv,clzObj,(C.jlong)(img1_nativeObj),(C.jlong)(keypoints1_mat_nativeObj),(C.jlong)(img2_nativeObj),(C.jlong)(keypoints2_mat_nativeObj),(C.jlong)(matches1to2_mat_nativeObj),(C.jlong)(outImg_nativeObj))
}

func FlannBasedMatcherNative_FlannBasedMatcher_0() int64 {
	return int64(C.Java_org_opencv_features2d_FlannBasedMatcher_FlannBasedMatcher_10(clzEnv,clzObj))
}

func FlannBasedMatcherNative_create_0() int64 {
	return int64(C.Java_org_opencv_features2d_FlannBasedMatcher_create_10(clzEnv,clzObj))
}

func FlannBasedMatcherNative_delete(nativeObj int64) {
	C.Java_org_opencv_features2d_FlannBasedMatcher_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func GFTTDetectorNative_create_0(maxCorners int,qualityLevel float64,minDistance float64,blockSize int,gradiantSize int,useHarrisDetector bool,k float64) int64 {
	return int64(C.Java_org_opencv_features2d_GFTTDetector_create_10(clzEnv,clzObj,(C.jint)(maxCorners),(C.jdouble)(qualityLevel),(C.jdouble)(minDistance),(C.jint)(blockSize),(C.jint)(gradiantSize),tojboolean(useHarrisDetector),(C.jdouble)(k)))
}

func GFTTDetectorNative_create_1(maxCorners int,qualityLevel float64,minDistance float64,blockSize int,gradiantSize int) int64 {
	return int64(C.Java_org_opencv_features2d_GFTTDetector_create_11(clzEnv,clzObj,(C.jint)(maxCorners),(C.jdouble)(qualityLevel),(C.jdouble)(minDistance),(C.jint)(blockSize),(C.jint)(gradiantSize)))
}

func GFTTDetectorNative_create_2(maxCorners int,qualityLevel float64,minDistance float64,blockSize int,useHarrisDetector bool,k float64) int64 {
	return int64(C.Java_org_opencv_features2d_GFTTDetector_create_12(clzEnv,clzObj,(C.jint)(maxCorners),(C.jdouble)(qualityLevel),(C.jdouble)(minDistance),(C.jint)(blockSize),tojboolean(useHarrisDetector),(C.jdouble)(k)))
}

func GFTTDetectorNative_create_3() int64 {
	return int64(C.Java_org_opencv_features2d_GFTTDetector_create_13(clzEnv,clzObj))
}

func GFTTDetectorNative_delete(nativeObj int64) {
	C.Java_org_opencv_features2d_GFTTDetector_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func GFTTDetectorNative_getBlockSize_0(nativeObj int64) int {
	return int(C.Java_org_opencv_features2d_GFTTDetector_getBlockSize_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func GFTTDetectorNative_getDefaultName_0(nativeObj int64) string {
	return togostring(
			C.Java_org_opencv_features2d_GFTTDetector_getDefaultName_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func GFTTDetectorNative_getHarrisDetector_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_features2d_GFTTDetector_getHarrisDetector_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func GFTTDetectorNative_getK_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_features2d_GFTTDetector_getK_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func GFTTDetectorNative_getMaxFeatures_0(nativeObj int64) int {
	return int(C.Java_org_opencv_features2d_GFTTDetector_getMaxFeatures_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func GFTTDetectorNative_getMinDistance_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_features2d_GFTTDetector_getMinDistance_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func GFTTDetectorNative_getQualityLevel_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_features2d_GFTTDetector_getQualityLevel_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func GFTTDetectorNative_setBlockSize_0(nativeObj int64,blockSize int) {
	C.Java_org_opencv_features2d_GFTTDetector_setBlockSize_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(blockSize))
}

func GFTTDetectorNative_setHarrisDetector_0(nativeObj int64,val bool) {
	C.Java_org_opencv_features2d_GFTTDetector_setHarrisDetector_10(clzEnv,clzObj,(C.jlong)(nativeObj),tojboolean(val))
}

func GFTTDetectorNative_setK_0(nativeObj int64,k float64) {
	C.Java_org_opencv_features2d_GFTTDetector_setK_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jdouble)(k))
}

func GFTTDetectorNative_setMaxFeatures_0(nativeObj int64,maxFeatures int) {
	C.Java_org_opencv_features2d_GFTTDetector_setMaxFeatures_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(maxFeatures))
}

func GFTTDetectorNative_setMinDistance_0(nativeObj int64,minDistance float64) {
	C.Java_org_opencv_features2d_GFTTDetector_setMinDistance_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jdouble)(minDistance))
}

func GFTTDetectorNative_setQualityLevel_0(nativeObj int64,qlevel float64) {
	C.Java_org_opencv_features2d_GFTTDetector_setQualityLevel_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jdouble)(qlevel))
}

func KAZENative_create_0(extended bool,upright bool,threshold float32,nOctaves int,nOctaveLayers int,diffusivity int) int64 {
	return int64(C.Java_org_opencv_features2d_KAZE_create_10(clzEnv,clzObj,tojboolean(extended),tojboolean(upright),(C.jfloat)(threshold),(C.jint)(nOctaves),(C.jint)(nOctaveLayers),(C.jint)(diffusivity)))
}

func KAZENative_create_1() int64 {
	return int64(C.Java_org_opencv_features2d_KAZE_create_11(clzEnv,clzObj))
}

func KAZENative_delete(nativeObj int64) {
	C.Java_org_opencv_features2d_KAZE_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func KAZENative_getDefaultName_0(nativeObj int64) string {
	return togostring(
			C.Java_org_opencv_features2d_KAZE_getDefaultName_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func KAZENative_getDiffusivity_0(nativeObj int64) int {
	return int(C.Java_org_opencv_features2d_KAZE_getDiffusivity_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func KAZENative_getExtended_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_features2d_KAZE_getExtended_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func KAZENative_getNOctaveLayers_0(nativeObj int64) int {
	return int(C.Java_org_opencv_features2d_KAZE_getNOctaveLayers_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func KAZENative_getNOctaves_0(nativeObj int64) int {
	return int(C.Java_org_opencv_features2d_KAZE_getNOctaves_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func KAZENative_getThreshold_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_features2d_KAZE_getThreshold_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func KAZENative_getUpright_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_features2d_KAZE_getUpright_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func KAZENative_setDiffusivity_0(nativeObj int64,diff int) {
	C.Java_org_opencv_features2d_KAZE_setDiffusivity_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(diff))
}

func KAZENative_setExtended_0(nativeObj int64,extended bool) {
	C.Java_org_opencv_features2d_KAZE_setExtended_10(clzEnv,clzObj,(C.jlong)(nativeObj),tojboolean(extended))
}

func KAZENative_setNOctaveLayers_0(nativeObj int64,octaveLayers int) {
	C.Java_org_opencv_features2d_KAZE_setNOctaveLayers_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(octaveLayers))
}

func KAZENative_setNOctaves_0(nativeObj int64,octaves int) {
	C.Java_org_opencv_features2d_KAZE_setNOctaves_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(octaves))
}

func KAZENative_setThreshold_0(nativeObj int64,threshold float64) {
	C.Java_org_opencv_features2d_KAZE_setThreshold_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jdouble)(threshold))
}

func KAZENative_setUpright_0(nativeObj int64,upright bool) {
	C.Java_org_opencv_features2d_KAZE_setUpright_10(clzEnv,clzObj,(C.jlong)(nativeObj),tojboolean(upright))
}

func MSERNative_create_0(_delta int,_min_area int,_max_area int,_max_variation float64,_min_diversity float64,_max_evolution int,_area_threshold float64,_min_margin float64,_edge_blur_size int) int64 {
	return int64(C.Java_org_opencv_features2d_MSER_create_10(clzEnv,clzObj,(C.jint)(_delta),(C.jint)(_min_area),(C.jint)(_max_area),(C.jdouble)(_max_variation),(C.jdouble)(_min_diversity),(C.jint)(_max_evolution),(C.jdouble)(_area_threshold),(C.jdouble)(_min_margin),(C.jint)(_edge_blur_size)))
}

func MSERNative_create_1() int64 {
	return int64(C.Java_org_opencv_features2d_MSER_create_11(clzEnv,clzObj))
}

func MSERNative_delete(nativeObj int64) {
	C.Java_org_opencv_features2d_MSER_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func MSERNative_detectRegions_0(nativeObj int64,image_nativeObj int64,msers_mat_nativeObj int64,bboxes_mat_nativeObj int64) {
	C.Java_org_opencv_features2d_MSER_detectRegions_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(image_nativeObj),(C.jlong)(msers_mat_nativeObj),(C.jlong)(bboxes_mat_nativeObj))
}

func MSERNative_getDefaultName_0(nativeObj int64) string {
	return togostring(
			C.Java_org_opencv_features2d_MSER_getDefaultName_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func MSERNative_getDelta_0(nativeObj int64) int {
	return int(C.Java_org_opencv_features2d_MSER_getDelta_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func MSERNative_getMaxArea_0(nativeObj int64) int {
	return int(C.Java_org_opencv_features2d_MSER_getMaxArea_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func MSERNative_getMinArea_0(nativeObj int64) int {
	return int(C.Java_org_opencv_features2d_MSER_getMinArea_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func MSERNative_getPass2Only_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_features2d_MSER_getPass2Only_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func MSERNative_setDelta_0(nativeObj int64,delta int) {
	C.Java_org_opencv_features2d_MSER_setDelta_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(delta))
}

func MSERNative_setMaxArea_0(nativeObj int64,maxArea int) {
	C.Java_org_opencv_features2d_MSER_setMaxArea_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(maxArea))
}

func MSERNative_setMinArea_0(nativeObj int64,minArea int) {
	C.Java_org_opencv_features2d_MSER_setMinArea_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(minArea))
}

func MSERNative_setPass2Only_0(nativeObj int64,f bool) {
	C.Java_org_opencv_features2d_MSER_setPass2Only_10(clzEnv,clzObj,(C.jlong)(nativeObj),tojboolean(f))
}

func ORBNative_create_0(nfeatures int,scaleFactor float32,nlevels int,edgeThreshold int,firstLevel int,WTA_K int,scoreType int,patchSize int,fastThreshold int) int64 {
	return int64(C.Java_org_opencv_features2d_ORB_create_10(clzEnv,clzObj,(C.jint)(nfeatures),(C.jfloat)(scaleFactor),(C.jint)(nlevels),(C.jint)(edgeThreshold),(C.jint)(firstLevel),(C.jint)(WTA_K),(C.jint)(scoreType),(C.jint)(patchSize),(C.jint)(fastThreshold)))
}

func ORBNative_create_1() int64 {
	return int64(C.Java_org_opencv_features2d_ORB_create_11(clzEnv,clzObj))
}

func ORBNative_delete(nativeObj int64) {
	C.Java_org_opencv_features2d_ORB_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func ORBNative_getDefaultName_0(nativeObj int64) string {
	return togostring(
			C.Java_org_opencv_features2d_ORB_getDefaultName_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ORBNative_getEdgeThreshold_0(nativeObj int64) int {
	return int(C.Java_org_opencv_features2d_ORB_getEdgeThreshold_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ORBNative_getFastThreshold_0(nativeObj int64) int {
	return int(C.Java_org_opencv_features2d_ORB_getFastThreshold_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ORBNative_getFirstLevel_0(nativeObj int64) int {
	return int(C.Java_org_opencv_features2d_ORB_getFirstLevel_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ORBNative_getMaxFeatures_0(nativeObj int64) int {
	return int(C.Java_org_opencv_features2d_ORB_getMaxFeatures_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ORBNative_getNLevels_0(nativeObj int64) int {
	return int(C.Java_org_opencv_features2d_ORB_getNLevels_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ORBNative_getPatchSize_0(nativeObj int64) int {
	return int(C.Java_org_opencv_features2d_ORB_getPatchSize_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ORBNative_getScaleFactor_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_features2d_ORB_getScaleFactor_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ORBNative_getScoreType_0(nativeObj int64) int {
	return int(C.Java_org_opencv_features2d_ORB_getScoreType_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ORBNative_getWTA_K_0(nativeObj int64) int {
	return int(C.Java_org_opencv_features2d_ORB_getWTA_1K_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ORBNative_setEdgeThreshold_0(nativeObj int64,edgeThreshold int) {
	C.Java_org_opencv_features2d_ORB_setEdgeThreshold_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(edgeThreshold))
}

func ORBNative_setFastThreshold_0(nativeObj int64,fastThreshold int) {
	C.Java_org_opencv_features2d_ORB_setFastThreshold_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(fastThreshold))
}

func ORBNative_setFirstLevel_0(nativeObj int64,firstLevel int) {
	C.Java_org_opencv_features2d_ORB_setFirstLevel_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(firstLevel))
}

func ORBNative_setMaxFeatures_0(nativeObj int64,maxFeatures int) {
	C.Java_org_opencv_features2d_ORB_setMaxFeatures_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(maxFeatures))
}

func ORBNative_setNLevels_0(nativeObj int64,nlevels int) {
	C.Java_org_opencv_features2d_ORB_setNLevels_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(nlevels))
}

func ORBNative_setPatchSize_0(nativeObj int64,patchSize int) {
	C.Java_org_opencv_features2d_ORB_setPatchSize_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(patchSize))
}

func ORBNative_setScaleFactor_0(nativeObj int64,scaleFactor float64) {
	C.Java_org_opencv_features2d_ORB_setScaleFactor_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jdouble)(scaleFactor))
}

func ORBNative_setScoreType_0(nativeObj int64,scoreType int) {
	C.Java_org_opencv_features2d_ORB_setScoreType_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(scoreType))
}

func ORBNative_setWTA_K_0(nativeObj int64,wta_k int) {
	C.Java_org_opencv_features2d_ORB_setWTA_1K_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(wta_k))
}

func ParamsNative_Params_0() int64 {
	return int64(C.Java_org_opencv_features2d_Params_Params_10(clzEnv,clzObj))
}

func ParamsNative_delete(nativeObj int64) {
	C.Java_org_opencv_features2d_Params_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func ParamsNative_get_filterByArea_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_features2d_Params_get_1filterByArea_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ParamsNative_get_filterByCircularity_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_features2d_Params_get_1filterByCircularity_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ParamsNative_get_filterByColor_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_features2d_Params_get_1filterByColor_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ParamsNative_get_filterByConvexity_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_features2d_Params_get_1filterByConvexity_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ParamsNative_get_filterByInertia_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_features2d_Params_get_1filterByInertia_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ParamsNative_get_maxArea_0(nativeObj int64) float32 {
	return float32(C.Java_org_opencv_features2d_Params_get_1maxArea_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ParamsNative_get_maxCircularity_0(nativeObj int64) float32 {
	return float32(C.Java_org_opencv_features2d_Params_get_1maxCircularity_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ParamsNative_get_maxConvexity_0(nativeObj int64) float32 {
	return float32(C.Java_org_opencv_features2d_Params_get_1maxConvexity_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ParamsNative_get_maxInertiaRatio_0(nativeObj int64) float32 {
	return float32(C.Java_org_opencv_features2d_Params_get_1maxInertiaRatio_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ParamsNative_get_maxThreshold_0(nativeObj int64) float32 {
	return float32(C.Java_org_opencv_features2d_Params_get_1maxThreshold_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ParamsNative_get_minArea_0(nativeObj int64) float32 {
	return float32(C.Java_org_opencv_features2d_Params_get_1minArea_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ParamsNative_get_minCircularity_0(nativeObj int64) float32 {
	return float32(C.Java_org_opencv_features2d_Params_get_1minCircularity_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ParamsNative_get_minConvexity_0(nativeObj int64) float32 {
	return float32(C.Java_org_opencv_features2d_Params_get_1minConvexity_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ParamsNative_get_minDistBetweenBlobs_0(nativeObj int64) float32 {
	return float32(C.Java_org_opencv_features2d_Params_get_1minDistBetweenBlobs_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ParamsNative_get_minInertiaRatio_0(nativeObj int64) float32 {
	return float32(C.Java_org_opencv_features2d_Params_get_1minInertiaRatio_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ParamsNative_get_minRepeatability_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_features2d_Params_get_1minRepeatability_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ParamsNative_get_minThreshold_0(nativeObj int64) float32 {
	return float32(C.Java_org_opencv_features2d_Params_get_1minThreshold_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ParamsNative_get_thresholdStep_0(nativeObj int64) float32 {
	return float32(C.Java_org_opencv_features2d_Params_get_1thresholdStep_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ParamsNative_set_filterByArea_0(nativeObj int64,filterByArea bool) {
	C.Java_org_opencv_features2d_Params_set_1filterByArea_10(clzEnv,clzObj,(C.jlong)(nativeObj),tojboolean(filterByArea))
}

func ParamsNative_set_filterByCircularity_0(nativeObj int64,filterByCircularity bool) {
	C.Java_org_opencv_features2d_Params_set_1filterByCircularity_10(clzEnv,clzObj,(C.jlong)(nativeObj),tojboolean(filterByCircularity))
}

func ParamsNative_set_filterByColor_0(nativeObj int64,filterByColor bool) {
	C.Java_org_opencv_features2d_Params_set_1filterByColor_10(clzEnv,clzObj,(C.jlong)(nativeObj),tojboolean(filterByColor))
}

func ParamsNative_set_filterByConvexity_0(nativeObj int64,filterByConvexity bool) {
	C.Java_org_opencv_features2d_Params_set_1filterByConvexity_10(clzEnv,clzObj,(C.jlong)(nativeObj),tojboolean(filterByConvexity))
}

func ParamsNative_set_filterByInertia_0(nativeObj int64,filterByInertia bool) {
	C.Java_org_opencv_features2d_Params_set_1filterByInertia_10(clzEnv,clzObj,(C.jlong)(nativeObj),tojboolean(filterByInertia))
}

func ParamsNative_set_maxArea_0(nativeObj int64,maxArea float32) {
	C.Java_org_opencv_features2d_Params_set_1maxArea_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jfloat)(maxArea))
}

func ParamsNative_set_maxCircularity_0(nativeObj int64,maxCircularity float32) {
	C.Java_org_opencv_features2d_Params_set_1maxCircularity_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jfloat)(maxCircularity))
}

func ParamsNative_set_maxConvexity_0(nativeObj int64,maxConvexity float32) {
	C.Java_org_opencv_features2d_Params_set_1maxConvexity_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jfloat)(maxConvexity))
}

func ParamsNative_set_maxInertiaRatio_0(nativeObj int64,maxInertiaRatio float32) {
	C.Java_org_opencv_features2d_Params_set_1maxInertiaRatio_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jfloat)(maxInertiaRatio))
}

func ParamsNative_set_maxThreshold_0(nativeObj int64,maxThreshold float32) {
	C.Java_org_opencv_features2d_Params_set_1maxThreshold_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jfloat)(maxThreshold))
}

func ParamsNative_set_minArea_0(nativeObj int64,minArea float32) {
	C.Java_org_opencv_features2d_Params_set_1minArea_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jfloat)(minArea))
}

func ParamsNative_set_minCircularity_0(nativeObj int64,minCircularity float32) {
	C.Java_org_opencv_features2d_Params_set_1minCircularity_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jfloat)(minCircularity))
}

func ParamsNative_set_minConvexity_0(nativeObj int64,minConvexity float32) {
	C.Java_org_opencv_features2d_Params_set_1minConvexity_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jfloat)(minConvexity))
}

func ParamsNative_set_minDistBetweenBlobs_0(nativeObj int64,minDistBetweenBlobs float32) {
	C.Java_org_opencv_features2d_Params_set_1minDistBetweenBlobs_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jfloat)(minDistBetweenBlobs))
}

func ParamsNative_set_minInertiaRatio_0(nativeObj int64,minInertiaRatio float32) {
	C.Java_org_opencv_features2d_Params_set_1minInertiaRatio_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jfloat)(minInertiaRatio))
}

func ParamsNative_set_minRepeatability_0(nativeObj int64,minRepeatability int64) {
	C.Java_org_opencv_features2d_Params_set_1minRepeatability_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(minRepeatability))
}

func ParamsNative_set_minThreshold_0(nativeObj int64,minThreshold float32) {
	C.Java_org_opencv_features2d_Params_set_1minThreshold_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jfloat)(minThreshold))
}

func ParamsNative_set_thresholdStep_0(nativeObj int64,thresholdStep float32) {
	C.Java_org_opencv_features2d_Params_set_1thresholdStep_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jfloat)(thresholdStep))
}
