package strategies

import (
"bitbucket.org/crashtest-security/google-hash-code-2020-team-golang/model"
)

// BasicStrategy implements the basic Strategy
type DoNotPrintBookTwice struct {
	Name string
}

// GetName returns the Name of the Strategy
func (b DoNotPrintBookTwice) GetName() string {
	return b.Name
}

// Run starts the calculation for the current Strategy
func (b DoNotPrintBookTwice) Run(dataset model.Dataset) model.Outputset {
	outputSet := model.Outputset{}

	return outputSet
}
