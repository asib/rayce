package point

import "math"

type Point struct {
	X float64
	Y float64
	Z float64
}

func New(x, y, z float64) *Point {
	return &Point{x, y, z}
}

func Zero() *Point {
	return New(0, 0, 0)
}

func Negate(p *Point) *Point {
	return &Point{-p.X, -p.Y, -p.Z}
}

func Add(p1, p2 *Point) *Point {
	return &Point{
		p1.X + p2.X,
		p1.Y + p2.Y,
		p1.Z + p2.Z,
	}
}

func Sub(p1, p2 *Point) *Point {
	return Add(p1, Negate(p2))
}

func Distance(p *Point) float64 {
	return math.Sqrt(math.Pow(p.X, 2) + math.Pow(p.Y, 2) + math.Pow(p.Z, 2))
}

// For convenience.

func (p1 *Point) Add(p2 *Point) *Point {
	return Add(p1, p2)
}

func (p *Point) Negate() *Point {
	return Negate(p)
}

func (p1 *Point) Sub(p2 *Point) *Point {
	return Sub(p1, p2)
}

func (p *Point) Distance() float64 {
	return Distance(p)
}
