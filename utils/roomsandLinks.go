package utils

import (
	"bufio"
	"fmt"
	"lem-in/structure"
	"strings"
)

func ParseRoomsAndLinks(scanner *bufio.Scanner, graph *structure.Graph) error {
	var startRoom, endRoom *structure.Room

	for scanner.Scan() {
		line := scanner.Text()

		// Handle Room Definitions
		if strings.HasPrefix(line, "##") {
			// Handle special commands (start/end)
			if line == "##start" {
				room, err := ParseRoom(scanner, true, false)
				if err != nil {
					return err
				}
				startRoom = room
				graph.Rooms[room.Name] = room
			} else if line == "##end" {
				room, err := ParseRoom(scanner, false, true)
				if err != nil {
					return err
				}
				endRoom = room
				graph.Rooms[room.Name] = room
			}
		} else if strings.Contains(line, " ") {
			// Handle normal room
			room, err := ParseRoomFromLine(line, false, false)
			if err != nil {
				return err
			}
			graph.Rooms[room.Name] = room
		} else if strings.Contains(line, "-") {
			// Handle link
			link, err := ParseSingleLink(line, graph)
			if err != nil {
				return err
			}
			graph.Links = append(graph.Links, link)
		}
	}

	if startRoom == nil || endRoom == nil {
		return fmt.Errorf("missing start or end room")
	}

	return nil
}
