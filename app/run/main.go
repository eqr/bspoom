package main

import (
	"bspoom/app/config"
	"bspoom/app/engine"
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
)

func main() {
	// Check if --profile is passed
	profileFlag := flag.Bool("profile", false, "Enable CPU profiling")
	flag.Parse()
	if *profileFlag {
		filename := "cpu.prof"
		// if exists, delete
		if _, err := os.Stat(filename); err == nil {
			if err := os.Remove(filename); err != nil {
				panic(err)
			}
		}

		// create file
		f, err := os.Create(filename)
		if err != nil {
			panic(err)
		}

		if err := pprof.StartCPUProfile(f); err != nil {
			panic(err)
		}

		defer func() {
			fmt.Println("Stopping CPU profile")
			pprof.StopCPUProfile()
		}()
	}

	c := config.NewConfig()
	a := engine.NewApp(c)
	a.Init()
	a.Run()
}
