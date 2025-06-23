package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// holds game state
type Game struct {}

// for logic updates
func (g *Game) Update() error {
	return nil
}

// render grid and other elements
func (g *Game) Draw(screen *ebiten.Image) {
	DrawGrid(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
} 