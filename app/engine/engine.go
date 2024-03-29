package engine

import (
	"bspoom/app/bsp"
	"bspoom/app/config"
	"bspoom/app/level"
	"bspoom/app/lmap"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Engine interface {
	Update()
	Draw()
}

func NewEngine(levelData level.LevelData, bspBuilder *bsp.Builder, cfg config.Config) Engine {
	//bspTree := bspBuilder.Build(levelData)
	tree := bspBuilder.Build(levelData)
	bspTraverser := bsp.NewTraverser(tree, levelData.Segments)
	return &engine{
		LevelData:   levelData,
		mapRenderer: lmap.NewRenderer(levelData, bspBuilder, bspTraverser, cfg),
		traverser:   bspTraverser,
	}
}

type engine struct {
	level.LevelData
	mapRenderer lmap.MapRenderer
	traverser   *bsp.Traverser
}

func (e *engine) Update() {
	e.traverser.Update()
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
