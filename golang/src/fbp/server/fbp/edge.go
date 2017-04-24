package fbp

import "fmt"

//Connection Error
type ConnectionError struct {
	error     string
	idVertex1 int64
	idVertex2 int64
}

func (e *ConnectionError) Error() string {
	return fmt.Sprintf("Vertex connection error: " + e.error)
}

type Edge interface {
	GetGraph() Graph
	GetID() int64
	GetIDVertex1() int64
	GetIDVertex2() int64
	Disconnect()
}

type edge struct {
	id        int64
	graph     Graph
	idVertex1 int64
	idVertex2 int64
}

func newEdge(id int64, graph Graph, idVertex1 int64, idVertex2 int64) Edge {
	return &edge{
		id,
		graph,
		idVertex1,
		idVertex2,
	}
}

func (e *edge) GetGraph() Graph {
	return e.graph
}
func (e *edge) GetID() int64 {
	return e.id
}
func (e *edge) GetIDVertex1() int64 {
	return e.idVertex1
}
func (e *edge) GetIDVertex2() int64 {
	return e.idVertex2
}
func (e *edge) Disconnect() {
	e.graph.Disconnect(e.idVertex1, e.idVertex2)
}
