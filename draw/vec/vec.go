package vec

import (
	"image/color"
	"math"
)

type Vec struct {
	X float64
	Y float64
	Z float64
	W float64
}

func New(x, y, z float64) *Vec {
	return &Vec{x, y, z, 1}
}

func FromCol(c color.NRGBA) *Vec {
	return &Vec{float64(c.R), float64(c.G), float64(c.B), 255}
}

func Zero() *Vec {
	return New(0, 0, 0)
}

func Negate(v *Vec) *Vec {
	return &Vec{-v.X, -v.Y, -v.Z, 1}
}

func Add(v1, v2 *Vec) *Vec {
	return &Vec{
		v1.X + v2.X,
		v1.Y + v2.Y,
		v1.Z + v2.Z,
		1,
	}
}

// Subtract v2 from v1.
func Sub(v1, v2 *Vec) *Vec {
	negV2 := Negate(v2)
	return Add(v1, negV2)
}

func Mul(c float64, v *Vec) *Vec {
	return &Vec{
		c * v.X,
		c * v.Y,
		c * v.Z,
		1,
	}
}

func Scale(v1, v2 *Vec) *Vec {
	return &Vec{
		v1.X * v2.X,
		v1.Y * v2.Y,
		v1.Z * v2.Z,
		1,
	}
}

func Dot(v1, v2 *Vec) float64 {
	return (v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z)
}

func Norm(v *Vec) *Vec {
	abs := v.Mod()
	return &Vec{
		v.X / abs,
		v.Y / abs,
		v.Z / abs,
		1,
	}
}

func Mod(v *Vec) float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2) + math.Pow(v.Z, 2))
}

func Reflect(v *Vec, n1 *Vec) *Vec {
	n := n1.Norm()
	return v.Sub(n.Mul(2 * v.Dot(n)))
}

func Cross(v1 *Vec, v2 *Vec) *Vec {
	return &Vec{
		v1.Y*v2.Z - v1.Z*v2.Y,
		v1.Z*v2.X - v1.X*v2.Z,
		v1.X*v2.Y - v1.Y*v2.X,
		1,
	}
}

// For convenience, provide methods on *Vec.
// None of these actually do any mutation.

func (v *Vec) Negate() *Vec {
	return Negate(v)
}

func (v1 *Vec) Add(v2 *Vec) *Vec {
	return Add(v1, v2)
}

// Subtract v2 from v1.
func (v1 *Vec) Sub(v2 *Vec) *Vec {
	return Sub(v1, v2)
}

func (v *Vec) Mul(c float64) *Vec {
	return Mul(c, v)
}

func (v1 *Vec) Scale(v2 *Vec) *Vec {
	return Scale(v1, v2)
}

func (v1 *Vec) Dot(v2 *Vec) float64 {
	return Dot(v1, v2)
}

func (v *Vec) Norm() *Vec {
	return Norm(v)
}

func (v *Vec) Mod() float64 {
	return Mod(v)
}

func (v *Vec) Reflect(n *Vec) *Vec {
	return Reflect(v, n)
}

func (v *Vec) Cross(v1 *Vec) *Vec {
	return Cross(v, v1)
}
