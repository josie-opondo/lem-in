package main

import (
	"fmt"
	"lem-in/utils"
)

func main() {
	graph, ants, err := utils.ParseFile("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Parsed Graph: %+v\n", graph)
	fmt.Printf("Number of Ants: %d\n", ants)
}
