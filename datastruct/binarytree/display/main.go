package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

// Putpixel describes a function expected to draw a point on a bitmap at (x, y) coordinates.
type Putpixel func(x, y int)

// 求绝对值
func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

// Bresenham's algorithm, http://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm
// https://github.com/akavel/polyclip-go/blob/9b07bdd6e0a784f7e5d9321bff03425ab3a98beb/polyutil/draw.go
// TODO: handle int overflow etc.
func drawline(x0, y0, x1, y1 int, brush Putpixel) {
	// dx := abs(x1 - x0)
	// dy := abs(y1 - y0)
	// sx, sy := 1, 1
	// if x0 >= x1 {
	// 	sx = -1
	// }
	// if y0 >= y1 {
	// 	sy = -1
	// }
	// err := dx - dy

	for {
		brush(x0, y0)
		if x0 == x1 && y0 == y1 {
			return
		}
		// e2 := err * 2
		// if e2 > -dy {
		// 	err -= dy
		// 	x0 += sx
		// }
		// if e2 < dx {
		// 	err += dx
		// 	y0 += sy
		// }
	}
}

func draw() {
	const (
		dx = 1000
		dy = 500
	)
	img := image.NewNRGBA(image.Rect(0, 0, dx, dy))

	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			img.Set(x, y, color.RGBA{1, 1, 1, 255})
		}
	}
	for x := 0; x <= 1000; x++ {
		y := (x - 250) * (x - 250)
		img.Set(x, y, color.RGBA{255, 255, 255, 255})
	}
	imgcounter := 250
	imgfile, _ := os.Create(fmt.Sprintf("%03d.png", imgcounter))
	defer imgfile.Close()

	// 以PNG格式保存文件
	err := png.Encode(imgfile, img)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	draw()
}
