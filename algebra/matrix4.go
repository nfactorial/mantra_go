package algebra

import (
	"fmt"
	"math"
)

type Matrix4 struct {
	M00 MnFloat
	M01 MnFloat
	M02 MnFloat
	M03 MnFloat
	M10 MnFloat
	M11 MnFloat
	M12 MnFloat
	M13 MnFloat
	M20 MnFloat
	M21 MnFloat
	M22 MnFloat
	M23 MnFloat
	M30 MnFloat
	M31 MnFloat
	M32 MnFloat
	M33 MnFloat
}

var IdentityMatrix4 = Matrix4{
	M00: MnOne,
	M11: MnOne,
	M22: MnOne,
	M33: MnOne,
}

func NewRotationMatrix4(right Vector3, up Vector3, forward Vector3) Matrix4 {
	return Matrix4{
		M00: right.X,
		M01: right.Y,
		M02: right.Z,
		M03: 0,
		M10: up.X,
		M11: up.Y,
		M12: up.Z,
		M13: 0,
		M20: forward.X,
		M21: forward.Y,
		M22: forward.Z,
		M23: 0,
		M30: 0,
		M31: 0,
		M32: 0,
		M33: 1,
	}
}

func NewScaleMatrix4(x MnFloat, y MnFloat, z MnFloat) Matrix4 {
	return Matrix4{
		M00: x,
		M11: y,
		M22: z,
		M33: MnOne,
	}
}

func NewTranslationMatrix4(x MnFloat, y MnFloat, z MnFloat) Matrix4 {
	return Matrix4{
		M00: MnOne,
		M11: MnOne,
		M22: MnOne,
		M30: x,
		M31: y,
		M32: z,
		M33: MnOne,
	}
}

func NewPerspectiveMatrix4(fov MnFloat, aspect MnFloat, near MnFloat, far MnFloat) Matrix4 {
	f := MnOne / math.Tan(fov / 2.0)
	nf := MnOne / (near - far)

	return Matrix4{
		M00: f / aspect,
		M11: f,
		M22: (far + near) * nf,
		M23: -MnOne,
		M32: 2 * far * near * nf,
	}
}

func (m Matrix4) Transpose() Matrix4 {
	return Matrix4{
		M00: m.M00,
		M01: m.M10,
		M02: m.M20,
		M03: m.M30,
		M10: m.M01,
		M11: m.M11,
		M12: m.M21,
		M13: m.M31,
		M20: m.M02,
		M21: m.M12,
		M22: m.M22,
		M23: m.M32,
		M30: m.M03,
		M31: m.M13,
		M32: m.M23,
		M33: m.M33,
	}
}

func (m Matrix4) Invert() Matrix4 {
	a00 := m.M00 * m.M11 - m.M01 * m.M10
	a01 := m.M00 * m.M12 - m.M02 * m.M10
	a02 := m.M00 * m.M13 - m.M03 * m.M10
	a03 := m.M01 * m.M12 - m.M02 * m.M11
	a04 := m.M01 * m.M13 - m.M03 * m.M11
	a05 := m.M02 * m.M13 - m.M03 * m.M12
	a06 := m.M20 * m.M31 - m.M21 * m.M30
	a07 := m.M20 * m.M32 - m.M22 * m.M30
	a08 := m.M20 * m.M33 - m.M23 * m.M30
	a09 := m.M21 * m.M32 - m.M22 * m.M31
	a10 := m.M21 * m.M33 - m.M23 * m.M31
	a11 := m.M22 * m.M33 - m.M23 * m.M32

	determinant := a00 * a11 - a01 * a10 + a02 * a09 + a03 * a08 - a04 * a07 + a05 * a06

	if determinant == MnZero {		// TODO: Shouldn't compare with 0
		return Matrix4{}
	}

	determinant = MnOne / determinant

	return Matrix4{
		M00: (m.M11 * a11 - m.M12 * a10 + m.M13 * a09) * determinant,
		M01: (m.M02 * a10 - m.M01 * a11 - m.M03 * a09) * determinant,
		M02: (m.M31 * a05 - m.M32 * a04 + m.M33 * a03) * determinant,
		M03: (m.M22 * a04 - m.M21 * a05 - m.M23 * a03) * determinant,
		M10: (m.M12 * a08 - m.M10 * a11 - m.M13 * a07) * determinant,
		M11: (m.M00 * a11 - m.M02 * a08 + m.M03 * a07) * determinant,
		M12: (m.M32 * a02 - m.M30 * a05 - m.M33 * a01) * determinant,
		M13: (m.M20 * a05 - m.M22 * a02 + m.M23 * a01) * determinant,
		M20: (m.M10 * a10 - m.M11 * a08 + m.M13 * a06) * determinant,
		M21: (m.M01 * a08 - m.M00 * a10 - m.M03 * a06) * determinant,
		M22: (m.M30 * a04 - m.M31 * a02 + m.M33 * a00) * determinant,
		M23: (m.M21 * a02 - m.M20 * a04 - m.M23 * a00) * determinant,
		M30: (m.M11 * a07 - m.M10 * a09 - m.M12 * a06) * determinant,
		M31: (m.M00 * a09 - m.M01 * a07 + m.M02 * a06) * determinant,
		M32: (m.M31 * a01 - m.M30 * a03 - m.M32 * a00) * determinant,
		M33: (m.M20 * a03 - m.M21 * a01 + m.M22 * a00) * determinant,
	}
}

func (m Matrix4) Determinant() MnFloat {
	a00 := m.M00 * m.M11 - m.M01 * m.M10
	a01 := m.M00 * m.M12 - m.M02 * m.M10
	a02 := m.M00 * m.M13 - m.M03 * m.M10
	a03 := m.M01 * m.M12 - m.M02 * m.M11
	a04 := m.M01 * m.M13 - m.M03 * m.M11
	a05 := m.M02 * m.M13 - m.M03 * m.M12
	a06 := m.M20 * m.M31 - m.M21 * m.M30
	a07 := m.M20 * m.M32 - m.M22 * m.M30
	a08 := m.M20 * m.M33 - m.M23 * m.M30
	a09 := m.M21 * m.M32 - m.M22 * m.M31
	a10 := m.M21 * m.M33 - m.M23 * m.M31
	a11 := m.M22 * m.M33 - m.M23 * m.M32

	return a00 * a11 - a01 * a10 + a02 * a09 + a03 * a08 - a04 * a07 + a05 * a06
}

func (m Matrix4) Dump() {
	fmt.Println(fmt.Sprintf("M00 = %f, M01 = %f, M02 = %f, M03 = %f", m.M00, m.M01, m.M02, m.M03))
	fmt.Println(fmt.Sprintf("M10 = %f, M11 = %f, M12 = %f, M13 = %f", m.M10, m.M11, m.M12, m.M13))
	fmt.Println(fmt.Sprintf("M20 = %f, M21 = %f, M22 = %f, M23 = %f", m.M20, m.M21, m.M22, m.M23))
	fmt.Println(fmt.Sprintf("M30 = %f, M31 = %f, M32 = %f, M33 = %f", m.M30, m.M31, m.M32, m.M33))
}
