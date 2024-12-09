package utils

import (
	"bufio"
	"fmt"
	"lem-in/structure"
	"os"
	"strconv"
	"strings"
)

func Parser() (*structure.Graph, int, error) {
	// Open the file
	file, err := os.OpenFile("file.txt", os.O_RDONLY, 0444)
	if err != nil {
		return nil, 0, fmt.Errorf("cannot read file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		return nil, 0, fmt.Errorf("cannot read file: %w", err)
	}
	antCount, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return nil, 0, fmt.Errorf("cannot parse antcount: %w", err)
	}
	var startRoom, endRoom *structure.Room
	graph := &structure.Graph{
		Rooms: make(map[string]*structure.Room),
		Links: []structure.Link{},
	}

	for scanner.Scan() {
		input := scanner.Text()

		if input == "##start" {
			// Move to the next line for the start room details
			if !scanner.Scan() {
				return nil, 0, fmt.Errorf("missing start room details")
			}
			fields := strings.Fields(scanner.Text())
			if len(fields) != 3 {
				return nil, 0, fmt.Errorf("invalid format for start room")
			}
			name := fields[0]
			x, _ := strconv.Atoi(fields[1])
			y, _ := strconv.Atoi(fields[2])
			startRoom = &structure.Room{Name: name, X: x, Y: y, IsStart: true}
			graph.Rooms[name] = startRoom
		} else if input == "##end" {
			// Move to the next line for the end room details
			if !scanner.Scan() {
				return nil, 0, fmt.Errorf("missing end room details")
			}
			fields := strings.Fields(scanner.Text())
			if len(fields) != 3 {
				return nil, 0, fmt.Errorf("invalid format for end room")
			}
			name := fields[0]
			x, _ := strconv.Atoi(fields[1])
			y, _ := strconv.Atoi(fields[2])
			endRoom = &structure.Room{Name: name, X: x, Y: y, IsEnd: true}
			graph.Rooms[name] = endRoom
		} else if strings.Contains(input, " ") { // Parse a normal room
			fields := strings.Fields(input)
			if len(fields) != 3 {
				return nil, 0, fmt.Errorf("invalid format for room: %s", input)
			}
			name := fields[0]
			x, _ := strconv.Atoi(fields[1])
			y, _ := strconv.Atoi(fields[2])
			room := &structure.Room{Name: name, X: x, Y: y, IsStart: false, IsEnd: false}
			graph.Rooms[name] = room
		} else if strings.Contains(input, "-") { // Parse a link
			parts := strings.Split(input, "-")
			if len(parts) != 2 {
				return nil, 0, fmt.Errorf("invalid format for link: %s", input)
			}
			room1 := parts[0]
			room2 := parts[1]
			graph.Links = append(graph.Links, structure.Link{Room1: room1, Room2: room2})
		}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return nil, 0, fmt.Errorf("error reading file: %w", err)
	}

	// Print the parsed graph for debugging
	fmt.Printf("Start Room: %+v\n", startRoom)
	fmt.Printf("End Room: %+v\n", endRoom)
	fmt.Printf("All Rooms: %+v\n", graph.Rooms)
	fmt.Printf("All Links: %+v\n", graph.Links)

	return graph, antCount, nil
}
