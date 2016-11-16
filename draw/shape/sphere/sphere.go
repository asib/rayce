package sphere

import (
	"image/color"
	"math"

	"github.com/asib/rayce/draw"
	"github.com/asib/rayce/draw/line"
	"github.com/asib/rayce/draw/point"
	"github.com/asib/rayce/draw/shape"
	"github.com/asib/rayce/draw/vec"
)

type Sphere struct {
	Centre *point.Point
	Radius float64
	*shape.BaseShape
}

func New(c *point.Point, r float64, dif, spec float64, rough float64, col color.NRGBA) *Sphere {
	return &Sphere{
		c, r,
		shape.NewBaseShape(dif, spec, rough, col),
	}
}

func (s *Sphere) NormalAt(p *point.Point) *vec.Vec {
	return draw.ToVec(p.Sub(s.Centre)).Norm()
}

func (s *Sphere) Intersect(l *line.Line) *shape.LineIntersection {
	// Used slides to calculate intersection.
	CO := draw.ToVec(l.P.Sub(s.Centre))

	a := l.Direction.Dot(l.Direction)
	b := 2 * l.Direction.Dot(CO)
	c := CO.Dot(CO) - math.Pow(s.Radius, 2)

	// This will be imaginary if there's no intersection.
	discriminant := math.Pow(b, 2) - 4*a*c
	if discriminant < 0 {
		return nil
	}
	d := math.Sqrt(discriminant)

	s1 := (d - b) / (2 * a)
	s2 := -(d + b) / (2 * a)

	if s1 > 0 && s1 < s2 {
		p := l.P.Add(draw.ToPoint(l.Direction.Mul(s1)))
		return &shape.LineIntersection{
			s,
			p,
			s.NormalAt(p),
			s1,
		}
	} else if s2 > 0 && s2 < s1 {
		p := l.P.Add(draw.ToPoint(l.Direction.Mul(s2)))
		return &shape.LineIntersection{
			s,
			p,
			s.NormalAt(p),
			s2,
		}
	}

	return nil
}
