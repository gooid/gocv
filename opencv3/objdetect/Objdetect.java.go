package objdetect

import (
	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

const ObjdetectCASCADE_DO_CANNY_PRUNING = 1
const ObjdetectCASCADE_SCALE_IMAGE = 2
const ObjdetectCASCADE_FIND_BIGGEST_OBJECT = 4
const ObjdetectCASCADE_DO_ROUGH_SEARCH = 8

func GroupRectangles(rectList *MatOfRect, weights *MatOfInt, groupThreshold int, eps float64) {
	rectList_mat := rectList
	weights_mat := weights
	ObjdetectNative_groupRectangles_0(rectList_mat.GetNativeObjAddr(), weights_mat.GetNativeObjAddr(), groupThreshold, eps)
	return
}
func GroupRectangles2(rectList *MatOfRect, weights *MatOfInt, groupThreshold int) {
	rectList_mat := rectList
	weights_mat := weights
	ObjdetectNative_groupRectangles_1(rectList_mat.GetNativeObjAddr(), weights_mat.GetNativeObjAddr(), groupThreshold)
	return
}
