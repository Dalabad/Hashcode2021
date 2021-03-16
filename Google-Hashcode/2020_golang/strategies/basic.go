package strategies

import (
	"bitbucket.org/crashtest-security/google-hash-code-2020-team-golang/model"
	"github.com/fatih/color"
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
func (b BasicStrategy) Run(dataset model.Dataset) model.Outputset {
	outputset := model.Outputset{}

	color.Black("Number of Libraries: %d", dataset.LibraryAmount)
	color.Black("Number of Books: %d", dataset.BookAmount)
	color.Black("Number of Days: %d", dataset.Days)

	outputset.LibraryAmount = 0

	// Iterate over all libraries and add them to the output set
	for i := 0; i < dataset.LibraryAmount; i++ {

		library := dataset.Libraries[i]

		libraryScanSchedule := model.LibraryScanSchedule{}

		libraryScanSchedule.ID = library.ID
		libraryScanSchedule.BookAmount = library.BooksInLibary
		libraryScanSchedule.Books = library.Books

		outputset.LibraryAmount++
		outputset.LibraryScanSchedules = append(outputset.LibraryScanSchedules, libraryScanSchedule)
	}

	return outputset
}
