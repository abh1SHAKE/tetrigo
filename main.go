package main 

import (
	"log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/abh1SHAKE/tetrigo/game"
)

func main() {
	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowTitle("Tetris")

	err := ebiten.RunGame(&game.Game{})
	if err != nil {
		log.Fatal("Error while running game: ",err)
	}
}