package main

import (
	grid "cmd/main.go/pkg/knights_tour/board"
	"fmt"
)

func main() {
	fmt.Println("Hello")
  grid := grid.GetGrid(8)
  fmt.Printf("%d\n", grid[0][0])
}
