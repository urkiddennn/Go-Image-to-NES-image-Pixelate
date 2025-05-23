package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"math"
	"os"
)

var NesPalette = []color.RGBA{
	{0x7C, 0x7C, 0x7C, 0xFF},
	{0x00, 0x00, 0xFC, 0xFF},
	{0x00, 0x00, 0xBC, 0xFF},
	{0x44, 0x28, 0xBC, 0xFF},
	{0x94, 0x00, 0x84, 0xFF},
	{0xA8, 0x00, 0x20, 0xFF},
	{0xA8, 0x10, 0x00, 0xFF},
	{0x88, 0x14, 0x00, 0xFF},
	{0x50, 0x30, 0x00, 0xFF},
	{0x00, 0x78, 0x00, 0xFF},
	{0x00, 0x68, 0x00, 0xFF},
	{0x00, 0x58, 0x00, 0xFF},
	{0x00, 0x40, 0x58, 0xFF},
	{0x00, 0x00, 0x00, 0xFF},
	{0x00, 0x00, 0x00, 0xFF},
	{0x00, 0x00, 0x00, 0xFF},
}

// Find the colosest color to the NES palette
func closestColor(c color.Color) color.Color {
	r, g, b, _ := c.RGBA()
	minDistance := math.MaxFloat64
	var closest color.RGBA

	for _, paletteColor := range NesPalette {
		pr, pg, pb := float64(paletteColor.R), float64(paletteColor.G), float64(paletteColor.B)

		distance := math.Sqrt(math.Pow(float64(r>>8)-pr, 2) + math.Pow(float64(g>>8)-pg, 2) + math.Pow(float64(b>>8)-pb, 2))
		if distance < minDistance {
			minDistance = distance
			closest = paletteColor
		}
	}

	return closest
}

// Apply the NesPalette
func applyNESColorPalette(img image.Image) image.Image {
	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalColor := img.At(x, y)
			newImg.Set(x, y, closestColor(originalColor))
		}
	}
	return newImg
}

func main() {
	// Define a Custom NES-like color palette
	// get the image file
	imgFile, err := os.Open("reck.jpg")
	if err != nil {
		panic(err)
	}
	defer imgFile.Close()
	// Decode the image file
	img, _, err := image.Decode(imgFile)
	if err != nil {
		panic(err)
	}

	nesImg := applyNESColorPalette(img)

	ouputFile, err := os.Create("output2.jpg")
	if err != nil {
		panic(err)
	}
	defer ouputFile.Close()

	jpegOptions := &jpeg.Options{Quality: 2}

	err = jpeg.Encode(ouputFile, nesImg, jpegOptions)
	if err != nil {
		panic(err)
	}

	fmt.Println("Image coverted to NES palette Successfully!")
}
