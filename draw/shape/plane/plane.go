package plane

import (
	"image/color"

	"github.com/asib/rayce/draw"
	"github.com/asib/rayce/draw/line"
	"github.com/asib/rayce/draw/point"
	"github.com/asib/rayce/draw/shape"
	"github.com/asib/rayce/draw/vec"
)

type Plane struct {
	A *point.Point
	N *vec.Vec
	*shape.BaseShape
}

func New(a *point.Point, n *vec.Vec, dif, spec float64, rough float64, col color.NRGBA) *Plane {
	return &Plane{
		a,
		n,
		shape.NewBaseShape(dif, spec, rough, col),
	}
}

func (p *Plane) Intersect(l *line.Line) *shape.LineIntersection {
	// Used slides to calculate intersection.
	s := draw.ToVec(p.A.Sub(l.P)).Dot(p.N) / p.N.Dot(l.Direction)
	if s > 0 {
		return &shape.LineIntersection{
			p,
			l.P.Add(draw.ToPoint(l.Direction.Mul(s))),
			p.N,
			s,
		}
	}
	return nil
}
