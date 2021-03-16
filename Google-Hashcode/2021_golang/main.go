package main

import (
	"fmt"
	"hashcode/src"
)

func main() {

	files := []string{"a", "b", "c", "d", "e", "f"}

	for _, file := range files {
		d := src.Dataset{}
		d.ReadInput(file)
		d.SetSchedules()
		//d.Simulate()
		d.WriteOutput(file)

		fmt.Printf("Input: %s - Final Score: %d\n", file, d.Score)
	}
}