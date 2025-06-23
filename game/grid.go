package game

import (
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func DrawGrid(screen *ebiten.Image) {
	gridColor := color.RGBA{80, 80, 80, 255}
	screen.Fill(color.RGBA64{0, 0, 0, 255})

	lineWidth := float32(1)

	for x := 0; x <= GridColumns; x++ {
		vector.StrokeLine(
			screen,
			float32(x * BlockSize), 0,
			float32(x * BlockSize), float32(ScreenHeight),
			lineWidth, gridColor, false)
	}

	for y := 0; y <= GridRows; y++ {
		vector.StrokeLine(
			screen,
			0, float32(y * BlockSize),
			float32(ScreenWidth), float32(y * BlockSize),
			lineWidth, gridColor, false)
	}
}