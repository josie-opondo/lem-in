package solution

import "lem-in/structure"

func InitializeAnts(antCount int, startRoom *structure.Room) map[int]string {
	antPositions := make(map[int]string)
	for i := 1; i <= antCount; i++ {
		antPositions[i] = startRoom.Name
	}
	startRoom.Occupied = antCount
	return antPositions
}
