package main

import (
	"fmt"
)

// find connected components using dfs algorithm
var groupNumber int
var groups [][]int

/*func main() {
	g := NewGraph(18)
	for i := 0; i < 18; i++ {
		g.InsertVertex(i)
	}

	g.InsertEdge(6, 7, 1)
	g.InsertEdge(6, 11, 1)

	g.InsertEdge(5, 1, 1)
	g.InsertEdge(5, 16, 1)
	g.InsertEdge(5, 17, 1)

	g.InsertEdge(15, 10, 1)
	g.InsertEdge(15, 2, 1)
	g.InsertEdge(15, 9, 1)
	g.InsertEdge(9, 3, 1)

	g.InsertEdge(4, 0, 1)
	g.InsertEdge(8, 0, 1)
	g.InsertEdge(13, 0, 1)
	g.InsertEdge(14, 0, 1)

	// we use undirected adjacent list for this question by passing false as "directed" value to set graph adjacent list
	list := g.GetAdjacencyList(false)
	g.AdjacenyList = list
	g.Dfs1()
	fmt.Println(groups)
}*/

// some little changes to dfs algorithm to solve this problem :

func (g *Graph) Dfs1() {
	groups = make([][]int, len(g.Vertices))
	fmt.Println(" DFS visit : ")
	status := make(map[*Vertex]int)
	for _, vertex := range g.Vertices {
		status[vertex] = 0
	}
	for i, vertex := range g.Vertices {
		if status[vertex] == 0 { // vertex is unvisited
			groupNumber = i
			g.Dfs_visit1(status, vertex)
		}
	}
}

func (g *Graph) Dfs_visit1(status map[*Vertex]int, vertex *Vertex) {
	status[vertex] = 1
	groups[groupNumber] = append(groups[groupNumber], vertex.Key)
	list := g.AdjacenyList
	temp := list[vertex]
	for temp != nil {
		if status[temp.EndPoint] == 0 {
			g.Dfs_visit1(status, temp.EndPoint)
		} else if status[temp.StartPoint] == 0 {
			g.Dfs_visit1(status, temp.StartPoint)
		}
		temp = temp.Next
	}
	status[vertex] = 2
}
