package fbp

/*
TODO:
-Inforce IsUnit Contraint (save num of connections by registering a listener on ConnectedEvent and DisconnectEvent
and creating an accept condition)
-

*/

import "errors"

type Event interface {
	GetEventType() string
}

type Vertex interface {
	GetID() int64
	GetGraph() (Graph, error)
	IsUnit() bool
	Accepts(to int64) bool
	AddAcceptCondition(handler func(int64, int64) bool)
	FireEvent(e Event, isAsync bool)
	AddEventListener(t string, listener func(e Event))
}

//Vertex
type vertex struct {
	id               int64
	graph            Graph
	isUnit           bool
	acceptConditions []func(int64, int64) bool
	listeners        map[string][]func(e Event)
}

func NewVertex(isUnit bool) Vertex {

	v := &vertex{
		id:               -1,
		graph:            nil,
		isUnit:           isUnit,
		acceptConditions: []func(int64, int64) bool{},
		listeners:        make(map[string][]func(e Event)),
	}

	v.AddEventListener("ChangedOwner", func(e Event) {
		connEvent := e.(*ChangedOwnerEvent)
		v.graph = connEvent.Graph
		v.id = connEvent.ID
	})

	return v
}

func (vertex *vertex) GetID() int64 {
	return vertex.id
}

func (vertex *vertex) GetGraph() (Graph, error) {

	if vertex.graph == nil {
		return nil, errors.New("This vertex doesnt have a graph owner")
	}

	return vertex.graph, nil
}

func (vertex *vertex) IsUnit() bool {
	return vertex.isUnit
}

func (vertex *vertex) Accepts(to int64) bool {

	for _, acceptFunc := range vertex.acceptConditions {

		if !acceptFunc(vertex.id, to) {
			return false
		}

	}

	return true
}

func (vertex *vertex) FireEvent(e Event, isAsync bool) {

	t := e.GetEventType()

	for _, listener := range vertex.listeners[t] {
		if isAsync {
			go listener(e)
		} else {
			listener(e)
		}

	}

}

func (vertex *vertex) AddAcceptCondition(handler func(int64, int64) bool) {

	if handler != nil {
		vertex.acceptConditions = append(vertex.acceptConditions, handler)
	}

}

func (vertex *vertex) AddEventListener(t string, listener func(e Event)) {
	if listener != nil {
		vertex.listeners[t] = append(vertex.listeners[t], listener)
	}
}
