package videoio

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

type VideoWriter struct {
	nativeObj int64
}

func NewVideoWriter(addr int64) (rcvr *VideoWriter) {
	rcvr = &VideoWriter{}
	rcvr.nativeObj = addr
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func NewVideoWriter2(filename string, apiPreference int, fourcc int, fps float64, frameSize *Size, isColor bool) (rcvr *VideoWriter) {
	rcvr = &VideoWriter{}
	rcvr.nativeObj = VideoWriterNative_VideoWriter_0(filename, apiPreference, fourcc, fps, frameSize.Width, frameSize.Height, isColor)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func NewVideoWriter3(filename string, apiPreference int, fourcc int, fps float64, frameSize *Size) (rcvr *VideoWriter) {
	rcvr = &VideoWriter{}
	rcvr.nativeObj = VideoWriterNative_VideoWriter_1(filename, apiPreference, fourcc, fps, frameSize.Width, frameSize.Height)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func NewVideoWriter4(filename string, fourcc int, fps float64, frameSize *Size, isColor bool) (rcvr *VideoWriter) {
	rcvr = &VideoWriter{}
	rcvr.nativeObj = VideoWriterNative_VideoWriter_2(filename, fourcc, fps, frameSize.Width, frameSize.Height, isColor)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
	return
}
func NewVideoWriter5(filename string, fourcc int, fps float64, frameSize *Size) (rcvr *VideoWriter) {
	rcvr = &VideoWriter{}
	rcvr.nativeObj = VideoWriterNative_VideoWriter_3(filename, fourcc, fps, frameSize.Width, frameSize.Height)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func NewVideoWriter6() (rcvr *VideoWriter) {
	rcvr = &VideoWriter{}
	rcvr.nativeObj = VideoWriterNative_VideoWriter_4()
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func (rcvr *VideoWriter) finalize() {
	VideoWriterNative_delete(rcvr.nativeObj)
}
func Fourcc(c1 rune, c2 rune, c3 rune, c4 rune) int {
	retVal := VideoWriterNative_fourcc_0(c1, c2, c3, c4)
	return retVal
}
func (rcvr *VideoWriter) Get(propId int) float64 {
	retVal := VideoWriterNative_get_0(rcvr.nativeObj, propId)
	return retVal
}
func (rcvr *VideoWriter) GetNativeObjAddr() int64 {
	return rcvr.nativeObj
}
func (rcvr *VideoWriter) IsOpened() bool {
	retVal := VideoWriterNative_isOpened_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *VideoWriter) Open(filename string, apiPreference int, fourcc int, fps float64, frameSize *Size, isColor bool) bool {
	retVal := VideoWriterNative_open_0(rcvr.nativeObj, filename, apiPreference, fourcc, fps, frameSize.Width, frameSize.Height, isColor)
	return retVal
}
func (rcvr *VideoWriter) Open2(filename string, apiPreference int, fourcc int, fps float64, frameSize *Size) bool {
	retVal := VideoWriterNative_open_1(rcvr.nativeObj, filename, apiPreference, fourcc, fps, frameSize.Width, frameSize.Height)
	return retVal
}
func (rcvr *VideoWriter) Open3(filename string, fourcc int, fps float64, frameSize *Size, isColor bool) bool {
	retVal := VideoWriterNative_open_2(rcvr.nativeObj, filename, fourcc, fps, frameSize.Width, frameSize.Height, isColor)
	return retVal
}
func (rcvr *VideoWriter) Open4(filename string, fourcc int, fps float64, frameSize *Size) bool {
	retVal := VideoWriterNative_open_3(rcvr.nativeObj, filename, fourcc, fps, frameSize.Width, frameSize.Height)
	return retVal
}
func (rcvr *VideoWriter) Release() {
	VideoWriterNative_release_0(rcvr.nativeObj)
	return
}
func (rcvr *VideoWriter) Set(propId int, value float64) bool {
	retVal := VideoWriterNative_set_0(rcvr.nativeObj, propId, value)
	return retVal
}
func (rcvr *VideoWriter) Write(image *Mat) {
	VideoWriterNative_write_0(rcvr.nativeObj, image.GetNativeObjAddr())
	return
}
