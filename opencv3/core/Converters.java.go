package core

import (
	"fmt"

	. "github.com/gooid/gocv/opencv3/internal/native"
)

func matGetRCS(m *Mat) (rows, cols int) {
	rows = m.Rows()
	cols = m.Cols()
	if cols != 1 && rows == 1 {
		return cols, rows
	}
	return
}
func ConvertersToVectorDMatch(m *Mat) (matches []*DMatch) {
	count, matcols := matGetRCS(m)
	if CvTypeCV_64FC4 != m.Type() || matcols != 1 {
		Throw(NewIllegalArgumentException(fmt.Sprintf("%v%v", "CvTypeCV_64FC4 != m.Type() || m.Cols()!=1\n", m)))
	}
	buff := make([]float64, 4*count)
	m.GetD(0, 0, buff)
	for i := 0; i < count; i++ {
		matches = append(matches, NewDMatch3(int(buff[4*i]), int(buff[4*i+1]), int(buff[4*i+2]), float32(buff[4*i+3])))
	}
	return
}
func ConvertersToVectorKeyPoint(m *Mat) (kps []*KeyPoint) {
	count, matcols := matGetRCS(m)
	if CvTypeCV_64FC(7) != m.Type() || matcols != 1 {
		Throw(NewIllegalArgumentException(fmt.Sprintf("%v%v", "CvTypeCV_64FC(7) != m.Type() || m.Cols()!=1\n", m)))
	}
	buff := make([]float64, 7*count)
	m.GetD(0, 0, buff)
	for i := 0; i < count; i++ {
		kps = append(kps, NewKeyPoint(float32(buff[7*i]), float32(buff[7*i+1]), float32(buff[7*i+2]), float32(buff[7*i+3]), float32(buff[7*i+4]), int(buff[7*i+5]), int(buff[7*i+6])))
	}
	return
}
func ConvertersToVectorMat(m *Mat) (mats []*Mat) {
	count, matcols := matGetRCS(m)
	if CvTypeCV_32SC2 != m.Type() || matcols != 1 {
		Throw(NewIllegalArgumentException(fmt.Sprintf("%v%v", "CvTypeCV_32SC2 != m.Type() || m.Cols()!=1\n", m)))
	}
	buff := make([]int32, count*2)
	m.GetI(0, 0, buff)
	for i := 0; i < count; i++ {
		addr := int64(buff[i*2])<<32 | int64(buff[i*2+1])&0xffffffff
		mats = append(mats, NewMat(addr))
	}
	return
}
func ConvertersToVectorPoint(m *Mat) (pts []*Point) {
	count, matcols := matGetRCS(m)
	rtype := m.Type()
	if matcols != 1 {
		Throw(NewIllegalArgumentException(fmt.Sprintf("%v%v", "Input Mat should have one column\n", m)))
	}
	if rtype == CvTypeCV_32SC2 {
		buff := make([]int32, 2*count)
		m.GetI(0, 0, buff)
		for i := 0; i < count; i++ {
			pts = append(pts, NewPoint(float64(buff[i*2]), float64(buff[i*2+1])))
		}
	} else if rtype == CvTypeCV_32FC2 {
		buff := make([]float32, 2*count)
		m.GetF(0, 0, buff)
		for i := 0; i < count; i++ {
			pts = append(pts, NewPoint(float64(buff[i*2]), float64(buff[i*2+1])))
		}
	} else if rtype == CvTypeCV_64FC2 {
		buff := make([]float64, 2*count)
		m.GetD(0, 0, buff)
		for i := 0; i < count; i++ {
			pts = append(pts, NewPoint(float64(buff[i*2]), float64(buff[i*2+1])))
		}
	} else {
		Throw(NewIllegalArgumentException(fmt.Sprintf("%v%v", "Input Mat should be of CV_32SC2, CV_32FC2 or CV_64FC2 type\n", m)))
	}
	return
}
func ConvertersToVectorPoint2d(m *Mat) []*Point {
	return ConvertersToVectorPoint(m)
}
func ConvertersToVectorPoint2f(m *Mat) []*Point {
	return ConvertersToVectorPoint(m)
}
func ConvertersToVectorPoint3(m *Mat) (pts []*Point3) {
	count, matcols := matGetRCS(m)
	rtype := m.Type()
	if matcols != 1 {
		Throw(NewIllegalArgumentException(fmt.Sprintf("%v%v", "Input Mat should have one column\n", m)))
	}
	if rtype == CvTypeCV_32SC3 {
		buff := make([]int32, 3*count)
		m.GetI(0, 0, buff)
		for i := 0; i < count; i++ {
			pts = append(pts, NewPoint31(float64(buff[i*3]), float64(buff[i*3+1]), float64(buff[i*3+2])))
		}
	} else if rtype == CvTypeCV_32FC3 {
		buff := make([]float32, 3*count)
		m.GetF(0, 0, buff)
		for i := 0; i < count; i++ {
			pts = append(pts, NewPoint31(float64(buff[i*3]), float64(buff[i*3+1]), float64(buff[i*3+2])))
		}
	} else if rtype == CvTypeCV_64FC3 {
		buff := make([]float64, 3*count)
		m.GetD(0, 0, buff)
		for i := 0; i < count; i++ {
			pts = append(pts, NewPoint31(float64(buff[i*3]), float64(buff[i*3+1]), float64(buff[i*3+2])))
		}
	} else {
		Throw(NewIllegalArgumentException(fmt.Sprintf("%v%v", "Input Mat should be of CV_32SC3, CV_32FC3 or CV_64FC3 type\n", m)))
	}
	return
}
func ConvertersToVectorPoint3d(m *Mat) []*Point3 {
	return ConvertersToVectorPoint3(m)
}
func ConvertersToVectorPoint3f(m *Mat) []*Point3 {
	return ConvertersToVectorPoint3(m)
}
func ConvertersToVectorPoint3i(m *Mat) []*Point3 {
	return ConvertersToVectorPoint3(m)
}
func ConvertersToVectorRect(m *Mat) (rs []*Rect) {
	count, matcols := matGetRCS(m)
	if CvTypeCV_32SC4 != m.Type() || matcols != 1 {
		Throw(NewIllegalArgumentException(fmt.Sprintf("%v%v", "CvTypeCV_32SC4 != m.Type() || m.Cols()!=1\n", m)))
	}
	buff := make([]int32, 4*count)
	m.GetI(0, 0, buff)
	for i := 0; i < count; i++ {
		rs = append(rs, NewRect(int(buff[4*i]), int(buff[4*i+1]), int(buff[4*i+2]), int(buff[4*i+3])))
	}
	return
}
func ConvertersToVectorRect2d(m *Mat) (rs []*Rect2d) {
	count, matcols := matGetRCS(m)
	if CvTypeCV_64FC4 != m.Type() || matcols != 1 {
		Throw(NewIllegalArgumentException(fmt.Sprintf("%v%v", "CvTypeCV_64FC4 != m.Type() || m.Cols()!=1\n", m)))
	}
	buff := make([]float64, 4*count)
	m.GetD(0, 0, buff)
	for i := 0; i < count; i++ {
		rs = append(rs, NewRect2d(buff[4*i], buff[4*i+1], buff[4*i+2], buff[4*i+3]))
	}
	return
}
func ConvertersToVectorByte(m *Mat) (bs []byte) {
	count, matcols := matGetRCS(m)
	if CvTypeCV_8SC1 != m.Type() || matcols != 1 {
		Throw(NewIllegalArgumentException(fmt.Sprintf("%v%v", "CvTypeCV_8SC1 != m.Type() || m.Cols()!=1\n", m)))
	}
	buff := make([]byte, count)
	m.GetB(0, 0, buff)
	for i := 0; i < count; i++ {
		bs = append(bs, buff[i])
	}
	return
}
func ConvertersToVectorDouble(m *Mat) (ds []float64) {
	count, matcols := matGetRCS(m)
	if CvTypeCV_64FC1 != m.Type() || matcols != 1 {
		Throw(NewIllegalArgumentException(fmt.Sprintf("%v%v", "CvTypeCV_64FC1 != m.Type() || m.Cols()!=1\n", m)))
	}
	buff := make([]float64, count)
	m.GetD(0, 0, buff)
	for i := 0; i < count; i++ {
		ds = append(ds, buff[i])
	}
	return
}
func ConvertersToVectorFloat(m *Mat) (fs []float32) {
	count, matcols := matGetRCS(m)
	if CvTypeCV_32FC1 != m.Type() || matcols != 1 {
		Throw(NewIllegalArgumentException(fmt.Sprintf("%v%v", "CvTypeCV_32FC1 != m.Type() || m.Cols()!=1\n", m)))
	}
	buff := make([]float32, count)
	m.GetF(0, 0, buff)
	for i := 0; i < count; i++ {
		fs = append(fs, buff[i])
	}
	return
}
func ConvertersToVectorInt(m *Mat) (is []int32) {
	count, matcols := matGetRCS(m)
	if CvTypeCV_32SC1 != m.Type() || matcols != 1 {
		Throw(NewIllegalArgumentException(fmt.Sprintf("%v%v", "CvTypeCV_32SC1 != m.Type() || m.Cols()!=1\n", m)))
	}
	buff := make([]int32, count)
	m.GetI(0, 0, buff)
	for i := 0; i < count; i++ {
		is = append(is, buff[i])
	}
	return
}
func ConvertersToVectorUint8(m *Mat) (us []byte) {
	count, matcols := matGetRCS(m)
	if CvTypeCV_8UC1 != m.Type() || matcols != 1 {
		Throw(NewIllegalArgumentException(fmt.Sprintf("%v%v", "CvTypeCV_8UC1 != m.Type() || m.Cols()!=1\n", m)))
	}
	buff := make([]byte, count)
	m.GetB(0, 0, buff)
	for i := 0; i < count; i++ {
		us = append(us, buff[i])
	}
	return
}
func ConvertersToVectorVectorDMatch(m *Mat) (lvdm []*MatOfDMatch) {
	if m == nil {
		Throw(NewIllegalArgumentException("Input Mat can't be null"))
	}
	mats := ConvertersToVectorMat(m)
	for _, mi := range mats {
		vdm := NewMatOfDMatch3(mi)
		lvdm = append(lvdm, vdm)
		mi.Release()
	}

	return
}
func ConvertersToVectorVectorKeyPoint(m *Mat) (kps []*MatOfKeyPoint) {
	if m == nil {
		Throw(NewIllegalArgumentException("Input Mat can't be null"))
	}
	mats := ConvertersToVectorMat(m)
	for _, mi := range mats {
		vkp := NewMatOfKeyPoint3(mi)
		kps = append(kps, vkp)
		mi.Release()
	}

	return
}
func ConvertersToVectorVectorPoint(m *Mat) (pts []*MatOfPoint) {
	if m == nil {
		Throw(NewIllegalArgumentException("Input Mat can't be null"))
	}
	mats := ConvertersToVectorMat(m)
	for _, mi := range mats {
		pt := NewMatOfPoint3(mi)
		pts = append(pts, pt)
		mi.Release()
	}
	return
}
func ConvertersToVectorVectorPoint2f(m *Mat) (pts []*MatOfPoint2f) {
	if m == nil {
		Throw(NewIllegalArgumentException("Input Mat can't be null"))
	}
	mats := ConvertersToVectorMat(m)
	for _, mi := range mats {
		pt := NewMatOfPoint2f3(mi)
		pts = append(pts, pt)
		mi.Release()
	}

	return
}
func ConvertersToVectorVectorPoint3f(m *Mat, pts []*MatOfPoint3f) {
	if m == nil {
		Throw(NewIllegalArgumentException("Input Mat can't be null"))
	}
	mats := ConvertersToVectorMat(m)
	for _, mi := range mats {
		pt := NewMatOfPoint3f3(mi)
		pts = append(pts, pt)
		mi.Release()
	}

	return
}
func ConvertersToVectorVectorByte(m *Mat) (llb [][]byte) {
	if m == nil {
		Throw(NewIllegalArgumentException("Input Mat can't be null"))
	}
	mats := ConvertersToVectorMat(m)
	for _, mi := range mats {
		lb := ConvertersToVectorByte(mi)
		llb = append(llb, lb)
		mi.Release()
	}

	return
}
func ConvertersVectorDMatch(matches []*DMatch) *Mat {
	var res *Mat
	count := len(matches)
	if count > 0 {
		res = NewMat3(count, 1, CvTypeCV_64FC4)
		buff := make([]float64, count*4)
		for i := 0; i < count; i++ {
			m := matches[i]
			buff[4*i] = float64(m.QueryIdx)
			buff[4*i+1] = float64(m.TrainIdx)
			buff[4*i+2] = float64(m.ImgIdx)
			buff[4*i+3] = float64(m.Distance)
		}
		res.PutD(0, 0, buff)
	} else {
		res = NewMat2()
	}
	return res
}
func ConvertersVectorKeyPoint(kps []*KeyPoint) *Mat {
	var res *Mat
	count := len(kps)
	if count > 0 {
		res = NewMat3(count, 1, CvTypeCV_64FC(7))
		buff := make([]float64, count*7)
		for i := 0; i < count; i++ {
			kp := kps[i]
			buff[7*i] = float64(kp.Pt.X)
			buff[7*i+1] = float64(kp.Pt.Y)
			buff[7*i+2] = float64(kp.Size)
			buff[7*i+3] = float64(kp.Angle)
			buff[7*i+4] = float64(kp.Response)
			buff[7*i+5] = float64(kp.Octave)
			buff[7*i+6] = float64(kp.Class_id)
		}
		res.PutD(0, 0, buff)
	} else {
		res = NewMat2()
	}
	return res
}
func ConvertersVectorMat(mats []*Mat) *Mat {
	var res *Mat
	count := len(mats)
	if count > 0 {
		res = NewMat3(count, 1, CvTypeCV_32SC2)
		buff := make([]int32, count*2)
		for i := 0; i < count; i++ {
			addr := mats[i].nativeObj
			buff[i*2] = int32(addr >> 32)
			buff[i*2+1] = int32(addr & 0xffffffff)
		}
		res.PutI(0, 0, buff)
	} else {
		res = NewMat2()
	}
	return res
}
func ConvertersVectorPoint2d(pts []*Point) *Mat {
	return ConvertersVectorPoint(pts, CvTypeCV_64F)
}
func ConvertersVectorPoint2f(pts []*Point) *Mat {
	return ConvertersVectorPoint(pts, CvTypeCV_32F)
}
func ConvertersVectorPoint3(pts []*Point3, typeDepth int) *Mat {
	var res *Mat
	count := len(pts)
	if count > 0 {
		switch typeDepth {
		case CvTypeCV_32S:
			{
				res = NewMat3(count, 1, CvTypeCV_32SC3)
				buff := make([]int32, count*3)
				for i := 0; i < count; i++ {
					p := pts[i]
					buff[i*3] = int32(p.X)
					buff[i*3+1] = int32(p.Y)
					buff[i*3+2] = int32(p.Z)
				}
				res.PutI(0, 0, buff)
			}
		case CvTypeCV_32F:
			{
				res = NewMat3(count, 1, CvTypeCV_32FC3)
				buff := make([]float32, count*3)
				for i := 0; i < count; i++ {
					p := pts[i]
					buff[i*3] = float32(p.X)
					buff[i*3+1] = float32(p.Y)
					buff[i*3+2] = float32(p.Z)
				}
				res.PutF(0, 0, buff)
			}
		case CvTypeCV_64F:
			{
				res = NewMat3(count, 1, CvTypeCV_64FC3)
				buff := make([]float64, count*3)
				for i := 0; i < count; i++ {
					p := pts[i]
					buff[i*3] = p.X
					buff[i*3+1] = p.Y
					buff[i*3+2] = p.Z
				}
				res.PutD(0, 0, buff)
			}
		default:
			Throw(NewIllegalArgumentException("'typeDepth' can be CV_32S, CV_32F or CV_64F"))
		}
	} else {
		res = NewMat2()
	}
	return res
}
func ConvertersVectorPoint3d(pts []*Point3) *Mat {
	return ConvertersVectorPoint3(pts, CvTypeCV_64F)
}
func ConvertersVectorPoint3f(pts []*Point3) *Mat {
	return ConvertersVectorPoint3(pts, CvTypeCV_32F)
}
func ConvertersVectorPoint3i(pts []*Point3) *Mat {
	return ConvertersVectorPoint3(pts, CvTypeCV_32S)
}
func ConvertersVectorPoint(pts []*Point, typeDepth int) *Mat {
	var res *Mat
	count := len(pts)
	if count > 0 {
		switch typeDepth {
		case CvTypeCV_32S:
			{
				res = NewMat3(count, 1, CvTypeCV_32SC2)
				buff := make([]int32, count*2)
				for i := 0; i < count; i++ {
					p := pts[i]
					buff[i*2] = int32(p.X)
					buff[i*2+1] = int32(p.Y)
				}
				res.PutI(0, 0, buff)
			}
		case CvTypeCV_32F:
			{
				res = NewMat3(count, 1, CvTypeCV_32FC2)
				buff := make([]float32, count*2)
				for i := 0; i < count; i++ {
					p := pts[i]
					buff[i*2] = float32(p.X)
					buff[i*2+1] = float32(p.Y)
				}
				res.PutF(0, 0, buff)
			}
		case CvTypeCV_64F:
			{
				res = NewMat3(count, 1, CvTypeCV_64FC2)
				buff := make([]float64, count*2)
				for i := 0; i < count; i++ {
					p := pts[i]
					buff[i*2] = p.X
					buff[i*2+1] = p.Y
				}
				res.PutD(0, 0, buff)
			}
		default:
			Throw(NewIllegalArgumentException("'typeDepth' can be CV_32S, CV_32F or CV_64F"))
		}
	} else {
		res = NewMat2()
	}
	return res
}
func ConvertersVectorPoint2(pts []*Point) *Mat {
	return ConvertersVectorPoint(pts, CvTypeCV_32S)
}
func ConvertersVectorRect2d(rs []*Rect2d) *Mat {
	var res *Mat
	count := len(rs)
	if count > 0 {
		res = NewMat3(count, 1, CvTypeCV_64FC4)
		buff := make([]float64, 4*count)
		for i := 0; i < count; i++ {
			r := rs[i]
			buff[4*i] = float64(r.X)
			buff[4*i+1] = float64(r.Y)
			buff[4*i+2] = float64(r.Width)
			buff[4*i+3] = float64(r.Height)
		}
		res.PutD(0, 0, buff)
	} else {
		res = NewMat2()
	}
	return res
}
func ConvertersVectorRect(rs []*Rect) *Mat {
	var res *Mat
	count := len(rs)
	if count > 0 {
		res = NewMat3(count, 1, CvTypeCV_32SC4)
		buff := make([]int32, 4*count)
		for i := 0; i < count; i++ {
			r := rs[i]
			buff[4*i] = int32(r.X)
			buff[4*i+1] = int32(r.Y)
			buff[4*i+2] = int32(r.Width)
			buff[4*i+3] = int32(r.Height)
		}
		res.PutI(0, 0, buff)
	} else {
		res = NewMat2()
	}
	return res
}
func ConvertersVectorByte(bs []byte) *Mat {
	var res *Mat
	count := len(bs)
	if count > 0 {
		res = NewMat3(count, 1, CvTypeCV_8SC1)
		buff := make([]byte, count)
		for i := 0; i < count; i++ {
			b := bs[i]
			buff[i] = b
		}
		res.PutB(0, 0, buff)
	} else {
		res = NewMat2()
	}
	return res
}
func ConvertersVectorDouble(ds []float64) *Mat {
	var res *Mat
	count := len(ds)
	if count > 0 {
		res = NewMat3(count, 1, CvTypeCV_64FC1)
		buff := make([]float64, count)
		for i := 0; i < count; i++ {
			v := ds[i]
			buff[i] = v
		}
		res.PutD(0, 0, buff)
	} else {
		res = NewMat2()
	}
	return res
}
func ConvertersVectorFloat(fs []float32) *Mat {
	var res *Mat
	count := len(fs)
	if count > 0 {
		res = NewMat3(count, 1, CvTypeCV_32FC1)
		buff := make([]float32, count)
		for i := 0; i < count; i++ {
			f := fs[i]
			buff[i] = f
		}
		res.PutF(0, 0, buff)
	} else {
		res = NewMat2()
	}
	return res
}
func ConvertersVectorInt(is []int32) *Mat {
	var res *Mat
	count := len(is)
	if count > 0 {
		res = NewMat3(count, 1, CvTypeCV_32SC1)
		buff := make([]int32, count)
		for i := 0; i < count; i++ {
			v := is[i]
			buff[i] = v
		}
		res.PutI(0, 0, buff)
	} else {
		res = NewMat2()
	}
	return res
}
func ConvertersVectorUint8(bs []uint8) *Mat {
	var res *Mat
	count := len(bs)
	if count > 0 {
		res = NewMat3(count, 1, CvTypeCV_8UC1)
		buff := make([]byte, count)
		for i := 0; i < count; i++ {
			b := bs[i]
			buff[i] = b
		}
		res.PutB(0, 0, buff)
	} else {
		res = NewMat2()
	}
	return res
}
func ConvertersVectorVectorDMatch(lvdm []*MatOfDMatch) *Mat {
	var res *Mat
	if len(lvdm) > 0 {
		var mats []*Mat
		for _, vdm := range lvdm {
			mats = append(mats, vdm.Mat)
		}
		res = ConvertersVectorMat(mats)
	} else {
		res = NewMat2()
	}
	return res
}
func ConvertersVectorVectorKeyPoint(kps []*MatOfKeyPoint) *Mat {
	var res *Mat
	if len(kps) > 0 {
		var mats []*Mat
		for _, vkp := range kps {
			mats = append(mats, vkp.Mat)
		}
		res = ConvertersVectorMat(mats)
	} else {
		res = NewMat2()
	}
	return res
}
func ConvertersVectorVectorPoint2f(pts []*MatOfPoint2f) *Mat {
	var res *Mat
	if len(pts) > 0 {
		var mats []*Mat
		for _, vpt := range pts {
			mats = append(mats, vpt.Mat)
		}
		res = ConvertersVectorMat(mats)
	} else {
		res = NewMat2()
	}
	return res
}
func ConvertersVectorVectorPoint3f(pts []*MatOfPoint3f) *Mat {
	var res *Mat
	if len(pts) > 0 {
		var mats []*Mat
		for _, vpt := range pts {
			mats = append(mats, vpt.Mat)
		}
		res = ConvertersVectorMat(mats)
	} else {
		res = NewMat2()
	}
	return res
}
func ConvertersVectorVectorPoint(pts []*MatOfPoint) *Mat {
	var res *Mat
	if len(pts) > 0 {
		var mats []*Mat
		for _, vpt := range pts {
			mats = append(mats, vpt.Mat)
		}
		res = ConvertersVectorMat(mats)
	} else {
		res = NewMat2()
	}
	return res
}
func ConvertersVectorVectorByte(lvb []*MatOfByte) *Mat {
	var res *Mat
	if len(lvb) > 0 {
		var mats []*Mat
		for _, vb := range lvb {
			mats = append(mats, vb.Mat)
		}
		res = ConvertersVectorMat(mats)
	} else {
		res = NewMat2()
	}
	return res
}
