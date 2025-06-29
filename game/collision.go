package game

import (
	"image/color"
)

func IsValidPosition(t Tetromino, grid[GridRows][GridColumns]color.RGBA) bool {
	for row := 0; row < len(t.Shape); row++ {
		for col := 0; col < len(t.Shape[row]); col++ {
			if t.Shape[row][col] == 0 {
				continue
			}

			gridRow := t.Row + row
			gridColumn := t.Column + col

			if gridRow < 0 || gridRow >= GridRows {
				return false
			}

			if gridColumn < 0 || gridColumn >= GridColumns {
				return false
			}

			if grid[gridRow][gridColumn] != (color.RGBA{}) {
				return false
			}
		}
	}

	return true
}