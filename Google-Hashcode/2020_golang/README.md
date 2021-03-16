# Start application

## Setup

To get all dependencies call `get dep`

## Launch
To start the application, provide a strategy and the name of the dataset file:

``go run ./main.go --strategy=BasicStrategy --data=simple.in``

The dataset must be provided in the `input` folder. The output of the simulation will be provided in the `output` folder with the same name as the input dataset.

To run all input files in one run, use `all` (default) as the dataset name

``go run ./main.go --strategy=BasicStrategy``

``go run ./main.go --strategy=BasicStrategy --data=all``