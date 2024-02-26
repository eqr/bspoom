package level

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Point struct {
	X, Y float32
}

func (p Point) ToVector() rl.Vector2 {
	return rl.NewVector2(p.X, p.Y)
}

type Segment struct {
	P1, P2 Point
}

var (
	P00      = Point{1.0, 1.0}
	P01      = Point{7.0, 1.0}
	P02      = Point{7.0, 8.0}
	P03      = Point{1.0, 8.0}
	Segments = []Segment{
		{P00, P01},
		{P01, P02},
		{P02, P03},
		{P03, P00},
	}
)
