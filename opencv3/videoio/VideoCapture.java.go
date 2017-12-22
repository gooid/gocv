package videoio

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

type VideoCapture struct {
	nativeObj int64
}

func NewVideoCapture(addr int64) (rcvr *VideoCapture) {
	rcvr = &VideoCapture{}
	rcvr.nativeObj = addr
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func NewVideoCapture2(filename string, apiPreference int) (rcvr *VideoCapture) {
	rcvr = &VideoCapture{}
	rcvr.nativeObj = VideoCaptureNative_VideoCapture_0(filename, apiPreference)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
	return
}
func NewVideoCapture3(filename string) (rcvr *VideoCapture) {
	rcvr = &VideoCapture{}
	rcvr.nativeObj = VideoCaptureNative_VideoCapture_1(filename)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
	return
}
func NewVideoCapture4(index int) (rcvr *VideoCapture) {
	rcvr = &VideoCapture{}
	rcvr.nativeObj = VideoCaptureNative_VideoCapture_2(index)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
	return
}
func NewVideoCapture5() (rcvr *VideoCapture) {
	rcvr = &VideoCapture{}
	rcvr.nativeObj = VideoCaptureNative_VideoCapture_3()
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
	return
}
func (rcvr *VideoCapture) finalize() {
	VideoCaptureNative_delete(rcvr.nativeObj)
}
func (rcvr *VideoCapture) Get(propId int) float64 {
	retVal := VideoCaptureNative_get_0(rcvr.nativeObj, propId)
	return retVal
}
func (rcvr *VideoCapture) GetNativeObjAddr() int64 {
	return rcvr.nativeObj
}
func (rcvr *VideoCapture) Grab() bool {
	retVal := VideoCaptureNative_grab_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *VideoCapture) IsOpened() bool {
	retVal := VideoCaptureNative_isOpened_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *VideoCapture) Open(filename string, apiPreference int) bool {
	retVal := VideoCaptureNative_open_0(rcvr.nativeObj, filename, apiPreference)
	return retVal
}
func (rcvr *VideoCapture) Open2(filename string) bool {
	retVal := VideoCaptureNative_open_1(rcvr.nativeObj, filename)
	return retVal
}
func (rcvr *VideoCapture) Open3(cameraNum int, apiPreference int) bool {
	retVal := VideoCaptureNative_open_2(rcvr.nativeObj, cameraNum, apiPreference)
	return retVal
}
func (rcvr *VideoCapture) Open4(index int) bool {
	retVal := VideoCaptureNative_open_3(rcvr.nativeObj, index)
	return retVal
}
func (rcvr *VideoCapture) Read(image *Mat) bool {
	retVal := VideoCaptureNative_read_0(rcvr.nativeObj, image.GetNativeObjAddr())
	return retVal
}
func (rcvr *VideoCapture) Release() {
	VideoCaptureNative_release_0(rcvr.nativeObj)
	return
}
func (rcvr *VideoCapture) Retrieve(image *Mat, flag int) bool {
	retVal := VideoCaptureNative_retrieve_0(rcvr.nativeObj, image.GetNativeObjAddr(), flag)
	return retVal
}
func (rcvr *VideoCapture) Retrieve2(image *Mat) bool {
	retVal := VideoCaptureNative_retrieve_1(rcvr.nativeObj, image.GetNativeObjAddr())
	return retVal
}
func (rcvr *VideoCapture) Set(propId int, value float64) bool {
	retVal := VideoCaptureNative_set_0(rcvr.nativeObj, propId, value)
	return retVal
}
