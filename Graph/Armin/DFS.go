package main

import (
	"fmt"
)

func (g *Graph) Dfs() {
	fmt.Println(" DFS visit : ")
	status := make(map[*Vertex]int) // 0 stands for unvisited ,
	//                              1 stands for visited but unfinished and 2 stands for visited and finished
	// set all vertices unvisited
	for _, vertex := range g.Vertices {
		status[vertex] = 0
	}

	for _, vertex := range g.Vertices {
		if status[vertex] == 0 { // vertex is unvisited
			g.Dfs_visit(status, vertex)
		}
	}
}

func (g *Graph) Dfs_visit(status map[*Vertex]int, vertex *Vertex) {
	status[vertex] = 1
	fmt.Println(vertex.Key)
	list := g.AdjacenyList
	temp := list[vertex]
	for temp != nil {
		if status[temp.EndPoint] == 0 {
			g.Dfs_visit(status, temp.EndPoint)
		}
		temp = temp.Next
	}
	status[vertex] = 2
}
