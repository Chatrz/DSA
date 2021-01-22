package main

import (
	"fmt"
)

type AdjacentItem struct {
	EndPoint   *Vertex
	StartPoint *Vertex
	Weight     int
	Next       *AdjacentItem
}

func NewAdjacentItem(endPoint, startPoint *Vertex, weight int) *AdjacentItem {
	return &AdjacentItem{
		EndPoint:   endPoint,
		StartPoint: startPoint,
		Weight:     weight,
		Next:       nil,
	}
}

// directed value determines if adjacenyList must be directed or not
func (g *Graph) GetAdjacencyList(directed bool) map[*Vertex]*AdjacentItem {
	list := make(map[*Vertex]*AdjacentItem)
	for _, vertex := range g.Vertices {
		counter := 0
		for _, edge := range vertex.Edges {
			if edge.Start == vertex || !directed { // if directed == > it depends on vertex to be starting point but if it is undirected == > it is true anyway
				if counter == 0 {
					list[vertex] = NewAdjacentItem(edge.End, edge.Start, edge.Weight)
					counter++
				} else {
					list[vertex].AddFront(NewAdjacentItem(edge.End, edge.Start, edge.Weight))
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
			if temp.StartPoint == v {
				fmt.Print(temp.EndPoint.Key)
			} else {
				fmt.Print(temp.StartPoint.Key)
			}
			fmt.Print(" , ")
			fmt.Print(temp.Weight)
			fmt.Print(" ) ")
			temp = temp.Next
		}
		fmt.Println()
	}
}
