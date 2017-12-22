package core

import (
	"fmt"
	"math"
)

type DMatch struct {
	QueryIdx int
	TrainIdx int
	ImgIdx   int
	Distance float32
}

func NewDMatch() (rcvr *DMatch) {
	rcvr = NewDMatch2(-1, -1, math.MaxFloat32)
	return
}
func NewDMatch2(_queryIdx int, _trainIdx int, _distance float32) (rcvr *DMatch) {
	rcvr = &DMatch{}
	rcvr.QueryIdx = _queryIdx
	rcvr.TrainIdx = _trainIdx
	rcvr.ImgIdx = -1
	rcvr.Distance = _distance
	return
}
func NewDMatch3(_queryIdx int, _trainIdx int, _imgIdx int, _distance float32) (rcvr *DMatch) {
	rcvr = &DMatch{}
	rcvr.QueryIdx = _queryIdx
	rcvr.TrainIdx = _trainIdx
	rcvr.ImgIdx = _imgIdx
	rcvr.Distance = _distance
	return
}
func (rcvr *DMatch) LessThan(it *DMatch) bool {
	return rcvr.Distance < it.Distance
}
func (rcvr *DMatch) String() string {
	return fmt.Sprintf("%v%v%v%v%v%v%v%v%v", "DMatch [queryIdx=", rcvr.QueryIdx, ", trainIdx=", rcvr.TrainIdx, ", imgIdx=", rcvr.ImgIdx, ", distance=", rcvr.Distance, "]")
}
