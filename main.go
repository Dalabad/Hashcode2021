package main

import "fmt"

func main() {

	files := []string{"a", "b", "c", "d", "e", "f"}

	d := Dataset{}
	d.readInput(files[0])
	d.simulate()
	d.writeOutput()

	fmt.Printf("Final Score: %d\n", d.Score)
}

func (d *Dataset) simulate() {
	for simulationTimestamp := 0; simulationTimestamp < d.Time; simulationTimestamp++ {
		for _, car := range d.Cars {
			currentStreet := car.Path[0]
			if currentStreet.EndIntersection.isGreen(currentStreet, simulationTimestamp) {
				// Set car to next street
				if len(car.Path) > 1 {
					car.Path = car.Path[1:]
				} else {
					// Car has completed its path, remove
					car.Delete()
					d.UpdateScore(simulationTimestamp)
				}
			}
		}
	}
}