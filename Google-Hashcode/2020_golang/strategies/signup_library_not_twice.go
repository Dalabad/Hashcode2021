package strategies

import (
	"sort"

	"bitbucket.org/crashtest-security/google-hash-code-2020-team-golang/model"
	"github.com/fatih/color"
)

// SignupLibraryFirstNoBooksTwice implements the highest score Strategy
type SignupLibraryFirstNoBooksTwice struct {
	Name string
}

// GetName returns the Name of the Strategy
func (b SignupLibraryFirstNoBooksTwice) GetName() string {
	return b.Name
}

// Run starts the calculation for the current Strategy
func (b SignupLibraryFirstNoBooksTwice) Run(dataset model.Dataset) model.Outputset {
	outputset := model.Outputset{}

	printedBooks := make(map[int]bool)

	var sortedLibrariesByScore []KeyValue

	librariesWithScore := getLibraryBySignup(dataset)
	for _, kv := range librariesWithScore {
		sortedLibrariesByScore = append(sortedLibrariesByScore, kv)
	}

	sort.Slice(sortedLibrariesByScore, func(i, j int) bool {
		return sortedLibrariesByScore[i].Score < sortedLibrariesByScore[j].Score
	})

	for _, kv := range sortedLibrariesByScore {

		library := getLibraryById(kv.LibID, dataset)
		allBooks := library.Books
		var booksToPrint []model.Book

		for _, book := range allBooks {
			if ok, _ := printedBooks[book.ID]; !ok {
				booksToPrint = append(booksToPrint, book)
				printedBooks[book.ID] = true
			}
		}

		if len(booksToPrint) == 0 {
			continue
		}

		outputset.LibraryScanSchedules = append(outputset.LibraryScanSchedules, model.LibraryScanSchedule{
			ID:         kv.LibID,
			BookAmount: library.BooksInLibary,
			Books:      booksToPrint,
		})
		outputset.LibraryAmount++
	}

	color.Black("Number of Libraries: %d", dataset.LibraryAmount)
	color.Black("Number of Books: %d", dataset.BookAmount)
	color.Black("Number of Days: %d", dataset.Days)

	return outputset
}

func getLibraryBySignup(dataset model.Dataset) []KeyValue {
	var resp []KeyValue

	for _, lib := range dataset.Libraries {
		resp = append(resp, KeyValue{
			LibID: lib.ID,
			Score: lib.SignupDuration,
		})
	}

	return resp
}
