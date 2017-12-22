package features2d

import (
	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

const Features2dDRAW_OVER_OUTIMG = 1
const Features2dNOT_DRAW_SINGLE_POINTS = 2
const Features2dDRAW_RICH_KEYPOINTS = 4

func DrawKeypoints(image *Mat, keypoints *MatOfKeyPoint, outImage *Mat, color *Scalar, flags int) {
	keypoints_mat := keypoints
	Features2dNative_drawKeypoints_0(image.GetNativeObjAddr(), keypoints_mat.GetNativeObjAddr(), outImage.GetNativeObjAddr(), color.Val[0], color.Val[1], color.Val[2], color.Val[3], flags)
	return
}
func DrawKeypoints2(image *Mat, keypoints *MatOfKeyPoint, outImage *Mat) {
	keypoints_mat := keypoints
	Features2dNative_drawKeypoints_1(image.GetNativeObjAddr(), keypoints_mat.GetNativeObjAddr(), outImage.GetNativeObjAddr())
	return
}
func DrawMatches(img1 *Mat, keypoints1 *MatOfKeyPoint, img2 *Mat, keypoints2 *MatOfKeyPoint, matches1to2 *MatOfDMatch, outImg *Mat, matchColor *Scalar, singlePointColor *Scalar, matchesMask *MatOfByte, flags int) {
	keypoints1_mat := keypoints1
	keypoints2_mat := keypoints2
	matches1to2_mat := matches1to2
	matchesMask_mat := matchesMask
	Features2dNative_drawMatches_0(img1.GetNativeObjAddr(), keypoints1_mat.GetNativeObjAddr(), img2.GetNativeObjAddr(), keypoints2_mat.GetNativeObjAddr(), matches1to2_mat.GetNativeObjAddr(), outImg.GetNativeObjAddr(), matchColor.Val[0], matchColor.Val[1], matchColor.Val[2], matchColor.Val[3], singlePointColor.Val[0], singlePointColor.Val[1], singlePointColor.Val[2], singlePointColor.Val[3], matchesMask_mat.GetNativeObjAddr(), flags)
	return
}
func DrawMatches2(img1 *Mat, keypoints1 *MatOfKeyPoint, img2 *Mat, keypoints2 *MatOfKeyPoint, matches1to2 *MatOfDMatch, outImg *Mat) {
	keypoints1_mat := keypoints1
	keypoints2_mat := keypoints2
	matches1to2_mat := matches1to2
	Features2dNative_drawMatches_1(img1.GetNativeObjAddr(), keypoints1_mat.GetNativeObjAddr(), img2.GetNativeObjAddr(), keypoints2_mat.GetNativeObjAddr(), matches1to2_mat.GetNativeObjAddr(), outImg.GetNativeObjAddr())
	return
}
func DrawMatches21(img1 *Mat, keypoints1 *MatOfKeyPoint, img2 *Mat, keypoints2 *MatOfKeyPoint, matches1to2 []*MatOfDMatch, outImg *Mat, matchColor *Scalar, singlePointColor *Scalar, matchesMask []*MatOfByte, flags int) {
	keypoints1_mat := keypoints1
	keypoints2_mat := keypoints2
	matches1to2_mat := ConvertersVectorVectorDMatch(matches1to2)
	matchesMask_mat := ConvertersVectorVectorByte(matchesMask)
	Features2dNative_drawMatches2_0(img1.GetNativeObjAddr(), keypoints1_mat.GetNativeObjAddr(), img2.GetNativeObjAddr(), keypoints2_mat.GetNativeObjAddr(), matches1to2_mat.GetNativeObjAddr(), outImg.GetNativeObjAddr(), matchColor.Val[0], matchColor.Val[1], matchColor.Val[2], matchColor.Val[3], singlePointColor.Val[0], singlePointColor.Val[1], singlePointColor.Val[2], singlePointColor.Val[3], matchesMask_mat.GetNativeObjAddr(), flags)
	return
}
func DrawMatches22(img1 *Mat, keypoints1 *MatOfKeyPoint, img2 *Mat, keypoints2 *MatOfKeyPoint, matches1to2 []*MatOfDMatch, outImg *Mat) {
	keypoints1_mat := keypoints1
	keypoints2_mat := keypoints2
	matches1to2_mat := ConvertersVectorVectorDMatch(matches1to2)
	Features2dNative_drawMatches2_1(img1.GetNativeObjAddr(), keypoints1_mat.GetNativeObjAddr(), img2.GetNativeObjAddr(), keypoints2_mat.GetNativeObjAddr(), matches1to2_mat.GetNativeObjAddr(), outImg.GetNativeObjAddr())
	return
}
func DrawMatchesKnn(img1 *Mat, keypoints1 *MatOfKeyPoint, img2 *Mat, keypoints2 *MatOfKeyPoint, matches1to2 []*MatOfDMatch, outImg *Mat, matchColor *Scalar, singlePointColor *Scalar, matchesMask []*MatOfByte, flags int) {
	keypoints1_mat := keypoints1
	keypoints2_mat := keypoints2
	matches1to2_mat := ConvertersVectorVectorDMatch(matches1to2)
	matchesMask_mat := ConvertersVectorVectorByte(matchesMask)
	Features2dNative_drawMatchesKnn_0(img1.GetNativeObjAddr(), keypoints1_mat.GetNativeObjAddr(), img2.GetNativeObjAddr(), keypoints2_mat.GetNativeObjAddr(), matches1to2_mat.GetNativeObjAddr(), outImg.GetNativeObjAddr(), matchColor.Val[0], matchColor.Val[1], matchColor.Val[2], matchColor.Val[3], singlePointColor.Val[0], singlePointColor.Val[1], singlePointColor.Val[2], singlePointColor.Val[3], matchesMask_mat.GetNativeObjAddr(), flags)
	return
}
func DrawMatchesKnn2(img1 *Mat, keypoints1 *MatOfKeyPoint, img2 *Mat, keypoints2 *MatOfKeyPoint, matches1to2 []*MatOfDMatch, outImg *Mat) {
	keypoints1_mat := keypoints1
	keypoints2_mat := keypoints2
	matches1to2_mat := ConvertersVectorVectorDMatch(matches1to2)
	Features2dNative_drawMatchesKnn_1(img1.GetNativeObjAddr(), keypoints1_mat.GetNativeObjAddr(), img2.GetNativeObjAddr(), keypoints2_mat.GetNativeObjAddr(), matches1to2_mat.GetNativeObjAddr(), outImg.GetNativeObjAddr())
	return
}
