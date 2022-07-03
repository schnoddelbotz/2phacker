package main

import (
	"schnoddelbotz/mm-webpage/walker"
)

func main() {
	w := walker.NewWalker(".")
	w.RunForest()
	// now exec() hugo itself, completing work
}
