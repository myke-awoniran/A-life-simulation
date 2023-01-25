package main

import "math"

type Vector struct {
	x float64
	y float64
}

// I create receiver functions for the vectors

func (v1 *Vector) Add(v2 Vector) Vector {
	return Vector{x: v1.x + v2.x, y: v2.y + v1.y}
}

func (v1 *Vector) Subtract(v2 Vector) Vector {
	return Vector{x: v1.x - v2.x, y: v1.y - v2.y}
}

func (v1 *Vector) Multiply(v2 Vector) Vector {
	return Vector{x: v1.x * v2.x, y: v1.y * v2.y}
}

func (v1 *Vector) AddV(d float64) Vector {
	return Vector{x: v1.x + d, y: v1.y + d}
}

func (v1 Vector) MultiplyV(d float64) Vector {
	return Vector{x: v1.x * d, y: v1.y * d}
}

func (v1 *Vector) DivisionV(d float64) Vector {
	return Vector{x: v1.x / d, y: v1.y / d}
}

func (v1 *Vector) limit(lower, upper float64) Vector {
	return Vector{x: math.Min(math.Max(v1.x, lower), upper), y: math.Min(math.Max(v1.y, lower), upper)}
}

func (v1 *Vector) Distance(v2 Vector) float64 {
	return math.Sqrt(math.Pow(v1.x-v2.x, 2) + math.Pow(v1.y-v2.y, 2))
}
