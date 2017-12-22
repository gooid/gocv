package imgproc

import (
	"fmt"
	"math"
)

type Moments struct {
	M00  float64
	M10  float64
	M01  float64
	M20  float64
	M11  float64
	M02  float64
	M30  float64
	M21  float64
	M12  float64
	M03  float64
	Mu20 float64
	Mu11 float64
	Mu02 float64
	Mu30 float64
	Mu21 float64
	Mu12 float64
	Mu03 float64
	Nu20 float64
	Nu11 float64
	Nu02 float64
	Nu30 float64
	Nu21 float64
	Nu12 float64
	Nu03 float64
}

func NewMoments(m00 float64, m10 float64, m01 float64, m20 float64, m11 float64, m02 float64, m30 float64, m21 float64, m12 float64, m03 float64) (rcvr *Moments) {
	rcvr = &Moments{}
	rcvr.M00 = m00
	rcvr.M10 = m10
	rcvr.M01 = m01
	rcvr.M20 = m20
	rcvr.M11 = m11
	rcvr.M02 = m02
	rcvr.M30 = m30
	rcvr.M21 = m21
	rcvr.M12 = m12
	rcvr.M03 = m03
	rcvr.CompleteState()
	return
}
func NewMoments2() (rcvr *Moments) {
	rcvr = NewMoments(0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
	return
}
func NewMoments3(vals []float64) (rcvr *Moments) {
	rcvr = &Moments{}
	rcvr.Set(vals)
	return
}
func (rcvr *Moments) CompleteState() {
	cx := float64(0)
	cy := float64(0)
	var mu20 float64
	var mu11 float64
	var mu02 float64
	inv_m00 := 0.0
	if math.Abs(rcvr.M00) > 0.00000001 {
		inv_m00 = 1. / rcvr.M00
		cx = rcvr.M10 * inv_m00
		cy = rcvr.M01 * inv_m00
	}
	mu20 = rcvr.M20 - rcvr.M10*cx
	mu11 = rcvr.M11 - rcvr.M10*cy
	mu02 = rcvr.M02 - rcvr.M01*cy
	rcvr.Mu20 = mu20
	rcvr.Mu11 = mu11
	rcvr.Mu02 = mu02
	rcvr.Mu30 = rcvr.M30 - cx*(3*mu20+cx*rcvr.M10)
	mu11 += mu11
	rcvr.Mu21 = rcvr.M21 - cx*(mu11+cx*rcvr.M01) - cy*mu20
	rcvr.Mu12 = rcvr.M12 - cy*(mu11+cy*rcvr.M10) - cx*mu02
	rcvr.Mu03 = rcvr.M03 - cy*(3*mu02+cy*rcvr.M01)
	inv_sqrt_m00 := math.Sqrt(math.Abs(inv_m00))
	s2 := inv_m00 * inv_m00
	s3 := s2 * inv_sqrt_m00
	rcvr.Nu20 = rcvr.Mu20 * s2
	rcvr.Nu11 = rcvr.Mu11 * s2
	rcvr.Nu02 = rcvr.Mu02 * s2
	rcvr.Nu30 = rcvr.Mu30 * s3
	rcvr.Nu21 = rcvr.Mu21 * s3
	rcvr.Nu12 = rcvr.Mu12 * s3
	rcvr.Nu03 = rcvr.Mu03 * s3
}
func (rcvr *Moments) Get_m00() float64 {
	return rcvr.M00
}
func (rcvr *Moments) Get_m01() float64 {
	return rcvr.M01
}
func (rcvr *Moments) Get_m02() float64 {
	return rcvr.M02
}
func (rcvr *Moments) Get_m03() float64 {
	return rcvr.M03
}
func (rcvr *Moments) Get_m10() float64 {
	return rcvr.M10
}
func (rcvr *Moments) Get_m11() float64 {
	return rcvr.M11
}
func (rcvr *Moments) Get_m12() float64 {
	return rcvr.M12
}
func (rcvr *Moments) Get_m20() float64 {
	return rcvr.M20
}
func (rcvr *Moments) Get_m21() float64 {
	return rcvr.M21
}
func (rcvr *Moments) Get_m30() float64 {
	return rcvr.M30
}
func (rcvr *Moments) Get_mu02() float64 {
	return rcvr.Mu02
}
func (rcvr *Moments) Get_mu03() float64 {
	return rcvr.Mu03
}
func (rcvr *Moments) Get_mu11() float64 {
	return rcvr.Mu11
}
func (rcvr *Moments) Get_mu12() float64 {
	return rcvr.Mu12
}
func (rcvr *Moments) Get_mu20() float64 {
	return rcvr.Mu20
}
func (rcvr *Moments) Get_mu21() float64 {
	return rcvr.Mu21
}
func (rcvr *Moments) Get_mu30() float64 {
	return rcvr.Mu30
}
func (rcvr *Moments) Get_nu02() float64 {
	return rcvr.Nu02
}
func (rcvr *Moments) Get_nu03() float64 {
	return rcvr.Nu03
}
func (rcvr *Moments) Get_nu11() float64 {
	return rcvr.Nu11
}
func (rcvr *Moments) Get_nu12() float64 {
	return rcvr.Nu12
}
func (rcvr *Moments) Get_nu20() float64 {
	return rcvr.Nu20
}
func (rcvr *Moments) Get_nu21() float64 {
	return rcvr.Nu21
}
func (rcvr *Moments) Get_nu30() float64 {
	return rcvr.Nu30
}
func (rcvr *Moments) Set(vals []float64) {
	if vals != nil {
		rcvr.M00 = func() float64 {
			if len(vals) > 0 {
				return vals[0]
			} else {
				return 0
			}
		}()
		rcvr.M10 = func() float64 {
			if len(vals) > 1 {
				return vals[1]
			} else {
				return 0
			}
		}()
		rcvr.M01 = func() float64 {
			if len(vals) > 2 {
				return vals[2]
			} else {
				return 0
			}
		}()
		rcvr.M20 = func() float64 {
			if len(vals) > 3 {
				return vals[3]
			} else {
				return 0
			}
		}()
		rcvr.M11 = func() float64 {
			if len(vals) > 4 {
				return vals[4]
			} else {
				return 0
			}
		}()
		rcvr.M02 = func() float64 {
			if len(vals) > 5 {
				return vals[5]
			} else {
				return 0
			}
		}()
		rcvr.M30 = func() float64 {
			if len(vals) > 6 {
				return vals[6]
			} else {
				return 0
			}
		}()
		rcvr.M21 = func() float64 {
			if len(vals) > 7 {
				return vals[7]
			} else {
				return 0
			}
		}()
		rcvr.M12 = func() float64 {
			if len(vals) > 8 {
				return vals[8]
			} else {
				return 0
			}
		}()
		rcvr.M03 = func() float64 {
			if len(vals) > 9 {
				return vals[9]
			} else {
				return 0
			}
		}()
		rcvr.CompleteState()
	} else {
		rcvr.M00 = 0
		rcvr.M10 = 0
		rcvr.M01 = 0
		rcvr.M20 = 0
		rcvr.M11 = 0
		rcvr.M02 = 0
		rcvr.M30 = 0
		rcvr.M21 = 0
		rcvr.M12 = 0
		rcvr.M03 = 0
		rcvr.Mu20 = 0
		rcvr.Mu11 = 0
		rcvr.Mu02 = 0
		rcvr.Mu30 = 0
		rcvr.Mu21 = 0
		rcvr.Mu12 = 0
		rcvr.Mu03 = 0
		rcvr.Nu20 = 0
		rcvr.Nu11 = 0
		rcvr.Nu02 = 0
		rcvr.Nu30 = 0
		rcvr.Nu21 = 0
		rcvr.Nu12 = 0
		rcvr.Nu03 = 0
	}
}
func (rcvr *Moments) Set_m00(m00 float64) {
	rcvr.M00 = m00
}
func (rcvr *Moments) Set_m01(m01 float64) {
	rcvr.M01 = m01
}
func (rcvr *Moments) Set_m02(m02 float64) {
	rcvr.M02 = m02
}
func (rcvr *Moments) Set_m03(m03 float64) {
	rcvr.M03 = m03
}
func (rcvr *Moments) Set_m10(m10 float64) {
	rcvr.M10 = m10
}
func (rcvr *Moments) Set_m11(m11 float64) {
	rcvr.M11 = m11
}
func (rcvr *Moments) Set_m12(m12 float64) {
	rcvr.M12 = m12
}
func (rcvr *Moments) Set_m20(m20 float64) {
	rcvr.M20 = m20
}
func (rcvr *Moments) Set_m21(m21 float64) {
	rcvr.M21 = m21
}
func (rcvr *Moments) Set_m30(m30 float64) {
	rcvr.M30 = m30
}
func (rcvr *Moments) Set_mu02(mu02 float64) {
	rcvr.Mu02 = mu02
}
func (rcvr *Moments) Set_mu03(mu03 float64) {
	rcvr.Mu03 = mu03
}
func (rcvr *Moments) Set_mu11(mu11 float64) {
	rcvr.Mu11 = mu11
}
func (rcvr *Moments) Set_mu12(mu12 float64) {
	rcvr.Mu12 = mu12
}
func (rcvr *Moments) Set_mu20(mu20 float64) {
	rcvr.Mu20 = mu20
}
func (rcvr *Moments) Set_mu21(mu21 float64) {
	rcvr.Mu21 = mu21
}
func (rcvr *Moments) Set_mu30(mu30 float64) {
	rcvr.Mu30 = mu30
}
func (rcvr *Moments) Set_nu02(nu02 float64) {
	rcvr.Nu02 = nu02
}
func (rcvr *Moments) Set_nu03(nu03 float64) {
	rcvr.Nu03 = nu03
}
func (rcvr *Moments) Set_nu11(nu11 float64) {
	rcvr.Nu11 = nu11
}
func (rcvr *Moments) Set_nu12(nu12 float64) {
	rcvr.Nu12 = nu12
}
func (rcvr *Moments) Set_nu20(nu20 float64) {
	rcvr.Nu20 = nu20
}
func (rcvr *Moments) Set_nu21(nu21 float64) {
	rcvr.Nu21 = nu21
}
func (rcvr *Moments) Set_nu30(nu30 float64) {
	rcvr.Nu30 = nu30
}
func (rcvr *Moments) String() string {
	return fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "Moments [ \nm00=", rcvr.M00, ", ", "\n", "m10=", rcvr.M10, ", ", "m01=", rcvr.M01, ", ", "\n", "m20=", rcvr.M20, ", ", "m11=", rcvr.M11, ", ", "m02=", rcvr.M02, ", ", "\n", "m30=", rcvr.M30, ", ", "m21=", rcvr.M21, ", ", "m12=", rcvr.M12, ", ", "m03=", rcvr.M03, ", ", "\n", "mu20=", rcvr.Mu20, ", ", "mu11=", rcvr.Mu11, ", ", "mu02=", rcvr.Mu02, ", ", "\n", "mu30=", rcvr.Mu30, ", ", "mu21=", rcvr.Mu21, ", ", "mu12=", rcvr.Mu12, ", ", "mu03=", rcvr.Mu03, ", ", "\n", "nu20=", rcvr.Nu20, ", ", "nu11=", rcvr.Nu11, ", ", "nu02=", rcvr.Nu02, ", ", "\n", "nu30=", rcvr.Nu30, ", ", "nu21=", rcvr.Nu21, ", ", "nu12=", rcvr.Nu12, ", ", "nu03=", rcvr.Nu03, ", ", "\n]")
}
