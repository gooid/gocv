package core

import (
	"fmt"
	"runtime"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

type Mat struct {
	nativeObj int64
}

func NewMat(addr int64) (rcvr *Mat) {
	rcvr = &Mat{}
	if addr == 0 {
		Throw(NewJavaLangUnsupportedOperationException("Native object address is NULL"))
	}
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	rcvr.nativeObj = addr
	return
}
func NewMat2() (rcvr *Mat) {
	rcvr = &Mat{}
	rcvr.nativeObj = MatNative_n_Mat()
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
	return
}
func NewMat3(rows int, cols int, rtype int) (rcvr *Mat) {
	rcvr = &Mat{}
	rcvr.nativeObj = MatNative_n_Mat2(rows, cols, rtype)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
	return
}
func NewMat4(size *Size, rtype int) (rcvr *Mat) {
	rcvr = &Mat{}
	rcvr.nativeObj = MatNative_n_Mat3(size.Width, size.Height, rtype)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
	return
}
func NewMat5(rows int, cols int, rtype int, s *Scalar) (rcvr *Mat) {
	rcvr = &Mat{}
	rcvr.nativeObj = MatNative_n_Mat4(rows, cols, rtype, s.Val[0], s.Val[1], s.Val[2], s.Val[3])
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
	return
}
func NewMat6(size *Size, rtype int, s *Scalar) (rcvr *Mat) {
	rcvr = &Mat{}
	rcvr.nativeObj = MatNative_n_Mat5(size.Width, size.Height, rtype, s.Val[0], s.Val[1], s.Val[2], s.Val[3])
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
	return
}
func NewMat7(m *Mat, rowRange *Range, colRange *Range) (rcvr *Mat) {
	rcvr = &Mat{}
	rcvr.nativeObj = MatNative_n_Mat6(m.nativeObj, rowRange.Start, rowRange.End, colRange.Start, colRange.End)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
	return
}
func NewMat8(m *Mat, rowRange *Range) (rcvr *Mat) {
	rcvr = &Mat{}
	rcvr.nativeObj = MatNative_n_Mat7(m.nativeObj, rowRange.Start, rowRange.End)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
	return
}
func NewMat9(m *Mat, roi *Rect) (rcvr *Mat) {
	rcvr = &Mat{}
	rcvr.nativeObj = MatNative_n_Mat6(m.nativeObj, roi.Y, roi.Y+roi.Height, roi.X, roi.X+roi.Width)
	runtime.SetFinalizer(rcvr, func(interface{}) { rcvr.finalize() })
	return
	return
}
func (rcvr *Mat) AdjustROI(dtop int, dbottom int, dleft int, dright int) *Mat {
	retVal := NewMat(MatNative_n_adjustROI(rcvr.nativeObj, dtop, dbottom, dleft, dright))
	return retVal
}
func (rcvr *Mat) AssignTo(m *Mat, rtype int) {
	MatNative_n_assignTo(rcvr.nativeObj, m.nativeObj, rtype)
	return
}
func (rcvr *Mat) AssignTo2(m *Mat) {
	MatNative_n_assignTo2(rcvr.nativeObj, m.nativeObj)
	return
}
func (rcvr *Mat) Channels() int {
	retVal := MatNative_n_channels(rcvr.nativeObj)
	return retVal
}
func (rcvr *Mat) CheckVector(elemChannels int, depth int, requireContinuous bool) int {
	retVal := MatNative_n_checkVector(rcvr.nativeObj, elemChannels, depth, requireContinuous)
	return retVal
}
func (rcvr *Mat) CheckVector2(elemChannels int, depth int) int {
	retVal := MatNative_n_checkVector2(rcvr.nativeObj, elemChannels, depth)
	return retVal
}
func (rcvr *Mat) CheckVector3(elemChannels int) int {
	retVal := MatNative_n_checkVector3(rcvr.nativeObj, elemChannels)
	return retVal
}
func (rcvr *Mat) Clone() *Mat {
	retVal := NewMat(MatNative_n_clone(rcvr.nativeObj))
	return retVal
}
func (rcvr *Mat) Col(x int) *Mat {
	retVal := NewMat(MatNative_n_col(rcvr.nativeObj, x))
	return retVal
}
func (rcvr *Mat) ColRange(startcol int, endcol int) *Mat {
	retVal := NewMat(MatNative_n_colRange(rcvr.nativeObj, startcol, endcol))
	return retVal
}
func (rcvr *Mat) ColRange2(r *Range) *Mat {
	retVal := NewMat(MatNative_n_colRange(rcvr.nativeObj, r.Start, r.End))
	return retVal
}
func (rcvr *Mat) Cols() int {
	retVal := MatNative_n_cols(rcvr.nativeObj)
	return retVal
}
func (rcvr *Mat) ConvertTo(m *Mat, rtype int, alpha float64, beta float64) {
	MatNative_n_convertTo(rcvr.nativeObj, m.nativeObj, rtype, alpha, beta)
	return
}
func (rcvr *Mat) ConvertTo2(m *Mat, rtype int, alpha float64) {
	MatNative_n_convertTo2(rcvr.nativeObj, m.nativeObj, rtype, alpha)
	return
}
func (rcvr *Mat) ConvertTo3(m *Mat, rtype int) {
	MatNative_n_convertTo3(rcvr.nativeObj, m.nativeObj, rtype)
	return
}
func (rcvr *Mat) CopyTo(m *Mat) {
	MatNative_n_copyTo(rcvr.nativeObj, m.nativeObj)
	return
}
func (rcvr *Mat) CopyTo2(m *Mat, mask *Mat) {
	MatNative_n_copyTo2(rcvr.nativeObj, m.nativeObj, mask.nativeObj)
	return
}
func (rcvr *Mat) Create(rows int, cols int, rtype int) {
	MatNative_n_create(rcvr.nativeObj, rows, cols, rtype)
	return
}
func (rcvr *Mat) Create2(size *Size, rtype int) {
	MatNative_n_create2(rcvr.nativeObj, size.Width, size.Height, rtype)
	return
}
func (rcvr *Mat) Cross(m *Mat) *Mat {
	retVal := NewMat(MatNative_n_cross(rcvr.nativeObj, m.nativeObj))
	return retVal
}
func (rcvr *Mat) DataAddr() int64 {
	retVal := MatNative_n_dataAddr(rcvr.nativeObj)
	return retVal
}
func (rcvr *Mat) Depth() int {
	retVal := MatNative_n_depth(rcvr.nativeObj)
	return retVal
}
func (rcvr *Mat) Diag(d int) *Mat {
	retVal := NewMat(MatNative_n_diag(rcvr.nativeObj, d))
	return retVal
}
func (rcvr *Mat) Diag2() *Mat {
	retVal := NewMat(MatNative_n_diag(rcvr.nativeObj, 0))
	return retVal
}
func Diag3(d *Mat) *Mat {
	retVal := NewMat(MatNative_n_diag2(d.nativeObj))
	return retVal
}
func (rcvr *Mat) Dims() int {
	retVal := MatNative_n_dims(rcvr.nativeObj)
	return retVal
}
func (rcvr *Mat) Dot(m *Mat) float64 {
	retVal := MatNative_n_dot(rcvr.nativeObj, m.nativeObj)
	return retVal
}
func (rcvr *Mat) Dump() string {
	return MatNative_nDump(rcvr.nativeObj)
}
func (rcvr *Mat) ElemSize() int64 {
	retVal := MatNative_n_elemSize(rcvr.nativeObj)
	return retVal
}
func (rcvr *Mat) ElemSize1() int64 {
	retVal := MatNative_n_elemSize1(rcvr.nativeObj)
	return retVal
}
func (rcvr *Mat) Empty() bool {
	retVal := MatNative_n_empty(rcvr.nativeObj)
	return retVal
}
func Eye(rows int, cols int, rtype int) *Mat {
	retVal := NewMat(MatNative_n_eye(rows, cols, rtype))
	return retVal
}
func Eye2(size *Size, rtype int) *Mat {
	retVal := NewMat(MatNative_n_eye2(size.Width, size.Height, rtype))
	return retVal
}
func (rcvr *Mat) finalize() {
	MatNative_n_delete(rcvr.nativeObj)
}
func (rcvr *Mat) GetB(row int, col int, data []byte) int {
	t := rcvr.Type()
	if data == nil || len(data)%CvTypeChannels(t) != 0 {
		Throw(NewJavaLangUnsupportedOperationException(fmt.Sprintf("%v%v%v%v%v", "Provided data element number (", func() int {
			if data == nil {
				return 0
			} else {
				return len(data)
			}
		}(), ") should be multiple of the Mat channels count (", CvTypeChannels(t), ")")))
	}
	if CvTypeDepth(t) == CvTypeCV_8U || CvTypeDepth(t) == CvTypeCV_8S {
		return MatNative_nGetB(rcvr.nativeObj, row, col, len(data), data)
	}
	Throw(NewJavaLangUnsupportedOperationException(fmt.Sprintf("%v%v", "Mat data type is not compatible: ", t)))
	return 0
}
func (rcvr *Mat) GetS(row int, col int, data []int16) int {
	t := rcvr.Type()
	if data == nil || len(data)%CvTypeChannels(t) != 0 {
		Throw(NewJavaLangUnsupportedOperationException(fmt.Sprintf("%v%v%v%v%v", "Provided data element number (", func() int {
			if data == nil {
				return 0
			} else {
				return len(data)
			}
		}(), ") should be multiple of the Mat channels count (", CvTypeChannels(t), ")")))
	}
	if CvTypeDepth(t) == CvTypeCV_16U || CvTypeDepth(t) == CvTypeCV_16S {
		return MatNative_nGetS(rcvr.nativeObj, row, col, len(data), data)
	}
	Throw(NewJavaLangUnsupportedOperationException(fmt.Sprintf("%v%v", "Mat data type is not compatible: ", t)))
	return 0
}
func (rcvr *Mat) GetI(row int, col int, data []int32) int {
	t := rcvr.Type()
	if data == nil || len(data)%CvTypeChannels(t) != 0 {
		Throw(NewJavaLangUnsupportedOperationException(fmt.Sprintf("%v%v%v%v%v", "Provided data element number (", func() int {
			if data == nil {
				return 0
			} else {
				return len(data)
			}
		}(), ") should be multiple of the Mat channels count (", CvTypeChannels(t), ")")))
	}
	if CvTypeDepth(t) == CvTypeCV_32S {
		return MatNative_nGetI(rcvr.nativeObj, row, col, len(data), data)
	}
	Throw(NewJavaLangUnsupportedOperationException(fmt.Sprintf("%v%v", "Mat data type is not compatible: ", t)))
	return 0
}
func (rcvr *Mat) GetF(row int, col int, data []float32) int {
	t := rcvr.Type()
	if data == nil || len(data)%CvTypeChannels(t) != 0 {
		Throw(NewJavaLangUnsupportedOperationException(fmt.Sprintf("%v%v%v%v%v", "Provided data element number (", func() int {
			if data == nil {
				return 0
			} else {
				return len(data)
			}
		}(), ") should be multiple of the Mat channels count (", CvTypeChannels(t), ")")))
	}
	if CvTypeDepth(t) == CvTypeCV_32F {
		return MatNative_nGetF(rcvr.nativeObj, row, col, len(data), data)
	}
	Throw(NewJavaLangUnsupportedOperationException(fmt.Sprintf("%v%v", "Mat data type is not compatible: ", t)))
	return 0
}
func (rcvr *Mat) GetD(row int, col int, data []float64) int {
	t := rcvr.Type()
	if data == nil || len(data)%CvTypeChannels(t) != 0 {
		Throw(NewJavaLangUnsupportedOperationException(fmt.Sprintf("%v%v%v%v%v", "Provided data element number (", func() int {
			if data == nil {
				return 0
			} else {
				return len(data)
			}
		}(), ") should be multiple of the Mat channels count (", CvTypeChannels(t), ")")))
	}
	if CvTypeDepth(t) == CvTypeCV_64F {
		return MatNative_nGetD(rcvr.nativeObj, row, col, len(data), data)
	}
	Throw(NewJavaLangUnsupportedOperationException(fmt.Sprintf("%v%v", "Mat data type is not compatible: ", t)))
	return 0
}
func (rcvr *Mat) Get6(row int, col int) []float64 {
	return MatNative_nGet(rcvr.nativeObj, row, col)
}
func (rcvr *Mat) GetNativeObjAddr() int64 {
	return rcvr.nativeObj
}
func (rcvr *Mat) Height() int {
	return rcvr.Rows()
}
func (rcvr *Mat) Inv(method int) *Mat {
	retVal := NewMat(MatNative_n_inv(rcvr.nativeObj, method))
	return retVal
}
func (rcvr *Mat) Inv2() *Mat {
	retVal := NewMat(MatNative_n_inv2(rcvr.nativeObj))
	return retVal
}
func (rcvr *Mat) IsContinuous() bool {
	retVal := MatNative_n_isContinuous(rcvr.nativeObj)
	return retVal
}
func (rcvr *Mat) IsSubmatrix() bool {
	retVal := MatNative_n_isSubmatrix(rcvr.nativeObj)
	return retVal
}
func (rcvr *Mat) LocateROI(wholeSize *Size, ofs *Point) {
	wholeSize_out := make([]float64, 2)
	ofs_out := make([]float64, 2)
	MatNative_locateROI_0(rcvr.nativeObj, wholeSize_out, ofs_out)
	if wholeSize != nil {
		wholeSize.Width = wholeSize_out[0]
		wholeSize.Height = wholeSize_out[1]
	}
	if ofs != nil {
		ofs.X = ofs_out[0]
		ofs.Y = ofs_out[1]
	}
	return
}
func MatNative_locateROI_0(nativeObj int64, wholeSize_out []float64, ofs_out []float64) {
}
func (rcvr *Mat) Mul(m *Mat, scale float64) *Mat {
	retVal := NewMat(MatNative_n_mul(rcvr.nativeObj, m.nativeObj, scale))
	return retVal
}
func (rcvr *Mat) Mul2(m *Mat) *Mat {
	retVal := NewMat(MatNative_n_mul2(rcvr.nativeObj, m.nativeObj))
	return retVal
}
func Ones(rows int, cols int, rtype int) *Mat {
	retVal := NewMat(MatNative_n_ones(rows, cols, rtype))
	return retVal
}
func Ones2(size *Size, rtype int) *Mat {
	retVal := NewMat(MatNative_n_ones2(size.Width, size.Height, rtype))
	return retVal
}
func (rcvr *Mat) Push_back(m *Mat) {
	MatNative_n_push_back(rcvr.nativeObj, m.nativeObj)
	return
}
func (rcvr *Mat) PutD(row int, col int, data []float64) int {
	t := rcvr.Type()
	if data == nil || len(data)%CvTypeChannels(t) != 0 {
		Throw(NewJavaLangUnsupportedOperationException(fmt.Sprintf("%v%v%v%v%v", "Provided data element number (", func() int {
			if data == nil {
				return 0
			} else {
				return len(data)
			}
		}(), ") should be multiple of the Mat channels count (", CvTypeChannels(t), ")")))
	}
	return MatNative_nPutD(rcvr.nativeObj, row, col, len(data), data)
}
func (rcvr *Mat) PutF(row int, col int, data []float32) int {
	t := rcvr.Type()
	if data == nil || len(data)%CvTypeChannels(t) != 0 {
		Throw(NewJavaLangUnsupportedOperationException(fmt.Sprintf("%v%v%v%v%v", "Provided data element number (", func() int {
			if data == nil {
				return 0
			} else {
				return len(data)
			}
		}(), ") should be multiple of the Mat channels count (", CvTypeChannels(t), ")")))
	}
	if CvTypeDepth(t) == CvTypeCV_32F {
		return MatNative_nPutF(rcvr.nativeObj, row, col, len(data), data)
	}
	Throw(NewJavaLangUnsupportedOperationException(fmt.Sprintf("%v%v", "Mat data type is not compatible: ", t)))
	return 0
}
func (rcvr *Mat) PutI(row int, col int, data []int32) int {
	t := rcvr.Type()
	if data == nil || len(data)%CvTypeChannels(t) != 0 {
		Throw(NewJavaLangUnsupportedOperationException(fmt.Sprintf("%v%v%v%v%v", "Provided data element number (", func() int {
			if data == nil {
				return 0
			} else {
				return len(data)
			}
		}(), ") should be multiple of the Mat channels count (", CvTypeChannels(t), ")")))
	}
	if CvTypeDepth(t) == CvTypeCV_32S {
		return MatNative_nPutI(rcvr.nativeObj, row, col, len(data), data)
	}
	Throw(NewJavaLangUnsupportedOperationException(fmt.Sprintf("%v%v", "Mat data type is not compatible: ", t)))
	return 0
}
func (rcvr *Mat) PutS(row int, col int, data []int16) int {
	t := rcvr.Type()
	if data == nil || len(data)%CvTypeChannels(t) != 0 {
		Throw(NewJavaLangUnsupportedOperationException(fmt.Sprintf("%v%v%v%v%v", "Provided data element number (", func() int {
			if data == nil {
				return 0
			} else {
				return len(data)
			}
		}(), ") should be multiple of the Mat channels count (", CvTypeChannels(t), ")")))
	}
	if CvTypeDepth(t) == CvTypeCV_16U || CvTypeDepth(t) == CvTypeCV_16S {
		return MatNative_nPutS(rcvr.nativeObj, row, col, len(data), data)
	}
	Throw(NewJavaLangUnsupportedOperationException(fmt.Sprintf("%v%v", "Mat data type is not compatible: ", t)))
	return 0
}
func (rcvr *Mat) PutB(row int, col int, data []byte) int {
	t := rcvr.Type()
	if data == nil || len(data)%CvTypeChannels(t) != 0 {
		Throw(NewJavaLangUnsupportedOperationException(fmt.Sprintf("%v%v%v%v%v", "Provided data element number (", func() int {
			if data == nil {
				return 0
			} else {
				return len(data)
			}
		}(), ") should be multiple of the Mat channels count (", CvTypeChannels(t), ")")))
	}
	if CvTypeDepth(t) == CvTypeCV_8U || CvTypeDepth(t) == CvTypeCV_8S {
		return MatNative_nPutB(rcvr.nativeObj, row, col, len(data), data)
	}
	Throw(NewJavaLangUnsupportedOperationException(fmt.Sprintf("%v%v", "Mat data type is not compatible: ", t)))
	return 0
}
func (rcvr *Mat) Release() {
	MatNative_n_release(rcvr.nativeObj)
	return
}
func (rcvr *Mat) Reshape(cn int, rows int) *Mat {
	retVal := NewMat(MatNative_n_reshape(rcvr.nativeObj, cn, rows))
	return retVal
}
func (rcvr *Mat) Reshape2(cn int) *Mat {
	retVal := NewMat(MatNative_n_reshape2(rcvr.nativeObj, cn))
	return retVal
}
func (rcvr *Mat) Row(y int) *Mat {
	retVal := NewMat(MatNative_n_row(rcvr.nativeObj, y))
	return retVal
}
func (rcvr *Mat) RowRange(startrow int, endrow int) *Mat {
	retVal := NewMat(MatNative_n_rowRange(rcvr.nativeObj, startrow, endrow))
	return retVal
}
func (rcvr *Mat) RowRange2(r *Range) *Mat {
	retVal := NewMat(MatNative_n_rowRange(rcvr.nativeObj, r.Start, r.End))
	return retVal
}
func (rcvr *Mat) Rows() int {
	retVal := MatNative_n_rows(rcvr.nativeObj)
	return retVal
}
func (rcvr *Mat) SetTo(s *Scalar) *Mat {
	retVal := NewMat(MatNative_n_setTo(rcvr.nativeObj, s.Val[0], s.Val[1], s.Val[2], s.Val[3]))
	return retVal
}
func (rcvr *Mat) SetTo2(value *Scalar, mask *Mat) *Mat {
	retVal := NewMat(MatNative_n_setTo2(rcvr.nativeObj, value.Val[0], value.Val[1], value.Val[2], value.Val[3], mask.nativeObj))
	return retVal
}
func (rcvr *Mat) SetTo3(value *Mat, mask *Mat) *Mat {
	retVal := NewMat(MatNative_n_setTo3(rcvr.nativeObj, value.nativeObj, mask.nativeObj))
	return retVal
}
func (rcvr *Mat) SetTo4(value *Mat) *Mat {
	retVal := NewMat(MatNative_n_setTo4(rcvr.nativeObj, value.nativeObj))
	return retVal
}
func (rcvr *Mat) Size() *Size {
	retVal := NewSize4(MatNative_n_size(rcvr.nativeObj))
	return retVal
}
func (rcvr *Mat) Step1(i int) int64 {
	retVal := MatNative_n_step1(rcvr.nativeObj, i)
	return retVal
}
func (rcvr *Mat) Step12() int64 {
	retVal := MatNative_n_step12(rcvr.nativeObj)
	return retVal
}
func (rcvr *Mat) Submat(rowStart int, rowEnd int, colStart int, colEnd int) *Mat {
	retVal := NewMat(MatNative_n_submat_rr(rcvr.nativeObj, rowStart, rowEnd, colStart, colEnd))
	return retVal
}
func (rcvr *Mat) Submat2(rowRange *Range, colRange *Range) *Mat {
	retVal := NewMat(MatNative_n_submat_rr(rcvr.nativeObj, rowRange.Start, rowRange.End, colRange.Start, colRange.End))
	return retVal
}
func (rcvr *Mat) Submat3(roi *Rect) *Mat {
	retVal := NewMat(MatNative_n_submat(rcvr.nativeObj, roi.X, roi.Y, roi.Width, roi.Height))
	return retVal
}
func (rcvr *Mat) T() *Mat {
	retVal := NewMat(MatNative_n_t(rcvr.nativeObj))
	return retVal
}
func (rcvr *Mat) String() string {
	return fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v", "Mat [ ", rcvr.Rows(), "*", rcvr.Cols(), "*", TypeToString(rcvr.Type()), ", isCont=", rcvr.IsContinuous(), ", isSubmat=", rcvr.IsSubmatrix(), fmt.Sprintf(", nativeObj=0x%08X", rcvr.nativeObj), fmt.Sprintf(", dataAddr=0x%08X", rcvr.DataAddr()), " ]")
}
func (rcvr *Mat) Total() int64 {
	retVal := MatNative_n_total(rcvr.nativeObj)
	return retVal
}
func (rcvr *Mat) Type() int {
	retVal := MatNative_n_type(rcvr.nativeObj)
	return retVal
}
func (rcvr *Mat) Width() int {
	return rcvr.Cols()
}
func Zeros(rows int, cols int, rtype int) *Mat {
	retVal := NewMat(MatNative_n_zeros(rows, cols, rtype))
	return retVal
}
func Zeros2(size *Size, rtype int) *Mat {
	retVal := NewMat(MatNative_n_zeros2(size.Width, size.Height, rtype))
	return retVal
}
