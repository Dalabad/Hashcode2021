package strategies

import (
	"bitbucket.org/crashtest-security/google-hash-code-2020-team-golang/practice_round/model"
)

// BasicStrategy implements the basic Strategy
type BasicStrategy struct {
	Name string
}

// GetName returns the Name of the Strategy
func (b BasicStrategy) GetName() string {
	return b.Name
}

// Run starts the calculation for the current Strategy
func (b BasicStrategy) Run(dataset model.Dataset) model.Dataset {

	totalSlices := 0

	// Iterate pizzas in reverse order
	for i := dataset.PizzaTypes - 1; i >= 0; i-- {
		numberOfSlices := dataset.PizzaSlices[i]
		if numberOfSlices+totalSlices > dataset.SliceOrders {
			continue
		}
		dataset.PizzaOrderCount++
		dataset.PizzaOrderTypes = append([]int{i}, dataset.PizzaOrderTypes...)
		totalSlices = totalSlices + numberOfSlices
	}

	return dataset
}
