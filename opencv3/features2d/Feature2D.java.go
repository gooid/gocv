package features2d

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

type Feature2D struct {
	*Algorithm
}

func NewFeature2D(addr int64) (rcvr *Feature2D) {
	rcvr = &Feature2D{}
	rcvr.Algorithm = NewAlgorithm(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func (rcvr *Feature2D) Compute(image *Mat, keypoints *MatOfKeyPoint, descriptors *Mat) {
	keypoints_mat := keypoints
	Feature2DNative_compute_0(rcvr.GetNativeObjAddr(), image.GetNativeObjAddr(), keypoints_mat.GetNativeObjAddr(), descriptors.GetNativeObjAddr())
	return
}
func (rcvr *Feature2D) Compute2(images []*Mat, keypoints []*MatOfKeyPoint) (keypointsout []*MatOfKeyPoint, descriptors []*Mat) {
	images_mat := ConvertersVectorMat(images)
	keypoints_mat := ConvertersVectorVectorKeyPoint(keypoints)
	descriptors_mat := NewMat2()
	Feature2DNative_compute_1(rcvr.GetNativeObjAddr(), images_mat.GetNativeObjAddr(), keypoints_mat.GetNativeObjAddr(), descriptors_mat.GetNativeObjAddr())
	keypointsout = ConvertersToVectorVectorKeyPoint(keypoints_mat)
	keypoints_mat.Release()
	descriptors = ConvertersToVectorMat(descriptors_mat)
	descriptors_mat.Release()
	return
}
func (rcvr *Feature2D) DefaultNorm() int {
	retVal := Feature2DNative_defaultNorm_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *Feature2D) DescriptorSize() int {
	retVal := Feature2DNative_descriptorSize_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *Feature2D) DescriptorType() int {
	retVal := Feature2DNative_descriptorType_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *Feature2D) Detect(image *Mat, keypoints *MatOfKeyPoint, mask *Mat) {
	keypoints_mat := keypoints
	Feature2DNative_detect_0(rcvr.GetNativeObjAddr(), image.GetNativeObjAddr(), keypoints_mat.GetNativeObjAddr(), mask.GetNativeObjAddr())
	return
}
func (rcvr *Feature2D) Detect2(image *Mat, keypoints *MatOfKeyPoint) {
	keypoints_mat := keypoints
	Feature2DNative_detect_1(rcvr.GetNativeObjAddr(), image.GetNativeObjAddr(), keypoints_mat.GetNativeObjAddr())
	return
}
func (rcvr *Feature2D) Detect3(images []*Mat, masks []*Mat) (keypoints []*MatOfKeyPoint) {
	images_mat := ConvertersVectorMat(images)
	keypoints_mat := NewMat2()
	masks_mat := ConvertersVectorMat(masks)
	Feature2DNative_detect_2(rcvr.GetNativeObjAddr(), images_mat.GetNativeObjAddr(), keypoints_mat.GetNativeObjAddr(), masks_mat.GetNativeObjAddr())
	keypoints = ConvertersToVectorVectorKeyPoint(keypoints_mat)
	keypoints_mat.Release()
	return
}
func (rcvr *Feature2D) Detect4(images []*Mat) (keypoints []*MatOfKeyPoint) {
	images_mat := ConvertersVectorMat(images)
	keypoints_mat := NewMat2()
	Feature2DNative_detect_3(rcvr.GetNativeObjAddr(), images_mat.GetNativeObjAddr(), keypoints_mat.GetNativeObjAddr())
	keypoints = ConvertersToVectorVectorKeyPoint(keypoints_mat)
	keypoints_mat.Release()
	return
}
func (rcvr *Feature2D) DetectAndCompute(image *Mat, mask *Mat, keypoints *MatOfKeyPoint, descriptors *Mat, useProvidedKeypoints bool) {
	keypoints_mat := keypoints
	Feature2DNative_detectAndCompute_0(rcvr.GetNativeObjAddr(), image.GetNativeObjAddr(), mask.GetNativeObjAddr(), keypoints_mat.GetNativeObjAddr(), descriptors.GetNativeObjAddr(), useProvidedKeypoints)
	return
}
func (rcvr *Feature2D) DetectAndCompute2(image *Mat, mask *Mat, keypoints *MatOfKeyPoint, descriptors *Mat) {
	keypoints_mat := keypoints
	Feature2DNative_detectAndCompute_1(rcvr.GetNativeObjAddr(), image.GetNativeObjAddr(), mask.GetNativeObjAddr(), keypoints_mat.GetNativeObjAddr(), descriptors.GetNativeObjAddr())
	return
}
func (rcvr *Feature2D) Empty() bool {
	retVal := Feature2DNative_empty_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *Feature2D) finalize() {
	Feature2DNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *Feature2D) GetDefaultName() string {
	retVal := Feature2DNative_getDefaultName_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *Feature2D) Read(fileName string) {
	Feature2DNative_read_0(rcvr.GetNativeObjAddr(), fileName)
	return
}
func (rcvr *Feature2D) Write(fileName string) {
	Feature2DNative_write_0(rcvr.GetNativeObjAddr(), fileName)
	return
}
