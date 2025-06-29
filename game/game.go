package game

import (
	"fmt"
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// holds game state
type Game struct {
	tickCount int
	ActivePiece Tetromino
	Grid [GridRows][GridColumns]color.RGBA

	// movement-related
	keyHeld bool
	moveDelay int
	moveDirection int
	softDropTimer int
	moveRepeatTimer int

	// rotation debounce
	rotateCooldown int

	// locking behaviour
	lockDelayMax int
	lockDelayCounter int
}

// for logic updates
func (g *Game) Update() error {
	g.handleInput()

	g.tickCount++

	if g.tickCount % 30 == 0 {
		ghost := g.ActivePiece
		ghost.Row++

		if IsValidPosition(ghost, g.Grid) {
			g.ActivePiece = ghost
			g.lockDelayCounter = 0
		}
	}

	ghost := g.ActivePiece
	ghost.Row++

	if !IsValidPosition(ghost, g.Grid) {
		g.lockDelayCounter++
		if g.lockDelayCounter >= g.lockDelayMax {
			g.lockPiece()
		}
	} else {
		g.lockDelayCounter = 0
	}

	return  nil
}

// render grid and other elements
func (g *Game) Draw(screen *ebiten.Image) {
	DrawGrid(screen)

	for row := 0; row < GridRows; row++ {
		for col := 0; col < GridColumns; col++ {
			cellColor := g.Grid[row][col]
			if cellColor == (color.RGBA{}) {
				continue
			}

			x := LeftPadding + col * BlockSize
			y := TopPadding + row * BlockSize

			border := ebiten.NewImage(BlockSize, BlockSize)
			border.Fill(color.RGBA{
				uint8(cellColor.R / 2),
				uint8(cellColor.G / 2),
				uint8(cellColor.B / 2),
				255,
			})

			geom := ebiten.GeoM{}
			geom.Translate(float64(x), float64(y))
			screen.DrawImage(border, &ebiten.DrawImageOptions{GeoM:  geom})

			padding := 3
			inner := ebiten.NewImage(BlockSize - padding * 2, BlockSize - padding * 2)
			inner.Fill(cellColor)
			innerGeom := ebiten.GeoM{}
			innerGeom.Translate(float64(x + padding), float64(y + padding))
			screen.DrawImage(inner, &ebiten.DrawImageOptions{GeoM: innerGeom})
		}
	}

	g.ActivePiece.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func NewGame() *Game {
	return &Game{
		ActivePiece: NewTetromino(),
		lockDelayMax: 30,
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

			if IsValidPosition(ghost, g.Grid) {
				g.ActivePiece = ghost
				g.lockDelayCounter = 0
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

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.hardDrop()
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
			if IsValidPosition(ghost, g.Grid) {
				g.ActivePiece = ghost
				g.lockDelayCounter = 0
			}
			g.moveRepeatTimer = 0
		}
	} else if g.moveDelay == 1 {
		ghost := g.ActivePiece
		ghost.Column += direction
		if IsValidPosition(ghost, g.Grid) {
			g.ActivePiece = ghost
			g.lockDelayCounter = 0
		}
	}
}

func (g *Game) rotateActivePiece() {
	rotated := g.ActivePiece
	rotated.RotateClockwise()

	if IsValidPosition(rotated, g.Grid) {
		g.ActivePiece = rotated
		g.lockDelayCounter = 0
	}
}

func (g *Game) hardDrop() {
	ghost := g.ActivePiece

	for {
		next := ghost
		next.Row++

		if IsValidPosition(next, g.Grid) {
			ghost = next
		} else {
			break
		}
	}

	g.ActivePiece = ghost
	g.lockPiece()
}

func (g *Game) lockPiece() {
	for row := 0; row < len(g.ActivePiece.Shape); row++ {
		for col := 0; col < len(g.ActivePiece.Shape[row]); col++ {
			if g.ActivePiece.Shape[row][col] == 0 {
				continue
			}

			gridRow := g.ActivePiece.Row + row
			gridColumn := g.ActivePiece.Column + col
			
			if gridRow >= 0 && gridRow < GridRows && gridColumn >= 0 && gridColumn < GridColumns {
				g.Grid[gridRow][gridColumn] = g.ActivePiece.Color
			}
		}
	}

	g.lockDelayCounter = 0

	g.ActivePiece = NewTetromino()

	if !IsValidPosition(g.ActivePiece, g.Grid) {
		fmt.Println("Game over")
	}
}