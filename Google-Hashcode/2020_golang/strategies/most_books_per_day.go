package strategies

import (
	"sort"

	"github.com/fatih/color"

	"bitbucket.org/crashtest-security/google-hash-code-2020-team-golang/model"
)

// MostBooksPerDayStrategy implements the most books per day Strategy
type MostBooksPerDayStrategy struct {
	Name string
}

// GetName returns the Name of the Strategy
func (b MostBooksPerDayStrategy) GetName() string {
	return b.Name
}

// Run starts the calculation for the current Strategy

func (b MostBooksPerDayStrategy) Run(dataset model.Dataset) model.Outputset {
	outputset := model.CreateOutputSet()

	sortedLibrariesByMostBooksPerDay := dataset.Libraries

	sort.Slice(sortedLibrariesByMostBooksPerDay, func(i, j int) bool {
		return sortedLibrariesByMostBooksPerDay[i].BooksPerDay < sortedLibrariesByMostBooksPerDay[j].BooksPerDay
	})

	for _, library := range sortedLibrariesByMostBooksPerDay {

		books := library.GetBooksWithoutDuplicate(outputset)
		//books := library.Books

		if len(books) == 0 {
			continue
		}

		outputset.LibraryScanSchedules = append(outputset.LibraryScanSchedules, model.LibraryScanSchedule{
			ID:         library.ID,
			BookAmount: len(books),
			Books:      books,
		})
		outputset.LibraryAmount++
	}

	color.Black("Number of Libraries: %d", dataset.LibraryAmount)
	color.Black("Number of Books: %d", dataset.BookAmount)
	color.Black("Number of Days: %d", dataset.Days)

	return outputset
}
