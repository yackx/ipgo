package dijkstra

type Edge struct {
	vertex1, vertex2 string
	weight int
}

type Graph struct {
	edges []*Edge
}

func NewGraph() *Graph {
	return &Graph{}
}

func (g *Graph) Add(vertex1, vertex2 string, weight int) {
	g.edges = append(g.edges, &Edge{vertex1, vertex2, weight})
}

func (g *Graph) Neighbors(label string) []*Edge {
	neighbors := make([]*Edge, 0)
	for _, edge := range g.edges {
		if edge.vertex1 == label {
			neighbors = append(neighbors, &Edge{edge.vertex1, edge.vertex2, edge.weight})
		} else if edge.vertex2 == label {
			neighbors = append(neighbors, &Edge{edge.vertex2, edge.vertex1, edge.weight})
		}
	}
	return neighbors
}

type Path struct {
	vertex string
	distance int
	previousVertex string
}

//func BuildShortestPaths(graph *Graph) []*Path {
//	reur
//}