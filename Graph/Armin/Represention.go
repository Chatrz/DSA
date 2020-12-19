package main

import (
	"errors"
	"fmt"
)

type Vertex struct {
	Key   interface{}
	Edges []*Edge
}

type Edge struct {
	Start  *Vertex
	End    *Vertex
	Weight int
}

type Graph struct {
	Vertices []*Vertex
	Edges    []*Edge
	Capacity int
}

func NewVertex(key interface{}) *Vertex {
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
		Vertices: []*Vertex{},
		Edges:    []*Edge{},
		Capacity: capacity,
	}
}

func (g *Graph) InsertVertex(key interface{}) {
	if len(g.Vertices) >= g.Capacity {
		fmt.Println("GRAPH IS FULL ! ")
	} else {
		g.Vertices = append(g.Vertices, NewVertex(key))
	}
}

func (g *Graph) InsertEdge(index1, index2 interface{}, weight int) error {
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

func main() {
	g := NewGraph(10)
	g.InsertVertex(10)
	g.InsertVertex(23)
	g.InsertEdge(10, 23, 5)
	fmt.Println(g.Edges[0].Weight)
}