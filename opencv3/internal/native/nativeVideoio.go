package core

/*

#include "jni.h"

extern jlong Java_org_opencv_videoio_VideoCapture_VideoCapture_10(JNIEnv*, jclass, jstring, jint);
extern jlong Java_org_opencv_videoio_VideoCapture_VideoCapture_11(JNIEnv*, jclass, jstring);
extern jlong Java_org_opencv_videoio_VideoCapture_VideoCapture_12(JNIEnv*, jclass, jint);
extern jlong Java_org_opencv_videoio_VideoCapture_VideoCapture_13(JNIEnv*, jclass);
extern void Java_org_opencv_videoio_VideoCapture_delete(JNIEnv*, jclass, jlong);
extern jdouble Java_org_opencv_videoio_VideoCapture_get_10(JNIEnv*, jclass, jlong, jint);
extern jboolean Java_org_opencv_videoio_VideoCapture_grab_10(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_videoio_VideoCapture_isOpened_10(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_videoio_VideoCapture_open_10(JNIEnv*, jclass, jlong, jstring, jint);
extern jboolean Java_org_opencv_videoio_VideoCapture_open_11(JNIEnv*, jclass, jlong, jstring);
extern jboolean Java_org_opencv_videoio_VideoCapture_open_12(JNIEnv*, jclass, jlong, jint, jint);
extern jboolean Java_org_opencv_videoio_VideoCapture_open_13(JNIEnv*, jclass, jlong, jint);
extern jboolean Java_org_opencv_videoio_VideoCapture_read_10(JNIEnv*, jclass, jlong, jlong);
extern void Java_org_opencv_videoio_VideoCapture_release_10(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_videoio_VideoCapture_retrieve_10(JNIEnv*, jclass, jlong, jlong, jint);
extern jboolean Java_org_opencv_videoio_VideoCapture_retrieve_11(JNIEnv*, jclass, jlong, jlong);
extern jboolean Java_org_opencv_videoio_VideoCapture_set_10(JNIEnv*, jclass, jlong, jint, jdouble);
extern jlong Java_org_opencv_videoio_VideoWriter_VideoWriter_10(JNIEnv*, jclass, jstring, jint, jint, jdouble, jdouble, jdouble, jboolean);
extern jlong Java_org_opencv_videoio_VideoWriter_VideoWriter_11(JNIEnv*, jclass, jstring, jint, jint, jdouble, jdouble, jdouble);
extern jlong Java_org_opencv_videoio_VideoWriter_VideoWriter_12(JNIEnv*, jclass, jstring, jint, jdouble, jdouble, jdouble, jboolean);
extern jlong Java_org_opencv_videoio_VideoWriter_VideoWriter_13(JNIEnv*, jclass, jstring, jint, jdouble, jdouble, jdouble);
extern jlong Java_org_opencv_videoio_VideoWriter_VideoWriter_14(JNIEnv*, jclass);
extern void Java_org_opencv_videoio_VideoWriter_delete(JNIEnv*, jclass, jlong);
extern jint Java_org_opencv_videoio_VideoWriter_fourcc_10(JNIEnv*, jclass, jshort, jshort, jshort, jshort);
extern jdouble Java_org_opencv_videoio_VideoWriter_get_10(JNIEnv*, jclass, jlong, jint);
extern jboolean Java_org_opencv_videoio_VideoWriter_isOpened_10(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_videoio_VideoWriter_open_10(JNIEnv*, jclass, jlong, jstring, jint, jint, jdouble, jdouble, jdouble, jboolean);
extern jboolean Java_org_opencv_videoio_VideoWriter_open_11(JNIEnv*, jclass, jlong, jstring, jint, jint, jdouble, jdouble, jdouble);
extern jboolean Java_org_opencv_videoio_VideoWriter_open_12(JNIEnv*, jclass, jlong, jstring, jint, jdouble, jdouble, jdouble, jboolean);
extern jboolean Java_org_opencv_videoio_VideoWriter_open_13(JNIEnv*, jclass, jlong, jstring, jint, jdouble, jdouble, jdouble);
extern void Java_org_opencv_videoio_VideoWriter_release_10(JNIEnv*, jclass, jlong);
extern jboolean Java_org_opencv_videoio_VideoWriter_set_10(JNIEnv*, jclass, jlong, jint, jdouble);
extern void Java_org_opencv_videoio_VideoWriter_write_10(JNIEnv*, jclass, jlong, jlong);

*/
import "C"

func VideoCaptureNative_VideoCapture_0(filename string, apiPreference int) int64 {
	defer ungointerface(filename)
	return int64(C.Java_org_opencv_videoio_VideoCapture_VideoCapture_10(clzEnv, clzObj, tojstring(filename), (C.jint)(apiPreference)))
}
func VideoCaptureNative_VideoCapture_1(filename string) int64 {
	defer ungointerface(filename)
	return int64(C.Java_org_opencv_videoio_VideoCapture_VideoCapture_11(clzEnv, clzObj, tojstring(filename)))
}
func VideoCaptureNative_VideoCapture_2(index int) int64 {
	return int64(C.Java_org_opencv_videoio_VideoCapture_VideoCapture_12(clzEnv, clzObj, (C.jint)(index)))
}

func VideoCaptureNative_VideoCapture_3() int64 {
	return int64(C.Java_org_opencv_videoio_VideoCapture_VideoCapture_13(clzEnv, clzObj))
}

func VideoCaptureNative_delete(nativeObj int64) {
	C.Java_org_opencv_videoio_VideoCapture_delete(clzEnv, clzObj, (C.jlong)(nativeObj))
}

func VideoCaptureNative_get_0(nativeObj int64, propId int) float64 {
	return float64(C.Java_org_opencv_videoio_VideoCapture_get_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(propId)))
}

func VideoCaptureNative_grab_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_videoio_VideoCapture_grab_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func VideoCaptureNative_isOpened_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_videoio_VideoCapture_isOpened_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func VideoCaptureNative_open_0(nativeObj int64, filename string, apiPreference int) bool {
	defer ungointerface(filename)
	return togobool(C.Java_org_opencv_videoio_VideoCapture_open_10(clzEnv, clzObj, (C.jlong)(nativeObj), tojstring(filename), (C.jint)(apiPreference)))
}

func VideoCaptureNative_open_1(nativeObj int64, filename string) bool {
	defer ungointerface(filename)
	return togobool(C.Java_org_opencv_videoio_VideoCapture_open_11(clzEnv, clzObj, (C.jlong)(nativeObj), tojstring(filename)))
}

func VideoCaptureNative_open_2(nativeObj int64, cameraNum int, apiPreference int) bool {
	return togobool(C.Java_org_opencv_videoio_VideoCapture_open_12(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(cameraNum), (C.jint)(apiPreference)))
}

func VideoCaptureNative_open_3(nativeObj int64, index int) bool {
	return togobool(C.Java_org_opencv_videoio_VideoCapture_open_13(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(index)))
}

func VideoCaptureNative_read_0(nativeObj int64, image_nativeObj int64) bool {
	return togobool(C.Java_org_opencv_videoio_VideoCapture_read_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(image_nativeObj)))
}

func VideoCaptureNative_release_0(nativeObj int64) {
	C.Java_org_opencv_videoio_VideoCapture_release_10(clzEnv, clzObj, (C.jlong)(nativeObj))
}

func VideoCaptureNative_retrieve_0(nativeObj int64, image_nativeObj int64, flag int) bool {
	return togobool(C.Java_org_opencv_videoio_VideoCapture_retrieve_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(image_nativeObj), (C.jint)(flag)))
}

func VideoCaptureNative_retrieve_1(nativeObj int64, image_nativeObj int64) bool {
	return togobool(C.Java_org_opencv_videoio_VideoCapture_retrieve_11(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(image_nativeObj)))
}

func VideoCaptureNative_set_0(nativeObj int64, propId int, value float64) bool {
	return togobool(C.Java_org_opencv_videoio_VideoCapture_set_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(propId), (C.jdouble)(value)))
}

func VideoWriterNative_VideoWriter_0(filename string, apiPreference int, fourcc int, fps float64, frameSize_width float64, frameSize_height float64, isColor bool) int64 {
	defer ungointerface(filename)
	return int64(C.Java_org_opencv_videoio_VideoWriter_VideoWriter_10(clzEnv, clzObj, tojstring(filename), (C.jint)(apiPreference), (C.jint)(fourcc), (C.jdouble)(fps), (C.jdouble)(frameSize_width), (C.jdouble)(frameSize_height), tojboolean(isColor)))
}

func VideoWriterNative_VideoWriter_1(filename string, apiPreference int, fourcc int, fps float64, frameSize_width float64, frameSize_height float64) int64 {
	defer ungointerface(filename)
	return int64(C.Java_org_opencv_videoio_VideoWriter_VideoWriter_11(clzEnv, clzObj, tojstring(filename), (C.jint)(apiPreference), (C.jint)(fourcc), (C.jdouble)(fps), (C.jdouble)(frameSize_width), (C.jdouble)(frameSize_height)))
}

func VideoWriterNative_VideoWriter_2(filename string, fourcc int, fps float64, frameSize_width float64, frameSize_height float64, isColor bool) int64 {
	defer ungointerface(filename)
	return int64(C.Java_org_opencv_videoio_VideoWriter_VideoWriter_12(clzEnv, clzObj, tojstring(filename), (C.jint)(fourcc), (C.jdouble)(fps), (C.jdouble)(frameSize_width), (C.jdouble)(frameSize_height), tojboolean(isColor)))
}

func VideoWriterNative_VideoWriter_3(filename string, fourcc int, fps float64, frameSize_width float64, frameSize_height float64) int64 {
	defer ungointerface(filename)
	return int64(C.Java_org_opencv_videoio_VideoWriter_VideoWriter_13(clzEnv, clzObj, tojstring(filename), (C.jint)(fourcc), (C.jdouble)(fps), (C.jdouble)(frameSize_width), (C.jdouble)(frameSize_height)))
}

func VideoWriterNative_VideoWriter_4() int64 {
	return int64(C.Java_org_opencv_videoio_VideoWriter_VideoWriter_14(clzEnv, clzObj))
}

func VideoWriterNative_delete(nativeObj int64) {
	C.Java_org_opencv_videoio_VideoWriter_delete(clzEnv, clzObj, (C.jlong)(nativeObj))
}

func VideoWriterNative_fourcc_0(c1 rune, c2 rune, c3 rune, c4 rune) int {
	return int(C.Java_org_opencv_videoio_VideoWriter_fourcc_10(clzEnv, clzObj, (C.jshort)(c1),
		(C.jshort)(c2), (C.jshort)(c3), (C.jshort)(c4)))
}

func VideoWriterNative_get_0(nativeObj int64, propId int) float64 {
	return float64(C.Java_org_opencv_videoio_VideoWriter_get_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(propId)))
}

func VideoWriterNative_isOpened_0(nativeObj int64) bool {
	return togobool(C.Java_org_opencv_videoio_VideoWriter_isOpened_10(clzEnv, clzObj, (C.jlong)(nativeObj)))
}

func VideoWriterNative_open_0(nativeObj int64, filename string, apiPreference int, fourcc int, fps float64, frameSize_width float64, frameSize_height float64, isColor bool) bool {
	defer ungointerface(filename)
	return togobool(C.Java_org_opencv_videoio_VideoWriter_open_10(clzEnv, clzObj, (C.jlong)(nativeObj), tojstring(filename), (C.jint)(apiPreference), (C.jint)(fourcc), (C.jdouble)(fps), (C.jdouble)(frameSize_width), (C.jdouble)(frameSize_height), tojboolean(isColor)))
}

func VideoWriterNative_open_1(nativeObj int64, filename string, apiPreference int, fourcc int, fps float64, frameSize_width float64, frameSize_height float64) bool {
	defer ungointerface(filename)
	return togobool(C.Java_org_opencv_videoio_VideoWriter_open_11(clzEnv, clzObj, (C.jlong)(nativeObj), tojstring(filename), (C.jint)(apiPreference), (C.jint)(fourcc), (C.jdouble)(fps), (C.jdouble)(frameSize_width), (C.jdouble)(frameSize_height)))
}

func VideoWriterNative_open_2(nativeObj int64, filename string, fourcc int, fps float64, frameSize_width float64, frameSize_height float64, isColor bool) bool {
	defer ungointerface(filename)
	return togobool(C.Java_org_opencv_videoio_VideoWriter_open_12(clzEnv, clzObj, (C.jlong)(nativeObj), tojstring(filename), (C.jint)(fourcc), (C.jdouble)(fps), (C.jdouble)(frameSize_width), (C.jdouble)(frameSize_height), tojboolean(isColor)))
}

func VideoWriterNative_open_3(nativeObj int64, filename string, fourcc int, fps float64, frameSize_width float64, frameSize_height float64) bool {
	defer ungointerface(filename)
	return togobool(C.Java_org_opencv_videoio_VideoWriter_open_13(clzEnv, clzObj, (C.jlong)(nativeObj), tojstring(filename), (C.jint)(fourcc), (C.jdouble)(fps), (C.jdouble)(frameSize_width), (C.jdouble)(frameSize_height)))
}

func VideoWriterNative_release_0(nativeObj int64) {
	C.Java_org_opencv_videoio_VideoWriter_release_10(clzEnv, clzObj, (C.jlong)(nativeObj))
}

func VideoWriterNative_set_0(nativeObj int64, propId int, value float64) bool {
	return togobool(C.Java_org_opencv_videoio_VideoWriter_set_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jint)(propId), (C.jdouble)(value)))
}

func VideoWriterNative_write_0(nativeObj int64, image_nativeObj int64) {
	C.Java_org_opencv_videoio_VideoWriter_write_10(clzEnv, clzObj, (C.jlong)(nativeObj), (C.jlong)(image_nativeObj))
}
