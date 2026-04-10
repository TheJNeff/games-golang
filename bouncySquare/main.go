package main

import (
	"log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"strconv"
)

const (
	windowWidth = 640
	windowHeight = 480
	windowTitle = "Bouncy Square!"
)

const (
	screenWidth = 320
	screenHeight = 240
)

const (
	squareSize = 20
)

var (
	squareImage *ebiten.Image = ebiten.NewImage(squareSize, squareSize)
	squareColor color.Color = color.RGBA{255, 0, 0, 255}
)

type Game struct {
	redSquare *ebiten.Image
	redSquareX int
	redSquareY int
	redSquareXDir int
	redSquareYDir int
}

func (g *Game) Update() error {
	// move the red square 1 pixel per frame in the current direction
	g.redSquareX = g.redSquareX + (1 * g.redSquareXDir)
	g.redSquareY = g.redSquareY + (1 * g.redSquareYDir)

	// if the square collides with the side of the screen, reverse the horizontal direction
	if (g.redSquareX + 20 > 320 || g.redSquareX <= 0) {
		g.redSquareXDir = g.redSquareXDir * -1;
	}
	// if the square collides with the top or bottom of the screen, reverse the vertical direction
	if (g.redSquareY + 20 > 240 || g.redSquareY <= 0) {
		g.redSquareYDir = g.redSquareYDir * -1;
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// print the location of the square for debugging 
	ebitenutil.DebugPrint(screen, "X: " + strconv.Itoa(g.redSquareX) + " Y: " + strconv.Itoa(g.redSquareY))
	
	// put the square in the appropriate location before drawing it
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.redSquareX), float64(g.redSquareY))

	// draw the square on the screen
	screen.DrawImage(g.redSquare, op)
	return
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle(windowTitle)
	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func NewGame() *Game {
  	squareImage.Fill(squareColor)
	return &Game{ 
		redSquare: squareImage, 
		redSquareX: 0, 
		redSquareY: 0, 
		redSquareXDir: 1, 
		redSquareYDir: 1}
}
