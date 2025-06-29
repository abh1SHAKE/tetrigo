package game

import (
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func DrawGrid(screen *ebiten.Image) {
	screen.Fill(color.RGBA64{0, 0, 0, 255})
	gridLineColor := color.RGBA{80, 80, 80, 255}
	gridBorderColor := color.RGBA{100, 100, 100, 255}

	gridLineWidth := float32(1)
	gridBorderWidth := float32(3)

	// internal grid
	for x := 0; x <= GridColumns; x++ {
		vector.StrokeLine(
			screen,
			float32(LeftPadding + x * BlockSize), float32(TopPadding),
			float32(LeftPadding + x * BlockSize), float32(GridHeight + TopPadding),
			gridLineWidth, gridLineColor, false)
	}

	for y := 0; y <= GridRows; y++ {
		vector.StrokeLine(
			screen,
			float32(LeftPadding), float32(TopPadding + y * BlockSize),
			float32(LeftPadding + GridWidth), float32(TopPadding + y * BlockSize),
			gridLineWidth, gridLineColor, false)
	}

	// grid border
	vector.StrokeLine(
		screen,
		float32(LeftPadding), float32(TopPadding),
		float32(LeftPadding), float32(TopPadding + GridHeight),
		gridBorderWidth, gridBorderColor, false)

	vector.StrokeLine(
		screen,
		float32(LeftPadding + GridWidth), float32(TopPadding),
		float32(LeftPadding + GridWidth), float32(TopPadding + GridHeight),
		gridBorderWidth, gridBorderColor, false)

	vector.StrokeLine(
		screen,
		float32(LeftPadding), float32(TopPadding),
		float32(LeftPadding + GridWidth), float32(TopPadding),
		gridBorderWidth, gridBorderColor, false)

	vector.StrokeLine(
		screen,
		float32(LeftPadding), float32(TopPadding + GridHeight),
		float32(LeftPadding + GridWidth), float32(TopPadding + GridHeight),
		gridBorderWidth, gridBorderColor, false)
}