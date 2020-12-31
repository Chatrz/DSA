package main

import (
	"fmt"
)

// undirected version inpired from == > https://www.youtube.com/watch?v=pVfj6mxhdMw
// directed version needed changes came in line **48** of this implementation.

func (g *Graph) DijkstraUndirected(startingPoint *Vertex) (map[*Vertex]int, map[*Vertex]*Vertex) {
	status := make(map[*Vertex]bool)        // true == > visited ,false == > unvisited
	shortPath := make(map[*Vertex]int)      // shortest path from starting point to current vertex
	prevVertex := make(map[*Vertex]*Vertex) // previous vertex in path from starting point to vertex
	// Mark all the vertices as not visited
	for _, v := range g.Vertices {
		status[v] = false
		shortPath[v] = 100000000
		prevVertex[v] = nil
	}
	shortPath[startingPoint] = 0

	// setting adjacency list for vertices
	g.AdjacenyList = g.GetAdjacencyList(false)
	g.PrintAdjacentList(g.AdjacenyList)

	for _, currentVertex := range g.Vertices {
		if status[currentVertex] {
			continue
		}
		status[currentVertex] = true
		for currentVertex != nil {
			//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
			// If we are only interested in a shortest path between startingPoint and a particular vertex like V we can end the program here by adding :
			/*
				 if currentVertex == V {
					 break
				 }
			*/
			//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
			status[currentVertex] = true
			currentEdge := g.AdjacenyList[currentVertex] // setting head of list as current edge
			minDistance := 100000
			var closeNeighbor *Vertex
			closeNeighbor = nil
			var neighbor *Vertex

			for currentEdge != nil {
				// if we wanted to do perform this algorithm to a directed graph the only change we had to make was
				// putting this piece of code instead of line  56  to 60  :
				/*if currentVertex == currentEdge.StartPoint {
					neighbor = currentEdge.EndPoint
				} else {
					currentEdge = currentEdge.Next
					continue
				}*/
				if currentVertex == currentEdge.EndPoint {
					neighbor = currentEdge.StartPoint
				} else {
					neighbor = currentEdge.EndPoint
				}
				distances := currentEdge.Weight + shortPath[currentVertex]
				if distances < shortPath[neighbor] {
					shortPath[neighbor] = distances
					prevVertex[neighbor] = currentVertex
					g.CheckNeighbors(neighbor, shortPath, prevVertex)
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

	}
	return shortPath, prevVertex
}

func (g *Graph) CheckNeighbors(currentVertex *Vertex, shortPath map[*Vertex]int, prevVertex map[*Vertex]*Vertex) {
	currentEdge := g.AdjacenyList[currentVertex] // setting head of list as current edge
	var neighbor *Vertex
	for currentEdge != nil {
		// if we wanted to do perform this algorithm to a directed graph the only change we had to make was
		// putting this piece of code instead of line  53  to 57  :
		/*if currentVertex == currentEdge.StartPoint {
			neighbor = currentEdge.EndPoint
		} else {
			currentEdge = currentEdge.Next
			continue
		}*/
		if currentVertex == currentEdge.EndPoint {
			neighbor = currentEdge.StartPoint
		} else {
			neighbor = currentEdge.EndPoint
		}
		distances := currentEdge.Weight + shortPath[currentVertex]
		if distances < shortPath[neighbor] {
			shortPath[neighbor] = distances
			prevVertex[neighbor] = currentVertex
			g.CheckNeighbors(neighbor, shortPath, prevVertex)
		}
		currentEdge = currentEdge.Next
	}
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

func PrintPath(startVertex, destVertex *Vertex, prev map[*Vertex]*Vertex) {
	currentVertex := destVertex
	fmt.Println(" PATH IS : (It is back ward version )")
	for currentVertex != nil {
		fmt.Print(currentVertex.Key)
		fmt.Print(" < ==  ")
		if currentVertex == startVertex {
			break
		}
		currentVertex = prev[currentVertex]
	}
	fmt.Println()
}

<<<<<<< HEAD
/*
15 21
1 2
1 3
3 8
8 10
7 8
7 10
8 9
1 7
1 6
2 5
2 4
5 6
6 7
6 11
11 5
7 9
12 4
9 15
15 14
14 13
13 12
*/
func main() {
	g := NewGraph(15)
	for i := 1; i <= 15; i++ { // 0 ==A , 1 == B , 2 == C , 3 == D , 4 == E
=======
/*func main() {
	g := NewGraph(5)
	for i := 0; i < 5; i++ { // 0 ==A , 1 == B , 2 == C , 3 == D , 4 == E
>>>>>>> armin
		g.InsertVertex(i)
	}
	g.InsertEdge(1, 2, 1)
	g.InsertEdge(1, 3, 1)
	g.InsertEdge(3, 8, 1)
	g.InsertEdge(8, 10, 1)
	g.InsertEdge(7, 8, 1)
	g.InsertEdge(7, 10, 1)
	g.InsertEdge(8, 9, 1)
	g.InsertEdge(1, 7, 1)
	g.InsertEdge(1, 6, 1)
	g.InsertEdge(2, 5, 1)
	g.InsertEdge(2, 4, 1)
	g.InsertEdge(5, 6, 1)
	g.InsertEdge(6, 7, 1)
	g.InsertEdge(6, 11, 1)
	g.InsertEdge(11, 5, 1)
	g.InsertEdge(7, 9, 1)
	g.InsertEdge(12, 4, 1)
	g.InsertEdge(9, 15, 1)
	g.InsertEdge(15, 14, 1)
	g.InsertEdge(14, 13, 1)
	g.InsertEdge(13, 12, 1)

	path, prev := g.DijkstraUndirected(g.Vertices[0])
	printDijsktraResult(path, prev)
	PrintPath(g.Vertices[0], g.Vertices[1], prev)
}*/
