package camera

import (
	"math"

	"github.com/asib/rayce/draw/line"
	"github.com/asib/rayce/draw/point"
	"github.com/asib/rayce/draw/vec"
)

const fov = float64(45)

type Camera struct {
	WidthPx, HeightPx     int
	WidthWld, HeightWld   float64
	AspectRatio           float64
	WidthStep, HeightStep float64
}

func New(width, height int) *Camera {
	fwidth := float64(width)
	fheight := float64(height)
	ar := fwidth / fheight
	widthWld := 2 * math.Tan(((fov/180)*math.Pi)/2)
	heightWld := (widthWld / ar)
	return &Camera{
		width, height,
		widthWld, heightWld,
		ar,
		(widthWld / fwidth), (heightWld / fheight),
	}
}

func (c *Camera) CastRay(x, y int) []*line.Line {
	xPos := (c.WidthStep-c.WidthWld)/2 + float64(x)*c.WidthStep
	yPos := (c.HeightStep+c.HeightWld)/2 - float64(y)*c.HeightStep

	lines := []*line.Line{line.New(point.Zero(), vec.New(xPos, yPos, 1).Norm())}
	return lines

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			lines = append(lines,
				line.New(point.Zero(),
					vec.New(xPos+(float64(i)*c.WidthStep/1), yPos+(float64(j)*c.HeightStep/1), 1).Norm()))
		}
	}

	return lines
}
