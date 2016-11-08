package draw

import (
	"github.com/asib/rayce/draw/point"
	"github.com/asib/rayce/draw/vec"
)

func ToVec(p *point.Point) *vec.Vec {
	return &vec.Vec{p.X, p.Y, p.Z, 1}
}

func ToPoint(v *vec.Vec) *point.Point {
	return &point.Point{v.X, v.Y, v.Z}
}
