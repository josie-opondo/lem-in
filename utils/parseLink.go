package utils

import (
	"bufio"
	"fmt"
	"lem-in/structure"
	"strings"
)

func ParseLink(scanner *bufio.Scanner, graph *structure.Graph) error {
	// Map to track existing links for duplicate detection
	existingLinks := make(map[string]bool)

	for scanner.Scan() {
		line := scanner.Text()

		// Only process lines containing a link
		if !strings.Contains(line, "-") {
			continue
		}

		// Parse the link
		link, err := ParseSingleLink(line, graph)
		if err != nil {
			return err
		}

		// Generate a unique key for the link (unordered)
		linkKey := fmt.Sprintf("%s-%s", link.Room1, link.Room2)
		if link.Room1 > link.Room2 {
			linkKey = fmt.Sprintf("%s-%s", link.Room2, link.Room1)
		}

		// Check for duplicate links
		if existingLinks[linkKey] {
			return fmt.Errorf("duplicate link: %s", line)
		}
		existingLinks[linkKey] = true

		// Add the link to the graph
		graph.Links = append(graph.Links, link)
	}

	return nil
}

func ParseSingleLink(line string, graph *structure.Graph) (structure.Link, error) {
	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		return structure.Link{}, fmt.Errorf("invalid format for link: %s", line)
	}

	room1 := parts[0]
	room2 := parts[1]

	// Self-link check
	if room1 == room2 {
		return structure.Link{}, fmt.Errorf("self-link detected: %s", line)
	}

	// Room existence check
	if _, exists := graph.Rooms[room1]; !exists {
		return structure.Link{}, fmt.Errorf("link references unknown room: %s", room1)
	}
	if _, exists := graph.Rooms[room2]; !exists {
		return structure.Link{}, fmt.Errorf("link references unknown room: %s", room2)
	}

	// Return parsed link
	return structure.Link{Room1: room1, Room2: room2}, nil
}
