package fbp

import (
	"errors"
	"strconv"
	"sync"
)

//http://www.jpaulmorrison.com/fbp/morrison_2005.htm

type Network interface {
	AddComponent(component Component) (int64, error)
	GetNamespace() int64
	AddInPort(portName string) error
	AddOutPort(portName string) error
	CloseInPort(portName string)
	RemovePort(ID int64) error
	Send(portName string, IP interface{}) error
	Recv(portName string) (interface{}, error)
	Connect(outputNSID int64, outPortName string, inputNSID int64, inPortName string) (int64, error)
	Disconnect(inPortName string, outPortName string) error
	Destroy(connectionID int64) error
	Start()
}

type namespace struct {
	net                 *network
	id                  int64
	inPorts             map[int64]InputPort
	outPorts            map[int64]OutputPort
	mapNameToInPort     map[string]InputPort
	mapNameToOutPort    map[string]OutputPort
	muxInPortReceivers  map[int64]MuxReceiver
	outPortToConnection map[int64]Connection
}

func (ns *namespace) addInPort(portName string, inPort InputPort) {
	ID := ns.net.graph.AddVertex(inPort)
	ns.inPorts[ID] = inPort
	ns.mapNameToInPort[portName] = inPort
	mux := NewMuxReceiver()
	ns.muxInPortReceivers[ID] = mux
	inPort.FireEvent(&InPortChangedEvent{Recv: mux.Recv}, false)

	inPort.AddAcceptCondition(func(v1 int64, v2 int64) bool {
		graph, err := inPort.GetGraph()
		if err == nil {
			vertex2, err := graph.GetVertex(v2)

			if err != nil {
				return false
			}

			_, ok := vertex2.(OutputPort)

			return ok
		}
		return false
	})

}

func (ns *namespace) addOutPort(portName string, outPort OutputPort) {
	ID := ns.net.graph.AddVertex(outPort)
	ns.outPorts[ID] = outPort
	ns.mapNameToOutPort[portName] = outPort
	outPort.FireEvent(&OutPortChangedEvent{Send: func(IP interface{}) error {
		conn, ok := ns.outPortToConnection[ID]
		if ok {
			return conn.Send(IP)
		}
		return nil
	}, Close: func() {
		conn, ok := ns.outPortToConnection[ID]
		if ok {
			conn.CloseOutPort()
		}
	}}, false)

	outPort.AddAcceptCondition(func(v1 int64, v2 int64) bool {
		graph, err := outPort.GetGraph()
		if err == nil {
			vertex2, err := graph.GetVertex(v2)

			if err != nil {
				return false
			}

			_, ok := vertex2.(InputPort)

			return ok
		}
		return false
	})
}

func newNamespace(net *network, ID int64) namespace {
	return namespace{
		net:                 net,
		id:                  ID,
		inPorts:             make(map[int64]InputPort),
		outPorts:            make(map[int64]OutputPort),
		mapNameToInPort:     make(map[string]InputPort),
		mapNameToOutPort:    make(map[string]OutputPort),
		muxInPortReceivers:  make(map[int64]MuxReceiver),
		outPortToConnection: make(map[int64]Connection),
	}
}

type network struct {
	graph              Graph
	connections        map[int64]Connection
	namespaces         map[int64]namespace
	components         map[int64]Component
	currentNamespaceID int64
	netNamespace       int64
	mutex              sync.Mutex
	hasStarted         bool
}

func NewNetwork() Network {
	net := &network{
		graph:              NewGraph(),
		connections:        make(map[int64]Connection),
		namespaces:         make(map[int64]namespace),
		components:         make(map[int64]Component),
		currentNamespaceID: 1,
		netNamespace:       0,
		mutex:              sync.Mutex{},
		hasStarted:         false,
	}

	ns := newNamespace(net, 0)
	net.namespaces[0] = ns

	return net
}

//Returns namespace ID
func (net *network) AddComponent(component Component) (int64, error) {

	ID := net.currentNamespaceID
	ns := newNamespace(net, ID)
	net.namespaces[ID] = ns
	net.components[ID] = component
	net.currentNamespaceID++

	for key, in := range component.GetInputs() {
		ns.addInPort(key, in)
	}

	for key, out := range component.GetOutputs() {
		ns.addOutPort(key, out)
	}

	return ID, nil
}

func (net *network) GetNamespace() int64 {
	return net.netNamespace
}

//Exposes a in port. In port is internally represented as an outputPort
func (net *network) AddInPort(portName string) error {
	ns := net.namespaces[net.netNamespace]
	ns.addOutPort(portName, NewSimpleOutputPort())
	return nil
}

//Exposes a out port. Out port is internally represented as an inputPort
func (net *network) AddOutPort(portName string) error {
	ns := net.namespaces[net.netNamespace]
	ns.addInPort(portName, NewSimpleInputPort())
	return nil
}

func (net *network) CloseInPort(portName string) {
	ns := net.namespaces[net.netNamespace]
	outPort, ok := ns.mapNameToOutPort[portName]
	if !ok {
		return
	}
	conn, ok := ns.outPortToConnection[outPort.GetID()]
	if !ok {
		return
	}
	conn.CloseOutPort()
}

func (net *network) RemovePort(ID int64) error {
	net.mutex.Lock()
	if !net.hasStarted {
		net.mutex.Unlock()
		return errors.New("The network has not started yet.")
	}
	net.mutex.Unlock()
	return nil
}
func (net *network) Send(portName string, IP interface{}) error {
	net.mutex.Lock()
	if !net.hasStarted {
		net.mutex.Unlock()
		return errors.New("The network has not started yet.")
	}
	net.mutex.Unlock()

	ns := net.namespaces[net.netNamespace]
	outPort, ok := ns.mapNameToOutPort[portName]
	if !ok {
		return errors.New("No such out port: " + portName)
	}
	conn, ok := ns.outPortToConnection[outPort.GetID()]
	if !ok {
		return errors.New("Out port has no connection")
	}
	conn.Send(IP)
	return nil

}
func (net *network) Recv(portName string) (interface{}, error) {
	net.mutex.Lock()
	if !net.hasStarted {
		net.mutex.Unlock()
		return nil, errors.New("The network has not started yet.")
	}
	net.mutex.Unlock()

	ns := net.namespaces[net.netNamespace]
	inPort, ok := ns.mapNameToInPort[portName]
	if !ok {
		return nil, errors.New("No such in port: " + portName)
	}
	mux, ok := ns.muxInPortReceivers[inPort.GetID()]
	if !ok {
		return nil, errors.New("No such multiplexer with ID: " + strconv.FormatInt(inPort.GetID(), 10))
	}
	return mux.Recv()

}
func (net *network) Connect(outputNSID int64, outPortName string, inputNSID int64, inPortName string) (int64, error) {

	inputNS, ok := net.namespaces[inputNSID]
	if !ok {
		return 0, errors.New("Input namespace not found")
	}
	outputNS, ok := net.namespaces[outputNSID]
	if !ok {
		return 0, errors.New("Output namespace not found")
	}

	inPort, ok := inputNS.mapNameToInPort[inPortName]
	if !ok {
		return 0, errors.New("No such in port name: " + inPortName)
	}

	outPort, ok := outputNS.mapNameToOutPort[outPortName]
	if !ok {
		return 0, errors.New("No such out port name: " + outPortName)
	}

	ID, err := net.graph.Connect(inPort.GetID(), outPort.GetID())
	if err == nil {
		conn := NewConnection()
		net.connections[ID] = conn
		inputNS.muxInPortReceivers[inPort.GetID()].AddReceiver(conn.Recv)
		outputNS.outPortToConnection[outPort.GetID()] = conn
	}
	return ID, err

}
func (net *network) Disconnect(inPortName string, outPortName string) error {
	net.mutex.Lock()
	defer net.mutex.Unlock()
	if net.hasStarted {
		return errors.New("Can disconnect ports when the network is running.")
	}

	//TODO
	return nil
}
func (net *network) Destroy(connectionID int64) error {
	net.mutex.Lock()
	defer net.mutex.Unlock()
	if net.hasStarted {
		return errors.New("Can disconnect ports when the network is running.")
	}

	//TODO
	return nil
}

func (net *network) Start() {
	net.mutex.Lock()
	net.hasStarted = true
	net.mutex.Unlock()
	for _, ns := range net.namespaces {
		for _, mux := range ns.muxInPortReceivers {
			mux.Listen()
		}
	}

	for _, component := range net.components {
		go component.Init()
	}

	//TODO: Block until end
}
