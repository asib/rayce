package scene

import (
	"github.com/asib/rayce/draw/light"
	"github.com/asib/rayce/draw/line"
	"github.com/asib/rayce/draw/shape"
)

type Scene struct {
	AmbientIntensity float64
	Objects          []shape.Shape
	Lights           []*light.PointLight
}

// Assume camera fixed at origin.
func (s *Scene) ClosestIntersection(l *line.Line) *shape.LineIntersection {
	var closest *shape.LineIntersection
	for _, o := range s.Objects {
		if intersection := o.Intersect(l); intersection != nil && closest != nil &&
			intersection.P.Distance() < closest.P.Distance() {
			closest = intersection
		} else if closest == nil {
			closest = intersection
		}
	}

	return closest
}
