package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	// Define a Custom NES-like color palette
	palette := color.Palette{
		color.RGBA{0x7C, 0x7C, 0x7C, 0xFF}, // Light gray
		color.RGBA{0x00, 0x00, 0xFC, 0xFF}, // Blue
		color.RGBA{0x00, 0x00, 0xBC, 0xFF}, // Dark blue
		color.RGBA{0x44, 0x28, 0xBC, 0xFF}, // Purple
		color.RGBA{0x94, 0x00, 0x84, 0xFF}, // Magenta
		color.RGBA{0xA8, 0x00, 0x20, 0xFF}, // Red
		color.RGBA{0xA8, 0x10, 0x00, 0xFF}, // Dark red
		color.RGBA{0x88, 0x14, 0x00, 0xFF}, // Brown
		color.RGBA{0x50, 0x30, 0x00, 0xFF}, // Dark brown
		color.RGBA{0x00, 0x78, 0x00, 0xFF}, // Green
		color.RGBA{0x00, 0x68, 0x00, 0xFF}, // Dark green
		color.RGBA{0x00, 0x58, 0x00, 0xFF}, // Olive
		color.RGBA{0x00, 0x40, 0x58, 0xFF}, // Teal
		color.RGBA{0x00, 0x00, 0x00, 0xFF},
	}
	img := image.NewPaletted(image.Rect(0, 0, 100, 100), palette)

	for y := 0; y < 100; y++ {
		for x := 0; x < 100; x++ {
			colorIndex := uint8((x + y) % len(palette))
			img.SetColorIndex(x, y, colorIndex)
		}
	}
	// handle Errors
	file, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}

	fmt.Println("Created the output")
}
