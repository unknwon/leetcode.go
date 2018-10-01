package leetcode

var (
	g       [][]byte
	height  int
	width   int
	checked [][]bool
)

func search(i, j int) {
	if i < 0 || i >= height || j < 0 || j >= width || checked[i][j] {
		return
	}
	checked[i][j] = true

	if g[i][j] == '1' {
		search(i+1, j)
		search(i-1, j)
		search(i, j+1)
		search(i, j-1)
	}
}

func numIslands(grid [][]byte) int {
	numIslands := 0
	g = grid
	height = len(grid)
	if height == 0 {
		return 0
	}

	width = len(grid[0])
	if width == 0 {
		return 0
	}

	checked = make([][]bool, height)
	for i := range checked {
		checked[i] = make([]bool, width)
	}

	for i := range grid {
		for j := range grid[i] {
			if checked[i][j] {
				continue
			}
			checked[i][j] = true

			if grid[i][j] == '1' {
				numIslands++
				search(i+1, j)
				search(i-1, j)
				search(i, j+1)
				search(i, j-1)
			}
		}
	}
	return numIslands
}
