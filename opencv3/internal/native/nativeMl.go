package core

/*

#include "jni.h"

extern jlong Java_org_opencv_ml_ANN_1MLP_create_10(JNIEnv*, jclass);
extern void Java_org_opencv_ml_ANN_1MLP_delete(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_ml_ANN_1MLP_getBackpropMomentumScale_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_ml_ANN_1MLP_getBackpropWeightScale_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_ml_ANN_1MLP_getLayerSizes_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_ml_ANN_1MLP_getRpropDW0_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_ml_ANN_1MLP_getRpropDWMax_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_ml_ANN_1MLP_getRpropDWMin_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_ml_ANN_1MLP_getRpropDWMinus_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_ml_ANN_1MLP_getRpropDWPlus_10(JNIEnv*, jclass, jlong);
extern jdoubleArray Java_org_opencv_ml_ANN_1MLP_getTermCriteria_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_ml_ANN_1MLP_getTrainMethod_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_ml_ANN_1MLP_getWeights_10(JNIEnv*, jclass, jlong, jint);
extern jlong Java_org_opencv_ml_ANN_1MLP_load_10(JNIEnv*, jclass, jstring);
extern void Java_org_opencv_ml_ANN_1MLP_setActivationFunction_10(JNIEnv*, jclass, jlong, jint, jdouble, jdouble);
extern void Java_org_opencv_ml_ANN_1MLP_setActivationFunction_11(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_ml_ANN_1MLP_setBackpropMomentumScale_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_ml_ANN_1MLP_setBackpropWeightScale_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_ml_ANN_1MLP_setLayerSizes_10(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_ml_ANN_1MLP_setRpropDW0_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_ml_ANN_1MLP_setRpropDWMax_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_ml_ANN_1MLP_setRpropDWMin_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_ml_ANN_1MLP_setRpropDWMinus_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_ml_ANN_1MLP_setRpropDWPlus_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_ml_ANN_1MLP_setTermCriteria_10(JNIEnv*, jclass, jlong, jint, jint, jdouble);
extern void Java_org_opencv_ml_ANN_1MLP_setTrainMethod_10(JNIEnv*, jclass, jlong, jint, jdouble, jdouble);
extern void Java_org_opencv_ml_ANN_1MLP_setTrainMethod_11(JNIEnv*, jclass, jlong, jint);
extern jlong Java_org_opencv_ml_Boost_create_10(JNIEnv*, jclass);
extern void Java_org_opencv_ml_Boost_delete(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_ml_Boost_getBoostType_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_ml_Boost_getWeakCount_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_ml_Boost_getWeightTrimRate_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_ml_Boost_load_10(JNIEnv*, jclass, jstring, jstring);
extern jlong Java_org_opencv_ml_Boost_load_11(JNIEnv*, jclass, jstring);
extern void Java_org_opencv_ml_Boost_setBoostType_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_ml_Boost_setWeakCount_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_ml_Boost_setWeightTrimRate_10(JNIEnv*, jclass, jlong, jdouble);
extern jlong Java_org_opencv_ml_DTrees_create_10(JNIEnv*, jclass);
extern void Java_org_opencv_ml_DTrees_delete(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_ml_DTrees_getCVFolds_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_ml_DTrees_getMaxCategories_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_ml_DTrees_getMaxDepth_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_ml_DTrees_getMinSampleCount_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_ml_DTrees_getPriors_10(JNIEnv*, jclass, jlong);
extern jfloat Java_org_opencv_ml_DTrees_getRegressionAccuracy_10(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_ml_DTrees_getTruncatePrunedTree_10(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_ml_DTrees_getUse1SERule_10(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_ml_DTrees_getUseSurrogates_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_ml_DTrees_load_10(JNIEnv*, jclass, jstring, jstring);
extern jlong Java_org_opencv_ml_DTrees_load_11(JNIEnv*, jclass, jstring);
extern void Java_org_opencv_ml_DTrees_setCVFolds_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_ml_DTrees_setMaxCategories_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_ml_DTrees_setMaxDepth_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_ml_DTrees_setMinSampleCount_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_ml_DTrees_setPriors_10(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_ml_DTrees_setRegressionAccuracy_10(JNIEnv*, jclass, jlong, jfloat);
extern void Java_org_opencv_ml_DTrees_setTruncatePrunedTree_10(JNIEnv*, jclass, jlong, jboolean);
extern void Java_org_opencv_ml_DTrees_setUse1SERule_10(JNIEnv*, jclass, jlong, jboolean);
extern void Java_org_opencv_ml_DTrees_setUseSurrogates_10(JNIEnv*, jclass, jlong, jboolean);
extern jlong Java_org_opencv_ml_EM_create_10(JNIEnv*, jclass);
extern void Java_org_opencv_ml_EM_delete(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_ml_EM_getClustersNumber_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_ml_EM_getCovarianceMatrixType_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_ml_EM_getCovs_10(JNIEnv*, jclass, jlong, jlong);
extern jlong Java_org_opencv_ml_EM_getMeans_10(JNIEnv*, jclass, jlong);
extern jdoubleArray Java_org_opencv_ml_EM_getTermCriteria_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_ml_EM_getWeights_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_ml_EM_load_10(JNIEnv*, jclass, jstring, jstring);
extern jlong Java_org_opencv_ml_EM_load_11(JNIEnv*, jclass, jstring);
extern jdoubleArray Java_org_opencv_ml_EM_predict2_10(JNIEnv*, jclass, jlong, jlong, jlong);
extern jfloat Java_org_opencv_ml_EM_predict_10(JNIEnv*, jclass, jlong, jlong, jlong, jint);
extern jfloat Java_org_opencv_ml_EM_predict_11(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_ml_EM_setClustersNumber_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_ml_EM_setCovarianceMatrixType_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_ml_EM_setTermCriteria_10(JNIEnv*, jclass, jlong, jint, jint, jdouble);
extern jboolean Java_org_opencv_ml_EM_trainEM_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong);
extern jboolean Java_org_opencv_ml_EM_trainEM_11(JNIEnv*, jclass, jlong, jlong);
extern jboolean Java_org_opencv_ml_EM_trainE_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong, jlong, jlong);
extern jboolean Java_org_opencv_ml_EM_trainE_11(JNIEnv*, jclass, jlong, jlong, jlong);
extern jboolean Java_org_opencv_ml_EM_trainM_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jlong, jlong);
extern jboolean Java_org_opencv_ml_EM_trainM_11(JNIEnv*, jclass, jlong, jlong, jlong);
extern jlong Java_org_opencv_ml_KNearest_create_10(JNIEnv*, jclass);
extern void Java_org_opencv_ml_KNearest_delete(JNIEnv*, jclass, jlong);
extern jfloat Java_org_opencv_ml_KNearest_findNearest_10(JNIEnv*, jclass, jlong, jlong, jint, jlong, jlong, jlong);
extern jfloat Java_org_opencv_ml_KNearest_findNearest_11(JNIEnv*, jclass, jlong, jlong, jint, jlong);
extern jint Java_org_opencv_ml_KNearest_getAlgorithmType_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_ml_KNearest_getDefaultK_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_ml_KNearest_getEmax_10(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_ml_KNearest_getIsClassifier_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_ml_KNearest_setAlgorithmType_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_ml_KNearest_setDefaultK_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_ml_KNearest_setEmax_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_ml_KNearest_setIsClassifier_10(JNIEnv*, jclass, jlong, jboolean);
extern jlong Java_org_opencv_ml_LogisticRegression_create_10(JNIEnv*, jclass);
extern void Java_org_opencv_ml_LogisticRegression_delete(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_ml_LogisticRegression_getIterations_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_ml_LogisticRegression_getLearningRate_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_ml_LogisticRegression_getMiniBatchSize_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_ml_LogisticRegression_getRegularization_10(JNIEnv*, jclass, jlong);
extern jdoubleArray Java_org_opencv_ml_LogisticRegression_getTermCriteria_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_ml_LogisticRegression_getTrainMethod_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_ml_LogisticRegression_get_1learnt_1thetas_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_ml_LogisticRegression_load_10(JNIEnv*, jclass, jstring, jstring);
extern jlong Java_org_opencv_ml_LogisticRegression_load_11(JNIEnv*, jclass, jstring);
extern jfloat Java_org_opencv_ml_LogisticRegression_predict_10(JNIEnv*, jclass, jlong, jlong, jlong, jint);
extern jfloat Java_org_opencv_ml_LogisticRegression_predict_11(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_ml_LogisticRegression_setIterations_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_ml_LogisticRegression_setLearningRate_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_ml_LogisticRegression_setMiniBatchSize_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_ml_LogisticRegression_setRegularization_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_ml_LogisticRegression_setTermCriteria_10(JNIEnv*, jclass, jlong, jint, jint, jdouble);
extern void Java_org_opencv_ml_LogisticRegression_setTrainMethod_10(JNIEnv*, jclass, jlong, jint);
extern jlong Java_org_opencv_ml_NormalBayesClassifier_create_10(JNIEnv*, jclass);
extern void Java_org_opencv_ml_NormalBayesClassifier_delete(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_ml_NormalBayesClassifier_load_10(JNIEnv*, jclass, jstring, jstring);
extern jlong Java_org_opencv_ml_NormalBayesClassifier_load_11(JNIEnv*, jclass, jstring);
extern jfloat Java_org_opencv_ml_NormalBayesClassifier_predictProb_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong, jint);
extern jfloat Java_org_opencv_ml_NormalBayesClassifier_predictProb_11(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern jlong Java_org_opencv_ml_ParamGrid_create_10(JNIEnv*, jclass, jdouble, jdouble, jdouble);
extern jlong Java_org_opencv_ml_ParamGrid_create_11(JNIEnv*, jclass);
extern void Java_org_opencv_ml_ParamGrid_delete(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_ml_ParamGrid_get_1logStep_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_ml_ParamGrid_get_1maxVal_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_ml_ParamGrid_get_1minVal_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_ml_ParamGrid_set_1logStep_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_ml_ParamGrid_set_1maxVal_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_ml_ParamGrid_set_1minVal_10(JNIEnv*, jclass, jlong, jdouble);
extern jlong Java_org_opencv_ml_RTrees_create_10(JNIEnv*, jclass);
extern void Java_org_opencv_ml_RTrees_delete(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_ml_RTrees_getActiveVarCount_10(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_ml_RTrees_getCalculateVarImportance_10(JNIEnv*, jclass, jlong);
extern jdoubleArray Java_org_opencv_ml_RTrees_getTermCriteria_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_ml_RTrees_getVarImportance_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_ml_RTrees_getVotes_10(JNIEnv*, jclass, jlong, jlong, jlong, jint);
extern jlong Java_org_opencv_ml_RTrees_load_10(JNIEnv*, jclass, jstring, jstring);
extern jlong Java_org_opencv_ml_RTrees_load_11(JNIEnv*, jclass, jstring);
extern void Java_org_opencv_ml_RTrees_setActiveVarCount_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_ml_RTrees_setCalculateVarImportance_10(JNIEnv*, jclass, jlong, jboolean);
extern void Java_org_opencv_ml_RTrees_setTermCriteria_10(JNIEnv*, jclass, jlong, jint, jint, jdouble);
extern jlong Java_org_opencv_ml_SVM_create_10(JNIEnv*, jclass);
extern void Java_org_opencv_ml_SVM_delete(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_ml_SVM_getC_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_ml_SVM_getClassWeights_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_ml_SVM_getCoef0_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_ml_SVM_getDecisionFunction_10(JNIEnv*, jclass, jlong, jint, jlong, jlong);
extern jlong Java_org_opencv_ml_SVM_getDefaultGridPtr_10(JNIEnv*, jclass, jint);
extern jdouble Java_org_opencv_ml_SVM_getDegree_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_ml_SVM_getGamma_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_ml_SVM_getKernelType_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_ml_SVM_getNu_10(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_ml_SVM_getP_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_ml_SVM_getSupportVectors_10(JNIEnv*, jclass, jlong);
extern jdoubleArray Java_org_opencv_ml_SVM_getTermCriteria_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_ml_SVM_getType_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_ml_SVM_getUncompressedSupportVectors_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_ml_SVM_load_10(JNIEnv*, jclass, jstring);
extern void Java_org_opencv_ml_SVM_setC_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_ml_SVM_setClassWeights_10(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_ml_SVM_setCoef0_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_ml_SVM_setDegree_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_ml_SVM_setGamma_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_ml_SVM_setKernel_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_ml_SVM_setNu_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_ml_SVM_setP_10(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_ml_SVM_setTermCriteria_10(JNIEnv*, jclass, jlong, jint, jint, jdouble);
extern void Java_org_opencv_ml_SVM_setType_10(JNIEnv*, jclass, jlong, jint);
extern jboolean Java_org_opencv_ml_SVM_trainAuto_10(JNIEnv*, jclass, jlong, jlong, jint, jlong, jint, jlong, jlong, jlong, jlong, jlong, jlong, jboolean);
extern jboolean Java_org_opencv_ml_SVM_trainAuto_11(JNIEnv*, jclass, jlong, jlong, jint, jlong);
extern jlong Java_org_opencv_ml_SVMSGD_create_10(JNIEnv*, jclass);
extern void Java_org_opencv_ml_SVMSGD_delete(JNIEnv*, jclass, jlong);
extern jfloat Java_org_opencv_ml_SVMSGD_getInitialStepSize_10(JNIEnv*, jclass, jlong);
extern jfloat Java_org_opencv_ml_SVMSGD_getMarginRegularization_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_ml_SVMSGD_getMarginType_10(JNIEnv*, jclass, jlong);
extern jfloat Java_org_opencv_ml_SVMSGD_getShift_10(JNIEnv*, jclass, jlong);
extern jfloat Java_org_opencv_ml_SVMSGD_getStepDecreasingPower_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_ml_SVMSGD_getSvmsgdType_10(JNIEnv*, jclass, jlong);
extern jdoubleArray Java_org_opencv_ml_SVMSGD_getTermCriteria_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_ml_SVMSGD_getWeights_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_ml_SVMSGD_load_10(JNIEnv*, jclass, jstring, jstring);
extern jlong Java_org_opencv_ml_SVMSGD_load_11(JNIEnv*, jclass, jstring);
extern void Java_org_opencv_ml_SVMSGD_setInitialStepSize_10(JNIEnv*, jclass, jlong, jfloat);
extern void Java_org_opencv_ml_SVMSGD_setMarginRegularization_10(JNIEnv*, jclass, jlong, jfloat);
extern void Java_org_opencv_ml_SVMSGD_setMarginType_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_ml_SVMSGD_setOptimalParameters_10(JNIEnv*, jclass, jlong, jint, jint);
extern void Java_org_opencv_ml_SVMSGD_setOptimalParameters_11(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_ml_SVMSGD_setStepDecreasingPower_10(JNIEnv*, jclass, jlong, jfloat);
extern void Java_org_opencv_ml_SVMSGD_setSvmsgdType_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_ml_SVMSGD_setTermCriteria_10(JNIEnv*, jclass, jlong, jint, jint, jdouble);
extern jfloat Java_org_opencv_ml_StatModel_calcError_10(JNIEnv*, jclass, jlong, jlong, jboolean, jlong);
extern void Java_org_opencv_ml_StatModel_delete(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_ml_StatModel_empty_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_ml_StatModel_getVarCount_10(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_ml_StatModel_isClassifier_10(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_ml_StatModel_isTrained_10(JNIEnv*, jclass, jlong);
extern jfloat Java_org_opencv_ml_StatModel_predict_10(JNIEnv*, jclass, jlong, jlong, jlong, jint);
extern jfloat Java_org_opencv_ml_StatModel_predict_11(JNIEnv*, jclass, jlong, jlong);
extern jboolean Java_org_opencv_ml_StatModel_train_10(JNIEnv*, jclass, jlong, jlong, jint, jlong);
extern jboolean Java_org_opencv_ml_StatModel_train_11(JNIEnv*, jclass, jlong, jlong, jint);
extern jboolean Java_org_opencv_ml_StatModel_train_12(JNIEnv*, jclass, jlong, jlong);
extern jlong Java_org_opencv_ml_TrainData_create_10(JNIEnv*, jclass, jlong, jint, jlong, jlong, jlong, jlong, jlong);
extern jlong Java_org_opencv_ml_TrainData_create_11(JNIEnv*, jclass, jlong, jint, jlong);
extern void Java_org_opencv_ml_TrainData_delete(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_ml_TrainData_getCatCount_10(JNIEnv*, jclass, jlong, jint);
extern jlong Java_org_opencv_ml_TrainData_getCatMap_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_ml_TrainData_getCatOfs_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_ml_TrainData_getClassLabels_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_ml_TrainData_getDefaultSubstValues_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_ml_TrainData_getLayout_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_ml_TrainData_getMissing_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_ml_TrainData_getNAllVars_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_ml_TrainData_getNSamples_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_ml_TrainData_getNTestSamples_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_ml_TrainData_getNTrainSamples_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_ml_TrainData_getNVars_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_ml_TrainData_getNames_10(JNIEnv*, jclass, jlong, jlong);
extern jlong Java_org_opencv_ml_TrainData_getNormCatResponses_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_ml_TrainData_getResponseType_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_ml_TrainData_getResponses_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_ml_TrainData_getSampleWeights_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_ml_TrainData_getSample_10(JNIEnv*, jclass, jlong, jlong, jint, jfloat);
extern jlong Java_org_opencv_ml_TrainData_getSamples_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_ml_TrainData_getSubVector_10(JNIEnv*, jclass, jlong, jlong);
extern jlong Java_org_opencv_ml_TrainData_getTestNormCatResponses_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_ml_TrainData_getTestResponses_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_ml_TrainData_getTestSampleIdx_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_ml_TrainData_getTestSampleWeights_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_ml_TrainData_getTestSamples_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_ml_TrainData_getTrainNormCatResponses_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_ml_TrainData_getTrainResponses_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_ml_TrainData_getTrainSampleIdx_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_ml_TrainData_getTrainSampleWeights_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_ml_TrainData_getTrainSamples_10(JNIEnv*, jclass, jlong, jint, jboolean, jboolean);
extern jlong Java_org_opencv_ml_TrainData_getTrainSamples_11(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_ml_TrainData_getValues_10(JNIEnv*, jclass, jlong, jint, jlong, jfloat);
extern jlong Java_org_opencv_ml_TrainData_getVarIdx_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_ml_TrainData_getVarSymbolFlags_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_ml_TrainData_getVarType_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_ml_TrainData_setTrainTestSplitRatio_10(JNIEnv*, jclass, jlong, jdouble, jboolean);
extern void Java_org_opencv_ml_TrainData_setTrainTestSplitRatio_11(JNIEnv*, jclass, jlong, jdouble);
extern void Java_org_opencv_ml_TrainData_setTrainTestSplit_10(JNIEnv*, jclass, jlong, jint, jboolean);
extern void Java_org_opencv_ml_TrainData_setTrainTestSplit_11(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_ml_TrainData_shuffleTrainTest_10(JNIEnv*, jclass, jlong);

*/
import "C"

func ANN_MLPNative_create_0() int64 {
	return int64(C.Java_org_opencv_ml_ANN_1MLP_create_10(clzEnv,clzObj))
}
func ANN_MLPNative_delete(nativeObj int64) {
	C.Java_org_opencv_ml_ANN_1MLP_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}
func ANN_MLPNative_getBackpropMomentumScale_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_ml_ANN_1MLP_getBackpropMomentumScale_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}
func ANN_MLPNative_getBackpropWeightScale_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_ml_ANN_1MLP_getBackpropWeightScale_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ANN_MLPNative_getLayerSizes_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_ml_ANN_1MLP_getLayerSizes_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ANN_MLPNative_getRpropDW0_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_ml_ANN_1MLP_getRpropDW0_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ANN_MLPNative_getRpropDWMax_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_ml_ANN_1MLP_getRpropDWMax_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ANN_MLPNative_getRpropDWMin_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_ml_ANN_1MLP_getRpropDWMin_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ANN_MLPNative_getRpropDWMinus_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_ml_ANN_1MLP_getRpropDWMinus_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ANN_MLPNative_getRpropDWPlus_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_ml_ANN_1MLP_getRpropDWPlus_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ANN_MLPNative_getTermCriteria_0(nativeObj int64) []float64 {
	return togofloat64Array(C.Java_org_opencv_ml_ANN_1MLP_getTermCriteria_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ANN_MLPNative_getTrainMethod_0(nativeObj int64) int {
	return int(C.Java_org_opencv_ml_ANN_1MLP_getTrainMethod_10(
			clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ANN_MLPNative_getWeights_0(nativeObj int64,layerIdx int) int64 {
	return int64(C.Java_org_opencv_ml_ANN_1MLP_getWeights_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(layerIdx)))
}

func ANN_MLPNative_load_0(filepath string) int64 {
	defer ungointerface(filepath)
	return int64(C.Java_org_opencv_ml_ANN_1MLP_load_10(clzEnv,clzObj,tojstring(filepath)))
}

func ANN_MLPNative_setActivationFunction_0(nativeObj int64,rtype int,param1 float64,param2 float64) {
	C.Java_org_opencv_ml_ANN_1MLP_setActivationFunction_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(rtype),(C.jdouble)(param1),(C.jdouble)(param2))
}

func ANN_MLPNative_setActivationFunction_1(nativeObj int64,rtype int) {
	C.Java_org_opencv_ml_ANN_1MLP_setActivationFunction_11(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(rtype))
}

func ANN_MLPNative_setBackpropMomentumScale_0(nativeObj int64,val float64) {
	C.Java_org_opencv_ml_ANN_1MLP_setBackpropMomentumScale_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jdouble)(val))
}

func ANN_MLPNative_setBackpropWeightScale_0(nativeObj int64,val float64) {
	C.Java_org_opencv_ml_ANN_1MLP_setBackpropWeightScale_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jdouble)(val))
}

func ANN_MLPNative_setLayerSizes_0(nativeObj int64,_layer_sizes_nativeObj int64) {
	C.Java_org_opencv_ml_ANN_1MLP_setLayerSizes_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(_layer_sizes_nativeObj))
}

func ANN_MLPNative_setRpropDW0_0(nativeObj int64,val float64) {
	C.Java_org_opencv_ml_ANN_1MLP_setRpropDW0_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jdouble)(val))
}

func ANN_MLPNative_setRpropDWMax_0(nativeObj int64,val float64) {
	C.Java_org_opencv_ml_ANN_1MLP_setRpropDWMax_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jdouble)(val))
}

func ANN_MLPNative_setRpropDWMin_0(nativeObj int64,val float64) {
	C.Java_org_opencv_ml_ANN_1MLP_setRpropDWMin_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jdouble)(val))
}

func ANN_MLPNative_setRpropDWMinus_0(nativeObj int64,val float64) {
	C.Java_org_opencv_ml_ANN_1MLP_setRpropDWMinus_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jdouble)(val))
}

func ANN_MLPNative_setRpropDWPlus_0(nativeObj int64,val float64) {
	C.Java_org_opencv_ml_ANN_1MLP_setRpropDWPlus_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jdouble)(val))
}

func ANN_MLPNative_setTermCriteria_0(nativeObj int64,val_type int,val_maxCount int,val_epsilon float64) {
	C.Java_org_opencv_ml_ANN_1MLP_setTermCriteria_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(val_type),(C.jint)(val_maxCount),(C.jdouble)(val_epsilon))
}

func ANN_MLPNative_setTrainMethod_0(nativeObj int64,method int,param1 float64,param2 float64) {
	C.Java_org_opencv_ml_ANN_1MLP_setTrainMethod_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(method),(C.jdouble)(param1),(C.jdouble)(param2))
}

func ANN_MLPNative_setTrainMethod_1(nativeObj int64,method int) {
	C.Java_org_opencv_ml_ANN_1MLP_setTrainMethod_11(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(method))
}

func BoostNative_create_0() int64 {
	return int64(C.Java_org_opencv_ml_Boost_create_10(clzEnv,clzObj))
}

func BoostNative_delete(nativeObj int64) {
	C.Java_org_opencv_ml_Boost_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func BoostNative_getBoostType_0(nativeObj int64) int {
	return int(C.Java_org_opencv_ml_Boost_getBoostType_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func BoostNative_getWeakCount_0(nativeObj int64) int {
	return int(C.Java_org_opencv_ml_Boost_getWeakCount_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func BoostNative_getWeightTrimRate_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_ml_Boost_getWeightTrimRate_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func BoostNative_load_0(filepath string,nodeName string) int64 {
	defer ungointerface(filepath)
	defer ungointerface(nodeName)
	return int64(C.Java_org_opencv_ml_Boost_load_10(clzEnv,clzObj,tojstring(filepath),tojstring(nodeName)))
}

func BoostNative_load_1(filepath string) int64 {
	defer ungointerface(filepath)
	return int64(C.Java_org_opencv_ml_Boost_load_11(clzEnv,clzObj,tojstring(filepath)))
}

func BoostNative_setBoostType_0(nativeObj int64,val int) {
	C.Java_org_opencv_ml_Boost_setBoostType_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(val))
}

func BoostNative_setWeakCount_0(nativeObj int64,val int) {
	C.Java_org_opencv_ml_Boost_setWeakCount_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(val))
}

func BoostNative_setWeightTrimRate_0(nativeObj int64,val float64) {
	C.Java_org_opencv_ml_Boost_setWeightTrimRate_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jdouble)(val))
}

func DTreesNative_create_0() int64 {
	return int64(C.Java_org_opencv_ml_DTrees_create_10(clzEnv,clzObj))
}

func DTreesNative_delete(nativeObj int64) {
	C.Java_org_opencv_ml_DTrees_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func DTreesNative_getCVFolds_0(nativeObj int64) int {
	return int(C.Java_org_opencv_ml_DTrees_getCVFolds_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func DTreesNative_getMaxCategories_0(nativeObj int64) int {
	return int(C.Java_org_opencv_ml_DTrees_getMaxCategories_10(
			clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func DTreesNative_getMaxDepth_0(nativeObj int64) int {
	return int(C.Java_org_opencv_ml_DTrees_getMaxDepth_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func DTreesNative_getMinSampleCount_0(nativeObj int64) int {
	return int(C.Java_org_opencv_ml_DTrees_getMinSampleCount_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func DTreesNative_getPriors_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_ml_DTrees_getPriors_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func DTreesNative_getRegressionAccuracy_0(nativeObj int64) float32 {
	return float32(C.Java_org_opencv_ml_DTrees_getRegressionAccuracy_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func DTreesNative_getTruncatePrunedTree_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_ml_DTrees_getTruncatePrunedTree_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func DTreesNative_getUse1SERule_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_ml_DTrees_getUse1SERule_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func DTreesNative_getUseSurrogates_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_ml_DTrees_getUseSurrogates_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func DTreesNative_load_0(filepath string,nodeName string) int64 {
	defer ungointerface(filepath)
	defer ungointerface(nodeName)
	return int64(C.Java_org_opencv_ml_DTrees_load_10(clzEnv,clzObj,tojstring(filepath),tojstring(nodeName)))
}

func DTreesNative_load_1(filepath string) int64 {
	defer ungointerface(filepath)
	return int64(C.Java_org_opencv_ml_DTrees_load_11(clzEnv,clzObj,tojstring(filepath)))
}

func DTreesNative_setCVFolds_0(nativeObj int64,val int) {
	C.Java_org_opencv_ml_DTrees_setCVFolds_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(val))
}

func DTreesNative_setMaxCategories_0(nativeObj int64,val int) {
	C.Java_org_opencv_ml_DTrees_setMaxCategories_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(val))
}

func DTreesNative_setMaxDepth_0(nativeObj int64,val int) {
	C.Java_org_opencv_ml_DTrees_setMaxDepth_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(val))
}

func DTreesNative_setMinSampleCount_0(nativeObj int64,val int) {
	C.Java_org_opencv_ml_DTrees_setMinSampleCount_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(val))
}

func DTreesNative_setPriors_0(nativeObj int64,val_nativeObj int64) {
	C.Java_org_opencv_ml_DTrees_setPriors_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(val_nativeObj))
}

func DTreesNative_setRegressionAccuracy_0(nativeObj int64,val float32) {
	C.Java_org_opencv_ml_DTrees_setRegressionAccuracy_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jfloat)(val))
}

func DTreesNative_setTruncatePrunedTree_0(nativeObj int64,val bool) {
	C.Java_org_opencv_ml_DTrees_setTruncatePrunedTree_10(clzEnv,clzObj,(C.jlong)(nativeObj),tojboolean(val))
}

func DTreesNative_setUse1SERule_0(nativeObj int64,val bool) {
	C.Java_org_opencv_ml_DTrees_setUse1SERule_10(clzEnv,clzObj,(C.jlong)(nativeObj),tojboolean(val))
}

func DTreesNative_setUseSurrogates_0(nativeObj int64,val bool) {
	C.Java_org_opencv_ml_DTrees_setUseSurrogates_10(clzEnv,clzObj,(C.jlong)(nativeObj),tojboolean(val))
}

func EMNative_create_0() int64 {
	return int64(C.Java_org_opencv_ml_EM_create_10(clzEnv,clzObj))
}

func EMNative_delete(nativeObj int64) {
	C.Java_org_opencv_ml_EM_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func EMNative_getClustersNumber_0(nativeObj int64) int {
	return int(C.Java_org_opencv_ml_EM_getClustersNumber_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func EMNative_getCovarianceMatrixType_0(nativeObj int64) int {
	return int(C.Java_org_opencv_ml_EM_getCovarianceMatrixType_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func EMNative_getCovs_0(nativeObj int64,covs_mat_nativeObj int64) {
	C.Java_org_opencv_ml_EM_getCovs_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(covs_mat_nativeObj))
}

func EMNative_getMeans_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_ml_EM_getMeans_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func EMNative_getTermCriteria_0(nativeObj int64) []float64 {
	return togofloat64Array(C.Java_org_opencv_ml_EM_getTermCriteria_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func EMNative_getWeights_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_ml_EM_getWeights_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func EMNative_load_0(filepath string,nodeName string) int64 {
	defer ungointerface(filepath)
	defer ungointerface(nodeName)
	return int64(C.Java_org_opencv_ml_EM_load_10(clzEnv,clzObj,tojstring(filepath),tojstring(nodeName)))
}

func EMNative_load_1(filepath string) int64 {
	defer ungointerface(filepath)
	return int64(C.Java_org_opencv_ml_EM_load_11(clzEnv,clzObj,tojstring(filepath)))
}

func EMNative_predict2_0(nativeObj int64,sample_nativeObj int64,probs_nativeObj int64) []float64 {
	return togofloat64Array(C.Java_org_opencv_ml_EM_predict2_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(sample_nativeObj),(C.jlong)(probs_nativeObj)))
}

func EMNative_predict_0(nativeObj int64,samples_nativeObj int64,results_nativeObj int64,flags int) float32 {
	return float32(C.Java_org_opencv_ml_EM_predict_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(samples_nativeObj),(C.jlong)(results_nativeObj),(C.jint)(flags)))
}

func EMNative_predict_1(nativeObj int64,samples_nativeObj int64) float32 {
	return float32(C.Java_org_opencv_ml_EM_predict_11(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(samples_nativeObj)))
}

func EMNative_setClustersNumber_0(nativeObj int64,val int) {
	C.Java_org_opencv_ml_EM_setClustersNumber_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(val))
}

func EMNative_setCovarianceMatrixType_0(nativeObj int64,val int) {
	C.Java_org_opencv_ml_EM_setCovarianceMatrixType_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(val))
}

func EMNative_setTermCriteria_0(nativeObj int64,val_type int,val_maxCount int,val_epsilon float64) {
	C.Java_org_opencv_ml_EM_setTermCriteria_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(val_type),(C.jint)(val_maxCount),(C.jdouble)(val_epsilon))
}

func EMNative_trainEM_0(nativeObj int64,samples_nativeObj int64,logLikelihoods_nativeObj int64,labels_nativeObj int64,probs_nativeObj int64) bool {
	return togobool(C.Java_org_opencv_ml_EM_trainEM_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(samples_nativeObj),(C.jlong)(logLikelihoods_nativeObj),(C.jlong)(labels_nativeObj),(C.jlong)(probs_nativeObj)))
}

func EMNative_trainEM_1(nativeObj int64,samples_nativeObj int64) bool {
	return togobool(C.Java_org_opencv_ml_EM_trainEM_11(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(samples_nativeObj)))
}

func EMNative_trainE_0(nativeObj int64,samples_nativeObj int64,means0_nativeObj int64,covs0_nativeObj int64,weights0_nativeObj int64,logLikelihoods_nativeObj int64,labels_nativeObj int64,probs_nativeObj int64) bool {
	return togobool(C.Java_org_opencv_ml_EM_trainE_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(samples_nativeObj),(C.jlong)(means0_nativeObj),(C.jlong)(covs0_nativeObj),(C.jlong)(weights0_nativeObj),(C.jlong)(logLikelihoods_nativeObj),(C.jlong)(labels_nativeObj),(C.jlong)(probs_nativeObj)))
}

func EMNative_trainE_1(nativeObj int64,samples_nativeObj int64,means0_nativeObj int64) bool {
	return togobool(C.Java_org_opencv_ml_EM_trainE_11(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(samples_nativeObj),(C.jlong)(means0_nativeObj)))
}

func EMNative_trainM_0(nativeObj int64,samples_nativeObj int64,probs0_nativeObj int64,logLikelihoods_nativeObj int64,labels_nativeObj int64,probs_nativeObj int64) bool {
	return togobool(C.Java_org_opencv_ml_EM_trainM_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(samples_nativeObj),(C.jlong)(probs0_nativeObj),(C.jlong)(logLikelihoods_nativeObj),(C.jlong)(labels_nativeObj),(C.jlong)(probs_nativeObj)))
}

func EMNative_trainM_1(nativeObj int64,samples_nativeObj int64,probs0_nativeObj int64) bool {
	return togobool(C.Java_org_opencv_ml_EM_trainM_11(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(samples_nativeObj),(C.jlong)(probs0_nativeObj)))
}

func KNearestNative_create_0() int64 {
	return int64(C.Java_org_opencv_ml_KNearest_create_10(clzEnv,clzObj))
}

func KNearestNative_delete(nativeObj int64) {
	C.Java_org_opencv_ml_KNearest_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func KNearestNative_findNearest_0(nativeObj int64,samples_nativeObj int64,k int,results_nativeObj int64,neighborResponses_nativeObj int64,dist_nativeObj int64) float32 {
	return float32(C.Java_org_opencv_ml_KNearest_findNearest_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(samples_nativeObj),(C.jint)(k),(C.jlong)(results_nativeObj),(C.jlong)(
				neighborResponses_nativeObj),(C.jlong)(dist_nativeObj)))
}

func KNearestNative_findNearest_1(nativeObj int64,samples_nativeObj int64,k int,results_nativeObj int64) float32 {
	return float32(C.Java_org_opencv_ml_KNearest_findNearest_11(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(samples_nativeObj),(C.jint)(k),(C.jlong)(results_nativeObj)))
}

func KNearestNative_getAlgorithmType_0(nativeObj int64) int {
	return int(C.Java_org_opencv_ml_KNearest_getAlgorithmType_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func KNearestNative_getDefaultK_0(nativeObj int64) int {
	return int(C.Java_org_opencv_ml_KNearest_getDefaultK_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func KNearestNative_getEmax_0(nativeObj int64) int {
	return int(C.Java_org_opencv_ml_KNearest_getEmax_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func KNearestNative_getIsClassifier_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_ml_KNearest_getIsClassifier_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func KNearestNative_setAlgorithmType_0(nativeObj int64,val int) {
	C.Java_org_opencv_ml_KNearest_setAlgorithmType_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(val))
}

func KNearestNative_setDefaultK_0(nativeObj int64,val int) {
	C.Java_org_opencv_ml_KNearest_setDefaultK_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(val))
}

func KNearestNative_setEmax_0(nativeObj int64,val int) {
	C.Java_org_opencv_ml_KNearest_setEmax_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(val))
}

func KNearestNative_setIsClassifier_0(nativeObj int64,val bool) {
	C.Java_org_opencv_ml_KNearest_setIsClassifier_10(clzEnv,clzObj,(C.jlong)(nativeObj),tojboolean(val))
}

func LogisticRegressionNative_create_0() int64 {
	return int64(C.Java_org_opencv_ml_LogisticRegression_create_10(clzEnv,clzObj))
}

func LogisticRegressionNative_delete(nativeObj int64) {
	C.Java_org_opencv_ml_LogisticRegression_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func LogisticRegressionNative_getIterations_0(nativeObj int64) int {
	return int(C.Java_org_opencv_ml_LogisticRegression_getIterations_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func LogisticRegressionNative_getLearningRate_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_ml_LogisticRegression_getLearningRate_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func LogisticRegressionNative_getMiniBatchSize_0(nativeObj int64) int {
	return int(C.Java_org_opencv_ml_LogisticRegression_getMiniBatchSize_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func LogisticRegressionNative_getRegularization_0(nativeObj int64) int {
	return int(C.Java_org_opencv_ml_LogisticRegression_getRegularization_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func LogisticRegressionNative_getTermCriteria_0(nativeObj int64) []float64 {
	return togofloat64Array(C.Java_org_opencv_ml_LogisticRegression_getTermCriteria_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func LogisticRegressionNative_getTrainMethod_0(nativeObj int64) int {
	return int(C.Java_org_opencv_ml_LogisticRegression_getTrainMethod_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func LogisticRegressionNative_get_learnt_thetas_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_ml_LogisticRegression_get_1learnt_1thetas_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func LogisticRegressionNative_load_0(filepath string,nodeName string) int64 {
	defer ungointerface(filepath)
	defer ungointerface(nodeName)
	return int64(C.Java_org_opencv_ml_LogisticRegression_load_10(clzEnv,clzObj,tojstring(filepath),tojstring(nodeName)))
}

func LogisticRegressionNative_load_1(filepath string) int64 {
	defer ungointerface(filepath)
	return int64(C.Java_org_opencv_ml_LogisticRegression_load_11(clzEnv,clzObj,tojstring(filepath)))
}

func LogisticRegressionNative_predict_0(nativeObj int64,samples_nativeObj int64,results_nativeObj int64,flags int) float32 {
	return float32(C.Java_org_opencv_ml_LogisticRegression_predict_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(samples_nativeObj),(C.jlong)(results_nativeObj),(C.jint)(flags)))
}

func LogisticRegressionNative_predict_1(nativeObj int64,samples_nativeObj int64) float32 {
	return float32(C.Java_org_opencv_ml_LogisticRegression_predict_11(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(samples_nativeObj)))
}

func LogisticRegressionNative_setIterations_0(nativeObj int64,val int) {
	C.Java_org_opencv_ml_LogisticRegression_setIterations_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(val))
}

func LogisticRegressionNative_setLearningRate_0(nativeObj int64,val float64) {
	C.Java_org_opencv_ml_LogisticRegression_setLearningRate_10(
		clzEnv,clzObj,(C.jlong)(nativeObj),(C.jdouble)(val))
}

func LogisticRegressionNative_setMiniBatchSize_0(nativeObj int64,val int) {
	C.Java_org_opencv_ml_LogisticRegression_setMiniBatchSize_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(val))
}

func LogisticRegressionNative_setRegularization_0(nativeObj int64,val int) {
	C.Java_org_opencv_ml_LogisticRegression_setRegularization_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(val))
}

func LogisticRegressionNative_setTermCriteria_0(nativeObj int64,val_type int,val_maxCount int,val_epsilon float64) {
	C.Java_org_opencv_ml_LogisticRegression_setTermCriteria_10(
		clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(val_type),(C.jint)(val_maxCount),(C.jdouble)(val_epsilon))
}

func LogisticRegressionNative_setTrainMethod_0(nativeObj int64,val int) {
	C.Java_org_opencv_ml_LogisticRegression_setTrainMethod_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(val))
}

func NormalBayesClassifierNative_create_0() int64 {
	return int64(C.Java_org_opencv_ml_NormalBayesClassifier_create_10(clzEnv,clzObj))
}

func NormalBayesClassifierNative_delete(nativeObj int64) {
	C.Java_org_opencv_ml_NormalBayesClassifier_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func NormalBayesClassifierNative_load_0(filepath string,nodeName string) int64 {
	defer ungointerface(filepath)
	defer ungointerface(nodeName)
	return int64(C.Java_org_opencv_ml_NormalBayesClassifier_load_10(clzEnv,clzObj,tojstring(filepath),tojstring(nodeName)))
}

func NormalBayesClassifierNative_load_1(filepath string) int64 {
	defer ungointerface(filepath)
	return int64(C.Java_org_opencv_ml_NormalBayesClassifier_load_11(clzEnv,clzObj,tojstring(filepath)))
}

func NormalBayesClassifierNative_predictProb_0(nativeObj int64,inputs_nativeObj int64,outputs_nativeObj int64,outputProbs_nativeObj int64,flags int) float32 {
	return float32(C.Java_org_opencv_ml_NormalBayesClassifier_predictProb_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(inputs_nativeObj),(C.jlong)(outputs_nativeObj),(C.jlong)(outputProbs_nativeObj),(C.jint)(flags)))
}

func NormalBayesClassifierNative_predictProb_1(nativeObj int64,inputs_nativeObj int64,outputs_nativeObj int64,outputProbs_nativeObj int64) float32 {
	return float32(C.Java_org_opencv_ml_NormalBayesClassifier_predictProb_11(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(inputs_nativeObj),(C.jlong)(outputs_nativeObj),(C.jlong)(outputProbs_nativeObj)))
}

func ParamGridNative_create_0(minVal float64,maxVal float64,logstep float64) int64 {
	return int64(C.Java_org_opencv_ml_ParamGrid_create_10(clzEnv,clzObj,(C.jdouble)(minVal),(C.jdouble)(maxVal),(C.jdouble)(logstep)))
}

func ParamGridNative_create_1() int64 {
	return int64(C.Java_org_opencv_ml_ParamGrid_create_11(clzEnv,clzObj))
}

func ParamGridNative_delete(nativeObj int64) {
	C.Java_org_opencv_ml_ParamGrid_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func ParamGridNative_get_logStep_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_ml_ParamGrid_get_1logStep_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ParamGridNative_get_maxVal_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_ml_ParamGrid_get_1maxVal_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ParamGridNative_get_minVal_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_ml_ParamGrid_get_1minVal_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func ParamGridNative_set_logStep_0(nativeObj int64,logStep float64) {
	C.Java_org_opencv_ml_ParamGrid_set_1logStep_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jdouble)(logStep))
}

func ParamGridNative_set_maxVal_0(nativeObj int64,maxVal float64) {
	C.Java_org_opencv_ml_ParamGrid_set_1maxVal_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jdouble)(maxVal))
}

func ParamGridNative_set_minVal_0(nativeObj int64,minVal float64) {
	C.Java_org_opencv_ml_ParamGrid_set_1minVal_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jdouble)(minVal))
}

func RTreesNative_create_0() int64 {
	return int64(C.Java_org_opencv_ml_RTrees_create_10(clzEnv,clzObj))
}

func RTreesNative_delete(nativeObj int64) {
	C.Java_org_opencv_ml_RTrees_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func RTreesNative_getActiveVarCount_0(nativeObj int64) int {
	return int(C.Java_org_opencv_ml_RTrees_getActiveVarCount_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func RTreesNative_getCalculateVarImportance_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_ml_RTrees_getCalculateVarImportance_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func RTreesNative_getTermCriteria_0(nativeObj int64) []float64 {
	return togofloat64Array(C.Java_org_opencv_ml_RTrees_getTermCriteria_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func RTreesNative_getVarImportance_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_ml_RTrees_getVarImportance_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func RTreesNative_getVotes_0(nativeObj int64,samples_nativeObj int64,results_nativeObj int64,flags int) {
	C.Java_org_opencv_ml_RTrees_getVotes_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(samples_nativeObj),(C.jlong)(results_nativeObj),(C.jint)(flags))
}

func RTreesNative_load_0(filepath string,nodeName string) int64 {
	defer ungointerface(filepath)
	defer ungointerface(nodeName)
	return int64(C.Java_org_opencv_ml_RTrees_load_10(clzEnv,clzObj,tojstring(filepath),tojstring(nodeName)))
}

func RTreesNative_load_1(filepath string) int64 {
	defer ungointerface(filepath)
	return int64(C.Java_org_opencv_ml_RTrees_load_11(clzEnv,clzObj,tojstring(filepath)))
}

func RTreesNative_setActiveVarCount_0(nativeObj int64,val int) {
	C.Java_org_opencv_ml_RTrees_setActiveVarCount_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(val))
}

func RTreesNative_setCalculateVarImportance_0(nativeObj int64,val bool) {
	C.Java_org_opencv_ml_RTrees_setCalculateVarImportance_10(clzEnv,clzObj,(C.jlong)(nativeObj),tojboolean(val))
}

func RTreesNative_setTermCriteria_0(nativeObj int64,val_type int,val_maxCount int,val_epsilon float64) {
	C.Java_org_opencv_ml_RTrees_setTermCriteria_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(val_type),(C.jint)(val_maxCount),(C.jdouble)(val_epsilon))
}

func SVMNative_create_0() int64 {
	return int64(C.Java_org_opencv_ml_SVM_create_10(clzEnv,clzObj))
}

func SVMNative_delete(nativeObj int64) {
	C.Java_org_opencv_ml_SVM_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func SVMNative_getC_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_ml_SVM_getC_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func SVMNative_getClassWeights_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_ml_SVM_getClassWeights_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func SVMNative_getCoef0_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_ml_SVM_getCoef0_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func SVMNative_getDecisionFunction_0(nativeObj int64,i int,alpha_nativeObj int64,svidx_nativeObj int64) float64 {
	return float64(C.Java_org_opencv_ml_SVM_getDecisionFunction_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(i),(C.jlong)(alpha_nativeObj),(C.jlong)(svidx_nativeObj)))
}

func SVMNative_getDefaultGridPtr_0(param_id int) int64 {
	return int64(C.Java_org_opencv_ml_SVM_getDefaultGridPtr_10(
			clzEnv,clzObj,(C.jint)(param_id)))
}

func SVMNative_getDegree_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_ml_SVM_getDegree_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func SVMNative_getGamma_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_ml_SVM_getGamma_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func SVMNative_getKernelType_0(nativeObj int64) int {
	return int(C.Java_org_opencv_ml_SVM_getKernelType_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func SVMNative_getNu_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_ml_SVM_getNu_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func SVMNative_getP_0(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_ml_SVM_getP_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func SVMNative_getSupportVectors_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_ml_SVM_getSupportVectors_10(
			clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func SVMNative_getTermCriteria_0(nativeObj int64) []float64 {
	return togofloat64Array(C.Java_org_opencv_ml_SVM_getTermCriteria_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func SVMNative_getType_0(nativeObj int64) int {
	return int(C.Java_org_opencv_ml_SVM_getType_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func SVMNative_getUncompressedSupportVectors_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_ml_SVM_getUncompressedSupportVectors_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func SVMNative_load_0(filepath string) int64 {
	defer ungointerface(filepath)
	return int64(C.Java_org_opencv_ml_SVM_load_10(clzEnv,clzObj,tojstring(filepath)))
}

func SVMNative_setC_0(nativeObj int64,val float64) {
	C.Java_org_opencv_ml_SVM_setC_10(clzEnv,clzObj,(C.jlong)(
		nativeObj),(C.jdouble)(val))
}

func SVMNative_setClassWeights_0(nativeObj int64,val_nativeObj int64) {
	C.Java_org_opencv_ml_SVM_setClassWeights_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(val_nativeObj))
}

func SVMNative_setCoef0_0(nativeObj int64,val float64) {
	C.Java_org_opencv_ml_SVM_setCoef0_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jdouble)(val))
}

func SVMNative_setDegree_0(nativeObj int64,val float64) {
	C.Java_org_opencv_ml_SVM_setDegree_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jdouble)(val))
}

func SVMNative_setGamma_0(nativeObj int64,val float64) {
	C.Java_org_opencv_ml_SVM_setGamma_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jdouble)(val))
}

func SVMNative_setKernel_0(nativeObj int64,kernelType int) {
	C.Java_org_opencv_ml_SVM_setKernel_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(kernelType))
}

func SVMNative_setNu_0(nativeObj int64,val float64) {
	C.Java_org_opencv_ml_SVM_setNu_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jdouble)(val))
}

func SVMNative_setP_0(nativeObj int64,val float64) {
	C.Java_org_opencv_ml_SVM_setP_10(clzEnv,clzObj,(C.jlong)(
		nativeObj),(C.jdouble)(val))
}

func SVMNative_setTermCriteria_0(nativeObj int64,val_type int,val_maxCount int,val_epsilon float64) {
	C.Java_org_opencv_ml_SVM_setTermCriteria_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(val_type),(C.jint)(val_maxCount),(C.jdouble)(val_epsilon))
}

func SVMNative_setType_0(nativeObj int64,val int) {
	C.Java_org_opencv_ml_SVM_setType_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(val))
}

func SVMNative_trainAuto_0(nativeObj int64,samples_nativeObj int64,layout int,responses_nativeObj int64,kFold int,Cgrid_nativeObj int64,gammaGrid_nativeObj int64,pGrid_nativeObj int64,nuGrid_nativeObj int64,coeffGrid_nativeObj int64,degreeGrid_nativeObj int64,balanced bool) bool {
	return togobool(C.Java_org_opencv_ml_SVM_trainAuto_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(samples_nativeObj),(C.jint)(layout),(C.jlong)(responses_nativeObj),(C.jint)(
					kFold),(C.jlong)(Cgrid_nativeObj),(C.jlong)(gammaGrid_nativeObj),(C.jlong)(pGrid_nativeObj),(C.jlong)(nuGrid_nativeObj),(C.jlong)(coeffGrid_nativeObj),(C.jlong)(degreeGrid_nativeObj),tojboolean(balanced)))
}

func SVMNative_trainAuto_1(nativeObj int64,samples_nativeObj int64,layout int,responses_nativeObj int64) bool {
	return togobool(C.Java_org_opencv_ml_SVM_trainAuto_11(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(samples_nativeObj),(C.jint)(layout),(C.jlong)(responses_nativeObj)))
}

func SVMSGDNative_create_0() int64 {
	return int64(C.Java_org_opencv_ml_SVMSGD_create_10(clzEnv,clzObj))
}

func SVMSGDNative_delete(nativeObj int64) {
	C.Java_org_opencv_ml_SVMSGD_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func SVMSGDNative_getInitialStepSize_0(nativeObj int64) float32 {
	return float32(C.Java_org_opencv_ml_SVMSGD_getInitialStepSize_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func SVMSGDNative_getMarginRegularization_0(nativeObj int64) float32 {
	return float32(C.Java_org_opencv_ml_SVMSGD_getMarginRegularization_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func SVMSGDNative_getMarginType_0(nativeObj int64) int {
	return int(C.Java_org_opencv_ml_SVMSGD_getMarginType_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func SVMSGDNative_getShift_0(nativeObj int64) float32 {
	return float32(C.Java_org_opencv_ml_SVMSGD_getShift_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func SVMSGDNative_getStepDecreasingPower_0(nativeObj int64) float32 {
	return float32(C.Java_org_opencv_ml_SVMSGD_getStepDecreasingPower_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func SVMSGDNative_getSvmsgdType_0(nativeObj int64) int {
	return int(C.Java_org_opencv_ml_SVMSGD_getSvmsgdType_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func SVMSGDNative_getTermCriteria_0(nativeObj int64) []float64 {
	return togofloat64Array(C.Java_org_opencv_ml_SVMSGD_getTermCriteria_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func SVMSGDNative_getWeights_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_ml_SVMSGD_getWeights_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func SVMSGDNative_load_0(filepath string,nodeName string) int64 {
	defer ungointerface(filepath)
	defer ungointerface(nodeName)
	return int64(C.Java_org_opencv_ml_SVMSGD_load_10(clzEnv,clzObj,tojstring(filepath),tojstring(nodeName)))
}

func SVMSGDNative_load_1(filepath string) int64 {
	defer ungointerface(filepath)
	return int64(C.Java_org_opencv_ml_SVMSGD_load_11(clzEnv,clzObj,tojstring(filepath)))
}

func SVMSGDNative_setInitialStepSize_0(nativeObj int64,InitialStepSize float32) {
	C.Java_org_opencv_ml_SVMSGD_setInitialStepSize_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jfloat)(InitialStepSize))
}

func SVMSGDNative_setMarginRegularization_0(nativeObj int64,marginRegularization float32) {
	C.Java_org_opencv_ml_SVMSGD_setMarginRegularization_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jfloat)(marginRegularization))
}

func SVMSGDNative_setMarginType_0(nativeObj int64,marginType int) {
	C.Java_org_opencv_ml_SVMSGD_setMarginType_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(marginType))
}

func SVMSGDNative_setOptimalParameters_0(nativeObj int64,svmsgdType int,marginType int) {
	C.Java_org_opencv_ml_SVMSGD_setOptimalParameters_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(svmsgdType),(C.jint)(marginType))
}

func SVMSGDNative_setOptimalParameters_1(nativeObj int64) {
	C.Java_org_opencv_ml_SVMSGD_setOptimalParameters_11(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func SVMSGDNative_setStepDecreasingPower_0(nativeObj int64,stepDecreasingPower float32) {
	C.Java_org_opencv_ml_SVMSGD_setStepDecreasingPower_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jfloat)(stepDecreasingPower))
}

func SVMSGDNative_setSvmsgdType_0(nativeObj int64,svmsgdType int) {
	C.Java_org_opencv_ml_SVMSGD_setSvmsgdType_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(svmsgdType))
}

func SVMSGDNative_setTermCriteria_0(nativeObj int64,val_type int,val_maxCount int,val_epsilon float64) {
	C.Java_org_opencv_ml_SVMSGD_setTermCriteria_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(val_type),(C.jint)(val_maxCount),(C.jdouble)(val_epsilon))
}

func StatModelNative_calcError_0(nativeObj int64,data_nativeObj int64,test bool,resp_nativeObj int64) float32 {
	return float32(C.Java_org_opencv_ml_StatModel_calcError_10(
				clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(data_nativeObj),tojboolean(test),(C.jlong)(resp_nativeObj)))
}

func StatModelNative_delete(nativeObj int64) {
	C.Java_org_opencv_ml_StatModel_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func StatModelNative_empty_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_ml_StatModel_empty_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func StatModelNative_getVarCount_0(nativeObj int64) int {
	return int(C.Java_org_opencv_ml_StatModel_getVarCount_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func StatModelNative_isClassifier_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_ml_StatModel_isClassifier_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func StatModelNative_isTrained_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_ml_StatModel_isTrained_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func StatModelNative_predict_0(nativeObj int64,samples_nativeObj int64,results_nativeObj int64,flags int) float32 {
	return float32(C.Java_org_opencv_ml_StatModel_predict_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(
					samples_nativeObj),(C.jlong)(results_nativeObj),(C.jint)(flags)))
}

func StatModelNative_predict_1(nativeObj int64,samples_nativeObj int64) float32 {
	return float32(C.Java_org_opencv_ml_StatModel_predict_11(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(
					samples_nativeObj)))
}

func StatModelNative_train_0(nativeObj int64,samples_nativeObj int64,layout int,responses_nativeObj int64) bool {
	return togobool(C.Java_org_opencv_ml_StatModel_train_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(samples_nativeObj),(C.jint)(layout),(C.jlong)(responses_nativeObj)))
}

func StatModelNative_train_1(nativeObj int64,trainData_nativeObj int64,flags int) bool {
	return togobool(C.Java_org_opencv_ml_StatModel_train_11(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(trainData_nativeObj),(C.jint)(flags)))
}

func StatModelNative_train_2(nativeObj int64,trainData_nativeObj int64) bool {
	return togobool(C.Java_org_opencv_ml_StatModel_train_12(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(trainData_nativeObj)))
}

func TrainDataNative_create_0(samples_nativeObj int64,layout int,responses_nativeObj int64,varIdx_nativeObj int64,sampleIdx_nativeObj int64,sampleWeights_nativeObj int64,varType_nativeObj int64) int64 {
	return int64(C.Java_org_opencv_ml_TrainData_create_10(clzEnv,clzObj,(C.jlong)(samples_nativeObj),(C.jint)(layout),(C.jlong)(responses_nativeObj),(C.jlong)(varIdx_nativeObj),(C.jlong)(sampleIdx_nativeObj),(C.jlong)(sampleWeights_nativeObj),(C.jlong)(varType_nativeObj)))
}

func TrainDataNative_create_1(samples_nativeObj int64,layout int,responses_nativeObj int64) int64 {
	return int64(C.Java_org_opencv_ml_TrainData_create_11(clzEnv,clzObj,(C.jlong)(samples_nativeObj),(C.jint)(layout),(C.jlong)(responses_nativeObj)))
}

func TrainDataNative_delete(nativeObj int64) {
	C.Java_org_opencv_ml_TrainData_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func TrainDataNative_getCatCount_0(nativeObj int64,vi int) int {
	return int(C.Java_org_opencv_ml_TrainData_getCatCount_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(vi)))
}

func TrainDataNative_getCatMap_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_ml_TrainData_getCatMap_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TrainDataNative_getCatOfs_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_ml_TrainData_getCatOfs_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TrainDataNative_getClassLabels_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_ml_TrainData_getClassLabels_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TrainDataNative_getDefaultSubstValues_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_ml_TrainData_getDefaultSubstValues_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TrainDataNative_getLayout_0(nativeObj int64) int {
	return int(C.Java_org_opencv_ml_TrainData_getLayout_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TrainDataNative_getMissing_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_ml_TrainData_getMissing_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TrainDataNative_getNAllVars_0(nativeObj int64) int {
	return int(C.Java_org_opencv_ml_TrainData_getNAllVars_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TrainDataNative_getNSamples_0(nativeObj int64) int {
	return int(C.Java_org_opencv_ml_TrainData_getNSamples_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TrainDataNative_getNTestSamples_0(nativeObj int64) int {
	return int(C.Java_org_opencv_ml_TrainData_getNTestSamples_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TrainDataNative_getNTrainSamples_0(nativeObj int64) int {
	return int(C.Java_org_opencv_ml_TrainData_getNTrainSamples_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TrainDataNative_getNVars_0(nativeObj int64) int {
	return int(C.Java_org_opencv_ml_TrainData_getNVars_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TrainDataNative_getNames_0(nativeObj int64,names int64) {
	C.Java_org_opencv_ml_TrainData_getNames_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(names))
}

func TrainDataNative_getNormCatResponses_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_ml_TrainData_getNormCatResponses_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TrainDataNative_getResponseType_0(nativeObj int64) int {
	return int(C.Java_org_opencv_ml_TrainData_getResponseType_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TrainDataNative_getResponses_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_ml_TrainData_getResponses_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TrainDataNative_getSampleWeights_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_ml_TrainData_getSampleWeights_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TrainDataNative_getSample_0(nativeObj int64,varIdx_nativeObj int64,sidx int,buf float32) {
	C.Java_org_opencv_ml_TrainData_getSample_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(varIdx_nativeObj),(C.jint)(sidx),(C.jfloat)(buf))
}

func TrainDataNative_getSamples_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_ml_TrainData_getSamples_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TrainDataNative_getSubVector_0(vec_nativeObj int64,idx_nativeObj int64) int64 {
	return int64(C.Java_org_opencv_ml_TrainData_getSubVector_10(clzEnv,clzObj,(C.jlong)(vec_nativeObj),(C.jlong)(idx_nativeObj)))
}

func TrainDataNative_getTestNormCatResponses_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_ml_TrainData_getTestNormCatResponses_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TrainDataNative_getTestResponses_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_ml_TrainData_getTestResponses_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TrainDataNative_getTestSampleIdx_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_ml_TrainData_getTestSampleIdx_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TrainDataNative_getTestSampleWeights_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_ml_TrainData_getTestSampleWeights_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TrainDataNative_getTestSamples_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_ml_TrainData_getTestSamples_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TrainDataNative_getTrainNormCatResponses_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_ml_TrainData_getTrainNormCatResponses_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TrainDataNative_getTrainResponses_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_ml_TrainData_getTrainResponses_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TrainDataNative_getTrainSampleIdx_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_ml_TrainData_getTrainSampleIdx_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TrainDataNative_getTrainSampleWeights_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_ml_TrainData_getTrainSampleWeights_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TrainDataNative_getTrainSamples_0(nativeObj int64,layout int,compressSamples bool,compressVars bool) int64 {
	return int64(C.Java_org_opencv_ml_TrainData_getTrainSamples_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(layout),tojboolean(compressSamples),tojboolean(compressVars)))
}

func TrainDataNative_getTrainSamples_1(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_ml_TrainData_getTrainSamples_11(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TrainDataNative_getValues_0(nativeObj int64,vi int,sidx_nativeObj int64,values float32) {
	C.Java_org_opencv_ml_TrainData_getValues_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(vi),(C.jlong)(sidx_nativeObj),(C.jfloat)(values))
}

func TrainDataNative_getVarIdx_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_ml_TrainData_getVarIdx_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TrainDataNative_getVarSymbolFlags_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_ml_TrainData_getVarSymbolFlags_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TrainDataNative_getVarType_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_ml_TrainData_getVarType_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func TrainDataNative_setTrainTestSplitRatio_0(nativeObj int64,ratio float64,shuffle bool) {
	C.Java_org_opencv_ml_TrainData_setTrainTestSplitRatio_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jdouble)(ratio),tojboolean(shuffle))
}

func TrainDataNative_setTrainTestSplitRatio_1(nativeObj int64,ratio float64) {
	C.Java_org_opencv_ml_TrainData_setTrainTestSplitRatio_11(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jdouble)(ratio))
}

func TrainDataNative_setTrainTestSplit_0(nativeObj int64,count int,shuffle bool) {
	C.Java_org_opencv_ml_TrainData_setTrainTestSplit_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(count),tojboolean(shuffle))
}

func TrainDataNative_setTrainTestSplit_1(nativeObj int64,count int) {
	C.Java_org_opencv_ml_TrainData_setTrainTestSplit_11(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(count))
}

func TrainDataNative_shuffleTrainTest_0(nativeObj int64) {
	C.Java_org_opencv_ml_TrainData_shuffleTrainTest_10(clzEnv,clzObj,(C.jlong)(nativeObj))
}
