package strategies

import (
	"sort"

	"bitbucket.org/crashtest-security/google-hash-code-2020-team-golang/model"
	"github.com/fatih/color"
)

// MaximizeLibraryScoreNoDuplicates implements the highest score Strategy
type MaximizeLibraryScoreNoDuplicates struct {
	Name string
}

// GetName returns the Name of the Strategy
func (b MaximizeLibraryScoreNoDuplicates) GetName() string {
	return b.Name
}

// Run starts the calculation for the current Strategy
func (b MaximizeLibraryScoreNoDuplicates) Run(dataset model.Dataset) model.Outputset {
	outputset := model.CreateOutputSet()

	var sortedLibrariesByScore []KeyValue

	librariesWithScore := getLibraryScore(dataset)
	for _, kv := range librariesWithScore {
		sortedLibrariesByScore = append(sortedLibrariesByScore, kv)
	}

	sort.Slice(sortedLibrariesByScore, func(i, j int) bool {
		return sortedLibrariesByScore[i].Score > sortedLibrariesByScore[j].Score
	})

	for _, kv := range sortedLibrariesByScore {
		
		library := getLibraryById(kv.LibID, dataset)
		
		books := library.GetBooksWithoutDuplicate(outputset)
		
		if len(books) == 0 {
			continue
		}
		
		outputset.LibraryScanSchedules = append(outputset.LibraryScanSchedules, model.LibraryScanSchedule{
			ID:         kv.LibID,
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
