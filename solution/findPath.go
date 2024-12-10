package solution

import "lem-in/structure"

func FindPaths(graph *structure.Graph, start, end string) [][]string {
	// BFS to find paths
	queue := [][]string{{start}}
	visited := map[string]bool{start: true}
	paths := [][]string{}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]

		last := path[len(path)-1]
		if last == end {
			paths = append(paths, path)
			continue
		}

		for _, link := range graph.Links {
			var next string
			if link.Room1 == last {
				next = link.Room2
			} else if link.Room2 == last {
				next = link.Room1
			} else {
				continue
			}

			if !visited[next] {
				visited[next] = true
				newPath := append([]string{}, path...)
				newPath = append(newPath, next)
				queue = append(queue, newPath)
			}
		}
	}

	return paths
}
