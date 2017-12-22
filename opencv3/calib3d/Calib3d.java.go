package calib3d

import (
	. "github.com/gooid/gocv/opencv3/core"
	. "github.com/gooid/gocv/opencv3/internal/native"
)

const CALIB_USE_INTRINSIC_GUESS = 1
const CALIB_RECOMPUTE_EXTRINSIC = 2
const CALIB_CHECK_COND = 4
const CALIB_FIX_SKEW = 8
const CALIB_FIX_K1 = 16
const CALIB_FIX_K2 = 32
const CALIB_FIX_K3 = 64
const CALIB_FIX_K4 = 128
const CALIB_FIX_INTRINSIC = 256
const CV_ITERATIVE = 0
const CV_EPNP = 1
const CV_P3P = 2
const CV_DLS = 3
const LMEDS = 4
const RANSAC = 8
const RHO = 16
const SOLVEPNP_ITERATIVE = 0
const SOLVEPNP_EPNP = 1
const SOLVEPNP_P3P = 2
const SOLVEPNP_DLS = 3
const SOLVEPNP_UPNP = 4
const SOLVEPNP_AP3P = 5
const CALIB_CB_ADAPTIVE_THRESH = 1
const CALIB_CB_NORMALIZE_IMAGE = 2
const CALIB_CB_FILTER_QUADS = 4
const CALIB_CB_FAST_CHECK = 8
const CALIB_CB_SYMMETRIC_GRID = 1
const CALIB_CB_ASYMMETRIC_GRID = 2
const CALIB_CB_CLUSTERING = 4
const CALIB_FIX_ASPECT_RATIO = 0x00002
const CALIB_FIX_PRINCIPAL_POINT = 0x00004
const CALIB_ZERO_TANGENT_DIST = 0x00008
const CALIB_FIX_FOCAL_LENGTH = 0x00010
const CALIB_FIX_K5 = 0x01000
const CALIB_FIX_K6 = 0x02000
const CALIB_RATIONAL_MODEL = 0x04000
const CALIB_THIN_PRISM_MODEL = 0x08000
const CALIB_FIX_S1_S2_S3_S4 = 0x10000
const CALIB_TILTED_MODEL = 0x40000
const CALIB_FIX_TAUX_TAUY = 0x80000
const CALIB_USE_QR = 0x100000
const CALIB_FIX_TANGENT_DIST = 0x200000
const CALIB_SAME_FOCAL_LENGTH = 0x00200
const CALIB_ZERO_DISPARITY = 0x00400
const FM_7POINT = 1
const FM_8POINT = 2
const FM_LMEDS = 4
const FM_RANSAC = 8

var Calib3dSOLVEPNP_MAX_COUNT = 5 + 1
var Calib3dCALIB_USE_LU = 1 << 17

func RQDecomp3x3(src *Mat, mtxR *Mat, mtxQ *Mat, Qx *Mat, Qy *Mat, Qz *Mat) []float64 {
	retVal := Calib3dNative_RQDecomp3x3_0(src.GetNativeObjAddr(), mtxR.GetNativeObjAddr(), mtxQ.GetNativeObjAddr(), Qx.GetNativeObjAddr(), Qy.GetNativeObjAddr(), Qz.GetNativeObjAddr())
	return retVal
}
func RQDecomp3x32(src *Mat, mtxR *Mat, mtxQ *Mat) []float64 {
	retVal := Calib3dNative_RQDecomp3x3_1(src.GetNativeObjAddr(), mtxR.GetNativeObjAddr(), mtxQ.GetNativeObjAddr())
	return retVal
}
func Rodrigues(src *Mat, dst *Mat, jacobian *Mat) {
	Calib3dNative_Rodrigues_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), jacobian.GetNativeObjAddr())
	return
}
func Rodrigues2(src *Mat, dst *Mat) {
	Calib3dNative_Rodrigues_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr())
	return
}
func Calibrate(objectPoints []*Mat, imagePoints []*Mat, image_size *Size, K *Mat, D *Mat, flags int, criteria *TermCriteria) ([]*Mat, []*Mat, float64) {
	objectPoints_mat := ConvertersVectorMat(objectPoints)
	imagePoints_mat := ConvertersVectorMat(imagePoints)
	rvecs_mat := NewMat2()
	tvecs_mat := NewMat2()
	retVal := Calib3dNative_calibrate_0(objectPoints_mat.GetNativeObjAddr(), imagePoints_mat.GetNativeObjAddr(), image_size.Width, image_size.Height, K.GetNativeObjAddr(), D.GetNativeObjAddr(), rvecs_mat.GetNativeObjAddr(), tvecs_mat.GetNativeObjAddr(), flags, criteria.Type, criteria.MaxCount, criteria.Epsilon)
	rvecs := ConvertersToVectorMat(rvecs_mat)
	rvecs_mat.Release()
	tvecs := ConvertersToVectorMat(tvecs_mat)
	tvecs_mat.Release()
	return rvecs, tvecs, retVal
}
func Calibrate2(objectPoints []*Mat, imagePoints []*Mat, image_size *Size, K *Mat, D *Mat, flags int) ([]*Mat, []*Mat, float64) {
	objectPoints_mat := ConvertersVectorMat(objectPoints)
	imagePoints_mat := ConvertersVectorMat(imagePoints)
	rvecs_mat := NewMat2()
	tvecs_mat := NewMat2()
	retVal := Calib3dNative_calibrate_1(objectPoints_mat.GetNativeObjAddr(), imagePoints_mat.GetNativeObjAddr(), image_size.Width, image_size.Height, K.GetNativeObjAddr(), D.GetNativeObjAddr(), rvecs_mat.GetNativeObjAddr(), tvecs_mat.GetNativeObjAddr(), flags)
	rvecs := ConvertersToVectorMat(rvecs_mat)
	rvecs_mat.Release()
	tvecs := ConvertersToVectorMat(tvecs_mat)
	tvecs_mat.Release()
	return rvecs, tvecs, retVal
}
func Calibrate3(objectPoints []*Mat, imagePoints []*Mat, image_size *Size, K *Mat, D *Mat) ([]*Mat, []*Mat, float64) {
	objectPoints_mat := ConvertersVectorMat(objectPoints)
	imagePoints_mat := ConvertersVectorMat(imagePoints)
	rvecs_mat := NewMat2()
	tvecs_mat := NewMat2()
	retVal := Calib3dNative_calibrate_2(objectPoints_mat.GetNativeObjAddr(), imagePoints_mat.GetNativeObjAddr(), image_size.Width, image_size.Height, K.GetNativeObjAddr(), D.GetNativeObjAddr(), rvecs_mat.GetNativeObjAddr(), tvecs_mat.GetNativeObjAddr())
	rvecs := ConvertersToVectorMat(rvecs_mat)
	rvecs_mat.Release()
	tvecs := ConvertersToVectorMat(tvecs_mat)
	tvecs_mat.Release()
	return rvecs, tvecs, retVal
}
func CalibrateCamera(objectPoints []*Mat, imagePoints []*Mat, imageSize *Size, cameraMatrix *Mat, distCoeffs *Mat, flags int, criteria *TermCriteria) ([]*Mat, []*Mat, float64) {
	objectPoints_mat := ConvertersVectorMat(objectPoints)
	imagePoints_mat := ConvertersVectorMat(imagePoints)
	rvecs_mat := NewMat2()
	tvecs_mat := NewMat2()
	retVal := Calib3dNative_calibrateCamera_0(objectPoints_mat.GetNativeObjAddr(), imagePoints_mat.GetNativeObjAddr(), imageSize.Width, imageSize.Height, cameraMatrix.GetNativeObjAddr(), distCoeffs.GetNativeObjAddr(), rvecs_mat.GetNativeObjAddr(), tvecs_mat.GetNativeObjAddr(), flags, criteria.Type, criteria.MaxCount, criteria.Epsilon)
	rvecs := ConvertersToVectorMat(rvecs_mat)
	rvecs_mat.Release()
	tvecs := ConvertersToVectorMat(tvecs_mat)
	tvecs_mat.Release()
	return rvecs, tvecs, retVal
}
func CalibrateCamera2(objectPoints []*Mat, imagePoints []*Mat, imageSize *Size, cameraMatrix *Mat, distCoeffs *Mat, flags int) ([]*Mat, []*Mat, float64) {
	objectPoints_mat := ConvertersVectorMat(objectPoints)
	imagePoints_mat := ConvertersVectorMat(imagePoints)
	rvecs_mat := NewMat2()
	tvecs_mat := NewMat2()
	retVal := Calib3dNative_calibrateCamera_1(objectPoints_mat.GetNativeObjAddr(), imagePoints_mat.GetNativeObjAddr(), imageSize.Width, imageSize.Height, cameraMatrix.GetNativeObjAddr(), distCoeffs.GetNativeObjAddr(), rvecs_mat.GetNativeObjAddr(), tvecs_mat.GetNativeObjAddr(), flags)
	rvecs := ConvertersToVectorMat(rvecs_mat)
	rvecs_mat.Release()
	tvecs := ConvertersToVectorMat(tvecs_mat)
	tvecs_mat.Release()
	return rvecs, tvecs, retVal
}
func CalibrateCamera3(objectPoints []*Mat, imagePoints []*Mat, imageSize *Size, cameraMatrix *Mat, distCoeffs *Mat) ([]*Mat, []*Mat, float64) {
	objectPoints_mat := ConvertersVectorMat(objectPoints)
	imagePoints_mat := ConvertersVectorMat(imagePoints)
	rvecs_mat := NewMat2()
	tvecs_mat := NewMat2()
	retVal := Calib3dNative_calibrateCamera_2(objectPoints_mat.GetNativeObjAddr(), imagePoints_mat.GetNativeObjAddr(), imageSize.Width, imageSize.Height, cameraMatrix.GetNativeObjAddr(), distCoeffs.GetNativeObjAddr(), rvecs_mat.GetNativeObjAddr(), tvecs_mat.GetNativeObjAddr())
	rvecs := ConvertersToVectorMat(rvecs_mat)
	rvecs_mat.Release()
	tvecs := ConvertersToVectorMat(tvecs_mat)
	tvecs_mat.Release()
	return rvecs, tvecs, retVal
}
func CalibrateCameraExtended(objectPoints []*Mat, imagePoints []*Mat, imageSize *Size, cameraMatrix *Mat, distCoeffs *Mat, stdDeviationsIntrinsics *Mat, stdDeviationsExtrinsics *Mat, perViewErrors *Mat, flags int, criteria *TermCriteria) ([]*Mat, []*Mat, float64) {
	objectPoints_mat := ConvertersVectorMat(objectPoints)
	imagePoints_mat := ConvertersVectorMat(imagePoints)
	rvecs_mat := NewMat2()
	tvecs_mat := NewMat2()
	retVal := Calib3dNative_calibrateCameraExtended_0(objectPoints_mat.GetNativeObjAddr(), imagePoints_mat.GetNativeObjAddr(), imageSize.Width, imageSize.Height, cameraMatrix.GetNativeObjAddr(), distCoeffs.GetNativeObjAddr(), rvecs_mat.GetNativeObjAddr(), tvecs_mat.GetNativeObjAddr(), stdDeviationsIntrinsics.GetNativeObjAddr(), stdDeviationsExtrinsics.GetNativeObjAddr(), perViewErrors.GetNativeObjAddr(), flags, criteria.Type, criteria.MaxCount, criteria.Epsilon)
	rvecs := ConvertersToVectorMat(rvecs_mat)
	rvecs_mat.Release()
	tvecs := ConvertersToVectorMat(tvecs_mat)
	tvecs_mat.Release()
	return rvecs, tvecs, retVal
}
func CalibrateCameraExtended2(objectPoints []*Mat, imagePoints []*Mat, imageSize *Size, cameraMatrix *Mat, distCoeffs *Mat, stdDeviationsIntrinsics *Mat, stdDeviationsExtrinsics *Mat, perViewErrors *Mat, flags int) ([]*Mat, []*Mat, float64) {
	objectPoints_mat := ConvertersVectorMat(objectPoints)
	imagePoints_mat := ConvertersVectorMat(imagePoints)
	rvecs_mat := NewMat2()
	tvecs_mat := NewMat2()
	retVal := Calib3dNative_calibrateCameraExtended_1(objectPoints_mat.GetNativeObjAddr(), imagePoints_mat.GetNativeObjAddr(), imageSize.Width, imageSize.Height, cameraMatrix.GetNativeObjAddr(), distCoeffs.GetNativeObjAddr(), rvecs_mat.GetNativeObjAddr(), tvecs_mat.GetNativeObjAddr(), stdDeviationsIntrinsics.GetNativeObjAddr(), stdDeviationsExtrinsics.GetNativeObjAddr(), perViewErrors.GetNativeObjAddr(), flags)
	rvecs := ConvertersToVectorMat(rvecs_mat)
	rvecs_mat.Release()
	tvecs := ConvertersToVectorMat(tvecs_mat)
	tvecs_mat.Release()
	return rvecs, tvecs, retVal
}
func CalibrateCameraExtended3(objectPoints []*Mat, imagePoints []*Mat, imageSize *Size, cameraMatrix *Mat, distCoeffs *Mat, stdDeviationsIntrinsics *Mat, stdDeviationsExtrinsics *Mat, perViewErrors *Mat) ([]*Mat, []*Mat, float64) {
	objectPoints_mat := ConvertersVectorMat(objectPoints)
	imagePoints_mat := ConvertersVectorMat(imagePoints)
	rvecs_mat := NewMat2()
	tvecs_mat := NewMat2()
	retVal := Calib3dNative_calibrateCameraExtended_2(objectPoints_mat.GetNativeObjAddr(), imagePoints_mat.GetNativeObjAddr(), imageSize.Width, imageSize.Height, cameraMatrix.GetNativeObjAddr(), distCoeffs.GetNativeObjAddr(), rvecs_mat.GetNativeObjAddr(), tvecs_mat.GetNativeObjAddr(), stdDeviationsIntrinsics.GetNativeObjAddr(), stdDeviationsExtrinsics.GetNativeObjAddr(), perViewErrors.GetNativeObjAddr())
	rvecs := ConvertersToVectorMat(rvecs_mat)
	rvecs_mat.Release()
	tvecs := ConvertersToVectorMat(tvecs_mat)
	tvecs_mat.Release()
	return rvecs, tvecs, retVal
}
func CalibrationMatrixValues(cameraMatrix *Mat, imageSize *Size, apertureWidth float64, apertureHeight float64, fovx []float64, fovy []float64, focalLength []float64, principalPoint *Point, aspectRatio []float64) {
	fovx_out := make([]float64, 1)
	fovy_out := make([]float64, 1)
	focalLength_out := make([]float64, 1)
	principalPoint_out := make([]float64, 2)
	aspectRatio_out := make([]float64, 1)
	Calib3dNative_calibrationMatrixValues_0(cameraMatrix.GetNativeObjAddr(), imageSize.Width, imageSize.Height, apertureWidth, apertureHeight, fovx_out, fovy_out, focalLength_out, principalPoint_out, aspectRatio_out)
	if fovx != nil {
		fovx[0] = fovx_out[0]
	}
	if fovy != nil {
		fovy[0] = fovy_out[0]
	}
	if focalLength != nil {
		focalLength[0] = focalLength_out[0]
	}
	if principalPoint != nil {
		principalPoint.X = principalPoint_out[0]
		principalPoint.Y = principalPoint_out[1]
	}
	if aspectRatio != nil {
		aspectRatio[0] = aspectRatio_out[0]
	}
	return
}
func ComposeRT(rvec1 *Mat, tvec1 *Mat, rvec2 *Mat, tvec2 *Mat, rvec3 *Mat, tvec3 *Mat, dr3dr1 *Mat, dr3dt1 *Mat, dr3dr2 *Mat, dr3dt2 *Mat, dt3dr1 *Mat, dt3dt1 *Mat, dt3dr2 *Mat, dt3dt2 *Mat) {
	Calib3dNative_composeRT_0(rvec1.GetNativeObjAddr(), tvec1.GetNativeObjAddr(), rvec2.GetNativeObjAddr(), tvec2.GetNativeObjAddr(), rvec3.GetNativeObjAddr(), tvec3.GetNativeObjAddr(), dr3dr1.GetNativeObjAddr(), dr3dt1.GetNativeObjAddr(), dr3dr2.GetNativeObjAddr(), dr3dt2.GetNativeObjAddr(), dt3dr1.GetNativeObjAddr(), dt3dt1.GetNativeObjAddr(), dt3dr2.GetNativeObjAddr(), dt3dt2.GetNativeObjAddr())
	return
}
func ComposeRT2(rvec1 *Mat, tvec1 *Mat, rvec2 *Mat, tvec2 *Mat, rvec3 *Mat, tvec3 *Mat) {
	Calib3dNative_composeRT_1(rvec1.GetNativeObjAddr(), tvec1.GetNativeObjAddr(), rvec2.GetNativeObjAddr(), tvec2.GetNativeObjAddr(), rvec3.GetNativeObjAddr(), tvec3.GetNativeObjAddr())
	return
}
func ComputeCorrespondEpilines(points *Mat, whichImage int, F *Mat, lines *Mat) {
	Calib3dNative_computeCorrespondEpilines_0(points.GetNativeObjAddr(), whichImage, F.GetNativeObjAddr(), lines.GetNativeObjAddr())
	return
}
func ConvertPointsFromHomogeneous(src *Mat, dst *Mat) {
	Calib3dNative_convertPointsFromHomogeneous_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr())
	return
}
func ConvertPointsToHomogeneous(src *Mat, dst *Mat) {
	Calib3dNative_convertPointsToHomogeneous_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr())
	return
}
func CorrectMatches(F *Mat, points1 *Mat, points2 *Mat, newPoints1 *Mat, newPoints2 *Mat) {
	Calib3dNative_correctMatches_0(F.GetNativeObjAddr(), points1.GetNativeObjAddr(), points2.GetNativeObjAddr(), newPoints1.GetNativeObjAddr(), newPoints2.GetNativeObjAddr())
	return
}
func DecomposeEssentialMat(E *Mat, R1 *Mat, R2 *Mat, t *Mat) {
	Calib3dNative_decomposeEssentialMat_0(E.GetNativeObjAddr(), R1.GetNativeObjAddr(), R2.GetNativeObjAddr(), t.GetNativeObjAddr())
	return
}
func DecomposeHomographyMat(H *Mat, K *Mat) ([]*Mat, []*Mat, []*Mat, int) {
	rotations_mat := NewMat2()
	translations_mat := NewMat2()
	normals_mat := NewMat2()
	retVal := Calib3dNative_decomposeHomographyMat_0(H.GetNativeObjAddr(), K.GetNativeObjAddr(), rotations_mat.GetNativeObjAddr(), translations_mat.GetNativeObjAddr(), normals_mat.GetNativeObjAddr())
	rotations := ConvertersToVectorMat(rotations_mat)
	rotations_mat.Release()
	translations := ConvertersToVectorMat(translations_mat)
	translations_mat.Release()
	normals := ConvertersToVectorMat(normals_mat)
	normals_mat.Release()
	return rotations, translations, normals, retVal
}
func DecomposeProjectionMatrix(projMatrix *Mat, cameraMatrix *Mat, rotMatrix *Mat, transVect *Mat, rotMatrixX *Mat, rotMatrixY *Mat, rotMatrixZ *Mat, eulerAngles *Mat) {
	Calib3dNative_decomposeProjectionMatrix_0(projMatrix.GetNativeObjAddr(), cameraMatrix.GetNativeObjAddr(), rotMatrix.GetNativeObjAddr(), transVect.GetNativeObjAddr(), rotMatrixX.GetNativeObjAddr(), rotMatrixY.GetNativeObjAddr(), rotMatrixZ.GetNativeObjAddr(), eulerAngles.GetNativeObjAddr())
	return
}
func DecomposeProjectionMatrix2(projMatrix *Mat, cameraMatrix *Mat, rotMatrix *Mat, transVect *Mat) {
	Calib3dNative_decomposeProjectionMatrix_1(projMatrix.GetNativeObjAddr(), cameraMatrix.GetNativeObjAddr(), rotMatrix.GetNativeObjAddr(), transVect.GetNativeObjAddr())
	return
}
func DistortPoints(undistorted *Mat, distorted *Mat, K *Mat, D *Mat, alpha float64) {
	Calib3dNative_distortPoints_0(undistorted.GetNativeObjAddr(), distorted.GetNativeObjAddr(), K.GetNativeObjAddr(), D.GetNativeObjAddr(), alpha)
	return
}
func DistortPoints2(undistorted *Mat, distorted *Mat, K *Mat, D *Mat) {
	Calib3dNative_distortPoints_1(undistorted.GetNativeObjAddr(), distorted.GetNativeObjAddr(), K.GetNativeObjAddr(), D.GetNativeObjAddr())
	return
}
func DrawChessboardCorners(image *Mat, patternSize *Size, corners *MatOfPoint2f, patternWasFound bool) {
	corners_mat := corners
	Calib3dNative_drawChessboardCorners_0(image.GetNativeObjAddr(), patternSize.Width, patternSize.Height, corners_mat.GetNativeObjAddr(), patternWasFound)
	return
}
func EstimateAffine2D(from *Mat, to *Mat, inliers *Mat, method int, ransacReprojThreshold float64, maxIters int64, confidence float64, refineIters int64) *Mat {
	retVal := NewMat(Calib3dNative_estimateAffine2D_0(from.GetNativeObjAddr(), to.GetNativeObjAddr(), inliers.GetNativeObjAddr(), method, ransacReprojThreshold, maxIters, confidence, refineIters))
	return retVal
}
func EstimateAffine2D2(from *Mat, to *Mat) *Mat {
	retVal := NewMat(Calib3dNative_estimateAffine2D_1(from.GetNativeObjAddr(), to.GetNativeObjAddr()))
	return retVal
}
func EstimateAffine3D(src *Mat, dst *Mat, out *Mat, inliers *Mat, ransacThreshold float64, confidence float64) int {
	retVal := Calib3dNative_estimateAffine3D_0(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), out.GetNativeObjAddr(), inliers.GetNativeObjAddr(), ransacThreshold, confidence)
	return retVal
}
func EstimateAffine3D2(src *Mat, dst *Mat, out *Mat, inliers *Mat) int {
	retVal := Calib3dNative_estimateAffine3D_1(src.GetNativeObjAddr(), dst.GetNativeObjAddr(), out.GetNativeObjAddr(), inliers.GetNativeObjAddr())
	return retVal
}
func EstimateAffinePartial2D(from *Mat, to *Mat, inliers *Mat, method int, ransacReprojThreshold float64, maxIters int64, confidence float64, refineIters int64) *Mat {
	retVal := NewMat(Calib3dNative_estimateAffinePartial2D_0(from.GetNativeObjAddr(), to.GetNativeObjAddr(), inliers.GetNativeObjAddr(), method, ransacReprojThreshold, maxIters, confidence, refineIters))
	return retVal
}
func EstimateAffinePartial2D2(from *Mat, to *Mat) *Mat {
	retVal := NewMat(Calib3dNative_estimateAffinePartial2D_1(from.GetNativeObjAddr(), to.GetNativeObjAddr()))
	return retVal
}
func EstimateNewCameraMatrixForUndistortRectify(K *Mat, D *Mat, image_size *Size, R *Mat, P *Mat, balance float64, new_size *Size, fov_scale float64) {
	Calib3dNative_estimateNewCameraMatrixForUndistortRectify_0(K.GetNativeObjAddr(), D.GetNativeObjAddr(), image_size.Width, image_size.Height, R.GetNativeObjAddr(), P.GetNativeObjAddr(), balance, new_size.Width, new_size.Height, fov_scale)
	return
}
func EstimateNewCameraMatrixForUndistortRectify2(K *Mat, D *Mat, image_size *Size, R *Mat, P *Mat) {
	Calib3dNative_estimateNewCameraMatrixForUndistortRectify_1(K.GetNativeObjAddr(), D.GetNativeObjAddr(), image_size.Width, image_size.Height, R.GetNativeObjAddr(), P.GetNativeObjAddr())
	return
}
func FilterSpeckles(img *Mat, newVal float64, maxSpeckleSize int, maxDiff float64, buf *Mat) {
	Calib3dNative_filterSpeckles_0(img.GetNativeObjAddr(), newVal, maxSpeckleSize, maxDiff, buf.GetNativeObjAddr())
	return
}
func FilterSpeckles2(img *Mat, newVal float64, maxSpeckleSize int, maxDiff float64) {
	Calib3dNative_filterSpeckles_1(img.GetNativeObjAddr(), newVal, maxSpeckleSize, maxDiff)
	return
}
func FindChessboardCorners(image *Mat, patternSize *Size, corners *MatOfPoint2f, flags int) bool {
	corners_mat := corners
	retVal := Calib3dNative_findChessboardCorners_0(image.GetNativeObjAddr(), patternSize.Width, patternSize.Height, corners_mat.GetNativeObjAddr(), flags)
	return retVal
}
func FindChessboardCorners2(image *Mat, patternSize *Size, corners *MatOfPoint2f) bool {
	corners_mat := corners
	retVal := Calib3dNative_findChessboardCorners_1(image.GetNativeObjAddr(), patternSize.Width, patternSize.Height, corners_mat.GetNativeObjAddr())
	return retVal
}
func FindCirclesGrid(image *Mat, patternSize *Size, centers *Mat, flags int) bool {
	retVal := Calib3dNative_findCirclesGrid_0(image.GetNativeObjAddr(), patternSize.Width, patternSize.Height, centers.GetNativeObjAddr(), flags)
	return retVal
}
func FindCirclesGrid2(image *Mat, patternSize *Size, centers *Mat) bool {
	retVal := Calib3dNative_findCirclesGrid_1(image.GetNativeObjAddr(), patternSize.Width, patternSize.Height, centers.GetNativeObjAddr())
	return retVal
}
func FindEssentialMat(points1 *Mat, points2 *Mat, cameraMatrix *Mat, method int, prob float64, threshold float64, mask *Mat) *Mat {
	retVal := NewMat(Calib3dNative_findEssentialMat_0(points1.GetNativeObjAddr(), points2.GetNativeObjAddr(), cameraMatrix.GetNativeObjAddr(), method, prob, threshold, mask.GetNativeObjAddr()))
	return retVal
}
func FindEssentialMat2(points1 *Mat, points2 *Mat, cameraMatrix *Mat, method int, prob float64, threshold float64) *Mat {
	retVal := NewMat(Calib3dNative_findEssentialMat_1(points1.GetNativeObjAddr(), points2.GetNativeObjAddr(), cameraMatrix.GetNativeObjAddr(), method, prob, threshold))
	return retVal
}
func FindEssentialMat3(points1 *Mat, points2 *Mat, cameraMatrix *Mat) *Mat {
	retVal := NewMat(Calib3dNative_findEssentialMat_2(points1.GetNativeObjAddr(), points2.GetNativeObjAddr(), cameraMatrix.GetNativeObjAddr()))
	return retVal
}
func FindEssentialMat4(points1 *Mat, points2 *Mat, focal float64, pp *Point, method int, prob float64, threshold float64, mask *Mat) *Mat {
	retVal := NewMat(Calib3dNative_findEssentialMat_3(points1.GetNativeObjAddr(), points2.GetNativeObjAddr(), focal, pp.X, pp.Y, method, prob, threshold, mask.GetNativeObjAddr()))
	return retVal
}
func FindEssentialMat5(points1 *Mat, points2 *Mat, focal float64, pp *Point, method int, prob float64, threshold float64) *Mat {
	retVal := NewMat(Calib3dNative_findEssentialMat_4(points1.GetNativeObjAddr(), points2.GetNativeObjAddr(), focal, pp.X, pp.Y, method, prob, threshold))
	return retVal
}
func FindEssentialMat6(points1 *Mat, points2 *Mat) *Mat {
	retVal := NewMat(Calib3dNative_findEssentialMat_5(points1.GetNativeObjAddr(), points2.GetNativeObjAddr()))
	return retVal
}
func FindFundamentalMat(points1 *MatOfPoint2f, points2 *MatOfPoint2f, method int, param1 float64, param2 float64, mask *Mat) *Mat {
	points1_mat := points1
	points2_mat := points2
	retVal := NewMat(Calib3dNative_findFundamentalMat_0(points1_mat.GetNativeObjAddr(), points2_mat.GetNativeObjAddr(), method, param1, param2, mask.GetNativeObjAddr()))
	return retVal
}
func FindFundamentalMat2(points1 *MatOfPoint2f, points2 *MatOfPoint2f, method int, param1 float64, param2 float64) *Mat {
	points1_mat := points1
	points2_mat := points2
	retVal := NewMat(Calib3dNative_findFundamentalMat_1(points1_mat.GetNativeObjAddr(), points2_mat.GetNativeObjAddr(), method, param1, param2))
	return retVal
}
func FindFundamentalMat3(points1 *MatOfPoint2f, points2 *MatOfPoint2f) *Mat {
	points1_mat := points1
	points2_mat := points2
	retVal := NewMat(Calib3dNative_findFundamentalMat_2(points1_mat.GetNativeObjAddr(), points2_mat.GetNativeObjAddr()))
	return retVal
}
func FindHomography(srcPoints *MatOfPoint2f, dstPoints *MatOfPoint2f, method int, ransacReprojThreshold float64, mask *Mat, maxIters int, confidence float64) *Mat {
	srcPoints_mat := srcPoints
	dstPoints_mat := dstPoints
	retVal := NewMat(Calib3dNative_findHomography_0(srcPoints_mat.GetNativeObjAddr(), dstPoints_mat.GetNativeObjAddr(), method, ransacReprojThreshold, mask.GetNativeObjAddr(), maxIters, confidence))
	return retVal
}
func FindHomography2(srcPoints *MatOfPoint2f, dstPoints *MatOfPoint2f, method int, ransacReprojThreshold float64) *Mat {
	srcPoints_mat := srcPoints
	dstPoints_mat := dstPoints
	retVal := NewMat(Calib3dNative_findHomography_1(srcPoints_mat.GetNativeObjAddr(), dstPoints_mat.GetNativeObjAddr(), method, ransacReprojThreshold))
	return retVal
}
func FindHomography3(srcPoints *MatOfPoint2f, dstPoints *MatOfPoint2f) *Mat {
	srcPoints_mat := srcPoints
	dstPoints_mat := dstPoints
	retVal := NewMat(Calib3dNative_findHomography_2(srcPoints_mat.GetNativeObjAddr(), dstPoints_mat.GetNativeObjAddr()))
	return retVal
}
func GetOptimalNewCameraMatrix(cameraMatrix *Mat, distCoeffs *Mat, imageSize *Size, alpha float64, newImgSize *Size, validPixROI *Rect, centerPrincipalPoint bool) *Mat {
	validPixROI_out := make([]float64, 4)
	retVal := NewMat(Calib3dNative_getOptimalNewCameraMatrix_0(cameraMatrix.GetNativeObjAddr(), distCoeffs.GetNativeObjAddr(), imageSize.Width, imageSize.Height, alpha, newImgSize.Width, newImgSize.Height, validPixROI_out, centerPrincipalPoint))
	if validPixROI != nil {
		validPixROI.X = int(validPixROI_out[0])
		validPixROI.Y = int(validPixROI_out[1])
		validPixROI.Width = int(validPixROI_out[2])
		validPixROI.Height = int(validPixROI_out[3])
	}
	return retVal
}
func GetOptimalNewCameraMatrix2(cameraMatrix *Mat, distCoeffs *Mat, imageSize *Size, alpha float64) *Mat {
	retVal := NewMat(Calib3dNative_getOptimalNewCameraMatrix_1(cameraMatrix.GetNativeObjAddr(), distCoeffs.GetNativeObjAddr(), imageSize.Width, imageSize.Height, alpha))
	return retVal
}
func GetValidDisparityROI(roi1 *Rect, roi2 *Rect, minDisparity int, numberOfDisparities int, SADWindowSize int) *Rect {
	retVal := NewRect5(Calib3dNative_getValidDisparityROI_0(roi1.X, roi1.Y, roi1.Width, roi1.Height, roi2.X, roi2.Y, roi2.Width, roi2.Height, minDisparity, numberOfDisparities, SADWindowSize))
	return retVal
}
func InitCameraMatrix2D(objectPoints []*MatOfPoint3f, imagePoints []*MatOfPoint2f, imageSize *Size, aspectRatio float64) *Mat {
	objectPoints_mat := ConvertersVectorVectorPoint3f(objectPoints)
	imagePoints_mat := ConvertersVectorVectorPoint2f(imagePoints)
	retVal := NewMat(Calib3dNative_initCameraMatrix2D_0(objectPoints_mat.GetNativeObjAddr(), imagePoints_mat.GetNativeObjAddr(), imageSize.Width, imageSize.Height, aspectRatio))
	return retVal
}
func InitCameraMatrix2D2(objectPoints []*MatOfPoint3f, imagePoints []*MatOfPoint2f, imageSize *Size) *Mat {
	objectPoints_mat := ConvertersVectorVectorPoint3f(objectPoints)
	imagePoints_mat := ConvertersVectorVectorPoint2f(imagePoints)
	retVal := NewMat(Calib3dNative_initCameraMatrix2D_1(objectPoints_mat.GetNativeObjAddr(), imagePoints_mat.GetNativeObjAddr(), imageSize.Width, imageSize.Height))
	return retVal
}
func InitUndistortRectifyMap(K *Mat, D *Mat, R *Mat, P *Mat, size *Size, m1type int, map1 *Mat, map2 *Mat) {
	Calib3dNative_initUndistortRectifyMap_0(K.GetNativeObjAddr(), D.GetNativeObjAddr(), R.GetNativeObjAddr(), P.GetNativeObjAddr(), size.Width, size.Height, m1type, map1.GetNativeObjAddr(), map2.GetNativeObjAddr())
	return
}
func MatMulDeriv(A *Mat, B *Mat, dABdA *Mat, dABdB *Mat) {
	Calib3dNative_matMulDeriv_0(A.GetNativeObjAddr(), B.GetNativeObjAddr(), dABdA.GetNativeObjAddr(), dABdB.GetNativeObjAddr())
	return
}
func ProjectPoints(objectPoints *MatOfPoint3f, rvec *Mat, tvec *Mat, cameraMatrix *Mat, distCoeffs *MatOfDouble, imagePoints *MatOfPoint2f, jacobian *Mat, aspectRatio float64) {
	objectPoints_mat := objectPoints
	distCoeffs_mat := distCoeffs
	imagePoints_mat := imagePoints
	Calib3dNative_projectPoints_0(objectPoints_mat.GetNativeObjAddr(), rvec.GetNativeObjAddr(), tvec.GetNativeObjAddr(), cameraMatrix.GetNativeObjAddr(), distCoeffs_mat.GetNativeObjAddr(), imagePoints_mat.GetNativeObjAddr(), jacobian.GetNativeObjAddr(), aspectRatio)
	return
}
func ProjectPoints2(objectPoints *MatOfPoint3f, rvec *Mat, tvec *Mat, cameraMatrix *Mat, distCoeffs *MatOfDouble, imagePoints *MatOfPoint2f) {
	objectPoints_mat := objectPoints
	distCoeffs_mat := distCoeffs
	imagePoints_mat := imagePoints
	Calib3dNative_projectPoints_1(objectPoints_mat.GetNativeObjAddr(), rvec.GetNativeObjAddr(), tvec.GetNativeObjAddr(), cameraMatrix.GetNativeObjAddr(), distCoeffs_mat.GetNativeObjAddr(), imagePoints_mat.GetNativeObjAddr())
	return
}
func ProjectPoints3(objectPoints *MatOfPoint3f, imagePoints *MatOfPoint2f, rvec *Mat, tvec *Mat, K *Mat, D *Mat, alpha float64, jacobian *Mat) {
	objectPoints_mat := objectPoints
	imagePoints_mat := imagePoints
	Calib3dNative_projectPoints_2(objectPoints_mat.GetNativeObjAddr(), imagePoints_mat.GetNativeObjAddr(), rvec.GetNativeObjAddr(), tvec.GetNativeObjAddr(), K.GetNativeObjAddr(), D.GetNativeObjAddr(), alpha, jacobian.GetNativeObjAddr())
	return
}
func ProjectPoints4(objectPoints *MatOfPoint3f, imagePoints *MatOfPoint2f, rvec *Mat, tvec *Mat, K *Mat, D *Mat) {
	objectPoints_mat := objectPoints
	imagePoints_mat := imagePoints
	Calib3dNative_projectPoints_3(objectPoints_mat.GetNativeObjAddr(), imagePoints_mat.GetNativeObjAddr(), rvec.GetNativeObjAddr(), tvec.GetNativeObjAddr(), K.GetNativeObjAddr(), D.GetNativeObjAddr())
	return
}
func RecoverPose(E *Mat, points1 *Mat, points2 *Mat, R *Mat, t *Mat, focal float64, pp *Point, mask *Mat) int {
	retVal := Calib3dNative_recoverPose_0(E.GetNativeObjAddr(), points1.GetNativeObjAddr(), points2.GetNativeObjAddr(), R.GetNativeObjAddr(), t.GetNativeObjAddr(), focal, pp.X, pp.Y, mask.GetNativeObjAddr())
	return retVal
}
func RecoverPose2(E *Mat, points1 *Mat, points2 *Mat, R *Mat, t *Mat, focal float64, pp *Point) int {
	retVal := Calib3dNative_recoverPose_1(E.GetNativeObjAddr(), points1.GetNativeObjAddr(), points2.GetNativeObjAddr(), R.GetNativeObjAddr(), t.GetNativeObjAddr(), focal, pp.X, pp.Y)
	return retVal
}
func RecoverPose3(E *Mat, points1 *Mat, points2 *Mat, R *Mat, t *Mat) int {
	retVal := Calib3dNative_recoverPose_2(E.GetNativeObjAddr(), points1.GetNativeObjAddr(), points2.GetNativeObjAddr(), R.GetNativeObjAddr(), t.GetNativeObjAddr())
	return retVal
}
func RecoverPose4(E *Mat, points1 *Mat, points2 *Mat, cameraMatrix *Mat, R *Mat, t *Mat, mask *Mat) int {
	retVal := Calib3dNative_recoverPose_3(E.GetNativeObjAddr(), points1.GetNativeObjAddr(), points2.GetNativeObjAddr(), cameraMatrix.GetNativeObjAddr(), R.GetNativeObjAddr(), t.GetNativeObjAddr(), mask.GetNativeObjAddr())
	return retVal
}
func RecoverPose5(E *Mat, points1 *Mat, points2 *Mat, cameraMatrix *Mat, R *Mat, t *Mat) int {
	retVal := Calib3dNative_recoverPose_4(E.GetNativeObjAddr(), points1.GetNativeObjAddr(), points2.GetNativeObjAddr(), cameraMatrix.GetNativeObjAddr(), R.GetNativeObjAddr(), t.GetNativeObjAddr())
	return retVal
}
func RecoverPose6(E *Mat, points1 *Mat, points2 *Mat, cameraMatrix *Mat, R *Mat, t *Mat, distanceThresh float64, mask *Mat, triangulatedPoints *Mat) int {
	retVal := Calib3dNative_recoverPose_5(E.GetNativeObjAddr(), points1.GetNativeObjAddr(), points2.GetNativeObjAddr(), cameraMatrix.GetNativeObjAddr(), R.GetNativeObjAddr(), t.GetNativeObjAddr(), distanceThresh, mask.GetNativeObjAddr(), triangulatedPoints.GetNativeObjAddr())
	return retVal
}
func RecoverPose7(E *Mat, points1 *Mat, points2 *Mat, cameraMatrix *Mat, R *Mat, t *Mat, distanceThresh float64) int {
	retVal := Calib3dNative_recoverPose_6(E.GetNativeObjAddr(), points1.GetNativeObjAddr(), points2.GetNativeObjAddr(), cameraMatrix.GetNativeObjAddr(), R.GetNativeObjAddr(), t.GetNativeObjAddr(), distanceThresh)
	return retVal
}
func Rectify3Collinear(cameraMatrix1 *Mat, distCoeffs1 *Mat, cameraMatrix2 *Mat, distCoeffs2 *Mat, cameraMatrix3 *Mat, distCoeffs3 *Mat, imgpt1 []*Mat, imgpt3 []*Mat, imageSize *Size, R12 *Mat, T12 *Mat, R13 *Mat, T13 *Mat, R1 *Mat, R2 *Mat, R3 *Mat, P1 *Mat, P2 *Mat, P3 *Mat, Q *Mat, alpha float64, newImgSize *Size, roi1 *Rect, roi2 *Rect, flags int) float32 {
	imgpt1_mat := ConvertersVectorMat(imgpt1)
	imgpt3_mat := ConvertersVectorMat(imgpt3)
	roi1_out := make([]float64, 4)
	roi2_out := make([]float64, 4)
	retVal := Calib3dNative_rectify3Collinear_0(cameraMatrix1.GetNativeObjAddr(), distCoeffs1.GetNativeObjAddr(), cameraMatrix2.GetNativeObjAddr(), distCoeffs2.GetNativeObjAddr(), cameraMatrix3.GetNativeObjAddr(), distCoeffs3.GetNativeObjAddr(), imgpt1_mat.GetNativeObjAddr(), imgpt3_mat.GetNativeObjAddr(), imageSize.Width, imageSize.Height, R12.GetNativeObjAddr(), T12.GetNativeObjAddr(), R13.GetNativeObjAddr(), T13.GetNativeObjAddr(), R1.GetNativeObjAddr(), R2.GetNativeObjAddr(), R3.GetNativeObjAddr(), P1.GetNativeObjAddr(), P2.GetNativeObjAddr(), P3.GetNativeObjAddr(), Q.GetNativeObjAddr(), alpha, newImgSize.Width, newImgSize.Height, roi1_out, roi2_out, flags)
	if roi1 != nil {
		roi1.X = int(roi1_out[0])
		roi1.Y = int(roi1_out[1])
		roi1.Width = int(roi1_out[2])
		roi1.Height = int(roi1_out[3])
	}
	if roi2 != nil {
		roi2.X = int(roi2_out[0])
		roi2.Y = int(roi2_out[1])
		roi2.Width = int(roi2_out[2])
		roi2.Height = int(roi2_out[3])
	}
	return retVal
}
func ReprojectImageTo3D(disparity *Mat, _3dImage *Mat, Q *Mat, handleMissingValues bool, ddepth int) {
	Calib3dNative_reprojectImageTo3D_0(disparity.GetNativeObjAddr(), _3dImage.GetNativeObjAddr(), Q.GetNativeObjAddr(), handleMissingValues, ddepth)
	return
}
func ReprojectImageTo3D2(disparity *Mat, _3dImage *Mat, Q *Mat, handleMissingValues bool) {
	Calib3dNative_reprojectImageTo3D_1(disparity.GetNativeObjAddr(), _3dImage.GetNativeObjAddr(), Q.GetNativeObjAddr(), handleMissingValues)
	return
}
func ReprojectImageTo3D3(disparity *Mat, _3dImage *Mat, Q *Mat) {
	Calib3dNative_reprojectImageTo3D_2(disparity.GetNativeObjAddr(), _3dImage.GetNativeObjAddr(), Q.GetNativeObjAddr())
	return
}
func SampsonDistance(pt1 *Mat, pt2 *Mat, F *Mat) float64 {
	retVal := Calib3dNative_sampsonDistance_0(pt1.GetNativeObjAddr(), pt2.GetNativeObjAddr(), F.GetNativeObjAddr())
	return retVal
}
func SolveP3P(objectPoints *Mat, imagePoints *Mat, cameraMatrix *Mat, distCoeffs *Mat, flags int) ([]*Mat, []*Mat, int) {
	rvecs_mat := NewMat2()
	tvecs_mat := NewMat2()
	retVal := Calib3dNative_solveP3P_0(objectPoints.GetNativeObjAddr(), imagePoints.GetNativeObjAddr(), cameraMatrix.GetNativeObjAddr(), distCoeffs.GetNativeObjAddr(), rvecs_mat.GetNativeObjAddr(), tvecs_mat.GetNativeObjAddr(), flags)
	rvecs := ConvertersToVectorMat(rvecs_mat)
	rvecs_mat.Release()
	tvecs := ConvertersToVectorMat(tvecs_mat)
	tvecs_mat.Release()
	return rvecs, tvecs, retVal
}
func SolvePnP(objectPoints *MatOfPoint3f, imagePoints *MatOfPoint2f, cameraMatrix *Mat, distCoeffs *MatOfDouble, rvec *Mat, tvec *Mat, useExtrinsicGuess bool, flags int) bool {
	objectPoints_mat := objectPoints
	imagePoints_mat := imagePoints
	distCoeffs_mat := distCoeffs
	retVal := Calib3dNative_solvePnP_0(objectPoints_mat.GetNativeObjAddr(), imagePoints_mat.GetNativeObjAddr(), cameraMatrix.GetNativeObjAddr(), distCoeffs_mat.GetNativeObjAddr(), rvec.GetNativeObjAddr(), tvec.GetNativeObjAddr(), useExtrinsicGuess, flags)
	return retVal
}
func SolvePnP2(objectPoints *MatOfPoint3f, imagePoints *MatOfPoint2f, cameraMatrix *Mat, distCoeffs *MatOfDouble, rvec *Mat, tvec *Mat) bool {
	objectPoints_mat := objectPoints
	imagePoints_mat := imagePoints
	distCoeffs_mat := distCoeffs
	retVal := Calib3dNative_solvePnP_1(objectPoints_mat.GetNativeObjAddr(), imagePoints_mat.GetNativeObjAddr(), cameraMatrix.GetNativeObjAddr(), distCoeffs_mat.GetNativeObjAddr(), rvec.GetNativeObjAddr(), tvec.GetNativeObjAddr())
	return retVal
}
func SolvePnPRansac(objectPoints *MatOfPoint3f, imagePoints *MatOfPoint2f, cameraMatrix *Mat, distCoeffs *MatOfDouble, rvec *Mat, tvec *Mat, useExtrinsicGuess bool, iterationsCount int, reprojectionError float32, confidence float64, inliers *Mat, flags int) bool {
	objectPoints_mat := objectPoints
	imagePoints_mat := imagePoints
	distCoeffs_mat := distCoeffs
	retVal := Calib3dNative_solvePnPRansac_0(objectPoints_mat.GetNativeObjAddr(), imagePoints_mat.GetNativeObjAddr(), cameraMatrix.GetNativeObjAddr(), distCoeffs_mat.GetNativeObjAddr(), rvec.GetNativeObjAddr(), tvec.GetNativeObjAddr(), useExtrinsicGuess, iterationsCount, reprojectionError, confidence, inliers.GetNativeObjAddr(), flags)
	return retVal
}
func SolvePnPRansac2(objectPoints *MatOfPoint3f, imagePoints *MatOfPoint2f, cameraMatrix *Mat, distCoeffs *MatOfDouble, rvec *Mat, tvec *Mat) bool {
	objectPoints_mat := objectPoints
	imagePoints_mat := imagePoints
	distCoeffs_mat := distCoeffs
	retVal := Calib3dNative_solvePnPRansac_1(objectPoints_mat.GetNativeObjAddr(), imagePoints_mat.GetNativeObjAddr(), cameraMatrix.GetNativeObjAddr(), distCoeffs_mat.GetNativeObjAddr(), rvec.GetNativeObjAddr(), tvec.GetNativeObjAddr())
	return retVal
}
func StereoCalibrate(objectPoints []*Mat, imagePoints1 []*Mat, imagePoints2 []*Mat, cameraMatrix1 *Mat, distCoeffs1 *Mat, cameraMatrix2 *Mat, distCoeffs2 *Mat, imageSize *Size, R *Mat, T *Mat, E *Mat, F *Mat, flags int, criteria *TermCriteria) float64 {
	objectPoints_mat := ConvertersVectorMat(objectPoints)
	imagePoints1_mat := ConvertersVectorMat(imagePoints1)
	imagePoints2_mat := ConvertersVectorMat(imagePoints2)
	retVal := Calib3dNative_stereoCalibrate_0(objectPoints_mat.GetNativeObjAddr(), imagePoints1_mat.GetNativeObjAddr(), imagePoints2_mat.GetNativeObjAddr(), cameraMatrix1.GetNativeObjAddr(), distCoeffs1.GetNativeObjAddr(), cameraMatrix2.GetNativeObjAddr(), distCoeffs2.GetNativeObjAddr(), imageSize.Width, imageSize.Height, R.GetNativeObjAddr(), T.GetNativeObjAddr(), E.GetNativeObjAddr(), F.GetNativeObjAddr(), flags, criteria.Type, criteria.MaxCount, criteria.Epsilon)
	return retVal
}
func StereoCalibrate2(objectPoints []*Mat, imagePoints1 []*Mat, imagePoints2 []*Mat, cameraMatrix1 *Mat, distCoeffs1 *Mat, cameraMatrix2 *Mat, distCoeffs2 *Mat, imageSize *Size, R *Mat, T *Mat, E *Mat, F *Mat, flags int) float64 {
	objectPoints_mat := ConvertersVectorMat(objectPoints)
	imagePoints1_mat := ConvertersVectorMat(imagePoints1)
	imagePoints2_mat := ConvertersVectorMat(imagePoints2)
	retVal := Calib3dNative_stereoCalibrate_1(objectPoints_mat.GetNativeObjAddr(), imagePoints1_mat.GetNativeObjAddr(), imagePoints2_mat.GetNativeObjAddr(), cameraMatrix1.GetNativeObjAddr(), distCoeffs1.GetNativeObjAddr(), cameraMatrix2.GetNativeObjAddr(), distCoeffs2.GetNativeObjAddr(), imageSize.Width, imageSize.Height, R.GetNativeObjAddr(), T.GetNativeObjAddr(), E.GetNativeObjAddr(), F.GetNativeObjAddr(), flags)
	return retVal
}
func StereoCalibrate3(objectPoints []*Mat, imagePoints1 []*Mat, imagePoints2 []*Mat, cameraMatrix1 *Mat, distCoeffs1 *Mat, cameraMatrix2 *Mat, distCoeffs2 *Mat, imageSize *Size, R *Mat, T *Mat, E *Mat, F *Mat) float64 {
	objectPoints_mat := ConvertersVectorMat(objectPoints)
	imagePoints1_mat := ConvertersVectorMat(imagePoints1)
	imagePoints2_mat := ConvertersVectorMat(imagePoints2)
	retVal := Calib3dNative_stereoCalibrate_2(objectPoints_mat.GetNativeObjAddr(), imagePoints1_mat.GetNativeObjAddr(), imagePoints2_mat.GetNativeObjAddr(), cameraMatrix1.GetNativeObjAddr(), distCoeffs1.GetNativeObjAddr(), cameraMatrix2.GetNativeObjAddr(), distCoeffs2.GetNativeObjAddr(), imageSize.Width, imageSize.Height, R.GetNativeObjAddr(), T.GetNativeObjAddr(), E.GetNativeObjAddr(), F.GetNativeObjAddr())
	return retVal
}
func StereoCalibrate4(objectPoints []*Mat, imagePoints1 []*Mat, imagePoints2 []*Mat, K1 *Mat, D1 *Mat, K2 *Mat, D2 *Mat, imageSize *Size, R *Mat, T *Mat, flags int, criteria *TermCriteria) float64 {
	objectPoints_mat := ConvertersVectorMat(objectPoints)
	imagePoints1_mat := ConvertersVectorMat(imagePoints1)
	imagePoints2_mat := ConvertersVectorMat(imagePoints2)
	retVal := Calib3dNative_stereoCalibrate_3(objectPoints_mat.GetNativeObjAddr(), imagePoints1_mat.GetNativeObjAddr(), imagePoints2_mat.GetNativeObjAddr(), K1.GetNativeObjAddr(), D1.GetNativeObjAddr(), K2.GetNativeObjAddr(), D2.GetNativeObjAddr(), imageSize.Width, imageSize.Height, R.GetNativeObjAddr(), T.GetNativeObjAddr(), flags, criteria.Type, criteria.MaxCount, criteria.Epsilon)
	return retVal
}
func StereoCalibrate5(objectPoints []*Mat, imagePoints1 []*Mat, imagePoints2 []*Mat, K1 *Mat, D1 *Mat, K2 *Mat, D2 *Mat, imageSize *Size, R *Mat, T *Mat, flags int) float64 {
	objectPoints_mat := ConvertersVectorMat(objectPoints)
	imagePoints1_mat := ConvertersVectorMat(imagePoints1)
	imagePoints2_mat := ConvertersVectorMat(imagePoints2)
	retVal := Calib3dNative_stereoCalibrate_4(objectPoints_mat.GetNativeObjAddr(), imagePoints1_mat.GetNativeObjAddr(), imagePoints2_mat.GetNativeObjAddr(), K1.GetNativeObjAddr(), D1.GetNativeObjAddr(), K2.GetNativeObjAddr(), D2.GetNativeObjAddr(), imageSize.Width, imageSize.Height, R.GetNativeObjAddr(), T.GetNativeObjAddr(), flags)
	return retVal
}
func StereoCalibrate6(objectPoints []*Mat, imagePoints1 []*Mat, imagePoints2 []*Mat, K1 *Mat, D1 *Mat, K2 *Mat, D2 *Mat, imageSize *Size, R *Mat, T *Mat) float64 {
	objectPoints_mat := ConvertersVectorMat(objectPoints)
	imagePoints1_mat := ConvertersVectorMat(imagePoints1)
	imagePoints2_mat := ConvertersVectorMat(imagePoints2)
	retVal := Calib3dNative_stereoCalibrate_5(objectPoints_mat.GetNativeObjAddr(), imagePoints1_mat.GetNativeObjAddr(), imagePoints2_mat.GetNativeObjAddr(), K1.GetNativeObjAddr(), D1.GetNativeObjAddr(), K2.GetNativeObjAddr(), D2.GetNativeObjAddr(), imageSize.Width, imageSize.Height, R.GetNativeObjAddr(), T.GetNativeObjAddr())
	return retVal
}
func StereoRectify(cameraMatrix1 *Mat, distCoeffs1 *Mat, cameraMatrix2 *Mat, distCoeffs2 *Mat, imageSize *Size, R *Mat, T *Mat, R1 *Mat, R2 *Mat, P1 *Mat, P2 *Mat, Q *Mat, flags int, alpha float64, newImageSize *Size, validPixROI1 *Rect, validPixROI2 *Rect) {
	validPixROI1_out := make([]float64, 4)
	validPixROI2_out := make([]float64, 4)
	Calib3dNative_stereoRectify_0(cameraMatrix1.GetNativeObjAddr(), distCoeffs1.GetNativeObjAddr(), cameraMatrix2.GetNativeObjAddr(), distCoeffs2.GetNativeObjAddr(), imageSize.Width, imageSize.Height, R.GetNativeObjAddr(), T.GetNativeObjAddr(), R1.GetNativeObjAddr(), R2.GetNativeObjAddr(), P1.GetNativeObjAddr(), P2.GetNativeObjAddr(), Q.GetNativeObjAddr(), flags, alpha, newImageSize.Width, newImageSize.Height, validPixROI1_out, validPixROI2_out)
	if validPixROI1 != nil {
		validPixROI1.X = int(validPixROI1_out[0])
		validPixROI1.Y = int(validPixROI1_out[1])
		validPixROI1.Width = int(validPixROI1_out[2])
		validPixROI1.Height = int(validPixROI1_out[3])
	}
	if validPixROI2 != nil {
		validPixROI2.X = int(validPixROI2_out[0])
		validPixROI2.Y = int(validPixROI2_out[1])
		validPixROI2.Width = int(validPixROI2_out[2])
		validPixROI2.Height = int(validPixROI2_out[3])
	}
	return
}
func StereoRectify2(cameraMatrix1 *Mat, distCoeffs1 *Mat, cameraMatrix2 *Mat, distCoeffs2 *Mat, imageSize *Size, R *Mat, T *Mat, R1 *Mat, R2 *Mat, P1 *Mat, P2 *Mat, Q *Mat) {
	Calib3dNative_stereoRectify_1(cameraMatrix1.GetNativeObjAddr(), distCoeffs1.GetNativeObjAddr(), cameraMatrix2.GetNativeObjAddr(), distCoeffs2.GetNativeObjAddr(), imageSize.Width, imageSize.Height, R.GetNativeObjAddr(), T.GetNativeObjAddr(), R1.GetNativeObjAddr(), R2.GetNativeObjAddr(), P1.GetNativeObjAddr(), P2.GetNativeObjAddr(), Q.GetNativeObjAddr())
	return
}
func StereoRectify3(K1 *Mat, D1 *Mat, K2 *Mat, D2 *Mat, imageSize *Size, R *Mat, tvec *Mat, R1 *Mat, R2 *Mat, P1 *Mat, P2 *Mat, Q *Mat, flags int, newImageSize *Size, balance float64, fov_scale float64) {
	Calib3dNative_stereoRectify_2(K1.GetNativeObjAddr(), D1.GetNativeObjAddr(), K2.GetNativeObjAddr(), D2.GetNativeObjAddr(), imageSize.Width, imageSize.Height, R.GetNativeObjAddr(), tvec.GetNativeObjAddr(), R1.GetNativeObjAddr(), R2.GetNativeObjAddr(), P1.GetNativeObjAddr(), P2.GetNativeObjAddr(), Q.GetNativeObjAddr(), flags, newImageSize.Width, newImageSize.Height, balance, fov_scale)
	return
}
func StereoRectify4(K1 *Mat, D1 *Mat, K2 *Mat, D2 *Mat, imageSize *Size, R *Mat, tvec *Mat, R1 *Mat, R2 *Mat, P1 *Mat, P2 *Mat, Q *Mat, flags int) {
	Calib3dNative_stereoRectify_3(K1.GetNativeObjAddr(), D1.GetNativeObjAddr(), K2.GetNativeObjAddr(), D2.GetNativeObjAddr(), imageSize.Width, imageSize.Height, R.GetNativeObjAddr(), tvec.GetNativeObjAddr(), R1.GetNativeObjAddr(), R2.GetNativeObjAddr(), P1.GetNativeObjAddr(), P2.GetNativeObjAddr(), Q.GetNativeObjAddr(), flags)
	return
}
func StereoRectifyUncalibrated(points1 *Mat, points2 *Mat, F *Mat, imgSize *Size, H1 *Mat, H2 *Mat, threshold float64) bool {
	retVal := Calib3dNative_stereoRectifyUncalibrated_0(points1.GetNativeObjAddr(), points2.GetNativeObjAddr(), F.GetNativeObjAddr(), imgSize.Width, imgSize.Height, H1.GetNativeObjAddr(), H2.GetNativeObjAddr(), threshold)
	return retVal
}
func StereoRectifyUncalibrated2(points1 *Mat, points2 *Mat, F *Mat, imgSize *Size, H1 *Mat, H2 *Mat) bool {
	retVal := Calib3dNative_stereoRectifyUncalibrated_1(points1.GetNativeObjAddr(), points2.GetNativeObjAddr(), F.GetNativeObjAddr(), imgSize.Width, imgSize.Height, H1.GetNativeObjAddr(), H2.GetNativeObjAddr())
	return retVal
}
func TriangulatePoints(projMatr1 *Mat, projMatr2 *Mat, projPoints1 *Mat, projPoints2 *Mat, points4D *Mat) {
	Calib3dNative_triangulatePoints_0(projMatr1.GetNativeObjAddr(), projMatr2.GetNativeObjAddr(), projPoints1.GetNativeObjAddr(), projPoints2.GetNativeObjAddr(), points4D.GetNativeObjAddr())
	return
}
func UndistortImage(distorted *Mat, undistorted *Mat, K *Mat, D *Mat, Knew *Mat, new_size *Size) {
	Calib3dNative_undistortImage_0(distorted.GetNativeObjAddr(), undistorted.GetNativeObjAddr(), K.GetNativeObjAddr(), D.GetNativeObjAddr(), Knew.GetNativeObjAddr(), new_size.Width, new_size.Height)
	return
}
func UndistortImage2(distorted *Mat, undistorted *Mat, K *Mat, D *Mat) {
	Calib3dNative_undistortImage_1(distorted.GetNativeObjAddr(), undistorted.GetNativeObjAddr(), K.GetNativeObjAddr(), D.GetNativeObjAddr())
	return
}
func UndistortPoints(distorted *Mat, undistorted *Mat, K *Mat, D *Mat, R *Mat, P *Mat) {
	Calib3dNative_undistortPoints_0(distorted.GetNativeObjAddr(), undistorted.GetNativeObjAddr(), K.GetNativeObjAddr(), D.GetNativeObjAddr(), R.GetNativeObjAddr(), P.GetNativeObjAddr())
	return
}
func UndistortPoints2(distorted *Mat, undistorted *Mat, K *Mat, D *Mat) {
	Calib3dNative_undistortPoints_1(distorted.GetNativeObjAddr(), undistorted.GetNativeObjAddr(), K.GetNativeObjAddr(), D.GetNativeObjAddr())
	return
}
func ValidateDisparity(disparity *Mat, cost *Mat, minDisparity int, numberOfDisparities int, disp12MaxDisp int) {
	Calib3dNative_validateDisparity_0(disparity.GetNativeObjAddr(), cost.GetNativeObjAddr(), minDisparity, numberOfDisparities, disp12MaxDisp)
	return
}
func ValidateDisparity2(disparity *Mat, cost *Mat, minDisparity int, numberOfDisparities int) {
	Calib3dNative_validateDisparity_1(disparity.GetNativeObjAddr(), cost.GetNativeObjAddr(), minDisparity, numberOfDisparities)
	return
}
