package Coordinate

import (
	"image/color"
)

// Bresenham's algorithm, http://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm
// https://github.com/akavel/polyclip-go/blob/9b07bdd6e0a784f7e5d9321bff03425ab3a98beb/polyutil/draw.go
// TODO: handle int overflow etc.
func Bresenham(x0, y0, x1, y1 int, c *Coordinate, cm Rgba) {
	dx := abs(x1 - x0)
	dy := abs(y1 - y0)
	sx, sy := 1, 1
	if x0 >= x1 {
		sx = -1
	}
	if y0 >= y1 {
		sy = -1
	}
	err := dx - dy

	for {
		brush(x0, y0, c, cm)
		if x0 == x1 && y0 == y1 {
			return
		}
		e2 := err * 2
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}
}

func drawline(x0, y0, x1, y1 int, c *Coordinate, cm Rgba) { //bold line
	Bresenham(x0, y0, x1, y1, c, cm)
	Bresenham(x0+1, y0+1, x1+1, y1+1, c, cm)
	Bresenham(x0+2, y0+2, x1+2, y1+2, c, cm)
	Bresenham(x0-1, y0-1, x1-1, y1-1, c, cm)
	Bresenham(x0-2, y0-2, x1-2, y1-2, c, cm)
}

// 求绝对值
func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func brush(x, y int, c *Coordinate, cm Rgba) {
	c.Img.Set(x, y, color.RGBA{cm.r, cm.g, cm.b, cm.a})
}
