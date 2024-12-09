package main

import (
	"fmt"
	"lem-in/utils"
)

func main() {
	graph, ants, err := utils.Parser()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Parsed Graph: %+v\n", graph)
	fmt.Printf("Number of Ants: %d\n", ants)
}
