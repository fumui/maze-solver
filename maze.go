package maze

func ShortestRoutes(size int, wall []int, start, target int) [][]int {

	return nil
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
