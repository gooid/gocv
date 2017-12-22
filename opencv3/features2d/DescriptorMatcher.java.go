package features2d

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

const DescriptorMatcherFLANNBASED = 1
const DescriptorMatcherBRUTEFORCE = 2
const DescriptorMatcherBRUTEFORCE_L1 = 3
const DescriptorMatcherBRUTEFORCE_HAMMING = 4
const DescriptorMatcherBRUTEFORCE_HAMMINGLUT = 5
const DescriptorMatcherBRUTEFORCE_SL2 = 6

type DescriptorMatcher struct {
	*Algorithm
}

func NewDescriptorMatcher(addr int64) (rcvr *DescriptorMatcher) {
	rcvr = &DescriptorMatcher{}
	rcvr.Algorithm = NewAlgorithm(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func (rcvr *DescriptorMatcher) Add(descriptors []*Mat) {
	descriptors_mat := ConvertersVectorMat(descriptors)
	DescriptorMatcherNative_add_0(rcvr.GetNativeObjAddr(), descriptors_mat.GetNativeObjAddr())
	return
}
func (rcvr *DescriptorMatcher) Clear() {
	DescriptorMatcherNative_clear_0(rcvr.GetNativeObjAddr())
	return
}
func (rcvr *DescriptorMatcher) Clone(emptyTrainData bool) *DescriptorMatcher {
	retVal := NewDescriptorMatcher(DescriptorMatcherNative_clone_0(rcvr.GetNativeObjAddr(), emptyTrainData))
	return retVal
}
func (rcvr *DescriptorMatcher) Clone2() *DescriptorMatcher {
	retVal := NewDescriptorMatcher(DescriptorMatcherNative_clone_1(rcvr.GetNativeObjAddr()))
	return retVal
}
func DescriptorMatcherCreate(descriptorMatcherType string) *DescriptorMatcher {
	retVal := NewDescriptorMatcher(DescriptorMatcherNative_create_0(descriptorMatcherType))
	return retVal
}
func DescriptorMatcherCreate2(matcherType int) *DescriptorMatcher {
	retVal := NewDescriptorMatcher(DescriptorMatcherNative_create_1(matcherType))
	return retVal
}
func (rcvr *DescriptorMatcher) Empty() bool {
	retVal := DescriptorMatcherNative_empty_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *DescriptorMatcher) finalize() {
	DescriptorMatcherNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *DescriptorMatcher) GetTrainDescriptors() []*Mat {
	retValMat := NewMat(DescriptorMatcherNative_getTrainDescriptors_0(rcvr.GetNativeObjAddr()))
	return ConvertersToVectorMat(retValMat)
}
func (rcvr *DescriptorMatcher) IsMaskSupported() bool {
	retVal := DescriptorMatcherNative_isMaskSupported_0(rcvr.GetNativeObjAddr())
	return retVal
}
func (rcvr *DescriptorMatcher) KnnMatch(queryDescriptors *Mat, trainDescriptors *Mat, k int, mask *Mat, compactResult bool) (matches []*MatOfDMatch) {
	matches_mat := NewMat2()
	DescriptorMatcherNative_knnMatch_0(rcvr.GetNativeObjAddr(), queryDescriptors.GetNativeObjAddr(), trainDescriptors.GetNativeObjAddr(), matches_mat.GetNativeObjAddr(), k, mask.GetNativeObjAddr(), compactResult)
	matches = ConvertersToVectorVectorDMatch(matches_mat)
	matches_mat.Release()
	return
}
func (rcvr *DescriptorMatcher) KnnMatch2(queryDescriptors *Mat, trainDescriptors *Mat, k int) (matches []*MatOfDMatch) {
	matches_mat := NewMat2()
	DescriptorMatcherNative_knnMatch_1(rcvr.GetNativeObjAddr(), queryDescriptors.GetNativeObjAddr(), trainDescriptors.GetNativeObjAddr(), matches_mat.GetNativeObjAddr(), k)
	matches = ConvertersToVectorVectorDMatch(matches_mat)
	matches_mat.Release()
	return
}
func (rcvr *DescriptorMatcher) KnnMatch3(queryDescriptors *Mat, k int, masks []*Mat, compactResult bool) (matches []*MatOfDMatch) {
	matches_mat := NewMat2()
	masks_mat := ConvertersVectorMat(masks)
	DescriptorMatcherNative_knnMatch_2(rcvr.GetNativeObjAddr(), queryDescriptors.GetNativeObjAddr(), matches_mat.GetNativeObjAddr(), k, masks_mat.GetNativeObjAddr(), compactResult)
	matches = ConvertersToVectorVectorDMatch(matches_mat)
	matches_mat.Release()
	return
}
func (rcvr *DescriptorMatcher) KnnMatch4(queryDescriptors *Mat, k int) (matches []*MatOfDMatch) {
	matches_mat := NewMat2()
	DescriptorMatcherNative_knnMatch_3(rcvr.GetNativeObjAddr(), queryDescriptors.GetNativeObjAddr(), matches_mat.GetNativeObjAddr(), k)
	matches = ConvertersToVectorVectorDMatch(matches_mat)
	matches_mat.Release()
	return
}
func (rcvr *DescriptorMatcher) Match(queryDescriptors *Mat, trainDescriptors *Mat, matches *MatOfDMatch, mask *Mat) {
	matches_mat := matches
	DescriptorMatcherNative_match_0(rcvr.GetNativeObjAddr(), queryDescriptors.GetNativeObjAddr(), trainDescriptors.GetNativeObjAddr(), matches_mat.GetNativeObjAddr(), mask.GetNativeObjAddr())
	return
}
func (rcvr *DescriptorMatcher) Match2(queryDescriptors *Mat, trainDescriptors *Mat, matches *MatOfDMatch) {
	matches_mat := matches
	DescriptorMatcherNative_match_1(rcvr.GetNativeObjAddr(), queryDescriptors.GetNativeObjAddr(), trainDescriptors.GetNativeObjAddr(), matches_mat.GetNativeObjAddr())
	return
}
func (rcvr *DescriptorMatcher) Match3(queryDescriptors *Mat, matches *MatOfDMatch, masks []*Mat) {
	matches_mat := matches
	masks_mat := ConvertersVectorMat(masks)
	DescriptorMatcherNative_match_2(rcvr.GetNativeObjAddr(), queryDescriptors.GetNativeObjAddr(), matches_mat.GetNativeObjAddr(), masks_mat.GetNativeObjAddr())
	return
}
func (rcvr *DescriptorMatcher) Match4(queryDescriptors *Mat, matches *MatOfDMatch) {
	matches_mat := matches
	DescriptorMatcherNative_match_3(rcvr.GetNativeObjAddr(), queryDescriptors.GetNativeObjAddr(), matches_mat.GetNativeObjAddr())
	return
}
func (rcvr *DescriptorMatcher) RadiusMatch(queryDescriptors *Mat, trainDescriptors *Mat, maxDistance float32, mask *Mat, compactResult bool) (matches []*MatOfDMatch) {
	matches_mat := NewMat2()
	DescriptorMatcherNative_radiusMatch_0(rcvr.GetNativeObjAddr(), queryDescriptors.GetNativeObjAddr(), trainDescriptors.GetNativeObjAddr(), matches_mat.GetNativeObjAddr(), maxDistance, mask.GetNativeObjAddr(), compactResult)
	matches = ConvertersToVectorVectorDMatch(matches_mat)
	matches_mat.Release()
	return
}
func (rcvr *DescriptorMatcher) RadiusMatch2(queryDescriptors *Mat, trainDescriptors *Mat, maxDistance float32) (matches []*MatOfDMatch) {
	matches_mat := NewMat2()
	DescriptorMatcherNative_radiusMatch_1(rcvr.GetNativeObjAddr(), queryDescriptors.GetNativeObjAddr(), trainDescriptors.GetNativeObjAddr(), matches_mat.GetNativeObjAddr(), maxDistance)
	matches = ConvertersToVectorVectorDMatch(matches_mat)
	matches_mat.Release()
	return
}
func (rcvr *DescriptorMatcher) RadiusMatch3(queryDescriptors *Mat, maxDistance float32, masks []*Mat, compactResult bool) (matches []*MatOfDMatch) {
	matches_mat := NewMat2()
	masks_mat := ConvertersVectorMat(masks)
	DescriptorMatcherNative_radiusMatch_2(rcvr.GetNativeObjAddr(), queryDescriptors.GetNativeObjAddr(), matches_mat.GetNativeObjAddr(), maxDistance, masks_mat.GetNativeObjAddr(), compactResult)
	matches = ConvertersToVectorVectorDMatch(matches_mat)
	matches_mat.Release()
	return
}
func (rcvr *DescriptorMatcher) RadiusMatch4(queryDescriptors *Mat, maxDistance float32) (matches []*MatOfDMatch) {
	matches_mat := NewMat2()
	DescriptorMatcherNative_radiusMatch_3(rcvr.GetNativeObjAddr(), queryDescriptors.GetNativeObjAddr(), matches_mat.GetNativeObjAddr(), maxDistance)
	matches = ConvertersToVectorVectorDMatch(matches_mat)
	matches_mat.Release()
	return
}
func (rcvr *DescriptorMatcher) Read(fileName string) {
	DescriptorMatcherNative_read_0(rcvr.GetNativeObjAddr(), fileName)
	return
}
func (rcvr *DescriptorMatcher) Train() {
	DescriptorMatcherNative_train_0(rcvr.GetNativeObjAddr())
	return
}
func (rcvr *DescriptorMatcher) Write(fileName string) {
	DescriptorMatcherNative_write_0(rcvr.GetNativeObjAddr(), fileName)
	return
}
