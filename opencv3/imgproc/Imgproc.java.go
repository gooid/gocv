package imgproc

import (
	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

const iPL_BORDER_CONSTANTImgproc = 0
const iPL_BORDER_REPLICATEImgproc = 1
const iPL_BORDER_REFLECTImgproc = 2
const iPL_BORDER_WRAPImgproc = 3
const iPL_BORDER_REFLECT_101Imgproc = 4
const iPL_BORDER_TRANSPARENTImgproc = 5
const cV_INTER_NNImgproc = 0
const cV_INTER_LINEARImgproc = 1
const cV_INTER_CUBICImgproc = 2
const cV_INTER_AREAImgproc = 3
const cV_INTER_LANCZOS4Imgproc = 4
const cV_MOP_ERODEImgproc = 0
const cV_MOP_DILATEImgproc = 1
const cV_MOP_OPENImgproc = 2
const cV_MOP_CLOSEImgproc = 3
const cV_MOP_GRADIENTImgproc = 4
const cV_MOP_TOPHATImgproc = 5
const cV_MOP_BLACKHATImgproc = 6
const cV_RETR_EXTERNALImgproc = 0
const cV_RETR_LISTImgproc = 1
const cV_RETR_CCOMPImgproc = 2
const cV_RETR_TREEImgproc = 3
const cV_RETR_FLOODFILLImgproc = 4
const cV_CHAIN_APPROX_NONEImgproc = 1
const cV_CHAIN_APPROX_SIMPLEImgproc = 2
const cV_CHAIN_APPROX_TC89_L1Imgproc = 3
const cV_CHAIN_APPROX_TC89_KCOSImgproc = 4
const cV_THRESH_BINARYImgproc = 0
const cV_THRESH_BINARY_INVImgproc = 1
const cV_THRESH_TRUNCImgproc = 2
const cV_THRESH_TOZEROImgproc = 3
const cV_THRESH_TOZERO_INVImgproc = 4
const cV_THRESH_MASKImgproc = 7
const cV_THRESH_OTSUImgproc = 8
const cV_THRESH_TRIANGLEImgproc = 16
const LINE_AA = 16
const LINE_8 = 8
const LINE_4 = 4
const CV_BLUR_NO_SCALE = 0
const CV_BLUR = 1
const CV_GAUSSIAN = 2
const CV_MEDIAN = 3
const CV_BILATERAL = 4
const CV_GAUSSIAN_5x5 = 7
const CV_MAX_SOBEL_KSIZE = 7
const CV_RGBA2mRGBA = 125
const CV_mRGBA2RGBA = 126
const CV_WARP_FILL_OUTLIERS = 8
const CV_WARP_INVERSE_MAP = 16
const CV_SHAPE_RECT = 0
const CV_SHAPE_CROSS = 1
const CV_SHAPE_ELLIPSE = 2
const CV_SHAPE_CUSTOM = 100
const CV_CHAIN_CODE = 0
const CV_LINK_RUNS = 5
const CV_POLY_APPROX_DP = 0
const CV_CONTOURS_MATCH_I1 = 1
const CV_CONTOURS_MATCH_I2 = 2
const CV_CONTOURS_MATCH_I3 = 3
const CV_CLOCKWISE = 1
const CV_COUNTER_CLOCKWISE = 2
const CV_COMP_CORREL = 0
const CV_COMP_CHISQR = 1
const CV_COMP_INTERSECT = 2
const CV_COMP_BHATTACHARYYA = 3
const CV_COMP_CHISQR_ALT = 4
const CV_COMP_KL_DIV = 5
const CV_DIST_MASK_3 = 3
const CV_DIST_MASK_5 = 5
const CV_DIST_MASK_PRECISE = 0
const CV_DIST_LABEL_CCOMP = 0
const CV_DIST_LABEL_PIXEL = 1
const CV_DIST_L1 = 1
const CV_DIST_L2 = 2
const CV_DIST_C = 3
const CV_DIST_L12 = 4
const CV_DIST_FAIR = 5
const CV_DIST_WELSCH = 6
const CV_DIST_HUBER = 7
const CV_HOUGH_STANDARD = 0
const CV_HOUGH_PROBABILISTIC = 1
const CV_HOUGH_MULTI_SCALE = 2
const CV_HOUGH_GRADIENT = 3
const MORPH_ERODE = 0
const MORPH_DILATE = 1
const MORPH_OPEN = 2
const MORPH_CLOSE = 3
const MORPH_GRADIENT = 4
const MORPH_TOPHAT = 5
const MORPH_BLACKHAT = 6
const MORPH_HITMISS = 7
const MORPH_RECT = 0
const MORPH_CROSS = 1
const MORPH_ELLIPSE = 2
const INTER_NEAREST = 0
const INTER_LINEAR = 1
const INTER_CUBIC = 2
const INTER_AREA = 3
const INTER_LANCZOS4 = 4
const INTER_MAX = 7
const WARP_FILL_OUTLIERS = 8
const WARP_INVERSE_MAP = 16
const INTER_BITS = 5
const DIST_L1 = 1
const DIST_L2 = 2
const DIST_C = 3
const DIST_L12 = 4
const DIST_FAIR = 5
const DIST_WELSCH = 6
const DIST_HUBER = 7
const DIST_MASK_3 = 3
const DIST_MASK_5 = 5
const DIST_MASK_PRECISE = 0
const THRESH_BINARY = 0
const THRESH_BINARY_INV = 1
const THRESH_TRUNC = 2
const THRESH_TOZERO = 3
const THRESH_TOZERO_INV = 4
const THRESH_MASK = 7
const THRESH_OTSU = 8
const THRESH_TRIANGLE = 16
const ADAPTIVE_THRESH_MEAN_C = 0
const ADAPTIVE_THRESH_GAUSSIAN_C = 1
const PROJ_SPHERICAL_ORTHO = 0
const PROJ_SPHERICAL_EQRECT = 1
const GC_BGD = 0
const GC_FGD = 1
const GC_PR_BGD = 2
const GC_PR_FGD = 3
const GC_INIT_WITH_RECT = 0
const GC_INIT_WITH_MASK = 1
const GC_EVAL = 2
const DIST_LABEL_CCOMP = 0
const DIST_LABEL_PIXEL = 1
const CC_STAT_LEFT = 0
const CC_STAT_TOP = 1
const CC_STAT_WIDTH = 2
const CC_STAT_HEIGHT = 3
const CC_STAT_AREA = 4
const CC_STAT_MAX = 5
const CCL_WU = 0
const CCL_GRANA = 1
const RETR_EXTERNAL = 0
const RETR_LIST = 1
const RETR_CCOMP = 2
const RETR_TREE = 3
const RETR_FLOODFILL = 4
const CHAIN_APPROX_NONE = 1
const CHAIN_APPROX_SIMPLE = 2
const CHAIN_APPROX_TC89_L1 = 3
const CHAIN_APPROX_TC89_KCOS = 4
const CONTOURS_MATCH_I1 = 1
const CONTOURS_MATCH_I2 = 2
const CONTOURS_MATCH_I3 = 3
const HOUGH_STANDARD = 0
const HOUGH_PROBABILISTIC = 1
const HOUGH_MULTI_SCALE = 2
const HOUGH_GRADIENT = 3
const LSD_REFINE_NONE = 0
const LSD_REFINE_STD = 1
const LSD_REFINE_ADV = 2
const HISTCMP_CORREL = 0
const HISTCMP_CHISQR = 1
const HISTCMP_INTERSECT = 2
const HISTCMP_BHATTACHARYYA = 3
const HISTCMP_CHISQR_ALT = 4
const HISTCMP_KL_DIV = 5
const COLOR_BGR2BGRA = 0
const COLOR_BGRA2BGR = 1
const COLOR_BGR2RGBA = 2
const COLOR_RGBA2BGR = 3
const COLOR_BGR2RGB = 4
const COLOR_BGRA2RGBA = 5
const COLOR_BGR2GRAY = 6
const COLOR_RGB2GRAY = 7
const COLOR_GRAY2BGR = 8
const COLOR_GRAY2BGRA = 9
const COLOR_BGRA2GRAY = 10
const COLOR_RGBA2GRAY = 11
const COLOR_BGR2BGR565 = 12
const COLOR_RGB2BGR565 = 13
const COLOR_BGR5652BGR = 14
const COLOR_BGR5652RGB = 15
const COLOR_BGRA2BGR565 = 16
const COLOR_RGBA2BGR565 = 17
const COLOR_BGR5652BGRA = 18
const COLOR_BGR5652RGBA = 19
const COLOR_GRAY2BGR565 = 20
const COLOR_BGR5652GRAY = 21
const COLOR_BGR2BGR555 = 22
const COLOR_RGB2BGR555 = 23
const COLOR_BGR5552BGR = 24
const COLOR_BGR5552RGB = 25
const COLOR_BGRA2BGR555 = 26
const COLOR_RGBA2BGR555 = 27
const COLOR_BGR5552BGRA = 28
const COLOR_BGR5552RGBA = 29
const COLOR_GRAY2BGR555 = 30
const COLOR_BGR5552GRAY = 31
const COLOR_BGR2XYZ = 32
const COLOR_RGB2XYZ = 33
const COLOR_XYZ2BGR = 34
const COLOR_XYZ2RGB = 35
const COLOR_BGR2YCrCb = 36
const COLOR_RGB2YCrCb = 37
const COLOR_YCrCb2BGR = 38
const COLOR_YCrCb2RGB = 39
const COLOR_BGR2HSV = 40
const COLOR_RGB2HSV = 41
const COLOR_BGR2Lab = 44
const COLOR_RGB2Lab = 45
const COLOR_BGR2Luv = 50
const COLOR_RGB2Luv = 51
const COLOR_BGR2HLS = 52
const COLOR_RGB2HLS = 53
const COLOR_HSV2BGR = 54
const COLOR_HSV2RGB = 55
const COLOR_Lab2BGR = 56
const COLOR_Lab2RGB = 57
const COLOR_Luv2BGR = 58
const COLOR_Luv2RGB = 59
const COLOR_HLS2BGR = 60
const COLOR_HLS2RGB = 61
const COLOR_BGR2HSV_FULL = 66
const COLOR_RGB2HSV_FULL = 67
const COLOR_BGR2HLS_FULL = 68
const COLOR_RGB2HLS_FULL = 69
const COLOR_HSV2BGR_FULL = 70
const COLOR_HSV2RGB_FULL = 71
const COLOR_HLS2BGR_FULL = 72
const COLOR_HLS2RGB_FULL = 73
const COLOR_LBGR2Lab = 74
const COLOR_LRGB2Lab = 75
const COLOR_LBGR2Luv = 76
const COLOR_LRGB2Luv = 77
const COLOR_Lab2LBGR = 78
const COLOR_Lab2LRGB = 79
const COLOR_Luv2LBGR = 80
const COLOR_Luv2LRGB = 81
const COLOR_BGR2YUV = 82
const COLOR_RGB2YUV = 83
const COLOR_YUV2BGR = 84
const COLOR_YUV2RGB = 85
const COLOR_YUV2RGB_NV12 = 90
const COLOR_YUV2BGR_NV12 = 91
const COLOR_YUV2RGB_NV21 = 92
const COLOR_YUV2BGR_NV21 = 93
const COLOR_YUV2RGBA_NV12 = 94
const COLOR_YUV2BGRA_NV12 = 95
const COLOR_YUV2RGBA_NV21 = 96
const COLOR_YUV2BGRA_NV21 = 97
const COLOR_YUV2RGB_YV12 = 98
const COLOR_YUV2BGR_YV12 = 99
const COLOR_YUV2RGB_IYUV = 100
const COLOR_YUV2BGR_IYUV = 101
const COLOR_YUV2RGBA_YV12 = 102
const COLOR_YUV2BGRA_YV12 = 103
const COLOR_YUV2RGBA_IYUV = 104
const COLOR_YUV2BGRA_IYUV = 105
const COLOR_YUV2GRAY_420 = 106
const COLOR_YUV2RGB_UYVY = 107
const COLOR_YUV2BGR_UYVY = 108
const COLOR_YUV2RGBA_UYVY = 111
const COLOR_YUV2BGRA_UYVY = 112
const COLOR_YUV2RGB_YUY2 = 115
const COLOR_YUV2BGR_YUY2 = 116
const COLOR_YUV2RGB_YVYU = 117
const COLOR_YUV2BGR_YVYU = 118
const COLOR_YUV2RGBA_YUY2 = 119
const COLOR_YUV2BGRA_YUY2 = 120
const COLOR_YUV2RGBA_YVYU = 121
const COLOR_YUV2BGRA_YVYU = 122
const COLOR_YUV2GRAY_UYVY = 123
const COLOR_YUV2GRAY_YUY2 = 124
const COLOR_RGBA2mRGBA = 125
const COLOR_mRGBA2RGBA = 126
const COLOR_RGB2YUV_I420 = 127
const COLOR_BGR2YUV_I420 = 128
const COLOR_RGBA2YUV_I420 = 129
const COLOR_BGRA2YUV_I420 = 130
const COLOR_RGB2YUV_YV12 = 131
const COLOR_BGR2YUV_YV12 = 132
const COLOR_RGBA2YUV_YV12 = 133
const COLOR_BGRA2YUV_YV12 = 134
const COLOR_BayerBG2BGR = 46
const COLOR_BayerGB2BGR = 47
const COLOR_BayerRG2BGR = 48
const COLOR_BayerGR2BGR = 49
const COLOR_BayerBG2GRAY = 86
const COLOR_BayerGB2GRAY = 87
const COLOR_BayerRG2GRAY = 88
const COLOR_BayerGR2GRAY = 89
const COLOR_BayerBG2BGR_VNG = 62
const COLOR_BayerGB2BGR_VNG = 63
const COLOR_BayerRG2BGR_VNG = 64
const COLOR_BayerGR2BGR_VNG = 65
const COLOR_BayerBG2BGR_EA = 135
const COLOR_BayerGB2BGR_EA = 136
const COLOR_BayerRG2BGR_EA = 137
const COLOR_BayerGR2BGR_EA = 138
const COLOR_BayerBG2BGRA = 139
const COLOR_BayerGB2BGRA = 140
const COLOR_BayerRG2BGRA = 141
const COLOR_BayerGR2BGRA = 142
const COLOR_COLORCVT_MAX = 143
const INTERSECT_NONE = 0
const INTERSECT_PARTIAL = 1
const INTERSECT_FULL = 2
const TM_SQDIFF = 0
const TM_SQDIFF_NORMED = 1
const TM_CCORR = 2
const TM_CCORR_NORMED = 3
const TM_CCOEFF = 4
const TM_CCOEFF_NORMED = 5
const COLORMAP_AUTUMN = 0
const COLORMAP_BONE = 1
const COLORMAP_JET = 2
const COLORMAP_WINTER = 3
const COLORMAP_RAINBOW = 4
const COLORMAP_OCEAN = 5
const COLORMAP_SUMMER = 6
const COLORMAP_SPRING = 7
const COLORMAP_COOL = 8
const COLORMAP_HSV = 9
const COLORMAP_PINK = 10
const COLORMAP_HOT = 11
const COLORMAP_PARULA = 12
const MARKER_CROSS = 0
const MARKER_TILTED_CROSS = 1
const MARKER_STAR = 2
const MARKER_DIAMOND = 3
const MARKER_SQUARE = 4
const MARKER_TRIANGLE_UP = 5
const MARKER_TRIANGLE_DOWN = 6

const CV_SCHARR = -1
const CV_COMP_HELLINGER = CV_COMP_BHATTACHARYYA
const CV_DIST_USER = -1
const CV_CANNY_L2_GRADIENT = 1 << 31
const INTER_BITS2 = INTER_BITS * 2
const INTER_TAB_SIZE = 1 << INTER_BITS
const INTER_TAB_SIZE2 = INTER_TAB_SIZE * INTER_TAB_SIZE
const DIST_USER = -1
const FLOODFILL_FIXED_RANGE = 1 << 16
const FLOODFILL_MASK_ONLY = 1 << 17
const CCL_DEFAULT = -1
const HISTCMP_HELLINGER = HISTCMP_BHATTACHARYYA
const COLOR_RGB2RGBA = COLOR_BGR2BGRA
const COLOR_RGBA2RGB = COLOR_BGRA2BGR
const COLOR_RGB2BGRA = COLOR_BGR2RGBA
const COLOR_BGRA2RGB = COLOR_RGBA2BGR
const COLOR_RGB2BGR = COLOR_BGR2RGB
const COLOR_RGBA2BGRA = COLOR_BGRA2RGBA
const COLOR_GRAY2RGB = COLOR_GRAY2BGR
const COLOR_GRAY2RGBA = COLOR_GRAY2BGRA
const COLOR_YUV420sp2RGB = COLOR_YUV2RGB_NV21
const COLOR_YUV420sp2BGR = COLOR_YUV2BGR_NV21
const COLOR_YUV420sp2RGBA = COLOR_YUV2RGBA_NV21
const COLOR_YUV420sp2BGRA = COLOR_YUV2BGRA_NV21
const COLOR_YUV2RGB_I420 = COLOR_YUV2RGB_IYUV
const COLOR_YUV2BGR_I420 = COLOR_YUV2BGR_IYUV
const COLOR_YUV420p2RGB = COLOR_YUV2RGB_YV12
const COLOR_YUV420p2BGR = COLOR_YUV2BGR_YV12
const COLOR_YUV2RGBA_I420 = COLOR_YUV2RGBA_IYUV
const COLOR_YUV2BGRA_I420 = COLOR_YUV2BGRA_IYUV
const COLOR_YUV420p2RGBA = COLOR_YUV2RGBA_YV12
const COLOR_YUV420p2BGRA = COLOR_YUV2BGRA_YV12
const COLOR_YUV2GRAY_NV21 = COLOR_YUV2GRAY_420
const COLOR_YUV2GRAY_NV12 = COLOR_YUV2GRAY_420
const COLOR_YUV2GRAY_YV12 = COLOR_YUV2GRAY_420
const COLOR_YUV2GRAY_IYUV = COLOR_YUV2GRAY_420
const COLOR_YUV2GRAY_I420 = COLOR_YUV2GRAY_420
const COLOR_YUV420sp2GRAY = COLOR_YUV2GRAY_420
const COLOR_YUV420p2GRAY = COLOR_YUV2GRAY_420
const COLOR_YUV2RGB_Y422 = COLOR_YUV2RGB_UYVY
const COLOR_YUV2BGR_Y422 = COLOR_YUV2BGR_UYVY
const COLOR_YUV2RGB_UYNV = COLOR_YUV2RGB_UYVY
const COLOR_YUV2BGR_UYNV = COLOR_YUV2BGR_UYVY
const COLOR_YUV2RGBA_Y422 = COLOR_YUV2RGBA_UYVY
const COLOR_YUV2BGRA_Y422 = COLOR_YUV2BGRA_UYVY
const COLOR_YUV2RGBA_UYNV = COLOR_YUV2RGBA_UYVY
const COLOR_YUV2BGRA_UYNV = COLOR_YUV2BGRA_UYVY
const COLOR_YUV2RGB_YUYV = COLOR_YUV2RGB_YUY2
const COLOR_YUV2BGR_YUYV = COLOR_YUV2BGR_YUY2
const COLOR_YUV2RGB_YUNV = COLOR_YUV2RGB_YUY2
const COLOR_YUV2BGR_YUNV = COLOR_YUV2BGR_YUY2
const COLOR_YUV2RGBA_YUYV = COLOR_YUV2RGBA_YUY2
const COLOR_YUV2BGRA_YUYV = COLOR_YUV2BGRA_YUY2
const COLOR_YUV2RGBA_YUNV = COLOR_YUV2RGBA_YUY2
const COLOR_YUV2BGRA_YUNV = COLOR_YUV2BGRA_YUY2
const COLOR_YUV2GRAY_Y422 = COLOR_YUV2GRAY_UYVY
const COLOR_YUV2GRAY_UYNV = COLOR_YUV2GRAY_UYVY
const COLOR_YUV2GRAY_YVYU = COLOR_YUV2GRAY_YUY2
const COLOR_YUV2GRAY_YUYV = COLOR_YUV2GRAY_YUY2
const COLOR_YUV2GRAY_YUNV = COLOR_YUV2GRAY_YUY2
const COLOR_RGB2YUV_IYUV = COLOR_RGB2YUV_I420
const COLOR_BGR2YUV_IYUV = COLOR_BGR2YUV_I420
const COLOR_RGBA2YUV_IYUV = COLOR_RGBA2YUV_I420
const COLOR_BGRA2YUV_IYUV = COLOR_BGRA2YUV_I420
const COLOR_BayerBG2RGB = COLOR_BayerRG2BGR
const COLOR_BayerGB2RGB = COLOR_BayerGR2BGR
const COLOR_BayerRG2RGB = COLOR_BayerBG2BGR
const COLOR_BayerGR2RGB = COLOR_BayerGB2BGR
const COLOR_BayerBG2RGB_VNG = COLOR_BayerRG2BGR_VNG
const COLOR_BayerGB2RGB_VNG = COLOR_BayerGR2BGR_VNG
const COLOR_BayerRG2RGB_VNG = COLOR_BayerBG2BGR_VNG
const COLOR_BayerGR2RGB_VNG = COLOR_BayerGB2BGR_VNG
const COLOR_BayerBG2RGB_EA = COLOR_BayerRG2BGR_EA
const COLOR_BayerGB2RGB_EA = COLOR_BayerGR2BGR_EA
const COLOR_BayerRG2RGB_EA = COLOR_BayerBG2BGR_EA
const COLOR_BayerGR2RGB_EA = COLOR_BayerGB2BGR_EA
const COLOR_BayerBG2RGBA = COLOR_BayerRG2BGRA
const COLOR_BayerGB2RGBA = COLOR_BayerGR2BGRA
const COLOR_BayerRG2RGBA = COLOR_BayerBG2BGRA
const COLOR_BayerGR2RGBA = COLOR_BayerGB2BGRA

func Canny(dx *Mat, dy *Mat, edges *Mat, threshold1 float64, threshold2 float64, L2gradient bool) {
	ImgprocNative_Canny_0(dx.GetNativeObjAddr(), dy.GetNativeObjAddr(), edges.GetNativeObjAddr(), threshold1, threshold2, L2gradient)
	return
}
func Canny2(dx *Mat, dy *Mat, edges *Mat, threshold1 float64, threshold2 float64) {
	ImgprocNative_Canny_1(dx.GetNativeObjAddr(), dy.GetNativeObjAddr(), edges.GetNativeObjAddr(), threshold1, threshold2)
	return
}
func Canny3(image *Mat, edges *Mat, threshold1 float64, threshold2 float64, apertureSize int, L2gradient bool) {
	ImgprocNative_Canny_2(image.GetNativeObjAddr(), edges.GetNativeObjAddr(), threshold1, threshold2, apertureSize, L2gradient)
	return
}
func Canny4(image *Mat, edges *Mat, threshold1 float64, threshold2 float64) {
	ImgprocNative_Canny_3(image.GetNativeObjAddr(), edges.GetNativeObjAddr(), threshold1, threshold2)
	return
}

func EMD(signature1 *Mat, signature2 *Mat, distType int, cost *Mat, flow *Mat) float32 {
	retVal := ImgprocNative_EMD_0(signature1.GetNativeObjAddr(), signature2.GetNativeObjAddr(), distType, cost.GetNativeObjAddr(), flow.GetNativeObjAddr())
	return retVal
}
func EMD2(signature1 *Mat, signature2 *Mat, distType int) float32 {
	retVal := ImgprocNative_EMD_1(signature1.GetNativeObjAddr(), signature2.GetNativeObjAddr(), distType)
	return retVal
}

func GaussianBlur(src *Mat, dst *Mat, ksize *Size, sigmaX float64, sigmaY float64, borderType int) {
	ImgprocNative_GaussianBlur_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), ksize.Width, ksize.Height, sigmaX, sigmaY, borderType)
	return
}
func GaussianBlur2(src *Mat, dst *Mat, ksize *Size, sigmaX float64, sigmaY float64) {
	ImgprocNative_GaussianBlur_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), ksize.Width, ksize.Height, sigmaX, sigmaY)
	return
}
func GaussianBlur3(src *Mat, dst *Mat, ksize *Size, sigmaX float64) {
	ImgprocNative_GaussianBlur_2(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), ksize.Width, ksize.Height, sigmaX)
	return
}

func HoughCircles(image *Mat, circles *Mat, method int, dp float64, minDist float64, param1 float64, param2 float64, minRadius int, maxRadius int) {
	ImgprocNative_HoughCircles_0(image.GetNativeObjAddr(), circles.GetNativeObjAddr(), method, dp, minDist, param1, param2, minRadius, maxRadius)
	return
}
func HoughCircles2(image *Mat, circles *Mat, method int, dp float64, minDist float64) {
	ImgprocNative_HoughCircles_1(image.GetNativeObjAddr(), circles.GetNativeObjAddr(), method, dp, minDist)
	return
}

func HoughLines(image *Mat, lines *Mat, rho float64, theta float64, threshold int, srn float64, stn float64, min_theta float64, max_theta float64) {
	ImgprocNative_HoughLines_0(image.GetNativeObjAddr(), lines.GetNativeObjAddr(), rho, theta, threshold, srn, stn, min_theta, max_theta)
	return
}
func HoughLines2(image *Mat, lines *Mat, rho float64, theta float64, threshold int) {
	ImgprocNative_HoughLines_1(image.GetNativeObjAddr(), lines.GetNativeObjAddr(), rho, theta, threshold)
	return
}
func HoughLinesP(image *Mat, lines *Mat, rho float64, theta float64, threshold int, minLineLength float64, maxLineGap float64) {
	ImgprocNative_HoughLinesP_0(image.GetNativeObjAddr(), lines.GetNativeObjAddr(), rho, theta, threshold, minLineLength, maxLineGap)
	return
}
func HoughLinesP2(image *Mat, lines *Mat, rho float64, theta float64, threshold int) {
	ImgprocNative_HoughLinesP_1(image.GetNativeObjAddr(), lines.GetNativeObjAddr(), rho, theta, threshold)
	return
}

func HuMoments(m *Moments, hu *Mat) {
	ImgprocNative_HuMoments_0(m.M00, m.M10, m.M01, m.M20, m.M11, m.M02, m.M30, m.M21, m.M12, m.M03, hu.GetNativeObjAddr())
	return
}

func Laplacian(src *Mat, dst *Mat, ddepth int, ksize int, scale float64, delta float64, borderType int) {
	ImgprocNative_Laplacian_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), ddepth, ksize, scale, delta, borderType)
	return
}
func Laplacian2(src *Mat, dst *Mat, ddepth int, ksize int, scale float64, delta float64) {
	ImgprocNative_Laplacian_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), ddepth, ksize, scale, delta)
	return
}
func Laplacian3(src *Mat, dst *Mat, ddepth int) {
	ImgprocNative_Laplacian_2(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), ddepth)
	return
}

func Scharr(src *Mat, dst *Mat, ddepth int, dx int, dy int, scale float64, delta float64, borderType int) {
	ImgprocNative_Scharr_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), ddepth, dx, dy, scale, delta, borderType)
	return
}
func Scharr2(src *Mat, dst *Mat, ddepth int, dx int, dy int, scale float64, delta float64) {
	ImgprocNative_Scharr_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), ddepth, dx, dy, scale, delta)
	return
}
func Scharr3(src *Mat, dst *Mat, ddepth int, dx int, dy int) {
	ImgprocNative_Scharr_2(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), ddepth, dx, dy)
	return
}

func Sobel(src *Mat, dst *Mat, ddepth int, dx int, dy int, ksize int, scale float64, delta float64, borderType int) {
	ImgprocNative_Sobel_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), ddepth, dx, dy, ksize, scale, delta, borderType)
	return
}
func Sobel2(src *Mat, dst *Mat, ddepth int, dx int, dy int, ksize int, scale float64, delta float64) {
	ImgprocNative_Sobel_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), ddepth, dx, dy, ksize, scale, delta)
	return
}
func Sobel3(src *Mat, dst *Mat, ddepth int, dx int, dy int) {
	ImgprocNative_Sobel_2(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), ddepth, dx, dy)
	return
}

func Accumulate(src *Mat, dst *Mat, mask *Mat) {
	ImgprocNative_accumulate_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), mask.GetNativeObjAddr())
	return
}
func Accumulate2(src *Mat, dst *Mat) {
	ImgprocNative_accumulate_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr())
	return
}
func AccumulateProduct(src1 *Mat, src2 *Mat, dst *Mat, mask *Mat) {
	ImgprocNative_accumulateProduct_0(src1.GetNativeObjAddr(), src2.GetNativeObjAddr(), dst.GetNativeObjAddr(), mask.GetNativeObjAddr())
	return
}
func AccumulateProduct2(src1 *Mat, src2 *Mat, dst *Mat) {
	ImgprocNative_accumulateProduct_1(src1.GetNativeObjAddr(), src2.GetNativeObjAddr(), dst.GetNativeObjAddr())
	return
}

func AccumulateSquare(src *Mat, dst *Mat, mask *Mat) {
	ImgprocNative_accumulateSquare_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), mask.GetNativeObjAddr())
	return
}
func AccumulateSquare2(src *Mat, dst *Mat) {
	ImgprocNative_accumulateSquare_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr())
	return
}

func AccumulateWeighted(src *Mat, dst *Mat, alpha float64, mask *Mat) {
	ImgprocNative_accumulateWeighted_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), alpha, mask.GetNativeObjAddr())
	return
}
func AccumulateWeighted2(src *Mat, dst *Mat, alpha float64) {
	ImgprocNative_accumulateWeighted_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), alpha)
	return
}

func AdaptiveThreshold(src *Mat, dst *Mat, maxValue float64, adaptiveMethod int, thresholdType int, blockSize int, C float64) {
	ImgprocNative_adaptiveThreshold_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), maxValue, adaptiveMethod, thresholdType, blockSize, C)
	return
}

func ApplyColorMap(src *Mat, dst *Mat, userColor *Mat) {
	ImgprocNative_applyColorMap_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), userColor.GetNativeObjAddr())
	return
}
func ApplyColorMap2(src *Mat, dst *Mat, colormap int) {
	ImgprocNative_applyColorMap_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), colormap)
	return
}

func ApproxPolyDP(curve *MatOfPoint2f, approxCurve *MatOfPoint2f, epsilon float64, closed bool) {
	curve_mat := curve
	approxCurve_mat := approxCurve
	ImgprocNative_approxPolyDP_0(curve_mat.GetNativeObjAddr(), approxCurve_mat.GetNativeObjAddr(), epsilon, closed)
	return
}

func ArcLength(curve *MatOfPoint2f, closed bool) float64 {
	curve_mat := curve
	retVal := ImgprocNative_arcLength_0(curve_mat.GetNativeObjAddr(), closed)
	return retVal
}

func ArrowedLine(img *Mat, pt1 *Point, pt2 *Point, color *Scalar, thickness int, line_type int, shift int, tipLength float64) {
	ImgprocNative_arrowedLine_0(img.GetNativeObjAddr(), pt1.X, pt1.Y, pt2.X, pt2.Y, color.Val[0], color.Val[1], color.Val[2], color.Val[3], thickness, line_type, shift, tipLength)
	return
}
func ArrowedLine2(img *Mat, pt1 *Point, pt2 *Point, color *Scalar) {
	ImgprocNative_arrowedLine_1(img.GetNativeObjAddr(), pt1.X, pt1.Y, pt2.X, pt2.Y, color.Val[0], color.Val[1], color.Val[2], color.Val[3])
	return
}

func BilateralFilter(src *Mat, dst *Mat, d int, sigmaColor float64, sigmaSpace float64, borderType int) {
	ImgprocNative_bilateralFilter_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), d, sigmaColor, sigmaSpace, borderType)
	return
}
func BilateralFilter2(src *Mat, dst *Mat, d int, sigmaColor float64, sigmaSpace float64) {
	ImgprocNative_bilateralFilter_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), d, sigmaColor, sigmaSpace)
	return
}

func Blur(src *Mat, dst *Mat, ksize *Size, anchor *Point, borderType int) {
	ImgprocNative_blur_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), ksize.Width, ksize.Height, anchor.X, anchor.Y, borderType)
	return
}
func Blur2(src *Mat, dst *Mat, ksize *Size, anchor *Point) {
	ImgprocNative_blur_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), ksize.Width, ksize.Height, anchor.X, anchor.Y)
	return
}
func Blur3(src *Mat, dst *Mat, ksize *Size) {
	ImgprocNative_blur_2(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), ksize.Width, ksize.Height)
	return
}

func BoundingRect(points *MatOfPoint) *Rect {
	points_mat := points
	retVal := NewRect5(ImgprocNative_boundingRect_0(points_mat.GetNativeObjAddr()))
	return retVal
}

func BoxFilter(src *Mat, dst *Mat, ddepth int, ksize *Size, anchor *Point, normalize bool, borderType int) {
	ImgprocNative_boxFilter_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), ddepth, ksize.Width, ksize.Height, anchor.X, anchor.Y, normalize, borderType)
	return
}
func BoxFilter2(src *Mat, dst *Mat, ddepth int, ksize *Size, anchor *Point, normalize bool) {
	ImgprocNative_boxFilter_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), ddepth, ksize.Width, ksize.Height, anchor.X, anchor.Y, normalize)
	return
}
func BoxFilter3(src *Mat, dst *Mat, ddepth int, ksize *Size) {
	ImgprocNative_boxFilter_2(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), ddepth, ksize.Width, ksize.Height)
	return
}

func BoxPoints(box *RotatedRect, points *Mat) {
	ImgprocNative_boxPoints_0(box.Center.X, box.Center.Y, box.Size.Width, box.Size.Height, box.Angle, points.GetNativeObjAddr())
	return
}

func CalcBackProject(images []*Mat, channels *MatOfInt, hist *Mat, dst *Mat, ranges *MatOfFloat, scale float64) {
	images_mat := ConvertersVectorMat(images)
	channels_mat := channels
	ranges_mat := ranges
	ImgprocNative_calcBackProject_0(images_mat.GetNativeObjAddr(), channels_mat.GetNativeObjAddr(), hist.GetNativeObjAddr(), dst.GetNativeObjAddr(), ranges_mat.GetNativeObjAddr(), scale)
	return
}

func CalcHist(images []*Mat, channels *MatOfInt, mask *Mat, hist *Mat, histSize *MatOfInt, ranges *MatOfFloat, accumulate bool) {
	images_mat := ConvertersVectorMat(images)
	channels_mat := channels
	histSize_mat := histSize
	ranges_mat := ranges
	ImgprocNative_calcHist_0(images_mat.GetNativeObjAddr(), channels_mat.GetNativeObjAddr(), mask.GetNativeObjAddr(), hist.GetNativeObjAddr(), histSize_mat.GetNativeObjAddr(), ranges_mat.GetNativeObjAddr(), accumulate)
	return
}
func CalcHist2(images []*Mat, channels *MatOfInt, mask *Mat, hist *Mat, histSize *MatOfInt, ranges *MatOfFloat) {
	images_mat := ConvertersVectorMat(images)
	channels_mat := channels
	histSize_mat := histSize
	ranges_mat := ranges
	ImgprocNative_calcHist_1(images_mat.GetNativeObjAddr(), channels_mat.GetNativeObjAddr(), mask.GetNativeObjAddr(), hist.GetNativeObjAddr(), histSize_mat.GetNativeObjAddr(), ranges_mat.GetNativeObjAddr())
	return
}

func Circle(img *Mat, center *Point, radius int, color *Scalar, thickness int, lineType int, shift int) {
	ImgprocNative_circle_0(img.GetNativeObjAddr(), center.X, center.Y, radius, color.Val[0], color.Val[1], color.Val[2], color.Val[3], thickness, lineType, shift)
	return
}
func Circle2(img *Mat, center *Point, radius int, color *Scalar, thickness int) {
	ImgprocNative_circle_1(img.GetNativeObjAddr(), center.X, center.Y, radius, color.Val[0], color.Val[1], color.Val[2], color.Val[3], thickness)
	return
}
func Circle3(img *Mat, center *Point, radius int, color *Scalar) {
	ImgprocNative_circle_2(img.GetNativeObjAddr(), center.X, center.Y, radius, color.Val[0], color.Val[1], color.Val[2], color.Val[3])
	return
}

func ClipLine(imgRect *Rect, pt1 *Point, pt2 *Point) bool {
	pt1_out := make([]float64, 2)
	pt2_out := make([]float64, 2)
	retVal := ImgprocNative_clipLine_0(imgRect.X, imgRect.Y, imgRect.Width, imgRect.Height, pt1.X, pt1.Y, pt1_out, pt2.X, pt2.Y, pt2_out)
	if pt1 != nil {
		pt1.X = pt1_out[0]
		pt1.Y = pt1_out[1]
	}
	if pt2 != nil {
		pt2.X = pt2_out[0]
		pt2.Y = pt2_out[1]
	}
	return retVal
}

func CompareHist(H1 *Mat, H2 *Mat, method int) float64 {
	retVal := ImgprocNative_compareHist_0(H1.GetNativeObjAddr(), H2.GetNativeObjAddr(), method)
	return retVal
}

func ConnectedComponents(image *Mat, labels *Mat, connectivity int, ltype int) int {
	retVal := ImgprocNative_connectedComponents_0(image.GetNativeObjAddr(), labels.GetNativeObjAddr(), connectivity, ltype)
	return retVal
}
func ConnectedComponents2(image *Mat, labels *Mat) int {
	retVal := ImgprocNative_connectedComponents_1(image.GetNativeObjAddr(), labels.GetNativeObjAddr())
	return retVal
}
func ConnectedComponentsWithAlgorithm(image *Mat, labels *Mat, connectivity int, ltype int, ccltype int) int {
	retVal := ImgprocNative_connectedComponentsWithAlgorithm_0(image.GetNativeObjAddr(), labels.GetNativeObjAddr(), connectivity, ltype, ccltype)
	return retVal
}

func ConnectedComponentsWithStats(image *Mat, labels *Mat, stats *Mat, centroids *Mat, connectivity int, ltype int) int {
	retVal := ImgprocNative_connectedComponentsWithStats_0(image.GetNativeObjAddr(), labels.GetNativeObjAddr(), stats.GetNativeObjAddr(), centroids.GetNativeObjAddr(), connectivity, ltype)
	return retVal
}
func ConnectedComponentsWithStats2(image *Mat, labels *Mat, stats *Mat, centroids *Mat) int {
	retVal := ImgprocNative_connectedComponentsWithStats_1(image.GetNativeObjAddr(), labels.GetNativeObjAddr(), stats.GetNativeObjAddr(), centroids.GetNativeObjAddr())
	return retVal
}
func ConnectedComponentsWithStatsWithAlgorithm(image *Mat, labels *Mat, stats *Mat, centroids *Mat, connectivity int, ltype int, ccltype int) int {
	retVal := ImgprocNative_connectedComponentsWithStatsWithAlgorithm_0(image.GetNativeObjAddr(), labels.GetNativeObjAddr(), stats.GetNativeObjAddr(), centroids.GetNativeObjAddr(), connectivity, ltype, ccltype)
	return retVal
}

func ContourArea(contour *Mat, oriented bool) float64 {
	retVal := ImgprocNative_contourArea_0(contour.GetNativeObjAddr(), oriented)
	return retVal
}
func ContourArea2(contour *Mat) float64 {
	retVal := ImgprocNative_contourArea_1(contour.GetNativeObjAddr())
	return retVal
}

func ConvertMaps(map1 *Mat, map2 *Mat, dstmap1 *Mat, dstmap2 *Mat, dstmap1type int, nninterpolation bool) {
	ImgprocNative_convertMaps_0(map1.GetNativeObjAddr(), map2.GetNativeObjAddr(), dstmap1.GetNativeObjAddr(), dstmap2.GetNativeObjAddr(), dstmap1type, nninterpolation)
	return
}
func ConvertMaps2(map1 *Mat, map2 *Mat, dstmap1 *Mat, dstmap2 *Mat, dstmap1type int) {
	ImgprocNative_convertMaps_1(map1.GetNativeObjAddr(), map2.GetNativeObjAddr(), dstmap1.GetNativeObjAddr(), dstmap2.GetNativeObjAddr(), dstmap1type)
	return
}

func ConvexHull(points *MatOfPoint, hull *MatOfInt, clockwise bool) {
	points_mat := points
	hull_mat := hull
	ImgprocNative_convexHull_0(points_mat.GetNativeObjAddr(), hull_mat.GetNativeObjAddr(), clockwise)
	return
}
func ConvexHull2(points *MatOfPoint, hull *MatOfInt) {
	points_mat := points
	hull_mat := hull
	ImgprocNative_convexHull_1(points_mat.GetNativeObjAddr(), hull_mat.GetNativeObjAddr())
	return
}

func ConvexityDefects(contour *MatOfPoint, convexhull *MatOfInt, convexityDefects *MatOfInt4) {
	contour_mat := contour
	convexhull_mat := convexhull
	convexityDefects_mat := convexityDefects
	ImgprocNative_convexityDefects_0(contour_mat.GetNativeObjAddr(), convexhull_mat.GetNativeObjAddr(), convexityDefects_mat.GetNativeObjAddr())
	return
}

func CornerEigenValsAndVecs(src *Mat, dst *Mat, blockSize int, ksize int, borderType int) {
	ImgprocNative_cornerEigenValsAndVecs_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), blockSize, ksize, borderType)
	return
}
func CornerEigenValsAndVecs2(src *Mat, dst *Mat, blockSize int, ksize int) {
	ImgprocNative_cornerEigenValsAndVecs_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), blockSize, ksize)
	return
}

func CornerHarris(src *Mat, dst *Mat, blockSize int, ksize int, k float64, borderType int) {
	ImgprocNative_cornerHarris_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), blockSize, ksize, k, borderType)
	return
}
func CornerHarris2(src *Mat, dst *Mat, blockSize int, ksize int, k float64) {
	ImgprocNative_cornerHarris_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), blockSize, ksize, k)
	return
}

func CornerMinEigenVal(src *Mat, dst *Mat, blockSize int, ksize int, borderType int) {
	ImgprocNative_cornerMinEigenVal_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), blockSize, ksize, borderType)
	return
}
func CornerMinEigenVal2(src *Mat, dst *Mat, blockSize int, ksize int) {
	ImgprocNative_cornerMinEigenVal_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), blockSize, ksize)
	return
}
func CornerMinEigenVal3(src *Mat, dst *Mat, blockSize int) {
	ImgprocNative_cornerMinEigenVal_2(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), blockSize)
	return
}

func CornerSubPix(image *Mat, corners *Mat, winSize *Size, zeroZone *Size, criteria *TermCriteria) {
	ImgprocNative_cornerSubPix_0(image.GetNativeObjAddr(), corners.GetNativeObjAddr(), winSize.Width, winSize.Height, zeroZone.Width, zeroZone.Height, criteria.Type, criteria.MaxCount, criteria.Epsilon)
	return
}

func CLAHECreate(clipLimit float64, tileGridSize *Size) *CLAHE {
	retVal := NewCLAHE(ImgprocNative_createCLAHE_0(clipLimit, tileGridSize.Width, tileGridSize.Height))
	return retVal
}
func CLAHECreate2() *CLAHE {
	retVal := NewCLAHE(ImgprocNative_createCLAHE_1())
	return retVal
}

func CreateHanningWindow(dst *Mat, winSize *Size, rtype int) {
	ImgprocNative_createHanningWindow_0(dst.GetNativeObjAddr(), winSize.Width, winSize.Height, rtype)
	return
}

func CreateLineSegmentDetector(_refine int, _scale float64, _sigma_scale float64, _quant float64, _ang_th float64, _log_eps float64, _density_th float64, _n_bins int) *LineSegmentDetector {
	retVal := NewLineSegmentDetector(ImgprocNative_createLineSegmentDetector_0(_refine, _scale, _sigma_scale, _quant, _ang_th, _log_eps, _density_th, _n_bins))
	return retVal
}
func CreateLineSegmentDetector2() *LineSegmentDetector {
	retVal := NewLineSegmentDetector(ImgprocNative_createLineSegmentDetector_1())
	return retVal
}

func CvtColor(src *Mat, dst *Mat, code int, dstCn int) {
	ImgprocNative_cvtColor_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), code, dstCn)
	return
}
func CvtColor2(src *Mat, dst *Mat, code int) {
	ImgprocNative_cvtColor_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), code)
	return
}

func Demosaicing(_src *Mat, _dst *Mat, code int, dcn int) {
	ImgprocNative_demosaicing_0(_src.GetNativeObjAddr(), _dst.GetNativeObjAddr(), code, dcn)
	return
}
func Demosaicing2(_src *Mat, _dst *Mat, code int) {
	ImgprocNative_demosaicing_1(_src.GetNativeObjAddr(), _dst.GetNativeObjAddr(), code)
	return
}

func Dilate(src *Mat, dst *Mat, kernel *Mat, anchor *Point, iterations int, borderType int, borderValue *Scalar) {
	ImgprocNative_dilate_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), kernel.GetNativeObjAddr(), anchor.X, anchor.Y, iterations, borderType, borderValue.Val[0], borderValue.Val[1], borderValue.Val[2], borderValue.Val[3])
	return
}
func Dilate2(src *Mat, dst *Mat, kernel *Mat, anchor *Point, iterations int) {
	ImgprocNative_dilate_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), kernel.GetNativeObjAddr(), anchor.X, anchor.Y, iterations)
	return
}
func Dilate3(src *Mat, dst *Mat, kernel *Mat) {
	ImgprocNative_dilate_2(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), kernel.GetNativeObjAddr())
	return
}

func DistanceTransform(src *Mat, dst *Mat, distanceType int, maskSize int, dstType int) {
	ImgprocNative_distanceTransform_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), distanceType, maskSize, dstType)
	return
}
func DistanceTransform2(src *Mat, dst *Mat, distanceType int, maskSize int) {
	ImgprocNative_distanceTransform_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), distanceType, maskSize)
	return
}
func DistanceTransformWithLabels(src *Mat, dst *Mat, labels *Mat, distanceType int, maskSize int, labelType int) {
	ImgprocNative_distanceTransformWithLabels_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), labels.GetNativeObjAddr(), distanceType, maskSize, labelType)
	return
}
func DistanceTransformWithLabels2(src *Mat, dst *Mat, labels *Mat, distanceType int, maskSize int) {
	ImgprocNative_distanceTransformWithLabels_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), labels.GetNativeObjAddr(), distanceType, maskSize)
	return
}

func DrawContours(image *Mat, contours []*MatOfPoint, contourIdx int, color *Scalar, thickness int, lineType int, hierarchy *Mat, maxLevel int, offset *Point) {
	contours_mat := ConvertersVectorVectorPoint(contours)
	ImgprocNative_drawContours_0(image.GetNativeObjAddr(), contours_mat.GetNativeObjAddr(), contourIdx, color.Val[0], color.Val[1], color.Val[2], color.Val[3], thickness, lineType, hierarchy.GetNativeObjAddr(), maxLevel, offset.X, offset.Y)
	return
}
func DrawContours2(image *Mat, contours []*MatOfPoint, contourIdx int, color *Scalar, thickness int) {
	contours_mat := ConvertersVectorVectorPoint(contours)
	ImgprocNative_drawContours_1(image.GetNativeObjAddr(), contours_mat.GetNativeObjAddr(), contourIdx, color.Val[0], color.Val[1], color.Val[2], color.Val[3], thickness)
	return
}
func DrawContours3(image *Mat, contours []*MatOfPoint, contourIdx int, color *Scalar) {
	contours_mat := ConvertersVectorVectorPoint(contours)
	ImgprocNative_drawContours_2(image.GetNativeObjAddr(), contours_mat.GetNativeObjAddr(), contourIdx, color.Val[0], color.Val[1], color.Val[2], color.Val[3])
	return
}

func DrawMarker(img *Mat, position *Point, color *Scalar, markerType int, markerSize int, thickness int, line_type int) {
	ImgprocNative_drawMarker_0(img.GetNativeObjAddr(), position.X, position.Y, color.Val[0], color.Val[1], color.Val[2], color.Val[3], markerType, markerSize, thickness, line_type)
	return
}
func DrawMarker2(img *Mat, position *Point, color *Scalar) {
	ImgprocNative_drawMarker_1(img.GetNativeObjAddr(), position.X, position.Y, color.Val[0], color.Val[1], color.Val[2], color.Val[3])
	return
}

func Ellipse(img *Mat, center *Point, axes *Size, angle float64, startAngle float64, endAngle float64, color *Scalar, thickness int, lineType int, shift int) {
	ImgprocNative_ellipse_0(img.GetNativeObjAddr(), center.X, center.Y, axes.Width, axes.Height, angle, startAngle, endAngle, color.Val[0], color.Val[1], color.Val[2], color.Val[3], thickness, lineType, shift)
	return
}
func Ellipse2(img *Mat, center *Point, axes *Size, angle float64, startAngle float64, endAngle float64, color *Scalar, thickness int) {
	ImgprocNative_ellipse_1(img.GetNativeObjAddr(), center.X, center.Y, axes.Width, axes.Height, angle, startAngle, endAngle, color.Val[0], color.Val[1], color.Val[2], color.Val[3], thickness)
	return
}
func Ellipse3(img *Mat, center *Point, axes *Size, angle float64, startAngle float64, endAngle float64, color *Scalar) {
	ImgprocNative_ellipse_2(img.GetNativeObjAddr(), center.X, center.Y, axes.Width, axes.Height, angle, startAngle, endAngle, color.Val[0], color.Val[1], color.Val[2], color.Val[3])
	return
}
func Ellipse4(img *Mat, box *RotatedRect, color *Scalar, thickness int, lineType int) {
	ImgprocNative_ellipse_3(img.GetNativeObjAddr(), box.Center.X, box.Center.Y, box.Size.Width, box.Size.Height, box.Angle, color.Val[0], color.Val[1], color.Val[2], color.Val[3], thickness, lineType)
	return
}
func Ellipse5(img *Mat, box *RotatedRect, color *Scalar, thickness int) {
	ImgprocNative_ellipse_4(img.GetNativeObjAddr(), box.Center.X, box.Center.Y, box.Size.Width, box.Size.Height, box.Angle, color.Val[0], color.Val[1], color.Val[2], color.Val[3], thickness)
	return
}
func Ellipse6(img *Mat, box *RotatedRect, color *Scalar) {
	ImgprocNative_ellipse_5(img.GetNativeObjAddr(), box.Center.X, box.Center.Y, box.Size.Width, box.Size.Height, box.Angle, color.Val[0], color.Val[1], color.Val[2], color.Val[3])
	return
}
func Ellipse2Poly(center *Point, axes *Size, angle int, arcStart int, arcEnd int, delta int, pts *MatOfPoint) {
	pts_mat := pts
	ImgprocNative_ellipse2Poly_0(center.X, center.Y, axes.Width, axes.Height, angle, arcStart, arcEnd, delta, pts_mat.GetNativeObjAddr())
	return
}

func EqualizeHist(src *Mat, dst *Mat) {
	ImgprocNative_equalizeHist_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr())
	return
}

func Erode(src *Mat, dst *Mat, kernel *Mat, anchor *Point, iterations int, borderType int, borderValue *Scalar) {
	ImgprocNative_erode_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), kernel.GetNativeObjAddr(), anchor.X, anchor.Y, iterations, borderType, borderValue.Val[0], borderValue.Val[1], borderValue.Val[2], borderValue.Val[3])
	return
}
func Erode2(src *Mat, dst *Mat, kernel *Mat, anchor *Point, iterations int) {
	ImgprocNative_erode_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), kernel.GetNativeObjAddr(), anchor.X, anchor.Y, iterations)
	return
}
func Erode3(src *Mat, dst *Mat, kernel *Mat) {
	ImgprocNative_erode_2(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), kernel.GetNativeObjAddr())
	return
}

func FillConvexPoly(img *Mat, points *MatOfPoint, color *Scalar, lineType int, shift int) {
	points_mat := points
	ImgprocNative_fillConvexPoly_0(img.GetNativeObjAddr(), points_mat.GetNativeObjAddr(), color.Val[0], color.Val[1], color.Val[2], color.Val[3], lineType, shift)
	return
}
func FillConvexPoly2(img *Mat, points *MatOfPoint, color *Scalar) {
	points_mat := points
	ImgprocNative_fillConvexPoly_1(img.GetNativeObjAddr(), points_mat.GetNativeObjAddr(), color.Val[0], color.Val[1], color.Val[2], color.Val[3])
	return
}

func FillPoly(img *Mat, pts []*MatOfPoint, color *Scalar, lineType int, shift int, offset *Point) {
	pts_mat := ConvertersVectorVectorPoint(pts)
	ImgprocNative_fillPoly_0(img.GetNativeObjAddr(), pts_mat.GetNativeObjAddr(), color.Val[0], color.Val[1], color.Val[2], color.Val[3], lineType, shift, offset.X, offset.Y)
	return
}
func FillPoly2(img *Mat, pts []*MatOfPoint, color *Scalar) {
	pts_mat := ConvertersVectorVectorPoint(pts)
	ImgprocNative_fillPoly_1(img.GetNativeObjAddr(), pts_mat.GetNativeObjAddr(), color.Val[0], color.Val[1], color.Val[2], color.Val[3])
	return
}

func Filter2D(src *Mat, dst *Mat, ddepth int, kernel *Mat, anchor *Point, delta float64, borderType int) {
	ImgprocNative_filter2D_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), ddepth, kernel.GetNativeObjAddr(), anchor.X, anchor.Y, delta, borderType)
	return
}
func Filter2D2(src *Mat, dst *Mat, ddepth int, kernel *Mat, anchor *Point, delta float64) {
	ImgprocNative_filter2D_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), ddepth, kernel.GetNativeObjAddr(), anchor.X, anchor.Y, delta)
	return
}
func Filter2D3(src *Mat, dst *Mat, ddepth int, kernel *Mat) {
	ImgprocNative_filter2D_2(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), ddepth, kernel.GetNativeObjAddr())
	return
}

func FindContours(image *Mat, hierarchy *Mat, mode int, method int, offset *Point) (contours []*MatOfPoint) {
	contours_mat := NewMat2()
	ImgprocNative_findContours_0(image.GetNativeObjAddr(), contours_mat.GetNativeObjAddr(), hierarchy.GetNativeObjAddr(), mode, method, offset.X, offset.Y)
	contours = ConvertersToVectorVectorPoint(contours_mat)
	contours_mat.Release()
	return
}
func FindContours2(image *Mat, hierarchy *Mat, mode int, method int) (contours []*MatOfPoint) {
	contours_mat := NewMat2()
	ImgprocNative_findContours_1(image.GetNativeObjAddr(), contours_mat.GetNativeObjAddr(), hierarchy.GetNativeObjAddr(), mode, method)
	contours = ConvertersToVectorVectorPoint(contours_mat)
	contours_mat.Release()
	return
}

func FitEllipse(points *MatOfPoint2f) *RotatedRect {
	points_mat := points
	retVal := NewRotatedRect3(ImgprocNative_fitEllipse_0(points_mat.GetNativeObjAddr()))
	return retVal
}
func FitEllipseAMS(points *Mat) *RotatedRect {
	retVal := NewRotatedRect3(ImgprocNative_fitEllipseAMS_0(points.GetNativeObjAddr()))
	return retVal
}

func FitEllipseDirect(points *Mat) *RotatedRect {
	retVal := NewRotatedRect3(ImgprocNative_fitEllipseDirect_0(points.GetNativeObjAddr()))
	return retVal
}

func FitLine(points *Mat, line *Mat, distType int, param float64, reps float64, aeps float64) {
	ImgprocNative_fitLine_0(points.GetNativeObjAddr(), line.GetNativeObjAddr(), distType, param, reps, aeps)
	return
}

func FloodFill(image *Mat, mask *Mat, seedPoint *Point, newVal *Scalar, rect *Rect, loDiff *Scalar, upDiff *Scalar, flags int) int {
	rect_out := make([]float64, 4)
	retVal := ImgprocNative_floodFill_0(image.GetNativeObjAddr(), mask.GetNativeObjAddr(), seedPoint.X, seedPoint.Y, newVal.Val[0], newVal.Val[1], newVal.Val[2], newVal.Val[3], rect_out, loDiff.Val[0], loDiff.Val[1], loDiff.Val[2], loDiff.Val[3], upDiff.Val[0], upDiff.Val[1], upDiff.Val[2], upDiff.Val[3], flags)
	if rect != nil {
		rect.X = int(rect_out[0])
		rect.Y = int(rect_out[1])
		rect.Width = int(rect_out[2])
		rect.Height = int(rect_out[3])
	}
	return retVal
}
func FloodFill2(image *Mat, mask *Mat, seedPoint *Point, newVal *Scalar) int {
	retVal := ImgprocNative_floodFill_1(image.GetNativeObjAddr(), mask.GetNativeObjAddr(), seedPoint.X, seedPoint.Y, newVal.Val[0], newVal.Val[1], newVal.Val[2], newVal.Val[3])
	return retVal
}

func GetAffineTransform(src *MatOfPoint2f, dst *MatOfPoint2f) *Mat {
	src_mat := src
	dst_mat := dst
	retVal := NewMat(ImgprocNative_getAffineTransform_0(src_mat.GetNativeObjAddr(), dst_mat.GetNativeObjAddr()))
	return retVal
}

func GetDefaultNewCameraMatrix(cameraMatrix *Mat, imgsize *Size, centerPrincipalPoint bool) *Mat {
	retVal := NewMat(ImgprocNative_getDefaultNewCameraMatrix_0(cameraMatrix.GetNativeObjAddr(), imgsize.Width, imgsize.Height, centerPrincipalPoint))
	return retVal
}
func GetDefaultNewCameraMatrix2(cameraMatrix *Mat) *Mat {
	retVal := NewMat(ImgprocNative_getDefaultNewCameraMatrix_1(cameraMatrix.GetNativeObjAddr()))
	return retVal
}

func GetDerivKernels(kx *Mat, ky *Mat, dx int, dy int, ksize int, normalize bool, ktype int) {
	ImgprocNative_getDerivKernels_0(kx.GetNativeObjAddr(), ky.GetNativeObjAddr(), dx, dy, ksize, normalize, ktype)
	return
}
func GetDerivKernels2(kx *Mat, ky *Mat, dx int, dy int, ksize int) {
	ImgprocNative_getDerivKernels_1(kx.GetNativeObjAddr(), ky.GetNativeObjAddr(), dx, dy, ksize)
	return
}

func GetGaborKernel(ksize *Size, sigma float64, theta float64, lambd float64, gamma float64, psi float64, ktype int) *Mat {
	retVal := NewMat(ImgprocNative_getGaborKernel_0(ksize.Width, ksize.Height, sigma, theta, lambd, gamma, psi, ktype))
	return retVal
}
func GetGaborKernel2(ksize *Size, sigma float64, theta float64, lambd float64, gamma float64) *Mat {
	retVal := NewMat(ImgprocNative_getGaborKernel_1(ksize.Width, ksize.Height, sigma, theta, lambd, gamma))
	return retVal
}

func GetGaussianKernel(ksize int, sigma float64, ktype int) *Mat {
	retVal := NewMat(ImgprocNative_getGaussianKernel_0(ksize, sigma, ktype))
	return retVal
}
func GetGaussianKernel2(ksize int, sigma float64) *Mat {
	retVal := NewMat(ImgprocNative_getGaussianKernel_1(ksize, sigma))
	return retVal
}

func GetPerspectiveTransform(src *Mat, dst *Mat) *Mat {
	retVal := NewMat(ImgprocNative_getPerspectiveTransform_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr()))
	return retVal
}

func GetRectSubPix(image *Mat, patchSize *Size, center *Point, patch *Mat, patchType int) {
	ImgprocNative_getRectSubPix_0(image.GetNativeObjAddr(), patchSize.Width, patchSize.Height, center.X, center.Y, patch.GetNativeObjAddr(), patchType)
	return
}
func GetRectSubPix2(image *Mat, patchSize *Size, center *Point, patch *Mat) {
	ImgprocNative_getRectSubPix_1(image.GetNativeObjAddr(), patchSize.Width, patchSize.Height, center.X, center.Y, patch.GetNativeObjAddr())
	return
}

func GetRotationMatrix2D(center *Point, angle float64, scale float64) *Mat {
	retVal := NewMat(ImgprocNative_getRotationMatrix2D_0(center.X, center.Y, angle, scale))
	return retVal
}

func GetStructuringElement(shape int, ksize *Size, anchor *Point) *Mat {
	retVal := NewMat(ImgprocNative_getStructuringElement_0(shape, ksize.Width, ksize.Height, anchor.X, anchor.Y))
	return retVal
}
func GetStructuringElement2(shape int, ksize *Size) *Mat {
	retVal := NewMat(ImgprocNative_getStructuringElement_1(shape, ksize.Width, ksize.Height))
	return retVal
}

func GetTextSize(text string, fontFace int, fontScale float64, thickness int, baseLine []int32) *Size {
	if baseLine != nil && len(baseLine) != 1 {
		Throw(NewIllegalArgumentException("'baseLine' must be 'int[1]' or 'null'."))
	}
	retVal := NewSize4(ImgprocNative_n_getTextSize(text, fontFace, fontScale, thickness, baseLine))
	return retVal
}
func GoodFeaturesToTrack(image *Mat, corners *MatOfPoint, maxCorners int, qualityLevel float64, minDistance float64, mask *Mat, blockSize int, gradientSize int, useHarrisDetector bool, k float64) {
	corners_mat := corners
	ImgprocNative_goodFeaturesToTrack_0(image.GetNativeObjAddr(), corners_mat.GetNativeObjAddr(), maxCorners, qualityLevel, minDistance, mask.GetNativeObjAddr(), blockSize, gradientSize, useHarrisDetector, k)
	return
}
func GoodFeaturesToTrack2(image *Mat, corners *MatOfPoint, maxCorners int, qualityLevel float64, minDistance float64, mask *Mat, blockSize int, gradientSize int) {
	corners_mat := corners
	ImgprocNative_goodFeaturesToTrack_1(image.GetNativeObjAddr(), corners_mat.GetNativeObjAddr(), maxCorners, qualityLevel, minDistance, mask.GetNativeObjAddr(), blockSize, gradientSize)
	return
}
func GoodFeaturesToTrack3(image *Mat, corners *MatOfPoint, maxCorners int, qualityLevel float64, minDistance float64, mask *Mat, blockSize int, useHarrisDetector bool, k float64) {
	corners_mat := corners
	ImgprocNative_goodFeaturesToTrack_2(image.GetNativeObjAddr(), corners_mat.GetNativeObjAddr(), maxCorners, qualityLevel, minDistance, mask.GetNativeObjAddr(), blockSize, useHarrisDetector, k)
	return
}
func GoodFeaturesToTrack4(image *Mat, corners *MatOfPoint, maxCorners int, qualityLevel float64, minDistance float64) {
	corners_mat := corners
	ImgprocNative_goodFeaturesToTrack_3(image.GetNativeObjAddr(), corners_mat.GetNativeObjAddr(), maxCorners, qualityLevel, minDistance)
	return
}

func GrabCut(img *Mat, mask *Mat, rect *Rect, bgdModel *Mat, fgdModel *Mat, iterCount int, mode int) {
	ImgprocNative_grabCut_0(img.GetNativeObjAddr(), mask.GetNativeObjAddr(), rect.X, rect.Y, rect.Width, rect.Height, bgdModel.GetNativeObjAddr(), fgdModel.GetNativeObjAddr(), iterCount, mode)
	return
}
func GrabCut2(img *Mat, mask *Mat, rect *Rect, bgdModel *Mat, fgdModel *Mat, iterCount int) {
	ImgprocNative_grabCut_1(img.GetNativeObjAddr(), mask.GetNativeObjAddr(), rect.X, rect.Y, rect.Width, rect.Height, bgdModel.GetNativeObjAddr(), fgdModel.GetNativeObjAddr(), iterCount)
	return
}

func InitUndistortRectifyMap(cameraMatrix *Mat, distCoeffs *Mat, R *Mat, newCameraMatrix *Mat, size *Size, m1type int, map1 *Mat, map2 *Mat) {
	ImgprocNative_initUndistortRectifyMap_0(cameraMatrix.GetNativeObjAddr(), distCoeffs.GetNativeObjAddr(), R.GetNativeObjAddr(), newCameraMatrix.GetNativeObjAddr(), size.Width, size.Height, m1type, map1.GetNativeObjAddr(), map2.GetNativeObjAddr())
	return
}

func InitWideAngleProjMap(cameraMatrix *Mat, distCoeffs *Mat, imageSize *Size, destImageWidth int, m1type int, map1 *Mat, map2 *Mat, projType int, alpha float64) float32 {
	retVal := ImgprocNative_initWideAngleProjMap_0(cameraMatrix.GetNativeObjAddr(), distCoeffs.GetNativeObjAddr(), imageSize.Width, imageSize.Height, destImageWidth, m1type, map1.GetNativeObjAddr(), map2.GetNativeObjAddr(), projType, alpha)
	return retVal
}
func InitWideAngleProjMap2(cameraMatrix *Mat, distCoeffs *Mat, imageSize *Size, destImageWidth int, m1type int, map1 *Mat, map2 *Mat) float32 {
	retVal := ImgprocNative_initWideAngleProjMap_1(cameraMatrix.GetNativeObjAddr(), distCoeffs.GetNativeObjAddr(), imageSize.Width, imageSize.Height, destImageWidth, m1type, map1.GetNativeObjAddr(), map2.GetNativeObjAddr())
	return retVal
}

func Integral(src *Mat, sum *Mat, sdepth int) {
	ImgprocNative_integral_0(src.GetNativeObjAddr(), sum.GetNativeObjAddr(), sdepth)
	return
}
func Integral2(src *Mat, sum *Mat) {
	ImgprocNative_integral_1(src.GetNativeObjAddr(), sum.GetNativeObjAddr())
	return
}
func Integral21(src *Mat, sum *Mat, sqsum *Mat, sdepth int, sqdepth int) {
	ImgprocNative_integral2_0(src.GetNativeObjAddr(), sum.GetNativeObjAddr(), sqsum.GetNativeObjAddr(), sdepth, sqdepth)
	return
}
func Integral22(src *Mat, sum *Mat, sqsum *Mat) {
	ImgprocNative_integral2_1(src.GetNativeObjAddr(), sum.GetNativeObjAddr(), sqsum.GetNativeObjAddr())
	return
}

func Integral3(src *Mat, sum *Mat, sqsum *Mat, tilted *Mat, sdepth int, sqdepth int) {
	ImgprocNative_integral3_0(src.GetNativeObjAddr(), sum.GetNativeObjAddr(), sqsum.GetNativeObjAddr(), tilted.GetNativeObjAddr(), sdepth, sqdepth)
	return
}
func Integral32(src *Mat, sum *Mat, sqsum *Mat, tilted *Mat) {
	ImgprocNative_integral3_1(src.GetNativeObjAddr(), sum.GetNativeObjAddr(), sqsum.GetNativeObjAddr(), tilted.GetNativeObjAddr())
	return
}

func IntersectConvexConvex(_p1 *Mat, _p2 *Mat, _p12 *Mat, handleNested bool) float32 {
	retVal := ImgprocNative_intersectConvexConvex_0(_p1.GetNativeObjAddr(), _p2.GetNativeObjAddr(), _p12.GetNativeObjAddr(), handleNested)
	return retVal
}
func IntersectConvexConvex2(_p1 *Mat, _p2 *Mat, _p12 *Mat) float32 {
	retVal := ImgprocNative_intersectConvexConvex_1(_p1.GetNativeObjAddr(), _p2.GetNativeObjAddr(), _p12.GetNativeObjAddr())
	return retVal
}

func InvertAffineTransform(M *Mat, iM *Mat) {
	ImgprocNative_invertAffineTransform_0(M.GetNativeObjAddr(), iM.GetNativeObjAddr())
	return
}

func IsContourConvex(contour *MatOfPoint) bool {
	contour_mat := contour
	retVal := ImgprocNative_isContourConvex_0(contour_mat.GetNativeObjAddr())
	return retVal
}

func Line(img *Mat, pt1 *Point, pt2 *Point, color *Scalar, thickness int, lineType int, shift int) {
	ImgprocNative_line_0(img.GetNativeObjAddr(), pt1.X, pt1.Y, pt2.X, pt2.Y, color.Val[0], color.Val[1], color.Val[2], color.Val[3], thickness, lineType, shift)
	return
}
func Line2(img *Mat, pt1 *Point, pt2 *Point, color *Scalar, thickness int) {
	ImgprocNative_line_1(img.GetNativeObjAddr(), pt1.X, pt1.Y, pt2.X, pt2.Y, color.Val[0], color.Val[1], color.Val[2], color.Val[3], thickness)
	return
}
func Line3(img *Mat, pt1 *Point, pt2 *Point, color *Scalar) {
	ImgprocNative_line_2(img.GetNativeObjAddr(), pt1.X, pt1.Y, pt2.X, pt2.Y, color.Val[0], color.Val[1], color.Val[2], color.Val[3])
	return
}

func LinearPolar(src *Mat, dst *Mat, center *Point, maxRadius float64, flags int) {
	ImgprocNative_linearPolar_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), center.X, center.Y, maxRadius, flags)
	return
}

func LogPolar(src *Mat, dst *Mat, center *Point, M float64, flags int) {
	ImgprocNative_logPolar_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), center.X, center.Y, M, flags)
	return
}

func MatchShapes(contour1 *Mat, contour2 *Mat, method int, parameter float64) float64 {
	retVal := ImgprocNative_matchShapes_0(contour1.GetNativeObjAddr(), contour2.GetNativeObjAddr(), method, parameter)
	return retVal
}

func MatchTemplate(image *Mat, templ *Mat, result *Mat, method int, mask *Mat) {
	ImgprocNative_matchTemplate_0(image.GetNativeObjAddr(), templ.GetNativeObjAddr(), result.GetNativeObjAddr(), method, mask.GetNativeObjAddr())
	return
}
func MatchTemplate2(image *Mat, templ *Mat, result *Mat, method int) {
	ImgprocNative_matchTemplate_1(image.GetNativeObjAddr(), templ.GetNativeObjAddr(), result.GetNativeObjAddr(), method)
	return
}

func MedianBlur(src *Mat, dst *Mat, ksize int) {
	ImgprocNative_medianBlur_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), ksize)
	return
}

func MinAreaRect(points *MatOfPoint2f) *RotatedRect {
	points_mat := points
	retVal := NewRotatedRect3(ImgprocNative_minAreaRect_0(points_mat.GetNativeObjAddr()))
	return retVal
}

func MinEnclosingCircle(points *MatOfPoint2f, center *Point, radius []float32) {
	points_mat := points
	center_out := make([]float64, 2)
	radius_out := make([]float64, 1)
	ImgprocNative_minEnclosingCircle_0(points_mat.GetNativeObjAddr(), center_out, radius_out)
	if center != nil {
		center.X = center_out[0]
		center.Y = center_out[1]
	}
	if radius != nil {
		radius[0] = float32(radius_out[0])
	}
	return
}

func MinEnclosingTriangle(points *Mat, triangle *Mat) float64 {
	retVal := ImgprocNative_minEnclosingTriangle_0(points.GetNativeObjAddr(), triangle.GetNativeObjAddr())
	return retVal
}

func MomentsCreate(array *Mat, binaryImage bool) *Moments {
	retVal := NewMoments3(ImgprocNative_moments_0(array.GetNativeObjAddr(), binaryImage))
	return retVal
}
func Moments2(array *Mat) *Moments {
	retVal := NewMoments3(ImgprocNative_moments_1(array.GetNativeObjAddr()))
	return retVal
}

func MorphologyEx(src *Mat, dst *Mat, op int, kernel *Mat, anchor *Point, iterations int, borderType int, borderValue *Scalar) {
	ImgprocNative_morphologyEx_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), op, kernel.GetNativeObjAddr(), anchor.X, anchor.Y, iterations, borderType, borderValue.Val[0], borderValue.Val[1], borderValue.Val[2], borderValue.Val[3])
	return
}
func MorphologyEx2(src *Mat, dst *Mat, op int, kernel *Mat, anchor *Point, iterations int) {
	ImgprocNative_morphologyEx_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), op, kernel.GetNativeObjAddr(), anchor.X, anchor.Y, iterations)
	return
}
func MorphologyEx3(src *Mat, dst *Mat, op int, kernel *Mat) {
	ImgprocNative_morphologyEx_2(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), op, kernel.GetNativeObjAddr())
	return
}

func PhaseCorrelate(src1 *Mat, src2 *Mat, window *Mat, response []float64) *Point {
	response_out := make([]float64, 1)
	retVal := NewPoint3(ImgprocNative_phaseCorrelate_0(src1.GetNativeObjAddr(), src2.GetNativeObjAddr(), window.GetNativeObjAddr(), response_out))
	if response != nil {
		response[0] = response_out[0]
	}
	return retVal
}
func PhaseCorrelate2(src1 *Mat, src2 *Mat) *Point {
	retVal := NewPoint3(ImgprocNative_phaseCorrelate_1(src1.GetNativeObjAddr(), src2.GetNativeObjAddr()))
	return retVal
}

func PointPolygonTest(contour *MatOfPoint2f, pt *Point, measureDist bool) float64 {
	contour_mat := contour
	retVal := ImgprocNative_pointPolygonTest_0(contour_mat.GetNativeObjAddr(), pt.X, pt.Y, measureDist)
	return retVal
}

func Polylines(img *Mat, pts []*MatOfPoint, isClosed bool, color *Scalar, thickness int, lineType int, shift int) {
	pts_mat := ConvertersVectorVectorPoint(pts)
	ImgprocNative_polylines_0(img.GetNativeObjAddr(), pts_mat.GetNativeObjAddr(), isClosed, color.Val[0], color.Val[1], color.Val[2], color.Val[3], thickness, lineType, shift)
	return
}
func Polylines2(img *Mat, pts []*MatOfPoint, isClosed bool, color *Scalar, thickness int) {
	pts_mat := ConvertersVectorVectorPoint(pts)
	ImgprocNative_polylines_1(img.GetNativeObjAddr(), pts_mat.GetNativeObjAddr(), isClosed, color.Val[0], color.Val[1], color.Val[2], color.Val[3], thickness)
	return
}
func Polylines3(img *Mat, pts []*MatOfPoint, isClosed bool, color *Scalar) {
	pts_mat := ConvertersVectorVectorPoint(pts)
	ImgprocNative_polylines_2(img.GetNativeObjAddr(), pts_mat.GetNativeObjAddr(), isClosed, color.Val[0], color.Val[1], color.Val[2], color.Val[3])
	return
}

func PreCornerDetect(src *Mat, dst *Mat, ksize int, borderType int) {
	ImgprocNative_preCornerDetect_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), ksize, borderType)
	return
}
func PreCornerDetect2(src *Mat, dst *Mat, ksize int) {
	ImgprocNative_preCornerDetect_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), ksize)
	return
}

func PutText(img *Mat, text string, org *Point, fontFace int, fontScale float64, color *Scalar, thickness int, lineType int, bottomLeftOrigin bool) {
	ImgprocNative_putText_0(img.GetNativeObjAddr(), text, org.X, org.Y, fontFace, fontScale, color.Val[0], color.Val[1], color.Val[2], color.Val[3], thickness, lineType, bottomLeftOrigin)
	return
}
func PutText2(img *Mat, text string, org *Point, fontFace int, fontScale float64, color *Scalar, thickness int) {
	ImgprocNative_putText_1(img.GetNativeObjAddr(), text, org.X, org.Y, fontFace, fontScale, color.Val[0], color.Val[1], color.Val[2], color.Val[3], thickness)
	return
}
func PutText3(img *Mat, text string, org *Point, fontFace int, fontScale float64, color *Scalar) {
	ImgprocNative_putText_2(img.GetNativeObjAddr(), text, org.X, org.Y, fontFace, fontScale, color.Val[0], color.Val[1], color.Val[2], color.Val[3])
	return
}

func PyrDown(src *Mat, dst *Mat, dstsize *Size, borderType int) {
	ImgprocNative_pyrDown_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), dstsize.Width, dstsize.Height, borderType)
	return
}
func PyrDown2(src *Mat, dst *Mat, dstsize *Size) {
	ImgprocNative_pyrDown_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), dstsize.Width, dstsize.Height)
	return
}
func PyrDown3(src *Mat, dst *Mat) {
	ImgprocNative_pyrDown_2(src.GetNativeObjAddr(), dst.GetNativeObjAddr())
	return
}

func PyrMeanShiftFiltering(src *Mat, dst *Mat, sp float64, sr float64, maxLevel int, termcrit *TermCriteria) {
	ImgprocNative_pyrMeanShiftFiltering_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), sp, sr, maxLevel, termcrit.Type, termcrit.MaxCount, termcrit.Epsilon)
	return
}
func PyrMeanShiftFiltering2(src *Mat, dst *Mat, sp float64, sr float64) {
	ImgprocNative_pyrMeanShiftFiltering_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), sp, sr)
	return
}

func PyrUp(src *Mat, dst *Mat, dstsize *Size, borderType int) {
	ImgprocNative_pyrUp_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), dstsize.Width, dstsize.Height, borderType)
	return
}
func PyrUp2(src *Mat, dst *Mat, dstsize *Size) {
	ImgprocNative_pyrUp_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), dstsize.Width, dstsize.Height)
	return
}
func PyrUp3(src *Mat, dst *Mat) {
	ImgprocNative_pyrUp_2(src.GetNativeObjAddr(), dst.GetNativeObjAddr())
	return
}

func Rectangle(img *Mat, pt1 *Point, pt2 *Point, color *Scalar, thickness int, lineType int, shift int) {
	ImgprocNative_rectangle_0(img.GetNativeObjAddr(), pt1.X, pt1.Y, pt2.X, pt2.Y, color.Val[0], color.Val[1], color.Val[2], color.Val[3], thickness, lineType, shift)
	return
}
func Rectangle2(img *Mat, pt1 *Point, pt2 *Point, color *Scalar, thickness int) {
	ImgprocNative_rectangle_1(img.GetNativeObjAddr(), pt1.X, pt1.Y, pt2.X, pt2.Y, color.Val[0], color.Val[1], color.Val[2], color.Val[3], thickness)
	return
}
func Rectangle3(img *Mat, pt1 *Point, pt2 *Point, color *Scalar) {
	ImgprocNative_rectangle_2(img.GetNativeObjAddr(), pt1.X, pt1.Y, pt2.X, pt2.Y, color.Val[0], color.Val[1], color.Val[2], color.Val[3])
	return
}

func Remap(src *Mat, dst *Mat, map1 *Mat, map2 *Mat, interpolation int, borderMode int, borderValue *Scalar) {
	ImgprocNative_remap_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), map1.GetNativeObjAddr(), map2.GetNativeObjAddr(), interpolation, borderMode, borderValue.Val[0], borderValue.Val[1], borderValue.Val[2], borderValue.Val[3])
	return
}
func Remap2(src *Mat, dst *Mat, map1 *Mat, map2 *Mat, interpolation int) {
	ImgprocNative_remap_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), map1.GetNativeObjAddr(), map2.GetNativeObjAddr(), interpolation)
	return
}

func Resize(src *Mat, dst *Mat, dsize *Size, fx float64, fy float64, interpolation int) {
	ImgprocNative_resize_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), dsize.Width, dsize.Height, fx, fy, interpolation)
	return
}
func Resize2(src *Mat, dst *Mat, dsize *Size) {
	ImgprocNative_resize_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), dsize.Width, dsize.Height)
	return
}

func RotatedRectangleIntersection(rect1 *RotatedRect, rect2 *RotatedRect, intersectingRegion *Mat) int {
	retVal := ImgprocNative_rotatedRectangleIntersection_0(rect1.Center.X, rect1.Center.Y, rect1.Size.Width, rect1.Size.Height, rect1.Angle, rect2.Center.X, rect2.Center.Y, rect2.Size.Width, rect2.Size.Height, rect2.Angle, intersectingRegion.GetNativeObjAddr())
	return retVal
}

func SepFilter2D(src *Mat, dst *Mat, ddepth int, kernelX *Mat, kernelY *Mat, anchor *Point, delta float64, borderType int) {
	ImgprocNative_sepFilter2D_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), ddepth, kernelX.GetNativeObjAddr(), kernelY.GetNativeObjAddr(), anchor.X, anchor.Y, delta, borderType)
	return
}
func SepFilter2D2(src *Mat, dst *Mat, ddepth int, kernelX *Mat, kernelY *Mat, anchor *Point, delta float64) {
	ImgprocNative_sepFilter2D_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), ddepth, kernelX.GetNativeObjAddr(), kernelY.GetNativeObjAddr(), anchor.X, anchor.Y, delta)
	return
}
func SepFilter2D3(src *Mat, dst *Mat, ddepth int, kernelX *Mat, kernelY *Mat) {
	ImgprocNative_sepFilter2D_2(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), ddepth, kernelX.GetNativeObjAddr(), kernelY.GetNativeObjAddr())
	return
}

func SpatialGradient(src *Mat, dx *Mat, dy *Mat, ksize int, borderType int) {
	ImgprocNative_spatialGradient_0(src.GetNativeObjAddr(), dx.GetNativeObjAddr(), dy.GetNativeObjAddr(), ksize, borderType)
	return
}
func SpatialGradient2(src *Mat, dx *Mat, dy *Mat, ksize int) {
	ImgprocNative_spatialGradient_1(src.GetNativeObjAddr(), dx.GetNativeObjAddr(), dy.GetNativeObjAddr(), ksize)
	return
}
func SpatialGradient3(src *Mat, dx *Mat, dy *Mat) {
	ImgprocNative_spatialGradient_2(src.GetNativeObjAddr(), dx.GetNativeObjAddr(), dy.GetNativeObjAddr())
	return
}

func SqrBoxFilter(_src *Mat, _dst *Mat, ddepth int, ksize *Size, anchor *Point, normalize bool, borderType int) {
	ImgprocNative_sqrBoxFilter_0(_src.GetNativeObjAddr(), _dst.GetNativeObjAddr(), ddepth, ksize.Width, ksize.Height, anchor.X, anchor.Y, normalize, borderType)
	return
}
func SqrBoxFilter2(_src *Mat, _dst *Mat, ddepth int, ksize *Size, anchor *Point, normalize bool) {
	ImgprocNative_sqrBoxFilter_1(_src.GetNativeObjAddr(), _dst.GetNativeObjAddr(), ddepth, ksize.Width, ksize.Height, anchor.X, anchor.Y, normalize)
	return
}
func SqrBoxFilter3(_src *Mat, _dst *Mat, ddepth int, ksize *Size) {
	ImgprocNative_sqrBoxFilter_2(_src.GetNativeObjAddr(), _dst.GetNativeObjAddr(), ddepth, ksize.Width, ksize.Height)
	return
}

func Threshold(src *Mat, dst *Mat, thresh float64, maxval float64, rtype int) float64 {
	retVal := ImgprocNative_threshold_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), thresh, maxval, rtype)
	return retVal
}

func Undistort(src *Mat, dst *Mat, cameraMatrix *Mat, distCoeffs *Mat, newCameraMatrix *Mat) {
	ImgprocNative_undistort_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), cameraMatrix.GetNativeObjAddr(), distCoeffs.GetNativeObjAddr(), newCameraMatrix.GetNativeObjAddr())
	return
}
func Undistort2(src *Mat, dst *Mat, cameraMatrix *Mat, distCoeffs *Mat) {
	ImgprocNative_undistort_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), cameraMatrix.GetNativeObjAddr(), distCoeffs.GetNativeObjAddr())
	return
}
func UndistortPoints(src *Mat, dst *Mat, cameraMatrix *Mat, distCoeffs *Mat, R *Mat, P *Mat) {
	ImgprocNative_undistortPoints_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), cameraMatrix.GetNativeObjAddr(), distCoeffs.GetNativeObjAddr(), R.GetNativeObjAddr(), P.GetNativeObjAddr())
	return
}
func UndistortPoints2(src *Mat, dst *Mat, cameraMatrix *Mat, distCoeffs *Mat) {
	ImgprocNative_undistortPoints_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), cameraMatrix.GetNativeObjAddr(), distCoeffs.GetNativeObjAddr())
	return
}
func UndistortPointsIter(src *Mat, dst *Mat, cameraMatrix *Mat, distCoeffs *Mat, R *Mat, P *Mat, criteria *TermCriteria) {
	ImgprocNative_undistortPointsIter_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), cameraMatrix.GetNativeObjAddr(), distCoeffs.GetNativeObjAddr(), R.GetNativeObjAddr(), P.GetNativeObjAddr(), criteria.Type, criteria.MaxCount, criteria.Epsilon)
	return
}

func WarpAffine(src *Mat, dst *Mat, M *Mat, dsize *Size, flags int, borderMode int, borderValue *Scalar) {
	ImgprocNative_warpAffine_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), M.GetNativeObjAddr(), dsize.Width, dsize.Height, flags, borderMode, borderValue.Val[0], borderValue.Val[1], borderValue.Val[2], borderValue.Val[3])
	return
}
func WarpAffine2(src *Mat, dst *Mat, M *Mat, dsize *Size, flags int) {
	ImgprocNative_warpAffine_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), M.GetNativeObjAddr(), dsize.Width, dsize.Height, flags)
	return
}
func WarpAffine3(src *Mat, dst *Mat, M *Mat, dsize *Size) {
	ImgprocNative_warpAffine_2(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), M.GetNativeObjAddr(), dsize.Width, dsize.Height)
	return
}

func WarpPerspective(src *Mat, dst *Mat, M *Mat, dsize *Size, flags int, borderMode int, borderValue *Scalar) {
	ImgprocNative_warpPerspective_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), M.GetNativeObjAddr(), dsize.Width, dsize.Height, flags, borderMode, borderValue.Val[0], borderValue.Val[1], borderValue.Val[2], borderValue.Val[3])
	return
}
func WarpPerspective2(src *Mat, dst *Mat, M *Mat, dsize *Size, flags int) {
	ImgprocNative_warpPerspective_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), M.GetNativeObjAddr(), dsize.Width, dsize.Height, flags)
	return
}
func WarpPerspective3(src *Mat, dst *Mat, M *Mat, dsize *Size) {
	ImgprocNative_warpPerspective_2(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), M.GetNativeObjAddr(), dsize.Width, dsize.Height)
	return
}

func Watershed(image *Mat, markers *Mat) {
	ImgprocNative_watershed_0(image.GetNativeObjAddr(), markers.GetNativeObjAddr())
	return
}
