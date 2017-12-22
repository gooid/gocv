package features2d

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

const AKAZEDESCRIPTOR_KAZE_UPRIGHT = 2
const AKAZEDESCRIPTOR_KAZE = 3
const AKAZEDESCRIPTOR_MLDB_UPRIGHT = 4
const AKAZEDESCRIPTOR_MLDB = 5

type AKAZE struct {
	*Feature2D
}

func NewAKAZE(addr int64) (rcvr *AKAZE) {
	rcvr = &AKAZE{}
	rcvr.Feature2D = NewFeature2D(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func AKAZECreate(descriptor_type int, descriptor_size int, descriptor_channels int, threshold float32, nOctaves int, nOctaveLayers int, diffusivity int) *AKAZE {
	retVal := NewAKAZE(AKAZENative_create_0(descriptor_type, descriptor_size, descriptor_channels, threshold, nOctaves, nOctaveLayers, diffusivity))
	return retVal
}
func AKAZECreate2() *AKAZE {
	retVal := NewAKAZE(AKAZENative_create_1())
	return retVal
}
func (rcvr *AKAZE) finalize() {
	AKAZENative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *AKAZE) GetDefaultName() string {
	retVal := AKAZENative_getDefaultName_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *AKAZE) GetDescriptorChannels() int {
	retVal := AKAZENative_getDescriptorChannels_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *AKAZE) GetDescriptorSize() int {
	retVal := AKAZENative_getDescriptorSize_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *AKAZE) GetDescriptorType() int {
	retVal := AKAZENative_getDescriptorType_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *AKAZE) GetDiffusivity() int {
	retVal := AKAZENative_getDiffusivity_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *AKAZE) GetNOctaveLayers() int {
	retVal := AKAZENative_getNOctaveLayers_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *AKAZE) GetNOctaves() int {
	retVal := AKAZENative_getNOctaves_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *AKAZE) GetThreshold() float64 {
	retVal := AKAZENative_getThreshold_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *AKAZE) SetDescriptorChannels(dch int) {
	AKAZENative_setDescriptorChannels_0(rcvr.GetNativeObjAddr(), dch)
	return
}
func (rcvr *AKAZE) SetDescriptorSize(dsize int) {
	AKAZENative_setDescriptorSize_0(rcvr.GetNativeObjAddr(), dsize)
	return
}
func (rcvr *AKAZE) SetDescriptorType(dtype int) {
	AKAZENative_setDescriptorType_0(rcvr.GetNativeObjAddr(), dtype)
	return
}
func (rcvr *AKAZE) SetDiffusivity(diff int) {
	AKAZENative_setDiffusivity_0(rcvr.GetNativeObjAddr(), diff)
	return
}
func (rcvr *AKAZE) SetNOctaveLayers(octaveLayers int) {
	AKAZENative_setNOctaveLayers_0(rcvr.GetNativeObjAddr(), octaveLayers)
	return
}
func (rcvr *AKAZE) SetNOctaves(octaves int) {
	AKAZENative_setNOctaves_0(rcvr.GetNativeObjAddr(), octaves)
	return
}
func (rcvr *AKAZE) SetThreshold(threshold float64) {
	AKAZENative_setThreshold_0(rcvr.GetNativeObjAddr(), threshold)
	return
}
