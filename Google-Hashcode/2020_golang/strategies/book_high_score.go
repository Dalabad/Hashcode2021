package strategies

import (
"bitbucket.org/crashtest-security/google-hash-code-2020-team-golang/model"
)

// BasicStrategy implements the basic Strategy
type BooksHighScore struct {
	Name string
}

// GetName returns the Name of the Strategy
func (b BooksHighScore) GetName() string {
	return b.Name
}

// Run starts the calculation for the current Strategy
func (b BooksHighScore) Run(dataset model.Dataset) model.Outputset {
	outputSet := model.Outputset{}

	for _, library := range dataset.Libraries {
		books := library.Books
		bookCount := 0
		// sort by book score such that the books with the highest score come first
		for i:=0; i<len(books); i++ {
			bookCount++
			for j:=0; j<len(books); j++ {
				if books[j+1].Score > books[j].Score {
					books[j], books[j+1] = books[j+1], books[j]
				}
			}
		}
		outputSet.LibraryScanSchedules = append(outputSet.LibraryScanSchedules, model.LibraryScanSchedule{
			ID:        	library.ID,
			BookAmount: bookCount,
			Books:      books,
		})
	}

	return outputSet
}