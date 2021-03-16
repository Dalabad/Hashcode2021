package simulation

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"bitbucket.org/crashtest-security/google-hash-code-2020-team-golang/model"
	"bitbucket.org/crashtest-security/google-hash-code-2020-team-golang/strategies"
	"github.com/fatih/color"
)

// Simulator reads the input data and executes the selected Strategy
type Simulator struct {
	Strategy strategies.StrategyInterface
	Datasets map[string]model.Dataset
}

// Execute starts the simulation based on the selected Strategy and dataset
func (s Simulator) Execute(strategy string, datasetName string) {
	s.Strategy = s.getSelectedStrategy(strategy)
	s.Datasets = s.readInputFile(datasetName)

	color.White("-----------------------------------------")
	for datasetName, dataset := range s.Datasets {
		color.Green("Starting the simulation")
		color.White("Strategy: %s", s.Strategy.GetName())
		color.White("Dataset: %s", datasetName)

		mdl := s.Strategy.Run(dataset)

		s.writeOutputFile(datasetName, mdl)
		color.White("-----------------------------------------")
	}
}

func (s Simulator) getSelectedStrategy(strategy string) strategies.StrategyInterface {
	color.White("Selecting Strategy")

	// TODO: Add strategies
	switch strategy {
	case "BasicStrategy":
		return strategies.BasicStrategy{Name: strategy}
	case "MostBooksPerDayStrategy":
		return strategies.MostBooksPerDayStrategy{Name: strategy}
	case "MaximizeScore":
		return strategies.MaximizeLibraryScore{Name: strategy}
	case "MaximizeLibraryScoreNoDuplicates":
		return strategies.MaximizeLibraryScoreNoDuplicates{Name: strategy}
	case "BooksHighScore":
		return strategies.BooksHighScore{Name: strategy}
	case "LibrariesLowSignup":
		return strategies.LibrariesLowSignup{Name: strategy}
	case "DoNotPrintBookTwice":
		return strategies.DoNotPrintBookTwice{Name: strategy}
	case "MaximizeLibraryScoreNoBookTwice":
		return strategies.MaximizeLibraryScoreNoBookTwice{Name: strategy}
	case "SignupLibraryFirstNoBooksTwice":
		return strategies.SignupLibraryFirstNoBooksTwice{Name: strategy}
	case "SignupLibraryFirstNoBooksTwiceBooksSorted":
		return strategies.SignupLibraryFirstNoBooksTwiceBooksSorted{Name: strategy}
	case "GetLibraryWeightedScore":
		return strategies.GetLibraryWeightedScore{Name: strategy}
	}

	err := fmt.Errorf("selected Strategy %q does not exist", strategy)
	CatchError(err)

	return nil
}

type ChnlResponse struct {
	Filename string
	Dataset  model.Dataset
}

func (s Simulator) readInputFile(datasetName string) map[string]model.Dataset {
	datasets := make(map[string]model.Dataset)

	files := []string{datasetName}

	if datasetName == "all" {
		files = getAllInputFiles()
	}

	chnl := make(chan ChnlResponse)

	for _, fileName := range files {
		go readSingleFile(fileName, chnl)
	}

	for range files {
		chnlResponse := <-chnl
		datasets[chnlResponse.Filename] = chnlResponse.Dataset

		color.White(fmt.Sprintf("Done reading %s", chnlResponse.Filename))
	}

	return datasets
}

func readSingleFile(fileName string, chnl chan ChnlResponse) {
	color.White("Reading Dataset %s", fileName)

	cwd, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/input/%s", cwd, fileName)

	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		err = fmt.Errorf("selected input fileName %q does not exist", filePath)
		CatchError(err)
	}

	inFile, _ := os.Open(filePath)

	var currentLibrary model.Library
	mdl := model.Dataset{}
	currentLine := 0
	currentLibraryNumber := 0

	// Start reading from the file with a reader.
	reader := bufio.NewReader(inFile)

	var line []byte
	for {
		line, _ = reader.ReadBytes('\n')
		lineString := strings.Trim(string(line), "\n")
		lineElements := strings.Split(lineString, " ")

		if lineElements[0] == "" {
			break
		}

		// color.Blue("%d. Line: %q", currentLine, lineElements)

		if currentLine == 0 {
			mdl.BookAmount = getIntValue(lineElements[0])
			mdl.LibraryAmount = getIntValue(lineElements[1])
			mdl.Days = getIntValue(lineElements[2])
		} else if currentLine == 1 {
			for index, score := range lineElements {
				mdl.Books = append(mdl.Books, model.Book{ID: index, Score: getIntValue(score)})
			}
		} else {
			if currentLine%2 == 0 {
				currentLibrary = model.Library{
					ID:             currentLibraryNumber,
					SignupDuration: getIntValue(lineElements[1]),
					BooksPerDay:    getIntValue(lineElements[2]),
					BooksInLibary:  getIntValue(lineElements[0]),
					Books:          nil,
				}

				currentLibraryNumber++
			} else {
				for _, bookID := range lineElements {
					currentLibrary.Books = append(currentLibrary.Books, getBookById(getIntValue(bookID), mdl))
				}
				mdl.Libraries = append(mdl.Libraries, currentLibrary)
			}
		}

		currentLine++
	}

	err = inFile.Close()
	CatchError(err)

	chnl <- ChnlResponse{
		Filename: fileName,
		Dataset:  mdl,
	}
}

func (s Simulator) writeOutputFile(datasetFileName string, data model.Outputset) {
	datasetName := strings.Split(datasetFileName, ".")[0]

	cwd, _ := os.Getwd()
	outputPath := fmt.Sprintf("%s/output/%s.out", cwd, datasetName)

	file, err := os.Create(outputPath)
	CatchError(err)

	_, err = fmt.Fprintln(file, fmt.Sprintf("%d", data.LibraryAmount))
	CatchError(err)
	for _, library := range data.LibraryScanSchedules {
		_, err = fmt.Fprintln(file, fmt.Sprintf("%d %d", library.ID, library.BookAmount))
		CatchError(err)
		for i, book := range library.Books {
			// Increase counter for correct line break at the end of the library
			i++
			_, err = fmt.Fprint(file, fmt.Sprintf("%d ", book.ID))
			if i == len(library.Books) {
				fmt.Fprint(file, "\n")
			}
			CatchError(err)
		}
	}

	err = file.Close()
	CatchError(err)

	color.Green("Output file `%s.out` written successfully", datasetName)
}

// CatchError takes an error, prints its message and ends the application
func CatchError(err error) {
	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}
}

// getIntValue takes a string and converts it to an integer value
func getIntValue(val string) int {
	intVal, err := strconv.Atoi(val)
	CatchError(err)

	return intVal
}

func getAllInputFiles() []string {
	var fileNames []string

	cwd, _ := os.Getwd()
	root := fmt.Sprintf("%s/input/", cwd)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(info.Name(), ".txt") {
			fmt.Printf("Detected new dataset file: %s\n", info.Name())
			fileNames = append(fileNames, info.Name())
		}
		return nil
	})
	if err != nil {
		CatchError(err)
	}

	return fileNames
}

func getBookById(id int, dataset model.Dataset) model.Book {
	for _, book := range dataset.Books {
		if book.ID == id {
			return book
		}
	}

	panic(fmt.Sprintf("Book with ID %d not valid", id))
	return model.Book{}
}
