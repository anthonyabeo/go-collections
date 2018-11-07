package tests

import (
	"go-collections/graphs"
	"testing"
)

// Setup
var dalg = graphs.NewDALGraph()

func TestDirectedAdjacencyListGraph_NumEdges(t *testing.T) {
	if dalg.NumEdges() != 0 {
		t.Errorf("Incorrect, got: %d, expected: %d.", dalg.NumEdges(), 0)
	}
}

func TestDirectedAdjacencyListGraph_NumVertices(t *testing.T) {
	if dalg.NumVertices() != 0 {
		t.Errorf("Incorrect, got: %d, expected: %d.", dalg.NumVertices(), 0)
	}
}

func TestDirectedAdjacencyListGraph_InsertVertex(t *testing.T) {
	v := dalg.InsertVertex("thor")
	if v.Item() != "thor" {
		t.Errorf("Incorrect, got: %d, expected: %s.", v.Item(), "thor")
	}

	v = dalg.InsertVertex("loki")
	if v.Item() != "loki" {
		t.Errorf("Incorrect, got: %d, expected: %s.", v.Item(), "loki")
	}

	v = dalg.InsertVertex(nil)
	if v != nil {
		t.Errorf("Incorrect, got: %d, expected: %v.", v, nil)
	}

	if dalg.NumVertices() != 2 {
		t.Errorf("Incorrect, got: %d, expected: %v.", dalg.NumVertices(), 1)
	}
}

func TestDirectedAdjacencyListGraph_InsertEdge(t *testing.T) {
	e, _ := dalg.InsertEdge(0, 1, "brother")
	if e.Weight() != "brother" {
		t.Errorf("Incorrect, got: %d, expected: %v.", e.Weight(), "brother")
	}

	if dalg.NumEdges() != 1 {
		t.Errorf("Incorrect, got: %d, expected: %v.", dalg.NumEdges(), 1)
	}

	if dalg.PrimaryStructure()[0].OutgoingEdges()[0].Weight() != "brother" ||
		dalg.PrimaryStructure()[1].IncomingEdges()[0].Weight() != "brother" {

		t.Error("edge between vertex 0 and 1 broken")
	}

	e, _ = dalg.InsertEdge(1, 4, "brother")
	if e != nil {
		t.Error("Incorrect, should be nil")
	}

	e, _ = dalg.InsertEdge(0, 1, "brother")
	if e != nil {
		t.Error("Incorrect, should be nil. Edge already exist")
	}
}

func TestDirectedAdjacencyListGraph_GetEdge(t *testing.T) {
	e, _ := dalg.GetEdge(0, 1)
	if e == nil {
		t.Error("Incorrect. Edge actually exist")
	}

	if e.Weight() != "brother" {
		t.Errorf("Incorrect, got: %d, expected: %v.", e.Weight(), "brother")
	}

	e, _ = dalg.GetEdge(1, 4)
	if e != nil {
		t.Error("Incorrect. Edge DOES NOT actually exist. e should be nil")
	}
}

func TestDirectedAdjacencyListGraph_RemoveEdge(t *testing.T) {
	e, _ := dalg.GetEdge(0, 1)
	if e == nil {
		t.Error("Incorrect. Edge actually exist")
	}

	dalg.RemoveEdge(e)

	e, _ = dalg.GetEdge(0, 1)
	if e != nil {
		t.Error("Incorrect. Edge does not exist. Should be nil")
	}

	if dalg.NumEdges() != 0 {
		t.Errorf("Incorrect, got: %d, expected: %v.", dalg.NumEdges(), 0)
	}

	e, _ = dalg.GetEdge(1, 4)
	if e != nil {
		t.Error("Incorrect. Edge DOES NOT actually exist. e should be nil")
	}

	if dalg.RemoveEdge(e) {
		t.Error("Incorrect. Edge DOES NOT actually exist.")
	}

	if dalg.RemoveEdge(graphs.NewEdgeWithDefaultData()) {
		t.Error("Incorrect. Edge DOES NOT actually exist.")
	}
}

func TestDirectedAdjacencyListGraph_RemoveVertex(t *testing.T) {
	dalg.RemoveVertex(0)
	if dalg.NumVertices() != 1 {
		t.Errorf("Incorrect, got: %d, expected: %v.", dalg.NumVertices(), 1)
	}

	if dalg.RemoveVertex(0) {
		t.Error("Incorrect. Vertex DOES NOT actually exist.")
	}
}

func TestIntegration(t *testing.T) {
	// Create a new graph
	g := graphs.NewDALGraph()

	// Insert all the vertices
	v1 := g.InsertVertex(1)
	v2 := g.InsertVertex(2)
	v3 := g.InsertVertex(3)
	v4 := g.InsertVertex(4)
	v5 := g.InsertVertex(5)

	// Insert all the edges
	e12, _ := g.InsertEdge(v1.Id(), v2.Id(), "12")
	g.InsertEdge(v1.Id(), v4.Id(), "14")
	g.InsertEdge(v3.Id(), v1.Id(), "31")
	e34, _ := g.InsertEdge(v3.Id(), v4.Id(), "34")
	g.InsertEdge(v4.Id(), v5.Id(), "45")
	g.InsertEdge(v5.Id(), v3.Id(), "53")

	e912, _ := g.InsertEdge(9, 12, "912")
	if e912 != nil {
		t.Errorf("Incorrect, cannot insert such edge")
	}

	// Checking for the number of edges and vertices
	if g.NumVertices() != 5 {
		t.Errorf("Incorrect, got: %d, expected: %v.", g.NumVertices(), 5)
	}
	if g.NumEdges() != 6 {
		t.Errorf("Incorrect, got: %d, expected: %v.", g.NumEdges(), 6)
	}

	// Testing GetEdge()
	e, _ := g.GetEdge(v1.Id(), v2.Id())
	if e.Id() != e12.Id() {
		t.Errorf("Incorrect, got: %d, expected: %v.", e.Id(), e12.Id())
	}
	if e.Weight() != e12.Weight() {
		t.Errorf("Incorrect, got: %d, expected: %v.", e.Weight(), e12.Weight())
	}

	e, _ = g.GetEdge(v3.Id(), v4.Id())
	if e.Id() != e34.Id() {
		t.Errorf("Incorrect, got: %d, expected: %v.", e.Id(), e34.Id())
	}
	if e.Weight() != e34.Weight() {
		t.Errorf("Incorrect, got: %d, expected: %v.", e.Weight(), e34.Weight())
	}

	e, _ = g.GetEdge(5, 3)
	if e != nil {
		t.Error("Incorrect. That edge is invalid")
	}

	// Testing OutDegree
	if g.OutDegree(v1.Id()) != 2 {
		t.Errorf("Incorrect, got: %d, expected: %v.", g.OutDegree(v1.Id()), 2)
	}
	if g.OutDegree(v3.Id()) != 2 {
		t.Errorf("Incorrect, got: %d, expected: %v.", g.OutDegree(v3.Id()), 2)
	}
	if g.OutDegree(v5.Id()) != 1 {
		t.Errorf("Incorrect, got: %d, expected: %v.", g.OutDegree(v5.Id()), 1)
	}
	if g.OutDegree(9) != -1 {
		t.Errorf("Incorrect, got: %d, expected: %v.", g.OutDegree(9), -1)
	}

	// Testing InDegree
	if g.InDegree(v4.Id()) != 2 {
		t.Errorf("Incorrect, got: %d, expected: %v.", g.InDegree(v4.Id()), 2)
	}
	if g.InDegree(v3.Id()) != 1 {
		t.Errorf("Incorrect, got: %d, expected: %v.", g.InDegree(v3.Id()), 1)
	}
	if g.InDegree(v2.Id()) != 1 {
		t.Errorf("Incorrect, got: %d, expected: %v.", g.InDegree(v2.Id()), 1)
	}
	if g.InDegree(9) != -1 {
		t.Errorf("Incorrect, got: %d, expected: %v.", g.InDegree(9), -1)
	}

	// Testing Breadth-First Search
	t.Log(g.BFS())
}
