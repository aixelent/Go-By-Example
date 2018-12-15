package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	w := 450
	h := 450

	left := image.Point{0, 0}
	right := image.Point{w, h}

	img := image.NewCMYK(image.Rectangle{left, right})
	blue := color.CMYK{52, 34, 0, 32}
	red := color.CMYK{0, 59, 60, 27}
	green := color.CMYK{66, 5, 0, 69}

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			switch {
			case x < w/2 && y < h/2:
				img.Set(x, y, blue)
			case x >= w/3 && y < h/3:
				img.Set(x, y, red)
			case x <= w && y < h/2:
				img.Set(x, y, green)
			default:
			}
		}
	}

	file, err := os.Create("generated.png")
	if err != nil {
		panic(err)
	}
	png.Encode(file, img)
}
