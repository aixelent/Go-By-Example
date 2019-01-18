package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
)

func main() {
	red := color.RGBA{250, 128, 114, 1}
	green := color.RGBA{144, 238, 144, 1}

	r := image.Rect(0, 0, 400, 400)
	i := image.NewNRGBA(r)

	draw.Draw(i, r, &image.Uniform{green}, image.ZP, draw.Src)

	i.Set(25, 50, red)
	i.Set(50, 75, red)
	i.Set(75, 100, color.Black)
	i.Set(100, 125, color.White)

	i1 := i.At(25, 50)
	i2 := i.At(50, 75)
	i3 := i.At(75, 100)
	i4 := i.At(100, 125)
	i5 := i.At(0, 0)

	var buffer bytes.Buffer
	err := png.Encode(&buffer, i)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Image size:", i.Bounds().Dx(), "x", i.Bounds().Dy())
	fmt.Printf("Pixel %v has the %v color.\n", image.Pt(25, 50), i1)
	fmt.Printf("Pixel %v has the %v color.\n", image.Pt(50, 75), i2)
	fmt.Printf("Pixel %v has the %v color.\n", image.Pt(75, 100), i3)
	fmt.Printf("Pixel %v has the %v color.\n", image.Pt(100, 125), i4)
	fmt.Printf("Pixel %v has the %v color.\n", image.Pt(0, 0), i5)
}
