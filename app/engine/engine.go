package engine

import (
	"bspoom/app/config"
	"bspoom/app/level"
	"bspoom/app/lmap"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Engine interface {
	Update()
	Draw()
}

func NewEngine(levelData level.LevelData, cfg config.Config) Engine {
	return &engine{
		LevelData:   levelData,
		mapRenderer: lmap.NewRenderer(levelData, cfg),
	}
}

type engine struct {
	level.LevelData
	mapRenderer lmap.MapRenderer
}

func (e *engine) Update() {
}

func (e *engine) Draw2D() {
	e.mapRenderer.Draw()
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
