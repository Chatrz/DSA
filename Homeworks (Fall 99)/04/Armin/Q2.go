package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var edgeState map[*Edge]bool
var counter int

var reader *bufio.Reader

// main function :
func main() {
	counter = 0
	reader = bufio.NewReader(os.Stdin)
	var v int
	var e int
	fmt.Scan(&v)
	fmt.Scan(&e)
	edgeState = make(map[*Edge]bool)
	g := NewGraph(v)
	for i := 1; i <= v; i++ {
		g.InsertVertex(i)
	}
	var start int
	var end int
	for i := 0; i < e; i++ {
		scan(&start, &end)
		g.InsertEdge(g.Vertices[start-1], g.Vertices[end-1], 1)
	}
	g.BfsWalk(g.Vertices[0])
	fmt.Println(counter)
	counter1 := 0
	for _, edge := range g.Edges {
		counter1++
		if edgeState[edge] {
			fmt.Print(counter1)
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

// scan function for scanning in large scale
func scan(x1, y1 *int) {
	line, _ := reader.ReadString('\n')
	line = strings.TrimSuffix(line, "\n")
	numbers := strings.Split(line, " ")
	*x1, _ = strconv.Atoi(numbers[0])
	*y1, _ = strconv.Atoi(numbers[1])
}

// implementation of graph in Go :
type Vertex struct {
	Key            int
	Edges          []*Edge
	NeccessaryEdge *Edge
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

func NewVertex(key int) *Vertex {
	return &Vertex{
		Key:            key,
		Edges:          []*Edge{},
		NeccessaryEdge: nil,
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

func (g *Graph) InsertVertex(key int) {
	g.Vertices = append(g.Vertices, NewVertex(key))
}

func (g *Graph) InsertEdge(start, end *Vertex, weight int) error {
	newEdge := NewEdge(start, end, weight)
	start.Edges = append(start.Edges, newEdge)
	end.Edges = append(end.Edges, newEdge)
	g.Edges = append(g.Edges, newEdge)
	edgeState[newEdge] = false
	return nil
}

////////////////////////////////////////////////////////// end of implementation

//////// bfs algorithm with some changes to solve this problem
func (g *Graph) BfsWalk(start *Vertex) {
	status := make(map[*Vertex]bool) // true == > visited ,false == > unvisited
	for _, v := range g.Vertices {
		status[v] = false
	}
	queue := NewQueue(10 * len(g.Vertices))
	status[start] = true
	queue.Enqueue(start)
	for !queue.IsEmpty() {
		start, _ = queue.Dequeue()
		for _, edge := range start.Edges {
			if !status[edge.End] {
				edgeState[edge] = true
				status[edge.End] = true
				counter++
				queue.Enqueue(edge.End)
			} else if !status[edge.Start] {
				edgeState[edge] = true
				counter++
				status[edge.Start] = true
				queue.Enqueue(edge.Start)
			}
		}
	}
}

/////////////////////////////////////////////////////////// implementation of queue
type Queue struct {
	queue    []*Vertex
	capacity int
	size     int
	head     int
	tail     int
}

func NewQueue(capacity int) *Queue {
	q := new(Queue)
	q.queue = make([]*Vertex, capacity)
	q.capacity = capacity
	q.tail = capacity - 1
	return q
}

func (q *Queue) Enqueue(key *Vertex) error {
	if q.IsFull() {
		return errors.New("queue overflow error")
	}
	q.tail = (q.tail + 1) % q.capacity
	q.queue[q.tail] = key
	q.size++
	return nil
}

func (q *Queue) Dequeue() (*Vertex, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue underfolow error")
	}
	value := q.queue[q.head]
	q.head = (q.head + 1) % q.capacity
	q.size--
	return value, nil
}

func (q *Queue) IsFull() bool {
	return q.size == q.capacity
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}
