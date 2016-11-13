package shape

import "image/color"

type BaseShape struct {
	diffuse, specular float64
	roughness         float64
	color             color.NRGBA
}

func NewBaseShape(d, s float64, r float64, c color.NRGBA) *BaseShape {
	return &BaseShape{d, s, r, c}
}

func (s *BaseShape) Diffuse() float64 {
	return s.diffuse
}

func (s *BaseShape) Specular() float64 {
	return s.specular
}

func (s *BaseShape) Roughness() float64 {
	return s.roughness
}

func (s *BaseShape) Color() color.NRGBA {
	return s.color
}
