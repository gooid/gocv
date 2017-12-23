package core

/*

#include "jni.h"

extern jlong Java_org_opencv_imgcodecs_Imgcodecs_imdecode_10(JNIEnv*, jclass, jlong, jint);
extern jboolean Java_org_opencv_imgcodecs_Imgcodecs_imencode_10(JNIEnv*, jclass, jstring, jlong, jlong, jlong);
extern jboolean Java_org_opencv_imgcodecs_Imgcodecs_imencode_11(JNIEnv*, jclass, jstring, jlong, jlong);
extern jlong Java_org_opencv_imgcodecs_Imgcodecs_imread_10(JNIEnv*, jclass, jstring, jint);
extern jlong Java_org_opencv_imgcodecs_Imgcodecs_imread_11(JNIEnv*, jclass, jstring);
extern jboolean Java_org_opencv_imgcodecs_Imgcodecs_imreadmulti_10(JNIEnv*, jclass, jstring, jlong, jint);
extern jboolean Java_org_opencv_imgcodecs_Imgcodecs_imreadmulti_11(JNIEnv*, jclass, jstring, jlong);
extern jboolean Java_org_opencv_imgcodecs_Imgcodecs_imwrite_10(JNIEnv*, jclass, jstring, jlong, jlong);
extern jboolean Java_org_opencv_imgcodecs_Imgcodecs_imwrite_11(JNIEnv*, jclass, jstring, jlong);

*/
import "C"

func ImgcodecsNative_imdecode_0(buf_nativeObj int64, flags int) int64 {
	return int64(C.Java_org_opencv_imgcodecs_Imgcodecs_imdecode_10(clzEnv, clzObj, (C.jlong)(buf_nativeObj), (C.jint)(flags)))
}
func ImgcodecsNative_imencode_0(ext string, img_nativeObj int64, buf_mat_nativeObj int64, params_mat_nativeObj int64) bool {
	defer ungointerface(ext)
	return togobool(C.Java_org_opencv_imgcodecs_Imgcodecs_imencode_10(clzEnv, clzObj, tojstring(ext), (C.jlong)(img_nativeObj), (C.jlong)(buf_mat_nativeObj), (C.jlong)(params_mat_nativeObj)))
}
func ImgcodecsNative_imencode_1(ext string, img_nativeObj int64, buf_mat_nativeObj int64) bool {
	defer ungointerface(ext)
	return togobool(C.Java_org_opencv_imgcodecs_Imgcodecs_imencode_11(clzEnv, clzObj, tojstring(ext), (C.jlong)(img_nativeObj), (C.jlong)(buf_mat_nativeObj)))
}

func ImgcodecsNative_imread_0(filename string, flags int) int64 {
	defer ungointerface(filename)
	return int64(C.Java_org_opencv_imgcodecs_Imgcodecs_imread_10(clzEnv, clzObj, tojstring(filename), (C.jint)(flags)))
}

func ImgcodecsNative_imread_1(filename string) int64 {
	defer ungointerface(filename)
	return int64(C.Java_org_opencv_imgcodecs_Imgcodecs_imread_11(clzEnv, clzObj, tojstring(filename)))
}

func ImgcodecsNative_imreadmulti_0(filename string, mats_mat_nativeObj int64, flags int) bool {
	defer ungointerface(filename)
	return togobool(C.Java_org_opencv_imgcodecs_Imgcodecs_imreadmulti_10(clzEnv, clzObj, tojstring(filename), (C.jlong)(mats_mat_nativeObj), (C.jint)(flags)))
}

func ImgcodecsNative_imreadmulti_1(filename string, mats_mat_nativeObj int64) bool {
	defer ungointerface(filename)
	return togobool(C.Java_org_opencv_imgcodecs_Imgcodecs_imreadmulti_11(clzEnv, clzObj, tojstring(filename), (C.jlong)(mats_mat_nativeObj)))
}

func ImgcodecsNative_imwrite_0(filename string, img_nativeObj int64, params_mat_nativeObj int64) bool {
	defer ungointerface(filename)
	return togobool(C.Java_org_opencv_imgcodecs_Imgcodecs_imwrite_10(clzEnv, clzObj, tojstring(filename), (C.jlong)(img_nativeObj), (C.jlong)(params_mat_nativeObj)))
}

func ImgcodecsNative_imwrite_1(filename string, img_nativeObj int64) bool {
	defer ungointerface(filename)
	return togobool(C.Java_org_opencv_imgcodecs_Imgcodecs_imwrite_11(clzEnv, clzObj, tojstring(filename), (C.jlong)(img_nativeObj)))
}
