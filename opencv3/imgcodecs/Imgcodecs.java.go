package imgcodecs

import (
	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

const CV_LOAD_IMAGE_GRAYSCALE = 0
const CV_LOAD_IMAGE_COLOR = 1
const CV_LOAD_IMAGE_ANYDEPTH = 2
const CV_LOAD_IMAGE_ANYCOLOR = 4
const CV_LOAD_IMAGE_IGNORE_ORIENTATION = 128
const CV_IMWRITE_JPEG_QUALITY = 1
const CV_IMWRITE_JPEG_PROGRESSIVE = 2
const CV_IMWRITE_JPEG_OPTIMIZE = 3
const CV_IMWRITE_JPEG_RST_INTERVAL = 4
const CV_IMWRITE_JPEG_LUMA_QUALITY = 5
const CV_IMWRITE_JPEG_CHROMA_QUALITY = 6
const CV_IMWRITE_PNG_COMPRESSION = 16
const CV_IMWRITE_PNG_STRATEGY = 17
const CV_IMWRITE_PNG_BILEVEL = 18
const CV_IMWRITE_PNG_STRATEGY_DEFAULT = 0
const CV_IMWRITE_PNG_STRATEGY_FILTERED = 1
const CV_IMWRITE_PNG_STRATEGY_HUFFMAN_ONLY = 2
const CV_IMWRITE_PNG_STRATEGY_RLE = 3
const CV_IMWRITE_PNG_STRATEGY_FIXED = 4
const CV_IMWRITE_PXM_BINARY = 32
const CV_IMWRITE_WEBP_QUALITY = 64
const CV_IMWRITE_PAM_TUPLETYPE = 128
const CV_IMWRITE_PAM_FORMAT_NULL = 0
const CV_IMWRITE_PAM_FORMAT_BLACKANDWHITE = 1
const CV_IMWRITE_PAM_FORMAT_GRAYSCALE = 2
const CV_IMWRITE_PAM_FORMAT_GRAYSCALE_ALPHA = 3
const CV_IMWRITE_PAM_FORMAT_RGB = 4
const CV_IMWRITE_PAM_FORMAT_RGB_ALPHA = 5
const CV_CVTIMG_FLIP = 1
const CV_CVTIMG_SWAP_RB = 2
const IMREAD_GRAYSCALE = 0
const IMREAD_COLOR = 1
const IMREAD_ANYDEPTH = 2
const IMREAD_ANYCOLOR = 4
const IMREAD_LOAD_GDAL = 8
const IMREAD_REDUCED_GRAYSCALE_2 = 16
const IMREAD_REDUCED_COLOR_2 = 17
const IMREAD_REDUCED_GRAYSCALE_4 = 32
const IMREAD_REDUCED_COLOR_4 = 33
const IMREAD_REDUCED_GRAYSCALE_8 = 64
const IMREAD_REDUCED_COLOR_8 = 65
const IMREAD_IGNORE_ORIENTATION = 128
const IMWRITE_JPEG_QUALITY = 1
const IMWRITE_JPEG_PROGRESSIVE = 2
const IMWRITE_JPEG_OPTIMIZE = 3
const IMWRITE_JPEG_RST_INTERVAL = 4
const IMWRITE_JPEG_LUMA_QUALITY = 5
const IMWRITE_JPEG_CHROMA_QUALITY = 6
const IMWRITE_PNG_COMPRESSION = 16
const IMWRITE_PNG_STRATEGY = 17
const IMWRITE_PNG_BILEVEL = 18
const IMWRITE_PXM_BINARY = 32
const IMWRITE_WEBP_QUALITY = 64
const IMWRITE_PAM_TUPLETYPE = 128
const IMWRITE_PNG_STRATEGY_DEFAULT = 0
const IMWRITE_PNG_STRATEGY_FILTERED = 1
const IMWRITE_PNG_STRATEGY_HUFFMAN_ONLY = 2
const IMWRITE_PNG_STRATEGY_RLE = 3
const IMWRITE_PNG_STRATEGY_FIXED = 4
const IMWRITE_PAM_FORMAT_NULL = 0
const IMWRITE_PAM_FORMAT_BLACKANDWHITE = 1
const IMWRITE_PAM_FORMAT_GRAYSCALE = 2
const IMWRITE_PAM_FORMAT_GRAYSCALE_ALPHA = 3
const IMWRITE_PAM_FORMAT_RGB = 4
const IMWRITE_PAM_FORMAT_RGB_ALPHA = 5

var CV_LOAD_IMAGE_UNCHANGED = -1
var IMREAD_UNCHANGED = -1

func Imdecode(buf *Mat, flags int) *Mat {
	retVal := NewMat(ImgcodecsNative_imdecode_0(buf.GetNativeObjAddr(), flags))
	return retVal
}

func Imencode(ext string, img *Mat, buf *MatOfByte, params *MatOfInt) bool {
	buf_mat := buf
	params_mat := params
	retVal := ImgcodecsNative_imencode_0(ext, img.GetNativeObjAddr(), buf_mat.GetNativeObjAddr(), params_mat.GetNativeObjAddr())
	return retVal
}
func Imencode2(ext string, img *Mat, buf *MatOfByte) bool {
	buf_mat := buf
	retVal := ImgcodecsNative_imencode_1(ext, img.GetNativeObjAddr(), buf_mat.GetNativeObjAddr())
	return retVal
}

func Imread(filename string, flags int) *Mat {
	retVal := NewMat(ImgcodecsNative_imread_0(filename, flags))
	return retVal
}
func Imread2(filename string) *Mat {
	retVal := NewMat(ImgcodecsNative_imread_1(filename))
	return retVal
}

func Imreadmulti(filename string, flags int) ([]*Mat, bool) {
	mats_mat := NewMat2()
	retVal := ImgcodecsNative_imreadmulti_0(filename, mats_mat.GetNativeObjAddr(), flags)
	mats := ConvertersToVectorMat(mats_mat)
	mats_mat.Release()
	return mats, retVal
}
func Imreadmulti2(filename string) ([]*Mat, bool) {
	mats_mat := NewMat2()
	retVal := ImgcodecsNative_imreadmulti_1(filename, mats_mat.GetNativeObjAddr())
	mats := ConvertersToVectorMat(mats_mat)
	mats_mat.Release()
	return mats, retVal
}

func Imwrite(filename string, img *Mat, params *MatOfInt) bool {
	params_mat := params
	retVal := ImgcodecsNative_imwrite_0(filename, img.GetNativeObjAddr(), params_mat.GetNativeObjAddr())
	return retVal
}
func Imwrite2(filename string, img *Mat) bool {
	retVal := ImgcodecsNative_imwrite_1(filename, img.GetNativeObjAddr())
	return retVal
}
