package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"reflect"

	gruvbox "github.com/mpstewart/go-gruvbox"
)

type swatch struct {
	name string
	c    color.RGBA
}

func main() {
	swatches := paletteSwatches(gruvbox.Color())
	if len(swatches) == 0 {
		fmt.Println("no colors in palette")
		return
	}

	const (
		margin       = 16
		swatchWidth  = 220
		swatchHeight = 36
		spacing      = 8
	)

	width := swatchWidth + margin*2
	height := margin*2 + len(swatches)*swatchHeight + (len(swatches)-1)*spacing
	canvas := image.NewRGBA(image.Rect(0, 0, width, height))

	for i, sw := range swatches {
		y := margin + i*(swatchHeight+spacing)
		rect := image.Rect(margin, y, margin+swatchWidth, y+swatchHeight)
		draw.Draw(canvas, rect, &image.Uniform{C: sw.c}, image.Point{}, draw.Src)
	}

	const outPath = "gruvbox_palette.png"
	if err := writePNG(outPath, canvas); err != nil {
		fmt.Fprintf(os.Stderr, "write PNG: %v\n", err)
		os.Exit(1)
	}

	for _, sw := range swatches {
		fmt.Printf("%-12s #%02x%02x%02x\n", sw.name, sw.c.R, sw.c.G, sw.c.B)
	}
	fmt.Printf("wrote %s with %d swatches\n", outPath, len(swatches))
}

func paletteSwatches(p gruvbox.ColorPalette) []swatch {
	val := reflect.ValueOf(p)
	typ := val.Type()

	swatches := make([]swatch, 0, typ.NumField())
	for i := 0; i < typ.NumField(); i++ {
		field := val.Field(i)
		col := color.RGBAModel.Convert(field.Interface().(color.Color)).(color.RGBA)
		swatches = append(swatches, swatch{
			name: typ.Field(i).Name,
			c:    col,
		})
	}
	return swatches
}

func writePNG(path string, img image.Image) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		return err
	}
	return nil
}
