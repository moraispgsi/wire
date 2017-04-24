package fbp

type PortChangedEvent struct {
	Name      string
	Component Component
}

func (e *PortChangedEvent) GetEventType() string {
	return "PortChanged"
}

type InPortChangedEvent struct {
	Recv func() (interface{}, error)
}

func (e *InPortChangedEvent) GetEventType() string {
	return "InPortChanged"
}

type OutPortChangedEvent struct {
	Send  func(IP interface{}) error
	Close func()
}

func (e *OutPortChangedEvent) GetEventType() string {
	return "OutPortChanged"
}
