package utils

import (
	"bufio"
	"fmt"
	"lem-in/structure"
	"os"
)

func ParseFile(filename string) (*structure.Graph, int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, 0, fmt.Errorf("cannot open file: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	scanner := bufio.NewScanner(file)

	// Parse Ant Count
	antCount, err := ParseAnts(scanner)
	if err != nil {
		return nil, 0, err
	}

	// Initialize Graph
	graph := &structure.Graph{
		Rooms: make(map[string]*structure.Room),
		Links: []structure.Link{},
	}

	// Parse Rooms and Links
	if err := ParseRoomsAndLinks(scanner, graph); err != nil {
		return nil, 0, err
	}

	return graph, antCount, nil
}
