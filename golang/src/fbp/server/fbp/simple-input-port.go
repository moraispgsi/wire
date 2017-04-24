package fbp

import "errors"

type simpleInputPort struct {
	simplePort
	recv func() (interface{}, error)
}

func (port *simpleInputPort) Recv() (interface{}, error) {
	if port.recv != nil {
		return port.recv()
	}
	return nil, errors.New("No receiver function assigned.")
}

func NewSimpleInputPort() InputPort {

	port := newSimplePort(false)

	inputPort := &simpleInputPort{
		simplePort: port,
		recv:       nil,
	}

	port.AddEventListener("InPortChanged", func(e Event) {
		event, _ := e.(*InPortChangedEvent)
		inputPort.recv = event.Recv
	})

	return inputPort
}
