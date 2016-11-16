package shape

import (
	"image/color"

	"github.com/asib/rayce/draw/line"
	"github.com/asib/rayce/draw/point"
	"github.com/asib/rayce/draw/vec"
)

type Shape interface {
	Intersect(v *line.Line) *LineIntersection
	Diffuse() float64
	Specular() float64
	Roughness() float64
	Color() color.NRGBA
}

type LineIntersection struct {
	S    Shape
	P    *point.Point // point of intersection
	N    *vec.Vec     // normal at that point
	Dist float64      // what we had to multiply the direction vector by - if negative, know the thing is behind us
}
