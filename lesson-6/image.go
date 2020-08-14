package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func main() {
	green := color.RGBA{0, 255, 0, 255}
	rectImg := image.NewRGBA(image.Rect(0, 0, 200, 200))
	var red color.Color = color.RGBA{0, 0, 200, 255}

	draw.Draw(rectImg, rectImg.Bounds(), &image.Uniform{green}, image.ZP, draw.Src)
	for x := 1; x < 380; x++ {
		y := x/210 + 10
		rectImg.Set(x, y, red)
	}
	for s := 1; s < 380; s++ {
		z := s/210 + 30
		rectImg.Set(z, s, red)
	}
	file, err := os.Create("recatngle.png")
	if err != nil {
		log.Fatalf("Failed create file:%s", err)
	}
	defer file.Close()
	png.Encode(file, rectImg)
}
