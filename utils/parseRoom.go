package utils

import (
	"bufio"
	"fmt"
	"lem-in/structure"
	"strconv"
	"strings"
)

func ParseRoom(scanner *bufio.Scanner, isStart, isEnd bool) (*structure.Room, error) {
	if scanner.Scan() {
		line := scanner.Text()
		return ParseRoomFromLine(line, isStart, isEnd)
	}
	return nil, fmt.Errorf("missing room definition after command")
}

func ParseRoomFromLine(line string, isStart, isEnd bool) (*structure.Room, error) {
	fields := strings.Fields(line)
	if len(fields) != 3 {
		return nil, fmt.Errorf("invalid room definition: %s", line)
	}

	name := fields[0]
	x, err := strconv.Atoi(fields[1])
	if err != nil {
		return nil, fmt.Errorf("invalid x coordinate for room %s", name)
	}
	y, err := strconv.Atoi(fields[2])
	if err != nil {
		return nil, fmt.Errorf("invalid y coordinate for room %s", name)
	}

	// Ensure the room name doesn't start with `L` or `#`
	if strings.HasPrefix(name, "L") || strings.HasPrefix(name, "#") {
		return nil, fmt.Errorf("invalid room name: %s", name)
	}

	return &structure.Room{Name: name, X: x, Y: y, IsStart: isStart, IsEnd: isEnd}, nil
}
