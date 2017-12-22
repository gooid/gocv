package features2d

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

const KAZEDIFF_PM_G1 = 0
const KAZEDIFF_PM_G2 = 1
const KAZEDIFF_WEICKERT = 2
const KAZEDIFF_CHARBONNIER = 3

type KAZE struct {
	*Feature2D
}

func NewKAZE(addr int64) (rcvr *KAZE) {
	rcvr = &KAZE{}
	rcvr.Feature2D = NewFeature2D(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func KAZECreate(extended bool, upright bool, threshold float32, nOctaves int, nOctaveLayers int, diffusivity int) *KAZE {
	retVal := NewKAZE(KAZENative_create_0(extended, upright, threshold, nOctaves, nOctaveLayers, diffusivity))
	return retVal
}
func KAZECreate2() *KAZE {
	retVal := NewKAZE(KAZENative_create_1())
	return retVal
}
func (rcvr *KAZE) finalize() {
	KAZENative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *KAZE) GetDefaultName() string {
	retVal := KAZENative_getDefaultName_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *KAZE) GetDiffusivity() int {
	retVal := KAZENative_getDiffusivity_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *KAZE) GetExtended() bool {
	retVal := KAZENative_getExtended_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *KAZE) GetNOctaveLayers() int {
	retVal := KAZENative_getNOctaveLayers_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *KAZE) GetNOctaves() int {
	retVal := KAZENative_getNOctaves_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *KAZE) GetThreshold() float64 {
	retVal := KAZENative_getThreshold_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *KAZE) GetUpright() bool {
	retVal := KAZENative_getUpright_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *KAZE) SetDiffusivity(diff int) {
	KAZENative_setDiffusivity_0(rcvr.GetNativeObjAddr(), diff)
	return
}
func (rcvr *KAZE) SetExtended(extended bool) {
	KAZENative_setExtended_0(rcvr.GetNativeObjAddr(), extended)
	return
}
func (rcvr *KAZE) SetNOctaveLayers(octaveLayers int) {
	KAZENative_setNOctaveLayers_0(rcvr.GetNativeObjAddr(), octaveLayers)
	return
}
func (rcvr *KAZE) SetNOctaves(octaves int) {
	KAZENative_setNOctaves_0(rcvr.GetNativeObjAddr(), octaves)
	return
}
func (rcvr *KAZE) SetThreshold(threshold float64) {
	KAZENative_setThreshold_0(rcvr.GetNativeObjAddr(), threshold)
	return
}
func (rcvr *KAZE) SetUpright(upright bool) {
	KAZENative_setUpright_0(rcvr.GetNativeObjAddr(), upright)
	return
}
