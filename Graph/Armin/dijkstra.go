package main

import (
	"fmt"
)

// undirected version inpired from == > https://www.youtube.com/watch?v=pVfj6mxhdMw

func (g *Graph) DijkstraUndirected(startingPoint *Vertex) (map[*Vertex]int, map[*Vertex]*Vertex) {
	status := make(map[*Vertex]bool)        // true == > visited ,false == > unvisited
	shortPath := make(map[*Vertex]int)      // shortest path from starting point to current vertex
	prevVertex := make(map[*Vertex]*Vertex) // previous vertex in path from starting point to vertex
	// Mark all the vertices as not visited
	for _, v := range g.Vertices {
		status[v] = false
		shortPath[v] = 1000000
		prevVertex[v] = nil
	}
	shortPath[startingPoint] = 0

	// setting adjacency list for vertices
	g.AdjacenyList = g.GetAdjacencyList(false)
	g.PrintAdjacentList(g.AdjacenyList)

	currentVertex := startingPoint
	for currentVertex != nil {
		status[currentVertex] = true
		currentEdge := g.AdjacenyList[currentVertex] // setting head of list as current edge
		minDistance := 100000
		var closeNeighbor *Vertex
		closeNeighbor = nil
		var neighbor *Vertex

		for currentEdge != nil {

			if currentVertex == currentEdge.EndPoint {
				neighbor = currentEdge.StartPoint
			} else {
				neighbor = currentEdge.EndPoint
			}

			distances := currentEdge.Weight + shortPath[currentVertex]
			if distances < shortPath[neighbor] {
				shortPath[neighbor] = distances
				prevVertex[neighbor] = currentVertex
			}
			if !status[neighbor] {
				if minDistance > currentEdge.Weight {
					minDistance = currentEdge.Weight
					closeNeighbor = neighbor
				}
			}

			currentEdge = currentEdge.Next
		}
		currentVertex = closeNeighbor
	}
	return shortPath, prevVertex
}

func printDijsktraResult(path map[*Vertex]int, prev map[*Vertex]*Vertex) {
	for key, value := range path {
		fmt.Print(" for ")
		fmt.Print(key.Key)
		fmt.Print(" shotest path value is :  ")
		fmt.Print(value)
		fmt.Print("   prev vertes : ")
		if prev[key] == nil {
			fmt.Println(" NONE !")
		} else {
			fmt.Println(prev[key].Key)
		}
		fmt.Println()
	}
}

func main() {
	g := NewGraph(5)

	for i := 0; i < 5; i++ { // 0 ==A , 1 == B , 2 == C , 3 == D , 4 == E
		g.InsertVertex(i)
	}

	g.InsertEdge(0, 1, 6)
	g.InsertEdge(0, 3, 1)

	g.InsertEdge(1, 3, 2)
	g.InsertEdge(1, 4, 2)
	g.InsertEdge(1, 2, 5)

	g.InsertEdge(3, 4, 1)
	g.InsertEdge(4, 2, 5)

	printDijsktraResult(g.DijkstraUndirected(g.Vertices[1]))

}
