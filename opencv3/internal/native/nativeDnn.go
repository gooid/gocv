package core

/*

#include "jni.h"

extern jlong Java_org_opencv_dnn_DictValue_DictValue_10(JNIEnv*, jclass, jstring);
extern jlong Java_org_opencv_dnn_DictValue_DictValue_11(JNIEnv*, jclass, jdouble);
extern jlong Java_org_opencv_dnn_DictValue_DictValue_12(JNIEnv*, jclass, jint);
extern void Java_org_opencv_dnn_DictValue_delete(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_dnn_DictValue_getIntValue_10(JNIEnv*, jclass, jlong, jint);
extern jint Java_org_opencv_dnn_DictValue_getIntValue_11(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_dnn_DictValue_getRealValue_10(JNIEnv*, jclass, jlong, jint);
extern jdouble Java_org_opencv_dnn_DictValue_getRealValue_11(JNIEnv*, jclass, jlong);
extern jstring Java_org_opencv_dnn_DictValue_getStringValue_10(JNIEnv*, jclass, jlong, jint);
extern jstring Java_org_opencv_dnn_DictValue_getStringValue_11(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_dnn_DictValue_isInt_10(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_dnn_DictValue_isReal_10(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_dnn_DictValue_isString_10(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_dnn_Dnn_blobFromImage_10(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jboolean, jboolean);
extern jlong Java_org_opencv_dnn_Dnn_blobFromImage_11(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_dnn_Dnn_blobFromImages_10(JNIEnv*, jclass, jlong, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jdouble, jboolean, jboolean);
extern jlong Java_org_opencv_dnn_Dnn_blobFromImages_11(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_dnn_Dnn_createCaffeImporter_10(JNIEnv*, jclass, jstring, jstring);
extern jlong Java_org_opencv_dnn_Dnn_createCaffeImporter_11(JNIEnv*, jclass, jstring);
extern jlong Java_org_opencv_dnn_Dnn_createTensorflowImporter_10(JNIEnv*, jclass, jstring);
extern jlong Java_org_opencv_dnn_Dnn_createTorchImporter_10(JNIEnv*, jclass, jstring, jboolean);
extern jlong Java_org_opencv_dnn_Dnn_createTorchImporter_11(JNIEnv*, jclass, jstring);
extern jlong Java_org_opencv_dnn_Dnn_readNetFromCaffe_10(JNIEnv*, jclass, jstring, jstring);
extern jlong Java_org_opencv_dnn_Dnn_readNetFromCaffe_11(JNIEnv*, jclass, jstring);
extern jlong Java_org_opencv_dnn_Dnn_readNetFromDarknet_10(JNIEnv*, jclass, jstring, jstring);
extern jlong Java_org_opencv_dnn_Dnn_readNetFromDarknet_11(JNIEnv*, jclass, jstring);
extern jlong Java_org_opencv_dnn_Dnn_readNetFromTensorflow_10(JNIEnv*, jclass, jstring, jstring);
extern jlong Java_org_opencv_dnn_Dnn_readNetFromTensorflow_11(JNIEnv*, jclass, jstring);
extern jlong Java_org_opencv_dnn_Dnn_readNetFromTorch_10(JNIEnv*, jclass, jstring, jboolean);
extern jlong Java_org_opencv_dnn_Dnn_readNetFromTorch_11(JNIEnv*, jclass, jstring);
extern jlong Java_org_opencv_dnn_Dnn_readTorchBlob_10(JNIEnv*, jclass, jstring, jboolean);
extern jlong Java_org_opencv_dnn_Dnn_readTorchBlob_11(JNIEnv*, jclass, jstring);
extern void Java_org_opencv_dnn_Dnn_shrinkCaffeModel_10(JNIEnv*, jclass, jstring, jstring);
extern void Java_org_opencv_dnn_Importer_delete(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_dnn_Importer_populateNet_10(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_dnn_Layer_delete(JNIEnv*, jclass, jlong);
extern jlong Java_org_opencv_dnn_Layer_finalize_10(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_dnn_Layer_finalize_11(JNIEnv*, jclass, jlong, jlong, jlong);
extern void Java_org_opencv_dnn_Layer_forward_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern jlong Java_org_opencv_dnn_Layer_get_1blobs_10(JNIEnv*, jclass, jlong);
extern jstring Java_org_opencv_dnn_Layer_get_1name_10(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_dnn_Layer_get_1preferableTarget_10(JNIEnv*, jclass, jlong);
extern jstring Java_org_opencv_dnn_Layer_get_1type_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_dnn_Layer_run_10(JNIEnv*, jclass, jlong, jlong, jlong, jlong);
extern void Java_org_opencv_dnn_Layer_set_1blobs_10(JNIEnv*, jclass, jlong, jlong);
extern jlong Java_org_opencv_dnn_Net_Net_10(JNIEnv*, jclass);
extern void Java_org_opencv_dnn_Net_connect_10(JNIEnv*, jclass, jlong, jstring, jstring);
extern void Java_org_opencv_dnn_Net_delete(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_dnn_Net_deleteLayer_10(JNIEnv*, jclass, jlong, jlong);
extern jboolean Java_org_opencv_dnn_Net_empty_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_dnn_Net_enableFusion_10(JNIEnv*, jclass, jlong, jboolean);
extern jlong Java_org_opencv_dnn_Net_forward_10(JNIEnv*, jclass, jlong, jstring);
extern jlong Java_org_opencv_dnn_Net_forward_11(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_dnn_Net_forward_12(JNIEnv*, jclass, jlong, jlong, jstring);
extern void Java_org_opencv_dnn_Net_forward_13(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_dnn_Net_forward_14(JNIEnv*, jclass, jlong, jlong, jlong);
extern jlong Java_org_opencv_dnn_Net_getFLOPS_10(JNIEnv*, jclass, jlong, jlong);
extern jlong Java_org_opencv_dnn_Net_getFLOPS_11(JNIEnv*, jclass, jlong, jint, jlong);
extern jlong Java_org_opencv_dnn_Net_getFLOPS_12(JNIEnv*, jclass, jlong, jint, jlong);
extern jlong Java_org_opencv_dnn_Net_getFLOPS_13(JNIEnv*, jclass, jlong, jlong);
extern jint Java_org_opencv_dnn_Net_getLayerId_10(JNIEnv*, jclass, jlong, jstring);
extern jlong Java_org_opencv_dnn_Net_getLayerNames_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_dnn_Net_getLayerTypes_10(JNIEnv*, jclass, jlong, jlong);
extern jlong Java_org_opencv_dnn_Net_getLayer_10(JNIEnv*, jclass, jlong, jlong);
extern jint Java_org_opencv_dnn_Net_getLayersCount_10(JNIEnv*, jclass, jlong, jstring);
extern void Java_org_opencv_dnn_Net_getMemoryConsumption_10(JNIEnv*, jclass, jlong, jlong, jdoubleArray, jdoubleArray);
extern void Java_org_opencv_dnn_Net_getMemoryConsumption_11(JNIEnv*, jclass, jlong, jint, jlong, jdoubleArray, jdoubleArray);
extern void Java_org_opencv_dnn_Net_getMemoryConsumption_12(JNIEnv*, jclass, jlong, jint, jlong, jdoubleArray, jdoubleArray);
extern jlong Java_org_opencv_dnn_Net_getParam_10(JNIEnv*, jclass, jlong, jlong, jint);
extern jlong Java_org_opencv_dnn_Net_getParam_11(JNIEnv*, jclass, jlong, jlong);
extern jlong Java_org_opencv_dnn_Net_getPerfProfile_10(JNIEnv*, jclass, jlong, jlong);
extern jlong Java_org_opencv_dnn_Net_getUnconnectedOutLayers_10(JNIEnv*, jclass, jlong);
extern void Java_org_opencv_dnn_Net_setHalideScheduler_10(JNIEnv*, jclass, jlong, jstring);
extern void Java_org_opencv_dnn_Net_setInput_10(JNIEnv*, jclass, jlong, jlong, jstring);
extern void Java_org_opencv_dnn_Net_setInput_11(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_dnn_Net_setInputsNames_10(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_dnn_Net_setParam_10(JNIEnv*, jclass, jlong, jlong, jint, jlong);
extern void Java_org_opencv_dnn_Net_setPreferableBackend_10(JNIEnv*, jclass, jlong, jint);
extern void Java_org_opencv_dnn_Net_setPreferableTarget_10(JNIEnv*, jclass, jlong, jint);

*/
import "C"

func DictValueNative_DictValue_0(s string) int64 {
	defer ungointerface(s)
	return int64(C.Java_org_opencv_dnn_DictValue_DictValue_10(clzEnv,clzObj,tojstring(s)))
}
func DictValueNative_DictValue_1(p float64) int64 {
	return int64(C.Java_org_opencv_dnn_DictValue_DictValue_11(clzEnv,clzObj,(C.jdouble)(p)))
}
func DictValueNative_DictValue_2(i int) int64 {
	return int64(C.Java_org_opencv_dnn_DictValue_DictValue_12(clzEnv,clzObj,(C.jint)(i)))
}

func DictValueNative_delete(nativeObj int64) {
	C.Java_org_opencv_dnn_DictValue_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func DictValueNative_getIntValue_0(nativeObj int64,idx int) int {
	return int(C.Java_org_opencv_dnn_DictValue_getIntValue_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(idx)))
}

func DictValueNative_getIntValue_1(nativeObj int64) int {
	return int(C.Java_org_opencv_dnn_DictValue_getIntValue_11(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func DictValueNative_getRealValue_0(nativeObj int64,idx int) float64 {
	return float64(C.Java_org_opencv_dnn_DictValue_getRealValue_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(idx)))
}

func DictValueNative_getRealValue_1(nativeObj int64) float64 {
	return float64(C.Java_org_opencv_dnn_DictValue_getRealValue_11(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func DictValueNative_getStringValue_0(nativeObj int64,idx int) string {
	return togostring(
			C.Java_org_opencv_dnn_DictValue_getStringValue_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(idx)))
}

func DictValueNative_getStringValue_1(nativeObj int64) string {
	return togostring(
			C.Java_org_opencv_dnn_DictValue_getStringValue_11(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func DictValueNative_isInt_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_dnn_DictValue_isInt_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func DictValueNative_isReal_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_dnn_DictValue_isReal_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func DictValueNative_isString_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_dnn_DictValue_isString_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func DnnNative_blobFromImage_0(image_nativeObj int64,scalefactor float64,size_width float64,size_height float64,mean_val0 float64,mean_val1 float64,mean_val2 float64,mean_val3 float64,swapRB bool,crop bool) int64 {
	return int64(C.Java_org_opencv_dnn_Dnn_blobFromImage_10(clzEnv,clzObj,(C.jlong)(image_nativeObj),(C.jdouble)(scalefactor),(C.jdouble)(size_width),(C.jdouble)(size_height),(C.jdouble)(mean_val0),(C.jdouble)(mean_val1),(C.jdouble)(mean_val2),(C.jdouble)(mean_val3),tojboolean(swapRB),tojboolean(crop)))
}

func DnnNative_blobFromImage_1(image_nativeObj int64) int64 {
	return int64(C.Java_org_opencv_dnn_Dnn_blobFromImage_11(clzEnv,clzObj,(C.jlong)(image_nativeObj)))
}

func DnnNative_blobFromImages_0(images_mat_nativeObj int64,scalefactor float64,size_width float64,size_height float64,mean_val0 float64,mean_val1 float64,mean_val2 float64,mean_val3 float64,swapRB bool,crop bool) int64 {
	return int64(C.Java_org_opencv_dnn_Dnn_blobFromImages_10(clzEnv,clzObj,(C.jlong)(images_mat_nativeObj),(C.jdouble)(scalefactor),(C.jdouble)(size_width),(C.jdouble)(size_height),(C.jdouble)(mean_val0),(C.jdouble)(mean_val1),(C.jdouble)(mean_val2),(C.jdouble)(mean_val3),tojboolean(swapRB),tojboolean(crop)))
}

func DnnNative_blobFromImages_1(images_mat_nativeObj int64) int64 {
	return int64(C.Java_org_opencv_dnn_Dnn_blobFromImages_11(clzEnv,clzObj,(C.jlong)(images_mat_nativeObj)))
}

func DnnNative_createCaffeImporter_0(prototxt string,caffeModel string) int64 {
	defer ungointerface(prototxt)
	defer ungointerface(caffeModel)
	return int64(C.Java_org_opencv_dnn_Dnn_createCaffeImporter_10(clzEnv,clzObj,tojstring(prototxt),tojstring(caffeModel)))
}

func DnnNative_createCaffeImporter_1(prototxt string) int64 {
	defer ungointerface(prototxt)
	return int64(C.Java_org_opencv_dnn_Dnn_createCaffeImporter_11(clzEnv,clzObj,tojstring(prototxt)))
}

func DnnNative_createTensorflowImporter_0(model string) int64 {
	defer ungointerface(model)
	return int64(C.Java_org_opencv_dnn_Dnn_createTensorflowImporter_10(clzEnv,clzObj,tojstring(model)))
}

func DnnNative_createTorchImporter_0(filename string,isBinary bool) int64 {
	defer ungointerface(filename)
	return int64(C.Java_org_opencv_dnn_Dnn_createTorchImporter_10(clzEnv,clzObj,tojstring(filename),tojboolean(isBinary)))
}

func DnnNative_createTorchImporter_1(filename string) int64 {
	defer ungointerface(filename)
	return int64(C.Java_org_opencv_dnn_Dnn_createTorchImporter_11(clzEnv,clzObj,tojstring(filename)))
}

func DnnNative_readNetFromCaffe_0(prototxt string,caffeModel string) int64 {
	defer ungointerface(prototxt)
	defer ungointerface(caffeModel)
	return int64(C.Java_org_opencv_dnn_Dnn_readNetFromCaffe_10(clzEnv,clzObj,tojstring(prototxt),tojstring(caffeModel)))
}

func DnnNative_readNetFromCaffe_1(prototxt string) int64 {
	defer ungointerface(prototxt)
	return int64(C.Java_org_opencv_dnn_Dnn_readNetFromCaffe_11(clzEnv,clzObj,tojstring(prototxt)))
}

func DnnNative_readNetFromDarknet_0(cfgFile string,darknetModel string) int64 {
	defer ungointerface(cfgFile)
	defer ungointerface(darknetModel)
	return int64(C.Java_org_opencv_dnn_Dnn_readNetFromDarknet_10(clzEnv,clzObj,tojstring(cfgFile),tojstring(darknetModel)))
}

func DnnNative_readNetFromDarknet_1(cfgFile string) int64 {
	defer ungointerface(cfgFile)
	return int64(C.Java_org_opencv_dnn_Dnn_readNetFromDarknet_11(clzEnv,clzObj,tojstring(cfgFile)))
}

func DnnNative_readNetFromTensorflow_0(model string,config string) int64 {
	defer ungointerface(model)
	defer ungointerface(config)
	return int64(C.Java_org_opencv_dnn_Dnn_readNetFromTensorflow_10(clzEnv,clzObj,tojstring(model),tojstring(config)))
}

func DnnNative_readNetFromTensorflow_1(model string) int64 {
	defer ungointerface(model)
	return int64(C.Java_org_opencv_dnn_Dnn_readNetFromTensorflow_11(clzEnv,clzObj,tojstring(model)))
}

func DnnNative_readNetFromTorch_0(model string,isBinary bool) int64 {
	defer ungointerface(model)
	return int64(C.Java_org_opencv_dnn_Dnn_readNetFromTorch_10(clzEnv,clzObj,tojstring(model),tojboolean(isBinary)))
}

func DnnNative_readNetFromTorch_1(model string) int64 {
	defer ungointerface(model)
	return int64(C.Java_org_opencv_dnn_Dnn_readNetFromTorch_11(clzEnv,clzObj,tojstring(model)))
}

func DnnNative_readTorchBlob_0(filename string,isBinary bool) int64 {
	defer ungointerface(filename)
	return int64(C.Java_org_opencv_dnn_Dnn_readTorchBlob_10(clzEnv,clzObj,tojstring(filename),tojboolean(isBinary)))
}

func DnnNative_readTorchBlob_1(filename string) int64 {
	defer ungointerface(filename)
	return int64(C.Java_org_opencv_dnn_Dnn_readTorchBlob_11(clzEnv,clzObj,tojstring(filename)))
}

func DnnNative_shrinkCaffeModel_0(src string,dst string) {
	defer ungointerface(src)
	defer ungointerface(dst)
	C.Java_org_opencv_dnn_Dnn_shrinkCaffeModel_10(clzEnv,clzObj,tojstring(src),tojstring(dst))
}

func ImporterNative_delete(nativeObj int64) {
	C.Java_org_opencv_dnn_Importer_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func ImporterNative_populateNet_0(nativeObj int64,net_nativeObj int64) {
	C.Java_org_opencv_dnn_Importer_populateNet_10(clzEnv,clzObj,(C.jlong)(
		nativeObj),(C.jlong)(net_nativeObj))
}

func LayerNative_delete(nativeObj int64) {
	C.Java_org_opencv_dnn_Layer_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func LayerNative_finalize_0(nativeObj int64,inputs_mat_nativeObj int64) int64 {
	return int64(C.Java_org_opencv_dnn_Layer_finalize_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(inputs_mat_nativeObj)))
}

func LayerNative_finalize_1(nativeObj int64,inputs_mat_nativeObj int64,outputs_mat_nativeObj int64) {
	C.Java_org_opencv_dnn_Layer_finalize_11(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(inputs_mat_nativeObj),(C.jlong)(outputs_mat_nativeObj))
}

func LayerNative_forward_0(nativeObj int64,inputs_mat_nativeObj int64,outputs_mat_nativeObj int64,internals_mat_nativeObj int64) {
	C.Java_org_opencv_dnn_Layer_forward_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(inputs_mat_nativeObj),(C.jlong)(outputs_mat_nativeObj),(C.jlong)(internals_mat_nativeObj))
}

func LayerNative_get_blobs_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_dnn_Layer_get_1blobs_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func LayerNative_get_name_0(nativeObj int64) string {
	return togostring(
			C.Java_org_opencv_dnn_Layer_get_1name_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func LayerNative_get_preferableTarget_0(nativeObj int64) int {
	return int(C.Java_org_opencv_dnn_Layer_get_1preferableTarget_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func LayerNative_get_type_0(nativeObj int64) string {
	return togostring(
			C.Java_org_opencv_dnn_Layer_get_1type_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func LayerNative_run_0(nativeObj int64,inputs_mat_nativeObj int64,outputs_mat_nativeObj int64,internals_mat_nativeObj int64) {
	C.Java_org_opencv_dnn_Layer_run_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(inputs_mat_nativeObj),(C.jlong)(outputs_mat_nativeObj),(C.jlong)(internals_mat_nativeObj))
}

func LayerNative_set_blobs_0(nativeObj int64,blobs_mat_nativeObj int64) {
	C.Java_org_opencv_dnn_Layer_set_1blobs_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(blobs_mat_nativeObj))
}

func NetNative_Net_0() int64 {
	return int64(C.Java_org_opencv_dnn_Net_Net_10(clzEnv,clzObj))
}

func NetNative_connect_0(nativeObj int64,outPin string,inpPin string) {
	defer ungointerface(outPin)
	defer ungointerface(inpPin)
	C.Java_org_opencv_dnn_Net_connect_10(clzEnv,clzObj,(C.jlong)(nativeObj),tojstring(outPin),tojstring(inpPin))
}

func NetNative_delete(nativeObj int64) {
	C.Java_org_opencv_dnn_Net_delete(clzEnv,clzObj,(C.jlong)(nativeObj))
}

func NetNative_deleteLayer_0(nativeObj int64,layer_nativeObj int64) {
	C.Java_org_opencv_dnn_Net_deleteLayer_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(layer_nativeObj))
}

func NetNative_empty_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_dnn_Net_empty_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func NetNative_enableFusion_0(nativeObj int64,fusion bool) {
	C.Java_org_opencv_dnn_Net_enableFusion_10(clzEnv,clzObj,(C.jlong)(nativeObj),tojboolean(fusion))
}

func NetNative_forward_0(nativeObj int64,outputName string) int64 {
	defer ungointerface(outputName)
	return int64(C.Java_org_opencv_dnn_Net_forward_10(clzEnv,clzObj,(C.jlong)(nativeObj),tojstring(outputName)))
}

func NetNative_forward_1(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_dnn_Net_forward_11(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func NetNative_forward_2(nativeObj int64,outputBlobs_mat_nativeObj int64,outputName string) {
	defer ungointerface(outputName)
	C.Java_org_opencv_dnn_Net_forward_12(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(outputBlobs_mat_nativeObj),tojstring(outputName))
}

func NetNative_forward_3(nativeObj int64,outputBlobs_mat_nativeObj int64) {
	C.Java_org_opencv_dnn_Net_forward_13(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(outputBlobs_mat_nativeObj))
}

func NetNative_forward_4(nativeObj int64,outputBlobs_mat_nativeObj int64,outBlobNames int64) {
	C.Java_org_opencv_dnn_Net_forward_14(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(outputBlobs_mat_nativeObj),(C.jlong)(outBlobNames))
}

func NetNative_getFLOPS_0(nativeObj int64,netInputShape_mat_nativeObj int64) int64 {
	return int64(C.Java_org_opencv_dnn_Net_getFLOPS_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(netInputShape_mat_nativeObj)))
}

func NetNative_getFLOPS_1(nativeObj int64,layerId int,netInputShape_mat_nativeObj int64) int64 {
	return int64(C.Java_org_opencv_dnn_Net_getFLOPS_11(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(layerId),(C.jlong)(netInputShape_mat_nativeObj)))
}

func NetNative_getFLOPS_2(nativeObj int64,layerId int,netInputShapes int64) int64 {
	return int64(C.Java_org_opencv_dnn_Net_getFLOPS_12(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(layerId),(C.jlong)(netInputShapes)))
}

func NetNative_getFLOPS_3(nativeObj int64,netInputShapes int64) int64 {
	return int64(C.Java_org_opencv_dnn_Net_getFLOPS_13(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(netInputShapes)))
}

func NetNative_getLayerId_0(nativeObj int64,layer string) int {
	defer ungointerface(layer)
	return int(C.Java_org_opencv_dnn_Net_getLayerId_10(clzEnv,clzObj,(C.jlong)(nativeObj),tojstring(layer)))
}

func NetNative_getLayerNames_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_dnn_Net_getLayerNames_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func NetNative_getLayerTypes_0(nativeObj int64,layersTypes int64) {
	C.Java_org_opencv_dnn_Net_getLayerTypes_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(layersTypes))
}

func NetNative_getLayer_0(nativeObj int64,layerId_nativeObj int64) int64 {
	return int64(C.Java_org_opencv_dnn_Net_getLayer_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(layerId_nativeObj)))
}

func NetNative_getLayersCount_0(nativeObj int64,layerType string) int {
	defer ungointerface(layerType)
	return int(C.Java_org_opencv_dnn_Net_getLayersCount_10(clzEnv,clzObj,(C.jlong)(nativeObj),tojstring(layerType)))
}

func NetNative_getMemoryConsumption_0(nativeObj int64,netInputShape_mat_nativeObj int64,weights_out []float64,blobs_out []float64) {
	defer ungointerface(weights_out)
	defer ungointerface(blobs_out)
	C.Java_org_opencv_dnn_Net_getMemoryConsumption_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(netInputShape_mat_nativeObj),tojdoubleArray(weights_out),tojdoubleArray(blobs_out))
}

func NetNative_getMemoryConsumption_1(nativeObj int64,layerId int,netInputShape_mat_nativeObj int64,weights_out []float64,blobs_out []float64) {
	defer ungointerface(weights_out)
	defer ungointerface(blobs_out)
	C.Java_org_opencv_dnn_Net_getMemoryConsumption_11(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(layerId),(C.jlong)(netInputShape_mat_nativeObj),tojdoubleArray(weights_out),tojdoubleArray(blobs_out))
}

func NetNative_getMemoryConsumption_2(nativeObj int64,layerId int,netInputShapes int64,weights_out []float64,blobs_out []float64) {
	defer ungointerface(weights_out)
	defer ungointerface(blobs_out)
	C.Java_org_opencv_dnn_Net_getMemoryConsumption_12(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(layerId),(C.jlong)(netInputShapes),tojdoubleArray(weights_out),tojdoubleArray(blobs_out))
}

func NetNative_getParam_0(nativeObj int64,layer_nativeObj int64,numParam int) int64 {
	return int64(C.Java_org_opencv_dnn_Net_getParam_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(layer_nativeObj),(C.jint)(numParam)))
}

func NetNative_getParam_1(nativeObj int64,layer_nativeObj int64) int64 {
	return int64(C.Java_org_opencv_dnn_Net_getParam_11(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(layer_nativeObj)))
}

func NetNative_getPerfProfile_0(nativeObj int64,timings_mat_nativeObj int64) int64 {
	return int64(C.Java_org_opencv_dnn_Net_getPerfProfile_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(timings_mat_nativeObj)))
}

func NetNative_getUnconnectedOutLayers_0(nativeObj int64) int64 {
	return int64(C.Java_org_opencv_dnn_Net_getUnconnectedOutLayers_10(clzEnv,clzObj,(C.jlong)(nativeObj)))
}

func NetNative_setHalideScheduler_0(nativeObj int64,scheduler string) {
	defer ungointerface(scheduler)
	C.Java_org_opencv_dnn_Net_setHalideScheduler_10(clzEnv,clzObj,(C.jlong)(nativeObj),tojstring(scheduler))
}

func NetNative_setInput_0(nativeObj int64,blob_nativeObj int64,name string) {
	defer ungointerface(name)
	C.Java_org_opencv_dnn_Net_setInput_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(blob_nativeObj),tojstring(name))
}

func NetNative_setInput_1(nativeObj int64,blob_nativeObj int64) {
	C.Java_org_opencv_dnn_Net_setInput_11(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(blob_nativeObj))
}

func NetNative_setInputsNames_0(nativeObj int64,inputBlobNames int64) {
	C.Java_org_opencv_dnn_Net_setInputsNames_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(inputBlobNames))
}

func NetNative_setParam_0(nativeObj int64,layer_nativeObj int64,numParam int,blob_nativeObj int64) {
	C.Java_org_opencv_dnn_Net_setParam_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jlong)(layer_nativeObj),(C.jint)(numParam),(C.jlong)(blob_nativeObj))
}

func NetNative_setPreferableBackend_0(nativeObj int64,backendId int) {
	C.Java_org_opencv_dnn_Net_setPreferableBackend_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(backendId))
}

func NetNative_setPreferableTarget_0(nativeObj int64,targetId int) {
	C.Java_org_opencv_dnn_Net_setPreferableTarget_10(clzEnv,clzObj,(C.jlong)(nativeObj),(C.jint)(targetId))
}
