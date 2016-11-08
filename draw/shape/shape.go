package shape

import (
	"github.com/asib/rayce/draw"
	"github.com/asib/rayce/draw/point"
	"github.com/asib/rayce/draw/vec"
)

type Shape interface {
	Intersect(v *draw.Line) *LineIntersection
}

type LineIntersection struct {
	P *point.Point // point of intersection
	N *vec.Vec     // normal at that point
}
