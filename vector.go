package main

import "math"

type Vector struct {
	X float64
	Y float64
}

func (v *Vector) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vector) Normalize() Vector {
	return Vector{
		X: v.X / v.Magnitude(),
		Y: v.Y / v.Magnitude(),
	}
}
