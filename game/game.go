package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// holds game state
type Game struct {
	tickCount int
	ActivePiece Tetromino
}

// for logic updates
func (g *Game) Update() error {
	g.tickCount++

	if g.tickCount % 30 == 0 {
		g.ActivePiece.Row++
	}

	return  nil
}

// render grid and other elements
func (g *Game) Draw(screen *ebiten.Image) {
	DrawGrid(screen)
	g.ActivePiece.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func NewGame() *Game {
	return &Game{
		ActivePiece: NewTetromino(),
	}
}