package main

import (
	"fmt"
	"io/ioutil"
)

type Car struct {
	Path []Street
}

type Street struct {
	Cars              []Car
	StartIntersection Intersection
	EndIntersection   Intersection
	Name              string
	Length            int
}

type Intersection struct {
	ID       int
	Schedule Schedule
}

type Schedule struct {
	Streets  []Street
	Duration []int
}

type Dataset struct {
	Time          int
	Score         int
	Intersections []Intersection
	Streets       []Street
	Cars          []Car
}

func (d *Dataset) writeOutput() {

}

func (d *Dataset) readInput(filename string) {
	file, err := ioutil.ReadFile(fmt.Sprintf("/input/%s.txt", filename))
}