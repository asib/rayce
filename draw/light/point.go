package light

import (
	"image/color"
	"math"

	"github.com/asib/rayce/draw/point"
)

type PointLight struct {
	Pos       *point.Point
	Color     color.NRGBA
	Intensity float64
}

func New(p *point.Point, c color.NRGBA, i float64) *PointLight {
	return &PointLight{p, c, i}
}

func (l *PointLight) IntensityAt(d float64) float64 {
	return l.Intensity / (4 * math.Pi * math.Pow(d, 2))
}
