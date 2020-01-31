package maze

import "testing"

func TestAvailablePaths(t *testing.T) {
	basicTesting := func(t *testing.T, size int, walls []int, position int, expectedOutput []int) {
		output := availablePaths(size, walls, position)
		for _, val := range output {
			exist := false
			for _, expectedVal := range expectedOutput {
				if expectedVal == val {
					exist = true
				}
			}
			if !exist || len(output) != len(expectedOutput) {
				t.Errorf(
					"Test Failed on input:\n\tsize:\t%v\n\twalls:\t%v\n\tposition:\t%v\nExpecting output:\n\t%v\nBut got\n\t%v",
					size,
					walls,
					position,
					expectedOutput,
					output,
				)
			}
		}
	}
	walls := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 21, 22, 23, 24, 25, 26, 28}
	basicTesting(t, 10, walls, 12, []int{11, 13})
	basicTesting(t, 10, walls, 27, []int{17, 37})
	basicTesting(t, 10, walls, 85, []int{84, 75, 86, 95})
	basicTesting(t, 10, walls, 80, []int{79, 70, 90})
}
