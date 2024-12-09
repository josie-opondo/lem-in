package utils

import (
	"bufio"
	"errors"
	"lem-in/structure"
	"os"
	"strconv"
	"strings"
)

func Parser() {
	contents, err := os.OpenFile("file.txt", os.O_RDONLY, 0444)
	if err != nil {
		errors.New("Cannot read file")
	}

	scanner := bufio.NewScanner(contents)
	allInput := []string{}
	for scanner.Scan() {
		input := scanner.Text()
		allInput = append(allInput, input)

		for i:= 0; i < len(allInput); i++ {
			if allInput[i] == "##start" {
				i++
				fields := strings.Fields(allInput[i])
				name := fields[0]
				x, _ := strconv.Atoi(fields[1])
				y, _ := strconv.Atoi(fields[2])
				startRoom := &structure.Room {
					Name: name,
					X: x,
					Y: y,
					IsStart: true,
				}
			} else if allInput[i] == "##end" {
				i++
				fields := strings.Fields(allInput[i])
				name := fields[0]
				x, _ := strconv.Atoi(fields[1])
				y, _ := strconv.Atoi(fields[2])
				endRoom := &structure.Room{Name: name, X: x, Y: y, IsEnd: true }
			} else if strings.Contains(allInput[i], " ") {
				fields := strings.Fields(allInput[i])
				name := fields[0]
				x, _ := strconv.Atoi(fields[1])
				y, _ := strconv.Atoi(fields[2])
				rooms[name] := &structure.Graph{Rooms: name}
			}
		}
	}
}