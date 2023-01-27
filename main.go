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
	for _, boid := range boids {
		screen.Set(int(boid.position.x+1), int(boid.position.y), green)
		screen.Set(int(boid.position.x-1), int(boid.position.y-1), green)
	}
}

func main() {
	// if err := ebiten.RunGame(&Game{}); err != nil {
	// 	log.Fatal(err)
	// }
}
