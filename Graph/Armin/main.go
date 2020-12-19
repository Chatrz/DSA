package main

func main() {
	g := NewGraph(12)
	for i := 0; i < 12; i++ {
		g.InsertVertex(i)
	}
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

	// we use undirected adjacent list for this question by passing false as "directed" value to set graph adjacent list
	list := g.GetAdjacencyList(false)
	g.AdjacenyList = list
  g.PrintAdjacentList(list)
  g.BfsWalk(g.Vertices[0])
}
