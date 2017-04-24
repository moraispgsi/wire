package fbp

import "errors"

type simpleOutputPort struct {
	simplePort
	send  func(IP interface{}) error
	close func()
}

func (port *simpleOutputPort) Send(IP interface{}) error {
	if port.send != nil {
		return port.send(IP)
	}
	return errors.New("No sender function assigned.")
}

func (port *simpleOutputPort) Close() {
	if port.close != nil {
		port.close()
	}
}

func NewSimpleOutputPort() OutputPort {

	port := newSimplePort(true)

	outputPort := &simpleOutputPort{
		simplePort: port,
	}

	port.AddEventListener("OutPortChanged", func(e Event) {
		event, _ := e.(*OutPortChangedEvent)
		outputPort.send = event.Send
		outputPort.close = event.Close
	})

	return outputPort

}
