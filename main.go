package main

import (
	"github.com/oyamo/rmx/src"
)

func main () {
	engine, err := src.NewEngine()
	if err != nil {
		panic(err)
	}
	
	engine.Run()
}