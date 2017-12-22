package photo

import (
	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

const cV_INPAINT_NSPhoto = 0
const cV_INPAINT_TELEAPhoto = 1
const INPAINT_NS = 0
const INPAINT_TELEA = 1
const NORMAL_CLONE = 1
const MIXED_CLONE = 2
const MONOCHROME_TRANSFER = 3
const RECURS_FILTER = 1
const NORMCONV_FILTER = 2
const LDR_SIZE = 256

func ColorChange(src *Mat, mask *Mat, dst *Mat, red_mul float32, green_mul float32, blue_mul float32) {
	PhotoNative_colorChange_0(src.GetNativeObjAddr(), mask.GetNativeObjAddr(), dst.GetNativeObjAddr(), red_mul, green_mul, blue_mul)
	return
}
func ColorChange2(src *Mat, mask *Mat, dst *Mat) {
	PhotoNative_colorChange_1(src.GetNativeObjAddr(), mask.GetNativeObjAddr(), dst.GetNativeObjAddr())
	return
}
func CreateAlignMTB(max_bits int, exclude_range int, cut bool) *AlignMTB {
	retVal := NewAlignMTB(PhotoNative_createAlignMTB_0(max_bits, exclude_range, cut))
	return retVal
}
func CreateAlignMTB2() *AlignMTB {
	retVal := NewAlignMTB(PhotoNative_createAlignMTB_1())
	return retVal
}
func CreateCalibrateDebevec(samples int, lambda float32, random bool) *CalibrateDebevec {
	retVal := NewCalibrateDebevec(PhotoNative_createCalibrateDebevec_0(samples, lambda, random))
	return retVal
}
func CreateCalibrateDebevec2() *CalibrateDebevec {
	retVal := NewCalibrateDebevec(PhotoNative_createCalibrateDebevec_1())
	return retVal
}
func CreateCalibrateRobertson(max_iter int, threshold float32) *CalibrateRobertson {
	retVal := NewCalibrateRobertson(PhotoNative_createCalibrateRobertson_0(max_iter, threshold))
	return retVal
}
func CreateCalibrateRobertson2() *CalibrateRobertson {
	retVal := NewCalibrateRobertson(PhotoNative_createCalibrateRobertson_1())
	return retVal
}
func CreateMergeDebevec() *MergeDebevec {
	retVal := NewMergeDebevec(PhotoNative_createMergeDebevec_0())
	return retVal
}
func CreateMergeMertens(contrast_weight float32, saturation_weight float32, exposure_weight float32) *MergeMertens {
	retVal := NewMergeMertens(PhotoNative_createMergeMertens_0(contrast_weight, saturation_weight, exposure_weight))
	return retVal
}
func CreateMergeMertens2() *MergeMertens {
	retVal := NewMergeMertens(PhotoNative_createMergeMertens_1())
	return retVal
}
func CreateMergeRobertson() *MergeRobertson {
	retVal := NewMergeRobertson(PhotoNative_createMergeRobertson_0())
	return retVal
}
func CreateTonemap(gamma float32) *Tonemap {
	retVal := NewTonemap(PhotoNative_createTonemap_0(gamma))
	return retVal
}
func CreateTonemap2() *Tonemap {
	retVal := NewTonemap(PhotoNative_createTonemap_1())
	return retVal
}
func CreateTonemapDrago(gamma float32, saturation float32, bias float32) *TonemapDrago {
	retVal := NewTonemapDrago(PhotoNative_createTonemapDrago_0(gamma, saturation, bias))
	return retVal
}
func CreateTonemapDrago2() *TonemapDrago {
	retVal := NewTonemapDrago(PhotoNative_createTonemapDrago_1())
	return retVal
}
func CreateTonemapDurand(gamma float32, contrast float32, saturation float32, sigma_space float32, sigma_color float32) *TonemapDurand {
	retVal := NewTonemapDurand(PhotoNative_createTonemapDurand_0(gamma, contrast, saturation, sigma_space, sigma_color))
	return retVal
}
func CreateTonemapDurand2() *TonemapDurand {
	retVal := NewTonemapDurand(PhotoNative_createTonemapDurand_1())
	return retVal
}
func CreateTonemapMantiuk(gamma float32, scale float32, saturation float32) *TonemapMantiuk {
	retVal := NewTonemapMantiuk(PhotoNative_createTonemapMantiuk_0(gamma, scale, saturation))
	return retVal
}
func CreateTonemapMantiuk2() *TonemapMantiuk {
	retVal := NewTonemapMantiuk(PhotoNative_createTonemapMantiuk_1())
	return retVal
}
func CreateTonemapReinhard(gamma float32, intensity float32, light_adapt float32, color_adapt float32) *TonemapReinhard {
	retVal := NewTonemapReinhard(PhotoNative_createTonemapReinhard_0(gamma, intensity, light_adapt, color_adapt))
	return retVal
}
func CreateTonemapReinhard2() *TonemapReinhard {
	retVal := NewTonemapReinhard(PhotoNative_createTonemapReinhard_1())
	return retVal
}
func Decolor(src *Mat, grayscale *Mat, color_boost *Mat) {
	PhotoNative_decolor_0(src.GetNativeObjAddr(), grayscale.GetNativeObjAddr(), color_boost.GetNativeObjAddr())
	return
}
func Denoise_TVL1(observations []*Mat, result *Mat, lambda float64, niters int) {
	observations_mat := ConvertersVectorMat(observations)
	PhotoNative_denoise_TVL1_0(observations_mat.GetNativeObjAddr(), result.GetNativeObjAddr(), lambda, niters)
	return
}
func Denoise_TVL12(observations []*Mat, result *Mat) {
	observations_mat := ConvertersVectorMat(observations)
	PhotoNative_denoise_TVL1_1(observations_mat.GetNativeObjAddr(), result.GetNativeObjAddr())
	return
}
func DetailEnhance(src *Mat, dst *Mat, sigma_s float32, sigma_r float32) {
	PhotoNative_detailEnhance_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), sigma_s, sigma_r)
	return
}
func DetailEnhance2(src *Mat, dst *Mat) {
	PhotoNative_detailEnhance_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr())
	return
}
func EdgePreservingFilter(src *Mat, dst *Mat, flags int, sigma_s float32, sigma_r float32) {
	PhotoNative_edgePreservingFilter_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), flags, sigma_s, sigma_r)
	return
}
func EdgePreservingFilter2(src *Mat, dst *Mat) {
	PhotoNative_edgePreservingFilter_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr())
	return
}
func FastNlMeansDenoising(src *Mat, dst *Mat, h float32, templateWindowSize int, searchWindowSize int) {
	PhotoNative_fastNlMeansDenoising_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), h, templateWindowSize, searchWindowSize)
	return
}
func FastNlMeansDenoising2(src *Mat, dst *Mat) {
	PhotoNative_fastNlMeansDenoising_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr())
	return
}
func FastNlMeansDenoising3(src *Mat, dst *Mat, h *MatOfFloat, templateWindowSize int, searchWindowSize int, normType int) {
	h_mat := h
	PhotoNative_fastNlMeansDenoising_2(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), h_mat.GetNativeObjAddr(), templateWindowSize, searchWindowSize, normType)
	return
}
func FastNlMeansDenoising4(src *Mat, dst *Mat, h *MatOfFloat) {
	h_mat := h
	PhotoNative_fastNlMeansDenoising_3(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), h_mat.GetNativeObjAddr())
	return
}
func FastNlMeansDenoisingColored(src *Mat, dst *Mat, h float32, hColor float32, templateWindowSize int, searchWindowSize int) {
	PhotoNative_fastNlMeansDenoisingColored_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), h, hColor, templateWindowSize, searchWindowSize)
	return
}
func FastNlMeansDenoisingColored2(src *Mat, dst *Mat) {
	PhotoNative_fastNlMeansDenoisingColored_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr())
	return
}
func FastNlMeansDenoisingColoredMulti(srcImgs []*Mat, dst *Mat, imgToDenoiseIndex int, temporalWindowSize int, h float32, hColor float32, templateWindowSize int, searchWindowSize int) {
	srcImgs_mat := ConvertersVectorMat(srcImgs)
	PhotoNative_fastNlMeansDenoisingColoredMulti_0(srcImgs_mat.GetNativeObjAddr(), dst.GetNativeObjAddr(), imgToDenoiseIndex, temporalWindowSize, h, hColor, templateWindowSize, searchWindowSize)
	return
}
func FastNlMeansDenoisingColoredMulti2(srcImgs []*Mat, dst *Mat, imgToDenoiseIndex int, temporalWindowSize int) {
	srcImgs_mat := ConvertersVectorMat(srcImgs)
	PhotoNative_fastNlMeansDenoisingColoredMulti_1(srcImgs_mat.GetNativeObjAddr(), dst.GetNativeObjAddr(), imgToDenoiseIndex, temporalWindowSize)
	return
}
func FastNlMeansDenoisingMulti(srcImgs []*Mat, dst *Mat, imgToDenoiseIndex int, temporalWindowSize int, h float32, templateWindowSize int, searchWindowSize int) {
	srcImgs_mat := ConvertersVectorMat(srcImgs)
	PhotoNative_fastNlMeansDenoisingMulti_0(srcImgs_mat.GetNativeObjAddr(), dst.GetNativeObjAddr(), imgToDenoiseIndex, temporalWindowSize, h, templateWindowSize, searchWindowSize)
	return
}
func FastNlMeansDenoisingMulti2(srcImgs []*Mat, dst *Mat, imgToDenoiseIndex int, temporalWindowSize int) {
	srcImgs_mat := ConvertersVectorMat(srcImgs)
	PhotoNative_fastNlMeansDenoisingMulti_1(srcImgs_mat.GetNativeObjAddr(), dst.GetNativeObjAddr(), imgToDenoiseIndex, temporalWindowSize)
	return
}
func FastNlMeansDenoisingMulti3(srcImgs []*Mat, dst *Mat, imgToDenoiseIndex int, temporalWindowSize int, h *MatOfFloat, templateWindowSize int, searchWindowSize int, normType int) {
	srcImgs_mat := ConvertersVectorMat(srcImgs)
	h_mat := h
	PhotoNative_fastNlMeansDenoisingMulti_2(srcImgs_mat.GetNativeObjAddr(), dst.GetNativeObjAddr(), imgToDenoiseIndex, temporalWindowSize, h_mat.GetNativeObjAddr(), templateWindowSize, searchWindowSize, normType)
	return
}
func FastNlMeansDenoisingMulti4(srcImgs []*Mat, dst *Mat, imgToDenoiseIndex int, temporalWindowSize int, h *MatOfFloat) {
	srcImgs_mat := ConvertersVectorMat(srcImgs)
	h_mat := h
	PhotoNative_fastNlMeansDenoisingMulti_3(srcImgs_mat.GetNativeObjAddr(), dst.GetNativeObjAddr(), imgToDenoiseIndex, temporalWindowSize, h_mat.GetNativeObjAddr())
	return
}
func IlluminationChange(src *Mat, mask *Mat, dst *Mat, alpha float32, beta float32) {
	PhotoNative_illuminationChange_0(src.GetNativeObjAddr(), mask.GetNativeObjAddr(), dst.GetNativeObjAddr(), alpha, beta)
	return
}
func IlluminationChange2(src *Mat, mask *Mat, dst *Mat) {
	PhotoNative_illuminationChange_1(src.GetNativeObjAddr(), mask.GetNativeObjAddr(), dst.GetNativeObjAddr())
	return
}
func Inpaint(src *Mat, inpaintMask *Mat, dst *Mat, inpaintRadius float64, flags int) {
	PhotoNative_inpaint_0(src.GetNativeObjAddr(), inpaintMask.GetNativeObjAddr(), dst.GetNativeObjAddr(), inpaintRadius, flags)
	return
}
func PencilSketch(src *Mat, dst1 *Mat, dst2 *Mat, sigma_s float32, sigma_r float32, shade_factor float32) {
	PhotoNative_pencilSketch_0(src.GetNativeObjAddr(), dst1.GetNativeObjAddr(), dst2.GetNativeObjAddr(), sigma_s, sigma_r, shade_factor)
	return
}
func PencilSketch2(src *Mat, dst1 *Mat, dst2 *Mat) {
	PhotoNative_pencilSketch_1(src.GetNativeObjAddr(), dst1.GetNativeObjAddr(), dst2.GetNativeObjAddr())
	return
}
func SeamlessClone(src *Mat, dst *Mat, mask *Mat, p *Point, blend *Mat, flags int) {
	PhotoNative_seamlessClone_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), mask.GetNativeObjAddr(), p.X, p.Y, blend.GetNativeObjAddr(), flags)
	return
}
func Stylization(src *Mat, dst *Mat, sigma_s float32, sigma_r float32) {
	PhotoNative_stylization_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), sigma_s, sigma_r)
	return
}
func Stylization2(src *Mat, dst *Mat) {
	PhotoNative_stylization_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr())
	return
}
func TextureFlattening(src *Mat, mask *Mat, dst *Mat, low_threshold float32, high_threshold float32, kernel_size int) {
	PhotoNative_textureFlattening_0(src.GetNativeObjAddr(), mask.GetNativeObjAddr(), dst.GetNativeObjAddr(), low_threshold, high_threshold, kernel_size)
	return
}
func TextureFlattening2(src *Mat, mask *Mat, dst *Mat) {
	PhotoNative_textureFlattening_1(src.GetNativeObjAddr(), mask.GetNativeObjAddr(), dst.GetNativeObjAddr())
	return
}
