package utils

import (
	"bufio"
	"fmt"
	"strconv"
)

func ParseAnts(scanner *bufio.Scanner) (int, error) {
	if scanner.Scan() {
		line := scanner.Text()
		antCount, err := strconv.Atoi(line)
		if err != nil || antCount <= 0 {
			return 0, fmt.Errorf("invalid number of ants: %s", line)
		}
		return antCount, nil
	}
	return 0, fmt.Errorf("missing ant count")
}
