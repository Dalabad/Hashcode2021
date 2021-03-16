package model

// Dataset contains all data read from the dataset
// and will be used throughout the simulation
type Dataset struct {
	SliceOrders int   // Number of pizza slices that are ordered
	PizzaTypes  int   // Number of different pizza types
	PizzaSlices []int // Number of slices per pizza type

	PizzaOrderCount int   // Number of different pizzas to order
	PizzaOrderTypes []int // Types of Pizzas
}
