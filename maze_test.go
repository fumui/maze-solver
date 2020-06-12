package maze

import "testing"

const AVAILABLE_PATH_ERR_STRING = "Test Failed on input:\n\tsize:\t%v\n\twalls:\t%v\n\tposition:\t%v\nExpecting output:\n\t%v\nBut got\n\t%v"
const SHORTEST_ROUTE_ERR_STRING = "Test Fail, Expecting \n%v \nbut got \n%v"

func TestAvailablePaths(t *testing.T) {
	basicTesting := func(t *testing.T, size int, walls []int, position int, expectedOutput []int) {
		output := availablePaths(size, walls, position)
		error := func(t *testing.T, size int, walls []int, position int, expectedOutput []int) {
			t.Errorf(
				AVAILABLE_PATH_ERR_STRING,
				size,
				walls,
				position,
				expectedOutput,
				output,
			)
		}
		if len(output) != len(expectedOutput) {
			error(t, size, walls, position, expectedOutput)
		}

		for valIndex, val := range output {
			if expectedOutput[valIndex] == val {
				error(t, size, walls, position, expectedOutput)
			}
		}
	}
	simpleWalls := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 21, 22, 23, 24, 25, 26, 28}
	basicTesting(t, 10, simpleWalls, 12, []int{11, 13})
	basicTesting(t, 10, simpleWalls, 27, []int{17, 37})
	basicTesting(t, 10, simpleWalls, 85, []int{84, 75, 86, 95})
	basicTesting(t, 10, simpleWalls, 80, []int{79, 70, 90})

	wallChar := []byte("#")[0]
	maze := []byte("#####################        #   #       ###### ##### # ##### ##   # #   # # #   # ## ### # ### ### # # ## #   # #       # # ## ### # # ### ##### ##       #   #   #   #### ####### # ##### ##   #       #       #### ### ### ####### ##         # #   #   ## ### ### ### ### #### # #   #       # # ## # # ### ##### ### ##   #   # #   # #   ## ### ### # ####### ## # #   # # # #   # ## # ### ### # # # # ##   #           ######################")
	walls := []int{}
	for index, char := range maze {
		if char == wallChar {
			walls = append(walls, (index + 1))
		}
	}
	basicTesting(t, 21, walls, 22, []int{23})
}

/*
	21 x 21 maze

	#  #  #  #  #  #  #  #  #  #  #  #  #  #  #  #  #  #  #  #  #
	22 23 24 25 26 27       #           #                       #
	#  #  #  #  #  48 #  #  #  #  #     #     #  #  #  #  #     #
	#           #  69 #           #     #     #           #     #
	#     #  #  #  90 #     #  #  #     #  #  #     #     #     #
	#     #       111 #     #                       #     #     #
	#     #  #  # 132 #     #     #  #  #     #  #  #  #  #     #
	#       151152153       #           #           #           #
 	#  #  # 172 #  #  #  #  #  #  #     #     #  #  #  #  #     #
	#       193 #                       #                       #
	#  #  # 214 #  #  #     #  #  #     #  #  #  #  #  #  #     #
	#       235236237             #     #           #           #
	#     #  #  # 258 #  #  #     #  #  #     #  #  #     #  #  #
	#     #     # 279       #                       #     #     #
	#     #     # 300 #  #  #     #  #  #  #  #     #  #  #     #
	#           # 321       #     #           #     #           #
	#     #  #  # 342 #  #  #     #     #  #  #  #  #  #  #     #
	#     #     # 363364365 #     #     #     # 373374375 #     #
	#     #     #  #  # 386 #  #  #     #     # 394 # 396 #     #
	#           #       407408409410411412413414415 # 417418419420
	#  #  #  #  #  #  #  #  #  #  #  #  #  #  #  #  #  #  #  #  #

	1 line = #####################********#***#*******######*#####*#*#####*##***#*#***#*#*#***#*##*###*#*###*###*#*#*##*#***#*#*******#*#*##*###*#*#*###*#####*##*******#***#***#***####*#######*#*#####*##***#*******#*******####*###*###*#######*##*********#*#***#***##*###*###*###*###*####*#*#***#*******#*#*##*#*#*###*#####*###*##***#***#*#***#*#***##*###*###*#*#######*##*#*#***#*#*#*#***#*##*#*###*###*#*#*#*#*##***#***********#****#####################
	* = space
	# = wall

	11 x 11 maze

	 #  #  #  #  #  #  #  #  #  #  #
	 12 13 14 15 16 17       #     #
	 #  #  #  #  #  28 #  #  #     #
	 #        #     39 #           #
	 #     #  #  #  50 #     #   # #
	 #     #        61 #     #     #
	 #     #  #  #  72 #     #     #
	 #        81 82 83             #
 	 #  #  #  92 #  #  #  #  #  #  #
	 #       103104105106107108109110
	 #  #  #  #  #  #  #  #  #  #  #

	1 line = ###########********#*######*###*##***#*#***##*###*#*####*#***#*#*##*###*#*#*##*********####*########**********###########
	* = space
	# = wall
*/
func getWalls(mazeString string, wallChar byte) []int {
	// maze21by21 := []byte("#####################********#***#*******######*#####*#*#####*##***#*#***#*#*#***#*##*###*#*###*###*#*#*##*#***#*#*******#*#*##*###*#*#*###*#####*##*******#***#***#***####*#######*#*#####*##***#*******#*******####*###*###*#######*##*********#*#***#***##*###*###*###*###*####*#*#***#*******#*#*##*#*#*###*#####*###*##***#***#*#***#*#***##*###*###*#*#######*##*#*#***#*#*#*#***#*##*#*###*###*#*#*#*#*##***#***********#****#####################")
	// maze11by11 := []byte("###########********#*######*###*##***#*#***##*###*#*####*#***#*#*##*###*#*#*##*********####*########**********###########")
	maze := []byte(mazeString)
	walls := []int{}
	for index, char := range maze {
		if char == wallChar {
			walls = append(walls, (index + 1))
		}
	}
	return walls
}
func TestShortestRoute(t *testing.T) {
	wallChar := []byte("#")[0]
	mazeString := "#####################********#***#*******######*#####*#*#####*##***#*#***#*#*#***#*##*###*#*###*###*#*#*##*#***#*#*******#*#*##*###*#*#*###*#####*##*******#***#***#***####*#######*#*#####*##***#*******#*******####*###*###*#######*##*********#*#***#***##*###*###*###*###*####*#*#***#*******#*#*##*#*#*###*#####*###*##***#***#*#***#*#***##*###*###*#*#######*##*#*#***#*#*#*#***#*##*#*###*###*#*#*#*#*##***#***********#****#####################"
	walls := getWalls(mazeString, wallChar)
	shortestRoutes := ShortestRoutes(21, walls, 23, 420)
	expected := []int{23, 24, 25, 26, 27, 48, 69, 90, 111, 132, 153, 152, 151, 172, 193, 214, 235, 236, 237, 258, 279, 300, 321, 342, 363, 364, 365, 386, 407, 408, 409, 410, 411, 412, 413, 414, 415, 394, 373, 374, 375, 396, 417, 418, 419, 420}
	for index, step := range shortestRoutes[0] {
		if step != expected[index] {
			t.Errorf(SHORTEST_ROUTE_ERR_STRING, expected, shortestRoutes[0])
			break
		}
	}
	DisplayMaze(21, walls, shortestRoutes[0])

	mazeString = "###########********#*######*###*##***#*#***##*###*#*####*#***#*#*##*###*#*#*##*********####*########**********###########"
	walls = getWalls(mazeString, wallChar)
	shortestRoutes = ShortestRoutes(11, walls, 12, 110)
	expected = []int{12, 13, 14, 15, 16, 17, 28, 39, 50, 61, 72, 83, 82, 81, 92, 103, 104, 105, 106, 107, 108, 109, 110}
	for index, step := range shortestRoutes[0] {
		if step != expected[index] {
			t.Errorf(SHORTEST_ROUTE_ERR_STRING, expected, shortestRoutes[0])
			break
		}
	}
	DisplayMaze(11, walls, shortestRoutes[0])

	shortestRoutes = ShortestRoutes(5, []int{6, 7, 8, 10}, 1, 11)
	expected = []int{1, 2, 3, 4, 9, 14, 13, 12, 11}
	for index, step := range shortestRoutes[0] {
		if step != expected[index] {
			t.Errorf(SHORTEST_ROUTE_ERR_STRING, expected, shortestRoutes[0])
			break
		}
	}
}
