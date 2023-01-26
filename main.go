package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// func greet() string {
// 	fmt.println("everything is working as expected")
// }

const (
	screenWidth, screenHeight = 640, 360
)

var (
	green = color.RGBA{R: 10, G: 255, B: 50, A: 255}
)

type Game struct {
}

func (g *Game) update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func main() {
	println("oh Lord, I want a coupe")
}
