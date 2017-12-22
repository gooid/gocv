package core

/*
#cgo windows,amd64 CFLAGS: -I${SRCDIR}/w64
#cgo windows,amd64 LDFLAGS: -L${SRCDIR}/w64

#cgo LDFLAGS: -lopencv_java3

#include "jni.h"
*/
import "C"

func Throw(ex interface{}) {
	panic(ex)
}

func NewJavaLangUnsupportedOperationException(str string) interface{} {
	return "UnsupportedOperationException:" + str
}
func NewIllegalArgumentException(str string) interface{} { return "IllegalArgumentException:" + str }
func NewRuntimeException(str string) interface{}         { return "RuntimeException:" + str }

// for native

var clzObj C.jclass

func togobool(b C.jboolean) bool {
	if b == 0 {
		return false
	}
	return true
}

func tojboolean(b bool) C.jboolean {
	if b {
		return C.jboolean(1)
	} else {
		return C.jboolean(0)
	}
}
