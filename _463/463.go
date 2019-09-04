package _463

func islandPerimeter(grid [][]int) int {
	var cost = 4
	total := 0
	xl := len(grid)
	if xl == 0 {
		return 0
	}

	yl := len(grid[0])
	for i, v1 := range grid {
		for j, v2 := range v1 {
			if v2 == 1 {
				total += cost

				if i > 0 && grid[i-1][j] == 1 {
					total -= 1
				}
				if i < xl-1 && grid[i+1][j] == 1 {
					total -= 1
				}

				if j > 0 && grid[i][j-1] == 1 {
					total -= 1
				}
				if j < yl-1 && grid[i][j+1] == 1 {
					total -= 1
				}
			}
		}
	}
	return total
}
