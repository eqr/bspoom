package engine

import (
	"bspoom/app/config"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type App struct {
	cfg       config.Config
	deltaTime float32
	engine    Engine
}

func NewApp(cfg config.Config) *App {
	return &App{
		cfg:       cfg,
		deltaTime: 0.0,
		engine:    NewEngine(),
	}
}

func (a *App) Init() {
	rl.InitWindow(int32(a.cfg.WinWidth), int32(a.cfg.WinHeight), "BSPoom")
}

func (a *App) Run() {
	for !rl.WindowShouldClose() {
		a.deltaTime = rl.GetFrameTime()
		a.engine.Update()
		a.engine.Draw()
	}
	rl.CloseWindow()
}
