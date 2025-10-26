# go-gruvbox

A tiny Go module that surfaces the full Gruvbox color palette as ergonomic constants for GUIs and TUIs.

## Features
- `gruvbox.Lipgloss()` returns hex strings ready for [`lipgloss.Color`](https://github.com/charmbracelet/lipgloss).
- `gruvbox.Color()` exposes the same palette as `image/color.Color` values.
- Examples showing Bubble Tea/Lip Gloss usage and `image/draw` swatches live under `examples/`.

## Install
```bash
go get github.com/mpstewart/go-gruvbox
```

## Usage
```go
package main

import (
    "fmt"
    gruvbox "github.com/mpstewart/go-gruvbox"
)

func main() {
    gb := gruvbox.Lipgloss()
    fmt.Println("Background color:", gb.BG)
}
```

For image work:
```go
canvas := image.NewRGBA(image.Rect(0, 0, 100, 100))
draw.Draw(canvas, canvas.Bounds(), &image.Uniform{C: gruvbox.Color().BG}, image.Point{}, draw.Src)
```

## Examples
- `go run ./examples/lipgloss` — interactive bubbletea table
- `go run ./examples/color` — renders PNG swatches using `image/draw` and prints the RGB values.

## License
MIT
