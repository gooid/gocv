package features2d

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

type MSER struct {
	*Feature2D
}

func NewMSER(addr int64) (rcvr *MSER) {
	rcvr = &MSER{}
	rcvr.Feature2D = NewFeature2D(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func MSERCreate(_delta int, _min_area int, _max_area int, _max_variation float64, _min_diversity float64, _max_evolution int, _area_threshold float64, _min_margin float64, _edge_blur_size int) *MSER {
	retVal := NewMSER(MSERNative_create_0(_delta, _min_area, _max_area, _max_variation, _min_diversity, _max_evolution, _area_threshold, _min_margin, _edge_blur_size))
	return retVal
}
func MSERCreate2() *MSER {
	retVal := NewMSER(MSERNative_create_1())
	return retVal
}
func (rcvr *MSER) DetectRegions(image *Mat, bboxes *MatOfRect) (msers []*MatOfPoint) {
	msers_mat := NewMat2()
	bboxes_mat := bboxes
	MSERNative_detectRegions_0(rcvr.GetNativeObjAddr(), image.GetNativeObjAddr(), msers_mat.GetNativeObjAddr(), bboxes_mat.GetNativeObjAddr())
	msers = ConvertersToVectorVectorPoint(msers_mat)
	msers_mat.Release()
	return
}
func (rcvr *MSER) finalize() {
	MSERNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *MSER) GetDefaultName() string {
	retVal := MSERNative_getDefaultName_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *MSER) GetDelta() int {
	retVal := MSERNative_getDelta_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *MSER) GetMaxArea() int {
	retVal := MSERNative_getMaxArea_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *MSER) GetMinArea() int {
	retVal := MSERNative_getMinArea_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *MSER) GetPass2Only() bool {
	retVal := MSERNative_getPass2Only_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *MSER) SetDelta(delta int) {
	MSERNative_setDelta_0(rcvr.GetNativeObjAddr(), delta)
	return
}
func (rcvr *MSER) SetMaxArea(maxArea int) {
	MSERNative_setMaxArea_0(rcvr.GetNativeObjAddr(), maxArea)
	return
}
func (rcvr *MSER) SetMinArea(minArea int) {
	MSERNative_setMinArea_0(rcvr.GetNativeObjAddr(), minArea)
	return
}
func (rcvr *MSER) SetPass2Only(f bool) {
	MSERNative_setPass2Only_0(rcvr.GetNativeObjAddr(), f)
	return
}
