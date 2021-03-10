package main

func main() {

	files := []string{"a","b","c","d","e","f"}

	d := Dataset{}
	d.readInput(files[0])

	// run simulation

	d.writeOutput()
}

