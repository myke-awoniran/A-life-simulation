package main

import (
	"math/rand"
	"time"
)

type Boid struct {
	position Vector2D
	velocity Vector2D
	id       int
}

func (boid *Boid) MoveOne() {
	boid.position = boid.position.Add(boid.velocity)
	next := boid.position.Add(boid.velocity)

	if next.x > screenWidth || next.x < 0 {
		boid.velocity = Vector2D{-boid.velocity.x, boid.velocity.y}
	}

	if next.y > screenHeight || next.y < 0 {
		boid.velocity = Vector2D{boid.velocity.x, -boid.velocity.y}
	}
}

func (boid *Boid) Start() {
	for {
		boid.MoveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func createBoid(boidId int) {
	boid := Boid{
		position: Vector2D{rand.Float64() * screenWidth, rand.Float64() * screenHeight},
		velocity: Vector2D{(rand.Float64() * 2) - 1.0, (rand.Float64() * 2) - 1.0},
		id:       boidId,
	}

	boids[boidId] = &boid
	go boid.Start()
}
