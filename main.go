package main

import (
	"os"
	"github.com/oyamo/rmx/src"
)

func main () {
	engine, err := src.NewEngine()
	if err != nil {
		src.PrintErr(err.Error())
		os.Exit(1)
	}

	err = engine.Run()
	if err != nil {
		src.PrintErr(err.Error())
		os.Exit(1)
	}
}