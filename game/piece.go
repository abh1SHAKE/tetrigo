package game

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Tetromino struct {
	Shape [][]int
	Color color.RGBA
	Row int
	Column int
}

func NewTetromino() Tetromino {
	shapes := [][][]int {
		{
			{0, 0, 0, 0},
			{1, 1, 1, 1},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		},
		{
			{0, 1, 1, 0},
			{0, 1, 1, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		},
		{
			{0, 1, 0, 0},
			{1, 1, 1, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		},
		{
			{0, 1, 1, 0},
			{1, 1, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		},
		{
			{1, 1, 0, 0},
			{0, 1, 1, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		},
		{
			{1, 0, 0, 0},
			{1, 1, 1, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		},
		{
			{0, 0, 1, 0},
			{1, 1, 1, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		},
	}

	colors := []color.RGBA {
		{0, 255, 255, 255}, 
		{255, 255, 0, 255},
		{128, 0, 128, 255},   
		{0, 255, 0, 255},   
		{255, 0, 0, 255},     
		{0, 0, 255, 255},    
		{255, 165, 0, 255}, 
	}

	index := rand.Intn(len(shapes))

	return Tetromino{
		Shape: shapes[index],
		Color: colors[index],
		Row: 0,
		Column: 3,
	}
}

func (t *Tetromino) Draw(screen *ebiten.Image) {
	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			if (t.Shape[row][col] == 1) {
				x := LeftPadding + (t.Column + col) * BlockSize
				y := TopPadding + (t.Row + row) * BlockSize

				border := ebiten.NewImage(BlockSize, BlockSize)
				border.Fill(color.RGBA{
					uint8(t.Color.R / 2),
					uint8(t.Color.G / 2),
					uint8(t.Color.B / 2),
					255,
				})

				geom := ebiten.GeoM{}
				geom.Translate(float64(x), float64(y))
				screen.DrawImage(border, &ebiten.DrawImageOptions{GeoM: geom})

				padding := 3
				inner := ebiten.NewImage(BlockSize - padding * 2, BlockSize - padding * 2)
				inner.Fill(t.Color)

				innerGeom := ebiten.GeoM{}
				innerGeom.Translate(float64(x + padding), float64(y + padding))
				screen.DrawImage(inner, &ebiten.DrawImageOptions{GeoM: innerGeom})
			}
		}
	}
}

func (t *Tetromino) RotateClockwise() {
	len := len(t.Shape)
	newShape := make([][]int, len)

	for i := 0; i < len; i++ {
		newShape[i] = make([]int, len)
		for j := 0; j < len; j++ {
			newShape[i][j] = t.Shape[len-j-1][i]
		}
	}

	t.Shape = newShape
}