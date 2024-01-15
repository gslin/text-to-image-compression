package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	data, err := os.ReadFile("pg1513.txt")
	if err != nil {
		panic(err)
	}
	datasize := len(data)
	fmt.Printf("File size: %d\n", datasize)

	img := image.NewGray(image.Rect(0, 0, datasize, 1))
	for i, p := range(data) {
		img.SetGray(i, 0, color.Gray{p})
	}

	f, _ := os.Create("pg1513.png")

	enc := &png.Encoder{CompressionLevel: png.BestCompression}
	if enc.Encode(f, img); err != nil {
		panic(err)
	}

	f.Close()
}
