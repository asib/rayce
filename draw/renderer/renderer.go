package renderer

import (
	"image"
	"image/color"
	"math"

	"github.com/asib/rayce/draw"
	"github.com/asib/rayce/draw/camera"
	"github.com/asib/rayce/draw/line"
	"github.com/asib/rayce/draw/point"
	"github.com/asib/rayce/draw/scene"
	"github.com/asib/rayce/draw/shape"
	"github.com/asib/rayce/draw/vec"
)

type Renderer struct {
	Width  int
	Height int
}

func (r *Renderer) Trace(sc *scene.Scene, l *line.Line) color.Color {
	closest := sc.ClosestIntersection(l)

	if closest == nil {
		return color.NRGBA{255, 255, 255, 0}
	}

	return r.Illuminate(sc, closest.S, closest.P, closest.N)
}

func (ren *Renderer) Illuminate(sc *scene.Scene, sh shape.Shape, p *point.Point, n *vec.Vec) color.Color {
	col := sh.Color()
	r, g, b := col.R, col.G, col.B
	intensity := sc.AmbientIntensity
	specIntensity := 0.0
	for _, l := range sc.Lights {
		// fire ray between intersection point and all lights
		lightRay := line.New(p, draw.ToVec(l.Pos.Sub(p)))
		if sc.ClosestIntersection(lightRay) == nil { // there aren't any objects in the way, so illuminate
			intensity += l.Intensity * sh.Diffuse() * math.Max(0, n.Dot(lightRay.Direction.Norm()))

			refl := lightRay.Direction.Reflect(n).Norm()
			v := draw.ToVec(p).Norm()
			specIntensity += l.Intensity * sh.Specular() * math.Pow(math.Max(0, v.Dot(refl)), float64(sh.Roughness()))
		}
	}

	specVal := math.Min(255, 255*specIntensity)
	rVal := float64(r) * intensity
	gVal := float64(g) * intensity
	bVal := float64(b) * intensity
	return color.NRGBA{uint8(math.Min(255, (specVal+rVal)/2)), uint8(math.Min(255, (specVal+gVal)/2)), uint8(math.Min(255, (specVal+bVal)/2)), 255}
}

func (r *Renderer) Render(sc *scene.Scene) image.Image {
	buffer := image.NewRGBA(image.Rect(0, 0, r.Width, r.Height))
	cam := camera.New(r.Width, r.Height)

	for i := 0; i < r.Width; i++ {
		for j := 0; j < r.Height; j++ {
			l := cam.CastRay(i, j)
			col := r.Trace(sc, l)
			buffer.Set(i, j, col)
		}
	}

	return buffer
}
