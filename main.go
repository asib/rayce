package main

import (
	"fmt"

	"github.com/asib/rayce/draw"
	"github.com/asib/rayce/draw/point"
	"github.com/asib/rayce/draw/shape/sphere"
	"github.com/asib/rayce/draw/vec"
)

func main() {
	zero := point.Zero()
	s := &sphere.Sphere{zero, 1}
	l := &draw.Line{zero.Sub(point.New(3, 0, 0)), vec.New(1, 0, 0)}
	fmt.Printf("%+v %+v\n", s.Intersect(l).P, s.Intersect(l).N)
}
