package main

import (
	"image/color"
	"image/png"
	"os"

	"github.com/asib/rayce/draw/light"
	"github.com/asib/rayce/draw/point"
	"github.com/asib/rayce/draw/renderer"
	"github.com/asib/rayce/draw/scene"
	"github.com/asib/rayce/draw/shape"
	"github.com/asib/rayce/draw/shape/plane"
	"github.com/asib/rayce/draw/shape/sphere"
	"github.com/asib/rayce/draw/vec"
)

func main() {
	sc := &scene.Scene{0.2, []shape.Shape{
		sphere.New(point.New(-0.95, -0.21884, 3.63261), 0.35, 0.6, 0.2, 5, color.NRGBA{255, 0, 0, 255}),
		sphere.New(point.New(-0.4, 0.5, 4.33013), 0.7, 0.4, 0.4, 4, color.NRGBA{0, 255, 0, 255}),
		sphere.New(point.New(0.72734, -0.35322, 3.19986), 0.45, 0.5, 0.3, 3, color.NRGBA{0, 0, 255, 255}),
		plane.New(point.New(0.0, -0.10622, 4.68013), vec.New(0, 4.2239089012146, -2.180126190185547),
			0.03, 0.0, 3, color.NRGBA{150, 150, 150, 5}),
	}, []*light.PointLight{
		light.New(point.New(2, 2, 4.5), color.NRGBA{255, 255, 255, 255}, 1.0),
		light.New(point.New(-2, 2.5, 1), color.NRGBA{255, 255, 255, 255}, 2.0),
	}}
	r := &renderer.Renderer{640, 480}
	im := r.Render(sc)

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	f, err := os.Create(wd + "/test.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	png.Encode(f, im)
}
