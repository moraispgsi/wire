package fbp

import (
	"errors"
	"fmt"
	"strconv"
)

type Graph interface {
	AddVertex(vertex Vertex) int64
	HasVertex(id int64) bool
	GetVertex(id int64) (Vertex, error)
	GetVerticesIDs() []int64
	RemoveVertex(id int64)
	HasEdge(idVertex1 int64, idVertex2 int64) bool
	GetEdge(idVertex1 int64, idVertex2 int64) (Edge, error)
	Connect(idVertex1 int64, idVertex2 int64) (int64, error)
	Disconnect(idVertex1 int64, idVertex2 int64)
}

type graph struct {
	verticesMap          map[int64]Vertex
	connections          map[int64]Edge
	mapStringToEdge      map[string]Edge
	mapVertexToNumEdges  map[int64]int64
	verticesCurrentID    int64 //needs to be atomic
	connectionsCurrentID int64 //needs to be atomic
}

func NewGraph() Graph {
	var graph graph
	graph.verticesMap = make(map[int64]Vertex)
	graph.connections = make(map[int64]Edge)
	graph.mapStringToEdge = make(map[string]Edge)
	graph.mapVertexToNumEdges = make(map[int64]int64)
	graph.verticesCurrentID = 1
	graph.connectionsCurrentID = 1
	return &graph
}

func (graph *graph) accepts(idVertex1 int64, idVertex2 int64) bool {

	if graph.HasEdge(idVertex1, idVertex2) {
		return false
	}

	vertex1, err := graph.GetVertex(idVertex1)

	if err != nil {
		return false
	}

	vertex2, err := graph.GetVertex(idVertex2)

	if err != nil {
		return false
	}

	return vertex1.Accepts(idVertex2) && vertex2.Accepts(idVertex1)

}

func (graph *graph) AddVertex(vertex Vertex) int64 {
	id := graph.verticesCurrentID
	e := ChangedOwnerEvent{
		ID:    id,
		Graph: graph,
	}
	vertex.FireEvent(&e, false)
	graph.verticesMap[id] = vertex
	graph.verticesCurrentID++
	graph.mapVertexToNumEdges[id] = 0
	return id
}

func (graph *graph) HasVertex(id int64) bool {
	_, ok := graph.verticesMap[id]
	return ok
}

func (graph *graph) GetVertex(id int64) (Vertex, error) {

	if _, ok := graph.verticesMap[id]; !ok {
		return nil, errors.New("The vertex id " + strconv.FormatInt(id, 10) + " does not exists.")
	}

	return graph.verticesMap[id], nil

}

func (graph *graph) GetVerticesIDs() []int64 {

	vertexIds := make([]int64, len(graph.verticesMap))

	i := 0
	for k := range graph.verticesMap {
		vertexIds[i] = k
		i++
	}

	return vertexIds

}

func (graph *graph) RemoveVertex(id int64) {
	v := graph.verticesMap[id]
	delete(graph.verticesMap, id)
	delete(graph.mapVertexToNumEdges, id)

	//TODO remove connections and run disconnect handlers

	e := ChangedOwnerEvent{
		ID:    0,
		Graph: nil,
	}

	v.FireEvent(&e, false)
}

func (graph *graph) HasEdge(idVertex1 int64, idVertex2 int64) bool {

	if !graph.HasVertex(idVertex1) || !graph.HasVertex(idVertex2) {
		return false
	}

	_, ok := graph.mapStringToEdge[fmt.Sprintf("%d-%d", idVertex1, idVertex2)]

	return ok

}

func (graph *graph) GetEdge(idVertex1 int64, idVertex2 int64) (Edge, error) {

	if !graph.HasEdge(idVertex1, idVertex2) {
		return nil, errors.New("Edge does not exists.")
	}
	edge := graph.mapStringToEdge[fmt.Sprintf("%d-%d", idVertex1, idVertex2)]
	return edge, nil
}

func (graph *graph) Connect(idVertex1 int64, idVertex2 int64) (int64, error) {

	if !graph.HasVertex(idVertex1) || !graph.HasVertex(idVertex2) {
		return 0, &ConnectionError{error: "Both vertices must exist in the graph"}
	}

	if graph.HasEdge(idVertex1, idVertex2) {
		return 0, &ConnectionError{
			error:     "Connection already exists",
			idVertex1: idVertex1,
			idVertex2: idVertex2,
		}
	}

	vertex1, _ := graph.GetVertex(idVertex1)
	if !vertex1.Accepts(idVertex2) {
		return 0, &ConnectionError{
			error:     "Vertex with Id" + strconv.FormatInt(idVertex1, 10) + " conditions were not met.",
			idVertex1: idVertex1,
			idVertex2: idVertex2,
		}
	}

	vertex2, _ := graph.GetVertex(idVertex2)
	if !vertex2.Accepts(idVertex1) {
		return 0, &ConnectionError{
			error:     "Vertex with Id" + strconv.FormatInt(idVertex2, 10) + " conditions were not met.",
			idVertex1: idVertex1,
			idVertex2: idVertex2,
		}
	}

	connectionID := graph.connectionsCurrentID
	graph.connectionsCurrentID++
	edge := newEdge(connectionID, graph, idVertex1, idVertex2)

	graph.connections[connectionID] = edge
	graph.mapVertexToNumEdges[idVertex1]++
	graph.mapVertexToNumEdges[idVertex2]++
	graph.mapStringToEdge[fmt.Sprintf("%d-%d", idVertex1, idVertex2)] = edge
	graph.mapStringToEdge[fmt.Sprintf("%d-%d", idVertex2, idVertex1)] = edge

	//lauch connection with event fire
	vertex1.FireEvent(&ConnectedEvent{
		V:    vertex1,
		Edge: edge,
	}, false)

	vertex2.FireEvent(&ConnectedEvent{
		V:    vertex2,
		Edge: edge,
	}, false)

	return connectionID, nil
}

func (graph *graph) Disconnect(idVertex1 int64, idVertex2 int64) {

	if !graph.HasVertex(idVertex1) || !graph.HasVertex(idVertex2) {
		return
	}

	edge, err := graph.GetEdge(idVertex1, idVertex2)
	if err != nil {
		return
	}

	vertex1, _ := graph.GetVertex(idVertex1)
	vertex2, _ := graph.GetVertex(idVertex2)

	delete(graph.connections, edge.GetID())
	delete(graph.mapStringToEdge, fmt.Sprintf("%d-%d", idVertex1, idVertex2))
	delete(graph.mapStringToEdge, fmt.Sprintf("%d-%d", idVertex2, idVertex1))

	vertex1.FireEvent(&DisconnectEvent{
		V: vertex2,
	}, false)
	vertex2.FireEvent(&DisconnectEvent{
		V: vertex1,
	}, false)

}
