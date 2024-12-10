package solution

import (
	"fmt"
	"lem-in/structure"
	"strings"
)

func SimulateAntMovement(graph *structure.Graph, antCount int, start, end string) {
	paths := FindPaths(graph, start, end)
	if len(paths) == 0 {
		fmt.Println("No paths found from start to end.")
		return
	}

	antPositions := InitializeAnts(antCount, graph.Rooms[start])
	steps := 0

	for len(antPositions) > 0 {
		stepMoves := []string{}

		// Move each ant
		for antID, currentRoom := range antPositions {
			for _, path := range paths {
				// Find the next room on the path
				idx := indexOf(path, currentRoom)
				if idx != -1 && idx+1 < len(path) {
					nextRoom := path[idx+1]

					// Move only if the room isn't occupied (or is the end room)
					if nextRoom == end || graph.Rooms[nextRoom].Occupied == 0 {
						graph.Rooms[currentRoom].Occupied--
						graph.Rooms[nextRoom].Occupied++
						antPositions[antID] = nextRoom
						stepMoves = append(stepMoves, fmt.Sprintf("L%d-%s", antID, nextRoom))

						// Remove ant from tracking if it reaches the end room
						if nextRoom == end {
							delete(antPositions, antID)
						}
						break
					}
				}
			}
		}

		steps++
		fmt.Println(strings.Join(stepMoves, " "))
	}

	fmt.Printf("Simulation completed in %d steps.\n", steps)
}

func indexOf(slice []string, value string) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}
