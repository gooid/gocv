package core

import (
	. "github.com/gooid/gocv/opencv3/internal/native"
)

const cV_8UCore = 0
const cV_8SCore = 1
const cV_16UCore = 2
const cV_16SCore = 3
const cV_32SCore = 4
const cV_32FCore = 5
const cV_64FCore = 6
const cV_USRTYPE1Core = 7
const CoreSVD_MODIFY_A = 1
const CoreSVD_NO_UV = 2
const CoreSVD_FULL_UV = 4
const CoreREDUCE_SUM = 0
const CoreREDUCE_AVG = 1
const CoreREDUCE_MAX = 2
const CoreREDUCE_MIN = 3
const CoreStsOk = 0
const CoreDECOMP_LU = 0
const CoreDECOMP_SVD = 1
const CoreDECOMP_EIG = 2
const CoreDECOMP_CHOLESKY = 3
const CoreDECOMP_QR = 4
const CoreDECOMP_NORMAL = 16
const CoreNORM_INF = 1
const CoreNORM_L1 = 2
const CoreNORM_L2 = 4
const CoreNORM_L2SQR = 5
const CoreNORM_HAMMING = 6
const CoreNORM_HAMMING2 = 7
const CoreNORM_TYPE_MASK = 7
const CoreNORM_RELATIVE = 8
const CoreNORM_MINMAX = 32
const CoreCMP_EQ = 0
const CoreCMP_GT = 1
const CoreCMP_GE = 2
const CoreCMP_LT = 3
const CoreCMP_LE = 4
const CoreCMP_NE = 5
const CoreGEMM_1_T = 1
const CoreGEMM_2_T = 2
const CoreGEMM_3_T = 4
const CoreDFT_INVERSE = 1
const CoreDFT_SCALE = 2
const CoreDFT_ROWS = 4
const CoreDFT_COMPLEX_OUTPUT = 16
const CoreDFT_REAL_OUTPUT = 32
const CoreDFT_COMPLEX_INPUT = 64
const CoreBORDER_CONSTANT = 0
const CoreBORDER_REPLICATE = 1
const CoreBORDER_REFLECT = 2
const CoreBORDER_WRAP = 3
const CoreBORDER_REFLECT_101 = 4
const CoreBORDER_TRANSPARENT = 5
const CoreBORDER_ISOLATED = 16
const CoreSORT_EVERY_ROW = 0
const CoreSORT_EVERY_COLUMN = 1
const CoreSORT_ASCENDING = 0
const CoreSORT_DESCENDING = 16
const CoreCOVAR_SCRAMBLED = 0
const CoreCOVAR_NORMAL = 1
const CoreCOVAR_USE_AVG = 2
const CoreCOVAR_SCALE = 4
const CoreCOVAR_ROWS = 8
const CoreCOVAR_COLS = 16
const CoreKMEANS_RANDOM_CENTERS = 0
const CoreKMEANS_PP_CENTERS = 2
const CoreKMEANS_USE_INITIAL_LABELS = 1
const CoreLINE_4 = 4
const CoreLINE_8 = 8
const CoreLINE_AA = 16
const CoreFONT_HERSHEY_SIMPLEX = 0
const CoreFONT_HERSHEY_PLAIN = 1
const CoreFONT_HERSHEY_DUPLEX = 2
const CoreFONT_HERSHEY_COMPLEX = 3
const CoreFONT_HERSHEY_TRIPLEX = 4
const CoreFONT_HERSHEY_COMPLEX_SMALL = 5
const CoreFONT_HERSHEY_SCRIPT_SIMPLEX = 6
const CoreFONT_HERSHEY_SCRIPT_COMPLEX = 7
const CoreFONT_ITALIC = 16
const CoreROTATE_90_CLOCKWISE = 0
const CoreROTATE_180 = 1
const CoreROTATE_90_COUNTERCLOCKWISE = 2
const CoreTYPE_GENERAL = 0
const CoreIMPL_PLAIN = 0
const CoreFLAGS_NONE = 0
const CoreFLAGS_MAPPING = 0x01
const CoreFLAGS_EXPAND_SAME_NAMES = 0x02

var CoreVERSION = getVersion()
var CoreNATIVE_LIBRARY_NAME = getNativeLibraryName()
var CoreVERSION_MAJOR = getVersionMajor()
var CoreVERSION_MINOR = getVersionMinor()
var CoreVERSION_REVISION = getVersionRevision()
var CoreVERSION_STATUS = getVersionStatus()
var CoreFILLED = -1
var CoreStsBackTrace = -1
var CoreStsError = -2
var CoreStsInternal = -3
var CoreStsNoMem = -4
var CoreStsBadArg = -5
var CoreStsBadFunc = -6
var CoreStsNoConv = -7
var CoreStsAutoTrace = -8
var CoreHeaderIsNull = -9
var CoreBadImageSize = -10
var CoreBadOffset = -11
var CoreBadDataPtr = -12
var CoreBadStep = -13
var CoreBadModelOrChSeq = -14
var CoreBadNumChannels = -15
var CoreBadNumChannel1U = -16
var CoreBadDepth = -17
var CoreBadAlphaChannel = -18
var CoreBadOrder = -19
var CoreBadOrigin = -20
var CoreBadAlign = -21
var CoreBadCallBack = -22
var CoreBadTileSize = -23
var CoreBadCOI = -24
var CoreBadROISize = -25
var CoreMaskIsTiled = -26
var CoreStsNullPtr = -27
var CoreStsVecLengthErr = -28
var CoreStsFilterStructContentErr = -29
var CoreStsKernelStructContentErr = -30
var CoreStsFilterOffsetErr = -31
var CoreStsBadSize = -201
var CoreStsDivByZero = -202
var CoreStsInplaceNotSupported = -203
var CoreStsObjectNotFound = -204
var CoreStsUnmatchedFormats = -205
var CoreStsBadFlag = -206
var CoreStsBadPoint = -207
var CoreStsBadMask = -208
var CoreStsUnmatchedSizes = -209
var CoreStsUnsupportedFormat = -210
var CoreStsOutOfRange = -211
var CoreStsParseError = -212
var CoreStsNotImplemented = -213
var CoreStsBadMemBlock = -214
var CoreStsAssert = -215
var CoreGpuNotSupported = -216
var CoreGpuApiCallError = -217
var CoreOpenGlNotSupported = -218
var CoreOpenGlApiCallError = -219
var CoreOpenCLApiCallError = -220
var CoreOpenCLDoubleNotSupported = -221
var CoreOpenCLInitError = -222
var CoreOpenCLNoAMDBlasFft = -223
var CoreDCT_INVERSE = CoreDFT_INVERSE
var CoreDCT_ROWS = CoreDFT_ROWS
var CoreBORDER_REFLECT101 = CoreBORDER_REFLECT_101
var CoreBORDER_DEFAULT = CoreBORDER_REFLECT_101
var CoreTYPE_MARKER = 0 + 1
var CoreTYPE_WRAPPER = 0 + 2
var CoreTYPE_FUN = 0 + 3
var CoreIMPL_IPP = 0 + 1
var CoreIMPL_OPENCL = 0 + 2

type Core struct{}

func NewCore() (rcvr *Core) {
	rcvr = &Core{}
	return
}
func LUT(src *Mat, lut *Mat, dst *Mat) {
	CoreNative_LUT_0(src.nativeObj, lut.nativeObj, dst.nativeObj)
	return
}

func Mahalanobis(v1 *Mat, v2 *Mat, icovar *Mat) float64 {
	retVal := CoreNative_Mahalanobis_0(v1.nativeObj, v2.nativeObj, icovar.nativeObj)
	return retVal
}

func PCABackProject(data *Mat, mean *Mat, eigenvectors *Mat, result *Mat) {
	CoreNative_PCABackProject_0(data.nativeObj, mean.nativeObj, eigenvectors.nativeObj, result.nativeObj)
	return
}

func PCACompute(data *Mat, mean *Mat, eigenvectors *Mat, retainedVariance float64) {
	CoreNative_PCACompute_0(data.nativeObj, mean.nativeObj, eigenvectors.nativeObj, retainedVariance)
	return
}
func PCACompute2(data *Mat, mean *Mat, eigenvectors *Mat, maxComponents int) {
	CoreNative_PCACompute_1(data.nativeObj, mean.nativeObj, eigenvectors.nativeObj, maxComponents)
	return
}
func PCACompute3(data *Mat, mean *Mat, eigenvectors *Mat) {
	CoreNative_PCACompute_2(data.nativeObj, mean.nativeObj, eigenvectors.nativeObj)
	return
}

func PCAProject(data *Mat, mean *Mat, eigenvectors *Mat, result *Mat) {
	CoreNative_PCAProject_0(data.nativeObj, mean.nativeObj, eigenvectors.nativeObj, result.nativeObj)
	return
}

func PSNR(src1 *Mat, src2 *Mat) float64 {
	retVal := CoreNative_PSNR_0(src1.nativeObj, src2.nativeObj)
	return retVal
}

func SVBackSubst(w *Mat, u *Mat, vt *Mat, rhs *Mat, dst *Mat) {
	CoreNative_SVBackSubst_0(w.nativeObj, u.nativeObj, vt.nativeObj, rhs.nativeObj, dst.nativeObj)
	return
}

func SVDecomp(src *Mat, w *Mat, u *Mat, vt *Mat, flags int) {
	CoreNative_SVDecomp_0(src.nativeObj, w.nativeObj, u.nativeObj, vt.nativeObj, flags)
	return
}
func SVDecomp2(src *Mat, w *Mat, u *Mat, vt *Mat) {
	CoreNative_SVDecomp_1(src.nativeObj, w.nativeObj, u.nativeObj, vt.nativeObj)
	return
}

func Absdiff(src1 *Mat, src2 *Mat, dst *Mat) {
	CoreNative_absdiff_0(src1.nativeObj, src2.nativeObj, dst.nativeObj)
	return
}
func Absdiff2(src1 *Mat, src2 *Scalar, dst *Mat) {
	CoreNative_absdiff_1(src1.nativeObj, src2.Val[0], src2.Val[1], src2.Val[2], src2.Val[3], dst.nativeObj)
	return
}

func Add(src1 *Mat, src2 *Mat, dst *Mat, mask *Mat, dtype int) {
	CoreNative_add_0(src1.nativeObj, src2.nativeObj, dst.nativeObj, mask.nativeObj, dtype)
	return
}
func Add2(src1 *Mat, src2 *Mat, dst *Mat, mask *Mat) {
	CoreNative_add_1(src1.nativeObj, src2.nativeObj, dst.nativeObj, mask.nativeObj)
	return
}
func Add3(src1 *Mat, src2 *Mat, dst *Mat) {
	CoreNative_add_2(src1.nativeObj, src2.nativeObj, dst.nativeObj)
	return
}
func Add4(src1 *Mat, src2 *Scalar, dst *Mat, mask *Mat, dtype int) {
	CoreNative_add_3(src1.nativeObj, src2.Val[0], src2.Val[1], src2.Val[2], src2.Val[3], dst.nativeObj, mask.nativeObj, dtype)
	return
}
func Add5(src1 *Mat, src2 *Scalar, dst *Mat, mask *Mat) {
	CoreNative_add_4(src1.nativeObj, src2.Val[0], src2.Val[1], src2.Val[2], src2.Val[3], dst.nativeObj, mask.nativeObj)
	return
}
func Add6(src1 *Mat, src2 *Scalar, dst *Mat) {
	CoreNative_add_5(src1.nativeObj, src2.Val[0], src2.Val[1], src2.Val[2], src2.Val[3], dst.nativeObj)
	return
}
func AddWeighted(src1 *Mat, alpha float64, src2 *Mat, beta float64, gamma float64, dst *Mat, dtype int) {
	CoreNative_addWeighted_0(src1.nativeObj, alpha, src2.nativeObj, beta, gamma, dst.nativeObj, dtype)
	return
}
func AddWeighted2(src1 *Mat, alpha float64, src2 *Mat, beta float64, gamma float64, dst *Mat) {
	CoreNative_addWeighted_1(src1.nativeObj, alpha, src2.nativeObj, beta, gamma, dst.nativeObj)
	return
}

func BatchDistance(src1 *Mat, src2 *Mat, dist *Mat, dtype int, nidx *Mat, normType int, K int, mask *Mat, update int, crosscheck bool) {
	CoreNative_batchDistance_0(src1.nativeObj, src2.nativeObj, dist.nativeObj, dtype, nidx.nativeObj, normType, K, mask.nativeObj, update, crosscheck)
	return
}
func BatchDistance2(src1 *Mat, src2 *Mat, dist *Mat, dtype int, nidx *Mat, normType int, K int) {
	CoreNative_batchDistance_1(src1.nativeObj, src2.nativeObj, dist.nativeObj, dtype, nidx.nativeObj, normType, K)
	return
}
func BatchDistance3(src1 *Mat, src2 *Mat, dist *Mat, dtype int, nidx *Mat) {
	CoreNative_batchDistance_2(src1.nativeObj, src2.nativeObj, dist.nativeObj, dtype, nidx.nativeObj)
	return
}

func Bitwise_and(src1 *Mat, src2 *Mat, dst *Mat, mask *Mat) {
	CoreNative_bitwise_and_0(src1.nativeObj, src2.nativeObj, dst.nativeObj, mask.nativeObj)
	return
}
func Bitwise_and2(src1 *Mat, src2 *Mat, dst *Mat) {
	CoreNative_bitwise_and_1(src1.nativeObj, src2.nativeObj, dst.nativeObj)
	return
}

func Bitwise_not(src *Mat, dst *Mat, mask *Mat) {
	CoreNative_bitwise_not_0(src.nativeObj, dst.nativeObj, mask.nativeObj)
	return
}
func Bitwise_not2(src *Mat, dst *Mat) {
	CoreNative_bitwise_not_1(src.nativeObj, dst.nativeObj)
	return
}

func Bitwise_or(src1 *Mat, src2 *Mat, dst *Mat, mask *Mat) {
	CoreNative_bitwise_or_0(src1.nativeObj, src2.nativeObj, dst.nativeObj, mask.nativeObj)
	return
}
func Bitwise_or2(src1 *Mat, src2 *Mat, dst *Mat) {
	CoreNative_bitwise_or_1(src1.nativeObj, src2.nativeObj, dst.nativeObj)
	return
}

func Bitwise_xor(src1 *Mat, src2 *Mat, dst *Mat, mask *Mat) {
	CoreNative_bitwise_xor_0(src1.nativeObj, src2.nativeObj, dst.nativeObj, mask.nativeObj)
	return
}
func Bitwise_xor2(src1 *Mat, src2 *Mat, dst *Mat) {
	CoreNative_bitwise_xor_1(src1.nativeObj, src2.nativeObj, dst.nativeObj)
	return
}

func BorderInterpolate(p int, length int, borderType int) int {
	retVal := CoreNative_borderInterpolate_0(p, length, borderType)
	return retVal
}

func CalcCovarMatrix(samples *Mat, covar *Mat, mean *Mat, flags int, ctype int) {
	CoreNative_calcCovarMatrix_0(samples.nativeObj, covar.nativeObj, mean.nativeObj, flags, ctype)
	return
}
func CalcCovarMatrix2(samples *Mat, covar *Mat, mean *Mat, flags int) {
	CoreNative_calcCovarMatrix_1(samples.nativeObj, covar.nativeObj, mean.nativeObj, flags)
	return
}

func CartToPolar(x *Mat, y *Mat, magnitude *Mat, angle *Mat, angleInDegrees bool) {
	CoreNative_cartToPolar_0(x.nativeObj, y.nativeObj, magnitude.nativeObj, angle.nativeObj, angleInDegrees)
	return
}
func CartToPolar2(x *Mat, y *Mat, magnitude *Mat, angle *Mat) {
	CoreNative_cartToPolar_1(x.nativeObj, y.nativeObj, magnitude.nativeObj, angle.nativeObj)
	return
}

func CheckRange(a *Mat, quiet bool, minVal float64, maxVal float64) bool {
	retVal := CoreNative_checkRange_0(a.nativeObj, quiet, minVal, maxVal)
	return retVal
}
func CheckRange2(a *Mat) bool {
	retVal := CoreNative_checkRange_1(a.nativeObj)
	return retVal
}

func Compare(src1 *Mat, src2 *Mat, dst *Mat, cmpop int) {
	CoreNative_compare_0(src1.nativeObj, src2.nativeObj, dst.nativeObj, cmpop)
	return
}
func Compare2(src1 *Mat, src2 *Scalar, dst *Mat, cmpop int) {
	CoreNative_compare_1(src1.nativeObj, src2.Val[0], src2.Val[1], src2.Val[2], src2.Val[3], dst.nativeObj, cmpop)
	return
}

func CompleteSymm(mtx *Mat, lowerToUpper bool) {
	CoreNative_completeSymm_0(mtx.nativeObj, lowerToUpper)
	return
}
func CompleteSymm2(mtx *Mat) {
	CoreNative_completeSymm_1(mtx.nativeObj)
	return
}

func ConvertFp16(src *Mat, dst *Mat) {
	CoreNative_convertFp16_0(src.nativeObj, dst.nativeObj)
	return
}

func ConvertScaleAbs(src *Mat, dst *Mat, alpha float64, beta float64) {
	CoreNative_convertScaleAbs_0(src.nativeObj, dst.nativeObj, alpha, beta)
	return
}
func ConvertScaleAbs2(src *Mat, dst *Mat) {
	CoreNative_convertScaleAbs_1(src.nativeObj, dst.nativeObj)
	return
}

func CopyMakeBorder(src *Mat, dst *Mat, top int, bottom int, left int, right int, borderType int, value *Scalar) {
	CoreNative_copyMakeBorder_0(src.nativeObj, dst.nativeObj, top, bottom, left, right, borderType, value.Val[0], value.Val[1], value.Val[2], value.Val[3])
	return
}
func CopyMakeBorder2(src *Mat, dst *Mat, top int, bottom int, left int, right int, borderType int) {
	CoreNative_copyMakeBorder_1(src.nativeObj, dst.nativeObj, top, bottom, left, right, borderType)
	return
}

func CountNonZero(src *Mat) int {
	retVal := CoreNative_countNonZero_0(src.nativeObj)
	return retVal
}

func CubeRoot(val float32) float32 {
	retVal := CoreNative_cubeRoot_0(val)
	return retVal
}

func Dct(src *Mat, dst *Mat, flags int) {
	CoreNative_dct_0(src.nativeObj, dst.nativeObj, flags)
	return
}
func Dct2(src *Mat, dst *Mat) {
	CoreNative_dct_1(src.nativeObj, dst.nativeObj)
	return
}

func Determinant(mtx *Mat) float64 {
	retVal := CoreNative_determinant_0(mtx.nativeObj)
	return retVal
}

func Dft(src *Mat, dst *Mat, flags int, nonzeroRows int) {
	CoreNative_dft_0(src.nativeObj, dst.nativeObj, flags, nonzeroRows)
	return
}
func Dft2(src *Mat, dst *Mat) {
	CoreNative_dft_1(src.nativeObj, dst.nativeObj)
	return
}

func Divide(src1 *Mat, src2 *Mat, dst *Mat, scale float64, dtype int) {
	CoreNative_divide_0(src1.nativeObj, src2.nativeObj, dst.nativeObj, scale, dtype)
	return
}
func Divide2(src1 *Mat, src2 *Mat, dst *Mat, scale float64) {
	CoreNative_divide_1(src1.nativeObj, src2.nativeObj, dst.nativeObj, scale)
	return
}
func Divide3(src1 *Mat, src2 *Mat, dst *Mat) {
	CoreNative_divide_2(src1.nativeObj, src2.nativeObj, dst.nativeObj)
	return
}
func Divide4(src1 *Mat, src2 *Scalar, dst *Mat, scale float64, dtype int) {
	CoreNative_divide_3(src1.nativeObj, src2.Val[0], src2.Val[1], src2.Val[2], src2.Val[3], dst.nativeObj, scale, dtype)
	return
}
func Divide5(src1 *Mat, src2 *Scalar, dst *Mat, scale float64) {
	CoreNative_divide_4(src1.nativeObj, src2.Val[0], src2.Val[1], src2.Val[2], src2.Val[3], dst.nativeObj, scale)
	return
}
func Divide6(src1 *Mat, src2 *Scalar, dst *Mat) {
	CoreNative_divide_5(src1.nativeObj, src2.Val[0], src2.Val[1], src2.Val[2], src2.Val[3], dst.nativeObj)
	return
}
func Divide7(scale float64, src2 *Mat, dst *Mat, dtype int) {
	CoreNative_divide_6(scale, src2.nativeObj, dst.nativeObj, dtype)
	return
}
func Divide8(scale float64, src2 *Mat, dst *Mat) {
	CoreNative_divide_7(scale, src2.nativeObj, dst.nativeObj)
	return
}

func Eigen(src *Mat, eigenvalues *Mat, eigenvectors *Mat) bool {
	retVal := CoreNative_eigen_0(src.nativeObj, eigenvalues.nativeObj, eigenvectors.nativeObj)
	return retVal
}
func Eigen2(src *Mat, eigenvalues *Mat) bool {
	retVal := CoreNative_eigen_1(src.nativeObj, eigenvalues.nativeObj)
	return retVal
}
func EigenNonSymmetric(src *Mat, eigenvalues *Mat, eigenvectors *Mat) {
	CoreNative_eigenNonSymmetric_0(src.nativeObj, eigenvalues.nativeObj, eigenvectors.nativeObj)
	return
}

func Exp(src *Mat, dst *Mat) {
	CoreNative_exp_0(src.nativeObj, dst.nativeObj)
	return
}

func ExtractChannel(src *Mat, dst *Mat, coi int) {
	CoreNative_extractChannel_0(src.nativeObj, dst.nativeObj, coi)
	return
}

func FastAtan2(y float32, x float32) float32 {
	retVal := CoreNative_fastAtan2_0(y, x)
	return retVal
}

func FindNonZero(src *Mat, idx *Mat) {
	CoreNative_findNonZero_0(src.nativeObj, idx.nativeObj)
	return
}

func Flip(src *Mat, dst *Mat, flipCode int) {
	CoreNative_flip_0(src.nativeObj, dst.nativeObj, flipCode)
	return
}

func Gemm(src1 *Mat, src2 *Mat, alpha float64, src3 *Mat, beta float64, dst *Mat, flags int) {
	CoreNative_gemm_0(src1.nativeObj, src2.nativeObj, alpha, src3.nativeObj, beta, dst.nativeObj, flags)
	return
}
func Gemm2(src1 *Mat, src2 *Mat, alpha float64, src3 *Mat, beta float64, dst *Mat) {
	CoreNative_gemm_1(src1.nativeObj, src2.nativeObj, alpha, src3.nativeObj, beta, dst.nativeObj)
	return
}

func GetBuildInformation() string {
	retVal := CoreNative_getBuildInformation_0()
	return retVal
}

func GetCPUTickCount() int64 {
	retVal := CoreNative_getCPUTickCount_0()
	return retVal
}

func GetIppVersion() string {
	retVal := CoreNative_getIppVersion_0()
	return retVal
}

func getNativeLibraryName() string {
	return "opencv_java331"
}
func GetNumThreads() int {
	retVal := CoreNative_getNumThreads_0()
	return retVal
}

func GetNumberOfCPUs() int {
	retVal := CoreNative_getNumberOfCPUs_0()
	return retVal
}

func GetOptimalDFTSize(vecsize int) int {
	retVal := CoreNative_getOptimalDFTSize_0(vecsize)
	return retVal
}

func GetThreadNum() int {
	retVal := CoreNative_getThreadNum_0()
	return retVal
}

func GetTickCount() int64 {
	retVal := CoreNative_getTickCount_0()
	return retVal
}

func GetTickFrequency() float64 {
	retVal := CoreNative_getTickFrequency_0()
	return retVal
}

func getVersion() string {
	return "3.3.1"
}
func getVersionMajor() int {
	return 3
}
func getVersionMinor() int {
	return 3
}
func getVersionRevision() int {
	return 1
}
func getVersionStatus() string {
	return ""
}
func Hconcat(src []*Mat, dst *Mat) {
	src_mat := ConvertersVectorMat(src)
	CoreNative_hconcat_0(src_mat.nativeObj, dst.nativeObj)
	return
}

func Idct(src *Mat, dst *Mat, flags int) {
	CoreNative_idct_0(src.nativeObj, dst.nativeObj, flags)
	return
}
func Idct2(src *Mat, dst *Mat) {
	CoreNative_idct_1(src.nativeObj, dst.nativeObj)
	return
}

func Idft(src *Mat, dst *Mat, flags int, nonzeroRows int) {
	CoreNative_idft_0(src.nativeObj, dst.nativeObj, flags, nonzeroRows)
	return
}
func Idft2(src *Mat, dst *Mat) {
	CoreNative_idft_1(src.nativeObj, dst.nativeObj)
	return
}

func InRange(src *Mat, lowerb *Scalar, upperb *Scalar, dst *Mat) {
	CoreNative_inRange_0(src.nativeObj, lowerb.Val[0], lowerb.Val[1], lowerb.Val[2], lowerb.Val[3], upperb.Val[0], upperb.Val[1], upperb.Val[2], upperb.Val[3], dst.nativeObj)
	return
}

func InsertChannel(src *Mat, dst *Mat, coi int) {
	CoreNative_insertChannel_0(src.nativeObj, dst.nativeObj, coi)
	return
}

func Invert(src *Mat, dst *Mat, flags int) float64 {
	retVal := CoreNative_invert_0(src.nativeObj, dst.nativeObj, flags)
	return retVal
}
func Invert2(src *Mat, dst *Mat) float64 {
	retVal := CoreNative_invert_1(src.nativeObj, dst.nativeObj)
	return retVal
}

func Kmeans(data *Mat, K int, bestLabels *Mat, criteria *TermCriteria, attempts int, flags int, centers *Mat) float64 {
	retVal := CoreNative_kmeans_0(data.nativeObj, K, bestLabels.nativeObj, criteria.Type, criteria.MaxCount, criteria.Epsilon, attempts, flags, centers.nativeObj)
	return retVal
}
func Kmeans2(data *Mat, K int, bestLabels *Mat, criteria *TermCriteria, attempts int, flags int) float64 {
	retVal := CoreNative_kmeans_1(data.nativeObj, K, bestLabels.nativeObj, criteria.Type, criteria.MaxCount, criteria.Epsilon, attempts, flags)
	return retVal
}

func Log(src *Mat, dst *Mat) {
	CoreNative_log_0(src.nativeObj, dst.nativeObj)
	return
}

func Magnitude(x *Mat, y *Mat, magnitude *Mat) {
	CoreNative_magnitude_0(x.nativeObj, y.nativeObj, magnitude.nativeObj)
	return
}

func Max(src1 *Mat, src2 *Mat, dst *Mat) {
	CoreNative_max_0(src1.nativeObj, src2.nativeObj, dst.nativeObj)
	return
}
func Max2(src1 *Mat, src2 *Scalar, dst *Mat) {
	CoreNative_max_1(src1.nativeObj, src2.Val[0], src2.Val[1], src2.Val[2], src2.Val[3], dst.nativeObj)
	return
}

func Mean(src *Mat, mask *Mat) *Scalar {
	retVal := NewScalar5(CoreNative_mean_0(src.nativeObj, mask.nativeObj))
	return retVal
}
func Mean2(src *Mat) *Scalar {
	retVal := NewScalar5(CoreNative_mean_1(src.nativeObj))
	return retVal
}
func MeanStdDev(src *Mat, mean *MatOfDouble, stddev *MatOfDouble, mask *Mat) {
	mean_mat := mean
	stddev_mat := stddev
	CoreNative_meanStdDev_0(src.nativeObj, mean_mat.nativeObj, stddev_mat.nativeObj, mask.nativeObj)
	return
}
func MeanStdDev2(src *Mat, mean *MatOfDouble, stddev *MatOfDouble) {
	mean_mat := mean
	stddev_mat := stddev
	CoreNative_meanStdDev_1(src.nativeObj, mean_mat.nativeObj, stddev_mat.nativeObj)
	return
}

func Merge(mv []*Mat, dst *Mat) {
	mv_mat := ConvertersVectorMat(mv)
	CoreNative_merge_0(mv_mat.nativeObj, dst.nativeObj)
	return
}

func Min(src1 *Mat, src2 *Mat, dst *Mat) {
	CoreNative_min_0(src1.nativeObj, src2.nativeObj, dst.nativeObj)
	return
}
func Min2(src1 *Mat, src2 *Scalar, dst *Mat) {
	CoreNative_min_1(src1.nativeObj, src2.Val[0], src2.Val[1], src2.Val[2], src2.Val[3], dst.nativeObj)
	return
}

func MinMaxLoc(src *Mat, mask *Mat) (minV float64, minP *Point, maxV float64, maxP *Point) {
	maskNativeObj := int64(0)
	if mask != nil {
		maskNativeObj = mask.nativeObj
	}
	resarr := CoreNative_n_minMaxLocManual(src.nativeObj, maskNativeObj)
	return resarr[0], NewPoint(resarr[2], resarr[3]),
		resarr[1], NewPoint(resarr[4], resarr[5])
}
func MinMaxLoc2(src *Mat) (minV float64, minP *Point, maxV float64, maxP *Point) {
	return MinMaxLoc(src, nil)
}

func MixChannels(src []*Mat, dst []*Mat, fromTo *MatOfInt) {
	src_mat := ConvertersVectorMat(src)
	dst_mat := ConvertersVectorMat(dst)
	fromTo_mat := fromTo
	CoreNative_mixChannels_0(src_mat.nativeObj, dst_mat.nativeObj, fromTo_mat.nativeObj)
	return
}

func MulSpectrums(a *Mat, b *Mat, c *Mat, flags int, conjB bool) {
	CoreNative_mulSpectrums_0(a.nativeObj, b.nativeObj, c.nativeObj, flags, conjB)
	return
}
func MulSpectrums2(a *Mat, b *Mat, c *Mat, flags int) {
	CoreNative_mulSpectrums_1(a.nativeObj, b.nativeObj, c.nativeObj, flags)
	return
}

func MulTransposed(src *Mat, dst *Mat, aTa bool, delta *Mat, scale float64, dtype int) {
	CoreNative_mulTransposed_0(src.nativeObj, dst.nativeObj, aTa, delta.nativeObj, scale, dtype)
	return
}
func MulTransposed2(src *Mat, dst *Mat, aTa bool, delta *Mat, scale float64) {
	CoreNative_mulTransposed_1(src.nativeObj, dst.nativeObj, aTa, delta.nativeObj, scale)
	return
}
func MulTransposed3(src *Mat, dst *Mat, aTa bool) {
	CoreNative_mulTransposed_2(src.nativeObj, dst.nativeObj, aTa)
	return
}

func Multiply(src1 *Mat, src2 *Mat, dst *Mat, scale float64, dtype int) {
	CoreNative_multiply_0(src1.nativeObj, src2.nativeObj, dst.nativeObj, scale, dtype)
	return
}
func Multiply2(src1 *Mat, src2 *Mat, dst *Mat, scale float64) {
	CoreNative_multiply_1(src1.nativeObj, src2.nativeObj, dst.nativeObj, scale)
	return
}
func Multiply3(src1 *Mat, src2 *Mat, dst *Mat) {
	CoreNative_multiply_2(src1.nativeObj, src2.nativeObj, dst.nativeObj)
	return
}
func Multiply4(src1 *Mat, src2 *Scalar, dst *Mat, scale float64, dtype int) {
	CoreNative_multiply_3(src1.nativeObj, src2.Val[0], src2.Val[1], src2.Val[2], src2.Val[3], dst.nativeObj, scale, dtype)
	return
}
func Multiply5(src1 *Mat, src2 *Scalar, dst *Mat, scale float64) {
	CoreNative_multiply_4(src1.nativeObj, src2.Val[0], src2.Val[1], src2.Val[2], src2.Val[3], dst.nativeObj, scale)
	return
}
func Multiply6(src1 *Mat, src2 *Scalar, dst *Mat) {
	CoreNative_multiply_5(src1.nativeObj, src2.Val[0], src2.Val[1], src2.Val[2], src2.Val[3], dst.nativeObj)
	return
}

func Norm(src1 *Mat, src2 *Mat, normType int, mask *Mat) float64 {
	retVal := CoreNative_norm_0(src1.nativeObj, src2.nativeObj, normType, mask.nativeObj)
	return retVal
}
func Norm2(src1 *Mat, src2 *Mat, normType int) float64 {
	retVal := CoreNative_norm_1(src1.nativeObj, src2.nativeObj, normType)
	return retVal
}
func Norm3(src1 *Mat, src2 *Mat) float64 {
	retVal := CoreNative_norm_2(src1.nativeObj, src2.nativeObj)
	return retVal
}
func Norm4(src1 *Mat, normType int, mask *Mat) float64 {
	retVal := CoreNative_norm_3(src1.nativeObj, normType, mask.nativeObj)
	return retVal
}
func Norm5(src1 *Mat, normType int) float64 {
	retVal := CoreNative_norm_4(src1.nativeObj, normType)
	return retVal
}
func Norm6(src1 *Mat) float64 {
	retVal := CoreNative_norm_5(src1.nativeObj)
	return retVal
}

func Normalize(src *Mat, dst *Mat, alpha float64, beta float64, norm_type int, dtype int, mask *Mat) {
	CoreNative_normalize_0(src.nativeObj, dst.nativeObj, alpha, beta, norm_type, dtype, mask.nativeObj)
	return
}
func Normalize2(src *Mat, dst *Mat, alpha float64, beta float64, norm_type int, dtype int) {
	CoreNative_normalize_1(src.nativeObj, dst.nativeObj, alpha, beta, norm_type, dtype)
	return
}
func Normalize3(src *Mat, dst *Mat, alpha float64, beta float64, norm_type int) {
	CoreNative_normalize_2(src.nativeObj, dst.nativeObj, alpha, beta, norm_type)
	return
}
func Normalize4(src *Mat, dst *Mat) {
	CoreNative_normalize_3(src.nativeObj, dst.nativeObj)
	return
}

func PatchNaNs(a *Mat, val float64) {
	CoreNative_patchNaNs_0(a.nativeObj, val)
	return
}
func PatchNaNs2(a *Mat) {
	CoreNative_patchNaNs_1(a.nativeObj)
	return
}

func PerspectiveTransform(src *Mat, dst *Mat, m *Mat) {
	CoreNative_perspectiveTransform_0(src.nativeObj, dst.nativeObj, m.nativeObj)
	return
}

func Phase(x *Mat, y *Mat, angle *Mat, angleInDegrees bool) {
	CoreNative_phase_0(x.nativeObj, y.nativeObj, angle.nativeObj, angleInDegrees)
	return
}
func Phase2(x *Mat, y *Mat, angle *Mat) {
	CoreNative_phase_1(x.nativeObj, y.nativeObj, angle.nativeObj)
	return
}

func PolarToCart(magnitude *Mat, angle *Mat, x *Mat, y *Mat, angleInDegrees bool) {
	CoreNative_polarToCart_0(magnitude.nativeObj, angle.nativeObj, x.nativeObj, y.nativeObj, angleInDegrees)
	return
}
func PolarToCart2(magnitude *Mat, angle *Mat, x *Mat, y *Mat) {
	CoreNative_polarToCart_1(magnitude.nativeObj, angle.nativeObj, x.nativeObj, y.nativeObj)
	return
}

func Pow(src *Mat, power float64, dst *Mat) {
	CoreNative_pow_0(src.nativeObj, power, dst.nativeObj)
	return
}

func RandShuffle(dst *Mat, iterFactor float64) {
	CoreNative_randShuffle_0(dst.nativeObj, iterFactor)
	return
}
func RandShuffle2(dst *Mat) {
	CoreNative_randShuffle_1(dst.nativeObj)
	return
}

func Randn(dst *Mat, mean float64, stddev float64) {
	CoreNative_randn_0(dst.nativeObj, mean, stddev)
	return
}

func Randu(dst *Mat, low float64, high float64) {
	CoreNative_randu_0(dst.nativeObj, low, high)
	return
}

func Reduce(src *Mat, dst *Mat, dim int, rtype int, dtype int) {
	CoreNative_reduce_0(src.nativeObj, dst.nativeObj, dim, rtype, dtype)
	return
}
func Reduce2(src *Mat, dst *Mat, dim int, rtype int) {
	CoreNative_reduce_1(src.nativeObj, dst.nativeObj, dim, rtype)
	return
}

func Repeat(src *Mat, ny int, nx int, dst *Mat) {
	CoreNative_repeat_0(src.nativeObj, ny, nx, dst.nativeObj)
	return
}

func Rotate(src *Mat, dst *Mat, rotateCode int) {
	CoreNative_rotate_0(src.nativeObj, dst.nativeObj, rotateCode)
	return
}

func ScaleAdd(src1 *Mat, alpha float64, src2 *Mat, dst *Mat) {
	CoreNative_scaleAdd_0(src1.nativeObj, alpha, src2.nativeObj, dst.nativeObj)
	return
}

func SetErrorVerbosity(verbose bool) {
	CoreNative_setErrorVerbosity_0(verbose)
	return
}

func SetIdentity(mtx *Mat, s *Scalar) {
	CoreNative_setIdentity_0(mtx.nativeObj, s.Val[0], s.Val[1], s.Val[2], s.Val[3])
	return
}
func SetIdentity2(mtx *Mat) {
	CoreNative_setIdentity_1(mtx.nativeObj)
	return
}

func SetNumThreads(nthreads int) {
	CoreNative_setNumThreads_0(nthreads)
	return
}

func SetRNGSeed(seed int) {
	CoreNative_setRNGSeed_0(seed)
	return
}

func SetUseIPP(flag bool) {
	CoreNative_setUseIPP_0(flag)
	return
}

func SetUseIPP_NE(flag bool) {
	CoreNative_setUseIPP_NE_0(flag)
	return
}

func Solve(src1 *Mat, src2 *Mat, dst *Mat, flags int) bool {
	retVal := CoreNative_solve_0(src1.nativeObj, src2.nativeObj, dst.nativeObj, flags)
	return retVal
}
func Solve2(src1 *Mat, src2 *Mat, dst *Mat) bool {
	retVal := CoreNative_solve_1(src1.nativeObj, src2.nativeObj, dst.nativeObj)
	return retVal
}
func SolveCubic(coeffs *Mat, roots *Mat) int {
	retVal := CoreNative_solveCubic_0(coeffs.nativeObj, roots.nativeObj)
	return retVal
}

func SolvePoly(coeffs *Mat, roots *Mat, maxIters int) float64 {
	retVal := CoreNative_solvePoly_0(coeffs.nativeObj, roots.nativeObj, maxIters)
	return retVal
}
func SolvePoly2(coeffs *Mat, roots *Mat) float64 {
	retVal := CoreNative_solvePoly_1(coeffs.nativeObj, roots.nativeObj)
	return retVal
}

func Sort(src *Mat, dst *Mat, flags int) {
	CoreNative_sort_0(src.nativeObj, dst.nativeObj, flags)
	return
}
func SortIdx(src *Mat, dst *Mat, flags int) {
	CoreNative_sortIdx_0(src.nativeObj, dst.nativeObj, flags)
	return
}

func Split(m *Mat) []*Mat {
	mv_mat := NewMat2()
	CoreNative_split_0(m.nativeObj, mv_mat.nativeObj)
	mv := ConvertersToVectorMat(mv_mat)
	mv_mat.Release()
	return mv
}

func Sqrt(src *Mat, dst *Mat) {
	CoreNative_sqrt_0(src.nativeObj, dst.nativeObj)
	return
}

func Subtract(src1 *Mat, src2 *Mat, dst *Mat, mask *Mat, dtype int) {
	CoreNative_subtract_0(src1.nativeObj, src2.nativeObj, dst.nativeObj, mask.nativeObj, dtype)
	return
}
func Subtract2(src1 *Mat, src2 *Mat, dst *Mat, mask *Mat) {
	CoreNative_subtract_1(src1.nativeObj, src2.nativeObj, dst.nativeObj, mask.nativeObj)
	return
}
func Subtract3(src1 *Mat, src2 *Mat, dst *Mat) {
	CoreNative_subtract_2(src1.nativeObj, src2.nativeObj, dst.nativeObj)
	return
}
func Subtract4(src1 *Mat, src2 *Scalar, dst *Mat, mask *Mat, dtype int) {
	CoreNative_subtract_3(src1.nativeObj, src2.Val[0], src2.Val[1], src2.Val[2], src2.Val[3], dst.nativeObj, mask.nativeObj, dtype)
	return
}
func Subtract5(src1 *Mat, src2 *Scalar, dst *Mat, mask *Mat) {
	CoreNative_subtract_4(src1.nativeObj, src2.Val[0], src2.Val[1], src2.Val[2], src2.Val[3], dst.nativeObj, mask.nativeObj)
	return
}
func Subtract6(src1 *Mat, src2 *Scalar, dst *Mat) {
	CoreNative_subtract_5(src1.nativeObj, src2.Val[0], src2.Val[1], src2.Val[2], src2.Val[3], dst.nativeObj)
	return
}

func SumElems(src *Mat) *Scalar {
	retVal := NewScalar5(CoreNative_sumElems_0(src.nativeObj))
	return retVal
}

func Trace(mtx *Mat) *Scalar {
	retVal := NewScalar5(CoreNative_trace_0(mtx.nativeObj))
	return retVal
}

func Transform(src *Mat, dst *Mat, m *Mat) {
	CoreNative_transform_0(src.nativeObj, dst.nativeObj, m.nativeObj)
	return
}

func Transpose(src *Mat, dst *Mat) {
	CoreNative_transpose_0(src.nativeObj, dst.nativeObj)
	return
}

func UseIPP() bool {
	retVal := CoreNative_useIPP_0()
	return retVal
}

func UseIPP_NE() bool {
	retVal := CoreNative_useIPP_NE_0()
	return retVal
}

func Vconcat(src []*Mat, dst *Mat) {
	src_mat := ConvertersVectorMat(src)
	CoreNative_vconcat_0(src_mat.nativeObj, dst.nativeObj)
	return
}
