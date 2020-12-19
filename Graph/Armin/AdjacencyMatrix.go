package main

import "fmt"

func (g *Graph) GatAdjacencyMatrix() map[int]map[int]int {
	matrix := make(map[int]map[int]int)
	for _, rawVertex := range g.Vertices {
		matrix[rawVertex.Key] = make(map[int]int)
		for _, vertex := range g.Vertices {
			matrix[rawVertex.Key][vertex.Key] = 0
		}
		for _, edge := range rawVertex.Edges {
			if edge.Start == rawVertex {
				matrix[rawVertex.Key][edge.End.Key] = edge.Weight
			}
		}
	}
	return matrix
}

func (g *Graph) printMatrix(matrix map[int]map[int]int) {
  fmt.Print("    ")
	for _, vertex := range g.Vertices {
		fmt.Print(vertex.Key)
		fmt.Print("  ")
	}
	fmt.Println()
	for _, vertex := range g.Vertices {
		fmt.Print(vertex.Key)
		fmt.Print("  ")
		for _, vertex2 := range g.Vertices {
			fmt.Print(matrix[vertex.Key][vertex2.Key])
      fmt.Print("   ")
		}
		fmt.Println()
	}
}
