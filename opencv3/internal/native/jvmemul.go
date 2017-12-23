package core

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
#include "jni.h"
#ifdef __ANDROID__
#define JNINativeInterface_ JNINativeInterface
#define JNIInvokeInterface_ JNIInvokeInterface
#endif
#include <stdlib.h>
typedef const char* cstr;
typedef const jdouble * cpjdouble;
typedef const jint * cpjint;

extern jclass FindClass(JNIEnv *env, const char *name);
extern jint ThrowNew(JNIEnv *env, jclass clazz, const char *msg);
extern jdoubleArray NewDoubleArray(JNIEnv *env, jsize len);
extern void SetDoubleArrayRegion(JNIEnv *env, jdoubleArray array, jsize start, jsize len, const jdouble *buf);
extern void SetIntArrayRegion(JNIEnv *env, jintArray array, jsize start, jsize len, const jint *buf);
extern void * iGetPrimitiveArrayCritical(JNIEnv *env, jarray array, jboolean *isCopy);
extern void ReleasePrimitiveArrayCritical(JNIEnv *env, jarray array, void *carray, jint mode);
extern jstring NewStringUTF(JNIEnv *env, const char *utf);
extern const char* GetStringUTFChars(JNIEnv *env, jstring str, jboolean *isCopy);
extern void ReleaseStringUTFChars(JNIEnv *env, jstring str, const char* chars);
static void * GetPrimitiveArrayCritical(JNIEnv *env, jarray array, jboolean *isCopy) {
	return (void *)(~((jlong)(iGetPrimitiveArrayCritical(env, array, isCopy))));
}

static JavaVM gccJVM;
static JNIEnv gccEnv;

static jint DestroyJavaVM(JavaVM *vm) {	return JNI_OK; }
static jint AttachCurrentThread(JavaVM *vm, void **penv, void *args) { return JNI_OK; }
static jint DetachCurrentThread(JavaVM *vm) { return JNI_OK; }
static jint GetEnv(JavaVM *vm, void **penv, jint version) { *penv = (void *)&gccEnv; return JNI_OK; }
static jint AttachCurrentThreadAsDaemon(JavaVM *vm, void **penv, void *args) { return JNI_OK; }

extern jint JNI_OnLoad(JavaVM *vm, void *reserved);
extern void JNI_OnUnload(JavaVM *vm, void *reserved);

static JNIEnv* setEnv() {
	struct JNINativeInterface_ *env =
		(struct JNINativeInterface_ *)malloc(sizeof(struct JNINativeInterface_));

	env->FindClass = FindClass;
	env->ThrowNew = ThrowNew;

	env->NewDoubleArray = NewDoubleArray;
	env->SetDoubleArrayRegion = SetDoubleArrayRegion;
	env->SetIntArrayRegion = SetIntArrayRegion;
	env->GetPrimitiveArrayCritical = GetPrimitiveArrayCritical;
	env->ReleasePrimitiveArrayCritical = ReleasePrimitiveArrayCritical;

	env->NewStringUTF = NewStringUTF;
	env->GetStringUTFChars = GetStringUTFChars;
	env->ReleaseStringUTFChars = ReleaseStringUTFChars;

	struct JNIInvokeInterface_ *jvm = (struct JNIInvokeInterface_ *)malloc(sizeof(struct JNIInvokeInterface_));
	jvm->DestroyJavaVM = DestroyJavaVM;
	jvm->AttachCurrentThread = AttachCurrentThread;
	jvm->DetachCurrentThread = DetachCurrentThread;
	jvm->GetEnv = GetEnv;
	jvm->AttachCurrentThreadAsDaemon = AttachCurrentThreadAsDaemon;

	gccJVM = (JavaVM)jvm;
	gccEnv = (JNIEnv)env;
	JNI_OnLoad((JavaVM*)&jvm, (void *)0);
	return &gccEnv;
}

*/
import "C"

func init() {
	clzEnv = C.setEnv()
}

var clzEnv *C.JNIEnv

func defaultJNIEnv() *C.JNIEnv {
	return clzEnv
}

//export FindClass
func FindClass(env *C.JNIEnv, name C.cstr) C.jclass {
	//fmt.Printf("FindClass(%v, %v)\n", env, C.GoString(name))
	return nil
}

//export ThrowNew
func ThrowNew(env *C.JNIEnv, clazz C.jclass, msg C.cstr) C.jint {
	panic(C.GoString(msg))
	return 0
}

//export NewDoubleArray
func NewDoubleArray(env *C.JNIEnv, size C.jsize) C.jdoubleArray {
	buf := make([]float64, size)
	return tojdoubleArray(buf)
}

//export SetDoubleArrayRegion
func SetDoubleArrayRegion(env *C.JNIEnv, array C.jdoubleArray, start C.jsize, size C.jsize, buf C.cpjdouble) {
	arr := togointerface(array, false).([]float64)
	b := (*[1 << 26]float64)(unsafe.Pointer(buf))[:size]
	copy(arr[start:start+size], b)
}

//export SetIntArrayRegion
func SetIntArrayRegion(env *C.JNIEnv, array C.jintArray, start C.jsize, size C.jsize, buf C.cpjint) {
	arr := togointerface(array, false).([]int32)
	b := (*[1 << 27]int32)(unsafe.Pointer(buf))[:size]
	copy(arr[start:start+size], b)
}

//export iGetPrimitiveArrayCritical
func iGetPrimitiveArrayCritical(env *C.JNIEnv, array C.jarray, isCopy *C.jboolean) unsafe.Pointer {
	return unsafe.Pointer(^uintptr(getPrimitiveArrayCritical(env, array, isCopy)))
}

//export ReleasePrimitiveArrayCritical
func ReleasePrimitiveArrayCritical(env *C.JNIEnv, array C.jarray, carray unsafe.Pointer, mode C.jint) {
	// nothing
}

//export NewStringUTF
func NewStringUTF(env *C.JNIEnv, utf C.cstr) C.jstring {
	str := C.GoString(utf)
	return tojstring(str)
}

//export GetStringUTFChars
func GetStringUTFChars(env *C.JNIEnv, str C.jstring, isCopy *C.jboolean) C.cstr {
	gstr := togointerface(str, false).(string)
	return C.CString(gstr)
}

//export ReleaseStringUTFChars
func ReleaseStringUTFChars(env *C.JNIEnv, str C.jstring, chars C.cstr) {
	C.free(unsafe.Pointer(chars))
}

// java / go convert
var keepLives = map[C.jobject]interface{}{}

func getPrimitiveArrayCritical(env *C.JNIEnv, array C.jarray, isCopy *C.jboolean) unsafe.Pointer {
	if obj, ok := keepLives[array]; ok {
		switch o := obj.(type) {
		case []byte:
			return unsafe.Pointer(&o[0])
		case []int16:
			return unsafe.Pointer(&o[0])
		case []int32:
			return unsafe.Pointer(&o[0])
		case []float32:
			return unsafe.Pointer(&o[0])
		case []float64:
			return unsafe.Pointer(&o[0])
		default:
			panic(fmt.Errorf("getPrimitiveArrayCritical: %t", obj))
		}
	}
	panic(fmt.Errorf("getPrimitiveArrayCritical: not found %p", array))
	return nil
}

// get interface
func togointerface(i C.jobject, release bool) interface{} {
	if obj, ok := keepLives[i]; ok {
		if release {
			delete(keepLives, i)
		}
		return obj
	}
	panic(fmt.Errorf("togointerface"))
}
func togostring(i C.jstring) string {
	return togointerface(i, true).(string)
}
func togointArray(i C.jobject) []int32 {
	return togointerface(i, true).([]int32)
}
func togofloat64Array(i C.jobject) []float64 {
	return togointerface(i, true).([]float64)
}

// convert to jobject
func tojobject(i interface{}) C.jobject {
	return C.jobject(unsafe.Pointer(^reflect.ValueOf(i).Pointer()))
}
// release keepLive
func ungointerface(i interface{}) {
	if s, ok := i.(string); ok {
		delete(keepLives, tojobject([]byte(s)))
	} else {
		delete(keepLives, tojobject(i))
	}
}
func tojstring(i string) C.jstring {
	jobj := tojobject([]byte(i))
	keepLives[jobj] = i
	return jobj
}

func tojbyteArray(i []byte) C.jobject {
	jobj := tojobject(i)
	keepLives[jobj] = i
	return jobj
}

func tojintArray(i []int32) C.jobject {
	jobj := tojobject(i)
	keepLives[jobj] = i
	return jobj
}
func tojshortArray(i []int16) C.jobject {
	jobj := tojobject(i)
	keepLives[jobj] = i
	return jobj
}
func tojfloatArray(i []float32) C.jobject {
	jobj := tojobject(i)
	keepLives[jobj] = i
	return jobj
}
func tojdoubleArray(i []float64) C.jobject {
	jobj := tojobject(i)
	keepLives[jobj] = i
	return jobj
}
