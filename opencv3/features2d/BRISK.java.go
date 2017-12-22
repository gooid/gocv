package features2d

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

type BRISK struct {
	*Feature2D
}

func NewBRISK(addr int64) (rcvr *BRISK) {
	rcvr = &BRISK{}
	rcvr.Feature2D = NewFeature2D(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func BRISKCreate(thresh int, octaves int, radiusList *MatOfFloat, numberList *MatOfInt, dMax float32, dMin float32, indexChange *MatOfInt) *BRISK {
	radiusList_mat := radiusList
	numberList_mat := numberList
	indexChange_mat := indexChange
	retVal := NewBRISK(BRISKNative_create_0(thresh, octaves, radiusList_mat.GetNativeObjAddr(), numberList_mat.GetNativeObjAddr(), dMax, dMin, indexChange_mat.GetNativeObjAddr()))
	return retVal
}
func BRISKCreate2(thresh int, octaves int, radiusList *MatOfFloat, numberList *MatOfInt) *BRISK {
	radiusList_mat := radiusList
	numberList_mat := numberList
	retVal := NewBRISK(BRISKNative_create_1(thresh, octaves, radiusList_mat.GetNativeObjAddr(), numberList_mat.GetNativeObjAddr()))
	return retVal
}
func BRISKCreate3(thresh int, octaves int, patternScale float32) *BRISK {
	retVal := NewBRISK(BRISKNative_create_2(thresh, octaves, patternScale))
	return retVal
}
func BRISKCreate4() *BRISK {
	retVal := NewBRISK(BRISKNative_create_3())
	return retVal
}
func BRISKCreate5(radiusList *MatOfFloat, numberList *MatOfInt, dMax float32, dMin float32, indexChange *MatOfInt) *BRISK {
	radiusList_mat := radiusList
	numberList_mat := numberList
	indexChange_mat := indexChange
	retVal := NewBRISK(BRISKNative_create_4(radiusList_mat.GetNativeObjAddr(), numberList_mat.GetNativeObjAddr(), dMax, dMin, indexChange_mat.GetNativeObjAddr()))
	return retVal
}
func BRISKCreate6(radiusList *MatOfFloat, numberList *MatOfInt) *BRISK {
	radiusList_mat := radiusList
	numberList_mat := numberList
	retVal := NewBRISK(BRISKNative_create_5(radiusList_mat.GetNativeObjAddr(), numberList_mat.GetNativeObjAddr()))
	return retVal
}
func (rcvr *BRISK) finalize() {
	BRISKNative_delete(rcvr.GetNativeObjAddr())
}
func (rcvr *BRISK) GetDefaultName() string {
	retVal := BRISKNative_getDefaultName_0(rcvr.GetNativeObjAddr())
	return retVal
}
