package strategies

import (
	"fmt"
	"sort"

	"bitbucket.org/crashtest-security/google-hash-code-2020-team-golang/model"
	"github.com/fatih/color"
)

// MaximizeLibraryScore implements the highest score Strategy
type MaximizeLibraryScore struct {
	Name string
}

// GetName returns the Name of the Strategy
func (b MaximizeLibraryScore) GetName() string {
	return b.Name
}

type KeyValue struct {
	LibID int
	Score int
}

// Run starts the calculation for the current Strategy
func (b MaximizeLibraryScore) Run(dataset model.Dataset) model.Outputset {
	outputset := model.Outputset{}

	var sortedLibrariesByScore []KeyValue

	librariesWithScore := getLibraryScore(dataset)
	for _, kv := range librariesWithScore {
		sortedLibrariesByScore = append(sortedLibrariesByScore, kv)
	}

	sort.Slice(sortedLibrariesByScore, func(i, j int) bool {
		return sortedLibrariesByScore[i].Score > sortedLibrariesByScore[j].Score
	})

	for _, kv := range sortedLibrariesByScore {
		outputset.LibraryAmount++

		library := getLibraryById(kv.LibID, dataset)

		outputset.LibraryScanSchedules = append(outputset.LibraryScanSchedules, model.LibraryScanSchedule{
			ID:         kv.LibID,
			BookAmount: library.BooksInLibary,
			Books:      library.Books,
		})
	}

	color.Black("Number of Libraries: %d", dataset.LibraryAmount)
	color.Black("Number of Books: %d", dataset.BookAmount)
	color.Black("Number of Days: %d", dataset.Days)

	return outputset
}

func getLibraryScore(dataset model.Dataset) []KeyValue {
	var resp []KeyValue

	for _, lib := range dataset.Libraries {
		libScore := 0

		for _, book := range lib.Books {
			libScore += book.Score
		}

		resp = append(resp, KeyValue{
			LibID: lib.ID,
			Score: libScore,
		})
	}

	return resp
}

func getLibraryById(id int, dataset model.Dataset) model.Library {
	for _, library := range dataset.Libraries {
		if library.ID == id {
			return library
		}
	}

	panic(fmt.Sprintf("Library with ID %d not valid", id))
	return model.Library{}
}
