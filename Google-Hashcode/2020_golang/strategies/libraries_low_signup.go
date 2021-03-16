package strategies

import (
"bitbucket.org/crashtest-security/google-hash-code-2020-team-golang/model"
)

// BasicStrategy implements the basic Strategy
type LibrariesLowSignup struct {
	Name string
}

// GetName returns the Name of the Strategy
func (b LibrariesLowSignup) GetName() string {
	return b.Name
}

// Run starts the calculation for the current Strategy
func (b LibrariesLowSignup) Run(dataset model.Dataset) model.Outputset {
	outputSet := model.Outputset{}

	// sort libraries such that the libraries with the smallest signup duration come first
	libraries := dataset.Libraries
	for i:=0; i<len(libraries); i++ {
		for j:=0; j<len(libraries); j++ {
			if libraries[j+1].SignupDuration < libraries[j].SignupDuration {
				libraries[j], libraries[j+1] = libraries[j+1], libraries[j]
			}
		}
	}

	for _, library := range libraries {
		outputSet.LibraryScanSchedules = append(outputSet.LibraryScanSchedules, model.LibraryScanSchedule{
			ID:         library.ID,
			BookAmount: library.BooksInLibary,
			Books:      library.Books,
		})
	}

	return outputSet
}