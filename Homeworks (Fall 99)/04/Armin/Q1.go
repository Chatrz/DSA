package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// funtion for scanning in large scale because of crashing fmt package in windows
var reader *bufio.Reader

func scan(x1, y1, x2, y2 *int) {
	line, _ := reader.ReadString('\n')
	line = strings.TrimSuffix(line, "\n")
	numbers := strings.Split(line, " ")
	*x1, _ = strconv.Atoi(numbers[0])
	*y1, _ = strconv.Atoi(numbers[1])
	*x2, _ = strconv.Atoi(numbers[2])
	*y2, _ = strconv.Atoi(numbers[3])
}

////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////// implemention of gragh in golang
type Vertex struct {
	Key   int
	Edges []*Edge
	Color int
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

type AdjacentItem struct {
	EndPoint   *Vertex
	StartPoint *Vertex
	Edge       *Edge
	Weight     int
	Next       *AdjacentItem
}

func NewVertex(key int) *Vertex {
	return &Vertex{
		Key:   key,
		Edges: []*Edge{},
		Color: -1,
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

func NewAdjacentItem(endPoint, startPoint *Vertex, edge *Edge, weight int) *AdjacentItem {
	return &AdjacentItem{
		EndPoint:   endPoint,
		StartPoint: startPoint,
		Edge:       edge,
		Weight:     weight,
		Next:       nil,
	}
}

func (g *Graph) InsertVertex(key int) {
	if len(g.Vertices) >= g.Capacity {
		fmt.Println("GRAPH IS FULL ! ")
	} else {
		g.Vertices = append(g.Vertices, NewVertex(key))
	}
}

func (g *Graph) InsertEdge(start, end *Vertex, weight int) error {
	newEdge := NewEdge(start, end, weight)
	start.Edges = append(start.Edges, newEdge)
	end.Edges = append(end.Edges, newEdge)
	g.Edges = append(g.Edges, newEdge)
	return nil
}

func (g *Graph) checkEdgeExistence(start, end *Vertex, weight int) bool {
	for _, edge := range g.Edges {
		if edge.Start == start && edge.End == end && edge.Weight == weight {
			return true
		}
	}
	return false
}

// directed value determines if adjacenyList must be directed or not
func (g *Graph) GetAdjacencyList(directed bool) map[*Vertex]*AdjacentItem {
	list := make(map[*Vertex]*AdjacentItem)
	for _, vertex := range g.Vertices {
		counter := 0
		for _, edge := range vertex.Edges {
			if edge.Start == vertex || !directed { // if directed == > it depends on vertex to be starting point but if it is undirected == > it is true anyway
				if counter == 0 {
					list[vertex] = NewAdjacentItem(edge.End, edge.Start, edge, edge.Weight)
					counter++
				} else {
					list[vertex].AddFront(NewAdjacentItem(edge.End, edge.Start, edge, edge.Weight))
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

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////// end of gragh implemention

//////////////////////////////////////////////////////////// algorithm that use dfs algorithm to set color value of connected Vertices a same value ( groupNumber )
var groupNumber int // a variable to hold color value for each group of connected vertices

func (g *Graph) Dfs() {
	status := make(map[*Vertex]int)
	for _, vertex := range g.Vertices {
		status[vertex] = 0
	}
	for i, vertex := range g.Vertices {
		if status[vertex] == 0 { // vertex is unvisited
			groupNumber = i
			g.Dfs_visit(status, vertex)
		}
	}
}

func (g *Graph) Dfs_visit(status map[*Vertex]int, vertex *Vertex) {
	status[vertex] = 1
	vertex.Color = groupNumber
	list := g.AdjacenyList
	temp := list[vertex]
	for temp != nil {
		if status[temp.EndPoint] == 0 {
			g.Dfs_visit(status, temp.EndPoint)
		} else if status[temp.StartPoint] == 0 {
			g.Dfs_visit(status, temp.StartPoint)
		}
		temp = temp.Next
	}
	status[vertex] = 2
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func main() {
	reader = bufio.NewReader(os.Stdin)
	var n int
	var m int
	fmt.Scan(&n)
	fmt.Scan(&m)
	g := NewGraph(n * m)
	var inp string
	// getting graph vertices and check up and left vertex of each new vertex to add edge if needed
	for i := 0; i < n; i++ {
		fmt.Scan(&inp)
		j := 0
		for _, t := range inp {
			if string(t) == "#" {
				g.InsertVertex(0)
			} else {
				g.InsertVertex(1)
				if j-1 >= 0 && g.Vertices[i*m+j-1].Key == 1 {
					g.InsertEdge(g.Vertices[i*m+j-1], g.Vertices[i*m+j], 1)
				}
				if i-1 >= 0 && g.Vertices[(i-1)*m+j].Key == 1 {
					g.InsertEdge(g.Vertices[(i-1)*m+j], g.Vertices[i*m+j], 1)
				}
			}
			j++
		}
	}
	// set the adjacenyList of vertices
	list := g.GetAdjacencyList(false)
	g.AdjacenyList = list
	// setting color of connected vertices
	g.Dfs()
	// asking questions part :
	var q int
	fmt.Scan(&q)
	var x1, y1, x2, y2 int
	for i := 0; i < q; i++ {
		scan(&x1, &y1, &x2, &y2)
		check := g.Vertices[(x1-1)*m+(y1-1)].Color == g.Vertices[(x2-1)*m+(y2-1)].Color
		if check {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
