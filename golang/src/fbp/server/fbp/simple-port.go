package fbp

type simplePort struct {
	Vertex
	name      string
	component Component
}

func (port *simplePort) GetComponent() Component {
	return port.component
}

func (port *simplePort) GetName() string {
	return port.name
}

func newSimplePort(isUnit bool) simplePort {
	vertex := NewVertex(isUnit)

	port := simplePort{
		Vertex:    vertex,
		name:      "",
		component: nil,
	}

	port.AddEventListener("PortChanged", func(e Event) {
		portChangeEvent, _ := e.(*PortChangedEvent)
		port.component = portChangeEvent.Component
		port.name = portChangeEvent.Name
	})

	return port
}
