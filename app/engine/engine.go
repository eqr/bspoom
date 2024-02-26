package engine

import rl "github.com/gen2brain/raylib-go/raylib"

type Engine interface {
	Update()
	Draw()
}

func NewEngine() Engine {
	return &engine{}
}

type engine struct {
}

func (e *engine) Update() {
}

func (e *engine) Draw2D() {
}

func (e *engine) Draw3D() {
}

func (e *engine) Draw() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.Black)
	e.Draw2D()
	e.Draw3D()

	rl.EndDrawing()
}
