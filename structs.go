package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
	f, err := os.Create("output/a.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	for _, intersection := range d.Intersections {
		// ID of the intersection
		_, err = f.WriteString(fmt.Sprintf("%d\n", intersection.ID))
		if err != nil {
			panic(err)
		}
		// Number of incoming streets
		_, err = f.WriteString(fmt.Sprintf("%d\n", len(intersection.Schedule.Streets)))
		if err != nil {
			panic(err)
		}
		// Schedule with the format: "street_name duration"
		for i, street := range intersection.Schedule.Streets {
			_, err = f.WriteString(fmt.Sprintf("%s %d\n", street.Name, intersection.Schedule.Duration[i]))
			if err != nil {
				log.Fatal(err)
			}
		}

	}


}

func (d *Dataset) readInput(filename string) {
	file, err := ioutil.ReadFile(fmt.Sprintf("/input/%s.txt", filename))
}
