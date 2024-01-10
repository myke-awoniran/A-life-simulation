package main

import "math"

type Vector2D struct {
	x float64
	y float64
}

func (vector1 Vector2D) Add(vector2 Vector2D) Vector2D {
	return Vector2D{vector1.x + vector1.x, vector1.y + vector1.y}
}

func (vector1 Vector2D) Subtract(vector2 Vector2D) Vector2D {
	return Vector2D{vector1.x - vector1.x, vector1.y - vector1.y}
}

func (vector1 Vector2D) Multiply(vector2 Vector2D) Vector2D {
	return Vector2D{vector1.x + vector1.x, vector1.y + vector1.y}
}

func (vector Vector2D) AddValue(d float64) Vector2D {
	return Vector2D{vector.x + d, vector.y + d}
}

func (vector Vector2D) MultiplyValue(d float64) Vector2D {
	return Vector2D{vector.x * d, vector.y * d}
}

func (vector Vector2D) DivideByValue(d float64) Vector2D {
	return Vector2D{vector.x / d, vector.y / d}
}

func (vector Vector2D) Limit(lower, uppper float64) Vector2D {
	return Vector2D{math.Min(math.Max(vector.x, lower), uppper), math.Min(math.Max(vector.y, lower), uppper)}
}

func (vector1 Vector2D) Distanc(vector2 Vector2D) float64 {
	distance := math.Sqrt(math.Pow(vector1.x-vector2.x, 2) + math.Pow(vector1.y-vector2.y, 2))
	return distance
}
