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

	// Parse the number of ants
	if !scanner.Scan() {
		return nil, 0, fmt.Errorf("file is empty or missing the number of ants")
	}
	antCount, err := strconv.Atoi(scanner.Text())
	if err != nil || antCount <= 0 {
		return nil, 0, fmt.Errorf("invalid number of ants: %s", scanner.Text())
	}

	// Initialize the graph
	var startRoom, endRoom *structure.Room
	graph := &structure.Graph{
		Rooms: make(map[string]*structure.Room),
		Links: []structure.Link{},
	}

	// Parse the rest of the input
	for scanner.Scan() {
		input := scanner.Text()

		if input == "##start" {
			if startRoom != nil {
				return nil, 0, fmt.Errorf("duplicate ##start directive")
			}
			// Move to the next line for the start room details
			if !scanner.Scan() {
				return nil, 0, fmt.Errorf("missing start room details")
			}
			fields := strings.Fields(scanner.Text())
			if len(fields) != 3 {
				return nil, 0, fmt.Errorf("invalid format for start room")
			}
			name := fields[0]
			x, err := strconv.Atoi(fields[1])
			y, err2 := strconv.Atoi(fields[2])
			if err != nil || err2 != nil {
				return nil, 0, fmt.Errorf("invalid coordinates for start room")
			}
			if _, exists := graph.Rooms[name]; exists {
				return nil, 0, fmt.Errorf("duplicate room name: %s", name)
			}
			startRoom = &structure.Room{Name: name, X: x, Y: y, IsStart: true}
			graph.Rooms[name] = startRoom
		} else if input == "##end" {
			if endRoom != nil {
				return nil, 0, fmt.Errorf("duplicate ##end directive")
			}
			// Move to the next line for the end room details
			if !scanner.Scan() {
				return nil, 0, fmt.Errorf("missing end room details")
			}
			fields := strings.Fields(scanner.Text())
			if len(fields) != 3 {
				return nil, 0, fmt.Errorf("invalid format for end room")
			}
			name := fields[0]
			x, err := strconv.Atoi(fields[1])
			y, err2 := strconv.Atoi(fields[2])
			if err != nil || err2 != nil {
				return nil, 0, fmt.Errorf("invalid coordinates for end room")
			}
			if _, exists := graph.Rooms[name]; exists {
				return nil, 0, fmt.Errorf("duplicate room name: %s", name)
			}
			endRoom = &structure.Room{Name: name, X: x, Y: y, IsEnd: true}
			graph.Rooms[name] = endRoom
		} else if strings.Contains(input, " ") { // Parse a normal room
			fields := strings.Fields(input)
			if len(fields) != 3 {
				return nil, 0, fmt.Errorf("invalid format for room: %s", input)
			}
			name := fields[0]
			x, err := strconv.Atoi(fields[1])
			y, err2 := strconv.Atoi(fields[2])
			if err != nil || err2 != nil {
				return nil, 0, fmt.Errorf("invalid coordinates for room: %s", input)
			}
			if strings.HasPrefix(name, "L") || strings.HasPrefix(name, "#") || strings.Contains(name, " ") {
				return nil, 0, fmt.Errorf("invalid room name: %s", name)
			}
			if _, exists := graph.Rooms[name]; exists {
				return nil, 0, fmt.Errorf("duplicate room name: %s", name)
			}
			room := &structure.Room{Name: name, X: x, Y: y, IsStart: false, IsEnd: false}
			graph.Rooms[name] = room
		} else if strings.Contains(input, "-") { // Parse a link
			parts := strings.Split(input, "-")
			if len(parts) != 2 {
				return nil, 0, fmt.Errorf("invalid format for link: %s", input)
			}
			room1 := parts[0]
			room2 := parts[1]
			if room1 == room2 {
				return nil, 0, fmt.Errorf("self-link detected: %s", input)
			}
			if _, exists := graph.Rooms[room1]; !exists {
				return nil, 0, fmt.Errorf("link references unknown room: %s", room1)
			}
			if _, exists := graph.Rooms[room2]; !exists {
				return nil, 0, fmt.Errorf("link references unknown room: %s", room2)
			}
			for _, link := range graph.Links {
				if (link.Room1 == room1 && link.Room2 == room2) || (link.Room1 == room2 && link.Room2 == room1) {
					return nil, 0, fmt.Errorf("duplicate link: %s", input)
				}
			}
			graph.Links = append(graph.Links, structure.Link{Room1: room1, Room2: room2})
		}
	}

	// Final validation: Ensure start and end rooms exist
	if startRoom == nil {
		return nil, 0, fmt.Errorf("missing start room")
	}
	if endRoom == nil {
		return nil, 0, fmt.Errorf("missing end room")
	}

	return graph, antCount, nil
}
