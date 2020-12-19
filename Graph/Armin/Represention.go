package main

import (
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

func main() {
	g := NewGraph(10)
	g.InsertVertex(10)
	g.InsertVertex(23)
	fmt.Println(g.Vertices[1].Key)
}
