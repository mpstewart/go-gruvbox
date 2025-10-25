package gruvbox

import "image/color"

// ColorPalette exposes Gruvbox colors as image/color values.
type ColorPalette = Palette[color.Color]

// Color returns the Gruvbox palette specialized for image/color usage.
func Color() ColorPalette {
	return colorPalette
}

var colorPalette = ColorPalette{
	BG: color.RGBA{0x28, 0x28, 0x28, 0xFF},
}
