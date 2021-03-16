package model

// Dataset contains all data read from the dataset
// and will be used throughout the simulation
type Dataset struct {
	Days          int
	LibraryAmount int
	BookAmount    int
	Libraries     []Library
	Books         []Book
}

type Library struct {
	ID             int
	SignupDuration int
	BooksPerDay    int
	Books          []Book
	BooksInLibary  int
}

func (library Library) GetBooksWithoutDuplicate(outputset Outputset) []Book {
	books := []Book{}

	for _, book := range library.Books {
		if _, ok := outputset.AlreadyScannedBooks[book.ID]; ok {
			continue
		} else {
			books = append(books, book)
			outputset.AlreadyScannedBooks[book.ID] = true
		}
	}

	return books
}

type Book struct {
	ID    int
	Score int
}

type Outputset struct {
	LibraryAmount        int                   // Number of Libraries to Signup
	LibraryScanSchedules []LibraryScanSchedule // Libraries containing their sending schedule. Ordered by Signup Date

	AlreadyScannedBooks map[int]bool // Helper storing which books were already registered for scanning
}

type LibraryScanSchedule struct {
	ID         int
	BookAmount int
	Books      []Book // Books to scan. Ordered by Scan Pirority?
}

func CreateOutputSet() Outputset {
	return Outputset{
		AlreadyScannedBooks: map[int]bool{},
	}
}
