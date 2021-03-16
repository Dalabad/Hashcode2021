package simulation

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"bitbucket.org/crashtest-security/google-hash-code-2020-team-golang/practice_round/model"
	"bitbucket.org/crashtest-security/google-hash-code-2020-team-golang/practice_round/strategies"
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
	color.Blue(strategy)

	// TODO: Add strategies
	switch strategy {
	case "BasicStrategy":
		return strategies.BasicStrategy{Name: strategy}
	}

	err := fmt.Errorf("selected Strategy %q does not exist", strategy)
	CatchError(err)

	return nil
}

func (s Simulator) readInputFile(datasetName string) map[string]model.Dataset {
	datasets := make(map[string]model.Dataset)

	files := []string{datasetName}

	if datasetName == "all" {
		files = getAllInputFiles()
	}

	for _, fileName := range files {
		color.White("Reading Dataset %s", fileName)

		cwd, _ := os.Getwd()
		filePath := fmt.Sprintf("%s/input/%s", cwd, fileName)

		_, err := os.Stat(filePath)
		if os.IsNotExist(err) {
			err = fmt.Errorf("selected input fileName %q does not exist", filePath)
			CatchError(err)
		}

		inFile, _ := os.Open(filePath)

		scanner := bufio.NewScanner(inFile)

		mdl := model.Dataset{}
		currentLine := 0
		for scanner.Scan() {
			line := scanner.Text()
			lineElements := strings.Split(line, " ")
			//color.Blue("%d. Line: %q", currentLine, lineElements)
			// mdl.Something(lineElements[0], lineElements[1], ...)

			if currentLine == 0 {
				mdl.SliceOrders, _ = strconv.Atoi(lineElements[0])
				mdl.PizzaTypes, _ = strconv.Atoi(lineElements[1])
			} else {
				mdl.PizzaSlices = make([]int, len(lineElements))
				for i, numberOfSlicesString := range lineElements {
					numberOfSlices, _ := strconv.Atoi(numberOfSlicesString)
					mdl.PizzaSlices[i] = numberOfSlices
				}
			}

			currentLine++
		}

		err = inFile.Close()
		CatchError(err)

		datasets[fileName] = mdl
	}

	return datasets
}

func (s Simulator) writeOutputFile(datasetFileName string, data model.Dataset) {
	datasetName := strings.Split(datasetFileName, ".")[0]

	cwd, _ := os.Getwd()
	outputPath := fmt.Sprintf("%s/output/%s.out", cwd, datasetName)

	file, err := os.Create(outputPath)
	CatchError(err)

	// Write first line - Number of pizzas to be ordered
	fmt.Fprint(file, fmt.Sprintf("%d\n", data.PizzaOrderCount))

	// Write second line - Pizza types to be ordered
	// Convert slice of integers to slice of strings
	pizzaOrderTypes := strings.Trim(strings.Join(strings.Split(fmt.Sprint(data.PizzaOrderTypes), " "), " "), "[]")
	fmt.Fprint(file, pizzaOrderTypes)

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
		if strings.HasSuffix(info.Name(), ".in") {
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
