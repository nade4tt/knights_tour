package grid

func GetGrid(size int) [][]int {
  grid := make([][]int, size)
  for i := range size {
    grid[i] = make([]int, size)
    for j := range size {
      grid[i][j] = -1
    }
  }
  return grid
}
