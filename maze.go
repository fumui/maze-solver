package maze

import (
	"fmt"
	"sort"
)

//ShortestRoutes search the most optimal route to reach the target on a given params
//size := the size of the grid (ex. 5)
//walls := grid numbers that cannot counts as an obstacle (ex. [6,7,8,10])
//start := the starting position (ex. 1)
//target := ex. 11
// 1  2  3  4  5
//######### 9 ###
//11 12 13 14 15
//16 17 18 19 20
//21 22 23 24 25
//return : [1,2,3,4,9,14,13,12,11]
func ShortestRoutes(size int, walls []int, start, target int) [][]int {
	//variable declarations
	possibleRoutes := make([][]int, 0, 100)
	possibleRoutes = append(possibleRoutes, []int{start})
	shortestRoutes := [][]int{}
	found := false

	//loop until reached the target
	for !found {
		//need to remove deadend route after the loop to prevent indexOutOfBound error
		toBeRemovedIndex := []int{}
		//to store newfound route at crossroads or T junctions
		newPossibleRoute := make([][]int, 0)
		//loop through all possible route and find their next available path,
		//if found crossroads or T junctions, then create new slices to store those new routes
		for possibleRouteIndex, possibleRoute := range possibleRoutes {
			position := possibleRoute[(len(possibleRoute) - 1)]
			availablePaths := availablePaths(size, walls, position)
			//if a route doesn't have available path anymore other than the previous position
			//then remove it
			if len(availablePaths) == 1 && len(possibleRoute) > 1 {
				toBeRemovedIndex = append(toBeRemovedIndex, possibleRouteIndex)
			}

			//loop through all possible path and replace current loop with the updated loop
			//and add new route (if any)
			firstAvailablePath := true
			for _, availablePath := range availablePaths {
				if len(possibleRoute) > 1 && availablePath == possibleRoute[(len(possibleRoute)-2)] {
					continue
				}
				newRoute := append([]int{}, possibleRoute...)
				newRoute = append(newRoute, availablePath)
				//if one of the available path is the target then we found the shortest route
				//continue to find another route to check if any other route can reach the target
				//with the same number of step
				if availablePath == target {
					found = true
					shortestRoutes = append(shortestRoutes, newRoute)
					break
				}
				//only need to be called once to remove current route
				if firstAvailablePath {
					toBeRemovedIndex = append(toBeRemovedIndex, possibleRouteIndex)
					firstAvailablePath = false
				}
				newPossibleRoute = append(newPossibleRoute, newRoute)
			}
		}
		//safely remove routes that need to be removed
		sort.Ints(toBeRemovedIndex)
		for timesRemoved, index := range toBeRemovedIndex {
			possibleRoutes = remove(possibleRoutes, index-timesRemoved)
		}
		//add new possible routes if any
		possibleRoutes = append(possibleRoutes, newPossibleRoute...)
	}
	return shortestRoutes
}

//helper function to see get the available path on a maze based on the position
func availablePaths(size int, walls []int, position int) []int {
	paths := make(map[string]int)

	//get all possible move then map it
	availablePaths := []int{}
	paths["up"] = position - size
	paths["right"] = position + 1
	paths["down"] = position + size
	paths["left"] = position - 1

	maxBlock := size * size
	//for all the allowed paths, if it out of bounds or a wall then it's not a valid path
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

//helper function to remove slice of slices element
func remove(slice [][]int, s int) [][]int {
	return append(slice[:s], slice[s+1:]...)
}

func DisplayMaze(size int, walls, path []int) {
	index := 0
	for row := 0; row < size; row++ {
		for collumn := 0; collumn < size; collumn++ {
			index++
			isWall := false
			for _, wall := range walls {
				if index == wall {
					isWall = true
					break
				}
			}
			isPath := false
			for _, step := range path {
				if index == step {
					isPath = true
					break
				}
			}
			if isWall {
				fmt.Print("###")
			} else if isPath {
				fmt.Print(" . ")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println()
	}
}
