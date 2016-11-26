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

func (r *Renderer) Trace(sc *scene.Scene, depth int, l *line.Line) color.NRGBA {
	closest := sc.ClosestIntersection(l)

	if closest == nil {
		return color.NRGBA{0, 0, 0, 255}
	}

	col := r.Illuminate(sc, depth, closest.S, closest.P, closest.N)

	// Reflection
	if depth < 5 {
		ref := l.Direction.Norm().Reflect(closest.N.Norm())
		reflect := line.New(closest.P.Add(draw.ToPoint(ref.Mul(1e-3))), ref)
		refCol := r.Trace(sc, depth+1, reflect)
		refVec := vec.FromCol(refCol).Mul(0.3)
		cVec := vec.FromCol(col).Mul(0.7)
		return color.NRGBA{uint8(math.Min(255, refVec.X+cVec.X)), uint8(math.Min(255, refVec.Y+cVec.Y)), uint8(math.Min(255, refVec.Z+cVec.Z)), 255}
	}
	return col
}

func (ren *Renderer) Illuminate(sc *scene.Scene, depth int, sh shape.Shape, p *point.Point, n *vec.Vec) color.NRGBA {
	col := sh.Color()
	r, g, b := col.R, col.G, col.B
	intensity := sc.AmbientIntensity
	specIntensity := 0.0
	for _, l := range sc.Lights {
		// fire ray between intersection point and all lights
		lightDir := draw.ToVec(l.Pos.Sub(p))
		lightRay := line.New(p.Add(draw.ToPoint(lightDir.Mul(1e-3))), lightDir)
		if ci := sc.ClosestIntersection(lightRay); ci == nil { // there aren't any objects in the way, so illuminate
			//intensity += l.IntensityAt(lightDir.Mod()) * sh.Diffuse() * math.Max(0, n.Dot(lightRay.Direction.Norm()))
			intensity += l.Intensity * sh.Diffuse() * math.Max(0, n.Dot(lightRay.Direction.Norm()))

			refl := lightRay.Direction.Reflect(n).Norm()
			v := draw.ToVec(p).Norm()
			//specIntensity += l.IntensityAt(lightDir.Mod()) * sh.Specular() * math.Pow(math.Max(0, v.Dot(refl)), float64(sh.Roughness()))
			specIntensity += l.Intensity * sh.Specular() * math.Pow(math.Max(0, v.Dot(refl)), float64(sh.Roughness()))
		}
	}

	specVal := math.Min(255, 255*specIntensity)
	rVal := float64(r) * intensity
	gVal := float64(g) * intensity
	bVal := float64(b) * intensity
	return color.NRGBA{uint8(math.Min(255, (specVal + rVal))), uint8(math.Min(255, (specVal + gVal))), uint8(math.Min(255, (specVal + bVal))), 255}
}

func (ren *Renderer) Render(sc *scene.Scene) image.Image {
	buffer := image.NewRGBA(image.Rect(0, 0, ren.Width, ren.Height))
	cam := camera.New(ren.Width, ren.Height)

	for i := 0; i < ren.Width; i++ {
		for j := 0; j < ren.Height; j++ {
			lines := cam.CastRay(i, j)
			var (
				r uint = 0
				g uint = 0
				b uint = 0
			)
			for _, l := range lines {
				col := ren.Trace(sc, 1, l)
				r += uint(col.R)
				g += uint(col.G)
				b += uint(col.B)
			}

			c := uint(len(lines))
			buffer.Set(i, j, &color.NRGBA{uint8(r / c), uint8(g / c), uint8(b / c), 255})
		}
	}

	return buffer
}
