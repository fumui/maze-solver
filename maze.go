package maze

import "fmt"

func ShortestRoutes(size int, walls []int, start, target int) [][]int {
	possibleRoutes := make([][]int, 0, 100)
	possibleRoutes = append(possibleRoutes, []int{start})
	shortestRoutes := [][]int{}
	found := false
	i := 0
	for !found {
		i++
		toBeRemovedIndex := []int{}
		newPossibleRoute := make([][]int, 0)
		for possibleRouteIndex, possibleRoute := range possibleRoutes {
			position := possibleRoute[(len(possibleRoute) - 1)]
			availablePaths := availablePaths(size, walls, position)
			if len(availablePaths) == 1 && len(possibleRoute) > 1 {
				toBeRemovedIndex = append(toBeRemovedIndex, possibleRouteIndex)
			}
			validPathTaken := 0
			for _, availablePath := range availablePaths {
				if len(possibleRoute) > 1 && availablePath == possibleRoute[(len(possibleRoute)-2)] {
					continue
				}
				newRoute := append([]int{}, possibleRoute...)
				newRoute = append(newRoute, availablePath)
				if availablePath == target {
					found = true
					shortestRoutes = append(shortestRoutes, newRoute)
					break
				}
				if validPathTaken == 0 {
					toBeRemovedIndex = append(toBeRemovedIndex, possibleRouteIndex)
					newPossibleRoute = append(newPossibleRoute, newRoute)
					validPathTaken++
				} else {
					newPossibleRoute = append(newPossibleRoute, newRoute)
				}
			}
			fmt.Printf("possible route %v:\n", possibleRouteIndex)
			fmt.Println(possibleRoute)
		}
		for timesRemoved, index := range toBeRemovedIndex {
			possibleRoutes = remove(possibleRoutes, index-timesRemoved)
		}
		possibleRoutes = append(possibleRoutes, newPossibleRoute...)
		fmt.Println("level:")
		fmt.Println(i)
	}
	return shortestRoutes
}

func availablePaths(size int, walls []int, position int) []int {
	paths := make(map[string]int)

	availablePaths := []int{}
	paths["up"] = position - size
	paths["right"] = position + 1
	paths["down"] = position + size
	paths["left"] = position - 1

	maxBlock := size * size
	for key, path := range paths {
		valid := true
		if (path < 0) || //Over the top maze boundary
			(path > maxBlock) || //Over the bottom maze boundary
			(position%size == 0 && key == "right") || //Over the right maze boundary
			((position-1)%size == 0 && key == "left") { //Over the left maze boundary
			valid = false
		} else {
			for _, wall := range walls {
				if path == wall {
					valid = false
					break
				}
			}
		}
		if valid {
			availablePaths = append(availablePaths, path)
		}
	}
	return availablePaths
}

func remove(slice [][]int, s int) [][]int {
	return append(slice[:s], slice[s+1:]...)
}
