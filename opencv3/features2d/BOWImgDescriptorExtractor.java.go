package features2d

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

type BOWImgDescriptorExtractor struct {
	nativeObj int64
}

func NewBOWImgDescriptorExtractor(addr int64) (rcvr *BOWImgDescriptorExtractor) {
	rcvr = &BOWImgDescriptorExtractor{}
	rcvr.nativeObj = addr
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func (rcvr *BOWImgDescriptorExtractor) Compute(image *Mat, keypoints *MatOfKeyPoint, imgDescriptor *Mat) {
	keypoints_mat := keypoints
	BOWImgDescriptorExtractorNative_compute_0(rcvr.nativeObj, image.GetNativeObjAddr(), keypoints_mat.GetNativeObjAddr(), imgDescriptor.GetNativeObjAddr())
	return
}
func (rcvr *BOWImgDescriptorExtractor) DescriptorSize() int {
	retVal := BOWImgDescriptorExtractorNative_descriptorSize_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *BOWImgDescriptorExtractor) DescriptorType() int {
	retVal := BOWImgDescriptorExtractorNative_descriptorType_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *BOWImgDescriptorExtractor) finalize() {
	BOWImgDescriptorExtractorNative_delete(rcvr.nativeObj)
}
func (rcvr *BOWImgDescriptorExtractor) GetNativeObjAddr() int64 {
	return rcvr.nativeObj
}
func (rcvr *BOWImgDescriptorExtractor) GetVocabulary() *Mat {
	retVal := NewMat(BOWImgDescriptorExtractorNative_getVocabulary_0(rcvr.nativeObj))
	return retVal
}
func (rcvr *BOWImgDescriptorExtractor) SetVocabulary(vocabulary *Mat) {
	BOWImgDescriptorExtractorNative_setVocabulary_0(rcvr.nativeObj, vocabulary.GetNativeObjAddr())
	return
}
