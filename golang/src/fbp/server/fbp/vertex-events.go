package fbp

type ChangedOwnerEvent struct {
	ID    int64
	Graph Graph
}

func (e *ChangedOwnerEvent) GetEventType() string {
	return "ChangedOwner"
}

type ConnectedEvent struct {
	V    Vertex
	Edge Edge
}

func (e *ConnectedEvent) GetEventType() string {
	return "ConnectedEvent"
}

type DisconnectEvent struct {
	V Vertex
}

func (e *DisconnectEvent) GetEventType() string {
	return "DisconnectEvent"
}
