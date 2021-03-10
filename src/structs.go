package src

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Car struct {
	Path                    []Street
	DurationOnCurrentStreet int
}

func (c *Car) Delete() {
	// TODO: Remove car
}

type Street struct {
	Cars              []*Car
	StartIntersection *Intersection
	EndIntersection   *Intersection
	Name              string
	Length            int
}

type Intersection struct {
	ID       int
	Schedule *Schedule
}

func (i *Intersection) isGreen(street Street, timestamp int) bool {
	if len(i.Schedule.Duration) == 0 {
		return false
	}

	overallDuration := 0
	for _, v := range i.Schedule.Duration {
		overallDuration += v
	}

	remaining := timestamp % overallDuration

	for streetIndex, duration := range i.Schedule.Duration {
		if remaining > duration {
			remaining -= duration
		} else {
			return i.Schedule.Streets[streetIndex].Name == street.Name
		}
	}

	return false
}

type Schedule struct {
	Streets  []*Street
	Duration []int
}

type Dataset struct {
	Time          int
	Score         int
	Bonus         int
	Intersections []*Intersection
	Streets       []Street
	Cars          []Car
}

func (d *Dataset) WriteOutput(filename string) {
	f, err := os.Create(fmt.Sprintf("output/%s.out", filename))
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

func (d *Dataset) ReadInput(filename string) {
	file, err := os.Open(fmt.Sprintf("input/%s.txt", filename))

	if err != nil {
		log.Fatalln("file not found")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	paramsStrings := strings.Split(scanner.Text(), " ")
	params := make([]int, 0)
	for _, paramString := range paramsStrings {
		param, _ := strconv.Atoi(paramString)
		params = append(params, param)
	}

	time, intersectionsCount, streetsCount, carsCount, bonusPoints := params[0], params[1], params[2], params[3], params[4]
	d.Bonus = time
	d.Score = bonusPoints

	for i := 0; i < streetsCount; i++ {
		scanner.Scan()
		streetData := strings.Split(scanner.Text(), " ")
		startIntersectionId, _ := strconv.Atoi(streetData[0])
		endIntersectionId, _ := strconv.Atoi(streetData[1])
		name := streetData[2]
		duration, _ := strconv.Atoi(streetData[3])

		if !ContainsIntersection(d.Intersections, startIntersectionId) {
			startIntersection := Intersection{
				ID: startIntersectionId,
				Schedule: &Schedule{
					Streets:  make([]*Street, 0),
					Duration: make([]int, 0),
				},
			}
			d.Intersections = append(d.Intersections, &startIntersection)
		}
		if !ContainsIntersection(d.Intersections, endIntersectionId) {
			endIntersection := Intersection{
				ID: endIntersectionId,
				Schedule: &Schedule{
					Streets:  make([]*Street, 0),
					Duration: make([]int, 0),
				},
			}
			d.Intersections = append(d.Intersections, &endIntersection)
		}

		endIntersection := d.FindIntersectionById(endIntersectionId)

		street := Street{
			Cars:              make([]*Car, 0),
			StartIntersection: d.FindIntersectionById(startIntersectionId),
			EndIntersection:   endIntersection,
			Name:              name,
			Length:            duration,
		}
		endIntersection.Schedule.Streets = append(endIntersection.Schedule.Streets, &street)

		d.Streets = append(d.Streets, street)
	}
	for i := 0; i < carsCount; i++ {
		scanner.Scan()
		carData := strings.Split(scanner.Text(), " ")
		pathLength, _ := strconv.Atoi(carData[0])
		car := Car{Path: make([]Street, pathLength), DurationOnCurrentStreet: 0}

		for i := 1; i <= pathLength; i++ {
			car.Path = append(car.Path, d.FindStreetByName(carData[i]))
		}
		d.Cars = append(d.Cars, car)
	}

	if len(d.Intersections) != intersectionsCount {
		panic("intersections count does not add up")
	}
}

func (d *Dataset) FindIntersectionById(id int) *Intersection {
	for _, i := range d.Intersections {
		if i.ID == id {
			return i
		}
	}
	panic(fmt.Sprintf("intersection %d not found", id))
}

func (d *Dataset) FindStreetByName(name string) Street {
	for _, street := range d.Streets {
		if street.Name == name {
			return street
		}
	}
	panic(fmt.Sprintf("street %s not found", name))
}

func ContainsIntersection(intersections []*Intersection, id int) bool {
	for _, intersection := range intersections {
		if id == intersection.ID {
			return true
		}
	}
	return false
}

func (d *Dataset) UpdateScore(timestamp int) {
	addScore := 1000 + (d.Time - timestamp)
	d.Score += addScore
	d.Score += d.Bonus
	fmt.Printf("Increase Score by %d\n", addScore)
}

func (d *Dataset) Simulate() {
	for simulationTimestamp := 0; simulationTimestamp < d.Time; simulationTimestamp++ {
		for _, street := range d.Streets {
			if len(street.Cars) <= 0 {
				continue
			}

			car := street.Cars[0]

			// Decrease duration of car on current street by 1
			if car.DurationOnCurrentStreet > 0 {
				car.DurationOnCurrentStreet--
				continue
			}

			if street.EndIntersection.isGreen(street, simulationTimestamp) {
				// Set car to next street
				if len(car.Path) > 1 {
					car.Path = car.Path[1:]
					car.DurationOnCurrentStreet = car.Path[0].Length
				} else {
					// Car has completed its path, remove
					car.Delete()
					d.UpdateScore(simulationTimestamp)
					continue
				}

				// Remove car from street
				if len(street.Cars) > 1 {
					street.Cars = street.Cars[0:]
				} else {
					street.Cars = []*Car{}
				}

				// Move car to next street
				nextStreet := car.Path[0]
				nextStreet.Cars = append(nextStreet.Cars, car)
			}
		}
	}
}

func (d *Dataset) SetSchedules() {
	for _, intersection := range d.Intersections {
		// Eliminate empty streets by setting traffic light to red
		streets := d.getAllUnUsedStreets(d.Cars)

		for _, street := range streets {
			for streetIndex, intersectionStreet := range intersection.Schedule.Streets {
				if street.Name == intersectionStreet.Name {
					intersection.Schedule.Duration[streetIndex] = 0
				}
			}
		}

		// Set traffic lights to green if only one road incoming
		if len(intersection.Schedule.Streets) == 1 {
			intersection.Schedule.Duration[0] = d.Time
		}
	}
}

func (d *Dataset) getAllUnUsedStreets(cars []Car) []Street {
	streetNames := make(map[string]bool)
	for _, car := range cars {
		for _, street := range car.Path {
			streetNames[street.Name] = true
		}
	}

	var streets []Street

	for name := range streetNames {
		for _, street := range d.Streets {
			if street.Name == name {
				continue
			}
			streets = append(streets, street)
		}
	}

	return streets
}

func (d *Dataset) getAllUsedStreets(cars []Car) []Street {
	streetNames := make(map[string]bool)
	for _, car := range cars {
		for _, street := range car.Path {
			streetNames[street.Name] = true
		}
	}

	var streets []Street

	for name := range streetNames {
		for _, street := range d.Streets {
			if street.Name == name {
			streets = append(streets, street)
			}
		}
	}

	return streets
}
