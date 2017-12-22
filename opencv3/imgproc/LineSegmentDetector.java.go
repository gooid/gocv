package imgproc

import (
	"runtime"

	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

type LineSegmentDetector struct {
	*Algorithm
}

func NewLineSegmentDetector(addr int64) (rcvr *LineSegmentDetector) {
	rcvr = &LineSegmentDetector{}
	rcvr.Algorithm = NewAlgorithm(addr)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
}
func (rcvr *LineSegmentDetector) CompareSegments(size *Size, lines1 *Mat, lines2 *Mat, _image *Mat) int {
	retVal := LineSegmentDetectorNative_compareSegments_0(rcvr.GetNativeObjAddr(), size.Width, size.Height, lines1.GetNativeObjAddr(), lines2.GetNativeObjAddr(), _image.GetNativeObjAddr())
	return retVal
}
func (rcvr *LineSegmentDetector) CompareSegments2(size *Size, lines1 *Mat, lines2 *Mat) int {
	retVal := LineSegmentDetectorNative_compareSegments_1(rcvr.GetNativeObjAddr(), size.Width, size.Height, lines1.GetNativeObjAddr(), lines2.GetNativeObjAddr())
	return retVal
}

func (rcvr *LineSegmentDetector) Detect(_image *Mat, _lines *Mat, width *Mat, prec *Mat, nfa *Mat) {
	LineSegmentDetectorNative_detect_0(rcvr.GetNativeObjAddr(), _image.GetNativeObjAddr(), _lines.GetNativeObjAddr(), width.GetNativeObjAddr(), prec.GetNativeObjAddr(), nfa.GetNativeObjAddr())
	return
}
func (rcvr *LineSegmentDetector) Detect2(_image *Mat, _lines *Mat) {
	LineSegmentDetectorNative_detect_1(rcvr.GetNativeObjAddr(), _image.GetNativeObjAddr(), _lines.GetNativeObjAddr())
	return
}

func (rcvr *LineSegmentDetector) DrawSegments(_image *Mat, lines *Mat) {
	LineSegmentDetectorNative_drawSegments_0(rcvr.GetNativeObjAddr(), _image.GetNativeObjAddr(), lines.GetNativeObjAddr())
	return
}

func (rcvr *LineSegmentDetector) finalize() {
	LineSegmentDetectorNative_delete(rcvr.GetNativeObjAddr())
}
