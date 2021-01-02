package main

import (
	"fmt"
)
/// this version of code inspired from == > https://www.geeksforgeeks.org/strongly-connected-components/

var stack *Stack
var connectedComponents []*Vertex

func (g *Graph) DFSUtil(v *Vertex, visited map[*Vertex]bool) {
	// Mark the current node as visited and print it
	visited[v] = true
	fmt.Println(v.Key)
	// Recur for all the vertices adjacent to this vertex
	temp := g.AdjacenyList[v]
	for temp != nil {
		if temp.StartPoint == v {
			if !visited[temp.EndPoint] {
				g.DFSUtil(temp.EndPoint, visited)
			}
		}
		temp = temp.Next
	}
}

func (g *Graph) getVertex(key int) *Vertex {
	for _, v := range g.Vertices {
		if v.Key == key {
			return v
		}
	}
	return nil
}

func (g *Graph) getTranspose(size int) *Graph {
	reverssedGraph := NewGraph(size)
	for _, v := range g.Vertices {
		reverssedGraph.Vertices = append(reverssedGraph.Vertices, v)
	}
	for _, e := range g.Edges {
		reverssedGraph.InsertEdge(e.End.Key, e.Start.Key, e.Weight)
	}
	return reverssedGraph
}

func (g *Graph) fillOrder(v *Vertex, visited map[*Vertex]bool) {
	// Mark the current node as visited
	visited[v] = true

	// Recur for all the vertices adjacent to this vertex
	temp := g.AdjacenyList[v]
	for temp != nil {
		if temp.StartPoint == v {
			if !visited[temp.EndPoint] {
				g.fillOrder(temp.EndPoint, visited)
			}
		}
		temp = temp.Next
	}
	// All vertices reachable from v are processed by now, push v
  fmt.Print("PUSHHH  ")
  fmt.Println(v.Key)
	stack.Push(v)
}

func (g *Graph) printSCCs(size int) {
	stack = CreateStack()
	// Mark all the vertices as not visited (For first DFS)
	visited := make(map[*Vertex]bool)
	for _, v := range g.Vertices {
		visited[v] = false
	}

	// Fill vertices in stack according to their finishing times
	for _, v := range g.Vertices {
		if !visited[v] {
			g.fillOrder(v, visited)
		}
	}

	// Create a reversed graph
	reverssedGraph := g.getTranspose(size)
	reverssedGraph.AdjacenyList = reverssedGraph.GetAdjacencyList(false)
	reverssedGraph.PrintAdjacentList(reverssedGraph.AdjacenyList)

	// Mark all the vertices as not visited (For second DFS)
	visitedReversed := make(map[*Vertex]bool)
	for _, v := range reverssedGraph.Vertices {
		visitedReversed[v] = false
	}

	// Now process all vertices in order defined by Stack
	for _, v := range stack.stack {
		fmt.Println(v.Key)
	}
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@")
	for !stack.IsEmpty() {

		// Pop a vertex from stack
		v := reverssedGraph.getVertex(stack.Pop().Key)
		fmt.Print("STRONGLY CONNECTED TO   ")
		fmt.Println(v.Key)
		// Print Strongly connected component of the popped vertex
		if !visitedReversed[v] {
			connectedComponents = []*Vertex{}
			reverssedGraph.DFSUtil(v, visitedReversed)
			fmt.Println("#######################")
		}
	}
}

func main() {
	g := NewGraph(5)
	for i := 0; i < 5; i++ {
		g.InsertVertex(i)
	}
	g.InsertEdge(1, 0, 1)
	g.InsertEdge(0, 3, 1)
	g.InsertEdge(0, 2, 1)
	g.InsertEdge(2, 1, 1)
	g.InsertEdge(3, 4, 1)
	g.AdjacenyList = g.GetAdjacencyList(false)
	g.printSCCs(5)
}
