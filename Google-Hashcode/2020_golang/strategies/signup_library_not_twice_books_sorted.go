package strategies

import (
	"sort"

	"bitbucket.org/crashtest-security/google-hash-code-2020-team-golang/model"
	"github.com/fatih/color"
)

// SignupLibraryFirstNoBooksTwiceBooksSorted implements the highest score Strategy
type SignupLibraryFirstNoBooksTwiceBooksSorted struct {
	Name string
}

// GetName returns the Name of the Strategy
func (b SignupLibraryFirstNoBooksTwiceBooksSorted) GetName() string {
	return b.Name
}

// Run starts the calculation for the current Strategy
func (b SignupLibraryFirstNoBooksTwiceBooksSorted) Run(dataset model.Dataset) model.Outputset {
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
		allBooks := sortBooksByScore(library.Books)
		var booksToPrint []model.Book

		for _, book := range allBooks {
			if ok, _ := printedBooks[book.ID]; !ok {
				booksToPrint = append(booksToPrint, book)
			}
		}

		if len(booksToPrint) == 0 {
			continue
		}

		outputset.LibraryScanSchedules = append(outputset.LibraryScanSchedules, model.LibraryScanSchedule{
			ID:         kv.LibID,
			BookAmount: len(booksToPrint),
			Books:      booksToPrint,
		})
		outputset.LibraryAmount++
	}

	color.Black("Number of Libraries: %d", dataset.LibraryAmount)
	color.Black("Number of Books: %d", dataset.BookAmount)
	color.Black("Number of Days: %d", dataset.Days)

	return outputset
}

func sortBooksByScore(books []model.Book) []model.Book {
	sort.Slice(books, func(i, j int) bool {
		return books[i].Score > books[j].Score
	})
	return books
}
