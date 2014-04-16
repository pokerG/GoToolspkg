package Coordinate

import (
	"image"
	"image/color"
	// "image/png"
	"image/draw"
	// "log"
	// "os"
)

const defaultx int = 100
const defaulty int = 100

type Coordinate struct {
	// file *os.File
	Img              *image.NRGBA
	originX, originY int //original point
}

type Rgba struct {
	r uint8
	g uint8
	b uint8
	a uint8
}

func NewCoordinate() *Coordinate {
	return NewCoordinateSize(-defaultx, -defaulty, defaultx, defaulty)
}

func NewCoordinateSize(x0, y0, x1, y1 int) *Coordinate {
	c := &Coordinate{}
	c.Img = image.NewNRGBA(image.Rect(0, 0, abs(x0)+abs(x1), abs(y0)+abs(y1)))
	white := color.RGBA{255, 255, 255, 255}
	draw.Draw(c.Img, c.Img.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)
	drawline(0, abs(y1), abs(x0)+abs(x1), abs(y1), c, Rgba{0, 0, 0, 255})
	drawline(abs(x0), 0, abs(x0), abs(y0)+abs(y1), c, Rgba{0, 0, 0, 255})
	c.originX = abs(x0)
	c.originY = abs(y1)
	return c
}

func NewRgba(r, g, b, a uint8) *Rgba { //color match
	c := &Rgba{r, g, b, a}
	return c
}

func NewNode(x, y int) *Node { // coordinate node
	return &Node{x, y}
}
