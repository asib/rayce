package line

import (
	"github.com/asib/rayce/draw/point"
	"github.com/asib/rayce/draw/vec"
)

type Line struct {
	P         *point.Point
	Direction *vec.Vec
}

func New(p *point.Point, d *vec.Vec) *Line {
	return &Line{p, d}
}
