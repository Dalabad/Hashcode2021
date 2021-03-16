package strategies

import (
	"bitbucket.org/crashtest-security/google-hash-code-2020-team-golang/model"
	"github.com/fatih/color"
	"sort"
)

// GetLibraryWeightedScore implements the highest score Strategy
type GetLibraryWeightedScore struct {
	Name string
}

// GetName returns the Name of the Strategy
func (b GetLibraryWeightedScore) GetName() string {
	return b.Name
}

// Run starts the calculation for the current Strategy
func (b GetLibraryWeightedScore) Run(dataset model.Dataset) model.Outputset {
	outputset := model.Outputset{}

	printedBooks := make(map[int]bool)

	var sortedLibrariesByScore []KeyValue

	librariesWithScore := getLibraryByWeight(dataset)
	for _, kv := range librariesWithScore {
		sortedLibrariesByScore = append(sortedLibrariesByScore, kv)
	}

	sort.Slice(sortedLibrariesByScore, func(i, j int) bool {
		return sortedLibrariesByScore[i].Score >= sortedLibrariesByScore[j].Score
	})

	for _, kv := range sortedLibrariesByScore {
		var booksToPrint []model.Book

		library := getLibraryById(kv.LibID, dataset)
		allBooks := sortBooksByScore(library.Books)

		for _, book := range allBooks {
			if ok, _ := printedBooks[book.ID]; !ok {
				booksToPrint = append(booksToPrint, book)
				printedBooks[book.ID] = true
			}
		}

		if len(booksToPrint) == 0 {
			continue
		}

		outputset.LibraryAmount++

		outputset.LibraryScanSchedules = append(outputset.LibraryScanSchedules, model.LibraryScanSchedule{
			ID:         kv.LibID,
			BookAmount: len(booksToPrint),
			Books:      booksToPrint,
		})
	}

	color.Black("Number of Libraries: %d", dataset.LibraryAmount)
	color.Black("Number of Books: %d", dataset.BookAmount)
	color.Black("Number of Days: %d", dataset.Days)

	return outputset
}

func getLibraryByWeight(dataset model.Dataset) []KeyValue {
	var resp []KeyValue

	for _, lib := range dataset.Libraries {
		bookScore := 0
		for _, book := range lib.Books {
			bookScore += book.Score
		}

		scanSpeedFactor := 1
		signupFactor := 1
		bookScoreFactor := 1

		resp = append(resp, KeyValue{
			LibID: lib.ID,
			Score: (bookScore*bookScoreFactor * lib.BooksPerDay*scanSpeedFactor)/(lib.SignupDuration*signupFactor),
		})
	}

	return resp
}
