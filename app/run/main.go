package main

import (
	"bspoom/app/config"
	"bspoom/app/engine"
)

func main() {
	c := config.NewConfig()
	a := engine.NewApp(c)
	a.Init()
	a.Run()
}
