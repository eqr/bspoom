package level

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

type Point struct {
	X, Y float32
}

// TODO consider using an OpenGL-powered method for this
// Magnitude method calculates the magnitude of the vector
func (p Point) Magnitude() float32 {
	return float32(math.Sqrt(float64(p.X*p.X + p.Y*p.Y)))
}

// Normalize method normalizes the vector
func (p Point) Normalize() Point {
	magnitude := p.Magnitude()
	return Point{X: p.X / magnitude, Y: p.Y / magnitude}
}

func (p Point) ToVector() rl.Vector2 {
	return rl.NewVector2(p.X, p.Y)
}

func (p Point) Cross2D(p2 Point) float32 {
	return p.X*p2.Y - p.Y*p2.X
}

func (p Point) Minus(p2 Point) Point {
	return Point{X: p.X - p2.X, Y: p.Y - p2.Y}
}

func (p Point) Multiply(p2 float32) Point {
	return Point{X: p.X * p2, Y: p.Y * p2}
}

func (p Point) Plus(p2 Point) Point {
	return Point{X: p.X + p2.X, Y: p.Y + p2.Y}
}

func (p Point) IsOnFront(p2 Point) bool {
	return p.X*p2.Y < p.Y*p2.X
}

func (p Point) IsOnBack(p2 Point) bool {
	return !p.IsOnFront(p2)
}

type Segment struct {
	P1, P2 Point
}

func (s Segment) Vector() Point {
	return Point{X: s.P2.X - s.P1.X, Y: s.P2.Y - s.P1.Y}
}

func (s Segment) Copy() Segment {
	return Segment{P1: s.P1, P2: s.P2}
}

var (
	P00      = Point{1.0, 1.0}
	P01      = Point{7.0, 1.0}
	P02      = Point{7.0, 8.0}
	P03      = Point{1.0, 8.0}
	P04      = Point{5.0, 2.0}
	P05      = Point{4.0, 4.0}
	P06      = Point{5.0, 6.0}
	P07      = Point{6.0, 4.0}
	P08      = Point{1.8, 1.8}
	P09      = Point{1.8, 4.2}
	P10      = Point{3.5, 4.2}
	P11      = Point{3.5, 1.8}
	P12      = Point{1.5, 5.5}
	P13      = Point{1.5, 6.5}
	P14      = Point{2.0, 6.5}
	P15      = Point{2.0, 5.5}
	P16      = Point{3.25, 4.8}
	P17      = Point{2.7, 7.1}
	P18      = Point{3.8, 7.}
	Segments = []Segment{
		{P04, P05},
		{P05, P06},
		{P06, P07},
		{P07, P04},
		{P00, P01},
		{P01, P02},
		{P02, P03},
		{P03, P00},
		{P08, P09},
		{P09, P10},
		{P10, P11},
		{P11, P08},
		{P12, P13},
		{P13, P14},
		{P14, P15},
		{P15, P12},
		{P16, P17},
		{P17, P18},
		{P18, P16},
	}
	Seed = int64(26284)
)
