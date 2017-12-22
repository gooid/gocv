package features2d

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

const oPPONENTEXTRACTORDescriptorExtractor = 1000
const DescriptorExtractorSIFT = 1
const DescriptorExtractorSURF = 2
const DescriptorExtractorORB = 3
const DescriptorExtractorBRIEF = 4
const DescriptorExtractorBRISK = 5
const DescriptorExtractorFREAK = 6
const DescriptorExtractorAKAZE = 7

var DescriptorExtractorOPPONENT_SIFT = oPPONENTEXTRACTORDescriptorExtractor + DescriptorExtractorSIFT
var DescriptorExtractorOPPONENT_SURF = oPPONENTEXTRACTORDescriptorExtractor + DescriptorExtractorSURF
var DescriptorExtractorOPPONENT_ORB = oPPONENTEXTRACTORDescriptorExtractor + DescriptorExtractorORB
var DescriptorExtractorOPPONENT_BRIEF = oPPONENTEXTRACTORDescriptorExtractor + DescriptorExtractorBRIEF
var DescriptorExtractorOPPONENT_BRISK = oPPONENTEXTRACTORDescriptorExtractor + DescriptorExtractorBRISK
var DescriptorExtractorOPPONENT_FREAK = oPPONENTEXTRACTORDescriptorExtractor + DescriptorExtractorFREAK
var DescriptorExtractorOPPONENT_AKAZE = oPPONENTEXTRACTORDescriptorExtractor + DescriptorExtractorAKAZE

type DescriptorExtractor struct {
	nativeObj int64
}

func NewDescriptorExtractor(addr int64) (rcvr *DescriptorExtractor) {
	rcvr = &DescriptorExtractor{}
	rcvr.nativeObj = addr
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func (rcvr *DescriptorExtractor) Compute(image *Mat, keypoints *MatOfKeyPoint, descriptors *Mat) {
	keypoints_mat := keypoints
	DescriptorExtractorNative_compute_0(rcvr.nativeObj, image.GetNativeObjAddr(), keypoints_mat.GetNativeObjAddr(), descriptors.GetNativeObjAddr())
	return
}
func (rcvr *DescriptorExtractor) Compute2(images []*Mat, keypoints []*MatOfKeyPoint) (keypointsout []*MatOfKeyPoint, descriptors []*Mat) {
	images_mat := ConvertersVectorMat(images)
	keypoints_mat := ConvertersVectorVectorKeyPoint(keypoints)
	descriptors_mat := NewMat2()
	DescriptorExtractorNative_compute_1(rcvr.nativeObj, images_mat.GetNativeObjAddr(), keypoints_mat.GetNativeObjAddr(), descriptors_mat.GetNativeObjAddr())
	keypointsout = ConvertersToVectorVectorKeyPoint(keypoints_mat)
	keypoints_mat.Release()
	descriptors = ConvertersToVectorMat(descriptors_mat)
	descriptors_mat.Release()
	return
}
func DescriptorExtractorCreate(extractorType int) *DescriptorExtractor {
	retVal := NewDescriptorExtractor(DescriptorExtractorNative_create_0(extractorType))
	return retVal
}
func (rcvr *DescriptorExtractor) DescriptorSize() int {
	retVal := DescriptorExtractorNative_descriptorSize_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *DescriptorExtractor) DescriptorType() int {
	retVal := DescriptorExtractorNative_descriptorType_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *DescriptorExtractor) Empty() bool {
	retVal := DescriptorExtractorNative_empty_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *DescriptorExtractor) finalize() {
	DescriptorExtractorNative_delete(rcvr.nativeObj)
}
func (rcvr *DescriptorExtractor) GetNativeObjAddr() int64 {
	return rcvr.nativeObj
}
func (rcvr *DescriptorExtractor) Read(fileName string) {
	DescriptorExtractorNative_read_0(rcvr.nativeObj, fileName)
	return
}
func (rcvr *DescriptorExtractor) Write(fileName string) {
	DescriptorExtractorNative_write_0(rcvr.nativeObj, fileName)
	return
}
