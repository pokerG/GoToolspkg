package Coordinate

import (
	"fmt"
)

type Node struct {
	x int
	y int
}

func (this *Coordinate) FoldLine(n []Node, r Rgba) {
	l := len(n)
	for i := 0; i < l-1; i++ {
		x0 := this.originX + n[i].x
		y0 := this.originY - n[i].y
		x1 := this.originX + n[i+1].x
		y1 := this.originY - n[i+1].y
		fmt.Println(x0, y0, x1, y1)
		drawline(x0, y0, x1, y1, this, r)
	}

}
