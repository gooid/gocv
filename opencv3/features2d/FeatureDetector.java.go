package features2d

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

const gRIDDETECTORFeatureDetector = 1000
const pYRAMIDDETECTORFeatureDetector = 2000
const dYNAMICDETECTORFeatureDetector = 3000
const FeatureDetectorFAST = 1
const FeatureDetectorSTAR = 2
const FeatureDetectorSIFT = 3
const FeatureDetectorSURF = 4
const FeatureDetectorORB = 5
const FeatureDetectorMSER = 6
const FeatureDetectorGFTT = 7
const FeatureDetectorHARRIS = 8
const FeatureDetectorSIMPLEBLOB = 9
const FeatureDetectorDENSE = 10
const FeatureDetectorBRISK = 11
const FeatureDetectorAKAZE = 12

var FeatureDetectorGRID_FAST = gRIDDETECTORFeatureDetector + FeatureDetectorFAST
var FeatureDetectorGRID_STAR = gRIDDETECTORFeatureDetector + FeatureDetectorSTAR
var FeatureDetectorGRID_SIFT = gRIDDETECTORFeatureDetector + FeatureDetectorSIFT
var FeatureDetectorGRID_SURF = gRIDDETECTORFeatureDetector + FeatureDetectorSURF
var FeatureDetectorGRID_ORB = gRIDDETECTORFeatureDetector + FeatureDetectorORB
var FeatureDetectorGRID_MSER = gRIDDETECTORFeatureDetector + FeatureDetectorMSER
var FeatureDetectorGRID_GFTT = gRIDDETECTORFeatureDetector + FeatureDetectorGFTT
var FeatureDetectorGRID_HARRIS = gRIDDETECTORFeatureDetector + FeatureDetectorHARRIS
var FeatureDetectorGRID_SIMPLEBLOB = gRIDDETECTORFeatureDetector + FeatureDetectorSIMPLEBLOB
var FeatureDetectorGRID_DENSE = gRIDDETECTORFeatureDetector + FeatureDetectorDENSE
var FeatureDetectorGRID_BRISK = gRIDDETECTORFeatureDetector + FeatureDetectorBRISK
var FeatureDetectorGRID_AKAZE = gRIDDETECTORFeatureDetector + FeatureDetectorAKAZE
var FeatureDetectorPYRAMID_FAST = pYRAMIDDETECTORFeatureDetector + FeatureDetectorFAST
var FeatureDetectorPYRAMID_STAR = pYRAMIDDETECTORFeatureDetector + FeatureDetectorSTAR
var FeatureDetectorPYRAMID_SIFT = pYRAMIDDETECTORFeatureDetector + FeatureDetectorSIFT
var FeatureDetectorPYRAMID_SURF = pYRAMIDDETECTORFeatureDetector + FeatureDetectorSURF
var FeatureDetectorPYRAMID_ORB = pYRAMIDDETECTORFeatureDetector + FeatureDetectorORB
var FeatureDetectorPYRAMID_MSER = pYRAMIDDETECTORFeatureDetector + FeatureDetectorMSER
var FeatureDetectorPYRAMID_GFTT = pYRAMIDDETECTORFeatureDetector + FeatureDetectorGFTT
var FeatureDetectorPYRAMID_HARRIS = pYRAMIDDETECTORFeatureDetector + FeatureDetectorHARRIS
var FeatureDetectorPYRAMID_SIMPLEBLOB = pYRAMIDDETECTORFeatureDetector + FeatureDetectorSIMPLEBLOB
var FeatureDetectorPYRAMID_DENSE = pYRAMIDDETECTORFeatureDetector + FeatureDetectorDENSE
var FeatureDetectorPYRAMID_BRISK = pYRAMIDDETECTORFeatureDetector + FeatureDetectorBRISK
var FeatureDetectorPYRAMID_AKAZE = pYRAMIDDETECTORFeatureDetector + FeatureDetectorAKAZE
var FeatureDetectorDYNAMIC_FAST = dYNAMICDETECTORFeatureDetector + FeatureDetectorFAST
var FeatureDetectorDYNAMIC_STAR = dYNAMICDETECTORFeatureDetector + FeatureDetectorSTAR
var FeatureDetectorDYNAMIC_SIFT = dYNAMICDETECTORFeatureDetector + FeatureDetectorSIFT
var FeatureDetectorDYNAMIC_SURF = dYNAMICDETECTORFeatureDetector + FeatureDetectorSURF
var FeatureDetectorDYNAMIC_ORB = dYNAMICDETECTORFeatureDetector + FeatureDetectorORB
var FeatureDetectorDYNAMIC_MSER = dYNAMICDETECTORFeatureDetector + FeatureDetectorMSER
var FeatureDetectorDYNAMIC_GFTT = dYNAMICDETECTORFeatureDetector + FeatureDetectorGFTT
var FeatureDetectorDYNAMIC_HARRIS = dYNAMICDETECTORFeatureDetector + FeatureDetectorHARRIS
var FeatureDetectorDYNAMIC_SIMPLEBLOB = dYNAMICDETECTORFeatureDetector + FeatureDetectorSIMPLEBLOB
var FeatureDetectorDYNAMIC_DENSE = dYNAMICDETECTORFeatureDetector + FeatureDetectorDENSE
var FeatureDetectorDYNAMIC_BRISK = dYNAMICDETECTORFeatureDetector + FeatureDetectorBRISK
var FeatureDetectorDYNAMIC_AKAZE = dYNAMICDETECTORFeatureDetector + FeatureDetectorAKAZE

type FeatureDetector struct {
	nativeObj int64
}

func NewFeatureDetector(addr int64) (rcvr *FeatureDetector) {
	rcvr = &FeatureDetector{}
	rcvr.nativeObj = addr
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func FeatureDetectorCreate(detectorType int) *FeatureDetector {
	retVal := NewFeatureDetector(FeatureDetectorNative_create_0(detectorType))
	return retVal
}
func (rcvr *FeatureDetector) Detect(image *Mat, keypoints *MatOfKeyPoint, mask *Mat) {
	keypoints_mat := keypoints
	FeatureDetectorNative_detect_0(rcvr.nativeObj, image.GetNativeObjAddr(), keypoints_mat.GetNativeObjAddr(), mask.GetNativeObjAddr())
	return
}
func (rcvr *FeatureDetector) Detect2(image *Mat, keypoints *MatOfKeyPoint) {
	keypoints_mat := keypoints
	FeatureDetectorNative_detect_1(rcvr.nativeObj, image.GetNativeObjAddr(), keypoints_mat.GetNativeObjAddr())
	return
}
func (rcvr *FeatureDetector) Detect3(images []*Mat, masks []*Mat) (keypoints []*MatOfKeyPoint) {
	images_mat := ConvertersVectorMat(images)
	keypoints_mat := NewMat2()
	masks_mat := ConvertersVectorMat(masks)
	FeatureDetectorNative_detect_2(rcvr.nativeObj, images_mat.GetNativeObjAddr(), keypoints_mat.GetNativeObjAddr(), masks_mat.GetNativeObjAddr())
	keypoints = ConvertersToVectorVectorKeyPoint(keypoints_mat)
	keypoints_mat.Release()
	return
}
func (rcvr *FeatureDetector) Detect4(images []*Mat) (keypoints []*MatOfKeyPoint) {
	images_mat := ConvertersVectorMat(images)
	keypoints_mat := NewMat2()
	FeatureDetectorNative_detect_3(rcvr.nativeObj, images_mat.GetNativeObjAddr(), keypoints_mat.GetNativeObjAddr())
	keypoints = ConvertersToVectorVectorKeyPoint(keypoints_mat)
	keypoints_mat.Release()
	return
}
func (rcvr *FeatureDetector) Empty() bool {
	retVal := FeatureDetectorNative_empty_0(rcvr.nativeObj)
	return retVal
}
func (rcvr *FeatureDetector) finalize() {
	FeatureDetectorNative_delete(rcvr.nativeObj)
}
func (rcvr *FeatureDetector) GetNativeObjAddr() int64 {
	return rcvr.nativeObj
}
func (rcvr *FeatureDetector) Read(fileName string) {
	FeatureDetectorNative_read_0(rcvr.nativeObj, fileName)
	return
}
func (rcvr *FeatureDetector) Write(fileName string) {
	FeatureDetectorNative_write_0(rcvr.nativeObj, fileName)
	return
}
