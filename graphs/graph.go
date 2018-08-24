package graphs

import (
	"errors"
	"go-collections/linear/lists/linked"
	"math"
)

// A struct to represent each vertex in the graph.
// A list contains an item and a list of all the edges
// the vertex belongs to.
type Vertex struct {
	item     interface{}
	edgeList *linked.DoublyLinkedList
}

// A struct to represent an edge in the graph. An edge contains
// references to the two endpoints (vertices) that make up the
// edge
type Edge struct {
	name    interface{}
	vertex1 *Vertex
	vertex2 *Vertex
}

// An implementation of a graph as an Adjacency List. The graph
// is represented by a list of vertices and edges
type AdjacencyListGraph struct {
	vertices *linked.DoublyLinkedList
	edges    *linked.DoublyLinkedList
}

// returns the number of edges in the graph
func (alg *AdjacencyListGraph) NumEdges() int {
	return alg.edges.Size()
}

// returns the number of vertices in the graph
func (alg *AdjacencyListGraph) NumVertices() int {
	return alg.vertices.Size()
}

// Returns an iteration of all the vertices of the graph
func (alg *AdjacencyListGraph) Vertices() *linked.DoublyLinkedList {
	return alg.vertices
}

// Returns an iteration of all the edge of the graph
func (alg *AdjacencyListGraph) Edges() *linked.DoublyLinkedList {
	return alg.edges
}

// Returns the edge from vertex u to vertex v, if one exists;
// otherwise return null.
func (alg *AdjacencyListGraph) GetEdge(u, v *Vertex) *Edge {
	min := math.Min(float64(u.edgeList.Size()), float64(v.edgeList.Size()))

	var minVertex *Vertex
	if min == float64(u.edgeList.Size()) {
		minVertex = u
	} else {
		minVertex = u
	}

	for curNode := minVertex.edgeList.Head(); curNode.Item() != "trailer"; curNode = curNode.Next() {
		if (curNode.Item().(*Edge).vertex1 == v && curNode.Item().(*Edge).vertex2 == u) ||
			(curNode.Item().(*Edge).vertex2 == v && curNode.Item().(*Edge).vertex1 == u) {
			return curNode.Item().(*Edge)
		}
	}

	return nil
}

// Returns an array containing the two endpoint vertices of
// edge e
func (alg *AdjacencyListGraph) EndVertices(e *Edge) [2]*Vertex {
	x := [2]*Vertex{nil, nil}

	edges := alg.edges
	for curNode := edges.Head(); curNode.Item() != "trailer"; curNode = curNode.Next() {
		if curNode.Item().(*Edge).name == e.name {
			x[0] = curNode.Item().(*Edge).vertex1
			x[1] = curNode.Item().(*Edge).vertex2
			break
		}
	}

	return x
}

// For edge e incident to vertex v, returns the other vertex of
// the edge; an error occurs if e is not incident to v.
func (alg *AdjacencyListGraph) Opposite(v *Vertex, e *Edge) (*Vertex, error) {
	if e.vertex1 == v {
		return e.vertex2, nil
	} else if e.vertex2 == v {
		return e.vertex1, nil
	}

	return nil, errors.New("edge e has no such vertices")
}

// returns the number of edges incident on the vertex v.
func (alg *AdjacencyListGraph) Degree(v *Vertex) int {
	return v.edgeList.Size()
}

// Returns an iteration of all the edge incident of vertex v
func (alg *AdjacencyListGraph) EdgesOf(v *Vertex) *linked.DoublyLinkedList {
	return v.edgeList
}

// Creates and returns a new Vertex storing element x
func (alg *AdjacencyListGraph) InsertVertex(x interface{}) *Vertex {
	v := &Vertex{
		x,
		new(linked.DoublyLinkedList),
	}
	alg.vertices.AddFirst(v)

	return v
}

// Creates and returns a new Edge from vertex u to vertex v,
// storing element x; an error occurs if there already exists an
// edge from u to v.
func (alg *AdjacencyListGraph) InsertEdge(u, v *Vertex, e *Edge, x interface{}) (*Edge, error) {
	if alg.GetEdge(u, v) != nil {
		return nil, errors.New("edge already exist")
	}

	edge := &Edge{x, u, v}
	alg.edges.AddFirst(edge)
	u.edgeList.AddFirst(edge)
	v.edgeList.AddFirst(edge)

	return e, nil
}

// Removes vertex v and all its incident edges from the graph.
func (alg *AdjacencyListGraph) RemoveVertex(v *Vertex) {
	edges := v.edgeList
	for curNode := edges.Head(); curNode.Item() != "trailer"; curNode = curNode.Next() {
		alg.RemoveEdge(curNode.Item().(*Edge))
	}

	v = nil
}

// Removes edge e from the graph
func (alg *AdjacencyListGraph) RemoveEdge(e *Edge) {
	edges := alg.edges
	for curNode := edges.Head(); curNode.Item() != "trailer"; curNode = curNode.Next() {
		if curNode.Next().Item().(*Edge).name == e.name {
			curNode.SetNext(curNode.Next().Next())
			curNode.Next().Next().SetPrev(curNode)
		}
	}
}
