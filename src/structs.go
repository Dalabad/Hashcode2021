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
	IsFinished              bool
}

func (c *Car) Delete() {
	c.IsFinished = true
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

	amountIntersections := 0
	for _, intersection := range d.Intersections {
		for _, dur := range intersection.Schedule.Duration {
			if dur > 0 {
				amountIntersections++
				break
			}
		}
	}

	_, err = f.WriteString(fmt.Sprintf("%d\n", amountIntersections))
	if err != nil {
		panic(err)
	}

	for _, intersection := range d.Intersections {

		amountOfSchedules := 0
		for _, dur := range intersection.Schedule.Duration {
			if dur > 0 {
				amountOfSchedules++
			}
		}

		if amountOfSchedules == 0 {
			continue
		}

		// ID of the intersection
		_, err = f.WriteString(fmt.Sprintf("%d\n", intersection.ID))
		if err != nil {
			panic(err)
		}
		// Number of incoming streets
		_, err = f.WriteString(fmt.Sprintf("%d\n", amountOfSchedules))
		if err != nil {
			panic(err)
		}
		// Schedule with the format: "street_name duration"
		for i, street := range intersection.Schedule.Streets {
			if intersection.Schedule.Duration[i] == 0 {
				continue
			}

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
	d.Time = time
	d.Bonus = bonusPoints

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
		endIntersection.Schedule.Duration = append(endIntersection.Schedule.Duration, 0)

		d.Streets = append(d.Streets, street)
	}
	for i := 0; i < carsCount; i++ {
		scanner.Scan()
		carData := strings.Split(scanner.Text(), " ")
		pathLength, _ := strconv.Atoi(carData[0])
		car := Car{Path: make([]Street, pathLength), DurationOnCurrentStreet: 0}

		for j := 1; j <= pathLength; j++ {
			car.Path[j-1] = d.FindStreetByName(carData[j])
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
	for simulationTimestamp := 1; simulationTimestamp <= d.Time; simulationTimestamp++ {
		fmt.Printf("Simulate step %d\n", simulationTimestamp)

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
		numStreets := len(intersection.Schedule.Streets)
		if numStreets == 1 {
			intersection.Schedule.Duration[0] = d.Time
		} else if numStreets > 1 {
			for i := 0; i < numStreets; i++ {
				intersection.Schedule.Duration[i] = 1
			}
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

outer:
	for _, street := range d.Streets {
		for name := range streetNames {
			if street.Name == name {
				break outer
			}
		}
		streets = append(streets, street)
	}

	return streets
}
