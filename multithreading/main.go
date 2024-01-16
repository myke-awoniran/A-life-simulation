package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
)

const (
	screenWidth, screenHeight = 600, 600
	boidCount                 = 500
	viewRad                   = 13
	adjRadius                 = 0.015
)

var (
	green   = color.RGBA{R: 10, G: 255, B: 50, A: 255}
	boids   [boidCount]*Boid
	boidMap [screenWidth + 1][screenHeight + 1]int
)

type Game struct {
}

func (game *Game) Update() error {
	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	for _, boid := range boids {
		screen.Set(int(boid.position.x+1), int(boid.position.y+1), green)
		screen.Set(int(boid.position.x-1), int(boid.position.y), green)
		screen.Set(int(boid.position.x), int(boid.position.y-1), green)
		screen.Set(int(boid.position.x), int(boid.position.y+1), green)

	}
}

func (game *Game) Layout(_, _ int) (width, height int) {
	return screenWidth, screenHeight
}

func main() {

	println(len(boids))
	for i := 0; i < boidCount; i++ {
		createBoid(i)
	}
	ebiten.SetWindowSize(screenWidth*2, screenHeight)
	ebiten.SetWindowTitle("Boids in a box")
	err := ebiten.RunGame(&Game{})

	if err != nil {
		log.Fatal(err)
	}
}
