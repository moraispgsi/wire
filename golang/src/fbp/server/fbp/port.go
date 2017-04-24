package fbp

type Port interface {
	Vertex
	GetName() string
	GetComponent() Component
}

type ArrayPort interface {
	Port
	GetIndex() int64
}

type InputPort interface {
	Port
	Recv() (interface{}, error)
}

type OutputPort interface {
	Port
	Send(IP interface{}) error
	Close()
}

type ArrayInputPort interface {
	ArrayPort
	Recv(index int64) (interface{}, error)
}
