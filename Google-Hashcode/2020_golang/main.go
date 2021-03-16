package main

import (
	"flag"

	"bitbucket.org/crashtest-security/google-hash-code-2020-team-golang/simulation"
)

func main() {
	strategy := flag.String("strategy", "", "Select the strategy to use for the simulation")
	inputFile := flag.String("data", "all", "Select the input file to use for the simulation")
	flag.Parse()

	simulator := simulation.Simulator{}
	simulator.Execute(*strategy, *inputFile)
}
