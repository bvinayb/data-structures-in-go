package graphs

import "fmt"

type Graph struct {
	Vertices []*Vertex
}

type Vertex struct {
	Key       int
	Adjacents []*Vertex
}

func (g *Graph) AddVertex(key int) {
	g.Vertices = append(g.Vertices, &Vertex{Key: key})
}

func (g *Graph) Print() {
	for _, v := range g.Vertices {
		fmt.Printf("Vertex %d:", v)
		for _, val := range v.Adjacents {
			fmt.Printf("%d", val.Key)
		}
	}
}

func (g *Graph) AddEdge(from, to int) {
	fromVertex := g.GetVertex(from)
	toVertex := g.GetVertex(to)

	fromVertex.Adjacents = append(fromVertex.Adjacents, toVertex)
}

func (g *Graph) GetVertex(vertexKey int) *Vertex {
	for i, v := range g.Vertices {
		if v.Key == vertexKey {
			return g.Vertices[i]
		}
	}
	return nil
}
