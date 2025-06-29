package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// holds game state
type Game struct {
	tickCount int
	keyHeld bool
	moveDelay int
	moveDirection int
	softDropTimer int
	rotateCooldown int
	moveRepeatTimer int
	ActivePiece Tetromino
}

// for logic updates
func (g *Game) Update() error {
	g.handleInput()

	g.tickCount++

	if g.tickCount % 30 == 0 {
		ghost := g.ActivePiece
		ghost.Row++

		if IsValidPosition(ghost) {
			g.ActivePiece = ghost
		} else {
			// LOCK THE PIECE
		}
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

func (g *Game) handleInput() {
	left := ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA)
	right := ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD)

	if left && !right {
		g.handleMovement(-1)
	} else if !left && right {
		g.handleMovement(1)
	} else {
		g.moveDirection = 0
		g.keyHeld = false
		g.moveDelay = 0
		g.moveRepeatTimer = 0
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		g.softDropTimer++
		if g.softDropTimer >= 5 {
			ghost := g.ActivePiece
			ghost.Row++
			
			if IsValidPosition(ghost) {
				g.ActivePiece = ghost
			}

			g.softDropTimer = 0
		}
	} else {
		g.softDropTimer = 5
	}

	if ebiten.IsKeyPressed(ebiten.KeyZ) {
		if g.rotateCooldown == 0 {
			g.rotateActivePiece()
			g.rotateCooldown = 10
		}
	} else {
		g.rotateCooldown = 0
	}
}

func (g *Game) handleMovement(direction int) {
	if g.moveDirection != direction {
		g.moveDirection = direction
		g.moveDelay = 0
		g.moveRepeatTimer = 0
		g.keyHeld = true
	}

	g.moveDelay++
	if g.moveDelay >= 10 {
		g.moveRepeatTimer++
		if g.moveRepeatTimer >= 3 {
			ghost := g.ActivePiece
			ghost.Column += direction
			if IsValidPosition(ghost) {
				g.ActivePiece = ghost
			}
			g.moveRepeatTimer = 0
		}
	} else if g.moveDelay == 1 {
		ghost := g.ActivePiece
		ghost.Column += direction
		if IsValidPosition(ghost) {
			g.ActivePiece = ghost
		}
	}
}

func (g *Game) rotateActivePiece() {
	g.ActivePiece.RotateClockwise()
}