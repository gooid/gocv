package core

import (
	"fmt"
	"unsafe"
)

/*
#cgo windows,amd64 CFLAGS: -I${SRCDIR}/w64
#cgo windows,amd64 LDFLAGS: -L${SRCDIR}/w64

// 动态加载运行库时，需重命名JNI_OnLoad，否则极易和其它库的JNI_OnLoad重名
#cgo android CFLAGS: -DDYNLOAD -DJNI_OnLoad=cvJNI_OnLoad
#cgo !android LDFLAGS: -lopencv_java3

extern void setLibPath(const char*);
extern int InitProcs();
#include <stdlib.h>
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

var bInit = false

func Load() error {
	if bInit {
		return nil
	}
	i := C.InitProcs()
	if i != 0 {
		return fmt.Errorf("LoadLib fail count =", i)
	}
	bInit = true
	Init()
	return nil
}

func LoadLib(path string) error {
	cstr := C.CString(path)
	defer C.free(unsafe.Pointer(cstr))
	C.setLibPath(cstr)
	return Load()
}
