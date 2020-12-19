package main

import (
	"fmt"
)

type AdjacentItem struct {
	EndPoint *Vertex
	Weight   int
	Next     *AdjacentItem
}

func NewAdjacentItem(endPoint *Vertex, weight int) *AdjacentItem {
	return &AdjacentItem{
		EndPoint: endPoint,
		Weight:   weight,
		Next:     nil,
	}
}

func (g *Graph) GetAdjacencyList() map[*Vertex]*AdjacentItem {
	list := make(map[*Vertex]*AdjacentItem)
	for _, vertex := range g.Vertices {
		counter := 0
		for _, edge := range vertex.Edges {
			if edge.Start == vertex {
				if counter == 0 {
					list[vertex] = NewAdjacentItem(edge.End, edge.Weight)
					counter++
				} else {
					list[vertex].AddFront(NewAdjacentItem(edge.End, edge.Weight))
				}
			}
		}
	}
	return list
}

func (head *AdjacentItem) AddFront(newAdjacentItem *AdjacentItem) {
	temp := head
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = newAdjacentItem
}

func (g *Graph) PrintAdjacentList(list map[*Vertex]*AdjacentItem) {
	for _, v := range g.Vertices {
		fmt.Print(v.Key)
		fmt.Print(" :  ")
		temp := list[v]
		for temp != nil {
			fmt.Print(" ( ")
			fmt.Print(temp.EndPoint.Key)
			fmt.Print(" , ")
			fmt.Print(temp.Weight)
			fmt.Print(" ) ")
			temp = temp.Next
		}
		fmt.Println()
	}
}
