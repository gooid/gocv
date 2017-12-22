package features2d

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

type Params struct {
	nativeObj int64
}

func NewParams(addr int64) (rcvr *Params) {
	rcvr = &Params{}
	rcvr.nativeObj = addr
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func NewParams2() (rcvr *Params) {
	rcvr = &Params{}
	rcvr.nativeObj = ParamsNative_Params_0()
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
	return
}
func (rcvr *Params) finalize() {
	ParamsNative_delete(rcvr.nativeObj)
}
func (rcvr *Params) GetNativeObjAddr() int64 {
	return rcvr.nativeObj
}
func (rcvr *Params) Get_filterByArea() bool {
	retVal := ParamsNative_get_filterByArea_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *Params) Get_filterByCircularity() bool {
	retVal := ParamsNative_get_filterByCircularity_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *Params) Get_filterByColor() bool {
	retVal := ParamsNative_get_filterByColor_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *Params) Get_filterByConvexity() bool {
	retVal := ParamsNative_get_filterByConvexity_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *Params) Get_filterByInertia() bool {
	retVal := ParamsNative_get_filterByInertia_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *Params) Get_maxArea() float32 {
	retVal := ParamsNative_get_maxArea_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *Params) Get_maxCircularity() float32 {
	retVal := ParamsNative_get_maxCircularity_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *Params) Get_maxConvexity() float32 {
	retVal := ParamsNative_get_maxConvexity_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *Params) Get_maxInertiaRatio() float32 {
	retVal := ParamsNative_get_maxInertiaRatio_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *Params) Get_maxThreshold() float32 {
	retVal := ParamsNative_get_maxThreshold_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *Params) Get_minArea() float32 {
	retVal := ParamsNative_get_minArea_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *Params) Get_minCircularity() float32 {
	retVal := ParamsNative_get_minCircularity_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *Params) Get_minConvexity() float32 {
	retVal := ParamsNative_get_minConvexity_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *Params) Get_minDistBetweenBlobs() float32 {
	retVal := ParamsNative_get_minDistBetweenBlobs_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *Params) Get_minInertiaRatio() float32 {
	retVal := ParamsNative_get_minInertiaRatio_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *Params) Get_minRepeatability() int64 {
	retVal := ParamsNative_get_minRepeatability_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *Params) Get_minThreshold() float32 {
	retVal := ParamsNative_get_minThreshold_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *Params) Get_thresholdStep() float32 {
	retVal := ParamsNative_get_thresholdStep_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *Params) Set_filterByArea(filterByArea bool) {
	ParamsNative_set_filterByArea_0(rcvr.nativeObj, filterByArea)
	return
}
func (rcvr *Params) Set_filterByCircularity(filterByCircularity bool) {
	ParamsNative_set_filterByCircularity_0(rcvr.nativeObj, filterByCircularity)
	return
}
func (rcvr *Params) Set_filterByColor(filterByColor bool) {
	ParamsNative_set_filterByColor_0(rcvr.nativeObj, filterByColor)
	return
}
func (rcvr *Params) Set_filterByConvexity(filterByConvexity bool) {
	ParamsNative_set_filterByConvexity_0(rcvr.nativeObj, filterByConvexity)
	return
}
func (rcvr *Params) Set_filterByInertia(filterByInertia bool) {
	ParamsNative_set_filterByInertia_0(rcvr.nativeObj, filterByInertia)
	return
}
func (rcvr *Params) Set_maxArea(maxArea float32) {
	ParamsNative_set_maxArea_0(rcvr.nativeObj, maxArea)
	return
}
func (rcvr *Params) Set_maxCircularity(maxCircularity float32) {
	ParamsNative_set_maxCircularity_0(rcvr.nativeObj, maxCircularity)
	return
}
func (rcvr *Params) Set_maxConvexity(maxConvexity float32) {
	ParamsNative_set_maxConvexity_0(rcvr.nativeObj, maxConvexity)
	return
}
func (rcvr *Params) Set_maxInertiaRatio(maxInertiaRatio float32) {
	ParamsNative_set_maxInertiaRatio_0(rcvr.nativeObj, maxInertiaRatio)
	return
}
func (rcvr *Params) Set_maxThreshold(maxThreshold float32) {
	ParamsNative_set_maxThreshold_0(rcvr.nativeObj, maxThreshold)
	return
}
func (rcvr *Params) Set_minArea(minArea float32) {
	ParamsNative_set_minArea_0(rcvr.nativeObj, minArea)
	return
}
func (rcvr *Params) Set_minCircularity(minCircularity float32) {
	ParamsNative_set_minCircularity_0(rcvr.nativeObj, minCircularity)
	return
}
func (rcvr *Params) Set_minConvexity(minConvexity float32) {
	ParamsNative_set_minConvexity_0(rcvr.nativeObj, minConvexity)
	return
}
func (rcvr *Params) Set_minDistBetweenBlobs(minDistBetweenBlobs float32) {
	ParamsNative_set_minDistBetweenBlobs_0(rcvr.nativeObj, minDistBetweenBlobs)
	return
}
func (rcvr *Params) Set_minInertiaRatio(minInertiaRatio float32) {
	ParamsNative_set_minInertiaRatio_0(rcvr.nativeObj, minInertiaRatio)
	return
}
func (rcvr *Params) Set_minRepeatability(minRepeatability int64) {
	ParamsNative_set_minRepeatability_0(rcvr.nativeObj, minRepeatability)
	return
}
func (rcvr *Params) Set_minThreshold(minThreshold float32) {
	ParamsNative_set_minThreshold_0(rcvr.nativeObj, minThreshold)
	return
}
func (rcvr *Params) Set_thresholdStep(thresholdStep float32) {
	ParamsNative_set_thresholdStep_0(rcvr.nativeObj, thresholdStep)
	return
}
