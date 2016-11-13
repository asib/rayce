package main

import (
	"image/color"
	"image/jpeg"
	"os"

	"github.com/asib/rayce/draw/light"
	"github.com/asib/rayce/draw/point"
	"github.com/asib/rayce/draw/renderer"
	"github.com/asib/rayce/draw/scene"
	"github.com/asib/rayce/draw/shape"
	"github.com/asib/rayce/draw/shape/sphere"
)

func main() {
	sc := &scene.Scene{0.3, []shape.Shape{
		sphere.New(point.New(0.55, 0.46, 3.5), 0.5, 0.8, 0.8, 5, color.NRGBA{255, 0, 0, 255}),
		sphere.New(point.New(-0.55, 0, 5), 0.9, 0.9, 1.0, 40, color.NRGBA{0, 255, 0, 255}),
		sphere.New(point.New(-0.1, 0.8, 3), 0.1, 0.9, 1.0, 30, color.NRGBA{0, 0, 255, 255}),
	}, []*light.PointLight{
		light.New(point.New(5, 8, -5), color.NRGBA{255, 255, 255, 255}, 2.0),
		light.New(point.New(-5, -3, -5), color.NRGBA{255, 255, 255, 255}, 0.9),
	}}
	r := &renderer.Renderer{640, 480}
	im := r.Render(sc)

	f, err := os.Create("/Users/jacobfenton/bla1.jpeg")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	jpeg.Encode(f, im, nil)
}
