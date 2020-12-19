package main

import (
	"errors"
	"fmt"
)

type Vertex struct {
	Key   int
	Edges []*Edge
}

type Edge struct {
	Start  *Vertex
	End    *Vertex
	Weight int
}

type Graph struct {
	Vertices     []*Vertex
	Edges        []*Edge
	AdjacenyList map[*Vertex]*AdjacentItem
	Capacity     int
}

func NewVertex(key int) *Vertex {
	return &Vertex{
		Key:   key,
		Edges: []*Edge{},
	}
}

func NewEdge(start, end *Vertex, weight int) *Edge {
	return &Edge{
		Start:  start,
		End:    end,
		Weight: weight,
	}
}

func NewGraph(capacity int) *Graph {
	return &Graph{
		Vertices:     []*Vertex{},
		Edges:        []*Edge{},
		AdjacenyList: nil,
		Capacity:     capacity,
	}
}

func (g *Graph) InsertVertex(key int) {
	if len(g.Vertices) >= g.Capacity {
		fmt.Println("GRAPH IS FULL ! ")
	} else {
		g.Vertices = append(g.Vertices, NewVertex(key))
	}
}

func (g *Graph) InsertEdge(index1, index2 int, weight int) error {
	var start *Vertex
	var end *Vertex
	for _, vertex := range g.Vertices {
		if vertex.Key == index1 {
			start = vertex
		} else if vertex.Key == index2 {
			end = vertex
		}
	}
	if start == nil || end == nil {
		return errors.New("unvalid start or end point for edge !")
	} else if g.checkEdgeExistence(start, end, weight) {
		return errors.New("EDGE ALREADY EXISTS!")
	} else {
		newEdge := NewEdge(start, end, weight)
		start.Edges = append(start.Edges, newEdge)
		end.Edges = append(end.Edges, newEdge)
		g.Edges = append(g.Edges, newEdge)
		return nil
	}
}

func (g *Graph) checkEdgeExistence(start, end *Vertex, weight int) bool {
	for _, edge := range g.Edges {
		if edge.Start == start && edge.End == end && edge.Weight == weight {
			return true
		}
	}
	return false
}

func (g *Graph) PrintEdgesList() {
	fmt.Println("Edges List : ")
	for _, edge := range g.Edges {
		fmt.Print(" ( ")
		fmt.Print(edge.Start.Key)
		fmt.Print(" , ")
		fmt.Print(edge.End.Key)
		fmt.Print(" , ")
		fmt.Print(edge.Weight)
		fmt.Print(" )  ")
	}
	fmt.Println()
}

func main() {
	g := NewGraph(12)
	g.InsertVertex(0)
	g.InsertVertex(1)
	g.InsertVertex(9)
	g.InsertVertex(8)
	g.InsertVertex(10)
	g.InsertVertex(11)
	g.InsertVertex(7)
	g.InsertVertex(6)
	g.InsertVertex(5)
	g.InsertVertex(3)
	g.InsertVertex(4)
	g.InsertVertex(2)
	g.InsertEdge(1, 0, 1)
	g.InsertEdge(0, 9, 1)
	g.InsertEdge(9, 8, 1)
	g.InsertEdge(8, 1, 1)
	g.InsertEdge(8, 7, 1)
	g.InsertEdge(7, 10, 1)
	g.InsertEdge(10, 11, 1)
	g.InsertEdge(11, 7, 1)
	g.InsertEdge(7, 3, 1)
	g.InsertEdge(6, 7, 1)
	g.InsertEdge(5, 6, 1)
	g.InsertEdge(3, 5, 1)
	g.InsertEdge(3, 2, 1)
	g.InsertEdge(3, 4, 1)

	matrix := g.GatAdjacencyMatrix()
	g.printMatrix(matrix)
	list := g.GetAdjacencyList()
	g.AdjacenyList = list
	g.PrintAdjacentList(list)
	g.PrintEdgesList()
	g.Dfs()
}
