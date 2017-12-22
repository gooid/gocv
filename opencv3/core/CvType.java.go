package core

import (
	"fmt"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

const CvTypeCV_8U = 0
const CvTypeCV_8S = 1
const CvTypeCV_16U = 2
const CvTypeCV_16S = 3
const CvTypeCV_32S = 4
const CvTypeCV_32F = 5
const CvTypeCV_64F = 6
const CvTypeCV_USRTYPE1 = 7
const cV_CN_MAXCvType = 512
const cV_CN_SHIFTCvType = 3

var CvTypeCV_8UC1 = CvTypeCV_8UC(1)
var CvTypeCV_8UC2 = CvTypeCV_8UC(2)
var CvTypeCV_8UC3 = CvTypeCV_8UC(3)
var CvTypeCV_8UC4 = CvTypeCV_8UC(4)
var CvTypeCV_8SC1 = CvTypeCV_8SC(1)
var CvTypeCV_8SC2 = CvTypeCV_8SC(2)
var CvTypeCV_8SC3 = CvTypeCV_8SC(3)
var CvTypeCV_8SC4 = CvTypeCV_8SC(4)
var CvTypeCV_16UC1 = CvTypeCV_16UC(1)
var CvTypeCV_16UC2 = CvTypeCV_16UC(2)
var CvTypeCV_16UC3 = CvTypeCV_16UC(3)
var CvTypeCV_16UC4 = CvTypeCV_16UC(4)
var CvTypeCV_16SC1 = CvTypeCV_16SC(1)
var CvTypeCV_16SC2 = CvTypeCV_16SC(2)
var CvTypeCV_16SC3 = CvTypeCV_16SC(3)
var CvTypeCV_16SC4 = CvTypeCV_16SC(4)
var CvTypeCV_32SC1 = CvTypeCV_32SC(1)
var CvTypeCV_32SC2 = CvTypeCV_32SC(2)
var CvTypeCV_32SC3 = CvTypeCV_32SC(3)
var CvTypeCV_32SC4 = CvTypeCV_32SC(4)
var CvTypeCV_32FC1 = CvTypeCV_32FC(1)
var CvTypeCV_32FC2 = CvTypeCV_32FC(2)
var CvTypeCV_32FC3 = CvTypeCV_32FC(3)
var CvTypeCV_32FC4 = CvTypeCV_32FC(4)
var CvTypeCV_64FC1 = CvTypeCV_64FC(1)
var CvTypeCV_64FC2 = CvTypeCV_64FC(2)
var CvTypeCV_64FC3 = CvTypeCV_64FC(3)
var CvTypeCV_64FC4 = CvTypeCV_64FC(4)
var cV_DEPTH_MAXCvType = 1 << cV_CN_SHIFTCvType

type CvType struct {
}

func CvTypeCV_16SC(ch int) int {
	return CvTypeMakeType(CvTypeCV_16S, ch)
}
func CvTypeCV_16UC(ch int) int {
	return CvTypeMakeType(CvTypeCV_16U, ch)
}
func CvTypeCV_32FC(ch int) int {
	return CvTypeMakeType(CvTypeCV_32F, ch)
}
func CvTypeCV_32SC(ch int) int {
	return CvTypeMakeType(CvTypeCV_32S, ch)
}
func CvTypeCV_64FC(ch int) int {
	return CvTypeMakeType(CvTypeCV_64F, ch)
}
func CvTypeCV_8SC(ch int) int {
	return CvTypeMakeType(CvTypeCV_8S, ch)
}
func CvTypeCV_8UC(ch int) int {
	return CvTypeMakeType(CvTypeCV_8U, ch)
}
func NewCvType() (rcvr *CvType) {
	rcvr = &CvType{}
	return
}
func CvTypeELEM_SIZE(rtype int) int {
	switch CvTypeDepth(rtype) {
	case CvTypeCV_8U:
		fallthrough
	case CvTypeCV_8S:
		return CvTypeChannels(rtype)
		fallthrough
	case CvTypeCV_16U:
		fallthrough
	case CvTypeCV_16S:
		return 2 * CvTypeChannels(rtype)
		fallthrough
	case CvTypeCV_32S:
		fallthrough
	case CvTypeCV_32F:
		return 4 * CvTypeChannels(rtype)
		fallthrough
	case CvTypeCV_64F:
		return 8 * CvTypeChannels(rtype)
		fallthrough
	default:
		Throw(NewJavaLangUnsupportedOperationException(fmt.Sprintf("%v%v", "Unsupported CvType value: ", rtype)))
		return 0
	}
}
func CvTypeChannels(rtype int) int {
	return int(rtype)>>cV_CN_SHIFTCvType + 1
}
func CvTypeDepth(rtype int) int {
	return rtype & (cV_DEPTH_MAXCvType - 1)
}
func CvTypeIsInteger(rtype int) bool {
	return CvTypeDepth(rtype) < CvTypeCV_32F
}
func CvTypeMakeType(depth int, channels int) int {
	if channels <= 0 || channels >= cV_CN_MAXCvType {
		Throw(NewJavaLangUnsupportedOperationException(fmt.Sprintf("%v%v", "Channels count should be 1..", cV_CN_MAXCvType-1)))
	}
	if depth < 0 || depth >= cV_DEPTH_MAXCvType {
		Throw(NewJavaLangUnsupportedOperationException(fmt.Sprintf("%v%v", "Data type depth should be 0..", cV_DEPTH_MAXCvType-1)))
	}
	return depth&(cV_DEPTH_MAXCvType-1) + (channels-1)<<cV_CN_SHIFTCvType
}
func TypeToString(rtype int) string {
	var s string
	switch CvTypeDepth(rtype) {
	case CvTypeCV_8U:
		s = "CV_8U"
	case CvTypeCV_8S:
		s = "CV_8S"
	case CvTypeCV_16U:
		s = "CV_16U"
	case CvTypeCV_16S:
		s = "CV_16S"
	case CvTypeCV_32S:
		s = "CV_32S"
	case CvTypeCV_32F:
		s = "CV_32F"
	case CvTypeCV_64F:
		s = "CV_64F"
	case CvTypeCV_USRTYPE1:
		s = "CV_USRTYPE1"
	default:
		Throw(NewJavaLangUnsupportedOperationException(fmt.Sprintf("%v%v", "Unsupported CvType value: ", rtype)))
	}
	ch := CvTypeChannels(rtype)
	if ch <= 4 {
		return fmt.Sprintf("%v%v%v", s, "C", ch)
	} else {
		return fmt.Sprintf("%v%v%v%v", s, "C(", ch, ")")
	}
}
